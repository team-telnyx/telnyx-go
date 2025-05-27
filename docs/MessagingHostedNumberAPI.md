# \MessagingHostedNumberAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckEligibilityNumbers**](MessagingHostedNumberAPI.md#CheckEligibilityNumbers) | **Get** /messaging_hosted_number_orders/eligibility_numbers_check | Check eligibility of phone numbers for hosted messaging
[**CreateMessagingHostedNumberOrder**](MessagingHostedNumberAPI.md#CreateMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders | Create a messaging hosted number order
[**CreateVerificationCodesForMessagingHostedNumberOrder**](MessagingHostedNumberAPI.md#CreateVerificationCodesForMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders/{id}/verification_codes | Create verification codes for the hosted numbers order
[**DeleteMessagingHostedNumber**](MessagingHostedNumberAPI.md#DeleteMessagingHostedNumber) | **Delete** /messaging_hosted_numbers/{id} | Delete a messaging hosted number
[**GetMessagingHostedNumberOrder**](MessagingHostedNumberAPI.md#GetMessagingHostedNumberOrder) | **Get** /messaging_hosted_number_orders/{id} | Retrieve a messaging hosted number order
[**ListMessagingHostedNumberOrders**](MessagingHostedNumberAPI.md#ListMessagingHostedNumberOrders) | **Get** /messaging_hosted_number_orders | List messaging hosted number orders
[**UploadMessagingHostedNumberOrderFile**](MessagingHostedNumberAPI.md#UploadMessagingHostedNumberOrderFile) | **Post** /messaging_hosted_number_orders/{id}/actions/file_upload | Upload file required for a messaging hosted number order
[**ValidateVerificationCodesForMessagingHostedNumberOrder**](MessagingHostedNumberAPI.md#ValidateVerificationCodesForMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders/{id}/validation_codes | Validate the verification codes for the hosted numbers order



## CheckEligibilityNumbers

> EligibilityNumbersResponse CheckEligibilityNumbers(ctx).PhoneNumbers(phoneNumbers).Execute()

Check eligibility of phone numbers for hosted messaging

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
	phoneNumbers := "phoneNumbers_example" // string | Comma-separated list of phone numbers to check eligibility

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.CheckEligibilityNumbers(context.Background()).PhoneNumbers(phoneNumbers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.CheckEligibilityNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckEligibilityNumbers`: EligibilityNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.CheckEligibilityNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckEligibilityNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumbers** | **string** | Comma-separated list of phone numbers to check eligibility | 

### Return type

[**EligibilityNumbersResponse**](EligibilityNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessagingHostedNumberOrder

> RetrieveMessagingHostedNumberOrderResponse CreateMessagingHostedNumberOrder(ctx).CreateMessagingHostedNumberOrderRequest(createMessagingHostedNumberOrderRequest).Execute()

Create a messaging hosted number order

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
	createMessagingHostedNumberOrderRequest := *openapiclient.NewCreateMessagingHostedNumberOrderRequest() // CreateMessagingHostedNumberOrderRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.CreateMessagingHostedNumberOrder(context.Background()).CreateMessagingHostedNumberOrderRequest(createMessagingHostedNumberOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.CreateMessagingHostedNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessagingHostedNumberOrder`: RetrieveMessagingHostedNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.CreateMessagingHostedNumberOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessagingHostedNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMessagingHostedNumberOrderRequest** | [**CreateMessagingHostedNumberOrderRequest**](CreateMessagingHostedNumberOrderRequest.md) | Message payload | 

### Return type

[**RetrieveMessagingHostedNumberOrderResponse**](RetrieveMessagingHostedNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationCodesForMessagingHostedNumberOrder

> interface{} CreateVerificationCodesForMessagingHostedNumberOrder(ctx, id).VerificationCodesRequest(verificationCodesRequest).Execute()

Create verification codes for the hosted numbers order



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
	id := "id_example" // string | Order ID to have a verification code created.
	verificationCodesRequest := *openapiclient.NewVerificationCodesRequest([]string{"PhoneNumbers_example"}, "VerificationMethod_example") // VerificationCodesRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.CreateVerificationCodesForMessagingHostedNumberOrder(context.Background(), id).VerificationCodesRequest(verificationCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.CreateVerificationCodesForMessagingHostedNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationCodesForMessagingHostedNumberOrder`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.CreateVerificationCodesForMessagingHostedNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Order ID to have a verification code created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationCodesForMessagingHostedNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verificationCodesRequest** | [**VerificationCodesRequest**](VerificationCodesRequest.md) | Message payload | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessagingHostedNumber

> RetrieveMessagingHostedNumberResponse1 DeleteMessagingHostedNumber(ctx, id).Execute()

Delete a messaging hosted number

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
	id := "id_example" // string | Identifies the type of resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.DeleteMessagingHostedNumber(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.DeleteMessagingHostedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessagingHostedNumber`: RetrieveMessagingHostedNumberResponse1
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.DeleteMessagingHostedNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessagingHostedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveMessagingHostedNumberResponse1**](RetrieveMessagingHostedNumberResponse1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagingHostedNumberOrder

> RetrieveMessagingHostedNumberOrderResponse GetMessagingHostedNumberOrder(ctx, id).Execute()

Retrieve a messaging hosted number order

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
	id := "id_example" // string | Identifies the type of resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.GetMessagingHostedNumberOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.GetMessagingHostedNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagingHostedNumberOrder`: RetrieveMessagingHostedNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.GetMessagingHostedNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagingHostedNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveMessagingHostedNumberOrderResponse**](RetrieveMessagingHostedNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingHostedNumberOrders

> ListMessagingHostedNumberOrderResponse ListMessagingHostedNumberOrders(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List messaging hosted number orders

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.ListMessagingHostedNumberOrders(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.ListMessagingHostedNumberOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessagingHostedNumberOrders`: ListMessagingHostedNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.ListMessagingHostedNumberOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagingHostedNumberOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListMessagingHostedNumberOrderResponse**](ListMessagingHostedNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMessagingHostedNumberOrderFile

> RetrieveMessagingHostedNumberOrderResponse UploadMessagingHostedNumberOrderFile(ctx, id).Loa(loa).Bill(bill).Execute()

Upload file required for a messaging hosted number order

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
	id := "id_example" // string | Identifies the type of resource.
	loa := os.NewFile(1234, "some_file") // *os.File | Must be a signed LOA for the numbers in the order in PDF format. (optional)
	bill := os.NewFile(1234, "some_file") // *os.File | Must be the last month's bill with proof of ownership of all of the numbers in the order in PDF format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.UploadMessagingHostedNumberOrderFile(context.Background(), id).Loa(loa).Bill(bill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.UploadMessagingHostedNumberOrderFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMessagingHostedNumberOrderFile`: RetrieveMessagingHostedNumberOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.UploadMessagingHostedNumberOrderFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadMessagingHostedNumberOrderFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loa** | ***os.File** | Must be a signed LOA for the numbers in the order in PDF format. | 
 **bill** | ***os.File** | Must be the last month&#39;s bill with proof of ownership of all of the numbers in the order in PDF format. | 

### Return type

[**RetrieveMessagingHostedNumberOrderResponse**](RetrieveMessagingHostedNumberOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateVerificationCodesForMessagingHostedNumberOrder

> RetrieveMessagingHostedNumberResponse ValidateVerificationCodesForMessagingHostedNumberOrder(ctx, id).ValidationCodesRequest(validationCodesRequest).Execute()

Validate the verification codes for the hosted numbers order



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
	id := "id_example" // string | Order ID related to the validation codes.
	validationCodesRequest := *openapiclient.NewValidationCodesRequest([]openapiclient.ValidationCodesRequestVerificationCodesInner{*openapiclient.NewValidationCodesRequestVerificationCodesInner("PhoneNumber_example", "Code_example")}) // ValidationCodesRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingHostedNumberAPI.ValidateVerificationCodesForMessagingHostedNumberOrder(context.Background(), id).ValidationCodesRequest(validationCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingHostedNumberAPI.ValidateVerificationCodesForMessagingHostedNumberOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateVerificationCodesForMessagingHostedNumberOrder`: RetrieveMessagingHostedNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingHostedNumberAPI.ValidateVerificationCodesForMessagingHostedNumberOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Order ID related to the validation codes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateVerificationCodesForMessagingHostedNumberOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validationCodesRequest** | [**ValidationCodesRequest**](ValidationCodesRequest.md) | Message payload | 

### Return type

[**RetrieveMessagingHostedNumberResponse**](RetrieveMessagingHostedNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

