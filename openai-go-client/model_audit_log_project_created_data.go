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

// checks if the AuditLogProjectCreatedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogProjectCreatedData{}

// AuditLogProjectCreatedData The payload used to create the project.
type AuditLogProjectCreatedData struct {
	// The project name.
	Name *string `json:"name,omitempty"`
	// The title of the project as seen on the dashboard.
	Title *string `json:"title,omitempty"`
}

// NewAuditLogProjectCreatedData instantiates a new AuditLogProjectCreatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogProjectCreatedData() *AuditLogProjectCreatedData {
	this := AuditLogProjectCreatedData{}
	return &this
}

// NewAuditLogProjectCreatedDataWithDefaults instantiates a new AuditLogProjectCreatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogProjectCreatedDataWithDefaults() *AuditLogProjectCreatedData {
	this := AuditLogProjectCreatedData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuditLogProjectCreatedData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogProjectCreatedData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuditLogProjectCreatedData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuditLogProjectCreatedData) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AuditLogProjectCreatedData) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogProjectCreatedData) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AuditLogProjectCreatedData) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AuditLogProjectCreatedData) SetTitle(v string) {
	o.Title = &v
}

func (o AuditLogProjectCreatedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogProjectCreatedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAuditLogProjectCreatedData struct {
	value *AuditLogProjectCreatedData
	isSet bool
}

func (v NullableAuditLogProjectCreatedData) Get() *AuditLogProjectCreatedData {
	return v.value
}

func (v *NullableAuditLogProjectCreatedData) Set(val *AuditLogProjectCreatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogProjectCreatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogProjectCreatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogProjectCreatedData(val *AuditLogProjectCreatedData) *NullableAuditLogProjectCreatedData {
	return &NullableAuditLogProjectCreatedData{value: val, isSet: true}
}

func (v NullableAuditLogProjectCreatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogProjectCreatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


