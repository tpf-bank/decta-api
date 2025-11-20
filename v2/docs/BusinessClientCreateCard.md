# BusinessClientCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | Business client company name | [optional] 
**RegistrationNumber** | **string** | Business client company registration number | 
**NameOnCard** | Pointer to **string** | Company name on card | [optional] 
**Address** | Pointer to [**AddressCreateCard**](AddressCreateCard.md) |  | [optional] 
**Contact** | Pointer to [**ContactCreateCard**](ContactCreateCard.md) |  | [optional] 

## Methods

### NewBusinessClientCreateCard

`func NewBusinessClientCreateCard(registrationNumber string, ) *BusinessClientCreateCard`

NewBusinessClientCreateCard instantiates a new BusinessClientCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessClientCreateCardWithDefaults

`func NewBusinessClientCreateCardWithDefaults() *BusinessClientCreateCard`

NewBusinessClientCreateCardWithDefaults instantiates a new BusinessClientCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *BusinessClientCreateCard) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BusinessClientCreateCard) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BusinessClientCreateCard) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BusinessClientCreateCard) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *BusinessClientCreateCard) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *BusinessClientCreateCard) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *BusinessClientCreateCard) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.


### GetNameOnCard

`func (o *BusinessClientCreateCard) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *BusinessClientCreateCard) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *BusinessClientCreateCard) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *BusinessClientCreateCard) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetAddress

`func (o *BusinessClientCreateCard) GetAddress() AddressCreateCard`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessClientCreateCard) GetAddressOk() (*AddressCreateCard, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessClientCreateCard) SetAddress(v AddressCreateCard)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BusinessClientCreateCard) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *BusinessClientCreateCard) GetContact() ContactCreateCard`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BusinessClientCreateCard) GetContactOk() (*ContactCreateCard, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BusinessClientCreateCard) SetContact(v ContactCreateCard)`

SetContact sets Contact field to given value.

### HasContact

`func (o *BusinessClientCreateCard) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


