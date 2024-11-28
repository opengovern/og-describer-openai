package provider

import (
	"errors"
	model "github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/configs"
	"github.com/opengovern/og-describer-openai/provider/describer"
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
		// Get organization ID
		organizationID := cfg.OrganizationID

		// Create openai handler
		openAIAPIHandler := describer.NewOpenAIAPIHandler(apiKey, organizationID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

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
		// Get organization ID
		organizationID := cfg.OrganizationID

		// Create openai handler
		openAIAPIHandler := describer.NewOpenAIAPIHandler(apiKey, organizationID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, openAIAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
