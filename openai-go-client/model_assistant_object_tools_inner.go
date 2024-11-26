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

// AssistantObjectToolsInner - struct for AssistantObjectToolsInner
type AssistantObjectToolsInner struct {
	AssistantToolsCode *AssistantToolsCode
	AssistantToolsFileSearch *AssistantToolsFileSearch
	AssistantToolsFunction *AssistantToolsFunction
}

// AssistantToolsCodeAsAssistantObjectToolsInner is a convenience function that returns AssistantToolsCode wrapped in AssistantObjectToolsInner
func AssistantToolsCodeAsAssistantObjectToolsInner(v *AssistantToolsCode) AssistantObjectToolsInner {
	return AssistantObjectToolsInner{
		AssistantToolsCode: v,
	}
}

// AssistantToolsFileSearchAsAssistantObjectToolsInner is a convenience function that returns AssistantToolsFileSearch wrapped in AssistantObjectToolsInner
func AssistantToolsFileSearchAsAssistantObjectToolsInner(v *AssistantToolsFileSearch) AssistantObjectToolsInner {
	return AssistantObjectToolsInner{
		AssistantToolsFileSearch: v,
	}
}

// AssistantToolsFunctionAsAssistantObjectToolsInner is a convenience function that returns AssistantToolsFunction wrapped in AssistantObjectToolsInner
func AssistantToolsFunctionAsAssistantObjectToolsInner(v *AssistantToolsFunction) AssistantObjectToolsInner {
	return AssistantObjectToolsInner{
		AssistantToolsFunction: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssistantObjectToolsInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into AssistantToolsFileSearch
	err = newStrictDecoder(data).Decode(&dst.AssistantToolsFileSearch)
	if err == nil {
		jsonAssistantToolsFileSearch, _ := json.Marshal(dst.AssistantToolsFileSearch)
		if string(jsonAssistantToolsFileSearch) == "{}" { // empty struct
			dst.AssistantToolsFileSearch = nil
		} else {
			if err = validator.Validate(dst.AssistantToolsFileSearch); err != nil {
				dst.AssistantToolsFileSearch = nil
			} else {
				match++
			}
		}
	} else {
		dst.AssistantToolsFileSearch = nil
	}

	// try to unmarshal data into AssistantToolsFunction
	err = newStrictDecoder(data).Decode(&dst.AssistantToolsFunction)
	if err == nil {
		jsonAssistantToolsFunction, _ := json.Marshal(dst.AssistantToolsFunction)
		if string(jsonAssistantToolsFunction) == "{}" { // empty struct
			dst.AssistantToolsFunction = nil
		} else {
			if err = validator.Validate(dst.AssistantToolsFunction); err != nil {
				dst.AssistantToolsFunction = nil
			} else {
				match++
			}
		}
	} else {
		dst.AssistantToolsFunction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AssistantToolsCode = nil
		dst.AssistantToolsFileSearch = nil
		dst.AssistantToolsFunction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AssistantObjectToolsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AssistantObjectToolsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssistantObjectToolsInner) MarshalJSON() ([]byte, error) {
	if src.AssistantToolsCode != nil {
		return json.Marshal(&src.AssistantToolsCode)
	}

	if src.AssistantToolsFileSearch != nil {
		return json.Marshal(&src.AssistantToolsFileSearch)
	}

	if src.AssistantToolsFunction != nil {
		return json.Marshal(&src.AssistantToolsFunction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssistantObjectToolsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AssistantToolsCode != nil {
		return obj.AssistantToolsCode
	}

	if obj.AssistantToolsFileSearch != nil {
		return obj.AssistantToolsFileSearch
	}

	if obj.AssistantToolsFunction != nil {
		return obj.AssistantToolsFunction
	}

	// all schemas are nil
	return nil
}

type NullableAssistantObjectToolsInner struct {
	value *AssistantObjectToolsInner
	isSet bool
}

func (v NullableAssistantObjectToolsInner) Get() *AssistantObjectToolsInner {
	return v.value
}

func (v *NullableAssistantObjectToolsInner) Set(val *AssistantObjectToolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantObjectToolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantObjectToolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantObjectToolsInner(val *AssistantObjectToolsInner) *NullableAssistantObjectToolsInner {
	return &NullableAssistantObjectToolsInner{value: val, isSet: true}
}

func (v NullableAssistantObjectToolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantObjectToolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


