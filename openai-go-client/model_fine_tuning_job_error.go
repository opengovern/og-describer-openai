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

// checks if the FineTuningJobError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FineTuningJobError{}

// FineTuningJobError For fine-tuning jobs that have `failed`, this will contain more information on the cause of the failure.
type FineTuningJobError struct {
	// A machine-readable error code.
	Code string `json:"code"`
	// A human-readable error message.
	Message string `json:"message"`
	// The parameter that was invalid, usually `training_file` or `validation_file`. This field will be null if the failure was not parameter-specific.
	Param NullableString `json:"param"`
}

type _FineTuningJobError FineTuningJobError

// NewFineTuningJobError instantiates a new FineTuningJobError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFineTuningJobError(code string, message string, param NullableString) *FineTuningJobError {
	this := FineTuningJobError{}
	this.Code = code
	this.Message = message
	this.Param = param
	return &this
}

// NewFineTuningJobErrorWithDefaults instantiates a new FineTuningJobError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFineTuningJobErrorWithDefaults() *FineTuningJobError {
	this := FineTuningJobError{}
	return &this
}

// GetCode returns the Code field value
func (o *FineTuningJobError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *FineTuningJobError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *FineTuningJobError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *FineTuningJobError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FineTuningJobError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FineTuningJobError) SetMessage(v string) {
	o.Message = v
}

// GetParam returns the Param field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FineTuningJobError) GetParam() string {
	if o == nil || o.Param.Get() == nil {
		var ret string
		return ret
	}

	return *o.Param.Get()
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FineTuningJobError) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Param.Get(), o.Param.IsSet()
}

// SetParam sets field value
func (o *FineTuningJobError) SetParam(v string) {
	o.Param.Set(&v)
}

func (o FineTuningJobError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FineTuningJobError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["param"] = o.Param.Get()
	return toSerialize, nil
}

func (o *FineTuningJobError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
		"param",
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

	varFineTuningJobError := _FineTuningJobError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFineTuningJobError)

	if err != nil {
		return err
	}

	*o = FineTuningJobError(varFineTuningJobError)

	return err
}

type NullableFineTuningJobError struct {
	value *FineTuningJobError
	isSet bool
}

func (v NullableFineTuningJobError) Get() *FineTuningJobError {
	return v.value
}

func (v *NullableFineTuningJobError) Set(val *FineTuningJobError) {
	v.value = val
	v.isSet = true
}

func (v NullableFineTuningJobError) IsSet() bool {
	return v.isSet
}

func (v *NullableFineTuningJobError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFineTuningJobError(val *FineTuningJobError) *NullableFineTuningJobError {
	return &NullableFineTuningJobError{value: val, isSet: true}
}

func (v NullableFineTuningJobError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFineTuningJobError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


