# Amount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | **string** | Currency alpha code according to ISO 4217 | 
**Value** | **string** | Operation amount | 

## Methods

### NewAmount

`func NewAmount(ccy string, value string, ) *Amount`

NewAmount instantiates a new Amount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountWithDefaults

`func NewAmountWithDefaults() *Amount`

NewAmountWithDefaults instantiates a new Amount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *Amount) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *Amount) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *Amount) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetValue

`func (o *Amount) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Amount) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Amount) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


