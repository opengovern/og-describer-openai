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

// checks if the ChatCompletionResponseMessageAudio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionResponseMessageAudio{}

// ChatCompletionResponseMessageAudio If the audio output modality is requested, this object contains data about the audio response from the model. [Learn more](/docs/guides/audio). 
type ChatCompletionResponseMessageAudio struct {
	// Unique identifier for this audio response.
	Id string `json:"id"`
	// The Unix timestamp (in seconds) for when this audio response will no longer be accessible on the server for use in multi-turn conversations. 
	ExpiresAt int32 `json:"expires_at"`
	// Base64 encoded audio bytes generated by the model, in the format specified in the request. 
	Data string `json:"data"`
	// Transcript of the audio generated by the model.
	Transcript string `json:"transcript"`
}

type _ChatCompletionResponseMessageAudio ChatCompletionResponseMessageAudio

// NewChatCompletionResponseMessageAudio instantiates a new ChatCompletionResponseMessageAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionResponseMessageAudio(id string, expiresAt int32, data string, transcript string) *ChatCompletionResponseMessageAudio {
	this := ChatCompletionResponseMessageAudio{}
	this.Id = id
	this.ExpiresAt = expiresAt
	this.Data = data
	this.Transcript = transcript
	return &this
}

// NewChatCompletionResponseMessageAudioWithDefaults instantiates a new ChatCompletionResponseMessageAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionResponseMessageAudioWithDefaults() *ChatCompletionResponseMessageAudio {
	this := ChatCompletionResponseMessageAudio{}
	return &this
}

// GetId returns the Id field value
func (o *ChatCompletionResponseMessageAudio) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionResponseMessageAudio) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChatCompletionResponseMessageAudio) SetId(v string) {
	o.Id = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *ChatCompletionResponseMessageAudio) GetExpiresAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionResponseMessageAudio) GetExpiresAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *ChatCompletionResponseMessageAudio) SetExpiresAt(v int32) {
	o.ExpiresAt = v
}

// GetData returns the Data field value
func (o *ChatCompletionResponseMessageAudio) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionResponseMessageAudio) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ChatCompletionResponseMessageAudio) SetData(v string) {
	o.Data = v
}

// GetTranscript returns the Transcript field value
func (o *ChatCompletionResponseMessageAudio) GetTranscript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transcript
}

// GetTranscriptOk returns a tuple with the Transcript field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionResponseMessageAudio) GetTranscriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transcript, true
}

// SetTranscript sets field value
func (o *ChatCompletionResponseMessageAudio) SetTranscript(v string) {
	o.Transcript = v
}

func (o ChatCompletionResponseMessageAudio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionResponseMessageAudio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["data"] = o.Data
	toSerialize["transcript"] = o.Transcript
	return toSerialize, nil
}

func (o *ChatCompletionResponseMessageAudio) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"expires_at",
		"data",
		"transcript",
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

	varChatCompletionResponseMessageAudio := _ChatCompletionResponseMessageAudio{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionResponseMessageAudio)

	if err != nil {
		return err
	}

	*o = ChatCompletionResponseMessageAudio(varChatCompletionResponseMessageAudio)

	return err
}

type NullableChatCompletionResponseMessageAudio struct {
	value *ChatCompletionResponseMessageAudio
	isSet bool
}

func (v NullableChatCompletionResponseMessageAudio) Get() *ChatCompletionResponseMessageAudio {
	return v.value
}

func (v *NullableChatCompletionResponseMessageAudio) Set(val *ChatCompletionResponseMessageAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionResponseMessageAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionResponseMessageAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionResponseMessageAudio(val *ChatCompletionResponseMessageAudio) *NullableChatCompletionResponseMessageAudio {
	return &NullableChatCompletionResponseMessageAudio{value: val, isSet: true}
}

func (v NullableChatCompletionResponseMessageAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionResponseMessageAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

