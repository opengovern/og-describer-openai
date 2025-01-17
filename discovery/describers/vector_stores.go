package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-openai/discovery/provider"
	"net/http"
	"net/url"
)

func ListVectorStores(ctx context.Context, handler *model.OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processVectorStores(ctx, handler)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, value := range results {
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

func GetVectorStore(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	vectorStore, err := processVectorStore(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          vectorStore.ID,
		Name:        vectorStore.Name,
		Description: vectorStore,
	}
	return &value, nil
}

func processVectorStores(ctx context.Context, handler *model.OpenAIAPIHandler) ([]models.Resource, error) {
	var vectorStores []model.VectorStoreDescription
	var vectorStoreResponse model.VectorStoreResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/vector_stores"

	var after string
	for {
		params := url.Values{}
		params.Set("limit", "100")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
		req, e := http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		req.Header.Set("OpenAI-Beta", "assistants=v2")
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&vectorStoreResponse); e != nil {
				return nil, e
			}
			vectorStores = append(vectorStores, vectorStoreResponse.Data...)

			return resp, e
		}
		err := handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		if !vectorStoreResponse.HasMore {
			break
		}
		after = vectorStoreResponse.LastID
	}

	var values []models.Resource
	for _, vectorStore := range vectorStores {
		value := models.Resource{
			ID:          vectorStore.ID,
			Name:        vectorStore.Name,
			Description: vectorStore,
		}
		values = append(values, value)
	}
	return values, nil
}

func processVectorStore(ctx context.Context, handler *model.OpenAIAPIHandler, resourceID string) (*model.VectorStoreDescription, error) {
	var vectorStore *model.VectorStoreDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/vector_stores/"
	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, e := http.NewRequest("GET", finalURL, nil)
	if e != nil {
		return nil, e
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(vectorStore); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	return vectorStore, nil
}
