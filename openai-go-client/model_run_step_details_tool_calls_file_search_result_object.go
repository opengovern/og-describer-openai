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

// checks if the RunStepDetailsToolCallsFileSearchResultObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsToolCallsFileSearchResultObject{}

// RunStepDetailsToolCallsFileSearchResultObject A result instance of the file search.
type RunStepDetailsToolCallsFileSearchResultObject struct {
	// The ID of the file that result was found in.
	FileId string `json:"file_id"`
	// The name of the file that result was found in.
	FileName string `json:"file_name"`
	// The score of the result. All values must be a floating point number between 0 and 1.
	Score float32 `json:"score"`
	// The content of the result that was found. The content is only included if requested via the include query parameter.
	Content []RunStepDetailsToolCallsFileSearchResultObjectContentInner `json:"content,omitempty"`
}

type _RunStepDetailsToolCallsFileSearchResultObject RunStepDetailsToolCallsFileSearchResultObject

// NewRunStepDetailsToolCallsFileSearchResultObject instantiates a new RunStepDetailsToolCallsFileSearchResultObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsToolCallsFileSearchResultObject(fileId string, fileName string, score float32) *RunStepDetailsToolCallsFileSearchResultObject {
	this := RunStepDetailsToolCallsFileSearchResultObject{}
	this.FileId = fileId
	this.FileName = fileName
	this.Score = score
	return &this
}

// NewRunStepDetailsToolCallsFileSearchResultObjectWithDefaults instantiates a new RunStepDetailsToolCallsFileSearchResultObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsToolCallsFileSearchResultObjectWithDefaults() *RunStepDetailsToolCallsFileSearchResultObject {
	this := RunStepDetailsToolCallsFileSearchResultObject{}
	return &this
}

// GetFileId returns the FileId field value
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *RunStepDetailsToolCallsFileSearchResultObject) SetFileId(v string) {
	o.FileId = v
}

// GetFileName returns the FileName field value
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *RunStepDetailsToolCallsFileSearchResultObject) SetFileName(v string) {
	o.FileName = v
}

// GetScore returns the Score field value
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *RunStepDetailsToolCallsFileSearchResultObject) SetScore(v float32) {
	o.Score = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetContent() []RunStepDetailsToolCallsFileSearchResultObjectContentInner {
	if o == nil || IsNil(o.Content) {
		var ret []RunStepDetailsToolCallsFileSearchResultObjectContentInner
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchResultObject) GetContentOk() ([]RunStepDetailsToolCallsFileSearchResultObjectContentInner, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *RunStepDetailsToolCallsFileSearchResultObject) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []RunStepDetailsToolCallsFileSearchResultObjectContentInner and assigns it to the Content field.
func (o *RunStepDetailsToolCallsFileSearchResultObject) SetContent(v []RunStepDetailsToolCallsFileSearchResultObjectContentInner) {
	o.Content = v
}

func (o RunStepDetailsToolCallsFileSearchResultObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsToolCallsFileSearchResultObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_id"] = o.FileId
	toSerialize["file_name"] = o.FileName
	toSerialize["score"] = o.Score
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

func (o *RunStepDetailsToolCallsFileSearchResultObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_id",
		"file_name",
		"score",
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

	varRunStepDetailsToolCallsFileSearchResultObject := _RunStepDetailsToolCallsFileSearchResultObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDetailsToolCallsFileSearchResultObject)

	if err != nil {
		return err
	}

	*o = RunStepDetailsToolCallsFileSearchResultObject(varRunStepDetailsToolCallsFileSearchResultObject)

	return err
}

type NullableRunStepDetailsToolCallsFileSearchResultObject struct {
	value *RunStepDetailsToolCallsFileSearchResultObject
	isSet bool
}

func (v NullableRunStepDetailsToolCallsFileSearchResultObject) Get() *RunStepDetailsToolCallsFileSearchResultObject {
	return v.value
}

func (v *NullableRunStepDetailsToolCallsFileSearchResultObject) Set(val *RunStepDetailsToolCallsFileSearchResultObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsToolCallsFileSearchResultObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsToolCallsFileSearchResultObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsToolCallsFileSearchResultObject(val *RunStepDetailsToolCallsFileSearchResultObject) *NullableRunStepDetailsToolCallsFileSearchResultObject {
	return &NullableRunStepDetailsToolCallsFileSearchResultObject{value: val, isSet: true}
}

func (v NullableRunStepDetailsToolCallsFileSearchResultObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsToolCallsFileSearchResultObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


