# \ManagedAccountsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedAccount**](ManagedAccountsAPI.md#CreateManagedAccount) | **Post** /managed_accounts | Create a new managed account.
[**DisableManagedAccount**](ManagedAccountsAPI.md#DisableManagedAccount) | **Post** /managed_accounts/{id}/actions/disable | Disables a managed account
[**EnableManagedAccount**](ManagedAccountsAPI.md#EnableManagedAccount) | **Post** /managed_accounts/{id}/actions/enable | Enables a managed account
[**ListAllocatableGlobalOutboundChannels**](ManagedAccountsAPI.md#ListAllocatableGlobalOutboundChannels) | **Get** /managed_accounts/allocatable_global_outbound_channels | Display information about allocatable global outbound channels for the current user.
[**ListManagedAccounts**](ManagedAccountsAPI.md#ListManagedAccounts) | **Get** /managed_accounts | Lists accounts managed by the current user.
[**RetrieveManagedAccount**](ManagedAccountsAPI.md#RetrieveManagedAccount) | **Get** /managed_accounts/{id} | Retrieve a managed account
[**UpdateManagedAccount**](ManagedAccountsAPI.md#UpdateManagedAccount) | **Patch** /managed_accounts/{id} | Update a managed account
[**UpdateManagedAccountGlobalChannelLimit**](ManagedAccountsAPI.md#UpdateManagedAccountGlobalChannelLimit) | **Patch** /managed_accounts/{id}/update_global_channel_limit | Update the amount of allocatable global outbound channels allocated to a specific managed account.



## CreateManagedAccount

> CreateManagedAccount200Response CreateManagedAccount(ctx).CreateManagedAccountRequest(createManagedAccountRequest).Execute()

Create a new managed account.



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
	createManagedAccountRequest := *openapiclient.NewCreateManagedAccountRequest("Larry's Cat Food Inc") // CreateManagedAccountRequest | Parameters that define the managed account to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.CreateManagedAccount(context.Background()).CreateManagedAccountRequest(createManagedAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.CreateManagedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManagedAccount`: CreateManagedAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.CreateManagedAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createManagedAccountRequest** | [**CreateManagedAccountRequest**](CreateManagedAccountRequest.md) | Parameters that define the managed account to be created | 

### Return type

[**CreateManagedAccount200Response**](CreateManagedAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableManagedAccount

> CreateManagedAccount200Response DisableManagedAccount(ctx, id).Execute()

Disables a managed account



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
	id := "id_example" // string | Managed Account User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.DisableManagedAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.DisableManagedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableManagedAccount`: CreateManagedAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.DisableManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Account User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateManagedAccount200Response**](CreateManagedAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableManagedAccount

> CreateManagedAccount200Response EnableManagedAccount(ctx, id).EnableManagedAccountRequest(enableManagedAccountRequest).Execute()

Enables a managed account



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
	id := "id_example" // string | Managed Account User ID
	enableManagedAccountRequest := *openapiclient.NewEnableManagedAccountRequest() // EnableManagedAccountRequest | Additional parameters for what to do when enabling the managed account (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.EnableManagedAccount(context.Background(), id).EnableManagedAccountRequest(enableManagedAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.EnableManagedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableManagedAccount`: CreateManagedAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.EnableManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Account User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableManagedAccountRequest** | [**EnableManagedAccountRequest**](EnableManagedAccountRequest.md) | Additional parameters for what to do when enabling the managed account | 

### Return type

[**CreateManagedAccount200Response**](CreateManagedAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllocatableGlobalOutboundChannels

> ListAllocatableGlobalOutboundChannels200Response ListAllocatableGlobalOutboundChannels(ctx).Execute()

Display information about allocatable global outbound channels for the current user.



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
	resp, r, err := apiClient.ManagedAccountsAPI.ListAllocatableGlobalOutboundChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.ListAllocatableGlobalOutboundChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllocatableGlobalOutboundChannels`: ListAllocatableGlobalOutboundChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.ListAllocatableGlobalOutboundChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllocatableGlobalOutboundChannelsRequest struct via the builder pattern


### Return type

[**ListAllocatableGlobalOutboundChannels200Response**](ListAllocatableGlobalOutboundChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagedAccounts

> ListManagedAccounts200Response ListManagedAccounts(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterEmailContains(filterEmailContains).FilterEmailEq(filterEmailEq).FilterOrganizationNameContains(filterOrganizationNameContains).FilterOrganizationNameEq(filterOrganizationNameEq).Sort(sort).IncludeCancelledAccounts(includeCancelledAccounts).Execute()

Lists accounts managed by the current user.



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
	filterEmailContains := "filterEmailContains_example" // string | If present, email containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional) (default to "null")
	filterEmailEq := "filterEmailEq_example" // string | If present, only returns results with the <code>email</code> matching exactly the value given. (optional) (default to "null")
	filterOrganizationNameContains := "filterOrganizationNameContains_example" // string | If present, only returns results with the <code>organization_name</code> containing the given value. Matching is not case-sensitive. Requires at least three characters. (optional) (default to "null")
	filterOrganizationNameEq := "filterOrganizationNameEq_example" // string | If present, only returns results with the <code>organization_name</code> matching exactly the value given. (optional) (default to "null")
	sort := "email" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>email</code>: sorts the result by the     <code>email</code> field in ascending order.   </li>    <li>     <code>-email</code>: sorts the result by the     <code>email</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")
	includeCancelledAccounts := true // bool | Specifies if cancelled accounts should be included in the results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.ListManagedAccounts(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterEmailContains(filterEmailContains).FilterEmailEq(filterEmailEq).FilterOrganizationNameContains(filterOrganizationNameContains).FilterOrganizationNameEq(filterOrganizationNameEq).Sort(sort).IncludeCancelledAccounts(includeCancelledAccounts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.ListManagedAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListManagedAccounts`: ListManagedAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.ListManagedAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListManagedAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterEmailContains** | **string** | If present, email containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to &quot;null&quot;]
 **filterEmailEq** | **string** | If present, only returns results with the &lt;code&gt;email&lt;/code&gt; matching exactly the value given. | [default to &quot;null&quot;]
 **filterOrganizationNameContains** | **string** | If present, only returns results with the &lt;code&gt;organization_name&lt;/code&gt; containing the given value. Matching is not case-sensitive. Requires at least three characters. | [default to &quot;null&quot;]
 **filterOrganizationNameEq** | **string** | If present, only returns results with the &lt;code&gt;organization_name&lt;/code&gt; matching exactly the value given. | [default to &quot;null&quot;]
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;email&lt;/code&gt;: sorts the result by the     &lt;code&gt;email&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-email&lt;/code&gt;: sorts the result by the     &lt;code&gt;email&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]
 **includeCancelledAccounts** | **bool** | Specifies if cancelled accounts should be included in the results. | [default to false]

### Return type

[**ListManagedAccounts200Response**](ListManagedAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveManagedAccount

> CreateManagedAccount200Response RetrieveManagedAccount(ctx, id).Execute()

Retrieve a managed account



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
	id := "id_example" // string | Managed Account User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.RetrieveManagedAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.RetrieveManagedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveManagedAccount`: CreateManagedAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.RetrieveManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Account User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateManagedAccount200Response**](CreateManagedAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedAccount

> CreateManagedAccount200Response UpdateManagedAccount(ctx, id).UpdateManagedAccountRequest(updateManagedAccountRequest).Execute()

Update a managed account



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
	id := "id_example" // string | Managed Account User ID
	updateManagedAccountRequest := *openapiclient.NewUpdateManagedAccountRequest() // UpdateManagedAccountRequest | Parameters that define the updates to the managed account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.UpdateManagedAccount(context.Background(), id).UpdateManagedAccountRequest(updateManagedAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.UpdateManagedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateManagedAccount`: CreateManagedAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.UpdateManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Account User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateManagedAccountRequest** | [**UpdateManagedAccountRequest**](UpdateManagedAccountRequest.md) | Parameters that define the updates to the managed account | 

### Return type

[**CreateManagedAccount200Response**](CreateManagedAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedAccountGlobalChannelLimit

> UpdateManagedAccountGlobalChannelLimit200Response UpdateManagedAccountGlobalChannelLimit(ctx, id).UpdateManagedAccountGlobalChannelLimitRequest(updateManagedAccountGlobalChannelLimitRequest).Execute()

Update the amount of allocatable global outbound channels allocated to a specific managed account.



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
	id := "id_example" // string | Managed Account User ID
	updateManagedAccountGlobalChannelLimitRequest := *openapiclient.NewUpdateManagedAccountGlobalChannelLimitRequest() // UpdateManagedAccountGlobalChannelLimitRequest | Parameters that define the changes to the global outbounds channels for the managed account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAccountsAPI.UpdateManagedAccountGlobalChannelLimit(context.Background(), id).UpdateManagedAccountGlobalChannelLimitRequest(updateManagedAccountGlobalChannelLimitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.UpdateManagedAccountGlobalChannelLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateManagedAccountGlobalChannelLimit`: UpdateManagedAccountGlobalChannelLimit200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.UpdateManagedAccountGlobalChannelLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Account User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedAccountGlobalChannelLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateManagedAccountGlobalChannelLimitRequest** | [**UpdateManagedAccountGlobalChannelLimitRequest**](UpdateManagedAccountGlobalChannelLimitRequest.md) | Parameters that define the changes to the global outbounds channels for the managed account | 

### Return type

[**UpdateManagedAccountGlobalChannelLimit200Response**](UpdateManagedAccountGlobalChannelLimit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

