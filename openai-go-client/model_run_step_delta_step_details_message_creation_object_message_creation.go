/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RunStepDeltaStepDetailsMessageCreationObjectMessageCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsMessageCreationObjectMessageCreation{}

// RunStepDeltaStepDetailsMessageCreationObjectMessageCreation struct for RunStepDeltaStepDetailsMessageCreationObjectMessageCreation
type RunStepDeltaStepDetailsMessageCreationObjectMessageCreation struct {
	// The ID of the message that was created by this run step.
	MessageId *string `json:"message_id,omitempty"`
}

// NewRunStepDeltaStepDetailsMessageCreationObjectMessageCreation instantiates a new RunStepDeltaStepDetailsMessageCreationObjectMessageCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsMessageCreationObjectMessageCreation() *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation {
	this := RunStepDeltaStepDetailsMessageCreationObjectMessageCreation{}
	return &this
}

// NewRunStepDeltaStepDetailsMessageCreationObjectMessageCreationWithDefaults instantiates a new RunStepDeltaStepDetailsMessageCreationObjectMessageCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsMessageCreationObjectMessageCreationWithDefaults() *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation {
	this := RunStepDeltaStepDetailsMessageCreationObjectMessageCreation{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) SetMessageId(v string) {
	o.MessageId = &v
}

func (o RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	return toSerialize, nil
}

type NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation struct {
	value *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation) Get() *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation) Set(val *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation(val *RunStepDeltaStepDetailsMessageCreationObjectMessageCreation) *NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation {
	return &NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsMessageCreationObjectMessageCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


