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

func ListAssistants(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	results, err := processAssistants(ctx, handler)
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

func GetAssistant(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*models.Resource, error) {
	assistant, err := processAssistant(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var name string
	if assistant.Name != nil {
		name = *assistant.Name
	}
	value := models.Resource{
		ID:   assistant.ID,
		Name: name,
		Description: JSONAllFieldsMarshaller{
			Value: assistant,
		},
	}
	return &value, nil
}

func processAssistants(ctx context.Context, handler *OpenAIAPIHandler) ([]models.Resource, error) {
	var assistants []model.AssistantDescription
	var assistantResponse *model.AssistantResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/assistants"
	for {
		var after string
		params := url.Values{}
		params.Set("limit", "100")
		params.Set("order", order)
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

			req.Header.Set("OpenAI-Beta", "assistants=v2")
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&assistantResponse); e != nil {
				return nil, e
			}
			assistants = append(assistants, assistantResponse.Data...)
			return resp, e
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, err
		}
		after = assistantResponse.LastID
		if !assistantResponse.HasMore {
			break
		}
	}

	var results []models.Resource
	for _, assistant := range assistants {

		var name string
		if assistant.Name != nil {
			name = *assistant.Name
		}
		value := models.Resource{
			ID:   assistant.ID,
			Name: name,
			Description: JSONAllFieldsMarshaller{
				Value: assistant,
			},
		}
		results = append(results, value)
	}
	return results, nil
}

func processAssistant(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*model.AssistantDescription, error) {
	var assistant *model.AssistantDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/assistants/"
	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error

		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(assistant); e != nil {
			return nil, e
		}
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	return assistant, nil
}
