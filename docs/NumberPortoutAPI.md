# \NumberPortoutAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortoutReport**](NumberPortoutAPI.md#CreatePortoutReport) | **Post** /portouts/reports | Create a port-out related report
[**FindPortoutComments**](NumberPortoutAPI.md#FindPortoutComments) | **Get** /portouts/{id}/comments | List all comments for a portout request
[**FindPortoutRequest**](NumberPortoutAPI.md#FindPortoutRequest) | **Get** /portouts/{id} | Get a portout request
[**GetPortRequestSupportingDocuments**](NumberPortoutAPI.md#GetPortRequestSupportingDocuments) | **Get** /portouts/{id}/supporting_documents | List supporting documents on a portout request
[**GetPortoutReport**](NumberPortoutAPI.md#GetPortoutReport) | **Get** /portouts/reports/{id} | Retrieve a report
[**ListPortoutEvents**](NumberPortoutAPI.md#ListPortoutEvents) | **Get** /portouts/events | List all port-out events
[**ListPortoutRejections**](NumberPortoutAPI.md#ListPortoutRejections) | **Get** /portouts/rejections/{portout_id} | List eligible port-out rejection codes for a specific order
[**ListPortoutReports**](NumberPortoutAPI.md#ListPortoutReports) | **Get** /portouts/reports | List port-out related reports
[**ListPortoutRequest**](NumberPortoutAPI.md#ListPortoutRequest) | **Get** /portouts | List portout requests
[**PostPortRequestComment**](NumberPortoutAPI.md#PostPortRequestComment) | **Post** /portouts/{id}/comments | Create a comment on a portout request
[**PostPortRequestSupportingDocuments**](NumberPortoutAPI.md#PostPortRequestSupportingDocuments) | **Post** /portouts/{id}/supporting_documents | Create a list of supporting documents on a portout request
[**RepublishPortoutEvent**](NumberPortoutAPI.md#RepublishPortoutEvent) | **Post** /portouts/events/{id}/republish | Republish a port-out event
[**ShowPortoutEvent**](NumberPortoutAPI.md#ShowPortoutEvent) | **Get** /portouts/events/{id} | Show a port-out event
[**UpdatePortoutStatus**](NumberPortoutAPI.md#UpdatePortoutStatus) | **Patch** /portouts/{id}/{status} | Update Status



## CreatePortoutReport

> CreatePortoutReport201Response CreatePortoutReport(ctx).CreatePortoutReportRequest(createPortoutReportRequest).Execute()

Create a port-out related report



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
	createPortoutReportRequest := *openapiclient.NewCreatePortoutReportRequest("export_portouts_csv", *openapiclient.NewExportPortoutsCSVReport(*openapiclient.NewExportPortoutsCSVReportFilters())) // CreatePortoutReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.CreatePortoutReport(context.Background()).CreatePortoutReportRequest(createPortoutReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.CreatePortoutReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortoutReport`: CreatePortoutReport201Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.CreatePortoutReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortoutReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortoutReportRequest** | [**CreatePortoutReportRequest**](CreatePortoutReportRequest.md) |  | 

### Return type

[**CreatePortoutReport201Response**](CreatePortoutReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortoutComments

> FindPortoutComments200Response FindPortoutComments(ctx, id).Execute()

List all comments for a portout request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Portout id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.FindPortoutComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.FindPortoutComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindPortoutComments`: FindPortoutComments200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.FindPortoutComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Portout id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortoutCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindPortoutComments200Response**](FindPortoutComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPortoutRequest

> FindPortoutRequest200Response FindPortoutRequest(ctx, id).Execute()

Get a portout request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Portout id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.FindPortoutRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.FindPortoutRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindPortoutRequest`: FindPortoutRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.FindPortoutRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Portout id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPortoutRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindPortoutRequest200Response**](FindPortoutRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortRequestSupportingDocuments

> GetPortRequestSupportingDocuments201Response GetPortRequestSupportingDocuments(ctx, id).Execute()

List supporting documents on a portout request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Portout id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.GetPortRequestSupportingDocuments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.GetPortRequestSupportingDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortRequestSupportingDocuments`: GetPortRequestSupportingDocuments201Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.GetPortRequestSupportingDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Portout id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortRequestSupportingDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPortRequestSupportingDocuments201Response**](GetPortRequestSupportingDocuments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortoutReport

> CreatePortoutReport201Response GetPortoutReport(ctx, id).Execute()

Retrieve a report



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies a report.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.GetPortoutReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.GetPortoutReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortoutReport`: CreatePortoutReport201Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.GetPortoutReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies a report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortoutReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreatePortoutReport201Response**](CreatePortoutReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortoutEvents

> ListPortoutEvents200Response ListPortoutEvents(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterEventType(filterEventType).FilterPortoutId(filterPortoutId).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtLte(filterCreatedAtLte).Execute()

List all port-out events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterEventType := "portout.status_changed" // string | Filter by event type. (optional)
	filterPortoutId := "34dc46a9-53ed-4e01-9454-26227ea13326" // string | Filter by port-out order ID. (optional)
	filterCreatedAtGte := time.Now() // time.Time | Filter by created at greater than or equal to. (optional)
	filterCreatedAtLte := time.Now() // time.Time | Filter by created at less than or equal to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.ListPortoutEvents(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterEventType(filterEventType).FilterPortoutId(filterPortoutId).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtLte(filterCreatedAtLte).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.ListPortoutEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortoutEvents`: ListPortoutEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.ListPortoutEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortoutEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterEventType** | **string** | Filter by event type. | 
 **filterPortoutId** | **string** | Filter by port-out order ID. | 
 **filterCreatedAtGte** | **time.Time** | Filter by created at greater than or equal to. | 
 **filterCreatedAtLte** | **time.Time** | Filter by created at less than or equal to. | 

### Return type

[**ListPortoutEvents200Response**](ListPortoutEvents200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortoutRejections

> ListPortoutRejections200Response ListPortoutRejections(ctx, portoutId).FilterCode(filterCode).FilterCodeIn(filterCodeIn).Execute()

List eligible port-out rejection codes for a specific order



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
	portoutId := "329d6658-8f93-405d-862f-648776e8afd7" // string | Identifies a port out order.
	filterCode := int32(1002) // int32 | Filter rejections of a specific code (optional)
	filterCodeIn := []int32{int32(123)} // []int32 | Filter rejections in a list of codes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.ListPortoutRejections(context.Background(), portoutId).FilterCode(filterCode).FilterCodeIn(filterCodeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.ListPortoutRejections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortoutRejections`: ListPortoutRejections200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.ListPortoutRejections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portoutId** | **string** | Identifies a port out order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortoutRejectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterCode** | **int32** | Filter rejections of a specific code | 
 **filterCodeIn** | **[]int32** | Filter rejections in a list of codes | 

### Return type

[**ListPortoutRejections200Response**](ListPortoutRejections200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortoutReports

> ListPortoutReports200Response ListPortoutReports(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterReportType(filterReportType).FilterStatus(filterStatus).Execute()

List port-out related reports



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
	filterReportType := "export_portouts_csv" // string | Filter reports of a specific type (optional)
	filterStatus := "completed" // string | Filter reports of a specific status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.ListPortoutReports(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterReportType(filterReportType).FilterStatus(filterStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.ListPortoutReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortoutReports`: ListPortoutReports200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.ListPortoutReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortoutReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterReportType** | **string** | Filter reports of a specific type | 
 **filterStatus** | **string** | Filter reports of a specific status | 

### Return type

[**ListPortoutReports200Response**](ListPortoutReports200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortoutRequest

> ListPortoutRequest200Response ListPortoutRequest(ctx).FilterCarrierName(filterCarrierName).FilterPon(filterPon).FilterSpid(filterSpid).FilterStatus(filterStatus).FilterStatusIn(filterStatusIn).FilterPortedOutAtGte(filterPortedOutAtGte).FilterPortedOutAtLte(filterPortedOutAtLte).FilterInsertedAtGte(filterInsertedAtGte).FilterInsertedAtLte(filterInsertedAtLte).FilterFocDate(filterFocDate).FilterPhoneNumber(filterPhoneNumber).FilterSupportKey(filterSupportKey).PageNumber(pageNumber).PageSize(pageSize).Execute()

List portout requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterCarrierName := "filterCarrierName_example" // string | Filter by new carrier name. (optional)
	filterPon := "filterPon_example" // string | Filter by Port Order Number (PON). (optional)
	filterSpid := "filterSpid_example" // string | Filter by new carrier spid. (optional)
	filterStatus := "filterStatus_example" // string | Filter by portout status. (optional)
	filterStatusIn := []string{"FilterStatusIn_example"} // []string | Filter by a list of portout statuses (optional)
	filterPortedOutAtGte := time.Now() // time.Time | Filter by ported_out_at date greater than or equal. (optional)
	filterPortedOutAtLte := time.Now() // time.Time | Filter by ported_out_at date less than or equal. (optional)
	filterInsertedAtGte := time.Now() // time.Time | Filter by inserted_at date greater than or equal. (optional)
	filterInsertedAtLte := time.Now() // time.Time | Filter by inserted_at date less than or equal. (optional)
	filterFocDate := time.Now() // time.Time | Filter by foc_date. Matches all portouts with the same date (optional)
	filterPhoneNumber := "+13035551212" // string | Filter by a phone number on the portout. Matches all portouts with the phone number (optional)
	filterSupportKey := "PO_abc123" // string | Filter by the portout's support_key (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.ListPortoutRequest(context.Background()).FilterCarrierName(filterCarrierName).FilterPon(filterPon).FilterSpid(filterSpid).FilterStatus(filterStatus).FilterStatusIn(filterStatusIn).FilterPortedOutAtGte(filterPortedOutAtGte).FilterPortedOutAtLte(filterPortedOutAtLte).FilterInsertedAtGte(filterInsertedAtGte).FilterInsertedAtLte(filterInsertedAtLte).FilterFocDate(filterFocDate).FilterPhoneNumber(filterPhoneNumber).FilterSupportKey(filterSupportKey).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.ListPortoutRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortoutRequest`: ListPortoutRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.ListPortoutRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortoutRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCarrierName** | **string** | Filter by new carrier name. | 
 **filterPon** | **string** | Filter by Port Order Number (PON). | 
 **filterSpid** | **string** | Filter by new carrier spid. | 
 **filterStatus** | **string** | Filter by portout status. | 
 **filterStatusIn** | **[]string** | Filter by a list of portout statuses | 
 **filterPortedOutAtGte** | **time.Time** | Filter by ported_out_at date greater than or equal. | 
 **filterPortedOutAtLte** | **time.Time** | Filter by ported_out_at date less than or equal. | 
 **filterInsertedAtGte** | **time.Time** | Filter by inserted_at date greater than or equal. | 
 **filterInsertedAtLte** | **time.Time** | Filter by inserted_at date less than or equal. | 
 **filterFocDate** | **time.Time** | Filter by foc_date. Matches all portouts with the same date | 
 **filterPhoneNumber** | **string** | Filter by a phone number on the portout. Matches all portouts with the phone number | 
 **filterSupportKey** | **string** | Filter by the portout&#39;s support_key | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListPortoutRequest200Response**](ListPortoutRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPortRequestComment

> PostPortRequestComment201Response PostPortRequestComment(ctx, id).PostPortRequestCommentRequest(postPortRequestCommentRequest).Execute()

Create a comment on a portout request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Portout id
	postPortRequestCommentRequest := *openapiclient.NewPostPortRequestCommentRequest() // PostPortRequestCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.PostPortRequestComment(context.Background(), id).PostPortRequestCommentRequest(postPortRequestCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.PostPortRequestComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPortRequestComment`: PostPortRequestComment201Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.PostPortRequestComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Portout id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPortRequestCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPortRequestCommentRequest** | [**PostPortRequestCommentRequest**](PostPortRequestCommentRequest.md) |  | 

### Return type

[**PostPortRequestComment201Response**](PostPortRequestComment201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPortRequestSupportingDocuments

> GetPortRequestSupportingDocuments201Response PostPortRequestSupportingDocuments(ctx, id).PostPortRequestSupportingDocumentsRequest(postPortRequestSupportingDocumentsRequest).Execute()

Create a list of supporting documents on a portout request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Portout id
	postPortRequestSupportingDocumentsRequest := *openapiclient.NewPostPortRequestSupportingDocumentsRequest() // PostPortRequestSupportingDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.PostPortRequestSupportingDocuments(context.Background(), id).PostPortRequestSupportingDocumentsRequest(postPortRequestSupportingDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.PostPortRequestSupportingDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPortRequestSupportingDocuments`: GetPortRequestSupportingDocuments201Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.PostPortRequestSupportingDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Portout id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPortRequestSupportingDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPortRequestSupportingDocumentsRequest** | [**PostPortRequestSupportingDocumentsRequest**](PostPortRequestSupportingDocumentsRequest.md) |  | 

### Return type

[**GetPortRequestSupportingDocuments201Response**](GetPortRequestSupportingDocuments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepublishPortoutEvent

> RepublishPortoutEvent(ctx, id).Execute()

Republish a port-out event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the port-out event.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NumberPortoutAPI.RepublishPortoutEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.RepublishPortoutEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the port-out event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepublishPortoutEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPortoutEvent

> ShowPortoutEvent200Response ShowPortoutEvent(ctx, id).Execute()

Show a port-out event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the port-out event.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.ShowPortoutEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.ShowPortoutEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPortoutEvent`: ShowPortoutEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.ShowPortoutEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the port-out event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPortoutEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShowPortoutEvent200Response**](ShowPortoutEvent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePortoutStatus

> FindPortoutRequest200Response UpdatePortoutStatus(ctx, id, status).UpdatePortoutStatusRequest(updatePortoutStatusRequest).Execute()

Update Status



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Portout id
	status := "status_example" // string | Updated portout status
	updatePortoutStatusRequest := *openapiclient.NewUpdatePortoutStatusRequest() // UpdatePortoutStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberPortoutAPI.UpdatePortoutStatus(context.Background(), id, status).UpdatePortoutStatusRequest(updatePortoutStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberPortoutAPI.UpdatePortoutStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortoutStatus`: FindPortoutRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberPortoutAPI.UpdatePortoutStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Portout id | 
**status** | **string** | Updated portout status | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortoutStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePortoutStatusRequest** | [**UpdatePortoutStatusRequest**](UpdatePortoutStatusRequest.md) |  | 

### Return type

[**FindPortoutRequest200Response**](FindPortoutRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

