# RealtimeConversationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the item, this can be generated by the client to help manage server-side context, but is not required because the server will generate one if not provided. | [optional] 
**Type** | Pointer to **string** | The type of the item (&#x60;message&#x60;, &#x60;function_call&#x60;, &#x60;function_call_output&#x60;). | [optional] 
**Status** | Pointer to **string** | The status of the item (&#x60;completed&#x60;, &#x60;incomplete&#x60;). These have no effect on the conversation, but are accepted for consistency with the &#x60;conversation.item.created&#x60; event. | [optional] 
**Role** | Pointer to **string** | The role of the message sender (&#x60;user&#x60;, &#x60;assistant&#x60;, &#x60;system&#x60;), only applicable for &#x60;message&#x60; items. | [optional] 
**Content** | Pointer to [**[]RealtimeConversationItemContentInner**](RealtimeConversationItemContentInner.md) | The content of the message, applicable for &#x60;message&#x60; items. Message items with a role of &#x60;system&#x60; support only &#x60;input_text&#x60; content, message items of role &#x60;user&#x60; support &#x60;input_text&#x60; and &#x60;input_audio&#x60; content, and message items of role &#x60;assistant&#x60; support &#x60;text&#x60; content. | [optional] 
**CallId** | Pointer to **string** | The ID of the function call (for &#x60;function_call&#x60; and &#x60;function_call_output&#x60; items). If passed on a &#x60;function_call_output&#x60; item, the server will check that a &#x60;function_call&#x60; item with the same ID exists in the conversation history. | [optional] 
**Name** | Pointer to **string** | The name of the function being called (for &#x60;function_call&#x60; items). | [optional] 
**Arguments** | Pointer to **string** | The arguments of the function call (for &#x60;function_call&#x60; items). | [optional] 
**Output** | Pointer to **string** | The output of the function call (for &#x60;function_call_output&#x60; items). | [optional] 

## Methods

### NewRealtimeConversationItem

`func NewRealtimeConversationItem() *RealtimeConversationItem`

NewRealtimeConversationItem instantiates a new RealtimeConversationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeConversationItemWithDefaults

`func NewRealtimeConversationItemWithDefaults() *RealtimeConversationItem`

NewRealtimeConversationItemWithDefaults instantiates a new RealtimeConversationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealtimeConversationItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealtimeConversationItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealtimeConversationItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RealtimeConversationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealtimeConversationItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeConversationItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeConversationItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealtimeConversationItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *RealtimeConversationItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RealtimeConversationItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RealtimeConversationItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RealtimeConversationItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *RealtimeConversationItem) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RealtimeConversationItem) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RealtimeConversationItem) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RealtimeConversationItem) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *RealtimeConversationItem) GetContent() []RealtimeConversationItemContentInner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RealtimeConversationItem) GetContentOk() (*[]RealtimeConversationItemContentInner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RealtimeConversationItem) SetContent(v []RealtimeConversationItemContentInner)`

SetContent sets Content field to given value.

### HasContent

`func (o *RealtimeConversationItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCallId

`func (o *RealtimeConversationItem) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *RealtimeConversationItem) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *RealtimeConversationItem) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *RealtimeConversationItem) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetName

`func (o *RealtimeConversationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealtimeConversationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealtimeConversationItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealtimeConversationItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArguments

`func (o *RealtimeConversationItem) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *RealtimeConversationItem) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *RealtimeConversationItem) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *RealtimeConversationItem) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetOutput

`func (o *RealtimeConversationItem) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *RealtimeConversationItem) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *RealtimeConversationItem) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *RealtimeConversationItem) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


