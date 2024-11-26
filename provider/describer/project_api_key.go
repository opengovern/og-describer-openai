package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
)

func ListProjectAPIKeys(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processProjectAPIKeys(ctx, handler, openaiChan, &wg)
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

func GetProjectAPIKey(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	projectAPIKey, err := processProjectAPIKey(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	createdAt := unixToTimestamp(projectAPIKey.CreatedAt)
	value := models.Resource{
		ID:   projectAPIKey.Id,
		Name: projectAPIKey.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.ProjectApiKeyDescription{
				Object:        projectAPIKey.Object,
				ID:            projectAPIKey.Id,
				Name:          projectAPIKey.Name,
				RedactedValue: projectAPIKey.RedactedValue,
				CreatedAt:     createdAt,
				Owner:         projectAPIKey.Owner,
			},
		},
	}
	return &value, nil
}

func processProjectAPIKeys(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projectAPIKeys *openai.ProjectApiKeyListResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projectAPIKeys, resp, e = handler.Client.ProjectsAPI.ListProjectApiKeys(ctx, handler.ProjectID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, projectAPIKey := range projectAPIKeys.Data {
		wg.Add(1)
		go func(projectAPIKey openai.ProjectApiKey) {
			defer wg.Done()
			createdAt := unixToTimestamp(projectAPIKey.CreatedAt)
			value := models.Resource{
				ID:   projectAPIKey.Id,
				Name: projectAPIKey.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.ProjectApiKeyDescription{
						Object:        projectAPIKey.Object,
						ID:            projectAPIKey.Id,
						Name:          projectAPIKey.Name,
						RedactedValue: projectAPIKey.RedactedValue,
						CreatedAt:     createdAt,
						Owner:         projectAPIKey.Owner,
					},
				},
			}
			openaiChan <- value
		}(projectAPIKey)
	}
}

func processProjectAPIKey(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.ProjectApiKey, error) {
	var projectAPIKey *openai.ProjectApiKey
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projectAPIKey, resp, e = handler.Client.ProjectsAPI.RetrieveProjectApiKey(ctx, handler.ProjectID, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return projectAPIKey, nil
}
