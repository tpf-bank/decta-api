# GiftCardOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | Decta API user controlled field | [optional] 
**Count** | **string** | Number of gift cards to create&lt;br&gt;min: 1&lt;br&gt;max: 10000 | 
**Template** | [**TemplateApiDto**](TemplateApiDto.md) |  | 

## Methods

### NewGiftCardOrder

`func NewGiftCardOrder(count string, template TemplateApiDto, ) *GiftCardOrder`

NewGiftCardOrder instantiates a new GiftCardOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardOrderWithDefaults

`func NewGiftCardOrderWithDefaults() *GiftCardOrder`

NewGiftCardOrderWithDefaults instantiates a new GiftCardOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *GiftCardOrder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GiftCardOrder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GiftCardOrder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GiftCardOrder) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCount

`func (o *GiftCardOrder) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GiftCardOrder) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GiftCardOrder) SetCount(v string)`

SetCount sets Count field to given value.


### GetTemplate

`func (o *GiftCardOrder) GetTemplate() TemplateApiDto`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GiftCardOrder) GetTemplateOk() (*TemplateApiDto, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GiftCardOrder) SetTemplate(v TemplateApiDto)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


