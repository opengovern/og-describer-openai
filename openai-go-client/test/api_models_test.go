/*
OpenAI API

Testing ModelsAPIService

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

func Test_openapi_ModelsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ModelsAPIService DeleteModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var model string

		resp, httpRes, err := apiClient.ModelsAPI.DeleteModel(context.Background(), model).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService ListModels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ModelsAPI.ListModels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelsAPIService RetrieveModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var model string

		resp, httpRes, err := apiClient.ModelsAPI.RetrieveModel(context.Background(), model).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
