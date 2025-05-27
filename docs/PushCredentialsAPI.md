# \PushCredentialsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePushCredential**](PushCredentialsAPI.md#CreatePushCredential) | **Post** /mobile_push_credentials | Creates a new mobile push credential
[**DeletePushCredentialById**](PushCredentialsAPI.md#DeletePushCredentialById) | **Delete** /mobile_push_credentials/{push_credential_id} | Deletes a mobile push credential
[**GetPushCredentialById**](PushCredentialsAPI.md#GetPushCredentialById) | **Get** /mobile_push_credentials/{push_credential_id} | Retrieves a mobile push credential
[**ListPushCredentials**](PushCredentialsAPI.md#ListPushCredentials) | **Get** /mobile_push_credentials | List mobile push credentials



## CreatePushCredential

> PushCredentialResponse CreatePushCredential(ctx).CreatePushCredentialRequest(createPushCredentialRequest).Execute()

Creates a new mobile push credential



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
	createPushCredentialRequest := openapiclient.CreatePushCredential_request{CreateAndroidPushCredentialRequest: openapiclient.NewCreateAndroidPushCredentialRequest("Type_example", map[string]interface{}({"private_key":"BBBB0J56jd8kda:APA91vjb11BCjvxx3Jxja...","client_email":"account@customer.org"}), "LucyAndroidCredential")} // CreatePushCredentialRequest | Mobile push credential parameters that need to be sent in the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PushCredentialsAPI.CreatePushCredential(context.Background()).CreatePushCredentialRequest(createPushCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushCredentialsAPI.CreatePushCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePushCredential`: PushCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `PushCredentialsAPI.CreatePushCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPushCredentialRequest** | [**CreatePushCredentialRequest**](CreatePushCredentialRequest.md) | Mobile push credential parameters that need to be sent in the request | 

### Return type

[**PushCredentialResponse**](PushCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePushCredentialById

> DeletePushCredentialById(ctx, pushCredentialId).Execute()

Deletes a mobile push credential



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
	pushCredentialId := "0ccc7b76-4df3-4bca-a05a-3da1ecc389f0" // string | The unique identifier of a mobile push credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PushCredentialsAPI.DeletePushCredentialById(context.Background(), pushCredentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushCredentialsAPI.DeletePushCredentialById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pushCredentialId** | **string** | The unique identifier of a mobile push credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePushCredentialByIdRequest struct via the builder pattern


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


## GetPushCredentialById

> PushCredentialResponse GetPushCredentialById(ctx, pushCredentialId).Execute()

Retrieves a mobile push credential



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
	pushCredentialId := "0ccc7b76-4df3-4bca-a05a-3da1ecc389f0" // string | The unique identifier of a mobile push credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PushCredentialsAPI.GetPushCredentialById(context.Background(), pushCredentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushCredentialsAPI.GetPushCredentialById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPushCredentialById`: PushCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `PushCredentialsAPI.GetPushCredentialById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pushCredentialId** | **string** | The unique identifier of a mobile push credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushCredentialByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PushCredentialResponse**](PushCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPushCredentials

> ListPushCredentialsResponse ListPushCredentials(ctx).FilterType(filterType).FilterAlias(filterAlias).PageSize(pageSize).PageNumber(pageNumber).Execute()

List mobile push credentials



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
	filterType := "ios" // string | type of mobile push credentials (optional)
	filterAlias := "LucyCredential" // string | Unique mobile push credential alias (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PushCredentialsAPI.ListPushCredentials(context.Background()).FilterType(filterType).FilterAlias(filterAlias).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushCredentialsAPI.ListPushCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPushCredentials`: ListPushCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `PushCredentialsAPI.ListPushCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPushCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterType** | **string** | type of mobile push credentials | 
 **filterAlias** | **string** | Unique mobile push credential alias | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**ListPushCredentialsResponse**](ListPushCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

