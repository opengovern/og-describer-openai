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

// checks if the PredictionContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredictionContent{}

// PredictionContent Static predicted output content, such as the content of a text file that is being regenerated. 
type PredictionContent struct {
	// The type of the predicted content you want to provide. This type is currently always `content`. 
	Type string `json:"type"`
	Content PredictionContentContent `json:"content"`
}

type _PredictionContent PredictionContent

// NewPredictionContent instantiates a new PredictionContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionContent(type_ string, content PredictionContentContent) *PredictionContent {
	this := PredictionContent{}
	this.Type = type_
	this.Content = content
	return &this
}

// NewPredictionContentWithDefaults instantiates a new PredictionContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionContentWithDefaults() *PredictionContent {
	this := PredictionContent{}
	return &this
}

// GetType returns the Type field value
func (o *PredictionContent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PredictionContent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PredictionContent) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *PredictionContent) GetContent() PredictionContentContent {
	if o == nil {
		var ret PredictionContentContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PredictionContent) GetContentOk() (*PredictionContentContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *PredictionContent) SetContent(v PredictionContentContent) {
	o.Content = v
}

func (o PredictionContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictionContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *PredictionContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"content",
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

	varPredictionContent := _PredictionContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPredictionContent)

	if err != nil {
		return err
	}

	*o = PredictionContent(varPredictionContent)

	return err
}

type NullablePredictionContent struct {
	value *PredictionContent
	isSet bool
}

func (v NullablePredictionContent) Get() *PredictionContent {
	return v.value
}

func (v *NullablePredictionContent) Set(val *PredictionContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionContent(val *PredictionContent) *NullablePredictionContent {
	return &NullablePredictionContent{value: val, isSet: true}
}

func (v NullablePredictionContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

