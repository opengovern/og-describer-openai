package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
	"time"
)

func ListProjects(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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

func GetProject(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	project, err := processProject(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	createdAt := unixToTimestamp(project.CreatedAt)
	var archivedAt time.Time
	if project.ArchivedAt.IsSet() {
		archived := *project.ArchivedAt.Get()
		archivedAt = unixToTimestamp(archived)
	}
	value := models.Resource{
		ID:   project.Id,
		Name: project.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.ProjectDescription{
				Object:     project.Object,
				ID:         project.Id,
				Name:       project.Name,
				CreatedAt:  createdAt,
				ArchivedAt: archivedAt,
				Status:     project.Status,
			},
		},
	}
	return &value, nil
}

func processProjects(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projects *openai.ProjectListResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projects, resp, e = handler.Client.ProjectsAPI.ListProjects(ctx).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, project := range projects.Data {
		wg.Add(1)
		go func(project openai.Project) {
			defer wg.Done()
			createdAt := unixToTimestamp(project.CreatedAt)
			var archivedAt time.Time
			if project.ArchivedAt.IsSet() {
				archived := *project.ArchivedAt.Get()
				archivedAt = unixToTimestamp(archived)
			}
			value := models.Resource{
				ID:   project.Id,
				Name: project.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.ProjectDescription{
						Object:     project.Object,
						ID:         project.Id,
						Name:       project.Name,
						CreatedAt:  createdAt,
						ArchivedAt: archivedAt,
						Status:     project.Status,
					},
				},
			}
			openaiChan <- value
		}(project)
	}
}

func processProject(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.Project, error) {
	var project *openai.Project
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		project, resp, e = handler.Client.ProjectsAPI.RetrieveProject(ctx, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return project, nil
}
