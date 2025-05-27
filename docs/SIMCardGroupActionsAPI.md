# \SIMCardGroupActionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSimCardGroupAction**](SIMCardGroupActionsAPI.md#GetSimCardGroupAction) | **Get** /sim_card_group_actions/{id} | Get SIM card group action details
[**GetSimCardGroupActions**](SIMCardGroupActionsAPI.md#GetSimCardGroupActions) | **Get** /sim_card_group_actions | List SIM card group actions



## GetSimCardGroupAction

> GetSimCardGroupAction200Response GetSimCardGroupAction(ctx, id).Execute()

Get SIM card group action details



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
	resp, r, err := apiClient.SIMCardGroupActionsAPI.GetSimCardGroupAction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupActionsAPI.GetSimCardGroupAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardGroupAction`: GetSimCardGroupAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupActionsAPI.GetSimCardGroupAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardGroupActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardGroupAction200Response**](GetSimCardGroupAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardGroupActions

> GetSimCardGroupActions200Response GetSimCardGroupActions(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterSimCardGroupId(filterSimCardGroupId).FilterStatus(filterStatus).FilterType(filterType).Execute()

List SIM card group actions



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
	filterSimCardGroupId := "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9" // string | A valid SIM card group ID. (optional)
	filterStatus := "in-progress" // string | Filter by a specific status of the resource's lifecycle. (optional)
	filterType := "set_private_wireless_gateway" // string | Filter by action type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupActionsAPI.GetSimCardGroupActions(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterSimCardGroupId(filterSimCardGroupId).FilterStatus(filterStatus).FilterType(filterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupActionsAPI.GetSimCardGroupActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardGroupActions`: GetSimCardGroupActions200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupActionsAPI.GetSimCardGroupActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardGroupActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterSimCardGroupId** | **string** | A valid SIM card group ID. | 
 **filterStatus** | **string** | Filter by a specific status of the resource&#39;s lifecycle. | 
 **filterType** | **string** | Filter by action type. | 

### Return type

[**GetSimCardGroupActions200Response**](GetSimCardGroupActions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

