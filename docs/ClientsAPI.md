# \ClientsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeDeliveryAddress**](ClientsAPI.md#ChangeDeliveryAddress) | **Patch** /v1/api/cards/{ppan}/delivery-address | Change clients delivery address
[**GetClient**](ClientsAPI.md#GetClient) | **Get** /v1/api/clients/{clientId} | Search client by ID
[**GetClients**](ClientsAPI.md#GetClients) | **Get** /v1/api/clients | Get client list
[**UpdateClientAddress**](ClientsAPI.md#UpdateClientAddress) | **Patch** /v1/api/clients/{clientId}/address | Change client address fields
[**UpdatePassphrase**](ClientsAPI.md#UpdatePassphrase) | **Put** /v1/api/cards/{ppan}/passphrase | Change card passphrase
[**UpdatePhone**](ClientsAPI.md#UpdatePhone) | **Put** /v1/api/clients/{clientId}/phone | Change clients phone number



## ChangeDeliveryAddress

> ChangeDeliveryAddress(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).DeliveryAddress(deliveryAddress).RequestId(requestId).Execute()

Change clients delivery address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ppan := "ppan_example" // string | masked card number
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	deliveryAddress := *openapiclient.NewDeliveryAddress("GBR", "London", "62 Bayswater Road") // DeliveryAddress | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.ChangeDeliveryAddress(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).DeliveryAddress(deliveryAddress).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.ChangeDeliveryAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeDeliveryAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **deliveryAddress** | [**DeliveryAddress**](DeliveryAddress.md) |  | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClient

> ClientInfo GetClient(ctx, clientId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Search client by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientId := "CR0000000000000001" // string | clientId is a unique client id number generated on the DECTA partner side. Length 18 symbols
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetClient(context.Background(), clientId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: ClientInfo
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client id number generated on the DECTA partner side. Length 18 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**ClientInfo**](ClientInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> ClientInfoDataArray GetClients(ctx).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Count(count).ClientType(clientType).Name(name).Surname(surname).CompanyName(companyName).RegistrationNumber(registrationNumber).StartId(startId).RequestId(requestId).Execute()

Get client list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	count := int32(10) // int32 | Number of clients (optional) (default to 10)
	clientType := "eq:PRIVATE" // string | Client type (optional)
	name := "eq:John" // string | Persons name (optional)
	surname := "eq:Doe" // string | Persons surname (optional)
	companyName := "eq:Decta" // string | Business client company name (optional)
	registrationNumber := "eq:123456789" // string | Business client company registration number (optional)
	startId := "CR0000000000001234" // string | Start the ordered list from client with selected client id (optional)
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetClients(context.Background()).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Count(count).ClientType(clientType).Name(name).Surname(surname).CompanyName(companyName).RegistrationNumber(registrationNumber).StartId(startId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: ClientInfoDataArray
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **count** | **int32** | Number of clients | [default to 10]
 **clientType** | **string** | Client type | 
 **name** | **string** | Persons name | 
 **surname** | **string** | Persons surname | 
 **companyName** | **string** | Business client company name | 
 **registrationNumber** | **string** | Business client company registration number | 
 **startId** | **string** | Start the ordered list from client with selected client id | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**ClientInfoDataArray**](ClientInfoDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientAddress

> UpdateClientAddress(ctx, clientId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientAddress(clientAddress).RequestId(requestId).Execute()

Change client address fields



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientId := "CR0000000000000001" // string | clientId is a unique client id number generated on the DECTA partner side. Length 18 symbols
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	clientAddress := *openapiclient.NewClientAddress() // ClientAddress | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.UpdateClientAddress(context.Background(), clientId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientAddress(clientAddress).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdateClientAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client id number generated on the DECTA partner side. Length 18 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **clientAddress** | [**ClientAddress**](ClientAddress.md) |  | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassphrase

> UpdatePassphrase(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).PassphraseValue(passphraseValue).Execute()

Change card passphrase



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ppan := "ppan_example" // string | masked card number
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)
	passphraseValue := *openapiclient.NewPassphraseValue("123456") // PassphraseValue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.UpdatePassphrase(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).PassphraseValue(passphraseValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdatePassphrase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePassphraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 
 **passphraseValue** | [**PassphraseValue**](PassphraseValue.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhone

> UpdatePhone(ctx, clientId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PhoneValue(phoneValue).RequestId(requestId).Execute()

Change clients phone number



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientId := "CR0000000000000001" // string | clientId is a unique client \"id\" number generated on Decta partner side. Length 18 symbols
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	phoneValue := *openapiclient.NewPhoneValue("123456") // PhoneValue | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.UpdatePhone(context.Background(), clientId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PhoneValue(phoneValue).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdatePhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client \&quot;id\&quot; number generated on Decta partner side. Length 18 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **phoneValue** | [**PhoneValue**](PhoneValue.md) |  | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

