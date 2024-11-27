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

// checks if the ChatCompletionRequestMessageContentPartAudioInputAudio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionRequestMessageContentPartAudioInputAudio{}

// ChatCompletionRequestMessageContentPartAudioInputAudio struct for ChatCompletionRequestMessageContentPartAudioInputAudio
type ChatCompletionRequestMessageContentPartAudioInputAudio struct {
	// Base64 encoded audio data.
	Data string `json:"data"`
	// The format of the encoded audio data. Currently supports \"wav\" and \"mp3\". 
	Format string `json:"format"`
}

type _ChatCompletionRequestMessageContentPartAudioInputAudio ChatCompletionRequestMessageContentPartAudioInputAudio

// NewChatCompletionRequestMessageContentPartAudioInputAudio instantiates a new ChatCompletionRequestMessageContentPartAudioInputAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionRequestMessageContentPartAudioInputAudio(data string, format string) *ChatCompletionRequestMessageContentPartAudioInputAudio {
	this := ChatCompletionRequestMessageContentPartAudioInputAudio{}
	this.Data = data
	this.Format = format
	return &this
}

// NewChatCompletionRequestMessageContentPartAudioInputAudioWithDefaults instantiates a new ChatCompletionRequestMessageContentPartAudioInputAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionRequestMessageContentPartAudioInputAudioWithDefaults() *ChatCompletionRequestMessageContentPartAudioInputAudio {
	this := ChatCompletionRequestMessageContentPartAudioInputAudio{}
	return &this
}

// GetData returns the Data field value
func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) SetData(v string) {
	o.Data = v
}

// GetFormat returns the Format field value
func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) SetFormat(v string) {
	o.Format = v
}

func (o ChatCompletionRequestMessageContentPartAudioInputAudio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionRequestMessageContentPartAudioInputAudio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["format"] = o.Format
	return toSerialize, nil
}

func (o *ChatCompletionRequestMessageContentPartAudioInputAudio) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"format",
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

	varChatCompletionRequestMessageContentPartAudioInputAudio := _ChatCompletionRequestMessageContentPartAudioInputAudio{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionRequestMessageContentPartAudioInputAudio)

	if err != nil {
		return err
	}

	*o = ChatCompletionRequestMessageContentPartAudioInputAudio(varChatCompletionRequestMessageContentPartAudioInputAudio)

	return err
}

type NullableChatCompletionRequestMessageContentPartAudioInputAudio struct {
	value *ChatCompletionRequestMessageContentPartAudioInputAudio
	isSet bool
}

func (v NullableChatCompletionRequestMessageContentPartAudioInputAudio) Get() *ChatCompletionRequestMessageContentPartAudioInputAudio {
	return v.value
}

func (v *NullableChatCompletionRequestMessageContentPartAudioInputAudio) Set(val *ChatCompletionRequestMessageContentPartAudioInputAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestMessageContentPartAudioInputAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestMessageContentPartAudioInputAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestMessageContentPartAudioInputAudio(val *ChatCompletionRequestMessageContentPartAudioInputAudio) *NullableChatCompletionRequestMessageContentPartAudioInputAudio {
	return &NullableChatCompletionRequestMessageContentPartAudioInputAudio{value: val, isSet: true}
}

func (v NullableChatCompletionRequestMessageContentPartAudioInputAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestMessageContentPartAudioInputAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

