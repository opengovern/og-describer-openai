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

// checks if the CreateChatCompletionResponseChoicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatCompletionResponseChoicesInner{}

// CreateChatCompletionResponseChoicesInner struct for CreateChatCompletionResponseChoicesInner
type CreateChatCompletionResponseChoicesInner struct {
	// The reason the model stopped generating tokens. This will be `stop` if the model hit a natural stop point or a provided stop sequence, `length` if the maximum number of tokens specified in the request was reached, `content_filter` if content was omitted due to a flag from our content filters, `tool_calls` if the model called a tool, or `function_call` (deprecated) if the model called a function. 
	FinishReason string `json:"finish_reason"`
	// The index of the choice in the list of choices.
	Index int32 `json:"index"`
	Message ChatCompletionResponseMessage `json:"message"`
	Logprobs NullableCreateChatCompletionResponseChoicesInnerLogprobs `json:"logprobs"`
}

type _CreateChatCompletionResponseChoicesInner CreateChatCompletionResponseChoicesInner

// NewCreateChatCompletionResponseChoicesInner instantiates a new CreateChatCompletionResponseChoicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatCompletionResponseChoicesInner(finishReason string, index int32, message ChatCompletionResponseMessage, logprobs NullableCreateChatCompletionResponseChoicesInnerLogprobs) *CreateChatCompletionResponseChoicesInner {
	this := CreateChatCompletionResponseChoicesInner{}
	this.FinishReason = finishReason
	this.Index = index
	this.Message = message
	this.Logprobs = logprobs
	return &this
}

// NewCreateChatCompletionResponseChoicesInnerWithDefaults instantiates a new CreateChatCompletionResponseChoicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatCompletionResponseChoicesInnerWithDefaults() *CreateChatCompletionResponseChoicesInner {
	this := CreateChatCompletionResponseChoicesInner{}
	return &this
}

// GetFinishReason returns the FinishReason field value
func (o *CreateChatCompletionResponseChoicesInner) GetFinishReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishReason
}

// GetFinishReasonOk returns a tuple with the FinishReason field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionResponseChoicesInner) GetFinishReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishReason, true
}

// SetFinishReason sets field value
func (o *CreateChatCompletionResponseChoicesInner) SetFinishReason(v string) {
	o.FinishReason = v
}

// GetIndex returns the Index field value
func (o *CreateChatCompletionResponseChoicesInner) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionResponseChoicesInner) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *CreateChatCompletionResponseChoicesInner) SetIndex(v int32) {
	o.Index = v
}

// GetMessage returns the Message field value
func (o *CreateChatCompletionResponseChoicesInner) GetMessage() ChatCompletionResponseMessage {
	if o == nil {
		var ret ChatCompletionResponseMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionResponseChoicesInner) GetMessageOk() (*ChatCompletionResponseMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateChatCompletionResponseChoicesInner) SetMessage(v ChatCompletionResponseMessage) {
	o.Message = v
}

// GetLogprobs returns the Logprobs field value
// If the value is explicit nil, the zero value for CreateChatCompletionResponseChoicesInnerLogprobs will be returned
func (o *CreateChatCompletionResponseChoicesInner) GetLogprobs() CreateChatCompletionResponseChoicesInnerLogprobs {
	if o == nil || o.Logprobs.Get() == nil {
		var ret CreateChatCompletionResponseChoicesInnerLogprobs
		return ret
	}

	return *o.Logprobs.Get()
}

// GetLogprobsOk returns a tuple with the Logprobs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionResponseChoicesInner) GetLogprobsOk() (*CreateChatCompletionResponseChoicesInnerLogprobs, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logprobs.Get(), o.Logprobs.IsSet()
}

// SetLogprobs sets field value
func (o *CreateChatCompletionResponseChoicesInner) SetLogprobs(v CreateChatCompletionResponseChoicesInnerLogprobs) {
	o.Logprobs.Set(&v)
}

func (o CreateChatCompletionResponseChoicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatCompletionResponseChoicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["finish_reason"] = o.FinishReason
	toSerialize["index"] = o.Index
	toSerialize["message"] = o.Message
	toSerialize["logprobs"] = o.Logprobs.Get()
	return toSerialize, nil
}

func (o *CreateChatCompletionResponseChoicesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"finish_reason",
		"index",
		"message",
		"logprobs",
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

	varCreateChatCompletionResponseChoicesInner := _CreateChatCompletionResponseChoicesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatCompletionResponseChoicesInner)

	if err != nil {
		return err
	}

	*o = CreateChatCompletionResponseChoicesInner(varCreateChatCompletionResponseChoicesInner)

	return err
}

type NullableCreateChatCompletionResponseChoicesInner struct {
	value *CreateChatCompletionResponseChoicesInner
	isSet bool
}

func (v NullableCreateChatCompletionResponseChoicesInner) Get() *CreateChatCompletionResponseChoicesInner {
	return v.value
}

func (v *NullableCreateChatCompletionResponseChoicesInner) Set(val *CreateChatCompletionResponseChoicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatCompletionResponseChoicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatCompletionResponseChoicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatCompletionResponseChoicesInner(val *CreateChatCompletionResponseChoicesInner) *NullableCreateChatCompletionResponseChoicesInner {
	return &NullableCreateChatCompletionResponseChoicesInner{value: val, isSet: true}
}

func (v NullableCreateChatCompletionResponseChoicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatCompletionResponseChoicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


