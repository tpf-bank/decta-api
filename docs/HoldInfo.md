# HoldInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Internal transaction code | [optional] 
**TypeDesc** | Pointer to **string** | Internal transaction description | [optional] 
**Time** | Pointer to **string** | Date and time of the authorization | [optional] 
**Description** | Pointer to **string** | transaction time | [optional] 
**Ppan** | Pointer to **string** | A card pseudo number | [optional] 
**ActionCode** | Pointer to **string** | A code that indicates action code of authorization request | [optional] 
**HoldNumber** | Pointer to **string** | A number identifying a hold | [optional] 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AccountCcyAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Fees** | Pointer to [**[]Fee**](Fee.md) | Transaction fees | [optional] 

## Methods

### NewHoldInfo

`func NewHoldInfo() *HoldInfo`

NewHoldInfo instantiates a new HoldInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldInfoWithDefaults

`func NewHoldInfoWithDefaults() *HoldInfo`

NewHoldInfoWithDefaults instantiates a new HoldInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HoldInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HoldInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HoldInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HoldInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeDesc

`func (o *HoldInfo) GetTypeDesc() string`

GetTypeDesc returns the TypeDesc field if non-nil, zero value otherwise.

### GetTypeDescOk

`func (o *HoldInfo) GetTypeDescOk() (*string, bool)`

GetTypeDescOk returns a tuple with the TypeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDesc

`func (o *HoldInfo) SetTypeDesc(v string)`

SetTypeDesc sets TypeDesc field to given value.

### HasTypeDesc

`func (o *HoldInfo) HasTypeDesc() bool`

HasTypeDesc returns a boolean if a field has been set.

### GetTime

`func (o *HoldInfo) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HoldInfo) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HoldInfo) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *HoldInfo) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDescription

`func (o *HoldInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPpan

`func (o *HoldInfo) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *HoldInfo) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *HoldInfo) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *HoldInfo) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetActionCode

`func (o *HoldInfo) GetActionCode() string`

GetActionCode returns the ActionCode field if non-nil, zero value otherwise.

### GetActionCodeOk

`func (o *HoldInfo) GetActionCodeOk() (*string, bool)`

GetActionCodeOk returns a tuple with the ActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCode

`func (o *HoldInfo) SetActionCode(v string)`

SetActionCode sets ActionCode field to given value.

### HasActionCode

`func (o *HoldInfo) HasActionCode() bool`

HasActionCode returns a boolean if a field has been set.

### GetHoldNumber

`func (o *HoldInfo) GetHoldNumber() string`

GetHoldNumber returns the HoldNumber field if non-nil, zero value otherwise.

### GetHoldNumberOk

`func (o *HoldInfo) GetHoldNumberOk() (*string, bool)`

GetHoldNumberOk returns a tuple with the HoldNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldNumber

`func (o *HoldInfo) SetHoldNumber(v string)`

SetHoldNumber sets HoldNumber field to given value.

### HasHoldNumber

`func (o *HoldInfo) HasHoldNumber() bool`

HasHoldNumber returns a boolean if a field has been set.

### GetMerchant

`func (o *HoldInfo) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *HoldInfo) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *HoldInfo) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *HoldInfo) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetAmount

`func (o *HoldInfo) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HoldInfo) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HoldInfo) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *HoldInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAccountCcyAmount

`func (o *HoldInfo) GetAccountCcyAmount() Amount`

GetAccountCcyAmount returns the AccountCcyAmount field if non-nil, zero value otherwise.

### GetAccountCcyAmountOk

`func (o *HoldInfo) GetAccountCcyAmountOk() (*Amount, bool)`

GetAccountCcyAmountOk returns a tuple with the AccountCcyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCcyAmount

`func (o *HoldInfo) SetAccountCcyAmount(v Amount)`

SetAccountCcyAmount sets AccountCcyAmount field to given value.

### HasAccountCcyAmount

`func (o *HoldInfo) HasAccountCcyAmount() bool`

HasAccountCcyAmount returns a boolean if a field has been set.

### GetFees

`func (o *HoldInfo) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *HoldInfo) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *HoldInfo) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *HoldInfo) HasFees() bool`

HasFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


