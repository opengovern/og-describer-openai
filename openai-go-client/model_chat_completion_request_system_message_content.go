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

// ChatCompletionRequestSystemMessageContent - The contents of the system message.
type ChatCompletionRequestSystemMessageContent struct {
	ArrayOfChatCompletionRequestSystemMessageContentPart *[]ChatCompletionRequestSystemMessageContentPart
	String *string
}

// []ChatCompletionRequestSystemMessageContentPartAsChatCompletionRequestSystemMessageContent is a convenience function that returns []ChatCompletionRequestSystemMessageContentPart wrapped in ChatCompletionRequestSystemMessageContent
func ArrayOfChatCompletionRequestSystemMessageContentPartAsChatCompletionRequestSystemMessageContent(v *[]ChatCompletionRequestSystemMessageContentPart) ChatCompletionRequestSystemMessageContent {
	return ChatCompletionRequestSystemMessageContent{
		ArrayOfChatCompletionRequestSystemMessageContentPart: v,
	}
}

// stringAsChatCompletionRequestSystemMessageContent is a convenience function that returns string wrapped in ChatCompletionRequestSystemMessageContent
func StringAsChatCompletionRequestSystemMessageContent(v *string) ChatCompletionRequestSystemMessageContent {
	return ChatCompletionRequestSystemMessageContent{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChatCompletionRequestSystemMessageContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfChatCompletionRequestSystemMessageContentPart
	err = newStrictDecoder(data).Decode(&dst.ArrayOfChatCompletionRequestSystemMessageContentPart)
	if err == nil {
		jsonArrayOfChatCompletionRequestSystemMessageContentPart, _ := json.Marshal(dst.ArrayOfChatCompletionRequestSystemMessageContentPart)
		if string(jsonArrayOfChatCompletionRequestSystemMessageContentPart) == "{}" { // empty struct
			dst.ArrayOfChatCompletionRequestSystemMessageContentPart = nil
		} else {
			if err = validator.Validate(dst.ArrayOfChatCompletionRequestSystemMessageContentPart); err != nil {
				dst.ArrayOfChatCompletionRequestSystemMessageContentPart = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfChatCompletionRequestSystemMessageContentPart = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfChatCompletionRequestSystemMessageContentPart = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChatCompletionRequestSystemMessageContent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChatCompletionRequestSystemMessageContent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChatCompletionRequestSystemMessageContent) MarshalJSON() ([]byte, error) {
	if src.ArrayOfChatCompletionRequestSystemMessageContentPart != nil {
		return json.Marshal(&src.ArrayOfChatCompletionRequestSystemMessageContentPart)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChatCompletionRequestSystemMessageContent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfChatCompletionRequestSystemMessageContentPart != nil {
		return obj.ArrayOfChatCompletionRequestSystemMessageContentPart
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableChatCompletionRequestSystemMessageContent struct {
	value *ChatCompletionRequestSystemMessageContent
	isSet bool
}

func (v NullableChatCompletionRequestSystemMessageContent) Get() *ChatCompletionRequestSystemMessageContent {
	return v.value
}

func (v *NullableChatCompletionRequestSystemMessageContent) Set(val *ChatCompletionRequestSystemMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestSystemMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestSystemMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestSystemMessageContent(val *ChatCompletionRequestSystemMessageContent) *NullableChatCompletionRequestSystemMessageContent {
	return &NullableChatCompletionRequestSystemMessageContent{value: val, isSet: true}
}

func (v NullableChatCompletionRequestSystemMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestSystemMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


