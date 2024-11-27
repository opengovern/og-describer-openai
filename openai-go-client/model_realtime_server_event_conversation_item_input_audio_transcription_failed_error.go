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

// checks if the RealtimeServerEventConversationItemInputAudioTranscriptionFailedError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventConversationItemInputAudioTranscriptionFailedError{}

// RealtimeServerEventConversationItemInputAudioTranscriptionFailedError Details of the transcription error.
type RealtimeServerEventConversationItemInputAudioTranscriptionFailedError struct {
	// The type of error.
	Type *string `json:"type,omitempty"`
	// Error code, if any.
	Code *string `json:"code,omitempty"`
	// A human-readable error message.
	Message *string `json:"message,omitempty"`
	// Parameter related to the error, if any.
	Param *string `json:"param,omitempty"`
}

// NewRealtimeServerEventConversationItemInputAudioTranscriptionFailedError instantiates a new RealtimeServerEventConversationItemInputAudioTranscriptionFailedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventConversationItemInputAudioTranscriptionFailedError() *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError {
	this := RealtimeServerEventConversationItemInputAudioTranscriptionFailedError{}
	return &this
}

// NewRealtimeServerEventConversationItemInputAudioTranscriptionFailedErrorWithDefaults instantiates a new RealtimeServerEventConversationItemInputAudioTranscriptionFailedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventConversationItemInputAudioTranscriptionFailedErrorWithDefaults() *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError {
	this := RealtimeServerEventConversationItemInputAudioTranscriptionFailedError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) SetMessage(v string) {
	o.Message = &v
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) SetParam(v string) {
	o.Param = &v
}

func (o RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	return toSerialize, nil
}

type NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError struct {
	value *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError
	isSet bool
}

func (v NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError) Get() *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError {
	return v.value
}

func (v *NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError) Set(val *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError(val *RealtimeServerEventConversationItemInputAudioTranscriptionFailedError) *NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError {
	return &NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError{value: val, isSet: true}
}

func (v NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventConversationItemInputAudioTranscriptionFailedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

