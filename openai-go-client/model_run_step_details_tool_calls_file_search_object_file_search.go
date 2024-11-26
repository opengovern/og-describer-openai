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

// checks if the RunStepDetailsToolCallsFileSearchObjectFileSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsToolCallsFileSearchObjectFileSearch{}

// RunStepDetailsToolCallsFileSearchObjectFileSearch For now, this is always going to be an empty object.
type RunStepDetailsToolCallsFileSearchObjectFileSearch struct {
	RankingOptions *RunStepDetailsToolCallsFileSearchRankingOptionsObject `json:"ranking_options,omitempty"`
	// The results of the file search.
	Results []RunStepDetailsToolCallsFileSearchResultObject `json:"results,omitempty"`
}

// NewRunStepDetailsToolCallsFileSearchObjectFileSearch instantiates a new RunStepDetailsToolCallsFileSearchObjectFileSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsToolCallsFileSearchObjectFileSearch() *RunStepDetailsToolCallsFileSearchObjectFileSearch {
	this := RunStepDetailsToolCallsFileSearchObjectFileSearch{}
	return &this
}

// NewRunStepDetailsToolCallsFileSearchObjectFileSearchWithDefaults instantiates a new RunStepDetailsToolCallsFileSearchObjectFileSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsToolCallsFileSearchObjectFileSearchWithDefaults() *RunStepDetailsToolCallsFileSearchObjectFileSearch {
	this := RunStepDetailsToolCallsFileSearchObjectFileSearch{}
	return &this
}

// GetRankingOptions returns the RankingOptions field value if set, zero value otherwise.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) GetRankingOptions() RunStepDetailsToolCallsFileSearchRankingOptionsObject {
	if o == nil || IsNil(o.RankingOptions) {
		var ret RunStepDetailsToolCallsFileSearchRankingOptionsObject
		return ret
	}
	return *o.RankingOptions
}

// GetRankingOptionsOk returns a tuple with the RankingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) GetRankingOptionsOk() (*RunStepDetailsToolCallsFileSearchRankingOptionsObject, bool) {
	if o == nil || IsNil(o.RankingOptions) {
		return nil, false
	}
	return o.RankingOptions, true
}

// HasRankingOptions returns a boolean if a field has been set.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) HasRankingOptions() bool {
	if o != nil && !IsNil(o.RankingOptions) {
		return true
	}

	return false
}

// SetRankingOptions gets a reference to the given RunStepDetailsToolCallsFileSearchRankingOptionsObject and assigns it to the RankingOptions field.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) SetRankingOptions(v RunStepDetailsToolCallsFileSearchRankingOptionsObject) {
	o.RankingOptions = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) GetResults() []RunStepDetailsToolCallsFileSearchResultObject {
	if o == nil || IsNil(o.Results) {
		var ret []RunStepDetailsToolCallsFileSearchResultObject
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) GetResultsOk() ([]RunStepDetailsToolCallsFileSearchResultObject, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []RunStepDetailsToolCallsFileSearchResultObject and assigns it to the Results field.
func (o *RunStepDetailsToolCallsFileSearchObjectFileSearch) SetResults(v []RunStepDetailsToolCallsFileSearchResultObject) {
	o.Results = v
}

func (o RunStepDetailsToolCallsFileSearchObjectFileSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsToolCallsFileSearchObjectFileSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RankingOptions) {
		toSerialize["ranking_options"] = o.RankingOptions
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableRunStepDetailsToolCallsFileSearchObjectFileSearch struct {
	value *RunStepDetailsToolCallsFileSearchObjectFileSearch
	isSet bool
}

func (v NullableRunStepDetailsToolCallsFileSearchObjectFileSearch) Get() *RunStepDetailsToolCallsFileSearchObjectFileSearch {
	return v.value
}

func (v *NullableRunStepDetailsToolCallsFileSearchObjectFileSearch) Set(val *RunStepDetailsToolCallsFileSearchObjectFileSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsToolCallsFileSearchObjectFileSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsToolCallsFileSearchObjectFileSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsToolCallsFileSearchObjectFileSearch(val *RunStepDetailsToolCallsFileSearchObjectFileSearch) *NullableRunStepDetailsToolCallsFileSearchObjectFileSearch {
	return &NullableRunStepDetailsToolCallsFileSearchObjectFileSearch{value: val, isSet: true}
}

func (v NullableRunStepDetailsToolCallsFileSearchObjectFileSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsToolCallsFileSearchObjectFileSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


