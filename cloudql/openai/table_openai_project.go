package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_project",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProject,
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
				Name:      "created_at",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "archived_at",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.ArchivedAt").NullIfZero().Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "status",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Status"),
			},
		}),
	}
}
