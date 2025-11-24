# DectaSecureUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | User name | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**LastModified** | Pointer to **string** | Date(YYYY-MM-DDTHH:mm:ss) of last modified | [optional] 

## Methods

### NewDectaSecureUserDto

`func NewDectaSecureUserDto() *DectaSecureUserDto`

NewDectaSecureUserDto instantiates a new DectaSecureUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDectaSecureUserDtoWithDefaults

`func NewDectaSecureUserDtoWithDefaults() *DectaSecureUserDto`

NewDectaSecureUserDtoWithDefaults instantiates a new DectaSecureUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *DectaSecureUserDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DectaSecureUserDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DectaSecureUserDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DectaSecureUserDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetActive

`func (o *DectaSecureUserDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DectaSecureUserDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DectaSecureUserDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DectaSecureUserDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastModified

`func (o *DectaSecureUserDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DectaSecureUserDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DectaSecureUserDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DectaSecureUserDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


