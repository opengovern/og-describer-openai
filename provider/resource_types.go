package provider

import (
	model "github.com/opengovern/og-describer-openai/pkg/sdk/models"
	"github.com/opengovern/og-describer-openai/provider/configs"
	"github.com/opengovern/og-describer-openai/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"OpenAI/Assistant": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Assistant",
		Tags: map[string][]string{
			"category": {"Assistant"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListAssistants),
		GetDescriber:  DescribeSingleByOpenAI(describer.GetAssistant),
	},

	"OpenAI/File": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/File",
		Tags: map[string][]string{
			"category": {"File"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListFiles),
		GetDescriber:  nil,
	},

	"OpenAI/Models": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Models",
		Tags: map[string][]string{
			"category": {"Model"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListModels),
		GetDescriber:  DescribeSingleByOpenAI(describer.GetModel),
	},

	"OpenAI/Project": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Project",
		Tags: map[string][]string{
			"category": {"Project"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListProjects),
		GetDescriber:  DescribeSingleByOpenAI(describer.GetProject),
	},

	"OpenAI/Project/APIKey": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Project/APIKey",
		Tags: map[string][]string{
			"category": {"Project"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListProjectAPIKeys),
		GetDescriber:  nil,
	},

	"OpenAI/Project/RateLimit": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Project/RateLimit",
		Tags: map[string][]string{
			"category": {"Project"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListProjectRateLimits),
		GetDescriber:  nil,
	},

	"OpenAI/Project/ServiceAccount": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Project/ServiceAccount",
		Tags: map[string][]string{
			"category": {"Project"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListProjectServiceAccounts),
		GetDescriber:  nil,
	},

	"OpenAI/Project/User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/Project/User",
		Tags: map[string][]string{
			"category": {"Project"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListProjectUsers),
		GetDescriber:  nil,
	},

	"OpenAI/VectorStore": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "OpenAI/VectorStore",
		Tags: map[string][]string{
			"category": {"VectorStore"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByOpenAI(describer.ListVectorStores),
		GetDescriber:  nil,
	},
}
