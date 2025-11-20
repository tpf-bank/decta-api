# LimitsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**Limits** | [**[]Limit**](Limit.md) |  | 

## Methods

### NewLimitsData

`func NewLimitsData(limits []Limit, ) *LimitsData`

NewLimitsData instantiates a new LimitsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsDataWithDefaults

`func NewLimitsDataWithDefaults() *LimitsData`

NewLimitsDataWithDefaults instantiates a new LimitsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *LimitsData) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *LimitsData) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *LimitsData) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *LimitsData) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetLimits

`func (o *LimitsData) GetLimits() []Limit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *LimitsData) GetLimitsOk() (*[]Limit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *LimitsData) SetLimits(v []Limit)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


