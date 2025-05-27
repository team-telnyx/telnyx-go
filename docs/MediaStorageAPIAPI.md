# \MediaStorageAPIAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMediaStorage**](MediaStorageAPIAPI.md#CreateMediaStorage) | **Post** /media | Upload media
[**DeleteMediaStorage**](MediaStorageAPIAPI.md#DeleteMediaStorage) | **Delete** /media/{media_name} | Deletes stored media
[**DownloadMedia**](MediaStorageAPIAPI.md#DownloadMedia) | **Get** /media/{media_name}/download | Download stored media
[**GetMediaStorage**](MediaStorageAPIAPI.md#GetMediaStorage) | **Get** /media/{media_name} | Retrieve stored media
[**ListMediaStorage**](MediaStorageAPIAPI.md#ListMediaStorage) | **Get** /media | List uploaded media
[**UpdateMediaStorage**](MediaStorageAPIAPI.md#UpdateMediaStorage) | **Put** /media/{media_name} | Update stored media



## CreateMediaStorage

> MediaResourceResponse CreateMediaStorage(ctx).UploadMediaRequest(uploadMediaRequest).Execute()

Upload media



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
	uploadMediaRequest := *openapiclient.NewUploadMediaRequest("http://www.example.com/audio.mp3") // UploadMediaRequest | Upload media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaStorageAPIAPI.CreateMediaStorage(context.Background()).UploadMediaRequest(uploadMediaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaStorageAPIAPI.CreateMediaStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMediaStorage`: MediaResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaStorageAPIAPI.CreateMediaStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMediaStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadMediaRequest** | [**UploadMediaRequest**](UploadMediaRequest.md) | Upload media request | 

### Return type

[**MediaResourceResponse**](MediaResourceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMediaStorage

> DeleteMediaStorage(ctx, mediaName).Execute()

Deletes stored media



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
	mediaName := "mediaName_example" // string | Uniquely identifies a media resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaStorageAPIAPI.DeleteMediaStorage(context.Background(), mediaName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaStorageAPIAPI.DeleteMediaStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaName** | **string** | Uniquely identifies a media resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaStorageRequest struct via the builder pattern


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


## DownloadMedia

> *os.File DownloadMedia(ctx, mediaName).Execute()

Download stored media



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
	mediaName := "mediaName_example" // string | Uniquely identifies a media resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaStorageAPIAPI.DownloadMedia(context.Background(), mediaName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaStorageAPIAPI.DownloadMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMedia`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MediaStorageAPIAPI.DownloadMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaName** | **string** | Uniquely identifies a media resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaStorage

> MediaResourceResponse GetMediaStorage(ctx, mediaName).Execute()

Retrieve stored media



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
	mediaName := "mediaName_example" // string | Uniquely identifies a media resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaStorageAPIAPI.GetMediaStorage(context.Background(), mediaName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaStorageAPIAPI.GetMediaStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaStorage`: MediaResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaStorageAPIAPI.GetMediaStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaName** | **string** | Uniquely identifies a media resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaResourceResponse**](MediaResourceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMediaStorage

> ListOfMediaResourcesResponse ListMediaStorage(ctx).FilterContentType(filterContentType).Execute()

List uploaded media



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
	filterContentType := "application_xml" // string | Filters files by given content types (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaStorageAPIAPI.ListMediaStorage(context.Background()).FilterContentType(filterContentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaStorageAPIAPI.ListMediaStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMediaStorage`: ListOfMediaResourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaStorageAPIAPI.ListMediaStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMediaStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterContentType** | **string** | Filters files by given content types | 

### Return type

[**ListOfMediaResourcesResponse**](ListOfMediaResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMediaStorage

> MediaResourceResponse UpdateMediaStorage(ctx, mediaName).UpdateMediaRequest(updateMediaRequest).Execute()

Update stored media



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
	mediaName := "mediaName_example" // string | Uniquely identifies a media resource.
	updateMediaRequest := *openapiclient.NewUpdateMediaRequest() // UpdateMediaRequest | Update media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaStorageAPIAPI.UpdateMediaStorage(context.Background(), mediaName).UpdateMediaRequest(updateMediaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaStorageAPIAPI.UpdateMediaStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMediaStorage`: MediaResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaStorageAPIAPI.UpdateMediaStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaName** | **string** | Uniquely identifies a media resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMediaStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMediaRequest** | [**UpdateMediaRequest**](UpdateMediaRequest.md) | Update media request | 

### Return type

[**MediaResourceResponse**](MediaResourceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

