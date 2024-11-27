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

// checks if the RealtimeClientEventConversationItemTruncate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeClientEventConversationItemTruncate{}

// RealtimeClientEventConversationItemTruncate Send this event to truncate a previous assistant message’s audio. The server will produce audio faster than realtime, so this event is useful when the user interrupts to truncate audio that has already been sent to the client but not yet played. This will synchronize the server's understanding of the audio with the client's playback. Truncating audio will delete the server-side text transcript to ensure there is not text in the context that hasn't been heard by the user. If successful, the server will respond with a `conversation.item.truncated` event. 
type RealtimeClientEventConversationItemTruncate struct {
	// Optional client-generated ID used to identify this event.
	EventId *string `json:"event_id,omitempty"`
	// The event type, must be \"conversation.item.truncate\".
	Type string `json:"type"`
	// The ID of the assistant message item to truncate. Only assistant message items can be truncated.
	ItemId string `json:"item_id"`
	// The index of the content part to truncate. Set this to 0.
	ContentIndex int32 `json:"content_index"`
	// Inclusive duration up to which audio is truncated, in milliseconds. If the audio_end_ms is greater than the actual audio duration, the server will respond with an error.
	AudioEndMs int32 `json:"audio_end_ms"`
}

type _RealtimeClientEventConversationItemTruncate RealtimeClientEventConversationItemTruncate

// NewRealtimeClientEventConversationItemTruncate instantiates a new RealtimeClientEventConversationItemTruncate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeClientEventConversationItemTruncate(type_ string, itemId string, contentIndex int32, audioEndMs int32) *RealtimeClientEventConversationItemTruncate {
	this := RealtimeClientEventConversationItemTruncate{}
	this.Type = type_
	this.ItemId = itemId
	this.ContentIndex = contentIndex
	this.AudioEndMs = audioEndMs
	return &this
}

// NewRealtimeClientEventConversationItemTruncateWithDefaults instantiates a new RealtimeClientEventConversationItemTruncate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeClientEventConversationItemTruncateWithDefaults() *RealtimeClientEventConversationItemTruncate {
	this := RealtimeClientEventConversationItemTruncate{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *RealtimeClientEventConversationItemTruncate) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemTruncate) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *RealtimeClientEventConversationItemTruncate) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *RealtimeClientEventConversationItemTruncate) SetEventId(v string) {
	o.EventId = &v
}

// GetType returns the Type field value
func (o *RealtimeClientEventConversationItemTruncate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemTruncate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeClientEventConversationItemTruncate) SetType(v string) {
	o.Type = v
}

// GetItemId returns the ItemId field value
func (o *RealtimeClientEventConversationItemTruncate) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemTruncate) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RealtimeClientEventConversationItemTruncate) SetItemId(v string) {
	o.ItemId = v
}

// GetContentIndex returns the ContentIndex field value
func (o *RealtimeClientEventConversationItemTruncate) GetContentIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentIndex
}

// GetContentIndexOk returns a tuple with the ContentIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemTruncate) GetContentIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentIndex, true
}

// SetContentIndex sets field value
func (o *RealtimeClientEventConversationItemTruncate) SetContentIndex(v int32) {
	o.ContentIndex = v
}

// GetAudioEndMs returns the AudioEndMs field value
func (o *RealtimeClientEventConversationItemTruncate) GetAudioEndMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AudioEndMs
}

// GetAudioEndMsOk returns a tuple with the AudioEndMs field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemTruncate) GetAudioEndMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudioEndMs, true
}

// SetAudioEndMs sets field value
func (o *RealtimeClientEventConversationItemTruncate) SetAudioEndMs(v int32) {
	o.AudioEndMs = v
}

func (o RealtimeClientEventConversationItemTruncate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeClientEventConversationItemTruncate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	toSerialize["type"] = o.Type
	toSerialize["item_id"] = o.ItemId
	toSerialize["content_index"] = o.ContentIndex
	toSerialize["audio_end_ms"] = o.AudioEndMs
	return toSerialize, nil
}

func (o *RealtimeClientEventConversationItemTruncate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"item_id",
		"content_index",
		"audio_end_ms",
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

	varRealtimeClientEventConversationItemTruncate := _RealtimeClientEventConversationItemTruncate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeClientEventConversationItemTruncate)

	if err != nil {
		return err
	}

	*o = RealtimeClientEventConversationItemTruncate(varRealtimeClientEventConversationItemTruncate)

	return err
}

type NullableRealtimeClientEventConversationItemTruncate struct {
	value *RealtimeClientEventConversationItemTruncate
	isSet bool
}

func (v NullableRealtimeClientEventConversationItemTruncate) Get() *RealtimeClientEventConversationItemTruncate {
	return v.value
}

func (v *NullableRealtimeClientEventConversationItemTruncate) Set(val *RealtimeClientEventConversationItemTruncate) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeClientEventConversationItemTruncate) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeClientEventConversationItemTruncate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeClientEventConversationItemTruncate(val *RealtimeClientEventConversationItemTruncate) *NullableRealtimeClientEventConversationItemTruncate {
	return &NullableRealtimeClientEventConversationItemTruncate{value: val, isSet: true}
}

func (v NullableRealtimeClientEventConversationItemTruncate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeClientEventConversationItemTruncate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

