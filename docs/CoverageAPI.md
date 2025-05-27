# \CoverageAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworkCoverage**](CoverageAPI.md#ListNetworkCoverage) | **Get** /network_coverage | List network coverage locations



## ListNetworkCoverage

> ListNetworkCoverage200Response ListNetworkCoverage(ctx).PageNumber(pageNumber).PageSize(pageSize).FiltersAvailableServicesContains(filtersAvailableServicesContains).FilterLocationRegion(filterLocationRegion).FilterLocationSite(filterLocationSite).FilterLocationPop(filterLocationPop).FilterLocationCode(filterLocationCode).Execute()

List network coverage locations



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
	filtersAvailableServicesContains := openapiclient.AvailableService("cloud_vpn") // AvailableService | The region of associated location to filter on. (optional)
	filterLocationRegion := "AMER" // string | The region of associated location to filter on. (optional)
	filterLocationSite := "SJC" // string | The site of associated location to filter on. (optional)
	filterLocationPop := "SV1" // string | The POP of associated location to filter on. (optional)
	filterLocationCode := "silicon_valley-ca" // string | The code of associated location to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoverageAPI.ListNetworkCoverage(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FiltersAvailableServicesContains(filtersAvailableServicesContains).FilterLocationRegion(filterLocationRegion).FilterLocationSite(filterLocationSite).FilterLocationPop(filterLocationPop).FilterLocationCode(filterLocationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoverageAPI.ListNetworkCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkCoverage`: ListNetworkCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `CoverageAPI.ListNetworkCoverage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filtersAvailableServicesContains** | [**AvailableService**](AvailableService.md) | The region of associated location to filter on. | 
 **filterLocationRegion** | **string** | The region of associated location to filter on. | 
 **filterLocationSite** | **string** | The site of associated location to filter on. | 
 **filterLocationPop** | **string** | The POP of associated location to filter on. | 
 **filterLocationCode** | **string** | The code of associated location to filter on. | 

### Return type

[**ListNetworkCoverage200Response**](ListNetworkCoverage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

