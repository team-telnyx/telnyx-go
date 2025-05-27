# \QueueCommandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQueueCalls**](QueueCommandsAPI.md#ListQueueCalls) | **Get** /queues/{queue_name}/calls | Retrieve calls from a queue
[**RetrieveCallFromQueue**](QueueCommandsAPI.md#RetrieveCallFromQueue) | **Get** /queues/{queue_name}/calls/{call_control_id} | Retrieve a call from a queue
[**RetrieveCallQueue**](QueueCommandsAPI.md#RetrieveCallQueue) | **Get** /queues/{queue_name} | Retrieve a call queue



## ListQueueCalls

> ListQueueCallsResponse ListQueueCalls(ctx, queueName).PageNumber(pageNumber).PageSize(pageSize).Execute()

Retrieve calls from a queue



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
	queueName := "queueName_example" // string | Uniquely identifies the queue by name
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueCommandsAPI.ListQueueCalls(context.Background(), queueName).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueCommandsAPI.ListQueueCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueueCalls`: ListQueueCallsResponse
	fmt.Fprintf(os.Stdout, "Response from `QueueCommandsAPI.ListQueueCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueName** | **string** | Uniquely identifies the queue by name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueueCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListQueueCallsResponse**](ListQueueCallsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCallFromQueue

> QueueCallResponse RetrieveCallFromQueue(ctx, queueName, callControlId).Execute()

Retrieve a call from a queue



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
	queueName := "queueName_example" // string | Uniquely identifies the queue by name
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueCommandsAPI.RetrieveCallFromQueue(context.Background(), queueName, callControlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueCommandsAPI.RetrieveCallFromQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCallFromQueue`: QueueCallResponse
	fmt.Fprintf(os.Stdout, "Response from `QueueCommandsAPI.RetrieveCallFromQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueName** | **string** | Uniquely identifies the queue by name | 
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCallFromQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QueueCallResponse**](QueueCallResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCallQueue

> QueueResponse RetrieveCallQueue(ctx, queueName).Execute()

Retrieve a call queue



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
	queueName := "queueName_example" // string | Uniquely identifies the queue by name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueCommandsAPI.RetrieveCallQueue(context.Background(), queueName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueCommandsAPI.RetrieveCallQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCallQueue`: QueueResponse
	fmt.Fprintf(os.Stdout, "Response from `QueueCommandsAPI.RetrieveCallQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueName** | **string** | Uniquely identifies the queue by name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCallQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueResponse**](QueueResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

