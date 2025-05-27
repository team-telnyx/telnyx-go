# \MessagingProfilesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingProfile**](MessagingProfilesAPI.md#CreateMessagingProfile) | **Post** /messaging_profiles | Create a messaging profile
[**DeleteMessagingProfile**](MessagingProfilesAPI.md#DeleteMessagingProfile) | **Delete** /messaging_profiles/{id} | Delete a messaging profile
[**GetMessagingProfileMetrics**](MessagingProfilesAPI.md#GetMessagingProfileMetrics) | **Get** /messaging_profiles/{id}/metrics | Retrieve messaging profile metrics
[**ListMessagingProfiles**](MessagingProfilesAPI.md#ListMessagingProfiles) | **Get** /messaging_profiles | List messaging profiles
[**ListProfileMetrics**](MessagingProfilesAPI.md#ListProfileMetrics) | **Get** /messaging_profile_metrics | List messaging profile metrics
[**ListProfilePhoneNumbers**](MessagingProfilesAPI.md#ListProfilePhoneNumbers) | **Get** /messaging_profiles/{id}/phone_numbers | List phone numbers associated with a messaging profile
[**ListProfileShortCodes**](MessagingProfilesAPI.md#ListProfileShortCodes) | **Get** /messaging_profiles/{id}/short_codes | List short codes associated with a messaging profile
[**RetrieveMessagingProfile**](MessagingProfilesAPI.md#RetrieveMessagingProfile) | **Get** /messaging_profiles/{id} | Retrieve a messaging profile
[**UpdateMessagingProfile**](MessagingProfilesAPI.md#UpdateMessagingProfile) | **Patch** /messaging_profiles/{id} | Update a messaging profile



## CreateMessagingProfile

> MessagingProfileResponse CreateMessagingProfile(ctx).CreateMessagingProfileRequest(createMessagingProfileRequest).Execute()

Create a messaging profile

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
	createMessagingProfileRequest := *openapiclient.NewCreateMessagingProfileRequest("Name_example", []string{"WhitelistedDestinations_example"}) // CreateMessagingProfileRequest | New Messaging Profile object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.CreateMessagingProfile(context.Background()).CreateMessagingProfileRequest(createMessagingProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.CreateMessagingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessagingProfile`: MessagingProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.CreateMessagingProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessagingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMessagingProfileRequest** | [**CreateMessagingProfileRequest**](CreateMessagingProfileRequest.md) | New Messaging Profile object | 

### Return type

[**MessagingProfileResponse**](MessagingProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessagingProfile

> MessagingProfileResponse DeleteMessagingProfile(ctx, id).Execute()

Delete a messaging profile

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.DeleteMessagingProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.DeleteMessagingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessagingProfile`: MessagingProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.DeleteMessagingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the messaging profile to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessagingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagingProfileResponse**](MessagingProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagingProfileMetrics

> RetrieveMessagingProfileMetricsResponse GetMessagingProfileMetrics(ctx, id).TimeFrame(timeFrame).Execute()

Retrieve messaging profile metrics

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile to retrieve
	timeFrame := "timeFrame_example" // string | The timeframe for which you'd like to retrieve metrics. (optional) (default to "24h")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.GetMessagingProfileMetrics(context.Background(), id).TimeFrame(timeFrame).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.GetMessagingProfileMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagingProfileMetrics`: RetrieveMessagingProfileMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.GetMessagingProfileMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the messaging profile to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagingProfileMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeFrame** | **string** | The timeframe for which you&#39;d like to retrieve metrics. | [default to &quot;24h&quot;]

### Return type

[**RetrieveMessagingProfileMetricsResponse**](RetrieveMessagingProfileMetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingProfiles

> ListMessagingProfilesResponse ListMessagingProfiles(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).Execute()

List messaging profiles

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
	filterName := "filterName_example" // string | Filter by name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.ListMessagingProfiles(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.ListMessagingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessagingProfiles`: ListMessagingProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.ListMessagingProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagingProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterName** | **string** | Filter by name | 

### Return type

[**ListMessagingProfilesResponse**](ListMessagingProfilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileMetrics

> ListMessagingProfileMetricsResponse ListProfileMetrics(ctx).PageNumber(pageNumber).PageSize(pageSize).Id(id).TimeFrame(timeFrame).Execute()

List messaging profile metrics

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile(s) to retrieve (optional)
	timeFrame := "timeFrame_example" // string | The timeframe for which you'd like to retrieve metrics. (optional) (default to "24h")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.ListProfileMetrics(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Id(id).TimeFrame(timeFrame).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.ListProfileMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfileMetrics`: ListMessagingProfileMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.ListProfileMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProfileMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **id** | **string** | The id of the messaging profile(s) to retrieve | 
 **timeFrame** | **string** | The timeframe for which you&#39;d like to retrieve metrics. | [default to &quot;24h&quot;]

### Return type

[**ListMessagingProfileMetricsResponse**](ListMessagingProfileMetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfilePhoneNumbers

> ListMessagingProfilePhoneNumbersResponse ListProfilePhoneNumbers(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

List phone numbers associated with a messaging profile

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile to retrieve
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.ListProfilePhoneNumbers(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.ListProfilePhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfilePhoneNumbers`: ListMessagingProfilePhoneNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.ListProfilePhoneNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the messaging profile to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProfilePhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListMessagingProfilePhoneNumbersResponse**](ListMessagingProfilePhoneNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileShortCodes

> ListMessagingProfileShortCodesResponse ListProfileShortCodes(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

List short codes associated with a messaging profile

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile to retrieve
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.ListProfileShortCodes(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.ListProfileShortCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfileShortCodes`: ListMessagingProfileShortCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.ListProfileShortCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the messaging profile to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProfileShortCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListMessagingProfileShortCodesResponse**](ListMessagingProfileShortCodesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMessagingProfile

> MessagingProfileResponse RetrieveMessagingProfile(ctx, id).Execute()

Retrieve a messaging profile

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.RetrieveMessagingProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.RetrieveMessagingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveMessagingProfile`: MessagingProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.RetrieveMessagingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the messaging profile to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMessagingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagingProfileResponse**](MessagingProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessagingProfile

> MessagingProfileResponse UpdateMessagingProfile(ctx, id).UpdateMessagingProfileRequest(updateMessagingProfileRequest).Execute()

Update a messaging profile

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the messaging profile to retrieve
	updateMessagingProfileRequest := *openapiclient.NewUpdateMessagingProfileRequest() // UpdateMessagingProfileRequest | New Messaging Profile object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingProfilesAPI.UpdateMessagingProfile(context.Background(), id).UpdateMessagingProfileRequest(updateMessagingProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingProfilesAPI.UpdateMessagingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessagingProfile`: MessagingProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingProfilesAPI.UpdateMessagingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the messaging profile to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessagingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMessagingProfileRequest** | [**UpdateMessagingProfileRequest**](UpdateMessagingProfileRequest.md) | New Messaging Profile object | 

### Return type

[**MessagingProfileResponse**](MessagingProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

