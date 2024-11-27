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

// checks if the ModifyThreadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyThreadRequest{}

// ModifyThreadRequest struct for ModifyThreadRequest
type ModifyThreadRequest struct {
	ToolResources NullableModifyThreadRequestToolResources `json:"tool_resources,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewModifyThreadRequest instantiates a new ModifyThreadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyThreadRequest() *ModifyThreadRequest {
	this := ModifyThreadRequest{}
	return &this
}

// NewModifyThreadRequestWithDefaults instantiates a new ModifyThreadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyThreadRequestWithDefaults() *ModifyThreadRequest {
	this := ModifyThreadRequest{}
	return &this
}

// GetToolResources returns the ToolResources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyThreadRequest) GetToolResources() ModifyThreadRequestToolResources {
	if o == nil || IsNil(o.ToolResources.Get()) {
		var ret ModifyThreadRequestToolResources
		return ret
	}
	return *o.ToolResources.Get()
}

// GetToolResourcesOk returns a tuple with the ToolResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyThreadRequest) GetToolResourcesOk() (*ModifyThreadRequestToolResources, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolResources.Get(), o.ToolResources.IsSet()
}

// HasToolResources returns a boolean if a field has been set.
func (o *ModifyThreadRequest) HasToolResources() bool {
	if o != nil && o.ToolResources.IsSet() {
		return true
	}

	return false
}

// SetToolResources gets a reference to the given NullableModifyThreadRequestToolResources and assigns it to the ToolResources field.
func (o *ModifyThreadRequest) SetToolResources(v ModifyThreadRequestToolResources) {
	o.ToolResources.Set(&v)
}
// SetToolResourcesNil sets the value for ToolResources to be an explicit nil
func (o *ModifyThreadRequest) SetToolResourcesNil() {
	o.ToolResources.Set(nil)
}

// UnsetToolResources ensures that no value is present for ToolResources, not even an explicit nil
func (o *ModifyThreadRequest) UnsetToolResources() {
	o.ToolResources.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyThreadRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyThreadRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModifyThreadRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ModifyThreadRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ModifyThreadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyThreadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ToolResources.IsSet() {
		toSerialize["tool_resources"] = o.ToolResources.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableModifyThreadRequest struct {
	value *ModifyThreadRequest
	isSet bool
}

func (v NullableModifyThreadRequest) Get() *ModifyThreadRequest {
	return v.value
}

func (v *NullableModifyThreadRequest) Set(val *ModifyThreadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyThreadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyThreadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyThreadRequest(val *ModifyThreadRequest) *NullableModifyThreadRequest {
	return &NullableModifyThreadRequest{value: val, isSet: true}
}

func (v NullableModifyThreadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyThreadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

