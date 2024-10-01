# TokenStatusChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenStatus** | **string** | Desirable token status | 
**Details** | **string** | Reason for token status change | 

## Methods

### NewTokenStatusChange

`func NewTokenStatusChange(tokenStatus string, details string, ) *TokenStatusChange`

NewTokenStatusChange instantiates a new TokenStatusChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenStatusChangeWithDefaults

`func NewTokenStatusChangeWithDefaults() *TokenStatusChange`

NewTokenStatusChangeWithDefaults instantiates a new TokenStatusChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenStatus

`func (o *TokenStatusChange) GetTokenStatus() string`

GetTokenStatus returns the TokenStatus field if non-nil, zero value otherwise.

### GetTokenStatusOk

`func (o *TokenStatusChange) GetTokenStatusOk() (*string, bool)`

GetTokenStatusOk returns a tuple with the TokenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStatus

`func (o *TokenStatusChange) SetTokenStatus(v string)`

SetTokenStatus sets TokenStatus field to given value.


### GetDetails

`func (o *TokenStatusChange) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TokenStatusChange) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TokenStatusChange) SetDetails(v string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


