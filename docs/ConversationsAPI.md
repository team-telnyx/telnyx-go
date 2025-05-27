# \ConversationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewConversationPublicConversationsPost**](ConversationsAPI.md#CreateNewConversationPublicConversationsPost) | **Post** /ai/conversations | Create a conversation
[**DeleteConversationByIdPublicConversationsDelete**](ConversationsAPI.md#DeleteConversationByIdPublicConversationsDelete) | **Delete** /ai/conversations/{conversation_id} | Delete a conversation
[**GetConversationByIdPublicConversationsGet**](ConversationsAPI.md#GetConversationByIdPublicConversationsGet) | **Get** /ai/conversations/{conversation_id} | Get a conversation
[**GetConversationsPublicConversationIdInsightsGet**](ConversationsAPI.md#GetConversationsPublicConversationIdInsightsGet) | **Get** /ai/conversations/{conversation_id}/conversations-insights | Get insights for a conversation
[**GetConversationsPublicConversationIdMessagesGet**](ConversationsAPI.md#GetConversationsPublicConversationIdMessagesGet) | **Get** /ai/conversations/{conversation_id}/messages | Get conversation messages
[**GetConversationsPublicConversationsGet**](ConversationsAPI.md#GetConversationsPublicConversationsGet) | **Get** /ai/conversations | List conversations
[**UpdateConversationByIdPublicConversationsPut**](ConversationsAPI.md#UpdateConversationByIdPublicConversationsPut) | **Put** /ai/conversations/{conversation_id} | Update conversation metadata



## CreateNewConversationPublicConversationsPost

> Conversation CreateNewConversationPublicConversationsPost(ctx).CreateConversationRequest(createConversationRequest).Execute()

Create a conversation



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
	createConversationRequest := *openapiclient.NewCreateConversationRequest() // CreateConversationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.CreateNewConversationPublicConversationsPost(context.Background()).CreateConversationRequest(createConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.CreateNewConversationPublicConversationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewConversationPublicConversationsPost`: Conversation
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.CreateNewConversationPublicConversationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewConversationPublicConversationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConversationRequest** | [**CreateConversationRequest**](CreateConversationRequest.md) |  | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversationByIdPublicConversationsDelete

> DeleteConversationByIdPublicConversationsDelete(ctx, conversationId).Execute()

Delete a conversation



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
	conversationId := "conversationId_example" // string | The ID of the conversation to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConversationsAPI.DeleteConversationByIdPublicConversationsDelete(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.DeleteConversationByIdPublicConversationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The ID of the conversation to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConversationByIdPublicConversationsDeleteRequest struct via the builder pattern


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


## GetConversationByIdPublicConversationsGet

> GetConversationByIdPublicConversationsGet200Response GetConversationByIdPublicConversationsGet(ctx, conversationId).Execute()

Get a conversation



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
	conversationId := "conversationId_example" // string | The ID of the conversation to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationByIdPublicConversationsGet(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationByIdPublicConversationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationByIdPublicConversationsGet`: GetConversationByIdPublicConversationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationByIdPublicConversationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The ID of the conversation to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationByIdPublicConversationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConversationByIdPublicConversationsGet200Response**](GetConversationByIdPublicConversationsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsPublicConversationIdInsightsGet

> ConversationInsightListData GetConversationsPublicConversationIdInsightsGet(ctx, conversationId).Execute()

Get insights for a conversation



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
	conversationId := "conversationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsPublicConversationIdInsightsGet(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsPublicConversationIdInsightsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsPublicConversationIdInsightsGet`: ConversationInsightListData
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsPublicConversationIdInsightsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsPublicConversationIdInsightsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConversationInsightListData**](ConversationInsightListData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsPublicConversationIdMessagesGet

> ConversationMessageListData GetConversationsPublicConversationIdMessagesGet(ctx, conversationId).Execute()

Get conversation messages



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
	conversationId := "conversationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsPublicConversationIdMessagesGet(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsPublicConversationIdMessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsPublicConversationIdMessagesGet`: ConversationMessageListData
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsPublicConversationIdMessagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsPublicConversationIdMessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConversationMessageListData**](ConversationMessageListData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsPublicConversationsGet

> ConversationsListData GetConversationsPublicConversationsGet(ctx).Id(id).Name(name).CreatedAt(createdAt).LastMessageAt(lastMessageAt).MetadataAssistantId(metadataAssistantId).MetadataCallControlId(metadataCallControlId).MetadataTelnyxAgentTarget(metadataTelnyxAgentTarget).MetadataTelnyxEndUserTarget(metadataTelnyxEndUserTarget).MetadataTelnyxConversationChannel(metadataTelnyxConversationChannel).Limit(limit).Order(order).Or(or).Execute()

List conversations



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
	id := "id_example" // string | Filter by conversation ID (e.g. id=eq.123) (optional)
	name := "name_example" // string | Filter by conversation Name (e.g. `name=like.Voice%`) (optional)
	createdAt := "createdAt_example" // string | Filter by creation datetime (e.g., `created_at=gte.2025-01-01`) (optional)
	lastMessageAt := "lastMessageAt_example" // string | Filter by last message datetime (e.g., `last_message_at=lte.2025-06-01`) (optional)
	metadataAssistantId := "metadataAssistantId_example" // string | Filter by assistant ID (e.g., `metadata->assistant_id=eq.assistant-123`) (optional)
	metadataCallControlId := "metadataCallControlId_example" // string | Filter by call control ID (e.g., `metadata->call_control_id=eq.v3:123`) (optional)
	metadataTelnyxAgentTarget := "metadataTelnyxAgentTarget_example" // string | Filter by the phone number, SIP URI, or other identifier for the agent (e.g., `metadata->telnyx_agent_target=eq.+13128675309`) (optional)
	metadataTelnyxEndUserTarget := "metadataTelnyxEndUserTarget_example" // string | Filter by the phone number, SIP URI, or other identifier for the end user (e.g., `metadata->telnyx_end_user_target=eq.+13128675309`) (optional)
	metadataTelnyxConversationChannel := "metadataTelnyxConversationChannel_example" // string | Filter by conversation channel (e.g., `metadata->telnyx_conversation_channel=eq.phone_call`) (optional)
	limit := int32(56) // int32 | Limit the number of returned conversations (e.g., `limit=10`) (optional)
	order := "order_example" // string | Order the results by specific fields (e.g., `order=created_at.desc` or `order=last_message_at.asc`) (optional)
	or := "or_example" // string | Apply OR conditions using PostgREST syntax (e.g., `or=(created_at.gte.2025-04-01,last_message_at.gte.2025-04-01)`) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsPublicConversationsGet(context.Background()).Id(id).Name(name).CreatedAt(createdAt).LastMessageAt(lastMessageAt).MetadataAssistantId(metadataAssistantId).MetadataCallControlId(metadataCallControlId).MetadataTelnyxAgentTarget(metadataTelnyxAgentTarget).MetadataTelnyxEndUserTarget(metadataTelnyxEndUserTarget).MetadataTelnyxConversationChannel(metadataTelnyxConversationChannel).Limit(limit).Order(order).Or(or).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsPublicConversationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsPublicConversationsGet`: ConversationsListData
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsPublicConversationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsPublicConversationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by conversation ID (e.g. id&#x3D;eq.123) | 
 **name** | **string** | Filter by conversation Name (e.g. &#x60;name&#x3D;like.Voice%&#x60;) | 
 **createdAt** | **string** | Filter by creation datetime (e.g., &#x60;created_at&#x3D;gte.2025-01-01&#x60;) | 
 **lastMessageAt** | **string** | Filter by last message datetime (e.g., &#x60;last_message_at&#x3D;lte.2025-06-01&#x60;) | 
 **metadataAssistantId** | **string** | Filter by assistant ID (e.g., &#x60;metadata-&gt;assistant_id&#x3D;eq.assistant-123&#x60;) | 
 **metadataCallControlId** | **string** | Filter by call control ID (e.g., &#x60;metadata-&gt;call_control_id&#x3D;eq.v3:123&#x60;) | 
 **metadataTelnyxAgentTarget** | **string** | Filter by the phone number, SIP URI, or other identifier for the agent (e.g., &#x60;metadata-&gt;telnyx_agent_target&#x3D;eq.+13128675309&#x60;) | 
 **metadataTelnyxEndUserTarget** | **string** | Filter by the phone number, SIP URI, or other identifier for the end user (e.g., &#x60;metadata-&gt;telnyx_end_user_target&#x3D;eq.+13128675309&#x60;) | 
 **metadataTelnyxConversationChannel** | **string** | Filter by conversation channel (e.g., &#x60;metadata-&gt;telnyx_conversation_channel&#x3D;eq.phone_call&#x60;) | 
 **limit** | **int32** | Limit the number of returned conversations (e.g., &#x60;limit&#x3D;10&#x60;) | 
 **order** | **string** | Order the results by specific fields (e.g., &#x60;order&#x3D;created_at.desc&#x60; or &#x60;order&#x3D;last_message_at.asc&#x60;) | 
 **or** | **string** | Apply OR conditions using PostgREST syntax (e.g., &#x60;or&#x3D;(created_at.gte.2025-04-01,last_message_at.gte.2025-04-01)&#x60;) | 

### Return type

[**ConversationsListData**](ConversationsListData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationByIdPublicConversationsPut

> GetConversationByIdPublicConversationsGet200Response UpdateConversationByIdPublicConversationsPut(ctx, conversationId).UpdateConversationRequest(updateConversationRequest).Execute()

Update conversation metadata



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
	conversationId := "conversationId_example" // string | The ID of the conversation to update
	updateConversationRequest := *openapiclient.NewUpdateConversationRequest() // UpdateConversationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.UpdateConversationByIdPublicConversationsPut(context.Background(), conversationId).UpdateConversationRequest(updateConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.UpdateConversationByIdPublicConversationsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConversationByIdPublicConversationsPut`: GetConversationByIdPublicConversationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.UpdateConversationByIdPublicConversationsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The ID of the conversation to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConversationByIdPublicConversationsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConversationRequest** | [**UpdateConversationRequest**](UpdateConversationRequest.md) |  | 

### Return type

[**GetConversationByIdPublicConversationsGet200Response**](GetConversationByIdPublicConversationsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

