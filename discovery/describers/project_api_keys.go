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

func ListProjectAPIKeys(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processProjectAPIKeys(ctx, handler)
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

func processProjectAPIKeys(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var projectAPIKeys []model.ProjectApiKey
	var projectAPIKeyResponse model.ProjectApiKeyResponse
	var resp *http.Response
	var after string
	baseURL := "https://api.openai.com/v1/organization/projects/"
	for {
		params := url.Values{}
		params.Set("limit", "100")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s%s/api_keys?%s", baseURL, handler.ProjectID, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectAPIKeyResponse); e != nil {
				return nil, e
			}
			projectAPIKeys = append(projectAPIKeys, projectAPIKeyResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		if !projectAPIKeyResponse.HasMore {
			break
		}
		after = projectAPIKeyResponse.LastID
	}
	var results []models.Resource
	for _, projectAPIKey := range projectAPIKeys {
		value := models.Resource{
			ID:   projectAPIKey.ID,
			Name: projectAPIKey.Name,
			Description: model.ProjectApiKeyDescription{
				ProjectApiKey: projectAPIKey,
				ProjectID:     handler.ProjectID,
			},
		}
		results = append(results, value)
	}
	return results, nil
}
