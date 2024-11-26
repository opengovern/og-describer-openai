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

// checks if the RealtimeServerEventErrorError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventErrorError{}

// RealtimeServerEventErrorError Details of the error.
type RealtimeServerEventErrorError struct {
	// The type of error (e.g., \"invalid_request_error\", \"server_error\").
	Type *string `json:"type,omitempty"`
	// Error code, if any.
	Code *string `json:"code,omitempty"`
	// A human-readable error message.
	Message *string `json:"message,omitempty"`
	// Parameter related to the error, if any.
	Param *string `json:"param,omitempty"`
	// The event_id of the client event that caused the error, if applicable.
	EventId *string `json:"event_id,omitempty"`
}

// NewRealtimeServerEventErrorError instantiates a new RealtimeServerEventErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventErrorError() *RealtimeServerEventErrorError {
	this := RealtimeServerEventErrorError{}
	return &this
}

// NewRealtimeServerEventErrorErrorWithDefaults instantiates a new RealtimeServerEventErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventErrorErrorWithDefaults() *RealtimeServerEventErrorError {
	this := RealtimeServerEventErrorError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RealtimeServerEventErrorError) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventErrorError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RealtimeServerEventErrorError) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RealtimeServerEventErrorError) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RealtimeServerEventErrorError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventErrorError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RealtimeServerEventErrorError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RealtimeServerEventErrorError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RealtimeServerEventErrorError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventErrorError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RealtimeServerEventErrorError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RealtimeServerEventErrorError) SetMessage(v string) {
	o.Message = &v
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *RealtimeServerEventErrorError) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventErrorError) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *RealtimeServerEventErrorError) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *RealtimeServerEventErrorError) SetParam(v string) {
	o.Param = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *RealtimeServerEventErrorError) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventErrorError) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *RealtimeServerEventErrorError) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *RealtimeServerEventErrorError) SetEventId(v string) {
	o.EventId = &v
}

func (o RealtimeServerEventErrorError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventErrorError) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	return toSerialize, nil
}

type NullableRealtimeServerEventErrorError struct {
	value *RealtimeServerEventErrorError
	isSet bool
}

func (v NullableRealtimeServerEventErrorError) Get() *RealtimeServerEventErrorError {
	return v.value
}

func (v *NullableRealtimeServerEventErrorError) Set(val *RealtimeServerEventErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventErrorError(val *RealtimeServerEventErrorError) *NullableRealtimeServerEventErrorError {
	return &NullableRealtimeServerEventErrorError{value: val, isSet: true}
}

func (v NullableRealtimeServerEventErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


