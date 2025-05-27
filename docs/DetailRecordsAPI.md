# \DetailRecordsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchDetailRecords**](DetailRecordsAPI.md#SearchDetailRecords) | **Get** /detail_records | Search detail records



## SearchDetailRecords

> DetailRecordsSearchResponse SearchDetailRecords(ctx).FilterRecordType(filterRecordType).FilterDateRange(filterDateRange).Filter(filter).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()

Search detail records



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
	filterRecordType := "messaging" // string | Filter by the given record type.
	filterDateRange := "today" // string | Filter by the given user-friendly date range. You can specify one of the following enum values, or a dynamic one using this format: last_N_days. (optional)
	filter := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Filter records on a given record attribute and value. <br/>Example: filter[status]=delivered (optional)
	pageNumber := int32(56) // int32 | Page number (optional) (default to 1)
	pageSize := int32(56) // int32 | Page size (optional) (default to 20)
	sort := []string{"Inner_example"} // []string | Specifies the sort order for results. <br/>Example: sort=-created_at (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DetailRecordsAPI.SearchDetailRecords(context.Background()).FilterRecordType(filterRecordType).FilterDateRange(filterDateRange).Filter(filter).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetailRecordsAPI.SearchDetailRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDetailRecords`: DetailRecordsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DetailRecordsAPI.SearchDetailRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDetailRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRecordType** | **string** | Filter by the given record type. | 
 **filterDateRange** | **string** | Filter by the given user-friendly date range. You can specify one of the following enum values, or a dynamic one using this format: last_N_days. | 
 **filter** | **map[string]interface{}** | Filter records on a given record attribute and value. &lt;br/&gt;Example: filter[status]&#x3D;delivered | 
 **pageNumber** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Page size | [default to 20]
 **sort** | **[]string** | Specifies the sort order for results. &lt;br/&gt;Example: sort&#x3D;-created_at | 

### Return type

[**DetailRecordsSearchResponse**](DetailRecordsSearchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

