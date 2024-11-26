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

// checks if the OpenAIFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenAIFile{}

// OpenAIFile The `File` object represents a document that has been uploaded to OpenAI.
type OpenAIFile struct {
	// The file identifier, which can be referenced in the API endpoints.
	Id string `json:"id"`
	// The size of the file, in bytes.
	Bytes int32 `json:"bytes"`
	// The Unix timestamp (in seconds) for when the file was created.
	CreatedAt int32 `json:"created_at"`
	// The name of the file.
	Filename string `json:"filename"`
	// The object type, which is always `file`.
	Object string `json:"object"`
	// The intended purpose of the file. Supported values are `assistants`, `assistants_output`, `batch`, `batch_output`, `fine-tune`, `fine-tune-results` and `vision`.
	Purpose string `json:"purpose"`
	// Deprecated. The current status of the file, which can be either `uploaded`, `processed`, or `error`.
	// Deprecated
	Status string `json:"status"`
	// Deprecated. For details on why a fine-tuning training file failed validation, see the `error` field on `fine_tuning.job`.
	// Deprecated
	StatusDetails *string `json:"status_details,omitempty"`
}

type _OpenAIFile OpenAIFile

// NewOpenAIFile instantiates a new OpenAIFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenAIFile(id string, bytes int32, createdAt int32, filename string, object string, purpose string, status string) *OpenAIFile {
	this := OpenAIFile{}
	this.Id = id
	this.Bytes = bytes
	this.CreatedAt = createdAt
	this.Filename = filename
	this.Object = object
	this.Purpose = purpose
	this.Status = status
	return &this
}

// NewOpenAIFileWithDefaults instantiates a new OpenAIFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenAIFileWithDefaults() *OpenAIFile {
	this := OpenAIFile{}
	return &this
}

// GetId returns the Id field value
func (o *OpenAIFile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenAIFile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenAIFile) SetId(v string) {
	o.Id = v
}

// GetBytes returns the Bytes field value
func (o *OpenAIFile) GetBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *OpenAIFile) GetBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *OpenAIFile) SetBytes(v int32) {
	o.Bytes = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OpenAIFile) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OpenAIFile) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OpenAIFile) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetFilename returns the Filename field value
func (o *OpenAIFile) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *OpenAIFile) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *OpenAIFile) SetFilename(v string) {
	o.Filename = v
}

// GetObject returns the Object field value
func (o *OpenAIFile) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *OpenAIFile) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *OpenAIFile) SetObject(v string) {
	o.Object = v
}

// GetPurpose returns the Purpose field value
func (o *OpenAIFile) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *OpenAIFile) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *OpenAIFile) SetPurpose(v string) {
	o.Purpose = v
}

// GetStatus returns the Status field value
// Deprecated
func (o *OpenAIFile) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *OpenAIFile) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
// Deprecated
func (o *OpenAIFile) SetStatus(v string) {
	o.Status = v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
// Deprecated
func (o *OpenAIFile) GetStatusDetails() string {
	if o == nil || IsNil(o.StatusDetails) {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *OpenAIFile) GetStatusDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetails) {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *OpenAIFile) HasStatusDetails() bool {
	if o != nil && !IsNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
// Deprecated
func (o *OpenAIFile) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

func (o OpenAIFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenAIFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["bytes"] = o.Bytes
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["filename"] = o.Filename
	toSerialize["object"] = o.Object
	toSerialize["purpose"] = o.Purpose
	toSerialize["status"] = o.Status
	if !IsNil(o.StatusDetails) {
		toSerialize["status_details"] = o.StatusDetails
	}
	return toSerialize, nil
}

func (o *OpenAIFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"bytes",
		"created_at",
		"filename",
		"object",
		"purpose",
		"status",
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

	varOpenAIFile := _OpenAIFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenAIFile)

	if err != nil {
		return err
	}

	*o = OpenAIFile(varOpenAIFile)

	return err
}

type NullableOpenAIFile struct {
	value *OpenAIFile
	isSet bool
}

func (v NullableOpenAIFile) Get() *OpenAIFile {
	return v.value
}

func (v *NullableOpenAIFile) Set(val *OpenAIFile) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenAIFile) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenAIFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenAIFile(val *OpenAIFile) *NullableOpenAIFile {
	return &NullableOpenAIFile{value: val, isSet: true}
}

func (v NullableOpenAIFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenAIFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


