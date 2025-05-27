# \ClustersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeNewClusterPublicTextClustersPost**](ClustersAPI.md#ComputeNewClusterPublicTextClustersPost) | **Post** /ai/clusters | Compute new clusters
[**DeleteClusterByTaskIdPublicTextClustersTaskIdDelete**](ClustersAPI.md#DeleteClusterByTaskIdPublicTextClustersTaskIdDelete) | **Delete** /ai/clusters/{task_id} | Delete a cluster
[**FetchClusterByTaskIdPublicTextClustersTaskIdGet**](ClustersAPI.md#FetchClusterByTaskIdPublicTextClustersTaskIdGet) | **Get** /ai/clusters/{task_id} | Fetch a cluster
[**FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet**](ClustersAPI.md#FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet) | **Get** /ai/clusters/{task_id}/graph | Fetch a cluster visualization
[**ListAllRequestedClustersPublicTextClustersGet**](ClustersAPI.md#ListAllRequestedClustersPublicTextClustersGet) | **Get** /ai/clusters | List all clusters



## ComputeNewClusterPublicTextClustersPost

> TextClusteringResponseData ComputeNewClusterPublicTextClustersPost(ctx).PublicTextClusteringRequest(publicTextClusteringRequest).Execute()

Compute new clusters



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
	publicTextClusteringRequest := *openapiclient.NewPublicTextClusteringRequest("Bucket_example") // PublicTextClusteringRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ComputeNewClusterPublicTextClustersPost(context.Background()).PublicTextClusteringRequest(publicTextClusteringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ComputeNewClusterPublicTextClustersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeNewClusterPublicTextClustersPost`: TextClusteringResponseData
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ComputeNewClusterPublicTextClustersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeNewClusterPublicTextClustersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicTextClusteringRequest** | [**PublicTextClusteringRequest**](PublicTextClusteringRequest.md) |  | 

### Return type

[**TextClusteringResponseData**](TextClusteringResponseData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterByTaskIdPublicTextClustersTaskIdDelete

> DeleteClusterByTaskIdPublicTextClustersTaskIdDelete(ctx, taskId).Execute()

Delete a cluster

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
	r, err := apiClient.ClustersAPI.DeleteClusterByTaskIdPublicTextClustersTaskIdDelete(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeleteClusterByTaskIdPublicTextClustersTaskIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterByTaskIdPublicTextClustersTaskIdDeleteRequest struct via the builder pattern


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


## FetchClusterByTaskIdPublicTextClustersTaskIdGet

> ClusteringStatusResponseData FetchClusterByTaskIdPublicTextClustersTaskIdGet(ctx, taskId).TopNNodes(topNNodes).ShowSubclusters(showSubclusters).Execute()

Fetch a cluster

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
	topNNodes := int32(56) // int32 | The number of nodes in the cluster to return in the response. Nodes will be sorted by their centrality within the cluster. (optional) (default to 0)
	showSubclusters := true // bool | Whether or not to include subclusters and their nodes in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.FetchClusterByTaskIdPublicTextClustersTaskIdGet(context.Background(), taskId).TopNNodes(topNNodes).ShowSubclusters(showSubclusters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.FetchClusterByTaskIdPublicTextClustersTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchClusterByTaskIdPublicTextClustersTaskIdGet`: ClusteringStatusResponseData
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.FetchClusterByTaskIdPublicTextClustersTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchClusterByTaskIdPublicTextClustersTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topNNodes** | **int32** | The number of nodes in the cluster to return in the response. Nodes will be sorted by their centrality within the cluster. | [default to 0]
 **showSubclusters** | **bool** | Whether or not to include subclusters and their nodes in the response. | [default to false]

### Return type

[**ClusteringStatusResponseData**](ClusteringStatusResponseData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet

> interface{} FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet(ctx, taskId).ClusterId(clusterId).Execute()

Fetch a cluster visualization

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
	clusterId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet(context.Background(), taskId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchClusterImageByTaskIdPublicTextClustersTaskIdImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterId** | **int32** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRequestedClustersPublicTextClustersGet

> ClusteringRequestInfoData ListAllRequestedClustersPublicTextClustersGet(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all clusters

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
	pageNumber := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListAllRequestedClustersPublicTextClustersGet(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListAllRequestedClustersPublicTextClustersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRequestedClustersPublicTextClustersGet`: ClusteringRequestInfoData
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListAllRequestedClustersPublicTextClustersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllRequestedClustersPublicTextClustersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**ClusteringRequestInfoData**](ClusteringRequestInfoData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

