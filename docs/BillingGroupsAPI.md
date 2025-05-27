# \BillingGroupsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingGroup**](BillingGroupsAPI.md#CreateBillingGroup) | **Post** /billing_groups | Create a billing group
[**DeleteBillingGroup**](BillingGroupsAPI.md#DeleteBillingGroup) | **Delete** /billing_groups/{id} | Delete a billing group
[**GetBillingGroup**](BillingGroupsAPI.md#GetBillingGroup) | **Get** /billing_groups/{id} | Get a billing group
[**ListBillingGroups**](BillingGroupsAPI.md#ListBillingGroups) | **Get** /billing_groups | List all billing groups
[**UpdateBillingGroup**](BillingGroupsAPI.md#UpdateBillingGroup) | **Patch** /billing_groups/{id} | Update a billing group



## CreateBillingGroup

> CreateBillingGroup200Response CreateBillingGroup(ctx).NewBillingGroup(newBillingGroup).Execute()

Create a billing group



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
	newBillingGroup := *openapiclient.NewNewBillingGroup() // NewBillingGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingGroupsAPI.CreateBillingGroup(context.Background()).NewBillingGroup(newBillingGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsAPI.CreateBillingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingGroup`: CreateBillingGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingGroupsAPI.CreateBillingGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBillingGroup** | [**NewBillingGroup**](NewBillingGroup.md) |  | 

### Return type

[**CreateBillingGroup200Response**](CreateBillingGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingGroup

> CreateBillingGroup200Response DeleteBillingGroup(ctx, id).Execute()

Delete a billing group



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
	id := "f5586561-8ff0-4291-a0ac-84fe544797bd" // string | The id of the billing group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingGroupsAPI.DeleteBillingGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsAPI.DeleteBillingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillingGroup`: CreateBillingGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingGroupsAPI.DeleteBillingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the billing group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateBillingGroup200Response**](CreateBillingGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingGroup

> CreateBillingGroup200Response GetBillingGroup(ctx, id).Execute()

Get a billing group



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
	id := "f5586561-8ff0-4291-a0ac-84fe544797bd" // string | The id of the billing group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingGroupsAPI.GetBillingGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsAPI.GetBillingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingGroup`: CreateBillingGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingGroupsAPI.GetBillingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the billing group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateBillingGroup200Response**](CreateBillingGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillingGroups

> ListBillingGroups200Response ListBillingGroups(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all billing groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingGroupsAPI.ListBillingGroups(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsAPI.ListBillingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingGroups`: ListBillingGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingGroupsAPI.ListBillingGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load | [default to 1]
 **pageSize** | **int32** | The size of the page | [default to 20]

### Return type

[**ListBillingGroups200Response**](ListBillingGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillingGroup

> CreateBillingGroup200Response UpdateBillingGroup(ctx, id).UpdateBillingGroup(updateBillingGroup).Execute()

Update a billing group



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
	id := "f5586561-8ff0-4291-a0ac-84fe544797bd" // string | The id of the billing group
	updateBillingGroup := *openapiclient.NewUpdateBillingGroup() // UpdateBillingGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingGroupsAPI.UpdateBillingGroup(context.Background(), id).UpdateBillingGroup(updateBillingGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingGroupsAPI.UpdateBillingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillingGroup`: CreateBillingGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingGroupsAPI.UpdateBillingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the billing group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBillingGroup** | [**UpdateBillingGroup**](UpdateBillingGroup.md) |  | 

### Return type

[**CreateBillingGroup200Response**](CreateBillingGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

