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

// checks if the RealtimeSessionTurnDetection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealtimeSessionTurnDetection{}

// RealtimeSessionTurnDetection Configuration for turn detection. Can be set to `null` to turn off. Server  VAD means that the model will detect the start and end of speech based on  audio volume and respond at the end of user speech. 
type RealtimeSessionTurnDetection struct {
	// Type of turn detection, only `server_vad` is currently supported. 
	Type *string `json:"type,omitempty"`
	// Activation threshold for VAD (0.0 to 1.0), this defaults to 0.5. A  higher threshold will require louder audio to activate the model, and  thus might perform better in noisy environments. 
	Threshold *float32 `json:"threshold,omitempty"`
	// Amount of audio to include before the VAD detected speech (in  milliseconds). Defaults to 300ms. 
	PrefixPaddingMs *int32 `json:"prefix_padding_ms,omitempty"`
	// Duration of silence to detect speech stop (in milliseconds). Defaults  to 500ms. With shorter values the model will respond more quickly,  but may jump in on short pauses from the user. 
	SilenceDurationMs *int32 `json:"silence_duration_ms,omitempty"`
}

// NewRealtimeSessionTurnDetection instantiates a new RealtimeSessionTurnDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeSessionTurnDetection() *RealtimeSessionTurnDetection {
	this := RealtimeSessionTurnDetection{}
	return &this
}

// NewRealtimeSessionTurnDetectionWithDefaults instantiates a new RealtimeSessionTurnDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeSessionTurnDetectionWithDefaults() *RealtimeSessionTurnDetection {
	this := RealtimeSessionTurnDetection{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RealtimeSessionTurnDetection) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionTurnDetection) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RealtimeSessionTurnDetection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RealtimeSessionTurnDetection) SetType(v string) {
	o.Type = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RealtimeSessionTurnDetection) GetThreshold() float32 {
	if o == nil || IsNil(o.Threshold) {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionTurnDetection) GetThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RealtimeSessionTurnDetection) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *RealtimeSessionTurnDetection) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetPrefixPaddingMs returns the PrefixPaddingMs field value if set, zero value otherwise.
func (o *RealtimeSessionTurnDetection) GetPrefixPaddingMs() int32 {
	if o == nil || IsNil(o.PrefixPaddingMs) {
		var ret int32
		return ret
	}
	return *o.PrefixPaddingMs
}

// GetPrefixPaddingMsOk returns a tuple with the PrefixPaddingMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionTurnDetection) GetPrefixPaddingMsOk() (*int32, bool) {
	if o == nil || IsNil(o.PrefixPaddingMs) {
		return nil, false
	}
	return o.PrefixPaddingMs, true
}

// HasPrefixPaddingMs returns a boolean if a field has been set.
func (o *RealtimeSessionTurnDetection) HasPrefixPaddingMs() bool {
	if o != nil && !IsNil(o.PrefixPaddingMs) {
		return true
	}

	return false
}

// SetPrefixPaddingMs gets a reference to the given int32 and assigns it to the PrefixPaddingMs field.
func (o *RealtimeSessionTurnDetection) SetPrefixPaddingMs(v int32) {
	o.PrefixPaddingMs = &v
}

// GetSilenceDurationMs returns the SilenceDurationMs field value if set, zero value otherwise.
func (o *RealtimeSessionTurnDetection) GetSilenceDurationMs() int32 {
	if o == nil || IsNil(o.SilenceDurationMs) {
		var ret int32
		return ret
	}
	return *o.SilenceDurationMs
}

// GetSilenceDurationMsOk returns a tuple with the SilenceDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeSessionTurnDetection) GetSilenceDurationMsOk() (*int32, bool) {
	if o == nil || IsNil(o.SilenceDurationMs) {
		return nil, false
	}
	return o.SilenceDurationMs, true
}

// HasSilenceDurationMs returns a boolean if a field has been set.
func (o *RealtimeSessionTurnDetection) HasSilenceDurationMs() bool {
	if o != nil && !IsNil(o.SilenceDurationMs) {
		return true
	}

	return false
}

// SetSilenceDurationMs gets a reference to the given int32 and assigns it to the SilenceDurationMs field.
func (o *RealtimeSessionTurnDetection) SetSilenceDurationMs(v int32) {
	o.SilenceDurationMs = &v
}

func (o RealtimeSessionTurnDetection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealtimeSessionTurnDetection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.PrefixPaddingMs) {
		toSerialize["prefix_padding_ms"] = o.PrefixPaddingMs
	}
	if !IsNil(o.SilenceDurationMs) {
		toSerialize["silence_duration_ms"] = o.SilenceDurationMs
	}
	return toSerialize, nil
}

type NullableRealtimeSessionTurnDetection struct {
	value *RealtimeSessionTurnDetection
	isSet bool
}

func (v NullableRealtimeSessionTurnDetection) Get() *RealtimeSessionTurnDetection {
	return v.value
}

func (v *NullableRealtimeSessionTurnDetection) Set(val *RealtimeSessionTurnDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableRealtimeSessionTurnDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableRealtimeSessionTurnDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealtimeSessionTurnDetection(val *RealtimeSessionTurnDetection) *NullableRealtimeSessionTurnDetection {
	return &NullableRealtimeSessionTurnDetection{value: val, isSet: true}
}

func (v NullableRealtimeSessionTurnDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealtimeSessionTurnDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


