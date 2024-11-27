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

// checks if the ListRunStepsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRunStepsResponse{}

// ListRunStepsResponse struct for ListRunStepsResponse
type ListRunStepsResponse struct {
	Object string `json:"object"`
	Data []RunStepObject `json:"data"`
	FirstId string `json:"first_id"`
	LastId string `json:"last_id"`
	HasMore bool `json:"has_more"`
}

type _ListRunStepsResponse ListRunStepsResponse

// NewListRunStepsResponse instantiates a new ListRunStepsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRunStepsResponse(object string, data []RunStepObject, firstId string, lastId string, hasMore bool) *ListRunStepsResponse {
	this := ListRunStepsResponse{}
	this.Object = object
	this.Data = data
	this.FirstId = firstId
	this.LastId = lastId
	this.HasMore = hasMore
	return &this
}

// NewListRunStepsResponseWithDefaults instantiates a new ListRunStepsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRunStepsResponseWithDefaults() *ListRunStepsResponse {
	this := ListRunStepsResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ListRunStepsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ListRunStepsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ListRunStepsResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ListRunStepsResponse) GetData() []RunStepObject {
	if o == nil {
		var ret []RunStepObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListRunStepsResponse) GetDataOk() ([]RunStepObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListRunStepsResponse) SetData(v []RunStepObject) {
	o.Data = v
}

// GetFirstId returns the FirstId field value
func (o *ListRunStepsResponse) GetFirstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value
// and a boolean to check if the value has been set.
func (o *ListRunStepsResponse) GetFirstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstId, true
}

// SetFirstId sets field value
func (o *ListRunStepsResponse) SetFirstId(v string) {
	o.FirstId = v
}

// GetLastId returns the LastId field value
func (o *ListRunStepsResponse) GetLastId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value
// and a boolean to check if the value has been set.
func (o *ListRunStepsResponse) GetLastIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastId, true
}

// SetLastId sets field value
func (o *ListRunStepsResponse) SetLastId(v string) {
	o.LastId = v
}

// GetHasMore returns the HasMore field value
func (o *ListRunStepsResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ListRunStepsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ListRunStepsResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o ListRunStepsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRunStepsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["first_id"] = o.FirstId
	toSerialize["last_id"] = o.LastId
	toSerialize["has_more"] = o.HasMore
	return toSerialize, nil
}

func (o *ListRunStepsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"data",
		"first_id",
		"last_id",
		"has_more",
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

	varListRunStepsResponse := _ListRunStepsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListRunStepsResponse)

	if err != nil {
		return err
	}

	*o = ListRunStepsResponse(varListRunStepsResponse)

	return err
}

type NullableListRunStepsResponse struct {
	value *ListRunStepsResponse
	isSet bool
}

func (v NullableListRunStepsResponse) Get() *ListRunStepsResponse {
	return v.value
}

func (v *NullableListRunStepsResponse) Set(val *ListRunStepsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRunStepsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRunStepsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRunStepsResponse(val *ListRunStepsResponse) *NullableListRunStepsResponse {
	return &NullableListRunStepsResponse{value: val, isSet: true}
}

func (v NullableListRunStepsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRunStepsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

