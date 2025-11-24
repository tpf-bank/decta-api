# CardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id | [optional] 
**Client** | Pointer to [**ClientInfo**](ClientInfo.md) |  | [optional] 
**CardAccount** | Pointer to [**CardAccountInfo**](CardAccountInfo.md) |  | [optional] 
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**CardName** | Pointer to **string** | Name on card | [optional] 
**Product** | Pointer to **string** | Product code | [optional] 
**State** | Pointer to **string** | Card state | [optional] 
**Expiry** | Pointer to **string** | Card expiry date | [optional] 
**Design** | Pointer to **string** | Design ID for plastic card product provided by Decta | [optional] 
**Priority** | Pointer to **string** | Priority flag for plastic card embossing | [optional] 

## Methods

### NewCardInfo

`func NewCardInfo() *CardInfo`

NewCardInfo instantiates a new CardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoWithDefaults

`func NewCardInfoWithDefaults() *CardInfo`

NewCardInfoWithDefaults instantiates a new CardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClient

`func (o *CardInfo) GetClient() ClientInfo`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CardInfo) GetClientOk() (*ClientInfo, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CardInfo) SetClient(v ClientInfo)`

SetClient sets Client field to given value.

### HasClient

`func (o *CardInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCardAccount

`func (o *CardInfo) GetCardAccount() CardAccountInfo`

GetCardAccount returns the CardAccount field if non-nil, zero value otherwise.

### GetCardAccountOk

`func (o *CardInfo) GetCardAccountOk() (*CardAccountInfo, bool)`

GetCardAccountOk returns a tuple with the CardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAccount

`func (o *CardInfo) SetCardAccount(v CardAccountInfo)`

SetCardAccount sets CardAccount field to given value.

### HasCardAccount

`func (o *CardInfo) HasCardAccount() bool`

HasCardAccount returns a boolean if a field has been set.

### GetPpan

`func (o *CardInfo) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *CardInfo) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *CardInfo) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *CardInfo) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetCardName

`func (o *CardInfo) GetCardName() string`

GetCardName returns the CardName field if non-nil, zero value otherwise.

### GetCardNameOk

`func (o *CardInfo) GetCardNameOk() (*string, bool)`

GetCardNameOk returns a tuple with the CardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardName

`func (o *CardInfo) SetCardName(v string)`

SetCardName sets CardName field to given value.

### HasCardName

`func (o *CardInfo) HasCardName() bool`

HasCardName returns a boolean if a field has been set.

### GetProduct

`func (o *CardInfo) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CardInfo) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CardInfo) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CardInfo) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetState

`func (o *CardInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CardInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetExpiry

`func (o *CardInfo) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CardInfo) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CardInfo) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *CardInfo) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetDesign

`func (o *CardInfo) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *CardInfo) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *CardInfo) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *CardInfo) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetPriority

`func (o *CardInfo) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CardInfo) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CardInfo) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CardInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


