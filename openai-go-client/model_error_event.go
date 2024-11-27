/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorEvent{}

// ErrorEvent Occurs when an [error](/docs/guides/error-codes#api-errors) occurs. This can happen due to an internal server error or a timeout.
type ErrorEvent struct {
	Event string `json:"event"`
	Data Error `json:"data"`
}

type _ErrorEvent ErrorEvent

// NewErrorEvent instantiates a new ErrorEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorEvent(event string, data Error) *ErrorEvent {
	this := ErrorEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewErrorEventWithDefaults instantiates a new ErrorEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorEventWithDefaults() *ErrorEvent {
	this := ErrorEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *ErrorEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *ErrorEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *ErrorEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *ErrorEvent) GetData() Error {
	if o == nil {
		var ret Error
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ErrorEvent) GetDataOk() (*Error, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ErrorEvent) SetData(v Error) {
	o.Data = v
}

func (o ErrorEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ErrorEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorEvent := _ErrorEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorEvent)

	if err != nil {
		return err
	}

	*o = ErrorEvent(varErrorEvent)

	return err
}

type NullableErrorEvent struct {
	value *ErrorEvent
	isSet bool
}

func (v NullableErrorEvent) Get() *ErrorEvent {
	return v.value
}

func (v *NullableErrorEvent) Set(val *ErrorEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorEvent(val *ErrorEvent) *NullableErrorEvent {
	return &NullableErrorEvent{value: val, isSet: true}
}

func (v NullableErrorEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

