package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiFile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_file",
		Description: "Files uploaded to OpenAI for fine-tuning.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFile,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetFile,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "ID of the file, e.g. davinci."},
			{
				Name:        "file_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileName"),
				Description: "Name of the file."},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
				Description: "Timestamp of when the file was created."},
			{
				Name:        "bytes",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Bytes"),
				Description: "Size of the file in bytes"},
			// Other columns
			{
				Name:        "object",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object"),
				Description: "The type of the uploaded document, e.g. file."},
			//{
			//	Name: "owner",
			//	Type: proto.ColumnType_STRING,
			//	Description: "Organization that owns the file, e.g. openai."},
			{
				Name:        "purpose",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Purpose"),
				Description: "The intended purpose of the uploaded documents."},
		}),
	}
}
