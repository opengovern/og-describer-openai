package describer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opengovern/og-describer-openai/provider/model"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	order = "desc"
)

type OpenAIAPIHandler struct {
	Client         *http.Client
	APIKey         string
	OrganizationID string
	RateLimiter    *rate.Limiter
	Semaphore      chan struct{}
	MaxRetries     int
	RetryBackoff   time.Duration
}

func NewOpenAIAPIHandler(apiKey string, orgID string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *OpenAIAPIHandler {
	return &OpenAIAPIHandler{
		Client:         http.DefaultClient,
		APIKey:         apiKey,
		OrganizationID: orgID,
		RateLimiter:    rate.NewLimiter(rateLimit, burst),
		Semaphore:      make(chan struct{}, maxConcurrency),
		MaxRetries:     maxRetries,
		RetryBackoff:   retryBackoff,
	}
}

func getProjects(ctx context.Context, handler *OpenAIAPIHandler) ([]model.ProjectDescription, error) {
	var projects []model.ProjectDescription
	var projectResponse model.ProjectResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		var after string
		for {
			params := url.Values{}
			params.Set("limit", "100")
			if after != "" {
				params.Set("after", after)
			}
			finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectResponse); e != nil {
				return nil, e
			}
			projects = append(projects, projectResponse.Data...)
			if !projectResponse.HasMore {
				break
			}
			after = projectResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// DoRequest executes the openai API request with rate limiting, retries, and concurrency control.
func (h *OpenAIAPIHandler) DoRequest(ctx context.Context, req *http.Request, requestFunc func(req *http.Request) (*http.Response, error)) error {
	h.Semaphore <- struct{}{}
	defer func() { <-h.Semaphore }()
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Set request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.APIKey))
		if h.OrganizationID != "" {
			req.Header.Set("OpenAI-Organization", h.OrganizationID)
		}
		// Execute the request function
		resp, err = requestFunc(req)
		if err == nil {
			return nil
		}
		// Handle rate limit errors
		if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
			// Wait until reset time
			resetTime := resp.Header.Get("x-ratelimit-reset-requests")
			var resetDuration time.Duration
			if resetTime != "" {
				resetUnix, parseErr := strconv.ParseInt(resetTime, 10, 64)
				if parseErr == nil {
					resetDuration = time.Until(time.Unix(resetUnix, 0))
				}
			}
			// Set rate limiter value according to rate limit requests header
			var remainRequests int
			remainRequestsStr := resp.Header.Get("x-ratelimit-remaining-requests")
			if remainRequestsStr != "" {
				remainRequests, err = strconv.Atoi(remainRequestsStr)
				if err == nil {
					h.RateLimiter = rate.NewLimiter(rate.Every(resetDuration/time.Duration(remainRequests)), 1)
				}
			}
			if resetDuration > 0 {
				time.Sleep(resetDuration)
				continue
			}
			// Exponential backoff if headers are missing
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		// Handle temporary network errors
		if isTemporary(err) {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		break
	}
	return err
}

// isTemporary checks if an error is temporary.
func isTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}
