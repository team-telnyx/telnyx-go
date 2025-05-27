# \IPRangesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessIpRangesAccessIpRangeIdDelete**](IPRangesAPI.md#AccessIpRangesAccessIpRangeIdDelete) | **Delete** /access_ip_ranges/{access_ip_range_id} | Delete access IP ranges
[**CreateAccessIPRange**](IPRangesAPI.md#CreateAccessIPRange) | **Post** /access_ip_ranges | Create new Access IP Range
[**ListAccessIpRanges**](IPRangesAPI.md#ListAccessIpRanges) | **Get** /access_ip_ranges | List all Access IP Ranges



## AccessIpRangesAccessIpRangeIdDelete

> AccessIPRangeResponseSchema AccessIpRangesAccessIpRangeIdDelete(ctx, accessIpRangeId).Execute()

Delete access IP ranges

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
	accessIpRangeId := "accessIpRangeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPRangesAPI.AccessIpRangesAccessIpRangeIdDelete(context.Background(), accessIpRangeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPRangesAPI.AccessIpRangesAccessIpRangeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessIpRangesAccessIpRangeIdDelete`: AccessIPRangeResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPRangesAPI.AccessIpRangesAccessIpRangeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessIpRangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessIpRangesAccessIpRangeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessIPRangeResponseSchema**](AccessIPRangeResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessIPRange

> AccessIPRangeResponseSchema CreateAccessIPRange(ctx).AccessIPRangePOST(accessIPRangePOST).Execute()

Create new Access IP Range

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
	accessIPRangePOST := *openapiclient.NewAccessIPRangePOST("CidrBlock_example") // AccessIPRangePOST | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPRangesAPI.CreateAccessIPRange(context.Background()).AccessIPRangePOST(accessIPRangePOST).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPRangesAPI.CreateAccessIPRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessIPRange`: AccessIPRangeResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPRangesAPI.CreateAccessIPRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessIPRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessIPRangePOST** | [**AccessIPRangePOST**](AccessIPRangePOST.md) |  | 

### Return type

[**AccessIPRangeResponseSchema**](AccessIPRangeResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessIpRanges

> AccessIPRangeListResponseSchema ListAccessIpRanges(ctx).FilterCidrBlock(filterCidrBlock).FilterCidrBlockStartswith(filterCidrBlockStartswith).FilterCidrBlockEndswith(filterCidrBlockEndswith).FilterCidrBlockContains(filterCidrBlockContains).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Access IP Ranges

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
	filterCidrBlock := "filterCidrBlock_example" // string |  (optional)
	filterCidrBlockStartswith := "filterCidrBlockStartswith_example" // string |  (optional)
	filterCidrBlockEndswith := "filterCidrBlockEndswith_example" // string |  (optional)
	filterCidrBlockContains := "filterCidrBlockContains_example" // string |  (optional)
	filterCreatedAtGt := time.Now() // time.Time |  (optional)
	filterCreatedAtLt := time.Now() // time.Time |  (optional)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPRangesAPI.ListAccessIpRanges(context.Background()).FilterCidrBlock(filterCidrBlock).FilterCidrBlockStartswith(filterCidrBlockStartswith).FilterCidrBlockEndswith(filterCidrBlockEndswith).FilterCidrBlockContains(filterCidrBlockContains).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPRangesAPI.ListAccessIpRanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessIpRanges`: AccessIPRangeListResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPRangesAPI.ListAccessIpRanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessIpRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCidrBlock** | **string** |  | 
 **filterCidrBlockStartswith** | **string** |  | 
 **filterCidrBlockEndswith** | **string** |  | 
 **filterCidrBlockContains** | **string** |  | 
 **filterCreatedAtGt** | **time.Time** |  | 
 **filterCreatedAtLt** | **time.Time** |  | 
 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**AccessIPRangeListResponseSchema**](AccessIPRangeListResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

