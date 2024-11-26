package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
)

func ListProjectServiceAccounts(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processProjectServiceAccounts(ctx, handler, openaiChan, &wg)
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

func GetProjectServiceAccount(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	projectServiceAccount, err := processProjectServiceAccount(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	createdAt := unixToTimestamp(projectServiceAccount.CreatedAt)
	value := models.Resource{
		ID:   projectServiceAccount.Id,
		Name: projectServiceAccount.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.ProjectServiceAccountDescription{
				Object:    projectServiceAccount.Object,
				ID:        projectServiceAccount.Id,
				Name:      projectServiceAccount.Name,
				Role:      projectServiceAccount.Role,
				CreatedAt: createdAt,
			},
		},
	}
	return &value, nil
}

func processProjectServiceAccounts(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projectServiceAccounts *openai.ProjectServiceAccountListResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projectServiceAccounts, resp, e = handler.Client.ProjectsAPI.ListProjectServiceAccounts(ctx, handler.ProjectID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, projectServiceAccount := range projectServiceAccounts.Data {
		wg.Add(1)
		go func(projectServiceAccount openai.ProjectServiceAccount) {
			defer wg.Done()
			createdAt := unixToTimestamp(projectServiceAccount.CreatedAt)
			value := models.Resource{
				ID:   projectServiceAccount.Id,
				Name: projectServiceAccount.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.ProjectServiceAccountDescription{
						Object:    projectServiceAccount.Object,
						ID:        projectServiceAccount.Id,
						Name:      projectServiceAccount.Name,
						Role:      projectServiceAccount.Role,
						CreatedAt: createdAt,
					},
				},
			}
			openaiChan <- value
		}(projectServiceAccount)
	}
}

func processProjectServiceAccount(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.ProjectServiceAccount, error) {
	var projectServiceAccount *openai.ProjectServiceAccount
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projectServiceAccount, resp, e = handler.Client.ProjectsAPI.RetrieveProjectServiceAccount(ctx, handler.ProjectID, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return projectServiceAccount, nil
}
