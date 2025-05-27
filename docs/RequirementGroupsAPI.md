# \RequirementGroupsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequirementGroup**](RequirementGroupsAPI.md#CreateRequirementGroup) | **Post** /requirement_groups | Create a new requirement group
[**DeleteRequirementGroup**](RequirementGroupsAPI.md#DeleteRequirementGroup) | **Delete** /requirement_groups/{id} | Delete a requirement group by ID
[**GetRequirementGroup**](RequirementGroupsAPI.md#GetRequirementGroup) | **Get** /requirement_groups/{id} | Get a single requirement group by ID
[**GetRequirementGroups**](RequirementGroupsAPI.md#GetRequirementGroups) | **Get** /requirement_groups | List requirement groups
[**SubmitRequirementGroup**](RequirementGroupsAPI.md#SubmitRequirementGroup) | **Post** /requirement_groups/{id}/submit_for_approval | Submit a Requirement Group for Approval
[**UpdateNumberOrderPhoneNumberRequirementGroup**](RequirementGroupsAPI.md#UpdateNumberOrderPhoneNumberRequirementGroup) | **Post** /number_order_phone_numbers/{id}/requirement_group | Update requirement group for a phone number order
[**UpdateRequirementGroup**](RequirementGroupsAPI.md#UpdateRequirementGroup) | **Patch** /requirement_groups/{id} | Update requirement values in requirement group
[**UpdateSubNumberOrderRequirementGroup**](RequirementGroupsAPI.md#UpdateSubNumberOrderRequirementGroup) | **Post** /sub_number_orders/{id}/requirement_group | Update requirement group for a sub number order



## CreateRequirementGroup

> RequirementGroup CreateRequirementGroup(ctx).CreateRequirementGroupRequest(createRequirementGroupRequest).Execute()

Create a new requirement group

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
	createRequirementGroupRequest := *openapiclient.NewCreateRequirementGroupRequest("US", "local", "ordering") // CreateRequirementGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.CreateRequirementGroup(context.Background()).CreateRequirementGroupRequest(createRequirementGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.CreateRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequirementGroup`: RequirementGroup
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.CreateRequirementGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequirementGroupRequest** | [**CreateRequirementGroupRequest**](CreateRequirementGroupRequest.md) |  | 

### Return type

[**RequirementGroup**](RequirementGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequirementGroup

> RequirementGroup DeleteRequirementGroup(ctx, id).Execute()

Delete a requirement group by ID

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
	id := "id_example" // string | ID of the requirement group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.DeleteRequirementGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.DeleteRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRequirementGroup`: RequirementGroup
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.DeleteRequirementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requirement group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequirementGroup**](RequirementGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequirementGroup

> RequirementGroup GetRequirementGroup(ctx, id).Execute()

Get a single requirement group by ID

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
	id := "id_example" // string | ID of the requirement group to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.GetRequirementGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.GetRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequirementGroup`: RequirementGroup
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.GetRequirementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requirement group to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequirementGroup**](RequirementGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequirementGroups

> []RequirementGroup GetRequirementGroups(ctx).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterAction(filterAction).FilterStatus(filterStatus).FilterCustomerReference(filterCustomerReference).Execute()

List requirement groups

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
	filterCountryCode := "FR" // string | Filter requirement groups by country code (iso alpha 2) (optional)
	filterPhoneNumberType := "filterPhoneNumberType_example" // string | Filter requirement groups by phone number type. (optional)
	filterAction := "filterAction_example" // string | Filter requirement groups by action type (optional)
	filterStatus := "filterStatus_example" // string | Filter requirement groups by status (optional)
	filterCustomerReference := "12345" // string | Filter requirement groups by customer reference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.GetRequirementGroups(context.Background()).FilterCountryCode(filterCountryCode).FilterPhoneNumberType(filterPhoneNumberType).FilterAction(filterAction).FilterStatus(filterStatus).FilterCustomerReference(filterCustomerReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.GetRequirementGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequirementGroups`: []RequirementGroup
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.GetRequirementGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequirementGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCountryCode** | **string** | Filter requirement groups by country code (iso alpha 2) | 
 **filterPhoneNumberType** | **string** | Filter requirement groups by phone number type. | 
 **filterAction** | **string** | Filter requirement groups by action type | 
 **filterStatus** | **string** | Filter requirement groups by status | 
 **filterCustomerReference** | **string** | Filter requirement groups by customer reference | 

### Return type

[**[]RequirementGroup**](RequirementGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRequirementGroup

> RequirementGroup SubmitRequirementGroup(ctx, id).Execute()

Submit a Requirement Group for Approval

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
	id := "id_example" // string | ID of the requirement group to submit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.SubmitRequirementGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.SubmitRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitRequirementGroup`: RequirementGroup
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.SubmitRequirementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requirement group to submit | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequirementGroup**](RequirementGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNumberOrderPhoneNumberRequirementGroup

> UpdateNumberOrderPhoneNumberRequirementGroup200Response UpdateNumberOrderPhoneNumberRequirementGroup(ctx, id).UpdateNumberOrderPhoneNumberRequirementGroupRequest(updateNumberOrderPhoneNumberRequirementGroupRequest).Execute()

Update requirement group for a phone number order

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the number order phone number
	updateNumberOrderPhoneNumberRequirementGroupRequest := *openapiclient.NewUpdateNumberOrderPhoneNumberRequirementGroupRequest("RequirementGroupId_example") // UpdateNumberOrderPhoneNumberRequirementGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.UpdateNumberOrderPhoneNumberRequirementGroup(context.Background(), id).UpdateNumberOrderPhoneNumberRequirementGroupRequest(updateNumberOrderPhoneNumberRequirementGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.UpdateNumberOrderPhoneNumberRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNumberOrderPhoneNumberRequirementGroup`: UpdateNumberOrderPhoneNumberRequirementGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.UpdateNumberOrderPhoneNumberRequirementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the number order phone number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNumberOrderPhoneNumberRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNumberOrderPhoneNumberRequirementGroupRequest** | [**UpdateNumberOrderPhoneNumberRequirementGroupRequest**](UpdateNumberOrderPhoneNumberRequirementGroupRequest.md) |  | 

### Return type

[**UpdateNumberOrderPhoneNumberRequirementGroup200Response**](UpdateNumberOrderPhoneNumberRequirementGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRequirementGroup

> RequirementGroup UpdateRequirementGroup(ctx, id).UpdateRequirementGroupRequest(updateRequirementGroupRequest).Execute()

Update requirement values in requirement group

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
	id := "id_example" // string | ID of the requirement group
	updateRequirementGroupRequest := *openapiclient.NewUpdateRequirementGroupRequest() // UpdateRequirementGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.UpdateRequirementGroup(context.Background(), id).UpdateRequirementGroupRequest(updateRequirementGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.UpdateRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRequirementGroup`: RequirementGroup
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.UpdateRequirementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the requirement group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequirementGroupRequest** | [**UpdateRequirementGroupRequest**](UpdateRequirementGroupRequest.md) |  | 

### Return type

[**RequirementGroup**](RequirementGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubNumberOrderRequirementGroup

> SubNumberOrderRequirementGroupResponse UpdateSubNumberOrderRequirementGroup(ctx, id).UpdateNumberOrderPhoneNumberRequirementGroupRequest(updateNumberOrderPhoneNumberRequirementGroupRequest).Execute()

Update requirement group for a sub number order

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the sub number order
	updateNumberOrderPhoneNumberRequirementGroupRequest := *openapiclient.NewUpdateNumberOrderPhoneNumberRequirementGroupRequest("RequirementGroupId_example") // UpdateNumberOrderPhoneNumberRequirementGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementGroupsAPI.UpdateSubNumberOrderRequirementGroup(context.Background(), id).UpdateNumberOrderPhoneNumberRequirementGroupRequest(updateNumberOrderPhoneNumberRequirementGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementGroupsAPI.UpdateSubNumberOrderRequirementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubNumberOrderRequirementGroup`: SubNumberOrderRequirementGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementGroupsAPI.UpdateSubNumberOrderRequirementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the sub number order | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubNumberOrderRequirementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNumberOrderPhoneNumberRequirementGroupRequest** | [**UpdateNumberOrderPhoneNumberRequirementGroupRequest**](UpdateNumberOrderPhoneNumberRequirementGroupRequest.md) |  | 

### Return type

[**SubNumberOrderRequirementGroupResponse**](SubNumberOrderRequirementGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

