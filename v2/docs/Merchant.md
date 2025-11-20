# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalId** | Pointer to **string** | Merchant specific terminal code | [optional] 
**Id** | Pointer to **string** | A code provided to merchants by a payment processor | [optional] 
**Name** | Pointer to **string** | Merchant’s name | [optional] 
**CategoryCode** | Pointer to **string** | A Merchant Category Code (MCC) - reflects the primary category in which the merchant does business | [optional] 
**Type** | Pointer to **string** | N – imprinter; P – POS including E-Commerce; A - ATM | [optional] 
**Country** | Pointer to **string** | Merchant’s location | [optional] 
**City** | Pointer to **string** | Merchant’s location | [optional] 

## Methods

### NewMerchant

`func NewMerchant() *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminalId

`func (o *Merchant) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *Merchant) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *Merchant) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *Merchant) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetId

`func (o *Merchant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Merchant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Merchant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Merchant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Merchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategoryCode

`func (o *Merchant) GetCategoryCode() string`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *Merchant) GetCategoryCodeOk() (*string, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *Merchant) SetCategoryCode(v string)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *Merchant) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### GetType

`func (o *Merchant) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Merchant) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Merchant) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Merchant) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *Merchant) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Merchant) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Merchant) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Merchant) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *Merchant) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Merchant) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Merchant) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Merchant) HasCity() bool`

HasCity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


