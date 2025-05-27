# \CustomerServiceRecordAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerServiceRecord**](CustomerServiceRecordAPI.md#CreateCustomerServiceRecord) | **Post** /customer_service_records | Create a customer service record
[**GetCustomerServiceRecord**](CustomerServiceRecordAPI.md#GetCustomerServiceRecord) | **Get** /customer_service_records/{customer_service_record_id} | Get a customer service record
[**ListCustomerServiceRecords**](CustomerServiceRecordAPI.md#ListCustomerServiceRecords) | **Get** /customer_service_records | List customer service records
[**VerifyPhoneNumberCoverage**](CustomerServiceRecordAPI.md#VerifyPhoneNumberCoverage) | **Post** /customer_service_records/phone_number_coverages | Verify CSR phone number coverage



## CreateCustomerServiceRecord

> CreateCustomerServiceRecord201Response CreateCustomerServiceRecord(ctx).CreateCustomerServiceRecordRequest(createCustomerServiceRecordRequest).Execute()

Create a customer service record



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
	createCustomerServiceRecordRequest := *openapiclient.NewCreateCustomerServiceRecordRequest("+1234567890") // CreateCustomerServiceRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerServiceRecordAPI.CreateCustomerServiceRecord(context.Background()).CreateCustomerServiceRecordRequest(createCustomerServiceRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceRecordAPI.CreateCustomerServiceRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerServiceRecord`: CreateCustomerServiceRecord201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerServiceRecordAPI.CreateCustomerServiceRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerServiceRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerServiceRecordRequest** | [**CreateCustomerServiceRecordRequest**](CreateCustomerServiceRecordRequest.md) |  | 

### Return type

[**CreateCustomerServiceRecord201Response**](CreateCustomerServiceRecord201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerServiceRecord

> GetCustomerServiceRecord201Response GetCustomerServiceRecord(ctx, customerServiceRecordId).Execute()

Get a customer service record



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
	customerServiceRecordId := "customerServiceRecordId_example" // string | The ID of the customer service record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerServiceRecordAPI.GetCustomerServiceRecord(context.Background(), customerServiceRecordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceRecordAPI.GetCustomerServiceRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerServiceRecord`: GetCustomerServiceRecord201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerServiceRecordAPI.GetCustomerServiceRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerServiceRecordId** | **string** | The ID of the customer service record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerServiceRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomerServiceRecord201Response**](GetCustomerServiceRecord201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerServiceRecords

> ListCustomerServiceRecords200Response ListCustomerServiceRecords(ctx).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberIn(filterPhoneNumberIn).FilterStatusEq(filterStatusEq).FilterStatusIn(filterStatusIn).FilterCreatedAtLt(filterCreatedAtLt).FilterCreatedAtGt(filterCreatedAtGt).Execute()

List customer service records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)
	filterPhoneNumberEq := "+12441239999" // string | Filters records to those with a specified number. (optional)
	filterPhoneNumberIn := []string{"+12441239999"} // []string | Filters records to those with at least one number in the list. (optional)
	filterStatusEq := "pending" // string | Filters records to those with a specific status. (optional)
	filterStatusIn := []string{"pending"} // []string | Filters records to those with a least one status in the list. (optional)
	filterCreatedAtLt := time.Now() // time.Time | Filters records to those created before a specific date. (optional)
	filterCreatedAtGt := time.Now() // time.Time | Filters records to those created after a specific date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerServiceRecordAPI.ListCustomerServiceRecords(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberIn(filterPhoneNumberIn).FilterStatusEq(filterStatusEq).FilterStatusIn(filterStatusIn).FilterCreatedAtLt(filterCreatedAtLt).FilterCreatedAtGt(filterCreatedAtGt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceRecordAPI.ListCustomerServiceRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomerServiceRecords`: ListCustomerServiceRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerServiceRecordAPI.ListCustomerServiceRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerServiceRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 
 **filterPhoneNumberEq** | **string** | Filters records to those with a specified number. | 
 **filterPhoneNumberIn** | **[]string** | Filters records to those with at least one number in the list. | 
 **filterStatusEq** | **string** | Filters records to those with a specific status. | 
 **filterStatusIn** | **[]string** | Filters records to those with a least one status in the list. | 
 **filterCreatedAtLt** | **time.Time** | Filters records to those created before a specific date. | 
 **filterCreatedAtGt** | **time.Time** | Filters records to those created after a specific date. | 

### Return type

[**ListCustomerServiceRecords200Response**](ListCustomerServiceRecords200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPhoneNumberCoverage

> VerifyPhoneNumberCoverage201Response VerifyPhoneNumberCoverage(ctx).VerifyPhoneNumberCoverageRequest(verifyPhoneNumberCoverageRequest).Execute()

Verify CSR phone number coverage



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
	verifyPhoneNumberCoverageRequest := *openapiclient.NewVerifyPhoneNumberCoverageRequest([]string{"+1234567890"}) // VerifyPhoneNumberCoverageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerServiceRecordAPI.VerifyPhoneNumberCoverage(context.Background()).VerifyPhoneNumberCoverageRequest(verifyPhoneNumberCoverageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceRecordAPI.VerifyPhoneNumberCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPhoneNumberCoverage`: VerifyPhoneNumberCoverage201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerServiceRecordAPI.VerifyPhoneNumberCoverage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPhoneNumberCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyPhoneNumberCoverageRequest** | [**VerifyPhoneNumberCoverageRequest**](VerifyPhoneNumberCoverageRequest.md) |  | 

### Return type

[**VerifyPhoneNumberCoverage201Response**](VerifyPhoneNumberCoverage201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

