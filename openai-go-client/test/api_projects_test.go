/*
OpenAI API

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/opengovern/og-describer-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService ArchiveProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ArchiveProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService CreateProjectServiceAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.CreateProjectServiceAccount(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService CreateProjectUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.CreateProjectUser(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProjectApiKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var keyId string

		resp, httpRes, err := apiClient.ProjectsAPI.DeleteProjectApiKey(context.Background(), projectId, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProjectServiceAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var serviceAccountId string

		resp, httpRes, err := apiClient.ProjectsAPI.DeleteProjectServiceAccount(context.Background(), projectId, serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProjectUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string

		resp, httpRes, err := apiClient.ProjectsAPI.DeleteProjectUser(context.Background(), projectId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListProjectApiKeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ListProjectApiKeys(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListProjectRateLimits", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ListProjectRateLimits(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListProjectServiceAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ListProjectServiceAccounts(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListProjectUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ListProjectUsers(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListProjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.ListProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ModifyProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.ModifyProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ModifyProjectUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string

		resp, httpRes, err := apiClient.ProjectsAPI.ModifyProjectUser(context.Background(), projectId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RetrieveProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.RetrieveProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RetrieveProjectApiKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var keyId string

		resp, httpRes, err := apiClient.ProjectsAPI.RetrieveProjectApiKey(context.Background(), projectId, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RetrieveProjectServiceAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var serviceAccountId string

		resp, httpRes, err := apiClient.ProjectsAPI.RetrieveProjectServiceAccount(context.Background(), projectId, serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService RetrieveProjectUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var userId string

		resp, httpRes, err := apiClient.ProjectsAPI.RetrieveProjectUser(context.Background(), projectId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService UpdateProjectRateLimits", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var rateLimitId string

		resp, httpRes, err := apiClient.ProjectsAPI.UpdateProjectRateLimits(context.Background(), projectId, rateLimitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}