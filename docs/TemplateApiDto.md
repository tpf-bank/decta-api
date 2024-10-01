# TemplateApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateClient** | [**PrivateClientGiftCardApiDto**](PrivateClientGiftCardApiDto.md) |  | 
**Card** | [**CardPreferencesGiftCardApiDto**](CardPreferencesGiftCardApiDto.md) |  | 

## Methods

### NewTemplateApiDto

`func NewTemplateApiDto(privateClient PrivateClientGiftCardApiDto, card CardPreferencesGiftCardApiDto, ) *TemplateApiDto`

NewTemplateApiDto instantiates a new TemplateApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateApiDtoWithDefaults

`func NewTemplateApiDtoWithDefaults() *TemplateApiDto`

NewTemplateApiDtoWithDefaults instantiates a new TemplateApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateClient

`func (o *TemplateApiDto) GetPrivateClient() PrivateClientGiftCardApiDto`

GetPrivateClient returns the PrivateClient field if non-nil, zero value otherwise.

### GetPrivateClientOk

`func (o *TemplateApiDto) GetPrivateClientOk() (*PrivateClientGiftCardApiDto, bool)`

GetPrivateClientOk returns a tuple with the PrivateClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateClient

`func (o *TemplateApiDto) SetPrivateClient(v PrivateClientGiftCardApiDto)`

SetPrivateClient sets PrivateClient field to given value.


### GetCard

`func (o *TemplateApiDto) GetCard() CardPreferencesGiftCardApiDto`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *TemplateApiDto) GetCardOk() (*CardPreferencesGiftCardApiDto, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *TemplateApiDto) SetCard(v CardPreferencesGiftCardApiDto)`

SetCard sets Card field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


