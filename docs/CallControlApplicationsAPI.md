# \CallControlApplicationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallControlApplication**](CallControlApplicationsAPI.md#CreateCallControlApplication) | **Post** /call_control_applications | Create a call control application
[**DeleteCallControlApplication**](CallControlApplicationsAPI.md#DeleteCallControlApplication) | **Delete** /call_control_applications/{id} | Delete a call control application
[**ListCallControlApplications**](CallControlApplicationsAPI.md#ListCallControlApplications) | **Get** /call_control_applications | List call control applications
[**RetrieveCallControlApplication**](CallControlApplicationsAPI.md#RetrieveCallControlApplication) | **Get** /call_control_applications/{id} | Retrieve a call control application
[**UpdateCallControlApplication**](CallControlApplicationsAPI.md#UpdateCallControlApplication) | **Patch** /call_control_applications/{id} | Update a call control application



## CreateCallControlApplication

> CallControlApplicationResponse CreateCallControlApplication(ctx).CreateCallControlApplicationRequest(createCallControlApplicationRequest).Execute()

Create a call control application



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
	createCallControlApplicationRequest := *openapiclient.NewCreateCallControlApplicationRequest("call-router", "https://example.com") // CreateCallControlApplicationRequest | Create call control application request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallControlApplicationsAPI.CreateCallControlApplication(context.Background()).CreateCallControlApplicationRequest(createCallControlApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallControlApplicationsAPI.CreateCallControlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCallControlApplication`: CallControlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `CallControlApplicationsAPI.CreateCallControlApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallControlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCallControlApplicationRequest** | [**CreateCallControlApplicationRequest**](CreateCallControlApplicationRequest.md) | Create call control application request. | 

### Return type

[**CallControlApplicationResponse**](CallControlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallControlApplication

> CallControlApplicationResponse DeleteCallControlApplication(ctx, id).Execute()

Delete a call control application



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
	resp, r, err := apiClient.CallControlApplicationsAPI.DeleteCallControlApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallControlApplicationsAPI.DeleteCallControlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCallControlApplication`: CallControlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `CallControlApplicationsAPI.DeleteCallControlApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallControlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CallControlApplicationResponse**](CallControlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallControlApplications

> ListCallControlApplicationsResponse ListCallControlApplications(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterApplicationNameContains(filterApplicationNameContains).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()

List call control applications



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
	filterApplicationNameContains := "fax-app" // string | If present, applications with <code>application_name</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional)
	filterOutboundVoiceProfileId := "1293384261075731499" // string | Identifies the associated outbound voice profile. (optional)
	sort := "connection_name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>connection_name</code>: sorts the result by the     <code>connection_name</code> field in ascending order.   </li>    <li>     <code>-connection_name</code>: sorts the result by the     <code>connection_name</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallControlApplicationsAPI.ListCallControlApplications(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterApplicationNameContains(filterApplicationNameContains).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallControlApplicationsAPI.ListCallControlApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCallControlApplications`: ListCallControlApplicationsResponse
	fmt.Fprintf(os.Stdout, "Response from `CallControlApplicationsAPI.ListCallControlApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCallControlApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterApplicationNameContains** | **string** | If present, applications with &lt;code&gt;application_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | 
 **filterOutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**ListCallControlApplicationsResponse**](ListCallControlApplicationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCallControlApplication

> CallControlApplicationResponse RetrieveCallControlApplication(ctx, id).Execute()

Retrieve a call control application



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
	resp, r, err := apiClient.CallControlApplicationsAPI.RetrieveCallControlApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallControlApplicationsAPI.RetrieveCallControlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCallControlApplication`: CallControlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `CallControlApplicationsAPI.RetrieveCallControlApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCallControlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CallControlApplicationResponse**](CallControlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallControlApplication

> CallControlApplicationResponse UpdateCallControlApplication(ctx, id).UpdateCallControlApplicationRequest(updateCallControlApplicationRequest).Execute()

Update a call control application



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
	updateCallControlApplicationRequest := *openapiclient.NewUpdateCallControlApplicationRequest("call-router", "https://example.com") // UpdateCallControlApplicationRequest | Update call control application request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallControlApplicationsAPI.UpdateCallControlApplication(context.Background(), id).UpdateCallControlApplicationRequest(updateCallControlApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallControlApplicationsAPI.UpdateCallControlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCallControlApplication`: CallControlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `CallControlApplicationsAPI.UpdateCallControlApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallControlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCallControlApplicationRequest** | [**UpdateCallControlApplicationRequest**](UpdateCallControlApplicationRequest.md) | Update call control application request. | 

### Return type

[**CallControlApplicationResponse**](CallControlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

