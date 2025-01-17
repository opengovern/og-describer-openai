package openai

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
