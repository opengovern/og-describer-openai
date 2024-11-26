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

// checks if the CreateImageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateImageRequest{}

// CreateImageRequest struct for CreateImageRequest
type CreateImageRequest struct {
	// A text description of the desired image(s). The maximum length is 1000 characters for `dall-e-2` and 4000 characters for `dall-e-3`.
	Prompt string `json:"prompt"`
	Model NullableCreateImageRequestModel `json:"model,omitempty"`
	// The number of images to generate. Must be between 1 and 10. For `dall-e-3`, only `n=1` is supported.
	N NullableInt32 `json:"n,omitempty"`
	// The quality of the image that will be generated. `hd` creates images with finer details and greater consistency across the image. This param is only supported for `dall-e-3`.
	Quality *string `json:"quality,omitempty"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`. URLs are only valid for 60 minutes after the image has been generated.
	ResponseFormat NullableString `json:"response_format,omitempty"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024` for `dall-e-2`. Must be one of `1024x1024`, `1792x1024`, or `1024x1792` for `dall-e-3` models.
	Size NullableString `json:"size,omitempty"`
	// The style of the generated images. Must be one of `vivid` or `natural`. Vivid causes the model to lean towards generating hyper-real and dramatic images. Natural causes the model to produce more natural, less hyper-real looking images. This param is only supported for `dall-e-3`.
	Style NullableString `json:"style,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices#end-user-ids). 
	User *string `json:"user,omitempty"`
}

type _CreateImageRequest CreateImageRequest

// NewCreateImageRequest instantiates a new CreateImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageRequest(prompt string) *CreateImageRequest {
	this := CreateImageRequest{}
	this.Prompt = prompt
	var model CreateImageRequestModel = dall-e-2
	this.Model = *NewNullableCreateImageRequestModel(&model)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var quality string = "standard"
	this.Quality = &quality
	var responseFormat RESPONSE_FORMAT = "url"
	this.ResponseFormat = *NewNullableString(&responseFormat)
	var size SIZE = "1024x1024"
	this.Size = *NewNullableString(&size)
	var style STYLE = "vivid"
	this.Style = *NewNullableString(&style)
	return &this
}

// NewCreateImageRequestWithDefaults instantiates a new CreateImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageRequestWithDefaults() *CreateImageRequest {
	this := CreateImageRequest{}
	var model CreateImageRequestModel = dall-e-2
	this.Model = *NewNullableCreateImageRequestModel(&model)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var quality string = "standard"
	this.Quality = &quality
	var responseFormat RESPONSE_FORMAT = "url"
	this.ResponseFormat = *NewNullableString(&responseFormat)
	var size SIZE = "1024x1024"
	this.Size = *NewNullableString(&size)
	var style STYLE = "vivid"
	this.Style = *NewNullableString(&style)
	return &this
}

// GetPrompt returns the Prompt field value
func (o *CreateImageRequest) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *CreateImageRequest) SetPrompt(v string) {
	o.Prompt = v
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetModel() CreateImageRequestModel {
	if o == nil || IsNil(o.Model.Get()) {
		var ret CreateImageRequestModel
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetModelOk() (*CreateImageRequestModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *CreateImageRequest) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableCreateImageRequestModel and assigns it to the Model field.
func (o *CreateImageRequest) SetModel(v CreateImageRequestModel) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *CreateImageRequest) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *CreateImageRequest) UnsetModel() {
	o.Model.Unset()
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetN() int32 {
	if o == nil || IsNil(o.N.Get()) {
		var ret int32
		return ret
	}
	return *o.N.Get()
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.N.Get(), o.N.IsSet()
}

// HasN returns a boolean if a field has been set.
func (o *CreateImageRequest) HasN() bool {
	if o != nil && o.N.IsSet() {
		return true
	}

	return false
}

// SetN gets a reference to the given NullableInt32 and assigns it to the N field.
func (o *CreateImageRequest) SetN(v int32) {
	o.N.Set(&v)
}
// SetNNil sets the value for N to be an explicit nil
func (o *CreateImageRequest) SetNNil() {
	o.N.Set(nil)
}

// UnsetN ensures that no value is present for N, not even an explicit nil
func (o *CreateImageRequest) UnsetN() {
	o.N.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *CreateImageRequest) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *CreateImageRequest) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *CreateImageRequest) SetQuality(v string) {
	o.Quality = &v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetResponseFormat() string {
	if o == nil || IsNil(o.ResponseFormat.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseFormat.Get()
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetResponseFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseFormat.Get(), o.ResponseFormat.IsSet()
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *CreateImageRequest) HasResponseFormat() bool {
	if o != nil && o.ResponseFormat.IsSet() {
		return true
	}

	return false
}

// SetResponseFormat gets a reference to the given NullableString and assigns it to the ResponseFormat field.
func (o *CreateImageRequest) SetResponseFormat(v string) {
	o.ResponseFormat.Set(&v)
}
// SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil
func (o *CreateImageRequest) SetResponseFormatNil() {
	o.ResponseFormat.Set(nil)
}

// UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
func (o *CreateImageRequest) UnsetResponseFormat() {
	o.ResponseFormat.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetSize() string {
	if o == nil || IsNil(o.Size.Get()) {
		var ret string
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *CreateImageRequest) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableString and assigns it to the Size field.
func (o *CreateImageRequest) SetSize(v string) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *CreateImageRequest) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *CreateImageRequest) UnsetSize() {
	o.Size.Unset()
}

// GetStyle returns the Style field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetStyle() string {
	if o == nil || IsNil(o.Style.Get()) {
		var ret string
		return ret
	}
	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetStyleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// HasStyle returns a boolean if a field has been set.
func (o *CreateImageRequest) HasStyle() bool {
	if o != nil && o.Style.IsSet() {
		return true
	}

	return false
}

// SetStyle gets a reference to the given NullableString and assigns it to the Style field.
func (o *CreateImageRequest) SetStyle(v string) {
	o.Style.Set(&v)
}
// SetStyleNil sets the value for Style to be an explicit nil
func (o *CreateImageRequest) SetStyleNil() {
	o.Style.Set(nil)
}

// UnsetStyle ensures that no value is present for Style, not even an explicit nil
func (o *CreateImageRequest) UnsetStyle() {
	o.Style.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateImageRequest) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateImageRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateImageRequest) SetUser(v string) {
	o.User = &v
}

func (o CreateImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateImageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prompt"] = o.Prompt
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.N.IsSet() {
		toSerialize["n"] = o.N.Get()
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.ResponseFormat.IsSet() {
		toSerialize["response_format"] = o.ResponseFormat.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Style.IsSet() {
		toSerialize["style"] = o.Style.Get()
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

func (o *CreateImageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prompt",
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

	varCreateImageRequest := _CreateImageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateImageRequest)

	if err != nil {
		return err
	}

	*o = CreateImageRequest(varCreateImageRequest)

	return err
}

type NullableCreateImageRequest struct {
	value *CreateImageRequest
	isSet bool
}

func (v NullableCreateImageRequest) Get() *CreateImageRequest {
	return v.value
}

func (v *NullableCreateImageRequest) Set(val *CreateImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageRequest(val *CreateImageRequest) *NullableCreateImageRequest {
	return &NullableCreateImageRequest{value: val, isSet: true}
}

func (v NullableCreateImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


