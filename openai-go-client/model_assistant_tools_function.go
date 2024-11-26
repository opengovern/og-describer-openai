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

// checks if the AssistantToolsFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantToolsFunction{}

// AssistantToolsFunction struct for AssistantToolsFunction
type AssistantToolsFunction struct {
	// The type of tool being defined: `function`
	Type string `json:"type"`
	Function FunctionObject `json:"function"`
}

type _AssistantToolsFunction AssistantToolsFunction

// NewAssistantToolsFunction instantiates a new AssistantToolsFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantToolsFunction(type_ string, function FunctionObject) *AssistantToolsFunction {
	this := AssistantToolsFunction{}
	this.Type = type_
	this.Function = function
	return &this
}

// NewAssistantToolsFunctionWithDefaults instantiates a new AssistantToolsFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantToolsFunctionWithDefaults() *AssistantToolsFunction {
	this := AssistantToolsFunction{}
	return &this
}

// GetType returns the Type field value
func (o *AssistantToolsFunction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssistantToolsFunction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssistantToolsFunction) SetType(v string) {
	o.Type = v
}

// GetFunction returns the Function field value
func (o *AssistantToolsFunction) GetFunction() FunctionObject {
	if o == nil {
		var ret FunctionObject
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *AssistantToolsFunction) GetFunctionOk() (*FunctionObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *AssistantToolsFunction) SetFunction(v FunctionObject) {
	o.Function = v
}

func (o AssistantToolsFunction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantToolsFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["function"] = o.Function
	return toSerialize, nil
}

func (o *AssistantToolsFunction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"function",
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

	varAssistantToolsFunction := _AssistantToolsFunction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssistantToolsFunction)

	if err != nil {
		return err
	}

	*o = AssistantToolsFunction(varAssistantToolsFunction)

	return err
}

type NullableAssistantToolsFunction struct {
	value *AssistantToolsFunction
	isSet bool
}

func (v NullableAssistantToolsFunction) Get() *AssistantToolsFunction {
	return v.value
}

func (v *NullableAssistantToolsFunction) Set(val *AssistantToolsFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantToolsFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantToolsFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantToolsFunction(val *AssistantToolsFunction) *NullableAssistantToolsFunction {
	return &NullableAssistantToolsFunction{value: val, isSet: true}
}

func (v NullableAssistantToolsFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantToolsFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


