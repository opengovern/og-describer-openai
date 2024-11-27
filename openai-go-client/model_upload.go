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

// checks if the Upload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Upload{}

// Upload The Upload object can accept byte chunks in the form of Parts. 
type Upload struct {
	// The Upload unique identifier, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The Unix timestamp (in seconds) for when the Upload was created.
	CreatedAt int32 `json:"created_at"`
	// The name of the file to be uploaded.
	Filename string `json:"filename"`
	// The intended number of bytes to be uploaded.
	Bytes int32 `json:"bytes"`
	// The intended purpose of the file. [Please refer here](/docs/api-reference/files/object#files/object-purpose) for acceptable values.
	Purpose string `json:"purpose"`
	// The status of the Upload.
	Status string `json:"status"`
	// The Unix timestamp (in seconds) for when the Upload was created.
	ExpiresAt int32 `json:"expires_at"`
	// The object type, which is always \"upload\".
	Object *string `json:"object,omitempty"`
	File *OpenAIFile `json:"file,omitempty"`
}

type _Upload Upload

// NewUpload instantiates a new Upload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpload(id string, createdAt int32, filename string, bytes int32, purpose string, status string, expiresAt int32) *Upload {
	this := Upload{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Filename = filename
	this.Bytes = bytes
	this.Purpose = purpose
	this.Status = status
	this.ExpiresAt = expiresAt
	return &this
}

// NewUploadWithDefaults instantiates a new Upload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadWithDefaults() *Upload {
	this := Upload{}
	return &this
}

// GetId returns the Id field value
func (o *Upload) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Upload) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Upload) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Upload) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Upload) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Upload) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetFilename returns the Filename field value
func (o *Upload) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *Upload) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *Upload) SetFilename(v string) {
	o.Filename = v
}

// GetBytes returns the Bytes field value
func (o *Upload) GetBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *Upload) GetBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *Upload) SetBytes(v int32) {
	o.Bytes = v
}

// GetPurpose returns the Purpose field value
func (o *Upload) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *Upload) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *Upload) SetPurpose(v string) {
	o.Purpose = v
}

// GetStatus returns the Status field value
func (o *Upload) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Upload) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Upload) SetStatus(v string) {
	o.Status = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *Upload) GetExpiresAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *Upload) GetExpiresAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *Upload) SetExpiresAt(v int32) {
	o.ExpiresAt = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Upload) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upload) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Upload) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Upload) SetObject(v string) {
	o.Object = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *Upload) GetFile() OpenAIFile {
	if o == nil || IsNil(o.File) {
		var ret OpenAIFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upload) GetFileOk() (*OpenAIFile, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *Upload) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given OpenAIFile and assigns it to the File field.
func (o *Upload) SetFile(v OpenAIFile) {
	o.File = &v
}

func (o Upload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Upload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["filename"] = o.Filename
	toSerialize["bytes"] = o.Bytes
	toSerialize["purpose"] = o.Purpose
	toSerialize["status"] = o.Status
	toSerialize["expires_at"] = o.ExpiresAt
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}

func (o *Upload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"filename",
		"bytes",
		"purpose",
		"status",
		"expires_at",
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

	varUpload := _Upload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpload)

	if err != nil {
		return err
	}

	*o = Upload(varUpload)

	return err
}

type NullableUpload struct {
	value *Upload
	isSet bool
}

func (v NullableUpload) Get() *Upload {
	return v.value
}

func (v *NullableUpload) Set(val *Upload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpload(val *Upload) *NullableUpload {
	return &NullableUpload{value: val, isSet: true}
}

func (v NullableUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

