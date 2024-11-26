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

// AudioResponseFormat The format of the output, in one of these options: `json`, `text`, `srt`, `verbose_json`, or `vtt`. 
type AudioResponseFormat string

// List of AudioResponseFormat
const (
	JSON AudioResponseFormat = "json"
	TEXT AudioResponseFormat = "text"
	SRT AudioResponseFormat = "srt"
	VERBOSE_JSON AudioResponseFormat = "verbose_json"
	VTT AudioResponseFormat = "vtt"
)

// All allowed values of AudioResponseFormat enum
var AllowedAudioResponseFormatEnumValues = []AudioResponseFormat{
	"json",
	"text",
	"srt",
	"verbose_json",
	"vtt",
}

func (v *AudioResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AudioResponseFormat(value)
	for _, existing := range AllowedAudioResponseFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AudioResponseFormat", value)
}

// NewAudioResponseFormatFromValue returns a pointer to a valid AudioResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudioResponseFormatFromValue(v string) (*AudioResponseFormat, error) {
	ev := AudioResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudioResponseFormat: valid values are %v", v, AllowedAudioResponseFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudioResponseFormat) IsValid() bool {
	for _, existing := range AllowedAudioResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AudioResponseFormat value
func (v AudioResponseFormat) Ptr() *AudioResponseFormat {
	return &v
}

type NullableAudioResponseFormat struct {
	value *AudioResponseFormat
	isSet bool
}

func (v NullableAudioResponseFormat) Get() *AudioResponseFormat {
	return v.value
}

func (v *NullableAudioResponseFormat) Set(val *AudioResponseFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableAudioResponseFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableAudioResponseFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudioResponseFormat(val *AudioResponseFormat) *NullableAudioResponseFormat {
	return &NullableAudioResponseFormat{value: val, isSet: true}
}

func (v NullableAudioResponseFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudioResponseFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

