# DeliveryAddressGiftCardApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | **string** | Type of a card delivery | 
**Name** | **string** | Name of a cardholder on an envelope | 
**Surname** | **string** | Surname of a cardholder on an envelope | 
**Country** | **string** | Country(ISO 3166-1 alpha-3) of address | 
**City** | **string** | City of address | 
**Street** | **string** | Number of a building and Street of address | 
**ZipCode** | **string** | ZIP code. Max length - 15 symbols. | 

## Methods

### NewDeliveryAddressGiftCardApiDto

`func NewDeliveryAddressGiftCardApiDto(shipment string, name string, surname string, country string, city string, street string, zipCode string, ) *DeliveryAddressGiftCardApiDto`

NewDeliveryAddressGiftCardApiDto instantiates a new DeliveryAddressGiftCardApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryAddressGiftCardApiDtoWithDefaults

`func NewDeliveryAddressGiftCardApiDtoWithDefaults() *DeliveryAddressGiftCardApiDto`

NewDeliveryAddressGiftCardApiDtoWithDefaults instantiates a new DeliveryAddressGiftCardApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *DeliveryAddressGiftCardApiDto) GetShipment() string`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *DeliveryAddressGiftCardApiDto) GetShipmentOk() (*string, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *DeliveryAddressGiftCardApiDto) SetShipment(v string)`

SetShipment sets Shipment field to given value.


### GetName

`func (o *DeliveryAddressGiftCardApiDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryAddressGiftCardApiDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryAddressGiftCardApiDto) SetName(v string)`

SetName sets Name field to given value.


### GetSurname

`func (o *DeliveryAddressGiftCardApiDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *DeliveryAddressGiftCardApiDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *DeliveryAddressGiftCardApiDto) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetCountry

`func (o *DeliveryAddressGiftCardApiDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DeliveryAddressGiftCardApiDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DeliveryAddressGiftCardApiDto) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *DeliveryAddressGiftCardApiDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DeliveryAddressGiftCardApiDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DeliveryAddressGiftCardApiDto) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *DeliveryAddressGiftCardApiDto) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *DeliveryAddressGiftCardApiDto) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *DeliveryAddressGiftCardApiDto) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *DeliveryAddressGiftCardApiDto) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *DeliveryAddressGiftCardApiDto) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *DeliveryAddressGiftCardApiDto) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


