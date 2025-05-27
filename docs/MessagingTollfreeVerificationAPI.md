# \MessagingTollfreeVerificationAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVerificationRequest**](MessagingTollfreeVerificationAPI.md#DeleteVerificationRequest) | **Delete** /messaging_tollfree/verification/requests/{id} | Delete Verification Request
[**GetVerificationRequest**](MessagingTollfreeVerificationAPI.md#GetVerificationRequest) | **Get** /messaging_tollfree/verification/requests/{id} | Get Verification Request
[**ListVerificationRequests**](MessagingTollfreeVerificationAPI.md#ListVerificationRequests) | **Get** /messaging_tollfree/verification/requests | List Verification Requests
[**SubmitVerificationRequest**](MessagingTollfreeVerificationAPI.md#SubmitVerificationRequest) | **Post** /messaging_tollfree/verification/requests | Submit Verification Request
[**UpdateVerificationRequest**](MessagingTollfreeVerificationAPI.md#UpdateVerificationRequest) | **Patch** /messaging_tollfree/verification/requests/{id} | Update Verification Request



## DeleteVerificationRequest

> interface{} DeleteVerificationRequest(ctx, id).Execute()

Delete Verification Request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingTollfreeVerificationAPI.DeleteVerificationRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingTollfreeVerificationAPI.DeleteVerificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVerificationRequest`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MessagingTollfreeVerificationAPI.DeleteVerificationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerificationRequest

> VerificationRequestStatus GetVerificationRequest(ctx, id).Execute()

Get Verification Request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingTollfreeVerificationAPI.GetVerificationRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingTollfreeVerificationAPI.GetVerificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerificationRequest`: VerificationRequestStatus
	fmt.Fprintf(os.Stdout, "Response from `MessagingTollfreeVerificationAPI.GetVerificationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerificationRequestStatus**](VerificationRequestStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerificationRequests

> PaginatedVerificationRequestStatus ListVerificationRequests(ctx).Page(page).PageSize(pageSize).DateStart(dateStart).DateEnd(dateEnd).Status(status).PhoneNumber(phoneNumber).Execute()

List Verification Requests



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |          Request this many records per page          This value is automatically clamped if the provided value is too large.         
	dateStart := time.Now() // time.Time |  (optional)
	dateEnd := time.Now() // time.Time |  (optional)
	status := openapiclient.TFVerificationStatus("Verified") // TFVerificationStatus |  (optional)
	phoneNumber := "phoneNumber_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingTollfreeVerificationAPI.ListVerificationRequests(context.Background()).Page(page).PageSize(pageSize).DateStart(dateStart).DateEnd(dateEnd).Status(status).PhoneNumber(phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingTollfreeVerificationAPI.ListVerificationRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVerificationRequests`: PaginatedVerificationRequestStatus
	fmt.Fprintf(os.Stdout, "Response from `MessagingTollfreeVerificationAPI.ListVerificationRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |          Request this many records per page          This value is automatically clamped if the provided value is too large.          | 
 **dateStart** | **time.Time** |  | 
 **dateEnd** | **time.Time** |  | 
 **status** | [**TFVerificationStatus**](TFVerificationStatus.md) |  | 
 **phoneNumber** | **string** |  | 

### Return type

[**PaginatedVerificationRequestStatus**](PaginatedVerificationRequestStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitVerificationRequest

> VerificationRequestEgress SubmitVerificationRequest(ctx).TFVerificationRequest(tFVerificationRequest).Execute()

Submit Verification Request



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
	tFVerificationRequest := *openapiclient.NewTFVerificationRequest("Telnyx LLC", "http://example.com", "600 Congress Avenue", "Austin", "Texas", "78701", "John", "Doe", "email@example.com", "+18005550100", openapiclient.Volume("10"), []openapiclient.TFPhoneNumber{*openapiclient.NewTFPhoneNumber("PhoneNumber_example")}, openapiclient.UseCaseCategories("2FA"), "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal", "Your Telnyx OTP is XXXX", "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset", []openapiclient.Url{*openapiclient.NewUrl("Url_example")}, "AdditionalInformation_example", "IsvReseller_example") // TFVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingTollfreeVerificationAPI.SubmitVerificationRequest(context.Background()).TFVerificationRequest(tFVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingTollfreeVerificationAPI.SubmitVerificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitVerificationRequest`: VerificationRequestEgress
	fmt.Fprintf(os.Stdout, "Response from `MessagingTollfreeVerificationAPI.SubmitVerificationRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tFVerificationRequest** | [**TFVerificationRequest**](TFVerificationRequest.md) |  | 

### Return type

[**VerificationRequestEgress**](VerificationRequestEgress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVerificationRequest

> VerificationRequestEgress UpdateVerificationRequest(ctx, id).TFVerificationRequest(tFVerificationRequest).Execute()

Update Verification Request



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tFVerificationRequest := *openapiclient.NewTFVerificationRequest("Telnyx LLC", "http://example.com", "600 Congress Avenue", "Austin", "Texas", "78701", "John", "Doe", "email@example.com", "+18005550100", openapiclient.Volume("10"), []openapiclient.TFPhoneNumber{*openapiclient.NewTFPhoneNumber("PhoneNumber_example")}, openapiclient.UseCaseCategories("2FA"), "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal", "Your Telnyx OTP is XXXX", "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset", []openapiclient.Url{*openapiclient.NewUrl("Url_example")}, "AdditionalInformation_example", "IsvReseller_example") // TFVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingTollfreeVerificationAPI.UpdateVerificationRequest(context.Background(), id).TFVerificationRequest(tFVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingTollfreeVerificationAPI.UpdateVerificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVerificationRequest`: VerificationRequestEgress
	fmt.Fprintf(os.Stdout, "Response from `MessagingTollfreeVerificationAPI.UpdateVerificationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tFVerificationRequest** | [**TFVerificationRequest**](TFVerificationRequest.md) |  | 

### Return type

[**VerificationRequestEgress**](VerificationRequestEgress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

