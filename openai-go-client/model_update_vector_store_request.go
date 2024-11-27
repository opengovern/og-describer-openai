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

// checks if the UpdateVectorStoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVectorStoreRequest{}

// UpdateVectorStoreRequest struct for UpdateVectorStoreRequest
type UpdateVectorStoreRequest struct {
	// The name of the vector store.
	Name NullableString `json:"name,omitempty"`
	ExpiresAfter *VectorStoreExpirationAfter `json:"expires_after,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewUpdateVectorStoreRequest instantiates a new UpdateVectorStoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVectorStoreRequest() *UpdateVectorStoreRequest {
	this := UpdateVectorStoreRequest{}
	return &this
}

// NewUpdateVectorStoreRequestWithDefaults instantiates a new UpdateVectorStoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVectorStoreRequestWithDefaults() *UpdateVectorStoreRequest {
	this := UpdateVectorStoreRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVectorStoreRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVectorStoreRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVectorStoreRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateVectorStoreRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateVectorStoreRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateVectorStoreRequest) UnsetName() {
	o.Name.Unset()
}

// GetExpiresAfter returns the ExpiresAfter field value if set, zero value otherwise.
func (o *UpdateVectorStoreRequest) GetExpiresAfter() VectorStoreExpirationAfter {
	if o == nil || IsNil(o.ExpiresAfter) {
		var ret VectorStoreExpirationAfter
		return ret
	}
	return *o.ExpiresAfter
}

// GetExpiresAfterOk returns a tuple with the ExpiresAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVectorStoreRequest) GetExpiresAfterOk() (*VectorStoreExpirationAfter, bool) {
	if o == nil || IsNil(o.ExpiresAfter) {
		return nil, false
	}
	return o.ExpiresAfter, true
}

// HasExpiresAfter returns a boolean if a field has been set.
func (o *UpdateVectorStoreRequest) HasExpiresAfter() bool {
	if o != nil && !IsNil(o.ExpiresAfter) {
		return true
	}

	return false
}

// SetExpiresAfter gets a reference to the given VectorStoreExpirationAfter and assigns it to the ExpiresAfter field.
func (o *UpdateVectorStoreRequest) SetExpiresAfter(v VectorStoreExpirationAfter) {
	o.ExpiresAfter = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVectorStoreRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVectorStoreRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateVectorStoreRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateVectorStoreRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UpdateVectorStoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVectorStoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ExpiresAfter) {
		toSerialize["expires_after"] = o.ExpiresAfter
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableUpdateVectorStoreRequest struct {
	value *UpdateVectorStoreRequest
	isSet bool
}

func (v NullableUpdateVectorStoreRequest) Get() *UpdateVectorStoreRequest {
	return v.value
}

func (v *NullableUpdateVectorStoreRequest) Set(val *UpdateVectorStoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVectorStoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVectorStoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVectorStoreRequest(val *UpdateVectorStoreRequest) *NullableUpdateVectorStoreRequest {
	return &NullableUpdateVectorStoreRequest{value: val, isSet: true}
}

func (v NullableUpdateVectorStoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVectorStoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

