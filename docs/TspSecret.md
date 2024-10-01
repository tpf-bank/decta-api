# TspSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**Secret** | Pointer to **string** | Secret | [optional] 

## Methods

### NewTspSecret

`func NewTspSecret() *TspSecret`

NewTspSecret instantiates a new TspSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTspSecretWithDefaults

`func NewTspSecretWithDefaults() *TspSecret`

NewTspSecretWithDefaults instantiates a new TspSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *TspSecret) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *TspSecret) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *TspSecret) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *TspSecret) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetSecret

`func (o *TspSecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TspSecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TspSecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TspSecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


