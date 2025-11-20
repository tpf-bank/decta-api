# \Class8AccountsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionsQueueInfo**](Class8AccountsAPI.md#GetTransactionsQueueInfo) | **Get** /v1/api/accounts/{cardAccount}/transactions-queue | Get account transactions and commissions in queue



## GetTransactionsQueueInfo

> DataArray GetTransactionsQueueInfo(ctx, cardAccount).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientId(clientId).Execute()

Get account transactions and commissions in queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/dectav2"
)

func main() {
	requestId := "requestId_example" // string | Request ID (UUID format)
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	cardAccount := "cardAccount_example" // string | Card account identification number
	clientId := "CR0000000000000001" // string | ClientId is a unique client id number generated on the DECTA partner side. Length 19 symbols

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class8AccountsAPI.GetTransactionsQueueInfo(context.Background(), cardAccount).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class8AccountsAPI.GetTransactionsQueueInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsQueueInfo`: DataArray
	fmt.Fprintf(os.Stdout, "Response from `Class8AccountsAPI.GetTransactionsQueueInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardAccount** | **string** | Card account identification number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsQueueInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **clientId** | **string** | ClientId is a unique client id number generated on the DECTA partner side. Length 19 symbols | 

### Return type

[**DataArray**](DataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

