/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AssistantToolsFileSearchTypeOnly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantToolsFileSearchTypeOnly{}

// AssistantToolsFileSearchTypeOnly struct for AssistantToolsFileSearchTypeOnly
type AssistantToolsFileSearchTypeOnly struct {
	// The type of tool being defined: `file_search`
	Type string `json:"type"`
}

type _AssistantToolsFileSearchTypeOnly AssistantToolsFileSearchTypeOnly

// NewAssistantToolsFileSearchTypeOnly instantiates a new AssistantToolsFileSearchTypeOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantToolsFileSearchTypeOnly(type_ string) *AssistantToolsFileSearchTypeOnly {
	this := AssistantToolsFileSearchTypeOnly{}
	this.Type = type_
	return &this
}

// NewAssistantToolsFileSearchTypeOnlyWithDefaults instantiates a new AssistantToolsFileSearchTypeOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantToolsFileSearchTypeOnlyWithDefaults() *AssistantToolsFileSearchTypeOnly {
	this := AssistantToolsFileSearchTypeOnly{}
	return &this
}

// GetType returns the Type field value
func (o *AssistantToolsFileSearchTypeOnly) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssistantToolsFileSearchTypeOnly) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssistantToolsFileSearchTypeOnly) SetType(v string) {
	o.Type = v
}

func (o AssistantToolsFileSearchTypeOnly) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantToolsFileSearchTypeOnly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AssistantToolsFileSearchTypeOnly) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAssistantToolsFileSearchTypeOnly := _AssistantToolsFileSearchTypeOnly{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssistantToolsFileSearchTypeOnly)

	if err != nil {
		return err
	}

	*o = AssistantToolsFileSearchTypeOnly(varAssistantToolsFileSearchTypeOnly)

	return err
}

type NullableAssistantToolsFileSearchTypeOnly struct {
	value *AssistantToolsFileSearchTypeOnly
	isSet bool
}

func (v NullableAssistantToolsFileSearchTypeOnly) Get() *AssistantToolsFileSearchTypeOnly {
	return v.value
}

func (v *NullableAssistantToolsFileSearchTypeOnly) Set(val *AssistantToolsFileSearchTypeOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantToolsFileSearchTypeOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantToolsFileSearchTypeOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantToolsFileSearchTypeOnly(val *AssistantToolsFileSearchTypeOnly) *NullableAssistantToolsFileSearchTypeOnly {
	return &NullableAssistantToolsFileSearchTypeOnly{value: val, isSet: true}
}

func (v NullableAssistantToolsFileSearchTypeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantToolsFileSearchTypeOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

