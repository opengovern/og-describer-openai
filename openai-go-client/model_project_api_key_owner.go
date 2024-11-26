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

// checks if the ProjectApiKeyOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectApiKeyOwner{}

// ProjectApiKeyOwner struct for ProjectApiKeyOwner
type ProjectApiKeyOwner struct {
	// `user` or `service_account`
	Type *string `json:"type,omitempty"`
	User *ProjectUser `json:"user,omitempty"`
	ServiceAccount *ProjectServiceAccount `json:"service_account,omitempty"`
}

// NewProjectApiKeyOwner instantiates a new ProjectApiKeyOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectApiKeyOwner() *ProjectApiKeyOwner {
	this := ProjectApiKeyOwner{}
	return &this
}

// NewProjectApiKeyOwnerWithDefaults instantiates a new ProjectApiKeyOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectApiKeyOwnerWithDefaults() *ProjectApiKeyOwner {
	this := ProjectApiKeyOwner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectApiKeyOwner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKeyOwner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectApiKeyOwner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectApiKeyOwner) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ProjectApiKeyOwner) GetUser() ProjectUser {
	if o == nil || IsNil(o.User) {
		var ret ProjectUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKeyOwner) GetUserOk() (*ProjectUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ProjectApiKeyOwner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given ProjectUser and assigns it to the User field.
func (o *ProjectApiKeyOwner) SetUser(v ProjectUser) {
	o.User = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *ProjectApiKeyOwner) GetServiceAccount() ProjectServiceAccount {
	if o == nil || IsNil(o.ServiceAccount) {
		var ret ProjectServiceAccount
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKeyOwner) GetServiceAccountOk() (*ProjectServiceAccount, bool) {
	if o == nil || IsNil(o.ServiceAccount) {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *ProjectApiKeyOwner) HasServiceAccount() bool {
	if o != nil && !IsNil(o.ServiceAccount) {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given ProjectServiceAccount and assigns it to the ServiceAccount field.
func (o *ProjectApiKeyOwner) SetServiceAccount(v ProjectServiceAccount) {
	o.ServiceAccount = &v
}

func (o ProjectApiKeyOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectApiKeyOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.ServiceAccount) {
		toSerialize["service_account"] = o.ServiceAccount
	}
	return toSerialize, nil
}

type NullableProjectApiKeyOwner struct {
	value *ProjectApiKeyOwner
	isSet bool
}

func (v NullableProjectApiKeyOwner) Get() *ProjectApiKeyOwner {
	return v.value
}

func (v *NullableProjectApiKeyOwner) Set(val *ProjectApiKeyOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectApiKeyOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectApiKeyOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectApiKeyOwner(val *ProjectApiKeyOwner) *NullableProjectApiKeyOwner {
	return &NullableProjectApiKeyOwner{value: val, isSet: true}
}

func (v NullableProjectApiKeyOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectApiKeyOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


