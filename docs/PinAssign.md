# PinAssign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | **string** | Card expiry in format [MMYY] | 
**EncryptedPINBlock** | **string** | PIN Block encrypted under the public key certificate. Pin length should be equal to 4 | 
**EncryptionCertificateThumbprint** | **string** | Public key certificate used for encryption thumbprint | 

## Methods

### NewPinAssign

`func NewPinAssign(expiry string, encryptedPINBlock string, encryptionCertificateThumbprint string, ) *PinAssign`

NewPinAssign instantiates a new PinAssign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinAssignWithDefaults

`func NewPinAssignWithDefaults() *PinAssign`

NewPinAssignWithDefaults instantiates a new PinAssign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *PinAssign) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PinAssign) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PinAssign) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### GetEncryptedPINBlock

`func (o *PinAssign) GetEncryptedPINBlock() string`

GetEncryptedPINBlock returns the EncryptedPINBlock field if non-nil, zero value otherwise.

### GetEncryptedPINBlockOk

`func (o *PinAssign) GetEncryptedPINBlockOk() (*string, bool)`

GetEncryptedPINBlockOk returns a tuple with the EncryptedPINBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPINBlock

`func (o *PinAssign) SetEncryptedPINBlock(v string)`

SetEncryptedPINBlock sets EncryptedPINBlock field to given value.


### GetEncryptionCertificateThumbprint

`func (o *PinAssign) GetEncryptionCertificateThumbprint() string`

GetEncryptionCertificateThumbprint returns the EncryptionCertificateThumbprint field if non-nil, zero value otherwise.

### GetEncryptionCertificateThumbprintOk

`func (o *PinAssign) GetEncryptionCertificateThumbprintOk() (*string, bool)`

GetEncryptionCertificateThumbprintOk returns a tuple with the EncryptionCertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCertificateThumbprint

`func (o *PinAssign) SetEncryptionCertificateThumbprint(v string)`

SetEncryptionCertificateThumbprint sets EncryptionCertificateThumbprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


