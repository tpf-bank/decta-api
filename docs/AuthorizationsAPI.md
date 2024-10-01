# \AuthorizationsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHoldsList**](AuthorizationsAPI.md#GetHoldsList) | **Get** /v1/api/cards/{ppan}/holds | Get card authorizations



## GetHoldsList

> HoldsDataArray GetHoldsList(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get card authorizations



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
	ppan := "123456ABCDEFG1234" // string | masked card number
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationsAPI.GetHoldsList(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsAPI.GetHoldsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHoldsList`: HoldsDataArray
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationsAPI.GetHoldsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**HoldsDataArray**](HoldsDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

