# \TransactionsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CardTransactions**](TransactionsAPI.md#CardTransactions) | **Get** /v1/api/cards/{ppan}/transactions | Get card transactions
[**DoTransaction**](TransactionsAPI.md#DoTransaction) | **Post** /v1/api/cards/{ppan}/transactions | Process card transaction



## CardTransactions

> TransactionInfoDataArray CardTransactions(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).IncludeAccountTransactions(includeAccountTransactions).HoldNumber(holdNumber).TrnTime(trnTime).Count(count).StartId(startId).RequestId(requestId).Execute()

Get card transactions



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
	includeAccountTransactions := true // bool | Include account transactions into result transactions list (optional)
	holdNumber := "eq:1341235" // string | Allows to select transactions from a  list, by providing an id of matched authorization lock which was released by receiving transaction information from the payment system. Allowed filter operator(s): eq, ne (optional)
	trnTime := []openapiclient.DateTimeFilter{*openapiclient.NewDateTimeFilter()} // []DateTimeFilter | The filter allows users to specify transaction time. Users can select multiple time frames. Allowed filter operator(s): eq, ne, gt, ge, lt, le (optional)
	count := int32(10) // int32 | Number of transactions (optional) (default to 10)
	startId := int64(4012) // int64 | Start the ordered list from transaction with selected transaction id (optional)
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.CardTransactions(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).IncludeAccountTransactions(includeAccountTransactions).HoldNumber(holdNumber).TrnTime(trnTime).Count(count).StartId(startId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CardTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CardTransactions`: TransactionInfoDataArray
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CardTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **includeAccountTransactions** | **bool** | Include account transactions into result transactions list | 
 **holdNumber** | **string** | Allows to select transactions from a  list, by providing an id of matched authorization lock which was released by receiving transaction information from the payment system. Allowed filter operator(s): eq, ne | 
 **trnTime** | [**[]DateTimeFilter**](DateTimeFilter.md) | The filter allows users to specify transaction time. Users can select multiple time frames. Allowed filter operator(s): eq, ne, gt, ge, lt, le | 
 **count** | **int32** | Number of transactions | [default to 10]
 **startId** | **int64** | Start the ordered list from transaction with selected transaction id | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**TransactionInfoDataArray**](TransactionInfoDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoTransaction

> TransactionInfo DoTransaction(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Transaction(transaction).RequestId(requestId).Execute()

Process card transaction



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
	transaction := *openapiclient.NewTransaction(float32(2.56), "USD", "Tea money", "TOP_UP") // Transaction | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.DoTransaction(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Transaction(transaction).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.DoTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoTransaction`: TransactionInfo
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.DoTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **transaction** | [**Transaction**](Transaction.md) |  | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**TransactionInfo**](TransactionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

