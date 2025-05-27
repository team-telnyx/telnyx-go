# \IntegrationSecretsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIntegrationSecret**](IntegrationSecretsAPI.md#CreateIntegrationSecret) | **Post** /integration_secrets | Create a secret
[**DeleteIntegrationSecret**](IntegrationSecretsAPI.md#DeleteIntegrationSecret) | **Delete** /integration_secrets/{id} | Delete an integration secret
[**ListIntegrationSecrets**](IntegrationSecretsAPI.md#ListIntegrationSecrets) | **Get** /integration_secrets | List integration secrets



## CreateIntegrationSecret

> IntegrationSecretCreatedResponse CreateIntegrationSecret(ctx).CreateIntegrationSecretRequest(createIntegrationSecretRequest).Execute()

Create a secret



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
	createIntegrationSecretRequest := *openapiclient.NewCreateIntegrationSecretRequest("Identifier_example", "Type_example") // CreateIntegrationSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationSecretsAPI.CreateIntegrationSecret(context.Background()).CreateIntegrationSecretRequest(createIntegrationSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationSecretsAPI.CreateIntegrationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIntegrationSecret`: IntegrationSecretCreatedResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationSecretsAPI.CreateIntegrationSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIntegrationSecretRequest** | [**CreateIntegrationSecretRequest**](CreateIntegrationSecretRequest.md) |  | 

### Return type

[**IntegrationSecretCreatedResponse**](IntegrationSecretCreatedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationSecret

> DeleteIntegrationSecret(ctx, id).Execute()

Delete an integration secret



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationSecretsAPI.DeleteIntegrationSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationSecretsAPI.DeleteIntegrationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationSecretRequest struct via the builder pattern


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


## ListIntegrationSecrets

> IntegrationSecretsListData ListIntegrationSecrets(ctx).PageSize(pageSize).PageNumber(pageNumber).FilterType(filterType).Execute()

List integration secrets



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
	pageSize := int32(25) // int32 |  (optional)
	pageNumber := int32(1) // int32 |  (optional)
	filterType := "bearer" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationSecretsAPI.ListIntegrationSecrets(context.Background()).PageSize(pageSize).PageNumber(pageNumber).FilterType(filterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationSecretsAPI.ListIntegrationSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationSecrets`: IntegrationSecretsListData
	fmt.Fprintf(os.Stdout, "Response from `IntegrationSecretsAPI.ListIntegrationSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | 
 **pageNumber** | **int32** |  | 
 **filterType** | **string** |  | 

### Return type

[**IntegrationSecretsListData**](IntegrationSecretsListData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

