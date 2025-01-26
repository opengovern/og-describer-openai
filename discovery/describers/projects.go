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

func ListProjects(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processProjects(ctx, handler)
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

func GetProject(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	project, err := processProject(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          project.ID,
		Name:        project.Name,
		Description: project,
	}
	return &value, nil
}

func processProjects(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var projects []model.ProjectDescription
	var projectResponse model.ProjectResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects"
	for {
		var after string
		params := url.Values{}
		params.Set("limit", "100")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectResponse); e != nil {
				return nil, e
			}
			projects = append(projects, projectResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		after = projectResponse.LastID
		if !projectResponse.HasMore {
			break
		}
	}

	var results []models.Resource
	for _, project := range projects {
		value := models.Resource{
			ID:          project.ID,
			Name:        project.Name,
			Description: project,
		}
		results = append(results, value)
	}
	return results, nil
}

func processProject(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*model.ProjectDescription, error) {
	var project *model.ProjectDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects/"
	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(project); e != nil {
			return nil, e
		}
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	return project, nil
}
