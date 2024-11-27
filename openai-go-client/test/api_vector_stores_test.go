/*
OpenAI API

Testing VectorStoresAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/opengovern/og-describer-openai/openai-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_VectorStoresAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VectorStoresAPIService CancelVectorStoreFileBatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string
		var batchId string

		resp, httpRes, err := apiClient.VectorStoresAPI.CancelVectorStoreFileBatch(context.Background(), vectorStoreId, batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService CreateVectorStore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VectorStoresAPI.CreateVectorStore(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService CreateVectorStoreFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string

		resp, httpRes, err := apiClient.VectorStoresAPI.CreateVectorStoreFile(context.Background(), vectorStoreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService CreateVectorStoreFileBatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string

		resp, httpRes, err := apiClient.VectorStoresAPI.CreateVectorStoreFileBatch(context.Background(), vectorStoreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService DeleteVectorStore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string

		resp, httpRes, err := apiClient.VectorStoresAPI.DeleteVectorStore(context.Background(), vectorStoreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService DeleteVectorStoreFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string
		var fileId string

		resp, httpRes, err := apiClient.VectorStoresAPI.DeleteVectorStoreFile(context.Background(), vectorStoreId, fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService GetVectorStore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string

		resp, httpRes, err := apiClient.VectorStoresAPI.GetVectorStore(context.Background(), vectorStoreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService GetVectorStoreFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string
		var fileId string

		resp, httpRes, err := apiClient.VectorStoresAPI.GetVectorStoreFile(context.Background(), vectorStoreId, fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService GetVectorStoreFileBatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string
		var batchId string

		resp, httpRes, err := apiClient.VectorStoresAPI.GetVectorStoreFileBatch(context.Background(), vectorStoreId, batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService ListFilesInVectorStoreBatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string
		var batchId string

		resp, httpRes, err := apiClient.VectorStoresAPI.ListFilesInVectorStoreBatch(context.Background(), vectorStoreId, batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService ListVectorStoreFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string

		resp, httpRes, err := apiClient.VectorStoresAPI.ListVectorStoreFiles(context.Background(), vectorStoreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService ListVectorStores", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VectorStoresAPI.ListVectorStores(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VectorStoresAPIService ModifyVectorStore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vectorStoreId string

		resp, httpRes, err := apiClient.VectorStoresAPI.ModifyVectorStore(context.Background(), vectorStoreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
