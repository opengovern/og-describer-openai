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

// checks if the ChatCompletionTokenLogprobTopLogprobsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionTokenLogprobTopLogprobsInner{}

// ChatCompletionTokenLogprobTopLogprobsInner struct for ChatCompletionTokenLogprobTopLogprobsInner
type ChatCompletionTokenLogprobTopLogprobsInner struct {
	// The token.
	Token string `json:"token"`
	// The log probability of this token, if it is within the top 20 most likely tokens. Otherwise, the value `-9999.0` is used to signify that the token is very unlikely.
	Logprob float32 `json:"logprob"`
	// A list of integers representing the UTF-8 bytes representation of the token. Useful in instances where characters are represented by multiple tokens and their byte representations must be combined to generate the correct text representation. Can be `null` if there is no bytes representation for the token.
	Bytes []int32 `json:"bytes"`
}

type _ChatCompletionTokenLogprobTopLogprobsInner ChatCompletionTokenLogprobTopLogprobsInner

// NewChatCompletionTokenLogprobTopLogprobsInner instantiates a new ChatCompletionTokenLogprobTopLogprobsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionTokenLogprobTopLogprobsInner(token string, logprob float32, bytes []int32) *ChatCompletionTokenLogprobTopLogprobsInner {
	this := ChatCompletionTokenLogprobTopLogprobsInner{}
	this.Token = token
	this.Logprob = logprob
	this.Bytes = bytes
	return &this
}

// NewChatCompletionTokenLogprobTopLogprobsInnerWithDefaults instantiates a new ChatCompletionTokenLogprobTopLogprobsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionTokenLogprobTopLogprobsInnerWithDefaults() *ChatCompletionTokenLogprobTopLogprobsInner {
	this := ChatCompletionTokenLogprobTopLogprobsInner{}
	return &this
}

// GetToken returns the Token field value
func (o *ChatCompletionTokenLogprobTopLogprobsInner) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionTokenLogprobTopLogprobsInner) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ChatCompletionTokenLogprobTopLogprobsInner) SetToken(v string) {
	o.Token = v
}

// GetLogprob returns the Logprob field value
func (o *ChatCompletionTokenLogprobTopLogprobsInner) GetLogprob() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Logprob
}

// GetLogprobOk returns a tuple with the Logprob field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionTokenLogprobTopLogprobsInner) GetLogprobOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logprob, true
}

// SetLogprob sets field value
func (o *ChatCompletionTokenLogprobTopLogprobsInner) SetLogprob(v float32) {
	o.Logprob = v
}

// GetBytes returns the Bytes field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *ChatCompletionTokenLogprobTopLogprobsInner) GetBytes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionTokenLogprobTopLogprobsInner) GetBytesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Bytes) {
		return nil, false
	}
	return o.Bytes, true
}

// SetBytes sets field value
func (o *ChatCompletionTokenLogprobTopLogprobsInner) SetBytes(v []int32) {
	o.Bytes = v
}

func (o ChatCompletionTokenLogprobTopLogprobsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionTokenLogprobTopLogprobsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["logprob"] = o.Logprob
	if o.Bytes != nil {
		toSerialize["bytes"] = o.Bytes
	}
	return toSerialize, nil
}

func (o *ChatCompletionTokenLogprobTopLogprobsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"logprob",
		"bytes",
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

	varChatCompletionTokenLogprobTopLogprobsInner := _ChatCompletionTokenLogprobTopLogprobsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionTokenLogprobTopLogprobsInner)

	if err != nil {
		return err
	}

	*o = ChatCompletionTokenLogprobTopLogprobsInner(varChatCompletionTokenLogprobTopLogprobsInner)

	return err
}

type NullableChatCompletionTokenLogprobTopLogprobsInner struct {
	value *ChatCompletionTokenLogprobTopLogprobsInner
	isSet bool
}

func (v NullableChatCompletionTokenLogprobTopLogprobsInner) Get() *ChatCompletionTokenLogprobTopLogprobsInner {
	return v.value
}

func (v *NullableChatCompletionTokenLogprobTopLogprobsInner) Set(val *ChatCompletionTokenLogprobTopLogprobsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionTokenLogprobTopLogprobsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionTokenLogprobTopLogprobsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionTokenLogprobTopLogprobsInner(val *ChatCompletionTokenLogprobTopLogprobsInner) *NullableChatCompletionTokenLogprobTopLogprobsInner {
	return &NullableChatCompletionTokenLogprobTopLogprobsInner{value: val, isSet: true}
}

func (v NullableChatCompletionTokenLogprobTopLogprobsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionTokenLogprobTopLogprobsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

