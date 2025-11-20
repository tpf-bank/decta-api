# PinTryReset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **string** | Card expiry date | [optional] 
**Details** | Pointer to **string** | Reason for resetting a card PIN counter | [optional] 

## Methods

### NewPinTryReset

`func NewPinTryReset() *PinTryReset`

NewPinTryReset instantiates a new PinTryReset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinTryResetWithDefaults

`func NewPinTryResetWithDefaults() *PinTryReset`

NewPinTryResetWithDefaults instantiates a new PinTryReset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *PinTryReset) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PinTryReset) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PinTryReset) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PinTryReset) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetDetails

`func (o *PinTryReset) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PinTryReset) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PinTryReset) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PinTryReset) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


