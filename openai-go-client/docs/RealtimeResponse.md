# RealtimeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the response. | [optional] 
**Object** | Pointer to **string** | The object type, must be &#x60;realtime.response&#x60;. | [optional] 
**Status** | Pointer to **string** | The final status of the response (&#x60;completed&#x60;, &#x60;cancelled&#x60;, &#x60;failed&#x60;, &#x60;incomplete&#x60;). | [optional] 
**StatusDetails** | Pointer to [**RealtimeResponseStatusDetails**](RealtimeResponseStatusDetails.md) |  | [optional] 
**Output** | Pointer to **[]map[string]interface{}** | The list of output items generated by the response. | [optional] 
**Usage** | Pointer to [**RealtimeResponseUsage**](RealtimeResponseUsage.md) |  | [optional] 

## Methods

### NewRealtimeResponse

`func NewRealtimeResponse() *RealtimeResponse`

NewRealtimeResponse instantiates a new RealtimeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeResponseWithDefaults

`func NewRealtimeResponseWithDefaults() *RealtimeResponse`

NewRealtimeResponseWithDefaults instantiates a new RealtimeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealtimeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealtimeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealtimeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RealtimeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *RealtimeResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RealtimeResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RealtimeResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RealtimeResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *RealtimeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RealtimeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RealtimeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RealtimeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *RealtimeResponse) GetStatusDetails() RealtimeResponseStatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *RealtimeResponse) GetStatusDetailsOk() (*RealtimeResponseStatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *RealtimeResponse) SetStatusDetails(v RealtimeResponseStatusDetails)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *RealtimeResponse) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetOutput

`func (o *RealtimeResponse) GetOutput() []map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *RealtimeResponse) GetOutputOk() (*[]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *RealtimeResponse) SetOutput(v []map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *RealtimeResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetUsage

`func (o *RealtimeResponse) GetUsage() RealtimeResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RealtimeResponse) GetUsageOk() (*RealtimeResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RealtimeResponse) SetUsage(v RealtimeResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *RealtimeResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

