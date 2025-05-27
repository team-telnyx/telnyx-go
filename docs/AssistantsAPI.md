# \AssistantsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewAssistantPublicAssistantsPost**](AssistantsAPI.md#CreateNewAssistantPublicAssistantsPost) | **Post** /ai/assistants | Create an assistant
[**CreateScheduledEvent**](AssistantsAPI.md#CreateScheduledEvent) | **Post** /ai/assistants/{assistant_id}/scheduled_events | Create a scheduled event
[**DeleteAssistantPublicAssistantsAssistantIdDelete**](AssistantsAPI.md#DeleteAssistantPublicAssistantsAssistantIdDelete) | **Delete** /ai/assistants/{assistant_id} | Delete an assistant
[**DeleteScheduledEvent**](AssistantsAPI.md#DeleteScheduledEvent) | **Delete** /ai/assistants/{assistant_id}/scheduled_events/{event_id} | Delete a scheduled event
[**GetAssistantPublicAssistantsAssistantIdGet**](AssistantsAPI.md#GetAssistantPublicAssistantsAssistantIdGet) | **Get** /ai/assistants/{assistant_id} | Get an assistant
[**GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet**](AssistantsAPI.md#GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet) | **Get** /ai/assistants/{assistant_id}/texml | Get assistant texml
[**GetAssistantsPublicAssistantsGet**](AssistantsAPI.md#GetAssistantsPublicAssistantsGet) | **Get** /ai/assistants | List assistants
[**GetScheduledEvent**](AssistantsAPI.md#GetScheduledEvent) | **Get** /ai/assistants/{assistant_id}/scheduled_events/{event_id} | Get a scheduled event
[**GetScheduledEvents**](AssistantsAPI.md#GetScheduledEvents) | **Get** /ai/assistants/{assistant_id}/scheduled_events | List scheduled events
[**UpdateAssistantPublicAssistantsAssistantIdPost**](AssistantsAPI.md#UpdateAssistantPublicAssistantsAssistantIdPost) | **Post** /ai/assistants/{assistant_id} | Update an assistant



## CreateNewAssistantPublicAssistantsPost

> Assistant CreateNewAssistantPublicAssistantsPost(ctx).CreateAssistantRequest(createAssistantRequest).Execute()

Create an assistant



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
	createAssistantRequest := *openapiclient.NewCreateAssistantRequest("Name_example", "Model_example", "Instructions_example") // CreateAssistantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.CreateNewAssistantPublicAssistantsPost(context.Background()).CreateAssistantRequest(createAssistantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.CreateNewAssistantPublicAssistantsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewAssistantPublicAssistantsPost`: Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.CreateNewAssistantPublicAssistantsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewAssistantPublicAssistantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssistantRequest** | [**CreateAssistantRequest**](CreateAssistantRequest.md) |  | 

### Return type

[**Assistant**](Assistant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScheduledEvent

> ScheduledEventResponse CreateScheduledEvent(ctx, assistantId).CreateScheduledEventRequest(createScheduledEventRequest).Execute()

Create a scheduled event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	assistantId := "assistantId_example" // string | 
	createScheduledEventRequest := *openapiclient.NewCreateScheduledEventRequest(openapiclient.ConversationChannelType("phone_call"), "TelnyxEndUserTarget_example", "TelnyxAgentTarget_example", time.Now()) // CreateScheduledEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.CreateScheduledEvent(context.Background(), assistantId).CreateScheduledEventRequest(createScheduledEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.CreateScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScheduledEvent`: ScheduledEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.CreateScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createScheduledEventRequest** | [**CreateScheduledEventRequest**](CreateScheduledEventRequest.md) |  | 

### Return type

[**ScheduledEventResponse**](ScheduledEventResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistantPublicAssistantsAssistantIdDelete

> AssistantDeletedResponse DeleteAssistantPublicAssistantsAssistantIdDelete(ctx, assistantId).Execute()

Delete an assistant



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
	assistantId := "assistantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.DeleteAssistantPublicAssistantsAssistantIdDelete(context.Background(), assistantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.DeleteAssistantPublicAssistantsAssistantIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssistantPublicAssistantsAssistantIdDelete`: AssistantDeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.DeleteAssistantPublicAssistantsAssistantIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssistantPublicAssistantsAssistantIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssistantDeletedResponse**](AssistantDeletedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledEvent

> interface{} DeleteScheduledEvent(ctx, assistantId, eventId).Execute()

Delete a scheduled event



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
	assistantId := "assistantId_example" // string | 
	eventId := "eventId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.DeleteScheduledEvent(context.Background(), assistantId, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.DeleteScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteScheduledEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.DeleteScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAssistantPublicAssistantsAssistantIdGet

> Assistant GetAssistantPublicAssistantsAssistantIdGet(ctx, assistantId).FetchDynamicVariablesFromWebhook(fetchDynamicVariablesFromWebhook).From(from).To(to).CallControlId(callControlId).Execute()

Get an assistant



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
	assistantId := "assistantId_example" // string | 
	fetchDynamicVariablesFromWebhook := true // bool |  (optional) (default to false)
	from := "from_example" // string |  (optional)
	to := "to_example" // string |  (optional)
	callControlId := "callControlId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantPublicAssistantsAssistantIdGet(context.Background(), assistantId).FetchDynamicVariablesFromWebhook(fetchDynamicVariablesFromWebhook).From(from).To(to).CallControlId(callControlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantPublicAssistantsAssistantIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantPublicAssistantsAssistantIdGet`: Assistant
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantPublicAssistantsAssistantIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantPublicAssistantsAssistantIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchDynamicVariablesFromWebhook** | **bool** |  | [default to false]
 **from** | **string** |  | 
 **to** | **string** |  | 
 **callControlId** | **string** |  | 

### Return type

[**Assistant**](Assistant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet

> string GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet(ctx, assistantId).Execute()

Get assistant texml



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
	assistantId := "assistantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet(context.Background(), assistantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet`: string
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantTexmlPublicAssistantsAssistantIdTexmlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantsPublicAssistantsGet

> AssistantsListData GetAssistantsPublicAssistantsGet(ctx).Execute()

List assistants



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
	resp, r, err := apiClient.AssistantsAPI.GetAssistantsPublicAssistantsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetAssistantsPublicAssistantsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantsPublicAssistantsGet`: AssistantsListData
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetAssistantsPublicAssistantsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantsPublicAssistantsGetRequest struct via the builder pattern


### Return type

[**AssistantsListData**](AssistantsListData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledEvent

> ScheduledEventResponse GetScheduledEvent(ctx, assistantId, eventId).Execute()

Get a scheduled event



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
	assistantId := "assistantId_example" // string | 
	eventId := "eventId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetScheduledEvent(context.Background(), assistantId, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledEvent`: ScheduledEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScheduledEventResponse**](ScheduledEventResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledEvents

> PaginatedScheduledEventList GetScheduledEvents(ctx, assistantId).PageSize(pageSize).PageNumber(pageNumber).FromDate(fromDate).ToDate(toDate).ConversationChannel(conversationChannel).Execute()

List scheduled events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	assistantId := "assistantId_example" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	fromDate := time.Now() // time.Time |  (optional)
	toDate := time.Now() // time.Time |  (optional)
	conversationChannel := openapiclient.ConversationChannelType("phone_call") // ConversationChannelType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.GetScheduledEvents(context.Background(), assistantId).PageSize(pageSize).PageNumber(pageNumber).FromDate(fromDate).ToDate(toDate).ConversationChannel(conversationChannel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.GetScheduledEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledEvents`: PaginatedScheduledEventList
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.GetScheduledEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 20]
 **pageNumber** | **int32** |  | [default to 1]
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 
 **conversationChannel** | [**ConversationChannelType**](ConversationChannelType.md) |  | 

### Return type

[**PaginatedScheduledEventList**](PaginatedScheduledEventList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistantPublicAssistantsAssistantIdPost

> interface{} UpdateAssistantPublicAssistantsAssistantIdPost(ctx, assistantId).UpdateAssistantRequest(updateAssistantRequest).Execute()

Update an assistant



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
	assistantId := "assistantId_example" // string | 
	updateAssistantRequest := *openapiclient.NewUpdateAssistantRequest() // UpdateAssistantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantsAPI.UpdateAssistantPublicAssistantsAssistantIdPost(context.Background(), assistantId).UpdateAssistantRequest(updateAssistantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantsAPI.UpdateAssistantPublicAssistantsAssistantIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssistantPublicAssistantsAssistantIdPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssistantsAPI.UpdateAssistantPublicAssistantsAssistantIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssistantPublicAssistantsAssistantIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAssistantRequest** | [**UpdateAssistantRequest**](UpdateAssistantRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

