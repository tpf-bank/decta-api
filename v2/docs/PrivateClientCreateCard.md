# PrivateClientCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of customer on Decta partner side (allows to bind multiple cards to one customer). | [optional] 
**Name** | Pointer to **string** | Person name | [optional] 
**Surname** | Pointer to **string** | Person surname | [optional] 
**MaidenName** | Pointer to **string** | Private client maiden name | [optional] 
**Document** | Pointer to [**DocumentCreateCard**](DocumentCreateCard.md) |  | [optional] 
**CurrentAddress** | Pointer to [**RegistrationAddressCreateApiDto**](RegistrationAddressCreateApiDto.md) |  | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**AdditionalContact** | Pointer to [**ContactCreateCard**](ContactCreateCard.md) |  | [optional] 
**UserDefinedField1** | Pointer to **string** | Free field for custom use, e.g. Backup mobile phone number used for courier deliveries, when main phone is not accessible. | [optional] 
**Comment** | Pointer to **string** | Additional information according to agreement with Decta | [optional] 

## Methods

### NewPrivateClientCreateCard

`func NewPrivateClientCreateCard() *PrivateClientCreateCard`

NewPrivateClientCreateCard instantiates a new PrivateClientCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateClientCreateCardWithDefaults

`func NewPrivateClientCreateCardWithDefaults() *PrivateClientCreateCard`

NewPrivateClientCreateCardWithDefaults instantiates a new PrivateClientCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateClientCreateCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateClientCreateCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateClientCreateCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateClientCreateCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrivateClientCreateCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateClientCreateCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateClientCreateCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateClientCreateCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *PrivateClientCreateCard) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *PrivateClientCreateCard) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *PrivateClientCreateCard) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *PrivateClientCreateCard) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMaidenName

`func (o *PrivateClientCreateCard) GetMaidenName() string`

GetMaidenName returns the MaidenName field if non-nil, zero value otherwise.

### GetMaidenNameOk

`func (o *PrivateClientCreateCard) GetMaidenNameOk() (*string, bool)`

GetMaidenNameOk returns a tuple with the MaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaidenName

`func (o *PrivateClientCreateCard) SetMaidenName(v string)`

SetMaidenName sets MaidenName field to given value.

### HasMaidenName

`func (o *PrivateClientCreateCard) HasMaidenName() bool`

HasMaidenName returns a boolean if a field has been set.

### GetDocument

`func (o *PrivateClientCreateCard) GetDocument() DocumentCreateCard`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *PrivateClientCreateCard) GetDocumentOk() (*DocumentCreateCard, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *PrivateClientCreateCard) SetDocument(v DocumentCreateCard)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *PrivateClientCreateCard) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *PrivateClientCreateCard) GetCurrentAddress() RegistrationAddressCreateApiDto`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *PrivateClientCreateCard) GetCurrentAddressOk() (*RegistrationAddressCreateApiDto, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *PrivateClientCreateCard) SetCurrentAddress(v RegistrationAddressCreateApiDto)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *PrivateClientCreateCard) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetMobilePhone

`func (o *PrivateClientCreateCard) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *PrivateClientCreateCard) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *PrivateClientCreateCard) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *PrivateClientCreateCard) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *PrivateClientCreateCard) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrivateClientCreateCard) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrivateClientCreateCard) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PrivateClientCreateCard) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdditionalContact

`func (o *PrivateClientCreateCard) GetAdditionalContact() ContactCreateCard`

GetAdditionalContact returns the AdditionalContact field if non-nil, zero value otherwise.

### GetAdditionalContactOk

`func (o *PrivateClientCreateCard) GetAdditionalContactOk() (*ContactCreateCard, bool)`

GetAdditionalContactOk returns a tuple with the AdditionalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContact

`func (o *PrivateClientCreateCard) SetAdditionalContact(v ContactCreateCard)`

SetAdditionalContact sets AdditionalContact field to given value.

### HasAdditionalContact

`func (o *PrivateClientCreateCard) HasAdditionalContact() bool`

HasAdditionalContact returns a boolean if a field has been set.

### GetUserDefinedField1

`func (o *PrivateClientCreateCard) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *PrivateClientCreateCard) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *PrivateClientCreateCard) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *PrivateClientCreateCard) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.

### GetComment

`func (o *PrivateClientCreateCard) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PrivateClientCreateCard) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PrivateClientCreateCard) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PrivateClientCreateCard) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


