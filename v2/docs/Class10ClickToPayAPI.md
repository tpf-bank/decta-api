# \Class10ClickToPayAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClickToPay**](Class10ClickToPayAPI.md#AddClickToPay) | **Post** /v2/api/click-to-pay/{ppan} | Add card to Click to Pay
[**ChangeClickToPay**](Class10ClickToPayAPI.md#ChangeClickToPay) | **Patch** /v2/api/click-to-pay/{ppan} | Change Click to Pay information using card
[**DeleteClickToPay**](Class10ClickToPayAPI.md#DeleteClickToPay) | **Delete** /v2/api/click-to-pay/{ppan} | Delete data from Click to Pay using card
[**GetClickToPay**](Class10ClickToPayAPI.md#GetClickToPay) | **Get** /v2/api/click-to-pay/{clientId} | Get Click to Pay information using client id



## AddClickToPay

> AddClickToPay(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClickToPayCrudApiDto(clickToPayCrudApiDto).Execute()

Add card to Click to Pay



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
	clickToPayCrudApiDto := *openapiclient.NewClickToPayCrudApiDto() // ClickToPayCrudApiDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class10ClickToPayAPI.AddClickToPay(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClickToPayCrudApiDto(clickToPayCrudApiDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class10ClickToPayAPI.AddClickToPay``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddClickToPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **clickToPayCrudApiDto** | [**ClickToPayCrudApiDto**](ClickToPayCrudApiDto.md) |  | 

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


## ChangeClickToPay

> ChangeClickToPay(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClickToPayCrudApiDto(clickToPayCrudApiDto).Execute()

Change Click to Pay information using card



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
	clickToPayCrudApiDto := *openapiclient.NewClickToPayCrudApiDto() // ClickToPayCrudApiDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class10ClickToPayAPI.ChangeClickToPay(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClickToPayCrudApiDto(clickToPayCrudApiDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class10ClickToPayAPI.ChangeClickToPay``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChangeClickToPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **clickToPayCrudApiDto** | [**ClickToPayCrudApiDto**](ClickToPayCrudApiDto.md) |  | 

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


## DeleteClickToPay

> DeleteClickToPay(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientId(clientId).Execute()

Delete data from Click to Pay using card



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
	clientId := "CR0000000000000001" // string | clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class10ClickToPayAPI.DeleteClickToPay(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class10ClickToPayAPI.DeleteClickToPay``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClickToPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **clientId** | **string** | clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClickToPay

> ClickToPayGetData GetClickToPay(ctx, clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Get Click to Pay information using client id



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
	clientId := "CR0000000000000001" // string | Click to Pay profile id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class10ClickToPayAPI.GetClickToPay(context.Background(), clientId).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class10ClickToPayAPI.GetClickToPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClickToPay`: ClickToPayGetData
	fmt.Fprintf(os.Stdout, "Response from `Class10ClickToPayAPI.GetClickToPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | Click to Pay profile id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClickToPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**ClickToPayGetData**](ClickToPayGetData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

