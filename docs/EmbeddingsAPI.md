# \EmbeddingsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete**](EmbeddingsAPI.md#EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete) | **Delete** /ai/embeddings/buckets/{bucket_name} | Disable AI for an Embedded Bucket
[**GetBucketName**](EmbeddingsAPI.md#GetBucketName) | **Get** /ai/embeddings/buckets/{bucket_name} | Get file-level embedding statuses for a bucket
[**GetEmbeddingBuckets**](EmbeddingsAPI.md#GetEmbeddingBuckets) | **Get** /ai/embeddings/buckets | List embedded buckets
[**GetEmbeddingTask**](EmbeddingsAPI.md#GetEmbeddingTask) | **Get** /ai/embeddings/{task_id} | Get an embedding task&#39;s status
[**GetTasksByStatus**](EmbeddingsAPI.md#GetTasksByStatus) | **Get** /ai/embeddings | Get Tasks by Status
[**PostEmbedding**](EmbeddingsAPI.md#PostEmbedding) | **Post** /ai/embeddings | Embed documents
[**PostEmbeddingSimilaritySearch**](EmbeddingsAPI.md#PostEmbeddingSimilaritySearch) | **Post** /ai/embeddings/similarity-search | Search for documents
[**PostEmbeddingUrl**](EmbeddingsAPI.md#PostEmbeddingUrl) | **Post** /ai/embeddings/url | Embed URL content



## EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete

> EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete(ctx, bucketName).Execute()

Disable AI for an Embedded Bucket



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
	bucketName := "bucketName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbeddingsAPI.EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete(context.Background(), bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketName

> EmbeddingsBucketFilesData GetBucketName(ctx, bucketName).Execute()

Get file-level embedding statuses for a bucket



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
	bucketName := "bucketName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.GetBucketName(context.Background(), bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.GetBucketName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBucketName`: EmbeddingsBucketFilesData
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.GetBucketName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmbeddingsBucketFilesData**](EmbeddingsBucketFilesData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmbeddingBuckets

> UserEmbeddedBucketsData GetEmbeddingBuckets(ctx).Execute()

List embedded buckets



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
	resp, r, err := apiClient.EmbeddingsAPI.GetEmbeddingBuckets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.GetEmbeddingBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbeddingBuckets`: UserEmbeddedBucketsData
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.GetEmbeddingBuckets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbeddingBucketsRequest struct via the builder pattern


### Return type

[**UserEmbeddedBucketsData**](UserEmbeddedBucketsData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmbeddingTask

> TaskStatusResponse GetEmbeddingTask(ctx, taskId).Execute()

Get an embedding task's status



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
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.GetEmbeddingTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.GetEmbeddingTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbeddingTask`: TaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.GetEmbeddingTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbeddingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskStatusResponse**](TaskStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksByStatus

> BackgroundTasksQueryResponseData GetTasksByStatus(ctx).Status(status).Execute()

Get Tasks by Status



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
	status := []string{"Inner_example"} // []string | List of task statuses i.e. `status=queued&status=processing` (optional) (default to ["processing","queued"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.GetTasksByStatus(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.GetTasksByStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasksByStatus`: BackgroundTasksQueryResponseData
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.GetTasksByStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | List of task statuses i.e. &#x60;status&#x3D;queued&amp;status&#x3D;processing&#x60; | [default to [&quot;processing&quot;,&quot;queued&quot;]]

### Return type

[**BackgroundTasksQueryResponseData**](BackgroundTasksQueryResponseData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEmbedding

> EmbeddingResponse PostEmbedding(ctx).EmbeddingBucketRequest(embeddingBucketRequest).Execute()

Embed documents



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
	embeddingBucketRequest := *openapiclient.NewEmbeddingBucketRequest("BucketName_example") // EmbeddingBucketRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.PostEmbedding(context.Background()).EmbeddingBucketRequest(embeddingBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.PostEmbedding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEmbedding`: EmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.PostEmbedding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEmbeddingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingBucketRequest** | [**EmbeddingBucketRequest**](EmbeddingBucketRequest.md) |  | 

### Return type

[**EmbeddingResponse**](EmbeddingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEmbeddingSimilaritySearch

> EmbeddingSimilaritySearchResponse PostEmbeddingSimilaritySearch(ctx).EmbeddingSimilaritySearchRequest(embeddingSimilaritySearchRequest).Execute()

Search for documents



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
	embeddingSimilaritySearchRequest := *openapiclient.NewEmbeddingSimilaritySearchRequest("BucketName_example", "Query_example") // EmbeddingSimilaritySearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.PostEmbeddingSimilaritySearch(context.Background()).EmbeddingSimilaritySearchRequest(embeddingSimilaritySearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.PostEmbeddingSimilaritySearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEmbeddingSimilaritySearch`: EmbeddingSimilaritySearchResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.PostEmbeddingSimilaritySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEmbeddingSimilaritySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingSimilaritySearchRequest** | [**EmbeddingSimilaritySearchRequest**](EmbeddingSimilaritySearchRequest.md) |  | 

### Return type

[**EmbeddingSimilaritySearchResponse**](EmbeddingSimilaritySearchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEmbeddingUrl

> EmbeddingResponse PostEmbeddingUrl(ctx).EmbeddingUrlRequest(embeddingUrlRequest).Execute()

Embed URL content



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
	embeddingUrlRequest := *openapiclient.NewEmbeddingUrlRequest("Url_example", "BucketName_example") // EmbeddingUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.PostEmbeddingUrl(context.Background()).EmbeddingUrlRequest(embeddingUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.PostEmbeddingUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEmbeddingUrl`: EmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.PostEmbeddingUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEmbeddingUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingUrlRequest** | [**EmbeddingUrlRequest**](EmbeddingUrlRequest.md) |  | 

### Return type

[**EmbeddingResponse**](EmbeddingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

