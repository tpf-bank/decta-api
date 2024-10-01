# \TokenAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTokenStatus**](TokenAPI.md#ChangeTokenStatus) | **Patch** /v1/api/tokens/{tokenNumber}/state | Change tokens status
[**GetApplePayReferenceData**](TokenAPI.md#GetApplePayReferenceData) | **Get** /v1/api/tokens/{ppan}/applepay-reference-data | Get ApplePay reference data
[**GetGooglePayReferenceData**](TokenAPI.md#GetGooglePayReferenceData) | **Get** /v1/api/tokens/{ppan}/googlepay-reference-data | Get GooglePay reference data
[**GetPanReferenceData**](TokenAPI.md#GetPanReferenceData) | **Get** /v1/api/tokens/{ppan}/pan-reference-data | 
[**GetTokenReferenceData**](TokenAPI.md#GetTokenReferenceData) | **Get** /v1/api/tokens/{ppan}/token-reference-data | 
[**GetTokens**](TokenAPI.md#GetTokens) | **Get** /v1/api/tokens/{ppan} | Get card token



## ChangeTokenStatus

> ChangeTokenStatus(ctx, tokenNumber).TokenHeader(tokenHeader).TokenSignature(tokenSignature).TokenStatusChange(tokenStatusChange).RequestId(requestId).Execute()

Change tokens status



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
	tokenNumber := "tokenNumber_example" // string | masked card number
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	tokenStatusChange := *openapiclient.NewTokenStatusChange("TokenStatus_example", "Details_example") // TokenStatusChange | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TokenAPI.ChangeTokenStatus(context.Background(), tokenNumber).TokenHeader(tokenHeader).TokenSignature(tokenSignature).TokenStatusChange(tokenStatusChange).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.ChangeTokenStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenNumber** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTokenStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **tokenStatusChange** | [**TokenStatusChange**](TokenStatusChange.md) |  | 
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


## GetApplePayReferenceData

> PanReferenceData GetApplePayReferenceData(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get ApplePay reference data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetApplePayReferenceData(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetApplePayReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplePayReferenceData`: PanReferenceData
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetApplePayReferenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplePayReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**PanReferenceData**](PanReferenceData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGooglePayReferenceData

> Data GetGooglePayReferenceData(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get GooglePay reference data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetGooglePayReferenceData(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetGooglePayReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGooglePayReferenceData`: Data
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetGooglePayReferenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGooglePayReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**Data**](Data.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPanReferenceData

> PanReferenceData GetPanReferenceData(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetPanReferenceData(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetPanReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPanReferenceData`: PanReferenceData
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetPanReferenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPanReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**PanReferenceData**](PanReferenceData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenReferenceData

> Data GetTokenReferenceData(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetTokenReferenceData(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTokenReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenReferenceData`: Data
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTokenReferenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**Data**](Data.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> TokenInfo GetTokens(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get card token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetTokens(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: TokenInfo
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**TokenInfo**](TokenInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

