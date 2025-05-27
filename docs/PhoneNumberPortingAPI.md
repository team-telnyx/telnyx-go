# \PhoneNumberPortingAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPortabilityCheck**](PhoneNumberPortingAPI.md#PostPortabilityCheck) | **Post** /portability_checks | Run a portability check



## PostPortabilityCheck

> PostPortabilityCheck201Response PostPortabilityCheck(ctx).PostPortabilityCheckRequest(postPortabilityCheckRequest).Execute()

Run a portability check



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
	postPortabilityCheckRequest := *openapiclient.NewPostPortabilityCheckRequest() // PostPortabilityCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberPortingAPI.PostPortabilityCheck(context.Background()).PostPortabilityCheckRequest(postPortabilityCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberPortingAPI.PostPortabilityCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPortabilityCheck`: PostPortabilityCheck201Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberPortingAPI.PostPortabilityCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPortabilityCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postPortabilityCheckRequest** | [**PostPortabilityCheckRequest**](PostPortabilityCheckRequest.md) |  | 

### Return type

[**PostPortabilityCheck201Response**](PostPortabilityCheck201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

