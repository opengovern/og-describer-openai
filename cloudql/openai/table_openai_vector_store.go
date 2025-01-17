package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiVectorStore(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_vector_store",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVectorStore,
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
				Name:      "usage_bytes",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.UsageBytes"),
			},
			{
				Name:      "file_counts",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.FileCounts"),
			},
			{
				Name:      "status",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Status"),
			},
			{
				Name:      "expires_after",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.ExpiresAfter"),
			},
			{
				Name:      "expires_at",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.ExpiresAt").NullIfZero().Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "last_active_at",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.LastActiveAt").NullIfZero().Transform(transform.UnixToTimestamp),
			},
			{
				Name:      "metadata",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Metadata"),
			},
		}),
	}
}
