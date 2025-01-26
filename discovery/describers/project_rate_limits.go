package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-openai/discovery/provider"
	"net/http"
	"net/url"
)

func ListProjectRateLimits(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processProjectRateLimits(ctx, handler)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, value := range results {
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func processProjectRateLimits(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var projectRateLimits []model.ProjectRateLimit
	var projectRateLimitResponse model.ProjectRateLimitResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects/"
	for {
		var after string
		params := url.Values{}
		params.Set("limit", "100")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s%s/rate_limits?%s", baseURL, handler.ProjectID, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectRateLimitResponse); e != nil {
				return nil, e
			}
			projectRateLimits = append(projectRateLimits, projectRateLimitResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		if !projectRateLimitResponse.HasMore {
			break
		}
		after = projectRateLimitResponse.LastID
	}

	var results []models.Resource
	for _, projectRateLimit := range projectRateLimits {
		value := models.Resource{
			ID:   projectRateLimit.ID,
			Name: projectRateLimit.Model,
			Description: model.ProjectRateLimitDescription{
				ProjectRateLimit: projectRateLimit,
				ProjectID:        handler.ProjectID,
			},
		}
		results = append(results, value)
	}
	return results, nil
}
