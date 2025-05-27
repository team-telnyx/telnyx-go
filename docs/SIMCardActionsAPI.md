# \SIMCardActionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBulkSimCardAction**](SIMCardActionsAPI.md#GetBulkSimCardAction) | **Get** /bulk_sim_card_actions/{id} | Get bulk SIM card action details
[**GetSimCardAction**](SIMCardActionsAPI.md#GetSimCardAction) | **Get** /sim_card_actions/{id} | Get SIM card action details
[**ListBulkSimCardActions**](SIMCardActionsAPI.md#ListBulkSimCardActions) | **Get** /bulk_sim_card_actions | List bulk SIM card actions
[**ListSimCardActions**](SIMCardActionsAPI.md#ListSimCardActions) | **Get** /sim_card_actions | List SIM card actions



## GetBulkSimCardAction

> GetBulkSimCardAction200Response GetBulkSimCardAction(ctx, id).Execute()

Get bulk SIM card action details



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardActionsAPI.GetBulkSimCardAction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardActionsAPI.GetBulkSimCardAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkSimCardAction`: GetBulkSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardActionsAPI.GetBulkSimCardAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkSimCardActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBulkSimCardAction200Response**](GetBulkSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardAction

> GetSimCardAction200Response GetSimCardAction(ctx, id).Execute()

Get SIM card action details



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardActionsAPI.GetSimCardAction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardActionsAPI.GetSimCardAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardAction`: GetSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardActionsAPI.GetSimCardAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardAction200Response**](GetSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBulkSimCardActions

> ListBulkSimCardActions200Response ListBulkSimCardActions(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterActionType(filterActionType).Execute()

List bulk SIM card actions



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
	filterActionType := "bulk_set_public_ips" // string | Filter by action type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardActionsAPI.ListBulkSimCardActions(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterActionType(filterActionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardActionsAPI.ListBulkSimCardActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBulkSimCardActions`: ListBulkSimCardActions200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardActionsAPI.ListBulkSimCardActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBulkSimCardActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterActionType** | **string** | Filter by action type. | 

### Return type

[**ListBulkSimCardActions200Response**](ListBulkSimCardActions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSimCardActions

> ListSimCardActions200Response ListSimCardActions(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterSimCardId(filterSimCardId).FilterStatus(filterStatus).FilterBulkSimCardActionId(filterBulkSimCardActionId).FilterActionType(filterActionType).Execute()

List SIM card actions



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
	filterSimCardId := "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9" // string | A valid SIM card ID. (optional)
	filterStatus := "in-progress" // string | Filter by a specific status of the resource's lifecycle. (optional)
	filterBulkSimCardActionId := "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9" // string | Filter by a bulk SIM card action ID. (optional)
	filterActionType := "disable" // string | Filter by action type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardActionsAPI.ListSimCardActions(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterSimCardId(filterSimCardId).FilterStatus(filterStatus).FilterBulkSimCardActionId(filterBulkSimCardActionId).FilterActionType(filterActionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardActionsAPI.ListSimCardActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSimCardActions`: ListSimCardActions200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardActionsAPI.ListSimCardActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSimCardActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterSimCardId** | **string** | A valid SIM card ID. | 
 **filterStatus** | **string** | Filter by a specific status of the resource&#39;s lifecycle. | 
 **filterBulkSimCardActionId** | **string** | Filter by a bulk SIM card action ID. | 
 **filterActionType** | **string** | Filter by action type. | 

### Return type

[**ListSimCardActions200Response**](ListSimCardActions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

