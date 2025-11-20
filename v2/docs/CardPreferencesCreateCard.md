# CardPreferencesCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **string** | Priority flag for plastic card embossing | [optional] 
**ProductCode** | **string** | Product code for virtual or plastic card product provided by Decta | 
**Design** | Pointer to **string** | Design ID for plastic card product provided by Decta. Mandatory only for plastic cards | [optional] 
**CardAccount** | Pointer to **string** | Card Account. In case of supplementary card will be used to attach card to it. | [optional] 
**Currencies** | **string** | Card account currencies list ordered by priority decreasing | 
**Holder** | Pointer to [**CardholderCreateCard**](CardholderCreateCard.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**DeliveryAddressCreate**](DeliveryAddressCreate.md) |  | [optional] 
**Supplementary** | **string** | Supplementary card mark | 
**Passphrase** | Pointer to **string** | Passphrase | [optional] 
**RangeId** | Pointer to **string** | Range id for virtual or plastic card product provided by Decta | [optional] 
**CondSet** | Pointer to **string** | Card condition for virtual or plastic card product provided by Decta | [optional] 
**AccountCondSet** | Pointer to **string** | Card account condition for virtual or plastic card product provided by Decta | [optional] 
**UserDefinedField8** | Pointer to **string** | Free field for custom use | [optional] 
**TranzAccount** | Pointer to **string** | Card transaction account | [optional] 
**UserDefinedField12** | Pointer to **string** | Free field for custom use | [optional] 
**UserDefinedField7** | Pointer to **string** | Free field for custom use | [optional] 

## Methods

### NewCardPreferencesCreateCard

`func NewCardPreferencesCreateCard(productCode string, currencies string, supplementary string, ) *CardPreferencesCreateCard`

NewCardPreferencesCreateCard instantiates a new CardPreferencesCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPreferencesCreateCardWithDefaults

`func NewCardPreferencesCreateCardWithDefaults() *CardPreferencesCreateCard`

NewCardPreferencesCreateCardWithDefaults instantiates a new CardPreferencesCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *CardPreferencesCreateCard) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CardPreferencesCreateCard) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CardPreferencesCreateCard) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CardPreferencesCreateCard) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProductCode

`func (o *CardPreferencesCreateCard) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *CardPreferencesCreateCard) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *CardPreferencesCreateCard) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetDesign

`func (o *CardPreferencesCreateCard) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *CardPreferencesCreateCard) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *CardPreferencesCreateCard) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *CardPreferencesCreateCard) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetCardAccount

`func (o *CardPreferencesCreateCard) GetCardAccount() string`

GetCardAccount returns the CardAccount field if non-nil, zero value otherwise.

### GetCardAccountOk

`func (o *CardPreferencesCreateCard) GetCardAccountOk() (*string, bool)`

GetCardAccountOk returns a tuple with the CardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAccount

`func (o *CardPreferencesCreateCard) SetCardAccount(v string)`

SetCardAccount sets CardAccount field to given value.

### HasCardAccount

`func (o *CardPreferencesCreateCard) HasCardAccount() bool`

HasCardAccount returns a boolean if a field has been set.

### GetCurrencies

`func (o *CardPreferencesCreateCard) GetCurrencies() string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *CardPreferencesCreateCard) GetCurrenciesOk() (*string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *CardPreferencesCreateCard) SetCurrencies(v string)`

SetCurrencies sets Currencies field to given value.


### GetHolder

`func (o *CardPreferencesCreateCard) GetHolder() CardholderCreateCard`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *CardPreferencesCreateCard) GetHolderOk() (*CardholderCreateCard, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *CardPreferencesCreateCard) SetHolder(v CardholderCreateCard)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *CardPreferencesCreateCard) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *CardPreferencesCreateCard) GetDeliveryAddress() DeliveryAddressCreate`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CardPreferencesCreateCard) GetDeliveryAddressOk() (*DeliveryAddressCreate, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CardPreferencesCreateCard) SetDeliveryAddress(v DeliveryAddressCreate)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *CardPreferencesCreateCard) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetSupplementary

`func (o *CardPreferencesCreateCard) GetSupplementary() string`

GetSupplementary returns the Supplementary field if non-nil, zero value otherwise.

### GetSupplementaryOk

`func (o *CardPreferencesCreateCard) GetSupplementaryOk() (*string, bool)`

GetSupplementaryOk returns a tuple with the Supplementary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementary

`func (o *CardPreferencesCreateCard) SetSupplementary(v string)`

SetSupplementary sets Supplementary field to given value.


### GetPassphrase

`func (o *CardPreferencesCreateCard) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CardPreferencesCreateCard) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CardPreferencesCreateCard) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CardPreferencesCreateCard) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetRangeId

