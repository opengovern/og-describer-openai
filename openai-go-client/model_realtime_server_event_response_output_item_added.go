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

// checks if the RealtimeServerEventResponseOutputItemAdded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventResponseOutputItemAdded{}

// RealtimeServerEventResponseOutputItemAdded Returned when a new Item is created during Response generation.
type RealtimeServerEventResponseOutputItemAdded struct {
	// The unique ID of the server event.
	EventId string `json:"event_id"`
	// The event type, must be `response.output_item.added`.
	Type string `json:"type"`
	// The ID of the Response to which the item belongs.
	ResponseId string `json:"response_id"`
	// The index of the output item in the Response.
	OutputIndex int32 `json:"output_index"`
	Item RealtimeConversationItem `json:"item"`
}

type _RealtimeServerEventResponseOutputItemAdded RealtimeServerEventResponseOutputItemAdded

// NewRealtimeServerEventResponseOutputItemAdded instantiates a new RealtimeServerEventResponseOutputItemAdded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventResponseOutputItemAdded(eventId string, type_ string, responseId string, outputIndex int32, item RealtimeConversationItem) *RealtimeServerEventResponseOutputItemAdded {
	this := RealtimeServerEventResponseOutputItemAdded{}
	this.EventId = eventId
	this.Type = type_
	this.ResponseId = responseId
	this.OutputIndex = outputIndex
	this.Item = item
	return &this
}

// NewRealtimeServerEventResponseOutputItemAddedWithDefaults instantiates a new RealtimeServerEventResponseOutputItemAdded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventResponseOutputItemAddedWithDefaults() *RealtimeServerEventResponseOutputItemAdded {
	this := RealtimeServerEventResponseOutputItemAdded{}
	return &this
}

// GetEventId returns the EventId field value
func (o *RealtimeServerEventResponseOutputItemAdded) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseOutputItemAdded) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *RealtimeServerEventResponseOutputItemAdded) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *RealtimeServerEventResponseOutputItemAdded) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseOutputItemAdded) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeServerEventResponseOutputItemAdded) SetType(v string) {
	o.Type = v
}

// GetResponseId returns the ResponseId field value
func (o *RealtimeServerEventResponseOutputItemAdded) GetResponseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseOutputItemAdded) GetResponseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseId, true
}

// SetResponseId sets field value
func (o *RealtimeServerEventResponseOutputItemAdded) SetResponseId(v string) {
	o.ResponseId = v
}

// GetOutputIndex returns the OutputIndex field value
func (o *RealtimeServerEventResponseOutputItemAdded) GetOutputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutputIndex
}

// GetOutputIndexOk returns a tuple with the OutputIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseOutputItemAdded) GetOutputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputIndex, true
}

// SetOutputIndex sets field value
func (o *RealtimeServerEventResponseOutputItemAdded) SetOutputIndex(v int32) {
	o.OutputIndex = v
}

// GetItem returns the Item field value
func (o *RealtimeServerEventResponseOutputItemAdded) GetItem() RealtimeConversationItem {
	if o == nil {
		var ret RealtimeConversationItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseOutputItemAdded) GetItemOk() (*RealtimeConversationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *RealtimeServerEventResponseOutputItemAdded) SetItem(v RealtimeConversationItem) {
	o.Item = v
}

func (o RealtimeServerEventResponseOutputItemAdded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventResponseOutputItemAdded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_id"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["response_id"] = o.ResponseId
	toSerialize["output_index"] = o.OutputIndex
	toSerialize["item"] = o.Item
	return toSerialize, nil
}

func (o *RealtimeServerEventResponseOutputItemAdded) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_id",
		"type",
		"response_id",
		"output_index",
		"item",
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

	varRealtimeServerEventResponseOutputItemAdded := _RealtimeServerEventResponseOutputItemAdded{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeServerEventResponseOutputItemAdded)

	if err != nil {
		return err
	}

	*o = RealtimeServerEventResponseOutputItemAdded(varRealtimeServerEventResponseOutputItemAdded)

	return err
}

type NullableRealtimeServerEventResponseOutputItemAdded struct {
	value *RealtimeServerEventResponseOutputItemAdded
	isSet bool
}

func (v NullableRealtimeServerEventResponseOutputItemAdded) Get() *RealtimeServerEventResponseOutputItemAdded {
	return v.value
}

func (v *NullableRealtimeServerEventResponseOutputItemAdded) Set(val *RealtimeServerEventResponseOutputItemAdded) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventResponseOutputItemAdded) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventResponseOutputItemAdded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventResponseOutputItemAdded(val *RealtimeServerEventResponseOutputItemAdded) *NullableRealtimeServerEventResponseOutputItemAdded {
	return &NullableRealtimeServerEventResponseOutputItemAdded{value: val, isSet: true}
}

func (v NullableRealtimeServerEventResponseOutputItemAdded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventResponseOutputItemAdded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

