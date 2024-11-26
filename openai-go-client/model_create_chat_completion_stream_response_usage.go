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

// checks if the CreateChatCompletionStreamResponseUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatCompletionStreamResponseUsage{}

// CreateChatCompletionStreamResponseUsage An optional field that will only be present when you set `stream_options: {\"include_usage\": true}` in your request. When present, it contains a null value except for the last chunk which contains the token usage statistics for the entire request. 
type CreateChatCompletionStreamResponseUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int32 `json:"completion_tokens"`
	// Number of tokens in the prompt.
	PromptTokens int32 `json:"prompt_tokens"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int32 `json:"total_tokens"`
}

type _CreateChatCompletionStreamResponseUsage CreateChatCompletionStreamResponseUsage

// NewCreateChatCompletionStreamResponseUsage instantiates a new CreateChatCompletionStreamResponseUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatCompletionStreamResponseUsage(completionTokens int32, promptTokens int32, totalTokens int32) *CreateChatCompletionStreamResponseUsage {
	this := CreateChatCompletionStreamResponseUsage{}
	this.CompletionTokens = completionTokens
	this.PromptTokens = promptTokens
	this.TotalTokens = totalTokens
	return &this
}

// NewCreateChatCompletionStreamResponseUsageWithDefaults instantiates a new CreateChatCompletionStreamResponseUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatCompletionStreamResponseUsageWithDefaults() *CreateChatCompletionStreamResponseUsage {
	this := CreateChatCompletionStreamResponseUsage{}
	return &this
}

// GetCompletionTokens returns the CompletionTokens field value
func (o *CreateChatCompletionStreamResponseUsage) GetCompletionTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompletionTokens
}

// GetCompletionTokensOk returns a tuple with the CompletionTokens field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionStreamResponseUsage) GetCompletionTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionTokens, true
}

// SetCompletionTokens sets field value
func (o *CreateChatCompletionStreamResponseUsage) SetCompletionTokens(v int32) {
	o.CompletionTokens = v
}

// GetPromptTokens returns the PromptTokens field value
func (o *CreateChatCompletionStreamResponseUsage) GetPromptTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromptTokens
}

// GetPromptTokensOk returns a tuple with the PromptTokens field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionStreamResponseUsage) GetPromptTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptTokens, true
}

// SetPromptTokens sets field value
func (o *CreateChatCompletionStreamResponseUsage) SetPromptTokens(v int32) {
	o.PromptTokens = v
}

// GetTotalTokens returns the TotalTokens field value
func (o *CreateChatCompletionStreamResponseUsage) GetTotalTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionStreamResponseUsage) GetTotalTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTokens, true
}

// SetTotalTokens sets field value
func (o *CreateChatCompletionStreamResponseUsage) SetTotalTokens(v int32) {
	o.TotalTokens = v
}

func (o CreateChatCompletionStreamResponseUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatCompletionStreamResponseUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completion_tokens"] = o.CompletionTokens
	toSerialize["prompt_tokens"] = o.PromptTokens
	toSerialize["total_tokens"] = o.TotalTokens
	return toSerialize, nil
}

func (o *CreateChatCompletionStreamResponseUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completion_tokens",
		"prompt_tokens",
		"total_tokens",
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

	varCreateChatCompletionStreamResponseUsage := _CreateChatCompletionStreamResponseUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatCompletionStreamResponseUsage)

	if err != nil {
		return err
	}

	*o = CreateChatCompletionStreamResponseUsage(varCreateChatCompletionStreamResponseUsage)

	return err
}

type NullableCreateChatCompletionStreamResponseUsage struct {
	value *CreateChatCompletionStreamResponseUsage
	isSet bool
}

func (v NullableCreateChatCompletionStreamResponseUsage) Get() *CreateChatCompletionStreamResponseUsage {
	return v.value
}

func (v *NullableCreateChatCompletionStreamResponseUsage) Set(val *CreateChatCompletionStreamResponseUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatCompletionStreamResponseUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatCompletionStreamResponseUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatCompletionStreamResponseUsage(val *CreateChatCompletionStreamResponseUsage) *NullableCreateChatCompletionStreamResponseUsage {
	return &NullableCreateChatCompletionStreamResponseUsage{value: val, isSet: true}
}

func (v NullableCreateChatCompletionStreamResponseUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatCompletionStreamResponseUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


