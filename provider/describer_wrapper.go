package provider

import (
	"errors"
	"fmt"
	openai "github.com/opengovern/og-describer-template/openai-go-client"
	model "github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/configs"
	"github.com/opengovern/og-describer-template/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

// DescribeListByOpenAI A wrapper to pass openai authorization to describer list functions
func DescribeListByOpenAI(describe func(context.Context, *describer.OpenAIAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Get apikey and check it has value
		apiKey := cfg.APIKey
		if apiKey == "" {
			return nil, errors.New("api_key must be configured")
		}
		organizationID := cfg.OrganizationID
		if organizationID == "" {
			return nil, errors.New("organization ID must be configured")
		}
		projectID := cfg.ProjectID
		// Create openai configs
		config := openai.NewConfiguration()
		config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		config.AddDefaultHeader("OpenAI-Organization", organizationID)
		config.AddDefaultHeader("OpenAI-Project", projectID)

		// Create openai client
		client := openai.NewAPIClient(config)

		// Create openai handler
		openAIAPIHandler := describer.NewOpenAIAPIHandler(client, projectID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, openAIAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByOpenAI A wrapper to pass openai authorization to describer get functions
func DescribeSingleByOpenAI(describe func(context.Context, *describer.OpenAIAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string) (*model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Get apikey and check it has value
		apiKey := cfg.APIKey
		if apiKey == "" {
			return nil, errors.New("api_key must be configured")
		}
		organizationID := cfg.OrganizationID
		if organizationID == "" {
			return nil, errors.New("organization ID must be configured")
		}
		projectID := cfg.ProjectID
		// Create openai configs
		config := openai.NewConfiguration()
		config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		config.AddDefaultHeader("OpenAI-Organization", organizationID)
		config.AddDefaultHeader("OpenAI-Project", projectID)

		// Create openai client
		client := openai.NewAPIClient(config)

		// Create openai handler
		openAIAPIHandler := describer.NewOpenAIAPIHandler(client, projectID, rate.Every(time.Second), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, openAIAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
