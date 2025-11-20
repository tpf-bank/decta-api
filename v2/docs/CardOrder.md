# CardOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | Decta API user controlled field | [optional] 
**PrivateClient** | Pointer to [**PrivateClient**](PrivateClient.md) |  | [optional] 
**BusinessClient** | Pointer to [**BusinessClient**](BusinessClient.md) |  | [optional] 
**Card** | [**CardPreferences**](CardPreferences.md) |  | 
**Attachments** | Pointer to [**Attachments**](Attachments.md) |  | [optional] 

## Methods

### NewCardOrder

`func NewCardOrder(card CardPreferences, ) *CardOrder`

NewCardOrder instantiates a new CardOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOrderWithDefaults

`func NewCardOrderWithDefaults() *CardOrder`

NewCardOrderWithDefaults instantiates a new CardOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CardOrder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CardOrder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CardOrder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CardOrder) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetPrivateClient

`func (o *CardOrder) GetPrivateClient() PrivateClient`

GetPrivateClient returns the PrivateClient field if non-nil, zero value otherwise.

### GetPrivateClientOk

`func (o *CardOrder) GetPrivateClientOk() (*PrivateClient, bool)`

GetPrivateClientOk returns a tuple with the PrivateClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateClient

`func (o *CardOrder) SetPrivateClient(v PrivateClient)`

SetPrivateClient sets PrivateClient field to given value.

### HasPrivateClient

`func (o *CardOrder) HasPrivateClient() bool`

HasPrivateClient returns a boolean if a field has been set.

### GetBusinessClient

`func (o *CardOrder) GetBusinessClient() BusinessClient`

GetBusinessClient returns the BusinessClient field if non-nil, zero value otherwise.

### GetBusinessClientOk

`func (o *CardOrder) GetBusinessClientOk() (*BusinessClient, bool)`

GetBusinessClientOk returns a tuple with the BusinessClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessClient

`func (o *CardOrder) SetBusinessClient(v BusinessClient)`

SetBusinessClient sets BusinessClient field to given value.

### HasBusinessClient

`func (o *CardOrder) HasBusinessClient() bool`

HasBusinessClient returns a boolean if a field has been set.

### GetCard

`func (o *CardOrder) GetCard() CardPreferences`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CardOrder) GetCardOk() (*CardPreferences, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CardOrder) SetCard(v CardPreferences)`

SetCard sets Card field to given value.


### GetAttachments

`func (o *CardOrder) GetAttachments() Attachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CardOrder) GetAttachmentsOk() (*Attachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CardOrder) SetAttachments(v Attachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CardOrder) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


