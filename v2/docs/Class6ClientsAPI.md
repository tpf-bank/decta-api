# \Class6ClientsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeDeliveryAddress**](Class6ClientsAPI.md#ChangeDeliveryAddress) | **Patch** /v1/api/cards/{ppan}/delivery-address | Change clients delivery address
[**GetClient**](Class6ClientsAPI.md#GetClient) | **Get** /v1/api/clients/{clientId} | Search client by ID
[**GetClients**](Class6ClientsAPI.md#GetClients) | **Get** /v1/api/clients | Get client list
[**UpdateClientAddress**](Class6ClientsAPI.md#UpdateClientAddress) | **Patch** /v1/api/clients/{clientId}/address | Change client address fields
[**UpdateEmail**](Class6ClientsAPI.md#UpdateEmail) | **Put** /v2/api/clients/{clientId}/email | Change client email address
[**UpdatePassphrase**](Class6ClientsAPI.md#UpdatePassphrase) | **Put** /v1/api/cards/{ppan}/passphrase | Change card passphrase
[**UpdatePhone**](Class6ClientsAPI.md#UpdatePhone) | **Put** /v1/api/clients/{clientId}/phone | Change clients phone number



## ChangeDeliveryAddress

> ChangeDeliveryAddress(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).DeliveryAddress(deliveryAddress).Execute()

Change clients delivery address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	ppan := "123456ABCDEF5678" // string | Masked card number
	deliveryAddress := *openapiclient.NewDeliveryAddress("GBR", "London", "62 Bayswater Road") // DeliveryAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class6ClientsAPI.ChangeDeliveryAddress(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).DeliveryAddress(deliveryAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.ChangeDeliveryAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeDeliveryAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **deliveryAddress** | [**DeliveryAddress**](DeliveryAddress.md) |  | 

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

> ClientInfo GetClient(ctx, clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Search client by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	clientId := "CR0000000000000001" // string | clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class6ClientsAPI.GetClient(context.Background(), clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: ClientInfo
	fmt.Fprintf(os.Stdout, "Response from `Class6ClientsAPI.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**ClientInfo**](ClientInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> ClientInfoDataArray GetClients(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Count(count).ClientType(clientType).Name(name).Surname(surname).CompanyName(companyName).RegistrationNumber(registrationNumber).StartId(startId).Execute()

Get client list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	count := int32(10) // int32 | Number of clients (optional) (default to 10)
	clientType := "eq:PRIVATE" // string | Client type (optional)
	name := "eq:John" // string | Persons name (optional)
	surname := "eq:Doe" // string | Persons surname (optional)
	companyName := "eq:Decta" // string | Business client company name (optional)
	registrationNumber := "eq:123456789" // string | Business client company registration number (optional)
	startId := "CR0000000000001234" // string | Start the ordered list from client with selected client id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class6ClientsAPI.GetClients(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Count(count).ClientType(clientType).Name(name).Surname(surname).CompanyName(companyName).RegistrationNumber(registrationNumber).StartId(startId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: ClientInfoDataArray
	fmt.Fprintf(os.Stdout, "Response from `Class6ClientsAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **count** | **int32** | Number of clients | [default to 10]
 **clientType** | **string** | Client type | 
 **name** | **string** | Persons name | 
 **surname** | **string** | Persons surname | 
 **companyName** | **string** | Business client company name | 
 **registrationNumber** | **string** | Business client company registration number | 
 **startId** | **string** | Start the ordered list from client with selected client id | 

### Return type

[**ClientInfoDataArray**](ClientInfoDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientAddress

> UpdateClientAddress(ctx, clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientAddress(clientAddress).Execute()

Change client address fields



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	clientId := "CR0000000000000001" // string | clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols
	clientAddress := *openapiclient.NewClientAddress() // ClientAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class6ClientsAPI.UpdateClientAddress(context.Background(), clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientAddress(clientAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.UpdateClientAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **clientAddress** | [**ClientAddress**](ClientAddress.md) |  | 

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


## UpdateEmail

> UpdateEmail(ctx, clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).EmailValue(emailValue).Execute()

Change client email address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	clientId := "CR0000000000000001" // string | clientId is a unique client \"id\" number generated on Decta partner side. Maximum length 19 symbols
	emailValue := *openapiclient.NewEmailValue("john.doe@mail.com") // EmailValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class6ClientsAPI.UpdateEmail(context.Background(), clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).EmailValue(emailValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.UpdateEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client \&quot;id\&quot; number generated on Decta partner side. Maximum length 19 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **emailValue** | [**EmailValue**](EmailValue.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassphrase

> UpdatePassphrase(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PassphraseValue(passphraseValue).Execute()

Change card passphrase



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	ppan := "123456ABCDEF5678" // string | Masked card number
	passphraseValue := *openapiclient.NewPassphraseValue("123456") // PassphraseValue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class6ClientsAPI.UpdatePassphrase(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PassphraseValue(passphraseValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.UpdatePassphrase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePassphraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

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

> UpdatePhone(ctx, clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PhoneValue(phoneValue).Execute()

Change clients phone number



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/v2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	clientId := "CR0000000000000001" // string | clientId is a unique client \"id\" number generated on Decta partner side. Maximum length 19 symbols
	phoneValue := *openapiclient.NewPhoneValue("123456") // PhoneValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class6ClientsAPI.UpdatePhone(context.Background(), clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PhoneValue(phoneValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class6ClientsAPI.UpdatePhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | clientId is a unique client \&quot;id\&quot; number generated on Decta partner side. Maximum length 19 symbols | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **phoneValue** | [**PhoneValue**](PhoneValue.md) |  | 

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

