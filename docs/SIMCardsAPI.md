# \SIMCardsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSimCard**](SIMCardsAPI.md#DeleteSimCard) | **Delete** /sim_cards/{id} | Deletes a SIM card
[**DeleteSimCardDataUsageNotifications**](SIMCardsAPI.md#DeleteSimCardDataUsageNotifications) | **Delete** /sim_card_data_usage_notifications/{id} | Delete SIM card data usage notifications
[**DisableSimCard**](SIMCardsAPI.md#DisableSimCard) | **Post** /sim_cards/{id}/actions/disable | Request a SIM card disable
[**EnableSimCard**](SIMCardsAPI.md#EnableSimCard) | **Post** /sim_cards/{id}/actions/enable | Request a SIM card enable
[**GetSimCard**](SIMCardsAPI.md#GetSimCard) | **Get** /sim_cards/{id} | Get SIM card
[**GetSimCardActivationCode**](SIMCardsAPI.md#GetSimCardActivationCode) | **Get** /sim_cards/{id}/activation_code | Get activation code for an eSIM
[**GetSimCardDataUsageNotification**](SIMCardsAPI.md#GetSimCardDataUsageNotification) | **Get** /sim_card_data_usage_notifications/{id} | Get a single SIM card data usage notification
[**GetSimCardDeviceDetails**](SIMCardsAPI.md#GetSimCardDeviceDetails) | **Get** /sim_cards/{id}/device_details | Get SIM card device details
[**GetSimCardPublicIp**](SIMCardsAPI.md#GetSimCardPublicIp) | **Get** /sim_cards/{id}/public_ip | Get SIM card public IP definition
[**GetSimCards**](SIMCardsAPI.md#GetSimCards) | **Get** /sim_cards | Get all SIM cards
[**GetWirelessConnectivityLogs**](SIMCardsAPI.md#GetWirelessConnectivityLogs) | **Get** /sim_cards/{id}/wireless_connectivity_logs | List wireless connectivity logs
[**ListDataUsageNotifications**](SIMCardsAPI.md#ListDataUsageNotifications) | **Get** /sim_card_data_usage_notifications | List SIM card data usage notifications
[**PatchSimCardDataUsageNotification**](SIMCardsAPI.md#PatchSimCardDataUsageNotification) | **Patch** /sim_card_data_usage_notifications/{id} | Updates information for a SIM Card Data Usage Notification
[**PostSimCardDataUsageNotification**](SIMCardsAPI.md#PostSimCardDataUsageNotification) | **Post** /sim_card_data_usage_notifications | Create a new SIM card data usage notification
[**PurchaseESim**](SIMCardsAPI.md#PurchaseESim) | **Post** /actions/purchase/esims | Purchase eSIMs
[**RegisterSimCards**](SIMCardsAPI.md#RegisterSimCards) | **Post** /actions/register/sim_cards | Register SIM cards
[**RemoveSimCardPublicIp**](SIMCardsAPI.md#RemoveSimCardPublicIp) | **Post** /sim_cards/{id}/actions/remove_public_ip | Request removing a SIM card public IP
[**SetPublicIPsBulk**](SIMCardsAPI.md#SetPublicIPsBulk) | **Post** /sim_cards/actions/bulk_set_public_ips | Request bulk setting SIM card public IPs.
[**SetSimCardPublicIp**](SIMCardsAPI.md#SetSimCardPublicIp) | **Post** /sim_cards/{id}/actions/set_public_ip | Request setting a SIM card public IP
[**SetSimCardStandby**](SIMCardsAPI.md#SetSimCardStandby) | **Post** /sim_cards/{id}/actions/set_standby | Request setting a SIM card to standby
[**UpdateSimCard**](SIMCardsAPI.md#UpdateSimCard) | **Patch** /sim_cards/{id} | Update a SIM card
[**ValidateRegistrationCodes**](SIMCardsAPI.md#ValidateRegistrationCodes) | **Post** /sim_cards/actions/validate_registration_codes | Validate SIM cards registration codes



## DeleteSimCard

> GetSimCard200Response DeleteSimCard(ctx, id).ReportLost(reportLost).Execute()

Deletes a SIM card



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.
	reportLost := true // bool | Enables deletion of disabled eSIMs that can't be uninstalled from a device. This is irreversible and the eSIM cannot be re-registered. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.DeleteSimCard(context.Background(), id).ReportLost(reportLost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.DeleteSimCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSimCard`: GetSimCard200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.DeleteSimCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSimCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportLost** | **bool** | Enables deletion of disabled eSIMs that can&#39;t be uninstalled from a device. This is irreversible and the eSIM cannot be re-registered. | [default to false]

### Return type

[**GetSimCard200Response**](GetSimCard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSimCardDataUsageNotifications

> PostSimCardDataUsageNotification201Response DeleteSimCardDataUsageNotifications(ctx, id).Execute()

Delete SIM card data usage notifications



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.DeleteSimCardDataUsageNotifications(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.DeleteSimCardDataUsageNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSimCardDataUsageNotifications`: PostSimCardDataUsageNotification201Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.DeleteSimCardDataUsageNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSimCardDataUsageNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostSimCardDataUsageNotification201Response**](PostSimCardDataUsageNotification201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableSimCard

> GetSimCardAction200Response DisableSimCard(ctx, id).Execute()

Request a SIM card disable



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.DisableSimCard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.DisableSimCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableSimCard`: GetSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.DisableSimCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableSimCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardAction200Response**](GetSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableSimCard

> GetSimCardAction200Response EnableSimCard(ctx, id).Execute()

Request a SIM card enable



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.EnableSimCard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.EnableSimCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableSimCard`: GetSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.EnableSimCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSimCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardAction200Response**](GetSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCard

> GetSimCard200Response GetSimCard(ctx, id).IncludeSimCardGroup(includeSimCardGroup).Execute()

Get SIM card



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.
	includeSimCardGroup := true // bool | It includes the associated SIM card group object in the response when present. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetSimCard(context.Background(), id).IncludeSimCardGroup(includeSimCardGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetSimCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCard`: GetSimCard200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetSimCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSimCardGroup** | **bool** | It includes the associated SIM card group object in the response when present. | [default to false]

### Return type

[**GetSimCard200Response**](GetSimCard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardActivationCode

> GetSimCardActivationCode200Response GetSimCardActivationCode(ctx, id).Execute()

Get activation code for an eSIM



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetSimCardActivationCode(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetSimCardActivationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardActivationCode`: GetSimCardActivationCode200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetSimCardActivationCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardActivationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardActivationCode200Response**](GetSimCardActivationCode200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardDataUsageNotification

> PostSimCardDataUsageNotification201Response GetSimCardDataUsageNotification(ctx, id).Execute()

Get a single SIM card data usage notification



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetSimCardDataUsageNotification(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetSimCardDataUsageNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardDataUsageNotification`: PostSimCardDataUsageNotification201Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetSimCardDataUsageNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardDataUsageNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostSimCardDataUsageNotification201Response**](PostSimCardDataUsageNotification201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardDeviceDetails

> GetSimCardDeviceDetails200Response GetSimCardDeviceDetails(ctx, id).Execute()

Get SIM card device details



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetSimCardDeviceDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetSimCardDeviceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardDeviceDetails`: GetSimCardDeviceDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetSimCardDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardDeviceDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardDeviceDetails200Response**](GetSimCardDeviceDetails200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardPublicIp

> GetSimCardPublicIp200Response GetSimCardPublicIp(ctx, id).Execute()

Get SIM card public IP definition



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetSimCardPublicIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetSimCardPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardPublicIp`: GetSimCardPublicIp200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetSimCardPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardPublicIp200Response**](GetSimCardPublicIp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCards

> GetSimCards200Response GetSimCards(ctx).PageNumber(pageNumber).PageSize(pageSize).IncludeSimCardGroup(includeSimCardGroup).FilterSimCardGroupId(filterSimCardGroupId).FilterTags(filterTags).FilterIccid(filterIccid).FilterStatus(filterStatus).Sort(sort).Execute()

Get all SIM cards



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
	includeSimCardGroup := true // bool | It includes the associated SIM card group object in the response when present. (optional) (default to false)
	filterSimCardGroupId := "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9" // string | A valid SIM card group ID. (optional)
	filterTags := []string{"Inner_example"} // []string | A list of SIM card tags to filter on.<br/><br/>  If the SIM card contains <b><i>all</i></b> of the given <code>tags</code> they will be found.<br/><br/> For example, if the SIM cards have the following tags: <ul>   <li><code>['customers', 'staff', 'test']</code>   <li><code>['test']</code></li>   <li><code>['customers']</code></li> </ul> Searching for <code>['customers', 'test']</code> returns only the first because it's the only one with both tags.<br/> Searching for <code>test</code> returns the first two SIMs, because both of them have such tag.<br/> Searching for <code>customers</code> returns the first and last SIMs.<br/>  (optional)
	filterIccid := "89310410106543789301" // string | A search string to partially match for the SIM card's ICCID. (optional)
	filterStatus := []string{"FilterStatus_example"} // []string | Filter by a SIM card's status. (optional)
	sort := "-current_billing_period_consumed_data.amount" // string | Sorts SIM cards by the given field. Defaults to ascending order unless field is prefixed with a minus sign. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetSimCards(context.Background()).PageNumber(pageNumber).PageSize(pageSize).IncludeSimCardGroup(includeSimCardGroup).FilterSimCardGroupId(filterSimCardGroupId).FilterTags(filterTags).FilterIccid(filterIccid).FilterStatus(filterStatus).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetSimCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCards`: GetSimCards200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetSimCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **includeSimCardGroup** | **bool** | It includes the associated SIM card group object in the response when present. | [default to false]
 **filterSimCardGroupId** | **string** | A valid SIM card group ID. | 
 **filterTags** | **[]string** | A list of SIM card tags to filter on.&lt;br/&gt;&lt;br/&gt;  If the SIM card contains &lt;b&gt;&lt;i&gt;all&lt;/i&gt;&lt;/b&gt; of the given &lt;code&gt;tags&lt;/code&gt; they will be found.&lt;br/&gt;&lt;br/&gt; For example, if the SIM cards have the following tags: &lt;ul&gt;   &lt;li&gt;&lt;code&gt;[&#39;customers&#39;, &#39;staff&#39;, &#39;test&#39;]&lt;/code&gt;   &lt;li&gt;&lt;code&gt;[&#39;test&#39;]&lt;/code&gt;&lt;/li&gt;   &lt;li&gt;&lt;code&gt;[&#39;customers&#39;]&lt;/code&gt;&lt;/li&gt; &lt;/ul&gt; Searching for &lt;code&gt;[&#39;customers&#39;, &#39;test&#39;]&lt;/code&gt; returns only the first because it&#39;s the only one with both tags.&lt;br/&gt; Searching for &lt;code&gt;test&lt;/code&gt; returns the first two SIMs, because both of them have such tag.&lt;br/&gt; Searching for &lt;code&gt;customers&lt;/code&gt; returns the first and last SIMs.&lt;br/&gt;  | 
 **filterIccid** | **string** | A search string to partially match for the SIM card&#39;s ICCID. | 
 **filterStatus** | **[]string** | Filter by a SIM card&#39;s status. | 
 **sort** | **string** | Sorts SIM cards by the given field. Defaults to ascending order unless field is prefixed with a minus sign. | 

### Return type

[**GetSimCards200Response**](GetSimCards200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessConnectivityLogs

> GetWirelessConnectivityLogs200Response GetWirelessConnectivityLogs(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

List wireless connectivity logs



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.GetWirelessConnectivityLogs(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.GetWirelessConnectivityLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWirelessConnectivityLogs`: GetWirelessConnectivityLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.GetWirelessConnectivityLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWirelessConnectivityLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**GetWirelessConnectivityLogs200Response**](GetWirelessConnectivityLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataUsageNotifications

> ListDataUsageNotifications200Response ListDataUsageNotifications(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterSimCardId(filterSimCardId).Execute()

List SIM card data usage notifications



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
	filterSimCardId := "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9" // string | A valid SIM card ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.ListDataUsageNotifications(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterSimCardId(filterSimCardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.ListDataUsageNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataUsageNotifications`: ListDataUsageNotifications200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.ListDataUsageNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataUsageNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterSimCardId** | **string** | A valid SIM card ID. | 

### Return type

[**ListDataUsageNotifications200Response**](ListDataUsageNotifications200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSimCardDataUsageNotification

> PostSimCardDataUsageNotification201Response PatchSimCardDataUsageNotification(ctx, id).SimCardDataUsageNotification(simCardDataUsageNotification).Execute()

Updates information for a SIM Card Data Usage Notification



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.
	simCardDataUsageNotification := *openapiclient.NewSimCardDataUsageNotification() // SimCardDataUsageNotification | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.PatchSimCardDataUsageNotification(context.Background(), id).SimCardDataUsageNotification(simCardDataUsageNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.PatchSimCardDataUsageNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSimCardDataUsageNotification`: PostSimCardDataUsageNotification201Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.PatchSimCardDataUsageNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSimCardDataUsageNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simCardDataUsageNotification** | [**SimCardDataUsageNotification**](SimCardDataUsageNotification.md) |  | 

### Return type

[**PostSimCardDataUsageNotification201Response**](PostSimCardDataUsageNotification201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimCardDataUsageNotification

> PostSimCardDataUsageNotification201Response PostSimCardDataUsageNotification(ctx).PostSimCardDataUsageNotificationRequest(postSimCardDataUsageNotificationRequest).Execute()

Create a new SIM card data usage notification



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
	postSimCardDataUsageNotificationRequest := *openapiclient.NewPostSimCardDataUsageNotificationRequest("6a09cdc3-8948-47f0-aa62-74ac943d6c58", *openapiclient.NewPostSimCardDataUsageNotificationRequestThreshold()) // PostSimCardDataUsageNotificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.PostSimCardDataUsageNotification(context.Background()).PostSimCardDataUsageNotificationRequest(postSimCardDataUsageNotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.PostSimCardDataUsageNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimCardDataUsageNotification`: PostSimCardDataUsageNotification201Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.PostSimCardDataUsageNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimCardDataUsageNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSimCardDataUsageNotificationRequest** | [**PostSimCardDataUsageNotificationRequest**](PostSimCardDataUsageNotificationRequest.md) |  | 

### Return type

[**PostSimCardDataUsageNotification201Response**](PostSimCardDataUsageNotification201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseESim

> PurchaseESim202Response PurchaseESim(ctx).ESimPurchase(eSimPurchase).Execute()

Purchase eSIMs



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
	eSimPurchase := *openapiclient.NewESimPurchase(int32(10)) // ESimPurchase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.PurchaseESim(context.Background()).ESimPurchase(eSimPurchase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.PurchaseESim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseESim`: PurchaseESim202Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.PurchaseESim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseESimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eSimPurchase** | [**ESimPurchase**](ESimPurchase.md) |  | 

### Return type

[**PurchaseESim202Response**](PurchaseESim202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSimCards

> PurchaseESim202Response RegisterSimCards(ctx).SIMCardRegistration(sIMCardRegistration).Execute()

Register SIM cards



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
	sIMCardRegistration := *openapiclient.NewSIMCardRegistration([]string{"RegistrationCodes_example"}) // SIMCardRegistration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.RegisterSimCards(context.Background()).SIMCardRegistration(sIMCardRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.RegisterSimCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterSimCards`: PurchaseESim202Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.RegisterSimCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSimCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sIMCardRegistration** | [**SIMCardRegistration**](SIMCardRegistration.md) |  | 

### Return type

[**PurchaseESim202Response**](PurchaseESim202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSimCardPublicIp

> GetSimCardAction200Response RemoveSimCardPublicIp(ctx, id).Execute()

Request removing a SIM card public IP



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.RemoveSimCardPublicIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.RemoveSimCardPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSimCardPublicIp`: GetSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.RemoveSimCardPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSimCardPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardAction200Response**](GetSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPublicIPsBulk

> SetPublicIPsBulk202Response SetPublicIPsBulk(ctx).SetPublicIPsBulkRequest(setPublicIPsBulkRequest).Execute()

Request bulk setting SIM card public IPs.



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
	setPublicIPsBulkRequest := *openapiclient.NewSetPublicIPsBulkRequest([]string{"6a09cdc3-8948-47f0-aa62-74ac943d6c58"}) // SetPublicIPsBulkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.SetPublicIPsBulk(context.Background()).SetPublicIPsBulkRequest(setPublicIPsBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.SetPublicIPsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPublicIPsBulk`: SetPublicIPsBulk202Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.SetPublicIPsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPublicIPsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPublicIPsBulkRequest** | [**SetPublicIPsBulkRequest**](SetPublicIPsBulkRequest.md) |  | 

### Return type

[**SetPublicIPsBulk202Response**](SetPublicIPsBulk202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSimCardPublicIp

> GetSimCardAction200Response SetSimCardPublicIp(ctx, id).RegionCode(regionCode).Execute()

Request setting a SIM card public IP



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.
	regionCode := "dc2" // string | The code of the region where the public IP should be assigned. A list of available regions can be found at the regions endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.SetSimCardPublicIp(context.Background(), id).RegionCode(regionCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.SetSimCardPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSimCardPublicIp`: GetSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.SetSimCardPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSimCardPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionCode** | **string** | The code of the region where the public IP should be assigned. A list of available regions can be found at the regions endpoint | 

### Return type

[**GetSimCardAction200Response**](GetSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSimCardStandby

> GetSimCardAction200Response SetSimCardStandby(ctx, id).Execute()

Request setting a SIM card to standby



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.SetSimCardStandby(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.SetSimCardStandby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSimCardStandby`: GetSimCardAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.SetSimCardStandby`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSimCardStandbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardAction200Response**](GetSimCardAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSimCard

> GetSimCard200Response UpdateSimCard(ctx, id).SIMCard(sIMCard).Execute()

Update a SIM card



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM.
	sIMCard := *openapiclient.NewSIMCard() // SIMCard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.UpdateSimCard(context.Background(), id).SIMCard(sIMCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.UpdateSimCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSimCard`: GetSimCard200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.UpdateSimCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSimCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sIMCard** | [**SIMCard**](SIMCard.md) |  | 

### Return type

[**GetSimCard200Response**](GetSimCard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRegistrationCodes

> SIMCardRegistrationCodeValidations ValidateRegistrationCodes(ctx).ValidateRegistrationCodesRequest(validateRegistrationCodesRequest).Execute()

Validate SIM cards registration codes



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
	validateRegistrationCodesRequest := *openapiclient.NewValidateRegistrationCodesRequest() // ValidateRegistrationCodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardsAPI.ValidateRegistrationCodes(context.Background()).ValidateRegistrationCodesRequest(validateRegistrationCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardsAPI.ValidateRegistrationCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateRegistrationCodes`: SIMCardRegistrationCodeValidations
	fmt.Fprintf(os.Stdout, "Response from `SIMCardsAPI.ValidateRegistrationCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRegistrationCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateRegistrationCodesRequest** | [**ValidateRegistrationCodesRequest**](ValidateRegistrationCodesRequest.md) |  | 

### Return type

[**SIMCardRegistrationCodeValidations**](SIMCardRegistrationCodeValidations.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

