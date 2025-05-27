# \ChatAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatPublicChatCompletionsPost**](ChatAPI.md#ChatPublicChatCompletionsPost) | **Post** /ai/chat/completions | Create a chat completion
[**GetModelsPublicModelsGet**](ChatAPI.md#GetModelsPublicModelsGet) | **Get** /ai/models | Get available models
[**PostSummary**](ChatAPI.md#PostSummary) | **Post** /ai/summarize | Summarize file content



## ChatPublicChatCompletionsPost

> interface{} ChatPublicChatCompletionsPost(ctx).ChatCompletionRequest(chatCompletionRequest).Execute()

Create a chat completion



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
	chatCompletionRequest := *openapiclient.NewChatCompletionRequest([]openapiclient.ChatCompletionSystemMessageParam{*openapiclient.NewChatCompletionSystemMessageParam(openapiclient.ChatCompletionSystemMessageParam_content{ArrayOfTextAndImageArrayInner: new([]TextAndImageArrayInner)}, "Role_example")}) // ChatCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.ChatPublicChatCompletionsPost(context.Background()).ChatCompletionRequest(chatCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.ChatPublicChatCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPublicChatCompletionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.ChatPublicChatCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPublicChatCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 

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


## GetModelsPublicModelsGet

> ModelsResponse GetModelsPublicModelsGet(ctx).Execute()

Get available models



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
	resp, r, err := apiClient.ChatAPI.GetModelsPublicModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.GetModelsPublicModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsPublicModelsGet`: ModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.GetModelsPublicModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsPublicModelsGetRequest struct via the builder pattern


### Return type

[**ModelsResponse**](ModelsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSummary

> SummaryResponseData PostSummary(ctx).SummaryRequest(summaryRequest).Execute()

Summarize file content



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
	summaryRequest := *openapiclient.NewSummaryRequest("Bucket_example", "Filename_example") // SummaryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.PostSummary(context.Background()).SummaryRequest(summaryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.PostSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSummary`: SummaryResponseData
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.PostSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **summaryRequest** | [**SummaryRequest**](SummaryRequest.md) |  | 

### Return type

[**SummaryResponseData**](SummaryResponseData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

