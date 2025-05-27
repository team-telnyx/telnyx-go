# \CDRUsageReportsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCDRUsageReportSync**](CDRUsageReportsAPI.md#GetCDRUsageReportSync) | **Get** /reports/cdr_usage_reports/sync | Generates and fetches CDR Usage Reports



## GetCDRUsageReportSync

> CdrGetSyncUsageReportResponse GetCDRUsageReportSync(ctx).AggregationType(aggregationType).ProductBreakdown(productBreakdown).StartDate(startDate).EndDate(endDate).Connections(connections).Execute()

Generates and fetches CDR Usage Reports



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
	aggregationType := "no_aggregation" // string | 
	productBreakdown := "no_breakdown" // string | 
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	connections := []float32{float32(1234567890123)} // []float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDRUsageReportsAPI.GetCDRUsageReportSync(context.Background()).AggregationType(aggregationType).ProductBreakdown(productBreakdown).StartDate(startDate).EndDate(endDate).Connections(connections).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDRUsageReportsAPI.GetCDRUsageReportSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCDRUsageReportSync`: CdrGetSyncUsageReportResponse
	fmt.Fprintf(os.Stdout, "Response from `CDRUsageReportsAPI.GetCDRUsageReportSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCDRUsageReportSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregationType** | **string** |  | 
 **productBreakdown** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **connections** | **[]float32** |  | 

### Return type

[**CdrGetSyncUsageReportResponse**](CdrGetSyncUsageReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

