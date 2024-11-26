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

// checks if the ProjectApiKeyDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectApiKeyDeleteResponse{}

// ProjectApiKeyDeleteResponse struct for ProjectApiKeyDeleteResponse
type ProjectApiKeyDeleteResponse struct {
	Object string `json:"object"`
	Id string `json:"id"`
	Deleted bool `json:"deleted"`
}

type _ProjectApiKeyDeleteResponse ProjectApiKeyDeleteResponse

// NewProjectApiKeyDeleteResponse instantiates a new ProjectApiKeyDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectApiKeyDeleteResponse(object string, id string, deleted bool) *ProjectApiKeyDeleteResponse {
	this := ProjectApiKeyDeleteResponse{}
	this.Object = object
	this.Id = id
	this.Deleted = deleted
	return &this
}

// NewProjectApiKeyDeleteResponseWithDefaults instantiates a new ProjectApiKeyDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectApiKeyDeleteResponseWithDefaults() *ProjectApiKeyDeleteResponse {
	this := ProjectApiKeyDeleteResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ProjectApiKeyDeleteResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKeyDeleteResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ProjectApiKeyDeleteResponse) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *ProjectApiKeyDeleteResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKeyDeleteResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectApiKeyDeleteResponse) SetId(v string) {
	o.Id = v
}

// GetDeleted returns the Deleted field value
func (o *ProjectApiKeyDeleteResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKeyDeleteResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *ProjectApiKeyDeleteResponse) SetDeleted(v bool) {
	o.Deleted = v
}

func (o ProjectApiKeyDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectApiKeyDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

func (o *ProjectApiKeyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"deleted",
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

	varProjectApiKeyDeleteResponse := _ProjectApiKeyDeleteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectApiKeyDeleteResponse)

	if err != nil {
		return err
	}

	*o = ProjectApiKeyDeleteResponse(varProjectApiKeyDeleteResponse)

	return err
}

type NullableProjectApiKeyDeleteResponse struct {
	value *ProjectApiKeyDeleteResponse
	isSet bool
}

func (v NullableProjectApiKeyDeleteResponse) Get() *ProjectApiKeyDeleteResponse {
	return v.value
}

func (v *NullableProjectApiKeyDeleteResponse) Set(val *ProjectApiKeyDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectApiKeyDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectApiKeyDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectApiKeyDeleteResponse(val *ProjectApiKeyDeleteResponse) *NullableProjectApiKeyDeleteResponse {
	return &NullableProjectApiKeyDeleteResponse{value: val, isSet: true}
}

func (v NullableProjectApiKeyDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectApiKeyDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


