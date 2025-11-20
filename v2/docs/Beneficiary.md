# Beneficiary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | First name of the beneficiary | 
**Surname** | **string** | Last name of the beneficiary | 
**BirthDate** | **string** | Date (YYYY-MM-DD) of birth of the beneficiary | 
**CountryCode** | **string** | Country (ISO 3166-1 alpha-3) of beneficiary&#39;s citizenship | 
**PercentageOfControl** | **string** | % of the shares or voting rights of the legal entity owned by the beneficiary | 
**RelationType** | **string** | Type of ownership | 

## Methods

### NewBeneficiary

`func NewBeneficiary(name string, surname string, birthDate string, countryCode string, percentageOfControl string, relationType string, ) *Beneficiary`

NewBeneficiary instantiates a new Beneficiary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficiaryWithDefaults

`func NewBeneficiaryWithDefaults() *Beneficiary`

NewBeneficiaryWithDefaults instantiates a new Beneficiary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Beneficiary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Beneficiary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Beneficiary) SetName(v string)`

SetName sets Name field to given value.


### GetSurname

`func (o *Beneficiary) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Beneficiary) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Beneficiary) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetBirthDate

`func (o *Beneficiary) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *Beneficiary) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *Beneficiary) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.


### GetCountryCode

`func (o *Beneficiary) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Beneficiary) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Beneficiary) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPercentageOfControl

`func (o *Beneficiary) GetPercentageOfControl() string`

GetPercentageOfControl returns the PercentageOfControl field if non-nil, zero value otherwise.

### GetPercentageOfControlOk

`func (o *Beneficiary) GetPercentageOfControlOk() (*string, bool)`

GetPercentageOfControlOk returns a tuple with the PercentageOfControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfControl

`func (o *Beneficiary) SetPercentageOfControl(v string)`

SetPercentageOfControl sets PercentageOfControl field to given value.


### GetRelationType

`func (o *Beneficiary) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *Beneficiary) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *Beneficiary) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


