package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiAssistant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_assistant",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAssistant,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "ID of the project."},
			{
				Name:      "name",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Name"),
			},
			{
				Name:      "description",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Description"),
			},
			{
				Name:      "model",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Model"),
			},
			{
				Name:      "instructions",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Instructions"),
			},
			{
				Name:      "tools",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Tools"),
			},
			{
				Name:      "tool_resources",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.ToolResources"),
			},
			{
				Name:      "metadata",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Metadata"),
			},
			{
				Name:      "temperature",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.Temperature"),
			},
			{
				Name:      "top_p",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.TopP"),
			},
			{
				Name:      "response_format",
				Type:      proto.ColumnType_UNKNOWN,
				Transform: transform.FromField("Description.ResponseFormat"),
			},
		}),
	}
}
