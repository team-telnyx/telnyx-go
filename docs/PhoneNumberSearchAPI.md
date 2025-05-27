# \PhoneNumberSearchAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumberBlocks**](PhoneNumberSearchAPI.md#ListAvailablePhoneNumberBlocks) | **Get** /available_phone_number_blocks | List available phone number blocks
[**ListAvailablePhoneNumbers**](PhoneNumberSearchAPI.md#ListAvailablePhoneNumbers) | **Get** /available_phone_numbers | List available phone numbers



## ListAvailablePhoneNumberBlocks

> ListAvailablePhoneNumbersBlocksResponse ListAvailablePhoneNumberBlocks(ctx).FilterLocality(filterLocality).FilterCountryCode(filterCountryCode).FilterNationalDestinationCode(filterNationalDestinationCode).FilterPhoneNumberType(filterPhoneNumberType).Execute()

List available phone number blocks

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
	filterLocality := "Chicago" // string | Filter phone numbers by city. (optional)
	filterCountryCode := "US" // string | Filter phone numbers by country. (optional)
	filterNationalDestinationCode := "312" // string | Filter by the national destination code of the number. (optional)
	filterPhoneNumberType := "filterPhoneNumberType_example" // string | Filter phone numbers by number type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberSearchAPI.ListAvailablePhoneNumberBlocks(context.Background()).FilterLocality(filterLocality).FilterCountryCode(filterCountryCode).FilterNationalDestinationCode(filterNationalDestinationCode).FilterPhoneNumberType(filterPhoneNumberType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberSearchAPI.ListAvailablePhoneNumberBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailablePhoneNumberBlocks`: ListAvailablePhoneNumbersBlocksResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberSearchAPI.ListAvailablePhoneNumberBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumberBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLocality** | **string** | Filter phone numbers by city. | 
 **filterCountryCode** | **string** | Filter phone numbers by country. | 
 **filterNationalDestinationCode** | **string** | Filter by the national destination code of the number. | 
 **filterPhoneNumberType** | **string** | Filter phone numbers by number type. | 

### Return type

[**ListAvailablePhoneNumbersBlocksResponse**](ListAvailablePhoneNumbersBlocksResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumbers

> ListAvailablePhoneNumbersResponse ListAvailablePhoneNumbers(ctx).FilterPhoneNumberStartsWith(filterPhoneNumberStartsWith).FilterPhoneNumberEndsWith(filterPhoneNumberEndsWith).FilterPhoneNumberContains(filterPhoneNumberContains).FilterLocality(filterLocality).FilterAdministrativeArea(filterAdministrativeArea).FilterCountryCode(filterCountryCode).FilterNationalDestinationCode(filterNationalDestinationCode).FilterRateCenter(filterRateCenter).FilterPhoneNumberType(filterPhoneNumberType).FilterFeatures(filterFeatures).FilterLimit(filterLimit).FilterBestEffort(filterBestEffort).FilterQuickship(filterQuickship).FilterReservable(filterReservable).FilterExcludeHeldNumbers(filterExcludeHeldNumbers).Execute()

List available phone numbers

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
	filterPhoneNumberStartsWith := "TACO" // string | Filter numbers starting with a pattern (excludes NDC if used with `national_destination_code` filter). (optional)
	filterPhoneNumberEndsWith := "TACO" // string | Filter numbers ending with a pattern (excludes NDC if used with `national_destination_code` filter). (optional)
	filterPhoneNumberContains := "TACO" // string | Filter numbers containing a pattern (excludes NDC if used with `national_destination_code` filter). (optional)
	filterLocality := "Chicago" // string | Filter phone numbers by city. (optional)
	filterAdministrativeArea := "IL" // string | Find numbers in a particular US state or CA province. (optional)
	filterCountryCode := "US" // string | Filter phone numbers by country. (optional)
	filterNationalDestinationCode := "312" // string | Filter by the national destination code of the number. (optional)
	filterRateCenter := "Chicago Heights" // string | Filter phone numbers by rate center. This filter is only applicable to USA and Canada numbers. (optional)
	filterPhoneNumberType := "filterPhoneNumberType_example" // string | Filter phone numbers by number type. (optional)
	filterFeatures := []string{"FilterFeatures_example"} // []string | Filter phone numbers with specific features. (optional)
	filterLimit := int32(100) // int32 | Limits the number of results. (optional)
	filterBestEffort := true // bool | Filter to determine if best effort results should be included. Only available in USA/CANADA. (optional)
	filterQuickship := true // bool | Filter to exclude phone numbers that need additional time after to purchase to activate. Only applicable for +1 toll_free numbers. (optional)
	filterReservable := true // bool | Filter to ensure only numbers that can be reserved are included in the results. (optional)
	filterExcludeHeldNumbers := true // bool | Filter to exclude phone numbers that are currently on hold/reserved for your account. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberSearchAPI.ListAvailablePhoneNumbers(context.Background()).FilterPhoneNumberStartsWith(filterPhoneNumberStartsWith).FilterPhoneNumberEndsWith(filterPhoneNumberEndsWith).FilterPhoneNumberContains(filterPhoneNumberContains).FilterLocality(filterLocality).FilterAdministrativeArea(filterAdministrativeArea).FilterCountryCode(filterCountryCode).FilterNationalDestinationCode(filterNationalDestinationCode).FilterRateCenter(filterRateCenter).FilterPhoneNumberType(filterPhoneNumberType).FilterFeatures(filterFeatures).FilterLimit(filterLimit).FilterBestEffort(filterBestEffort).FilterQuickship(filterQuickship).FilterReservable(filterReservable).FilterExcludeHeldNumbers(filterExcludeHeldNumbers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberSearchAPI.ListAvailablePhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailablePhoneNumbers`: ListAvailablePhoneNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberSearchAPI.ListAvailablePhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPhoneNumberStartsWith** | **string** | Filter numbers starting with a pattern (excludes NDC if used with &#x60;national_destination_code&#x60; filter). | 
 **filterPhoneNumberEndsWith** | **string** | Filter numbers ending with a pattern (excludes NDC if used with &#x60;national_destination_code&#x60; filter). | 
 **filterPhoneNumberContains** | **string** | Filter numbers containing a pattern (excludes NDC if used with &#x60;national_destination_code&#x60; filter). | 
 **filterLocality** | **string** | Filter phone numbers by city. | 
 **filterAdministrativeArea** | **string** | Find numbers in a particular US state or CA province. | 
 **filterCountryCode** | **string** | Filter phone numbers by country. | 
 **filterNationalDestinationCode** | **string** | Filter by the national destination code of the number. | 
 **filterRateCenter** | **string** | Filter phone numbers by rate center. This filter is only applicable to USA and Canada numbers. | 
 **filterPhoneNumberType** | **string** | Filter phone numbers by number type. | 
 **filterFeatures** | **[]string** | Filter phone numbers with specific features. | 
 **filterLimit** | **int32** | Limits the number of results. | 
 **filterBestEffort** | **bool** | Filter to determine if best effort results should be included. Only available in USA/CANADA. | 
 **filterQuickship** | **bool** | Filter to exclude phone numbers that need additional time after to purchase to activate. Only applicable for +1 toll_free numbers. | 
 **filterReservable** | **bool** | Filter to ensure only numbers that can be reserved are included in the results. | 
 **filterExcludeHeldNumbers** | **bool** | Filter to exclude phone numbers that are currently on hold/reserved for your account. | 

### Return type

[**ListAvailablePhoneNumbersResponse**](ListAvailablePhoneNumbersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

