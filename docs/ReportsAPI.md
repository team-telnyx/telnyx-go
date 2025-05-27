# \ReportsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingGroupReport**](ReportsAPI.md#CreateBillingGroupReport) | **Post** /ledger_billing_group_reports | Create a ledger billing group report
[**GetBillingGroupReport**](ReportsAPI.md#GetBillingGroupReport) | **Get** /ledger_billing_group_reports/{id} | Get a ledger billing group report



## CreateBillingGroupReport

> CreateBillingGroupReport200Response CreateBillingGroupReport(ctx).NewLedgerBillingGroupReport(newLedgerBillingGroupReport).Execute()

Create a ledger billing group report



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
	newLedgerBillingGroupReport := *openapiclient.NewNewLedgerBillingGroupReport() // NewLedgerBillingGroupReport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.CreateBillingGroupReport(context.Background()).NewLedgerBillingGroupReport(newLedgerBillingGroupReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CreateBillingGroupReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingGroupReport`: CreateBillingGroupReport200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.CreateBillingGroupReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingGroupReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newLedgerBillingGroupReport** | [**NewLedgerBillingGroupReport**](NewLedgerBillingGroupReport.md) |  | 

### Return type

[**CreateBillingGroupReport200Response**](CreateBillingGroupReport200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingGroupReport

> CreateBillingGroupReport200Response GetBillingGroupReport(ctx, id).Execute()

Get a ledger billing group report



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
	id := "f5586561-8ff0-4291-a0ac-84fe544797bd" // string | The id of the ledger billing group report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetBillingGroupReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetBillingGroupReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingGroupReport`: CreateBillingGroupReport200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetBillingGroupReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the ledger billing group report | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingGroupReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateBillingGroupReport200Response**](CreateBillingGroupReport200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

