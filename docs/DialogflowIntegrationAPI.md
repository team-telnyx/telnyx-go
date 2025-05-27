# \DialogflowIntegrationAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDialogflowConnection**](DialogflowIntegrationAPI.md#CreateDialogflowConnection) | **Post** /dialogflow_connections/{connection_id} | Create a Dialogflow Connection
[**DeleteDialogflowConnection**](DialogflowIntegrationAPI.md#DeleteDialogflowConnection) | **Delete** /dialogflow_connections/{connection_id} | Delete stored Dialogflow Connection
[**GetDialogflowConnection**](DialogflowIntegrationAPI.md#GetDialogflowConnection) | **Get** /dialogflow_connections/{connection_id} | Retrieve stored Dialogflow Connection
[**UpdateDialogflowConnection**](DialogflowIntegrationAPI.md#UpdateDialogflowConnection) | **Put** /dialogflow_connections/{connection_id} | Update stored Dialogflow Connection



## CreateDialogflowConnection

> DialogflowConnectionResponse CreateDialogflowConnection(ctx, connectionId).DialogflowConnection(dialogflowConnection).Execute()

Create a Dialogflow Connection



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
	dialogflowConnection := *openapiclient.NewDialogflowConnection() // DialogflowConnection | The params expected to create/update a Dialogflow Connection for given connection_id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DialogflowIntegrationAPI.CreateDialogflowConnection(context.Background(), connectionId).DialogflowConnection(dialogflowConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DialogflowIntegrationAPI.CreateDialogflowConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDialogflowConnection`: DialogflowConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DialogflowIntegrationAPI.CreateDialogflowConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDialogflowConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dialogflowConnection** | [**DialogflowConnection**](DialogflowConnection.md) | The params expected to create/update a Dialogflow Connection for given connection_id. | 

### Return type

[**DialogflowConnectionResponse**](DialogflowConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDialogflowConnection

> DeleteDialogflowConnection(ctx, connectionId).Execute()

Delete stored Dialogflow Connection



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DialogflowIntegrationAPI.DeleteDialogflowConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DialogflowIntegrationAPI.DeleteDialogflowConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDialogflowConnectionRequest struct via the builder pattern


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


## GetDialogflowConnection

> DialogflowConnectionResponse GetDialogflowConnection(ctx, connectionId).Execute()

Retrieve stored Dialogflow Connection



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DialogflowIntegrationAPI.GetDialogflowConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DialogflowIntegrationAPI.GetDialogflowConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDialogflowConnection`: DialogflowConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DialogflowIntegrationAPI.GetDialogflowConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDialogflowConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DialogflowConnectionResponse**](DialogflowConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDialogflowConnection

> DialogflowConnectionResponse UpdateDialogflowConnection(ctx, connectionId).DialogflowConnection(dialogflowConnection).Execute()

Update stored Dialogflow Connection



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
	dialogflowConnection := *openapiclient.NewDialogflowConnection() // DialogflowConnection | The params expected to create/update a Dialogflow Connection for given connection_id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DialogflowIntegrationAPI.UpdateDialogflowConnection(context.Background(), connectionId).DialogflowConnection(dialogflowConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DialogflowIntegrationAPI.UpdateDialogflowConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDialogflowConnection`: DialogflowConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DialogflowIntegrationAPI.UpdateDialogflowConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDialogflowConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dialogflowConnection** | [**DialogflowConnection**](DialogflowConnection.md) | The params expected to create/update a Dialogflow Connection for given connection_id. | 

### Return type

[**DialogflowConnectionResponse**](DialogflowConnectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

