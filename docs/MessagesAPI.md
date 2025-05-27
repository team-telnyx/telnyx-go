# \MessagesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupMmsMessage**](MessagesAPI.md#CreateGroupMmsMessage) | **Post** /messages/group_mms | Send a group MMS message
[**CreateLongCodeMessage**](MessagesAPI.md#CreateLongCodeMessage) | **Post** /messages/long_code | Send a long code message
[**CreateNumberPoolMessage**](MessagesAPI.md#CreateNumberPoolMessage) | **Post** /messages/number_pool | Send a message using number pool
[**CreateShortCodeMessage**](MessagesAPI.md#CreateShortCodeMessage) | **Post** /messages/short_code | Send a short code message
[**GetMessage**](MessagesAPI.md#GetMessage) | **Get** /messages/{id} | Retrieve a message
[**SendMessage**](MessagesAPI.md#SendMessage) | **Post** /messages | Send a message



## CreateGroupMmsMessage

> MessageResponse CreateGroupMmsMessage(ctx).CreateGroupMMSMessageRequest(createGroupMMSMessageRequest).Execute()

Send a group MMS message

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
	createGroupMMSMessageRequest := *openapiclient.NewCreateGroupMMSMessageRequest("From_example", []string{"+E.164"}) // CreateGroupMMSMessageRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateGroupMmsMessage(context.Background()).CreateGroupMMSMessageRequest(createGroupMMSMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateGroupMmsMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupMmsMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateGroupMmsMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMmsMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupMMSMessageRequest** | [**CreateGroupMMSMessageRequest**](CreateGroupMMSMessageRequest.md) | Message payload | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLongCodeMessage

> MessageResponse CreateLongCodeMessage(ctx).CreateLongCodeMessageRequest(createLongCodeMessageRequest).Execute()

Send a long code message

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
	createLongCodeMessageRequest := *openapiclient.NewCreateLongCodeMessageRequest("From_example", "+E.164") // CreateLongCodeMessageRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateLongCodeMessage(context.Background()).CreateLongCodeMessageRequest(createLongCodeMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateLongCodeMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLongCodeMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateLongCodeMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLongCodeMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLongCodeMessageRequest** | [**CreateLongCodeMessageRequest**](CreateLongCodeMessageRequest.md) | Message payload | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNumberPoolMessage

> MessageResponse CreateNumberPoolMessage(ctx).CreateNumberPoolMessageRequest(createNumberPoolMessageRequest).Execute()

Send a message using number pool

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
	createNumberPoolMessageRequest := *openapiclient.NewCreateNumberPoolMessageRequest("MessagingProfileId_example", "+E.164") // CreateNumberPoolMessageRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateNumberPoolMessage(context.Background()).CreateNumberPoolMessageRequest(createNumberPoolMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateNumberPoolMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNumberPoolMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateNumberPoolMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNumberPoolMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNumberPoolMessageRequest** | [**CreateNumberPoolMessageRequest**](CreateNumberPoolMessageRequest.md) | Message payload | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShortCodeMessage

> MessageResponse CreateShortCodeMessage(ctx).CreateShortCodeMessageRequest(createShortCodeMessageRequest).Execute()

Send a short code message

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
	createShortCodeMessageRequest := *openapiclient.NewCreateShortCodeMessageRequest("From_example", "+E.164") // CreateShortCodeMessageRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateShortCodeMessage(context.Background()).CreateShortCodeMessageRequest(createShortCodeMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateShortCodeMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShortCodeMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateShortCodeMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShortCodeMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShortCodeMessageRequest** | [**CreateShortCodeMessageRequest**](CreateShortCodeMessageRequest.md) | Message payload | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessage

> GetMessage200Response GetMessage(ctx, id).Execute()

Retrieve a message



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.GetMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessage`: GetMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the message | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMessage200Response**](GetMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> MessageResponse SendMessage(ctx).CreateMessageRequest(createMessageRequest).Execute()

Send a message



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
	createMessageRequest := *openapiclient.NewCreateMessageRequest("+E.164") // CreateMessageRequest | Message payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.SendMessage(context.Background()).CreateMessageRequest(createMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SendMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMessageRequest** | [**CreateMessageRequest**](CreateMessageRequest.md) | Message payload | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

