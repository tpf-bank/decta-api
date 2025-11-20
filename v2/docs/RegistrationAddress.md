# RegistrationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country(ISO 3166-1 alpha-3) of address | [optional] 
**City** | **string** | City of address | 
**Street** | **string** | Number of a building and Street of address | 
**ZipCode** | Pointer to **string** | ZIP code. Max length - 15 symbols. | [optional] 
**Months** | **int32** | Period of living at address in months | 

## Methods

### NewRegistrationAddress

`func NewRegistrationAddress(city string, street string, months int32, ) *RegistrationAddress`

NewRegistrationAddress instantiates a new RegistrationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationAddressWithDefaults

`func NewRegistrationAddressWithDefaults() *RegistrationAddress`

NewRegistrationAddressWithDefaults instantiates a new RegistrationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *RegistrationAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RegistrationAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RegistrationAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RegistrationAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *RegistrationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RegistrationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RegistrationAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *RegistrationAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *RegistrationAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *RegistrationAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *RegistrationAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *RegistrationAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *RegistrationAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *RegistrationAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetMonths

`func (o *RegistrationAddress) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *RegistrationAddress) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *RegistrationAddress) SetMonths(v int32)`

SetMonths sets Months field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


