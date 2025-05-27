# \WebhooksAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookDeliveries**](WebhooksAPI.md#GetWebhookDeliveries) | **Get** /webhook_deliveries | List webhook deliveries
[**GetWebhookDelivery**](WebhooksAPI.md#GetWebhookDelivery) | **Get** /webhook_deliveries/{id} | Find webhook_delivery details by ID



## GetWebhookDeliveries

> GetWebhookDeliveries200Response GetWebhookDeliveries(ctx).FilterStatusEq(filterStatusEq).FilterEventType(filterEventType).FilterWebhookContains(filterWebhookContains).FilterAttemptsContains(filterAttemptsContains).FilterStartedAtGte(filterStartedAtGte).FilterStartedAtLte(filterStartedAtLte).FilterFinishedAtGte(filterFinishedAtGte).FilterFinishedAtLte(filterFinishedAtLte).PageNumber(pageNumber).PageSize(pageSize).Execute()

List webhook deliveries



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
	filterStatusEq := "delivered" // string | Return only webhook_deliveries matching the given `status` (optional)
	filterEventType := "call_initiated,call.initiated" // string | Return only webhook_deliveries matching the given value of `event_type`. Accepts multiple values separated by a `,`. (optional)
	filterWebhookContains := "call.initiated" // string | Return only webhook deliveries whose `webhook` component contains the given text (optional)
	filterAttemptsContains := "https://fallback.example.com/webhooks" // string | Return only webhook_deliveries whose `attempts` component contains the given text (optional)
	filterStartedAtGte := "2019-03-29T11:10:00Z" // string | Return only webhook_deliveries whose delivery started later than or at given ISO 8601 datetime (optional)
	filterStartedAtLte := "2019-03-29T11:10:00Z" // string | Return only webhook_deliveries whose delivery started earlier than or at given ISO 8601 datetime (optional)
	filterFinishedAtGte := "2019-03-29T11:10:00Z" // string | Return only webhook_deliveries whose delivery finished later than or at given ISO 8601 datetime (optional)
	filterFinishedAtLte := "2019-03-29T11:10:00Z" // string | Return only webhook_deliveries whose delivery finished earlier than or at given ISO 8601 datetime (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhookDeliveries(context.Background()).FilterStatusEq(filterStatusEq).FilterEventType(filterEventType).FilterWebhookContains(filterWebhookContains).FilterAttemptsContains(filterAttemptsContains).FilterStartedAtGte(filterStartedAtGte).FilterStartedAtLte(filterStartedAtLte).FilterFinishedAtGte(filterFinishedAtGte).FilterFinishedAtLte(filterFinishedAtLte).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookDeliveries`: GetWebhookDeliveries200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatusEq** | **string** | Return only webhook_deliveries matching the given &#x60;status&#x60; | 
 **filterEventType** | **string** | Return only webhook_deliveries matching the given value of &#x60;event_type&#x60;. Accepts multiple values separated by a &#x60;,&#x60;. | 
 **filterWebhookContains** | **string** | Return only webhook deliveries whose &#x60;webhook&#x60; component contains the given text | 
 **filterAttemptsContains** | **string** | Return only webhook_deliveries whose &#x60;attempts&#x60; component contains the given text | 
 **filterStartedAtGte** | **string** | Return only webhook_deliveries whose delivery started later than or at given ISO 8601 datetime | 
 **filterStartedAtLte** | **string** | Return only webhook_deliveries whose delivery started earlier than or at given ISO 8601 datetime | 
 **filterFinishedAtGte** | **string** | Return only webhook_deliveries whose delivery finished later than or at given ISO 8601 datetime | 
 **filterFinishedAtLte** | **string** | Return only webhook_deliveries whose delivery finished earlier than or at given ISO 8601 datetime | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**GetWebhookDeliveries200Response**](GetWebhookDeliveries200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookDelivery

> GetWebhookDelivery200Response GetWebhookDelivery(ctx, id).Execute()

Find webhook_delivery details by ID



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
	id := "c9c0797e-901d-4349-a33c-c2c8f31a92c2" // string | Uniquely identifies the webhook_delivery.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhookDelivery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookDelivery`: GetWebhookDelivery200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the webhook_delivery. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWebhookDelivery200Response**](GetWebhookDelivery200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

