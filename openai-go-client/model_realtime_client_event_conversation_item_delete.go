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

// checks if the RealtimeClientEventConversationItemDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeClientEventConversationItemDelete{}

// RealtimeClientEventConversationItemDelete Send this event when you want to remove any item from the conversation history. The server will respond with a `conversation.item.deleted` event, unless the item does not exist in the conversation history, in which case the server will respond with an error.
type RealtimeClientEventConversationItemDelete struct {
	// Optional client-generated ID used to identify this event.
	EventId *string `json:"event_id,omitempty"`
	// The event type, must be \"conversation.item.delete\".
	Type string `json:"type"`
	// The ID of the item to delete.
	ItemId string `json:"item_id"`
}

type _RealtimeClientEventConversationItemDelete RealtimeClientEventConversationItemDelete

// NewRealtimeClientEventConversationItemDelete instantiates a new RealtimeClientEventConversationItemDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeClientEventConversationItemDelete(type_ string, itemId string) *RealtimeClientEventConversationItemDelete {
	this := RealtimeClientEventConversationItemDelete{}
	this.Type = type_
	this.ItemId = itemId
	return &this
}

// NewRealtimeClientEventConversationItemDeleteWithDefaults instantiates a new RealtimeClientEventConversationItemDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeClientEventConversationItemDeleteWithDefaults() *RealtimeClientEventConversationItemDelete {
	this := RealtimeClientEventConversationItemDelete{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *RealtimeClientEventConversationItemDelete) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemDelete) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *RealtimeClientEventConversationItemDelete) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *RealtimeClientEventConversationItemDelete) SetEventId(v string) {
	o.EventId = &v
}

// GetType returns the Type field value
func (o *RealtimeClientEventConversationItemDelete) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemDelete) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeClientEventConversationItemDelete) SetType(v string) {
	o.Type = v
}

// GetItemId returns the ItemId field value
func (o *RealtimeClientEventConversationItemDelete) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventConversationItemDelete) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RealtimeClientEventConversationItemDelete) SetItemId(v string) {
	o.ItemId = v
}

func (o RealtimeClientEventConversationItemDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeClientEventConversationItemDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	toSerialize["type"] = o.Type
	toSerialize["item_id"] = o.ItemId
	return toSerialize, nil
}

func (o *RealtimeClientEventConversationItemDelete) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"item_id",
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

	varRealtimeClientEventConversationItemDelete := _RealtimeClientEventConversationItemDelete{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeClientEventConversationItemDelete)

	if err != nil {
		return err
	}

	*o = RealtimeClientEventConversationItemDelete(varRealtimeClientEventConversationItemDelete)

	return err
}

type NullableRealtimeClientEventConversationItemDelete struct {
	value *RealtimeClientEventConversationItemDelete
	isSet bool
}

func (v NullableRealtimeClientEventConversationItemDelete) Get() *RealtimeClientEventConversationItemDelete {
	return v.value
}

func (v *NullableRealtimeClientEventConversationItemDelete) Set(val *RealtimeClientEventConversationItemDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeClientEventConversationItemDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeClientEventConversationItemDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeClientEventConversationItemDelete(val *RealtimeClientEventConversationItemDelete) *NullableRealtimeClientEventConversationItemDelete {
	return &NullableRealtimeClientEventConversationItemDelete{value: val, isSet: true}
}

func (v NullableRealtimeClientEventConversationItemDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeClientEventConversationItemDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


