# \SIPRECConnectorsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiprecConnector**](SIPRECConnectorsAPI.md#CreateSiprecConnector) | **Post** /siprec_connectors | Create a SIPREC connector
[**DeleteSiprecConnector**](SIPRECConnectorsAPI.md#DeleteSiprecConnector) | **Delete** /siprec_connectors/{connector_name} | Delete a SIPREC connector
[**GetSiprecConnector**](SIPRECConnectorsAPI.md#GetSiprecConnector) | **Get** /siprec_connectors/{connector_name} | Retrieve a SIPREC connector
[**UpdateSiprecConnector**](SIPRECConnectorsAPI.md#UpdateSiprecConnector) | **Put** /siprec_connectors/{connector_name} | Update a SIPREC connector



## CreateSiprecConnector

> SiprecConnectorResponse CreateSiprecConnector(ctx).SiprecConnector(siprecConnector).Execute()

Create a SIPREC connector



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
	siprecConnector := *openapiclient.NewSiprecConnector("siprec.telnyx.com", int32(5060), "my-siprec-connector") // SiprecConnector | The parameters required to create or update a SIPREC connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIPRECConnectorsAPI.CreateSiprecConnector(context.Background()).SiprecConnector(siprecConnector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIPRECConnectorsAPI.CreateSiprecConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiprecConnector`: SiprecConnectorResponse
	fmt.Fprintf(os.Stdout, "Response from `SIPRECConnectorsAPI.CreateSiprecConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiprecConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siprecConnector** | [**SiprecConnector**](SiprecConnector.md) | The parameters required to create or update a SIPREC connector. | 

### Return type

[**SiprecConnectorResponse**](SiprecConnectorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiprecConnector

> DeleteSiprecConnector(ctx, connectorName).Execute()

Delete a SIPREC connector



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
	connectorName := "connectorName_example" // string | Uniquely identifies a SIPREC connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SIPRECConnectorsAPI.DeleteSiprecConnector(context.Background(), connectorName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIPRECConnectorsAPI.DeleteSiprecConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorName** | **string** | Uniquely identifies a SIPREC connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiprecConnectorRequest struct via the builder pattern


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


## GetSiprecConnector

> SiprecConnectorResponse GetSiprecConnector(ctx, connectorName).Execute()

Retrieve a SIPREC connector



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
	connectorName := "connectorName_example" // string | Uniquely identifies a SIPREC connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIPRECConnectorsAPI.GetSiprecConnector(context.Background(), connectorName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIPRECConnectorsAPI.GetSiprecConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiprecConnector`: SiprecConnectorResponse
	fmt.Fprintf(os.Stdout, "Response from `SIPRECConnectorsAPI.GetSiprecConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorName** | **string** | Uniquely identifies a SIPREC connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiprecConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiprecConnectorResponse**](SiprecConnectorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiprecConnector

> SiprecConnectorResponse UpdateSiprecConnector(ctx, connectorName).SiprecConnector(siprecConnector).Execute()

Update a SIPREC connector



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
	connectorName := "connectorName_example" // string | Uniquely identifies a SIPREC connector.
	siprecConnector := *openapiclient.NewSiprecConnector("siprec.telnyx.com", int32(5060), "my-siprec-connector") // SiprecConnector | The parameters required to create or update a SIPREC connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIPRECConnectorsAPI.UpdateSiprecConnector(context.Background(), connectorName).SiprecConnector(siprecConnector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIPRECConnectorsAPI.UpdateSiprecConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiprecConnector`: SiprecConnectorResponse
	fmt.Fprintf(os.Stdout, "Response from `SIPRECConnectorsAPI.UpdateSiprecConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorName** | **string** | Uniquely identifies a SIPREC connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiprecConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siprecConnector** | [**SiprecConnector**](SiprecConnector.md) | The parameters required to create or update a SIPREC connector. | 

### Return type

[**SiprecConnectorResponse**](SiprecConnectorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

