//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

package model

import (
	openai "github.com/opengovern/og-describer-openai/openai-go-client"
	"time"
)

type Metadata struct{}

type FileDescription struct {
	ID        string
	FileName  string
	CreatedAt time.Time
	Bytes     int32
	Object    string
	Purpose   string
}

type ModelsDescription struct {
	ID        string
	CreatedAt time.Time
	Object    string
	OwnedBy   string
}

type AssistantDescription struct {
	ID string
	//Object         string
	//CreatedAt      time.Time
	Name           string
	Description    string
	Model          string
	Instructions   string
	Tools          []openai.AssistantObjectToolsInner
	ToolResources  openai.AssistantObjectToolResources
	Metadata       map[string]interface{}
	Temperature    float32
	TopP           float32
	ResponseFormat *openai.AssistantsApiResponseFormatOption
}

type VectorStoreDescription struct {
	Id           string
	Object       string
	CreatedAt    time.Time
	Name         string
	UsageBytes   int32
	FileCounts   openai.VectorStoreObjectFileCounts
	Status       string
	ExpiresAfter *openai.VectorStoreExpirationAfter
	ExpiresAt    time.Time
	LastActiveAt time.Time
	Metadata     map[string]interface{}
}

type ProjectDescription struct {
	ID         string
	Object     string
	Name       string
	CreatedAt  time.Time
	ArchivedAt time.Time
	Status     string
}

type ProjectUserDescription struct {
	UserID    string
	ProjectID string
	//Object  string
	//Name    string
	//Email   string
	//Role    string
	//AddedAt time.Time
}

type ProjectServiceAccountDescription struct {
	Object    string
	ID        string
	Name      string
	Role      string
	CreatedAt time.Time
}

type ProjectApiKeyDescription struct {
	Object        string
	ID            string
	Name          string
	RedactedValue string
	CreatedAt     time.Time
	Owner         openai.ProjectApiKeyOwner
}

type ProjectRateLimitDescription struct {
	Object                      string
	ID                          string
	Model                       string
	MaxRequestsPer1Minute       int32
	MaxTokensPer1Minute         int32
	MaxImagesPer1Minute         *int32
	MaxAudioMegabytesPer1Minute *int32
	MaxRequestsPer1Day          *int32
	Batch1DayMaxInputTokens     *int32
}
