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

// checks if the CreateAssistantRequestToolResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssistantRequestToolResources{}

// CreateAssistantRequestToolResources A set of resources that are used by the assistant's tools. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs. 
type CreateAssistantRequestToolResources struct {
	CodeInterpreter *CreateAssistantRequestToolResourcesCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch NullableCreateAssistantRequestToolResourcesFileSearch `json:"file_search,omitempty"`
}

// NewCreateAssistantRequestToolResources instantiates a new CreateAssistantRequestToolResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssistantRequestToolResources() *CreateAssistantRequestToolResources {
	this := CreateAssistantRequestToolResources{}
	return &this
}

// NewCreateAssistantRequestToolResourcesWithDefaults instantiates a new CreateAssistantRequestToolResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssistantRequestToolResourcesWithDefaults() *CreateAssistantRequestToolResources {
	this := CreateAssistantRequestToolResources{}
	return &this
}

// GetCodeInterpreter returns the CodeInterpreter field value if set, zero value otherwise.
func (o *CreateAssistantRequestToolResources) GetCodeInterpreter() CreateAssistantRequestToolResourcesCodeInterpreter {
	if o == nil || IsNil(o.CodeInterpreter) {
		var ret CreateAssistantRequestToolResourcesCodeInterpreter
		return ret
	}
	return *o.CodeInterpreter
}

// GetCodeInterpreterOk returns a tuple with the CodeInterpreter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssistantRequestToolResources) GetCodeInterpreterOk() (*CreateAssistantRequestToolResourcesCodeInterpreter, bool) {
	if o == nil || IsNil(o.CodeInterpreter) {
		return nil, false
	}
	return o.CodeInterpreter, true
}

// HasCodeInterpreter returns a boolean if a field has been set.
func (o *CreateAssistantRequestToolResources) HasCodeInterpreter() bool {
	if o != nil && !IsNil(o.CodeInterpreter) {
		return true
	}

	return false
}

// SetCodeInterpreter gets a reference to the given CreateAssistantRequestToolResourcesCodeInterpreter and assigns it to the CodeInterpreter field.
func (o *CreateAssistantRequestToolResources) SetCodeInterpreter(v CreateAssistantRequestToolResourcesCodeInterpreter) {
	o.CodeInterpreter = &v
}

// GetFileSearch returns the FileSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAssistantRequestToolResources) GetFileSearch() CreateAssistantRequestToolResourcesFileSearch {
	if o == nil || IsNil(o.FileSearch.Get()) {
		var ret CreateAssistantRequestToolResourcesFileSearch
		return ret
	}
	return *o.FileSearch.Get()
}

// GetFileSearchOk returns a tuple with the FileSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAssistantRequestToolResources) GetFileSearchOk() (*CreateAssistantRequestToolResourcesFileSearch, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSearch.Get(), o.FileSearch.IsSet()
}

// HasFileSearch returns a boolean if a field has been set.
func (o *CreateAssistantRequestToolResources) HasFileSearch() bool {
	if o != nil && o.FileSearch.IsSet() {
		return true
	}

	return false
}

// SetFileSearch gets a reference to the given NullableCreateAssistantRequestToolResourcesFileSearch and assigns it to the FileSearch field.
func (o *CreateAssistantRequestToolResources) SetFileSearch(v CreateAssistantRequestToolResourcesFileSearch) {
	o.FileSearch.Set(&v)
}
// SetFileSearchNil sets the value for FileSearch to be an explicit nil
func (o *CreateAssistantRequestToolResources) SetFileSearchNil() {
	o.FileSearch.Set(nil)
}

// UnsetFileSearch ensures that no value is present for FileSearch, not even an explicit nil
func (o *CreateAssistantRequestToolResources) UnsetFileSearch() {
	o.FileSearch.Unset()
}

func (o CreateAssistantRequestToolResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssistantRequestToolResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeInterpreter) {
		toSerialize["code_interpreter"] = o.CodeInterpreter
	}
	if o.FileSearch.IsSet() {
		toSerialize["file_search"] = o.FileSearch.Get()
	}
	return toSerialize, nil
}

type NullableCreateAssistantRequestToolResources struct {
	value *CreateAssistantRequestToolResources
	isSet bool
}

func (v NullableCreateAssistantRequestToolResources) Get() *CreateAssistantRequestToolResources {
	return v.value
}

func (v *NullableCreateAssistantRequestToolResources) Set(val *CreateAssistantRequestToolResources) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssistantRequestToolResources) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssistantRequestToolResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssistantRequestToolResources(val *CreateAssistantRequestToolResources) *NullableCreateAssistantRequestToolResources {
	return &NullableCreateAssistantRequestToolResources{value: val, isSet: true}
}

func (v NullableCreateAssistantRequestToolResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssistantRequestToolResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


