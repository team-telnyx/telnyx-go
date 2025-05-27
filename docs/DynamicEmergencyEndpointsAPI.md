# \DynamicEmergencyEndpointsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicEmergencyEndpoint**](DynamicEmergencyEndpointsAPI.md#CreateDynamicEmergencyEndpoint) | **Post** /dynamic_emergency_endpoints | Create a dynamic emergency endpoint.
[**DeleteDynamicEmergencyEndpoint**](DynamicEmergencyEndpointsAPI.md#DeleteDynamicEmergencyEndpoint) | **Delete** /dynamic_emergency_endpoints/{id} | Delete a dynamic emergency endpoint
[**GetDynamicEmergencyEndpoint**](DynamicEmergencyEndpointsAPI.md#GetDynamicEmergencyEndpoint) | **Get** /dynamic_emergency_endpoints/{id} | Get a dynamic emergency endpoint
[**ListDynamicEmergencyEndpoints**](DynamicEmergencyEndpointsAPI.md#ListDynamicEmergencyEndpoints) | **Get** /dynamic_emergency_endpoints | List dynamic emergency endpoints



## CreateDynamicEmergencyEndpoint

> CreateDynamicEmergencyEndpoint201Response CreateDynamicEmergencyEndpoint(ctx).DynamicEmergencyEndpoint(dynamicEmergencyEndpoint).Execute()

Create a dynamic emergency endpoint.



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
	dynamicEmergencyEndpoint := *openapiclient.NewDynamicEmergencyEndpoint("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0", "+13125550000", "Jane Doe Desk Phone") // DynamicEmergencyEndpoint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyEndpointsAPI.CreateDynamicEmergencyEndpoint(context.Background()).DynamicEmergencyEndpoint(dynamicEmergencyEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyEndpointsAPI.CreateDynamicEmergencyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicEmergencyEndpoint`: CreateDynamicEmergencyEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyEndpointsAPI.CreateDynamicEmergencyEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicEmergencyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicEmergencyEndpoint** | [**DynamicEmergencyEndpoint**](DynamicEmergencyEndpoint.md) |  | 

### Return type

[**CreateDynamicEmergencyEndpoint201Response**](CreateDynamicEmergencyEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDynamicEmergencyEndpoint

> CreateDynamicEmergencyEndpoint201Response DeleteDynamicEmergencyEndpoint(ctx, id).Execute()

Delete a dynamic emergency endpoint



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Dynamic Emergency Endpoint id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyEndpointsAPI.DeleteDynamicEmergencyEndpoint(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyEndpointsAPI.DeleteDynamicEmergencyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDynamicEmergencyEndpoint`: CreateDynamicEmergencyEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyEndpointsAPI.DeleteDynamicEmergencyEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Emergency Endpoint id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicEmergencyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDynamicEmergencyEndpoint201Response**](CreateDynamicEmergencyEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicEmergencyEndpoint

> CreateDynamicEmergencyEndpoint201Response GetDynamicEmergencyEndpoint(ctx, id).Execute()

Get a dynamic emergency endpoint



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Dynamic Emergency Endpoint id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyEndpointsAPI.GetDynamicEmergencyEndpoint(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyEndpointsAPI.GetDynamicEmergencyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicEmergencyEndpoint`: CreateDynamicEmergencyEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyEndpointsAPI.GetDynamicEmergencyEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Emergency Endpoint id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicEmergencyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDynamicEmergencyEndpoint201Response**](CreateDynamicEmergencyEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDynamicEmergencyEndpoints

> ListDynamicEmergencyEndpoints200Response ListDynamicEmergencyEndpoints(ctx).FilterStatus(filterStatus).FilterCountryCode(filterCountryCode).PageNumber(pageNumber).PageSize(pageSize).Execute()

List dynamic emergency endpoints



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
	filterStatus := "filterStatus_example" // string | Filter by status. (optional)
	filterCountryCode := "filterCountryCode_example" // string | Filter by country code. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyEndpointsAPI.ListDynamicEmergencyEndpoints(context.Background()).FilterStatus(filterStatus).FilterCountryCode(filterCountryCode).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyEndpointsAPI.ListDynamicEmergencyEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDynamicEmergencyEndpoints`: ListDynamicEmergencyEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyEndpointsAPI.ListDynamicEmergencyEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDynamicEmergencyEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Filter by status. | 
 **filterCountryCode** | **string** | Filter by country code. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListDynamicEmergencyEndpoints200Response**](ListDynamicEmergencyEndpoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

