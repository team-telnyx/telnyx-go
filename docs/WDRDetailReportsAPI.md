# \WDRDetailReportsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaginatedWdrs**](WDRDetailReportsAPI.md#GetPaginatedWdrs) | **Get** /reports/wdrs | Fetches all Wdr records



## GetPaginatedWdrs

> ExternalWdrGetDetailResponse GetPaginatedWdrs(ctx).StartDate(startDate).EndDate(endDate).Id(id).Mcc(mcc).Mnc(mnc).Imsi(imsi).SimGroupName(simGroupName).SimGroupId(simGroupId).SimCardId(simCardId).PhoneNumber(phoneNumber).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()

Fetches all Wdr records



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
	startDate := "2021-05-01T00:00:00Z" // string | Start date (optional)
	endDate := "2021-06-01T00:00:00Z" // string | End date (optional)
	id := "e093fbe0-5bde-11eb-ae93-0242ac130002" // string |  (optional)
	mcc := "204" // string |  (optional)
	mnc := "01" // string |  (optional)
	imsi := "123456" // string |  (optional)
	simGroupName := "sim name" // string |  (optional)
	simGroupId := "f05a189f-7c46-4531-ac56-1460dc465a42" // string |  (optional)
	simCardId := "877f80a6-e5b2-4687-9a04-88076265720f" // string |  (optional)
	phoneNumber := "+12345678910" // string |  (optional)
	pageNumber := int32(56) // int32 | Page number (optional) (default to 1)
	pageSize := int32(56) // int32 | Size of the page (optional) (default to 20)
	sort := []string{"Inner_example"} // []string |  (optional) (default to ["created_at"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WDRDetailReportsAPI.GetPaginatedWdrs(context.Background()).StartDate(startDate).EndDate(endDate).Id(id).Mcc(mcc).Mnc(mnc).Imsi(imsi).SimGroupName(simGroupName).SimGroupId(simGroupId).SimCardId(simCardId).PhoneNumber(phoneNumber).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WDRDetailReportsAPI.GetPaginatedWdrs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaginatedWdrs`: ExternalWdrGetDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `WDRDetailReportsAPI.GetPaginatedWdrs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaginatedWdrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Start date | 
 **endDate** | **string** | End date | 
 **id** | **string** |  | 
 **mcc** | **string** |  | 
 **mnc** | **string** |  | 
 **imsi** | **string** |  | 
 **simGroupName** | **string** |  | 
 **simGroupId** | **string** |  | 
 **simCardId** | **string** |  | 
 **phoneNumber** | **string** |  | 
 **pageNumber** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Size of the page | [default to 20]
 **sort** | **[]string** |  | [default to [&quot;created_at&quot;]]

### Return type

[**ExternalWdrGetDetailResponse**](ExternalWdrGetDetailResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

