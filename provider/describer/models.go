package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
)

func ListModels(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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
	value := models.Resource{
		ID:   modelData.ID,
		Name: modelData.ID,
		Description: JSONAllFieldsMarshaller{
			Value: modelData,
		},
	}
	return &value, nil
}

func processModels(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var modelResponse model.ModelsResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/models"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		req, e = http.NewRequest("GET", baseURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(&modelResponse); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, modelData := range modelResponse.Data {
		wg.Add(1)
		go func(modelData model.ModelsDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   modelData.ID,
				Name: modelData.ID,
				Description: JSONAllFieldsMarshaller{
					Value: modelData,
				},
			}
			openaiChan <- value
		}(modelData)
	}
}

func processModel(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*model.ModelsDescription, error) {
	var modelData *model.ModelsDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/models/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(modelData); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return modelData, nil
}
