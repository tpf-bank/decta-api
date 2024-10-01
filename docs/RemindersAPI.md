# \RemindersAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPinTryCounter**](RemindersAPI.md#GetPinTryCounter) | **Get** /v1/api/cards/{ppan}/pin-try-counter | Receive count of PIN code entry tries
[**RemindPanCvv**](RemindersAPI.md#RemindPanCvv) | **Get** /v1/api/cards/{ppan}/pan-cvv | Remind card number
[**RemindPin**](RemindersAPI.md#RemindPin) | **Get** /v1/api/cards/{ppan}/pin | Remind PIN code
[**ResetPinTryCounter**](RemindersAPI.md#ResetPinTryCounter) | **Post** /v1/api/cards/{ppan}/pin-try-counter | Reset count of PIN code entry tries



## GetPinTryCounter

> PinTryCounterInfo GetPinTryCounter(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Receive count of PIN code entry tries



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
	resp, r, err := apiClient.RemindersAPI.GetPinTryCounter(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.GetPinTryCounter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPinTryCounter`: PinTryCounterInfo
	fmt.Fprintf(os.Stdout, "Response from `RemindersAPI.GetPinTryCounter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPinTryCounterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**PinTryCounterInfo**](PinTryCounterInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemindPanCvv

> RemindPanCvv(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Remind card number



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
	r, err := apiClient.RemindersAPI.RemindPanCvv(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.RemindPanCvv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemindPanCvvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemindPin

> RemindPin(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Remind PIN code



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
	r, err := apiClient.RemindersAPI.RemindPin(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.RemindPin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemindPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPinTryCounter

> ResetPinTryCounter(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinTryReset(pinTryReset).RequestId(requestId).Execute()

Reset count of PIN code entry tries



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
	pinTryReset := *openapiclient.NewPinTryReset() // PinTryReset | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemindersAPI.ResetPinTryCounter(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinTryReset(pinTryReset).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.ResetPinTryCounter``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResetPinTryCounterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **pinTryReset** | [**PinTryReset**](PinTryReset.md) |  | 
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

