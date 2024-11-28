package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"net/url"
	"sync"
)

func ListProjectAPIKeys(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	projects, err := getProjects(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, project := range projects {
			processProjectAPIKeys(ctx, handler, project.ID, openaiChan, &wg)
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

func processProjectAPIKeys(ctx context.Context, handler *OpenAIAPIHandler, projectID string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projectAPIKeys []model.ProjectApiKey
	var projectAPIKeyResponse model.ProjectApiKeyResponse
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
			finalURL := fmt.Sprintf("%s%s/api_keys?%s", baseURL, projectID, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectAPIKeyResponse); e != nil {
				return nil, e
			}
			projectAPIKeys = append(projectAPIKeys, projectAPIKeyResponse.Data...)
			if !projectAPIKeyResponse.HasMore {
				break
			}
			after = projectAPIKeyResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, projectAPIKey := range projectAPIKeys {
		wg.Add(1)
		go func(projectAPIKey model.ProjectApiKey) {
			defer wg.Done()
			value := models.Resource{
				ID:   projectAPIKey.ID,
				Name: projectAPIKey.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.ProjectApiKeyDescription{
						ProjectApiKey: projectAPIKey,
						ProjectID:     projectID,
					},
				},
			}
			openaiChan <- value
		}(projectAPIKey)
	}
}
