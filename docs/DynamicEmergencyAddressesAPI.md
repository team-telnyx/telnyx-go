# \DynamicEmergencyAddressesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicEmergencyAddress**](DynamicEmergencyAddressesAPI.md#CreateDynamicEmergencyAddress) | **Post** /dynamic_emergency_addresses | Create a dynamic emergency address.
[**DeleteDynamicEmergencyAddress**](DynamicEmergencyAddressesAPI.md#DeleteDynamicEmergencyAddress) | **Delete** /dynamic_emergency_addresses/{id} | Delete a dynamic emergency address
[**GetDynamicEmergencyAddress**](DynamicEmergencyAddressesAPI.md#GetDynamicEmergencyAddress) | **Get** /dynamic_emergency_addresses/{id} | Get a dynamic emergency address
[**ListDynamicEmergencyAddresses**](DynamicEmergencyAddressesAPI.md#ListDynamicEmergencyAddresses) | **Get** /dynamic_emergency_addresses | List dynamic emergency addresses



## CreateDynamicEmergencyAddress

> CreateDynamicEmergencyAddress201Response CreateDynamicEmergencyAddress(ctx).DynamicEmergencyAddress(dynamicEmergencyAddress).Execute()

Create a dynamic emergency address.



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
	dynamicEmergencyAddress := *openapiclient.NewDynamicEmergencyAddress("600", "Congress", "Austin", "TX", "78701") // DynamicEmergencyAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyAddressesAPI.CreateDynamicEmergencyAddress(context.Background()).DynamicEmergencyAddress(dynamicEmergencyAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyAddressesAPI.CreateDynamicEmergencyAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicEmergencyAddress`: CreateDynamicEmergencyAddress201Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyAddressesAPI.CreateDynamicEmergencyAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicEmergencyAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicEmergencyAddress** | [**DynamicEmergencyAddress**](DynamicEmergencyAddress.md) |  | 

### Return type

[**CreateDynamicEmergencyAddress201Response**](CreateDynamicEmergencyAddress201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDynamicEmergencyAddress

> CreateDynamicEmergencyAddress201Response DeleteDynamicEmergencyAddress(ctx, id).Execute()

Delete a dynamic emergency address



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Dynamic Emergency Address id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyAddressesAPI.DeleteDynamicEmergencyAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyAddressesAPI.DeleteDynamicEmergencyAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDynamicEmergencyAddress`: CreateDynamicEmergencyAddress201Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyAddressesAPI.DeleteDynamicEmergencyAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Emergency Address id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicEmergencyAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDynamicEmergencyAddress201Response**](CreateDynamicEmergencyAddress201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicEmergencyAddress

> CreateDynamicEmergencyAddress201Response GetDynamicEmergencyAddress(ctx, id).Execute()

Get a dynamic emergency address



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Dynamic Emergency Address id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyAddressesAPI.GetDynamicEmergencyAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyAddressesAPI.GetDynamicEmergencyAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicEmergencyAddress`: CreateDynamicEmergencyAddress201Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyAddressesAPI.GetDynamicEmergencyAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Emergency Address id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicEmergencyAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDynamicEmergencyAddress201Response**](CreateDynamicEmergencyAddress201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDynamicEmergencyAddresses

> ListDynamicEmergencyAddresses200Response ListDynamicEmergencyAddresses(ctx).FilterStatus(filterStatus).FilterCountryCode(filterCountryCode).PageNumber(pageNumber).PageSize(pageSize).Execute()

List dynamic emergency addresses



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
	filterStatus := "filterStatus_example" // string | Filter by status. (optional)
	filterCountryCode := "filterCountryCode_example" // string | Filter by country code. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicEmergencyAddressesAPI.ListDynamicEmergencyAddresses(context.Background()).FilterStatus(filterStatus).FilterCountryCode(filterCountryCode).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicEmergencyAddressesAPI.ListDynamicEmergencyAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDynamicEmergencyAddresses`: ListDynamicEmergencyAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicEmergencyAddressesAPI.ListDynamicEmergencyAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDynamicEmergencyAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Filter by status. | 
 **filterCountryCode** | **string** | Filter by country code. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListDynamicEmergencyAddresses200Response**](ListDynamicEmergencyAddresses200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

