# \Class1CardOrderAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrdersList**](Class1CardOrderAPI.md#GetOrdersList) | **Get** /v1/api/orders | Order status
[**OrderCard**](Class1CardOrderAPI.md#OrderCard) | **Post** /v1/api/cards/order | Order card
[**OrderGiftCard**](Class1CardOrderAPI.md#OrderGiftCard) | **Post** /v1/api/cards/order-gift-card | Create gift cards



## GetOrdersList

> []Order GetOrdersList(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Id(id).ExternalId(externalId).Count(count).Execute()

Order status



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
	id := "CRDECTA11201812132210" // string | start from document with id (optional) (default to "")
	externalId := "91857628-8103057" // string | get documents with externalId (optional) (default to "")
	count := int32(10) // int32 | number of card orders</br>max: 10000 (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1CardOrderAPI.GetOrdersList(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Id(id).ExternalId(externalId).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1CardOrderAPI.GetOrdersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersList`: []Order
	fmt.Fprintf(os.Stdout, "Response from `Class1CardOrderAPI.GetOrdersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **id** | **string** | start from document with id | [default to &quot;&quot;]
 **externalId** | **string** | get documents with externalId | [default to &quot;&quot;]
 **count** | **int32** | number of card orders&lt;/br&gt;max: 10000 | [default to 1]

### Return type

[**[]Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCard

> OrderCardResponse OrderCard(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardOrder(cardOrder).Execute()

Order card



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
	cardOrder := *openapiclient.NewCardOrder(*openapiclient.NewCardPreferences("112", "["EUR"]", "true")) // CardOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1CardOrderAPI.OrderCard(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardOrder(cardOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1CardOrderAPI.OrderCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderCard`: OrderCardResponse
	fmt.Fprintf(os.Stdout, "Response from `Class1CardOrderAPI.OrderCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **cardOrder** | [**CardOrder**](CardOrder.md) |  | 

### Return type

[**OrderCardResponse**](OrderCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGiftCard

> GiftOrderCardResponse OrderGiftCard(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).GiftCardOrder(giftCardOrder).Execute()

Create gift cards



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
	giftCardOrder := *openapiclient.NewGiftCardOrder("1", *openapiclient.NewTemplateApiDto(*openapiclient.NewPrivateClientGiftCardApiDto("John", "Doe", "Schwedak", *openapiclient.NewRegistrationAddressGiftCardApiDto("Months_example", "GBR", "London", "62 Bayswater Road", "W2 3PH"), "en"), *openapiclient.NewCardPreferencesGiftCardApiDto("112", "List [ "EUR", "USD", "RUB" ]", *openapiclient.NewDeliveryAddressGiftCardApiDto("Standard", "Janis", "Smith", "GBR", "London", "62 Bayswater Road", "W2 3PH")))) // GiftCardOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1CardOrderAPI.OrderGiftCard(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).GiftCardOrder(giftCardOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1CardOrderAPI.OrderGiftCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderGiftCard`: GiftOrderCardResponse
	fmt.Fprintf(os.Stdout, "Response from `Class1CardOrderAPI.OrderGiftCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **giftCardOrder** | [**GiftCardOrder**](GiftCardOrder.md) |  | 

### Return type

[**GiftOrderCardResponse**](GiftOrderCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

