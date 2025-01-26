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

func ListProjectUsers(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processProjectUsers(ctx, handler)
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

func processProjectUsers(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var projectUsers []model.ProjectUser
	var projectUserResponse model.ProjectUserResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects/"
	for {
		var after string
		params := url.Values{}
		params.Set("limit", "100")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s%s/users?%s", baseURL, handler.ProjectID, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectUserResponse); e != nil {
				return nil, e
			}
			projectUsers = append(projectUsers, projectUserResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		after = projectUserResponse.LastID
		if !projectUserResponse.HasMore {
			break
		}
	}
	var results []models.Resource
	for _, projectUser := range projectUsers {
		value := models.Resource{
			ID:   projectUser.ID,
			Name: projectUser.Name,
			Description: model.ProjectUserDescription{
				UserID:    projectUser.ID,
				ProjectID: handler.ProjectID,
				Object:    projectUser.Object,
				Name:      projectUser.Name,
				Email:     projectUser.Email,
				Role:      projectUser.Role,
			},
		}
		results = append(results, value)
	}
	return results, nil
}
