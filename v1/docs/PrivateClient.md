# PrivateClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Person name | [optional] 
**Surname** | Pointer to **string** | Person surname | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 
**Language** | Pointer to **string** | Communication language(ISO 639-1) | [optional] 
**Id** | **string** | ID of customer on Decta partner side (allows to bind multiple cards to one customer). Length 18 symbols. | 
**Document** | [**Document**](Document.md) |  | 
**Comment** | Pointer to **string** | Additional information according to agreement with Decta | [optional] 
**CurrentAddress** | Pointer to [**RegistrationAddress**](RegistrationAddress.md) |  | [optional] 
**MaidenName** | Pointer to **string** | Private client maiden name | [optional] 
**AdditionalContact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**UserDefinedField1** | Pointer to **string** | Free field for custom use, e.g. Backup mobile phone number used for courier deliveries, when main phone is not accessible. | [optional] 

## Methods

### NewPrivateClient

`func NewPrivateClient(id string, document Document, ) *PrivateClient`

NewPrivateClient instantiates a new PrivateClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateClientWithDefaults

`func NewPrivateClientWithDefaults() *PrivateClient`

NewPrivateClientWithDefaults instantiates a new PrivateClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *PrivateClient) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *PrivateClient) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *PrivateClient) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *PrivateClient) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *PrivateClient) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *PrivateClient) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *PrivateClient) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *PrivateClient) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *PrivateClient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrivateClient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrivateClient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PrivateClient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLanguage

`func (o *PrivateClient) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PrivateClient) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PrivateClient) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PrivateClient) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetId

`func (o *PrivateClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateClient) SetId(v string)`

SetId sets Id field to given value.


### GetDocument

`func (o *PrivateClient) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *PrivateClient) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *PrivateClient) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetComment

`func (o *PrivateClient) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PrivateClient) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PrivateClient) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PrivateClient) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *PrivateClient) GetCurrentAddress() RegistrationAddress`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *PrivateClient) GetCurrentAddressOk() (*RegistrationAddress, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *PrivateClient) SetCurrentAddress(v RegistrationAddress)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *PrivateClient) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetMaidenName

`func (o *PrivateClient) GetMaidenName() string`

GetMaidenName returns the MaidenName field if non-nil, zero value otherwise.

### GetMaidenNameOk

`func (o *PrivateClient) GetMaidenNameOk() (*string, bool)`

GetMaidenNameOk returns a tuple with the MaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaidenName

`func (o *PrivateClient) SetMaidenName(v string)`

SetMaidenName sets MaidenName field to given value.

### HasMaidenName

`func (o *PrivateClient) HasMaidenName() bool`

HasMaidenName returns a boolean if a field has been set.

### GetAdditionalContact

`func (o *PrivateClient) GetAdditionalContact() Contact`

GetAdditionalContact returns the AdditionalContact field if non-nil, zero value otherwise.

### GetAdditionalContactOk

`func (o *PrivateClient) GetAdditionalContactOk() (*Contact, bool)`

GetAdditionalContactOk returns a tuple with the AdditionalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContact

`func (o *PrivateClient) SetAdditionalContact(v Contact)`

SetAdditionalContact sets AdditionalContact field to given value.

### HasAdditionalContact

`func (o *PrivateClient) HasAdditionalContact() bool`

HasAdditionalContact returns a boolean if a field has been set.

### GetUserDefinedField1

`func (o *PrivateClient) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *PrivateClient) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *PrivateClient) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *PrivateClient) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


