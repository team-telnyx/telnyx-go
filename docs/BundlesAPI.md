# \BundlesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingBundleById**](BundlesAPI.md#GetBillingBundleById) | **Get** /bundle_pricing/billing_bundles/{bundle_id} | Get Bundle By Id
[**GetUserBillingBundles**](BundlesAPI.md#GetUserBillingBundles) | **Get** /bundle_pricing/billing_bundles | Retrieve Bundles



## GetBillingBundleById

> BillingBundleResponse GetBillingBundleById(ctx, bundleId).AuthorizationBearer(authorizationBearer).Execute()

Get Bundle By Id



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
	bundleId := "8661948c-a386-4385-837f-af00f40f111a" // string | 
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.GetBillingBundleById(context.Background(), bundleId).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetBillingBundleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingBundleById`: BillingBundleResponse
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetBillingBundleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingBundleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**BillingBundleResponse**](BillingBundleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingBundles

> PaginatedBillingBundlesResponse GetUserBillingBundles(ctx).FilterCountryIso(filterCountryIso).PageNumber(pageNumber).PageSize(pageSize).AuthorizationBearer(authorizationBearer).Execute()

Retrieve Bundles



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	authorizationBearer := "authorizationBearer_example" // string | Format: Bearer <TOKEN> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.GetUserBillingBundles(context.Background()).FilterCountryIso(filterCountryIso).PageNumber(pageNumber).PageSize(pageSize).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.GetUserBillingBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingBundles`: PaginatedBillingBundlesResponse
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.GetUserBillingBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCountryIso** | **[]string** | Filter by country code. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **authorizationBearer** | **string** | Format: Bearer &lt;TOKEN&gt; | 

### Return type

[**PaginatedBillingBundlesResponse**](PaginatedBillingBundlesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

