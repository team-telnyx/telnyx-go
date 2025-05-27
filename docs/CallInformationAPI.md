# \CallInformationAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListConnectionActiveCalls**](CallInformationAPI.md#ListConnectionActiveCalls) | **Get** /connections/{connection_id}/active_calls | List all active calls for given connection
[**RetrieveCallStatus**](CallInformationAPI.md#RetrieveCallStatus) | **Get** /calls/{call_control_id} | Retrieve a call status



## ListConnectionActiveCalls

> ActiveCallsResponse ListConnectionActiveCalls(ctx, connectionId).PageLimit(pageLimit).PageAfter(pageAfter).PageBefore(pageBefore).Execute()

List all active calls for given connection



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
	connectionId := "connectionId_example" // string | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource.
	pageLimit := int32(56) // int32 | Limit of records per single page (optional) (default to 20)
	pageAfter := "pageAfter_example" // string | Opaque identifier of next page (optional) (default to "null")
	pageBefore := "pageBefore_example" // string | Opaque identifier of previous page (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallInformationAPI.ListConnectionActiveCalls(context.Background(), connectionId).PageLimit(pageLimit).PageAfter(pageAfter).PageBefore(pageBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallInformationAPI.ListConnectionActiveCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectionActiveCalls`: ActiveCallsResponse
	fmt.Fprintf(os.Stdout, "Response from `CallInformationAPI.ListConnectionActiveCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionActiveCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageLimit** | **int32** | Limit of records per single page | [default to 20]
 **pageAfter** | **string** | Opaque identifier of next page | [default to &quot;null&quot;]
 **pageBefore** | **string** | Opaque identifier of previous page | [default to &quot;null&quot;]

### Return type

[**ActiveCallsResponse**](ActiveCallsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCallStatus

> RetrieveCallStatusResponse RetrieveCallStatus(ctx, callControlId).Execute()

Retrieve a call status



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallInformationAPI.RetrieveCallStatus(context.Background(), callControlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallInformationAPI.RetrieveCallStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCallStatus`: RetrieveCallStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CallInformationAPI.RetrieveCallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCallStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveCallStatusResponse**](RetrieveCallStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

