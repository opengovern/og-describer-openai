package openai

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-openai",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"openai_file":                    tableOpenAiFile(ctx),
			"openai_model":                   tableOpenAiModel(ctx),
			"openai_project":                 tableOpenAiProject(ctx),
			"openai_project_user":            tableOpenAiProjectUser(ctx),
			"openai_project_api_key":         tableOpenAiProjectApiKey(ctx),
			"openai_project_rate_limit":      tableOpenAiProjectRateLimit(ctx),
			"openai_project_service_account": tableOpenAiProjectServiceAccount(ctx),
			"openai_vector_store":            tableOpenAiVectorStore(ctx),
			"openai_assistant":               tableOpenAiAssistant(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_account_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
