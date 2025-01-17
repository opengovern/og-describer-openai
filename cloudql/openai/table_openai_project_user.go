package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiProjectUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_project_user",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectUser,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "user_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserID"),
				Description: "User ID"},
			{
				Name:      "project_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.ProjectID"),
			},
		}),
	}
}
