# FinetuneCompletionRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **string** | The input prompt for this training example. | [optional] 
**Completion** | Pointer to **string** | The desired completion for this training example. | [optional] 

## Methods

### NewFinetuneCompletionRequestInput

`func NewFinetuneCompletionRequestInput() *FinetuneCompletionRequestInput`

NewFinetuneCompletionRequestInput instantiates a new FinetuneCompletionRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinetuneCompletionRequestInputWithDefaults

`func NewFinetuneCompletionRequestInputWithDefaults() *FinetuneCompletionRequestInput`

NewFinetuneCompletionRequestInputWithDefaults instantiates a new FinetuneCompletionRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *FinetuneCompletionRequestInput) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FinetuneCompletionRequestInput) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FinetuneCompletionRequestInput) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *FinetuneCompletionRequestInput) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetCompletion

`func (o *FinetuneCompletionRequestInput) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *FinetuneCompletionRequestInput) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *FinetuneCompletionRequestInput) SetCompletion(v string)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *FinetuneCompletionRequestInput) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


