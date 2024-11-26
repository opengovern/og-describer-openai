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

// checks if the AuditLogOrganizationUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogOrganizationUpdated{}

// AuditLogOrganizationUpdated The details for events with this `type`.
type AuditLogOrganizationUpdated struct {
	// The organization ID.
	Id *string `json:"id,omitempty"`
	ChangesRequested *AuditLogOrganizationUpdatedChangesRequested `json:"changes_requested,omitempty"`
}

// NewAuditLogOrganizationUpdated instantiates a new AuditLogOrganizationUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogOrganizationUpdated() *AuditLogOrganizationUpdated {
	this := AuditLogOrganizationUpdated{}
	return &this
}

// NewAuditLogOrganizationUpdatedWithDefaults instantiates a new AuditLogOrganizationUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogOrganizationUpdatedWithDefaults() *AuditLogOrganizationUpdated {
	this := AuditLogOrganizationUpdated{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogOrganizationUpdated) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogOrganizationUpdated) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogOrganizationUpdated) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditLogOrganizationUpdated) SetId(v string) {
	o.Id = &v
}

// GetChangesRequested returns the ChangesRequested field value if set, zero value otherwise.
func (o *AuditLogOrganizationUpdated) GetChangesRequested() AuditLogOrganizationUpdatedChangesRequested {
	if o == nil || IsNil(o.ChangesRequested) {
		var ret AuditLogOrganizationUpdatedChangesRequested
		return ret
	}
	return *o.ChangesRequested
}

// GetChangesRequestedOk returns a tuple with the ChangesRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogOrganizationUpdated) GetChangesRequestedOk() (*AuditLogOrganizationUpdatedChangesRequested, bool) {
	if o == nil || IsNil(o.ChangesRequested) {
		return nil, false
	}
	return o.ChangesRequested, true
}

// HasChangesRequested returns a boolean if a field has been set.
func (o *AuditLogOrganizationUpdated) HasChangesRequested() bool {
	if o != nil && !IsNil(o.ChangesRequested) {
		return true
	}

	return false
}

// SetChangesRequested gets a reference to the given AuditLogOrganizationUpdatedChangesRequested and assigns it to the ChangesRequested field.
func (o *AuditLogOrganizationUpdated) SetChangesRequested(v AuditLogOrganizationUpdatedChangesRequested) {
	o.ChangesRequested = &v
}

func (o AuditLogOrganizationUpdated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogOrganizationUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ChangesRequested) {
		toSerialize["changes_requested"] = o.ChangesRequested
	}
	return toSerialize, nil
}

type NullableAuditLogOrganizationUpdated struct {
	value *AuditLogOrganizationUpdated
	isSet bool
}

func (v NullableAuditLogOrganizationUpdated) Get() *AuditLogOrganizationUpdated {
	return v.value
}

func (v *NullableAuditLogOrganizationUpdated) Set(val *AuditLogOrganizationUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogOrganizationUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogOrganizationUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogOrganizationUpdated(val *AuditLogOrganizationUpdated) *NullableAuditLogOrganizationUpdated {
	return &NullableAuditLogOrganizationUpdated{value: val, isSet: true}
}

func (v NullableAuditLogOrganizationUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogOrganizationUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


