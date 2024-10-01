# ClientAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country(ISO 3166-1 alpha-3) of address | [optional] 
**City** | Pointer to **string** | City of address | [optional] 
**Street** | Pointer to **string** | Number of a building and Street of address | [optional] 
**ZipCode** | Pointer to **string** | ZIP code | [optional] 

## Methods

### NewClientAddress

`func NewClientAddress() *ClientAddress`

NewClientAddress instantiates a new ClientAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAddressWithDefaults

`func NewClientAddressWithDefaults() *ClientAddress`

NewClientAddressWithDefaults instantiates a new ClientAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *ClientAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClientAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClientAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ClientAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *ClientAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ClientAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ClientAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ClientAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreet

`func (o *ClientAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ClientAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ClientAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ClientAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetZipCode

`func (o *ClientAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ClientAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ClientAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ClientAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


