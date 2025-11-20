# CardData3HolderApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**CardData** | Pointer to [**[]CardData3**](CardData3.md) | CardData | [optional] 

## Methods

### NewCardData3HolderApiDto

`func NewCardData3HolderApiDto() *CardData3HolderApiDto`

NewCardData3HolderApiDto instantiates a new CardData3HolderApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardData3HolderApiDtoWithDefaults

`func NewCardData3HolderApiDtoWithDefaults() *CardData3HolderApiDto`

NewCardData3HolderApiDtoWithDefaults instantiates a new CardData3HolderApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *CardData3HolderApiDto) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *CardData3HolderApiDto) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *CardData3HolderApiDto) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *CardData3HolderApiDto) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetCardData

`func (o *CardData3HolderApiDto) GetCardData() []CardData3`

GetCardData returns the CardData field if non-nil, zero value otherwise.

### GetCardDataOk

`func (o *CardData3HolderApiDto) GetCardDataOk() (*[]CardData3, bool)`

GetCardDataOk returns a tuple with the CardData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardData

`func (o *CardData3HolderApiDto) SetCardData(v []CardData3)`

SetCardData sets CardData field to given value.

### HasCardData

`func (o *CardData3HolderApiDto) HasCardData() bool`

HasCardData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


