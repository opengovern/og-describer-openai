/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RealtimeSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeSession{}

// RealtimeSession Realtime session object configuration.
type RealtimeSession struct {
	// The set of modalities the model can respond with. To disable audio, set this to [\"text\"]. 
	Modalities []string `json:"modalities,omitempty"`
	// The default system instructions (i.e. system message) prepended to model  calls. This field allows the client to guide the model on desired  responses. The model can be instructed on response content and format,  (e.g. \"be extremely succinct\", \"act friendly\", \"here are examples of good  responses\") and on audio behavior (e.g. \"talk quickly\", \"inject emotion  into your voice\", \"laugh frequently\"). The instructions are not guaranteed  to be followed by the model, but they provide guidance to the model on the  desired behavior.  Note that the server sets default instructions which will be used if this  field is not set and are visible in the `session.created` event at the  start of the session. 
	Instructions *string `json:"instructions,omitempty"`
	// The voice the model uses to respond. Supported voices are `alloy`, `ash`, `ballad`, `coral`, `echo`, `sage`, `shimmer`, and `verse`. Cannot be  changed once the model has responded with audio at least once. 
	Voice *string `json:"voice,omitempty"`
	// The format of input audio. Options are `pcm16`, `g711_ulaw`, or `g711_alaw`. 
	InputAudioFormat *string `json:"input_audio_format,omitempty"`
	// The format of output audio. Options are `pcm16`, `g711_ulaw`, or `g711_alaw`. 
	OutputAudioFormat *string `json:"output_audio_format,omitempty"`
	InputAudioTranscription *RealtimeSessionInputAudioTranscription `json:"input_audio_transcription,omitempty"`
	TurnDetection *RealtimeSessionTurnDetection `json:"turn_detection,omitempty"`
	// Tools (functions) available to the model.
	Tools []RealtimeSessionToolsInner `json:"tools,omitempty"`
	// How the model chooses tools. Options are `auto`, `none`, `required`, or  specify a function. 
	ToolChoice *string `json:"tool_choice,omitempty"`
	// Sampling temperature for the model, limited to [0.6, 1.2]. Defaults to 0.8. 
	Temperature *float32 `json:"temperature,omitempty"`
	MaxResponseOutputTokens *RealtimeSessionMaxResponseOutputTokens `json:"max_response_output_tokens,omitempty"`
}

// NewRealtimeSession instantiates a new RealtimeSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeSession() *RealtimeSession {
	this := RealtimeSession{}
	return &this
}

// NewRealtimeSessionWithDefaults instantiates a new RealtimeSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeSessionWithDefaults() *RealtimeSession {
	this := RealtimeSession{}
	return &this
}

// GetModalities returns the Modalities field value if set, zero value otherwise.
func (o *RealtimeSession) GetModalities() []string {
	if o == nil || IsNil(o.Modalities) {
		var ret []string
		return ret
	}
	return o.Modalities
}

// GetModalitiesOk returns a tuple with the Modalities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetModalitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Modalities) {
		return nil, false
	}
	return o.Modalities, true
}

// HasModalities returns a boolean if a field has been set.
func (o *RealtimeSession) HasModalities() bool {
	if o != nil && !IsNil(o.Modalities) {
		return true
	}

	return false
}

