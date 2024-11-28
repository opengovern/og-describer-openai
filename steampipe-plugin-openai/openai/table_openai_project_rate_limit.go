package openai

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-openai/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiProjectRateLimit(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_project_rate_limit",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectRateLimit,
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
				Name:      "model",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Model"),
			},
			{
				Name:      "max_requests_per_1_minute",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.MaxRequestsPer1Minute"),
			},
			{
				Name:      "max_tokens_per_1_minute",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.MaxTokensPer1Minute"),
			},
			{
				Name:      "max_images_per_1_minute",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.MaxImagesPer1Minute"),
			},
			{
				Name:      "max_audio_megabytes_per_1_minute",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.MaxAudioMegabytesPer1Minute"),
			},
			{
				Name:      "max_requests_per_1_day",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.MaxRequestsPer1Day"),
			},
			{
				Name:      "batch_1_day_max_input_tokens",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("Description.Batch1DayMaxInputTokens"),
			},
			{
				Name:      "project_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.ProjectID"),
			},
		}),
	}
}
