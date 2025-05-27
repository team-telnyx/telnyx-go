# \NumberConfigurationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateMessagingSettingsOnPhoneNumbers**](NumberConfigurationsAPI.md#BulkUpdateMessagingSettingsOnPhoneNumbers) | **Post** /messaging_numbers_bulk_updates | Update the messaging profile of multiple phone numbers
[**GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus**](NumberConfigurationsAPI.md#GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus) | **Get** /messaging_numbers_bulk_updates/{order_id} | Retrieve bulk update status
[**GetPhoneNumberMessagingSettings**](NumberConfigurationsAPI.md#GetPhoneNumberMessagingSettings) | **Get** /phone_numbers/{id}/messaging | Retrieve a phone number with messaging settings
[**ListPhoneNumbersWithMessagingSettings**](NumberConfigurationsAPI.md#ListPhoneNumbersWithMessagingSettings) | **Get** /phone_numbers/messaging | List phone numbers with messaging settings
[**UpdatePhoneNumberMessagingSettings**](NumberConfigurationsAPI.md#UpdatePhoneNumberMessagingSettings) | **Patch** /phone_numbers/{id}/messaging | Update the messaging profile and/or messaging product of a phone number



## BulkUpdateMessagingSettingsOnPhoneNumbers

> RetrieveBulkUpdateMessagingSettingsResponse BulkUpdateMessagingSettingsOnPhoneNumbers(ctx).BulkMessagingSettingsUpdatePhoneNumbersRequest(bulkMessagingSettingsUpdatePhoneNumbersRequest).Execute()

Update the messaging profile of multiple phone numbers

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
	bulkMessagingSettingsUpdatePhoneNumbersRequest := *openapiclient.NewBulkMessagingSettingsUpdatePhoneNumbersRequest("MessagingProfileId_example", []string{"Numbers_example"}) // BulkMessagingSettingsUpdatePhoneNumbersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberConfigurationsAPI.BulkUpdateMessagingSettingsOnPhoneNumbers(context.Background()).BulkMessagingSettingsUpdatePhoneNumbersRequest(bulkMessagingSettingsUpdatePhoneNumbersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberConfigurationsAPI.BulkUpdateMessagingSettingsOnPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateMessagingSettingsOnPhoneNumbers`: RetrieveBulkUpdateMessagingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `NumberConfigurationsAPI.BulkUpdateMessagingSettingsOnPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateMessagingSettingsOnPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkMessagingSettingsUpdatePhoneNumbersRequest** | [**BulkMessagingSettingsUpdatePhoneNumbersRequest**](BulkMessagingSettingsUpdatePhoneNumbersRequest.md) |  | 

### Return type

[**RetrieveBulkUpdateMessagingSettingsResponse**](RetrieveBulkUpdateMessagingSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus

> RetrieveBulkUpdateMessagingSettingsResponse GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus(ctx, orderId).Execute()

Retrieve bulk update status

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
	orderId := "orderId_example" // string | Order ID to verify bulk update status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberConfigurationsAPI.GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberConfigurationsAPI.GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus`: RetrieveBulkUpdateMessagingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `NumberConfigurationsAPI.GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Order ID to verify bulk update status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkUpdateMessagingSettingsOnPhoneNumbersStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveBulkUpdateMessagingSettingsResponse**](RetrieveBulkUpdateMessagingSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhoneNumberMessagingSettings

> RetrieveMessagingSettingsResponse GetPhoneNumberMessagingSettings(ctx, id).Execute()

Retrieve a phone number with messaging settings

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
	id := "id_example" // string | Identifies the type of resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberConfigurationsAPI.GetPhoneNumberMessagingSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberConfigurationsAPI.GetPhoneNumberMessagingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPhoneNumberMessagingSettings`: RetrieveMessagingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `NumberConfigurationsAPI.GetPhoneNumberMessagingSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the type of resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhoneNumberMessagingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveMessagingSettingsResponse**](RetrieveMessagingSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumbersWithMessagingSettings

> ListMessagingSettingsResponse ListPhoneNumbersWithMessagingSettings(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List phone numbers with messaging settings

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
	resp, r, err := apiClient.NumberConfigurationsAPI.ListPhoneNumbersWithMessagingSettings(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberConfigurationsAPI.ListPhoneNumbersWithMessagingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPhoneNumbersWithMessagingSettings`: ListMessagingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `NumberConfigurationsAPI.ListPhoneNumbersWithMessagingSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhoneNumbersWithMessagingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListMessagingSettingsResponse**](ListMessagingSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumberMessagingSettings

> RetrieveMessagingSettingsResponse UpdatePhoneNumberMessagingSettings(ctx, id).UpdatePhoneNumberMessagingSettingsRequest(updatePhoneNumberMessagingSettingsRequest).Execute()

Update the messaging profile and/or messaging product of a phone number

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
	id := "id_example" // string | The phone number to update.
	updatePhoneNumberMessagingSettingsRequest := *openapiclient.NewUpdatePhoneNumberMessagingSettingsRequest() // UpdatePhoneNumberMessagingSettingsRequest | The new configuration to set for this phone number.  To avoid modifying a value, either omit the field or set its value to `null`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberConfigurationsAPI.UpdatePhoneNumberMessagingSettings(context.Background(), id).UpdatePhoneNumberMessagingSettingsRequest(updatePhoneNumberMessagingSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberConfigurationsAPI.UpdatePhoneNumberMessagingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePhoneNumberMessagingSettings`: RetrieveMessagingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `NumberConfigurationsAPI.UpdatePhoneNumberMessagingSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The phone number to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneNumberMessagingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePhoneNumberMessagingSettingsRequest** | [**UpdatePhoneNumberMessagingSettingsRequest**](UpdatePhoneNumberMessagingSettingsRequest.md) | The new configuration to set for this phone number.  To avoid modifying a value, either omit the field or set its value to &#x60;null&#x60;. | 

### Return type

[**RetrieveMessagingSettingsResponse**](RetrieveMessagingSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

