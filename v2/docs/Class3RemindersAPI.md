# \Class3RemindersAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPinTryCounter**](Class3RemindersAPI.md#GetPinTryCounter) | **Get** /v1/api/cards/{ppan}/pin-try-counter | Receive count of PIN code entry tries
[**RemindPanCvv**](Class3RemindersAPI.md#RemindPanCvv) | **Get** /v1/api/cards/{ppan}/pan-cvv | Remind card number
[**RemindPin**](Class3RemindersAPI.md#RemindPin) | **Get** /v1/api/cards/{ppan}/pin | Remind PIN code
[**ResetPinTryCounter**](Class3RemindersAPI.md#ResetPinTryCounter) | **Post** /v1/api/cards/{ppan}/pin-try-counter | Reset count of PIN code entry tries



## GetPinTryCounter

> PinTryCounterInfo GetPinTryCounter(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Receive count of PIN code entry tries



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
	resp, r, err := apiClient.Class3RemindersAPI.GetPinTryCounter(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3RemindersAPI.GetPinTryCounter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPinTryCounter`: PinTryCounterInfo
	fmt.Fprintf(os.Stdout, "Response from `Class3RemindersAPI.GetPinTryCounter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPinTryCounterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**PinTryCounterInfo**](PinTryCounterInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemindPanCvv

> RemindPanCvv(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Remind card number



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
	r, err := apiClient.Class3RemindersAPI.RemindPanCvv(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3RemindersAPI.RemindPanCvv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemindPanCvvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


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

> RemindPin(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Remind PIN code



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
	r, err := apiClient.Class3RemindersAPI.RemindPin(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3RemindersAPI.RemindPin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemindPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


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

> ResetPinTryCounter(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinTryReset(pinTryReset).Execute()

Reset count of PIN code entry tries



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
	pinTryReset := *openapiclient.NewPinTryReset("2020-01-01") // PinTryReset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class3RemindersAPI.ResetPinTryCounter(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinTryReset(pinTryReset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3RemindersAPI.ResetPinTryCounter``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResetPinTryCounterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **pinTryReset** | [**PinTryReset**](PinTryReset.md) |  | 

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

