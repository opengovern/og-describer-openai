package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-openai/discovery/provider"
	"net/http"
)

func ListModels(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	resources, err := processModels(ctx, handler)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, value := range resources {
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

func GetModel(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	modelData, err := processModel(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          modelData.ID,
		Name:        modelData.ID,
		Description: modelData,
	}
	return &value, nil
}

func processModels(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var modelResponse model.ModelsResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/models"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(&modelResponse); e != nil {
			return nil, e
		}
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	var resources []models.Resource
	for _, modelData := range modelResponse.Data {
		value := models.Resource{
			ID:          modelData.ID,
			Name:        modelData.ID,
			Description: modelData,
		}
		resources = append(resources, value)
	}
	return resources, nil
}

func processModel(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*model.ModelsDescription, error) {
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
