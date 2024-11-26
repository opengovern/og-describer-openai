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

// checks if the CreateAssistantRequestToolResourcesFileSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssistantRequestToolResourcesFileSearch{}

// CreateAssistantRequestToolResourcesFileSearch struct for CreateAssistantRequestToolResourcesFileSearch
type CreateAssistantRequestToolResourcesFileSearch struct {
	// The [vector store](/docs/api-reference/vector-stores/object) attached to this assistant. There can be a maximum of 1 vector store attached to the assistant. 
	VectorStoreIds []string
	// A helper to create a [vector store](/docs/api-reference/vector-stores/object) with file_ids and attach it to this assistant. There can be a maximum of 1 vector store attached to the assistant. 
	VectorStores []CreateAssistantRequestToolResourcesFileSearchVectorStoresInner
}

// NewCreateAssistantRequestToolResourcesFileSearch instantiates a new CreateAssistantRequestToolResourcesFileSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssistantRequestToolResourcesFileSearch() *CreateAssistantRequestToolResourcesFileSearch {
	this := CreateAssistantRequestToolResourcesFileSearch{}
	return &this
}

// NewCreateAssistantRequestToolResourcesFileSearchWithDefaults instantiates a new CreateAssistantRequestToolResourcesFileSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssistantRequestToolResourcesFileSearchWithDefaults() *CreateAssistantRequestToolResourcesFileSearch {
	this := CreateAssistantRequestToolResourcesFileSearch{}
	return &this
}

// GetVectorStoreIds returns the VectorStoreIds field value if set, zero value otherwise.
func (o *CreateAssistantRequestToolResourcesFileSearch) GetVectorStoreIds() []string {
	if o == nil || IsNil(o.VectorStoreIds) {
		var ret []string
		return ret
	}
	return o.VectorStoreIds
}

// GetVectorStoreIdsOk returns a tuple with the VectorStoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssistantRequestToolResourcesFileSearch) GetVectorStoreIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VectorStoreIds) {
		return nil, false
	}
	return o.VectorStoreIds, true
}

// HasVectorStoreIds returns a boolean if a field has been set.
func (o *CreateAssistantRequestToolResourcesFileSearch) HasVectorStoreIds() bool {
	if o != nil && !IsNil(o.VectorStoreIds) {
		return true
	}

	return false
}

// SetVectorStoreIds gets a reference to the given []string and assigns it to the VectorStoreIds field.
func (o *CreateAssistantRequestToolResourcesFileSearch) SetVectorStoreIds(v []string) {
	o.VectorStoreIds = v
}

// GetVectorStores returns the VectorStores field value if set, zero value otherwise.
func (o *CreateAssistantRequestToolResourcesFileSearch) GetVectorStores() []CreateAssistantRequestToolResourcesFileSearchVectorStoresInner {
	if o == nil || IsNil(o.VectorStores) {
		var ret []CreateAssistantRequestToolResourcesFileSearchVectorStoresInner
		return ret
	}
	return o.VectorStores
}

// GetVectorStoresOk returns a tuple with the VectorStores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssistantRequestToolResourcesFileSearch) GetVectorStoresOk() ([]CreateAssistantRequestToolResourcesFileSearchVectorStoresInner, bool) {
	if o == nil || IsNil(o.VectorStores) {
		return nil, false
	}
	return o.VectorStores, true
}

// HasVectorStores returns a boolean if a field has been set.
func (o *CreateAssistantRequestToolResourcesFileSearch) HasVectorStores() bool {
	if o != nil && !IsNil(o.VectorStores) {
		return true
	}

	return false
}

// SetVectorStores gets a reference to the given []CreateAssistantRequestToolResourcesFileSearchVectorStoresInner and assigns it to the VectorStores field.
func (o *CreateAssistantRequestToolResourcesFileSearch) SetVectorStores(v []CreateAssistantRequestToolResourcesFileSearchVectorStoresInner) {
	o.VectorStores = v
}

func (o CreateAssistantRequestToolResourcesFileSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssistantRequestToolResourcesFileSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VectorStoreIds) {
		toSerialize["vector_store_ids"] = o.VectorStoreIds
	}
	if !IsNil(o.VectorStores) {
		toSerialize["vector_stores"] = o.VectorStores
	}
	return toSerialize, nil
}

type NullableCreateAssistantRequestToolResourcesFileSearch struct {
	value *CreateAssistantRequestToolResourcesFileSearch
	isSet bool
}

func (v NullableCreateAssistantRequestToolResourcesFileSearch) Get() *CreateAssistantRequestToolResourcesFileSearch {
	return v.value
}

func (v *NullableCreateAssistantRequestToolResourcesFileSearch) Set(val *CreateAssistantRequestToolResourcesFileSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssistantRequestToolResourcesFileSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssistantRequestToolResourcesFileSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssistantRequestToolResourcesFileSearch(val *CreateAssistantRequestToolResourcesFileSearch) *NullableCreateAssistantRequestToolResourcesFileSearch {
	return &NullableCreateAssistantRequestToolResourcesFileSearch{value: val, isSet: true}
}

func (v NullableCreateAssistantRequestToolResourcesFileSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssistantRequestToolResourcesFileSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


