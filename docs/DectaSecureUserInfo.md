# DectaSecureUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**DectaSecureUserDto**](DectaSecureUserDto.md) |  | [optional] 
**Devices** | Pointer to [**[]DectaSecureDeviceDto**](DectaSecureDeviceDto.md) | Devices | [optional] 

## Methods

### NewDectaSecureUserInfo

`func NewDectaSecureUserInfo() *DectaSecureUserInfo`

NewDectaSecureUserInfo instantiates a new DectaSecureUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDectaSecureUserInfoWithDefaults

`func NewDectaSecureUserInfoWithDefaults() *DectaSecureUserInfo`

NewDectaSecureUserInfoWithDefaults instantiates a new DectaSecureUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *DectaSecureUserInfo) GetUser() DectaSecureUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DectaSecureUserInfo) GetUserOk() (*DectaSecureUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DectaSecureUserInfo) SetUser(v DectaSecureUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *DectaSecureUserInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDevices

`func (o *DectaSecureUserInfo) GetDevices() []DectaSecureDeviceDto`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DectaSecureUserInfo) GetDevicesOk() (*[]DectaSecureDeviceDto, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DectaSecureUserInfo) SetDevices(v []DectaSecureDeviceDto)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DectaSecureUserInfo) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


