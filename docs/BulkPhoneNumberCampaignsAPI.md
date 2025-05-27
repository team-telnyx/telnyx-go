# \BulkPhoneNumberCampaignsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssignmentTaskStatus**](BulkPhoneNumberCampaignsAPI.md#GetAssignmentTaskStatus) | **Get** /phoneNumberAssignmentByProfile/{taskId} | Get Assignment Task Status
[**GetPhoneNumberStatus**](BulkPhoneNumberCampaignsAPI.md#GetPhoneNumberStatus) | **Get** /phoneNumberAssignmentByProfile/{taskId}/phoneNumbers | Get Phone Number Status
[**PostAssignMessagingProfileToCampaign**](BulkPhoneNumberCampaignsAPI.md#PostAssignMessagingProfileToCampaign) | **Post** /phoneNumberAssignmentByProfile | Assign Messaging Profile To Campaign



## GetAssignmentTaskStatus

> AssignmentTaskStatusResponse GetAssignmentTaskStatus(ctx, taskId).Execute()

Get Assignment Task Status



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
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberCampaignsAPI.GetAssignmentTaskStatus(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberCampaignsAPI.GetAssignmentTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssignmentTaskStatus`: AssignmentTaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberCampaignsAPI.GetAssignmentTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignmentTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssignmentTaskStatusResponse**](AssignmentTaskStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhoneNumberStatus

> PhoneNumberStatusResponsePaginated GetPhoneNumberStatus(ctx, taskId).RecordsPerPage(recordsPerPage).Page(page).Execute()

Get Phone Number Status



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
	taskId := "taskId_example" // string | 
	recordsPerPage := TODO // interface{} |  (optional) (default to 20)
	page := TODO // interface{} |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberCampaignsAPI.GetPhoneNumberStatus(context.Background(), taskId).RecordsPerPage(recordsPerPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberCampaignsAPI.GetPhoneNumberStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPhoneNumberStatus`: PhoneNumberStatusResponsePaginated
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberCampaignsAPI.GetPhoneNumberStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhoneNumberStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordsPerPage** | [**interface{}**](interface{}.md) |  | [default to 20]
 **page** | [**interface{}**](interface{}.md) |  | [default to 1]

### Return type

[**PhoneNumberStatusResponsePaginated**](PhoneNumberStatusResponsePaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssignMessagingProfileToCampaign

> ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost PostAssignMessagingProfileToCampaign(ctx).AssignProfileToCampaignRequest(assignProfileToCampaignRequest).Execute()

Assign Messaging Profile To Campaign



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
	assignProfileToCampaignRequest := *openapiclient.NewAssignProfileToCampaignRequest("4001767e-ce0f-4cae-9d5f-0d5e636e7809") // AssignProfileToCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPhoneNumberCampaignsAPI.PostAssignMessagingProfileToCampaign(context.Background()).AssignProfileToCampaignRequest(assignProfileToCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPhoneNumberCampaignsAPI.PostAssignMessagingProfileToCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAssignMessagingProfileToCampaign`: ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost
	fmt.Fprintf(os.Stdout, "Response from `BulkPhoneNumberCampaignsAPI.PostAssignMessagingProfileToCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAssignMessagingProfileToCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignProfileToCampaignRequest** | [**AssignProfileToCampaignRequest**](AssignProfileToCampaignRequest.md) |  | 

### Return type

[**ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost**](ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

