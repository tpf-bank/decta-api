# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Person name &lt;i style&#x3D;\&quot;color:grey\&quot;&gt;[name length + surname length - max 24 symbols]&lt;i/&gt; | [optional] 
**Surname** | Pointer to **string** | Person surname &lt;i style&#x3D;\&quot;color:grey\&quot;&gt;[name length + surname length - max 24 symbols]&lt;i/&gt; | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**Language** | Pointer to **string** | Communication language(ISO 639-1) | [optional] 
**PositionRelation** | Pointer to **string** | Contact relation for private client or position for business client | [optional] 
**MailingAddress** | Pointer to [**DeliveryAddress**](DeliveryAddress.md) |  | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Contact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Contact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *Contact) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Contact) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Contact) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *Contact) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *Contact) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Contact) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *Contact) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *Contact) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLanguage

`func (o *Contact) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Contact) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Contact) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Contact) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPositionRelation

`func (o *Contact) GetPositionRelation() string`

GetPositionRelation returns the PositionRelation field if non-nil, zero value otherwise.

### GetPositionRelationOk

`func (o *Contact) GetPositionRelationOk() (*string, bool)`

GetPositionRelationOk returns a tuple with the PositionRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionRelation

`func (o *Contact) SetPositionRelation(v string)`

SetPositionRelation sets PositionRelation field to given value.

### HasPositionRelation

`func (o *Contact) HasPositionRelation() bool`

HasPositionRelation returns a boolean if a field has been set.

### GetMailingAddress

`func (o *Contact) GetMailingAddress() DeliveryAddress`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *Contact) GetMailingAddressOk() (*DeliveryAddress, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddress

`func (o *Contact) SetMailingAddress(v DeliveryAddress)`

SetMailingAddress sets MailingAddress field to given value.

### HasMailingAddress

`func (o *Contact) HasMailingAddress() bool`

HasMailingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


