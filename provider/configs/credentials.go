package configs

type IntegrationCredentials struct {
	APIKey         string `json:"api_key"`
	OrganizationID string `json:"organization_id"`
	ProjectID      string `json:"project_id"`
}
