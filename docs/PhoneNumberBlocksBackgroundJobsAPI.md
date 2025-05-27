# \PhoneNumberBlocksBackgroundJobsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneNumberBlockDeletionJob**](PhoneNumberBlocksBackgroundJobsAPI.md#CreatePhoneNumberBlockDeletionJob) | **Post** /phone_number_blocks/jobs/delete_phone_number_block | Deletes all numbers associated with a phone number block
[**GetPhoneNumberBlocksJob**](PhoneNumberBlocksBackgroundJobsAPI.md#GetPhoneNumberBlocksJob) | **Get** /phone_number_blocks/jobs/{id} | Retrieves a phone number blocks job
[**ListPhoneNumberBlocksJobs**](PhoneNumberBlocksBackgroundJobsAPI.md#ListPhoneNumberBlocksJobs) | **Get** /phone_number_blocks/jobs | Lists the phone number blocks jobs



## CreatePhoneNumberBlockDeletionJob

> PhoneNumberBlocksJobDeletePhoneNumberBlock CreatePhoneNumberBlockDeletionJob(ctx).PhoneNumberBlocksJobDeletePhoneNumberBlockRequest(phoneNumberBlocksJobDeletePhoneNumberBlockRequest).Execute()

Deletes all numbers associated with a phone number block



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
	phoneNumberBlocksJobDeletePhoneNumberBlockRequest := *openapiclient.NewPhoneNumberBlocksJobDeletePhoneNumberBlockRequest("PhoneNumberBlockId_example") // PhoneNumberBlocksJobDeletePhoneNumberBlockRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberBlocksBackgroundJobsAPI.CreatePhoneNumberBlockDeletionJob(context.Background()).PhoneNumberBlocksJobDeletePhoneNumberBlockRequest(phoneNumberBlocksJobDeletePhoneNumberBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberBlocksBackgroundJobsAPI.CreatePhoneNumberBlockDeletionJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePhoneNumberBlockDeletionJob`: PhoneNumberBlocksJobDeletePhoneNumberBlock
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberBlocksBackgroundJobsAPI.CreatePhoneNumberBlockDeletionJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhoneNumberBlockDeletionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumberBlocksJobDeletePhoneNumberBlockRequest** | [**PhoneNumberBlocksJobDeletePhoneNumberBlockRequest**](PhoneNumberBlocksJobDeletePhoneNumberBlockRequest.md) |  | 

### Return type

[**PhoneNumberBlocksJobDeletePhoneNumberBlock**](PhoneNumberBlocksJobDeletePhoneNumberBlock.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhoneNumberBlocksJob

> PhoneNumberBlocksJob GetPhoneNumberBlocksJob(ctx, id).Execute()

Retrieves a phone number blocks job

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
	id := "id_example" // string | Identifies the Phone Number Blocks Job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberBlocksBackgroundJobsAPI.GetPhoneNumberBlocksJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberBlocksBackgroundJobsAPI.GetPhoneNumberBlocksJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPhoneNumberBlocksJob`: PhoneNumberBlocksJob
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberBlocksBackgroundJobsAPI.GetPhoneNumberBlocksJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the Phone Number Blocks Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhoneNumberBlocksJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneNumberBlocksJob**](PhoneNumberBlocksJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumberBlocksJobs

> ListPhoneNumberBlocksBackgroundJobsResponse ListPhoneNumberBlocksJobs(ctx).FilterType(filterType).FilterStatus(filterStatus).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()

Lists the phone number blocks jobs

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
	filterType := "delete_phone_number_block" // string | Filter the phone number blocks jobs by type. (optional)
	filterStatus := "in_progress" // string | Filter the phone number blocks jobs by status. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberBlocksBackgroundJobsAPI.ListPhoneNumberBlocksJobs(context.Background()).FilterType(filterType).FilterStatus(filterStatus).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberBlocksBackgroundJobsAPI.ListPhoneNumberBlocksJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPhoneNumberBlocksJobs`: ListPhoneNumberBlocksBackgroundJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberBlocksBackgroundJobsAPI.ListPhoneNumberBlocksJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhoneNumberBlocksJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterType** | **string** | Filter the phone number blocks jobs by type. | 
 **filterStatus** | **string** | Filter the phone number blocks jobs by status. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumberBlocksBackgroundJobsResponse**](ListPhoneNumberBlocksBackgroundJobsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

