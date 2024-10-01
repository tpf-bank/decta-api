# \CardsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignPin**](CardsAPI.md#AssignPin) | **Post** /v1/api/cards/{ppan}/assign-pin | Change PIN code
[**GetCardData1**](CardsAPI.md#GetCardData1) | **Get** /v1/api/cards/{ppan}/card-data1 | Get sensitive card information 1
[**GetCardData2**](CardsAPI.md#GetCardData2) | **Get** /v1/api/cards/{ppan}/card-data2 | Get sensitive card information 2
[**GetCardData3**](CardsAPI.md#GetCardData3) | **Get** /v1/api/cards/{ppan}/card-data3 | Get sensitive card information for tokenization-3
[**GetCardData4**](CardsAPI.md#GetCardData4) | **Get** /v1/api/cards/{ppan}/card-data4 | Get sensitive card information for tokenization-4
[**GetCardData5**](CardsAPI.md#GetCardData5) | **Get** /v1/api/cards/{ppan}/card-data5 | Get sensitive card information 5
[**GetInfo**](CardsAPI.md#GetInfo) | **Get** /v1/api/cards/{ppan} | Get card information
[**GetList**](CardsAPI.md#GetList) | **Get** /v1/api/cards | Search cards
[**GetTspSecret**](CardsAPI.md#GetTspSecret) | **Get** /v1/api/cards/{ppan}/otp-secret | Get OTP secret for tokenization
[**RenewCard**](CardsAPI.md#RenewCard) | **Post** /v1/api/cards/{ppan}/renew | Renew card
[**ReplaceCard**](CardsAPI.md#ReplaceCard) | **Post** /v1/api/cards/{ppan}/replace | Replace card
[**UpdateCardUserDefinedFields**](CardsAPI.md#UpdateCardUserDefinedFields) | **Patch** /v1/api/cards/{ppan}/user-defined-fields | Change card user defined fields
[**UpdateState**](CardsAPI.md#UpdateState) | **Patch** /v1/api/cards/{ppan}/state | Change card status



## AssignPin

> AssignPin(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinAssign(pinAssign).RequestId(requestId).Execute()

Change PIN code



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
	pinAssign := *openapiclient.NewPinAssign("0527", "EncryptedPINBlock_example", "EncryptionCertificateThumbprint_example") // PinAssign | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CardsAPI.AssignPin(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinAssign(pinAssign).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.AssignPin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAssignPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **pinAssign** | [**PinAssign**](PinAssign.md) |  | 
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


## GetCardData1

> CardData1ApiDto GetCardData1(ctx, ppan).Param1(param1).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get sensitive card information 1



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
	param1 := "param1_example" // string | parameter 1
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCardData1(context.Background(), ppan).Param1(param1).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardData1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData1`: CardData1ApiDto
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardData1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **param1** | **string** | parameter 1 | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardData1ApiDto**](CardData1ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData2

> CardData2ApiDto GetCardData2(ctx, ppan).Param1(param1).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get sensitive card information 2



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
	param1 := "param1_example" // string | parameter 1
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCardData2(context.Background(), ppan).Param1(param1).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardData2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData2`: CardData2ApiDto
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardData2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **param1** | **string** | parameter 1 | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardData2ApiDto**](CardData2ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData3

> CardData3HolderApiDto GetCardData3(ctx, ppan).Secret(secret).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get sensitive card information for tokenization-3



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
	secret := "secret_example" // string | OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCardData3(context.Background(), ppan).Secret(secret).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardData3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData3`: CardData3HolderApiDto
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardData3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secret** | **string** | OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardData3HolderApiDto**](CardData3HolderApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData4

> Data4 GetCardData4(ctx, ppan).Secret(secret).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get sensitive card information for tokenization-4



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
	secret := "secret_example" // string | OTP secret value generated by card issuer and passed to the token service provider to grant permission to fetch for the sensitive info
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCardData4(context.Background(), ppan).Secret(secret).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardData4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData4`: Data4
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardData4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secret** | **string** | OTP secret value generated by card issuer and passed to the token service provider to grant permission to fetch for the sensitive info | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**Data4**](Data4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData5

> CardData2ApiDto GetCardData5(ctx, ppan).Param1(param1).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get sensitive card information 5



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
	param1 := "param1_example" // string | parameter 1
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetCardData5(context.Background(), ppan).Param1(param1).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardData5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData5`: CardData2ApiDto
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardData5`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **param1** | **string** | parameter 1 | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardData2ApiDto**](CardData2ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> CardInfo GetInfo(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get card information



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
	resp, r, err := apiClient.CardsAPI.GetInfo(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfo`: CardInfo
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardInfo**](CardInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList

> CardInfoDataArray GetList(ctx).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardName(cardName).Expiry(expiry).CardState(cardState).CardAccount(cardAccount).ClientId(clientId).Product(product).Count(count).StartId(startId).RequestId(requestId).Execute()

Search cards



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
	cardName := "eq:Vasily Pupkin" // string | The filter allows users to get a list of cards by the selected name on the card. Allowed filter operator(s): eq, ne (optional)
	expiry := []openapiclient.DateFilter{*openapiclient.NewDateFilter()} // []DateFilter | The filter allows users to sort cards by the selected expiration date in the list. Allowed filter operator(s): eq, ne, gt, ge, lt, le (optional)
	cardState := "eq:BLOCKED_BY_HOLDER" // string | The filter allows users to get a list of cards by the selected  card status. Allowed filter operator(s): eq (optional)
	cardAccount := "eq:00000001" // string | The filter allows users to get a list of cards by the selected card account. Allowed filter operator(s): eq, ne (optional)
	clientId := "eq:CR0000000000000001" // string | The filter allows users to get a list of cards by the selected client identification number. Allowed filter operator(s): eq, ne (optional)
	product := "eq:200" // string | The filter allows users to get a list of cards by selected product code. Allowed filter operator(s): eq, ne (optional)
	count := int32(10) // int32 | Number of cards in the list (optional) (default to 10)
	startId := int64(4012) // int64 | Start the ordered list from card with selected id (optional)
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardsAPI.GetList(context.Background()).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardName(cardName).Expiry(expiry).CardState(cardState).CardAccount(cardAccount).ClientId(clientId).Product(product).Count(count).StartId(startId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetList`: CardInfoDataArray
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **cardName** | **string** | The filter allows users to get a list of cards by the selected name on the card. Allowed filter operator(s): eq, ne | 
 **expiry** | [**[]DateFilter**](DateFilter.md) | The filter allows users to sort cards by the selected expiration date in the list. Allowed filter operator(s): eq, ne, gt, ge, lt, le | 
 **cardState** | **string** | The filter allows users to get a list of cards by the selected  card status. Allowed filter operator(s): eq | 
 **cardAccount** | **string** | The filter allows users to get a list of cards by the selected card account. Allowed filter operator(s): eq, ne | 
 **clientId** | **string** | The filter allows users to get a list of cards by the selected client identification number. Allowed filter operator(s): eq, ne | 
 **product** | **string** | The filter allows users to get a list of cards by selected product code. Allowed filter operator(s): eq, ne | 
 **count** | **int32** | Number of cards in the list | [default to 10]
 **startId** | **int64** | Start the ordered list from card with selected id | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardInfoDataArray**](CardInfoDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTspSecret

> TspSecret GetTspSecret(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Get OTP secret for tokenization



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
	resp, r, err := apiClient.CardsAPI.GetTspSecret(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetTspSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTspSecret`: TspSecret
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetTspSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTspSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**TspSecret**](TspSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewCard

> CardRenewInfo RenewCard(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Renew card



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
	resp, r, err := apiClient.CardsAPI.RenewCard(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.RenewCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewCard`: CardRenewInfo
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.RenewCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardRenewInfo**](CardRenewInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCard

> CardInfo ReplaceCard(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()

Replace card



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
	resp, r, err := apiClient.CardsAPI.ReplaceCard(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.ReplaceCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCard`: CardInfo
	fmt.Fprintf(os.Stdout, "Response from `CardsAPI.ReplaceCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **requestId** | **string** | Request ID (UUID format) | 

### Return type

[**CardInfo**](CardInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCardUserDefinedFields

> UpdateCardUserDefinedFields(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardUserDefinedFields(cardUserDefinedFields).RequestId(requestId).Execute()

Change card user defined fields



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
	ppan := "123456ABCDEF5678" // string | Masked card number
	tokenHeader := "tokenHeader_example" // string | URL64Encoded without padding Header part of JWS token
	tokenSignature := "tokenSignature_example" // string | URL64Encoded without padding Signature part of JWS
	cardUserDefinedFields := *openapiclient.NewCardUserDefinedFields() // CardUserDefinedFields | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CardsAPI.UpdateCardUserDefinedFields(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardUserDefinedFields(cardUserDefinedFields).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.UpdateCardUserDefinedFields``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCardUserDefinedFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **cardUserDefinedFields** | [**CardUserDefinedFields**](CardUserDefinedFields.md) |  | 
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


## UpdateState

> UpdateState(ctx, ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).StatusChange(statusChange).RequestId(requestId).Execute()

Change card status



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
	statusChange := *openapiclient.NewStatusChange("CardState_example", "Details_example") // StatusChange | 
	requestId := "requestId_example" // string | Request ID (UUID format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CardsAPI.UpdateState(context.Background(), ppan).TokenHeader(tokenHeader).TokenSignature(tokenSignature).StatusChange(statusChange).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.UpdateState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **statusChange** | [**StatusChange**](StatusChange.md) |  | 
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

