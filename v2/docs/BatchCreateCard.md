# BatchCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **string** | Number of cards to create | 
**Template** | [**BatchCreateCardTemplateApiDto**](BatchCreateCardTemplateApiDto.md) |  | 

## Methods

### NewBatchCreateCard

`func NewBatchCreateCard(count string, template BatchCreateCardTemplateApiDto, ) *BatchCreateCard`

NewBatchCreateCard instantiates a new BatchCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCreateCardWithDefaults

`func NewBatchCreateCardWithDefaults() *BatchCreateCard`

NewBatchCreateCardWithDefaults instantiates a new BatchCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BatchCreateCard) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BatchCreateCard) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BatchCreateCard) SetCount(v string)`

SetCount sets Count field to given value.


### GetTemplate

`func (o *BatchCreateCard) GetTemplate() BatchCreateCardTemplateApiDto`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *BatchCreateCard) GetTemplateOk() (*BatchCreateCardTemplateApiDto, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *BatchCreateCard) SetTemplate(v BatchCreateCardTemplateApiDto)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


