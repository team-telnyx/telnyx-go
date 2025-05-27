# \ProgrammableFaxApplicationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFaxApplication**](ProgrammableFaxApplicationsAPI.md#CreateFaxApplication) | **Post** /fax_applications | Creates a Fax Application
[**DeleteFaxApplication**](ProgrammableFaxApplicationsAPI.md#DeleteFaxApplication) | **Delete** /fax_applications/{id} | Deletes a Fax Application
[**GetFaxApplication**](ProgrammableFaxApplicationsAPI.md#GetFaxApplication) | **Get** /fax_applications/{id} | Retrieve a Fax Application
[**ListFaxApplications**](ProgrammableFaxApplicationsAPI.md#ListFaxApplications) | **Get** /fax_applications | List all Fax Applications
[**UpdateFaxApplication**](ProgrammableFaxApplicationsAPI.md#UpdateFaxApplication) | **Patch** /fax_applications/{id} | Update a Fax Application



## CreateFaxApplication

> FaxApplicationResponse CreateFaxApplication(ctx).CreateFaxApplicationRequest(createFaxApplicationRequest).Execute()

Creates a Fax Application



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
	createFaxApplicationRequest := *openapiclient.NewCreateFaxApplicationRequest("call-router", "https://example.com") // CreateFaxApplicationRequest | Parameters that can be set when creating a Fax Application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxApplicationsAPI.CreateFaxApplication(context.Background()).CreateFaxApplicationRequest(createFaxApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxApplicationsAPI.CreateFaxApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFaxApplication`: FaxApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxApplicationsAPI.CreateFaxApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFaxApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFaxApplicationRequest** | [**CreateFaxApplicationRequest**](CreateFaxApplicationRequest.md) | Parameters that can be set when creating a Fax Application | 

### Return type

[**FaxApplicationResponse**](FaxApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFaxApplication

> FaxApplicationResponse DeleteFaxApplication(ctx, id).Execute()

Deletes a Fax Application



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
	resp, r, err := apiClient.ProgrammableFaxApplicationsAPI.DeleteFaxApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxApplicationsAPI.DeleteFaxApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFaxApplication`: FaxApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxApplicationsAPI.DeleteFaxApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFaxApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FaxApplicationResponse**](FaxApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFaxApplication

> FaxApplicationResponse GetFaxApplication(ctx, id).Execute()

Retrieve a Fax Application



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
	resp, r, err := apiClient.ProgrammableFaxApplicationsAPI.GetFaxApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxApplicationsAPI.GetFaxApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFaxApplication`: FaxApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxApplicationsAPI.GetFaxApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFaxApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FaxApplicationResponse**](FaxApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFaxApplications

> GetAllFaxApplicationsResponse ListFaxApplications(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterApplicationNameContains(filterApplicationNameContains).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()

List all Fax Applications



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
	sort := "friendly_name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>friendly_name</code>: sorts the result by the     <code>friendly_name</code> field in ascending order.   </li>    <li>     <code>-friendly_name</code>: sorts the result by the     <code>friendly_name</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxApplicationsAPI.ListFaxApplications(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterApplicationNameContains(filterApplicationNameContains).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxApplicationsAPI.ListFaxApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFaxApplications`: GetAllFaxApplicationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxApplicationsAPI.ListFaxApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFaxApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterApplicationNameContains** | **string** | If present, applications with &lt;code&gt;application_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | 
 **filterOutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;friendly_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;friendly_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-friendly_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;friendly_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**GetAllFaxApplicationsResponse**](GetAllFaxApplicationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFaxApplication

> FaxApplicationResponse UpdateFaxApplication(ctx, id).UpdateFaxApplicationRequest(updateFaxApplicationRequest).Execute()

Update a Fax Application



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
	updateFaxApplicationRequest := *openapiclient.NewUpdateFaxApplicationRequest("call-router", "https://example.com") // UpdateFaxApplicationRequest | Parameters to be updated for the Fax Application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxApplicationsAPI.UpdateFaxApplication(context.Background(), id).UpdateFaxApplicationRequest(updateFaxApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxApplicationsAPI.UpdateFaxApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFaxApplication`: FaxApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxApplicationsAPI.UpdateFaxApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFaxApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFaxApplicationRequest** | [**UpdateFaxApplicationRequest**](UpdateFaxApplicationRequest.md) | Parameters to be updated for the Fax Application | 

### Return type

[**FaxApplicationResponse**](FaxApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

