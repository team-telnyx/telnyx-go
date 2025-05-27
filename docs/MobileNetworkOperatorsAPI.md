# \MobileNetworkOperatorsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMobileNetworkOperators**](MobileNetworkOperatorsAPI.md#GetMobileNetworkOperators) | **Get** /mobile_network_operators | List mobile network operators



## GetMobileNetworkOperators

> GetMobileNetworkOperators200Response GetMobileNetworkOperators(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterNameStartsWith(filterNameStartsWith).FilterNameContains(filterNameContains).FilterNameEndsWith(filterNameEndsWith).FilterCountryCode(filterCountryCode).FilterMcc(filterMcc).FilterMnc(filterMnc).FilterTadig(filterTadig).FilterNetworkPreferencesEnabled(filterNetworkPreferencesEnabled).Execute()

List mobile network operators



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
	filterNameStartsWith := "AT" // string | Filter by name starting with. (optional)
	filterNameContains := "T&T" // string | Filter by name containing match. (optional)
	filterNameEndsWith := "T" // string | Filter by name ending with. (optional)
	filterCountryCode := "US" // string | Filter by exact country_code. (optional)
	filterMcc := "310" // string | Filter by exact MCC. (optional)
	filterMnc := "410" // string | Filter by exact MNC. (optional)
	filterTadig := "USACG" // string | Filter by exact TADIG. (optional)
	filterNetworkPreferencesEnabled := true // bool | Filter by network_preferences_enabled. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileNetworkOperatorsAPI.GetMobileNetworkOperators(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterNameStartsWith(filterNameStartsWith).FilterNameContains(filterNameContains).FilterNameEndsWith(filterNameEndsWith).FilterCountryCode(filterCountryCode).FilterMcc(filterMcc).FilterMnc(filterMnc).FilterTadig(filterTadig).FilterNetworkPreferencesEnabled(filterNetworkPreferencesEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileNetworkOperatorsAPI.GetMobileNetworkOperators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMobileNetworkOperators`: GetMobileNetworkOperators200Response
	fmt.Fprintf(os.Stdout, "Response from `MobileNetworkOperatorsAPI.GetMobileNetworkOperators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileNetworkOperatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterNameStartsWith** | **string** | Filter by name starting with. | 
 **filterNameContains** | **string** | Filter by name containing match. | 
 **filterNameEndsWith** | **string** | Filter by name ending with. | 
 **filterCountryCode** | **string** | Filter by exact country_code. | 
 **filterMcc** | **string** | Filter by exact MCC. | 
 **filterMnc** | **string** | Filter by exact MNC. | 
 **filterTadig** | **string** | Filter by exact TADIG. | 
 **filterNetworkPreferencesEnabled** | **bool** | Filter by network_preferences_enabled. | 

### Return type

[**GetMobileNetworkOperators200Response**](GetMobileNetworkOperators200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

