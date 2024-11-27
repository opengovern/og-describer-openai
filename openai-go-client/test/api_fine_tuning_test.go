/*
OpenAI API

Testing FineTuningAPIService

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

func Test_openapi_FineTuningAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FineTuningAPIService CancelFineTuningJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fineTuningJobId string

		resp, httpRes, err := apiClient.FineTuningAPI.CancelFineTuningJob(context.Background(), fineTuningJobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FineTuningAPIService CreateFineTuningJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FineTuningAPI.CreateFineTuningJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FineTuningAPIService ListFineTuningEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fineTuningJobId string

		resp, httpRes, err := apiClient.FineTuningAPI.ListFineTuningEvents(context.Background(), fineTuningJobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FineTuningAPIService ListFineTuningJobCheckpoints", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fineTuningJobId string

		resp, httpRes, err := apiClient.FineTuningAPI.ListFineTuningJobCheckpoints(context.Background(), fineTuningJobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FineTuningAPIService ListPaginatedFineTuningJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FineTuningAPI.ListPaginatedFineTuningJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FineTuningAPIService RetrieveFineTuningJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fineTuningJobId string

		resp, httpRes, err := apiClient.FineTuningAPI.RetrieveFineTuningJob(context.Background(), fineTuningJobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
