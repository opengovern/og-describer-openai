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

// checks if the ThreadObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadObject{}

// ThreadObject Represents a thread that contains [messages](/docs/api-reference/messages).
type ThreadObject struct {
	// The identifier, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The object type, which is always `thread`.
	Object string `json:"object"`
	// The Unix timestamp (in seconds) for when the thread was created.
	CreatedAt int32 `json:"created_at"`
	ToolResources NullableModifyThreadRequestToolResources `json:"tool_resources"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long. 
	Metadata map[string]interface{} `json:"metadata"`
}

type _ThreadObject ThreadObject

// NewThreadObject instantiates a new ThreadObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadObject(id string, object string, createdAt int32, toolResources NullableModifyThreadRequestToolResources, metadata map[string]interface{}) *ThreadObject {
	this := ThreadObject{}
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	this.ToolResources = toolResources
	this.Metadata = metadata
	return &this
}

// NewThreadObjectWithDefaults instantiates a new ThreadObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadObjectWithDefaults() *ThreadObject {
	this := ThreadObject{}
	return &this
}

// GetId returns the Id field value
func (o *ThreadObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThreadObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThreadObject) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *ThreadObject) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ThreadObject) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ThreadObject) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ThreadObject) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ThreadObject) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ThreadObject) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetToolResources returns the ToolResources field value
// If the value is explicit nil, the zero value for ModifyThreadRequestToolResources will be returned
func (o *ThreadObject) GetToolResources() ModifyThreadRequestToolResources {
	if o == nil || o.ToolResources.Get() == nil {
		var ret ModifyThreadRequestToolResources
		return ret
	}

	return *o.ToolResources.Get()
}

// GetToolResourcesOk returns a tuple with the ToolResources field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadObject) GetToolResourcesOk() (*ModifyThreadRequestToolResources, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolResources.Get(), o.ToolResources.IsSet()
}

// SetToolResources sets field value
func (o *ThreadObject) SetToolResources(v ModifyThreadRequestToolResources) {
	o.ToolResources.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ThreadObject) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadObject) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ThreadObject) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ThreadObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["tool_resources"] = o.ToolResources.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *ThreadObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"created_at",
		"tool_resources",
		"metadata",
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

	varThreadObject := _ThreadObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThreadObject)

	if err != nil {
		return err
	}

	*o = ThreadObject(varThreadObject)

	return err
}

type NullableThreadObject struct {
	value *ThreadObject
	isSet bool
}

func (v NullableThreadObject) Get() *ThreadObject {
	return v.value
}

func (v *NullableThreadObject) Set(val *ThreadObject) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadObject) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadObject(val *ThreadObject) *NullableThreadObject {
	return &NullableThreadObject{value: val, isSet: true}
}

func (v NullableThreadObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

