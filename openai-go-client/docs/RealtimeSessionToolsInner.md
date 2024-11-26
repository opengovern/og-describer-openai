# RealtimeSessionToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the tool, i.e. &#x60;function&#x60;. | [optional] 
**Name** | Pointer to **string** | The name of the function. | [optional] 
**Description** | Pointer to **string** | The description of the function, including guidance on when and how  to call it, and guidance about what to tell the user when calling  (if anything).  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | Parameters of the function in JSON Schema. | [optional] 

## Methods

### NewRealtimeSessionToolsInner

`func NewRealtimeSessionToolsInner() *RealtimeSessionToolsInner`

NewRealtimeSessionToolsInner instantiates a new RealtimeSessionToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeSessionToolsInnerWithDefaults

`func NewRealtimeSessionToolsInnerWithDefaults() *RealtimeSessionToolsInner`

NewRealtimeSessionToolsInnerWithDefaults instantiates a new RealtimeSessionToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RealtimeSessionToolsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeSessionToolsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeSessionToolsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealtimeSessionToolsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *RealtimeSessionToolsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealtimeSessionToolsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealtimeSessionToolsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealtimeSessionToolsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RealtimeSessionToolsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RealtimeSessionToolsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RealtimeSessionToolsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RealtimeSessionToolsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *RealtimeSessionToolsInner) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RealtimeSessionToolsInner) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RealtimeSessionToolsInner) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RealtimeSessionToolsInner) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


