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

// checks if the RunStreamEventOneOf8 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStreamEventOneOf8{}

// RunStreamEventOneOf8 Occurs when a [run](/docs/api-reference/runs/object) is cancelled.
type RunStreamEventOneOf8 struct {
	Event string `json:"event"`
	Data RunObject `json:"data"`
}

type _RunStreamEventOneOf8 RunStreamEventOneOf8

// NewRunStreamEventOneOf8 instantiates a new RunStreamEventOneOf8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStreamEventOneOf8(event string, data RunObject) *RunStreamEventOneOf8 {
	this := RunStreamEventOneOf8{}
	this.Event = event
	this.Data = data
	return &this
}

// NewRunStreamEventOneOf8WithDefaults instantiates a new RunStreamEventOneOf8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStreamEventOneOf8WithDefaults() *RunStreamEventOneOf8 {
	this := RunStreamEventOneOf8{}
	return &this
}

// GetEvent returns the Event field value
func (o *RunStreamEventOneOf8) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *RunStreamEventOneOf8) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *RunStreamEventOneOf8) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *RunStreamEventOneOf8) GetData() RunObject {
	if o == nil {
		var ret RunObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RunStreamEventOneOf8) GetDataOk() (*RunObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RunStreamEventOneOf8) SetData(v RunObject) {
	o.Data = v
}

func (o RunStreamEventOneOf8) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStreamEventOneOf8) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RunStreamEventOneOf8) UnmarshalJSON(data []byte) (err error) {
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

	varRunStreamEventOneOf8 := _RunStreamEventOneOf8{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStreamEventOneOf8)

	if err != nil {
		return err
	}

	*o = RunStreamEventOneOf8(varRunStreamEventOneOf8)

	return err
}

type NullableRunStreamEventOneOf8 struct {
	value *RunStreamEventOneOf8
	isSet bool
}

func (v NullableRunStreamEventOneOf8) Get() *RunStreamEventOneOf8 {
	return v.value
}

func (v *NullableRunStreamEventOneOf8) Set(val *RunStreamEventOneOf8) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStreamEventOneOf8) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStreamEventOneOf8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStreamEventOneOf8(val *RunStreamEventOneOf8) *NullableRunStreamEventOneOf8 {
	return &NullableRunStreamEventOneOf8{value: val, isSet: true}
}

func (v NullableRunStreamEventOneOf8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStreamEventOneOf8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


