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

// checks if the RealtimeServerEventResponseContentPartAddedPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventResponseContentPartAddedPart{}

// RealtimeServerEventResponseContentPartAddedPart The content part that was added.
type RealtimeServerEventResponseContentPartAddedPart struct {
	// The content type (\"text\", \"audio\").
	Type *string `json:"type,omitempty"`
	// The text content (if type is \"text\").
	Text *string `json:"text,omitempty"`
	// Base64-encoded audio data (if type is \"audio\").
	Audio *string `json:"audio,omitempty"`
	// The transcript of the audio (if type is \"audio\").
	Transcript *string `json:"transcript,omitempty"`
}

// NewRealtimeServerEventResponseContentPartAddedPart instantiates a new RealtimeServerEventResponseContentPartAddedPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventResponseContentPartAddedPart() *RealtimeServerEventResponseContentPartAddedPart {
	this := RealtimeServerEventResponseContentPartAddedPart{}
	return &this
}

// NewRealtimeServerEventResponseContentPartAddedPartWithDefaults instantiates a new RealtimeServerEventResponseContentPartAddedPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventResponseContentPartAddedPartWithDefaults() *RealtimeServerEventResponseContentPartAddedPart {
	this := RealtimeServerEventResponseContentPartAddedPart{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RealtimeServerEventResponseContentPartAddedPart) SetType(v string) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RealtimeServerEventResponseContentPartAddedPart) SetText(v string) {
	o.Text = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetAudio() string {
	if o == nil || IsNil(o.Audio) {
		var ret string
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetAudioOk() (*string, bool) {
	if o == nil || IsNil(o.Audio) {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) HasAudio() bool {
	if o != nil && !IsNil(o.Audio) {
		return true
	}

	return false
}

// SetAudio gets a reference to the given string and assigns it to the Audio field.
func (o *RealtimeServerEventResponseContentPartAddedPart) SetAudio(v string) {
	o.Audio = &v
}

// GetTranscript returns the Transcript field value if set, zero value otherwise.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetTranscript() string {
	if o == nil || IsNil(o.Transcript) {
		var ret string
		return ret
	}
	return *o.Transcript
}

// GetTranscriptOk returns a tuple with the Transcript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) GetTranscriptOk() (*string, bool) {
	if o == nil || IsNil(o.Transcript) {
		return nil, false
	}
	return o.Transcript, true
}

// HasTranscript returns a boolean if a field has been set.
func (o *RealtimeServerEventResponseContentPartAddedPart) HasTranscript() bool {
	if o != nil && !IsNil(o.Transcript) {
		return true
	}

	return false
}

// SetTranscript gets a reference to the given string and assigns it to the Transcript field.
func (o *RealtimeServerEventResponseContentPartAddedPart) SetTranscript(v string) {
	o.Transcript = &v
}

func (o RealtimeServerEventResponseContentPartAddedPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventResponseContentPartAddedPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Audio) {
		toSerialize["audio"] = o.Audio
	}
	if !IsNil(o.Transcript) {
		toSerialize["transcript"] = o.Transcript
	}
	return toSerialize, nil
}

type NullableRealtimeServerEventResponseContentPartAddedPart struct {
	value *RealtimeServerEventResponseContentPartAddedPart
	isSet bool
}

func (v NullableRealtimeServerEventResponseContentPartAddedPart) Get() *RealtimeServerEventResponseContentPartAddedPart {
	return v.value
}

func (v *NullableRealtimeServerEventResponseContentPartAddedPart) Set(val *RealtimeServerEventResponseContentPartAddedPart) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventResponseContentPartAddedPart) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventResponseContentPartAddedPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventResponseContentPartAddedPart(val *RealtimeServerEventResponseContentPartAddedPart) *NullableRealtimeServerEventResponseContentPartAddedPart {
	return &NullableRealtimeServerEventResponseContentPartAddedPart{value: val, isSet: true}
}

func (v NullableRealtimeServerEventResponseContentPartAddedPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventResponseContentPartAddedPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

