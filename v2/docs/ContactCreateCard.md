# ContactCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Person name &lt;i style&#x3D;\&quot;color:grey\&quot;&gt;[name length + surname length - max 24 symbols]&lt;i/&gt; | [optional] 
**Surname** | Pointer to **string** | Person surname &lt;i style&#x3D;\&quot;color:grey\&quot;&gt;[name length + surname length - max 24 symbols]&lt;i/&gt; | [optional] 
**MobilePhone** | Pointer to **string** | Person Mobile phone number | [optional] 
**Email** | Pointer to **string** | Person E-mail address | [optional] 

## Methods

### NewContactCreateCard

`func NewContactCreateCard() *ContactCreateCard`

NewContactCreateCard instantiates a new ContactCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreateCardWithDefaults

`func NewContactCreateCardWithDefaults() *ContactCreateCard`

NewContactCreateCardWithDefaults instantiates a new ContactCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContactCreateCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactCreateCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactCreateCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactCreateCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *ContactCreateCard) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ContactCreateCard) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ContactCreateCard) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *ContactCreateCard) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *ContactCreateCard) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ContactCreateCard) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ContactCreateCard) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ContactCreateCard) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *ContactCreateCard) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactCreateCard) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactCreateCard) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactCreateCard) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


