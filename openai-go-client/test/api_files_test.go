/*
OpenAI API

Testing FilesAPIService

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

func Test_openapi_FilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FilesAPIService CreateFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.CreateFile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService DeleteFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.DeleteFile(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService DownloadFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.DownloadFile(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService ListFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.ListFiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService RetrieveFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.RetrieveFile(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}