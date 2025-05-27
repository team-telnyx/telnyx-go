# \AdvancedNumberOrdersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdvancedOrderV2**](AdvancedNumberOrdersAPI.md#CreateAdvancedOrderV2) | **Post** /advanced_orders | Create Advanced Order
[**GetAdvancedOrderV2**](AdvancedNumberOrdersAPI.md#GetAdvancedOrderV2) | **Get** /advanced_orders/{order_id} | Get Advanced Order
[**ListAdvancedOrdersV2**](AdvancedNumberOrdersAPI.md#ListAdvancedOrdersV2) | **Get** /advanced_orders | List Advanced Orders



## CreateAdvancedOrderV2

> AdvancedOrderResponse CreateAdvancedOrderV2(ctx).AdvancedOrderRequest(advancedOrderRequest).Execute()

Create Advanced Order

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
	advancedOrderRequest := *openapiclient.NewAdvancedOrderRequest() // AdvancedOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedNumberOrdersAPI.CreateAdvancedOrderV2(context.Background()).AdvancedOrderRequest(advancedOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedNumberOrdersAPI.CreateAdvancedOrderV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdvancedOrderV2`: AdvancedOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `AdvancedNumberOrdersAPI.CreateAdvancedOrderV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdvancedOrderV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advancedOrderRequest** | [**AdvancedOrderRequest**](AdvancedOrderRequest.md) |  | 

### Return type

[**AdvancedOrderResponse**](AdvancedOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdvancedOrderV2

> AdvancedOrderResponse GetAdvancedOrderV2(ctx, orderId).Execute()

Get Advanced Order

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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedNumberOrdersAPI.GetAdvancedOrderV2(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedNumberOrdersAPI.GetAdvancedOrderV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdvancedOrderV2`: AdvancedOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `AdvancedNumberOrdersAPI.GetAdvancedOrderV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedOrderV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdvancedOrderResponse**](AdvancedOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdvancedOrdersV2

> ListAdvancedOrderResponse ListAdvancedOrdersV2(ctx).Execute()

List Advanced Orders

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedNumberOrdersAPI.ListAdvancedOrdersV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedNumberOrdersAPI.ListAdvancedOrdersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdvancedOrdersV2`: ListAdvancedOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `AdvancedNumberOrdersAPI.ListAdvancedOrdersV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdvancedOrdersV2Request struct via the builder pattern


### Return type

[**ListAdvancedOrderResponse**](ListAdvancedOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

