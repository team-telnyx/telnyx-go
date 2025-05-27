# \CredentialsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTelephonyCredential**](CredentialsAPI.md#CreateTelephonyCredential) | **Post** /telephony_credentials | Create a credential
[**DeleteTelephonyCredential**](CredentialsAPI.md#DeleteTelephonyCredential) | **Delete** /telephony_credentials/{id} | Delete a credential
[**FindTelephonyCredentials**](CredentialsAPI.md#FindTelephonyCredentials) | **Get** /telephony_credentials | List all credentials
[**GetTelephonyCredential**](CredentialsAPI.md#GetTelephonyCredential) | **Get** /telephony_credentials/{id} | Get a credential
[**UpdateTelephonyCredential**](CredentialsAPI.md#UpdateTelephonyCredential) | **Patch** /telephony_credentials/{id} | Update a credential



## CreateTelephonyCredential

> TelephonyCredentialResponse CreateTelephonyCredential(ctx).TelephonyCredentialCreateRequest(telephonyCredentialCreateRequest).Execute()

Create a credential



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
	telephonyCredentialCreateRequest := *openapiclient.NewTelephonyCredentialCreateRequest("1234567890") // TelephonyCredentialCreateRequest | Parameters that can be defined during credential creation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateTelephonyCredential(context.Background()).TelephonyCredentialCreateRequest(telephonyCredentialCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTelephonyCredential`: TelephonyCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telephonyCredentialCreateRequest** | [**TelephonyCredentialCreateRequest**](TelephonyCredentialCreateRequest.md) | Parameters that can be defined during credential creation | 

### Return type

[**TelephonyCredentialResponse**](TelephonyCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTelephonyCredential

> TelephonyCredentialResponse DeleteTelephonyCredential(ctx, id).Execute()

Delete a credential



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
	id := "id_example" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.DeleteTelephonyCredential(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTelephonyCredential`: TelephonyCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.DeleteTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TelephonyCredentialResponse**](TelephonyCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTelephonyCredentials

> GetAllTelephonyCredentialResponse FindTelephonyCredentials(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterTag(filterTag).FilterName(filterName).FilterStatus(filterStatus).FilterResourceId(filterResourceId).FilterSipUsername(filterSipUsername).Execute()

List all credentials



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterTag := "filterTag_example" // string | Filter by tag (optional)
	filterName := "filterName_example" // string | Filter by name (optional)
	filterStatus := "filterStatus_example" // string | Filter by status (optional)
	filterResourceId := "filterResourceId_example" // string | Filter by resource_id (optional)
	filterSipUsername := "filterSipUsername_example" // string | Filter by sip_username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.FindTelephonyCredentials(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterTag(filterTag).FilterName(filterName).FilterStatus(filterStatus).FilterResourceId(filterResourceId).FilterSipUsername(filterSipUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.FindTelephonyCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindTelephonyCredentials`: GetAllTelephonyCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.FindTelephonyCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindTelephonyCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterTag** | **string** | Filter by tag | 
 **filterName** | **string** | Filter by name | 
 **filterStatus** | **string** | Filter by status | 
 **filterResourceId** | **string** | Filter by resource_id | 
 **filterSipUsername** | **string** | Filter by sip_username | 

### Return type

[**GetAllTelephonyCredentialResponse**](GetAllTelephonyCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelephonyCredential

> TelephonyCredentialResponse GetTelephonyCredential(ctx, id).Execute()

Get a credential



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
	id := "id_example" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetTelephonyCredential(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelephonyCredential`: TelephonyCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TelephonyCredentialResponse**](TelephonyCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTelephonyCredential

> TelephonyCredentialResponse UpdateTelephonyCredential(ctx, id).TelephonyCredentialUpdateRequest(telephonyCredentialUpdateRequest).Execute()

Update a credential



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
	id := "id_example" // string | Identifies the resource.
	telephonyCredentialUpdateRequest := *openapiclient.NewTelephonyCredentialUpdateRequest() // TelephonyCredentialUpdateRequest | Parameters that can be updated in a credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateTelephonyCredential(context.Background(), id).TelephonyCredentialUpdateRequest(telephonyCredentialUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTelephonyCredential`: TelephonyCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **telephonyCredentialUpdateRequest** | [**TelephonyCredentialUpdateRequest**](TelephonyCredentialUpdateRequest.md) | Parameters that can be updated in a credential | 

### Return type

[**TelephonyCredentialResponse**](TelephonyCredentialResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

