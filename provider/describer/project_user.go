package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
)

func ListProjectUsers(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	projects, err := getProjects(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, project := range projects {
			processProjectUsers(ctx, handler, project.Id, openaiChan, &wg)
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

//func GetProjectUser(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
//	projectUser, err := processProjectUser(ctx, handler, resourceID)
//	if err != nil {
//		return nil, err
//	}
//	addedAt := unixToTimestamp(projectUser.AddedAt)
//	value := models.Resource{
//		ID:   projectUser.Id,
//		Name: projectUser.Name,
//		Description: JSONAllFieldsMarshaller{
//			Value: model.ProjectUserDescription{
//				Object:  projectUser.Object,
//				ID:      projectUser.Id,
//				Name:    projectUser.Name,
//				Email:   projectUser.Email,
//				Role:    projectUser.Role,
//				AddedAt: addedAt,
//			},
//		},
//	}
//	return &value, nil
//}

func processProjectUsers(ctx context.Context, handler *OpenAIAPIHandler, projectID string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projectUsers *openai.ProjectUserListResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projectUsers, resp, e = handler.Client.ProjectsAPI.ListProjectUsers(ctx, projectID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, projectUser := range projectUsers.Data {
		wg.Add(1)
		go func(projectUser openai.ProjectUser) {
			defer wg.Done()
			//addedAt := unixToTimestamp(projectUser.AddedAt)
			value := models.Resource{
				ID:   projectUser.Id,
				Name: projectUser.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.ProjectUserDescription{
						UserID:    projectUser.Id,
						ProjectID: projectID,
						//Object:  projectUser.Object,
						//ID:      projectUser.Id,
						//Name:    projectUser.Name,
						//Email:   projectUser.Email,
						//Role:    projectUser.Role,
						//AddedAt: addedAt,
					},
				},
			}
			openaiChan <- value
		}(projectUser)
	}
}

//func processProjectUser(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.ProjectUser, error) {
//	var projectUser *openai.ProjectUser
//	var resp *http.Response
//	requestFunc := func() (*http.Response, error) {
//		var e error
//		projectUser, resp, e = handler.Client.ProjectsAPI.RetrieveProjectUser(ctx, handler.ProjectID, resourceID).Execute()
//		return resp, e
//	}
//	err := handler.DoRequest(ctx, requestFunc)
//	if err != nil {
//		return nil, err
//	}
//	return projectUser, nil
//}
