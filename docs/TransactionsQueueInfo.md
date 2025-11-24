# TransactionsQueueInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An identification number of the transaction | [optional] 
**Time** | Pointer to **string** | Date and time of the transaction | [optional] 
**Type** | Pointer to **string** | Internal transaction code | [optional] 
**TypeDesc** | Pointer to **string** | Internal transaction description | [optional] 
**Description** | Pointer to **string** | Deal description | [optional] 
**Ppan** | Pointer to **string** | A card pseudo number | [optional] 
**CardAccount** | Pointer to **string** | Card account identification number | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewTransactionsQueueInfo

`func NewTransactionsQueueInfo() *TransactionsQueueInfo`

NewTransactionsQueueInfo instantiates a new TransactionsQueueInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsQueueInfoWithDefaults

`func NewTransactionsQueueInfoWithDefaults() *TransactionsQueueInfo`

NewTransactionsQueueInfoWithDefaults instantiates a new TransactionsQueueInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionsQueueInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionsQueueInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionsQueueInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionsQueueInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *TransactionsQueueInfo) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TransactionsQueueInfo) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TransactionsQueueInfo) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *TransactionsQueueInfo) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *TransactionsQueueInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionsQueueInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionsQueueInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionsQueueInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeDesc

`func (o *TransactionsQueueInfo) GetTypeDesc() string`

GetTypeDesc returns the TypeDesc field if non-nil, zero value otherwise.

### GetTypeDescOk

`func (o *TransactionsQueueInfo) GetTypeDescOk() (*string, bool)`

GetTypeDescOk returns a tuple with the TypeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDesc

`func (o *TransactionsQueueInfo) SetTypeDesc(v string)`

SetTypeDesc sets TypeDesc field to given value.

### HasTypeDesc

`func (o *TransactionsQueueInfo) HasTypeDesc() bool`

HasTypeDesc returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionsQueueInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionsQueueInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionsQueueInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionsQueueInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPpan

`func (o *TransactionsQueueInfo) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *TransactionsQueueInfo) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *TransactionsQueueInfo) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *TransactionsQueueInfo) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetCardAccount

`func (o *TransactionsQueueInfo) GetCardAccount() string`

GetCardAccount returns the CardAccount field if non-nil, zero value otherwise.

### GetCardAccountOk

`func (o *TransactionsQueueInfo) GetCardAccountOk() (*string, bool)`

GetCardAccountOk returns a tuple with the CardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAccount

`func (o *TransactionsQueueInfo) SetCardAccount(v string)`

SetCardAccount sets CardAccount field to given value.

### HasCardAccount

`func (o *TransactionsQueueInfo) HasCardAccount() bool`

HasCardAccount returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionsQueueInfo) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionsQueueInfo) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionsQueueInfo) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionsQueueInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


