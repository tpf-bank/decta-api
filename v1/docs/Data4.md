# Data4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**CardData** | Pointer to [**[]CardData4**](CardData4.md) | CardData | [optional] 

## Methods

### NewData4

`func NewData4() *Data4`

NewData4 instantiates a new Data4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewData4WithDefaults

`func NewData4WithDefaults() *Data4`

NewData4WithDefaults instantiates a new Data4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *Data4) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *Data4) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *Data4) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *Data4) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetCardData

`func (o *Data4) GetCardData() []CardData4`

GetCardData returns the CardData field if non-nil, zero value otherwise.

### GetCardDataOk

`func (o *Data4) GetCardDataOk() (*[]CardData4, bool)`

GetCardDataOk returns a tuple with the CardData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardData

`func (o *Data4) SetCardData(v []CardData4)`

SetCardData sets CardData field to given value.

### HasCardData

`func (o *Data4) HasCardData() bool`

HasCardData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


