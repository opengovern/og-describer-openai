package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"sync"
)

func ListFiles(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processFiles(ctx, handler, openaiChan, &wg)
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

func GetFile(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	file, err := processFile(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	createdAt := unixToTimestamp(file.CreatedAt)
	value := models.Resource{
		ID:   file.Id,
		Name: file.Filename,
		Description: JSONAllFieldsMarshaller{
			Value: model.FileDescription{
				ID:        file.Id,
				FileName:  file.Filename,
				CreatedAt: createdAt,
				Bytes:     file.Bytes,
				Object:    file.Object,
				Purpose:   file.Purpose,
			},
		},
	}
	return &value, nil
}

func processFiles(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var files *openai.ListFilesResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		files, resp, e = handler.Client.FilesAPI.ListFiles(ctx).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, file := range files.Data {
		wg.Add(1)
		go func(file openai.OpenAIFile) {
			defer wg.Done()
			createdAt := unixToTimestamp(file.CreatedAt)
			value := models.Resource{
				ID:   file.Id,
				Name: file.Filename,
				Description: JSONAllFieldsMarshaller{
					Value: model.FileDescription{
						ID:        file.Id,
						FileName:  file.Filename,
						CreatedAt: createdAt,
						Bytes:     file.Bytes,
						Object:    file.Object,
						Purpose:   file.Purpose,
					},
				},
			}
			openaiChan <- value
		}(file)
	}
}

func processFile(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.OpenAIFile, error) {
	var file *openai.OpenAIFile
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		file, resp, e = handler.Client.FilesAPI.RetrieveFile(ctx, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return file, nil
}
