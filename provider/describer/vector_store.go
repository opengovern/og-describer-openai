package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-template/openai-go-client"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"net/http"
	"sync"
	"time"
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
	createdAt := unixToTimestamp(vectorStore.CreatedAt)
	var expiresAt time.Time
	if vectorStore.ExpiresAt.IsSet() {
		expires := *vectorStore.ExpiresAt.Get()
		expiresAt = unixToTimestamp(expires)
	}
	var lastActiveAt time.Time
	if vectorStore.LastActiveAt.IsSet() {
		lastActive := *vectorStore.LastActiveAt.Get()
		lastActiveAt = unixToTimestamp(lastActive)
	}
	value := models.Resource{
		ID:   vectorStore.Id,
		Name: vectorStore.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.VectorStoreDescription{
				Id:           vectorStore.Id,
				Object:       vectorStore.Object,
				CreatedAt:    createdAt,
				Name:         vectorStore.Name,
				UsageBytes:   vectorStore.UsageBytes,
				FileCounts:   vectorStore.FileCounts,
				Status:       vectorStore.Status,
				ExpiresAfter: vectorStore.ExpiresAfter,
				ExpiresAt:    expiresAt,
				LastActiveAt: lastActiveAt,
				Metadata:     vectorStore.Metadata,
			},
		},
	}
	return &value, nil
}

func processVectorStores(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var vectorStores *openai.ListVectorStoresResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		vectorStores, resp, e = handler.Client.VectorStoresAPI.ListVectorStores(ctx).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, vectorStore := range vectorStores.Data {
		wg.Add(1)
		go func(vectorStore openai.VectorStoreObject) {
			defer wg.Done()
			createdAt := unixToTimestamp(vectorStore.CreatedAt)
			var expiresAt time.Time
			if vectorStore.ExpiresAt.IsSet() {
				expires := *vectorStore.ExpiresAt.Get()
				expiresAt = unixToTimestamp(expires)
			}
			var lastActiveAt time.Time
			if vectorStore.LastActiveAt.IsSet() {
				lastActive := *vectorStore.LastActiveAt.Get()
				lastActiveAt = unixToTimestamp(lastActive)
			}
			value := models.Resource{
				ID:   vectorStore.Id,
				Name: vectorStore.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.VectorStoreDescription{
						Id:           vectorStore.Id,
						Object:       vectorStore.Object,
						CreatedAt:    createdAt,
						Name:         vectorStore.Name,
						UsageBytes:   vectorStore.UsageBytes,
						FileCounts:   vectorStore.FileCounts,
						Status:       vectorStore.Status,
						ExpiresAfter: vectorStore.ExpiresAfter,
						ExpiresAt:    expiresAt,
						LastActiveAt: lastActiveAt,
						Metadata:     vectorStore.Metadata,
					},
				},
			}
			openaiChan <- value
		}(vectorStore)
	}
}

func processVectorStore(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.VectorStoreObject, error) {
	var vectorStore *openai.VectorStoreObject
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		vectorStore, resp, e = handler.Client.VectorStoresAPI.GetVectorStore(ctx, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return vectorStore, nil
}
