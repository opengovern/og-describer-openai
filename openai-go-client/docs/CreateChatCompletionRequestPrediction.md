# CreateChatCompletionRequestPrediction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the predicted content you want to provide. This type is currently always &#x60;content&#x60;.  | 
**Content** | [**PredictionContentContent**](PredictionContentContent.md) |  | 

## Methods

### NewCreateChatCompletionRequestPrediction

`func NewCreateChatCompletionRequestPrediction(type_ string, content PredictionContentContent, ) *CreateChatCompletionRequestPrediction`

NewCreateChatCompletionRequestPrediction instantiates a new CreateChatCompletionRequestPrediction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChatCompletionRequestPredictionWithDefaults

`func NewCreateChatCompletionRequestPredictionWithDefaults() *CreateChatCompletionRequestPrediction`

NewCreateChatCompletionRequestPredictionWithDefaults instantiates a new CreateChatCompletionRequestPrediction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateChatCompletionRequestPrediction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateChatCompletionRequestPrediction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateChatCompletionRequestPrediction) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *CreateChatCompletionRequestPrediction) GetContent() PredictionContentContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateChatCompletionRequestPrediction) GetContentOk() (*PredictionContentContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateChatCompletionRequestPrediction) SetContent(v PredictionContentContent)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


