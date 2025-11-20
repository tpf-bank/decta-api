# BusinessClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | Business client company name | [optional] 
**RegistrationNumber** | **string** | Business client company registration number | 
**NameOnCard** | Pointer to **string** | Company name on card | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Beneficiaries** | Pointer to [**[]Beneficiary**](Beneficiary.md) | List of company beneficiaries - individuals who directly or indirectly controls at least 25% of the shares or voting rights of legal entity. | [optional] 

## Methods

### NewBusinessClient

`func NewBusinessClient(registrationNumber string, ) *BusinessClient`

NewBusinessClient instantiates a new BusinessClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessClientWithDefaults

`func NewBusinessClientWithDefaults() *BusinessClient`

NewBusinessClientWithDefaults instantiates a new BusinessClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *BusinessClient) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BusinessClient) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BusinessClient) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BusinessClient) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *BusinessClient) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *BusinessClient) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *BusinessClient) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.


### GetNameOnCard

`func (o *BusinessClient) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *BusinessClient) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *BusinessClient) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *BusinessClient) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetAddress

`func (o *BusinessClient) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessClient) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessClient) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BusinessClient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *BusinessClient) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BusinessClient) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BusinessClient) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *BusinessClient) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetBeneficiaries

`func (o *BusinessClient) GetBeneficiaries() []Beneficiary`

GetBeneficiaries returns the Beneficiaries field if non-nil, zero value otherwise.

### GetBeneficiariesOk

`func (o *BusinessClient) GetBeneficiariesOk() (*[]Beneficiary, bool)`

GetBeneficiariesOk returns a tuple with the Beneficiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaries

`func (o *BusinessClient) SetBeneficiaries(v []Beneficiary)`

SetBeneficiaries sets Beneficiaries field to given value.

### HasBeneficiaries

`func (o *BusinessClient) HasBeneficiaries() bool`

HasBeneficiaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


