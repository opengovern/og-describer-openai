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

// checks if the RealtimeServerEventConversationCreatedConversation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventConversationCreatedConversation{}

// RealtimeServerEventConversationCreatedConversation The conversation resource.
type RealtimeServerEventConversationCreatedConversation struct {
	// The unique ID of the conversation.
	Id *string `json:"id,omitempty"`
	// The object type, must be \"realtime.conversation\".
	Object *string `json:"object,omitempty"`
}

// NewRealtimeServerEventConversationCreatedConversation instantiates a new RealtimeServerEventConversationCreatedConversation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventConversationCreatedConversation() *RealtimeServerEventConversationCreatedConversation {
	this := RealtimeServerEventConversationCreatedConversation{}
	return &this
}

// NewRealtimeServerEventConversationCreatedConversationWithDefaults instantiates a new RealtimeServerEventConversationCreatedConversation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventConversationCreatedConversationWithDefaults() *RealtimeServerEventConversationCreatedConversation {
	this := RealtimeServerEventConversationCreatedConversation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RealtimeServerEventConversationCreatedConversation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationCreatedConversation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RealtimeServerEventConversationCreatedConversation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RealtimeServerEventConversationCreatedConversation) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *RealtimeServerEventConversationCreatedConversation) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationCreatedConversation) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *RealtimeServerEventConversationCreatedConversation) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *RealtimeServerEventConversationCreatedConversation) SetObject(v string) {
	o.Object = &v
}

func (o RealtimeServerEventConversationCreatedConversation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventConversationCreatedConversation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableRealtimeServerEventConversationCreatedConversation struct {
	value *RealtimeServerEventConversationCreatedConversation
	isSet bool
}

func (v NullableRealtimeServerEventConversationCreatedConversation) Get() *RealtimeServerEventConversationCreatedConversation {
	return v.value
}

func (v *NullableRealtimeServerEventConversationCreatedConversation) Set(val *RealtimeServerEventConversationCreatedConversation) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventConversationCreatedConversation) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventConversationCreatedConversation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventConversationCreatedConversation(val *RealtimeServerEventConversationCreatedConversation) *NullableRealtimeServerEventConversationCreatedConversation {
	return &NullableRealtimeServerEventConversationCreatedConversation{value: val, isSet: true}
}

func (v NullableRealtimeServerEventConversationCreatedConversation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventConversationCreatedConversation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


