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

// checks if the ChatCompletionRequestAssistantMessageAudio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionRequestAssistantMessageAudio{}

// ChatCompletionRequestAssistantMessageAudio Data about a previous audio response from the model.  [Learn more](/docs/guides/audio). 
type ChatCompletionRequestAssistantMessageAudio struct {
	// Unique identifier for a previous audio response from the model. 
	Id string `json:"id"`
}

type _ChatCompletionRequestAssistantMessageAudio ChatCompletionRequestAssistantMessageAudio

// NewChatCompletionRequestAssistantMessageAudio instantiates a new ChatCompletionRequestAssistantMessageAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionRequestAssistantMessageAudio(id string) *ChatCompletionRequestAssistantMessageAudio {
	this := ChatCompletionRequestAssistantMessageAudio{}
	this.Id = id
	return &this
}

// NewChatCompletionRequestAssistantMessageAudioWithDefaults instantiates a new ChatCompletionRequestAssistantMessageAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionRequestAssistantMessageAudioWithDefaults() *ChatCompletionRequestAssistantMessageAudio {
	this := ChatCompletionRequestAssistantMessageAudio{}
	return &this
}

// GetId returns the Id field value
func (o *ChatCompletionRequestAssistantMessageAudio) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequestAssistantMessageAudio) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChatCompletionRequestAssistantMessageAudio) SetId(v string) {
	o.Id = v
}

func (o ChatCompletionRequestAssistantMessageAudio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionRequestAssistantMessageAudio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ChatCompletionRequestAssistantMessageAudio) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varChatCompletionRequestAssistantMessageAudio := _ChatCompletionRequestAssistantMessageAudio{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionRequestAssistantMessageAudio)

	if err != nil {
		return err
	}

	*o = ChatCompletionRequestAssistantMessageAudio(varChatCompletionRequestAssistantMessageAudio)

	return err
}

type NullableChatCompletionRequestAssistantMessageAudio struct {
	value *ChatCompletionRequestAssistantMessageAudio
	isSet bool
}

func (v NullableChatCompletionRequestAssistantMessageAudio) Get() *ChatCompletionRequestAssistantMessageAudio {
	return v.value
}

func (v *NullableChatCompletionRequestAssistantMessageAudio) Set(val *ChatCompletionRequestAssistantMessageAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestAssistantMessageAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestAssistantMessageAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestAssistantMessageAudio(val *ChatCompletionRequestAssistantMessageAudio) *NullableChatCompletionRequestAssistantMessageAudio {
	return &NullableChatCompletionRequestAssistantMessageAudio{value: val, isSet: true}
}

func (v NullableChatCompletionRequestAssistantMessageAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestAssistantMessageAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


