# \Class2CardsAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignPin**](Class2CardsAPI.md#AssignPin) | **Post** /v1/api/cards/{ppan}/assign-pin | Change PIN code
[**Create**](Class2CardsAPI.md#Create) | **Post** /v2/api/cards | Create card V2
[**Create1**](Class2CardsAPI.md#Create1) | **Post** /v1/api/cards | Create card
[**CreateBatchCard**](Class2CardsAPI.md#CreateBatchCard) | **Post** /v2/api/cards/batch | Batch create card
[**GetCardData1**](Class2CardsAPI.md#GetCardData1) | **Get** /v1/api/cards/{ppan}/card-data1 | Get sensitive card information 1
[**GetCardData2**](Class2CardsAPI.md#GetCardData2) | **Get** /v1/api/cards/{ppan}/card-data2 | Get sensitive card information 2
[**GetCardData3**](Class2CardsAPI.md#GetCardData3) | **Get** /v1/api/cards/{ppan}/card-data3 | Get sensitive card information 3
[**GetCardData4**](Class2CardsAPI.md#GetCardData4) | **Get** /v1/api/cards/{ppan}/card-data4 | Get sensitive card information 4
[**GetCardData5**](Class2CardsAPI.md#GetCardData5) | **Get** /v1/api/cards/{ppan}/card-data5 | Get sensitive card information 5
[**GetCardData6**](Class2CardsAPI.md#GetCardData6) | **Get** /v2/api/cards/{ppan}/card-data6 | Get sensitive card information 6
[**GetInfo**](Class2CardsAPI.md#GetInfo) | **Get** /v1/api/cards/{ppan} | Get card information
[**GetList**](Class2CardsAPI.md#GetList) | **Get** /v1/api/cards | Search cards
[**GetTspSecret**](Class2CardsAPI.md#GetTspSecret) | **Get** /v1/api/cards/{ppan}/otp-secret | Get OTP secret
[**RenewCard**](Class2CardsAPI.md#RenewCard) | **Post** /v1/api/cards/{ppan}/renew | Renew card
[**ReplaceCard**](Class2CardsAPI.md#ReplaceCard) | **Post** /v1/api/cards/{ppan}/replace | Replace card
[**UpdateCardUserDefinedFields**](Class2CardsAPI.md#UpdateCardUserDefinedFields) | **Patch** /v1/api/cards/{ppan}/user-defined-fields | Change card user defined fields
[**UpdateState**](Class2CardsAPI.md#UpdateState) | **Patch** /v1/api/cards/{ppan}/state | Change card status



## AssignPin

> AssignPin(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinAssign(pinAssign).Execute()

Change PIN code



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
	pinAssign := *openapiclient.NewPinAssign("0527", "EncryptedPINBlock_example", "EncryptionCertificateThumbprint_example") // PinAssign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class2CardsAPI.AssignPin(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).PinAssign(pinAssign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.AssignPin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAssignPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **pinAssign** | [**PinAssign**](PinAssign.md) |  | 

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


## Create

> CreateCardResponse Create(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CreateCard(createCard).Execute()

Create card V2



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
	createCard := *openapiclient.NewCreateCard(*openapiclient.NewCardPreferencesCreateCard("112", "["EUR"]", "true")) // CreateCard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.Create(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CreateCard(createCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateCardResponse
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **createCard** | [**CreateCard**](CreateCard.md) |  | 

### Return type

[**CreateCardResponse**](CreateCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create1

> CreateCardResponse Create1(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CreateCard(createCard).Execute()

Create card



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
	createCard := *openapiclient.NewCreateCard(*openapiclient.NewCardPreferencesCreateCard("112", "["EUR"]", "true")) // CreateCard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.Create1(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CreateCard(createCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.Create1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create1`: CreateCardResponse
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.Create1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **createCard** | [**CreateCard**](CreateCard.md) |  | 

### Return type

[**CreateCardResponse**](CreateCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchCard

> CreateBatchCard(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).BatchCreateCard(batchCreateCard).Execute()

Batch create card



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
	batchCreateCard := *openapiclient.NewBatchCreateCard("1", *openapiclient.NewBatchCreateCardTemplateApiDto(*openapiclient.NewBatchCardPrivateClientApiDto("John", "Doe", "Schwedak"), *openapiclient.NewBatchCardPreferencesApiDto("112", "["EUR"]", *openapiclient.NewBatchCardDeliveryAddressApiDto("GBR", "London", "62 Bayswater Road"), "BANKID123456"))) // BatchCreateCard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class2CardsAPI.CreateBatchCard(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).BatchCreateCard(batchCreateCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.CreateBatchCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **batchCreateCard** | [**BatchCreateCard**](BatchCreateCard.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData1

> CardData1ApiDto GetCardData1(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Execute()

Get sensitive card information 1



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
	param1 := "param1_example" // string | parameter 1

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetCardData1(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetCardData1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData1`: CardData1ApiDto
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetCardData1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **param1** | **string** | parameter 1 | 

### Return type

[**CardData1ApiDto**](CardData1ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData2

> CardData2ApiDto GetCardData2(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Execute()

Get sensitive card information 2



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
	param1 := "param1_example" // string | parameter 1

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetCardData2(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetCardData2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData2`: CardData2ApiDto
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetCardData2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **param1** | **string** | parameter 1 | 

### Return type

[**CardData2ApiDto**](CardData2ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData3

> CardData3HolderApiDto GetCardData3(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Secret(secret).Execute()

Get sensitive card information 3



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
	secret := "secret_example" // string | OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetCardData3(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetCardData3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData3`: CardData3HolderApiDto
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetCardData3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **secret** | **string** | OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info | 

### Return type

[**CardData3HolderApiDto**](CardData3HolderApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData4

> Data4 GetCardData4(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Secret(secret).Execute()

Get sensitive card information 4



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
	secret := "secret_example" // string | OTP secret value generated by card issuer and passed to the token service provider to grant permission to fetch for the sensitive info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetCardData4(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetCardData4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData4`: Data4
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetCardData4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **secret** | **string** | OTP secret value generated by card issuer and passed to the token service provider to grant permission to fetch for the sensitive info | 

### Return type

[**Data4**](Data4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData5

> CardData2ApiDto GetCardData5(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Execute()

Get sensitive card information 5



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
	param1 := "param1_example" // string | parameter 1

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetCardData5(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetCardData5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData5`: CardData2ApiDto
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetCardData5`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **param1** | **string** | parameter 1 | 

### Return type

[**CardData2ApiDto**](CardData2ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardData6

> CardData6ApiDto GetCardData6(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Secret(secret).Execute()

Get sensitive card information 6



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
	param1 := "param1_example" // string | One time user side generated asymmetric keys public part. One-line key header and footer should be removed or separated with newline character.
	secret := "secret_example" // string | OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetCardData6(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Param1(param1).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetCardData6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardData6`: CardData6ApiDto
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetCardData6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardData6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **param1** | **string** | One time user side generated asymmetric keys public part. One-line key header and footer should be removed or separated with newline character. | 
 **secret** | **string** | OTP secret value generated by card issuer and passed to the token service provider to grant permission for the sensitive info. | 

### Return type

[**CardData6ApiDto**](CardData6ApiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> CardInfo GetInfo(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Get card information



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
	resp, r, err := apiClient.Class2CardsAPI.GetInfo(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfo`: CardInfo
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**CardInfo**](CardInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList

> CardInfoDataArray GetList(ctx).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardName(cardName).Expiry(expiry).CardState(cardState).CardAccount(cardAccount).ClientId(clientId).Product(product).UserDefinedField12(userDefinedField12).Count(count).StartId(startId).Execute()

Search cards



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
	cardName := "eq:John Doe" // string | The filter allows users to get a list of cards by the selected name on the card. Allowed filter operator(s): eq, ne (optional)
	expiry := []openapiclient.DateFilter{*openapiclient.NewDateFilter()} // []DateFilter | The filter allows users to sort cards by the selected expiration date in the list. Allowed filter operator(s): eq, ne, gt, ge, lt, le (optional)
	cardState := "eq:BLOCKED_BY_HOLDER" // string | The filter allows users to get a list of cards by the selected  card status. Allowed filter operator(s): eq (optional)
	cardAccount := "eq:00000001" // string | The filter allows users to get a list of cards by the selected card account. Allowed filter operator(s): eq, ne (optional)
	clientId := "eq:CR0000000000000001" // string | The filter allows users to get a list of cards by the selected client identification number. Allowed filter operator(s): eq, ne (optional)
	product := "eq:200" // string | The filter allows users to get a list of cards by selected product code. Allowed filter operator(s): eq, ne (optional)
	userDefinedField12 := "eq:BANKID123456" // string | The filter allows users to get a list of cards by selected userDefinedField12. Allowed filter operator(s): eq (optional)
	count := int32(10) // int32 | Number of cards in the list (optional) (default to 10)
	startId := int64(4012) // int64 | Start the ordered list from card with selected id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class2CardsAPI.GetList(context.Background()).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardName(cardName).Expiry(expiry).CardState(cardState).CardAccount(cardAccount).ClientId(clientId).Product(product).UserDefinedField12(userDefinedField12).Count(count).StartId(startId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetList`: CardInfoDataArray
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 
 **cardName** | **string** | The filter allows users to get a list of cards by the selected name on the card. Allowed filter operator(s): eq, ne | 
 **expiry** | [**[]DateFilter**](DateFilter.md) | The filter allows users to sort cards by the selected expiration date in the list. Allowed filter operator(s): eq, ne, gt, ge, lt, le | 
 **cardState** | **string** | The filter allows users to get a list of cards by the selected  card status. Allowed filter operator(s): eq | 
 **cardAccount** | **string** | The filter allows users to get a list of cards by the selected card account. Allowed filter operator(s): eq, ne | 
 **clientId** | **string** | The filter allows users to get a list of cards by the selected client identification number. Allowed filter operator(s): eq, ne | 
 **product** | **string** | The filter allows users to get a list of cards by selected product code. Allowed filter operator(s): eq, ne | 
 **userDefinedField12** | **string** | The filter allows users to get a list of cards by selected userDefinedField12. Allowed filter operator(s): eq | 
 **count** | **int32** | Number of cards in the list | [default to 10]
 **startId** | **int64** | Start the ordered list from card with selected id | 

### Return type

[**CardInfoDataArray**](CardInfoDataArray.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTspSecret

> TspSecret GetTspSecret(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Get OTP secret



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
	resp, r, err := apiClient.Class2CardsAPI.GetTspSecret(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.GetTspSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTspSecret`: TspSecret
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.GetTspSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTspSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**TspSecret**](TspSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewCard

> CardRenewInfo RenewCard(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Renew card



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
	resp, r, err := apiClient.Class2CardsAPI.RenewCard(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.RenewCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewCard`: CardRenewInfo
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.RenewCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**CardRenewInfo**](CardRenewInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCard

> CardInfo ReplaceCard(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()

Replace card



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
	resp, r, err := apiClient.Class2CardsAPI.ReplaceCard(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.ReplaceCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCard`: CardInfo
	fmt.Fprintf(os.Stdout, "Response from `Class2CardsAPI.ReplaceCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ppan** | **string** | Masked card number | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 


### Return type

[**CardInfo**](CardInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCardUserDefinedFields

> UpdateCardUserDefinedFields(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardUserDefinedFields(cardUserDefinedFields).Execute()

Change card user defined fields



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
	cardUserDefinedFields := *openapiclient.NewCardUserDefinedFields() // CardUserDefinedFields | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class2CardsAPI.UpdateCardUserDefinedFields(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).CardUserDefinedFields(cardUserDefinedFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.UpdateCardUserDefinedFields``: %v\n", err)
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
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **cardUserDefinedFields** | [**CardUserDefinedFields**](CardUserDefinedFields.md) |  | 

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

> UpdateState(ctx, ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).StatusChange(statusChange).Execute()

Change card status



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
	statusChange := *openapiclient.NewStatusChange("CardState_example", "Details_example") // StatusChange | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class2CardsAPI.UpdateState(context.Background(), ppan).RequestId(requestId).TokenHeader(tokenHeader).TokenSignature(tokenSignature).StatusChange(statusChange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class2CardsAPI.UpdateState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID (UUID format) | 
 **tokenHeader** | **string** | URL64Encoded without padding Header part of JWS token | 
 **tokenSignature** | **string** | URL64Encoded without padding Signature part of JWS | 

 **statusChange** | [**StatusChange**](StatusChange.md) |  | 

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

