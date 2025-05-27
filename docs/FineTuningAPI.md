# \FineTuningAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelNewFinetuningjobPublicFinetuningPost**](FineTuningAPI.md#CancelNewFinetuningjobPublicFinetuningPost) | **Post** /ai/fine_tuning/jobs/{job_id}/cancel | Cancel a fine tuning job
[**CreateNewFinetuningjobPublicFinetuningPost**](FineTuningAPI.md#CreateNewFinetuningjobPublicFinetuningPost) | **Post** /ai/fine_tuning/jobs | Create a fine tuning job
[**GetFinetuningjobPublicFinetuningGet**](FineTuningAPI.md#GetFinetuningjobPublicFinetuningGet) | **Get** /ai/fine_tuning/jobs | List fine tuning jobs
[**GetFinetuningjobPublicFinetuningJobIdGet**](FineTuningAPI.md#GetFinetuningjobPublicFinetuningJobIdGet) | **Get** /ai/fine_tuning/jobs/{job_id} | Get a fine tuning job



## CancelNewFinetuningjobPublicFinetuningPost

> FineTuningJob CancelNewFinetuningjobPublicFinetuningPost(ctx, jobId).Execute()

Cancel a fine tuning job



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.CancelNewFinetuningjobPublicFinetuningPost(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.CancelNewFinetuningjobPublicFinetuningPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelNewFinetuningjobPublicFinetuningPost`: FineTuningJob
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.CancelNewFinetuningjobPublicFinetuningPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelNewFinetuningjobPublicFinetuningPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FineTuningJob**](FineTuningJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewFinetuningjobPublicFinetuningPost

> FineTuningJob CreateNewFinetuningjobPublicFinetuningPost(ctx).CreateFineTuningJobRequest(createFineTuningJobRequest).Execute()

Create a fine tuning job



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
	createFineTuningJobRequest := *openapiclient.NewCreateFineTuningJobRequest("Model_example", "TrainingFile_example") // CreateFineTuningJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.CreateNewFinetuningjobPublicFinetuningPost(context.Background()).CreateFineTuningJobRequest(createFineTuningJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.CreateNewFinetuningjobPublicFinetuningPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewFinetuningjobPublicFinetuningPost`: FineTuningJob
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.CreateNewFinetuningjobPublicFinetuningPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewFinetuningjobPublicFinetuningPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFineTuningJobRequest** | [**CreateFineTuningJobRequest**](CreateFineTuningJobRequest.md) |  | 

### Return type

[**FineTuningJob**](FineTuningJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinetuningjobPublicFinetuningGet

> FineTuningJobsListData GetFinetuningjobPublicFinetuningGet(ctx).Execute()

List fine tuning jobs



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.GetFinetuningjobPublicFinetuningGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.GetFinetuningjobPublicFinetuningGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinetuningjobPublicFinetuningGet`: FineTuningJobsListData
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.GetFinetuningjobPublicFinetuningGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuningjobPublicFinetuningGetRequest struct via the builder pattern


### Return type

[**FineTuningJobsListData**](FineTuningJobsListData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinetuningjobPublicFinetuningJobIdGet

> FineTuningJob GetFinetuningjobPublicFinetuningJobIdGet(ctx, jobId).Execute()

Get a fine tuning job



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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FineTuningAPI.GetFinetuningjobPublicFinetuningJobIdGet(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FineTuningAPI.GetFinetuningjobPublicFinetuningJobIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinetuningjobPublicFinetuningJobIdGet`: FineTuningJob
	fmt.Fprintf(os.Stdout, "Response from `FineTuningAPI.GetFinetuningjobPublicFinetuningJobIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuningjobPublicFinetuningJobIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FineTuningJob**](FineTuningJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

