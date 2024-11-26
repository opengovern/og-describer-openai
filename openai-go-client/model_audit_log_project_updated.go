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

// checks if the AuditLogProjectUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogProjectUpdated{}

// AuditLogProjectUpdated The details for events with this `type`.
type AuditLogProjectUpdated struct {
	// The project ID.
	Id *string `json:"id,omitempty"`
	ChangesRequested *AuditLogProjectUpdatedChangesRequested `json:"changes_requested,omitempty"`
}

// NewAuditLogProjectUpdated instantiates a new AuditLogProjectUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogProjectUpdated() *AuditLogProjectUpdated {
	this := AuditLogProjectUpdated{}
	return &this
}

// NewAuditLogProjectUpdatedWithDefaults instantiates a new AuditLogProjectUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogProjectUpdatedWithDefaults() *AuditLogProjectUpdated {
	this := AuditLogProjectUpdated{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogProjectUpdated) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogProjectUpdated) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogProjectUpdated) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditLogProjectUpdated) SetId(v string) {
	o.Id = &v
}

// GetChangesRequested returns the ChangesRequested field value if set, zero value otherwise.
func (o *AuditLogProjectUpdated) GetChangesRequested() AuditLogProjectUpdatedChangesRequested {
	if o == nil || IsNil(o.ChangesRequested) {
		var ret AuditLogProjectUpdatedChangesRequested
		return ret
	}
	return *o.ChangesRequested
}

// GetChangesRequestedOk returns a tuple with the ChangesRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogProjectUpdated) GetChangesRequestedOk() (*AuditLogProjectUpdatedChangesRequested, bool) {
	if o == nil || IsNil(o.ChangesRequested) {
		return nil, false
	}
	return o.ChangesRequested, true
}

// HasChangesRequested returns a boolean if a field has been set.
func (o *AuditLogProjectUpdated) HasChangesRequested() bool {
	if o != nil && !IsNil(o.ChangesRequested) {
		return true
	}

	return false
}

// SetChangesRequested gets a reference to the given AuditLogProjectUpdatedChangesRequested and assigns it to the ChangesRequested field.
func (o *AuditLogProjectUpdated) SetChangesRequested(v AuditLogProjectUpdatedChangesRequested) {
	o.ChangesRequested = &v
}

func (o AuditLogProjectUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogProjectUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ChangesRequested) {
		toSerialize["changes_requested"] = o.ChangesRequested
	}
	return toSerialize, nil
}

type NullableAuditLogProjectUpdated struct {
	value *AuditLogProjectUpdated
	isSet bool
}

func (v NullableAuditLogProjectUpdated) Get() *AuditLogProjectUpdated {
	return v.value
}

func (v *NullableAuditLogProjectUpdated) Set(val *AuditLogProjectUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogProjectUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogProjectUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogProjectUpdated(val *AuditLogProjectUpdated) *NullableAuditLogProjectUpdated {
	return &NullableAuditLogProjectUpdated{value: val, isSet: true}
}

func (v NullableAuditLogProjectUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogProjectUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


