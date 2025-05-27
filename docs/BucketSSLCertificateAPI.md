# \BucketSSLCertificateAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStorageSSLCertificate**](BucketSSLCertificateAPI.md#AddStorageSSLCertificate) | **Put** /storage/buckets/{bucketName}/ssl_certificate | Add SSL Certificate
[**GetStorageSSLCertificates**](BucketSSLCertificateAPI.md#GetStorageSSLCertificates) | **Get** /storage/buckets/{bucketName}/ssl_certificate | Get Bucket SSL Certificate
[**RemoveStorageSSLCertificate**](BucketSSLCertificateAPI.md#RemoveStorageSSLCertificate) | **Delete** /storage/buckets/{bucketName}/ssl_certificate | Remove SSL Certificate



## AddStorageSSLCertificate

> GetStorageSSLCertificates200Response AddStorageSSLCertificate(ctx, bucketName).Certificate(certificate).PrivateKey(privateKey).Execute()

Add SSL Certificate



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
	certificate := os.NewFile(1234, "some_file") // *os.File | The SSL certificate file (optional)
	privateKey := os.NewFile(1234, "some_file") // *os.File | The private key file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketSSLCertificateAPI.AddStorageSSLCertificate(context.Background(), bucketName).Certificate(certificate).PrivateKey(privateKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketSSLCertificateAPI.AddStorageSSLCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddStorageSSLCertificate`: GetStorageSSLCertificates200Response
	fmt.Fprintf(os.Stdout, "Response from `BucketSSLCertificateAPI.AddStorageSSLCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddStorageSSLCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificate** | ***os.File** | The SSL certificate file | 
 **privateKey** | ***os.File** | The private key file | 

### Return type

[**GetStorageSSLCertificates200Response**](GetStorageSSLCertificates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageSSLCertificates

> GetStorageSSLCertificates200Response GetStorageSSLCertificates(ctx, bucketName).Execute()

Get Bucket SSL Certificate



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
	resp, r, err := apiClient.BucketSSLCertificateAPI.GetStorageSSLCertificates(context.Background(), bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketSSLCertificateAPI.GetStorageSSLCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageSSLCertificates`: GetStorageSSLCertificates200Response
	fmt.Fprintf(os.Stdout, "Response from `BucketSSLCertificateAPI.GetStorageSSLCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageSSLCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStorageSSLCertificates200Response**](GetStorageSSLCertificates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveStorageSSLCertificate

> GetStorageSSLCertificates200Response RemoveStorageSSLCertificate(ctx, bucketName).Execute()

Remove SSL Certificate



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
	bucketName := "bucketName_example" // string | Bucket Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketSSLCertificateAPI.RemoveStorageSSLCertificate(context.Background(), bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketSSLCertificateAPI.RemoveStorageSSLCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveStorageSSLCertificate`: GetStorageSSLCertificates200Response
	fmt.Fprintf(os.Stdout, "Response from `BucketSSLCertificateAPI.RemoveStorageSSLCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveStorageSSLCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStorageSSLCertificates200Response**](GetStorageSSLCertificates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

