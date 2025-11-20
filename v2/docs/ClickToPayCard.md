# ClickToPayCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | A card pseudo number | [optional] 
**Address** | Pointer to [**C2PAddressCreateCard**](C2PAddressCreateCard.md) |  | [optional] 

## Methods

### NewClickToPayCard

`func NewClickToPayCard() *ClickToPayCard`

NewClickToPayCard instantiates a new ClickToPayCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickToPayCardWithDefaults

`func NewClickToPayCardWithDefaults() *ClickToPayCard`

NewClickToPayCardWithDefaults instantiates a new ClickToPayCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *ClickToPayCard) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *ClickToPayCard) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *ClickToPayCard) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *ClickToPayCard) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetAddress

`func (o *ClickToPayCard) GetAddress() C2PAddressCreateCard`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClickToPayCard) GetAddressOk() (*C2PAddressCreateCard, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClickToPayCard) SetAddress(v C2PAddressCreateCard)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClickToPayCard) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


