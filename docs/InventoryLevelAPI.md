# \InventoryLevelAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInventoryCoverage**](InventoryLevelAPI.md#CreateInventoryCoverage) | **Get** /inventory_coverage | Create an inventory coverage request



## CreateInventoryCoverage

> CreateInventoryCoverage200Response CreateInventoryCoverage(ctx).FilterGroupBy(filterGroupBy).FilterNpa(filterNpa).FilterNxx(filterNxx).FilterAdministrativeArea(filterAdministrativeArea).FilterPhoneNumberType(filterPhoneNumberType).FilterCountryCode(filterCountryCode).FilterCount(filterCount).FilterFeatures(filterFeatures).Execute()

Create an inventory coverage request



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
	filterGroupBy := "nxx" // string | 
	filterNpa := int32(318) // int32 |  (optional)
	filterNxx := int32(202) // int32 |  (optional)
	filterAdministrativeArea := "LA" // string |  (optional)
	filterPhoneNumberType := "local" // string |  (optional)
	filterCountryCode := "US" // string |  (optional)
	filterCount := true // bool |  (optional)
	filterFeatures := []string{"FilterFeatures_example"} // []string | Filter if the phone number should be used for voice, fax, mms, sms, emergency. Returns features in the response when used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLevelAPI.CreateInventoryCoverage(context.Background()).FilterGroupBy(filterGroupBy).FilterNpa(filterNpa).FilterNxx(filterNxx).FilterAdministrativeArea(filterAdministrativeArea).FilterPhoneNumberType(filterPhoneNumberType).FilterCountryCode(filterCountryCode).FilterCount(filterCount).FilterFeatures(filterFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLevelAPI.CreateInventoryCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInventoryCoverage`: CreateInventoryCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryLevelAPI.CreateInventoryCoverage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInventoryCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGroupBy** | **string** |  | 
 **filterNpa** | **int32** |  | 
 **filterNxx** | **int32** |  | 
 **filterAdministrativeArea** | **string** |  | 
 **filterPhoneNumberType** | **string** |  | 
 **filterCountryCode** | **string** |  | 
 **filterCount** | **bool** |  | 
 **filterFeatures** | **[]string** | Filter if the phone number should be used for voice, fax, mms, sms, emergency. Returns features in the response when used. | 

### Return type

[**CreateInventoryCoverage200Response**](CreateInventoryCoverage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

