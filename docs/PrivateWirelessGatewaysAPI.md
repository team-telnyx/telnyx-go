# \PrivateWirelessGatewaysAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateWirelessGateway**](PrivateWirelessGatewaysAPI.md#CreatePrivateWirelessGateway) | **Post** /private_wireless_gateways | Create a Private Wireless Gateway
[**DeleteWirelessGateway**](PrivateWirelessGatewaysAPI.md#DeleteWirelessGateway) | **Delete** /private_wireless_gateways/{id} | Delete a Private Wireless Gateway
[**GetPrivateWirelessGateway**](PrivateWirelessGatewaysAPI.md#GetPrivateWirelessGateway) | **Get** /private_wireless_gateways/{id} | Get a Private Wireless Gateway
[**GetPrivateWirelessGateways**](PrivateWirelessGatewaysAPI.md#GetPrivateWirelessGateways) | **Get** /private_wireless_gateways | Get all Private Wireless Gateways



## CreatePrivateWirelessGateway

> CreatePrivateWirelessGateway202Response CreatePrivateWirelessGateway(ctx).CreatePrivateWirelessGatewayRequest(createPrivateWirelessGatewayRequest).Execute()

Create a Private Wireless Gateway



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
	createPrivateWirelessGatewayRequest := *openapiclient.NewCreatePrivateWirelessGatewayRequest("6a09cdc3-8948-47f0-aa62-74ac943d6c58", "My private wireless gateway") // CreatePrivateWirelessGatewayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateWirelessGatewaysAPI.CreatePrivateWirelessGateway(context.Background()).CreatePrivateWirelessGatewayRequest(createPrivateWirelessGatewayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateWirelessGatewaysAPI.CreatePrivateWirelessGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateWirelessGateway`: CreatePrivateWirelessGateway202Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateWirelessGatewaysAPI.CreatePrivateWirelessGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateWirelessGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPrivateWirelessGatewayRequest** | [**CreatePrivateWirelessGatewayRequest**](CreatePrivateWirelessGatewayRequest.md) |  | 

### Return type

[**CreatePrivateWirelessGateway202Response**](CreatePrivateWirelessGateway202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWirelessGateway

> GetPrivateWirelessGateway200Response DeleteWirelessGateway(ctx, id).Execute()

Delete a Private Wireless Gateway



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the private wireless gateway.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateWirelessGatewaysAPI.DeleteWirelessGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateWirelessGatewaysAPI.DeleteWirelessGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWirelessGateway`: GetPrivateWirelessGateway200Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateWirelessGatewaysAPI.DeleteWirelessGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the private wireless gateway. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWirelessGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPrivateWirelessGateway200Response**](GetPrivateWirelessGateway200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateWirelessGateway

> GetPrivateWirelessGateway200Response GetPrivateWirelessGateway(ctx, id).Execute()

Get a Private Wireless Gateway



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the private wireless gateway.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateWirelessGatewaysAPI.GetPrivateWirelessGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateWirelessGatewaysAPI.GetPrivateWirelessGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateWirelessGateway`: GetPrivateWirelessGateway200Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateWirelessGatewaysAPI.GetPrivateWirelessGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the private wireless gateway. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateWirelessGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPrivateWirelessGateway200Response**](GetPrivateWirelessGateway200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateWirelessGateways

> GetPrivateWirelessGateways200Response GetPrivateWirelessGateways(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).FilterIpRange(filterIpRange).FilterRegionCode(filterRegionCode).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).Execute()

Get all Private Wireless Gateways



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
	filterName := "my private gateway" // string | The name of the Private Wireless Gateway. (optional)
	filterIpRange := "192.168.0.0/24" // string | The IP address range of the Private Wireless Gateway. (optional)
	filterRegionCode := "dc2" // string | The name of the region where the Private Wireless Gateway is deployed. (optional)
	filterCreatedAt := "2018-02-02T22:25:27.521Z" // string | Private Wireless Gateway resource creation date. (optional)
	filterUpdatedAt := "2018-02-02T22:25:27.521Z" // string | When the Private Wireless Gateway was last updated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateWirelessGatewaysAPI.GetPrivateWirelessGateways(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).FilterIpRange(filterIpRange).FilterRegionCode(filterRegionCode).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateWirelessGatewaysAPI.GetPrivateWirelessGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateWirelessGateways`: GetPrivateWirelessGateways200Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateWirelessGatewaysAPI.GetPrivateWirelessGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateWirelessGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterName** | **string** | The name of the Private Wireless Gateway. | 
 **filterIpRange** | **string** | The IP address range of the Private Wireless Gateway. | 
 **filterRegionCode** | **string** | The name of the region where the Private Wireless Gateway is deployed. | 
 **filterCreatedAt** | **string** | Private Wireless Gateway resource creation date. | 
 **filterUpdatedAt** | **string** | When the Private Wireless Gateway was last updated. | 

### Return type

[**GetPrivateWirelessGateways200Response**](GetPrivateWirelessGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

