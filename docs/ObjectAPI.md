# \ObjectAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObject**](ObjectAPI.md#DeleteObject) | **Delete** /{bucketName}/{objectName} | DeleteObject
[**DeleteObjects**](ObjectAPI.md#DeleteObjects) | **Post** /{bucketName} | DeleteObjects
[**GetObject**](ObjectAPI.md#GetObject) | **Get** /{bucketName}/{objectName} | GetObject
[**HeadObject**](ObjectAPI.md#HeadObject) | **Head** /{bucketName}/{objectName} | HeadObject
[**ListObjects**](ObjectAPI.md#ListObjects) | **Get** /{bucketName} | ListObjectsV2
[**PutObject**](ObjectAPI.md#PutObject) | **Put** /{bucketName}/{objectName} | PutObject



## DeleteObject

> DeleteObject(ctx, bucketName, objectName).Execute()

DeleteObject



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
	bucketName := "bucketName_example" // string | The bucket name.
	objectName := "objectName_example" // string | The object name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAPI.DeleteObject(context.Background(), bucketName, objectName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.DeleteObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The bucket name. | 
**objectName** | **string** | The object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjects

> map[string]interface{} DeleteObjects(ctx, bucketName).Delete(delete).DeleteObjectsRequestInner(deleteObjectsRequestInner).Execute()

DeleteObjects



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
	bucketName := "bucketName_example" // string | The bucket name.
	delete := true // bool | 
	deleteObjectsRequestInner := []openapiclient.DeleteObjectsRequestInner{*openapiclient.NewDeleteObjectsRequestInner()} // []DeleteObjectsRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.DeleteObjects(context.Background(), bucketName).Delete(delete).DeleteObjectsRequestInner(deleteObjectsRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.DeleteObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjects`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.DeleteObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | **bool** |  | 
 **deleteObjectsRequestInner** | [**[]DeleteObjectsRequestInner**](DeleteObjectsRequestInner.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObject

> *os.File GetObject(ctx, bucketName, objectName).UploadId(uploadId).Execute()

GetObject



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
	bucketName := "bucketName_example" // string | The bucket name.
	objectName := "objectName_example" // string | The object name.
	uploadId := "uploadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.GetObject(context.Background(), bucketName, objectName).UploadId(uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObject`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The bucket name. | 
**objectName** | **string** | The object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadId** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadObject

> HeadObject(ctx, bucketName, objectName).Execute()

HeadObject



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
	bucketName := "bucketName_example" // string | The bucket name.
	objectName := "objectName_example" // string | The object name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAPI.HeadObject(context.Background(), bucketName, objectName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.HeadObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The bucket name. | 
**objectName** | **string** | The object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjects

> ListObjectsResponse ListObjects(ctx, bucketName).ListType(listType).Execute()

ListObjectsV2



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
	bucketName := "bucketName_example" // string | The name of the bucket.
	listType := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAPI.ListObjects(context.Background(), bucketName).ListType(listType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.ListObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListObjects`: ListObjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.ListObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listType** | **int32** |  | 

### Return type

[**ListObjectsResponse**](ListObjectsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObject

> PutObject(ctx, bucketName, objectName).Body(body).ContentType(contentType).PartNumber(partNumber).UploadId(uploadId).Execute()

PutObject



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
	bucketName := "bucketName_example" // string | The bucket name.
	objectName := "objectName_example" // string | The object name.
	body := os.NewFile(1234, "some_file") // *os.File | 
	contentType := "contentType_example" // string |  (optional)
	partNumber := "partNumber_example" // string |  (optional)
	uploadId := "uploadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAPI.PutObject(context.Background(), bucketName, objectName).Body(body).ContentType(contentType).PartNumber(partNumber).UploadId(uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.PutObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The bucket name. | 
**objectName** | **string** | The object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 
 **contentType** | **string** |  | 
 **partNumber** | **string** |  | 
 **uploadId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

