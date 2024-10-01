# TransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An identification number of the transaction | [optional] 
**PostDate** | Pointer to **string** | Financial transactions settlement date in the system | [optional] 
**Type** | Pointer to **string** | Internal transaction code | [optional] 
**TypeDesc** | Pointer to **string** | Internal transaction description | [optional] 
**Stan** | Pointer to **string** | System Trace Audit Number - a unique number identifying a payment transaction | [optional] 
**Time** | Pointer to **string** | Date and time of the authorization | [optional] 
**Description** | Pointer to **string** | transaction time | [optional] 
**Ppan** | Pointer to **string** | A card pseudo number | [optional] 
**ReferenceData** | Pointer to **string** | Acquirer reference number | [optional] 
**ReferenceNumber** | Pointer to **string** | to be defined later | [optional] 
**SlipNumber** | Pointer to **string** | to be defined later | [optional] 
**ApprovalCode** | Pointer to **string** | A code that indicates approval or denial for an authorization request | [optional] 
**MsgType** | Pointer to **string** | to be defined later | [optional] 
**HoldNumber** | Pointer to **string** | A number identifying a hold | [optional] 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AccountCcyAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Fees** | Pointer to [**[]Fee**](Fee.md) | Transaction fees | [optional] 
**ExchangeRate** | Pointer to **string** | Transaction exchange rate value | [optional] 
**Markup** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewTransactionInfo

`func NewTransactionInfo() *TransactionInfo`

NewTransactionInfo instantiates a new TransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoWithDefaults

`func NewTransactionInfoWithDefaults() *TransactionInfo`

NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPostDate

`func (o *TransactionInfo) GetPostDate() string`

GetPostDate returns the PostDate field if non-nil, zero value otherwise.

### GetPostDateOk

`func (o *TransactionInfo) GetPostDateOk() (*string, bool)`

GetPostDateOk returns a tuple with the PostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDate

`func (o *TransactionInfo) SetPostDate(v string)`

SetPostDate sets PostDate field to given value.

### HasPostDate

`func (o *TransactionInfo) HasPostDate() bool`

HasPostDate returns a boolean if a field has been set.

### GetType

`func (o *TransactionInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeDesc

`func (o *TransactionInfo) GetTypeDesc() string`

GetTypeDesc returns the TypeDesc field if non-nil, zero value otherwise.

### GetTypeDescOk

`func (o *TransactionInfo) GetTypeDescOk() (*string, bool)`

GetTypeDescOk returns a tuple with the TypeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDesc

`func (o *TransactionInfo) SetTypeDesc(v string)`

SetTypeDesc sets TypeDesc field to given value.

### HasTypeDesc

`func (o *TransactionInfo) HasTypeDesc() bool`

HasTypeDesc returns a boolean if a field has been set.

### GetStan

`func (o *TransactionInfo) GetStan() string`

GetStan returns the Stan field if non-nil, zero value otherwise.

### GetStanOk

`func (o *TransactionInfo) GetStanOk() (*string, bool)`

GetStanOk returns a tuple with the Stan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStan

`func (o *TransactionInfo) SetStan(v string)`

SetStan sets Stan field to given value.

### HasStan

`func (o *TransactionInfo) HasStan() bool`

HasStan returns a boolean if a field has been set.

### GetTime

`func (o *TransactionInfo) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TransactionInfo) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TransactionInfo) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *TransactionInfo) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPpan

`func (o *TransactionInfo) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *TransactionInfo) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *TransactionInfo) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *TransactionInfo) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetReferenceData

`func (o *TransactionInfo) GetReferenceData() string`

GetReferenceData returns the ReferenceData field if non-nil, zero value otherwise.

### GetReferenceDataOk

`func (o *TransactionInfo) GetReferenceDataOk() (*string, bool)`

GetReferenceDataOk returns a tuple with the ReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceData

`func (o *TransactionInfo) SetReferenceData(v string)`

SetReferenceData sets ReferenceData field to given value.

### HasReferenceData

`func (o *TransactionInfo) HasReferenceData() bool`

HasReferenceData returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *TransactionInfo) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *TransactionInfo) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *TransactionInfo) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *TransactionInfo) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetSlipNumber

`func (o *TransactionInfo) GetSlipNumber() string`

GetSlipNumber returns the SlipNumber field if non-nil, zero value otherwise.

### GetSlipNumberOk

`func (o *TransactionInfo) GetSlipNumberOk() (*string, bool)`

GetSlipNumberOk returns a tuple with the SlipNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlipNumber

`func (o *TransactionInfo) SetSlipNumber(v string)`

SetSlipNumber sets SlipNumber field to given value.

### HasSlipNumber

`func (o *TransactionInfo) HasSlipNumber() bool`

HasSlipNumber returns a boolean if a field has been set.

### GetApprovalCode

`func (o *TransactionInfo) GetApprovalCode() string`

GetApprovalCode returns the ApprovalCode field if non-nil, zero value otherwise.

### GetApprovalCodeOk

`func (o *TransactionInfo) GetApprovalCodeOk() (*string, bool)`

GetApprovalCodeOk returns a tuple with the ApprovalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalCode

`func (o *TransactionInfo) SetApprovalCode(v string)`

SetApprovalCode sets ApprovalCode field to given value.

### HasApprovalCode

`func (o *TransactionInfo) HasApprovalCode() bool`

HasApprovalCode returns a boolean if a field has been set.

### GetMsgType

`func (o *TransactionInfo) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *TransactionInfo) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *TransactionInfo) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *TransactionInfo) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetHoldNumber

`func (o *TransactionInfo) GetHoldNumber() string`

GetHoldNumber returns the HoldNumber field if non-nil, zero value otherwise.

### GetHoldNumberOk

`func (o *TransactionInfo) GetHoldNumberOk() (*string, bool)`

GetHoldNumberOk returns a tuple with the HoldNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldNumber

`func (o *TransactionInfo) SetHoldNumber(v string)`

SetHoldNumber sets HoldNumber field to given value.

### HasHoldNumber

`func (o *TransactionInfo) HasHoldNumber() bool`

HasHoldNumber returns a boolean if a field has been set.

### GetMerchant

`func (o *TransactionInfo) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransactionInfo) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransactionInfo) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransactionInfo) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionInfo) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionInfo) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionInfo) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAccountCcyAmount

`func (o *TransactionInfo) GetAccountCcyAmount() Amount`

GetAccountCcyAmount returns the AccountCcyAmount field if non-nil, zero value otherwise.

### GetAccountCcyAmountOk

`func (o *TransactionInfo) GetAccountCcyAmountOk() (*Amount, bool)`

GetAccountCcyAmountOk returns a tuple with the AccountCcyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCcyAmount

`func (o *TransactionInfo) SetAccountCcyAmount(v Amount)`

SetAccountCcyAmount sets AccountCcyAmount field to given value.

### HasAccountCcyAmount

`func (o *TransactionInfo) HasAccountCcyAmount() bool`

HasAccountCcyAmount returns a boolean if a field has been set.

### GetFees

`func (o *TransactionInfo) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *TransactionInfo) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *TransactionInfo) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *TransactionInfo) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetExchangeRate

`func (o *TransactionInfo) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *TransactionInfo) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *TransactionInfo) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *TransactionInfo) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetMarkup

`func (o *TransactionInfo) GetMarkup() Amount`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *TransactionInfo) GetMarkupOk() (*Amount, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *TransactionInfo) SetMarkup(v Amount)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *TransactionInfo) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


