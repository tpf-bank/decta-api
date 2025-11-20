# BusinessClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | Business client company name | [optional] 
**RegistrationNumber** | **string** | Business client company registration number | 
**NameOnCard** | Pointer to **string** | Company name on card | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Beneficiaries** | Pointer to [**[]Beneficiary**](Beneficiary.md) | List of company beneficiaries - individuals who directly or indirectly controls at least 25% of the shares or voting rights of legal entity. | [optional] 
**CardHolder** | Pointer to [**Cardholder**](Cardholder.md) |  | [optional] 

## Methods

### NewBusinessClientInfo

`func NewBusinessClientInfo(registrationNumber string, ) *BusinessClientInfo`

NewBusinessClientInfo instantiates a new BusinessClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessClientInfoWithDefaults

`func NewBusinessClientInfoWithDefaults() *BusinessClientInfo`

NewBusinessClientInfoWithDefaults instantiates a new BusinessClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *BusinessClientInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BusinessClientInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BusinessClientInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BusinessClientInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *BusinessClientInfo) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *BusinessClientInfo) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *BusinessClientInfo) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.


### GetNameOnCard

`func (o *BusinessClientInfo) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *BusinessClientInfo) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *BusinessClientInfo) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *BusinessClientInfo) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetAddress

`func (o *BusinessClientInfo) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessClientInfo) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessClientInfo) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BusinessClientInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *BusinessClientInfo) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BusinessClientInfo) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BusinessClientInfo) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *BusinessClientInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetBeneficiaries

`func (o *BusinessClientInfo) GetBeneficiaries() []Beneficiary`

GetBeneficiaries returns the Beneficiaries field if non-nil, zero value otherwise.

### GetBeneficiariesOk

`func (o *BusinessClientInfo) GetBeneficiariesOk() (*[]Beneficiary, bool)`

GetBeneficiariesOk returns a tuple with the Beneficiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaries

`func (o *BusinessClientInfo) SetBeneficiaries(v []Beneficiary)`

SetBeneficiaries sets Beneficiaries field to given value.

### HasBeneficiaries

`func (o *BusinessClientInfo) HasBeneficiaries() bool`

HasBeneficiaries returns a boolean if a field has been set.

### GetCardHolder

`func (o *BusinessClientInfo) GetCardHolder() Cardholder`

GetCardHolder returns the CardHolder field if non-nil, zero value otherwise.

### GetCardHolderOk

`func (o *BusinessClientInfo) GetCardHolderOk() (*Cardholder, bool)`

GetCardHolderOk returns a tuple with the CardHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolder

`func (o *BusinessClientInfo) SetCardHolder(v Cardholder)`

SetCardHolder sets CardHolder field to given value.

### HasCardHolder

`func (o *BusinessClientInfo) HasCardHolder() bool`

HasCardHolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


