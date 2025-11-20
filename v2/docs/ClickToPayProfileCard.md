# ClickToPayProfileCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of customer on Decta partner side (allows to bind multiple cards to one customer) | 
**Name** | Pointer to **string** | Name | [optional] 
**Surname** | Pointer to **string** | Surname | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Country** | Pointer to **string** | Country(ISO 3166-1 alpha-3) of address | [optional] 
**Cards** | Pointer to [**[]ClickToPayCard**](ClickToPayCard.md) |  | [optional] 

## Methods

### NewClickToPayProfileCard

`func NewClickToPayProfileCard(id string, ) *ClickToPayProfileCard`

NewClickToPayProfileCard instantiates a new ClickToPayProfileCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickToPayProfileCardWithDefaults

`func NewClickToPayProfileCardWithDefaults() *ClickToPayProfileCard`

NewClickToPayProfileCardWithDefaults instantiates a new ClickToPayProfileCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClickToPayProfileCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClickToPayProfileCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClickToPayProfileCard) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ClickToPayProfileCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClickToPayProfileCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClickToPayProfileCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClickToPayProfileCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *ClickToPayProfileCard) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ClickToPayProfileCard) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ClickToPayProfileCard) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *ClickToPayProfileCard) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetEmail

`func (o *ClickToPayProfileCard) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClickToPayProfileCard) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClickToPayProfileCard) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClickToPayProfileCard) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *ClickToPayProfileCard) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ClickToPayProfileCard) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ClickToPayProfileCard) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ClickToPayProfileCard) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetCountry

`func (o *ClickToPayProfileCard) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClickToPayProfileCard) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClickToPayProfileCard) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ClickToPayProfileCard) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCards

`func (o *ClickToPayProfileCard) GetCards() []ClickToPayCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *ClickToPayProfileCard) GetCardsOk() (*[]ClickToPayCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *ClickToPayProfileCard) SetCards(v []ClickToPayCard)`

SetCards sets Cards field to given value.

### HasCards

`func (o *ClickToPayProfileCard) HasCards() bool`

HasCards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


