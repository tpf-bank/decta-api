# Attachments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | [**[]File**](File.md) | Scan of ID document | 
**Bills** | Pointer to [**[]File**](File.md) | Utility bills | [optional] 
**AmlChecks** | Pointer to [**[]File**](File.md) | Document with sanction screening results | [optional] 
**CardAgreements** | Pointer to [**[]File**](File.md) | Cardholder agreements | [optional] 
**EmployeeContracts** | Pointer to [**[]File**](File.md) | Employee contract | [optional] 
**WorksPermits** | Pointer to [**[]File**](File.md) | EEA work permit for NON-EEA cardholders | [optional] 
**ClientsFunds** | Pointer to [**[]File**](File.md) | Document with a proof of funds and a source of funds | [optional] 

## Methods

### NewAttachments

`func NewAttachments(documents []File, ) *Attachments`

NewAttachments instantiates a new Attachments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsWithDefaults

`func NewAttachmentsWithDefaults() *Attachments`

NewAttachmentsWithDefaults instantiates a new Attachments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *Attachments) GetDocuments() []File`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Attachments) GetDocumentsOk() (*[]File, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Attachments) SetDocuments(v []File)`

SetDocuments sets Documents field to given value.


### GetBills

`func (o *Attachments) GetBills() []File`

GetBills returns the Bills field if non-nil, zero value otherwise.

### GetBillsOk

`func (o *Attachments) GetBillsOk() (*[]File, bool)`

GetBillsOk returns a tuple with the Bills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBills

`func (o *Attachments) SetBills(v []File)`

SetBills sets Bills field to given value.

### HasBills

`func (o *Attachments) HasBills() bool`

HasBills returns a boolean if a field has been set.

### GetAmlChecks

`func (o *Attachments) GetAmlChecks() []File`

GetAmlChecks returns the AmlChecks field if non-nil, zero value otherwise.

### GetAmlChecksOk

`func (o *Attachments) GetAmlChecksOk() (*[]File, bool)`

GetAmlChecksOk returns a tuple with the AmlChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmlChecks

`func (o *Attachments) SetAmlChecks(v []File)`

SetAmlChecks sets AmlChecks field to given value.

### HasAmlChecks

`func (o *Attachments) HasAmlChecks() bool`

HasAmlChecks returns a boolean if a field has been set.

### GetCardAgreements

`func (o *Attachments) GetCardAgreements() []File`

GetCardAgreements returns the CardAgreements field if non-nil, zero value otherwise.

### GetCardAgreementsOk

`func (o *Attachments) GetCardAgreementsOk() (*[]File, bool)`

GetCardAgreementsOk returns a tuple with the CardAgreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAgreements

`func (o *Attachments) SetCardAgreements(v []File)`

SetCardAgreements sets CardAgreements field to given value.

### HasCardAgreements

`func (o *Attachments) HasCardAgreements() bool`

HasCardAgreements returns a boolean if a field has been set.

### GetEmployeeContracts

`func (o *Attachments) GetEmployeeContracts() []File`

GetEmployeeContracts returns the EmployeeContracts field if non-nil, zero value otherwise.

### GetEmployeeContractsOk

`func (o *Attachments) GetEmployeeContractsOk() (*[]File, bool)`

GetEmployeeContractsOk returns a tuple with the EmployeeContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeContracts

`func (o *Attachments) SetEmployeeContracts(v []File)`

SetEmployeeContracts sets EmployeeContracts field to given value.

### HasEmployeeContracts

`func (o *Attachments) HasEmployeeContracts() bool`

HasEmployeeContracts returns a boolean if a field has been set.

### GetWorksPermits

`func (o *Attachments) GetWorksPermits() []File`

GetWorksPermits returns the WorksPermits field if non-nil, zero value otherwise.

### GetWorksPermitsOk

`func (o *Attachments) GetWorksPermitsOk() (*[]File, bool)`

GetWorksPermitsOk returns a tuple with the WorksPermits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksPermits

`func (o *Attachments) SetWorksPermits(v []File)`

SetWorksPermits sets WorksPermits field to given value.

### HasWorksPermits

`func (o *Attachments) HasWorksPermits() bool`

HasWorksPermits returns a boolean if a field has been set.

### GetClientsFunds

`func (o *Attachments) GetClientsFunds() []File`

GetClientsFunds returns the ClientsFunds field if non-nil, zero value otherwise.

### GetClientsFundsOk

`func (o *Attachments) GetClientsFundsOk() (*[]File, bool)`

GetClientsFundsOk returns a tuple with the ClientsFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsFunds

`func (o *Attachments) SetClientsFunds(v []File)`

SetClientsFunds sets ClientsFunds field to given value.

### HasClientsFunds

`func (o *Attachments) HasClientsFunds() bool`

HasClientsFunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


