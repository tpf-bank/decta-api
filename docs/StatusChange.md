# StatusChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardState** | **string** | Desirable card state | 
**Details** | **string** | For &lt;em&gt;Block/Unblock&lt;/em&gt; operation status change reason. For card &lt;em&gt;Activation&lt;/em&gt; must contain last 6 digits of card pan. | 

## Methods

### NewStatusChange

`func NewStatusChange(cardState string, details string, ) *StatusChange`

NewStatusChange instantiates a new StatusChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusChangeWithDefaults

`func NewStatusChangeWithDefaults() *StatusChange`

NewStatusChangeWithDefaults instantiates a new StatusChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardState

`func (o *StatusChange) GetCardState() string`

GetCardState returns the CardState field if non-nil, zero value otherwise.

### GetCardStateOk

`func (o *StatusChange) GetCardStateOk() (*string, bool)`

GetCardStateOk returns a tuple with the CardState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardState

`func (o *StatusChange) SetCardState(v string)`

SetCardState sets CardState field to given value.


### GetDetails

`func (o *StatusChange) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *StatusChange) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *StatusChange) SetDetails(v string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


