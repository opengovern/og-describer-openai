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

// checks if the RunStreamEventOneOf9 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStreamEventOneOf9{}

// RunStreamEventOneOf9 Occurs when a [run](/docs/api-reference/runs/object) expires.
type RunStreamEventOneOf9 struct {
	Event string `json:"event"`
	Data RunObject `json:"data"`
}

type _RunStreamEventOneOf9 RunStreamEventOneOf9

// NewRunStreamEventOneOf9 instantiates a new RunStreamEventOneOf9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStreamEventOneOf9(event string, data RunObject) *RunStreamEventOneOf9 {
	this := RunStreamEventOneOf9{}
	this.Event = event
	this.Data = data
	return &this
}

// NewRunStreamEventOneOf9WithDefaults instantiates a new RunStreamEventOneOf9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStreamEventOneOf9WithDefaults() *RunStreamEventOneOf9 {
	this := RunStreamEventOneOf9{}
	return &this
}

// GetEvent returns the Event field value
func (o *RunStreamEventOneOf9) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *RunStreamEventOneOf9) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *RunStreamEventOneOf9) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *RunStreamEventOneOf9) GetData() RunObject {
	if o == nil {
		var ret RunObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RunStreamEventOneOf9) GetDataOk() (*RunObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RunStreamEventOneOf9) SetData(v RunObject) {
	o.Data = v
}

func (o RunStreamEventOneOf9) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStreamEventOneOf9) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RunStreamEventOneOf9) UnmarshalJSON(data []byte) (err error) {
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

	varRunStreamEventOneOf9 := _RunStreamEventOneOf9{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStreamEventOneOf9)

	if err != nil {
		return err
	}

	*o = RunStreamEventOneOf9(varRunStreamEventOneOf9)

	return err
}

type NullableRunStreamEventOneOf9 struct {
	value *RunStreamEventOneOf9
	isSet bool
}

func (v NullableRunStreamEventOneOf9) Get() *RunStreamEventOneOf9 {
	return v.value
}

func (v *NullableRunStreamEventOneOf9) Set(val *RunStreamEventOneOf9) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStreamEventOneOf9) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStreamEventOneOf9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStreamEventOneOf9(val *RunStreamEventOneOf9) *NullableRunStreamEventOneOf9 {
	return &NullableRunStreamEventOneOf9{value: val, isSet: true}
}

func (v NullableRunStreamEventOneOf9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStreamEventOneOf9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

