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

// checks if the AuditLogActorApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogActorApiKey{}

// AuditLogActorApiKey The API Key used to perform the audit logged action.
type AuditLogActorApiKey struct {
	// The tracking id of the API key.
	Id *string `json:"id,omitempty"`
	// The type of API key. Can be either `user` or `service_account`.
	Type *string `json:"type,omitempty"`
	User *AuditLogActorUser `json:"user,omitempty"`
	ServiceAccount *AuditLogActorServiceAccount `json:"service_account,omitempty"`
}

// NewAuditLogActorApiKey instantiates a new AuditLogActorApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogActorApiKey() *AuditLogActorApiKey {
	this := AuditLogActorApiKey{}
	return &this
}

// NewAuditLogActorApiKeyWithDefaults instantiates a new AuditLogActorApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogActorApiKeyWithDefaults() *AuditLogActorApiKey {
	this := AuditLogActorApiKey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogActorApiKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActorApiKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditLogActorApiKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditLogActorApiKey) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditLogActorApiKey) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActorApiKey) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditLogActorApiKey) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditLogActorApiKey) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuditLogActorApiKey) GetUser() AuditLogActorUser {
	if o == nil || IsNil(o.User) {
		var ret AuditLogActorUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActorApiKey) GetUserOk() (*AuditLogActorUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuditLogActorApiKey) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given AuditLogActorUser and assigns it to the User field.
func (o *AuditLogActorApiKey) SetUser(v AuditLogActorUser) {
	o.User = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *AuditLogActorApiKey) GetServiceAccount() AuditLogActorServiceAccount {
	if o == nil || IsNil(o.ServiceAccount) {
		var ret AuditLogActorServiceAccount
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActorApiKey) GetServiceAccountOk() (*AuditLogActorServiceAccount, bool) {
	if o == nil || IsNil(o.ServiceAccount) {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *AuditLogActorApiKey) HasServiceAccount() bool {
	if o != nil && !IsNil(o.ServiceAccount) {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given AuditLogActorServiceAccount and assigns it to the ServiceAccount field.
func (o *AuditLogActorApiKey) SetServiceAccount(v AuditLogActorServiceAccount) {
	o.ServiceAccount = &v
}

func (o AuditLogActorApiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogActorApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
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

type NullableAuditLogActorApiKey struct {
	value *AuditLogActorApiKey
	isSet bool
}

func (v NullableAuditLogActorApiKey) Get() *AuditLogActorApiKey {
	return v.value
}

func (v *NullableAuditLogActorApiKey) Set(val *AuditLogActorApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogActorApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogActorApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogActorApiKey(val *AuditLogActorApiKey) *NullableAuditLogActorApiKey {
	return &NullableAuditLogActorApiKey{value: val, isSet: true}
}

func (v NullableAuditLogActorApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogActorApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