// SetModalities gets a reference to the given []string and assigns it to the Modalities field.
func (o *RealtimeSession) SetModalities(v []string) {
	o.Modalities = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *RealtimeSession) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *RealtimeSession) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *RealtimeSession) SetInstructions(v string) {
	o.Instructions = &v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *RealtimeSession) GetVoice() string {
	if o == nil || IsNil(o.Voice) {
		var ret string
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetVoiceOk() (*string, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *RealtimeSession) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given string and assigns it to the Voice field.
func (o *RealtimeSession) SetVoice(v string) {
	o.Voice = &v
}

// GetInputAudioFormat returns the InputAudioFormat field value if set, zero value otherwise.
func (o *RealtimeSession) GetInputAudioFormat() string {
	if o == nil || IsNil(o.InputAudioFormat) {
		var ret string
		return ret
	}
	return *o.InputAudioFormat
}

// GetInputAudioFormatOk returns a tuple with the InputAudioFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetInputAudioFormatOk() (*string, bool) {
	if o == nil || IsNil(o.InputAudioFormat) {
		return nil, false
	}
	return o.InputAudioFormat, true
}

// HasInputAudioFormat returns a boolean if a field has been set.
func (o *RealtimeSession) HasInputAudioFormat() bool {
	if o != nil && !IsNil(o.InputAudioFormat) {
		return true
	}

	return false
}

// SetInputAudioFormat gets a reference to the given string and assigns it to the InputAudioFormat field.
func (o *RealtimeSession) SetInputAudioFormat(v string) {
	o.InputAudioFormat = &v
}

// GetOutputAudioFormat returns the OutputAudioFormat field value if set, zero value otherwise.
func (o *RealtimeSession) GetOutputAudioFormat() string {
	if o == nil || IsNil(o.OutputAudioFormat) {
		var ret string
		return ret
	}
	return *o.OutputAudioFormat
}

// GetOutputAudioFormatOk returns a tuple with the OutputAudioFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetOutputAudioFormatOk() (*string, bool) {
	if o == nil || IsNil(o.OutputAudioFormat) {
		return nil, false
	}
	return o.OutputAudioFormat, true
}

// HasOutputAudioFormat returns a boolean if a field has been set.
func (o *RealtimeSession) HasOutputAudioFormat() bool {
	if o != nil && !IsNil(o.OutputAudioFormat) {
		return true
	}

	return false
}

// SetOutputAudioFormat gets a reference to the given string and assigns it to the OutputAudioFormat field.
func (o *RealtimeSession) SetOutputAudioFormat(v string) {
	o.OutputAudioFormat = &v
}

// GetInputAudioTranscription returns the InputAudioTranscription field value if set, zero value otherwise.
func (o *RealtimeSession) GetInputAudioTranscription() RealtimeSessionInputAudioTranscription {
	if o == nil || IsNil(o.InputAudioTranscription) {
		var ret RealtimeSessionInputAudioTranscription
		return ret
	}
	return *o.InputAudioTranscription
}

// GetInputAudioTranscriptionOk returns a tuple with the InputAudioTranscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetInputAudioTranscriptionOk() (*RealtimeSessionInputAudioTranscription, bool) {
	if o == nil || IsNil(o.InputAudioTranscription) {
		return nil, false
	}
	return o.InputAudioTranscription, true
}

// HasInputAudioTranscription returns a boolean if a field has been set.
func (o *RealtimeSession) HasInputAudioTranscription() bool {
	if o != nil && !IsNil(o.InputAudioTranscription) {
		return true
	}

	return false
}

// SetInputAudioTranscription gets a reference to the given RealtimeSessionInputAudioTranscription and assigns it to the InputAudioTranscription field.
func (o *RealtimeSession) SetInputAudioTranscription(v RealtimeSessionInputAudioTranscription) {
	o.InputAudioTranscription = &v
}

// GetTurnDetection returns the TurnDetection field value if set, zero value otherwise.
func (o *RealtimeSession) GetTurnDetection() RealtimeSessionTurnDetection {
	if o == nil || IsNil(o.TurnDetection) {
		var ret RealtimeSessionTurnDetection
		return ret
	}
	return *o.TurnDetection
}

// GetTurnDetectionOk returns a tuple with the TurnDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetTurnDetectionOk() (*RealtimeSessionTurnDetection, bool) {
	if o == nil || IsNil(o.TurnDetection) {
		return nil, false
	}
	return o.TurnDetection, true
}

// HasTurnDetection returns a boolean if a field has been set.
func (o *RealtimeSession) HasTurnDetection() bool {
	if o != nil && !IsNil(o.TurnDetection) {
		return true
	}

	return false
}

