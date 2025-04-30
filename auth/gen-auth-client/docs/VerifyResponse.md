# VerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | Результат проверки токена | [optional] 
**Username** | Pointer to **string** | Имя пользователя | [optional] 

## Methods

### NewVerifyResponse

`func NewVerifyResponse() *VerifyResponse`

NewVerifyResponse instantiates a new VerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyResponseWithDefaults

`func NewVerifyResponseWithDefaults() *VerifyResponse`

NewVerifyResponseWithDefaults instantiates a new VerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *VerifyResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *VerifyResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *VerifyResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *VerifyResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetUsername

`func (o *VerifyResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VerifyResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VerifyResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VerifyResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


