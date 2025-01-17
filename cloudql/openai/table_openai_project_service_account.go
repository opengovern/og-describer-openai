package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiProjectServiceAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_project_service_account",
		Description: "Service accounts in a Project.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectServiceAccount,
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
				Name:      "role",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Role"),
			},
			{
				Name:      "created_at",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "project_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.ProjectID"),
			},
		}),
	}
}
