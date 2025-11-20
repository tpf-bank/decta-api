# Limit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Limit identifier | 
**Desc** | Pointer to **string** | Limit identifier description | [optional] 
**Key** | **string** | Limit additional key | 
**Individual** | Pointer to [**LimitValues**](LimitValues.md) |  | [optional] 
**Group** | Pointer to [**LimitValues**](LimitValues.md) |  | [optional] 

## Methods

### NewLimit

`func NewLimit(id string, key string, ) *Limit`

NewLimit instantiates a new Limit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitWithDefaults

`func NewLimitWithDefaults() *Limit`

NewLimitWithDefaults instantiates a new Limit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Limit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Limit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Limit) SetId(v string)`

SetId sets Id field to given value.


### GetDesc

`func (o *Limit) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Limit) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Limit) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Limit) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetKey

`func (o *Limit) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Limit) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Limit) SetKey(v string)`

SetKey sets Key field to given value.


### GetIndividual

`func (o *Limit) GetIndividual() LimitValues`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *Limit) GetIndividualOk() (*LimitValues, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *Limit) SetIndividual(v LimitValues)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *Limit) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetGroup

`func (o *Limit) GetGroup() LimitValues`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Limit) GetGroupOk() (*LimitValues, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Limit) SetGroup(v LimitValues)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Limit) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


