# C2PAddressCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country(ISO 3166-1 alpha-3) of address | [optional] 
**City** | Pointer to **string** | City of address | [optional] 
**Street** | Pointer to **string** | Number of a building and Street of address | [optional] 
**ZipCode** | Pointer to **string** | ZIP code | [optional] 

## Methods

### NewC2PAddressCreateCard

`func NewC2PAddressCreateCard() *C2PAddressCreateCard`

NewC2PAddressCreateCard instantiates a new C2PAddressCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC2PAddressCreateCardWithDefaults

`func NewC2PAddressCreateCardWithDefaults() *C2PAddressCreateCard`

NewC2PAddressCreateCardWithDefaults instantiates a new C2PAddressCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *C2PAddressCreateCard) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *C2PAddressCreateCard) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *C2PAddressCreateCard) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *C2PAddressCreateCard) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *C2PAddressCreateCard) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *C2PAddressCreateCard) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *C2PAddressCreateCard) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *C2PAddressCreateCard) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreet

`func (o *C2PAddressCreateCard) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *C2PAddressCreateCard) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *C2PAddressCreateCard) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *C2PAddressCreateCard) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetZipCode

`func (o *C2PAddressCreateCard) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *C2PAddressCreateCard) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *C2PAddressCreateCard) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *C2PAddressCreateCard) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


