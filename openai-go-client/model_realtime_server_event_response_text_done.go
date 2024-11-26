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

// checks if the RealtimeServerEventResponseTextDone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventResponseTextDone{}

// RealtimeServerEventResponseTextDone Returned when the text value of a \"text\" content part is done streaming. Also emitted when a Response is interrupted, incomplete, or cancelled.
type RealtimeServerEventResponseTextDone struct {
	// The unique ID of the server event.
	EventId string `json:"event_id"`
	// The event type, must be \"response.text.done\".
	Type string `json:"type"`
	// The ID of the response.
	ResponseId string `json:"response_id"`
	// The ID of the item.
	ItemId string `json:"item_id"`
	// The index of the output item in the response.
	OutputIndex int32 `json:"output_index"`
	// The index of the content part in the item's content array.
	ContentIndex int32 `json:"content_index"`
	// The final text content.
	Text string `json:"text"`
}

type _RealtimeServerEventResponseTextDone RealtimeServerEventResponseTextDone

// NewRealtimeServerEventResponseTextDone instantiates a new RealtimeServerEventResponseTextDone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventResponseTextDone(eventId string, type_ string, responseId string, itemId string, outputIndex int32, contentIndex int32, text string) *RealtimeServerEventResponseTextDone {
	this := RealtimeServerEventResponseTextDone{}
	this.EventId = eventId
	this.Type = type_
	this.ResponseId = responseId
	this.ItemId = itemId
	this.OutputIndex = outputIndex
	this.ContentIndex = contentIndex
	this.Text = text
	return &this
}

// NewRealtimeServerEventResponseTextDoneWithDefaults instantiates a new RealtimeServerEventResponseTextDone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventResponseTextDoneWithDefaults() *RealtimeServerEventResponseTextDone {
	this := RealtimeServerEventResponseTextDone{}
	return &this
}

// GetEventId returns the EventId field value
func (o *RealtimeServerEventResponseTextDone) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *RealtimeServerEventResponseTextDone) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *RealtimeServerEventResponseTextDone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeServerEventResponseTextDone) SetType(v string) {
	o.Type = v
}

// GetResponseId returns the ResponseId field value
func (o *RealtimeServerEventResponseTextDone) GetResponseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetResponseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseId, true
}

// SetResponseId sets field value
func (o *RealtimeServerEventResponseTextDone) SetResponseId(v string) {
	o.ResponseId = v
}

// GetItemId returns the ItemId field value
func (o *RealtimeServerEventResponseTextDone) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RealtimeServerEventResponseTextDone) SetItemId(v string) {
	o.ItemId = v
}

// GetOutputIndex returns the OutputIndex field value
func (o *RealtimeServerEventResponseTextDone) GetOutputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutputIndex
}

// GetOutputIndexOk returns a tuple with the OutputIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetOutputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputIndex, true
}

// SetOutputIndex sets field value
func (o *RealtimeServerEventResponseTextDone) SetOutputIndex(v int32) {
	o.OutputIndex = v
}

// GetContentIndex returns the ContentIndex field value
func (o *RealtimeServerEventResponseTextDone) GetContentIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentIndex
}

// GetContentIndexOk returns a tuple with the ContentIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetContentIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentIndex, true
}

// SetContentIndex sets field value
func (o *RealtimeServerEventResponseTextDone) SetContentIndex(v int32) {
	o.ContentIndex = v
}

// GetText returns the Text field value
func (o *RealtimeServerEventResponseTextDone) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseTextDone) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *RealtimeServerEventResponseTextDone) SetText(v string) {
	o.Text = v
}

func (o RealtimeServerEventResponseTextDone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventResponseTextDone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_id"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["response_id"] = o.ResponseId
	toSerialize["item_id"] = o.ItemId
	toSerialize["output_index"] = o.OutputIndex
	toSerialize["content_index"] = o.ContentIndex
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

func (o *RealtimeServerEventResponseTextDone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_id",
		"type",
		"response_id",
		"item_id",
		"output_index",
		"content_index",
		"text",
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

	varRealtimeServerEventResponseTextDone := _RealtimeServerEventResponseTextDone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeServerEventResponseTextDone)

	if err != nil {
		return err
	}

	*o = RealtimeServerEventResponseTextDone(varRealtimeServerEventResponseTextDone)

	return err
}

type NullableRealtimeServerEventResponseTextDone struct {
	value *RealtimeServerEventResponseTextDone
	isSet bool
}

func (v NullableRealtimeServerEventResponseTextDone) Get() *RealtimeServerEventResponseTextDone {
	return v.value
}

func (v *NullableRealtimeServerEventResponseTextDone) Set(val *RealtimeServerEventResponseTextDone) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventResponseTextDone) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventResponseTextDone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventResponseTextDone(val *RealtimeServerEventResponseTextDone) *NullableRealtimeServerEventResponseTextDone {
	return &NullableRealtimeServerEventResponseTextDone{value: val, isSet: true}
}

func (v NullableRealtimeServerEventResponseTextDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventResponseTextDone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


