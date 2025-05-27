# \WireGuardInterfacesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWireguardInterface**](WireGuardInterfacesAPI.md#CreateWireguardInterface) | **Post** /wireguard_interfaces | Create a WireGuard Interface
[**CreateWireguardPeer**](WireGuardInterfacesAPI.md#CreateWireguardPeer) | **Post** /wireguard_peers | Create a WireGuard Peer
[**DeleteWireguardInterface**](WireGuardInterfacesAPI.md#DeleteWireguardInterface) | **Delete** /wireguard_interfaces/{id} | Delete a WireGuard Interface
[**DeleteWireguardPeer**](WireGuardInterfacesAPI.md#DeleteWireguardPeer) | **Delete** /wireguard_peers/{id} | Delete the WireGuard Peer
[**GetWireguardInterface**](WireGuardInterfacesAPI.md#GetWireguardInterface) | **Get** /wireguard_interfaces/{id} | Retrieve a WireGuard Interfaces
[**GetWireguardPeer**](WireGuardInterfacesAPI.md#GetWireguardPeer) | **Get** /wireguard_peers/{id} | Retrieve the WireGuard Peer
[**GetWireguardPeerConfig**](WireGuardInterfacesAPI.md#GetWireguardPeerConfig) | **Get** /wireguard_peers/{id}/config | Retrieve Wireguard config template for Peer
[**ListWireguardInterfaces**](WireGuardInterfacesAPI.md#ListWireguardInterfaces) | **Get** /wireguard_interfaces | List all WireGuard Interfaces
[**ListWireguardPeers**](WireGuardInterfacesAPI.md#ListWireguardPeers) | **Get** /wireguard_peers | List all WireGuard Peers
[**UpdateWireguardPeer**](WireGuardInterfacesAPI.md#UpdateWireguardPeer) | **Patch** /wireguard_peers/{id} | Update the WireGuard Peer



## CreateWireguardInterface

> CreateWireguardInterface202Response CreateWireguardInterface(ctx).WireguardInterfaceCreate(wireguardInterfaceCreate).Execute()

Create a WireGuard Interface



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
	wireguardInterfaceCreate := *openapiclient.NewWireguardInterfaceCreate("6a09cdc3-8948-47f0-aa62-74ac943d6c58", "ashburn-va") // WireguardInterfaceCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireGuardInterfacesAPI.CreateWireguardInterface(context.Background()).WireguardInterfaceCreate(wireguardInterfaceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.CreateWireguardInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWireguardInterface`: CreateWireguardInterface202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.CreateWireguardInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWireguardInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireguardInterfaceCreate** | [**WireguardInterfaceCreate**](WireguardInterfaceCreate.md) |  | 

### Return type

[**CreateWireguardInterface202Response**](CreateWireguardInterface202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWireguardPeer

> CreateWireguardPeer202Response CreateWireguardPeer(ctx).WireguardPeerCreate(wireguardPeerCreate).Execute()

Create a WireGuard Peer



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
	wireguardPeerCreate := *openapiclient.NewWireguardPeerCreate("6a09cdc3-8948-47f0-aa62-74ac943d6c58") // WireguardPeerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireGuardInterfacesAPI.CreateWireguardPeer(context.Background()).WireguardPeerCreate(wireguardPeerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.CreateWireguardPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWireguardPeer`: CreateWireguardPeer202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.CreateWireguardPeer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWireguardPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireguardPeerCreate** | [**WireguardPeerCreate**](WireguardPeerCreate.md) |  | 

### Return type

[**CreateWireguardPeer202Response**](CreateWireguardPeer202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWireguardInterface

> CreateWireguardInterface202Response DeleteWireguardInterface(ctx, id).Execute()

Delete a WireGuard Interface



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
	resp, r, err := apiClient.WireGuardInterfacesAPI.DeleteWireguardInterface(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.DeleteWireguardInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWireguardInterface`: CreateWireguardInterface202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.DeleteWireguardInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWireguardInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWireguardInterface202Response**](CreateWireguardInterface202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWireguardPeer

> CreateWireguardPeer202Response DeleteWireguardPeer(ctx, id).Execute()

Delete the WireGuard Peer



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
	resp, r, err := apiClient.WireGuardInterfacesAPI.DeleteWireguardPeer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.DeleteWireguardPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWireguardPeer`: CreateWireguardPeer202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.DeleteWireguardPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWireguardPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWireguardPeer202Response**](CreateWireguardPeer202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireguardInterface

> CreateWireguardInterface202Response GetWireguardInterface(ctx, id).Execute()

Retrieve a WireGuard Interfaces



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
	resp, r, err := apiClient.WireGuardInterfacesAPI.GetWireguardInterface(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.GetWireguardInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWireguardInterface`: CreateWireguardInterface202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.GetWireguardInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireguardInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWireguardInterface202Response**](CreateWireguardInterface202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireguardPeer

> CreateWireguardPeer202Response GetWireguardPeer(ctx, id).Execute()

Retrieve the WireGuard Peer



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
	resp, r, err := apiClient.WireGuardInterfacesAPI.GetWireguardPeer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.GetWireguardPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWireguardPeer`: CreateWireguardPeer202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.GetWireguardPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireguardPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWireguardPeer202Response**](CreateWireguardPeer202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireguardPeerConfig

> string GetWireguardPeerConfig(ctx, id).Execute()

Retrieve Wireguard config template for Peer



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
	resp, r, err := apiClient.WireGuardInterfacesAPI.GetWireguardPeerConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.GetWireguardPeerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWireguardPeerConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.GetWireguardPeerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireguardPeerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWireguardInterfaces

> ListWireguardInterfaces200Response ListWireguardInterfaces(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterNetworkId(filterNetworkId).Execute()

List all WireGuard Interfaces



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
	resp, r, err := apiClient.WireGuardInterfacesAPI.ListWireguardInterfaces(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterNetworkId(filterNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.ListWireguardInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWireguardInterfaces`: ListWireguardInterfaces200Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.ListWireguardInterfaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWireguardInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterNetworkId** | **string** | The associated network id to filter on. | 

### Return type

[**ListWireguardInterfaces200Response**](ListWireguardInterfaces200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWireguardPeers

> ListWireguardPeers200Response ListWireguardPeers(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterWireguardInterfaceId(filterWireguardInterfaceId).Execute()

List all WireGuard Peers



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
	filterWireguardInterfaceId := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | The id of the associated WireGuard interface to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireGuardInterfacesAPI.ListWireguardPeers(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterWireguardInterfaceId(filterWireguardInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.ListWireguardPeers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWireguardPeers`: ListWireguardPeers200Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.ListWireguardPeers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWireguardPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterWireguardInterfaceId** | **string** | The id of the associated WireGuard interface to filter on. | 

### Return type

[**ListWireguardPeers200Response**](ListWireguardPeers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWireguardPeer

> CreateWireguardPeer202Response UpdateWireguardPeer(ctx, id).WireguardPeerPatch(wireguardPeerPatch).Execute()

Update the WireGuard Peer



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
	wireguardPeerPatch := *openapiclient.NewWireguardPeerPatch() // WireguardPeerPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireGuardInterfacesAPI.UpdateWireguardPeer(context.Background(), id).WireguardPeerPatch(wireguardPeerPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireGuardInterfacesAPI.UpdateWireguardPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWireguardPeer`: CreateWireguardPeer202Response
	fmt.Fprintf(os.Stdout, "Response from `WireGuardInterfacesAPI.UpdateWireguardPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWireguardPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wireguardPeerPatch** | [**WireguardPeerPatch**](WireguardPeerPatch.md) |  | 

### Return type

[**CreateWireguardPeer202Response**](CreateWireguardPeer202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

