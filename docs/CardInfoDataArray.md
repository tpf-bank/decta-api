# CardInfoDataArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CardInfo**](CardInfo.md) | Requested card info data array | [optional] 
**NextId** | Pointer to **string** | Next item cursor. No more data if cursor is empty | [optional] 

## Methods

### NewCardInfoDataArray

`func NewCardInfoDataArray() *CardInfoDataArray`

NewCardInfoDataArray instantiates a new CardInfoDataArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoDataArrayWithDefaults

`func NewCardInfoDataArrayWithDefaults() *CardInfoDataArray`

NewCardInfoDataArrayWithDefaults instantiates a new CardInfoDataArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CardInfoDataArray) GetData() []CardInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CardInfoDataArray) GetDataOk() (*[]CardInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CardInfoDataArray) SetData(v []CardInfo)`

SetData sets Data field to given value.

### HasData

`func (o *CardInfoDataArray) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextId

`func (o *CardInfoDataArray) GetNextId() string`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *CardInfoDataArray) GetNextIdOk() (*string, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *CardInfoDataArray) SetNextId(v string)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *CardInfoDataArray) HasNextId() bool`

HasNextId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


