/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// CreateImageEditRequestModel The model to use for image generation. Only `dall-e-2` is supported at this time.
type CreateImageEditRequestModel struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateImageEditRequestModel) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreateImageEditRequestModel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CreateImageEditRequestModel) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCreateImageEditRequestModel struct {
	value *CreateImageEditRequestModel
	isSet bool
}

func (v NullableCreateImageEditRequestModel) Get() *CreateImageEditRequestModel {
	return v.value
}

func (v *NullableCreateImageEditRequestModel) Set(val *CreateImageEditRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageEditRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageEditRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageEditRequestModel(val *CreateImageEditRequestModel) *NullableCreateImageEditRequestModel {
	return &NullableCreateImageEditRequestModel{value: val, isSet: true}
}

func (v NullableCreateImageEditRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageEditRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


