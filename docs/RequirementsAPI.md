# \RequirementsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRequirements**](RequirementsAPI.md#ListRequirements) | **Get** /requirements | List all requirements
[**RetrieveDocumentRequirements**](RequirementsAPI.md#RetrieveDocumentRequirements) | **Get** /requirements/{id} | Retrieve a document requirement



## ListRequirements

> ListRequirements200Response ListRequirements(ctx).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterAction(filterAction).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all requirements



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
	filterCountryCode := "US" // string | Filters results to those applying to a 2-character (ISO 3166-1 alpha-2) country code (optional)
	filterPhoneNumberType := "local" // string | Filters results to those applying to a specific `phone_number_type` (optional)
	filterAction := "porting" // string | Filters requirements to those applying to a specific action. (optional)
	sort := "country_code" // string | Specifies the sort order for results. If you want to sort by a field in ascending order, include it as a sort parameter. If you want to sort in descending order, prepend a `-` in front of the field name. (optional)
	pageNumber := int32(56) // int32 | The page number to load (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.ListRequirements(context.Background()).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterAction(filterAction).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.ListRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRequirements`: ListRequirements200Response
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.ListRequirements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCountryCode** | **string** | Filters results to those applying to a 2-character (ISO 3166-1 alpha-2) country code | 
 **filterPhoneNumberType** | **string** | Filters results to those applying to a specific &#x60;phone_number_type&#x60; | 
 **filterAction** | **string** | Filters requirements to those applying to a specific action. | 
 **sort** | **string** | Specifies the sort order for results. If you want to sort by a field in ascending order, include it as a sort parameter. If you want to sort in descending order, prepend a &#x60;-&#x60; in front of the field name. | 
 **pageNumber** | **int32** | The page number to load | [default to 1]
 **pageSize** | **int32** | The size of the page | [default to 20]

### Return type

[**ListRequirements200Response**](ListRequirements200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDocumentRequirements

> RetrieveDocumentRequirements200Response RetrieveDocumentRequirements(ctx, id).Execute()

Retrieve a document requirement



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
	id := "a9dad8d5-fdbd-49d7-aa23-39bb08a5ebaa" // string | Uniquely identifies the requirement_type record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.RetrieveDocumentRequirements(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.RetrieveDocumentRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDocumentRequirements`: RetrieveDocumentRequirements200Response
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.RetrieveDocumentRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the requirement_type record | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDocumentRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveDocumentRequirements200Response**](RetrieveDocumentRequirements200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

