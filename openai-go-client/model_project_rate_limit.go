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

// checks if the ProjectRateLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectRateLimit{}

// ProjectRateLimit Represents a project rate limit config.
type ProjectRateLimit struct {
	// The object type, which is always `project.rate_limit`
	Object string `json:"object"`
	// The identifier, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The model this rate limit applies to.
	Model string `json:"model"`
	// The maximum requests per minute.
	MaxRequestsPer1Minute int32 `json:"max_requests_per_1_minute"`
	// The maximum tokens per minute.
	MaxTokensPer1Minute int32 `json:"max_tokens_per_1_minute"`
	// The maximum images per minute. Only present for relevant models.
	MaxImagesPer1Minute *int32 `json:"max_images_per_1_minute,omitempty"`
	// The maximum audio megabytes per minute. Only present for relevant models.
	MaxAudioMegabytesPer1Minute *int32 `json:"max_audio_megabytes_per_1_minute,omitempty"`
	// The maximum requests per day. Only present for relevant models.
	MaxRequestsPer1Day *int32 `json:"max_requests_per_1_day,omitempty"`
	// The maximum batch input tokens per day. Only present for relevant models.
	Batch1DayMaxInputTokens *int32 `json:"batch_1_day_max_input_tokens,omitempty"`
}

type _ProjectRateLimit ProjectRateLimit

// NewProjectRateLimit instantiates a new ProjectRateLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRateLimit(object string, id string, model string, maxRequestsPer1Minute int32, maxTokensPer1Minute int32) *ProjectRateLimit {
	this := ProjectRateLimit{}
	this.Object = object
	this.Id = id
	this.Model = model
	this.MaxRequestsPer1Minute = maxRequestsPer1Minute
	this.MaxTokensPer1Minute = maxTokensPer1Minute
	return &this
}

// NewProjectRateLimitWithDefaults instantiates a new ProjectRateLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRateLimitWithDefaults() *ProjectRateLimit {
	this := ProjectRateLimit{}
	return &this
}

// GetObject returns the Object field value
func (o *ProjectRateLimit) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ProjectRateLimit) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *ProjectRateLimit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectRateLimit) SetId(v string) {
	o.Id = v
}

// GetModel returns the Model field value
func (o *ProjectRateLimit) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ProjectRateLimit) SetModel(v string) {
	o.Model = v
}

// GetMaxRequestsPer1Minute returns the MaxRequestsPer1Minute field value
func (o *ProjectRateLimit) GetMaxRequestsPer1Minute() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxRequestsPer1Minute
}

// GetMaxRequestsPer1MinuteOk returns a tuple with the MaxRequestsPer1Minute field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetMaxRequestsPer1MinuteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRequestsPer1Minute, true
}

// SetMaxRequestsPer1Minute sets field value
func (o *ProjectRateLimit) SetMaxRequestsPer1Minute(v int32) {
	o.MaxRequestsPer1Minute = v
}

// GetMaxTokensPer1Minute returns the MaxTokensPer1Minute field value
func (o *ProjectRateLimit) GetMaxTokensPer1Minute() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxTokensPer1Minute
}

// GetMaxTokensPer1MinuteOk returns a tuple with the MaxTokensPer1Minute field value
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetMaxTokensPer1MinuteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTokensPer1Minute, true
}

// SetMaxTokensPer1Minute sets field value
func (o *ProjectRateLimit) SetMaxTokensPer1Minute(v int32) {
	o.MaxTokensPer1Minute = v
}

// GetMaxImagesPer1Minute returns the MaxImagesPer1Minute field value if set, zero value otherwise.
func (o *ProjectRateLimit) GetMaxImagesPer1Minute() int32 {
	if o == nil || IsNil(o.MaxImagesPer1Minute) {
		var ret int32
		return ret
	}
	return *o.MaxImagesPer1Minute
}

// GetMaxImagesPer1MinuteOk returns a tuple with the MaxImagesPer1Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetMaxImagesPer1MinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxImagesPer1Minute) {
		return nil, false
	}
	return o.MaxImagesPer1Minute, true
}

// HasMaxImagesPer1Minute returns a boolean if a field has been set.
func (o *ProjectRateLimit) HasMaxImagesPer1Minute() bool {
	if o != nil && !IsNil(o.MaxImagesPer1Minute) {
		return true
	}

	return false
}

// SetMaxImagesPer1Minute gets a reference to the given int32 and assigns it to the MaxImagesPer1Minute field.
func (o *ProjectRateLimit) SetMaxImagesPer1Minute(v int32) {
	o.MaxImagesPer1Minute = &v
}

// GetMaxAudioMegabytesPer1Minute returns the MaxAudioMegabytesPer1Minute field value if set, zero value otherwise.
func (o *ProjectRateLimit) GetMaxAudioMegabytesPer1Minute() int32 {
	if o == nil || IsNil(o.MaxAudioMegabytesPer1Minute) {
		var ret int32
		return ret
	}
	return *o.MaxAudioMegabytesPer1Minute
}

