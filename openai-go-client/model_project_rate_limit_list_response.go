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

// checks if the ProjectRateLimitListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectRateLimitListResponse{}

// ProjectRateLimitListResponse struct for ProjectRateLimitListResponse
type ProjectRateLimitListResponse struct {
	Object string `json:"object"`
	Data []ProjectRateLimit `json:"data"`
	FirstId string `json:"first_id"`
	LastId string `json:"last_id"`
	HasMore bool `json:"has_more"`
}

type _ProjectRateLimitListResponse ProjectRateLimitListResponse

// NewProjectRateLimitListResponse instantiates a new ProjectRateLimitListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRateLimitListResponse(object string, data []ProjectRateLimit, firstId string, lastId string, hasMore bool) *ProjectRateLimitListResponse {
	this := ProjectRateLimitListResponse{}
	this.Object = object
	this.Data = data
	this.FirstId = firstId
	this.LastId = lastId
	this.HasMore = hasMore
	return &this
}

// NewProjectRateLimitListResponseWithDefaults instantiates a new ProjectRateLimitListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRateLimitListResponseWithDefaults() *ProjectRateLimitListResponse {
	this := ProjectRateLimitListResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ProjectRateLimitListResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimitListResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ProjectRateLimitListResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ProjectRateLimitListResponse) GetData() []ProjectRateLimit {
	if o == nil {
		var ret []ProjectRateLimit
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimitListResponse) GetDataOk() ([]ProjectRateLimit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ProjectRateLimitListResponse) SetData(v []ProjectRateLimit) {
	o.Data = v
}

// GetFirstId returns the FirstId field value
func (o *ProjectRateLimitListResponse) GetFirstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimitListResponse) GetFirstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstId, true
}

// SetFirstId sets field value
func (o *ProjectRateLimitListResponse) SetFirstId(v string) {
	o.FirstId = v
}

// GetLastId returns the LastId field value
func (o *ProjectRateLimitListResponse) GetLastId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimitListResponse) GetLastIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastId, true
}

// SetLastId sets field value
func (o *ProjectRateLimitListResponse) SetLastId(v string) {
	o.LastId = v
}

// GetHasMore returns the HasMore field value
func (o *ProjectRateLimitListResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimitListResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ProjectRateLimitListResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o ProjectRateLimitListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectRateLimitListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["first_id"] = o.FirstId
	toSerialize["last_id"] = o.LastId
	toSerialize["has_more"] = o.HasMore
	return toSerialize, nil
}

func (o *ProjectRateLimitListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varProjectRateLimitListResponse := _ProjectRateLimitListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectRateLimitListResponse)

	if err != nil {
		return err
	}

	*o = ProjectRateLimitListResponse(varProjectRateLimitListResponse)

	return err
}

type NullableProjectRateLimitListResponse struct {
	value *ProjectRateLimitListResponse
	isSet bool
}

func (v NullableProjectRateLimitListResponse) Get() *ProjectRateLimitListResponse {
	return v.value
}

func (v *NullableProjectRateLimitListResponse) Set(val *ProjectRateLimitListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRateLimitListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRateLimitListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRateLimitListResponse(val *ProjectRateLimitListResponse) *NullableProjectRateLimitListResponse {
	return &NullableProjectRateLimitListResponse{value: val, isSet: true}
}

func (v NullableProjectRateLimitListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRateLimitListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

