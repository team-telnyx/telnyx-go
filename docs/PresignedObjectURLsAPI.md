# \PresignedObjectURLsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePresignedObjectUrl**](PresignedObjectURLsAPI.md#CreatePresignedObjectUrl) | **Post** /storage/buckets/{bucketName}/{objectName}/presigned_url | Create Presigned Object URL



## CreatePresignedObjectUrl

> PresignedObjectUrl CreatePresignedObjectUrl(ctx, bucketName, objectName).PresignedObjectUrlParams(presignedObjectUrlParams).Execute()

Create Presigned Object URL



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
	objectName := "objectName_example" // string | The name of the object
	presignedObjectUrlParams := *openapiclient.NewPresignedObjectUrlParams() // PresignedObjectUrlParams |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PresignedObjectURLsAPI.CreatePresignedObjectUrl(context.Background(), bucketName, objectName).PresignedObjectUrlParams(presignedObjectUrlParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PresignedObjectURLsAPI.CreatePresignedObjectUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePresignedObjectUrl`: PresignedObjectUrl
	fmt.Fprintf(os.Stdout, "Response from `PresignedObjectURLsAPI.CreatePresignedObjectUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 
**objectName** | **string** | The name of the object | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePresignedObjectUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **presignedObjectUrlParams** | [**PresignedObjectUrlParams**](PresignedObjectUrlParams.md) |  | 

### Return type

[**PresignedObjectUrl**](PresignedObjectUrl.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

