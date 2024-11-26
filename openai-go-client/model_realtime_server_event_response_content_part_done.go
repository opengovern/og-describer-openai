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

// checks if the RealtimeServerEventResponseContentPartDone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventResponseContentPartDone{}

// RealtimeServerEventResponseContentPartDone Returned when a content part is done streaming in an assistant message item. Also emitted when a Response is interrupted, incomplete, or cancelled.
type RealtimeServerEventResponseContentPartDone struct {
	// The unique ID of the server event.
	EventId string `json:"event_id"`
	// The event type, must be \"response.content_part.done\".
	Type string `json:"type"`
	// The ID of the response.
	ResponseId string `json:"response_id"`
	// The ID of the item.
	ItemId string `json:"item_id"`
	// The index of the output item in the response.
	OutputIndex int32 `json:"output_index"`
	// The index of the content part in the item's content array.
	ContentIndex int32 `json:"content_index"`
	Part RealtimeServerEventResponseContentPartDonePart `json:"part"`
}

type _RealtimeServerEventResponseContentPartDone RealtimeServerEventResponseContentPartDone

// NewRealtimeServerEventResponseContentPartDone instantiates a new RealtimeServerEventResponseContentPartDone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventResponseContentPartDone(eventId string, type_ string, responseId string, itemId string, outputIndex int32, contentIndex int32, part RealtimeServerEventResponseContentPartDonePart) *RealtimeServerEventResponseContentPartDone {
	this := RealtimeServerEventResponseContentPartDone{}
	this.EventId = eventId
	this.Type = type_
	this.ResponseId = responseId
	this.ItemId = itemId
	this.OutputIndex = outputIndex
	this.ContentIndex = contentIndex
	this.Part = part
	return &this
}

// NewRealtimeServerEventResponseContentPartDoneWithDefaults instantiates a new RealtimeServerEventResponseContentPartDone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventResponseContentPartDoneWithDefaults() *RealtimeServerEventResponseContentPartDone {
	this := RealtimeServerEventResponseContentPartDone{}
	return &this
}

// GetEventId returns the EventId field value
func (o *RealtimeServerEventResponseContentPartDone) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *RealtimeServerEventResponseContentPartDone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetType(v string) {
	o.Type = v
}

// GetResponseId returns the ResponseId field value
func (o *RealtimeServerEventResponseContentPartDone) GetResponseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetResponseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseId, true
}

// SetResponseId sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetResponseId(v string) {
	o.ResponseId = v
}

// GetItemId returns the ItemId field value
func (o *RealtimeServerEventResponseContentPartDone) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetItemId(v string) {
	o.ItemId = v
}

// GetOutputIndex returns the OutputIndex field value
func (o *RealtimeServerEventResponseContentPartDone) GetOutputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutputIndex
}

// GetOutputIndexOk returns a tuple with the OutputIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetOutputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputIndex, true
}

// SetOutputIndex sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetOutputIndex(v int32) {
	o.OutputIndex = v
}

// GetContentIndex returns the ContentIndex field value
func (o *RealtimeServerEventResponseContentPartDone) GetContentIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentIndex
}

// GetContentIndexOk returns a tuple with the ContentIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetContentIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentIndex, true
}

// SetContentIndex sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetContentIndex(v int32) {
	o.ContentIndex = v
}

// GetPart returns the Part field value
func (o *RealtimeServerEventResponseContentPartDone) GetPart() RealtimeServerEventResponseContentPartDonePart {
	if o == nil {
		var ret RealtimeServerEventResponseContentPartDonePart
		return ret
	}

	return o.Part
}

// GetPartOk returns a tuple with the Part field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseContentPartDone) GetPartOk() (*RealtimeServerEventResponseContentPartDonePart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Part, true
}

// SetPart sets field value
func (o *RealtimeServerEventResponseContentPartDone) SetPart(v RealtimeServerEventResponseContentPartDonePart) {
	o.Part = v
}

func (o RealtimeServerEventResponseContentPartDone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventResponseContentPartDone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_id"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["response_id"] = o.ResponseId
	toSerialize["item_id"] = o.ItemId
	toSerialize["output_index"] = o.OutputIndex
	toSerialize["content_index"] = o.ContentIndex
	toSerialize["part"] = o.Part
	return toSerialize, nil
}

func (o *RealtimeServerEventResponseContentPartDone) UnmarshalJSON(data []byte) (err error) {
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
		"part",
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

	varRealtimeServerEventResponseContentPartDone := _RealtimeServerEventResponseContentPartDone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeServerEventResponseContentPartDone)

	if err != nil {
		return err
	}

	*o = RealtimeServerEventResponseContentPartDone(varRealtimeServerEventResponseContentPartDone)

	return err
}

type NullableRealtimeServerEventResponseContentPartDone struct {
	value *RealtimeServerEventResponseContentPartDone
	isSet bool
}

func (v NullableRealtimeServerEventResponseContentPartDone) Get() *RealtimeServerEventResponseContentPartDone {
	return v.value
}

func (v *NullableRealtimeServerEventResponseContentPartDone) Set(val *RealtimeServerEventResponseContentPartDone) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventResponseContentPartDone) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventResponseContentPartDone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventResponseContentPartDone(val *RealtimeServerEventResponseContentPartDone) *NullableRealtimeServerEventResponseContentPartDone {
	return &NullableRealtimeServerEventResponseContentPartDone{value: val, isSet: true}
}

func (v NullableRealtimeServerEventResponseContentPartDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventResponseContentPartDone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


