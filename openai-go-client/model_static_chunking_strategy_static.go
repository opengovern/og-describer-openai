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

// checks if the StaticChunkingStrategyStatic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticChunkingStrategyStatic{}

// StaticChunkingStrategyStatic struct for StaticChunkingStrategyStatic
type StaticChunkingStrategyStatic struct {
	// The maximum number of tokens in each chunk. The default value is `800`. The minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens int32 `json:"max_chunk_size_tokens"`
	// The number of tokens that overlap between chunks. The default value is `400`.  Note that the overlap must not exceed half of `max_chunk_size_tokens`. 
	ChunkOverlapTokens int32 `json:"chunk_overlap_tokens"`
}

type _StaticChunkingStrategyStatic StaticChunkingStrategyStatic

// NewStaticChunkingStrategyStatic instantiates a new StaticChunkingStrategyStatic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticChunkingStrategyStatic(maxChunkSizeTokens int32, chunkOverlapTokens int32) *StaticChunkingStrategyStatic {
	this := StaticChunkingStrategyStatic{}
	this.MaxChunkSizeTokens = maxChunkSizeTokens
	this.ChunkOverlapTokens = chunkOverlapTokens
	return &this
}

// NewStaticChunkingStrategyStaticWithDefaults instantiates a new StaticChunkingStrategyStatic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticChunkingStrategyStaticWithDefaults() *StaticChunkingStrategyStatic {
	this := StaticChunkingStrategyStatic{}
	return &this
}

// GetMaxChunkSizeTokens returns the MaxChunkSizeTokens field value
func (o *StaticChunkingStrategyStatic) GetMaxChunkSizeTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxChunkSizeTokens
}

// GetMaxChunkSizeTokensOk returns a tuple with the MaxChunkSizeTokens field value
// and a boolean to check if the value has been set.
func (o *StaticChunkingStrategyStatic) GetMaxChunkSizeTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxChunkSizeTokens, true
}

// SetMaxChunkSizeTokens sets field value
func (o *StaticChunkingStrategyStatic) SetMaxChunkSizeTokens(v int32) {
	o.MaxChunkSizeTokens = v
}

// GetChunkOverlapTokens returns the ChunkOverlapTokens field value
func (o *StaticChunkingStrategyStatic) GetChunkOverlapTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChunkOverlapTokens
}

// GetChunkOverlapTokensOk returns a tuple with the ChunkOverlapTokens field value
// and a boolean to check if the value has been set.
func (o *StaticChunkingStrategyStatic) GetChunkOverlapTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChunkOverlapTokens, true
}

// SetChunkOverlapTokens sets field value
func (o *StaticChunkingStrategyStatic) SetChunkOverlapTokens(v int32) {
	o.ChunkOverlapTokens = v
}

func (o StaticChunkingStrategyStatic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticChunkingStrategyStatic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_chunk_size_tokens"] = o.MaxChunkSizeTokens
	toSerialize["chunk_overlap_tokens"] = o.ChunkOverlapTokens
	return toSerialize, nil
}

func (o *StaticChunkingStrategyStatic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_chunk_size_tokens",
		"chunk_overlap_tokens",
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

	varStaticChunkingStrategyStatic := _StaticChunkingStrategyStatic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticChunkingStrategyStatic)

	if err != nil {
		return err
	}

	*o = StaticChunkingStrategyStatic(varStaticChunkingStrategyStatic)

	return err
}

type NullableStaticChunkingStrategyStatic struct {
	value *StaticChunkingStrategyStatic
	isSet bool
}

func (v NullableStaticChunkingStrategyStatic) Get() *StaticChunkingStrategyStatic {
	return v.value
}

func (v *NullableStaticChunkingStrategyStatic) Set(val *StaticChunkingStrategyStatic) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticChunkingStrategyStatic) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticChunkingStrategyStatic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticChunkingStrategyStatic(val *StaticChunkingStrategyStatic) *NullableStaticChunkingStrategyStatic {
	return &NullableStaticChunkingStrategyStatic{value: val, isSet: true}
}

func (v NullableStaticChunkingStrategyStatic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticChunkingStrategyStatic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

