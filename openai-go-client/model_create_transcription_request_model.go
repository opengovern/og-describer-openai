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


// CreateTranscriptionRequestModel ID of the model to use. Only `whisper-1` (which is powered by our open source Whisper V2 model) is currently available. 
type CreateTranscriptionRequestModel struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateTranscriptionRequestModel) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(CreateTranscriptionRequestModel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CreateTranscriptionRequestModel) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCreateTranscriptionRequestModel struct {
	value *CreateTranscriptionRequestModel
	isSet bool
}

func (v NullableCreateTranscriptionRequestModel) Get() *CreateTranscriptionRequestModel {
	return v.value
}

func (v *NullableCreateTranscriptionRequestModel) Set(val *CreateTranscriptionRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTranscriptionRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTranscriptionRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTranscriptionRequestModel(val *CreateTranscriptionRequestModel) *NullableCreateTranscriptionRequestModel {
	return &NullableCreateTranscriptionRequestModel{value: val, isSet: true}
}

func (v NullableCreateTranscriptionRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTranscriptionRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


