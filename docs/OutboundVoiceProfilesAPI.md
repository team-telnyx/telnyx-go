# \OutboundVoiceProfilesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoiceProfile**](OutboundVoiceProfilesAPI.md#CreateVoiceProfile) | **Post** /outbound_voice_profiles | Create an outbound voice profile
[**DeleteOutboundVoiceProfile**](OutboundVoiceProfilesAPI.md#DeleteOutboundVoiceProfile) | **Delete** /outbound_voice_profiles/{id} | Delete an outbound voice profile
[**GetOutboundVoiceProfile**](OutboundVoiceProfilesAPI.md#GetOutboundVoiceProfile) | **Get** /outbound_voice_profiles/{id} | Retrieve an outbound voice profile
[**ListOutboundVoiceProfiles**](OutboundVoiceProfilesAPI.md#ListOutboundVoiceProfiles) | **Get** /outbound_voice_profiles | Get all outbound voice profiles
[**UpdateOutboundVoiceProfile**](OutboundVoiceProfilesAPI.md#UpdateOutboundVoiceProfile) | **Patch** /outbound_voice_profiles/{id} | Updates an existing outbound voice profile.



## CreateVoiceProfile

> OutboundVoiceProfileResponse CreateVoiceProfile(ctx).CreateOutboundVoiceProfileRequest(createOutboundVoiceProfileRequest).Execute()

Create an outbound voice profile



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
	createOutboundVoiceProfileRequest := *openapiclient.NewCreateOutboundVoiceProfileRequest("office") // CreateOutboundVoiceProfileRequest | Parameters that can be defined when creating an outbound voice profile

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboundVoiceProfilesAPI.CreateVoiceProfile(context.Background()).CreateOutboundVoiceProfileRequest(createOutboundVoiceProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundVoiceProfilesAPI.CreateVoiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoiceProfile`: OutboundVoiceProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundVoiceProfilesAPI.CreateVoiceProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOutboundVoiceProfileRequest** | [**CreateOutboundVoiceProfileRequest**](CreateOutboundVoiceProfileRequest.md) | Parameters that can be defined when creating an outbound voice profile | 

### Return type

[**OutboundVoiceProfileResponse**](OutboundVoiceProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutboundVoiceProfile

> OutboundVoiceProfileResponse DeleteOutboundVoiceProfile(ctx, id).Execute()

Delete an outbound voice profile



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
	resp, r, err := apiClient.OutboundVoiceProfilesAPI.DeleteOutboundVoiceProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundVoiceProfilesAPI.DeleteOutboundVoiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOutboundVoiceProfile`: OutboundVoiceProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundVoiceProfilesAPI.DeleteOutboundVoiceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutboundVoiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutboundVoiceProfileResponse**](OutboundVoiceProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboundVoiceProfile

> OutboundVoiceProfileResponse GetOutboundVoiceProfile(ctx, id).Execute()

Retrieve an outbound voice profile



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
	resp, r, err := apiClient.OutboundVoiceProfilesAPI.GetOutboundVoiceProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundVoiceProfilesAPI.GetOutboundVoiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboundVoiceProfile`: OutboundVoiceProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundVoiceProfilesAPI.GetOutboundVoiceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboundVoiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutboundVoiceProfileResponse**](OutboundVoiceProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutboundVoiceProfiles

> ListOutboundVoiceProfilesResponse ListOutboundVoiceProfiles(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterNameContains(filterNameContains).Sort(sort).Execute()

Get all outbound voice profiles



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
	filterNameContains := "office-profile" // string | Optional filter on outbound voice profile name. (optional)
	sort := "name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code>-</code> prefix.<br/><br/> That is: <ul>   <li>     <code>name</code>: sorts the result by the     <code>name</code> field in ascending order.   </li>    <li>     <code>-name</code>: sorts the result by the     <code>name</code> field in descending order.   </li> </ul> <br/> (optional) (default to "-created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboundVoiceProfilesAPI.ListOutboundVoiceProfiles(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterNameContains(filterNameContains).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundVoiceProfilesAPI.ListOutboundVoiceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOutboundVoiceProfiles`: ListOutboundVoiceProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundVoiceProfilesAPI.ListOutboundVoiceProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOutboundVoiceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterNameContains** | **string** | Optional filter on outbound voice profile name. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt;-&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;name&lt;/code&gt;: sorts the result by the     &lt;code&gt;name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-name&lt;/code&gt;: sorts the result by the     &lt;code&gt;name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; | [default to &quot;-created_at&quot;]

### Return type

[**ListOutboundVoiceProfilesResponse**](ListOutboundVoiceProfilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutboundVoiceProfile

> OutboundVoiceProfileResponse UpdateOutboundVoiceProfile(ctx, id).UpdateOutboundVoiceProfileRequest(updateOutboundVoiceProfileRequest).Execute()

Updates an existing outbound voice profile.



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
	updateOutboundVoiceProfileRequest := *openapiclient.NewUpdateOutboundVoiceProfileRequest("office") // UpdateOutboundVoiceProfileRequest | Parameters that can be updated on an outbound voice profile

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboundVoiceProfilesAPI.UpdateOutboundVoiceProfile(context.Background(), id).UpdateOutboundVoiceProfileRequest(updateOutboundVoiceProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundVoiceProfilesAPI.UpdateOutboundVoiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOutboundVoiceProfile`: OutboundVoiceProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundVoiceProfilesAPI.UpdateOutboundVoiceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutboundVoiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOutboundVoiceProfileRequest** | [**UpdateOutboundVoiceProfileRequest**](UpdateOutboundVoiceProfileRequest.md) | Parameters that can be updated on an outbound voice profile | 

### Return type

[**OutboundVoiceProfileResponse**](OutboundVoiceProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

