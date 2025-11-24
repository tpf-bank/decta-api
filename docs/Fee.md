# Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Internal fee code | [optional] 
**Description** | Pointer to **string** | Internal fee description | [optional] 
**Ccy** | Pointer to **string** | Currency code of the deducted fee | [optional] 
**Value** | Pointer to **string** | Amount of the deducted fee | [optional] 

## Methods

### NewFee

`func NewFee() *Fee`

NewFee instantiates a new Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeWithDefaults

`func NewFeeWithDefaults() *Fee`

NewFeeWithDefaults instantiates a new Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Fee) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Fee) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Fee) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Fee) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *Fee) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Fee) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Fee) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Fee) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCcy

`func (o *Fee) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *Fee) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *Fee) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *Fee) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetValue

`func (o *Fee) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Fee) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Fee) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Fee) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


