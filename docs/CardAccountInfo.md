# CardAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Card account identification number | [optional] 
**Currencies** | Pointer to [**[]CardAccountCurrencyInfo**](CardAccountCurrencyInfo.md) | Card account currencies | [optional] 

## Methods

### NewCardAccountInfo

`func NewCardAccountInfo() *CardAccountInfo`

NewCardAccountInfo instantiates a new CardAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAccountInfoWithDefaults

`func NewCardAccountInfoWithDefaults() *CardAccountInfo`

NewCardAccountInfoWithDefaults instantiates a new CardAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardAccountInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardAccountInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardAccountInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardAccountInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCurrencies

`func (o *CardAccountInfo) GetCurrencies() []CardAccountCurrencyInfo`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *CardAccountInfo) GetCurrenciesOk() (*[]CardAccountCurrencyInfo, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *CardAccountInfo) SetCurrencies(v []CardAccountCurrencyInfo)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *CardAccountInfo) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


