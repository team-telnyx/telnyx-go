# \PortingOrdersAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePortingOrder**](PortingOrdersAPI.md#ActivatePortingOrder) | **Post** /porting_orders/{id}/actions/activate | Activate every number in a porting order asynchronously.
[**CancelPortingOrder**](PortingOrdersAPI.md#CancelPortingOrder) | **Post** /porting_orders/{id}/actions/cancel | Cancel a porting order
[**ConfirmPortingOrder**](PortingOrdersAPI.md#ConfirmPortingOrder) | **Post** /porting_orders/{id}/actions/confirm | Submit a porting order.
[**CreateAdditionalDocuments**](PortingOrdersAPI.md#CreateAdditionalDocuments) | **Post** /porting_orders/{id}/additional_documents | Create a list of additional documents
[**CreateLoaConfiguration**](PortingOrdersAPI.md#CreateLoaConfiguration) | **Post** /porting/loa_configurations | Create a LOA configuration
[**CreatePhoneNumberConfigurations**](PortingOrdersAPI.md#CreatePhoneNumberConfigurations) | **Post** /porting_orders/phone_number_configurations | Create a list of phone number configurations
[**CreatePortingOrder**](PortingOrdersAPI.md#CreatePortingOrder) | **Post** /porting_orders | Create a porting order
[**CreatePortingOrderComment**](PortingOrdersAPI.md#CreatePortingOrderComment) | **Post** /porting_orders/{id}/comments | Create a comment for a porting order
[**CreatePortingPhoneNumberBlock**](PortingOrdersAPI.md#CreatePortingPhoneNumberBlock) | **Post** /porting_orders/{porting_order_id}/phone_number_blocks | Create a phone number block
[**CreatePortingPhoneNumberExtension**](PortingOrdersAPI.md#CreatePortingPhoneNumberExtension) | **Post** /porting_orders/{porting_order_id}/phone_number_extensions | Create a phone number extension
[**CreatePortingReport**](PortingOrdersAPI.md#CreatePortingReport) | **Post** /porting/reports | Create a porting related report
[**DeleteAdditionalDocument**](PortingOrdersAPI.md#DeleteAdditionalDocument) | **Delete** /porting_orders/{id}/additional_documents/{additional_document_id} | Delete an additional document
[**DeleteLoaConfiguration**](PortingOrdersAPI.md#DeleteLoaConfiguration) | **Delete** /porting/loa_configurations/{id} | Delete a LOA configuration
[**DeletePortingOrder**](PortingOrdersAPI.md#DeletePortingOrder) | **Delete** /porting_orders/{id} | Delete a porting order
[**DeletePortingPhoneNumberBlock**](PortingOrdersAPI.md#DeletePortingPhoneNumberBlock) | **Delete** /porting_orders/{porting_order_id}/phone_number_blocks/{id} | Delete a phone number block
[**DeletePortingPhoneNumberExtension**](PortingOrdersAPI.md#DeletePortingPhoneNumberExtension) | **Delete** /porting_orders/{porting_order_id}/phone_number_extensions/{id} | Delete a phone number extension
[**GetLoaConfiguration**](PortingOrdersAPI.md#GetLoaConfiguration) | **Get** /porting/loa_configurations/{id} | Retrieve a LOA configuration
[**GetPortingOrder**](PortingOrdersAPI.md#GetPortingOrder) | **Get** /porting_orders/{id} | Retrieve a porting order
[**GetPortingOrderLoaTemplate**](PortingOrdersAPI.md#GetPortingOrderLoaTemplate) | **Get** /porting_orders/{id}/loa_template | Download a porting order loa template
[**GetPortingOrderSubRequest**](PortingOrdersAPI.md#GetPortingOrderSubRequest) | **Get** /porting_orders/{id}/sub_request | Retrieve the associated V1 sub_request_id and port_request_id
[**GetPortingOrdersActivationJob**](PortingOrdersAPI.md#GetPortingOrdersActivationJob) | **Get** /porting_orders/{id}/activation_jobs/{activationJobId} | Retrieve a porting activation job
[**GetPortingReport**](PortingOrdersAPI.md#GetPortingReport) | **Get** /porting/reports/{id} | Retrieve a report
[**ListAdditionalDocuments**](PortingOrdersAPI.md#ListAdditionalDocuments) | **Get** /porting_orders/{id}/additional_documents | List additional documents
[**ListAllowedFocWindows**](PortingOrdersAPI.md#ListAllowedFocWindows) | **Get** /porting_orders/{id}/allowed_foc_windows | List allowed FOC dates
[**ListExceptionTypes**](PortingOrdersAPI.md#ListExceptionTypes) | **Get** /porting_orders/exception_types | List all exception types
[**ListLoaConfigurations**](PortingOrdersAPI.md#ListLoaConfigurations) | **Get** /porting/loa_configurations | List LOA configurations
[**ListPhoneNumberConfigurations**](PortingOrdersAPI.md#ListPhoneNumberConfigurations) | **Get** /porting_orders/phone_number_configurations | List all phone number configurations
[**ListPortingEvents**](PortingOrdersAPI.md#ListPortingEvents) | **Get** /porting/events | List all porting events
[**ListPortingOrderActivationJobs**](PortingOrdersAPI.md#ListPortingOrderActivationJobs) | **Get** /porting_orders/{id}/activation_jobs | List all porting activation jobs
[**ListPortingOrderComments**](PortingOrdersAPI.md#ListPortingOrderComments) | **Get** /porting_orders/{id}/comments | List all comments of a porting order
[**ListPortingOrderRequirements**](PortingOrdersAPI.md#ListPortingOrderRequirements) | **Get** /porting_orders/{id}/requirements | List porting order requirements
[**ListPortingOrders**](PortingOrdersAPI.md#ListPortingOrders) | **Get** /porting_orders | List all porting orders
[**ListPortingPhoneNumberBlocks**](PortingOrdersAPI.md#ListPortingPhoneNumberBlocks) | **Get** /porting_orders/{porting_order_id}/phone_number_blocks | List all phone number blocks
[**ListPortingPhoneNumberExtensions**](PortingOrdersAPI.md#ListPortingPhoneNumberExtensions) | **Get** /porting_orders/{porting_order_id}/phone_number_extensions | List all phone number extensions
[**ListPortingPhoneNumbers**](PortingOrdersAPI.md#ListPortingPhoneNumbers) | **Get** /porting_phone_numbers | List all porting phone numbers
[**ListPortingReports**](PortingOrdersAPI.md#ListPortingReports) | **Get** /porting/reports | List porting related reports
[**ListVerificationCodes**](PortingOrdersAPI.md#ListVerificationCodes) | **Get** /porting_orders/{id}/verification_codes | List verification codes
[**PreviewLoaConfiguration**](PortingOrdersAPI.md#PreviewLoaConfiguration) | **Get** /porting/loa_configurations/{id}/preview | Preview a LOA configuration
[**PreviewLoaConfigurationParams**](PortingOrdersAPI.md#PreviewLoaConfigurationParams) | **Post** /porting/loa_configuration/preview | Preview the LOA configuration parameters
[**RepublishPortingEvent**](PortingOrdersAPI.md#RepublishPortingEvent) | **Post** /porting/events/{id}/republish | Republish a porting event
[**SendPortingVerificationCodes**](PortingOrdersAPI.md#SendPortingVerificationCodes) | **Post** /porting_orders/{id}/verification_codes/send | Send the verification codes
[**SharePortingOrder**](PortingOrdersAPI.md#SharePortingOrder) | **Post** /porting_orders/{id}/actions/share | Share a porting order
[**ShowPortingEvent**](PortingOrdersAPI.md#ShowPortingEvent) | **Get** /porting/events/{id} | Show a porting event
[**UpdateLoaConfiguration**](PortingOrdersAPI.md#UpdateLoaConfiguration) | **Patch** /porting/loa_configurations/{id} | Update a LOA configuration
[**UpdatePortingOrder**](PortingOrdersAPI.md#UpdatePortingOrder) | **Patch** /porting_orders/{id} | Edit a porting order
[**UpdatePortingOrdersActivationJob**](PortingOrdersAPI.md#UpdatePortingOrdersActivationJob) | **Patch** /porting_orders/{id}/activation_jobs/{activationJobId} | Update a porting activation job
[**VerifyPortingVerificationCodes**](PortingOrdersAPI.md#VerifyPortingVerificationCodes) | **Post** /porting_orders/{id}/verification_codes/verify | Verify the verification code for a list of phone numbers



## ActivatePortingOrder

> ActivatePortingOrder202Response ActivatePortingOrder(ctx, id).Execute()

Activate every number in a porting order asynchronously.



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ActivatePortingOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ActivatePortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivatePortingOrder`: ActivatePortingOrder202Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ActivatePortingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivatePortingOrder202Response**](ActivatePortingOrder202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPortingOrder

> CancelPortingOrder200Response CancelPortingOrder(ctx, id).Execute()

Cancel a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CancelPortingOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CancelPortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPortingOrder`: CancelPortingOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CancelPortingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelPortingOrder200Response**](CancelPortingOrder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPortingOrder

> ConfirmPortingOrder200Response ConfirmPortingOrder(ctx, id).Execute()

Submit a porting order.



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ConfirmPortingOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ConfirmPortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmPortingOrder`: ConfirmPortingOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ConfirmPortingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfirmPortingOrder200Response**](ConfirmPortingOrder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdditionalDocuments

> CreateAdditionalDocuments201Response CreateAdditionalDocuments(ctx, id).CreateAdditionalDocumentsRequest(createAdditionalDocumentsRequest).Execute()

Create a list of additional documents



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	createAdditionalDocumentsRequest := *openapiclient.NewCreateAdditionalDocumentsRequest() // CreateAdditionalDocumentsRequest | A list of additional document parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreateAdditionalDocuments(context.Background(), id).CreateAdditionalDocumentsRequest(createAdditionalDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreateAdditionalDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdditionalDocuments`: CreateAdditionalDocuments201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreateAdditionalDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdditionalDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAdditionalDocumentsRequest** | [**CreateAdditionalDocumentsRequest**](CreateAdditionalDocumentsRequest.md) | A list of additional document parameters | 

### Return type

[**CreateAdditionalDocuments201Response**](CreateAdditionalDocuments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoaConfiguration

> CreateLoaConfiguration201Response CreateLoaConfiguration(ctx).PreviewLoaConfigurationParamsRequest(previewLoaConfigurationParamsRequest).Execute()

Create a LOA configuration



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
	previewLoaConfigurationParamsRequest := *openapiclient.NewPreviewLoaConfigurationParamsRequest("My LOA Configuration", *openapiclient.NewPreviewLoaConfigurationParamsRequestLogo("DocumentId_example"), "Telnyx", *openapiclient.NewPreviewLoaConfigurationParamsRequestAddress("600 Congress Avenue", "US"), *openapiclient.NewPreviewLoaConfigurationParamsRequestContact("testing@telnyx.com", "+12003270001")) // PreviewLoaConfigurationParamsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreateLoaConfiguration(context.Background()).PreviewLoaConfigurationParamsRequest(previewLoaConfigurationParamsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreateLoaConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoaConfiguration`: CreateLoaConfiguration201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreateLoaConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewLoaConfigurationParamsRequest** | [**PreviewLoaConfigurationParamsRequest**](PreviewLoaConfigurationParamsRequest.md) |  | 

### Return type

[**CreateLoaConfiguration201Response**](CreateLoaConfiguration201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumberConfigurations

> CreatePhoneNumberConfigurations201Response CreatePhoneNumberConfigurations(ctx).CreatePhoneNumberConfigurationsRequest(createPhoneNumberConfigurationsRequest).Execute()

Create a list of phone number configurations



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
	createPhoneNumberConfigurationsRequest := *openapiclient.NewCreatePhoneNumberConfigurationsRequest() // CreatePhoneNumberConfigurationsRequest | A list of phone number configuration parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreatePhoneNumberConfigurations(context.Background()).CreatePhoneNumberConfigurationsRequest(createPhoneNumberConfigurationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreatePhoneNumberConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePhoneNumberConfigurations`: CreatePhoneNumberConfigurations201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreatePhoneNumberConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhoneNumberConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPhoneNumberConfigurationsRequest** | [**CreatePhoneNumberConfigurationsRequest**](CreatePhoneNumberConfigurationsRequest.md) | A list of phone number configuration parameters | 

### Return type

[**CreatePhoneNumberConfigurations201Response**](CreatePhoneNumberConfigurations201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortingOrder

> CreatePortingOrder201Response CreatePortingOrder(ctx).CreatePortingOrder(createPortingOrder).Execute()

Create a porting order



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
	createPortingOrder := *openapiclient.NewCreatePortingOrder([]string{"PhoneNumbers_example"}) // CreatePortingOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreatePortingOrder(context.Background()).CreatePortingOrder(createPortingOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreatePortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortingOrder`: CreatePortingOrder201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreatePortingOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortingOrder** | [**CreatePortingOrder**](CreatePortingOrder.md) |  | 

### Return type

[**CreatePortingOrder201Response**](CreatePortingOrder201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortingOrderComment

> CreatePortingOrderComment201Response CreatePortingOrderComment(ctx, id).CreatePortingOrderComment(createPortingOrderComment).Execute()

Create a comment for a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	createPortingOrderComment := *openapiclient.NewCreatePortingOrderComment() // CreatePortingOrderComment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreatePortingOrderComment(context.Background(), id).CreatePortingOrderComment(createPortingOrderComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreatePortingOrderComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortingOrderComment`: CreatePortingOrderComment201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreatePortingOrderComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortingOrderCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPortingOrderComment** | [**CreatePortingOrderComment**](CreatePortingOrderComment.md) |  | 

### Return type

[**CreatePortingOrderComment201Response**](CreatePortingOrderComment201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortingPhoneNumberBlock

> CreatePortingPhoneNumberBlock201Response CreatePortingPhoneNumberBlock(ctx, portingOrderId).CreatePortingPhoneNumberBlockRequest(createPortingPhoneNumberBlockRequest).Execute()

Create a phone number block



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
	portingOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the Porting Order associated with the phone number block
	createPortingPhoneNumberBlockRequest := *openapiclient.NewCreatePortingPhoneNumberBlockRequest(*openapiclient.NewCreatePortingPhoneNumberBlockRequestPhoneNumberRange("+4930244999901", "+4930244999910"), []openapiclient.CreatePortingPhoneNumberBlockRequestActivationRangesInner{*openapiclient.NewCreatePortingPhoneNumberBlockRequestActivationRangesInner("+4930244999901", "+4930244999910")}) // CreatePortingPhoneNumberBlockRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreatePortingPhoneNumberBlock(context.Background(), portingOrderId).CreatePortingPhoneNumberBlockRequest(createPortingPhoneNumberBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreatePortingPhoneNumberBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortingPhoneNumberBlock`: CreatePortingPhoneNumberBlock201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreatePortingPhoneNumberBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingOrderId** | **string** | Identifies the Porting Order associated with the phone number block | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortingPhoneNumberBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPortingPhoneNumberBlockRequest** | [**CreatePortingPhoneNumberBlockRequest**](CreatePortingPhoneNumberBlockRequest.md) |  | 

### Return type

[**CreatePortingPhoneNumberBlock201Response**](CreatePortingPhoneNumberBlock201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortingPhoneNumberExtension

> CreatePortingPhoneNumberExtension201Response CreatePortingPhoneNumberExtension(ctx, portingOrderId).CreatePortingPhoneNumberExtensionRequest(createPortingPhoneNumberExtensionRequest).Execute()

Create a phone number extension



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
	portingOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the Porting Order associated with the phone number extension
	createPortingPhoneNumberExtensionRequest := *openapiclient.NewCreatePortingPhoneNumberExtensionRequest("f24151b6-3389-41d3-8747-7dd8c681e5e2", *openapiclient.NewCreatePortingPhoneNumberExtensionRequestExtensionRange(int32(1), int32(10)), []openapiclient.CreatePortingPhoneNumberExtensionRequestActivationRangesInner{*openapiclient.NewCreatePortingPhoneNumberExtensionRequestActivationRangesInner(int32(1), int32(10))}) // CreatePortingPhoneNumberExtensionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreatePortingPhoneNumberExtension(context.Background(), portingOrderId).CreatePortingPhoneNumberExtensionRequest(createPortingPhoneNumberExtensionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreatePortingPhoneNumberExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortingPhoneNumberExtension`: CreatePortingPhoneNumberExtension201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreatePortingPhoneNumberExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingOrderId** | **string** | Identifies the Porting Order associated with the phone number extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortingPhoneNumberExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPortingPhoneNumberExtensionRequest** | [**CreatePortingPhoneNumberExtensionRequest**](CreatePortingPhoneNumberExtensionRequest.md) |  | 

### Return type

[**CreatePortingPhoneNumberExtension201Response**](CreatePortingPhoneNumberExtension201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortingReport

> CreatePortingReport201Response CreatePortingReport(ctx).CreatePortingReportRequest(createPortingReportRequest).Execute()

Create a porting related report



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
	createPortingReportRequest := *openapiclient.NewCreatePortingReportRequest("export_porting_orders_csv", *openapiclient.NewExportPortingOrdersCSVReport(*openapiclient.NewExportPortingOrdersCSVReportFilters())) // CreatePortingReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.CreatePortingReport(context.Background()).CreatePortingReportRequest(createPortingReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.CreatePortingReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortingReport`: CreatePortingReport201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.CreatePortingReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortingReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortingReportRequest** | [**CreatePortingReportRequest**](CreatePortingReportRequest.md) |  | 

### Return type

[**CreatePortingReport201Response**](CreatePortingReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdditionalDocument

> DeleteAdditionalDocument(ctx, id, additionalDocumentId).Execute()

Delete an additional document



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	additionalDocumentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Additional document identification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortingOrdersAPI.DeleteAdditionalDocument(context.Background(), id, additionalDocumentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.DeleteAdditionalDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 
**additionalDocumentId** | **string** | Additional document identification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdditionalDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoaConfiguration

> DeleteLoaConfiguration(ctx, id).Execute()

Delete a LOA configuration



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies a LOA configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortingOrdersAPI.DeleteLoaConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.DeleteLoaConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies a LOA configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortingOrder

> DeletePortingOrder(ctx, id).Execute()

Delete a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortingOrdersAPI.DeletePortingOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.DeletePortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortingPhoneNumberBlock

> CreatePortingPhoneNumberBlock201Response DeletePortingPhoneNumberBlock(ctx, portingOrderId, id).Execute()

Delete a phone number block



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
	portingOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the Porting Order associated with the phone number block
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the phone number block to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.DeletePortingPhoneNumberBlock(context.Background(), portingOrderId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.DeletePortingPhoneNumberBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortingPhoneNumberBlock`: CreatePortingPhoneNumberBlock201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.DeletePortingPhoneNumberBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingOrderId** | **string** | Identifies the Porting Order associated with the phone number block | 
**id** | **string** | Identifies the phone number block to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortingPhoneNumberBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreatePortingPhoneNumberBlock201Response**](CreatePortingPhoneNumberBlock201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortingPhoneNumberExtension

> CreatePortingPhoneNumberExtension201Response DeletePortingPhoneNumberExtension(ctx, portingOrderId, id).Execute()

Delete a phone number extension



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
	portingOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the Porting Order associated with the phone number extension
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the phone number extension to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.DeletePortingPhoneNumberExtension(context.Background(), portingOrderId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.DeletePortingPhoneNumberExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortingPhoneNumberExtension`: CreatePortingPhoneNumberExtension201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.DeletePortingPhoneNumberExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingOrderId** | **string** | Identifies the Porting Order associated with the phone number extension | 
**id** | **string** | Identifies the phone number extension to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortingPhoneNumberExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreatePortingPhoneNumberExtension201Response**](CreatePortingPhoneNumberExtension201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoaConfiguration

> CreateLoaConfiguration201Response GetLoaConfiguration(ctx, id).Execute()

Retrieve a LOA configuration



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies a LOA configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.GetLoaConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.GetLoaConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoaConfiguration`: CreateLoaConfiguration201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.GetLoaConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies a LOA configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateLoaConfiguration201Response**](CreateLoaConfiguration201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortingOrder

> GetPortingOrder200Response GetPortingOrder(ctx, id).IncludePhoneNumbers(includePhoneNumbers).Execute()

Retrieve a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	includePhoneNumbers := true // bool | Include the first 50 phone number objects in the results (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.GetPortingOrder(context.Background(), id).IncludePhoneNumbers(includePhoneNumbers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.GetPortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortingOrder`: GetPortingOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.GetPortingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePhoneNumbers** | **bool** | Include the first 50 phone number objects in the results | [default to true]

### Return type

[**GetPortingOrder200Response**](GetPortingOrder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortingOrderLoaTemplate

> *os.File GetPortingOrderLoaTemplate(ctx, id).LoaConfigurationId(loaConfigurationId).Execute()

Download a porting order loa template



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	loaConfigurationId := "a36c2277-446b-4d11-b4ea-322e02a5c08d" // string | The identifier of the LOA configuration to use for the template. If not provided, the default LOA configuration will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.GetPortingOrderLoaTemplate(context.Background(), id).LoaConfigurationId(loaConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.GetPortingOrderLoaTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortingOrderLoaTemplate`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.GetPortingOrderLoaTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingOrderLoaTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loaConfigurationId** | **string** | The identifier of the LOA configuration to use for the template. If not provided, the default LOA configuration will be used. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortingOrderSubRequest

> GetPortingOrderSubRequest200Response GetPortingOrderSubRequest(ctx, id).Execute()

Retrieve the associated V1 sub_request_id and port_request_id



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.GetPortingOrderSubRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.GetPortingOrderSubRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortingOrderSubRequest`: GetPortingOrderSubRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.GetPortingOrderSubRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingOrderSubRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPortingOrderSubRequest200Response**](GetPortingOrderSubRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortingOrdersActivationJob

> ActivatePortingOrder202Response GetPortingOrdersActivationJob(ctx, id, activationJobId).Execute()

Retrieve a porting activation job



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	activationJobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Activation Job Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.GetPortingOrdersActivationJob(context.Background(), id, activationJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.GetPortingOrdersActivationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortingOrdersActivationJob`: ActivatePortingOrder202Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.GetPortingOrdersActivationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 
**activationJobId** | **string** | Activation Job Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingOrdersActivationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivatePortingOrder202Response**](ActivatePortingOrder202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortingReport

> CreatePortingReport201Response GetPortingReport(ctx, id).Execute()

Retrieve a report



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies a report.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.GetPortingReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.GetPortingReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortingReport`: CreatePortingReport201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.GetPortingReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies a report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreatePortingReport201Response**](CreatePortingReport201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdditionalDocuments

> ListAdditionalDocuments200Response ListAdditionalDocuments(ctx, id).PageNumber(pageNumber).PageSize(pageSize).FilterDocumentType(filterDocumentType).FilterDocumentTypeIn(filterDocumentTypeIn).Sort(sort).Execute()

List additional documents



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterDocumentType := "loa" // string | Filter additional documents of a specific document type (optional)
	filterDocumentTypeIn := []string{"loa"} // []string | Filter additional documents by a list of document types (optional)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListAdditionalDocuments(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).FilterDocumentType(filterDocumentType).FilterDocumentTypeIn(filterDocumentTypeIn).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListAdditionalDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdditionalDocuments`: ListAdditionalDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListAdditionalDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAdditionalDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterDocumentType** | **string** | Filter additional documents of a specific document type | 
 **filterDocumentTypeIn** | **[]string** | Filter additional documents by a list of document types | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListAdditionalDocuments200Response**](ListAdditionalDocuments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllowedFocWindows

> ListAllowedFocWindows200Response ListAllowedFocWindows(ctx, id).Execute()

List allowed FOC dates



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListAllowedFocWindows(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListAllowedFocWindows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllowedFocWindows`: ListAllowedFocWindows200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListAllowedFocWindows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllowedFocWindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAllowedFocWindows200Response**](ListAllowedFocWindows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExceptionTypes

> ListExceptionTypes200Response ListExceptionTypes(ctx).Execute()

List all exception types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListExceptionTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListExceptionTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExceptionTypes`: ListExceptionTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListExceptionTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExceptionTypesRequest struct via the builder pattern


### Return type

[**ListExceptionTypes200Response**](ListExceptionTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoaConfigurations

> ListLoaConfigurations200Response ListLoaConfigurations(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List LOA configurations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListLoaConfigurations(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListLoaConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoaConfigurations`: ListLoaConfigurations200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListLoaConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoaConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListLoaConfigurations200Response**](ListLoaConfigurations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumberConfigurations

> ListPhoneNumberConfigurations200Response ListPhoneNumberConfigurations(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterPortingOrderId(filterPortingOrderId).FilterPortingOrderIdIn(filterPortingOrderIdIn).FilterPortingOrderStatus(filterPortingOrderStatus).FilterPortingOrderStatusIn(filterPortingOrderStatusIn).FilterPortingPhoneNumber(filterPortingPhoneNumber).FilterPortingPhoneNumberIn(filterPortingPhoneNumberIn).FilterUserBundleId(filterUserBundleId).FilterUserBundleIdIn(filterUserBundleIdIn).Sort(sort).Execute()

List all phone number configurations



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
	filterPortingOrderId := "f3575e15-32ce-400e-a4c0-dd78800c20b0" // string | Filter results by porting order id (optional)
	filterPortingOrderIdIn := []string{"f3575e15-32ce-400e-a4c0-dd78800c20b0"} // []string | Filter results by a list of porting order ids (optional)
	filterPortingOrderStatus := "foc-date-confirmed" // string | Filter results by a specific porting order status (optional)
	filterPortingOrderStatusIn := []string{"["in-process","submitted"]"} // []string | Filter results by specific porting order statuses (optional)
	filterPortingPhoneNumber := "5d6f7ede-1961-4717-bfb5-db392c5efc2d" // string | Filter results by a specific porting phone number ID (optional)
	filterPortingPhoneNumberIn := []string{"5d6f7ede-1961-4717-bfb5-db392c5efc2d"} // []string | Filter results by a list of porting phone number IDs (optional)
	filterUserBundleId := "5d6f7ede-1961-4717-bfb5-db392c5efc2d" // string | Filter results by a specific user bundle ID (optional)
	filterUserBundleIdIn := []string{"5d6f7ede-1961-4717-bfb5-db392c5efc2d"} // []string | Filter results by a list of user bundle IDs (optional)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPhoneNumberConfigurations(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterPortingOrderId(filterPortingOrderId).FilterPortingOrderIdIn(filterPortingOrderIdIn).FilterPortingOrderStatus(filterPortingOrderStatus).FilterPortingOrderStatusIn(filterPortingOrderStatusIn).FilterPortingPhoneNumber(filterPortingPhoneNumber).FilterPortingPhoneNumberIn(filterPortingPhoneNumberIn).FilterUserBundleId(filterUserBundleId).FilterUserBundleIdIn(filterUserBundleIdIn).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPhoneNumberConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPhoneNumberConfigurations`: ListPhoneNumberConfigurations200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPhoneNumberConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhoneNumberConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterPortingOrderId** | **string** | Filter results by porting order id | 
 **filterPortingOrderIdIn** | **[]string** | Filter results by a list of porting order ids | 
 **filterPortingOrderStatus** | **string** | Filter results by a specific porting order status | 
 **filterPortingOrderStatusIn** | **[]string** | Filter results by specific porting order statuses | 
 **filterPortingPhoneNumber** | **string** | Filter results by a specific porting phone number ID | 
 **filterPortingPhoneNumberIn** | **[]string** | Filter results by a list of porting phone number IDs | 
 **filterUserBundleId** | **string** | Filter results by a specific user bundle ID | 
 **filterUserBundleIdIn** | **[]string** | Filter results by a list of user bundle IDs | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumberConfigurations200Response**](ListPhoneNumberConfigurations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingEvents

> ListPortingEvents200Response ListPortingEvents(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterType(filterType).FilterPortingOrderId(filterPortingOrderId).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtLte(filterCreatedAtLte).Execute()

List all porting events



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterType := "porting_order.deleted" // string | Filter by event type. (optional)
	filterPortingOrderId := "34dc46a9-53ed-4e01-9454-26227ea13326" // string | Filter by porting order ID. (optional)
	filterCreatedAtGte := time.Now() // time.Time | Filter by created at greater than or equal to. (optional)
	filterCreatedAtLte := time.Now() // time.Time | Filter by created at less than or equal to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingEvents(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterType(filterType).FilterPortingOrderId(filterPortingOrderId).FilterCreatedAtGte(filterCreatedAtGte).FilterCreatedAtLte(filterCreatedAtLte).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingEvents`: ListPortingEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterType** | **string** | Filter by event type. | 
 **filterPortingOrderId** | **string** | Filter by porting order ID. | 
 **filterCreatedAtGte** | **time.Time** | Filter by created at greater than or equal to. | 
 **filterCreatedAtLte** | **time.Time** | Filter by created at less than or equal to. | 

### Return type

[**ListPortingEvents200Response**](ListPortingEvents200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingOrderActivationJobs

> ListPortingOrderActivationJobs200Response ListPortingOrderActivationJobs(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all porting activation jobs



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingOrderActivationJobs(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingOrderActivationJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingOrderActivationJobs`: ListPortingOrderActivationJobs200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingOrderActivationJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortingOrderActivationJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListPortingOrderActivationJobs200Response**](ListPortingOrderActivationJobs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingOrderComments

> ListPortingOrderComments200Response ListPortingOrderComments(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all comments of a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingOrderComments(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingOrderComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingOrderComments`: ListPortingOrderComments200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingOrderComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortingOrderCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListPortingOrderComments200Response**](ListPortingOrderComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingOrderRequirements

> ListPortingOrderRequirements200Response ListPortingOrderRequirements(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

List porting order requirements



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingOrderRequirements(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingOrderRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingOrderRequirements`: ListPortingOrderRequirements200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingOrderRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortingOrderRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListPortingOrderRequirements200Response**](ListPortingOrderRequirements200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingOrders

> ListPortingOrders200Response ListPortingOrders(ctx).PageNumber(pageNumber).PageSize(pageSize).IncludePhoneNumbers(includePhoneNumbers).FilterStatus(filterStatus).FilterStatusIn(filterStatusIn).FilterCustomerReference(filterCustomerReference).FilterParentSupportKey(filterParentSupportKey).FilterPhoneNumbersCountryCode(filterPhoneNumbersCountryCode).FilterPhoneNumbersCarrierName(filterPhoneNumbersCarrierName).FilterMiscType(filterMiscType).FilterEndUserAdminEntityName(filterEndUserAdminEntityName).FilterEndUserAdminAuthPersonName(filterEndUserAdminAuthPersonName).FilterActivationSettingsFastPortEligible(filterActivationSettingsFastPortEligible).FilterActivationSettingsFocDatetimeRequestedGt(filterActivationSettingsFocDatetimeRequestedGt).FilterActivationSettingsFocDatetimeRequestedLt(filterActivationSettingsFocDatetimeRequestedLt).FilterPhoneNumbersPhoneNumberContains(filterPhoneNumbersPhoneNumberContains).Sort(sort).Execute()

List all porting orders



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
	includePhoneNumbers := true // bool | Include the first 50 phone number objects in the results (optional) (default to true)
	filterStatus := "in-process" // string | Filter results by status (optional)
	filterStatusIn := "in-process" // string | Filter porting orders by multiple statuses (optional)
	filterCustomerReference := "123abc" // string | Filter results by customer_reference (optional)
	filterParentSupportKey := "123abc" // string | Filter results by parent_support_key (optional)
	filterPhoneNumbersCountryCode := "US" // string | Filter results by country ISO 3166-1 alpha-2 code (optional)
	filterPhoneNumbersCarrierName := "Telnyx" // string | Filter results by old service provider (optional)
	filterMiscType := openapiclient.PortingOrderType("full") // PortingOrderType | Filter results by porting order type (optional)
	filterEndUserAdminEntityName := "Porter McPortersen" // string | Filter results by person or company name (optional)
	filterEndUserAdminAuthPersonName := "Admin McPortersen" // string | Filter results by authorized person (optional)
	filterActivationSettingsFastPortEligible := false // bool | Filter results by fast port eligible (optional)
	filterActivationSettingsFocDatetimeRequestedGt := "2021-03-25T10:00:00.000Z" // string | Filter results by foc date later than this value (optional)
	filterActivationSettingsFocDatetimeRequestedLt := "2021-03-25T10:00:00.000Z" // string | Filter results by foc date earlier than this value (optional)
	filterPhoneNumbersPhoneNumberContains := "+13038675309" // string | Filter results by full or partial phone_number (optional)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingOrders(context.Background()).PageNumber(pageNumber).PageSize(pageSize).IncludePhoneNumbers(includePhoneNumbers).FilterStatus(filterStatus).FilterStatusIn(filterStatusIn).FilterCustomerReference(filterCustomerReference).FilterParentSupportKey(filterParentSupportKey).FilterPhoneNumbersCountryCode(filterPhoneNumbersCountryCode).FilterPhoneNumbersCarrierName(filterPhoneNumbersCarrierName).FilterMiscType(filterMiscType).FilterEndUserAdminEntityName(filterEndUserAdminEntityName).FilterEndUserAdminAuthPersonName(filterEndUserAdminAuthPersonName).FilterActivationSettingsFastPortEligible(filterActivationSettingsFastPortEligible).FilterActivationSettingsFocDatetimeRequestedGt(filterActivationSettingsFocDatetimeRequestedGt).FilterActivationSettingsFocDatetimeRequestedLt(filterActivationSettingsFocDatetimeRequestedLt).FilterPhoneNumbersPhoneNumberContains(filterPhoneNumbersPhoneNumberContains).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingOrders`: ListPortingOrders200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortingOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **includePhoneNumbers** | **bool** | Include the first 50 phone number objects in the results | [default to true]
 **filterStatus** | **string** | Filter results by status | 
 **filterStatusIn** | **string** | Filter porting orders by multiple statuses | 
 **filterCustomerReference** | **string** | Filter results by customer_reference | 
 **filterParentSupportKey** | **string** | Filter results by parent_support_key | 
 **filterPhoneNumbersCountryCode** | **string** | Filter results by country ISO 3166-1 alpha-2 code | 
 **filterPhoneNumbersCarrierName** | **string** | Filter results by old service provider | 
 **filterMiscType** | [**PortingOrderType**](PortingOrderType.md) | Filter results by porting order type | 
 **filterEndUserAdminEntityName** | **string** | Filter results by person or company name | 
 **filterEndUserAdminAuthPersonName** | **string** | Filter results by authorized person | 
 **filterActivationSettingsFastPortEligible** | **bool** | Filter results by fast port eligible | 
 **filterActivationSettingsFocDatetimeRequestedGt** | **string** | Filter results by foc date later than this value | 
 **filterActivationSettingsFocDatetimeRequestedLt** | **string** | Filter results by foc date earlier than this value | 
 **filterPhoneNumbersPhoneNumberContains** | **string** | Filter results by full or partial phone_number | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPortingOrders200Response**](ListPortingOrders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingPhoneNumberBlocks

> ListPortingPhoneNumberBlocks200Response ListPortingPhoneNumberBlocks(ctx, portingOrderId).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all phone number blocks



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
	portingOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the Porting Order associated with the phone number blocks
	filterPhoneNumber := "+12003151212" // string | Filter results by phone number (optional)
	filterPhoneNumberIn := []string{"+12003151212"} // []string | Filter results by a list of phone numbers (optional)
	sort := "false" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingPhoneNumberBlocks(context.Background(), portingOrderId).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingPhoneNumberBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingPhoneNumberBlocks`: ListPortingPhoneNumberBlocks200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingPhoneNumberBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingOrderId** | **string** | Identifies the Porting Order associated with the phone number blocks | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortingPhoneNumberBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPhoneNumber** | **string** | Filter results by phone number | 
 **filterPhoneNumberIn** | **[]string** | Filter results by a list of phone numbers | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListPortingPhoneNumberBlocks200Response**](ListPortingPhoneNumberBlocks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingPhoneNumberExtensions

> ListPortingPhoneNumberExtensions200Response ListPortingPhoneNumberExtensions(ctx, portingOrderId).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).FilterPortingPhoneNumberId(filterPortingPhoneNumberId).FilterPortingPhoneNumberIdIn(filterPortingPhoneNumberIdIn).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all phone number extensions



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
	portingOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the Porting Order associated with the phone number extensions
	filterPhoneNumber := "+12003151212" // string | Filter results by phone number (optional)
	filterPhoneNumberIn := []string{"+12003151212"} // []string | Filter results by a list of phone numbers (optional)
	filterPortingPhoneNumberId := "04f8f1b9-310c-4a3c-963e-7dfc54765140" // string | Filter results by porting phone number id (optional)
	filterPortingPhoneNumberIdIn := "04f8f1b9-310c-4a3c-963e-7dfc54765140" // string | Filter results by a list of porting phone number ids (optional)
	sort := "false" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingPhoneNumberExtensions(context.Background(), portingOrderId).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).FilterPortingPhoneNumberId(filterPortingPhoneNumberId).FilterPortingPhoneNumberIdIn(filterPortingPhoneNumberIdIn).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingPhoneNumberExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingPhoneNumberExtensions`: ListPortingPhoneNumberExtensions200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingPhoneNumberExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingOrderId** | **string** | Identifies the Porting Order associated with the phone number extensions | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortingPhoneNumberExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPhoneNumber** | **string** | Filter results by phone number | 
 **filterPhoneNumberIn** | **[]string** | Filter results by a list of phone numbers | 
 **filterPortingPhoneNumberId** | **string** | Filter results by porting phone number id | 
 **filterPortingPhoneNumberIdIn** | **string** | Filter results by a list of porting phone number ids | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListPortingPhoneNumberExtensions200Response**](ListPortingPhoneNumberExtensions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingPhoneNumbers

> ListPortingPhoneNumbers200Response ListPortingPhoneNumbers(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterPortingOrderId(filterPortingOrderId).FilterPortingOrderIdIn(filterPortingOrderIdIn).FilterSupportKeyEq(filterSupportKeyEq).FilterSupportKeyIn(filterSupportKeyIn).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).FilterPortingOrderStatus(filterPortingOrderStatus).FilterActivationStatus(filterActivationStatus).FilterPortabilityStatus(filterPortabilityStatus).Execute()

List all porting phone numbers



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
	filterPortingOrderId := "f3575e15-32ce-400e-a4c0-dd78800c20b0" // string | Filter results by porting order id (optional)
	filterPortingOrderIdIn := []string{"f3575e15-32ce-400e-a4c0-dd78800c20b0"} // []string | Filter results by a list of porting order ids (optional)
	filterSupportKeyEq := "sr_a12345" // string | Filter results by support key (optional)
	filterSupportKeyIn := []string{"sr_a12345"} // []string | Filter results by a list of support keys (optional)
	filterPhoneNumber := "+12003151212" // string | Filter results by phone number (optional)
	filterPhoneNumberIn := []string{"+12003151212"} // []string | Filter results by a list of phone numbers (optional)
	filterPortingOrderStatus := "in-process" // string | Filter results by porting order status (optional)
	filterActivationStatus := openapiclient.PortingOrderActivationStatus("New") // PortingOrderActivationStatus | Filter results by activation status (optional)
	filterPortabilityStatus := openapiclient.PortabilityStatus("pending") // PortabilityStatus | Filter results by portability status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingPhoneNumbers(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterPortingOrderId(filterPortingOrderId).FilterPortingOrderIdIn(filterPortingOrderIdIn).FilterSupportKeyEq(filterSupportKeyEq).FilterSupportKeyIn(filterSupportKeyIn).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).FilterPortingOrderStatus(filterPortingOrderStatus).FilterActivationStatus(filterActivationStatus).FilterPortabilityStatus(filterPortabilityStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingPhoneNumbers`: ListPortingPhoneNumbers200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortingPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterPortingOrderId** | **string** | Filter results by porting order id | 
 **filterPortingOrderIdIn** | **[]string** | Filter results by a list of porting order ids | 
 **filterSupportKeyEq** | **string** | Filter results by support key | 
 **filterSupportKeyIn** | **[]string** | Filter results by a list of support keys | 
 **filterPhoneNumber** | **string** | Filter results by phone number | 
 **filterPhoneNumberIn** | **[]string** | Filter results by a list of phone numbers | 
 **filterPortingOrderStatus** | **string** | Filter results by porting order status | 
 **filterActivationStatus** | [**PortingOrderActivationStatus**](PortingOrderActivationStatus.md) | Filter results by activation status | 
 **filterPortabilityStatus** | [**PortabilityStatus**](PortabilityStatus.md) | Filter results by portability status | 

### Return type

[**ListPortingPhoneNumbers200Response**](ListPortingPhoneNumbers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortingReports

> ListPortingReports200Response ListPortingReports(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterReportType(filterReportType).FilterStatus(filterStatus).Execute()

List porting related reports



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
	filterReportType := "export_porting_orders_csv" // string | Filter reports of a specific type (optional)
	filterStatus := "completed" // string | Filter reports of a specific status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListPortingReports(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterReportType(filterReportType).FilterStatus(filterStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListPortingReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortingReports`: ListPortingReports200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListPortingReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortingReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterReportType** | **string** | Filter reports of a specific type | 
 **filterStatus** | **string** | Filter reports of a specific status | 

### Return type

[**ListPortingReports200Response**](ListPortingReports200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerificationCodes

> ListVerificationCodes200Response ListVerificationCodes(ctx, id).PageNumber(pageNumber).PageSize(pageSize).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).FilterVerified(filterVerified).Sort(sort).Execute()

List verification codes



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterPhoneNumber := "+12003151212" // string | Filter results by phone number (optional)
	filterPhoneNumberIn := []string{"+12003151212"} // []string | Filter results by a list of phone numbers (optional)
	filterVerified := true // bool | Filter verification codes that have been verified or not (optional)
	sort := "created_at" // string | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ListVerificationCodes(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).FilterPhoneNumber(filterPhoneNumber).FilterPhoneNumberIn(filterPhoneNumberIn).FilterVerified(filterVerified).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ListVerificationCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVerificationCodes`: ListVerificationCodes200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ListVerificationCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterPhoneNumber** | **string** | Filter results by phone number | 
 **filterPhoneNumberIn** | **[]string** | Filter results by a list of phone numbers | 
 **filterVerified** | **bool** | Filter verification codes that have been verified or not | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListVerificationCodes200Response**](ListVerificationCodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewLoaConfiguration

> *os.File PreviewLoaConfiguration(ctx, id).Execute()

Preview a LOA configuration



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies a LOA configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.PreviewLoaConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.PreviewLoaConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewLoaConfiguration`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.PreviewLoaConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies a LOA configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewLoaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewLoaConfigurationParams

> *os.File PreviewLoaConfigurationParams(ctx).PreviewLoaConfigurationParamsRequest(previewLoaConfigurationParamsRequest).Execute()

Preview the LOA configuration parameters



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
	previewLoaConfigurationParamsRequest := *openapiclient.NewPreviewLoaConfigurationParamsRequest("My LOA Configuration", *openapiclient.NewPreviewLoaConfigurationParamsRequestLogo("DocumentId_example"), "Telnyx", *openapiclient.NewPreviewLoaConfigurationParamsRequestAddress("600 Congress Avenue", "US"), *openapiclient.NewPreviewLoaConfigurationParamsRequestContact("testing@telnyx.com", "+12003270001")) // PreviewLoaConfigurationParamsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.PreviewLoaConfigurationParams(context.Background()).PreviewLoaConfigurationParamsRequest(previewLoaConfigurationParamsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.PreviewLoaConfigurationParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewLoaConfigurationParams`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.PreviewLoaConfigurationParams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewLoaConfigurationParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewLoaConfigurationParamsRequest** | [**PreviewLoaConfigurationParamsRequest**](PreviewLoaConfigurationParamsRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepublishPortingEvent

> RepublishPortingEvent(ctx, id).Execute()

Republish a porting event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the porting event.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortingOrdersAPI.RepublishPortingEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.RepublishPortingEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the porting event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepublishPortingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPortingVerificationCodes

> SendPortingVerificationCodes(ctx, id).SendPortingVerificationCodesRequest(sendPortingVerificationCodesRequest).Execute()

Send the verification codes



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	sendPortingVerificationCodesRequest := *openapiclient.NewSendPortingVerificationCodesRequest() // SendPortingVerificationCodesRequest | A list of phone numbers to send the verification codes to and the method to send them by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortingOrdersAPI.SendPortingVerificationCodes(context.Background(), id).SendPortingVerificationCodesRequest(sendPortingVerificationCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.SendPortingVerificationCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendPortingVerificationCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendPortingVerificationCodesRequest** | [**SendPortingVerificationCodesRequest**](SendPortingVerificationCodesRequest.md) | A list of phone numbers to send the verification codes to and the method to send them by | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharePortingOrder

> SharePortingOrder201Response SharePortingOrder(ctx, id).SharePortingOrderRequest(sharePortingOrderRequest).Execute()

Share a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	sharePortingOrderRequest := *openapiclient.NewSharePortingOrderRequest() // SharePortingOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.SharePortingOrder(context.Background(), id).SharePortingOrderRequest(sharePortingOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.SharePortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharePortingOrder`: SharePortingOrder201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.SharePortingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharePortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharePortingOrderRequest** | [**SharePortingOrderRequest**](SharePortingOrderRequest.md) |  | 

### Return type

[**SharePortingOrder201Response**](SharePortingOrder201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPortingEvent

> ShowPortingEvent200Response ShowPortingEvent(ctx, id).Execute()

Show a porting event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies the porting event.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.ShowPortingEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.ShowPortingEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPortingEvent`: ShowPortingEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.ShowPortingEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the porting event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPortingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShowPortingEvent200Response**](ShowPortingEvent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoaConfiguration

> CreateLoaConfiguration201Response UpdateLoaConfiguration(ctx, id).PreviewLoaConfigurationParamsRequest(previewLoaConfigurationParamsRequest).Execute()

Update a LOA configuration



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifies a LOA configuration.
	previewLoaConfigurationParamsRequest := *openapiclient.NewPreviewLoaConfigurationParamsRequest("My LOA Configuration", *openapiclient.NewPreviewLoaConfigurationParamsRequestLogo("DocumentId_example"), "Telnyx", *openapiclient.NewPreviewLoaConfigurationParamsRequestAddress("600 Congress Avenue", "US"), *openapiclient.NewPreviewLoaConfigurationParamsRequestContact("testing@telnyx.com", "+12003270001")) // PreviewLoaConfigurationParamsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.UpdateLoaConfiguration(context.Background(), id).PreviewLoaConfigurationParamsRequest(previewLoaConfigurationParamsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.UpdateLoaConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoaConfiguration`: CreateLoaConfiguration201Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.UpdateLoaConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies a LOA configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewLoaConfigurationParamsRequest** | [**PreviewLoaConfigurationParamsRequest**](PreviewLoaConfigurationParamsRequest.md) |  | 

### Return type

[**CreateLoaConfiguration201Response**](CreateLoaConfiguration201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePortingOrder

> UpdatePortingOrder200Response UpdatePortingOrder(ctx, id).UpdatePortingOrder(updatePortingOrder).Execute()

Edit a porting order



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	updatePortingOrder := *openapiclient.NewUpdatePortingOrder() // UpdatePortingOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.UpdatePortingOrder(context.Background(), id).UpdatePortingOrder(updatePortingOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.UpdatePortingOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortingOrder`: UpdatePortingOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.UpdatePortingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePortingOrder** | [**UpdatePortingOrder**](UpdatePortingOrder.md) |  | 

### Return type

[**UpdatePortingOrder200Response**](UpdatePortingOrder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePortingOrdersActivationJob

> ActivatePortingOrder202Response UpdatePortingOrdersActivationJob(ctx, id, activationJobId).UpdatePortingOrdersActivationJobRequest(updatePortingOrdersActivationJobRequest).Execute()

Update a porting activation job



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	activationJobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Activation Job Identifier
	updatePortingOrdersActivationJobRequest := *openapiclient.NewUpdatePortingOrdersActivationJobRequest() // UpdatePortingOrdersActivationJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.UpdatePortingOrdersActivationJob(context.Background(), id, activationJobId).UpdatePortingOrdersActivationJobRequest(updatePortingOrdersActivationJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.UpdatePortingOrdersActivationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortingOrdersActivationJob`: ActivatePortingOrder202Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.UpdatePortingOrdersActivationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 
**activationJobId** | **string** | Activation Job Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortingOrdersActivationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePortingOrdersActivationJobRequest** | [**UpdatePortingOrdersActivationJobRequest**](UpdatePortingOrdersActivationJobRequest.md) |  | 

### Return type

[**ActivatePortingOrder202Response**](ActivatePortingOrder202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPortingVerificationCodes

> VerifyPortingVerificationCodes200Response VerifyPortingVerificationCodes(ctx, id).VerifyPortingVerificationCodesRequest(verifyPortingVerificationCodesRequest).Execute()

Verify the verification code for a list of phone numbers



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Porting Order id
	verifyPortingVerificationCodesRequest := *openapiclient.NewVerifyPortingVerificationCodesRequest() // VerifyPortingVerificationCodesRequest | A list of phone numbers and their verification codes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortingOrdersAPI.VerifyPortingVerificationCodes(context.Background(), id).VerifyPortingVerificationCodesRequest(verifyPortingVerificationCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortingOrdersAPI.VerifyPortingVerificationCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPortingVerificationCodes`: VerifyPortingVerificationCodes200Response
	fmt.Fprintf(os.Stdout, "Response from `PortingOrdersAPI.VerifyPortingVerificationCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Porting Order id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPortingVerificationCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyPortingVerificationCodesRequest** | [**VerifyPortingVerificationCodesRequest**](VerifyPortingVerificationCodesRequest.md) | A list of phone numbers and their verification codes | 

### Return type

[**VerifyPortingVerificationCodes200Response**](VerifyPortingVerificationCodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

