# RealtimeServerEventErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of error (e.g., \&quot;invalid_request_error\&quot;, \&quot;server_error\&quot;). | [optional] 
**Code** | Pointer to **string** | Error code, if any. | [optional] 
**Message** | Pointer to **string** | A human-readable error message. | [optional] 
**Param** | Pointer to **string** | Parameter related to the error, if any. | [optional] 
**EventId** | Pointer to **string** | The event_id of the client event that caused the error, if applicable. | [optional] 

## Methods

### NewRealtimeServerEventErrorError

`func NewRealtimeServerEventErrorError() *RealtimeServerEventErrorError`

NewRealtimeServerEventErrorError instantiates a new RealtimeServerEventErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeServerEventErrorErrorWithDefaults

`func NewRealtimeServerEventErrorErrorWithDefaults() *RealtimeServerEventErrorError`

NewRealtimeServerEventErrorErrorWithDefaults instantiates a new RealtimeServerEventErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RealtimeServerEventErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeServerEventErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeServerEventErrorError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealtimeServerEventErrorError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *RealtimeServerEventErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RealtimeServerEventErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RealtimeServerEventErrorError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RealtimeServerEventErrorError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *RealtimeServerEventErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RealtimeServerEventErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RealtimeServerEventErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RealtimeServerEventErrorError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParam

`func (o *RealtimeServerEventErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *RealtimeServerEventErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *RealtimeServerEventErrorError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *RealtimeServerEventErrorError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetEventId

`func (o *RealtimeServerEventErrorError) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *RealtimeServerEventErrorError) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *RealtimeServerEventErrorError) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *RealtimeServerEventErrorError) HasEventId() bool`

HasEventId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


