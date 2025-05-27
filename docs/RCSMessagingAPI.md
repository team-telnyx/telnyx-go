# \RCSMessagingAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MesssagesRcsPost**](RCSMessagingAPI.md#MesssagesRcsPost) | **Post** /messsages/rcs | Send an RCS message



## MesssagesRcsPost

> RCSResponse MesssagesRcsPost(ctx).RCSMessage(rCSMessage).Execute()

Send an RCS message

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
	rCSMessage := *openapiclient.NewRCSMessage("Agent007", "+13125551234", "MessagingProfileId_example", *openapiclient.NewRCSAgentMessage()) // RCSMessage | RCS message body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RCSMessagingAPI.MesssagesRcsPost(context.Background()).RCSMessage(rCSMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RCSMessagingAPI.MesssagesRcsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MesssagesRcsPost`: RCSResponse
	fmt.Fprintf(os.Stdout, "Response from `RCSMessagingAPI.MesssagesRcsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMesssagesRcsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rCSMessage** | [**RCSMessage**](RCSMessage.md) | RCS message body | 

### Return type

[**RCSResponse**](RCSResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

