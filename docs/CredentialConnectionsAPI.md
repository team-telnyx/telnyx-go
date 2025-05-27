# \CredentialConnectionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRegistrationStatus**](CredentialConnectionsAPI.md#CheckRegistrationStatus) | **Post** /credential_connections/{id}/actions/check_registration_status | Check a Credential Connection Registration Status
[**CreateCredentialConnection**](CredentialConnectionsAPI.md#CreateCredentialConnection) | **Post** /credential_connections | Create a credential connection
[**DeleteCredentialConnection**](CredentialConnectionsAPI.md#DeleteCredentialConnection) | **Delete** /credential_connections/{id} | Delete a credential connection
[**ListCredentialConnections**](CredentialConnectionsAPI.md#ListCredentialConnections) | **Get** /credential_connections | List credential connections
[**RetrieveCredentialConnection**](CredentialConnectionsAPI.md#RetrieveCredentialConnection) | **Get** /credential_connections/{id} | Retrieve a credential connection
[**UpdateCredentialConnection**](CredentialConnectionsAPI.md#UpdateCredentialConnection) | **Patch** /credential_connections/{id} | Update a credential connection



## CheckRegistrationStatus

> RegistrationStatusResponse CheckRegistrationStatus(ctx, id).Execute()

Check a Credential Connection Registration Status



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
	id := "id_example" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialConnectionsAPI.CheckRegistrationStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialConnectionsAPI.CheckRegistrationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRegistrationStatus`: RegistrationStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialConnectionsAPI.CheckRegistrationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRegistrationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistrationStatusResponse**](RegistrationStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredentialConnection

> CredentialConnectionResponse CreateCredentialConnection(ctx).CreateCredentialConnectionRequest(createCredentialConnectionRequest).Execute()

Create a credential connection



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
	createCredentialConnectionRequest := *openapiclient.NewCreateCredentialConnectionRequest("myusername123", "my123secure456password789", "office-connection") // CreateCredentialConnectionRequest | Parameters that can be defined during credential connection creation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialConnectionsAPI.CreateCredentialConnection(context.Background()).CreateCredentialConnectionRequest(createCredentialConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialConnectionsAPI.CreateCredentialConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredentialConnection`: CredentialConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialConnectionsAPI.CreateCredentialConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCredentialConnectionRequest** | [**CreateCredentialConnectionRequest**](CreateCredentialConnectionRequest.md) | Parameters that can be defined during credential connection creation | 

### Return type

[**CredentialConnectionResponse**](CredentialConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialConnection

> CredentialConnectionResponse DeleteCredentialConnection(ctx, id).Execute()

Delete a credential connection



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
	id := "id_example" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialConnectionsAPI.DeleteCredentialConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialConnectionsAPI.DeleteCredentialConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCredentialConnection`: CredentialConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialConnectionsAPI.DeleteCredentialConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialConnectionResponse**](CredentialConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialConnections

> ListCredentialConnectionsResponse ListCredentialConnections(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterOutboundOutboundVoiceProfileId(filterOutboundOutboundVoiceProfileId).Sort(sort).Execute()

List credential connections



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
	resp, r, err := apiClient.CredentialConnectionsAPI.ListCredentialConnections(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterOutboundOutboundVoiceProfileId(filterOutboundOutboundVoiceProfileId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialConnectionsAPI.ListCredentialConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCredentialConnections`: ListCredentialConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialConnectionsAPI.ListCredentialConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterConnectionNameContains** | **string** | If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | 
 **filterOutboundOutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**ListCredentialConnectionsResponse**](ListCredentialConnectionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCredentialConnection

> CredentialConnectionResponse RetrieveCredentialConnection(ctx, id).Execute()

Retrieve a credential connection



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
	id := "id_example" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialConnectionsAPI.RetrieveCredentialConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialConnectionsAPI.RetrieveCredentialConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCredentialConnection`: CredentialConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialConnectionsAPI.RetrieveCredentialConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCredentialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialConnectionResponse**](CredentialConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialConnection

> CredentialConnectionResponse UpdateCredentialConnection(ctx, id).UpdateCredentialConnectionRequest(updateCredentialConnectionRequest).Execute()

Update a credential connection



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
	id := "id_example" // string | Identifies the resource.
	updateCredentialConnectionRequest := *openapiclient.NewUpdateCredentialConnectionRequest() // UpdateCredentialConnectionRequest | Parameters that can be updated in a credential connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialConnectionsAPI.UpdateCredentialConnection(context.Background(), id).UpdateCredentialConnectionRequest(updateCredentialConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialConnectionsAPI.UpdateCredentialConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredentialConnection`: CredentialConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialConnectionsAPI.UpdateCredentialConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCredentialConnectionRequest** | [**UpdateCredentialConnectionRequest**](UpdateCredentialConnectionRequest.md) | Parameters that can be updated in a credential connection | 

### Return type

[**CredentialConnectionResponse**](CredentialConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

