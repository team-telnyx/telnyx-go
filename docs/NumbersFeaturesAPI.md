# \NumbersFeaturesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostNumbersFeatures**](NumbersFeaturesAPI.md#PostNumbersFeatures) | **Post** /numbers_features | Retrieve the features for a list of numbers



## PostNumbersFeatures

> PostNumbersFeatures200Response PostNumbersFeatures(ctx).PostNumbersFeaturesRequest(postNumbersFeaturesRequest).Execute()

Retrieve the features for a list of numbers



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
	postNumbersFeaturesRequest := *openapiclient.NewPostNumbersFeaturesRequest([]string{"PhoneNumbers_example"}) // PostNumbersFeaturesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumbersFeaturesAPI.PostNumbersFeatures(context.Background()).PostNumbersFeaturesRequest(postNumbersFeaturesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumbersFeaturesAPI.PostNumbersFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNumbersFeatures`: PostNumbersFeatures200Response
	fmt.Fprintf(os.Stdout, "Response from `NumbersFeaturesAPI.PostNumbersFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNumbersFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postNumbersFeaturesRequest** | [**PostNumbersFeaturesRequest**](PostNumbersFeaturesRequest.md) |  | 

### Return type

[**PostNumbersFeatures200Response**](PostNumbersFeatures200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

