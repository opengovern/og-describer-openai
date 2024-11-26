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
	"reflect"
)


// AuditLogsAPIService AuditLogsAPI service
type AuditLogsAPIService service

type ApiListAuditLogsRequest struct {
	ctx context.Context
	ApiService *AuditLogsAPIService
	effectiveAt *ListAuditLogsEffectiveAtParameter
	projectIds *[]string
	eventTypes *[]AuditLogEventType
	actorIds *[]string
	actorEmails *[]string
	resourceIds *[]string
	limit *int32
	after *string
	before *string
}

// Return only events whose &#x60;effective_at&#x60; (Unix seconds) is in this range.
func (r ApiListAuditLogsRequest) EffectiveAt(effectiveAt ListAuditLogsEffectiveAtParameter) ApiListAuditLogsRequest {
	r.effectiveAt = &effectiveAt
	return r
}

// Return only events for these projects.
func (r ApiListAuditLogsRequest) ProjectIds(projectIds []string) ApiListAuditLogsRequest {
	r.projectIds = &projectIds
	return r
}

// Return only events with a &#x60;type&#x60; in one of these values. For example, &#x60;project.created&#x60;. For all options, see the documentation for the [audit log object](/docs/api-reference/audit-logs/object).
func (r ApiListAuditLogsRequest) EventTypes(eventTypes []AuditLogEventType) ApiListAuditLogsRequest {
	r.eventTypes = &eventTypes
	return r
}

// Return only events performed by these actors. Can be a user ID, a service account ID, or an api key tracking ID.
func (r ApiListAuditLogsRequest) ActorIds(actorIds []string) ApiListAuditLogsRequest {
	r.actorIds = &actorIds
	return r
}

// Return only events performed by users with these emails.
func (r ApiListAuditLogsRequest) ActorEmails(actorEmails []string) ApiListAuditLogsRequest {
	r.actorEmails = &actorEmails
	return r
}

// Return only events performed on these targets. For example, a project ID updated.
func (r ApiListAuditLogsRequest) ResourceIds(resourceIds []string) ApiListAuditLogsRequest {
	r.resourceIds = &resourceIds
	return r
}

// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. 
func (r ApiListAuditLogsRequest) Limit(limit int32) ApiListAuditLogsRequest {
	r.limit = &limit
	return r
}

// A cursor for use in pagination. &#x60;after&#x60; is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after&#x3D;obj_foo in order to fetch the next page of the list. 
func (r ApiListAuditLogsRequest) After(after string) ApiListAuditLogsRequest {
	r.after = &after
	return r
}

// A cursor for use in pagination. &#x60;before&#x60; is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with obj_foo, your subsequent call can include before&#x3D;obj_foo in order to fetch the previous page of the list. 
func (r ApiListAuditLogsRequest) Before(before string) ApiListAuditLogsRequest {
	r.before = &before
	return r
}

func (r ApiListAuditLogsRequest) Execute() (*ListAuditLogsResponse, *http.Response, error) {
	return r.ApiService.ListAuditLogsExecute(r)
}

/*
ListAuditLogs List user actions and configuration changes within this organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAuditLogsRequest
*/
func (a *AuditLogsAPIService) ListAuditLogs(ctx context.Context) ApiListAuditLogsRequest {
	return ApiListAuditLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListAuditLogsResponse
func (a *AuditLogsAPIService) ListAuditLogsExecute(r ApiListAuditLogsRequest) (*ListAuditLogsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAuditLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsAPIService.ListAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/audit_logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.effectiveAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "effective_at", r.effectiveAt, "form", "")
	}
	if r.projectIds != nil {
		t := *r.projectIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "project_ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "project_ids[]", t, "form", "multi")
		}
	}
	if r.eventTypes != nil {
		t := *r.eventTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "event_types[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "event_types[]", t, "form", "multi")
		}
	}
	if r.actorIds != nil {
		t := *r.actorIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "actor_ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "actor_ids[]", t, "form", "multi")
		}
	}
	if r.actorEmails != nil {
		t := *r.actorEmails
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "actor_emails[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "actor_emails[]", t, "form", "multi")
		}
	}
	if r.resourceIds != nil {
		t := *r.resourceIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "resource_ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "resource_ids[]", t, "form", "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 20
		r.limit = &defaultValue
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
