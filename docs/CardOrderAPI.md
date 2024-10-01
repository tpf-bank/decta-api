# \CardOrderAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrdersList**](CardOrderAPI.md#GetOrdersList) | **Get** /v1/api/orders | Order status
[**OrderCard**](CardOrderAPI.md#OrderCard) | **Post** /v1/api/cards/order | Order card
[**OrderGiftCard**](CardOrderAPI.md#OrderGiftCard) | **Post** /v1/api/cards/order-gift-card | Create gift cards



## GetOrdersList

> []Order GetOrdersList(ctx).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Id(id).ExternalId(externalId).Count(count).RequestId(requestId).Execute()

Order status



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
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	id := "CRDECTA11201812132210" // string | start from document with id (optional) (default to "")
	externalId := "91857628-8103057" // string | get documents with externalId (optional) (default to "")
	count := int32(10) // int32 | number of card orders</br>max: 10000 (optional) (default to 1)
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardOrderAPI.GetOrdersList(context.Background()).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Id(id).ExternalId(externalId).Count(count).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardOrderAPI.GetOrdersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersList`: []Order
	fmt.Fprintf(os.Stdout, "Response from `CardOrderAPI.GetOrdersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **id** | **string** | start from document with id | [default to &quot;&quot;]
 **externalId** | **string** | get documents with externalId | [default to &quot;&quot;]
 **count** | **int32** | number of card orders&lt;/br&gt;max: 10000 | [default to 1]
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**[]Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCard

> OrderCardResponse OrderCard(ctx).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardOrder(cardOrder).RequestId(requestId).Execute()

Order card



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
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	cardOrder := *openapiclient.NewCardOrder(*openapiclient.NewCardPreferences("112", []string{"["EUR","USD","RUB"]"}, "true")) // CardOrder | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardOrderAPI.OrderCard(context.Background()).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardOrder(cardOrder).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardOrderAPI.OrderCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderCard`: OrderCardResponse
	fmt.Fprintf(os.Stdout, "Response from `CardOrderAPI.OrderCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **cardOrder** | [**CardOrder**](CardOrder.md) |  | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**OrderCardResponse**](OrderCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGiftCard

> GiftOrderCardResponse OrderGiftCard(ctx).TokenHeader(tokenHeader).TokenSignature(tokenSignature).GiftCardOrder(giftCardOrder).RequestId(requestId).Execute()

Create gift cards



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
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	giftCardOrder := *openapiclient.NewGiftCardOrder("1", *openapiclient.NewTemplateApiDto(*openapiclient.NewPrivateClientGiftCardApiDto("John", "Doe", "Schwedak", *openapiclient.NewRegistrationAddressGiftCardApiDto("Months_example", "GBR", "London", "62 Bayswater Road", "W2 3PH"), "en"), *openapiclient.NewCardPreferencesGiftCardApiDto("112", []string{"List [ "EUR", "USD", "RUB" ]"}, *openapiclient.NewDeliveryAddressGiftCardApiDto("Standard", "Janis", "Smith", "GBR", "London", "62 Bayswater Road", "W2 3PH")))) // GiftCardOrder | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardOrderAPI.OrderGiftCard(context.Background()).TokenHeader(tokenHeader).TokenSignature(tokenSignature).GiftCardOrder(giftCardOrder).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardOrderAPI.OrderGiftCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderGiftCard`: GiftOrderCardResponse
	fmt.Fprintf(os.Stdout, "Response from `CardOrderAPI.OrderGiftCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **giftCardOrder** | [**GiftCardOrder**](GiftCardOrder.md) |  | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**GiftOrderCardResponse**](GiftOrderCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

