# BatchCardDeliveryAddressApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | Country(ISO 3166-1 alpha-3) of address | 
**City** | **string** | City of address | 
**Street** | **string** | Number of a building and Street of address | 
**ZipCode** | Pointer to **string** | ZIP code. Max length - 15 symbols. | [optional] 
**Shipment** | Pointer to **string** | Shipment | [optional] 
**Name** | Pointer to **string** | Name of a cardholder on an envelope | [optional] 
**Surname** | Pointer to **string** | Surname of a cardholder on an envelope | [optional] 
**Language** | Pointer to **string** | Communication language (ISO 639-1) | [optional] 

## Methods

### NewBatchCardDeliveryAddressApiDto

`func NewBatchCardDeliveryAddressApiDto(country string, city string, street string, ) *BatchCardDeliveryAddressApiDto`

NewBatchCardDeliveryAddressApiDto instantiates a new BatchCardDeliveryAddressApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCardDeliveryAddressApiDtoWithDefaults

`func NewBatchCardDeliveryAddressApiDtoWithDefaults() *BatchCardDeliveryAddressApiDto`

NewBatchCardDeliveryAddressApiDtoWithDefaults instantiates a new BatchCardDeliveryAddressApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *BatchCardDeliveryAddressApiDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BatchCardDeliveryAddressApiDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BatchCardDeliveryAddressApiDto) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *BatchCardDeliveryAddressApiDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BatchCardDeliveryAddressApiDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BatchCardDeliveryAddressApiDto) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *BatchCardDeliveryAddressApiDto) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *BatchCardDeliveryAddressApiDto) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *BatchCardDeliveryAddressApiDto) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *BatchCardDeliveryAddressApiDto) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *BatchCardDeliveryAddressApiDto) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *BatchCardDeliveryAddressApiDto) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *BatchCardDeliveryAddressApiDto) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetShipment

`func (o *BatchCardDeliveryAddressApiDto) GetShipment() string`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *BatchCardDeliveryAddressApiDto) GetShipmentOk() (*string, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *BatchCardDeliveryAddressApiDto) SetShipment(v string)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *BatchCardDeliveryAddressApiDto) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetName

`func (o *BatchCardDeliveryAddressApiDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchCardDeliveryAddressApiDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchCardDeliveryAddressApiDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchCardDeliveryAddressApiDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *BatchCardDeliveryAddressApiDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *BatchCardDeliveryAddressApiDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *BatchCardDeliveryAddressApiDto) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *BatchCardDeliveryAddressApiDto) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetLanguage

`func (o *BatchCardDeliveryAddressApiDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BatchCardDeliveryAddressApiDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BatchCardDeliveryAddressApiDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BatchCardDeliveryAddressApiDto) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


