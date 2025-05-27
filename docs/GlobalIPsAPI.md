# \GlobalIPsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalIp**](GlobalIPsAPI.md#CreateGlobalIp) | **Post** /global_ips | Create a Global IP
[**CreateGlobalIpAssignment**](GlobalIPsAPI.md#CreateGlobalIpAssignment) | **Post** /global_ip_assignments | Create a Global IP assignment
[**CreateGlobalIpHealthCheck**](GlobalIPsAPI.md#CreateGlobalIpHealthCheck) | **Post** /global_ip_health_checks | Create a Global IP health check
[**DeleteGlobalIp**](GlobalIPsAPI.md#DeleteGlobalIp) | **Delete** /global_ips/{id} | Delete a Global IP
[**DeleteGlobalIpAssignment**](GlobalIPsAPI.md#DeleteGlobalIpAssignment) | **Delete** /global_ip_assignments/{id} | Delete a Global IP assignment
[**DeleteGlobalIpHealthCheck**](GlobalIPsAPI.md#DeleteGlobalIpHealthCheck) | **Delete** /global_ip_health_checks/{id} | Delete a Global IP health check
[**GetGlobalIp**](GlobalIPsAPI.md#GetGlobalIp) | **Get** /global_ips/{id} | Retrieve a Global IP
[**GetGlobalIpAssignment**](GlobalIPsAPI.md#GetGlobalIpAssignment) | **Get** /global_ip_assignments/{id} | Retrieve a Global IP
[**GetGlobalIpAssignmentHealth**](GlobalIPsAPI.md#GetGlobalIpAssignmentHealth) | **Get** /global_ip_assignment_health | Global IP Assignment Health Check Metrics
[**GetGlobalIpAssignmentUsage**](GlobalIPsAPI.md#GetGlobalIpAssignmentUsage) | **Get** /global_ip_assignments_usage | Global IP Assignment Usage Metrics
[**GetGlobalIpHealthCheck**](GlobalIPsAPI.md#GetGlobalIpHealthCheck) | **Get** /global_ip_health_checks/{id} | Retrieve a Global IP health check
[**GetGlobalIpLatency**](GlobalIPsAPI.md#GetGlobalIpLatency) | **Get** /global_ip_latency | Global IP Latency Metrics
[**GetGlobalIpUsage**](GlobalIPsAPI.md#GetGlobalIpUsage) | **Get** /global_ip_usage | Global IP Usage Metrics
[**ListGlobalIpAllowedPorts**](GlobalIPsAPI.md#ListGlobalIpAllowedPorts) | **Get** /global_ip_allowed_ports | List all Global IP Allowed Ports
[**ListGlobalIpAssignments**](GlobalIPsAPI.md#ListGlobalIpAssignments) | **Get** /global_ip_assignments | List all Global IP assignments
[**ListGlobalIpHealthCheckTypes**](GlobalIPsAPI.md#ListGlobalIpHealthCheckTypes) | **Get** /global_ip_health_check_types | List all Global IP Health check types
[**ListGlobalIpHealthChecks**](GlobalIPsAPI.md#ListGlobalIpHealthChecks) | **Get** /global_ip_health_checks | List all Global IP health checks
[**ListGlobalIpProtocols**](GlobalIPsAPI.md#ListGlobalIpProtocols) | **Get** /global_ip_protocols | List all Global IP Protocols
[**ListGlobalIps**](GlobalIPsAPI.md#ListGlobalIps) | **Get** /global_ips | List all Global IPs
[**UpdateGlobalIpAssignment**](GlobalIPsAPI.md#UpdateGlobalIpAssignment) | **Patch** /global_ip_assignments/{id} | Update a Global IP assignment



## CreateGlobalIp

> CreateGlobalIp202Response CreateGlobalIp(ctx).GlobalIP(globalIP).Execute()

Create a Global IP



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
	globalIP := *openapiclient.NewGlobalIP() // GlobalIP | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.CreateGlobalIp(context.Background()).GlobalIP(globalIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.CreateGlobalIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlobalIp`: CreateGlobalIp202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.CreateGlobalIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalIP** | [**GlobalIP**](GlobalIP.md) |  | 

### Return type

[**CreateGlobalIp202Response**](CreateGlobalIp202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalIpAssignment

> CreateGlobalIpAssignment202Response CreateGlobalIpAssignment(ctx).GlobalIpAssignment(globalIpAssignment).Execute()

Create a Global IP assignment



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
	globalIpAssignment := *openapiclient.NewGlobalIpAssignment() // GlobalIpAssignment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.CreateGlobalIpAssignment(context.Background()).GlobalIpAssignment(globalIpAssignment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.CreateGlobalIpAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlobalIpAssignment`: CreateGlobalIpAssignment202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.CreateGlobalIpAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalIpAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalIpAssignment** | [**GlobalIpAssignment**](GlobalIpAssignment.md) |  | 

### Return type

[**CreateGlobalIpAssignment202Response**](CreateGlobalIpAssignment202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalIpHealthCheck

> CreateGlobalIpHealthCheck202Response CreateGlobalIpHealthCheck(ctx).GlobalIPHealthCheck(globalIPHealthCheck).Execute()

Create a Global IP health check



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
	globalIPHealthCheck := *openapiclient.NewGlobalIPHealthCheck() // GlobalIPHealthCheck | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.CreateGlobalIpHealthCheck(context.Background()).GlobalIPHealthCheck(globalIPHealthCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.CreateGlobalIpHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlobalIpHealthCheck`: CreateGlobalIpHealthCheck202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.CreateGlobalIpHealthCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalIpHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalIPHealthCheck** | [**GlobalIPHealthCheck**](GlobalIPHealthCheck.md) |  | 

### Return type

[**CreateGlobalIpHealthCheck202Response**](CreateGlobalIpHealthCheck202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalIp

> CreateGlobalIp202Response DeleteGlobalIp(ctx, id).Execute()

Delete a Global IP



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
	resp, r, err := apiClient.GlobalIPsAPI.DeleteGlobalIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.DeleteGlobalIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlobalIp`: CreateGlobalIp202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.DeleteGlobalIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateGlobalIp202Response**](CreateGlobalIp202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalIpAssignment

> CreateGlobalIpAssignment202Response DeleteGlobalIpAssignment(ctx, id).Execute()

Delete a Global IP assignment



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
	resp, r, err := apiClient.GlobalIPsAPI.DeleteGlobalIpAssignment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.DeleteGlobalIpAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlobalIpAssignment`: CreateGlobalIpAssignment202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.DeleteGlobalIpAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalIpAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateGlobalIpAssignment202Response**](CreateGlobalIpAssignment202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalIpHealthCheck

> CreateGlobalIpHealthCheck202Response DeleteGlobalIpHealthCheck(ctx, id).Execute()

Delete a Global IP health check



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
	resp, r, err := apiClient.GlobalIPsAPI.DeleteGlobalIpHealthCheck(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.DeleteGlobalIpHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlobalIpHealthCheck`: CreateGlobalIpHealthCheck202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.DeleteGlobalIpHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalIpHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateGlobalIpHealthCheck202Response**](CreateGlobalIpHealthCheck202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIp

> CreateGlobalIp202Response GetGlobalIp(ctx, id).Execute()

Retrieve a Global IP



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
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIp`: CreateGlobalIp202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateGlobalIp202Response**](CreateGlobalIp202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIpAssignment

> CreateGlobalIpAssignment202Response GetGlobalIpAssignment(ctx, id).Execute()

Retrieve a Global IP



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
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIpAssignment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIpAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIpAssignment`: CreateGlobalIpAssignment202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIpAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateGlobalIpAssignment202Response**](CreateGlobalIpAssignment202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIpAssignmentHealth

> GetGlobalIpAssignmentHealth200Response GetGlobalIpAssignmentHealth(ctx).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterGlobalIpAssignmentIdIn(filterGlobalIpAssignmentIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()

Global IP Assignment Health Check Metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterGlobalIpIdIn := "filterGlobalIpIdIn_example" // string | Filter by Global IP ID(s) separated by commas (optional)
	filterGlobalIpAssignmentIdIn := "filterGlobalIpAssignmentIdIn_example" // string | Filter by Global IP Assignment ID(s) separated by commas (optional)
	filterTimestampGt := time.Now() // time.Time | Filter by timestamp greater than (optional)
	filterTimestampLt := time.Now() // time.Time | Filter by timestamp less than (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIpAssignmentHealth(context.Background()).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterGlobalIpAssignmentIdIn(filterGlobalIpAssignmentIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIpAssignmentHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIpAssignmentHealth`: GetGlobalIpAssignmentHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIpAssignmentHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpAssignmentHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGlobalIpIdIn** | **string** | Filter by Global IP ID(s) separated by commas | 
 **filterGlobalIpAssignmentIdIn** | **string** | Filter by Global IP Assignment ID(s) separated by commas | 
 **filterTimestampGt** | **time.Time** | Filter by timestamp greater than | 
 **filterTimestampLt** | **time.Time** | Filter by timestamp less than | 

### Return type

[**GetGlobalIpAssignmentHealth200Response**](GetGlobalIpAssignmentHealth200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIpAssignmentUsage

> GetGlobalIpAssignmentUsage200Response GetGlobalIpAssignmentUsage(ctx).FilterGlobalIpAssignmentIdIn(filterGlobalIpAssignmentIdIn).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()

Global IP Assignment Usage Metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterGlobalIpAssignmentIdIn := "filterGlobalIpAssignmentIdIn_example" // string | Filter by Global IP Assignment ID(s) separated by commas (optional)
	filterGlobalIpIdIn := "filterGlobalIpIdIn_example" // string | Filter by Global IP ID(s), separated by commas (optional)
	filterTimestampGt := time.Now() // time.Time | Filter by timestamp greater than (optional)
	filterTimestampLt := time.Now() // time.Time | Filter by timestamp less than (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIpAssignmentUsage(context.Background()).FilterGlobalIpAssignmentIdIn(filterGlobalIpAssignmentIdIn).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIpAssignmentUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIpAssignmentUsage`: GetGlobalIpAssignmentUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIpAssignmentUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpAssignmentUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGlobalIpAssignmentIdIn** | **string** | Filter by Global IP Assignment ID(s) separated by commas | 
 **filterGlobalIpIdIn** | **string** | Filter by Global IP ID(s), separated by commas | 
 **filterTimestampGt** | **time.Time** | Filter by timestamp greater than | 
 **filterTimestampLt** | **time.Time** | Filter by timestamp less than | 

### Return type

[**GetGlobalIpAssignmentUsage200Response**](GetGlobalIpAssignmentUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIpHealthCheck

> CreateGlobalIpHealthCheck202Response GetGlobalIpHealthCheck(ctx, id).Execute()

Retrieve a Global IP health check



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
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIpHealthCheck(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIpHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIpHealthCheck`: CreateGlobalIpHealthCheck202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIpHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateGlobalIpHealthCheck202Response**](CreateGlobalIpHealthCheck202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIpLatency

> GetGlobalIpLatency200Response GetGlobalIpLatency(ctx).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()

Global IP Latency Metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterGlobalIpIdIn := "filterGlobalIpIdIn_example" // string | Filter by Global IP ID(s) separated by commas (optional)
	filterTimestampGt := time.Now() // time.Time | Filter by timestamp greater than (optional)
	filterTimestampLt := time.Now() // time.Time | Filter by timestamp less than (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIpLatency(context.Background()).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIpLatency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIpLatency`: GetGlobalIpLatency200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIpLatency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGlobalIpIdIn** | **string** | Filter by Global IP ID(s) separated by commas | 
 **filterTimestampGt** | **time.Time** | Filter by timestamp greater than | 
 **filterTimestampLt** | **time.Time** | Filter by timestamp less than | 

### Return type

[**GetGlobalIpLatency200Response**](GetGlobalIpLatency200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIpUsage

> GetGlobalIpUsage200Response GetGlobalIpUsage(ctx).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()

Global IP Usage Metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterGlobalIpIdIn := "filterGlobalIpIdIn_example" // string | Filter by Global IP ID(s) separated by commas (optional)
	filterTimestampGt := time.Now() // time.Time | Filter by timestamp greater than (optional)
	filterTimestampLt := time.Now() // time.Time | Filter by timestamp less than (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.GetGlobalIpUsage(context.Background()).FilterGlobalIpIdIn(filterGlobalIpIdIn).FilterTimestampGt(filterTimestampGt).FilterTimestampLt(filterTimestampLt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.GetGlobalIpUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIpUsage`: GetGlobalIpUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.GetGlobalIpUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIpUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGlobalIpIdIn** | **string** | Filter by Global IP ID(s) separated by commas | 
 **filterTimestampGt** | **time.Time** | Filter by timestamp greater than | 
 **filterTimestampLt** | **time.Time** | Filter by timestamp less than | 

### Return type

[**GetGlobalIpUsage200Response**](GetGlobalIpUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalIpAllowedPorts

> ListGlobalIpAllowedPorts200Response ListGlobalIpAllowedPorts(ctx).Execute()

List all Global IP Allowed Ports



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
	resp, r, err := apiClient.GlobalIPsAPI.ListGlobalIpAllowedPorts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.ListGlobalIpAllowedPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalIpAllowedPorts`: ListGlobalIpAllowedPorts200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.ListGlobalIpAllowedPorts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalIpAllowedPortsRequest struct via the builder pattern


### Return type

[**ListGlobalIpAllowedPorts200Response**](ListGlobalIpAllowedPorts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalIpAssignments

> ListGlobalIpAssignments200Response ListGlobalIpAssignments(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Global IP assignments



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.ListGlobalIpAssignments(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.ListGlobalIpAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalIpAssignments`: ListGlobalIpAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.ListGlobalIpAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalIpAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListGlobalIpAssignments200Response**](ListGlobalIpAssignments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalIpHealthCheckTypes

> ListGlobalIpHealthCheckTypes200Response ListGlobalIpHealthCheckTypes(ctx).Execute()

List all Global IP Health check types



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
	resp, r, err := apiClient.GlobalIPsAPI.ListGlobalIpHealthCheckTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.ListGlobalIpHealthCheckTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalIpHealthCheckTypes`: ListGlobalIpHealthCheckTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.ListGlobalIpHealthCheckTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalIpHealthCheckTypesRequest struct via the builder pattern


### Return type

[**ListGlobalIpHealthCheckTypes200Response**](ListGlobalIpHealthCheckTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalIpHealthChecks

> ListGlobalIpHealthChecks200Response ListGlobalIpHealthChecks(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Global IP health checks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.ListGlobalIpHealthChecks(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.ListGlobalIpHealthChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalIpHealthChecks`: ListGlobalIpHealthChecks200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.ListGlobalIpHealthChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalIpHealthChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListGlobalIpHealthChecks200Response**](ListGlobalIpHealthChecks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalIpProtocols

> ListGlobalIpProtocols200Response ListGlobalIpProtocols(ctx).Execute()

List all Global IP Protocols



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
	resp, r, err := apiClient.GlobalIPsAPI.ListGlobalIpProtocols(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.ListGlobalIpProtocols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalIpProtocols`: ListGlobalIpProtocols200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.ListGlobalIpProtocols`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalIpProtocolsRequest struct via the builder pattern


### Return type

[**ListGlobalIpProtocols200Response**](ListGlobalIpProtocols200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalIps

> ListGlobalIps200Response ListGlobalIps(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Global IPs



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.ListGlobalIps(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.ListGlobalIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalIps`: ListGlobalIps200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.ListGlobalIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListGlobalIps200Response**](ListGlobalIps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalIpAssignment

> CreateGlobalIpAssignment202Response UpdateGlobalIpAssignment(ctx, id).GlobalIpAssignmentUpdate(globalIpAssignmentUpdate).Execute()

Update a Global IP assignment



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
	globalIpAssignmentUpdate := *openapiclient.NewGlobalIpAssignmentUpdate() // GlobalIpAssignmentUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalIPsAPI.UpdateGlobalIpAssignment(context.Background(), id).GlobalIpAssignmentUpdate(globalIpAssignmentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalIPsAPI.UpdateGlobalIpAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGlobalIpAssignment`: CreateGlobalIpAssignment202Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalIPsAPI.UpdateGlobalIpAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalIpAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalIpAssignmentUpdate** | [**GlobalIpAssignmentUpdate**](GlobalIpAssignmentUpdate.md) |  | 

### Return type

[**CreateGlobalIpAssignment202Response**](CreateGlobalIpAssignment202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

