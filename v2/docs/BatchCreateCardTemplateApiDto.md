# BatchCreateCardTemplateApiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateClient** | [**BatchCardPrivateClientApiDto**](BatchCardPrivateClientApiDto.md) |  | 
**Card** | [**BatchCardPreferencesApiDto**](BatchCardPreferencesApiDto.md) |  | 

## Methods

### NewBatchCreateCardTemplateApiDto

`func NewBatchCreateCardTemplateApiDto(privateClient BatchCardPrivateClientApiDto, card BatchCardPreferencesApiDto, ) *BatchCreateCardTemplateApiDto`

NewBatchCreateCardTemplateApiDto instantiates a new BatchCreateCardTemplateApiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCreateCardTemplateApiDtoWithDefaults

`func NewBatchCreateCardTemplateApiDtoWithDefaults() *BatchCreateCardTemplateApiDto`

NewBatchCreateCardTemplateApiDtoWithDefaults instantiates a new BatchCreateCardTemplateApiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateClient

`func (o *BatchCreateCardTemplateApiDto) GetPrivateClient() BatchCardPrivateClientApiDto`

GetPrivateClient returns the PrivateClient field if non-nil, zero value otherwise.

### GetPrivateClientOk

`func (o *BatchCreateCardTemplateApiDto) GetPrivateClientOk() (*BatchCardPrivateClientApiDto, bool)`

GetPrivateClientOk returns a tuple with the PrivateClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateClient

`func (o *BatchCreateCardTemplateApiDto) SetPrivateClient(v BatchCardPrivateClientApiDto)`

SetPrivateClient sets PrivateClient field to given value.


### GetCard

`func (o *BatchCreateCardTemplateApiDto) GetCard() BatchCardPreferencesApiDto`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *BatchCreateCardTemplateApiDto) GetCardOk() (*BatchCardPreferencesApiDto, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *BatchCreateCardTemplateApiDto) SetCard(v BatchCardPreferencesApiDto)`

SetCard sets Card field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


