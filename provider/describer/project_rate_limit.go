package describer

import (
	"context"
	openai "github.com/opengovern/og-describer-template/openai-go-client"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"net/http"
	"sync"
)

func ListProjectRateLimits(ctx context.Context, handler *OpenAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	openaiChan := make(chan models.Resource)
	go func() {
		processProjectRateLimits(ctx, handler, openaiChan, &wg)
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

func processProjectRateLimits(ctx context.Context, handler *OpenAIAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var projectRateLimits *openai.ProjectRateLimitListResponse
	var resp *http.Response
	requestFunc := func() (*http.Response, error) {
		var e error
		projectRateLimits, resp, e = handler.Client.ProjectsAPI.ListProjectRateLimits(ctx, handler.ProjectID).Execute()
		return resp, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, projectRateLimit := range projectRateLimits.Data {
		wg.Add(1)
		go func(projectRateLimit openai.ProjectRateLimit) {
			defer wg.Done()
			value := models.Resource{
				ID:   projectRateLimit.Id,
				Name: projectRateLimit.Model,
				Description: JSONAllFieldsMarshaller{
					Value: model.ProjectRateLimitDescription{
						Object:                      projectRateLimit.Object,
						ID:                          projectRateLimit.Id,
						Model:                       projectRateLimit.Model,
						MaxRequestsPer1Minute:       projectRateLimit.MaxRequestsPer1Minute,
						MaxTokensPer1Minute:         projectRateLimit.MaxTokensPer1Minute,
						MaxImagesPer1Minute:         projectRateLimit.MaxImagesPer1Minute,
						MaxAudioMegabytesPer1Minute: projectRateLimit.MaxAudioMegabytesPer1Minute,
						MaxRequestsPer1Day:          projectRateLimit.MaxRequestsPer1Day,
						Batch1DayMaxInputTokens:     projectRateLimit.Batch1DayMaxInputTokens,
					},
				},
			}
			openaiChan <- value
		}(projectRateLimit)
	}
}
