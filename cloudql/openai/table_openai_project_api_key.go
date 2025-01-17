package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiProjectApiKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_project_api_key",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectApiKey,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "ID of the project."},
			{
				Name:      "object",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Object"),
			},
			{
				Name:      "name",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Name"),
			},
			{
				Name:      "redacted_value",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.RedactedValue"),
			},
			{
				Name:      "created_at",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "owner",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Owner"),
			},
			{
				Name:      "project_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.ProjectID"),
			},
		}),
	}
}
