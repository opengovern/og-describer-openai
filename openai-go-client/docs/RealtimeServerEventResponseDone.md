# RealtimeServerEventResponseDone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The unique ID of the server event. | 
**Type** | **string** | The event type, must be \&quot;response.done\&quot;. | 
**Response** | [**RealtimeResponse**](RealtimeResponse.md) |  | 

## Methods

### NewRealtimeServerEventResponseDone

`func NewRealtimeServerEventResponseDone(eventId string, type_ string, response RealtimeResponse, ) *RealtimeServerEventResponseDone`

NewRealtimeServerEventResponseDone instantiates a new RealtimeServerEventResponseDone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeServerEventResponseDoneWithDefaults

`func NewRealtimeServerEventResponseDoneWithDefaults() *RealtimeServerEventResponseDone`

NewRealtimeServerEventResponseDoneWithDefaults instantiates a new RealtimeServerEventResponseDone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *RealtimeServerEventResponseDone) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *RealtimeServerEventResponseDone) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *RealtimeServerEventResponseDone) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *RealtimeServerEventResponseDone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeServerEventResponseDone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeServerEventResponseDone) SetType(v string)`

SetType sets Type field to given value.


### GetResponse

`func (o *RealtimeServerEventResponseDone) GetResponse() RealtimeResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RealtimeServerEventResponseDone) GetResponseOk() (*RealtimeResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RealtimeServerEventResponseDone) SetResponse(v RealtimeResponse)`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

