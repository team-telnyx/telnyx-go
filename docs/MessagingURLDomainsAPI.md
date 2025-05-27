# \MessagingURLDomainsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMessagingUrlDomains**](MessagingURLDomainsAPI.md#ListMessagingUrlDomains) | **Get** /messaging_url_domains | List messaging URL domains



## ListMessagingUrlDomains

> ListMessagingProfileUrlDomainsResponse ListMessagingUrlDomains(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List messaging URL domains

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingURLDomainsAPI.ListMessagingUrlDomains(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingURLDomainsAPI.ListMessagingUrlDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessagingUrlDomains`: ListMessagingProfileUrlDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingURLDomainsAPI.ListMessagingUrlDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagingUrlDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListMessagingProfileUrlDomainsResponse**](ListMessagingProfileUrlDomainsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

