# \PhoneNumberConfigurationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePhoneNumber**](PhoneNumberConfigurationsAPI.md#DeletePhoneNumber) | **Delete** /phone_numbers/{id} | Delete a phone number
[**EnablePhoneNumberEmergency**](PhoneNumberConfigurationsAPI.md#EnablePhoneNumberEmergency) | **Post** /phone_numbers/{id}/actions/enable_emergency | Enable emergency for a phone number
[**GetPhoneNumberVoiceSettings**](PhoneNumberConfigurationsAPI.md#GetPhoneNumberVoiceSettings) | **Get** /phone_numbers/{id}/voice | Retrieve a phone number with voice settings
[**ListPhoneNumbers**](PhoneNumberConfigurationsAPI.md#ListPhoneNumbers) | **Get** /phone_numbers | List phone numbers
[**ListPhoneNumbersWithVoiceSettings**](PhoneNumberConfigurationsAPI.md#ListPhoneNumbersWithVoiceSettings) | **Get** /phone_numbers/voice | List phone numbers with voice settings
[**PhoneNumberBundleStatusChange**](PhoneNumberConfigurationsAPI.md#PhoneNumberBundleStatusChange) | **Patch** /phone_numbers/{id}/actions/bundle_status_change | Change the bundle status for a phone number (set to being in a bundle or remove from a bundle)
[**RetrievePhoneNumber**](PhoneNumberConfigurationsAPI.md#RetrievePhoneNumber) | **Get** /phone_numbers/{id} | Retrieve a phone number
[**SlimListPhoneNumbers**](PhoneNumberConfigurationsAPI.md#SlimListPhoneNumbers) | **Get** /phone_numbers/slim | Slim List phone numbers
[**UpdatePhoneNumber**](PhoneNumberConfigurationsAPI.md#UpdatePhoneNumber) | **Patch** /phone_numbers/{id} | Update a phone number
[**UpdatePhoneNumberVoiceSettings**](PhoneNumberConfigurationsAPI.md#UpdatePhoneNumberVoiceSettings) | **Patch** /phone_numbers/{id}/voice | Update a phone number with voice settings



## DeletePhoneNumber

> PhoneNumberResponse1 DeletePhoneNumber(ctx, id).Execute()

Delete a phone number

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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.DeletePhoneNumber(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.DeletePhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePhoneNumber`: PhoneNumberResponse1
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.DeletePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneNumberResponse1**](PhoneNumberResponse1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePhoneNumberEmergency

> PhoneNumberEnableEmergency EnablePhoneNumberEmergency(ctx, id).PhoneNumberEnableEmergencyRequest(phoneNumberEnableEmergencyRequest).Execute()

Enable emergency for a phone number

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
	id := "1293384261075731499" // string | Identifies the resource.
	phoneNumberEnableEmergencyRequest := *openapiclient.NewPhoneNumberEnableEmergencyRequest(false, "EmergencyAddressId_example") // PhoneNumberEnableEmergencyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.EnablePhoneNumberEmergency(context.Background(), id).PhoneNumberEnableEmergencyRequest(phoneNumberEnableEmergencyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.EnablePhoneNumberEmergency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnablePhoneNumberEmergency`: PhoneNumberEnableEmergency
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.EnablePhoneNumberEmergency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePhoneNumberEmergencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumberEnableEmergencyRequest** | [**PhoneNumberEnableEmergencyRequest**](PhoneNumberEnableEmergencyRequest.md) |  | 

### Return type

[**PhoneNumberEnableEmergency**](PhoneNumberEnableEmergency.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhoneNumberVoiceSettings

> RetrievePhoneNumberVoiceResponse GetPhoneNumberVoiceSettings(ctx, id).Execute()

Retrieve a phone number with voice settings

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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.GetPhoneNumberVoiceSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.GetPhoneNumberVoiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPhoneNumberVoiceSettings`: RetrievePhoneNumberVoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.GetPhoneNumberVoiceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhoneNumberVoiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrievePhoneNumberVoiceResponse**](RetrievePhoneNumberVoiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumbers

> ListPhoneNumbersResponse ListPhoneNumbers(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterTag(filterTag).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterCountryIsoAlpha2(filterCountryIsoAlpha2).FilterConnectionId(filterConnectionId).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceConnectionNameStartsWith(filterVoiceConnectionNameStartsWith).FilterVoiceConnectionNameEndsWith(filterVoiceConnectionNameEndsWith).FilterVoiceConnectionNameEq(filterVoiceConnectionNameEq).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).FilterNumberTypeEq(filterNumberTypeEq).FilterSource(filterSource).Sort(sort).Execute()

List phone numbers

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
	filterTag := "filterTag_example" // string | Filter by phone number tags. (optional)
	filterPhoneNumber := "filterPhoneNumber_example" // string | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. (optional)
	filterStatus := "active" // string | Filter by phone number status. (optional)
	filterCountryIsoAlpha2 := openapiclient.ListPhoneNumbers_filter_country_iso_alpha2__parameter{ArrayOfString: new([]string)} // ListPhoneNumbersFilterCountryIsoAlpha2Parameter | Filter by phone number country ISO alpha-2 code. Can be a single value or an array of values. (optional)
	filterConnectionId := "1521916448077776306" // string | Filter by connection_id. (optional)
	filterVoiceConnectionNameContains := "test" // string | Filter contains connection name. Requires at least three characters. (optional)
	filterVoiceConnectionNameStartsWith := "test" // string | Filter starts with connection name. Requires at least three characters. (optional)
	filterVoiceConnectionNameEndsWith := "test" // string | Filter ends with connection name. Requires at least three characters. (optional)
	filterVoiceConnectionNameEq := "test" // string | Filter by connection name. (optional)
	filterVoiceUsagePaymentMethod := "channel" // string | Filter by usage_payment_method. (optional)
	filterBillingGroupId := "62e4bf2e-c278-4282-b524-488d9c9c43b2" // string | Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string 'null'. (optional)
	filterEmergencyAddressId := "9102160989215728032" // string | Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string 'null'. (optional)
	filterCustomerReference := "filterCustomerReference_example" // string | Filter numbers via the customer_reference set. (optional)
	filterNumberTypeEq := "filterNumberTypeEq_example" // string | Filter phone numbers by phone number type. (optional)
	filterSource := "filterSource_example" // string | Filter phone numbers by their source. Use 'ported' for numbers ported from other carriers, or 'purchased' for numbers bought directly from Telnyx. (optional)
	sort := "connection_name" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.ListPhoneNumbers(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterTag(filterTag).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterCountryIsoAlpha2(filterCountryIsoAlpha2).FilterConnectionId(filterConnectionId).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceConnectionNameStartsWith(filterVoiceConnectionNameStartsWith).FilterVoiceConnectionNameEndsWith(filterVoiceConnectionNameEndsWith).FilterVoiceConnectionNameEq(filterVoiceConnectionNameEq).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).FilterNumberTypeEq(filterNumberTypeEq).FilterSource(filterSource).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.ListPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPhoneNumbers`: ListPhoneNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.ListPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterTag** | **string** | Filter by phone number tags. | 
 **filterPhoneNumber** | **string** | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterStatus** | **string** | Filter by phone number status. | 
 **filterCountryIsoAlpha2** | [**ListPhoneNumbersFilterCountryIsoAlpha2Parameter**](ListPhoneNumbersFilterCountryIsoAlpha2Parameter.md) | Filter by phone number country ISO alpha-2 code. Can be a single value or an array of values. | 
 **filterConnectionId** | **string** | Filter by connection_id. | 
 **filterVoiceConnectionNameContains** | **string** | Filter contains connection name. Requires at least three characters. | 
 **filterVoiceConnectionNameStartsWith** | **string** | Filter starts with connection name. Requires at least three characters. | 
 **filterVoiceConnectionNameEndsWith** | **string** | Filter ends with connection name. Requires at least three characters. | 
 **filterVoiceConnectionNameEq** | **string** | Filter by connection name. | 
 **filterVoiceUsagePaymentMethod** | **string** | Filter by usage_payment_method. | 
 **filterBillingGroupId** | **string** | Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string &#39;null&#39;. | 
 **filterEmergencyAddressId** | **string** | Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string &#39;null&#39;. | 
 **filterCustomerReference** | **string** | Filter numbers via the customer_reference set. | 
 **filterNumberTypeEq** | **string** | Filter phone numbers by phone number type. | 
 **filterSource** | **string** | Filter phone numbers by their source. Use &#39;ported&#39; for numbers ported from other carriers, or &#39;purchased&#39; for numbers bought directly from Telnyx. | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersResponse**](ListPhoneNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumbersWithVoiceSettings

> ListPhoneNumbersWithVoiceSettingsResponse ListPhoneNumbersWithVoiceSettings(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterPhoneNumber(filterPhoneNumber).FilterConnectionNameContains(filterConnectionNameContains).FilterCustomerReference(filterCustomerReference).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).Sort(sort).Execute()

List phone numbers with voice settings

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
	filterPhoneNumber := "filterPhoneNumber_example" // string | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. (optional)
	filterConnectionNameContains := "test" // string | Filter contains connection name. Requires at least three characters. (optional)
	filterCustomerReference := "filterCustomerReference_example" // string | Filter numbers via the customer_reference set. (optional)
	filterVoiceUsagePaymentMethod := "channel" // string | Filter by usage_payment_method. (optional)
	sort := "connection_name" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.ListPhoneNumbersWithVoiceSettings(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterPhoneNumber(filterPhoneNumber).FilterConnectionNameContains(filterConnectionNameContains).FilterCustomerReference(filterCustomerReference).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.ListPhoneNumbersWithVoiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPhoneNumbersWithVoiceSettings`: ListPhoneNumbersWithVoiceSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.ListPhoneNumbersWithVoiceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhoneNumbersWithVoiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterPhoneNumber** | **string** | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterConnectionNameContains** | **string** | Filter contains connection name. Requires at least three characters. | 
 **filterCustomerReference** | **string** | Filter numbers via the customer_reference set. | 
 **filterVoiceUsagePaymentMethod** | **string** | Filter by usage_payment_method. | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersWithVoiceSettingsResponse**](ListPhoneNumbersWithVoiceSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PhoneNumberBundleStatusChange

> PhoneNumberBundleStatusChange PhoneNumberBundleStatusChange(ctx, id).PhoneNumberBundleStatusChangeRequest(phoneNumberBundleStatusChangeRequest).Execute()

Change the bundle status for a phone number (set to being in a bundle or remove from a bundle)

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
	id := "1293384261075731499" // string | Identifies the resource.
	phoneNumberBundleStatusChangeRequest := *openapiclient.NewPhoneNumberBundleStatusChangeRequest("BundleId_example") // PhoneNumberBundleStatusChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.PhoneNumberBundleStatusChange(context.Background(), id).PhoneNumberBundleStatusChangeRequest(phoneNumberBundleStatusChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.PhoneNumberBundleStatusChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PhoneNumberBundleStatusChange`: PhoneNumberBundleStatusChange
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.PhoneNumberBundleStatusChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPhoneNumberBundleStatusChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumberBundleStatusChangeRequest** | [**PhoneNumberBundleStatusChangeRequest**](PhoneNumberBundleStatusChangeRequest.md) |  | 

### Return type

[**PhoneNumberBundleStatusChange**](PhoneNumberBundleStatusChange.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePhoneNumber

> PhoneNumberResponse RetrievePhoneNumber(ctx, id).Execute()

Retrieve a phone number

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
	id := "1293384261075731499" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.RetrievePhoneNumber(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.RetrievePhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievePhoneNumber`: PhoneNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.RetrievePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneNumberResponse**](PhoneNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlimListPhoneNumbers

> ListPhoneNumbersResponse1 SlimListPhoneNumbers(ctx).PageNumber(pageNumber).PageSize(pageSize).IncludeConnection(includeConnection).IncludeTags(includeTags).FilterTag(filterTag).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterCountryIsoAlpha2(filterCountryIsoAlpha2).FilterConnectionId(filterConnectionId).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceConnectionNameStartsWith(filterVoiceConnectionNameStartsWith).FilterVoiceConnectionNameEndsWith(filterVoiceConnectionNameEndsWith).FilterVoiceConnectionName(filterVoiceConnectionName).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).FilterNumberTypeEq(filterNumberTypeEq).FilterSource(filterSource).Sort(sort).Execute()

Slim List phone numbers



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
	includeConnection := true // bool | Include the connection associated with the phone number. (optional) (default to false)
	includeTags := true // bool | Include the tags associated with the phone number. (optional) (default to false)
	filterTag := "filterTag_example" // string | Filter by phone number tags. (This requires the include_tags param) (optional)
	filterPhoneNumber := "filterPhoneNumber_example" // string | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. (optional)
	filterStatus := "active" // string | Filter by phone number status. (optional)
	filterCountryIsoAlpha2 := openapiclient.ListPhoneNumbers_filter_country_iso_alpha2__parameter{ArrayOfString: new([]string)} // ListPhoneNumbersFilterCountryIsoAlpha2Parameter | Filter by phone number country ISO alpha-2 code. Can be a single value or an array of values. (optional)
	filterConnectionId := "1521916448077776306" // string | Filter by connection_id. (optional)
	filterVoiceConnectionNameContains := "test" // string | Filter contains connection name. Requires at least three characters and the include_connection param. (optional)
	filterVoiceConnectionNameStartsWith := "test" // string | Filter starts with connection name. Requires at least three characters and the include_connection param. (optional)
	filterVoiceConnectionNameEndsWith := "test" // string | Filter ends with connection name. Requires at least three characters and the include_connection param. (optional)
	filterVoiceConnectionName := "test" // string | Filter by connection name , requires the include_connection param and the include_connection param. (optional)
	filterVoiceUsagePaymentMethod := "channel" // string | Filter by usage_payment_method. (optional)
	filterBillingGroupId := "62e4bf2e-c278-4282-b524-488d9c9c43b2" // string | Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string 'null'. (optional)
	filterEmergencyAddressId := "9102160989215728032" // string | Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string 'null'. (optional)
	filterCustomerReference := "filterCustomerReference_example" // string | Filter numbers via the customer_reference set. (optional)
	filterNumberTypeEq := "filterNumberTypeEq_example" // string | Filter phone numbers by phone number type. (optional)
	filterSource := "filterSource_example" // string | Filter phone numbers by their source. Use 'ported' for numbers ported from other carriers, or 'purchased' for numbers bought directly from Telnyx. (optional)
	sort := "connection_name" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.SlimListPhoneNumbers(context.Background()).PageNumber(pageNumber).PageSize(pageSize).IncludeConnection(includeConnection).IncludeTags(includeTags).FilterTag(filterTag).FilterPhoneNumber(filterPhoneNumber).FilterStatus(filterStatus).FilterCountryIsoAlpha2(filterCountryIsoAlpha2).FilterConnectionId(filterConnectionId).FilterVoiceConnectionNameContains(filterVoiceConnectionNameContains).FilterVoiceConnectionNameStartsWith(filterVoiceConnectionNameStartsWith).FilterVoiceConnectionNameEndsWith(filterVoiceConnectionNameEndsWith).FilterVoiceConnectionName(filterVoiceConnectionName).FilterVoiceUsagePaymentMethod(filterVoiceUsagePaymentMethod).FilterBillingGroupId(filterBillingGroupId).FilterEmergencyAddressId(filterEmergencyAddressId).FilterCustomerReference(filterCustomerReference).FilterNumberTypeEq(filterNumberTypeEq).FilterSource(filterSource).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.SlimListPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlimListPhoneNumbers`: ListPhoneNumbersResponse1
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.SlimListPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlimListPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **includeConnection** | **bool** | Include the connection associated with the phone number. | [default to false]
 **includeTags** | **bool** | Include the tags associated with the phone number. | [default to false]
 **filterTag** | **string** | Filter by phone number tags. (This requires the include_tags param) | 
 **filterPhoneNumber** | **string** | Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterStatus** | **string** | Filter by phone number status. | 
 **filterCountryIsoAlpha2** | [**ListPhoneNumbersFilterCountryIsoAlpha2Parameter**](ListPhoneNumbersFilterCountryIsoAlpha2Parameter.md) | Filter by phone number country ISO alpha-2 code. Can be a single value or an array of values. | 
 **filterConnectionId** | **string** | Filter by connection_id. | 
 **filterVoiceConnectionNameContains** | **string** | Filter contains connection name. Requires at least three characters and the include_connection param. | 
 **filterVoiceConnectionNameStartsWith** | **string** | Filter starts with connection name. Requires at least three characters and the include_connection param. | 
 **filterVoiceConnectionNameEndsWith** | **string** | Filter ends with connection name. Requires at least three characters and the include_connection param. | 
 **filterVoiceConnectionName** | **string** | Filter by connection name , requires the include_connection param and the include_connection param. | 
 **filterVoiceUsagePaymentMethod** | **string** | Filter by usage_payment_method. | 
 **filterBillingGroupId** | **string** | Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string &#39;null&#39;. | 
 **filterEmergencyAddressId** | **string** | Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string &#39;null&#39;. | 
 **filterCustomerReference** | **string** | Filter numbers via the customer_reference set. | 
 **filterNumberTypeEq** | **string** | Filter phone numbers by phone number type. | 
 **filterSource** | **string** | Filter phone numbers by their source. Use &#39;ported&#39; for numbers ported from other carriers, or &#39;purchased&#39; for numbers bought directly from Telnyx. | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersResponse1**](ListPhoneNumbersResponse1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumber

> PhoneNumberResponse UpdatePhoneNumber(ctx, id).UpdatePhoneNumberRequest(updatePhoneNumberRequest).Execute()

Update a phone number

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
	id := "1293384261075731499" // string | Identifies the resource.
	updatePhoneNumberRequest := *openapiclient.NewUpdatePhoneNumberRequest() // UpdatePhoneNumberRequest | Updated settings for the phone number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.UpdatePhoneNumber(context.Background(), id).UpdatePhoneNumberRequest(updatePhoneNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.UpdatePhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePhoneNumber`: PhoneNumberResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.UpdatePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePhoneNumberRequest** | [**UpdatePhoneNumberRequest**](UpdatePhoneNumberRequest.md) | Updated settings for the phone number. | 

### Return type

[**PhoneNumberResponse**](PhoneNumberResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumberVoiceSettings

> RetrievePhoneNumberVoiceResponse UpdatePhoneNumberVoiceSettings(ctx, id).UpdatePhoneNumberVoiceSettingsRequest(updatePhoneNumberVoiceSettingsRequest).Execute()

Update a phone number with voice settings

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
	id := "1293384261075731499" // string | Identifies the resource.
	updatePhoneNumberVoiceSettingsRequest := *openapiclient.NewUpdatePhoneNumberVoiceSettingsRequest() // UpdatePhoneNumberVoiceSettingsRequest | Updated voice settings for the phone number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberConfigurationsAPI.UpdatePhoneNumberVoiceSettings(context.Background(), id).UpdatePhoneNumberVoiceSettingsRequest(updatePhoneNumberVoiceSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberConfigurationsAPI.UpdatePhoneNumberVoiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePhoneNumberVoiceSettings`: RetrievePhoneNumberVoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberConfigurationsAPI.UpdatePhoneNumberVoiceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneNumberVoiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePhoneNumberVoiceSettingsRequest** | [**UpdatePhoneNumberVoiceSettingsRequest**](UpdatePhoneNumberVoiceSettingsRequest.md) | Updated voice settings for the phone number. | 

### Return type

[**RetrievePhoneNumberVoiceResponse**](RetrievePhoneNumberVoiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

