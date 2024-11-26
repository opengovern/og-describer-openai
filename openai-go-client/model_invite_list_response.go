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

// checks if the InviteListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteListResponse{}

// InviteListResponse struct for InviteListResponse
type InviteListResponse struct {
	// The object type, which is always `list`
	Object string `json:"object"`
	Data []Invite `json:"data"`
	// The first `invite_id` in the retrieved `list`
	FirstId *string `json:"first_id,omitempty"`
	// The last `invite_id` in the retrieved `list`
	LastId *string `json:"last_id,omitempty"`
	// The `has_more` property is used for pagination to indicate there are additional results.
	HasMore *bool `json:"has_more,omitempty"`
}

type _InviteListResponse InviteListResponse

// NewInviteListResponse instantiates a new InviteListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteListResponse(object string, data []Invite) *InviteListResponse {
	this := InviteListResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewInviteListResponseWithDefaults instantiates a new InviteListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteListResponseWithDefaults() *InviteListResponse {
	this := InviteListResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *InviteListResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *InviteListResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *InviteListResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *InviteListResponse) GetData() []Invite {
	if o == nil {
		var ret []Invite
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InviteListResponse) GetDataOk() ([]Invite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InviteListResponse) SetData(v []Invite) {
	o.Data = v
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *InviteListResponse) GetFirstId() string {
	if o == nil || IsNil(o.FirstId) {
		var ret string
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteListResponse) GetFirstIdOk() (*string, bool) {
	if o == nil || IsNil(o.FirstId) {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *InviteListResponse) HasFirstId() bool {
	if o != nil && !IsNil(o.FirstId) {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given string and assigns it to the FirstId field.
func (o *InviteListResponse) SetFirstId(v string) {
	o.FirstId = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *InviteListResponse) GetLastId() string {
	if o == nil || IsNil(o.LastId) {
		var ret string
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteListResponse) GetLastIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastId) {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *InviteListResponse) HasLastId() bool {
	if o != nil && !IsNil(o.LastId) {
		return true
	}

	return false
}

// SetLastId gets a reference to the given string and assigns it to the LastId field.
func (o *InviteListResponse) SetLastId(v string) {
	o.LastId = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *InviteListResponse) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteListResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *InviteListResponse) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *InviteListResponse) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o InviteListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	if !IsNil(o.FirstId) {
		toSerialize["first_id"] = o.FirstId
	}
	if !IsNil(o.LastId) {
		toSerialize["last_id"] = o.LastId
	}
	if !IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	return toSerialize, nil
}

func (o *InviteListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"data",
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

	varInviteListResponse := _InviteListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInviteListResponse)

	if err != nil {
		return err
	}

	*o = InviteListResponse(varInviteListResponse)

	return err
}

type NullableInviteListResponse struct {
	value *InviteListResponse
	isSet bool
}

func (v NullableInviteListResponse) Get() *InviteListResponse {
	return v.value
}

func (v *NullableInviteListResponse) Set(val *InviteListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteListResponse(val *InviteListResponse) *NullableInviteListResponse {
	return &NullableInviteListResponse{value: val, isSet: true}
}

func (v NullableInviteListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


