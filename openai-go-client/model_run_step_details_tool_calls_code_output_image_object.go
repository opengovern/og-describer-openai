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

// checks if the RunStepDetailsToolCallsCodeOutputImageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsToolCallsCodeOutputImageObject{}

// RunStepDetailsToolCallsCodeOutputImageObject struct for RunStepDetailsToolCallsCodeOutputImageObject
type RunStepDetailsToolCallsCodeOutputImageObject struct {
	// Always `image`.
	Type string `json:"type"`
	Image RunStepDetailsToolCallsCodeOutputImageObjectImage `json:"image"`
}

type _RunStepDetailsToolCallsCodeOutputImageObject RunStepDetailsToolCallsCodeOutputImageObject

// NewRunStepDetailsToolCallsCodeOutputImageObject instantiates a new RunStepDetailsToolCallsCodeOutputImageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsToolCallsCodeOutputImageObject(type_ string, image RunStepDetailsToolCallsCodeOutputImageObjectImage) *RunStepDetailsToolCallsCodeOutputImageObject {
	this := RunStepDetailsToolCallsCodeOutputImageObject{}
	this.Type = type_
	this.Image = image
	return &this
}

// NewRunStepDetailsToolCallsCodeOutputImageObjectWithDefaults instantiates a new RunStepDetailsToolCallsCodeOutputImageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsToolCallsCodeOutputImageObjectWithDefaults() *RunStepDetailsToolCallsCodeOutputImageObject {
	this := RunStepDetailsToolCallsCodeOutputImageObject{}
	return &this
}

// GetType returns the Type field value
func (o *RunStepDetailsToolCallsCodeOutputImageObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeOutputImageObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDetailsToolCallsCodeOutputImageObject) SetType(v string) {
	o.Type = v
}

// GetImage returns the Image field value
func (o *RunStepDetailsToolCallsCodeOutputImageObject) GetImage() RunStepDetailsToolCallsCodeOutputImageObjectImage {
	if o == nil {
		var ret RunStepDetailsToolCallsCodeOutputImageObjectImage
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeOutputImageObject) GetImageOk() (*RunStepDetailsToolCallsCodeOutputImageObjectImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *RunStepDetailsToolCallsCodeOutputImageObject) SetImage(v RunStepDetailsToolCallsCodeOutputImageObjectImage) {
	o.Image = v
}

func (o RunStepDetailsToolCallsCodeOutputImageObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsToolCallsCodeOutputImageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["image"] = o.Image
	return toSerialize, nil
}

func (o *RunStepDetailsToolCallsCodeOutputImageObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"image",
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

	varRunStepDetailsToolCallsCodeOutputImageObject := _RunStepDetailsToolCallsCodeOutputImageObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDetailsToolCallsCodeOutputImageObject)

	if err != nil {
		return err
	}

	*o = RunStepDetailsToolCallsCodeOutputImageObject(varRunStepDetailsToolCallsCodeOutputImageObject)

	return err
}

type NullableRunStepDetailsToolCallsCodeOutputImageObject struct {
	value *RunStepDetailsToolCallsCodeOutputImageObject
	isSet bool
}

func (v NullableRunStepDetailsToolCallsCodeOutputImageObject) Get() *RunStepDetailsToolCallsCodeOutputImageObject {
	return v.value
}

func (v *NullableRunStepDetailsToolCallsCodeOutputImageObject) Set(val *RunStepDetailsToolCallsCodeOutputImageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsToolCallsCodeOutputImageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsToolCallsCodeOutputImageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsToolCallsCodeOutputImageObject(val *RunStepDetailsToolCallsCodeOutputImageObject) *NullableRunStepDetailsToolCallsCodeOutputImageObject {
	return &NullableRunStepDetailsToolCallsCodeOutputImageObject{value: val, isSet: true}
}

func (v NullableRunStepDetailsToolCallsCodeOutputImageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsToolCallsCodeOutputImageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


