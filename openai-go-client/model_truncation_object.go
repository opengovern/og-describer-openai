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

// checks if the TruncationObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TruncationObject{}

// TruncationObject Controls for how a thread will be truncated prior to the run. Use this to control the intial context window of the run.
type TruncationObject struct {
	// The truncation strategy to use for the thread. The default is `auto`. If set to `last_messages`, the thread will be truncated to the n most recent messages in the thread. When set to `auto`, messages in the middle of the thread will be dropped to fit the context length of the model, `max_prompt_tokens`.
	Type string `json:"type"`
	// The number of most recent messages from the thread when constructing the context for the run.
	LastMessages NullableInt32 `json:"last_messages,omitempty"`
}

type _TruncationObject TruncationObject

// NewTruncationObject instantiates a new TruncationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTruncationObject(type_ string) *TruncationObject {
	this := TruncationObject{}
	this.Type = type_
	return &this
}

// NewTruncationObjectWithDefaults instantiates a new TruncationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTruncationObjectWithDefaults() *TruncationObject {
	this := TruncationObject{}
	return &this
}

// GetType returns the Type field value
func (o *TruncationObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TruncationObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TruncationObject) SetType(v string) {
	o.Type = v
}

// GetLastMessages returns the LastMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TruncationObject) GetLastMessages() int32 {
	if o == nil || IsNil(o.LastMessages.Get()) {
		var ret int32
		return ret
	}
	return *o.LastMessages.Get()
}

// GetLastMessagesOk returns a tuple with the LastMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TruncationObject) GetLastMessagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMessages.Get(), o.LastMessages.IsSet()
}

// HasLastMessages returns a boolean if a field has been set.
func (o *TruncationObject) HasLastMessages() bool {
	if o != nil && o.LastMessages.IsSet() {
		return true
	}

	return false
}

// SetLastMessages gets a reference to the given NullableInt32 and assigns it to the LastMessages field.
func (o *TruncationObject) SetLastMessages(v int32) {
	o.LastMessages.Set(&v)
}
// SetLastMessagesNil sets the value for LastMessages to be an explicit nil
func (o *TruncationObject) SetLastMessagesNil() {
	o.LastMessages.Set(nil)
}

// UnsetLastMessages ensures that no value is present for LastMessages, not even an explicit nil
func (o *TruncationObject) UnsetLastMessages() {
	o.LastMessages.Unset()
}

func (o TruncationObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TruncationObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.LastMessages.IsSet() {
		toSerialize["last_messages"] = o.LastMessages.Get()
	}
	return toSerialize, nil
}

func (o *TruncationObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varTruncationObject := _TruncationObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTruncationObject)

	if err != nil {
		return err
	}

	*o = TruncationObject(varTruncationObject)

	return err
}

type NullableTruncationObject struct {
	value *TruncationObject
	isSet bool
}

func (v NullableTruncationObject) Get() *TruncationObject {
	return v.value
}

func (v *NullableTruncationObject) Set(val *TruncationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTruncationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTruncationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTruncationObject(val *TruncationObject) *NullableTruncationObject {
	return &NullableTruncationObject{value: val, isSet: true}
}

func (v NullableTruncationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTruncationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


