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

// checks if the ChatCompletionTool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionTool{}

// ChatCompletionTool struct for ChatCompletionTool
type ChatCompletionTool struct {
	// The type of the tool. Currently, only `function` is supported.
	Type string `json:"type"`
	Function FunctionObject `json:"function"`
}

type _ChatCompletionTool ChatCompletionTool

// NewChatCompletionTool instantiates a new ChatCompletionTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionTool(type_ string, function FunctionObject) *ChatCompletionTool {
	this := ChatCompletionTool{}
	this.Type = type_
	this.Function = function
	return &this
}

// NewChatCompletionToolWithDefaults instantiates a new ChatCompletionTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionToolWithDefaults() *ChatCompletionTool {
	this := ChatCompletionTool{}
	return &this
}

// GetType returns the Type field value
func (o *ChatCompletionTool) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionTool) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChatCompletionTool) SetType(v string) {
	o.Type = v
}

// GetFunction returns the Function field value
func (o *ChatCompletionTool) GetFunction() FunctionObject {
	if o == nil {
		var ret FunctionObject
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionTool) GetFunctionOk() (*FunctionObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *ChatCompletionTool) SetFunction(v FunctionObject) {
	o.Function = v
}

func (o ChatCompletionTool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionTool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["function"] = o.Function
	return toSerialize, nil
}

func (o *ChatCompletionTool) UnmarshalJSON(data []byte) (err error) {
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

	varChatCompletionTool := _ChatCompletionTool{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionTool)

	if err != nil {
		return err
	}

	*o = ChatCompletionTool(varChatCompletionTool)

	return err
}

type NullableChatCompletionTool struct {
	value *ChatCompletionTool
	isSet bool
}

func (v NullableChatCompletionTool) Get() *ChatCompletionTool {
	return v.value
}

func (v *NullableChatCompletionTool) Set(val *ChatCompletionTool) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionTool) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionTool(val *ChatCompletionTool) *NullableChatCompletionTool {
	return &NullableChatCompletionTool{value: val, isSet: true}
}

func (v NullableChatCompletionTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


