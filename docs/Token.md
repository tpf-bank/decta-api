# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ppan** | Pointer to **string** | Card pseudo-pan | [optional] 
**TokenNumber** | Pointer to **string** | Token pan | [optional] 
**TokenExpiry** | Pointer to **string** | Expiry date (YYMM) of token | [optional] 
**TokenStatus** | Pointer to **string** | Token status | [optional] 
**TokenType** | Pointer to **string** | Token type | [optional] 
**TokenServiceProviderId** | Pointer to **string** | Token service provider identifier | [optional] 
**WalletId** | Pointer to **string** | Wallet service provider identifier | [optional] 
**DeviceId** | Pointer to **string** | Device ID | [optional] 
**DeviceName** | Pointer to **string** | Device name | [optional] 
**PanReferenceId** | Pointer to **string** | PAN reference ID | [optional] 
**Created** | Pointer to **string** | Wallet assignment to card and date | [optional] 
**Changed** | Pointer to **string** | Date and time of last changes | [optional] 
**PreviousStatus** | Pointer to **string** | Previous status of token | [optional] 
**MessageTraceId** | Pointer to **string** | Unique ID for messages of token | [optional] 
**TokenRefId** | Pointer to **string** | Token reference ID | [optional] 
**RequestorId** | Pointer to **string** | Token requestor ID | [optional] 
**StatusChanged** | Pointer to **string** | Date and time of token status last changes | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpan

`func (o *Token) GetPpan() string`

GetPpan returns the Ppan field if non-nil, zero value otherwise.

### GetPpanOk

`func (o *Token) GetPpanOk() (*string, bool)`

GetPpanOk returns a tuple with the Ppan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpan

`func (o *Token) SetPpan(v string)`

SetPpan sets Ppan field to given value.

### HasPpan

`func (o *Token) HasPpan() bool`

HasPpan returns a boolean if a field has been set.

### GetTokenNumber

`func (o *Token) GetTokenNumber() string`

GetTokenNumber returns the TokenNumber field if non-nil, zero value otherwise.

### GetTokenNumberOk

`func (o *Token) GetTokenNumberOk() (*string, bool)`

GetTokenNumberOk returns a tuple with the TokenNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumber

`func (o *Token) SetTokenNumber(v string)`

SetTokenNumber sets TokenNumber field to given value.

### HasTokenNumber

`func (o *Token) HasTokenNumber() bool`

HasTokenNumber returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *Token) GetTokenExpiry() string`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *Token) GetTokenExpiryOk() (*string, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *Token) SetTokenExpiry(v string)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *Token) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.

### GetTokenStatus

`func (o *Token) GetTokenStatus() string`

GetTokenStatus returns the TokenStatus field if non-nil, zero value otherwise.

### GetTokenStatusOk

`func (o *Token) GetTokenStatusOk() (*string, bool)`

GetTokenStatusOk returns a tuple with the TokenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStatus

`func (o *Token) SetTokenStatus(v string)`

SetTokenStatus sets TokenStatus field to given value.

### HasTokenStatus

`func (o *Token) HasTokenStatus() bool`

HasTokenStatus returns a boolean if a field has been set.

### GetTokenType

`func (o *Token) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Token) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Token) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Token) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTokenServiceProviderId

`func (o *Token) GetTokenServiceProviderId() string`

GetTokenServiceProviderId returns the TokenServiceProviderId field if non-nil, zero value otherwise.

### GetTokenServiceProviderIdOk

`func (o *Token) GetTokenServiceProviderIdOk() (*string, bool)`

GetTokenServiceProviderIdOk returns a tuple with the TokenServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenServiceProviderId

`func (o *Token) SetTokenServiceProviderId(v string)`

SetTokenServiceProviderId sets TokenServiceProviderId field to given value.

### HasTokenServiceProviderId

`func (o *Token) HasTokenServiceProviderId() bool`

HasTokenServiceProviderId returns a boolean if a field has been set.

### GetWalletId

`func (o *Token) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Token) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Token) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *Token) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetDeviceId

`func (o *Token) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Token) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Token) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Token) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *Token) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *Token) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *Token) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *Token) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetPanReferenceId

`func (o *Token) GetPanReferenceId() string`

GetPanReferenceId returns the PanReferenceId field if non-nil, zero value otherwise.

### GetPanReferenceIdOk

`func (o *Token) GetPanReferenceIdOk() (*string, bool)`

GetPanReferenceIdOk returns a tuple with the PanReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanReferenceId

`func (o *Token) SetPanReferenceId(v string)`

SetPanReferenceId sets PanReferenceId field to given value.

### HasPanReferenceId

`func (o *Token) HasPanReferenceId() bool`

HasPanReferenceId returns a boolean if a field has been set.

### GetCreated

`func (o *Token) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Token) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Token) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Token) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetChanged

`func (o *Token) GetChanged() string`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *Token) GetChangedOk() (*string, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *Token) SetChanged(v string)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *Token) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetPreviousStatus

`func (o *Token) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *Token) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *Token) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.

### HasPreviousStatus

`func (o *Token) HasPreviousStatus() bool`

HasPreviousStatus returns a boolean if a field has been set.

### GetMessageTraceId

`func (o *Token) GetMessageTraceId() string`

GetMessageTraceId returns the MessageTraceId field if non-nil, zero value otherwise.

### GetMessageTraceIdOk

`func (o *Token) GetMessageTraceIdOk() (*string, bool)`

GetMessageTraceIdOk returns a tuple with the MessageTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTraceId

`func (o *Token) SetMessageTraceId(v string)`

SetMessageTraceId sets MessageTraceId field to given value.

### HasMessageTraceId

`func (o *Token) HasMessageTraceId() bool`

HasMessageTraceId returns a boolean if a field has been set.

### GetTokenRefId

`func (o *Token) GetTokenRefId() string`

GetTokenRefId returns the TokenRefId field if non-nil, zero value otherwise.

### GetTokenRefIdOk

`func (o *Token) GetTokenRefIdOk() (*string, bool)`

GetTokenRefIdOk returns a tuple with the TokenRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRefId

`func (o *Token) SetTokenRefId(v string)`

SetTokenRefId sets TokenRefId field to given value.

### HasTokenRefId

`func (o *Token) HasTokenRefId() bool`

HasTokenRefId returns a boolean if a field has been set.

### GetRequestorId

`func (o *Token) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *Token) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *Token) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *Token) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### GetStatusChanged

`func (o *Token) GetStatusChanged() string`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *Token) GetStatusChangedOk() (*string, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *Token) SetStatusChanged(v string)`

SetStatusChanged sets StatusChanged field to given value.

### HasStatusChanged

`func (o *Token) HasStatusChanged() bool`

HasStatusChanged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


