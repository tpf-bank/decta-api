# ClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqId** | Pointer to **string** | Client uniq identification number. Deprecated, use id instead | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Type** | Pointer to **string** | Client type. Private or Business | [optional] 
**PrivateClientInfo** | Pointer to [**PrivateClient**](PrivateClient.md) |  | [optional] 
**BusinessClientInfo** | Pointer to [**BusinessClientInfo**](BusinessClientInfo.md) |  | [optional] 

## Methods

### NewClientInfo

`func NewClientInfo() *ClientInfo`

NewClientInfo instantiates a new ClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientInfoWithDefaults

`func NewClientInfoWithDefaults() *ClientInfo`

NewClientInfoWithDefaults instantiates a new ClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqId

`func (o *ClientInfo) GetUniqId() string`

GetUniqId returns the UniqId field if non-nil, zero value otherwise.

### GetUniqIdOk

`func (o *ClientInfo) GetUniqIdOk() (*string, bool)`

GetUniqIdOk returns a tuple with the UniqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqId

`func (o *ClientInfo) SetUniqId(v string)`

SetUniqId sets UniqId field to given value.

### HasUniqId

`func (o *ClientInfo) HasUniqId() bool`

HasUniqId returns a boolean if a field has been set.

### GetId

`func (o *ClientInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ClientInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrivateClientInfo

`func (o *ClientInfo) GetPrivateClientInfo() PrivateClient`

GetPrivateClientInfo returns the PrivateClientInfo field if non-nil, zero value otherwise.

### GetPrivateClientInfoOk

`func (o *ClientInfo) GetPrivateClientInfoOk() (*PrivateClient, bool)`

GetPrivateClientInfoOk returns a tuple with the PrivateClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateClientInfo

`func (o *ClientInfo) SetPrivateClientInfo(v PrivateClient)`

SetPrivateClientInfo sets PrivateClientInfo field to given value.

### HasPrivateClientInfo

`func (o *ClientInfo) HasPrivateClientInfo() bool`

HasPrivateClientInfo returns a boolean if a field has been set.

### GetBusinessClientInfo

`func (o *ClientInfo) GetBusinessClientInfo() BusinessClientInfo`

GetBusinessClientInfo returns the BusinessClientInfo field if non-nil, zero value otherwise.

### GetBusinessClientInfoOk

`func (o *ClientInfo) GetBusinessClientInfoOk() (*BusinessClientInfo, bool)`

GetBusinessClientInfoOk returns a tuple with the BusinessClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessClientInfo

`func (o *ClientInfo) SetBusinessClientInfo(v BusinessClientInfo)`

SetBusinessClientInfo sets BusinessClientInfo field to given value.

### HasBusinessClientInfo

`func (o *ClientInfo) HasBusinessClientInfo() bool`

HasBusinessClientInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


