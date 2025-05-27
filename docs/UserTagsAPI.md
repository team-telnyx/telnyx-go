# \UserTagsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserTags**](UserTagsAPI.md#GetUserTags) | **Get** /user_tags | List User Tags



## GetUserTags

> GetUserTags200Response GetUserTags(ctx).FilterStartsWith(filterStartsWith).Execute()

List User Tags



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
	filterStartsWith := "my-tag" // string | Filter tags by prefix (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTagsAPI.GetUserTags(context.Background()).FilterStartsWith(filterStartsWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagsAPI.GetUserTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTags`: GetUserTags200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTagsAPI.GetUserTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStartsWith** | **string** | Filter tags by prefix | 

### Return type

[**GetUserTags200Response**](GetUserTags200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

