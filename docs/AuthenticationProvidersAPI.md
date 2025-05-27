# \AuthenticationProvidersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationProvider**](AuthenticationProvidersAPI.md#CreateAuthenticationProvider) | **Post** /authentication_providers | Creates an authentication provider
[**DeleteAuthenticationProvider**](AuthenticationProvidersAPI.md#DeleteAuthenticationProvider) | **Delete** /authentication_providers/{id} | Deletes an authentication provider
[**FindAuthenticationProviders**](AuthenticationProvidersAPI.md#FindAuthenticationProviders) | **Get** /authentication_providers | List all SSO authentication providers
[**GetAuthenticationProvider**](AuthenticationProvidersAPI.md#GetAuthenticationProvider) | **Get** /authentication_providers/{id} | Retrieve an authentication provider
[**UpdateAuthenticationProvider**](AuthenticationProvidersAPI.md#UpdateAuthenticationProvider) | **Patch** /authentication_providers/{id} | Update an authentication provider



## CreateAuthenticationProvider

> CreateAuthenticationProvider200Response CreateAuthenticationProvider(ctx).AuthenticationProviderCreate(authenticationProviderCreate).Execute()

Creates an authentication provider



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
	authenticationProviderCreate := *openapiclient.NewAuthenticationProviderCreate("Okta", "myorg", *openapiclient.NewSettings("https://myorg.myidp.com/saml/metadata", "https://myorg.myidp.com/trust/saml2/http-post/sso", "13:38:C7:BB:C9:FF:4A:70:38:3A:E3:D9:5C:CD:DB:2E:50:1E:80:A7")) // AuthenticationProviderCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationProvidersAPI.CreateAuthenticationProvider(context.Background()).AuthenticationProviderCreate(authenticationProviderCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationProvidersAPI.CreateAuthenticationProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthenticationProvider`: CreateAuthenticationProvider200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationProvidersAPI.CreateAuthenticationProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationProviderCreate** | [**AuthenticationProviderCreate**](AuthenticationProviderCreate.md) |  | 

### Return type

[**CreateAuthenticationProvider200Response**](CreateAuthenticationProvider200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticationProvider

> CreateAuthenticationProvider200Response DeleteAuthenticationProvider(ctx, id).Execute()

Deletes an authentication provider



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
	id := "id_example" // string | authentication provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationProvidersAPI.DeleteAuthenticationProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationProvidersAPI.DeleteAuthenticationProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAuthenticationProvider`: CreateAuthenticationProvider200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationProvidersAPI.DeleteAuthenticationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | authentication provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAuthenticationProvider200Response**](CreateAuthenticationProvider200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAuthenticationProviders

> FindAuthenticationProviders200Response FindAuthenticationProviders(ctx).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()

List all SSO authentication providers



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
	pageNumber := int32(56) // int32 | The page number to load (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page (optional) (default to 20)
	sort := "name" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code>-</code> prefix.<br/><br/> That is: <ul>   <li>     <code>name</code>: sorts the result by the     <code>name</code> field in ascending order.   </li>   <li>     <code>-name</code>: sorts the result by the     <code>name</code> field in descending order.   </li> </ul><br/>If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "-created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationProvidersAPI.FindAuthenticationProviders(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationProvidersAPI.FindAuthenticationProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAuthenticationProviders`: FindAuthenticationProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationProvidersAPI.FindAuthenticationProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAuthenticationProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load | [default to 1]
 **pageSize** | **int32** | The size of the page | [default to 20]
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt;-&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;name&lt;/code&gt;: sorts the result by the     &lt;code&gt;name&lt;/code&gt; field in ascending order.   &lt;/li&gt;   &lt;li&gt;     &lt;code&gt;-name&lt;/code&gt;: sorts the result by the     &lt;code&gt;name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt;&lt;br/&gt;If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;-created_at&quot;]

### Return type

[**FindAuthenticationProviders200Response**](FindAuthenticationProviders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationProvider

> CreateAuthenticationProvider200Response GetAuthenticationProvider(ctx, id).Execute()

Retrieve an authentication provider



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
	id := "id_example" // string | authentication provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationProvidersAPI.GetAuthenticationProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationProvidersAPI.GetAuthenticationProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticationProvider`: CreateAuthenticationProvider200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationProvidersAPI.GetAuthenticationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | authentication provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAuthenticationProvider200Response**](CreateAuthenticationProvider200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationProvider

> CreateAuthenticationProvider200Response UpdateAuthenticationProvider(ctx, id).UpdateAuthenticationProviderRequest(updateAuthenticationProviderRequest).Execute()

Update an authentication provider



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
	updateAuthenticationProviderRequest := *openapiclient.NewUpdateAuthenticationProviderRequest() // UpdateAuthenticationProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationProvidersAPI.UpdateAuthenticationProvider(context.Background(), id).UpdateAuthenticationProviderRequest(updateAuthenticationProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationProvidersAPI.UpdateAuthenticationProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAuthenticationProvider`: CreateAuthenticationProvider200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationProvidersAPI.UpdateAuthenticationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAuthenticationProviderRequest** | [**UpdateAuthenticationProviderRequest**](UpdateAuthenticationProviderRequest.md) |  | 

### Return type

[**CreateAuthenticationProvider200Response**](CreateAuthenticationProvider200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

