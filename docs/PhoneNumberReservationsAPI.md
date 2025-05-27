# \PhoneNumberReservationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberReservation**](PhoneNumberReservationsAPI.md#CreateNumberReservation) | **Post** /number_reservations | Create a number reservation
[**ExtendNumberReservationExpiryTime**](PhoneNumberReservationsAPI.md#ExtendNumberReservationExpiryTime) | **Post** /number_reservations/{number_reservation_id}/actions/extend | Extend a number reservation
[**ListNumberReservations**](PhoneNumberReservationsAPI.md#ListNumberReservations) | **Get** /number_reservations | List number reservations
[**RetrieveNumberReservation**](PhoneNumberReservationsAPI.md#RetrieveNumberReservation) | **Get** /number_reservations/{number_reservation_id} | Retrieve a number reservation



## CreateNumberReservation

> NumberReservationResponse CreateNumberReservation(ctx).CreateNumberReservationRequest(createNumberReservationRequest).Execute()

Create a number reservation



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
	createNumberReservationRequest := *openapiclient.NewCreateNumberReservationRequest() // CreateNumberReservationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberReservationsAPI.CreateNumberReservation(context.Background()).CreateNumberReservationRequest(createNumberReservationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberReservationsAPI.CreateNumberReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNumberReservation`: NumberReservationResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberReservationsAPI.CreateNumberReservation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNumberReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNumberReservationRequest** | [**CreateNumberReservationRequest**](CreateNumberReservationRequest.md) |  | 

### Return type

[**NumberReservationResponse**](NumberReservationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendNumberReservationExpiryTime

> NumberReservationResponse ExtendNumberReservationExpiryTime(ctx, numberReservationId).Execute()

Extend a number reservation



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
	numberReservationId := "numberReservationId_example" // string | The number reservation ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberReservationsAPI.ExtendNumberReservationExpiryTime(context.Background(), numberReservationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberReservationsAPI.ExtendNumberReservationExpiryTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendNumberReservationExpiryTime`: NumberReservationResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberReservationsAPI.ExtendNumberReservationExpiryTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberReservationId** | **string** | The number reservation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendNumberReservationExpiryTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NumberReservationResponse**](NumberReservationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNumberReservations

> ListNumberReservationsResponse ListNumberReservations(ctx).FilterStatus(filterStatus).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).FilterPhoneNumbersPhoneNumber(filterPhoneNumbersPhoneNumber).FilterCustomerReference(filterCustomerReference).PageNumber(pageNumber).PageSize(pageSize).Execute()

List number reservations



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
	filterStatus := "filterStatus_example" // string | Filter number reservations by status. (optional)
	filterCreatedAtGt := "filterCreatedAtGt_example" // string | Filter number reservations later than this value. (optional)
	filterCreatedAtLt := "filterCreatedAtLt_example" // string | Filter number reservations earlier than this value. (optional)
	filterPhoneNumbersPhoneNumber := "filterPhoneNumbersPhoneNumber_example" // string | Filter number reservations having these phone numbers. (optional)
	filterCustomerReference := "filterCustomerReference_example" // string | Filter number reservations via the customer reference set. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberReservationsAPI.ListNumberReservations(context.Background()).FilterStatus(filterStatus).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).FilterPhoneNumbersPhoneNumber(filterPhoneNumbersPhoneNumber).FilterCustomerReference(filterCustomerReference).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberReservationsAPI.ListNumberReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNumberReservations`: ListNumberReservationsResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberReservationsAPI.ListNumberReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNumberReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **string** | Filter number reservations by status. | 
 **filterCreatedAtGt** | **string** | Filter number reservations later than this value. | 
 **filterCreatedAtLt** | **string** | Filter number reservations earlier than this value. | 
 **filterPhoneNumbersPhoneNumber** | **string** | Filter number reservations having these phone numbers. | 
 **filterCustomerReference** | **string** | Filter number reservations via the customer reference set. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListNumberReservationsResponse**](ListNumberReservationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNumberReservation

> NumberReservationResponse RetrieveNumberReservation(ctx, numberReservationId).Execute()

Retrieve a number reservation



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
	numberReservationId := "numberReservationId_example" // string | The number reservation ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberReservationsAPI.RetrieveNumberReservation(context.Background(), numberReservationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberReservationsAPI.RetrieveNumberReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveNumberReservation`: NumberReservationResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberReservationsAPI.RetrieveNumberReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberReservationId** | **string** | The number reservation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNumberReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NumberReservationResponse**](NumberReservationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

