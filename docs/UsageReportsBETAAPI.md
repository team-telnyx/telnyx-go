# \UsageReportsBETAAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageReports**](UsageReportsBETAAPI.md#GetUsageReports) | **Get** /usage_reports | Get Telnyx product usage data (BETA)
[**ListUsageReportsOptions**](UsageReportsBETAAPI.md#ListUsageReportsOptions) | **Get** /usage_reports/options | Get Usage Reports query options (BETA)



## GetUsageReports

> UsageReportsResponse GetUsageReports(ctx).Product(product).Dimensions(dimensions).Metrics(metrics).StartDate(startDate).EndDate(endDate).DateRange(dateRange).Filter(filter).ManagedAccounts(managedAccounts).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Format(format).AuthorizationBearer(authorizationBearer).Execute()

Get Telnyx product usage data (BETA)



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
	product := "wireless" // string | Telnyx product
	dimensions := []string{"Inner_example"} // []string | Breakout by specified product dimensions
	metrics := []string{"Inner_example"} // []string | Specified product usage values
	startDate := "2024-02-01T00:00:00Z" // string | The start date for the time range you are interested in. The maximum time range is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ (optional)
	endDate := "2024-03-01T00:00:00Z" // string | The end date for the time range you are interested in. The maximum time range is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ (optional)
	dateRange := "today" // string | A more user-friendly way to specify the timespan you want to filter by. More options can be found in the Telnyx API Reference docs. (optional)
	filter := "filter[currency]=USD" // string | Filter records on dimensions (optional)
	managedAccounts := false // bool | Return the aggregations for all Managed Accounts under the user making the request. (optional)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	sort := []string{"Inner_example"} // []string | Specifies the sort order for results (optional)
	format := "json" // string | Specify the response format (csv or json). JSON is returned by default, even if not specified. (optional)
	authorizationBearer := "authorizationBearer_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageReportsBETAAPI.GetUsageReports(context.Background()).Product(product).Dimensions(dimensions).Metrics(metrics).StartDate(startDate).EndDate(endDate).DateRange(dateRange).Filter(filter).ManagedAccounts(managedAccounts).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Format(format).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageReportsBETAAPI.GetUsageReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageReports`: UsageReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `UsageReportsBETAAPI.GetUsageReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Telnyx product | 
 **dimensions** | **[]string** | Breakout by specified product dimensions | 
 **metrics** | **[]string** | Specified product usage values | 
 **startDate** | **string** | The start date for the time range you are interested in. The maximum time range is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ | 
 **endDate** | **string** | The end date for the time range you are interested in. The maximum time range is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ | 
 **dateRange** | **string** | A more user-friendly way to specify the timespan you want to filter by. More options can be found in the Telnyx API Reference docs. | 
 **filter** | **string** | Filter records on dimensions | 
 **managedAccounts** | **bool** | Return the aggregations for all Managed Accounts under the user making the request. | 
 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **sort** | **[]string** | Specifies the sort order for results | 
 **format** | **string** | Specify the response format (csv or json). JSON is returned by default, even if not specified. | 
 **authorizationBearer** | **string** |  | 

### Return type

[**UsageReportsResponse**](UsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageReportsOptions

> UsageReportsOptionsResponse ListUsageReportsOptions(ctx).Product(product).AuthorizationBearer(authorizationBearer).Execute()

Get Usage Reports query options (BETA)



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
	product := "wireless" // string | Options (dimensions and metrics) for a given product. If none specified, all products will be returned. (optional)
	authorizationBearer := "authorizationBearer_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageReportsBETAAPI.ListUsageReportsOptions(context.Background()).Product(product).AuthorizationBearer(authorizationBearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageReportsBETAAPI.ListUsageReportsOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsageReportsOptions`: UsageReportsOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UsageReportsBETAAPI.ListUsageReportsOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsageReportsOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Options (dimensions and metrics) for a given product. If none specified, all products will be returned. | 
 **authorizationBearer** | **string** |  | 

### Return type

[**UsageReportsOptionsResponse**](UsageReportsOptionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

