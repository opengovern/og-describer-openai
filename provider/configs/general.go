package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "openai"                                    // example: aws, azure
	IntegrationName      = integration.Type("OPENAI")                  // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-openai" // example: github.com/opengovern/og-describer-aws
)
