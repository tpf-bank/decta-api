# ApiErrorV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Field with error | [optional] 
**Code** | Pointer to **string** | Error code | [optional] 
**Message** | Pointer to **string** | Description message | [optional] 
**Details** | Pointer to **string** | Details | [optional] 
**Link** | Pointer to **string** | Link | [optional] 

## Methods

### NewApiErrorV2

`func NewApiErrorV2() *ApiErrorV2`

NewApiErrorV2 instantiates a new ApiErrorV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorV2WithDefaults

`func NewApiErrorV2WithDefaults() *ApiErrorV2`

NewApiErrorV2WithDefaults instantiates a new ApiErrorV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ApiErrorV2) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ApiErrorV2) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ApiErrorV2) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ApiErrorV2) HasField() bool`

HasField returns a boolean if a field has been set.

### GetCode

`func (o *ApiErrorV2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorV2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorV2) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiErrorV2) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ApiErrorV2) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiErrorV2) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiErrorV2) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiErrorV2) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *ApiErrorV2) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApiErrorV2) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApiErrorV2) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApiErrorV2) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLink

`func (o *ApiErrorV2) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ApiErrorV2) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ApiErrorV2) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ApiErrorV2) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


