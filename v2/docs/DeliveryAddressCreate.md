# DeliveryAddressCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to **string** | Type of a card delivery | [optional] 
**Country** | **string** | Country(ISO 3166-1 alpha-3) of address | 
**City** | **string** | City of address | 
**Street** | **string** | Number of a building and Street of address | 
**ZipCode** | **string** | ZIP code. Max length - 15 symbols. | 
**Language** | Pointer to **string** | Communication language (ISO 639-1) | [optional] 
**Name** | Pointer to **string** | Name of a cardholder on an envelope | [optional] 
**Surname** | Pointer to **string** | Surname of a cardholder on an envelope | [optional] 

## Methods

### NewDeliveryAddressCreate

`func NewDeliveryAddressCreate(country string, city string, street string, zipCode string, ) *DeliveryAddressCreate`

NewDeliveryAddressCreate instantiates a new DeliveryAddressCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryAddressCreateWithDefaults

`func NewDeliveryAddressCreateWithDefaults() *DeliveryAddressCreate`

NewDeliveryAddressCreateWithDefaults instantiates a new DeliveryAddressCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *DeliveryAddressCreate) GetShipment() string`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *DeliveryAddressCreate) GetShipmentOk() (*string, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *DeliveryAddressCreate) SetShipment(v string)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *DeliveryAddressCreate) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetCountry

`func (o *DeliveryAddressCreate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DeliveryAddressCreate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DeliveryAddressCreate) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *DeliveryAddressCreate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DeliveryAddressCreate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DeliveryAddressCreate) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *DeliveryAddressCreate) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *DeliveryAddressCreate) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *DeliveryAddressCreate) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *DeliveryAddressCreate) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *DeliveryAddressCreate) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *DeliveryAddressCreate) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetLanguage

`func (o *DeliveryAddressCreate) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DeliveryAddressCreate) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DeliveryAddressCreate) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *DeliveryAddressCreate) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *DeliveryAddressCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryAddressCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryAddressCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeliveryAddressCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *DeliveryAddressCreate) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *DeliveryAddressCreate) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *DeliveryAddressCreate) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *DeliveryAddressCreate) HasSurname() bool`

HasSurname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


