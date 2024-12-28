package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
	"net/url"
)

func ListFiles(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {

	results, err := processFiles(ctx, handler)
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

func processFiles(ctx context.Context, handler *OpenAIAPIHandler) ([]models.Resource, error) {
	var files []model.FileDescription
	var fileResponse *model.FileResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/files"
	for {
		var after string

		params := url.Values{}
		params.Set("limit", "10000")
		if after != "" {
			params.Set("after", after)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error

			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&fileResponse); e != nil {
				return nil, e
			}
			files = append(files, fileResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		if !fileResponse.HasMore {
			break
		}
		after = fileResponse.LastID
	}

	var values []models.Resource
	for _, file := range files {
		value := models.Resource{
			ID:          file.ID,
			Name:        file.FileName,
			Description: file,
		}
		values = append(values, value)
	}
	return values, nil
}
