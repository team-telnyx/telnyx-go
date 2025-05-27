# \ExternalConnectionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalConnection**](ExternalConnectionsAPI.md#CreateExternalConnection) | **Post** /external_connections | Creates an External Connection
[**CreateExternalConnectionUpload**](ExternalConnectionsAPI.md#CreateExternalConnectionUpload) | **Post** /external_connections/{id}/uploads | Creates an Upload request
[**DeleteExternalConnection**](ExternalConnectionsAPI.md#DeleteExternalConnection) | **Delete** /external_connections/{id} | Deletes an External Connection
[**DeleteExternalConnectionLogMessage**](ExternalConnectionsAPI.md#DeleteExternalConnectionLogMessage) | **Delete** /external_connections/log_messages/{id} | Dismiss a log message
[**GetExternalConnection**](ExternalConnectionsAPI.md#GetExternalConnection) | **Get** /external_connections/{id} | Retrieve an External Connection
[**GetExternalConnectionCivicAddress**](ExternalConnectionsAPI.md#GetExternalConnectionCivicAddress) | **Get** /external_connections/{id}/civic_addresses/{address_id} | Retrieve a Civic Address
[**GetExternalConnectionLogMessage**](ExternalConnectionsAPI.md#GetExternalConnectionLogMessage) | **Get** /external_connections/log_messages/{id} | Retrieve a log message
[**GetExternalConnectionPhoneNumber**](ExternalConnectionsAPI.md#GetExternalConnectionPhoneNumber) | **Get** /external_connections/{id}/phone_numbers/{phone_number_id} | Retrieve a phone number
[**GetExternalConnectionRelease**](ExternalConnectionsAPI.md#GetExternalConnectionRelease) | **Get** /external_connections/{id}/releases/{release_id} | Retrieve a Release request
[**GetExternalConnectionUpload**](ExternalConnectionsAPI.md#GetExternalConnectionUpload) | **Get** /external_connections/{id}/uploads/{ticket_id} | Retrieve an Upload request
[**GetExternalConnectionUploadsStatus**](ExternalConnectionsAPI.md#GetExternalConnectionUploadsStatus) | **Get** /external_connections/{id}/uploads/status | Get the count of pending upload requests
[**ListCivicAddresses**](ExternalConnectionsAPI.md#ListCivicAddresses) | **Get** /external_connections/{id}/civic_addresses | List all civic addresses and locations
[**ListExternalConnectionLogMessages**](ExternalConnectionsAPI.md#ListExternalConnectionLogMessages) | **Get** /external_connections/log_messages | List all log messages
[**ListExternalConnectionPhoneNumbers**](ExternalConnectionsAPI.md#ListExternalConnectionPhoneNumbers) | **Get** /external_connections/{id}/phone_numbers | List all phone numbers
[**ListExternalConnectionReleases**](ExternalConnectionsAPI.md#ListExternalConnectionReleases) | **Get** /external_connections/{id}/releases | List all Releases
[**ListExternalConnectionUploads**](ExternalConnectionsAPI.md#ListExternalConnectionUploads) | **Get** /external_connections/{id}/uploads | List all Upload requests
[**ListExternalConnections**](ExternalConnectionsAPI.md#ListExternalConnections) | **Get** /external_connections | List all External Connections
[**OperatorConnectRefresh**](ExternalConnectionsAPI.md#OperatorConnectRefresh) | **Post** /operator_connect/actions/refresh | Refresh Operator Connect integration
[**RefreshExternalConnectionUploads**](ExternalConnectionsAPI.md#RefreshExternalConnectionUploads) | **Post** /external_connections/{id}/uploads/refresh | Refresh the status of all Upload requests
[**RetryUpload**](ExternalConnectionsAPI.md#RetryUpload) | **Post** /external_connections/{id}/uploads/{ticket_id}/retry | Retry an Upload request
[**UpdateExternalConnection**](ExternalConnectionsAPI.md#UpdateExternalConnection) | **Patch** /external_connections/{id} | Update an External Connection
[**UpdateExternalConnectionPhoneNumber**](ExternalConnectionsAPI.md#UpdateExternalConnectionPhoneNumber) | **Patch** /external_connections/{id}/phone_numbers/{phone_number_id} | Update a phone number
[**UpdateLocation**](ExternalConnectionsAPI.md#UpdateLocation) | **Patch** /external_connections/{id}/locations/{location_id} | Update a location&#39;s static emergency address



## CreateExternalConnection

> ExternalConnectionResponse CreateExternalConnection(ctx).CreateExternalConnectionRequest(createExternalConnectionRequest).Execute()

Creates an External Connection



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
	createExternalConnectionRequest := *openapiclient.NewCreateExternalConnectionRequest(openapiclient.ExternalSipConnectionZoomOnly("zoom")) // CreateExternalConnectionRequest | Parameters that can be set when creating a External Connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.CreateExternalConnection(context.Background()).CreateExternalConnectionRequest(createExternalConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.CreateExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalConnection`: ExternalConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.CreateExternalConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExternalConnectionRequest** | [**CreateExternalConnectionRequest**](CreateExternalConnectionRequest.md) | Parameters that can be set when creating a External Connection | 

### Return type

[**ExternalConnectionResponse**](ExternalConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalConnectionUpload

> CreateUploadRequestResponse CreateExternalConnectionUpload(ctx, id).CreateExternalConnectionUploadRequest(createExternalConnectionUploadRequest).Execute()

Creates an Upload request



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
	id := "1293384261075731499" // string | Identifies the resource.
	createExternalConnectionUploadRequest := *openapiclient.NewCreateExternalConnectionUploadRequest() // CreateExternalConnectionUploadRequest | Parameters that can be set when creating an Upload request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.CreateExternalConnectionUpload(context.Background(), id).CreateExternalConnectionUploadRequest(createExternalConnectionUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.CreateExternalConnectionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalConnectionUpload`: CreateUploadRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.CreateExternalConnectionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalConnectionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createExternalConnectionUploadRequest** | [**CreateExternalConnectionUploadRequest**](CreateExternalConnectionUploadRequest.md) | Parameters that can be set when creating an Upload request. | 

### Return type

[**CreateUploadRequestResponse**](CreateUploadRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalConnection

> ExternalConnectionResponse DeleteExternalConnection(ctx, id).Execute()

Deletes an External Connection



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.DeleteExternalConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.DeleteExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalConnection`: ExternalConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.DeleteExternalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalConnectionResponse**](ExternalConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalConnectionLogMessage

> DismissRequestWasSuccessful DeleteExternalConnectionLogMessage(ctx, id).Execute()

Dismiss a log message



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.DeleteExternalConnectionLogMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.DeleteExternalConnectionLogMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalConnectionLogMessage`: DismissRequestWasSuccessful
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.DeleteExternalConnectionLogMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalConnectionLogMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DismissRequestWasSuccessful**](DismissRequestWasSuccessful.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnection

> ExternalConnectionResponse GetExternalConnection(ctx, id).Execute()

Retrieve an External Connection



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnection`: ExternalConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalConnectionResponse**](ExternalConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionCivicAddress

> GetCivicAddressResponse GetExternalConnectionCivicAddress(ctx, id, addressId).Execute()

Retrieve a Civic Address



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
	id := "1293384261075731499" // string | Identifies the resource.
	addressId := "318fb664-d341-44d2-8405-e6bfb9ced6d9" // string | Identifies a civic address or a location.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnectionCivicAddress(context.Background(), id, addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnectionCivicAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionCivicAddress`: GetCivicAddressResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnectionCivicAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 
**addressId** | **string** | Identifies a civic address or a location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionCivicAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCivicAddressResponse**](GetCivicAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionLogMessage

> GetLogMessageResponse GetExternalConnectionLogMessage(ctx, id).Execute()

Retrieve a log message



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnectionLogMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnectionLogMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionLogMessage`: GetLogMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnectionLogMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionLogMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLogMessageResponse**](GetLogMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionPhoneNumber

> GetExternalConnectionPhoneNumberResponse GetExternalConnectionPhoneNumber(ctx, id, phoneNumberId).Execute()

Retrieve a phone number



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
	id := "1293384261075731499" // string | Identifies the resource.
	phoneNumberId := "1234567889" // string | A phone number's ID via the Telnyx API

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnectionPhoneNumber(context.Background(), id, phoneNumberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnectionPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionPhoneNumber`: GetExternalConnectionPhoneNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnectionPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 
**phoneNumberId** | **string** | A phone number&#39;s ID via the Telnyx API | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetExternalConnectionPhoneNumberResponse**](GetExternalConnectionPhoneNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionRelease

> GetReleaseResponse GetExternalConnectionRelease(ctx, id, releaseId).Execute()

Retrieve a Release request



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
	id := "1293384261075731499" // string | Identifies the resource.
	releaseId := "7b6a6449-b055-45a6-81f6-f6f0dffa4cc6" // string | Identifies a Release request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnectionRelease(context.Background(), id, releaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnectionRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionRelease`: GetReleaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnectionRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 
**releaseId** | **string** | Identifies a Release request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetReleaseResponse**](GetReleaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionUpload

> GetUploadResponse GetExternalConnectionUpload(ctx, id, ticketId).Execute()

Retrieve an Upload request



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
	id := "1293384261075731499" // string | Identifies the resource.
	ticketId := "7b6a6449-b055-45a6-81f6-f6f0dffa4cc6" // string | Identifies an Upload request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnectionUpload(context.Background(), id, ticketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnectionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionUpload`: GetUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnectionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 
**ticketId** | **string** | Identifies an Upload request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUploadResponse**](GetUploadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionUploadsStatus

> GetUploadsStatusResponse GetExternalConnectionUploadsStatus(ctx, id).Execute()

Get the count of pending upload requests



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.GetExternalConnectionUploadsStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.GetExternalConnectionUploadsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionUploadsStatus`: GetUploadsStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.GetExternalConnectionUploadsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionUploadsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUploadsStatusResponse**](GetUploadsStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCivicAddresses

> GetAllCivicAddressesResponse ListCivicAddresses(ctx, id).FilterCountry(filterCountry).Execute()

List all civic addresses and locations



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
	id := "1293384261075731499" // string | Identifies the resource.
	filterCountry := []string{"Inner_example"} // []string | The country (or countries) to filter addresses by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.ListCivicAddresses(context.Background(), id).FilterCountry(filterCountry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.ListCivicAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCivicAddresses`: GetAllCivicAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.ListCivicAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCivicAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterCountry** | **[]string** | The country (or countries) to filter addresses by. | 

### Return type

[**GetAllCivicAddressesResponse**](GetAllCivicAddressesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalConnectionLogMessages

> ListLogMessagesResponse ListExternalConnectionLogMessages(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterExternalConnectionId(filterExternalConnectionId).FilterTelephoneNumberContains(filterTelephoneNumberContains).FilterTelephoneNumberEq(filterTelephoneNumberEq).Execute()

List all log messages



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
	filterExternalConnectionId := "67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The external connection ID to filter by or \"null\" to filter for logs without an external connection ID (optional)
	filterTelephoneNumberContains := "+123" // string | The partial phone number to filter log messages for. Requires 3-15 digits. (optional)
	filterTelephoneNumberEq := "+1234567890" // string | The phone number to filter log messages for or \"null\" to filter for logs without a phone number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.ListExternalConnectionLogMessages(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterExternalConnectionId(filterExternalConnectionId).FilterTelephoneNumberContains(filterTelephoneNumberContains).FilterTelephoneNumberEq(filterTelephoneNumberEq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.ListExternalConnectionLogMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalConnectionLogMessages`: ListLogMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.ListExternalConnectionLogMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalConnectionLogMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterExternalConnectionId** | **string** | The external connection ID to filter by or \&quot;null\&quot; to filter for logs without an external connection ID | 
 **filterTelephoneNumberContains** | **string** | The partial phone number to filter log messages for. Requires 3-15 digits. | 
 **filterTelephoneNumberEq** | **string** | The phone number to filter log messages for or \&quot;null\&quot; to filter for logs without a phone number | 

### Return type

[**ListLogMessagesResponse**](ListLogMessagesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalConnectionPhoneNumbers

> ListExternalConnectionPhoneNumbersResponse ListExternalConnectionPhoneNumbers(ctx, id).PageNumber(pageNumber).PageSize(pageSize).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberContains(filterPhoneNumberContains).FilterCivicAddressIdEq(filterCivicAddressIdEq).FilterLocationIdEq(filterLocationIdEq).Execute()

List all phone numbers



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
	id := "1293384261075731499" // string | Identifies the resource.
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterPhoneNumberEq := "+1234567890" // string | The phone number to filter by (optional)
	filterPhoneNumberContains := "+123" // string | The partial phone number to filter by. Requires 3-15 digits. (optional)
	filterCivicAddressIdEq := "67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The civic address ID to filter by (optional)
	filterLocationIdEq := "52545f6f-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The location ID to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.ListExternalConnectionPhoneNumbers(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberContains(filterPhoneNumberContains).FilterCivicAddressIdEq(filterCivicAddressIdEq).FilterLocationIdEq(filterLocationIdEq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.ListExternalConnectionPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalConnectionPhoneNumbers`: ListExternalConnectionPhoneNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.ListExternalConnectionPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalConnectionPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterPhoneNumberEq** | **string** | The phone number to filter by | 
 **filterPhoneNumberContains** | **string** | The partial phone number to filter by. Requires 3-15 digits. | 
 **filterCivicAddressIdEq** | **string** | The civic address ID to filter by | 
 **filterLocationIdEq** | **string** | The location ID to filter by | 

### Return type

[**ListExternalConnectionPhoneNumbersResponse**](ListExternalConnectionPhoneNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalConnectionReleases

> ListReleasesResponse ListExternalConnectionReleases(ctx, id).PageNumber(pageNumber).PageSize(pageSize).FilterStatusEq(filterStatusEq).FilterCivicAddressIdEq(filterCivicAddressIdEq).FilterLocationIdEq(filterLocationIdEq).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberContains(filterPhoneNumberContains).Execute()

List all Releases



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
	id := "1293384261075731499" // string | Identifies the resource.
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterStatusEq := []string{"FilterStatusEq_example"} // []string | The status of the release to filter by (optional)
	filterCivicAddressIdEq := "67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The civic address ID to filter by (optional)
	filterLocationIdEq := "52545f6f-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The location ID to filter by (optional)
	filterPhoneNumberEq := "+1234567890" // string | The phone number to filter by (optional)
	filterPhoneNumberContains := "+123" // string | The partial phone number to filter by. Requires 3-15 digits. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.ListExternalConnectionReleases(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).FilterStatusEq(filterStatusEq).FilterCivicAddressIdEq(filterCivicAddressIdEq).FilterLocationIdEq(filterLocationIdEq).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberContains(filterPhoneNumberContains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.ListExternalConnectionReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalConnectionReleases`: ListReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.ListExternalConnectionReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalConnectionReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterStatusEq** | **[]string** | The status of the release to filter by | 
 **filterCivicAddressIdEq** | **string** | The civic address ID to filter by | 
 **filterLocationIdEq** | **string** | The location ID to filter by | 
 **filterPhoneNumberEq** | **string** | The phone number to filter by | 
 **filterPhoneNumberContains** | **string** | The partial phone number to filter by. Requires 3-15 digits. | 

### Return type

[**ListReleasesResponse**](ListReleasesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalConnectionUploads

> ListUploadsResponse ListExternalConnectionUploads(ctx, id).PageNumber(pageNumber).PageSize(pageSize).FilterStatusEq(filterStatusEq).FilterCivicAddressIdEq(filterCivicAddressIdEq).FilterLocationIdEq(filterLocationIdEq).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberContains(filterPhoneNumberContains).Execute()

List all Upload requests



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
	id := "1293384261075731499" // string | Identifies the resource.
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterStatusEq := []string{"FilterStatusEq_example"} // []string | The status of the upload to filter by (optional)
	filterCivicAddressIdEq := "67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The civic address ID to filter by (optional)
	filterLocationIdEq := "52545f6f-9cd5-4a68-8c76-abb3aa5bf5d2" // string | The location ID to filter by (optional)
	filterPhoneNumberEq := "+1234567890" // string | The phone number to filter by (optional)
	filterPhoneNumberContains := "+123" // string | The partial phone number to filter by. Requires 3-15 digits. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.ListExternalConnectionUploads(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).FilterStatusEq(filterStatusEq).FilterCivicAddressIdEq(filterCivicAddressIdEq).FilterLocationIdEq(filterLocationIdEq).FilterPhoneNumberEq(filterPhoneNumberEq).FilterPhoneNumberContains(filterPhoneNumberContains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.ListExternalConnectionUploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalConnectionUploads`: ListUploadsResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.ListExternalConnectionUploads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalConnectionUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterStatusEq** | **[]string** | The status of the upload to filter by | 
 **filterCivicAddressIdEq** | **string** | The civic address ID to filter by | 
 **filterLocationIdEq** | **string** | The location ID to filter by | 
 **filterPhoneNumberEq** | **string** | The phone number to filter by | 
 **filterPhoneNumberContains** | **string** | The partial phone number to filter by. Requires 3-15 digits. | 

### Return type

[**ListUploadsResponse**](ListUploadsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalConnections

> GetAllExternalConnectionsResponse ListExternalConnections(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterExternalSipConnection(filterExternalSipConnection).FilterId(filterId).FilterCreatedAt(filterCreatedAt).FilterPhoneNumberEq(filterPhoneNumberEq).Execute()

List all External Connections



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
	filterConnectionNameContains := "filterConnectionNameContains_example" // string | If present, connections with <code>connection_name</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional)
	filterExternalSipConnection := "zoom" // string | If present, connections with <code>external_sip_connection</code> matching the given value will be returned. (optional)
	filterId := "1930241863466354012" // string | If present, connections with <code>id</code> matching the given value will be returned. (optional)
	filterCreatedAt := "2018-02-02T22:25:27.521Z" // string | Filter by ISO 8601 formatted date-time string matching resource creation date-time. (optional)
	filterPhoneNumberEq := "+15555555555" // string | If present, connections associated with the given phone_number will be returned. A full match is necessary with a e164 format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.ListExternalConnections(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterConnectionNameContains(filterConnectionNameContains).FilterExternalSipConnection(filterExternalSipConnection).FilterId(filterId).FilterCreatedAt(filterCreatedAt).FilterPhoneNumberEq(filterPhoneNumberEq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.ListExternalConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalConnections`: GetAllExternalConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.ListExternalConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterConnectionNameContains** | **string** | If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | 
 **filterExternalSipConnection** | **string** | If present, connections with &lt;code&gt;external_sip_connection&lt;/code&gt; matching the given value will be returned. | 
 **filterId** | **string** | If present, connections with &lt;code&gt;id&lt;/code&gt; matching the given value will be returned. | 
 **filterCreatedAt** | **string** | Filter by ISO 8601 formatted date-time string matching resource creation date-time. | 
 **filterPhoneNumberEq** | **string** | If present, connections associated with the given phone_number will be returned. A full match is necessary with a e164 format. | 

### Return type

[**GetAllExternalConnectionsResponse**](GetAllExternalConnectionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorConnectRefresh

> OperatorConnectRefreshResponse OperatorConnectRefresh(ctx).Execute()

Refresh Operator Connect integration



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
	resp, r, err := apiClient.ExternalConnectionsAPI.OperatorConnectRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.OperatorConnectRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperatorConnectRefresh`: OperatorConnectRefreshResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.OperatorConnectRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorConnectRefreshRequest struct via the builder pattern


### Return type

[**OperatorConnectRefreshResponse**](OperatorConnectRefreshResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshExternalConnectionUploads

> CreateUploadRequestResponse1 RefreshExternalConnectionUploads(ctx, id).Execute()

Refresh the status of all Upload requests



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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.RefreshExternalConnectionUploads(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.RefreshExternalConnectionUploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshExternalConnectionUploads`: CreateUploadRequestResponse1
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.RefreshExternalConnectionUploads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshExternalConnectionUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateUploadRequestResponse1**](CreateUploadRequestResponse1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryUpload

> GetUploadResponse RetryUpload(ctx, id, ticketId).Execute()

Retry an Upload request



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
	id := "1293384261075731499" // string | Identifies the resource.
	ticketId := "7b6a6449-b055-45a6-81f6-f6f0dffa4cc6" // string | Identifies an Upload request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.RetryUpload(context.Background(), id, ticketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.RetryUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryUpload`: GetUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.RetryUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 
**ticketId** | **string** | Identifies an Upload request | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUploadResponse**](GetUploadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalConnection

> ExternalConnectionResponse UpdateExternalConnection(ctx, id).UpdateExternalConnectionRequest(updateExternalConnectionRequest).Execute()

Update an External Connection



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
	id := "1293384261075731499" // string | Identifies the resource.
	updateExternalConnectionRequest := *openapiclient.NewUpdateExternalConnectionRequest() // UpdateExternalConnectionRequest | Parameters to be updated for the External Connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.UpdateExternalConnection(context.Background(), id).UpdateExternalConnectionRequest(updateExternalConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.UpdateExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalConnection`: ExternalConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.UpdateExternalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalConnectionRequest** | [**UpdateExternalConnectionRequest**](UpdateExternalConnectionRequest.md) | Parameters to be updated for the External Connection | 

### Return type

[**ExternalConnectionResponse**](ExternalConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalConnectionPhoneNumber

> GetExternalConnectionPhoneNumberResponse UpdateExternalConnectionPhoneNumber(ctx, id, phoneNumberId).UpdateExternalConnectionPhoneNumberRequest(updateExternalConnectionPhoneNumberRequest).Execute()

Update a phone number



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
	id := "1293384261075731499" // string | Identifies the resource.
	phoneNumberId := "1234567889" // string | A phone number's ID via the Telnyx API
	updateExternalConnectionPhoneNumberRequest := *openapiclient.NewUpdateExternalConnectionPhoneNumberRequest() // UpdateExternalConnectionPhoneNumberRequest | Values that can be set when updating a phone number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.UpdateExternalConnectionPhoneNumber(context.Background(), id, phoneNumberId).UpdateExternalConnectionPhoneNumberRequest(updateExternalConnectionPhoneNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.UpdateExternalConnectionPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalConnectionPhoneNumber`: GetExternalConnectionPhoneNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.UpdateExternalConnectionPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 
**phoneNumberId** | **string** | A phone number&#39;s ID via the Telnyx API | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalConnectionPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateExternalConnectionPhoneNumberRequest** | [**UpdateExternalConnectionPhoneNumberRequest**](UpdateExternalConnectionPhoneNumberRequest.md) | Values that can be set when updating a phone number | 

### Return type

[**GetExternalConnectionPhoneNumberResponse**](GetExternalConnectionPhoneNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> LocationResponse UpdateLocation(ctx, id, locationId).UpdateLocationRequest(updateLocationRequest).Execute()

Update a location's static emergency address

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the external connection
	locationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the location to update
	updateLocationRequest := *openapiclient.NewUpdateLocationRequest("StaticEmergencyAddressId_example") // UpdateLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionsAPI.UpdateLocation(context.Background(), id, locationId).UpdateLocationRequest(updateLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionsAPI.UpdateLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocation`: LocationResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionsAPI.UpdateLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the external connection | 
**locationId** | **string** | The ID of the location to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLocationRequest** | [**UpdateLocationRequest**](UpdateLocationRequest.md) |  | 

### Return type

[**LocationResponse**](LocationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

