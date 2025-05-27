# \OTAUpdatesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOtaUpdate**](OTAUpdatesAPI.md#GetOtaUpdate) | **Get** /ota_updates/{id} | Get OTA update
[**ListOtaUpdates**](OTAUpdatesAPI.md#ListOtaUpdates) | **Get** /ota_updates | List OTA updates



## GetOtaUpdate

> GetOtaUpdate200Response GetOtaUpdate(ctx, id).Execute()

Get OTA update



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
	resp, r, err := apiClient.OTAUpdatesAPI.GetOtaUpdate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OTAUpdatesAPI.GetOtaUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOtaUpdate`: GetOtaUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `OTAUpdatesAPI.GetOtaUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOtaUpdate200Response**](GetOtaUpdate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOtaUpdates

> ListOtaUpdates200Response ListOtaUpdates(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterStatus(filterStatus).FilterSimCardId(filterSimCardId).FilterType(filterType).Execute()

List OTA updates

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
	filterStatus := "in-progress" // string | Filter by a specific status of the resource's lifecycle. (optional)
	filterSimCardId := "filterSimCardId_example" // string | The SIM card identification UUID. (optional)
	filterType := "sim_card_network_preferences" // string | Filter by type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OTAUpdatesAPI.ListOtaUpdates(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterStatus(filterStatus).FilterSimCardId(filterSimCardId).FilterType(filterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OTAUpdatesAPI.ListOtaUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOtaUpdates`: ListOtaUpdates200Response
	fmt.Fprintf(os.Stdout, "Response from `OTAUpdatesAPI.ListOtaUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOtaUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterStatus** | **string** | Filter by a specific status of the resource&#39;s lifecycle. | 
 **filterSimCardId** | **string** | The SIM card identification UUID. | 
 **filterType** | **string** | Filter by type. | 

### Return type

[**ListOtaUpdates200Response**](ListOtaUpdates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

