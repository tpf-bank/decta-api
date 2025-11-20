# CreateCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Masked card number | [optional] 
**CardAccount** | Pointer to **string** | Card account number | [optional] 
**ClientId** | Pointer to **string** | ID of customer on Decta partner side (allows to bind multiple cards to one customer). | [optional] 

## Methods

### NewCreateCardResponse

`func NewCreateCardResponse() *CreateCardResponse`

NewCreateCardResponse instantiates a new CreateCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCardResponseWithDefaults

`func NewCreateCardResponseWithDefaults() *CreateCardResponse`

NewCreateCardResponseWithDefaults instantiates a new CreateCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *CreateCardResponse) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *CreateCardResponse) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *CreateCardResponse) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *CreateCardResponse) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetCardAccount

`func (o *CreateCardResponse) GetCardAccount() string`

GetCardAccount returns the CardAccount field if non-nil, zero value otherwise.

### GetCardAccountOk

`func (o *CreateCardResponse) GetCardAccountOk() (*string, bool)`

GetCardAccountOk returns a tuple with the CardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAccount

`func (o *CreateCardResponse) SetCardAccount(v string)`

SetCardAccount sets CardAccount field to given value.

### HasCardAccount

`func (o *CreateCardResponse) HasCardAccount() bool`

HasCardAccount returns a boolean if a field has been set.

### GetClientId

`func (o *CreateCardResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateCardResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateCardResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateCardResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


