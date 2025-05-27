# \ProgrammableFaxCommandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFax**](ProgrammableFaxCommandsAPI.md#CancelFax) | **Post** /faxes/{id}/actions/cancel | Cancel a fax
[**DeleteFax**](ProgrammableFaxCommandsAPI.md#DeleteFax) | **Delete** /faxes/{id} | Delete a fax
[**ListFaxes**](ProgrammableFaxCommandsAPI.md#ListFaxes) | **Get** /faxes | View a list of faxes
[**RefreshFax**](ProgrammableFaxCommandsAPI.md#RefreshFax) | **Post** /faxes/{id}/actions/refresh | Refresh a fax
[**SendFax**](ProgrammableFaxCommandsAPI.md#SendFax) | **Post** /faxes | Send a fax
[**ViewFax**](ProgrammableFaxCommandsAPI.md#ViewFax) | **Get** /faxes/{id} | View a fax



## CancelFax

> SuccessfulResponseUponAcceptingCancelFaxCommand CancelFax(ctx, id).Execute()

Cancel a fax



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of a fax.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxCommandsAPI.CancelFax(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxCommandsAPI.CancelFax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelFax`: SuccessfulResponseUponAcceptingCancelFaxCommand
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxCommandsAPI.CancelFax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of a fax. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessfulResponseUponAcceptingCancelFaxCommand**](SuccessfulResponseUponAcceptingCancelFaxCommand.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFax

> DeleteFax(ctx, id).Execute()

Delete a fax

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of a fax.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProgrammableFaxCommandsAPI.DeleteFax(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxCommandsAPI.DeleteFax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of a fax. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFaxes

> ListFaxesResponse ListFaxes(ctx).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLte(filterCreatedAtLte).FilterCreatedAtLt(filterCreatedAtLt).FilterDirectionEq(filterDirectionEq).FilterFromEq(filterFromEq).FilterToEq(filterToEq).PageSize(pageSize).PageNumber(pageNumber).Execute()

View a list of faxes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterCreatedAtGte := time.Now() // time.Time | ISO 8601 date time for filtering faxes created after or on that date (optional)
	filterCreatedAtGt := time.Now() // time.Time | ISO 8601 date time for filtering faxes created after that date (optional)
	filterCreatedAtLte := time.Now() // time.Time | ISO 8601 formatted date time for filtering faxes created on or before that date (optional)
	filterCreatedAtLt := time.Now() // time.Time | ISO 8601 formatted date time for filtering faxes created before that date (optional)
	filterDirectionEq := "inbound" // string | The direction, inbound or outbound, for filtering faxes sent from this account (optional)
	filterFromEq := "+13127367276" // string | The phone number, in E.164 format for filtering faxes sent from this number (optional)
	filterToEq := "+13127367276" // string | The phone number, in E.164 format for filtering faxes sent to this number (optional)
	pageSize := int32(2) // int32 | Number of fax resourcxes for the single page returned (optional)
	pageNumber := int32(2) // int32 | Number of the page to be retrieved (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxCommandsAPI.ListFaxes(context.Background()).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLte(filterCreatedAtLte).FilterCreatedAtLt(filterCreatedAtLt).FilterDirectionEq(filterDirectionEq).FilterFromEq(filterFromEq).FilterToEq(filterToEq).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxCommandsAPI.ListFaxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFaxes`: ListFaxesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxCommandsAPI.ListFaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCreatedAtGte** | **time.Time** | ISO 8601 date time for filtering faxes created after or on that date | 
 **filterCreatedAtGt** | **time.Time** | ISO 8601 date time for filtering faxes created after that date | 
 **filterCreatedAtLte** | **time.Time** | ISO 8601 formatted date time for filtering faxes created on or before that date | 
 **filterCreatedAtLt** | **time.Time** | ISO 8601 formatted date time for filtering faxes created before that date | 
 **filterDirectionEq** | **string** | The direction, inbound or outbound, for filtering faxes sent from this account | 
 **filterFromEq** | **string** | The phone number, in E.164 format for filtering faxes sent from this number | 
 **filterToEq** | **string** | The phone number, in E.164 format for filtering faxes sent to this number | 
 **pageSize** | **int32** | Number of fax resourcxes for the single page returned | 
 **pageNumber** | **int32** | Number of the page to be retrieved | 

### Return type

[**ListFaxesResponse**](ListFaxesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshFax

> RefreshFaxResponse RefreshFax(ctx, id).Execute()

Refresh a fax



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of a fax.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxCommandsAPI.RefreshFax(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxCommandsAPI.RefreshFax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshFax`: RefreshFaxResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxCommandsAPI.RefreshFax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of a fax. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshFaxResponse**](RefreshFaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFax

> SendFaxResponse SendFax(ctx).SendFaxRequest(sendFaxRequest).Execute()

Send a fax



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
	sendFaxRequest := *openapiclient.NewSendFaxRequest("234423", "+13127367276", "+13125790015") // SendFaxRequest | Send fax request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxCommandsAPI.SendFax(context.Background()).SendFaxRequest(sendFaxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxCommandsAPI.SendFax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFax`: SendFaxResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxCommandsAPI.SendFax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendFaxRequest** | [**SendFaxRequest**](SendFaxRequest.md) | Send fax request | 

### Return type

[**SendFaxResponse**](SendFaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewFax

> GetFaxResponse ViewFax(ctx, id).Execute()

View a fax

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of a fax.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammableFaxCommandsAPI.ViewFax(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammableFaxCommandsAPI.ViewFax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewFax`: GetFaxResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammableFaxCommandsAPI.ViewFax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of a fax. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFaxResponse**](GetFaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

