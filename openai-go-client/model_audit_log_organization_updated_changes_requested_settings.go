/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuditLogOrganizationUpdatedChangesRequestedSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogOrganizationUpdatedChangesRequestedSettings{}

// AuditLogOrganizationUpdatedChangesRequestedSettings struct for AuditLogOrganizationUpdatedChangesRequestedSettings
type AuditLogOrganizationUpdatedChangesRequestedSettings struct {
	// Visibility of the threads page which shows messages created with the Assistants API and Playground. One of `ANY_ROLE`, `OWNERS`, or `NONE`.
	ThreadsUiVisibility *string `json:"threads_ui_visibility,omitempty"`
	// Visibility of the usage dashboard which shows activity and costs for your organization. One of `ANY_ROLE` or `OWNERS`.
	UsageDashboardVisibility *string `json:"usage_dashboard_visibility,omitempty"`
}

// NewAuditLogOrganizationUpdatedChangesRequestedSettings instantiates a new AuditLogOrganizationUpdatedChangesRequestedSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogOrganizationUpdatedChangesRequestedSettings() *AuditLogOrganizationUpdatedChangesRequestedSettings {
	this := AuditLogOrganizationUpdatedChangesRequestedSettings{}
	return &this
}

// NewAuditLogOrganizationUpdatedChangesRequestedSettingsWithDefaults instantiates a new AuditLogOrganizationUpdatedChangesRequestedSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogOrganizationUpdatedChangesRequestedSettingsWithDefaults() *AuditLogOrganizationUpdatedChangesRequestedSettings {
	this := AuditLogOrganizationUpdatedChangesRequestedSettings{}
	return &this
}

// GetThreadsUiVisibility returns the ThreadsUiVisibility field value if set, zero value otherwise.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) GetThreadsUiVisibility() string {
	if o == nil || IsNil(o.ThreadsUiVisibility) {
		var ret string
		return ret
	}
	return *o.ThreadsUiVisibility
}

// GetThreadsUiVisibilityOk returns a tuple with the ThreadsUiVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) GetThreadsUiVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadsUiVisibility) {
		return nil, false
	}
	return o.ThreadsUiVisibility, true
}

// HasThreadsUiVisibility returns a boolean if a field has been set.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) HasThreadsUiVisibility() bool {
	if o != nil && !IsNil(o.ThreadsUiVisibility) {
		return true
	}

	return false
}

// SetThreadsUiVisibility gets a reference to the given string and assigns it to the ThreadsUiVisibility field.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) SetThreadsUiVisibility(v string) {
	o.ThreadsUiVisibility = &v
}

// GetUsageDashboardVisibility returns the UsageDashboardVisibility field value if set, zero value otherwise.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) GetUsageDashboardVisibility() string {
	if o == nil || IsNil(o.UsageDashboardVisibility) {
		var ret string
		return ret
	}
	return *o.UsageDashboardVisibility
}

// GetUsageDashboardVisibilityOk returns a tuple with the UsageDashboardVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) GetUsageDashboardVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.UsageDashboardVisibility) {
		return nil, false
	}
	return o.UsageDashboardVisibility, true
}

// HasUsageDashboardVisibility returns a boolean if a field has been set.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) HasUsageDashboardVisibility() bool {
	if o != nil && !IsNil(o.UsageDashboardVisibility) {
		return true
	}

	return false
}

// SetUsageDashboardVisibility gets a reference to the given string and assigns it to the UsageDashboardVisibility field.
func (o *AuditLogOrganizationUpdatedChangesRequestedSettings) SetUsageDashboardVisibility(v string) {
	o.UsageDashboardVisibility = &v
}

func (o AuditLogOrganizationUpdatedChangesRequestedSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogOrganizationUpdatedChangesRequestedSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThreadsUiVisibility) {
		toSerialize["threads_ui_visibility"] = o.ThreadsUiVisibility
	}
	if !IsNil(o.UsageDashboardVisibility) {
		toSerialize["usage_dashboard_visibility"] = o.UsageDashboardVisibility
	}
	return toSerialize, nil
}

type NullableAuditLogOrganizationUpdatedChangesRequestedSettings struct {
	value *AuditLogOrganizationUpdatedChangesRequestedSettings
	isSet bool
}

func (v NullableAuditLogOrganizationUpdatedChangesRequestedSettings) Get() *AuditLogOrganizationUpdatedChangesRequestedSettings {
	return v.value
}

func (v *NullableAuditLogOrganizationUpdatedChangesRequestedSettings) Set(val *AuditLogOrganizationUpdatedChangesRequestedSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogOrganizationUpdatedChangesRequestedSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogOrganizationUpdatedChangesRequestedSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogOrganizationUpdatedChangesRequestedSettings(val *AuditLogOrganizationUpdatedChangesRequestedSettings) *NullableAuditLogOrganizationUpdatedChangesRequestedSettings {
	return &NullableAuditLogOrganizationUpdatedChangesRequestedSettings{value: val, isSet: true}
}

func (v NullableAuditLogOrganizationUpdatedChangesRequestedSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogOrganizationUpdatedChangesRequestedSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


