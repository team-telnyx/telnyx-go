# \FQDNConnectionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFqdnConnection**](FQDNConnectionsAPI.md#CreateFqdnConnection) | **Post** /fqdn_connections | Create an FQDN connection
[**DeleteFqdnConnection**](FQDNConnectionsAPI.md#DeleteFqdnConnection) | **Delete** /fqdn_connections/{id} | Delete an FQDN connection
[**ListFqdnConnections**](FQDNConnectionsAPI.md#ListFqdnConnections) | **Get** /fqdn_connections | List FQDN connections
[**RetrieveFqdnConnection**](FQDNConnectionsAPI.md#RetrieveFqdnConnection) | **Get** /fqdn_connections/{id} | Retrieve an FQDN connection
[**UpdateFqdnConnection**](FQDNConnectionsAPI.md#UpdateFqdnConnection) | **Patch** /fqdn_connections/{id} | Update an FQDN connection



## CreateFqdnConnection

> FQDNConnectionResponse CreateFqdnConnection(ctx).CreateFqdnConnectionRequest(createFqdnConnectionRequest).Execute()

Create an FQDN connection



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
	createFqdnConnectionRequest := *openapiclient.NewCreateFqdnConnectionRequest("office-connection") // CreateFqdnConnectionRequest | Parameters that can be defined during FQDN connection creation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNConnectionsAPI.CreateFqdnConnection(context.Background()).CreateFqdnConnectionRequest(createFqdnConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNConnectionsAPI.CreateFqdnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFqdnConnection`: FQDNConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNConnectionsAPI.CreateFqdnConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFqdnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFqdnConnectionRequest** | [**CreateFqdnConnectionRequest**](CreateFqdnConnectionRequest.md) | Parameters that can be defined during FQDN connection creation | 

### Return type

[**FQDNConnectionResponse**](FQDNConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFqdnConnection

> FQDNConnectionResponse DeleteFqdnConnection(ctx, id).Execute()

Delete an FQDN connection



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNConnectionsAPI.DeleteFqdnConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNConnectionsAPI.DeleteFqdnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFqdnConnection`: FQDNConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNConnectionsAPI.DeleteFqdnConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFqdnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FQDNConnectionResponse**](FQDNConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFqdnConnections

> ListFQDNConnectionsResponse ListFqdnConnections(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterFqdn(filterFqdn).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()

List FQDN connections



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
	filterConnectionNameContains := "filterConnectionNameContains_example" // string | If present, connections with <code>connection_name</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional)
	filterFqdn := "filterFqdn_example" // string | If present, connections with an `fqdn` that equals the given value will be returned. Matching is case-sensitive, and the full string must match. (optional)
	filterOutboundVoiceProfileId := "1293384261075731499" // string | Identifies the associated outbound voice profile. (optional)
	sort := "connection_name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>connection_name</code>: sorts the result by the     <code>connection_name</code> field in ascending order.   </li>    <li>     <code>-connection_name</code>: sorts the result by the     <code>connection_name</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNConnectionsAPI.ListFqdnConnections(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterFqdn(filterFqdn).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNConnectionsAPI.ListFqdnConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFqdnConnections`: ListFQDNConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNConnectionsAPI.ListFqdnConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFqdnConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterConnectionNameContains** | **string** | If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | 
 **filterFqdn** | **string** | If present, connections with an &#x60;fqdn&#x60; that equals the given value will be returned. Matching is case-sensitive, and the full string must match. | 
 **filterOutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**ListFQDNConnectionsResponse**](ListFQDNConnectionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFqdnConnection

> FQDNConnectionResponse RetrieveFqdnConnection(ctx, id).Execute()

Retrieve an FQDN connection



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNConnectionsAPI.RetrieveFqdnConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNConnectionsAPI.RetrieveFqdnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFqdnConnection`: FQDNConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNConnectionsAPI.RetrieveFqdnConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFqdnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FQDNConnectionResponse**](FQDNConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFqdnConnection

> FQDNConnectionResponse UpdateFqdnConnection(ctx, id).UpdateFqdnConnectionRequest(updateFqdnConnectionRequest).Execute()

Update an FQDN connection



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
	id := "1293384261075731499" // string | Identifies the resource.
	updateFqdnConnectionRequest := *openapiclient.NewUpdateFqdnConnectionRequest() // UpdateFqdnConnectionRequest | Parameters that can be updated in a FQDN connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FQDNConnectionsAPI.UpdateFqdnConnection(context.Background(), id).UpdateFqdnConnectionRequest(updateFqdnConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FQDNConnectionsAPI.UpdateFqdnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFqdnConnection`: FQDNConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `FQDNConnectionsAPI.UpdateFqdnConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFqdnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFqdnConnectionRequest** | [**UpdateFqdnConnectionRequest**](UpdateFqdnConnectionRequest.md) | Parameters that can be updated in a FQDN connection | 

### Return type

[**FQDNConnectionResponse**](FQDNConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

