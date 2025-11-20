# Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**TokenReferenceData** | Pointer to [**[]TokenReferenceData**](TokenReferenceData.md) | Token Reference Data | [optional] 

## Methods

### NewData

`func NewData() *Data`

NewData instantiates a new Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWithDefaults

`func NewDataWithDefaults() *Data`

NewDataWithDefaults instantiates a new Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *Data) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *Data) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *Data) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *Data) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetTokenReferenceData

`func (o *Data) GetTokenReferenceData() []TokenReferenceData`

GetTokenReferenceData returns the TokenReferenceData field if non-nil, zero value otherwise.

### GetTokenReferenceDataOk

`func (o *Data) GetTokenReferenceDataOk() (*[]TokenReferenceData, bool)`

GetTokenReferenceDataOk returns a tuple with the TokenReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReferenceData

`func (o *Data) SetTokenReferenceData(v []TokenReferenceData)`

SetTokenReferenceData sets TokenReferenceData field to given value.

### HasTokenReferenceData

`func (o *Data) HasTokenReferenceData() bool`

HasTokenReferenceData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


