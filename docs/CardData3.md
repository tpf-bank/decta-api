# CardData3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pan** | Pointer to **string** | Clear text primary account number | [optional] 
**CardName** | Pointer to **string** | Cardholder name up to 26 symbols | [optional] 
**Expiry** | Pointer to **string** | Card expiry in format MMyy | [optional] 

## Methods

### NewCardData3

`func NewCardData3() *CardData3`

NewCardData3 instantiates a new CardData3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardData3WithDefaults

`func NewCardData3WithDefaults() *CardData3`

NewCardData3WithDefaults instantiates a new CardData3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPan

`func (o *CardData3) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *CardData3) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *CardData3) SetPan(v string)`

SetPan sets Pan field to given value.

### HasPan

`func (o *CardData3) HasPan() bool`

HasPan returns a boolean if a field has been set.

### GetCardName

`func (o *CardData3) GetCardName() string`

GetCardName returns the CardName field if non-nil, zero value otherwise.

### GetCardNameOk

`func (o *CardData3) GetCardNameOk() (*string, bool)`

GetCardNameOk returns a tuple with the CardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardName

`func (o *CardData3) SetCardName(v string)`

SetCardName sets CardName field to given value.

### HasCardName

`func (o *CardData3) HasCardName() bool`

HasCardName returns a boolean if a field has been set.

### GetExpiry

`func (o *CardData3) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CardData3) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CardData3) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *CardData3) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


