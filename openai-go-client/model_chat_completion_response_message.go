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

// checks if the ChatCompletionResponseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionResponseMessage{}

// ChatCompletionResponseMessage A chat completion message generated by the model.
type ChatCompletionResponseMessage struct {
	// The contents of the message.
	Content NullableString `json:"content"`
	// The refusal message generated by the model.
	Refusal NullableString `json:"refusal"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []ChatCompletionMessageToolCall `json:"tool_calls,omitempty"`
	// The role of the author of this message.
	Role string `json:"role"`
	// Deprecated
	FunctionCall *ChatCompletionResponseMessageFunctionCall `json:"function_call,omitempty"`
	Audio NullableChatCompletionResponseMessageAudio `json:"audio,omitempty"`
}

type _ChatCompletionResponseMessage ChatCompletionResponseMessage

// NewChatCompletionResponseMessage instantiates a new ChatCompletionResponseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionResponseMessage(content NullableString, refusal NullableString, role string) *ChatCompletionResponseMessage {
	this := ChatCompletionResponseMessage{}
	this.Content = content
	this.Refusal = refusal
	this.Role = role
	return &this
}

// NewChatCompletionResponseMessageWithDefaults instantiates a new ChatCompletionResponseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionResponseMessageWithDefaults() *ChatCompletionResponseMessage {
	this := ChatCompletionResponseMessage{}
	return &this
}

// GetContent returns the Content field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChatCompletionResponseMessage) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}

	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionResponseMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// SetContent sets field value
func (o *ChatCompletionResponseMessage) SetContent(v string) {
	o.Content.Set(&v)
}

// GetRefusal returns the Refusal field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChatCompletionResponseMessage) GetRefusal() string {
	if o == nil || o.Refusal.Get() == nil {
		var ret string
		return ret
	}

	return *o.Refusal.Get()
}

// GetRefusalOk returns a tuple with the Refusal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionResponseMessage) GetRefusalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refusal.Get(), o.Refusal.IsSet()
}

// SetRefusal sets field value
func (o *ChatCompletionResponseMessage) SetRefusal(v string) {
	o.Refusal.Set(&v)
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *ChatCompletionResponseMessage) GetToolCalls() []ChatCompletionMessageToolCall {
	if o == nil || IsNil(o.ToolCalls) {
		var ret []ChatCompletionMessageToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatCompletionResponseMessage) GetToolCallsOk() ([]ChatCompletionMessageToolCall, bool) {
	if o == nil || IsNil(o.ToolCalls) {
		return nil, false
	}
	return o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *ChatCompletionResponseMessage) HasToolCalls() bool {
	if o != nil && !IsNil(o.ToolCalls) {
		return true
	}

	return false
}

// SetToolCalls gets a reference to the given []ChatCompletionMessageToolCall and assigns it to the ToolCalls field.
func (o *ChatCompletionResponseMessage) SetToolCalls(v []ChatCompletionMessageToolCall) {
	o.ToolCalls = v
}

// GetRole returns the Role field value
func (o *ChatCompletionResponseMessage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionResponseMessage) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ChatCompletionResponseMessage) SetRole(v string) {
	o.Role = v
}

// GetFunctionCall returns the FunctionCall field value if set, zero value otherwise.
// Deprecated
func (o *ChatCompletionResponseMessage) GetFunctionCall() ChatCompletionResponseMessageFunctionCall {
	if o == nil || IsNil(o.FunctionCall) {
		var ret ChatCompletionResponseMessageFunctionCall
		return ret
	}
	return *o.FunctionCall
}

// GetFunctionCallOk returns a tuple with the FunctionCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ChatCompletionResponseMessage) GetFunctionCallOk() (*ChatCompletionResponseMessageFunctionCall, bool) {
	if o == nil || IsNil(o.FunctionCall) {
		return nil, false
	}
	return o.FunctionCall, true
}

// HasFunctionCall returns a boolean if a field has been set.
func (o *ChatCompletionResponseMessage) HasFunctionCall() bool {
	if o != nil && !IsNil(o.FunctionCall) {
		return true
	}

	return false
}

// SetFunctionCall gets a reference to the given ChatCompletionResponseMessageFunctionCall and assigns it to the FunctionCall field.
// Deprecated
func (o *ChatCompletionResponseMessage) SetFunctionCall(v ChatCompletionResponseMessageFunctionCall) {
	o.FunctionCall = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionResponseMessage) GetAudio() ChatCompletionResponseMessageAudio {
	if o == nil || IsNil(o.Audio.Get()) {
		var ret ChatCompletionResponseMessageAudio
		return ret
	}
	return *o.Audio.Get()
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionResponseMessage) GetAudioOk() (*ChatCompletionResponseMessageAudio, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audio.Get(), o.Audio.IsSet()
}

// HasAudio returns a boolean if a field has been set.
func (o *ChatCompletionResponseMessage) HasAudio() bool {
	if o != nil && o.Audio.IsSet() {
		return true
	}

	return false
}

// SetAudio gets a reference to the given NullableChatCompletionResponseMessageAudio and assigns it to the Audio field.
func (o *ChatCompletionResponseMessage) SetAudio(v ChatCompletionResponseMessageAudio) {
	o.Audio.Set(&v)
}
// SetAudioNil sets the value for Audio to be an explicit nil
func (o *ChatCompletionResponseMessage) SetAudioNil() {
	o.Audio.Set(nil)
}

// UnsetAudio ensures that no value is present for Audio, not even an explicit nil
func (o *ChatCompletionResponseMessage) UnsetAudio() {
	o.Audio.Unset()
}

func (o ChatCompletionResponseMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionResponseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content.Get()
	toSerialize["refusal"] = o.Refusal.Get()
	if !IsNil(o.ToolCalls) {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	toSerialize["role"] = o.Role
	if !IsNil(o.FunctionCall) {
		toSerialize["function_call"] = o.FunctionCall
	}
	if o.Audio.IsSet() {
		toSerialize["audio"] = o.Audio.Get()
	}
	return toSerialize, nil
}

func (o *ChatCompletionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"refusal",
		"role",
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

	varChatCompletionResponseMessage := _ChatCompletionResponseMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionResponseMessage)

	if err != nil {
		return err
	}

	*o = ChatCompletionResponseMessage(varChatCompletionResponseMessage)

	return err
}

type NullableChatCompletionResponseMessage struct {
	value *ChatCompletionResponseMessage
	isSet bool
}

func (v NullableChatCompletionResponseMessage) Get() *ChatCompletionResponseMessage {
	return v.value
}

func (v *NullableChatCompletionResponseMessage) Set(val *ChatCompletionResponseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionResponseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionResponseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionResponseMessage(val *ChatCompletionResponseMessage) *NullableChatCompletionResponseMessage {
	return &NullableChatCompletionResponseMessage{value: val, isSet: true}
}

func (v NullableChatCompletionResponseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionResponseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


