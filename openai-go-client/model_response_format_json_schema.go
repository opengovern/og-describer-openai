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

// checks if the ResponseFormatJsonSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseFormatJsonSchema{}

// ResponseFormatJsonSchema struct for ResponseFormatJsonSchema
type ResponseFormatJsonSchema struct {
	// The type of response format being defined: `json_schema`
	Type string `json:"type"`
	JsonSchema ResponseFormatJsonSchemaJsonSchema `json:"json_schema"`
}

type _ResponseFormatJsonSchema ResponseFormatJsonSchema

// NewResponseFormatJsonSchema instantiates a new ResponseFormatJsonSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseFormatJsonSchema(type_ string, jsonSchema ResponseFormatJsonSchemaJsonSchema) *ResponseFormatJsonSchema {
	this := ResponseFormatJsonSchema{}
	this.Type = type_
	this.JsonSchema = jsonSchema
	return &this
}

// NewResponseFormatJsonSchemaWithDefaults instantiates a new ResponseFormatJsonSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseFormatJsonSchemaWithDefaults() *ResponseFormatJsonSchema {
	this := ResponseFormatJsonSchema{}
	return &this
}

// GetType returns the Type field value
func (o *ResponseFormatJsonSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResponseFormatJsonSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResponseFormatJsonSchema) SetType(v string) {
	o.Type = v
}

// GetJsonSchema returns the JsonSchema field value
func (o *ResponseFormatJsonSchema) GetJsonSchema() ResponseFormatJsonSchemaJsonSchema {
	if o == nil {
		var ret ResponseFormatJsonSchemaJsonSchema
		return ret
	}

	return o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value
// and a boolean to check if the value has been set.
func (o *ResponseFormatJsonSchema) GetJsonSchemaOk() (*ResponseFormatJsonSchemaJsonSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonSchema, true
}

// SetJsonSchema sets field value
func (o *ResponseFormatJsonSchema) SetJsonSchema(v ResponseFormatJsonSchemaJsonSchema) {
	o.JsonSchema = v
}

func (o ResponseFormatJsonSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseFormatJsonSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["json_schema"] = o.JsonSchema
	return toSerialize, nil
}

func (o *ResponseFormatJsonSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"json_schema",
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

	varResponseFormatJsonSchema := _ResponseFormatJsonSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseFormatJsonSchema)

	if err != nil {
		return err
	}

	*o = ResponseFormatJsonSchema(varResponseFormatJsonSchema)

	return err
}

type NullableResponseFormatJsonSchema struct {
	value *ResponseFormatJsonSchema
	isSet bool
}

func (v NullableResponseFormatJsonSchema) Get() *ResponseFormatJsonSchema {
	return v.value
}

func (v *NullableResponseFormatJsonSchema) Set(val *ResponseFormatJsonSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseFormatJsonSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseFormatJsonSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseFormatJsonSchema(val *ResponseFormatJsonSchema) *NullableResponseFormatJsonSchema {
	return &NullableResponseFormatJsonSchema{value: val, isSet: true}
}

func (v NullableResponseFormatJsonSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseFormatJsonSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


