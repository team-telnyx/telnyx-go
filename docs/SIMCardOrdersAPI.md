# \SIMCardOrdersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSimCardOrder**](SIMCardOrdersAPI.md#CreateSimCardOrder) | **Post** /sim_card_orders | Create a SIM card order
[**GetSimCardOrder**](SIMCardOrdersAPI.md#GetSimCardOrder) | **Get** /sim_card_orders/{id} | Get a single SIM card order
[**GetSimCardOrders**](SIMCardOrdersAPI.md#GetSimCardOrders) | **Get** /sim_card_orders | Get all SIM card orders
[**PreviewSimCardOrders**](SIMCardOrdersAPI.md#PreviewSimCardOrders) | **Post** /sim_card_order_preview | Preview SIM card orders



## CreateSimCardOrder

> CreateSimCardOrder200Response CreateSimCardOrder(ctx).SimCardOrderCreate(simCardOrderCreate).Execute()

Create a SIM card order



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
	simCardOrderCreate := *openapiclient.NewSimCardOrderCreate("1293384261075731499", int32(23)) // SimCardOrderCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardOrdersAPI.CreateSimCardOrder(context.Background()).SimCardOrderCreate(simCardOrderCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardOrdersAPI.CreateSimCardOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimCardOrder`: CreateSimCardOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardOrdersAPI.CreateSimCardOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimCardOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simCardOrderCreate** | [**SimCardOrderCreate**](SimCardOrderCreate.md) |  | 

### Return type

[**CreateSimCardOrder200Response**](CreateSimCardOrder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardOrder

> CreateSimCardOrder200Response GetSimCardOrder(ctx, id).Execute()

Get a single SIM card order



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
	resp, r, err := apiClient.SIMCardOrdersAPI.GetSimCardOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardOrdersAPI.GetSimCardOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardOrder`: CreateSimCardOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardOrdersAPI.GetSimCardOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSimCardOrder200Response**](CreateSimCardOrder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardOrders

> GetSimCardOrders200Response GetSimCardOrders(ctx).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).FilterQuantity(filterQuantity).FilterCostAmount(filterCostAmount).FilterCostCurrency(filterCostCurrency).FilterAddressId(filterAddressId).FilterAddressStreetAddress(filterAddressStreetAddress).FilterAddressExtendedAddress(filterAddressExtendedAddress).FilterAddressLocality(filterAddressLocality).FilterAddressAdministrativeArea(filterAddressAdministrativeArea).FilterAddressCountryCode(filterAddressCountryCode).FilterAddressPostalCode(filterAddressPostalCode).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get all SIM card orders



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
	filterCreatedAt := "2018-02-02T22:25:27.521Z" // string | Filter by ISO 8601 formatted date-time string matching resource creation date-time. (optional)
	filterUpdatedAt := "2018-02-02T22:25:27.521Z" // string | Filter by ISO 8601 formatted date-time string matching resource last update date-time. (optional)
	filterQuantity := int32(21) // int32 | Filter orders by how many SIM cards were ordered. (optional)
	filterCostAmount := "2.53" // string | The total monetary amount of the order. (optional)
	filterCostCurrency := "USD" // string | Filter by ISO 4217 currency string. (optional)
	filterAddressId := "1293384261075731499" // string | Uniquely identifies the address for the order. (optional)
	filterAddressStreetAddress := "600 Congress Avenue" // string | Returns entries with matching name of the street where the address is located. (optional)
	filterAddressExtendedAddress := "14th Floor" // string | Returns entries with matching name of the supplemental field for address information. (optional)
	filterAddressLocality := "Austin" // string | Filter by the name of the city where the address is located. (optional)
	filterAddressAdministrativeArea := "TX" // string | Filter by state or province where the address is located. (optional)
	filterAddressCountryCode := "US" // string | Filter by the mobile operator two-character (ISO 3166-1 alpha-2) origin country code. (optional)
	filterAddressPostalCode := "78701" // string | Filter by postal code for the address. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardOrdersAPI.GetSimCardOrders(context.Background()).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).FilterQuantity(filterQuantity).FilterCostAmount(filterCostAmount).FilterCostCurrency(filterCostCurrency).FilterAddressId(filterAddressId).FilterAddressStreetAddress(filterAddressStreetAddress).FilterAddressExtendedAddress(filterAddressExtendedAddress).FilterAddressLocality(filterAddressLocality).FilterAddressAdministrativeArea(filterAddressAdministrativeArea).FilterAddressCountryCode(filterAddressCountryCode).FilterAddressPostalCode(filterAddressPostalCode).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardOrdersAPI.GetSimCardOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardOrders`: GetSimCardOrders200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardOrdersAPI.GetSimCardOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCreatedAt** | **string** | Filter by ISO 8601 formatted date-time string matching resource creation date-time. | 
 **filterUpdatedAt** | **string** | Filter by ISO 8601 formatted date-time string matching resource last update date-time. | 
 **filterQuantity** | **int32** | Filter orders by how many SIM cards were ordered. | 
 **filterCostAmount** | **string** | The total monetary amount of the order. | 
 **filterCostCurrency** | **string** | Filter by ISO 4217 currency string. | 
 **filterAddressId** | **string** | Uniquely identifies the address for the order. | 
 **filterAddressStreetAddress** | **string** | Returns entries with matching name of the street where the address is located. | 
 **filterAddressExtendedAddress** | **string** | Returns entries with matching name of the supplemental field for address information. | 
 **filterAddressLocality** | **string** | Filter by the name of the city where the address is located. | 
 **filterAddressAdministrativeArea** | **string** | Filter by state or province where the address is located. | 
 **filterAddressCountryCode** | **string** | Filter by the mobile operator two-character (ISO 3166-1 alpha-2) origin country code. | 
 **filterAddressPostalCode** | **string** | Filter by postal code for the address. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**GetSimCardOrders200Response**](GetSimCardOrders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewSimCardOrders

> PreviewSimCardOrders202Response PreviewSimCardOrders(ctx).PreviewSimCardOrdersRequest(previewSimCardOrdersRequest).Execute()

Preview SIM card orders



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
	previewSimCardOrdersRequest := *openapiclient.NewPreviewSimCardOrdersRequest(int32(21), "1293384261075731499") // PreviewSimCardOrdersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardOrdersAPI.PreviewSimCardOrders(context.Background()).PreviewSimCardOrdersRequest(previewSimCardOrdersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardOrdersAPI.PreviewSimCardOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewSimCardOrders`: PreviewSimCardOrders202Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardOrdersAPI.PreviewSimCardOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewSimCardOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewSimCardOrdersRequest** | [**PreviewSimCardOrdersRequest**](PreviewSimCardOrdersRequest.md) |  | 

### Return type

[**PreviewSimCardOrders202Response**](PreviewSimCardOrders202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

