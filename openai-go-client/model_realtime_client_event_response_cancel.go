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

// checks if the RealtimeClientEventResponseCancel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeClientEventResponseCancel{}

// RealtimeClientEventResponseCancel Send this event to cancel an in-progress response. The server will respond with a `response.cancelled` event or an error if there is no response to cancel.
type RealtimeClientEventResponseCancel struct {
	// Optional client-generated ID used to identify this event.
	EventId *string `json:"event_id,omitempty"`
	// The event type, must be `response.cancel`.
	Type string `json:"type"`
}

type _RealtimeClientEventResponseCancel RealtimeClientEventResponseCancel

// NewRealtimeClientEventResponseCancel instantiates a new RealtimeClientEventResponseCancel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeClientEventResponseCancel(type_ string) *RealtimeClientEventResponseCancel {
	this := RealtimeClientEventResponseCancel{}
	this.Type = type_
	return &this
}

// NewRealtimeClientEventResponseCancelWithDefaults instantiates a new RealtimeClientEventResponseCancel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeClientEventResponseCancelWithDefaults() *RealtimeClientEventResponseCancel {
	this := RealtimeClientEventResponseCancel{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *RealtimeClientEventResponseCancel) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventResponseCancel) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *RealtimeClientEventResponseCancel) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *RealtimeClientEventResponseCancel) SetEventId(v string) {
	o.EventId = &v
}

// GetType returns the Type field value
func (o *RealtimeClientEventResponseCancel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeClientEventResponseCancel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeClientEventResponseCancel) SetType(v string) {
	o.Type = v
}

func (o RealtimeClientEventResponseCancel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeClientEventResponseCancel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *RealtimeClientEventResponseCancel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varRealtimeClientEventResponseCancel := _RealtimeClientEventResponseCancel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeClientEventResponseCancel)

	if err != nil {
		return err
	}

	*o = RealtimeClientEventResponseCancel(varRealtimeClientEventResponseCancel)

	return err
}

type NullableRealtimeClientEventResponseCancel struct {
	value *RealtimeClientEventResponseCancel
	isSet bool
}

func (v NullableRealtimeClientEventResponseCancel) Get() *RealtimeClientEventResponseCancel {
	return v.value
}

func (v *NullableRealtimeClientEventResponseCancel) Set(val *RealtimeClientEventResponseCancel) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeClientEventResponseCancel) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeClientEventResponseCancel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeClientEventResponseCancel(val *RealtimeClientEventResponseCancel) *NullableRealtimeClientEventResponseCancel {
	return &NullableRealtimeClientEventResponseCancel{value: val, isSet: true}
}

func (v NullableRealtimeClientEventResponseCancel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeClientEventResponseCancel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

