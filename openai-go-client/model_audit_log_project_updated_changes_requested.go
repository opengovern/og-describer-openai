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

// checks if the AuditLogProjectUpdatedChangesRequested type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogProjectUpdatedChangesRequested{}

// AuditLogProjectUpdatedChangesRequested The payload used to update the project.
type AuditLogProjectUpdatedChangesRequested struct {
	// The title of the project as seen on the dashboard.
	Title *string `json:"title,omitempty"`
}

// NewAuditLogProjectUpdatedChangesRequested instantiates a new AuditLogProjectUpdatedChangesRequested object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogProjectUpdatedChangesRequested() *AuditLogProjectUpdatedChangesRequested {
	this := AuditLogProjectUpdatedChangesRequested{}
	return &this
}

// NewAuditLogProjectUpdatedChangesRequestedWithDefaults instantiates a new AuditLogProjectUpdatedChangesRequested object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogProjectUpdatedChangesRequestedWithDefaults() *AuditLogProjectUpdatedChangesRequested {
	this := AuditLogProjectUpdatedChangesRequested{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AuditLogProjectUpdatedChangesRequested) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogProjectUpdatedChangesRequested) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AuditLogProjectUpdatedChangesRequested) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AuditLogProjectUpdatedChangesRequested) SetTitle(v string) {
	o.Title = &v
}

func (o AuditLogProjectUpdatedChangesRequested) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogProjectUpdatedChangesRequested) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAuditLogProjectUpdatedChangesRequested struct {
	value *AuditLogProjectUpdatedChangesRequested
	isSet bool
}

func (v NullableAuditLogProjectUpdatedChangesRequested) Get() *AuditLogProjectUpdatedChangesRequested {
	return v.value
}

func (v *NullableAuditLogProjectUpdatedChangesRequested) Set(val *AuditLogProjectUpdatedChangesRequested) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogProjectUpdatedChangesRequested) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogProjectUpdatedChangesRequested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogProjectUpdatedChangesRequested(val *AuditLogProjectUpdatedChangesRequested) *NullableAuditLogProjectUpdatedChangesRequested {
	return &NullableAuditLogProjectUpdatedChangesRequested{value: val, isSet: true}
}

func (v NullableAuditLogProjectUpdatedChangesRequested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogProjectUpdatedChangesRequested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


