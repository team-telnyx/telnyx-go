# \BucketUsageAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBucketUsage**](BucketUsageAPI.md#GetBucketUsage) | **Get** /storage/buckets/{bucketName}/usage/storage | Get Bucket Usage
[**GetStorageAPIUsage**](BucketUsageAPI.md#GetStorageAPIUsage) | **Get** /storage/buckets/{bucketName}/usage/api | Get API Usage



## GetBucketUsage

> GetBucketUsage200Response GetBucketUsage(ctx, bucketName).Execute()

Get Bucket Usage



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
	bucketName := "bucketName_example" // string | The name of the bucket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketUsageAPI.GetBucketUsage(context.Background(), bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketUsageAPI.GetBucketUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBucketUsage`: GetBucketUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `BucketUsageAPI.GetBucketUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBucketUsage200Response**](GetBucketUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageAPIUsage

> GetStorageAPIUsage200Response GetStorageAPIUsage(ctx, bucketName).FilterStartTime(filterStartTime).FilterEndTime(filterEndTime).Execute()

Get API Usage



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
	bucketName := "bucketName_example" // string | The name of the bucket
	filterStartTime := time.Now() // time.Time | The start time of the period to filter the usage (ISO microsecond format)
	filterEndTime := time.Now() // time.Time | The end time of the period to filter the usage (ISO microsecond format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketUsageAPI.GetStorageAPIUsage(context.Background(), bucketName).FilterStartTime(filterStartTime).FilterEndTime(filterEndTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketUsageAPI.GetStorageAPIUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageAPIUsage`: GetStorageAPIUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `BucketUsageAPI.GetStorageAPIUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageAPIUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterStartTime** | **time.Time** | The start time of the period to filter the usage (ISO microsecond format) | 
 **filterEndTime** | **time.Time** | The end time of the period to filter the usage (ISO microsecond format) | 

### Return type

[**GetStorageAPIUsage200Response**](GetStorageAPIUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

