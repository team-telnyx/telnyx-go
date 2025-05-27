# \UserBundlesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserBundlesBulk**](UserBundlesAPI.md#CreateUserBundlesBulk) | **Post** /bundle_pricing/user_bundles/bulk | Create User Bundles
[**DeactivateUserBundle**](UserBundlesAPI.md#DeactivateUserBundle) | **Delete** /bundle_pricing/user_bundles/{user_bundle_id} | Deactivate User Bundle
[**GetUnusedUserBundles**](UserBundlesAPI.md#GetUnusedUserBundles) | **Get** /bundle_pricing/user_bundles/unused | Get Unused User Bundles
[**GetUserBundleById**](UserBundlesAPI.md#GetUserBundleById) | **Get** /bundle_pricing/user_bundles/{user_bundle_id} | Get User Bundle by Id
[**GetUserBundleResources**](UserBundlesAPI.md#GetUserBundleResources) | **Get** /bundle_pricing/user_bundles/{user_bundle_id}/resources | Get User Bundle Resources
[**GetUserBundles**](UserBundlesAPI.md#GetUserBundles) | **Get** /bundle_pricing/user_bundles | Get User Bundles



## CreateUserBundlesBulk

> CreatedUserBundlesResponse CreateUserBundlesBulk(ctx).CreateUserBundlesBulkRequest(createUserBundlesBulkRequest).AuthorizationBearer(authorizationBearer).Execute()

Create User Bundles



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
	createUserBundlesBulkRequest := *openapiclient.NewCreateUserBundlesBulkRequest() // CreateUserBundlesBulkRequest | 
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserBundlesAPI.CreateUserBundlesBulk(context.Background()).CreateUserBundlesBulkRequest(createUserBundlesBulkRequest).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserBundlesAPI.CreateUserBundlesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserBundlesBulk`: CreatedUserBundlesResponse
	fmt.Fprintf(os.Stdout, "Response from `UserBundlesAPI.CreateUserBundlesBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserBundlesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserBundlesBulkRequest** | [**CreateUserBundlesBulkRequest**](CreateUserBundlesBulkRequest.md) |  | 
 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**CreatedUserBundlesResponse**](CreatedUserBundlesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateUserBundle

> UserBundleCreateResponse DeactivateUserBundle(ctx, userBundleId).AuthorizationBearer(authorizationBearer).Execute()

Deactivate User Bundle



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
	userBundleId := "ca1d2263-d1f1-43ac-ba53-248e7a4bb26a" // string | 
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserBundlesAPI.DeactivateUserBundle(context.Background(), userBundleId).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserBundlesAPI.DeactivateUserBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateUserBundle`: UserBundleCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `UserBundlesAPI.DeactivateUserBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userBundleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**UserBundleCreateResponse**](UserBundleCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnusedUserBundles

> UnusedUserBundlesResponse GetUnusedUserBundles(ctx).FilterCountryIso(filterCountryIso).AuthorizationBearer(authorizationBearer).Execute()

Get Unused User Bundles



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
	filterCountryIso := []string{"US"} // []string | Filter by country code. (optional)
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserBundlesAPI.GetUnusedUserBundles(context.Background()).FilterCountryIso(filterCountryIso).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserBundlesAPI.GetUnusedUserBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnusedUserBundles`: UnusedUserBundlesResponse
	fmt.Fprintf(os.Stdout, "Response from `UserBundlesAPI.GetUnusedUserBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUnusedUserBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCountryIso** | **[]string** | Filter by country code. | 
 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**UnusedUserBundlesResponse**](UnusedUserBundlesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBundleById

> UserBundleResponse GetUserBundleById(ctx, userBundleId).AuthorizationBearer(authorizationBearer).Execute()

Get User Bundle by Id



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
	userBundleId := "ca1d2263-d1f1-43ac-ba53-248e7a4bb26a" // string | 
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserBundlesAPI.GetUserBundleById(context.Background(), userBundleId).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserBundlesAPI.GetUserBundleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBundleById`: UserBundleResponse
	fmt.Fprintf(os.Stdout, "Response from `UserBundlesAPI.GetUserBundleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userBundleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBundleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**UserBundleResponse**](UserBundleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBundleResources

> UserBundleResourcesResponse GetUserBundleResources(ctx, userBundleId).AuthorizationBearer(authorizationBearer).Execute()

Get User Bundle Resources



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
	userBundleId := "ca1d2263-d1f1-43ac-ba53-248e7a4bb26a" // string | 
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserBundlesAPI.GetUserBundleResources(context.Background(), userBundleId).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserBundlesAPI.GetUserBundleResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBundleResources`: UserBundleResourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `UserBundlesAPI.GetUserBundleResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userBundleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBundleResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**UserBundleResourcesResponse**](UserBundleResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBundles

> PaginatedUserBundlesResponse GetUserBundles(ctx).FilterCountryIso(filterCountryIso).FilterResource(filterResource).PageNumber(pageNumber).PageSize(pageSize).AuthorizationBearer(authorizationBearer).Execute()

Get User Bundles



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
	filterCountryIso := []string{"US"} // []string | Filter by country code. (optional)
	filterResource := []string{"+15617819942"} // []string | Filter by resource. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserBundlesAPI.GetUserBundles(context.Background()).FilterCountryIso(filterCountryIso).FilterResource(filterResource).PageNumber(pageNumber).PageSize(pageSize).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserBundlesAPI.GetUserBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBundles`: PaginatedUserBundlesResponse
	fmt.Fprintf(os.Stdout, "Response from `UserBundlesAPI.GetUserBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCountryIso** | **[]string** | Filter by country code. | 
 **filterResource** | **[]string** | Filter by resource. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**PaginatedUserBundlesResponse**](PaginatedUserBundlesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

