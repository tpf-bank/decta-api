# DateTimeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDateTimeFilter

`func NewDateTimeFilter() *DateTimeFilter`

NewDateTimeFilter instantiates a new DateTimeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeFilterWithDefaults

`func NewDateTimeFilterWithDefaults() *DateTimeFilter`

NewDateTimeFilterWithDefaults instantiates a new DateTimeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *DateTimeFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *DateTimeFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *DateTimeFilter) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *DateTimeFilter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOperation

`func (o *DateTimeFilter) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *DateTimeFilter) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *DateTimeFilter) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *DateTimeFilter) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *DateTimeFilter) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DateTimeFilter) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DateTimeFilter) SetValue(v time.Time)`

SetValue sets Value field to given value.

### HasValue

`func (o *DateTimeFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


