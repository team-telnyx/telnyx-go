# \CSVDownloadsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCsvDownload**](CSVDownloadsAPI.md#CreateCsvDownload) | **Post** /phone_numbers/csv_downloads | Create a CSV download
[**GetCsvDownload**](CSVDownloadsAPI.md#GetCsvDownload) | **Get** /phone_numbers/csv_downloads/{id} | Retrieve a CSV download
[**ListCsvDownloads**](CSVDownloadsAPI.md#ListCsvDownloads) | **Get** /phone_numbers/csv_downloads | List CSV downloads



## CreateCsvDownload

> CSVDownloadResponse CreateCsvDownload(ctx).CsvFormat(csvFormat).FilterHasBundle(filterHasBundle).FilterTag(filterTag).FilterConnectionId(filterConnectionId).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).Execute()

Create a CSV download

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
	csvFormat := "V2" // string | Which format to use when generating the CSV file. The default for backwards compatibility is 'V1' (optional) (default to "V1")
	filterHasBundle := "filterHasBundle_example" // string | Filter by phone number that have bundles. (optional)
	filterTag := "filterTag_example" // string | Filter by phone number tags. (optional)
	filterConnectionId := "1521916448077776306" // string | Filter by connection_id. (optional)
	filterPhoneNumber := "filterPhoneNumber_example" // string | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. (optional)
	filterStatus := "active" // string | Filter by phone number status. (optional)
	filterVoiceConnectionNameContains := "test" // string | Filter contains connection name. Requires at least three characters. (optional)
	filterVoiceUsagePaymentMethod := "channel" // string | Filter by usage_payment_method. (optional)
	filterBillingGroupId := "62e4bf2e-c278-4282-b524-488d9c9c43b2" // string | Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string 'null'. (optional)
	filterEmergencyAddressId := "9102160989215728032" // string | Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string 'null'. (optional)
	filterCustomerReference := "filterCustomerReference_example" // string | Filter numbers via the customer_reference set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CSVDownloadsAPI.CreateCsvDownload(context.Background()).CsvFormat(csvFormat).FilterHasBundle(filterHasBundle).FilterTag(filterTag).FilterConnectionId(filterConnectionId).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSVDownloadsAPI.CreateCsvDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCsvDownload`: CSVDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `CSVDownloadsAPI.CreateCsvDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCsvDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csvFormat** | **string** | Which format to use when generating the CSV file. The default for backwards compatibility is &#39;V1&#39; | [default to &quot;V1&quot;]
 **filterHasBundle** | **string** | Filter by phone number that have bundles. | 
 **filterTag** | **string** | Filter by phone number tags. | 
 **filterConnectionId** | **string** | Filter by connection_id. | 
 **filterPhoneNumber** | **string** | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterStatus** | **string** | Filter by phone number status. | 
 **filterVoiceConnectionNameContains** | **string** | Filter contains connection name. Requires at least three characters. | 
 **filterVoiceUsagePaymentMethod** | **string** | Filter by usage_payment_method. | 
 **filterBillingGroupId** | **string** | Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string &#39;null&#39;. | 
 **filterEmergencyAddressId** | **string** | Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string &#39;null&#39;. | 
 **filterCustomerReference** | **string** | Filter numbers via the customer_reference set. | 

### Return type

[**CSVDownloadResponse**](CSVDownloadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsvDownload

> CSVDownloadResponse GetCsvDownload(ctx, id).Execute()

Retrieve a CSV download

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
	id := "id_example" // string | Identifies the CSV download.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CSVDownloadsAPI.GetCsvDownload(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSVDownloadsAPI.GetCsvDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsvDownload`: CSVDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `CSVDownloadsAPI.GetCsvDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the CSV download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsvDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CSVDownloadResponse**](CSVDownloadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCsvDownloads

> ListCSVDownloadsResponse ListCsvDownloads(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List CSV downloads

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
	resp, r, err := apiClient.CSVDownloadsAPI.ListCsvDownloads(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSVDownloadsAPI.ListCsvDownloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCsvDownloads`: ListCSVDownloadsResponse
	fmt.Fprintf(os.Stdout, "Response from `CSVDownloadsAPI.ListCsvDownloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCsvDownloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListCSVDownloadsResponse**](ListCSVDownloadsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

