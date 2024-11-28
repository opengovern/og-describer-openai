package openai

import (
	"context"
	"encoding/json"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append(c, []*plugin.Column{
		{
			Name:        "platform_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Platform Account ID in which the resource is located.",
			Transform:   transform.FromField("IntegrationID"),
		},
		{
			Name:        "platform_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("PlatformID"),
		},
		{
			Name:        "platform_metadata",
			Type:        proto.ColumnType_JSON,
			Description: "The metadata of the resource",
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
		{
			Name:        "platform_resource_description",
			Type:        proto.ColumnType_JSON,
			Description: "The full model description of the resource",
			Transform:   transform.FromField("Description").Transform(marshalJSON),
		},
	}...)
}

func marshalJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	b, err := json.Marshal(d.Value)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Profile represents the structure of our profile data.
type Profile struct {
	UserID    string `json:"id"`
	FirstName string `json:"name"`
	Email     string `json:"email"`
	Orgs      struct {
		Data []OrganizationInfo `json:"data"`
	} `json:"orgs"`
}

type OrganizationInfo struct {
	Object      string `json:"object"`
	OrgId       string `json:"id"`
	Created     int64  `json:"created"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Personal    bool   `json:"personal"`
	ParentOrgId string `json:"parent_org_id"`
}
