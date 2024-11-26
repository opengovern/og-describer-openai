package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/model"
	"net/http"
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
	//createdAt := unixToTimestamp(assistant.CreatedAt)
	var name string
	if assistant.Name.IsSet() {
		name = *assistant.Name.Get()
	}
	var description string
	if assistant.Description.IsSet() {
		description = *assistant.Description.Get()
	}
	var instructions string
	if assistant.Instructions.IsSet() {
		instructions = *assistant.Instructions.Get()
	}
	var toolResources openai.AssistantObjectToolResources
	if assistant.ToolResources.IsSet() {
		toolResources = *assistant.ToolResources.Get()
	}
	var temperature float32
	if assistant.Temperature.IsSet() {
		temperature = *assistant.Temperature.Get()
	}
	var topP float32
	if assistant.TopP.IsSet() {
		topP = *assistant.TopP.Get()
	}
	value := models.Resource{
		ID:   assistant.Id,
		Name: name,
		Description: JSONAllFieldsMarshaller{
			Value: model.AssistantDescription{
				ID: assistant.Id,
				//Object:         assistant.Object,
				//CreatedAt:      createdAt,
				Name:           name,
				Description:    description,
				Model:          assistant.Model,
				Instructions:   instructions,
				Tools:          assistant.Tools,
				ToolResources:  toolResources,
				Metadata:       assistant.Metadata,
				Temperature:    temperature,
				TopP:           topP,
				ResponseFormat: assistant.ResponseFormat,
			},
		},
	}
	return &value, nil
}

func processAssistants(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var assistants *openai.ListAssistantsResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		assistants, resp, e = handler.Client.AssistantsAPI.ListAssistants(ctx).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, assistant := range assistants.Data {
		wg.Add(1)
		go func(assistant openai.AssistantObject) {
			defer wg.Done()
			//createdAt := unixToTimestamp(assistant.CreatedAt)
			var name string
			if assistant.Name.IsSet() {
				name = *assistant.Name.Get()
			}
			var description string
			if assistant.Description.IsSet() {
				description = *assistant.Description.Get()
			}
			var instructions string
			if assistant.Instructions.IsSet() {
				instructions = *assistant.Instructions.Get()
			}
			var toolResources openai.AssistantObjectToolResources
			if assistant.ToolResources.IsSet() {
				toolResources = *assistant.ToolResources.Get()
			}
			var temperature float32
			if assistant.Temperature.IsSet() {
				temperature = *assistant.Temperature.Get()
			}
			var topP float32
			if assistant.TopP.IsSet() {
				topP = *assistant.TopP.Get()
			}
			value := models.Resource{
				ID:   assistant.Id,
				Name: name,
				Description: JSONAllFieldsMarshaller{
					Value: model.AssistantDescription{
						ID: assistant.Id,
						//Object:         assistant.Object,
						//CreatedAt:      createdAt,
						Name:           name,
						Description:    description,
						Model:          assistant.Model,
						Instructions:   instructions,
						Tools:          assistant.Tools,
						ToolResources:  toolResources,
						Metadata:       assistant.Metadata,
						Temperature:    temperature,
						TopP:           topP,
						ResponseFormat: assistant.ResponseFormat,
					},
				},
			}
			openaiChan <- value
		}(assistant)
	}
}

func processAssistant(ctx context.Context, handler *OpenAIAPIHandler, resourceID string) (*openai.AssistantObject, error) {
	var assistant *openai.AssistantObject
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		assistant, resp, e = handler.Client.AssistantsAPI.GetAssistant(ctx, resourceID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return assistant, nil
}
