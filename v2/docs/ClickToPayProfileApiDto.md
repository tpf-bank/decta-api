# ClickToPayProfileApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of customer on Decta partner side (allows to bind multiple cards to one customer) | 
**Name** | Pointer to **string** | Name | [optional] 
**Surname** | Pointer to **string** | Surname | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Country** | Pointer to **string** | Country(ISO 3166-1 alpha-3) of address | [optional] 
**Address** | Pointer to [**C2PAddressCreateCard**](C2PAddressCreateCard.md) |  | [optional] 

## Methods

### NewClickToPayProfileApiDto

`func NewClickToPayProfileApiDto(id string, ) *ClickToPayProfileApiDto`

NewClickToPayProfileApiDto instantiates a new ClickToPayProfileApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickToPayProfileApiDtoWithDefaults

`func NewClickToPayProfileApiDtoWithDefaults() *ClickToPayProfileApiDto`

NewClickToPayProfileApiDtoWithDefaults instantiates a new ClickToPayProfileApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClickToPayProfileApiDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClickToPayProfileApiDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClickToPayProfileApiDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ClickToPayProfileApiDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClickToPayProfileApiDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClickToPayProfileApiDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClickToPayProfileApiDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *ClickToPayProfileApiDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ClickToPayProfileApiDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ClickToPayProfileApiDto) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *ClickToPayProfileApiDto) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetEmail

`func (o *ClickToPayProfileApiDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClickToPayProfileApiDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClickToPayProfileApiDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClickToPayProfileApiDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *ClickToPayProfileApiDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ClickToPayProfileApiDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ClickToPayProfileApiDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ClickToPayProfileApiDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetCountry

`func (o *ClickToPayProfileApiDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClickToPayProfileApiDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClickToPayProfileApiDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ClickToPayProfileApiDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAddress

`func (o *ClickToPayProfileApiDto) GetAddress() C2PAddressCreateCard`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClickToPayProfileApiDto) GetAddressOk() (*C2PAddressCreateCard, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClickToPayProfileApiDto) SetAddress(v C2PAddressCreateCard)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClickToPayProfileApiDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


