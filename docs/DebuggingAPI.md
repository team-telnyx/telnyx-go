# \DebuggingAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCallEvents**](DebuggingAPI.md#ListCallEvents) | **Get** /call_events | List call events



## ListCallEvents

> ListCallEventsResponse ListCallEvents(ctx).FilterLegId(filterLegId).FilterApplicationSessionId(filterApplicationSessionId).FilterConnectionId(filterConnectionId).FilterProduct(filterProduct).FilterFrom(filterFrom).FilterTo(filterTo).FilterFailed(filterFailed).FilterType(filterType).FilterName(filterName).FilterOccurredAtGt(filterOccurredAtGt).FilterOccurredAtGte(filterOccurredAtGte).FilterOccurredAtLt(filterOccurredAtLt).FilterOccurredAtLte(filterOccurredAtLte).FilterOccurredAtEq(filterOccurredAtEq).PageNumber(pageNumber).PageSize(pageSize).Execute()

List call events



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
	filterLegId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of an individual call leg. (optional)
	filterApplicationSessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the call session. A session may include multiple call leg events. (optional)
	filterConnectionId := "filterConnectionId_example" // string | The unique identifier of the conection. (optional)
	filterProduct := "texml" // string | Filter by product. (optional)
	filterFrom := "+12025550142" // string | Filter by From number. (optional)
	filterTo := "+12025550142" // string | Filter by To number. (optional)
	filterFailed := false // bool | Delivery failed or not. (optional)
	filterType := "webhook" // string | Event type (optional)
	filterName := "webhook" // string | Event name (optional)
	filterOccurredAtGt := "2019-03-29T11:10:00Z" // string | Event occurred_at: greater than (optional)
	filterOccurredAtGte := "2019-03-29T11:10:00Z" // string | Event occurred_at: greater than or equal (optional)
	filterOccurredAtLt := "2019-03-29T11:10:00Z" // string | Event occurred_at: lower than (optional)
	filterOccurredAtLte := "2019-03-29T11:10:00Z" // string | Event occurred_at: lower than or equal (optional)
	filterOccurredAtEq := "2019-03-29T11:10:00Z" // string | Event occurred_at: equal (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DebuggingAPI.ListCallEvents(context.Background()).FilterLegId(filterLegId).FilterApplicationSessionId(filterApplicationSessionId).FilterConnectionId(filterConnectionId).FilterProduct(filterProduct).FilterFrom(filterFrom).FilterTo(filterTo).FilterFailed(filterFailed).FilterType(filterType).FilterName(filterName).FilterOccurredAtGt(filterOccurredAtGt).FilterOccurredAtGte(filterOccurredAtGte).FilterOccurredAtLt(filterOccurredAtLt).FilterOccurredAtLte(filterOccurredAtLte).FilterOccurredAtEq(filterOccurredAtEq).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DebuggingAPI.ListCallEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCallEvents`: ListCallEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `DebuggingAPI.ListCallEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCallEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLegId** | **string** | The unique identifier of an individual call leg. | 
 **filterApplicationSessionId** | **string** | The unique identifier of the call session. A session may include multiple call leg events. | 
 **filterConnectionId** | **string** | The unique identifier of the conection. | 
 **filterProduct** | **string** | Filter by product. | 
 **filterFrom** | **string** | Filter by From number. | 
 **filterTo** | **string** | Filter by To number. | 
 **filterFailed** | **bool** | Delivery failed or not. | 
 **filterType** | **string** | Event type | 
 **filterName** | **string** | Event name | 
 **filterOccurredAtGt** | **string** | Event occurred_at: greater than | 
 **filterOccurredAtGte** | **string** | Event occurred_at: greater than or equal | 
 **filterOccurredAtLt** | **string** | Event occurred_at: lower than | 
 **filterOccurredAtLte** | **string** | Event occurred_at: lower than or equal | 
 **filterOccurredAtEq** | **string** | Event occurred_at: equal | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListCallEventsResponse**](ListCallEventsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

