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

// checks if the RunStepDetailsMessageCreationObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsMessageCreationObject{}

// RunStepDetailsMessageCreationObject Details of the message creation by the run step.
type RunStepDetailsMessageCreationObject struct {
	// Always `message_creation`.
	Type string `json:"type"`
	MessageCreation RunStepDetailsMessageCreationObjectMessageCreation `json:"message_creation"`
}

type _RunStepDetailsMessageCreationObject RunStepDetailsMessageCreationObject

// NewRunStepDetailsMessageCreationObject instantiates a new RunStepDetailsMessageCreationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsMessageCreationObject(type_ string, messageCreation RunStepDetailsMessageCreationObjectMessageCreation) *RunStepDetailsMessageCreationObject {
	this := RunStepDetailsMessageCreationObject{}
	this.Type = type_
	this.MessageCreation = messageCreation
	return &this
}

// NewRunStepDetailsMessageCreationObjectWithDefaults instantiates a new RunStepDetailsMessageCreationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsMessageCreationObjectWithDefaults() *RunStepDetailsMessageCreationObject {
	this := RunStepDetailsMessageCreationObject{}
	return &this
}

// GetType returns the Type field value
func (o *RunStepDetailsMessageCreationObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsMessageCreationObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDetailsMessageCreationObject) SetType(v string) {
	o.Type = v
}

// GetMessageCreation returns the MessageCreation field value
func (o *RunStepDetailsMessageCreationObject) GetMessageCreation() RunStepDetailsMessageCreationObjectMessageCreation {
	if o == nil {
		var ret RunStepDetailsMessageCreationObjectMessageCreation
		return ret
	}

	return o.MessageCreation
}

// GetMessageCreationOk returns a tuple with the MessageCreation field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsMessageCreationObject) GetMessageCreationOk() (*RunStepDetailsMessageCreationObjectMessageCreation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageCreation, true
}

// SetMessageCreation sets field value
func (o *RunStepDetailsMessageCreationObject) SetMessageCreation(v RunStepDetailsMessageCreationObjectMessageCreation) {
	o.MessageCreation = v
}

func (o RunStepDetailsMessageCreationObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsMessageCreationObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["message_creation"] = o.MessageCreation
	return toSerialize, nil
}

func (o *RunStepDetailsMessageCreationObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"message_creation",
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

	varRunStepDetailsMessageCreationObject := _RunStepDetailsMessageCreationObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDetailsMessageCreationObject)

	if err != nil {
		return err
	}

	*o = RunStepDetailsMessageCreationObject(varRunStepDetailsMessageCreationObject)

	return err
}

type NullableRunStepDetailsMessageCreationObject struct {
	value *RunStepDetailsMessageCreationObject
	isSet bool
}

func (v NullableRunStepDetailsMessageCreationObject) Get() *RunStepDetailsMessageCreationObject {
	return v.value
}

func (v *NullableRunStepDetailsMessageCreationObject) Set(val *RunStepDetailsMessageCreationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsMessageCreationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsMessageCreationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsMessageCreationObject(val *RunStepDetailsMessageCreationObject) *NullableRunStepDetailsMessageCreationObject {
	return &NullableRunStepDetailsMessageCreationObject{value: val, isSet: true}
}

func (v NullableRunStepDetailsMessageCreationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsMessageCreationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


