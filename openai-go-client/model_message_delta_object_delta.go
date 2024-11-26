/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MessageDeltaObjectDelta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDeltaObjectDelta{}

// MessageDeltaObjectDelta The delta containing the fields that have changed on the Message.
type MessageDeltaObjectDelta struct {
	// The entity that produced the message. One of `user` or `assistant`.
	Role *string `json:"role,omitempty"`
	// The content of the message in array of text and/or images.
	Content []MessageDeltaObjectDeltaContentInner `json:"content,omitempty"`
}

// NewMessageDeltaObjectDelta instantiates a new MessageDeltaObjectDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDeltaObjectDelta() *MessageDeltaObjectDelta {
	this := MessageDeltaObjectDelta{}
	return &this
}

// NewMessageDeltaObjectDeltaWithDefaults instantiates a new MessageDeltaObjectDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDeltaObjectDeltaWithDefaults() *MessageDeltaObjectDelta {
	this := MessageDeltaObjectDelta{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MessageDeltaObjectDelta) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeltaObjectDelta) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MessageDeltaObjectDelta) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MessageDeltaObjectDelta) SetRole(v string) {
	o.Role = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *MessageDeltaObjectDelta) GetContent() []MessageDeltaObjectDeltaContentInner {
	if o == nil || IsNil(o.Content) {
		var ret []MessageDeltaObjectDeltaContentInner
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeltaObjectDelta) GetContentOk() ([]MessageDeltaObjectDeltaContentInner, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MessageDeltaObjectDelta) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []MessageDeltaObjectDeltaContentInner and assigns it to the Content field.
func (o *MessageDeltaObjectDelta) SetContent(v []MessageDeltaObjectDeltaContentInner) {
	o.Content = v
}

func (o MessageDeltaObjectDelta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDeltaObjectDelta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableMessageDeltaObjectDelta struct {
	value *MessageDeltaObjectDelta
	isSet bool
}

func (v NullableMessageDeltaObjectDelta) Get() *MessageDeltaObjectDelta {
	return v.value
}

func (v *NullableMessageDeltaObjectDelta) Set(val *MessageDeltaObjectDelta) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDeltaObjectDelta) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDeltaObjectDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDeltaObjectDelta(val *MessageDeltaObjectDelta) *NullableMessageDeltaObjectDelta {
	return &NullableMessageDeltaObjectDelta{value: val, isSet: true}
}

func (v NullableMessageDeltaObjectDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDeltaObjectDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


