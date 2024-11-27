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

// checks if the VectorStoreFileObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VectorStoreFileObject{}

// VectorStoreFileObject A list of files attached to a vector store.
type VectorStoreFileObject struct {
	// The identifier, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The object type, which is always `vector_store.file`.
	Object string `json:"object"`
	// The total vector store usage in bytes. Note that this may be different from the original file size.
	UsageBytes int32 `json:"usage_bytes"`
	// The Unix timestamp (in seconds) for when the vector store file was created.
	CreatedAt int32 `json:"created_at"`
	// The ID of the [vector store](/docs/api-reference/vector-stores/object) that the [File](/docs/api-reference/files) is attached to.
	VectorStoreId string `json:"vector_store_id"`
	// The status of the vector store file, which can be either `in_progress`, `completed`, `cancelled`, or `failed`. The status `completed` indicates that the vector store file is ready for use.
	Status string `json:"status"`
	LastError NullableVectorStoreFileObjectLastError `json:"last_error"`
	ChunkingStrategy *VectorStoreFileObjectChunkingStrategy `json:"chunking_strategy,omitempty"`
}

type _VectorStoreFileObject VectorStoreFileObject

// NewVectorStoreFileObject instantiates a new VectorStoreFileObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVectorStoreFileObject(id string, object string, usageBytes int32, createdAt int32, vectorStoreId string, status string, lastError NullableVectorStoreFileObjectLastError) *VectorStoreFileObject {
	this := VectorStoreFileObject{}
	this.Id = id
	this.Object = object
	this.UsageBytes = usageBytes
	this.CreatedAt = createdAt
	this.VectorStoreId = vectorStoreId
	this.Status = status
	this.LastError = lastError
	return &this
}

// NewVectorStoreFileObjectWithDefaults instantiates a new VectorStoreFileObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVectorStoreFileObjectWithDefaults() *VectorStoreFileObject {
	this := VectorStoreFileObject{}
	return &this
}

// GetId returns the Id field value
func (o *VectorStoreFileObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VectorStoreFileObject) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *VectorStoreFileObject) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *VectorStoreFileObject) SetObject(v string) {
	o.Object = v
}

// GetUsageBytes returns the UsageBytes field value
func (o *VectorStoreFileObject) GetUsageBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageBytes
}

// GetUsageBytesOk returns a tuple with the UsageBytes field value
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetUsageBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageBytes, true
}

// SetUsageBytes sets field value
func (o *VectorStoreFileObject) SetUsageBytes(v int32) {
	o.UsageBytes = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VectorStoreFileObject) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VectorStoreFileObject) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetVectorStoreId returns the VectorStoreId field value
func (o *VectorStoreFileObject) GetVectorStoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VectorStoreId
}

// GetVectorStoreIdOk returns a tuple with the VectorStoreId field value
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetVectorStoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VectorStoreId, true
}

// SetVectorStoreId sets field value
func (o *VectorStoreFileObject) SetVectorStoreId(v string) {
	o.VectorStoreId = v
}

// GetStatus returns the Status field value
func (o *VectorStoreFileObject) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VectorStoreFileObject) SetStatus(v string) {
	o.Status = v
}

// GetLastError returns the LastError field value
// If the value is explicit nil, the zero value for VectorStoreFileObjectLastError will be returned
func (o *VectorStoreFileObject) GetLastError() VectorStoreFileObjectLastError {
	if o == nil || o.LastError.Get() == nil {
		var ret VectorStoreFileObjectLastError
		return ret
	}

	return *o.LastError.Get()
}

// GetLastErrorOk returns a tuple with the LastError field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VectorStoreFileObject) GetLastErrorOk() (*VectorStoreFileObjectLastError, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastError.Get(), o.LastError.IsSet()
}

// SetLastError sets field value
func (o *VectorStoreFileObject) SetLastError(v VectorStoreFileObjectLastError) {
	o.LastError.Set(&v)
}

// GetChunkingStrategy returns the ChunkingStrategy field value if set, zero value otherwise.
func (o *VectorStoreFileObject) GetChunkingStrategy() VectorStoreFileObjectChunkingStrategy {
	if o == nil || IsNil(o.ChunkingStrategy) {
		var ret VectorStoreFileObjectChunkingStrategy
		return ret
	}
	return *o.ChunkingStrategy
}

// GetChunkingStrategyOk returns a tuple with the ChunkingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VectorStoreFileObject) GetChunkingStrategyOk() (*VectorStoreFileObjectChunkingStrategy, bool) {
	if o == nil || IsNil(o.ChunkingStrategy) {
		return nil, false
	}
	return o.ChunkingStrategy, true
}

// HasChunkingStrategy returns a boolean if a field has been set.
func (o *VectorStoreFileObject) HasChunkingStrategy() bool {
	if o != nil && !IsNil(o.ChunkingStrategy) {
		return true
	}

	return false
}

// SetChunkingStrategy gets a reference to the given VectorStoreFileObjectChunkingStrategy and assigns it to the ChunkingStrategy field.
func (o *VectorStoreFileObject) SetChunkingStrategy(v VectorStoreFileObjectChunkingStrategy) {
	o.ChunkingStrategy = &v
}

func (o VectorStoreFileObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VectorStoreFileObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["usage_bytes"] = o.UsageBytes
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["vector_store_id"] = o.VectorStoreId
	toSerialize["status"] = o.Status
	toSerialize["last_error"] = o.LastError.Get()
	if !IsNil(o.ChunkingStrategy) {
		toSerialize["chunking_strategy"] = o.ChunkingStrategy
	}
	return toSerialize, nil
}

func (o *VectorStoreFileObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"usage_bytes",
		"created_at",
		"vector_store_id",
		"status",
		"last_error",
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

	varVectorStoreFileObject := _VectorStoreFileObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVectorStoreFileObject)

	if err != nil {
		return err
	}

	*o = VectorStoreFileObject(varVectorStoreFileObject)

	return err
}

type NullableVectorStoreFileObject struct {
	value *VectorStoreFileObject
	isSet bool
}

func (v NullableVectorStoreFileObject) Get() *VectorStoreFileObject {
	return v.value
}

func (v *NullableVectorStoreFileObject) Set(val *VectorStoreFileObject) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorStoreFileObject) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorStoreFileObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorStoreFileObject(val *VectorStoreFileObject) *NullableVectorStoreFileObject {
	return &NullableVectorStoreFileObject{value: val, isSet: true}
}

func (v NullableVectorStoreFileObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorStoreFileObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

