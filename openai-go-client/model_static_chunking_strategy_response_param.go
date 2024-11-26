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

// checks if the StaticChunkingStrategyResponseParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticChunkingStrategyResponseParam{}

// StaticChunkingStrategyResponseParam struct for StaticChunkingStrategyResponseParam
type StaticChunkingStrategyResponseParam struct {
	// Always `static`.
	Type string `json:"type"`
	Static StaticChunkingStrategy `json:"static"`
}

type _StaticChunkingStrategyResponseParam StaticChunkingStrategyResponseParam

// NewStaticChunkingStrategyResponseParam instantiates a new StaticChunkingStrategyResponseParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticChunkingStrategyResponseParam(type_ string, static StaticChunkingStrategy) *StaticChunkingStrategyResponseParam {
	this := StaticChunkingStrategyResponseParam{}
	this.Type = type_
	this.Static = static
	return &this
}

// NewStaticChunkingStrategyResponseParamWithDefaults instantiates a new StaticChunkingStrategyResponseParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticChunkingStrategyResponseParamWithDefaults() *StaticChunkingStrategyResponseParam {
	this := StaticChunkingStrategyResponseParam{}
	return &this
}

// GetType returns the Type field value
func (o *StaticChunkingStrategyResponseParam) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StaticChunkingStrategyResponseParam) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StaticChunkingStrategyResponseParam) SetType(v string) {
	o.Type = v
}

// GetStatic returns the Static field value
func (o *StaticChunkingStrategyResponseParam) GetStatic() StaticChunkingStrategy {
	if o == nil {
		var ret StaticChunkingStrategy
		return ret
	}

	return o.Static
}

// GetStaticOk returns a tuple with the Static field value
// and a boolean to check if the value has been set.
func (o *StaticChunkingStrategyResponseParam) GetStaticOk() (*StaticChunkingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Static, true
}

// SetStatic sets field value
func (o *StaticChunkingStrategyResponseParam) SetStatic(v StaticChunkingStrategy) {
	o.Static = v
}

func (o StaticChunkingStrategyResponseParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticChunkingStrategyResponseParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["static"] = o.Static
	return toSerialize, nil
}

func (o *StaticChunkingStrategyResponseParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"static",
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

	varStaticChunkingStrategyResponseParam := _StaticChunkingStrategyResponseParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticChunkingStrategyResponseParam)

	if err != nil {
		return err
	}

	*o = StaticChunkingStrategyResponseParam(varStaticChunkingStrategyResponseParam)

	return err
}

type NullableStaticChunkingStrategyResponseParam struct {
	value *StaticChunkingStrategyResponseParam
	isSet bool
}

func (v NullableStaticChunkingStrategyResponseParam) Get() *StaticChunkingStrategyResponseParam {
	return v.value
}

func (v *NullableStaticChunkingStrategyResponseParam) Set(val *StaticChunkingStrategyResponseParam) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticChunkingStrategyResponseParam) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticChunkingStrategyResponseParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticChunkingStrategyResponseParam(val *StaticChunkingStrategyResponseParam) *NullableStaticChunkingStrategyResponseParam {
	return &NullableStaticChunkingStrategyResponseParam{value: val, isSet: true}
}

func (v NullableStaticChunkingStrategyResponseParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticChunkingStrategyResponseParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


