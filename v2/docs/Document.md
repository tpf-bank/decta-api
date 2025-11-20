# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of ID document of a person | 
**Subtype** | Pointer to **string** | Subtype of ID document of a person. | [optional] 
**Number** | **string** | Number of ID document | 
**CountryCode** | **string** | Country(ISO 3166-1 alpha-3) of ID document | 
**IssuingDate** | **string** | Date(YYYY-MM-DD) of issue of person ID document | 
**ExpiryDate** | **string** | Date(YYYY-MM-DD) of expiry of person ID document | 
**BirthDate** | **string** | Date(YYYY-MM-DD) of birth of a person | 

## Methods

### NewDocument

`func NewDocument(type_ string, number string, countryCode string, issuingDate string, expiryDate string, birthDate string, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Document) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document) SetType(v string)`

SetType sets Type field to given value.


### GetSubtype

`func (o *Document) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Document) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Document) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Document) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetNumber

`func (o *Document) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Document) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Document) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCountryCode

`func (o *Document) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Document) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Document) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetIssuingDate

`func (o *Document) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *Document) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *Document) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.


### GetExpiryDate

`func (o *Document) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Document) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Document) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetBirthDate

`func (o *Document) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *Document) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *Document) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


