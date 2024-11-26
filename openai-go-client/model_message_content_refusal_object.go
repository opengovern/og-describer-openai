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

// checks if the MessageContentRefusalObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageContentRefusalObject{}

// MessageContentRefusalObject The refusal content generated by the assistant.
type MessageContentRefusalObject struct {
	// Always `refusal`.
	Type string `json:"type"`
	Refusal string `json:"refusal"`
}

type _MessageContentRefusalObject MessageContentRefusalObject

// NewMessageContentRefusalObject instantiates a new MessageContentRefusalObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageContentRefusalObject(type_ string, refusal string) *MessageContentRefusalObject {
	this := MessageContentRefusalObject{}
	this.Type = type_
	this.Refusal = refusal
	return &this
}

// NewMessageContentRefusalObjectWithDefaults instantiates a new MessageContentRefusalObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageContentRefusalObjectWithDefaults() *MessageContentRefusalObject {
	this := MessageContentRefusalObject{}
	return &this
}

// GetType returns the Type field value
func (o *MessageContentRefusalObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageContentRefusalObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageContentRefusalObject) SetType(v string) {
	o.Type = v
}

// GetRefusal returns the Refusal field value
func (o *MessageContentRefusalObject) GetRefusal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Refusal
}

// GetRefusalOk returns a tuple with the Refusal field value
// and a boolean to check if the value has been set.
func (o *MessageContentRefusalObject) GetRefusalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Refusal, true
}

// SetRefusal sets field value
func (o *MessageContentRefusalObject) SetRefusal(v string) {
	o.Refusal = v
}

func (o MessageContentRefusalObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageContentRefusalObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["refusal"] = o.Refusal
	return toSerialize, nil
}

func (o *MessageContentRefusalObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"refusal",
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

	varMessageContentRefusalObject := _MessageContentRefusalObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageContentRefusalObject)

	if err != nil {
		return err
	}

	*o = MessageContentRefusalObject(varMessageContentRefusalObject)

	return err
}

type NullableMessageContentRefusalObject struct {
	value *MessageContentRefusalObject
	isSet bool
}

func (v NullableMessageContentRefusalObject) Get() *MessageContentRefusalObject {
	return v.value
}

func (v *NullableMessageContentRefusalObject) Set(val *MessageContentRefusalObject) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageContentRefusalObject) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageContentRefusalObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageContentRefusalObject(val *MessageContentRefusalObject) *NullableMessageContentRefusalObject {
	return &NullableMessageContentRefusalObject{value: val, isSet: true}
}

func (v NullableMessageContentRefusalObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageContentRefusalObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


