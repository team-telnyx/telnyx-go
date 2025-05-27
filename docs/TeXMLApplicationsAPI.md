# \TeXMLApplicationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTexmlApplication**](TeXMLApplicationsAPI.md#CreateTexmlApplication) | **Post** /texml_applications | Creates a TeXML Application
[**DeleteTexmlApplication**](TeXMLApplicationsAPI.md#DeleteTexmlApplication) | **Delete** /texml_applications/{id} | Deletes a TeXML Application
[**FindTexmlApplications**](TeXMLApplicationsAPI.md#FindTexmlApplications) | **Get** /texml_applications | List all TeXML Applications
[**GetTexmlApplication**](TeXMLApplicationsAPI.md#GetTexmlApplication) | **Get** /texml_applications/{id} | Retrieve a TeXML Application
[**UpdateTexmlApplication**](TeXMLApplicationsAPI.md#UpdateTexmlApplication) | **Patch** /texml_applications/{id} | Update a TeXML Application



## CreateTexmlApplication

> TexmlApplicationResponse CreateTexmlApplication(ctx).CreateTexmlApplicationRequest(createTexmlApplicationRequest).Execute()

Creates a TeXML Application



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
	createTexmlApplicationRequest := *openapiclient.NewCreateTexmlApplicationRequest("call-router", "https://example.com") // CreateTexmlApplicationRequest | Parameters that can be set when creating a TeXML Application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLApplicationsAPI.CreateTexmlApplication(context.Background()).CreateTexmlApplicationRequest(createTexmlApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLApplicationsAPI.CreateTexmlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTexmlApplication`: TexmlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLApplicationsAPI.CreateTexmlApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTexmlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTexmlApplicationRequest** | [**CreateTexmlApplicationRequest**](CreateTexmlApplicationRequest.md) | Parameters that can be set when creating a TeXML Application | 

### Return type

[**TexmlApplicationResponse**](TexmlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTexmlApplication

> TexmlApplicationResponse DeleteTexmlApplication(ctx, id).Execute()

Deletes a TeXML Application



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
	resp, r, err := apiClient.TeXMLApplicationsAPI.DeleteTexmlApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLApplicationsAPI.DeleteTexmlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTexmlApplication`: TexmlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLApplicationsAPI.DeleteTexmlApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTexmlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TexmlApplicationResponse**](TexmlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTexmlApplications

> GetAllTexmlApplicationsResponse FindTexmlApplications(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterFriendlyNameContains(filterFriendlyNameContains).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()

List all TeXML Applications



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
	pageSize := int32(56) // int32 | The size of the page (optional) (default to 250)
	filterFriendlyNameContains := "filterFriendlyNameContains_example" // string | If present, applications with <code>friendly_name</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional) (default to "null")
	filterOutboundVoiceProfileId := "1293384261075731499" // string | Identifies the associated outbound voice profile. (optional)
	sort := "friendly_name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>friendly_name</code>: sorts the result by the     <code>friendly_name</code> field in ascending order.   </li>    <li>     <code>-friendly_name</code>: sorts the result by the     <code>friendly_name</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLApplicationsAPI.FindTexmlApplications(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterFriendlyNameContains(filterFriendlyNameContains).FilterOutboundVoiceProfileId(filterOutboundVoiceProfileId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLApplicationsAPI.FindTexmlApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindTexmlApplications`: GetAllTexmlApplicationsResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLApplicationsAPI.FindTexmlApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindTexmlApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page | [default to 250]
 **filterFriendlyNameContains** | **string** | If present, applications with &lt;code&gt;friendly_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to &quot;null&quot;]
 **filterOutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;friendly_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;friendly_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-friendly_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;friendly_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**GetAllTexmlApplicationsResponse**](GetAllTexmlApplicationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlApplication

> TexmlApplicationResponse GetTexmlApplication(ctx, id).Execute()

Retrieve a TeXML Application



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
	resp, r, err := apiClient.TeXMLApplicationsAPI.GetTexmlApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLApplicationsAPI.GetTexmlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlApplication`: TexmlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLApplicationsAPI.GetTexmlApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TexmlApplicationResponse**](TexmlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTexmlApplication

> TexmlApplicationResponse UpdateTexmlApplication(ctx, id).UpdateTexmlApplicationRequest(updateTexmlApplicationRequest).Execute()

Update a TeXML Application



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
	updateTexmlApplicationRequest := *openapiclient.NewUpdateTexmlApplicationRequest("call-router", "https://example.com") // UpdateTexmlApplicationRequest | Parameters that can be updated in a TeXML Application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLApplicationsAPI.UpdateTexmlApplication(context.Background(), id).UpdateTexmlApplicationRequest(updateTexmlApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLApplicationsAPI.UpdateTexmlApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTexmlApplication`: TexmlApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLApplicationsAPI.UpdateTexmlApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTexmlApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTexmlApplicationRequest** | [**UpdateTexmlApplicationRequest**](UpdateTexmlApplicationRequest.md) | Parameters that can be updated in a TeXML Application | 

### Return type

[**TexmlApplicationResponse**](TexmlApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

