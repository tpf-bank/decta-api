# CardPreferencesGiftCardApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | productCode | 
**Currencies** | **[]string** | currencies | 
**DeliveryAddress** | [**DeliveryAddressGiftCardApiDto**](DeliveryAddressGiftCardApiDto.md) |  | 

## Methods

### NewCardPreferencesGiftCardApiDto

`func NewCardPreferencesGiftCardApiDto(productCode string, currencies []string, deliveryAddress DeliveryAddressGiftCardApiDto, ) *CardPreferencesGiftCardApiDto`

NewCardPreferencesGiftCardApiDto instantiates a new CardPreferencesGiftCardApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPreferencesGiftCardApiDtoWithDefaults

`func NewCardPreferencesGiftCardApiDtoWithDefaults() *CardPreferencesGiftCardApiDto`

NewCardPreferencesGiftCardApiDtoWithDefaults instantiates a new CardPreferencesGiftCardApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *CardPreferencesGiftCardApiDto) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *CardPreferencesGiftCardApiDto) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *CardPreferencesGiftCardApiDto) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetCurrencies

`func (o *CardPreferencesGiftCardApiDto) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *CardPreferencesGiftCardApiDto) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *CardPreferencesGiftCardApiDto) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.


### GetDeliveryAddress

`func (o *CardPreferencesGiftCardApiDto) GetDeliveryAddress() DeliveryAddressGiftCardApiDto`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CardPreferencesGiftCardApiDto) GetDeliveryAddressOk() (*DeliveryAddressGiftCardApiDto, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CardPreferencesGiftCardApiDto) SetDeliveryAddress(v DeliveryAddressGiftCardApiDto)`

SetDeliveryAddress sets DeliveryAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