// GetMaxAudioMegabytesPer1MinuteOk returns a tuple with the MaxAudioMegabytesPer1Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetMaxAudioMegabytesPer1MinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAudioMegabytesPer1Minute) {
		return nil, false
	}
	return o.MaxAudioMegabytesPer1Minute, true
}

// HasMaxAudioMegabytesPer1Minute returns a boolean if a field has been set.
func (o *ProjectRateLimit) HasMaxAudioMegabytesPer1Minute() bool {
	if o != nil && !IsNil(o.MaxAudioMegabytesPer1Minute) {
		return true
	}

	return false
}

// SetMaxAudioMegabytesPer1Minute gets a reference to the given int32 and assigns it to the MaxAudioMegabytesPer1Minute field.
func (o *ProjectRateLimit) SetMaxAudioMegabytesPer1Minute(v int32) {
	o.MaxAudioMegabytesPer1Minute = &v
}

// GetMaxRequestsPer1Day returns the MaxRequestsPer1Day field value if set, zero value otherwise.
func (o *ProjectRateLimit) GetMaxRequestsPer1Day() int32 {
	if o == nil || IsNil(o.MaxRequestsPer1Day) {
		var ret int32
		return ret
	}
	return *o.MaxRequestsPer1Day
}

// GetMaxRequestsPer1DayOk returns a tuple with the MaxRequestsPer1Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetMaxRequestsPer1DayOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRequestsPer1Day) {
		return nil, false
	}
	return o.MaxRequestsPer1Day, true
}

// HasMaxRequestsPer1Day returns a boolean if a field has been set.
func (o *ProjectRateLimit) HasMaxRequestsPer1Day() bool {
	if o != nil && !IsNil(o.MaxRequestsPer1Day) {
		return true
	}

	return false
}

// SetMaxRequestsPer1Day gets a reference to the given int32 and assigns it to the MaxRequestsPer1Day field.
func (o *ProjectRateLimit) SetMaxRequestsPer1Day(v int32) {
	o.MaxRequestsPer1Day = &v
}

// GetBatch1DayMaxInputTokens returns the Batch1DayMaxInputTokens field value if set, zero value otherwise.
func (o *ProjectRateLimit) GetBatch1DayMaxInputTokens() int32 {
	if o == nil || IsNil(o.Batch1DayMaxInputTokens) {
		var ret int32
		return ret
	}
	return *o.Batch1DayMaxInputTokens
}

// GetBatch1DayMaxInputTokensOk returns a tuple with the Batch1DayMaxInputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRateLimit) GetBatch1DayMaxInputTokensOk() (*int32, bool) {
	if o == nil || IsNil(o.Batch1DayMaxInputTokens) {
		return nil, false
	}
	return o.Batch1DayMaxInputTokens, true
}

// HasBatch1DayMaxInputTokens returns a boolean if a field has been set.
func (o *ProjectRateLimit) HasBatch1DayMaxInputTokens() bool {
	if o != nil && !IsNil(o.Batch1DayMaxInputTokens) {
		return true
	}

	return false
}

// SetBatch1DayMaxInputTokens gets a reference to the given int32 and assigns it to the Batch1DayMaxInputTokens field.
func (o *ProjectRateLimit) SetBatch1DayMaxInputTokens(v int32) {
	o.Batch1DayMaxInputTokens = &v
}

func (o ProjectRateLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectRateLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["model"] = o.Model
	toSerialize["max_requests_per_1_minute"] = o.MaxRequestsPer1Minute
	toSerialize["max_tokens_per_1_minute"] = o.MaxTokensPer1Minute
	if !IsNil(o.MaxImagesPer1Minute) {
		toSerialize["max_images_per_1_minute"] = o.MaxImagesPer1Minute
	}
	if !IsNil(o.MaxAudioMegabytesPer1Minute) {
		toSerialize["max_audio_megabytes_per_1_minute"] = o.MaxAudioMegabytesPer1Minute
	}
	if !IsNil(o.MaxRequestsPer1Day) {
		toSerialize["max_requests_per_1_day"] = o.MaxRequestsPer1Day
	}
	if !IsNil(o.Batch1DayMaxInputTokens) {
		toSerialize["batch_1_day_max_input_tokens"] = o.Batch1DayMaxInputTokens
	}
	return toSerialize, nil
}

func (o *ProjectRateLimit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"model",
		"max_requests_per_1_minute",
		"max_tokens_per_1_minute",
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

	varProjectRateLimit := _ProjectRateLimit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectRateLimit)

	if err != nil {
		return err
	}

	*o = ProjectRateLimit(varProjectRateLimit)

	return err
}

type NullableProjectRateLimit struct {
	value *ProjectRateLimit
	isSet bool
}

func (v NullableProjectRateLimit) Get() *ProjectRateLimit {
	return v.value
}

func (v *NullableProjectRateLimit) Set(val *ProjectRateLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRateLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRateLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRateLimit(val *ProjectRateLimit) *NullableProjectRateLimit {
	return &NullableProjectRateLimit{value: val, isSet: true}
}

func (v NullableProjectRateLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRateLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


