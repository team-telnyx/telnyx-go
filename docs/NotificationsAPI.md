# \NotificationsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationChannels**](NotificationsAPI.md#CreateNotificationChannels) | **Post** /notification_channels | Create a notification channel
[**CreateNotificationProfile**](NotificationsAPI.md#CreateNotificationProfile) | **Post** /notification_profiles | Create a notification profile
[**CreateNotificationSetting**](NotificationsAPI.md#CreateNotificationSetting) | **Post** /notification_settings | Add a Notification Setting
[**DeleteNotificationChannel**](NotificationsAPI.md#DeleteNotificationChannel) | **Delete** /notification_channels/{id} | Delete a notification channel
[**DeleteNotificationProfile**](NotificationsAPI.md#DeleteNotificationProfile) | **Delete** /notification_profiles/{id} | Delete a notification profile
[**DeleteNotificationSetting**](NotificationsAPI.md#DeleteNotificationSetting) | **Delete** /notification_settings/{id} | Delete a notification setting
[**FindNotificationsEvents**](NotificationsAPI.md#FindNotificationsEvents) | **Get** /notification_events | List all Notifications Events
[**FindNotificationsEventsConditions**](NotificationsAPI.md#FindNotificationsEventsConditions) | **Get** /notification_event_conditions | List all Notifications Events Conditions
[**FindNotificationsProfiles**](NotificationsAPI.md#FindNotificationsProfiles) | **Get** /notification_profiles | List all Notifications Profiles
[**GetNotificationChannel**](NotificationsAPI.md#GetNotificationChannel) | **Get** /notification_channels/{id} | Get a notification channel
[**GetNotificationProfile**](NotificationsAPI.md#GetNotificationProfile) | **Get** /notification_profiles/{id} | Get a notification profile
[**GetNotificationSetting**](NotificationsAPI.md#GetNotificationSetting) | **Get** /notification_settings/{id} | Get a notification setting
[**ListNotificationChannels**](NotificationsAPI.md#ListNotificationChannels) | **Get** /notification_channels | List notification channels
[**ListNotificationSettings**](NotificationsAPI.md#ListNotificationSettings) | **Get** /notification_settings | List notification settings
[**UpdateNotificationChannel**](NotificationsAPI.md#UpdateNotificationChannel) | **Patch** /notification_channels/{id} | Update a notification channel
[**UpdateNotificationProfile**](NotificationsAPI.md#UpdateNotificationProfile) | **Patch** /notification_profiles/{id} | Update a notification profile



## CreateNotificationChannels

> CreateNotificationChannels200Response CreateNotificationChannels(ctx).NotificationChannel(notificationChannel).Execute()

Create a notification channel



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
	notificationChannel := *openapiclient.NewNotificationChannel() // NotificationChannel | Add a Notification Channel (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateNotificationChannels(context.Background()).NotificationChannel(notificationChannel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationChannels`: CreateNotificationChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateNotificationChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationChannel** | [**NotificationChannel**](NotificationChannel.md) | Add a Notification Channel | 

### Return type

[**CreateNotificationChannels200Response**](CreateNotificationChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationProfile

> CreateNotificationProfile200Response CreateNotificationProfile(ctx).NotificationProfile(notificationProfile).Execute()

Create a notification profile



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
	notificationProfile := *openapiclient.NewNotificationProfile() // NotificationProfile | Add a Notification Profile (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateNotificationProfile(context.Background()).NotificationProfile(notificationProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateNotificationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationProfile`: CreateNotificationProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateNotificationProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationProfile** | [**NotificationProfile**](NotificationProfile.md) | Add a Notification Profile | 

### Return type

[**CreateNotificationProfile200Response**](CreateNotificationProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationSetting

> CreateNotificationSetting200Response CreateNotificationSetting(ctx).NotificationSetting(notificationSetting).Execute()

Add a Notification Setting



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
	notificationSetting := *openapiclient.NewNotificationSetting() // NotificationSetting |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateNotificationSetting(context.Background()).NotificationSetting(notificationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationSetting`: CreateNotificationSetting200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateNotificationSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationSetting** | [**NotificationSetting**](NotificationSetting.md) |  | 

### Return type

[**CreateNotificationSetting200Response**](CreateNotificationSetting200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationChannel

> CreateNotificationChannels200Response DeleteNotificationChannel(ctx, id).Execute()

Delete a notification channel



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteNotificationChannel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationChannel`: CreateNotificationChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationChannels200Response**](CreateNotificationChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationProfile

> CreateNotificationProfile200Response DeleteNotificationProfile(ctx, id).Execute()

Delete a notification profile



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteNotificationProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationProfile`: CreateNotificationProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteNotificationProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationProfile200Response**](CreateNotificationProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationSetting

> CreateNotificationSetting200Response DeleteNotificationSetting(ctx, id).Execute()

Delete a notification setting



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteNotificationSetting(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationSetting`: CreateNotificationSetting200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationSetting200Response**](CreateNotificationSetting200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindNotificationsEvents

> FindNotificationsEvents200Response FindNotificationsEvents(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Notifications Events



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
	resp, r, err := apiClient.NotificationsAPI.FindNotificationsEvents(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.FindNotificationsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindNotificationsEvents`: FindNotificationsEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.FindNotificationsEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindNotificationsEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**FindNotificationsEvents200Response**](FindNotificationsEvents200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindNotificationsEventsConditions

> FindNotificationsEventsConditions200Response FindNotificationsEventsConditions(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterAssociatedRecordTypeEq(filterAssociatedRecordTypeEq).Execute()

List all Notifications Events Conditions



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
	filterAssociatedRecordTypeEq := "phone_number" // string | Filter by the associated record type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.FindNotificationsEventsConditions(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterAssociatedRecordTypeEq(filterAssociatedRecordTypeEq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.FindNotificationsEventsConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindNotificationsEventsConditions`: FindNotificationsEventsConditions200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.FindNotificationsEventsConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindNotificationsEventsConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterAssociatedRecordTypeEq** | **string** | Filter by the associated record type | 

### Return type

[**FindNotificationsEventsConditions200Response**](FindNotificationsEventsConditions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindNotificationsProfiles

> FindNotificationsProfiles200Response FindNotificationsProfiles(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Notifications Profiles



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
	resp, r, err := apiClient.NotificationsAPI.FindNotificationsProfiles(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.FindNotificationsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindNotificationsProfiles`: FindNotificationsProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.FindNotificationsProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindNotificationsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**FindNotificationsProfiles200Response**](FindNotificationsProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationChannel

> CreateNotificationChannels200Response GetNotificationChannel(ctx, id).Execute()

Get a notification channel



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationChannel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationChannel`: CreateNotificationChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationChannels200Response**](CreateNotificationChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationProfile

> CreateNotificationProfile200Response GetNotificationProfile(ctx, id).Execute()

Get a notification profile



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationProfile`: CreateNotificationProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationProfile200Response**](CreateNotificationProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSetting

> CreateNotificationSetting200Response GetNotificationSetting(ctx, id).Execute()

Get a notification setting



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationSetting(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSetting`: CreateNotificationSetting200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationSetting200Response**](CreateNotificationSetting200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationChannels

> ListNotificationChannels200Response ListNotificationChannels(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterChannelTypeIdEq(filterChannelTypeIdEq).Execute()

List notification channels



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
	filterChannelTypeIdEq := "webhook" // string | Filter by the id of a channel type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListNotificationChannels(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterChannelTypeIdEq(filterChannelTypeIdEq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationChannels`: ListNotificationChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListNotificationChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterChannelTypeIdEq** | **string** | Filter by the id of a channel type | 

### Return type

[**ListNotificationChannels200Response**](ListNotificationChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationSettings

> ListNotificationSettings200Response ListNotificationSettings(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterNotificationProfileIdEq(filterNotificationProfileIdEq).FilterNotificationChannelEq(filterNotificationChannelEq).FilterNotificationEventConditionIdEq(filterNotificationEventConditionIdEq).FilterAssociatedRecordTypeEq(filterAssociatedRecordTypeEq).Status(status).Execute()

List notification settings



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
	filterNotificationProfileIdEq := "12455643-3cf1-4683-ad23-1cd32f7d5e0a" // string | Filter by the id of a notification profile (optional)
	filterNotificationChannelEq := "12455643-3cf1-4683-ad23-1cd32f7d5e0a" // string | Filter by the id of a notification channel (optional)
	filterNotificationEventConditionIdEq := "12455643-3cf1-4683-ad23-1cd32f7d5e0a" // string | Filter by the id of a notification channel (optional)
	filterAssociatedRecordTypeEq := "phone_number" // string | Filter by the associated record type (optional)
	status := "no-answer" // string | Filters calls by status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListNotificationSettings(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterNotificationProfileIdEq(filterNotificationProfileIdEq).FilterNotificationChannelEq(filterNotificationChannelEq).FilterNotificationEventConditionIdEq(filterNotificationEventConditionIdEq).FilterAssociatedRecordTypeEq(filterAssociatedRecordTypeEq).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationSettings`: ListNotificationSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListNotificationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterNotificationProfileIdEq** | **string** | Filter by the id of a notification profile | 
 **filterNotificationChannelEq** | **string** | Filter by the id of a notification channel | 
 **filterNotificationEventConditionIdEq** | **string** | Filter by the id of a notification channel | 
 **filterAssociatedRecordTypeEq** | **string** | Filter by the associated record type | 
 **status** | **string** | Filters calls by status. | 

### Return type

[**ListNotificationSettings200Response**](ListNotificationSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationChannel

> CreateNotificationChannels200Response UpdateNotificationChannel(ctx, id).NotificationChannel(notificationChannel).Execute()

Update a notification channel



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.
	notificationChannel := *openapiclient.NewNotificationChannel() // NotificationChannel | Update notification channel object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.UpdateNotificationChannel(context.Background(), id).NotificationChannel(notificationChannel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationChannel`: CreateNotificationChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationChannel** | [**NotificationChannel**](NotificationChannel.md) | Update notification channel object | 

### Return type

[**CreateNotificationChannels200Response**](CreateNotificationChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationProfile

> CreateNotificationProfile200Response UpdateNotificationProfile(ctx, id).NotificationProfile(notificationProfile).Execute()

Update a notification profile



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the resource.
	notificationProfile := *openapiclient.NewNotificationProfile() // NotificationProfile | Update notification profile object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.UpdateNotificationProfile(context.Background(), id).NotificationProfile(notificationProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateNotificationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationProfile`: CreateNotificationProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateNotificationProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationProfile** | [**NotificationProfile**](NotificationProfile.md) | Update notification profile object | 

### Return type

[**CreateNotificationProfile200Response**](CreateNotificationProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

