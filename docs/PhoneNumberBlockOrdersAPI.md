# \PhoneNumberBlockOrdersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberBlockOrder**](PhoneNumberBlockOrdersAPI.md#CreateNumberBlockOrder) | **Post** /number_block_orders | Create a number block order
[**ListNumberBlockOrders**](PhoneNumberBlockOrdersAPI.md#ListNumberBlockOrders) | **Get** /number_block_orders | List number block orders
[**RetrieveNumberBlockOrder**](PhoneNumberBlockOrdersAPI.md#RetrieveNumberBlockOrder) | **Get** /number_block_orders/{number_block_order_id} | Retrieve a number block order



## CreateNumberBlockOrder

> NumberBlockOrderResponse CreateNumberBlockOrder(ctx).CreateNumberBlockOrderRequest(createNumberBlockOrderRequest).Execute()

Create a number block order



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
	createNumberBlockOrderRequest := *openapiclient.NewCreateNumberBlockOrderRequest("+19705555000", int32(10)) // CreateNumberBlockOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberBlockOrdersAPI.CreateNumberBlockOrder(context.Background()).CreateNumberBlockOrderRequest(createNumberBlockOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberBlockOrdersAPI.CreateNumberBlockOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNumberBlockOrder`: NumberBlockOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberBlockOrdersAPI.CreateNumberBlockOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNumberBlockOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNumberBlockOrderRequest** | [**CreateNumberBlockOrderRequest**](CreateNumberBlockOrderRequest.md) |  | 

### Return type

[**NumberBlockOrderResponse**](NumberBlockOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNumberBlockOrders

> ListNumberBlockOrdersResponse ListNumberBlockOrders(ctx).FilterStatus(filterStatus).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).FilterPhoneNumbersStartingNumber(filterPhoneNumbersStartingNumber).PageNumber(pageNumber).PageSize(pageSize).Execute()

List number block orders



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
	filterStatus := "pending" // string | Filter number block orders by status. (optional)
	filterCreatedAtGt := "2018-01-01T00:00:00.000000Z" // string | Filter number block orders later than this value. (optional)
	filterCreatedAtLt := "2018-01-01T00:00:00.000000Z" // string | Filter number block orders earlier than this value. (optional)
	filterPhoneNumbersStartingNumber := "+19705555000" // string | Filter number block  orders having these phone numbers. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberBlockOrdersAPI.ListNumberBlockOrders(context.Background()).FilterStatus(filterStatus).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).FilterPhoneNumbersStartingNumber(filterPhoneNumbersStartingNumber).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberBlockOrdersAPI.ListNumberBlockOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNumberBlockOrders`: ListNumberBlockOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberBlockOrdersAPI.ListNumberBlockOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNumberBlockOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Filter number block orders by status. | 
 **filterCreatedAtGt** | **string** | Filter number block orders later than this value. | 
 **filterCreatedAtLt** | **string** | Filter number block orders earlier than this value. | 
 **filterPhoneNumbersStartingNumber** | **string** | Filter number block  orders having these phone numbers. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListNumberBlockOrdersResponse**](ListNumberBlockOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNumberBlockOrder

> NumberBlockOrderResponse RetrieveNumberBlockOrder(ctx, numberBlockOrderId).Execute()

Retrieve a number block order



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
	numberBlockOrderId := "numberBlockOrderId_example" // string | The number block order ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberBlockOrdersAPI.RetrieveNumberBlockOrder(context.Background(), numberBlockOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberBlockOrdersAPI.RetrieveNumberBlockOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveNumberBlockOrder`: NumberBlockOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberBlockOrdersAPI.RetrieveNumberBlockOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberBlockOrderId** | **string** | The number block order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNumberBlockOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NumberBlockOrderResponse**](NumberBlockOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

