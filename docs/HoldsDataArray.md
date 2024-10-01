# HoldsDataArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]HoldInfo**](HoldInfo.md) | Requested holds data array | [optional] 
**NextId** | Pointer to **string** | Next item cursor. No more data if cursor is empty | [optional] 

## Methods

### NewHoldsDataArray

`func NewHoldsDataArray() *HoldsDataArray`

NewHoldsDataArray instantiates a new HoldsDataArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldsDataArrayWithDefaults

`func NewHoldsDataArrayWithDefaults() *HoldsDataArray`

NewHoldsDataArrayWithDefaults instantiates a new HoldsDataArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HoldsDataArray) GetData() []HoldInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HoldsDataArray) GetDataOk() (*[]HoldInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HoldsDataArray) SetData(v []HoldInfo)`

SetData sets Data field to given value.

### HasData

`func (o *HoldsDataArray) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextId

`func (o *HoldsDataArray) GetNextId() string`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *HoldsDataArray) GetNextIdOk() (*string, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *HoldsDataArray) SetNextId(v string)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *HoldsDataArray) HasNextId() bool`

HasNextId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


