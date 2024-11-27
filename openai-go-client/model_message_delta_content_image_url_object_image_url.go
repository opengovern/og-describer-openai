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

// checks if the MessageDeltaContentImageUrlObjectImageUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDeltaContentImageUrlObjectImageUrl{}

// MessageDeltaContentImageUrlObjectImageUrl struct for MessageDeltaContentImageUrlObjectImageUrl
type MessageDeltaContentImageUrlObjectImageUrl struct {
	// The URL of the image, must be a supported image types: jpeg, jpg, png, gif, webp.
	Url *string `json:"url,omitempty"`
	// Specifies the detail level of the image. `low` uses fewer tokens, you can opt in to high resolution using `high`.
	Detail *string `json:"detail,omitempty"`
}

// NewMessageDeltaContentImageUrlObjectImageUrl instantiates a new MessageDeltaContentImageUrlObjectImageUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDeltaContentImageUrlObjectImageUrl() *MessageDeltaContentImageUrlObjectImageUrl {
	this := MessageDeltaContentImageUrlObjectImageUrl{}
	var detail string = "auto"
	this.Detail = &detail
	return &this
}

// NewMessageDeltaContentImageUrlObjectImageUrlWithDefaults instantiates a new MessageDeltaContentImageUrlObjectImageUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDeltaContentImageUrlObjectImageUrlWithDefaults() *MessageDeltaContentImageUrlObjectImageUrl {
	this := MessageDeltaContentImageUrlObjectImageUrl{}
	var detail string = "auto"
	this.Detail = &detail
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MessageDeltaContentImageUrlObjectImageUrl) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeltaContentImageUrlObjectImageUrl) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageDeltaContentImageUrlObjectImageUrl) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MessageDeltaContentImageUrlObjectImageUrl) SetUrl(v string) {
	o.Url = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *MessageDeltaContentImageUrlObjectImageUrl) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeltaContentImageUrlObjectImageUrl) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *MessageDeltaContentImageUrlObjectImageUrl) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *MessageDeltaContentImageUrlObjectImageUrl) SetDetail(v string) {
	o.Detail = &v
}

func (o MessageDeltaContentImageUrlObjectImageUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDeltaContentImageUrlObjectImageUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableMessageDeltaContentImageUrlObjectImageUrl struct {
	value *MessageDeltaContentImageUrlObjectImageUrl
	isSet bool
}

func (v NullableMessageDeltaContentImageUrlObjectImageUrl) Get() *MessageDeltaContentImageUrlObjectImageUrl {
	return v.value
}

func (v *NullableMessageDeltaContentImageUrlObjectImageUrl) Set(val *MessageDeltaContentImageUrlObjectImageUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDeltaContentImageUrlObjectImageUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDeltaContentImageUrlObjectImageUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDeltaContentImageUrlObjectImageUrl(val *MessageDeltaContentImageUrlObjectImageUrl) *NullableMessageDeltaContentImageUrlObjectImageUrl {
	return &NullableMessageDeltaContentImageUrlObjectImageUrl{value: val, isSet: true}
}

func (v NullableMessageDeltaContentImageUrlObjectImageUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDeltaContentImageUrlObjectImageUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

