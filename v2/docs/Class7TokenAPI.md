# \Class7TokenAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTokenStatus**](Class7TokenAPI.md#ChangeTokenStatus) | **Patch** /v1/api/tokens/{tokenNumber}/state | Change tokens status
[**GetApplePayReferenceData**](Class7TokenAPI.md#GetApplePayReferenceData) | **Get** /v1/api/tokens/{ppan}/applepay-reference-data | Get ApplePay reference data
[**GetGooglePayReferenceData**](Class7TokenAPI.md#GetGooglePayReferenceData) | **Get** /v1/api/tokens/{ppan}/googlepay-reference-data | Get GooglePay reference data
[**GetTokens**](Class7TokenAPI.md#GetTokens) | **Get** /v1/api/tokens/{ppan} | Get card token



## ChangeTokenStatus

> ChangeTokenStatus(ctx, tokenNumber).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).TokenStatusChange(tokenStatusChange).Execute()

Change tokens status



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
	tokenNumber := "tokenNumber_example" // string | Token number
	tokenStatusChange := *openapiclient.NewTokenStatusChange("TokenStatus_example", "Details_example") // TokenStatusChange | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class7TokenAPI.ChangeTokenStatus(context.Background(), tokenNumber).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).TokenStatusChange(tokenStatusChange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class7TokenAPI.ChangeTokenStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenNumber** | **string** | Token number | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTokenStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **tokenStatusChange** | [**TokenStatusChange**](TokenStatusChange.md) |  | 

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

> PanReferenceData GetApplePayReferenceData(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Get ApplePay reference data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class7TokenAPI.GetApplePayReferenceData(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class7TokenAPI.GetApplePayReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplePayReferenceData`: PanReferenceData
	fmt.Fprintf(os.Stdout, "Response from `Class7TokenAPI.GetApplePayReferenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplePayReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**PanReferenceData**](PanReferenceData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGooglePayReferenceData

> Data GetGooglePayReferenceData(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Get GooglePay reference data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class7TokenAPI.GetGooglePayReferenceData(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class7TokenAPI.GetGooglePayReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGooglePayReferenceData`: Data
	fmt.Fprintf(os.Stdout, "Response from `Class7TokenAPI.GetGooglePayReferenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGooglePayReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**Data**](Data.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> TokenInfo GetTokens(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Get card token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class7TokenAPI.GetTokens(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class7TokenAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: TokenInfo
	fmt.Fprintf(os.Stdout, "Response from `Class7TokenAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**TokenInfo**](TokenInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

