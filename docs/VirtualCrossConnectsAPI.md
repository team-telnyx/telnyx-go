# \VirtualCrossConnectsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualCrossConnect**](VirtualCrossConnectsAPI.md#CreateVirtualCrossConnect) | **Post** /virtual_cross_connects | Create a Virtual Cross Connect
[**DeleteVirtualCrossConnect**](VirtualCrossConnectsAPI.md#DeleteVirtualCrossConnect) | **Delete** /virtual_cross_connects/{id} | Delete a Virtual Cross Connect
[**GetVirtualCrossConnect**](VirtualCrossConnectsAPI.md#GetVirtualCrossConnect) | **Get** /virtual_cross_connects/{id} | Retrieve a Virtual Cross Connect
[**ListVirtualCrossConnectCoverage**](VirtualCrossConnectsAPI.md#ListVirtualCrossConnectCoverage) | **Get** /virtual_cross_connects_coverage | List Virtual Cross Connect Cloud Coverage
[**ListVirtualCrossConnects**](VirtualCrossConnectsAPI.md#ListVirtualCrossConnects) | **Get** /virtual_cross_connects | List all Virtual Cross Connects
[**UpdateVirtualCrossConnect**](VirtualCrossConnectsAPI.md#UpdateVirtualCrossConnect) | **Patch** /virtual_cross_connects/{id} | Update the Virtual Cross Connect



## CreateVirtualCrossConnect

> CreateVirtualCrossConnect200Response CreateVirtualCrossConnect(ctx).VirtualCrossConnectCreate(virtualCrossConnectCreate).Execute()

Create a Virtual Cross Connect



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
	virtualCrossConnectCreate := *openapiclient.NewVirtualCrossConnectCreate("6a09cdc3-8948-47f0-aa62-74ac943d6c58", "aws", "us-east-1", float32(1234), "123456789012", "ashburn-va") // VirtualCrossConnectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualCrossConnectsAPI.CreateVirtualCrossConnect(context.Background()).VirtualCrossConnectCreate(virtualCrossConnectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualCrossConnectsAPI.CreateVirtualCrossConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualCrossConnect`: CreateVirtualCrossConnect200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualCrossConnectsAPI.CreateVirtualCrossConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualCrossConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualCrossConnectCreate** | [**VirtualCrossConnectCreate**](VirtualCrossConnectCreate.md) |  | 

### Return type

[**CreateVirtualCrossConnect200Response**](CreateVirtualCrossConnect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualCrossConnect

> CreateVirtualCrossConnect200Response DeleteVirtualCrossConnect(ctx, id).Execute()

Delete a Virtual Cross Connect



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
	resp, r, err := apiClient.VirtualCrossConnectsAPI.DeleteVirtualCrossConnect(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualCrossConnectsAPI.DeleteVirtualCrossConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualCrossConnect`: CreateVirtualCrossConnect200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualCrossConnectsAPI.DeleteVirtualCrossConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualCrossConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateVirtualCrossConnect200Response**](CreateVirtualCrossConnect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualCrossConnect

> CreateVirtualCrossConnect200Response GetVirtualCrossConnect(ctx, id).Execute()

Retrieve a Virtual Cross Connect



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
	resp, r, err := apiClient.VirtualCrossConnectsAPI.GetVirtualCrossConnect(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualCrossConnectsAPI.GetVirtualCrossConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualCrossConnect`: CreateVirtualCrossConnect200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualCrossConnectsAPI.GetVirtualCrossConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualCrossConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateVirtualCrossConnect200Response**](CreateVirtualCrossConnect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualCrossConnectCoverage

> ListVirtualCrossConnectCoverage200Response ListVirtualCrossConnectCoverage(ctx).PageNumber(pageNumber).PageSize(pageSize).FiltersAvailableBandwidthContains(filtersAvailableBandwidthContains).FilterCloudProvider(filterCloudProvider).FilterCloudProviderRegion(filterCloudProviderRegion).FilterLocationRegion(filterLocationRegion).FilterLocationSite(filterLocationSite).FilterLocationPop(filterLocationPop).FilterLocationCode(filterLocationCode).Execute()

List Virtual Cross Connect Cloud Coverage



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
	filtersAvailableBandwidthContains := int32(50) // int32 | The available bandwidth to filter on. (optional)
	filterCloudProvider := "aws" // string | The Telnyx region code (optional)
	filterCloudProviderRegion := "us-east-1" // string | The cloud provider region code to filter on (optional)
	filterLocationRegion := "AMER" // string | The region of associated location to filter on. (optional)
	filterLocationSite := "SJC" // string | The site of associated location to filter on. (optional)
	filterLocationPop := "SV1" // string | The POP of associated location to filter on. (optional)
	filterLocationCode := "silicon_valley-ca" // string | The code of associated location to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualCrossConnectsAPI.ListVirtualCrossConnectCoverage(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FiltersAvailableBandwidthContains(filtersAvailableBandwidthContains).FilterCloudProvider(filterCloudProvider).FilterCloudProviderRegion(filterCloudProviderRegion).FilterLocationRegion(filterLocationRegion).FilterLocationSite(filterLocationSite).FilterLocationPop(filterLocationPop).FilterLocationCode(filterLocationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualCrossConnectsAPI.ListVirtualCrossConnectCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualCrossConnectCoverage`: ListVirtualCrossConnectCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualCrossConnectsAPI.ListVirtualCrossConnectCoverage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualCrossConnectCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filtersAvailableBandwidthContains** | **int32** | The available bandwidth to filter on. | 
 **filterCloudProvider** | **string** | The Telnyx region code | 
 **filterCloudProviderRegion** | **string** | The cloud provider region code to filter on | 
 **filterLocationRegion** | **string** | The region of associated location to filter on. | 
 **filterLocationSite** | **string** | The site of associated location to filter on. | 
 **filterLocationPop** | **string** | The POP of associated location to filter on. | 
 **filterLocationCode** | **string** | The code of associated location to filter on. | 

### Return type

[**ListVirtualCrossConnectCoverage200Response**](ListVirtualCrossConnectCoverage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualCrossConnects

> ListVirtualCrossConnects200Response ListVirtualCrossConnects(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterNetworkId(filterNetworkId).Execute()

List all Virtual Cross Connects



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
	resp, r, err := apiClient.VirtualCrossConnectsAPI.ListVirtualCrossConnects(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterNetworkId(filterNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualCrossConnectsAPI.ListVirtualCrossConnects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualCrossConnects`: ListVirtualCrossConnects200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualCrossConnectsAPI.ListVirtualCrossConnects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualCrossConnectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterNetworkId** | **string** | The associated network id to filter on. | 

### Return type

[**ListVirtualCrossConnects200Response**](ListVirtualCrossConnects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualCrossConnect

> CreateVirtualCrossConnect200Response UpdateVirtualCrossConnect(ctx, id).VirtualCrossConnectPatch(virtualCrossConnectPatch).Execute()

Update the Virtual Cross Connect



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
	virtualCrossConnectPatch := *openapiclient.NewVirtualCrossConnectPatch() // VirtualCrossConnectPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualCrossConnectsAPI.UpdateVirtualCrossConnect(context.Background(), id).VirtualCrossConnectPatch(virtualCrossConnectPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualCrossConnectsAPI.UpdateVirtualCrossConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVirtualCrossConnect`: CreateVirtualCrossConnect200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualCrossConnectsAPI.UpdateVirtualCrossConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualCrossConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualCrossConnectPatch** | [**VirtualCrossConnectPatch**](VirtualCrossConnectPatch.md) |  | 

### Return type

[**CreateVirtualCrossConnect200Response**](CreateVirtualCrossConnect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

