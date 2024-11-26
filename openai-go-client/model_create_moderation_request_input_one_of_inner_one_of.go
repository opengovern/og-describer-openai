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

// checks if the CreateModerationRequestInputOneOfInnerOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateModerationRequestInputOneOfInnerOneOf{}

// CreateModerationRequestInputOneOfInnerOneOf An object describing an image to classify.
type CreateModerationRequestInputOneOfInnerOneOf struct {
	// Always `image_url`.
	Type string `json:"type"`
	ImageUrl CreateModerationRequestInputOneOfInnerOneOfImageUrl `json:"image_url"`
}

type _CreateModerationRequestInputOneOfInnerOneOf CreateModerationRequestInputOneOfInnerOneOf

// NewCreateModerationRequestInputOneOfInnerOneOf instantiates a new CreateModerationRequestInputOneOfInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateModerationRequestInputOneOfInnerOneOf(type_ string, imageUrl CreateModerationRequestInputOneOfInnerOneOfImageUrl) *CreateModerationRequestInputOneOfInnerOneOf {
	this := CreateModerationRequestInputOneOfInnerOneOf{}
	this.Type = type_
	this.ImageUrl = imageUrl
	return &this
}

// NewCreateModerationRequestInputOneOfInnerOneOfWithDefaults instantiates a new CreateModerationRequestInputOneOfInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateModerationRequestInputOneOfInnerOneOfWithDefaults() *CreateModerationRequestInputOneOfInnerOneOf {
	this := CreateModerationRequestInputOneOfInnerOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *CreateModerationRequestInputOneOfInnerOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateModerationRequestInputOneOfInnerOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateModerationRequestInputOneOfInnerOneOf) SetType(v string) {
	o.Type = v
}

// GetImageUrl returns the ImageUrl field value
func (o *CreateModerationRequestInputOneOfInnerOneOf) GetImageUrl() CreateModerationRequestInputOneOfInnerOneOfImageUrl {
	if o == nil {
		var ret CreateModerationRequestInputOneOfInnerOneOfImageUrl
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *CreateModerationRequestInputOneOfInnerOneOf) GetImageUrlOk() (*CreateModerationRequestInputOneOfInnerOneOfImageUrl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *CreateModerationRequestInputOneOfInnerOneOf) SetImageUrl(v CreateModerationRequestInputOneOfInnerOneOfImageUrl) {
	o.ImageUrl = v
}

func (o CreateModerationRequestInputOneOfInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateModerationRequestInputOneOfInnerOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["image_url"] = o.ImageUrl
	return toSerialize, nil
}

func (o *CreateModerationRequestInputOneOfInnerOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"image_url",
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

	varCreateModerationRequestInputOneOfInnerOneOf := _CreateModerationRequestInputOneOfInnerOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateModerationRequestInputOneOfInnerOneOf)

	if err != nil {
		return err
	}

	*o = CreateModerationRequestInputOneOfInnerOneOf(varCreateModerationRequestInputOneOfInnerOneOf)

	return err
}

type NullableCreateModerationRequestInputOneOfInnerOneOf struct {
	value *CreateModerationRequestInputOneOfInnerOneOf
	isSet bool
}

func (v NullableCreateModerationRequestInputOneOfInnerOneOf) Get() *CreateModerationRequestInputOneOfInnerOneOf {
	return v.value
}

func (v *NullableCreateModerationRequestInputOneOfInnerOneOf) Set(val *CreateModerationRequestInputOneOfInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateModerationRequestInputOneOfInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateModerationRequestInputOneOfInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateModerationRequestInputOneOfInnerOneOf(val *CreateModerationRequestInputOneOfInnerOneOf) *NullableCreateModerationRequestInputOneOfInnerOneOf {
	return &NullableCreateModerationRequestInputOneOfInnerOneOf{value: val, isSet: true}
}

func (v NullableCreateModerationRequestInputOneOfInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateModerationRequestInputOneOfInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


