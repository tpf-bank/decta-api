# \Class4TransactionsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CardTransactions**](Class4TransactionsAPI.md#CardTransactions) | **Get** /v1/api/cards/{ppan}/transactions | Get card transactions
[**DoTransaction**](Class4TransactionsAPI.md#DoTransaction) | **Post** /v1/api/cards/{ppan}/transactions | Process card transaction



## CardTransactions

> TransactionInfoDataArray CardTransactions(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).IncludeAccountTransactions(includeAccountTransactions).HoldNumber(holdNumber).TrnTime(trnTime).Count(count).StartId(startId).Execute()

Get card transactions



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
	ppan := "123456ABCDEF5678" // string | Masked card number
	includeAccountTransactions := true // bool | Include account transactions into result transactions list (optional)
	holdNumber := "eq:1341235" // string | Allows to select transactions from a  list, by providing an id of matched authorization lock which was released by receiving transaction information from the payment system. Allowed filter operator(s): eq, ne (optional)
	trnTime := []openapiclient.DateTimeFilter{*openapiclient.NewDateTimeFilter()} // []DateTimeFilter | The filter allows users to specify transaction time. Users can select multiple time frames. Allowed filter operator(s): eq, ne, gt, ge, lt, le (optional)
	count := int32(10) // int32 | Number of transactions (optional) (default to 10)
	startId := int64(4012) // int64 | Start the ordered list from transaction with selected transaction id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class4TransactionsAPI.CardTransactions(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).IncludeAccountTransactions(includeAccountTransactions).HoldNumber(holdNumber).TrnTime(trnTime).Count(count).StartId(startId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class4TransactionsAPI.CardTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CardTransactions`: TransactionInfoDataArray
	fmt.Fprintf(os.Stdout, "Response from `Class4TransactionsAPI.CardTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **includeAccountTransactions** | **bool** | Include account transactions into result transactions list | 
 **holdNumber** | **string** | Allows to select transactions from a  list, by providing an id of matched authorization lock which was released by receiving transaction information from the payment system. Allowed filter operator(s): eq, ne | 
 **trnTime** | [**[]DateTimeFilter**](DateTimeFilter.md) | The filter allows users to specify transaction time. Users can select multiple time frames. Allowed filter operator(s): eq, ne, gt, ge, lt, le | 
 **count** | **int32** | Number of transactions | [default to 10]
 **startId** | **int64** | Start the ordered list from transaction with selected transaction id | 

### Return type

[**TransactionInfoDataArray**](TransactionInfoDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoTransaction

> TransactionInfo DoTransaction(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Transaction(transaction).Execute()

Process card transaction



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
	ppan := "123456ABCDEF5678" // string | Masked card number
	transaction := *openapiclient.NewTransaction(float32(2.56), "USD", "Tea money", "TOP_UP") // Transaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class4TransactionsAPI.DoTransaction(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Transaction(transaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class4TransactionsAPI.DoTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoTransaction`: TransactionInfo
	fmt.Fprintf(os.Stdout, "Response from `Class4TransactionsAPI.DoTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **transaction** | [**Transaction**](Transaction.md) |  | 

### Return type

[**TransactionInfo**](TransactionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

