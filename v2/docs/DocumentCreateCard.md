# DocumentCreateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of ID document of a person | [optional] 
**Subtype** | Pointer to **string** | Subtype of ID document of a person. | [optional] 
**Number** | Pointer to **string** | Number of ID document | [optional] 
**IssuingDate** | Pointer to **string** | Date(YYYY-MM-DD) of issue of person ID document | [optional] 
**BirthDate** | Pointer to **string** | Date(YYYY-MM-DD) of birth of a person | [optional] 

## Methods

### NewDocumentCreateCard

`func NewDocumentCreateCard() *DocumentCreateCard`

NewDocumentCreateCard instantiates a new DocumentCreateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentCreateCardWithDefaults

`func NewDocumentCreateCardWithDefaults() *DocumentCreateCard`

NewDocumentCreateCardWithDefaults instantiates a new DocumentCreateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DocumentCreateCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentCreateCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentCreateCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentCreateCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *DocumentCreateCard) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DocumentCreateCard) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DocumentCreateCard) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *DocumentCreateCard) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetNumber

`func (o *DocumentCreateCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DocumentCreateCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DocumentCreateCard) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DocumentCreateCard) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetIssuingDate

`func (o *DocumentCreateCard) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *DocumentCreateCard) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *DocumentCreateCard) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *DocumentCreateCard) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *DocumentCreateCard) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *DocumentCreateCard) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *DocumentCreateCard) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *DocumentCreateCard) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


