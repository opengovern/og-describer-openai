package describer

import (
	"context"
	"errors"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

type OpenAIAPIHandler struct {
	Client *openai.APIClient
	//ProjectID    string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewOpenAIAPIHandler(client *openai.APIClient, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *OpenAIAPIHandler {
	return &OpenAIAPIHandler{
		Client:       client,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}

func getProjects(ctx context.Context, handler *OpenAIAPIHandler) ([]openai.Project, error) {
	var projects *openai.ProjectListResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projects, resp, e = handler.Client.ProjectsAPI.ListProjects(ctx).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return projects.Data, nil
}

// DoRequest executes the openai API request with rate limiting, retries, and concurrency control.
func (h *OpenAIAPIHandler) DoRequest(ctx context.Context, requestFunc func() (*http.Response, error)) error {
	h.Semaphore <- struct{}{}
	defer func() { <-h.Semaphore }()
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Execute the request function
		resp, err = requestFunc()
		if err == nil {
			return nil
		}
		// Handle rate limit errors
		if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
			// Set rate limiter value according to rate limit requests header
			var limitRequests int
			limitRequestsStr := resp.Header.Get("x-ratelimit-limit-requests")
			if limitRequestsStr != "" {
				limitRequests, err = strconv.Atoi(limitRequestsStr)
				if err == nil {
					h.RateLimiter = rate.NewLimiter(rate.Every(time.Minute/time.Duration(limitRequests)), 1)
				}
			}
			// Wait until reset time
			resetTime := resp.Header.Get("x-ratelimit-reset-requests")
			if resetTime != "" {
				resetUnix, parseErr := strconv.ParseInt(resetTime, 10, 64)
				if parseErr == nil {
					sleepDuration := time.Until(time.Unix(resetUnix, 0))
					if sleepDuration > 0 {
						time.Sleep(sleepDuration)
						continue
					}
				}
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

func unixToTimestamp(timeInt int32) time.Time {
	if timeInt == 0 {
		return time.Time{}
	}
	t := time.Unix(int64(timeInt), 0)
	return t
}
