# \VerifiedNumbersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerifiedNumber**](VerifiedNumbersAPI.md#CreateVerifiedNumber) | **Post** /verified_numbers | Request phone number verification
[**DeleteVerifiedNumber**](VerifiedNumbersAPI.md#DeleteVerifiedNumber) | **Delete** /verified_numbers/{phone_number} | Delete a verified number
[**GetVerifiedNumber**](VerifiedNumbersAPI.md#GetVerifiedNumber) | **Get** /verified_numbers/{phone_number} | Retrieve a verified number
[**ListVerifiedNumbers**](VerifiedNumbersAPI.md#ListVerifiedNumbers) | **Get** /verified_numbers | List all Verified Numbers
[**VerifyVerificationCode**](VerifiedNumbersAPI.md#VerifyVerificationCode) | **Post** /verified_numbers/{phone_number}/actions/verify | Submit verification code



## CreateVerifiedNumber

> CreateVerifiedNumberResponse CreateVerifiedNumber(ctx).CreateVerifiedNumberRequest(createVerifiedNumberRequest).Execute()

Request phone number verification



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	createVerifiedNumberRequest := *openapiclient.NewCreateVerifiedNumberRequest("+15551234567", "sms") // CreateVerifiedNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedNumbersAPI.CreateVerifiedNumber(context.Background()).CreateVerifiedNumberRequest(createVerifiedNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedNumbersAPI.CreateVerifiedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerifiedNumber`: CreateVerifiedNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifiedNumbersAPI.CreateVerifiedNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerifiedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerifiedNumberRequest** | [**CreateVerifiedNumberRequest**](CreateVerifiedNumberRequest.md) |  | 

### Return type

[**CreateVerifiedNumberResponse**](CreateVerifiedNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVerifiedNumber

> VerifiedNumberResponseDataWrapper DeleteVerifiedNumber(ctx, phoneNumber).Execute()

Delete a verified number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	phoneNumber := "+15551234567" // string | The phone number being deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedNumbersAPI.DeleteVerifiedNumber(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedNumbersAPI.DeleteVerifiedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVerifiedNumber`: VerifiedNumberResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifiedNumbersAPI.DeleteVerifiedNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | The phone number being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVerifiedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerifiedNumberResponseDataWrapper**](VerifiedNumberResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerifiedNumber

> VerifiedNumberResponseDataWrapper GetVerifiedNumber(ctx, phoneNumber).Execute()

Retrieve a verified number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	phoneNumber := "+15551234567" // string | The phone number being requested.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedNumbersAPI.GetVerifiedNumber(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedNumbersAPI.GetVerifiedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerifiedNumber`: VerifiedNumberResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifiedNumbersAPI.GetVerifiedNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | The phone number being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerifiedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerifiedNumberResponseDataWrapper**](VerifiedNumberResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerifiedNumbers

> ListVerifiedNumbersResponse ListVerifiedNumbers(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()

List all Verified Numbers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	pageSize := int32(56) // int32 |  (optional) (default to 25)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedNumbersAPI.ListVerifiedNumbers(context.Background()).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedNumbersAPI.ListVerifiedNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVerifiedNumbers`: ListVerifiedNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifiedNumbersAPI.ListVerifiedNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVerifiedNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | [default to 25]
 **pageNumber** | **int32** |  | [default to 1]

### Return type

[**ListVerifiedNumbersResponse**](ListVerifiedNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationCode

> VerifiedNumberResponseDataWrapper VerifyVerificationCode(ctx, phoneNumber).VerifyVerificationCodeRequest(verifyVerificationCodeRequest).Execute()

Submit verification code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	phoneNumber := "+15551234567" // string | The phone number being verified.
	verifyVerificationCodeRequest := *openapiclient.NewVerifyVerificationCodeRequest("123456") // VerifyVerificationCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedNumbersAPI.VerifyVerificationCode(context.Background(), phoneNumber).VerifyVerificationCodeRequest(verifyVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedNumbersAPI.VerifyVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyVerificationCode`: VerifiedNumberResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifiedNumbersAPI.VerifyVerificationCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | The phone number being verified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyVerificationCodeRequest** | [**VerifyVerificationCodeRequest**](VerifyVerificationCodeRequest.md) |  | 

### Return type

[**VerifiedNumberResponseDataWrapper**](VerifiedNumberResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

