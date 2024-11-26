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

// checks if the RealtimeServerEventConversationItemInputAudioTranscriptionCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventConversationItemInputAudioTranscriptionCompleted{}

// RealtimeServerEventConversationItemInputAudioTranscriptionCompleted This event is the output of audio transcription for user audio written to the user audio buffer. Transcription begins when the input audio buffer is committed by the client or server (in `server_vad` mode). Transcription runs asynchronously with Response creation, so this event may come before or after the Response events. Realtime API models accept audio natively, and thus input transcription is a separate process run on a separate ASR (Automatic Speech Recognition) model, currently always `whisper-1`. Thus the transcript may diverge somewhat from the model's interpretation, and should be treated as a rough guide.
type RealtimeServerEventConversationItemInputAudioTranscriptionCompleted struct {
	// The unique ID of the server event.
	EventId string `json:"event_id"`
	// The event type, must be `conversation.item.input_audio_transcription.completed`.
	Type string `json:"type"`
	// The ID of the user message item containing the audio.
	ItemId string `json:"item_id"`
	// The index of the content part containing the audio.
	ContentIndex int32 `json:"content_index"`
	// The transcribed text.
	Transcript string `json:"transcript"`
}

type _RealtimeServerEventConversationItemInputAudioTranscriptionCompleted RealtimeServerEventConversationItemInputAudioTranscriptionCompleted

// NewRealtimeServerEventConversationItemInputAudioTranscriptionCompleted instantiates a new RealtimeServerEventConversationItemInputAudioTranscriptionCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventConversationItemInputAudioTranscriptionCompleted(eventId string, type_ string, itemId string, contentIndex int32, transcript string) *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted {
	this := RealtimeServerEventConversationItemInputAudioTranscriptionCompleted{}
	this.EventId = eventId
	this.Type = type_
	this.ItemId = itemId
	this.ContentIndex = contentIndex
	this.Transcript = transcript
	return &this
}

// NewRealtimeServerEventConversationItemInputAudioTranscriptionCompletedWithDefaults instantiates a new RealtimeServerEventConversationItemInputAudioTranscriptionCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventConversationItemInputAudioTranscriptionCompletedWithDefaults() *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted {
	this := RealtimeServerEventConversationItemInputAudioTranscriptionCompleted{}
	return &this
}

// GetEventId returns the EventId field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) SetType(v string) {
	o.Type = v
}

// GetItemId returns the ItemId field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) SetItemId(v string) {
	o.ItemId = v
}

// GetContentIndex returns the ContentIndex field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetContentIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentIndex
}

// GetContentIndexOk returns a tuple with the ContentIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetContentIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentIndex, true
}

// SetContentIndex sets field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) SetContentIndex(v int32) {
	o.ContentIndex = v
}

// GetTranscript returns the Transcript field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetTranscript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transcript
}

// GetTranscriptOk returns a tuple with the Transcript field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) GetTranscriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transcript, true
}

// SetTranscript sets field value
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) SetTranscript(v string) {
	o.Transcript = v
}

func (o RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_id"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["item_id"] = o.ItemId
	toSerialize["content_index"] = o.ContentIndex
	toSerialize["transcript"] = o.Transcript
	return toSerialize, nil
}

func (o *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_id",
		"type",
		"item_id",
		"content_index",
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

	varRealtimeServerEventConversationItemInputAudioTranscriptionCompleted := _RealtimeServerEventConversationItemInputAudioTranscriptionCompleted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeServerEventConversationItemInputAudioTranscriptionCompleted)

	if err != nil {
		return err
	}

	*o = RealtimeServerEventConversationItemInputAudioTranscriptionCompleted(varRealtimeServerEventConversationItemInputAudioTranscriptionCompleted)

	return err
}

type NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted struct {
	value *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted
	isSet bool
}

func (v NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted) Get() *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted {
	return v.value
}

func (v *NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted) Set(val *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted(val *RealtimeServerEventConversationItemInputAudioTranscriptionCompleted) *NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted {
	return &NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted{value: val, isSet: true}
}

func (v NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventConversationItemInputAudioTranscriptionCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


