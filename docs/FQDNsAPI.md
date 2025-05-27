# \FQDNsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFqdn**](FQDNsAPI.md#CreateFqdn) | **Post** /fqdns | Create an FQDN
[**DeleteFqdn**](FQDNsAPI.md#DeleteFqdn) | **Delete** /fqdns/{id} | Delete an FQDN
[**ListFqdns**](FQDNsAPI.md#ListFqdns) | **Get** /fqdns | List FQDNs
[**RetrieveFqdn**](FQDNsAPI.md#RetrieveFqdn) | **Get** /fqdns/{id} | Retrieve an FQDN
[**UpdateFqdn**](FQDNsAPI.md#UpdateFqdn) | **Patch** /fqdns/{id} | Update an FQDN



## CreateFqdn

> FQDNResponse CreateFqdn(ctx).CreateFqdnRequest(createFqdnRequest).Execute()

Create an FQDN



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
	createFqdnRequest := *openapiclient.NewCreateFqdnRequest("ConnectionId_example", "example.com", "a") // CreateFqdnRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNsAPI.CreateFqdn(context.Background()).CreateFqdnRequest(createFqdnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNsAPI.CreateFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFqdn`: FQDNResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNsAPI.CreateFqdn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFqdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFqdnRequest** | [**CreateFqdnRequest**](CreateFqdnRequest.md) |  | 

### Return type

[**FQDNResponse**](FQDNResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFqdn

> FQDNResponse DeleteFqdn(ctx, id).Execute()

Delete an FQDN



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
	id := "1517907029795014409" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNsAPI.DeleteFqdn(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNsAPI.DeleteFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFqdn`: FQDNResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNsAPI.DeleteFqdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFqdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FQDNResponse**](FQDNResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFqdns

> ListFQDNsResponse ListFqdns(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionId(filterConnectionId).FilterFqdn(filterFqdn).FilterPort(filterPort).FilterDnsRecordType(filterDnsRecordType).Execute()

List FQDNs



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
	filterConnectionId := "filterConnectionId_example" // string | ID of the FQDN connection to which the FQDN belongs. (optional)
	filterFqdn := "example.com" // string | FQDN represented by the resource. (optional)
	filterPort := int32(5060) // int32 | Port to use when connecting to the FQDN. (optional)
	filterDnsRecordType := "a" // string | DNS record type used by the FQDN. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNsAPI.ListFqdns(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionId(filterConnectionId).FilterFqdn(filterFqdn).FilterPort(filterPort).FilterDnsRecordType(filterDnsRecordType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNsAPI.ListFqdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFqdns`: ListFQDNsResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNsAPI.ListFqdns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFqdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterConnectionId** | **string** | ID of the FQDN connection to which the FQDN belongs. | 
 **filterFqdn** | **string** | FQDN represented by the resource. | 
 **filterPort** | **int32** | Port to use when connecting to the FQDN. | 
 **filterDnsRecordType** | **string** | DNS record type used by the FQDN. | 

### Return type

[**ListFQDNsResponse**](ListFQDNsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFqdn

> FQDNResponse RetrieveFqdn(ctx, id).Execute()

Retrieve an FQDN



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
	id := "1517907029795014409" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNsAPI.RetrieveFqdn(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNsAPI.RetrieveFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFqdn`: FQDNResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNsAPI.RetrieveFqdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFqdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FQDNResponse**](FQDNResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFqdn

> FQDNResponse UpdateFqdn(ctx, id).UpdateFqdnRequest(updateFqdnRequest).Execute()

Update an FQDN



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
	id := "1517907029795014409" // string | Identifies the resource.
	updateFqdnRequest := *openapiclient.NewUpdateFqdnRequest() // UpdateFqdnRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNsAPI.UpdateFqdn(context.Background(), id).UpdateFqdnRequest(updateFqdnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNsAPI.UpdateFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFqdn`: FQDNResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNsAPI.UpdateFqdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFqdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFqdnRequest** | [**UpdateFqdnRequest**](UpdateFqdnRequest.md) |  | 

### Return type

[**FQDNResponse**](FQDNResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

