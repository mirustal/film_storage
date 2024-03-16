# RegisterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Admin** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegisterPostRequest

`func NewRegisterPostRequest(username string, password string, ) *RegisterPostRequest`

NewRegisterPostRequest instantiates a new RegisterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPostRequestWithDefaults

`func NewRegisterPostRequestWithDefaults() *RegisterPostRequest`

NewRegisterPostRequestWithDefaults instantiates a new RegisterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RegisterPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *RegisterPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAdmin

`func (o *RegisterPostRequest) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *RegisterPostRequest) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *RegisterPostRequest) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *RegisterPostRequest) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


