# \RequirementTypesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRequirementTypes**](RequirementTypesAPI.md#ListRequirementTypes) | **Get** /requirement_types | List all requirement types
[**RetrieveRequirementType**](RequirementTypesAPI.md#RetrieveRequirementType) | **Get** /requirement_types/{id} | Retrieve a requirement type



## ListRequirementTypes

> ListRequirementTypes200Response ListRequirementTypes(ctx).FilterNameContains(filterNameContains).Sort(sort).Execute()

List all requirement types



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
	filterNameContains := "utility bill" // string | Filters requirement types to those whose name contains a certain string. (optional)
	sort := "country_code" // string | Specifies the sort order for results. If you want to sort by a field in ascending order, include it as a sort parameter. If you want to sort in descending order, prepend a `-` in front of the field name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementTypesAPI.ListRequirementTypes(context.Background()).FilterNameContains(filterNameContains).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementTypesAPI.ListRequirementTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRequirementTypes`: ListRequirementTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `RequirementTypesAPI.ListRequirementTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequirementTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterNameContains** | **string** | Filters requirement types to those whose name contains a certain string. | 
 **sort** | **string** | Specifies the sort order for results. If you want to sort by a field in ascending order, include it as a sort parameter. If you want to sort in descending order, prepend a &#x60;-&#x60; in front of the field name. | 

### Return type

[**ListRequirementTypes200Response**](ListRequirementTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRequirementType

> RetrieveRequirementType200Response RetrieveRequirementType(ctx, id).Execute()

Retrieve a requirement type



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
	id := "a38c217a-8019-48f8-bff6-0fdd9939075b" // string | Uniquely identifies the requirement_type record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementTypesAPI.RetrieveRequirementType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementTypesAPI.RetrieveRequirementType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveRequirementType`: RetrieveRequirementType200Response
	fmt.Fprintf(os.Stdout, "Response from `RequirementTypesAPI.RetrieveRequirementType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the requirement_type record | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequirementTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveRequirementType200Response**](RetrieveRequirementType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

