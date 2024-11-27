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

// checks if the ProjectServiceAccountApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectServiceAccountApiKey{}

// ProjectServiceAccountApiKey struct for ProjectServiceAccountApiKey
type ProjectServiceAccountApiKey struct {
	// The object type, which is always `organization.project.service_account.api_key`
	Object string `json:"object"`
	Value string `json:"value"`
	Name string `json:"name"`
	CreatedAt int32 `json:"created_at"`
	Id string `json:"id"`
}

type _ProjectServiceAccountApiKey ProjectServiceAccountApiKey

// NewProjectServiceAccountApiKey instantiates a new ProjectServiceAccountApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServiceAccountApiKey(object string, value string, name string, createdAt int32, id string) *ProjectServiceAccountApiKey {
	this := ProjectServiceAccountApiKey{}
	this.Object = object
	this.Value = value
	this.Name = name
	this.CreatedAt = createdAt
	this.Id = id
	return &this
}

// NewProjectServiceAccountApiKeyWithDefaults instantiates a new ProjectServiceAccountApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServiceAccountApiKeyWithDefaults() *ProjectServiceAccountApiKey {
	this := ProjectServiceAccountApiKey{}
	return &this
}

// GetObject returns the Object field value
func (o *ProjectServiceAccountApiKey) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceAccountApiKey) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ProjectServiceAccountApiKey) SetObject(v string) {
	o.Object = v
}

// GetValue returns the Value field value
func (o *ProjectServiceAccountApiKey) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceAccountApiKey) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ProjectServiceAccountApiKey) SetValue(v string) {
	o.Value = v
}

// GetName returns the Name field value
func (o *ProjectServiceAccountApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceAccountApiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectServiceAccountApiKey) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectServiceAccountApiKey) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceAccountApiKey) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectServiceAccountApiKey) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ProjectServiceAccountApiKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceAccountApiKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectServiceAccountApiKey) SetId(v string) {
	o.Id = v
}

func (o ProjectServiceAccountApiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectServiceAccountApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["value"] = o.Value
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ProjectServiceAccountApiKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"value",
		"name",
		"created_at",
		"id",
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

	varProjectServiceAccountApiKey := _ProjectServiceAccountApiKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectServiceAccountApiKey)

	if err != nil {
		return err
	}

	*o = ProjectServiceAccountApiKey(varProjectServiceAccountApiKey)

	return err
}

type NullableProjectServiceAccountApiKey struct {
	value *ProjectServiceAccountApiKey
	isSet bool
}

func (v NullableProjectServiceAccountApiKey) Get() *ProjectServiceAccountApiKey {
	return v.value
}

func (v *NullableProjectServiceAccountApiKey) Set(val *ProjectServiceAccountApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServiceAccountApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServiceAccountApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServiceAccountApiKey(val *ProjectServiceAccountApiKey) *NullableProjectServiceAccountApiKey {
	return &NullableProjectServiceAccountApiKey{value: val, isSet: true}
}

func (v NullableProjectServiceAccountApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServiceAccountApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

