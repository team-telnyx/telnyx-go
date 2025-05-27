# \PhoneNumberOrdersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubNumberOrder**](PhoneNumberOrdersAPI.md#CancelSubNumberOrder) | **Patch** /sub_number_orders/{sub_number_order_id}/cancel | Cancel a sub number order
[**CreateComment**](PhoneNumberOrdersAPI.md#CreateComment) | **Post** /comments | Create a comment
[**CreateNumberOrder**](PhoneNumberOrdersAPI.md#CreateNumberOrder) | **Post** /number_orders | Create a number order
[**GetNumberOrderPhoneNumber**](PhoneNumberOrdersAPI.md#GetNumberOrderPhoneNumber) | **Get** /number_order_phone_numbers/{number_order_phone_number_id} | Retrieve a single phone number within a number order.
[**GetSubNumberOrder**](PhoneNumberOrdersAPI.md#GetSubNumberOrder) | **Get** /sub_number_orders/{sub_number_order_id} | Retrieve a sub number order
[**ListComments**](PhoneNumberOrdersAPI.md#ListComments) | **Get** /comments | Retrieve all comments
[**ListNumberOrders**](PhoneNumberOrdersAPI.md#ListNumberOrders) | **Get** /number_orders | List number orders
[**ListSubNumberOrders**](PhoneNumberOrdersAPI.md#ListSubNumberOrders) | **Get** /sub_number_orders | List sub number orders
[**MarkCommentRead**](PhoneNumberOrdersAPI.md#MarkCommentRead) | **Patch** /comments/{id}/read | Mark a comment as read
[**RetrieveComment**](PhoneNumberOrdersAPI.md#RetrieveComment) | **Get** /comments/{id} | Retrieve a comment
[**RetrieveNumberOrder**](PhoneNumberOrdersAPI.md#RetrieveNumberOrder) | **Get** /number_orders/{number_order_id} | Retrieve a number order
[**RetrieveOrderPhoneNumbers**](PhoneNumberOrdersAPI.md#RetrieveOrderPhoneNumbers) | **Get** /number_order_phone_numbers | Retrieve a list of phone numbers associated to orders
[**UpdateNumberOrder**](PhoneNumberOrdersAPI.md#UpdateNumberOrder) | **Patch** /number_orders/{number_order_id} | Update a number order
[**UpdateNumberOrderPhoneNumber**](PhoneNumberOrdersAPI.md#UpdateNumberOrderPhoneNumber) | **Patch** /number_order_phone_numbers/{number_order_phone_number_id} | Update requirements for a single phone number within a number order.
[**UpdateSubNumberOrder**](PhoneNumberOrdersAPI.md#UpdateSubNumberOrder) | **Patch** /sub_number_orders/{sub_number_order_id} | Update a sub number order&#39;s requirements



## CancelSubNumberOrder

> SubNumberOrderResponse CancelSubNumberOrder(ctx, subNumberOrderId).Execute()

Cancel a sub number order



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
	subNumberOrderId := "subNumberOrderId_example" // string | The ID of the sub number order.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.CancelSubNumberOrder(context.Background(), subNumberOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.CancelSubNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSubNumberOrder`: SubNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.CancelSubNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subNumberOrderId** | **string** | The ID of the sub number order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubNumberOrderResponse**](SubNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComment

> CreateComment200Response CreateComment(ctx).Comment(comment).Execute()

Create a comment

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
	comment := *openapiclient.NewComment() // Comment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.CreateComment(context.Background()).Comment(comment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.CreateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment`: CreateComment200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.CreateComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | [**Comment**](Comment.md) |  | 

### Return type

[**CreateComment200Response**](CreateComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNumberOrder

> NumberOrderResponse CreateNumberOrder(ctx).CreateNumberOrderRequest(createNumberOrderRequest).Execute()

Create a number order



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
	createNumberOrderRequest := *openapiclient.NewCreateNumberOrderRequest() // CreateNumberOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.CreateNumberOrder(context.Background()).CreateNumberOrderRequest(createNumberOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.CreateNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNumberOrder`: NumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.CreateNumberOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNumberOrderRequest** | [**CreateNumberOrderRequest**](CreateNumberOrderRequest.md) |  | 

### Return type

[**NumberOrderResponse**](NumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNumberOrderPhoneNumber

> NumberOrderPhoneNumberResponse GetNumberOrderPhoneNumber(ctx, numberOrderPhoneNumberId).Execute()

Retrieve a single phone number within a number order.



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
	numberOrderPhoneNumberId := "numberOrderPhoneNumberId_example" // string | The number order phone number ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.GetNumberOrderPhoneNumber(context.Background(), numberOrderPhoneNumberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.GetNumberOrderPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNumberOrderPhoneNumber`: NumberOrderPhoneNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.GetNumberOrderPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberOrderPhoneNumberId** | **string** | The number order phone number ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNumberOrderPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NumberOrderPhoneNumberResponse**](NumberOrderPhoneNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubNumberOrder

> SubNumberOrderResponse GetSubNumberOrder(ctx, subNumberOrderId).FilterIncludePhoneNumbers(filterIncludePhoneNumbers).Execute()

Retrieve a sub number order



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
	subNumberOrderId := "subNumberOrderId_example" // string | The sub number order ID.
	filterIncludePhoneNumbers := true // bool | Include the first 50 phone number objects in the results (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.GetSubNumberOrder(context.Background(), subNumberOrderId).FilterIncludePhoneNumbers(filterIncludePhoneNumbers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.GetSubNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubNumberOrder`: SubNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.GetSubNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subNumberOrderId** | **string** | The sub number order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterIncludePhoneNumbers** | **bool** | Include the first 50 phone number objects in the results | [default to false]

### Return type

[**SubNumberOrderResponse**](SubNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComments

> ListComments200Response ListComments(ctx).FilterCommentRecordType(filterCommentRecordType).FilterCommentRecordId(filterCommentRecordId).Execute()

Retrieve all comments

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
	filterCommentRecordType := "sub_number_order" // string | Record type that the comment relates to
	filterCommentRecordId := "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576" // string | ID of the record the comments relate to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.ListComments(context.Background()).FilterCommentRecordType(filterCommentRecordType).FilterCommentRecordId(filterCommentRecordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.ListComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComments`: ListComments200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.ListComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCommentRecordType** | **string** | Record type that the comment relates to | 
 **filterCommentRecordId** | **string** | ID of the record the comments relate to | 

### Return type

[**ListComments200Response**](ListComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNumberOrders

> ListNumberOrdersResponse ListNumberOrders(ctx).FilterStatus(filterStatus).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).FilterPhoneNumbersCount(filterPhoneNumbersCount).FilterCustomerReference(filterCustomerReference).FilterRequirementsMet(filterRequirementsMet).PageNumber(pageNumber).PageSize(pageSize).Execute()

List number orders



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
	filterStatus := "filterStatus_example" // string | Filter number orders by status. (optional)
	filterCreatedAtGt := "filterCreatedAtGt_example" // string | Filter number orders later than this value. (optional)
	filterCreatedAtLt := "filterCreatedAtLt_example" // string | Filter number orders earlier than this value. (optional)
	filterPhoneNumbersCount := "filterPhoneNumbersCount_example" // string | Filter number order with this amount of numbers (optional)
	filterCustomerReference := "filterCustomerReference_example" // string | Filter number orders via the customer reference set. (optional)
	filterRequirementsMet := true // bool | Filter number orders by requirements met. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.ListNumberOrders(context.Background()).FilterStatus(filterStatus).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).FilterPhoneNumbersCount(filterPhoneNumbersCount).FilterCustomerReference(filterCustomerReference).FilterRequirementsMet(filterRequirementsMet).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.ListNumberOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNumberOrders`: ListNumberOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.ListNumberOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNumberOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Filter number orders by status. | 
 **filterCreatedAtGt** | **string** | Filter number orders later than this value. | 
 **filterCreatedAtLt** | **string** | Filter number orders earlier than this value. | 
 **filterPhoneNumbersCount** | **string** | Filter number order with this amount of numbers | 
 **filterCustomerReference** | **string** | Filter number orders via the customer reference set. | 
 **filterRequirementsMet** | **bool** | Filter number orders by requirements met. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListNumberOrdersResponse**](ListNumberOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubNumberOrders

> ListSubNumberOrdersResponse ListSubNumberOrders(ctx).FilterStatus(filterStatus).FilterOrderRequestId(filterOrderRequestId).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterPhoneNumbersCount(filterPhoneNumbersCount).Execute()

List sub number orders



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
	filterStatus := "filterStatus_example" // string | Filter sub number orders by status. (optional)
	filterOrderRequestId := "12ade33a-21c0-473b-b055-b3c836e1c293" // string | ID of the number order the sub number order belongs to (optional)
	filterCountryCode := "US" // string | ISO alpha-2 country code. (optional)
	filterPhoneNumberType := "local" // string | Phone Number Type (optional)
	filterPhoneNumbersCount := int32(1) // int32 | Amount of numbers in the sub number order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.ListSubNumberOrders(context.Background()).FilterStatus(filterStatus).FilterOrderRequestId(filterOrderRequestId).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterPhoneNumbersCount(filterPhoneNumbersCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.ListSubNumberOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubNumberOrders`: ListSubNumberOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.ListSubNumberOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubNumberOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Filter sub number orders by status. | 
 **filterOrderRequestId** | **string** | ID of the number order the sub number order belongs to | 
 **filterCountryCode** | **string** | ISO alpha-2 country code. | 
 **filterPhoneNumberType** | **string** | Phone Number Type | 
 **filterPhoneNumbersCount** | **int32** | Amount of numbers in the sub number order | 

### Return type

[**ListSubNumberOrdersResponse**](ListSubNumberOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkCommentRead

> CreateComment200Response MarkCommentRead(ctx, id).Execute()

Mark a comment as read

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
	id := "id_example" // string | The comment ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.MarkCommentRead(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.MarkCommentRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkCommentRead`: CreateComment200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.MarkCommentRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The comment ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkCommentReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateComment200Response**](CreateComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveComment

> CreateComment200Response RetrieveComment(ctx, id).Execute()

Retrieve a comment

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
	id := "id_example" // string | The comment ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.RetrieveComment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.RetrieveComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveComment`: CreateComment200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.RetrieveComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The comment ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateComment200Response**](CreateComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNumberOrder

> NumberOrderResponse RetrieveNumberOrder(ctx, numberOrderId).Execute()

Retrieve a number order



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
	numberOrderId := "numberOrderId_example" // string | The number order ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.RetrieveNumberOrder(context.Background(), numberOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.RetrieveNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveNumberOrder`: NumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.RetrieveNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberOrderId** | **string** | The number order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NumberOrderResponse**](NumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOrderPhoneNumbers

> ListNumberOrderPhoneNumbersResponse RetrieveOrderPhoneNumbers(ctx).FilterCountryCode(filterCountryCode).Execute()

Retrieve a list of phone numbers associated to orders



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
	filterCountryCode := "US" // string | Country code of the order phone number. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.RetrieveOrderPhoneNumbers(context.Background()).FilterCountryCode(filterCountryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.RetrieveOrderPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveOrderPhoneNumbers`: ListNumberOrderPhoneNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.RetrieveOrderPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOrderPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCountryCode** | **string** | Country code of the order phone number. | 

### Return type

[**ListNumberOrderPhoneNumbersResponse**](ListNumberOrderPhoneNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNumberOrder

> NumberOrderResponse UpdateNumberOrder(ctx, numberOrderId).UpdateNumberOrderRequest(updateNumberOrderRequest).Execute()

Update a number order



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
	numberOrderId := "numberOrderId_example" // string | The number order ID.
	updateNumberOrderRequest := *openapiclient.NewUpdateNumberOrderRequest() // UpdateNumberOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.UpdateNumberOrder(context.Background(), numberOrderId).UpdateNumberOrderRequest(updateNumberOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.UpdateNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNumberOrder`: NumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.UpdateNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberOrderId** | **string** | The number order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNumberOrderRequest** | [**UpdateNumberOrderRequest**](UpdateNumberOrderRequest.md) |  | 

### Return type

[**NumberOrderResponse**](NumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNumberOrderPhoneNumber

> NumberOrderPhoneNumberResponse UpdateNumberOrderPhoneNumber(ctx, numberOrderPhoneNumberId).UpdateNumberOrderPhoneNumberRequest(updateNumberOrderPhoneNumberRequest).Execute()

Update requirements for a single phone number within a number order.



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
	numberOrderPhoneNumberId := "numberOrderPhoneNumberId_example" // string | The number order phone number ID.
	updateNumberOrderPhoneNumberRequest := *openapiclient.NewUpdateNumberOrderPhoneNumberRequest() // UpdateNumberOrderPhoneNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.UpdateNumberOrderPhoneNumber(context.Background(), numberOrderPhoneNumberId).UpdateNumberOrderPhoneNumberRequest(updateNumberOrderPhoneNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.UpdateNumberOrderPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNumberOrderPhoneNumber`: NumberOrderPhoneNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.UpdateNumberOrderPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberOrderPhoneNumberId** | **string** | The number order phone number ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNumberOrderPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNumberOrderPhoneNumberRequest** | [**UpdateNumberOrderPhoneNumberRequest**](UpdateNumberOrderPhoneNumberRequest.md) |  | 

### Return type

[**NumberOrderPhoneNumberResponse**](NumberOrderPhoneNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubNumberOrder

> SubNumberOrderResponse UpdateSubNumberOrder(ctx, subNumberOrderId).UpdateSubNumberOrderRequest(updateSubNumberOrderRequest).Execute()

Update a sub number order's requirements



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
	subNumberOrderId := "subNumberOrderId_example" // string | The sub number order ID.
	updateSubNumberOrderRequest := *openapiclient.NewUpdateSubNumberOrderRequest() // UpdateSubNumberOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberOrdersAPI.UpdateSubNumberOrder(context.Background(), subNumberOrderId).UpdateSubNumberOrderRequest(updateSubNumberOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberOrdersAPI.UpdateSubNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubNumberOrder`: SubNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberOrdersAPI.UpdateSubNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subNumberOrderId** | **string** | The sub number order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubNumberOrderRequest** | [**UpdateSubNumberOrderRequest**](UpdateSubNumberOrderRequest.md) |  | 

### Return type

[**SubNumberOrderResponse**](SubNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

