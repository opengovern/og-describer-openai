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

// checks if the ChatCompletionMessageToolCallChunkFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionMessageToolCallChunkFunction{}

// ChatCompletionMessageToolCallChunkFunction struct for ChatCompletionMessageToolCallChunkFunction
type ChatCompletionMessageToolCallChunkFunction struct {
	// The name of the function to call.
	Name *string `json:"name,omitempty"`
	// The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.
	Arguments *string `json:"arguments,omitempty"`
}

// NewChatCompletionMessageToolCallChunkFunction instantiates a new ChatCompletionMessageToolCallChunkFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionMessageToolCallChunkFunction() *ChatCompletionMessageToolCallChunkFunction {
	this := ChatCompletionMessageToolCallChunkFunction{}
	return &this
}

// NewChatCompletionMessageToolCallChunkFunctionWithDefaults instantiates a new ChatCompletionMessageToolCallChunkFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionMessageToolCallChunkFunctionWithDefaults() *ChatCompletionMessageToolCallChunkFunction {
	this := ChatCompletionMessageToolCallChunkFunction{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChatCompletionMessageToolCallChunkFunction) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionMessageToolCallChunkFunction) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChatCompletionMessageToolCallChunkFunction) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChatCompletionMessageToolCallChunkFunction) SetName(v string) {
	o.Name = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ChatCompletionMessageToolCallChunkFunction) GetArguments() string {
	if o == nil || IsNil(o.Arguments) {
		var ret string
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionMessageToolCallChunkFunction) GetArgumentsOk() (*string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ChatCompletionMessageToolCallChunkFunction) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given string and assigns it to the Arguments field.
func (o *ChatCompletionMessageToolCallChunkFunction) SetArguments(v string) {
	o.Arguments = &v
}

func (o ChatCompletionMessageToolCallChunkFunction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionMessageToolCallChunkFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	return toSerialize, nil
}

type NullableChatCompletionMessageToolCallChunkFunction struct {
	value *ChatCompletionMessageToolCallChunkFunction
	isSet bool
}

func (v NullableChatCompletionMessageToolCallChunkFunction) Get() *ChatCompletionMessageToolCallChunkFunction {
	return v.value
}

func (v *NullableChatCompletionMessageToolCallChunkFunction) Set(val *ChatCompletionMessageToolCallChunkFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionMessageToolCallChunkFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionMessageToolCallChunkFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionMessageToolCallChunkFunction(val *ChatCompletionMessageToolCallChunkFunction) *NullableChatCompletionMessageToolCallChunkFunction {
	return &NullableChatCompletionMessageToolCallChunkFunction{value: val, isSet: true}
}

func (v NullableChatCompletionMessageToolCallChunkFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionMessageToolCallChunkFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


