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

func ListProjectServiceAccounts(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processProjectServiceAccounts(ctx, handler)
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

func processProjectServiceAccounts(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var projectServiceAccounts []model.ProjectServiceAccount
	var projectServiceAccountResponse model.ProjectServiceAccountResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects/"
	for {
		var after string
		params := url.Values{}
		params.Set("limit", "100")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s%s/service_accounts?%s", baseURL, handler.ProjectID, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectServiceAccountResponse); e != nil {
				return nil, e
			}
			projectServiceAccounts = append(projectServiceAccounts, projectServiceAccountResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		if !projectServiceAccountResponse.HasMore {
			break
		}
		after = projectServiceAccountResponse.LastID
	}

	var results []models.Resource
	for _, projectServiceAccount := range projectServiceAccounts {
		value := models.Resource{
			ID:   projectServiceAccount.ID,
			Name: projectServiceAccount.Name,
			Description: model.ProjectServiceAccountDescription{
				ProjectServiceAccount: projectServiceAccount,
				ProjectID:             handler.ProjectID,
			},
		}
		results = append(results, value)
	}
	return results, nil
}
