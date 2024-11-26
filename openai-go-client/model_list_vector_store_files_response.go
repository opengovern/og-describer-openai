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

// checks if the ListVectorStoreFilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVectorStoreFilesResponse{}

// ListVectorStoreFilesResponse struct for ListVectorStoreFilesResponse
type ListVectorStoreFilesResponse struct {
	Object string `json:"object"`
	Data []VectorStoreFileObject `json:"data"`
	FirstId string `json:"first_id"`
	LastId string `json:"last_id"`
	HasMore bool `json:"has_more"`
}

type _ListVectorStoreFilesResponse ListVectorStoreFilesResponse

// NewListVectorStoreFilesResponse instantiates a new ListVectorStoreFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVectorStoreFilesResponse(object string, data []VectorStoreFileObject, firstId string, lastId string, hasMore bool) *ListVectorStoreFilesResponse {
	this := ListVectorStoreFilesResponse{}
	this.Object = object
	this.Data = data
	this.FirstId = firstId
	this.LastId = lastId
	this.HasMore = hasMore
	return &this
}

// NewListVectorStoreFilesResponseWithDefaults instantiates a new ListVectorStoreFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVectorStoreFilesResponseWithDefaults() *ListVectorStoreFilesResponse {
	this := ListVectorStoreFilesResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ListVectorStoreFilesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ListVectorStoreFilesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ListVectorStoreFilesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ListVectorStoreFilesResponse) GetData() []VectorStoreFileObject {
	if o == nil {
		var ret []VectorStoreFileObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListVectorStoreFilesResponse) GetDataOk() ([]VectorStoreFileObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListVectorStoreFilesResponse) SetData(v []VectorStoreFileObject) {
	o.Data = v
}

// GetFirstId returns the FirstId field value
func (o *ListVectorStoreFilesResponse) GetFirstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value
// and a boolean to check if the value has been set.
func (o *ListVectorStoreFilesResponse) GetFirstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstId, true
}

// SetFirstId sets field value
func (o *ListVectorStoreFilesResponse) SetFirstId(v string) {
	o.FirstId = v
}

// GetLastId returns the LastId field value
func (o *ListVectorStoreFilesResponse) GetLastId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value
// and a boolean to check if the value has been set.
func (o *ListVectorStoreFilesResponse) GetLastIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastId, true
}

// SetLastId sets field value
func (o *ListVectorStoreFilesResponse) SetLastId(v string) {
	o.LastId = v
}

// GetHasMore returns the HasMore field value
func (o *ListVectorStoreFilesResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ListVectorStoreFilesResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ListVectorStoreFilesResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o ListVectorStoreFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVectorStoreFilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["first_id"] = o.FirstId
	toSerialize["last_id"] = o.LastId
	toSerialize["has_more"] = o.HasMore
	return toSerialize, nil
}

func (o *ListVectorStoreFilesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListVectorStoreFilesResponse := _ListVectorStoreFilesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListVectorStoreFilesResponse)

	if err != nil {
		return err
	}

	*o = ListVectorStoreFilesResponse(varListVectorStoreFilesResponse)

	return err
}

type NullableListVectorStoreFilesResponse struct {
	value *ListVectorStoreFilesResponse
	isSet bool
}

func (v NullableListVectorStoreFilesResponse) Get() *ListVectorStoreFilesResponse {
	return v.value
}

func (v *NullableListVectorStoreFilesResponse) Set(val *ListVectorStoreFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVectorStoreFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVectorStoreFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVectorStoreFilesResponse(val *ListVectorStoreFilesResponse) *NullableListVectorStoreFilesResponse {
	return &NullableListVectorStoreFilesResponse{value: val, isSet: true}
}

func (v NullableListVectorStoreFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVectorStoreFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


