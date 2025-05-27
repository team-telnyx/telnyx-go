# \ReportingAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWdrReport**](ReportingAPI.md#CreateWdrReport) | **Post** /wireless/detail_records_reports | Create a Wireless Detail Records (WDRs) Report
[**DeleteWdrReport**](ReportingAPI.md#DeleteWdrReport) | **Delete** /wireless/detail_records_reports/{id} | Delete a Wireless Detail Record (WDR) Report
[**GetWdrReport**](ReportingAPI.md#GetWdrReport) | **Get** /wireless/detail_records_reports/{id} | Get a Wireless Detail Record (WDR) Report
[**GetWdrReports**](ReportingAPI.md#GetWdrReports) | **Get** /wireless/detail_records_reports | Get all Wireless Detail Records (WDRs) Reports



## CreateWdrReport

> CreateWdrReport201Response CreateWdrReport(ctx).WdrReportRequest(wdrReportRequest).Execute()

Create a Wireless Detail Records (WDRs) Report



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
	wdrReportRequest := *openapiclient.NewWdrReportRequest() // WdrReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.CreateWdrReport(context.Background()).WdrReportRequest(wdrReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.CreateWdrReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWdrReport`: CreateWdrReport201Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.CreateWdrReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWdrReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wdrReportRequest** | [**WdrReportRequest**](WdrReportRequest.md) |  | 

### Return type

[**CreateWdrReport201Response**](CreateWdrReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWdrReport

> CreateWdrReport201Response DeleteWdrReport(ctx, id).Execute()

Delete a Wireless Detail Record (WDR) Report



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
	resp, r, err := apiClient.ReportingAPI.DeleteWdrReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.DeleteWdrReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWdrReport`: CreateWdrReport201Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.DeleteWdrReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWdrReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWdrReport201Response**](CreateWdrReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWdrReport

> CreateWdrReport201Response GetWdrReport(ctx, id).Execute()

Get a Wireless Detail Record (WDR) Report



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
	resp, r, err := apiClient.ReportingAPI.GetWdrReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetWdrReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWdrReport`: CreateWdrReport201Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetWdrReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWdrReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWdrReport201Response**](CreateWdrReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWdrReports

> GetWdrReports200Response GetWdrReports(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get all Wireless Detail Records (WDRs) Reports



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
	resp, r, err := apiClient.ReportingAPI.GetWdrReports(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetWdrReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWdrReports`: GetWdrReports200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetWdrReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWdrReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**GetWdrReports200Response**](GetWdrReports200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

