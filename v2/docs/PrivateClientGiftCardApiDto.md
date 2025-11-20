# PrivateClientGiftCardApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Surname** | **string** | Surname | 
**MaidenName** | **string** | Maiden Name | 
**CurrentAddress** | [**RegistrationAddressGiftCardApiDto**](RegistrationAddressGiftCardApiDto.md) |  | 
**Language** | **string** | Communication language (ISO 639-1) | 

## Methods

### NewPrivateClientGiftCardApiDto

`func NewPrivateClientGiftCardApiDto(name string, surname string, maidenName string, currentAddress RegistrationAddressGiftCardApiDto, language string, ) *PrivateClientGiftCardApiDto`

NewPrivateClientGiftCardApiDto instantiates a new PrivateClientGiftCardApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateClientGiftCardApiDtoWithDefaults

`func NewPrivateClientGiftCardApiDtoWithDefaults() *PrivateClientGiftCardApiDto`

NewPrivateClientGiftCardApiDtoWithDefaults instantiates a new PrivateClientGiftCardApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateClientGiftCardApiDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateClientGiftCardApiDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateClientGiftCardApiDto) SetName(v string)`

SetName sets Name field to given value.


### GetSurname

`func (o *PrivateClientGiftCardApiDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *PrivateClientGiftCardApiDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *PrivateClientGiftCardApiDto) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetMaidenName

`func (o *PrivateClientGiftCardApiDto) GetMaidenName() string`

GetMaidenName returns the MaidenName field if non-nil, zero value otherwise.

### GetMaidenNameOk

`func (o *PrivateClientGiftCardApiDto) GetMaidenNameOk() (*string, bool)`

GetMaidenNameOk returns a tuple with the MaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaidenName

`func (o *PrivateClientGiftCardApiDto) SetMaidenName(v string)`

SetMaidenName sets MaidenName field to given value.


### GetCurrentAddress

`func (o *PrivateClientGiftCardApiDto) GetCurrentAddress() RegistrationAddressGiftCardApiDto`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *PrivateClientGiftCardApiDto) GetCurrentAddressOk() (*RegistrationAddressGiftCardApiDto, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *PrivateClientGiftCardApiDto) SetCurrentAddress(v RegistrationAddressGiftCardApiDto)`

SetCurrentAddress sets CurrentAddress field to given value.


### GetLanguage

`func (o *PrivateClientGiftCardApiDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PrivateClientGiftCardApiDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PrivateClientGiftCardApiDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


