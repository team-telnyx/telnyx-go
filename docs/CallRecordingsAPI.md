# \CallRecordingsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomStorageCredentials**](CallRecordingsAPI.md#CreateCustomStorageCredentials) | **Post** /custom_storage_credentials/{connection_id} | Create a custom storage credential
[**DeleteCustomStorageCredentials**](CallRecordingsAPI.md#DeleteCustomStorageCredentials) | **Delete** /custom_storage_credentials/{connection_id} | Delete a stored credential
[**DeleteRecording**](CallRecordingsAPI.md#DeleteRecording) | **Delete** /recordings/{recording_id} | Delete a call recording
[**DeleteRecordingTranscription**](CallRecordingsAPI.md#DeleteRecordingTranscription) | **Delete** /recording_transcriptions/{recording_transcription_id} | Delete a recording transcription
[**DeleteRecordings**](CallRecordingsAPI.md#DeleteRecordings) | **Delete** /recordings/actions/delete | Delete a list of call recordings
[**GetCustomStorageCredentials**](CallRecordingsAPI.md#GetCustomStorageCredentials) | **Get** /custom_storage_credentials/{connection_id} | Retrieve a stored credential
[**GetRecording**](CallRecordingsAPI.md#GetRecording) | **Get** /recordings/{recording_id} | Retrieve a call recording
[**GetRecordingTranscription**](CallRecordingsAPI.md#GetRecordingTranscription) | **Get** /recording_transcriptions/{recording_transcription_id} | Retrieve a recording transcription
[**GetRecordingTranscriptions**](CallRecordingsAPI.md#GetRecordingTranscriptions) | **Get** /recording_transcriptions | List all recording transcriptions
[**GetRecordings**](CallRecordingsAPI.md#GetRecordings) | **Get** /recordings | List all call recordings
[**UpdateCustomStorageCredentials**](CallRecordingsAPI.md#UpdateCustomStorageCredentials) | **Put** /custom_storage_credentials/{connection_id} | Update a stored credential



## CreateCustomStorageCredentials

> CredentialsResponse CreateCustomStorageCredentials(ctx, connectionId).CustomStorageConfiguration(customStorageConfiguration).Execute()

Create a custom storage credential



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
	connectionId := "connectionId_example" // string | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource.
	customStorageConfiguration := *openapiclient.NewCustomStorageConfiguration("gcs", openapiclient.CustomStorageConfiguration_configuration{AzureConfigurationData: openapiclient.NewAzureConfigurationData()}) // CustomStorageConfiguration | Creates new credentials resource for the specified connection_id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.CreateCustomStorageCredentials(context.Background(), connectionId).CustomStorageConfiguration(customStorageConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.CreateCustomStorageCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomStorageCredentials`: CredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.CreateCustomStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customStorageConfiguration** | [**CustomStorageConfiguration**](CustomStorageConfiguration.md) | Creates new credentials resource for the specified connection_id. | 

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomStorageCredentials

> DeleteCustomStorageCredentials(ctx, connectionId).Execute()

Delete a stored credential



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
	connectionId := "connectionId_example" // string | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CallRecordingsAPI.DeleteCustomStorageCredentials(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.DeleteCustomStorageCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomStorageCredentialsRequest struct via the builder pattern


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


## DeleteRecording

> RecordingResponse DeleteRecording(ctx, recordingId).Execute()

Delete a call recording



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
	recordingId := "recordingId_example" // string | Uniquely identifies the recording by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.DeleteRecording(context.Background(), recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.DeleteRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecording`: RecordingResponse
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.DeleteRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | Uniquely identifies the recording by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecordingResponse**](RecordingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingTranscription

> GetRecordingTranscription200Response DeleteRecordingTranscription(ctx, recordingTranscriptionId).Execute()

Delete a recording transcription



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
	recordingTranscriptionId := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording transcription by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.DeleteRecordingTranscription(context.Background(), recordingTranscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.DeleteRecordingTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecordingTranscription`: GetRecordingTranscription200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.DeleteRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingTranscriptionId** | **string** | Uniquely identifies the recording transcription by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRecordingTranscription200Response**](GetRecordingTranscription200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordings

> DeleteRecordings(ctx).RequestBody(requestBody).Execute()

Delete a list of call recordings



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
	requestBody := []string{"Property_example"} // []string | Deletes recordings for the given list of IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CallRecordingsAPI.DeleteRecordings(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.DeleteRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Deletes recordings for the given list of IDs. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomStorageCredentials

> CredentialsResponse GetCustomStorageCredentials(ctx, connectionId).Execute()

Retrieve a stored credential



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
	connectionId := "connectionId_example" // string | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.GetCustomStorageCredentials(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.GetCustomStorageCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomStorageCredentials`: CredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.GetCustomStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecording

> RecordingResponse GetRecording(ctx, recordingId).Execute()

Retrieve a call recording



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
	recordingId := "recordingId_example" // string | Uniquely identifies the recording by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.GetRecording(context.Background(), recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.GetRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecording`: RecordingResponse
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.GetRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | Uniquely identifies the recording by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecordingResponse**](RecordingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingTranscription

> GetRecordingTranscription200Response GetRecordingTranscription(ctx, recordingTranscriptionId).Execute()

Retrieve a recording transcription



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
	recordingTranscriptionId := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording transcription by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.GetRecordingTranscription(context.Background(), recordingTranscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.GetRecordingTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordingTranscription`: GetRecordingTranscription200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.GetRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingTranscriptionId** | **string** | Uniquely identifies the recording transcription by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRecordingTranscription200Response**](GetRecordingTranscription200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingTranscriptions

> GetRecordingTranscriptions200Response GetRecordingTranscriptions(ctx).Execute()

List all recording transcriptions



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
	resp, r, err := apiClient.CallRecordingsAPI.GetRecordingTranscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.GetRecordingTranscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordingTranscriptions`: GetRecordingTranscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.GetRecordingTranscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingTranscriptionsRequest struct via the builder pattern


### Return type

[**GetRecordingTranscriptions200Response**](GetRecordingTranscriptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordings

> GetRecordings200Response GetRecordings(ctx).FilterConferenceId(filterConferenceId).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtLte(filterCreatedAtLte).FilterCallLegId(filterCallLegId).FilterCallSessionId(filterCallSessionId).FilterFrom(filterFrom).FilterTo(filterTo).FilterConnectionId(filterConnectionId).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all call recordings



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
	filterConferenceId := "428c31b6-7af4-4bcb-b7f5-5013ef9657c1" // string | Returns only recordings associated with a given conference. (optional)
	filterCreatedAtGte := "2019-03-29T11:10:00Z" // string | Returns only recordings created later than or at given ISO 8601 datetime. (optional)
	filterCreatedAtLte := "2019-03-29T11:10:00Z" // string | Returns only recordings created earlier than or at given ISO 8601 datetime. (optional)
	filterCallLegId := "428c31b6-7af4-4bcb-b7f5-5013ef9657c1" // string | If present, recordings will be filtered to those with a matching call_leg_id. (optional)
	filterCallSessionId := "428c31b6-7af4-4bcb-b7f5-5013ef9657c1" // string | If present, recordings will be filtered to those with a matching call_session_id. (optional)
	filterFrom := "1234567890" // string | If present, recordings will be filtered to those with a matching `from` attribute (case-sensitive). (optional)
	filterTo := "1234567890" // string | If present, recordings will be filtered to those with a matching `to` attribute (case-sensitive). (optional)
	filterConnectionId := "175237942907135762" // string | If present, recordings will be filtered to those with a matching `connection_id` attribute (case-sensitive). (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.GetRecordings(context.Background()).FilterConferenceId(filterConferenceId).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtLte(filterCreatedAtLte).FilterCallLegId(filterCallLegId).FilterCallSessionId(filterCallSessionId).FilterFrom(filterFrom).FilterTo(filterTo).FilterConnectionId(filterConnectionId).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.GetRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordings`: GetRecordings200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.GetRecordings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterConferenceId** | **string** | Returns only recordings associated with a given conference. | 
 **filterCreatedAtGte** | **string** | Returns only recordings created later than or at given ISO 8601 datetime. | 
 **filterCreatedAtLte** | **string** | Returns only recordings created earlier than or at given ISO 8601 datetime. | 
 **filterCallLegId** | **string** | If present, recordings will be filtered to those with a matching call_leg_id. | 
 **filterCallSessionId** | **string** | If present, recordings will be filtered to those with a matching call_session_id. | 
 **filterFrom** | **string** | If present, recordings will be filtered to those with a matching &#x60;from&#x60; attribute (case-sensitive). | 
 **filterTo** | **string** | If present, recordings will be filtered to those with a matching &#x60;to&#x60; attribute (case-sensitive). | 
 **filterConnectionId** | **string** | If present, recordings will be filtered to those with a matching &#x60;connection_id&#x60; attribute (case-sensitive). | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**GetRecordings200Response**](GetRecordings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomStorageCredentials

> CredentialsResponse UpdateCustomStorageCredentials(ctx, connectionId).CustomStorageConfiguration(customStorageConfiguration).Execute()

Update a stored credential



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
	connectionId := "connectionId_example" // string | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource.
	customStorageConfiguration := *openapiclient.NewCustomStorageConfiguration("gcs", openapiclient.CustomStorageConfiguration_configuration{AzureConfigurationData: openapiclient.NewAzureConfigurationData()}) // CustomStorageConfiguration | Creates new credentials resource for the specified connection_id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.UpdateCustomStorageCredentials(context.Background(), connectionId).CustomStorageConfiguration(customStorageConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.UpdateCustomStorageCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomStorageCredentials`: CredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.UpdateCustomStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customStorageConfiguration** | [**CustomStorageConfiguration**](CustomStorageConfiguration.md) | Creates new credentials resource for the specified connection_id. | 

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

