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

// checks if the RunStepDeltaObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaObject{}

// RunStepDeltaObject Represents a run step delta i.e. any changed fields on a run step during streaming. 
type RunStepDeltaObject struct {
	// The identifier of the run step, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The object type, which is always `thread.run.step.delta`.
	Object string `json:"object"`
	Delta RunStepDeltaObjectDelta `json:"delta"`
}

type _RunStepDeltaObject RunStepDeltaObject

// NewRunStepDeltaObject instantiates a new RunStepDeltaObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaObject(id string, object string, delta RunStepDeltaObjectDelta) *RunStepDeltaObject {
	this := RunStepDeltaObject{}
	this.Id = id
	this.Object = object
	this.Delta = delta
	return &this
}

// NewRunStepDeltaObjectWithDefaults instantiates a new RunStepDeltaObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaObjectWithDefaults() *RunStepDeltaObject {
	this := RunStepDeltaObject{}
	return &this
}

// GetId returns the Id field value
func (o *RunStepDeltaObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunStepDeltaObject) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *RunStepDeltaObject) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaObject) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RunStepDeltaObject) SetObject(v string) {
	o.Object = v
}

// GetDelta returns the Delta field value
func (o *RunStepDeltaObject) GetDelta() RunStepDeltaObjectDelta {
	if o == nil {
		var ret RunStepDeltaObjectDelta
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaObject) GetDeltaOk() (*RunStepDeltaObjectDelta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *RunStepDeltaObject) SetDelta(v RunStepDeltaObjectDelta) {
	o.Delta = v
}

func (o RunStepDeltaObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["delta"] = o.Delta
	return toSerialize, nil
}

func (o *RunStepDeltaObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"delta",
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

	varRunStepDeltaObject := _RunStepDeltaObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDeltaObject)

	if err != nil {
		return err
	}

	*o = RunStepDeltaObject(varRunStepDeltaObject)

	return err
}

type NullableRunStepDeltaObject struct {
	value *RunStepDeltaObject
	isSet bool
}

func (v NullableRunStepDeltaObject) Get() *RunStepDeltaObject {
	return v.value
}

func (v *NullableRunStepDeltaObject) Set(val *RunStepDeltaObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaObject(val *RunStepDeltaObject) *NullableRunStepDeltaObject {
	return &NullableRunStepDeltaObject{value: val, isSet: true}
}

func (v NullableRunStepDeltaObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

