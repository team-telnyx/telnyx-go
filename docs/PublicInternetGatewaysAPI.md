# \PublicInternetGatewaysAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicInternetGateway**](PublicInternetGatewaysAPI.md#CreatePublicInternetGateway) | **Post** /public_internet_gateways | Create a Public Internet Gateway
[**DeletePublicInternetGateway**](PublicInternetGatewaysAPI.md#DeletePublicInternetGateway) | **Delete** /public_internet_gateways/{id} | Delete a Public Internet Gateway
[**GetPublicInternetGateway**](PublicInternetGatewaysAPI.md#GetPublicInternetGateway) | **Get** /public_internet_gateways/{id} | Retrieve a Public Internet Gateway
[**ListPublicInternetGateways**](PublicInternetGatewaysAPI.md#ListPublicInternetGateways) | **Get** /public_internet_gateways | List all Public Internet Gateways



## CreatePublicInternetGateway

> CreatePublicInternetGateway202Response CreatePublicInternetGateway(ctx).PublicInternetGatewayCreate(publicInternetGatewayCreate).Execute()

Create a Public Internet Gateway



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
	publicInternetGatewayCreate := *openapiclient.NewPublicInternetGatewayCreate("6a09cdc3-8948-47f0-aa62-74ac943d6c58") // PublicInternetGatewayCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicInternetGatewaysAPI.CreatePublicInternetGateway(context.Background()).PublicInternetGatewayCreate(publicInternetGatewayCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicInternetGatewaysAPI.CreatePublicInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicInternetGateway`: CreatePublicInternetGateway202Response
	fmt.Fprintf(os.Stdout, "Response from `PublicInternetGatewaysAPI.CreatePublicInternetGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicInternetGatewayCreate** | [**PublicInternetGatewayCreate**](PublicInternetGatewayCreate.md) |  | 

### Return type

[**CreatePublicInternetGateway202Response**](CreatePublicInternetGateway202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicInternetGateway

> CreatePublicInternetGateway202Response DeletePublicInternetGateway(ctx, id).Execute()

Delete a Public Internet Gateway



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicInternetGatewaysAPI.DeletePublicInternetGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicInternetGatewaysAPI.DeletePublicInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePublicInternetGateway`: CreatePublicInternetGateway202Response
	fmt.Fprintf(os.Stdout, "Response from `PublicInternetGatewaysAPI.DeletePublicInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreatePublicInternetGateway202Response**](CreatePublicInternetGateway202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicInternetGateway

> CreatePublicInternetGateway202Response GetPublicInternetGateway(ctx, id).Execute()

Retrieve a Public Internet Gateway



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicInternetGatewaysAPI.GetPublicInternetGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicInternetGatewaysAPI.GetPublicInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicInternetGateway`: CreatePublicInternetGateway202Response
	fmt.Fprintf(os.Stdout, "Response from `PublicInternetGatewaysAPI.GetPublicInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreatePublicInternetGateway202Response**](CreatePublicInternetGateway202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicInternetGateways

> ListPublicInternetGateways200Response ListPublicInternetGateways(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterNetworkId(filterNetworkId).Execute()

List all Public Internet Gateways



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterNetworkId := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | The associated network id to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicInternetGatewaysAPI.ListPublicInternetGateways(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterNetworkId(filterNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicInternetGatewaysAPI.ListPublicInternetGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicInternetGateways`: ListPublicInternetGateways200Response
	fmt.Fprintf(os.Stdout, "Response from `PublicInternetGatewaysAPI.ListPublicInternetGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicInternetGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterNetworkId** | **string** | The associated network id to filter on. | 

### Return type

[**ListPublicInternetGateways200Response**](ListPublicInternetGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

