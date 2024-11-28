package provider
import (
	"github.com/opengovern/og-describer-openai/provider/describer"
	"github.com/opengovern/og-describer-openai/provider/configs"
	model "github.com/opengovern/og-describer-openai/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"OpenAI/Project": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListProjects),
		GetDescriber:         DescribeSingleByOpenAI(describer.GetProject),
	},

	"OpenAI/Project/ApiKey": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Project/ApiKey",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListProjectAPIKeys),
		GetDescriber:         nil,
	},

	"OpenAI/Project/RateLimit": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Project/RateLimit",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListProjectRateLimits),
		GetDescriber:         nil,
	},

	"OpenAI/Project/ServiceAccount": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Project/ServiceAccount",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListProjectServiceAccounts),
		GetDescriber:         nil,
	},

	"OpenAI/Project/User": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Project/User",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListProjectUsers),
		GetDescriber:         nil,
	},

	"OpenAI/Model": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Model",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListModels),
		GetDescriber:         DescribeSingleByOpenAI(describer.GetModel),
	},

	"OpenAI/File": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/File",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListFiles),
		GetDescriber:         nil,
	},

	"OpenAI/VectorStore": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/VectorStore",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListVectorStores),
		GetDescriber:         DescribeSingleByOpenAI(describer.GetVectorStore),
	},

	"OpenAI/Assistant": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "OpenAI/Assistant",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByOpenAI(describer.ListAssistants),
		GetDescriber:         DescribeSingleByOpenAI(describer.GetAssistant),
	},
}
