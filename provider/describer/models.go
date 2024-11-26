package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
)

func ListModel(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processModels(ctx, handler, openaiChan, &wg)
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

func GetModel(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	modelData, err := processModel(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	createdAt := unixToTimestamp(modelData.Created)
	value := models.Resource{
		ID:   modelData.Id,
		Name: modelData.Id,
		Description: JSONAllFieldsMarshaller{
			Value: model.ModelsDescription{
				ID:        modelData.Id,
				CreatedAt: createdAt,
				Object:    modelData.Object,
				OwnedBy:   modelData.OwnedBy,
			},
		},
	}
	return &value, nil
}

func processModels(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var modelsData *openai.ListModelsResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		modelsData, resp, e = handler.Client.ModelsAPI.ListModels(ctx).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, modelData := range modelsData.Data {
		wg.Add(1)
		go func(modelData openai.Model) {
			defer wg.Done()
			createdAt := unixToTimestamp(modelData.Created)
			value := models.Resource{
				ID:   modelData.Id,
				Name: modelData.Id,
				Description: JSONAllFieldsMarshaller{
					Value: model.ModelsDescription{
						ID:        modelData.Id,
						CreatedAt: createdAt,
						Object:    modelData.Object,
						OwnedBy:   modelData.OwnedBy,
					},
				},
			}
			openaiChan <- value
		}(modelData)
	}
}

func processModel(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.Model, error) {
	var modelData *openai.Model
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		modelData, resp, e = handler.Client.ModelsAPI.RetrieveModel(ctx, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return modelData, nil
}
