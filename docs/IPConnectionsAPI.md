# \IPConnectionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpConnection**](IPConnectionsAPI.md#CreateIpConnection) | **Post** /ip_connections | Create an Ip connection
[**DeleteIpConnection**](IPConnectionsAPI.md#DeleteIpConnection) | **Delete** /ip_connections/{id} | Delete an Ip connection
[**ListIpConnections**](IPConnectionsAPI.md#ListIpConnections) | **Get** /ip_connections | List Ip connections
[**RetrieveIpConnection**](IPConnectionsAPI.md#RetrieveIpConnection) | **Get** /ip_connections/{id} | Retrieve an Ip connection
[**UpdateIpConnection**](IPConnectionsAPI.md#UpdateIpConnection) | **Patch** /ip_connections/{id} | Update an Ip connection



## CreateIpConnection

> IpConnectionResponse CreateIpConnection(ctx).CreateIpConnectionRequest(createIpConnectionRequest).Execute()

Create an Ip connection



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
	createIpConnectionRequest := *openapiclient.NewCreateIpConnectionRequest() // CreateIpConnectionRequest | Parameters that can be defined during IP connection creation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPConnectionsAPI.CreateIpConnection(context.Background()).CreateIpConnectionRequest(createIpConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPConnectionsAPI.CreateIpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIpConnection`: IpConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IPConnectionsAPI.CreateIpConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIpConnectionRequest** | [**CreateIpConnectionRequest**](CreateIpConnectionRequest.md) | Parameters that can be defined during IP connection creation | 

### Return type

[**IpConnectionResponse**](IpConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIpConnection

> IpConnectionResponse DeleteIpConnection(ctx, id).Execute()

Delete an Ip connection



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
	id := "id_example" // string | Identifies the type of resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPConnectionsAPI.DeleteIpConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPConnectionsAPI.DeleteIpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIpConnection`: IpConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IPConnectionsAPI.DeleteIpConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpConnectionResponse**](IpConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpConnections

> ListIpConnectionsResponse ListIpConnections(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterOutboundOutboundVoiceProfileId(filterOutboundOutboundVoiceProfileId).Sort(sort).Execute()

List Ip connections



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
	filterOutboundOutboundVoiceProfileId := "1293384261075731499" // string | Identifies the associated outbound voice profile. (optional)
	sort := "connection_name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>connection_name</code>: sorts the result by the     <code>connection_name</code> field in ascending order.   </li>    <li>     <code>-connection_name</code>: sorts the result by the     <code>connection_name</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPConnectionsAPI.ListIpConnections(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterOutboundOutboundVoiceProfileId(filterOutboundOutboundVoiceProfileId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPConnectionsAPI.ListIpConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIpConnections`: ListIpConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `IPConnectionsAPI.ListIpConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIpConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterConnectionNameContains** | **string** | If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | 
 **filterOutboundOutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**ListIpConnectionsResponse**](ListIpConnectionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIpConnection

> IpConnectionResponse RetrieveIpConnection(ctx, id).Execute()

Retrieve an Ip connection



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
	id := "id_example" // string | IP Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPConnectionsAPI.RetrieveIpConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPConnectionsAPI.RetrieveIpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveIpConnection`: IpConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IPConnectionsAPI.RetrieveIpConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IP Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpConnectionResponse**](IpConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpConnection

> IpConnectionResponse UpdateIpConnection(ctx, id).UpdateIpConnectionRequest(updateIpConnectionRequest).Execute()

Update an Ip connection



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
	id := "id_example" // string | Identifies the type of resource.
	updateIpConnectionRequest := *openapiclient.NewUpdateIpConnectionRequest() // UpdateIpConnectionRequest | Parameters that can be updated in a IP connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPConnectionsAPI.UpdateIpConnection(context.Background(), id).UpdateIpConnectionRequest(updateIpConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPConnectionsAPI.UpdateIpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIpConnection`: IpConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IPConnectionsAPI.UpdateIpConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIpConnectionRequest** | [**UpdateIpConnectionRequest**](UpdateIpConnectionRequest.md) | Parameters that can be updated in a IP connection | 

### Return type

[**IpConnectionResponse**](IpConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

