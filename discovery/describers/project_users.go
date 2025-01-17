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

func ListProjectUsers(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	projects, err := getProjects(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, project := range projects {
			processProjectUsers(ctx, handler, project.ID, openaiChan, &wg)
		}
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

func processProjectUsers(ctx context.Context, handler *model.OpenAIAPIHandler, projectID string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projectUsers []model.ProjectUser
	var projectUserResponse model.ProjectUserResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		var after string
		for {
			params := url.Values{}
			params.Set("limit", "100")
			if after != "" {
				params.Set("after", after)
			}
			finalURL := fmt.Sprintf("%s%s/users?%s", baseURL, projectID, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectUserResponse); e != nil {
				return nil, e
			}
			projectUsers = append(projectUsers, projectUserResponse.Data...)
			if !projectUserResponse.HasMore {
				break
			}
			after = projectUserResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, projectUser := range projectUsers {
		wg.Add(1)
		go func(projectUser model.ProjectUser) {
			defer wg.Done()
			value := models.Resource{
				ID:   projectUser.ID,
				Name: projectUser.Name,
				Description: model.ProjectUserDescription{
					UserID:    projectUser.ID,
					ProjectID: projectID,
					Object:    projectUser.Object,
					Name:      projectUser.Name,
					Email:     projectUser.Email,
					Role:      projectUser.Role,
				},
			}
			openaiChan <- value
		}(projectUser)
	}
}
