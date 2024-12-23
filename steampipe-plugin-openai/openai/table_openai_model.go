package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiModel(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_model",
		Description: "Models available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListModels,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "ID of the model, e.g. davinci."},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
				Description: "Timestamp of when the model was created."},
			{
				Name:        "owned_by",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OwnedBy"),
				Description: "Organization that owns the model, e.g. openai."},
			// Other columns
			{
				Name:        "object",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object"),
				Description: "Type of the object, e.g. model."},
			{
				Name:        "root",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Root"),
				Description: "Root of this model."},
			{
				Name:        "permission",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Permission"),
				Description: "Permissions for the model."},
			// Always null in testing? {Name: "parent", Type: proto.ColumnType_STRING, Description: ""},
		}),
	}
}
