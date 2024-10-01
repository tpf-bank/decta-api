# CardData4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pan** | Pointer to **string** | Clear text primary account number | [optional] 
**CardName** | Pointer to **string** | Cardholder name up to 26 symbols | [optional] 
**Expiry** | Pointer to **string** | Card expiry in format MMyy | [optional] 
**Cvv2** | Pointer to **string** | Generated Card CVV | [optional] 

## Methods

### NewCardData4

`func NewCardData4() *CardData4`

NewCardData4 instantiates a new CardData4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardData4WithDefaults

`func NewCardData4WithDefaults() *CardData4`

NewCardData4WithDefaults instantiates a new CardData4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPan

`func (o *CardData4) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *CardData4) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *CardData4) SetPan(v string)`

SetPan sets Pan field to given value.

### HasPan

`func (o *CardData4) HasPan() bool`

HasPan returns a boolean if a field has been set.

### GetCardName

`func (o *CardData4) GetCardName() string`

GetCardName returns the CardName field if non-nil, zero value otherwise.

### GetCardNameOk

`func (o *CardData4) GetCardNameOk() (*string, bool)`

GetCardNameOk returns a tuple with the CardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardName

`func (o *CardData4) SetCardName(v string)`

SetCardName sets CardName field to given value.

### HasCardName

`func (o *CardData4) HasCardName() bool`

HasCardName returns a boolean if a field has been set.

### GetExpiry

`func (o *CardData4) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CardData4) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CardData4) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *CardData4) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetCvv2

`func (o *CardData4) GetCvv2() string`

GetCvv2 returns the Cvv2 field if non-nil, zero value otherwise.

### GetCvv2Ok

`func (o *CardData4) GetCvv2Ok() (*string, bool)`

GetCvv2Ok returns a tuple with the Cvv2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv2

`func (o *CardData4) SetCvv2(v string)`

SetCvv2 sets Cvv2 field to given value.

### HasCvv2

`func (o *CardData4) HasCvv2() bool`

HasCvv2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


