# \MDRUsageReportsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsageReport**](MDRUsageReportsAPI.md#DeleteUsageReport) | **Delete** /reports/mdr_usage_reports/{id} | Delete MDR Usage Report
[**GetMDRUsageReportSync**](MDRUsageReportsAPI.md#GetMDRUsageReportSync) | **Get** /reports/mdr_usage_reports/sync | Generate and fetch MDR Usage Report
[**GetMdrUsageReports**](MDRUsageReportsAPI.md#GetMdrUsageReports) | **Get** /reports/mdr_usage_reports | Fetch all Messaging usage reports
[**GetUsageReport**](MDRUsageReportsAPI.md#GetUsageReport) | **Get** /reports/mdr_usage_reports/{id} | Retrieve messaging report
[**SubmitUsageReport**](MDRUsageReportsAPI.md#SubmitUsageReport) | **Post** /reports/mdr_usage_reports | Create MDR Usage Report



## DeleteUsageReport

> MdrDeleteUsageReportsResponse DeleteUsageReport(ctx, id).Execute()

Delete MDR Usage Report



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDRUsageReportsAPI.DeleteUsageReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDRUsageReportsAPI.DeleteUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsageReport`: MdrDeleteUsageReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `MDRUsageReportsAPI.DeleteUsageReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MdrDeleteUsageReportsResponse**](MdrDeleteUsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMDRUsageReportSync

> MdrGetSyncUsageReportResponse GetMDRUsageReportSync(ctx).AggregationType(aggregationType).StartDate(startDate).EndDate(endDate).Profiles(profiles).Execute()

Generate and fetch MDR Usage Report



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
	aggregationType := "profile" // string | 
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	profiles := []string{"My profile"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDRUsageReportsAPI.GetMDRUsageReportSync(context.Background()).AggregationType(aggregationType).StartDate(startDate).EndDate(endDate).Profiles(profiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDRUsageReportsAPI.GetMDRUsageReportSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMDRUsageReportSync`: MdrGetSyncUsageReportResponse
	fmt.Fprintf(os.Stdout, "Response from `MDRUsageReportsAPI.GetMDRUsageReportSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMDRUsageReportSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregationType** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **profiles** | **[]string** |  | 

### Return type

[**MdrGetSyncUsageReportResponse**](MdrGetSyncUsageReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMdrUsageReports

> MdrGetUsageReportsResponse GetMdrUsageReports(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

Fetch all Messaging usage reports



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
	pageNumber := int32(56) // int32 | Page number (optional) (default to 1)
	pageSize := int32(56) // int32 | Size of the page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDRUsageReportsAPI.GetMdrUsageReports(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDRUsageReportsAPI.GetMdrUsageReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMdrUsageReports`: MdrGetUsageReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `MDRUsageReportsAPI.GetMdrUsageReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMdrUsageReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Size of the page | [default to 20]

### Return type

[**MdrGetUsageReportsResponse**](MdrGetUsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageReport

> MdrGetUsageReportsByIdResponse GetUsageReport(ctx, id).Execute()

Retrieve messaging report



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDRUsageReportsAPI.GetUsageReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDRUsageReportsAPI.GetUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageReport`: MdrGetUsageReportsByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `MDRUsageReportsAPI.GetUsageReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MdrGetUsageReportsByIdResponse**](MdrGetUsageReportsByIdResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUsageReport

> MdrPostUsageReportsResponse SubmitUsageReport(ctx).MdrPostUsageReportRequest(mdrPostUsageReportRequest).Execute()

Create MDR Usage Report



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
	mdrPostUsageReportRequest := *openapiclient.NewMdrPostUsageReportRequest(time.Now(), time.Now(), "AggregationType_example") // MdrPostUsageReportRequest | Mdr usage report data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDRUsageReportsAPI.SubmitUsageReport(context.Background()).MdrPostUsageReportRequest(mdrPostUsageReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDRUsageReportsAPI.SubmitUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitUsageReport`: MdrPostUsageReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `MDRUsageReportsAPI.SubmitUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mdrPostUsageReportRequest** | [**MdrPostUsageReportRequest**](MdrPostUsageReportRequest.md) | Mdr usage report data | 

### Return type

[**MdrPostUsageReportsResponse**](MdrPostUsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

