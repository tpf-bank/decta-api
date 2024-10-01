# RegistrationAddressGiftCardApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Months** | **string** | Period of living at address in months | 
**Country** | **string** | Country(ISO 3166-1 alpha-3) of address | 
**City** | **string** | City of address | 
**Street** | **string** | Number of a building and Street of address | 
**ZipCode** | **string** | ZIP code. Max length - 15 symbols. | 

## Methods

### NewRegistrationAddressGiftCardApiDto

`func NewRegistrationAddressGiftCardApiDto(months string, country string, city string, street string, zipCode string, ) *RegistrationAddressGiftCardApiDto`

NewRegistrationAddressGiftCardApiDto instantiates a new RegistrationAddressGiftCardApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationAddressGiftCardApiDtoWithDefaults

`func NewRegistrationAddressGiftCardApiDtoWithDefaults() *RegistrationAddressGiftCardApiDto`

NewRegistrationAddressGiftCardApiDtoWithDefaults instantiates a new RegistrationAddressGiftCardApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonths

`func (o *RegistrationAddressGiftCardApiDto) GetMonths() string`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *RegistrationAddressGiftCardApiDto) GetMonthsOk() (*string, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *RegistrationAddressGiftCardApiDto) SetMonths(v string)`

SetMonths sets Months field to given value.


### GetCountry

`func (o *RegistrationAddressGiftCardApiDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RegistrationAddressGiftCardApiDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RegistrationAddressGiftCardApiDto) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *RegistrationAddressGiftCardApiDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RegistrationAddressGiftCardApiDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RegistrationAddressGiftCardApiDto) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *RegistrationAddressGiftCardApiDto) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *RegistrationAddressGiftCardApiDto) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *RegistrationAddressGiftCardApiDto) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *RegistrationAddressGiftCardApiDto) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *RegistrationAddressGiftCardApiDto) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *RegistrationAddressGiftCardApiDto) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


