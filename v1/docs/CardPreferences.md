# CardPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | Product code for virtual or plastic card product provided by Decta | 
**Design** | Pointer to **string** | Design ID for plastic card product provided by Decta. Mandatory only for plastic cards | [optional] 
**CardAccount** | Pointer to **string** | Card Account. In case of supplementary card will be used to attach card to it. | [optional] 
**Currencies** | **[]string** | Card account currencies list ordered by priority decreasing | 
**Holder** | Pointer to [**Cardholder**](Cardholder.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**DeliveryAddress**](DeliveryAddress.md) |  | [optional] 
**Supplementary** | **string** | Supplementary card mark | 
**AccountOwnerRelation** | Pointer to **string** | Card account owner relation | [optional] 
**Passphrase** | Pointer to **string** | Passphrase | [optional] 
**Priority** | Pointer to **string** | Priority flag for plastic card embossing | [optional] 
**RangeId** | Pointer to **string** | Range id for virtual or plastic card product provided by Decta | [optional] 
**UserDefinedField7** | Pointer to **string** | Free field for custom use | [optional] 

## Methods

### NewCardPreferences

`func NewCardPreferences(productCode string, currencies []string, supplementary string, ) *CardPreferences`

NewCardPreferences instantiates a new CardPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPreferencesWithDefaults

`func NewCardPreferencesWithDefaults() *CardPreferences`

NewCardPreferencesWithDefaults instantiates a new CardPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *CardPreferences) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *CardPreferences) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *CardPreferences) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetDesign

`func (o *CardPreferences) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *CardPreferences) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *CardPreferences) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *CardPreferences) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetCardAccount

`func (o *CardPreferences) GetCardAccount() string`

GetCardAccount returns the CardAccount field if non-nil, zero value otherwise.

### GetCardAccountOk

`func (o *CardPreferences) GetCardAccountOk() (*string, bool)`

GetCardAccountOk returns a tuple with the CardAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAccount

`func (o *CardPreferences) SetCardAccount(v string)`

SetCardAccount sets CardAccount field to given value.

### HasCardAccount

`func (o *CardPreferences) HasCardAccount() bool`

HasCardAccount returns a boolean if a field has been set.

### GetCurrencies

`func (o *CardPreferences) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *CardPreferences) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *CardPreferences) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.


### GetHolder

`func (o *CardPreferences) GetHolder() Cardholder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *CardPreferences) GetHolderOk() (*Cardholder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *CardPreferences) SetHolder(v Cardholder)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *CardPreferences) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *CardPreferences) GetDeliveryAddress() DeliveryAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CardPreferences) GetDeliveryAddressOk() (*DeliveryAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CardPreferences) SetDeliveryAddress(v DeliveryAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *CardPreferences) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetSupplementary

`func (o *CardPreferences) GetSupplementary() string`

GetSupplementary returns the Supplementary field if non-nil, zero value otherwise.

### GetSupplementaryOk

`func (o *CardPreferences) GetSupplementaryOk() (*string, bool)`

GetSupplementaryOk returns a tuple with the Supplementary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementary

`func (o *CardPreferences) SetSupplementary(v string)`

SetSupplementary sets Supplementary field to given value.


### GetAccountOwnerRelation

`func (o *CardPreferences) GetAccountOwnerRelation() string`

GetAccountOwnerRelation returns the AccountOwnerRelation field if non-nil, zero value otherwise.

### GetAccountOwnerRelationOk

`func (o *CardPreferences) GetAccountOwnerRelationOk() (*string, bool)`

GetAccountOwnerRelationOk returns a tuple with the AccountOwnerRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerRelation

`func (o *CardPreferences) SetAccountOwnerRelation(v string)`

SetAccountOwnerRelation sets AccountOwnerRelation field to given value.

### HasAccountOwnerRelation

`func (o *CardPreferences) HasAccountOwnerRelation() bool`

HasAccountOwnerRelation returns a boolean if a field has been set.

### GetPassphrase

`func (o *CardPreferences) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CardPreferences) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CardPreferences) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CardPreferences) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPriority

`func (o *CardPreferences) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CardPreferences) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CardPreferences) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CardPreferences) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRangeId

`func (o *CardPreferences) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *CardPreferences) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *CardPreferences) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *CardPreferences) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetUserDefinedField7

`func (o *CardPreferences) GetUserDefinedField7() string`

GetUserDefinedField7 returns the UserDefinedField7 field if non-nil, zero value otherwise.

### GetUserDefinedField7Ok

`func (o *CardPreferences) GetUserDefinedField7Ok() (*string, bool)`

GetUserDefinedField7Ok returns a tuple with the UserDefinedField7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField7

`func (o *CardPreferences) SetUserDefinedField7(v string)`

SetUserDefinedField7 sets UserDefinedField7 field to given value.

### HasUserDefinedField7

`func (o *CardPreferences) HasUserDefinedField7() bool`

HasUserDefinedField7 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


