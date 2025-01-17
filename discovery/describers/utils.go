package describers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/opengovern/og-describer-openai/discovery/provider"
	"net/http"
	"net/url"
)

const (
	order = "desc"
)



func getProjects(ctx context.Context, handler *model.OpenAIAPIHandler) ([]model.ProjectDescription, error) {
	var projects []model.ProjectDescription
	var projectResponse model.ProjectResponse
	var resp *http.Response
	baseURL := "https://api.openai.com/v1/organization/projects"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		var after string
		for {
			params := url.Values{}
			params.Set("limit", "100")
			if after != "" {
				params.Set("after", after)
			}
			finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&projectResponse); e != nil {
				return nil, e
			}
			projects = append(projects, projectResponse.Data...)
			if !projectResponse.HasMore {
				break
			}
			after = projectResponse.LastID
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return projects, nil
}



// isTemporary checks if an error is temporary.
func isTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}
