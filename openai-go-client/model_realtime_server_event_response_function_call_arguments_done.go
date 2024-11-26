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

// checks if the RealtimeServerEventResponseFunctionCallArgumentsDone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeServerEventResponseFunctionCallArgumentsDone{}

// RealtimeServerEventResponseFunctionCallArgumentsDone Returned when the model-generated function call arguments are done streaming. Also emitted when a Response is interrupted, incomplete, or cancelled.
type RealtimeServerEventResponseFunctionCallArgumentsDone struct {
	// The unique ID of the server event.
	EventId string `json:"event_id"`
	// The event type, must be \"response.function_call_arguments.done\".
	Type string `json:"type"`
	// The ID of the response.
	ResponseId string `json:"response_id"`
	// The ID of the function call item.
	ItemId string `json:"item_id"`
	// The index of the output item in the response.
	OutputIndex int32 `json:"output_index"`
	// The ID of the function call.
	CallId string `json:"call_id"`
	// The final arguments as a JSON string.
	Arguments string `json:"arguments"`
}

type _RealtimeServerEventResponseFunctionCallArgumentsDone RealtimeServerEventResponseFunctionCallArgumentsDone

// NewRealtimeServerEventResponseFunctionCallArgumentsDone instantiates a new RealtimeServerEventResponseFunctionCallArgumentsDone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeServerEventResponseFunctionCallArgumentsDone(eventId string, type_ string, responseId string, itemId string, outputIndex int32, callId string, arguments string) *RealtimeServerEventResponseFunctionCallArgumentsDone {
	this := RealtimeServerEventResponseFunctionCallArgumentsDone{}
	this.EventId = eventId
	this.Type = type_
	this.ResponseId = responseId
	this.ItemId = itemId
	this.OutputIndex = outputIndex
	this.CallId = callId
	this.Arguments = arguments
	return &this
}

// NewRealtimeServerEventResponseFunctionCallArgumentsDoneWithDefaults instantiates a new RealtimeServerEventResponseFunctionCallArgumentsDone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeServerEventResponseFunctionCallArgumentsDoneWithDefaults() *RealtimeServerEventResponseFunctionCallArgumentsDone {
	this := RealtimeServerEventResponseFunctionCallArgumentsDone{}
	return &this
}

// GetEventId returns the EventId field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetType(v string) {
	o.Type = v
}

// GetResponseId returns the ResponseId field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetResponseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetResponseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseId, true
}

// SetResponseId sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetResponseId(v string) {
	o.ResponseId = v
}

// GetItemId returns the ItemId field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetItemId(v string) {
	o.ItemId = v
}

// GetOutputIndex returns the OutputIndex field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetOutputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutputIndex
}

// GetOutputIndexOk returns a tuple with the OutputIndex field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetOutputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputIndex, true
}

// SetOutputIndex sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetOutputIndex(v int32) {
	o.OutputIndex = v
}

// GetCallId returns the CallId field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallId, true
}

// SetCallId sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetCallId(v string) {
	o.CallId = v
}

// GetArguments returns the Arguments field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetArguments() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) GetArgumentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value
func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) SetArguments(v string) {
	o.Arguments = v
}

func (o RealtimeServerEventResponseFunctionCallArgumentsDone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeServerEventResponseFunctionCallArgumentsDone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_id"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["response_id"] = o.ResponseId
	toSerialize["item_id"] = o.ItemId
	toSerialize["output_index"] = o.OutputIndex
	toSerialize["call_id"] = o.CallId
	toSerialize["arguments"] = o.Arguments
	return toSerialize, nil
}

func (o *RealtimeServerEventResponseFunctionCallArgumentsDone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_id",
		"type",
		"response_id",
		"item_id",
		"output_index",
		"call_id",
		"arguments",
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

	varRealtimeServerEventResponseFunctionCallArgumentsDone := _RealtimeServerEventResponseFunctionCallArgumentsDone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRealtimeServerEventResponseFunctionCallArgumentsDone)

	if err != nil {
		return err
	}

	*o = RealtimeServerEventResponseFunctionCallArgumentsDone(varRealtimeServerEventResponseFunctionCallArgumentsDone)

	return err
}

type NullableRealtimeServerEventResponseFunctionCallArgumentsDone struct {
	value *RealtimeServerEventResponseFunctionCallArgumentsDone
	isSet bool
}

func (v NullableRealtimeServerEventResponseFunctionCallArgumentsDone) Get() *RealtimeServerEventResponseFunctionCallArgumentsDone {
	return v.value
}

func (v *NullableRealtimeServerEventResponseFunctionCallArgumentsDone) Set(val *RealtimeServerEventResponseFunctionCallArgumentsDone) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeServerEventResponseFunctionCallArgumentsDone) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeServerEventResponseFunctionCallArgumentsDone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeServerEventResponseFunctionCallArgumentsDone(val *RealtimeServerEventResponseFunctionCallArgumentsDone) *NullableRealtimeServerEventResponseFunctionCallArgumentsDone {
	return &NullableRealtimeServerEventResponseFunctionCallArgumentsDone{value: val, isSet: true}
}

func (v NullableRealtimeServerEventResponseFunctionCallArgumentsDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeServerEventResponseFunctionCallArgumentsDone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


