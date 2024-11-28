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

func ListAssistants(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processAssistants(ctx, handler, openaiChan, &wg)
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

func processAssistants(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var assistants []model.AssistantDescription
	var assistantResponse *model.AssistantResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/assistants"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		var after string
		for {
			params := url.Values{}
			params.Set("limit", "100")
			params.Set("order", order)
			if after != "" {
				params.Set("after", after)
			}
			finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			req.Header.Set("OpenAI-Beta", "assistants=v2")
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&assistantResponse); e != nil {
				return nil, e
			}
			assistants = append(assistants, assistantResponse.Data...)
			if !assistantResponse.HasMore {
				break
			}
			after = assistantResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, assistant := range assistants {
		wg.Add(1)
		go func(assistant model.AssistantDescription) {
			defer wg.Done()
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
			openaiChan <- value
		}(assistant)
	}
}

func processAssistant(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*model.AssistantDescription, error) {
	var assistant *model.AssistantDescription
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/assistants/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(assistant); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return assistant, nil
}
