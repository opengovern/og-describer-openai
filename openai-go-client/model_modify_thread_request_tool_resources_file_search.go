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

// checks if the ModifyThreadRequestToolResourcesFileSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyThreadRequestToolResourcesFileSearch{}

// ModifyThreadRequestToolResourcesFileSearch struct for ModifyThreadRequestToolResourcesFileSearch
type ModifyThreadRequestToolResourcesFileSearch struct {
	// The [vector store](/docs/api-reference/vector-stores/object) attached to this thread. There can be a maximum of 1 vector store attached to the thread. 
	VectorStoreIds []string `json:"vector_store_ids,omitempty"`
}

// NewModifyThreadRequestToolResourcesFileSearch instantiates a new ModifyThreadRequestToolResourcesFileSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyThreadRequestToolResourcesFileSearch() *ModifyThreadRequestToolResourcesFileSearch {
	this := ModifyThreadRequestToolResourcesFileSearch{}
	return &this
}

// NewModifyThreadRequestToolResourcesFileSearchWithDefaults instantiates a new ModifyThreadRequestToolResourcesFileSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyThreadRequestToolResourcesFileSearchWithDefaults() *ModifyThreadRequestToolResourcesFileSearch {
	this := ModifyThreadRequestToolResourcesFileSearch{}
	return &this
}

// GetVectorStoreIds returns the VectorStoreIds field value if set, zero value otherwise.
func (o *ModifyThreadRequestToolResourcesFileSearch) GetVectorStoreIds() []string {
	if o == nil || IsNil(o.VectorStoreIds) {
		var ret []string
		return ret
	}
	return o.VectorStoreIds
}

// GetVectorStoreIdsOk returns a tuple with the VectorStoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyThreadRequestToolResourcesFileSearch) GetVectorStoreIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VectorStoreIds) {
		return nil, false
	}
	return o.VectorStoreIds, true
}

// HasVectorStoreIds returns a boolean if a field has been set.
func (o *ModifyThreadRequestToolResourcesFileSearch) HasVectorStoreIds() bool {
	if o != nil && !IsNil(o.VectorStoreIds) {
		return true
	}

	return false
}

// SetVectorStoreIds gets a reference to the given []string and assigns it to the VectorStoreIds field.
func (o *ModifyThreadRequestToolResourcesFileSearch) SetVectorStoreIds(v []string) {
	o.VectorStoreIds = v
}

func (o ModifyThreadRequestToolResourcesFileSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyThreadRequestToolResourcesFileSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VectorStoreIds) {
		toSerialize["vector_store_ids"] = o.VectorStoreIds
	}
	return toSerialize, nil
}

type NullableModifyThreadRequestToolResourcesFileSearch struct {
	value *ModifyThreadRequestToolResourcesFileSearch
	isSet bool
}

func (v NullableModifyThreadRequestToolResourcesFileSearch) Get() *ModifyThreadRequestToolResourcesFileSearch {
	return v.value
}

func (v *NullableModifyThreadRequestToolResourcesFileSearch) Set(val *ModifyThreadRequestToolResourcesFileSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyThreadRequestToolResourcesFileSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyThreadRequestToolResourcesFileSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyThreadRequestToolResourcesFileSearch(val *ModifyThreadRequestToolResourcesFileSearch) *NullableModifyThreadRequestToolResourcesFileSearch {
	return &NullableModifyThreadRequestToolResourcesFileSearch{value: val, isSet: true}
}

func (v NullableModifyThreadRequestToolResourcesFileSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyThreadRequestToolResourcesFileSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

