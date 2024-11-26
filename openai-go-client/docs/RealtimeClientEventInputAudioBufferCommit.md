# RealtimeClientEventInputAudioBufferCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Optional client-generated ID used to identify this event. | [optional] 
**Type** | **string** | The event type, must be \&quot;input_audio_buffer.commit\&quot;. | 

## Methods

### NewRealtimeClientEventInputAudioBufferCommit

`func NewRealtimeClientEventInputAudioBufferCommit(type_ string, ) *RealtimeClientEventInputAudioBufferCommit`

NewRealtimeClientEventInputAudioBufferCommit instantiates a new RealtimeClientEventInputAudioBufferCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeClientEventInputAudioBufferCommitWithDefaults

`func NewRealtimeClientEventInputAudioBufferCommitWithDefaults() *RealtimeClientEventInputAudioBufferCommit`

NewRealtimeClientEventInputAudioBufferCommitWithDefaults instantiates a new RealtimeClientEventInputAudioBufferCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *RealtimeClientEventInputAudioBufferCommit) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *RealtimeClientEventInputAudioBufferCommit) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *RealtimeClientEventInputAudioBufferCommit) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *RealtimeClientEventInputAudioBufferCommit) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetType

`func (o *RealtimeClientEventInputAudioBufferCommit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeClientEventInputAudioBufferCommit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeClientEventInputAudioBufferCommit) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


