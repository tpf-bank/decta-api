# ClientInfoDataArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ClientInfo**](ClientInfo.md) | Requested client info data array | [optional] 
**NextId** | Pointer to **string** | Next item cursor. No more data if cursor is empty | [optional] 

## Methods

### NewClientInfoDataArray

`func NewClientInfoDataArray() *ClientInfoDataArray`

NewClientInfoDataArray instantiates a new ClientInfoDataArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientInfoDataArrayWithDefaults

`func NewClientInfoDataArrayWithDefaults() *ClientInfoDataArray`

NewClientInfoDataArrayWithDefaults instantiates a new ClientInfoDataArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ClientInfoDataArray) GetData() []ClientInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientInfoDataArray) GetDataOk() (*[]ClientInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientInfoDataArray) SetData(v []ClientInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ClientInfoDataArray) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextId

`func (o *ClientInfoDataArray) GetNextId() string`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *ClientInfoDataArray) GetNextIdOk() (*string, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *ClientInfoDataArray) SetNextId(v string)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *ClientInfoDataArray) HasNextId() bool`

HasNextId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


