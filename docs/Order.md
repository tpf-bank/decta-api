# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Order id | [optional] 
**Created** | Pointer to **string** | Order registration date and time | [optional] 
**Status** | Pointer to **string** | Order status | [optional] 
**Ppan** | Pointer to **string** | masked card number | [optional] 
**CardAccount** | Pointer to **string** | Card account number | [optional] 
**ClientId** | Pointer to **string** | Decta client id number | [optional] 
**ExternalId** | Pointer to **string** | Decta API user own order identification number | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Order) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Order) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Order) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Order) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPpan

`func (o *Order) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *Order) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *Order) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *Order) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetCardAccount

`func (o *Order) GetCardAccount() string`

GetCardAccount returns the CardAccount field if non-nil, zero value otherwise.

### GetCardAccountOk

`func (o *Order) GetCardAccountOk() (*string, bool)`

GetCardAccountOk returns a tuple with the CardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAccount

`func (o *Order) SetCardAccount(v string)`

SetCardAccount sets CardAccount field to given value.

### HasCardAccount

`func (o *Order) HasCardAccount() bool`

HasCardAccount returns a boolean if a field has been set.

### GetClientId

`func (o *Order) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Order) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Order) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Order) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetExternalId

`func (o *Order) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Order) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Order) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Order) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


