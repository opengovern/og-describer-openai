package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-openai/discovery/provider"
	"net/http"
	"net/url"
	"sync"
)

func ListProjects(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processProjects(ctx, handler, openaiChan, &wg)
		wg.Wait()
		close(openaiChan)
	}()
	var values []models.Resource
	for value := range openaiChan {
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

func processProjects(ctx context.Context, handler *model.OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
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
		return
	}
	for _, project := range projects {
		wg.Add(1)
		go func(project model.ProjectDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:          project.ID,
				Name:        project.Name,
				Description: project,
			}
			openaiChan <- value
		}(project)
	}
}

func processProject(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*model.ProjectDescription, error) {
	var project *model.ProjectDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(project); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return project, nil
}
