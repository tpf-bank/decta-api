# CreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateClient** | Pointer to [**PrivateClientCreateCard**](PrivateClientCreateCard.md) |  | [optional] 
**BusinessClient** | Pointer to [**BusinessClientCreateCard**](BusinessClientCreateCard.md) |  | [optional] 
**Card** | [**CardPreferencesCreateCard**](CardPreferencesCreateCard.md) |  | 

## Methods

### NewCreateCard

`func NewCreateCard(card CardPreferencesCreateCard, ) *CreateCard`

NewCreateCard instantiates a new CreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCardWithDefaults

`func NewCreateCardWithDefaults() *CreateCard`

NewCreateCardWithDefaults instantiates a new CreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateClient

`func (o *CreateCard) GetPrivateClient() PrivateClientCreateCard`

GetPrivateClient returns the PrivateClient field if non-nil, zero value otherwise.

### GetPrivateClientOk

`func (o *CreateCard) GetPrivateClientOk() (*PrivateClientCreateCard, bool)`

GetPrivateClientOk returns a tuple with the PrivateClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateClient

`func (o *CreateCard) SetPrivateClient(v PrivateClientCreateCard)`

SetPrivateClient sets PrivateClient field to given value.

### HasPrivateClient

`func (o *CreateCard) HasPrivateClient() bool`

HasPrivateClient returns a boolean if a field has been set.

### GetBusinessClient

`func (o *CreateCard) GetBusinessClient() BusinessClientCreateCard`

GetBusinessClient returns the BusinessClient field if non-nil, zero value otherwise.

### GetBusinessClientOk

`func (o *CreateCard) GetBusinessClientOk() (*BusinessClientCreateCard, bool)`

GetBusinessClientOk returns a tuple with the BusinessClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessClient

`func (o *CreateCard) SetBusinessClient(v BusinessClientCreateCard)`

SetBusinessClient sets BusinessClient field to given value.

### HasBusinessClient

`func (o *CreateCard) HasBusinessClient() bool`

HasBusinessClient returns a boolean if a field has been set.

### GetCard

`func (o *CreateCard) GetCard() CardPreferencesCreateCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CreateCard) GetCardOk() (*CardPreferencesCreateCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CreateCard) SetCard(v CardPreferencesCreateCard)`

SetCard sets Card field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


