# PinTryCounterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PinTryCount** | Pointer to **string** | PIN try count | [optional] 
**UpdateDate** | Pointer to **string** | Last unsuccessful date of online pin try | [optional] 

## Methods

### NewPinTryCounterInfo

`func NewPinTryCounterInfo() *PinTryCounterInfo`

NewPinTryCounterInfo instantiates a new PinTryCounterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinTryCounterInfoWithDefaults

`func NewPinTryCounterInfoWithDefaults() *PinTryCounterInfo`

NewPinTryCounterInfoWithDefaults instantiates a new PinTryCounterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPinTryCount

`func (o *PinTryCounterInfo) GetPinTryCount() string`

GetPinTryCount returns the PinTryCount field if non-nil, zero value otherwise.

### GetPinTryCountOk

`func (o *PinTryCounterInfo) GetPinTryCountOk() (*string, bool)`

GetPinTryCountOk returns a tuple with the PinTryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinTryCount

`func (o *PinTryCounterInfo) SetPinTryCount(v string)`

SetPinTryCount sets PinTryCount field to given value.

### HasPinTryCount

`func (o *PinTryCounterInfo) HasPinTryCount() bool`

HasPinTryCount returns a boolean if a field has been set.

### GetUpdateDate

`func (o *PinTryCounterInfo) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *PinTryCounterInfo) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *PinTryCounterInfo) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *PinTryCounterInfo) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


