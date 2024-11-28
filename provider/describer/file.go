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

func processFiles(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var files []model.FileDescription
	var fileResponse *model.FileResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/files"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		var after string
		for {
			params := url.Values{}
			params.Set("limit", "10000")
			if after != "" {
				params.Set("after", after)
			}
			finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&fileResponse); e != nil {
				return nil, e
			}
			files = append(files, fileResponse.Data...)
			if !fileResponse.HasMore {
				break
			}
			after = fileResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, file := range files {
		wg.Add(1)
		go func(file model.FileDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   file.ID,
				Name: file.FileName,
				Description: JSONAllFieldsMarshaller{
					Value: file,
				},
			}
			openaiChan <- value
		}(file)
	}
}
