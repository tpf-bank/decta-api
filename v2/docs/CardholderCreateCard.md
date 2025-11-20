# CardholderCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An Existing client id | [optional] 
**Name** | Pointer to **string** | Person name | [optional] 
**Surname** | Pointer to **string** | Person surname | [optional] 
**MaidenName** | Pointer to **string** | Holder maiden name | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**Document** | Pointer to [**DocumentCreateCard**](DocumentCreateCard.md) |  | [optional] 
**CurrentAddress** | Pointer to [**AddressCreateCard**](AddressCreateCard.md) |  | [optional] 
**UserDefinedField1** | Pointer to **string** | Free field for custom use, e.g. Backup mobile phone number used for courier deliveries, when main phone is not accessible. | [optional] 
**NameOnCard** | Pointer to **string** | Card holder name on card | [optional] 

## Methods

### NewCardholderCreateCard

`func NewCardholderCreateCard() *CardholderCreateCard`

NewCardholderCreateCard instantiates a new CardholderCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderCreateCardWithDefaults

`func NewCardholderCreateCardWithDefaults() *CardholderCreateCard`

NewCardholderCreateCardWithDefaults instantiates a new CardholderCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardholderCreateCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardholderCreateCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardholderCreateCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardholderCreateCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CardholderCreateCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardholderCreateCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardholderCreateCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CardholderCreateCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *CardholderCreateCard) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CardholderCreateCard) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CardholderCreateCard) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *CardholderCreateCard) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMaidenName

`func (o *CardholderCreateCard) GetMaidenName() string`

GetMaidenName returns the MaidenName field if non-nil, zero value otherwise.

### GetMaidenNameOk

`func (o *CardholderCreateCard) GetMaidenNameOk() (*string, bool)`

GetMaidenNameOk returns a tuple with the MaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaidenName

`func (o *CardholderCreateCard) SetMaidenName(v string)`

SetMaidenName sets MaidenName field to given value.

### HasMaidenName

`func (o *CardholderCreateCard) HasMaidenName() bool`

HasMaidenName returns a boolean if a field has been set.

### GetMobilePhone

`func (o *CardholderCreateCard) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *CardholderCreateCard) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *CardholderCreateCard) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *CardholderCreateCard) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *CardholderCreateCard) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CardholderCreateCard) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CardholderCreateCard) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CardholderCreateCard) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDocument

`func (o *CardholderCreateCard) GetDocument() DocumentCreateCard`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CardholderCreateCard) GetDocumentOk() (*DocumentCreateCard, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CardholderCreateCard) SetDocument(v DocumentCreateCard)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *CardholderCreateCard) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *CardholderCreateCard) GetCurrentAddress() AddressCreateCard`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *CardholderCreateCard) GetCurrentAddressOk() (*AddressCreateCard, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *CardholderCreateCard) SetCurrentAddress(v AddressCreateCard)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *CardholderCreateCard) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetUserDefinedField1

`func (o *CardholderCreateCard) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *CardholderCreateCard) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *CardholderCreateCard) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *CardholderCreateCard) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.

### GetNameOnCard

`func (o *CardholderCreateCard) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *CardholderCreateCard) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *CardholderCreateCard) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *CardholderCreateCard) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


