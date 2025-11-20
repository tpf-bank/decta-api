# PrivateClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Person name | [optional] 
**Surname** | Pointer to **string** | Person surname | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**Language** | Pointer to **string** | Communication language(ISO 639-1) | [optional] 
**Id** | Pointer to **string** | ID of customer on Decta partner side (allows to bind multiple cards to one customer). Length 18 symbols. | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Comment** | Pointer to **string** | Additional information according to agreement with Decta | [optional] 
**CurrentAddress** | Pointer to [**RegistrationAddress**](RegistrationAddress.md) |  | [optional] 
**MaidenName** | Pointer to **string** | Private client maiden name | [optional] 
**AdditionalContact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**UserDefinedField1** | Pointer to **string** | Free field for custom use, e.g. Backup mobile phone number used for courier deliveries, when main phone is not accessible. | [optional] 

## Methods

### NewPrivateClientInfo

`func NewPrivateClientInfo() *PrivateClientInfo`

NewPrivateClientInfo instantiates a new PrivateClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateClientInfoWithDefaults

`func NewPrivateClientInfoWithDefaults() *PrivateClientInfo`

NewPrivateClientInfoWithDefaults instantiates a new PrivateClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateClientInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateClientInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateClientInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateClientInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *PrivateClientInfo) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *PrivateClientInfo) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *PrivateClientInfo) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *PrivateClientInfo) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *PrivateClientInfo) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *PrivateClientInfo) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *PrivateClientInfo) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *PrivateClientInfo) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *PrivateClientInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrivateClientInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrivateClientInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PrivateClientInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLanguage

`func (o *PrivateClientInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PrivateClientInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PrivateClientInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PrivateClientInfo) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetId

`func (o *PrivateClientInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateClientInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateClientInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateClientInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocument

`func (o *PrivateClientInfo) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *PrivateClientInfo) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *PrivateClientInfo) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *PrivateClientInfo) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetComment

`func (o *PrivateClientInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PrivateClientInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PrivateClientInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PrivateClientInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *PrivateClientInfo) GetCurrentAddress() RegistrationAddress`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *PrivateClientInfo) GetCurrentAddressOk() (*RegistrationAddress, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *PrivateClientInfo) SetCurrentAddress(v RegistrationAddress)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *PrivateClientInfo) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetMaidenName

`func (o *PrivateClientInfo) GetMaidenName() string`

GetMaidenName returns the MaidenName field if non-nil, zero value otherwise.

### GetMaidenNameOk

`func (o *PrivateClientInfo) GetMaidenNameOk() (*string, bool)`

GetMaidenNameOk returns a tuple with the MaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaidenName

`func (o *PrivateClientInfo) SetMaidenName(v string)`

SetMaidenName sets MaidenName field to given value.

### HasMaidenName

`func (o *PrivateClientInfo) HasMaidenName() bool`

HasMaidenName returns a boolean if a field has been set.

### GetAdditionalContact

`func (o *PrivateClientInfo) GetAdditionalContact() Contact`

GetAdditionalContact returns the AdditionalContact field if non-nil, zero value otherwise.

### GetAdditionalContactOk

`func (o *PrivateClientInfo) GetAdditionalContactOk() (*Contact, bool)`

GetAdditionalContactOk returns a tuple with the AdditionalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContact

`func (o *PrivateClientInfo) SetAdditionalContact(v Contact)`

SetAdditionalContact sets AdditionalContact field to given value.

### HasAdditionalContact

`func (o *PrivateClientInfo) HasAdditionalContact() bool`

HasAdditionalContact returns a boolean if a field has been set.

### GetUserDefinedField1

`func (o *PrivateClientInfo) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *PrivateClientInfo) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *PrivateClientInfo) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *PrivateClientInfo) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


