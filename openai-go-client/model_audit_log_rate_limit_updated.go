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

// checks if the AuditLogRateLimitUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogRateLimitUpdated{}

// AuditLogRateLimitUpdated The details for events with this `type`.
type AuditLogRateLimitUpdated struct {
	// The rate limit ID
	Id *string `json:"id,omitempty"`
	ChangesRequested *AuditLogRateLimitUpdatedChangesRequested `json:"changes_requested,omitempty"`
}

// NewAuditLogRateLimitUpdated instantiates a new AuditLogRateLimitUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogRateLimitUpdated() *AuditLogRateLimitUpdated {
	this := AuditLogRateLimitUpdated{}
	return &this
}

// NewAuditLogRateLimitUpdatedWithDefaults instantiates a new AuditLogRateLimitUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogRateLimitUpdatedWithDefaults() *AuditLogRateLimitUpdated {
	this := AuditLogRateLimitUpdated{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogRateLimitUpdated) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogRateLimitUpdated) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogRateLimitUpdated) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditLogRateLimitUpdated) SetId(v string) {
	o.Id = &v
}

// GetChangesRequested returns the ChangesRequested field value if set, zero value otherwise.
func (o *AuditLogRateLimitUpdated) GetChangesRequested() AuditLogRateLimitUpdatedChangesRequested {
	if o == nil || IsNil(o.ChangesRequested) {
		var ret AuditLogRateLimitUpdatedChangesRequested
		return ret
	}
	return *o.ChangesRequested
}

// GetChangesRequestedOk returns a tuple with the ChangesRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogRateLimitUpdated) GetChangesRequestedOk() (*AuditLogRateLimitUpdatedChangesRequested, bool) {
	if o == nil || IsNil(o.ChangesRequested) {
		return nil, false
	}
	return o.ChangesRequested, true
}

// HasChangesRequested returns a boolean if a field has been set.
func (o *AuditLogRateLimitUpdated) HasChangesRequested() bool {
	if o != nil && !IsNil(o.ChangesRequested) {
		return true
	}

	return false
}

// SetChangesRequested gets a reference to the given AuditLogRateLimitUpdatedChangesRequested and assigns it to the ChangesRequested field.
func (o *AuditLogRateLimitUpdated) SetChangesRequested(v AuditLogRateLimitUpdatedChangesRequested) {
	o.ChangesRequested = &v
}

func (o AuditLogRateLimitUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogRateLimitUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ChangesRequested) {
		toSerialize["changes_requested"] = o.ChangesRequested
	}
	return toSerialize, nil
}

type NullableAuditLogRateLimitUpdated struct {
	value *AuditLogRateLimitUpdated
	isSet bool
}

func (v NullableAuditLogRateLimitUpdated) Get() *AuditLogRateLimitUpdated {
	return v.value
}

func (v *NullableAuditLogRateLimitUpdated) Set(val *AuditLogRateLimitUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogRateLimitUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogRateLimitUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogRateLimitUpdated(val *AuditLogRateLimitUpdated) *NullableAuditLogRateLimitUpdated {
	return &NullableAuditLogRateLimitUpdated{value: val, isSet: true}
}

func (v NullableAuditLogRateLimitUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogRateLimitUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

