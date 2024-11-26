# ProjectUserCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The ID of the user. | 
**Role** | **string** | &#x60;owner&#x60; or &#x60;member&#x60; | 

## Methods

### NewProjectUserCreateRequest

`func NewProjectUserCreateRequest(userId string, role string, ) *ProjectUserCreateRequest`

NewProjectUserCreateRequest instantiates a new ProjectUserCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUserCreateRequestWithDefaults

`func NewProjectUserCreateRequestWithDefaults() *ProjectUserCreateRequest`

NewProjectUserCreateRequestWithDefaults instantiates a new ProjectUserCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ProjectUserCreateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProjectUserCreateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProjectUserCreateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRole

`func (o *ProjectUserCreateRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectUserCreateRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectUserCreateRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


