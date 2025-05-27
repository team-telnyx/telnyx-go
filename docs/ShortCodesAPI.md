# \ShortCodesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListShortCodes**](ShortCodesAPI.md#ListShortCodes) | **Get** /short_codes | List short codes
[**RetrieveShortCode**](ShortCodesAPI.md#RetrieveShortCode) | **Get** /short_codes/{id} | Retrieve a short code
[**UpdateShortCode**](ShortCodesAPI.md#UpdateShortCode) | **Patch** /short_codes/{id} | Update short code



## ListShortCodes

> ListShortCodesResponse ListShortCodes(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterMessagingProfileId(filterMessagingProfileId).Execute()

List short codes

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
	filterMessagingProfileId := "filterMessagingProfileId_example" // string | Filter by Messaging Profile ID. Use the string `null` for phone numbers without assigned profiles. A synonym for the `/messaging_profiles/{id}/short_codes` endpoint when querying about an extant profile. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortCodesAPI.ListShortCodes(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterMessagingProfileId(filterMessagingProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortCodesAPI.ListShortCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListShortCodes`: ListShortCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortCodesAPI.ListShortCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListShortCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterMessagingProfileId** | **string** | Filter by Messaging Profile ID. Use the string &#x60;null&#x60; for phone numbers without assigned profiles. A synonym for the &#x60;/messaging_profiles/{id}/short_codes&#x60; endpoint when querying about an extant profile. | 

### Return type

[**ListShortCodesResponse**](ListShortCodesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveShortCode

> ShortCodeResponse RetrieveShortCode(ctx, id).Execute()

Retrieve a short code

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the short code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortCodesAPI.RetrieveShortCode(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortCodesAPI.RetrieveShortCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveShortCode`: ShortCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortCodesAPI.RetrieveShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the short code | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShortCodeResponse**](ShortCodeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ShortCodeResponse UpdateShortCode(ctx, id).UpdateShortCodeRequest(updateShortCodeRequest).Execute()

Update short code



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the short code
	updateShortCodeRequest := *openapiclient.NewUpdateShortCodeRequest("MessagingProfileId_example") // UpdateShortCodeRequest | Short code update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortCodesAPI.UpdateShortCode(context.Background(), id).UpdateShortCodeRequest(updateShortCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortCodesAPI.UpdateShortCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShortCode`: ShortCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortCodesAPI.UpdateShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the short code | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateShortCodeRequest** | [**UpdateShortCodeRequest**](UpdateShortCodeRequest.md) | Short code update | 

### Return type

[**ShortCodeResponse**](ShortCodeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

