# \EnumAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnumEndpoint**](EnumAPI.md#GetEnumEndpoint) | **Get** /enum/{endpoint} | Get Enum



## GetEnumEndpoint

> GetEnumEndpoint200Response GetEnumEndpoint(ctx, endpoint).Execute()

Get Enum

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
	endpoint := "endpoint_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnumAPI.GetEnumEndpoint(context.Background(), endpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnumAPI.GetEnumEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnumEndpoint`: GetEnumEndpoint200Response
	fmt.Fprintf(os.Stdout, "Response from `EnumAPI.GetEnumEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpoint** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnumEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnumEndpoint200Response**](GetEnumEndpoint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

