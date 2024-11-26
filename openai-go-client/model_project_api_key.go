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

// checks if the ProjectApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectApiKey{}

// ProjectApiKey Represents an individual API key in a project.
type ProjectApiKey struct {
	// The object type, which is always `organization.project.api_key`
	Object string `json:"object"`
	// The redacted value of the API key
	RedactedValue string `json:"redacted_value"`
	// The name of the API key
	Name string `json:"name"`
	// The Unix timestamp (in seconds) of when the API key was created
	CreatedAt int32 `json:"created_at"`
	// The identifier, which can be referenced in API endpoints
	Id string `json:"id"`
	Owner ProjectApiKeyOwner `json:"owner"`
}

type _ProjectApiKey ProjectApiKey

// NewProjectApiKey instantiates a new ProjectApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectApiKey(object string, redactedValue string, name string, createdAt int32, id string, owner ProjectApiKeyOwner) *ProjectApiKey {
	this := ProjectApiKey{}
	this.Object = object
	this.RedactedValue = redactedValue
	this.Name = name
	this.CreatedAt = createdAt
	this.Id = id
	this.Owner = owner
	return &this
}

// NewProjectApiKeyWithDefaults instantiates a new ProjectApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectApiKeyWithDefaults() *ProjectApiKey {
	this := ProjectApiKey{}
	return &this
}

// GetObject returns the Object field value
func (o *ProjectApiKey) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ProjectApiKey) SetObject(v string) {
	o.Object = v
}

// GetRedactedValue returns the RedactedValue field value
func (o *ProjectApiKey) GetRedactedValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedactedValue
}

// GetRedactedValueOk returns a tuple with the RedactedValue field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetRedactedValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedactedValue, true
}

// SetRedactedValue sets field value
func (o *ProjectApiKey) SetRedactedValue(v string) {
	o.RedactedValue = v
}

// GetName returns the Name field value
func (o *ProjectApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectApiKey) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectApiKey) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectApiKey) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ProjectApiKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectApiKey) SetId(v string) {
	o.Id = v
}

// GetOwner returns the Owner field value
func (o *ProjectApiKey) GetOwner() ProjectApiKeyOwner {
	if o == nil {
		var ret ProjectApiKeyOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetOwnerOk() (*ProjectApiKeyOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *ProjectApiKey) SetOwner(v ProjectApiKeyOwner) {
	o.Owner = v
}

func (o ProjectApiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["redacted_value"] = o.RedactedValue
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["owner"] = o.Owner
	return toSerialize, nil
}

func (o *ProjectApiKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"redacted_value",
		"name",
		"created_at",
		"id",
		"owner",
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

	varProjectApiKey := _ProjectApiKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectApiKey)

	if err != nil {
		return err
	}

	*o = ProjectApiKey(varProjectApiKey)

	return err
}

type NullableProjectApiKey struct {
	value *ProjectApiKey
	isSet bool
}

func (v NullableProjectApiKey) Get() *ProjectApiKey {
	return v.value
}

func (v *NullableProjectApiKey) Set(val *ProjectApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectApiKey(val *ProjectApiKey) *NullableProjectApiKey {
	return &NullableProjectApiKey{value: val, isSet: true}
}

func (v NullableProjectApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


