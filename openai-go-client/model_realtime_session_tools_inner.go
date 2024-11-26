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

// checks if the RealtimeSessionToolsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeSessionToolsInner{}

// RealtimeSessionToolsInner struct for RealtimeSessionToolsInner
type RealtimeSessionToolsInner struct {
	// The type of the tool, i.e. `function`.
	Type *string `json:"type,omitempty"`
	// The name of the function.
	Name *string `json:"name,omitempty"`
	// The description of the function, including guidance on when and how  to call it, and guidance about what to tell the user when calling  (if anything). 
	Description *string `json:"description,omitempty"`
	// Parameters of the function in JSON Schema.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// NewRealtimeSessionToolsInner instantiates a new RealtimeSessionToolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeSessionToolsInner() *RealtimeSessionToolsInner {
	this := RealtimeSessionToolsInner{}
	return &this
}

// NewRealtimeSessionToolsInnerWithDefaults instantiates a new RealtimeSessionToolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeSessionToolsInnerWithDefaults() *RealtimeSessionToolsInner {
	this := RealtimeSessionToolsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RealtimeSessionToolsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionToolsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RealtimeSessionToolsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RealtimeSessionToolsInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RealtimeSessionToolsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionToolsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RealtimeSessionToolsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RealtimeSessionToolsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RealtimeSessionToolsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionToolsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RealtimeSessionToolsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RealtimeSessionToolsInner) SetDescription(v string) {
	o.Description = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *RealtimeSessionToolsInner) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionToolsInner) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *RealtimeSessionToolsInner) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *RealtimeSessionToolsInner) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

func (o RealtimeSessionToolsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeSessionToolsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableRealtimeSessionToolsInner struct {
	value *RealtimeSessionToolsInner
	isSet bool
}

func (v NullableRealtimeSessionToolsInner) Get() *RealtimeSessionToolsInner {
	return v.value
}

func (v *NullableRealtimeSessionToolsInner) Set(val *RealtimeSessionToolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeSessionToolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeSessionToolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeSessionToolsInner(val *RealtimeSessionToolsInner) *NullableRealtimeSessionToolsInner {
	return &NullableRealtimeSessionToolsInner{value: val, isSet: true}
}

func (v NullableRealtimeSessionToolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeSessionToolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


