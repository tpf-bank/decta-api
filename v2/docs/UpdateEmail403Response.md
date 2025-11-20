# UpdateEmail403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ApiErrorV2**](ApiErrorV2.md) | Errors list | [optional] 

## Methods

### NewUpdateEmail403Response

`func NewUpdateEmail403Response() *UpdateEmail403Response`

NewUpdateEmail403Response instantiates a new UpdateEmail403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmail403ResponseWithDefaults

`func NewUpdateEmail403ResponseWithDefaults() *UpdateEmail403Response`

NewUpdateEmail403ResponseWithDefaults instantiates a new UpdateEmail403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *UpdateEmail403Response) GetErrors() []ApiErrorV2`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateEmail403Response) GetErrorsOk() (*[]ApiErrorV2, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateEmail403Response) SetErrors(v []ApiErrorV2)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UpdateEmail403Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


