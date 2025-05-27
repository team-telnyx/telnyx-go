# \IPsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIp**](IPsAPI.md#CreateIp) | **Post** /ips | Create an Ip
[**DeleteIp**](IPsAPI.md#DeleteIp) | **Delete** /ips/{id} | Delete an Ip
[**ListIps**](IPsAPI.md#ListIps) | **Get** /ips | List Ips
[**RetrieveIp**](IPsAPI.md#RetrieveIp) | **Get** /ips/{id} | Retrieve an Ip
[**UpdateIp**](IPsAPI.md#UpdateIp) | **Patch** /ips/{id} | Update an Ip



## CreateIp

> IpResponse CreateIp(ctx).CreateIpRequest(createIpRequest).Execute()

Create an Ip



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
	createIpRequest := *openapiclient.NewCreateIpRequest("192.168.0.0") // CreateIpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPsAPI.CreateIp(context.Background()).CreateIpRequest(createIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPsAPI.CreateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIp`: IpResponse
	fmt.Fprintf(os.Stdout, "Response from `IPsAPI.CreateIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIpRequest** | [**CreateIpRequest**](CreateIpRequest.md) |  | 

### Return type

[**IpResponse**](IpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIp

> IpResponse DeleteIp(ctx, id).Execute()

Delete an Ip



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the type of resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPsAPI.DeleteIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPsAPI.DeleteIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIp`: IpResponse
	fmt.Fprintf(os.Stdout, "Response from `IPsAPI.DeleteIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpResponse**](IpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIps

> ListIpsResponse ListIps(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionId(filterConnectionId).FilterIpAddress(filterIpAddress).FilterPort(filterPort).Execute()

List Ips



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
	filterConnectionId := "filterConnectionId_example" // string | ID of the IP Connection to which this IP should be attached. (optional)
	filterIpAddress := "192.168.0.0" // string | IP adddress represented by this resource. (optional)
	filterPort := int32(5060) // int32 | Port to use when connecting to this IP. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPsAPI.ListIps(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionId(filterConnectionId).FilterIpAddress(filterIpAddress).FilterPort(filterPort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPsAPI.ListIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIps`: ListIpsResponse
	fmt.Fprintf(os.Stdout, "Response from `IPsAPI.ListIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterConnectionId** | **string** | ID of the IP Connection to which this IP should be attached. | 
 **filterIpAddress** | **string** | IP adddress represented by this resource. | 
 **filterPort** | **int32** | Port to use when connecting to this IP. | 

### Return type

[**ListIpsResponse**](ListIpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIp

> IpResponse RetrieveIp(ctx, id).Execute()

Retrieve an Ip



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the type of resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPsAPI.RetrieveIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPsAPI.RetrieveIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveIp`: IpResponse
	fmt.Fprintf(os.Stdout, "Response from `IPsAPI.RetrieveIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpResponse**](IpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIp

> IpResponse UpdateIp(ctx, id).UpdateIpRequest(updateIpRequest).Execute()

Update an Ip



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the type of resource.
	updateIpRequest := *openapiclient.NewUpdateIpRequest("192.168.0.0") // UpdateIpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPsAPI.UpdateIp(context.Background(), id).UpdateIpRequest(updateIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPsAPI.UpdateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIp`: IpResponse
	fmt.Fprintf(os.Stdout, "Response from `IPsAPI.UpdateIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIpRequest** | [**UpdateIpRequest**](UpdateIpRequest.md) |  | 

### Return type

[**IpResponse**](IpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

