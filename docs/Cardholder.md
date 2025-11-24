# Cardholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Person name | [optional] 
**Surname** | Pointer to **string** | Person surname | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**Language** | Pointer to **string** | Communication language(ISO 639-1) | [optional] 
**Id** | Pointer to **string** | An Existing client id | [optional] 
**MaidenName** | Pointer to **string** | Holder maiden name | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**CurrentAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Role** | Pointer to **string** | Position of a cardholder in a company - for payroll and Business card issuing  | [optional] 
**UserDefinedField1** | Pointer to **string** | Free field for custom use, e.g. Backup mobile phone number used for courier deliveries, when main phone is not accessible. | [optional] 

## Methods

### NewCardholder

`func NewCardholder() *Cardholder`

NewCardholder instantiates a new Cardholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderWithDefaults

`func NewCardholderWithDefaults() *Cardholder`

NewCardholderWithDefaults instantiates a new Cardholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cardholder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cardholder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cardholder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cardholder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *Cardholder) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Cardholder) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Cardholder) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *Cardholder) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *Cardholder) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Cardholder) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *Cardholder) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *Cardholder) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *Cardholder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Cardholder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Cardholder) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Cardholder) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLanguage

`func (o *Cardholder) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Cardholder) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Cardholder) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Cardholder) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetId

`func (o *Cardholder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cardholder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cardholder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cardholder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaidenName

`func (o *Cardholder) GetMaidenName() string`

GetMaidenName returns the MaidenName field if non-nil, zero value otherwise.

### GetMaidenNameOk

`func (o *Cardholder) GetMaidenNameOk() (*string, bool)`

GetMaidenNameOk returns a tuple with the MaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaidenName

`func (o *Cardholder) SetMaidenName(v string)`

SetMaidenName sets MaidenName field to given value.

### HasMaidenName

`func (o *Cardholder) HasMaidenName() bool`

HasMaidenName returns a boolean if a field has been set.

### GetDocument

`func (o *Cardholder) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *Cardholder) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *Cardholder) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *Cardholder) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *Cardholder) GetCurrentAddress() Address`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *Cardholder) GetCurrentAddressOk() (*Address, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *Cardholder) SetCurrentAddress(v Address)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *Cardholder) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetRole

`func (o *Cardholder) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Cardholder) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Cardholder) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Cardholder) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserDefinedField1

`func (o *Cardholder) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *Cardholder) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *Cardholder) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *Cardholder) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


