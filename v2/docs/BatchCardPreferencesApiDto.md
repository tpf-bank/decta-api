# BatchCardPreferencesApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | Product code for virtual or plastic card product provided by Decta | 
**Currencies** | **string** | Card account currencies list ordered by priority decreasing | 
**DeliveryAddress** | [**BatchCardDeliveryAddressApiDto**](BatchCardDeliveryAddressApiDto.md) |  | 
**UserDefinedField12** | **string** | Free field for custom use | 

## Methods

### NewBatchCardPreferencesApiDto

`func NewBatchCardPreferencesApiDto(productCode string, currencies string, deliveryAddress BatchCardDeliveryAddressApiDto, userDefinedField12 string, ) *BatchCardPreferencesApiDto`

NewBatchCardPreferencesApiDto instantiates a new BatchCardPreferencesApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCardPreferencesApiDtoWithDefaults

`func NewBatchCardPreferencesApiDtoWithDefaults() *BatchCardPreferencesApiDto`

NewBatchCardPreferencesApiDtoWithDefaults instantiates a new BatchCardPreferencesApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *BatchCardPreferencesApiDto) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *BatchCardPreferencesApiDto) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *BatchCardPreferencesApiDto) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetCurrencies

`func (o *BatchCardPreferencesApiDto) GetCurrencies() string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *BatchCardPreferencesApiDto) GetCurrenciesOk() (*string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *BatchCardPreferencesApiDto) SetCurrencies(v string)`

SetCurrencies sets Currencies field to given value.


### GetDeliveryAddress

`func (o *BatchCardPreferencesApiDto) GetDeliveryAddress() BatchCardDeliveryAddressApiDto`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *BatchCardPreferencesApiDto) GetDeliveryAddressOk() (*BatchCardDeliveryAddressApiDto, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *BatchCardPreferencesApiDto) SetDeliveryAddress(v BatchCardDeliveryAddressApiDto)`

SetDeliveryAddress sets DeliveryAddress field to given value.


### GetUserDefinedField12

`func (o *BatchCardPreferencesApiDto) GetUserDefinedField12() string`

GetUserDefinedField12 returns the UserDefinedField12 field if non-nil, zero value otherwise.

### GetUserDefinedField12Ok

`func (o *BatchCardPreferencesApiDto) GetUserDefinedField12Ok() (*string, bool)`

GetUserDefinedField12Ok returns a tuple with the UserDefinedField12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField12

`func (o *BatchCardPreferencesApiDto) SetUserDefinedField12(v string)`

SetUserDefinedField12 sets UserDefinedField12 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


