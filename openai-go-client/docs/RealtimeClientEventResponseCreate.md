# RealtimeClientEventResponseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Optional client-generated ID used to identify this event. | [optional] 
**Type** | **string** | The event type, must be &#x60;response.create&#x60;. | 
**Response** | [**RealtimeResponse**](RealtimeResponse.md) |  | 

## Methods

### NewRealtimeClientEventResponseCreate

`func NewRealtimeClientEventResponseCreate(type_ string, response RealtimeResponse, ) *RealtimeClientEventResponseCreate`

NewRealtimeClientEventResponseCreate instantiates a new RealtimeClientEventResponseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeClientEventResponseCreateWithDefaults

`func NewRealtimeClientEventResponseCreateWithDefaults() *RealtimeClientEventResponseCreate`

NewRealtimeClientEventResponseCreateWithDefaults instantiates a new RealtimeClientEventResponseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *RealtimeClientEventResponseCreate) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *RealtimeClientEventResponseCreate) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *RealtimeClientEventResponseCreate) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *RealtimeClientEventResponseCreate) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetType

`func (o *RealtimeClientEventResponseCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeClientEventResponseCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeClientEventResponseCreate) SetType(v string)`

SetType sets Type field to given value.


### GetResponse

`func (o *RealtimeClientEventResponseCreate) GetResponse() RealtimeResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RealtimeClientEventResponseCreate) GetResponseOk() (*RealtimeResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RealtimeClientEventResponseCreate) SetResponse(v RealtimeResponse)`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

