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

// checks if the AuditLogActor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogActor{}

// AuditLogActor The actor who performed the audit logged action.
type AuditLogActor struct {
	// The type of actor. Is either `session` or `api_key`.
	Type *string `json:"type,omitempty"`
	Session AuditLogActorSession `json:"session,omitempty"`
	ApiKey AuditLogActorApiKey `json:"api_key,omitempty"`
}

// NewAuditLogActor instantiates a new AuditLogActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogActor() *AuditLogActor {
	this := AuditLogActor{}
	return &this
}

// NewAuditLogActorWithDefaults instantiates a new AuditLogActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogActorWithDefaults() *AuditLogActor {
	this := AuditLogActor{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditLogActor) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditLogActor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditLogActor) SetType(v string) {
	o.Type = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *AuditLogActor) GetSession() AuditLogActorSession {
	if o == nil || IsNil(o.Session) {
		var ret AuditLogActorSession
		return ret
	}
	return o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActor) GetSessionOk() (AuditLogActorSession, bool) {
	if o == nil || IsNil(o.Session) {
		return AuditLogActorSession{}, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *AuditLogActor) HasSession() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given AuditLogActorSession and assigns it to the Session field.
func (o *AuditLogActor) SetSession(v AuditLogActorSession) {
	o.Session = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *AuditLogActor) GetApiKey() AuditLogActorApiKey {
	if o == nil || IsNil(o.ApiKey) {
		var ret AuditLogActorApiKey
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogActor) GetApiKeyOk() (AuditLogActorApiKey, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return AuditLogActorApiKey{}, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *AuditLogActor) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given AuditLogActorApiKey and assigns it to the ApiKey field.
func (o *AuditLogActor) SetApiKey(v AuditLogActorApiKey) {
	o.ApiKey = v
}

func (o AuditLogActor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogActor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullableAuditLogActor struct {
	value *AuditLogActor
	isSet bool
}

func (v NullableAuditLogActor) Get() *AuditLogActor {
	return v.value
}

func (v *NullableAuditLogActor) Set(val *AuditLogActor) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogActor) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogActor(val *AuditLogActor) *NullableAuditLogActor {
	return &NullableAuditLogActor{value: val, isSet: true}
}

func (v NullableAuditLogActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