// SetTurnDetection gets a reference to the given RealtimeSessionTurnDetection and assigns it to the TurnDetection field.
func (o *RealtimeSession) SetTurnDetection(v RealtimeSessionTurnDetection) {
	o.TurnDetection = &v
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *RealtimeSession) GetTools() []RealtimeSessionToolsInner {
	if o == nil || IsNil(o.Tools) {
		var ret []RealtimeSessionToolsInner
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetToolsOk() ([]RealtimeSessionToolsInner, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *RealtimeSession) HasTools() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []RealtimeSessionToolsInner and assigns it to the Tools field.
func (o *RealtimeSession) SetTools(v []RealtimeSessionToolsInner) {
	o.Tools = v
}

// GetToolChoice returns the ToolChoice field value if set, zero value otherwise.
func (o *RealtimeSession) GetToolChoice() string {
	if o == nil || IsNil(o.ToolChoice) {
		var ret string
		return ret
	}
	return *o.ToolChoice
}

// GetToolChoiceOk returns a tuple with the ToolChoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetToolChoiceOk() (*string, bool) {
	if o == nil || IsNil(o.ToolChoice) {
		return nil, false
	}
	return o.ToolChoice, true
}

// HasToolChoice returns a boolean if a field has been set.
func (o *RealtimeSession) HasToolChoice() bool {
	if o != nil && !IsNil(o.ToolChoice) {
		return true
	}

	return false
}

// SetToolChoice gets a reference to the given string and assigns it to the ToolChoice field.
func (o *RealtimeSession) SetToolChoice(v string) {
	o.ToolChoice = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *RealtimeSession) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature) {
		var ret float32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetTemperatureOk() (*float32, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *RealtimeSession) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given float32 and assigns it to the Temperature field.
func (o *RealtimeSession) SetTemperature(v float32) {
	o.Temperature = &v
}

// GetMaxResponseOutputTokens returns the MaxResponseOutputTokens field value if set, zero value otherwise.
func (o *RealtimeSession) GetMaxResponseOutputTokens() RealtimeSessionMaxResponseOutputTokens {
	if o == nil || IsNil(o.MaxResponseOutputTokens) {
		var ret RealtimeSessionMaxResponseOutputTokens
		return ret
	}
	return *o.MaxResponseOutputTokens
}

// GetMaxResponseOutputTokensOk returns a tuple with the MaxResponseOutputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSession) GetMaxResponseOutputTokensOk() (*RealtimeSessionMaxResponseOutputTokens, bool) {
	if o == nil || IsNil(o.MaxResponseOutputTokens) {
		return nil, false
	}
	return o.MaxResponseOutputTokens, true
}

// HasMaxResponseOutputTokens returns a boolean if a field has been set.
func (o *RealtimeSession) HasMaxResponseOutputTokens() bool {
	if o != nil && !IsNil(o.MaxResponseOutputTokens) {
		return true
	}

	return false
}

// SetMaxResponseOutputTokens gets a reference to the given RealtimeSessionMaxResponseOutputTokens and assigns it to the MaxResponseOutputTokens field.
func (o *RealtimeSession) SetMaxResponseOutputTokens(v RealtimeSessionMaxResponseOutputTokens) {
	o.MaxResponseOutputTokens = &v
}

func (o RealtimeSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Modalities) {
		toSerialize["modalities"] = o.Modalities
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	if !IsNil(o.InputAudioFormat) {
		toSerialize["input_audio_format"] = o.InputAudioFormat
	}
	if !IsNil(o.OutputAudioFormat) {
		toSerialize["output_audio_format"] = o.OutputAudioFormat
	}
	if !IsNil(o.InputAudioTranscription) {
		toSerialize["input_audio_transcription"] = o.InputAudioTranscription
	}
	if !IsNil(o.TurnDetection) {
		toSerialize["turn_detection"] = o.TurnDetection
	}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.ToolChoice) {
		toSerialize["tool_choice"] = o.ToolChoice
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !IsNil(o.MaxResponseOutputTokens) {
		toSerialize["max_response_output_tokens"] = o.MaxResponseOutputTokens
	}
	return toSerialize, nil
}

type NullableRealtimeSession struct {
	value *RealtimeSession
	isSet bool
}

func (v NullableRealtimeSession) Get() *RealtimeSession {
	return v.value
}

func (v *NullableRealtimeSession) Set(val *RealtimeSession) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeSession) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeSession(val *RealtimeSession) *NullableRealtimeSession {
	return &NullableRealtimeSession{value: val, isSet: true}
}

func (v NullableRealtimeSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


