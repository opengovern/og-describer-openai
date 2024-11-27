/*
OpenAI API

Testing AssistantsAPIService

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

func Test_openapi_AssistantsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssistantsAPIService CancelRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.AssistantsAPI.CancelRun(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService CreateAssistant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AssistantsAPI.CreateAssistant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService CreateMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.CreateMessage(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService CreateRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.CreateRun(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService CreateThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AssistantsAPI.CreateThread(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService CreateThreadAndRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AssistantsAPI.CreateThreadAndRun(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService DeleteAssistant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var assistantId string

		resp, httpRes, err := apiClient.AssistantsAPI.DeleteAssistant(context.Background(), assistantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService DeleteMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var messageId string

		resp, httpRes, err := apiClient.AssistantsAPI.DeleteMessage(context.Background(), threadId, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService DeleteThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.DeleteThread(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService GetAssistant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var assistantId string

		resp, httpRes, err := apiClient.AssistantsAPI.GetAssistant(context.Background(), assistantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService GetMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var messageId string

		resp, httpRes, err := apiClient.AssistantsAPI.GetMessage(context.Background(), threadId, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService GetRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.AssistantsAPI.GetRun(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService GetRunStep", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string
		var stepId string

		resp, httpRes, err := apiClient.AssistantsAPI.GetRunStep(context.Background(), threadId, runId, stepId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService GetThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.GetThread(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ListAssistants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AssistantsAPI.ListAssistants(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ListMessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.ListMessages(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ListRunSteps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.AssistantsAPI.ListRunSteps(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ListRuns", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.ListRuns(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ModifyAssistant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var assistantId string

		resp, httpRes, err := apiClient.AssistantsAPI.ModifyAssistant(context.Background(), assistantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ModifyMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var messageId string

		resp, httpRes, err := apiClient.AssistantsAPI.ModifyMessage(context.Background(), threadId, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ModifyRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.AssistantsAPI.ModifyRun(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService ModifyThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.AssistantsAPI.ModifyThread(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssistantsAPIService SubmitToolOuputsToRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.AssistantsAPI.SubmitToolOuputsToRun(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}