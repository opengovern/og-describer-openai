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

// checks if the ChatCompletionRequestAssistantMessageFunctionCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionRequestAssistantMessageFunctionCall{}

// ChatCompletionRequestAssistantMessageFunctionCall Deprecated and replaced by `tool_calls`. The name and arguments of a function that should be called, as generated by the model.
type ChatCompletionRequestAssistantMessageFunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.
	Arguments string `json:"arguments"`
	// The name of the function to call.
	Name string `json:"name"`
}

type _ChatCompletionRequestAssistantMessageFunctionCall ChatCompletionRequestAssistantMessageFunctionCall

// NewChatCompletionRequestAssistantMessageFunctionCall instantiates a new ChatCompletionRequestAssistantMessageFunctionCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionRequestAssistantMessageFunctionCall(arguments string, name string) *ChatCompletionRequestAssistantMessageFunctionCall {
	this := ChatCompletionRequestAssistantMessageFunctionCall{}
	this.Arguments = arguments
	this.Name = name
	return &this
}

// NewChatCompletionRequestAssistantMessageFunctionCallWithDefaults instantiates a new ChatCompletionRequestAssistantMessageFunctionCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionRequestAssistantMessageFunctionCallWithDefaults() *ChatCompletionRequestAssistantMessageFunctionCall {
	this := ChatCompletionRequestAssistantMessageFunctionCall{}
	return &this
}

// GetArguments returns the Arguments field value
func (o *ChatCompletionRequestAssistantMessageFunctionCall) GetArguments() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestAssistantMessageFunctionCall) GetArgumentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value
func (o *ChatCompletionRequestAssistantMessageFunctionCall) SetArguments(v string) {
	o.Arguments = v
}

// GetName returns the Name field value
func (o *ChatCompletionRequestAssistantMessageFunctionCall) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestAssistantMessageFunctionCall) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChatCompletionRequestAssistantMessageFunctionCall) SetName(v string) {
	o.Name = v
}

func (o ChatCompletionRequestAssistantMessageFunctionCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionRequestAssistantMessageFunctionCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["arguments"] = o.Arguments
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ChatCompletionRequestAssistantMessageFunctionCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"arguments",
		"name",
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

	varChatCompletionRequestAssistantMessageFunctionCall := _ChatCompletionRequestAssistantMessageFunctionCall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionRequestAssistantMessageFunctionCall)

	if err != nil {
		return err
	}

	*o = ChatCompletionRequestAssistantMessageFunctionCall(varChatCompletionRequestAssistantMessageFunctionCall)

	return err
}

type NullableChatCompletionRequestAssistantMessageFunctionCall struct {
	value *ChatCompletionRequestAssistantMessageFunctionCall
	isSet bool
}

func (v NullableChatCompletionRequestAssistantMessageFunctionCall) Get() *ChatCompletionRequestAssistantMessageFunctionCall {
	return v.value
}

func (v *NullableChatCompletionRequestAssistantMessageFunctionCall) Set(val *ChatCompletionRequestAssistantMessageFunctionCall) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestAssistantMessageFunctionCall) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestAssistantMessageFunctionCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestAssistantMessageFunctionCall(val *ChatCompletionRequestAssistantMessageFunctionCall) *NullableChatCompletionRequestAssistantMessageFunctionCall {
	return &NullableChatCompletionRequestAssistantMessageFunctionCall{value: val, isSet: true}
}

func (v NullableChatCompletionRequestAssistantMessageFunctionCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestAssistantMessageFunctionCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


