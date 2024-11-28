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

func ListVectorStores(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processVectorStores(ctx, handler, openaiChan, &wg)
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

func GetVectorStore(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	vectorStore, err := processVectorStore(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   vectorStore.ID,
		Name: vectorStore.Name,
		Description: JSONAllFieldsMarshaller{
			Value: vectorStore,
		},
	}
	return &value, nil
}

func processVectorStores(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var vectorStores []model.VectorStoreDescription
	var vectorStoreResponse model.VectorStoreResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/vector_stores"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		var after string
		for {
			params := url.Values{}
			params.Set("limit", "100")
			if after != "" {
				params.Set("after", after)
			}
			finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&vectorStoreResponse); e != nil {
				return nil, e
			}
			vectorStores = append(vectorStores, vectorStoreResponse.Data...)
			if !vectorStoreResponse.HasMore {
				break
			}
			after = vectorStoreResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, vectorStore := range vectorStores {
		wg.Add(1)
		go func(vectorStore model.VectorStoreDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   vectorStore.ID,
				Name: vectorStore.Name,
				Description: JSONAllFieldsMarshaller{
					Value: vectorStore,
				},
			}
			openaiChan <- value
		}(vectorStore)
	}
}

func processVectorStore(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*model.VectorStoreDescription, error) {
	var vectorStore *model.VectorStoreDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/vector_stores/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(vectorStore); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return vectorStore, nil
}
