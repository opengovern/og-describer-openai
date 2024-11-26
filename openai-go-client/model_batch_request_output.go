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

// checks if the BatchRequestOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchRequestOutput{}

// BatchRequestOutput The per-line object of the batch output and error files
type BatchRequestOutput struct {
	Id *string `json:"id,omitempty"`
	// A developer-provided per-request id that will be used to match outputs to inputs.
	CustomId *string `json:"custom_id,omitempty"`
	Response NullableBatchRequestOutputResponse `json:"response,omitempty"`
	Error NullableBatchRequestOutputError `json:"error,omitempty"`
}

// NewBatchRequestOutput instantiates a new BatchRequestOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRequestOutput() *BatchRequestOutput {
	this := BatchRequestOutput{}
	return &this
}

// NewBatchRequestOutputWithDefaults instantiates a new BatchRequestOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRequestOutputWithDefaults() *BatchRequestOutput {
	this := BatchRequestOutput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BatchRequestOutput) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestOutput) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BatchRequestOutput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BatchRequestOutput) SetId(v string) {
	o.Id = &v
}

// GetCustomId returns the CustomId field value if set, zero value otherwise.
func (o *BatchRequestOutput) GetCustomId() string {
	if o == nil || IsNil(o.CustomId) {
		var ret string
		return ret
	}
	return *o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestOutput) GetCustomIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomId) {
		return nil, false
	}
	return o.CustomId, true
}

// HasCustomId returns a boolean if a field has been set.
func (o *BatchRequestOutput) HasCustomId() bool {
	if o != nil && !IsNil(o.CustomId) {
		return true
	}

	return false
}

// SetCustomId gets a reference to the given string and assigns it to the CustomId field.
func (o *BatchRequestOutput) SetCustomId(v string) {
	o.CustomId = &v
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestOutput) GetResponse() BatchRequestOutputResponse {
	if o == nil || IsNil(o.Response.Get()) {
		var ret BatchRequestOutputResponse
		return ret
	}
	return *o.Response.Get()
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestOutput) GetResponseOk() (*BatchRequestOutputResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Response.Get(), o.Response.IsSet()
}

// HasResponse returns a boolean if a field has been set.
func (o *BatchRequestOutput) HasResponse() bool {
	if o != nil && o.Response.IsSet() {
		return true
	}

	return false
}

// SetResponse gets a reference to the given NullableBatchRequestOutputResponse and assigns it to the Response field.
func (o *BatchRequestOutput) SetResponse(v BatchRequestOutputResponse) {
	o.Response.Set(&v)
}
// SetResponseNil sets the value for Response to be an explicit nil
func (o *BatchRequestOutput) SetResponseNil() {
	o.Response.Set(nil)
}

// UnsetResponse ensures that no value is present for Response, not even an explicit nil
func (o *BatchRequestOutput) UnsetResponse() {
	o.Response.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestOutput) GetError() BatchRequestOutputError {
	if o == nil || IsNil(o.Error.Get()) {
		var ret BatchRequestOutputError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestOutput) GetErrorOk() (*BatchRequestOutputError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *BatchRequestOutput) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableBatchRequestOutputError and assigns it to the Error field.
func (o *BatchRequestOutput) SetError(v BatchRequestOutputError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *BatchRequestOutput) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *BatchRequestOutput) UnsetError() {
	o.Error.Unset()
}

func (o BatchRequestOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRequestOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CustomId) {
		toSerialize["custom_id"] = o.CustomId
	}
	if o.Response.IsSet() {
		toSerialize["response"] = o.Response.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

type NullableBatchRequestOutput struct {
	value *BatchRequestOutput
	isSet bool
}

func (v NullableBatchRequestOutput) Get() *BatchRequestOutput {
	return v.value
}

func (v *NullableBatchRequestOutput) Set(val *BatchRequestOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRequestOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRequestOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRequestOutput(val *BatchRequestOutput) *NullableBatchRequestOutput {
	return &NullableBatchRequestOutput{value: val, isSet: true}
}

func (v NullableBatchRequestOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRequestOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


