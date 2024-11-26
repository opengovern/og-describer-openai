/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// CreateMessageRequestAttachmentsInnerToolsInner - struct for CreateMessageRequestAttachmentsInnerToolsInner
type CreateMessageRequestAttachmentsInnerToolsInner struct {
	AssistantToolsCode *AssistantToolsCode
	AssistantToolsFileSearchTypeOnly *AssistantToolsFileSearchTypeOnly
}

// AssistantToolsCodeAsCreateMessageRequestAttachmentsInnerToolsInner is a convenience function that returns AssistantToolsCode wrapped in CreateMessageRequestAttachmentsInnerToolsInner
func AssistantToolsCodeAsCreateMessageRequestAttachmentsInnerToolsInner(v *AssistantToolsCode) CreateMessageRequestAttachmentsInnerToolsInner {
	return CreateMessageRequestAttachmentsInnerToolsInner{
		AssistantToolsCode: v,
	}
}

// AssistantToolsFileSearchTypeOnlyAsCreateMessageRequestAttachmentsInnerToolsInner is a convenience function that returns AssistantToolsFileSearchTypeOnly wrapped in CreateMessageRequestAttachmentsInnerToolsInner
func AssistantToolsFileSearchTypeOnlyAsCreateMessageRequestAttachmentsInnerToolsInner(v *AssistantToolsFileSearchTypeOnly) CreateMessageRequestAttachmentsInnerToolsInner {
	return CreateMessageRequestAttachmentsInnerToolsInner{
		AssistantToolsFileSearchTypeOnly: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateMessageRequestAttachmentsInnerToolsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AssistantToolsCode
	err = newStrictDecoder(data).Decode(&dst.AssistantToolsCode)
	if err == nil {
		jsonAssistantToolsCode, _ := json.Marshal(dst.AssistantToolsCode)
		if string(jsonAssistantToolsCode) == "{}" { // empty struct
			dst.AssistantToolsCode = nil
		} else {
			if err = validator.Validate(dst.AssistantToolsCode); err != nil {
				dst.AssistantToolsCode = nil
			} else {
				match++
			}
		}
	} else {
		dst.AssistantToolsCode = nil
	}

	// try to unmarshal data into AssistantToolsFileSearchTypeOnly
	err = newStrictDecoder(data).Decode(&dst.AssistantToolsFileSearchTypeOnly)
	if err == nil {
		jsonAssistantToolsFileSearchTypeOnly, _ := json.Marshal(dst.AssistantToolsFileSearchTypeOnly)
		if string(jsonAssistantToolsFileSearchTypeOnly) == "{}" { // empty struct
			dst.AssistantToolsFileSearchTypeOnly = nil
		} else {
			if err = validator.Validate(dst.AssistantToolsFileSearchTypeOnly); err != nil {
				dst.AssistantToolsFileSearchTypeOnly = nil
			} else {
				match++
			}
		}
	} else {
		dst.AssistantToolsFileSearchTypeOnly = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AssistantToolsCode = nil
		dst.AssistantToolsFileSearchTypeOnly = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateMessageRequestAttachmentsInnerToolsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateMessageRequestAttachmentsInnerToolsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateMessageRequestAttachmentsInnerToolsInner) MarshalJSON() ([]byte, error) {
	if src.AssistantToolsCode != nil {
		return json.Marshal(&src.AssistantToolsCode)
	}

	if src.AssistantToolsFileSearchTypeOnly != nil {
		return json.Marshal(&src.AssistantToolsFileSearchTypeOnly)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateMessageRequestAttachmentsInnerToolsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AssistantToolsCode != nil {
		return obj.AssistantToolsCode
	}

	if obj.AssistantToolsFileSearchTypeOnly != nil {
		return obj.AssistantToolsFileSearchTypeOnly
	}

	// all schemas are nil
	return nil
}

type NullableCreateMessageRequestAttachmentsInnerToolsInner struct {
	value *CreateMessageRequestAttachmentsInnerToolsInner
	isSet bool
}

func (v NullableCreateMessageRequestAttachmentsInnerToolsInner) Get() *CreateMessageRequestAttachmentsInnerToolsInner {
	return v.value
}

func (v *NullableCreateMessageRequestAttachmentsInnerToolsInner) Set(val *CreateMessageRequestAttachmentsInnerToolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageRequestAttachmentsInnerToolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageRequestAttachmentsInnerToolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageRequestAttachmentsInnerToolsInner(val *CreateMessageRequestAttachmentsInnerToolsInner) *NullableCreateMessageRequestAttachmentsInnerToolsInner {
	return &NullableCreateMessageRequestAttachmentsInnerToolsInner{value: val, isSet: true}
}

func (v NullableCreateMessageRequestAttachmentsInnerToolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageRequestAttachmentsInnerToolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


