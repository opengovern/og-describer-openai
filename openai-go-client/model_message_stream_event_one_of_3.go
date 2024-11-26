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

// checks if the MessageStreamEventOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageStreamEventOneOf3{}

// MessageStreamEventOneOf3 Occurs when a [message](/docs/api-reference/messages/object) is completed.
type MessageStreamEventOneOf3 struct {
	Event string `json:"event"`
	Data MessageObject `json:"data"`
}

type _MessageStreamEventOneOf3 MessageStreamEventOneOf3

// NewMessageStreamEventOneOf3 instantiates a new MessageStreamEventOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageStreamEventOneOf3(event string, data MessageObject) *MessageStreamEventOneOf3 {
	this := MessageStreamEventOneOf3{}
	this.Event = event
	this.Data = data
	return &this
}

// NewMessageStreamEventOneOf3WithDefaults instantiates a new MessageStreamEventOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageStreamEventOneOf3WithDefaults() *MessageStreamEventOneOf3 {
	this := MessageStreamEventOneOf3{}
	return &this
}

// GetEvent returns the Event field value
func (o *MessageStreamEventOneOf3) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *MessageStreamEventOneOf3) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *MessageStreamEventOneOf3) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *MessageStreamEventOneOf3) GetData() MessageObject {
	if o == nil {
		var ret MessageObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MessageStreamEventOneOf3) GetDataOk() (*MessageObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MessageStreamEventOneOf3) SetData(v MessageObject) {
	o.Data = v
}

func (o MessageStreamEventOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageStreamEventOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *MessageStreamEventOneOf3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"data",
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

	varMessageStreamEventOneOf3 := _MessageStreamEventOneOf3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageStreamEventOneOf3)

	if err != nil {
		return err
	}

	*o = MessageStreamEventOneOf3(varMessageStreamEventOneOf3)

	return err
}

type NullableMessageStreamEventOneOf3 struct {
	value *MessageStreamEventOneOf3
	isSet bool
}

func (v NullableMessageStreamEventOneOf3) Get() *MessageStreamEventOneOf3 {
	return v.value
}

func (v *NullableMessageStreamEventOneOf3) Set(val *MessageStreamEventOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageStreamEventOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageStreamEventOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageStreamEventOneOf3(val *MessageStreamEventOneOf3) *NullableMessageStreamEventOneOf3 {
	return &NullableMessageStreamEventOneOf3{value: val, isSet: true}
}

func (v NullableMessageStreamEventOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageStreamEventOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


