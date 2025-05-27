# \MDRDetailReportsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaginatedMdrs**](MDRDetailReportsAPI.md#GetPaginatedMdrs) | **Get** /reports/mdrs | Fetch all Mdr records



## GetPaginatedMdrs

> MdrGetDetailResponse GetPaginatedMdrs(ctx).StartDate(startDate).EndDate(endDate).Id(id).Direction(direction).Profile(profile).Cld(cld).Cli(cli).Status(status).MessageType(messageType).Execute()

Fetch all Mdr records



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
	startDate := "startDate_example" // string | Pagination start date (optional)
	endDate := "endDate_example" // string | Pagination end date (optional)
	id := "e093fbe0-5bde-11eb-ae93-0242ac130002" // string |  (optional)
	direction := "INBOUND" // string |  (optional)
	profile := "My profile" // string |  (optional)
	cld := "+15551237654" // string |  (optional)
	cli := "+15551237654" // string |  (optional)
	status := "DELIVERED" // string |  (optional)
	messageType := "SMS" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDRDetailReportsAPI.GetPaginatedMdrs(context.Background()).StartDate(startDate).EndDate(endDate).Id(id).Direction(direction).Profile(profile).Cld(cld).Cli(cli).Status(status).MessageType(messageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDRDetailReportsAPI.GetPaginatedMdrs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaginatedMdrs`: MdrGetDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `MDRDetailReportsAPI.GetPaginatedMdrs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaginatedMdrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Pagination start date | 
 **endDate** | **string** | Pagination end date | 
 **id** | **string** |  | 
 **direction** | **string** |  | 
 **profile** | **string** |  | 
 **cld** | **string** |  | 
 **cli** | **string** |  | 
 **status** | **string** |  | 
 **messageType** | **string** |  | 

### Return type

[**MdrGetDetailResponse**](MdrGetDetailResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

