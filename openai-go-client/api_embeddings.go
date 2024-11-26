/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// EmbeddingsAPIService EmbeddingsAPI service
type EmbeddingsAPIService service

type ApiCreateEmbeddingRequest struct {
	ctx context.Context
	ApiService *EmbeddingsAPIService
	createEmbeddingRequest *CreateEmbeddingRequest
}

func (r ApiCreateEmbeddingRequest) CreateEmbeddingRequest(createEmbeddingRequest CreateEmbeddingRequest) ApiCreateEmbeddingRequest {
	r.createEmbeddingRequest = &createEmbeddingRequest
	return r
}

func (r ApiCreateEmbeddingRequest) Execute() (*CreateEmbeddingResponse, *http.Response, error) {
	return r.ApiService.CreateEmbeddingExecute(r)
}

/*
CreateEmbedding Creates an embedding vector representing the input text.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateEmbeddingRequest
*/
func (a *EmbeddingsAPIService) CreateEmbedding(ctx context.Context) ApiCreateEmbeddingRequest {
	return ApiCreateEmbeddingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateEmbeddingResponse
func (a *EmbeddingsAPIService) CreateEmbeddingExecute(r ApiCreateEmbeddingRequest) (*CreateEmbeddingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateEmbeddingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbeddingsAPIService.CreateEmbedding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/embeddings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createEmbeddingRequest == nil {
		return localVarReturnValue, nil, reportError("createEmbeddingRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createEmbeddingRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
