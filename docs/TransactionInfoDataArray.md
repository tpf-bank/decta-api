# TransactionInfoDataArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TransactionInfo**](TransactionInfo.md) | Requested transaction info data array | [optional] 
**NextId** | Pointer to **string** | Next item cursor. No more data if cursor is empty | [optional] 

## Methods

### NewTransactionInfoDataArray

`func NewTransactionInfoDataArray() *TransactionInfoDataArray`

NewTransactionInfoDataArray instantiates a new TransactionInfoDataArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoDataArrayWithDefaults

`func NewTransactionInfoDataArrayWithDefaults() *TransactionInfoDataArray`

NewTransactionInfoDataArrayWithDefaults instantiates a new TransactionInfoDataArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionInfoDataArray) GetData() []TransactionInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionInfoDataArray) GetDataOk() (*[]TransactionInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionInfoDataArray) SetData(v []TransactionInfo)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionInfoDataArray) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextId

`func (o *TransactionInfoDataArray) GetNextId() string`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *TransactionInfoDataArray) GetNextIdOk() (*string, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *TransactionInfoDataArray) SetNextId(v string)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *TransactionInfoDataArray) HasNextId() bool`

HasNextId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


