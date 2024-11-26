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

// MessageDeltaObjectDeltaContentInner - struct for MessageDeltaObjectDeltaContentInner
type MessageDeltaObjectDeltaContentInner struct {
	MessageDeltaContentImageFileObject *MessageDeltaContentImageFileObject
	MessageDeltaContentImageUrlObject *MessageDeltaContentImageUrlObject
	MessageDeltaContentRefusalObject *MessageDeltaContentRefusalObject
	MessageDeltaContentTextObject *MessageDeltaContentTextObject
}

// MessageDeltaContentImageFileObjectAsMessageDeltaObjectDeltaContentInner is a convenience function that returns MessageDeltaContentImageFileObject wrapped in MessageDeltaObjectDeltaContentInner
func MessageDeltaContentImageFileObjectAsMessageDeltaObjectDeltaContentInner(v *MessageDeltaContentImageFileObject) MessageDeltaObjectDeltaContentInner {
	return MessageDeltaObjectDeltaContentInner{
		MessageDeltaContentImageFileObject: v,
	}
}

// MessageDeltaContentImageUrlObjectAsMessageDeltaObjectDeltaContentInner is a convenience function that returns MessageDeltaContentImageUrlObject wrapped in MessageDeltaObjectDeltaContentInner
func MessageDeltaContentImageUrlObjectAsMessageDeltaObjectDeltaContentInner(v *MessageDeltaContentImageUrlObject) MessageDeltaObjectDeltaContentInner {
	return MessageDeltaObjectDeltaContentInner{
		MessageDeltaContentImageUrlObject: v,
	}
}

// MessageDeltaContentRefusalObjectAsMessageDeltaObjectDeltaContentInner is a convenience function that returns MessageDeltaContentRefusalObject wrapped in MessageDeltaObjectDeltaContentInner
func MessageDeltaContentRefusalObjectAsMessageDeltaObjectDeltaContentInner(v *MessageDeltaContentRefusalObject) MessageDeltaObjectDeltaContentInner {
	return MessageDeltaObjectDeltaContentInner{
		MessageDeltaContentRefusalObject: v,
	}
}

// MessageDeltaContentTextObjectAsMessageDeltaObjectDeltaContentInner is a convenience function that returns MessageDeltaContentTextObject wrapped in MessageDeltaObjectDeltaContentInner
func MessageDeltaContentTextObjectAsMessageDeltaObjectDeltaContentInner(v *MessageDeltaContentTextObject) MessageDeltaObjectDeltaContentInner {
	return MessageDeltaObjectDeltaContentInner{
		MessageDeltaContentTextObject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageDeltaObjectDeltaContentInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MessageDeltaContentImageFileObject
	err = newStrictDecoder(data).Decode(&dst.MessageDeltaContentImageFileObject)
	if err == nil {
		jsonMessageDeltaContentImageFileObject, _ := json.Marshal(dst.MessageDeltaContentImageFileObject)
		if string(jsonMessageDeltaContentImageFileObject) == "{}" { // empty struct
			dst.MessageDeltaContentImageFileObject = nil
		} else {
			if err = validator.Validate(dst.MessageDeltaContentImageFileObject); err != nil {
				dst.MessageDeltaContentImageFileObject = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageDeltaContentImageFileObject = nil
	}

	// try to unmarshal data into MessageDeltaContentImageUrlObject
	err = newStrictDecoder(data).Decode(&dst.MessageDeltaContentImageUrlObject)
	if err == nil {
		jsonMessageDeltaContentImageUrlObject, _ := json.Marshal(dst.MessageDeltaContentImageUrlObject)
		if string(jsonMessageDeltaContentImageUrlObject) == "{}" { // empty struct
			dst.MessageDeltaContentImageUrlObject = nil
		} else {
			if err = validator.Validate(dst.MessageDeltaContentImageUrlObject); err != nil {
				dst.MessageDeltaContentImageUrlObject = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageDeltaContentImageUrlObject = nil
	}

	// try to unmarshal data into MessageDeltaContentRefusalObject
	err = newStrictDecoder(data).Decode(&dst.MessageDeltaContentRefusalObject)
	if err == nil {
		jsonMessageDeltaContentRefusalObject, _ := json.Marshal(dst.MessageDeltaContentRefusalObject)
		if string(jsonMessageDeltaContentRefusalObject) == "{}" { // empty struct
			dst.MessageDeltaContentRefusalObject = nil
		} else {
			if err = validator.Validate(dst.MessageDeltaContentRefusalObject); err != nil {
				dst.MessageDeltaContentRefusalObject = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageDeltaContentRefusalObject = nil
	}

	// try to unmarshal data into MessageDeltaContentTextObject
	err = newStrictDecoder(data).Decode(&dst.MessageDeltaContentTextObject)
	if err == nil {
		jsonMessageDeltaContentTextObject, _ := json.Marshal(dst.MessageDeltaContentTextObject)
		if string(jsonMessageDeltaContentTextObject) == "{}" { // empty struct
			dst.MessageDeltaContentTextObject = nil
		} else {
			if err = validator.Validate(dst.MessageDeltaContentTextObject); err != nil {
				dst.MessageDeltaContentTextObject = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageDeltaContentTextObject = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MessageDeltaContentImageFileObject = nil
		dst.MessageDeltaContentImageUrlObject = nil
		dst.MessageDeltaContentRefusalObject = nil
		dst.MessageDeltaContentTextObject = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MessageDeltaObjectDeltaContentInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MessageDeltaObjectDeltaContentInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageDeltaObjectDeltaContentInner) MarshalJSON() ([]byte, error) {
	if src.MessageDeltaContentImageFileObject != nil {
		return json.Marshal(&src.MessageDeltaContentImageFileObject)
	}

	if src.MessageDeltaContentImageUrlObject != nil {
		return json.Marshal(&src.MessageDeltaContentImageUrlObject)
	}

	if src.MessageDeltaContentRefusalObject != nil {
		return json.Marshal(&src.MessageDeltaContentRefusalObject)
	}

	if src.MessageDeltaContentTextObject != nil {
		return json.Marshal(&src.MessageDeltaContentTextObject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MessageDeltaObjectDeltaContentInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MessageDeltaContentImageFileObject != nil {
		return obj.MessageDeltaContentImageFileObject
	}

	if obj.MessageDeltaContentImageUrlObject != nil {
		return obj.MessageDeltaContentImageUrlObject
	}

	if obj.MessageDeltaContentRefusalObject != nil {
		return obj.MessageDeltaContentRefusalObject
	}

	if obj.MessageDeltaContentTextObject != nil {
		return obj.MessageDeltaContentTextObject
	}

	// all schemas are nil
	return nil
}

type NullableMessageDeltaObjectDeltaContentInner struct {
	value *MessageDeltaObjectDeltaContentInner
	isSet bool
}

func (v NullableMessageDeltaObjectDeltaContentInner) Get() *MessageDeltaObjectDeltaContentInner {
	return v.value
}

func (v *NullableMessageDeltaObjectDeltaContentInner) Set(val *MessageDeltaObjectDeltaContentInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDeltaObjectDeltaContentInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDeltaObjectDeltaContentInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDeltaObjectDeltaContentInner(val *MessageDeltaObjectDeltaContentInner) *NullableMessageDeltaObjectDeltaContentInner {
	return &NullableMessageDeltaObjectDeltaContentInner{value: val, isSet: true}
}

func (v NullableMessageDeltaObjectDeltaContentInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDeltaObjectDeltaContentInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