`func (o *CardPreferencesCreateCard) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *CardPreferencesCreateCard) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *CardPreferencesCreateCard) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *CardPreferencesCreateCard) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetCondSet

`func (o *CardPreferencesCreateCard) GetCondSet() string`

GetCondSet returns the CondSet field if non-nil, zero value otherwise.

### GetCondSetOk

`func (o *CardPreferencesCreateCard) GetCondSetOk() (*string, bool)`

GetCondSetOk returns a tuple with the CondSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondSet

`func (o *CardPreferencesCreateCard) SetCondSet(v string)`

SetCondSet sets CondSet field to given value.

### HasCondSet

`func (o *CardPreferencesCreateCard) HasCondSet() bool`

HasCondSet returns a boolean if a field has been set.

### GetAccountCondSet

`func (o *CardPreferencesCreateCard) GetAccountCondSet() string`

GetAccountCondSet returns the AccountCondSet field if non-nil, zero value otherwise.

### GetAccountCondSetOk

`func (o *CardPreferencesCreateCard) GetAccountCondSetOk() (*string, bool)`

GetAccountCondSetOk returns a tuple with the AccountCondSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCondSet

`func (o *CardPreferencesCreateCard) SetAccountCondSet(v string)`

SetAccountCondSet sets AccountCondSet field to given value.

### HasAccountCondSet

`func (o *CardPreferencesCreateCard) HasAccountCondSet() bool`

HasAccountCondSet returns a boolean if a field has been set.

### GetUserDefinedField8

`func (o *CardPreferencesCreateCard) GetUserDefinedField8() string`

GetUserDefinedField8 returns the UserDefinedField8 field if non-nil, zero value otherwise.

### GetUserDefinedField8Ok

`func (o *CardPreferencesCreateCard) GetUserDefinedField8Ok() (*string, bool)`

GetUserDefinedField8Ok returns a tuple with the UserDefinedField8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField8

`func (o *CardPreferencesCreateCard) SetUserDefinedField8(v string)`

SetUserDefinedField8 sets UserDefinedField8 field to given value.

### HasUserDefinedField8

`func (o *CardPreferencesCreateCard) HasUserDefinedField8() bool`

HasUserDefinedField8 returns a boolean if a field has been set.

### GetTranzAccount

`func (o *CardPreferencesCreateCard) GetTranzAccount() string`

GetTranzAccount returns the TranzAccount field if non-nil, zero value otherwise.

### GetTranzAccountOk

`func (o *CardPreferencesCreateCard) GetTranzAccountOk() (*string, bool)`

GetTranzAccountOk returns a tuple with the TranzAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranzAccount

`func (o *CardPreferencesCreateCard) SetTranzAccount(v string)`

SetTranzAccount sets TranzAccount field to given value.

### HasTranzAccount

`func (o *CardPreferencesCreateCard) HasTranzAccount() bool`

HasTranzAccount returns a boolean if a field has been set.

### GetUserDefinedField12

`func (o *CardPreferencesCreateCard) GetUserDefinedField12() string`

GetUserDefinedField12 returns the UserDefinedField12 field if non-nil, zero value otherwise.

### GetUserDefinedField12Ok

`func (o *CardPreferencesCreateCard) GetUserDefinedField12Ok() (*string, bool)`

GetUserDefinedField12Ok returns a tuple with the UserDefinedField12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField12

`func (o *CardPreferencesCreateCard) SetUserDefinedField12(v string)`

SetUserDefinedField12 sets UserDefinedField12 field to given value.

### HasUserDefinedField12

`func (o *CardPreferencesCreateCard) HasUserDefinedField12() bool`

HasUserDefinedField12 returns a boolean if a field has been set.

### GetUserDefinedField7

`func (o *CardPreferencesCreateCard) GetUserDefinedField7() string`

GetUserDefinedField7 returns the UserDefinedField7 field if non-nil, zero value otherwise.

### GetUserDefinedField7Ok

`func (o *CardPreferencesCreateCard) GetUserDefinedField7Ok() (*string, bool)`

GetUserDefinedField7Ok returns a tuple with the UserDefinedField7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField7

`func (o *CardPreferencesCreateCard) SetUserDefinedField7(v string)`

SetUserDefinedField7 sets UserDefinedField7 field to given value.

### HasUserDefinedField7

`func (o *CardPreferencesCreateCard) HasUserDefinedField7() bool`

HasUserDefinedField7 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


