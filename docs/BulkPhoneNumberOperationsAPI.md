# \BulkPhoneNumberOperationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeletePhoneNumbersJob**](BulkPhoneNumberOperationsAPI.md#CreateDeletePhoneNumbersJob) | **Post** /phone_numbers/jobs/delete_phone_numbers | Delete a batch of numbers
[**CreatePhoneNumbersJobUpdateEmergencySettings**](BulkPhoneNumberOperationsAPI.md#CreatePhoneNumbersJobUpdateEmergencySettings) | **Post** /phone_numbers/jobs/update_emergency_settings | Update the emergency settings from a batch of numbers
[**CreateUpdatePhoneNumbersJob**](BulkPhoneNumberOperationsAPI.md#CreateUpdatePhoneNumbersJob) | **Post** /phone_numbers/jobs/update_phone_numbers | Update a batch of numbers
[**ListPhoneNumbersJobs**](BulkPhoneNumberOperationsAPI.md#ListPhoneNumbersJobs) | **Get** /phone_numbers/jobs | Lists the phone numbers jobs
[**RetrievePhoneNumbersJob**](BulkPhoneNumberOperationsAPI.md#RetrievePhoneNumbersJob) | **Get** /phone_numbers/jobs/{id} | Retrieve a phone numbers job



## CreateDeletePhoneNumbersJob

> PhoneNumbersJobDeletePhoneNumbers CreateDeletePhoneNumbersJob(ctx).PhoneNumbersJobDeletePhoneNumbersRequest(phoneNumbersJobDeletePhoneNumbersRequest).Execute()

Delete a batch of numbers



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
	phoneNumbersJobDeletePhoneNumbersRequest := *openapiclient.NewPhoneNumbersJobDeletePhoneNumbersRequest([]string{"PhoneNumbers_example"}) // PhoneNumbersJobDeletePhoneNumbersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberOperationsAPI.CreateDeletePhoneNumbersJob(context.Background()).PhoneNumbersJobDeletePhoneNumbersRequest(phoneNumbersJobDeletePhoneNumbersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberOperationsAPI.CreateDeletePhoneNumbersJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeletePhoneNumbersJob`: PhoneNumbersJobDeletePhoneNumbers
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberOperationsAPI.CreateDeletePhoneNumbersJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeletePhoneNumbersJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumbersJobDeletePhoneNumbersRequest** | [**PhoneNumbersJobDeletePhoneNumbersRequest**](PhoneNumbersJobDeletePhoneNumbersRequest.md) |  | 

### Return type

[**PhoneNumbersJobDeletePhoneNumbers**](PhoneNumbersJobDeletePhoneNumbers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumbersJobUpdateEmergencySettings

> PhoneNumbersEnableEmergency CreatePhoneNumbersJobUpdateEmergencySettings(ctx).PhoneNumbersJobUpdateEmergencySettingsRequest(phoneNumbersJobUpdateEmergencySettingsRequest).Execute()

Update the emergency settings from a batch of numbers



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
	phoneNumbersJobUpdateEmergencySettingsRequest := *openapiclient.NewPhoneNumbersJobUpdateEmergencySettingsRequest([]string{"PhoneNumbers_example"}, false) // PhoneNumbersJobUpdateEmergencySettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberOperationsAPI.CreatePhoneNumbersJobUpdateEmergencySettings(context.Background()).PhoneNumbersJobUpdateEmergencySettingsRequest(phoneNumbersJobUpdateEmergencySettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberOperationsAPI.CreatePhoneNumbersJobUpdateEmergencySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePhoneNumbersJobUpdateEmergencySettings`: PhoneNumbersEnableEmergency
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberOperationsAPI.CreatePhoneNumbersJobUpdateEmergencySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhoneNumbersJobUpdateEmergencySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumbersJobUpdateEmergencySettingsRequest** | [**PhoneNumbersJobUpdateEmergencySettingsRequest**](PhoneNumbersJobUpdateEmergencySettingsRequest.md) |  | 

### Return type

[**PhoneNumbersEnableEmergency**](PhoneNumbersEnableEmergency.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpdatePhoneNumbersJob

> PhoneNumbersJobUpdatePhoneNumbers CreateUpdatePhoneNumbersJob(ctx).PhoneNumbersJobUpdatePhoneNumbersRequest(phoneNumbersJobUpdatePhoneNumbersRequest).FilterHasBundle(filterHasBundle).FilterTag(filterTag).FilterConnectionId(filterConnectionId).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).Execute()

Update a batch of numbers



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
	phoneNumbersJobUpdatePhoneNumbersRequest := *openapiclient.NewPhoneNumbersJobUpdatePhoneNumbersRequest([]string{"PhoneNumbers_example"}) // PhoneNumbersJobUpdatePhoneNumbersRequest | 
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
	resp, r, err := apiClient.BulkPhoneNumberOperationsAPI.CreateUpdatePhoneNumbersJob(context.Background()).PhoneNumbersJobUpdatePhoneNumbersRequest(phoneNumbersJobUpdatePhoneNumbersRequest).FilterHasBundle(filterHasBundle).FilterTag(filterTag).FilterConnectionId(filterConnectionId).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberOperationsAPI.CreateUpdatePhoneNumbersJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpdatePhoneNumbersJob`: PhoneNumbersJobUpdatePhoneNumbers
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberOperationsAPI.CreateUpdatePhoneNumbersJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdatePhoneNumbersJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumbersJobUpdatePhoneNumbersRequest** | [**PhoneNumbersJobUpdatePhoneNumbersRequest**](PhoneNumbersJobUpdatePhoneNumbersRequest.md) |  | 
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

[**PhoneNumbersJobUpdatePhoneNumbers**](PhoneNumbersJobUpdatePhoneNumbers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumbersJobs

> ListPhoneNumbersBackgroundJobsResponse ListPhoneNumbersJobs(ctx).FilterType(filterType).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()

Lists the phone numbers jobs

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
	filterType := "update_emergency_settings" // string | Filter the phone number jobs by type. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberOperationsAPI.ListPhoneNumbersJobs(context.Background()).FilterType(filterType).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberOperationsAPI.ListPhoneNumbersJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPhoneNumbersJobs`: ListPhoneNumbersBackgroundJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberOperationsAPI.ListPhoneNumbersJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhoneNumbersJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterType** | **string** | Filter the phone number jobs by type. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersBackgroundJobsResponse**](ListPhoneNumbersBackgroundJobsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePhoneNumbersJob

> PhoneNumbersJob RetrievePhoneNumbersJob(ctx, id).Execute()

Retrieve a phone numbers job

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
	id := "id_example" // string | Identifies the Phone Numbers Job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberOperationsAPI.RetrievePhoneNumbersJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberOperationsAPI.RetrievePhoneNumbersJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievePhoneNumbersJob`: PhoneNumbersJob
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberOperationsAPI.RetrievePhoneNumbersJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the Phone Numbers Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePhoneNumbersJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneNumbersJob**](PhoneNumbersJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

