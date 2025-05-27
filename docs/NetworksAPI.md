# \NetworksAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefaultGateway**](NetworksAPI.md#CreateDefaultGateway) | **Post** /networks/{id}/default_gateway | Create Default Gateway.
[**CreateNetwork**](NetworksAPI.md#CreateNetwork) | **Post** /networks | Create a Network
[**DeleteDefaultGateway**](NetworksAPI.md#DeleteDefaultGateway) | **Delete** /networks/{id}/default_gateway | Delete Default Gateway.
[**DeleteNetwork**](NetworksAPI.md#DeleteNetwork) | **Delete** /networks/{id} | Delete a Network
[**GetDefaultGateway**](NetworksAPI.md#GetDefaultGateway) | **Get** /networks/{id}/default_gateway | Get Default Gateway status.
[**GetNetwork**](NetworksAPI.md#GetNetwork) | **Get** /networks/{id} | Retrieve a Network
[**ListNetworkInterfaces**](NetworksAPI.md#ListNetworkInterfaces) | **Get** /networks/{id}/network_interfaces | List all Interfaces for a Network.
[**ListNetworks**](NetworksAPI.md#ListNetworks) | **Get** /networks | List all Networks
[**UpdateNetwork**](NetworksAPI.md#UpdateNetwork) | **Patch** /networks/{id} | Update a Network



## CreateDefaultGateway

> GetDefaultGateway200Response CreateDefaultGateway(ctx, id).DefaultGateway(defaultGateway).Execute()

Create Default Gateway.



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
	defaultGateway := *openapiclient.NewDefaultGateway() // DefaultGateway | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateDefaultGateway(context.Background(), id).DefaultGateway(defaultGateway).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateDefaultGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDefaultGateway`: GetDefaultGateway200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateDefaultGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDefaultGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultGateway** | [**DefaultGateway**](DefaultGateway.md) |  | 

### Return type

[**GetDefaultGateway200Response**](GetDefaultGateway200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetwork

> CreateNetwork200Response CreateNetwork(ctx).NetworkCreate(networkCreate).Execute()

Create a Network



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
	networkCreate := *openapiclient.NewNetworkCreate("test network") // NetworkCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetwork(context.Background()).NetworkCreate(networkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetwork`: CreateNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkCreate** | [**NetworkCreate**](NetworkCreate.md) |  | 

### Return type

[**CreateNetwork200Response**](CreateNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDefaultGateway

> GetDefaultGateway200Response DeleteDefaultGateway(ctx, id).Execute()

Delete Default Gateway.



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
	resp, r, err := apiClient.NetworksAPI.DeleteDefaultGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteDefaultGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDefaultGateway`: GetDefaultGateway200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteDefaultGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDefaultGateway200Response**](GetDefaultGateway200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> CreateNetwork200Response DeleteNetwork(ctx, id).Execute()

Delete a Network



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
	resp, r, err := apiClient.NetworksAPI.DeleteNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetwork`: CreateNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNetwork200Response**](CreateNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultGateway

> GetDefaultGateway200Response GetDefaultGateway(ctx, id).Execute()

Get Default Gateway status.



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
	resp, r, err := apiClient.NetworksAPI.GetDefaultGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetDefaultGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultGateway`: GetDefaultGateway200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetDefaultGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDefaultGateway200Response**](GetDefaultGateway200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork

> CreateNetwork200Response GetNetwork(ctx, id).Execute()

Retrieve a Network



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
	resp, r, err := apiClient.NetworksAPI.GetNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetwork`: CreateNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNetwork200Response**](CreateNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterfaces

> ListNetworkInterfaces200Response ListNetworkInterfaces(ctx, id).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).FilterType(filterType).FilterStatus(filterStatus).Execute()

List all Interfaces for a Network.



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterName := "test interface" // string | The interface name to filter on. (optional)
	filterType := "wireguard_interface" // string | The interface type to filter on. (optional)
	filterStatus := openapiclient.InterfaceStatus("created") // InterfaceStatus | The interface status to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkInterfaces(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).FilterType(filterType).FilterStatus(filterStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkInterfaces`: ListNetworkInterfaces200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterName** | **string** | The interface name to filter on. | 
 **filterType** | **string** | The interface type to filter on. | 
 **filterStatus** | [**InterfaceStatus**](InterfaceStatus.md) | The interface status to filter on. | 

### Return type

[**ListNetworkInterfaces200Response**](ListNetworkInterfaces200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworks

> ListNetworks200Response ListNetworks(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).Execute()

List all Networks



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
	filterName := "test network" // string | The network name to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworks(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworks`: ListNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterName** | **string** | The network name to filter on. | 

### Return type

[**ListNetworks200Response**](ListNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> CreateNetwork200Response UpdateNetwork(ctx, id).NetworkCreate(networkCreate).Execute()

Update a Network



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
	networkCreate := *openapiclient.NewNetworkCreate("test network") // NetworkCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetwork(context.Background(), id).NetworkCreate(networkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetwork`: CreateNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkCreate** | [**NetworkCreate**](NetworkCreate.md) |  | 

### Return type

[**CreateNetwork200Response**](CreateNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

