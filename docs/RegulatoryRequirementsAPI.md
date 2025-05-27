# \RegulatoryRequirementsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRegulatoryRequirements**](RegulatoryRequirementsAPI.md#ListRegulatoryRequirements) | **Get** /regulatory_requirements | Retrieve regulatory requirements
[**ListRegulatoryRequirementsPhoneNumbers**](RegulatoryRequirementsAPI.md#ListRegulatoryRequirementsPhoneNumbers) | **Get** /phone_numbers_regulatory_requirements | Retrieve regulatory requirements for a list of phone numbers



## ListRegulatoryRequirements

> ListRegulatoryRequirements200Response ListRegulatoryRequirements(ctx).FilterPhoneNumber(filterPhoneNumber).FilterRequirementGroupId(filterRequirementGroupId).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterAction(filterAction).Execute()

Retrieve regulatory requirements

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
	filterPhoneNumber := "+41215470622" // string | Phone number to check requirements for (optional)
	filterRequirementGroupId := "d4c0b4a6-7bd2-40c5-a3b9-2acd99e212b2" // string | ID of requirement group to check requirements for (optional)
	filterCountryCode := "DE" // string | Country code(iso2) to check requirements for (optional)
	filterPhoneNumberType := "filterPhoneNumberType_example" // string | Phone number type to check requirements for (optional)
	filterAction := "filterAction_example" // string | Action to check requirements for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryRequirementsAPI.ListRegulatoryRequirements(context.Background()).FilterPhoneNumber(filterPhoneNumber).FilterRequirementGroupId(filterRequirementGroupId).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterAction(filterAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryRequirementsAPI.ListRegulatoryRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRegulatoryRequirements`: ListRegulatoryRequirements200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryRequirementsAPI.ListRegulatoryRequirements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegulatoryRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPhoneNumber** | **string** | Phone number to check requirements for | 
 **filterRequirementGroupId** | **string** | ID of requirement group to check requirements for | 
 **filterCountryCode** | **string** | Country code(iso2) to check requirements for | 
 **filterPhoneNumberType** | **string** | Phone number type to check requirements for | 
 **filterAction** | **string** | Action to check requirements for | 

### Return type

[**ListRegulatoryRequirements200Response**](ListRegulatoryRequirements200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegulatoryRequirementsPhoneNumbers

> ListRegulatoryRequirementsPhoneNumbers200Response ListRegulatoryRequirementsPhoneNumbers(ctx).FilterPhoneNumber(filterPhoneNumber).Execute()

Retrieve regulatory requirements for a list of phone numbers

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
	filterPhoneNumber := "+41215470622,+41215470633" // string | Record type phone number/ phone numbers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryRequirementsAPI.ListRegulatoryRequirementsPhoneNumbers(context.Background()).FilterPhoneNumber(filterPhoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryRequirementsAPI.ListRegulatoryRequirementsPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRegulatoryRequirementsPhoneNumbers`: ListRegulatoryRequirementsPhoneNumbers200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryRequirementsAPI.ListRegulatoryRequirementsPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegulatoryRequirementsPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPhoneNumber** | **string** | Record type phone number/ phone numbers | 

### Return type

[**ListRegulatoryRequirementsPhoneNumbers200Response**](ListRegulatoryRequirementsPhoneNumbers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

