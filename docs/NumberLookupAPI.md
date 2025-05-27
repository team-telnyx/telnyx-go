# \NumberLookupAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupNumber**](NumberLookupAPI.md#LookupNumber) | **Get** /number_lookup/{phone_number} | Lookup phone number data



## LookupNumber

> NumberLookupResponse LookupNumber(ctx, phoneNumber).Type_(type_).Execute()

Lookup phone number data



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
	phoneNumber := "+18665552368" // string | The phone number to be looked up
	type_ := "type__example" // string | Specifies the type of number lookup to be performed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberLookupAPI.LookupNumber(context.Background(), phoneNumber).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberLookupAPI.LookupNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupNumber`: NumberLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `NumberLookupAPI.LookupNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | The phone number to be looked up | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Specifies the type of number lookup to be performed | 

### Return type

[**NumberLookupResponse**](NumberLookupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

