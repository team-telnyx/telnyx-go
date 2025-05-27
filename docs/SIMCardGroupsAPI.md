# \SIMCardGroupsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSimCardGroup**](SIMCardGroupsAPI.md#CreateSimCardGroup) | **Post** /sim_card_groups | Create a SIM card group
[**DeleteSimCardGroup**](SIMCardGroupsAPI.md#DeleteSimCardGroup) | **Delete** /sim_card_groups/{id} | Delete a SIM card group
[**GetAllSimCardGroups**](SIMCardGroupsAPI.md#GetAllSimCardGroups) | **Get** /sim_card_groups | Get all SIM card groups
[**GetSimCardGroup**](SIMCardGroupsAPI.md#GetSimCardGroup) | **Get** /sim_card_groups/{id} | Get SIM card group
[**RemoveSimCardGroupPrivateWirelessGateway**](SIMCardGroupsAPI.md#RemoveSimCardGroupPrivateWirelessGateway) | **Post** /sim_card_groups/{id}/actions/remove_private_wireless_gateway | Request Private Wireless Gateway removal from SIM card group
[**SetPrivateWirelessGatewayForSimCardGroup**](SIMCardGroupsAPI.md#SetPrivateWirelessGatewayForSimCardGroup) | **Post** /sim_card_groups/{id}/actions/set_private_wireless_gateway | Request Private Wireless Gateway assignment for SIM card group
[**UpdateSimCardGroup**](SIMCardGroupsAPI.md#UpdateSimCardGroup) | **Patch** /sim_card_groups/{id} | Update a SIM card group



## CreateSimCardGroup

> CreateSimCardGroup200Response CreateSimCardGroup(ctx).SIMCardGroupCreate(sIMCardGroupCreate).Execute()

Create a SIM card group



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
	sIMCardGroupCreate := *openapiclient.NewSIMCardGroupCreate("My Test Group") // SIMCardGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.CreateSimCardGroup(context.Background()).SIMCardGroupCreate(sIMCardGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.CreateSimCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimCardGroup`: CreateSimCardGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.CreateSimCardGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sIMCardGroupCreate** | [**SIMCardGroupCreate**](SIMCardGroupCreate.md) |  | 

### Return type

[**CreateSimCardGroup200Response**](CreateSimCardGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSimCardGroup

> CreateSimCardGroup200Response DeleteSimCardGroup(ctx, id).Execute()

Delete a SIM card group



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.DeleteSimCardGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.DeleteSimCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSimCardGroup`: CreateSimCardGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.DeleteSimCardGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSimCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSimCardGroup200Response**](CreateSimCardGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSimCardGroups

> GetAllSimCardGroups200Response GetAllSimCardGroups(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).FilterPrivateWirelessGatewayId(filterPrivateWirelessGatewayId).Execute()

Get all SIM card groups



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
	filterName := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A valid SIM card group name. (optional)
	filterPrivateWirelessGatewayId := "7606c6d3-ff7c-49c1-943d-68879e9d584d" // string | A Private Wireless Gateway ID associated with the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.GetAllSimCardGroups(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterName(filterName).FilterPrivateWirelessGatewayId(filterPrivateWirelessGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.GetAllSimCardGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSimCardGroups`: GetAllSimCardGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.GetAllSimCardGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSimCardGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterName** | **string** | A valid SIM card group name. | 
 **filterPrivateWirelessGatewayId** | **string** | A Private Wireless Gateway ID associated with the group. | 

### Return type

[**GetAllSimCardGroups200Response**](GetAllSimCardGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardGroup

> CreateSimCardGroup200Response GetSimCardGroup(ctx, id).IncludeIccids(includeIccids).Execute()

Get SIM card group



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM group.
	includeIccids := true // bool | It includes a list of associated ICCIDs. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.GetSimCardGroup(context.Background(), id).IncludeIccids(includeIccids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.GetSimCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardGroup`: CreateSimCardGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.GetSimCardGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeIccids** | **bool** | It includes a list of associated ICCIDs. | [default to false]

### Return type

[**CreateSimCardGroup200Response**](CreateSimCardGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSimCardGroupPrivateWirelessGateway

> GetSimCardGroupAction200Response RemoveSimCardGroupPrivateWirelessGateway(ctx, id).Execute()

Request Private Wireless Gateway removal from SIM card group



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.RemoveSimCardGroupPrivateWirelessGateway(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.RemoveSimCardGroupPrivateWirelessGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSimCardGroupPrivateWirelessGateway`: GetSimCardGroupAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.RemoveSimCardGroupPrivateWirelessGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSimCardGroupPrivateWirelessGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSimCardGroupAction200Response**](GetSimCardGroupAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPrivateWirelessGatewayForSimCardGroup

> GetSimCardGroupAction200Response SetPrivateWirelessGatewayForSimCardGroup(ctx, id).SetPrivateWirelessGatewayForSimCardGroupRequest(setPrivateWirelessGatewayForSimCardGroupRequest).Execute()

Request Private Wireless Gateway assignment for SIM card group



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM group.
	setPrivateWirelessGatewayForSimCardGroupRequest := *openapiclient.NewSetPrivateWirelessGatewayForSimCardGroupRequest("6a09cdc3-8948-47f0-aa62-74ac943d6c58") // SetPrivateWirelessGatewayForSimCardGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.SetPrivateWirelessGatewayForSimCardGroup(context.Background(), id).SetPrivateWirelessGatewayForSimCardGroupRequest(setPrivateWirelessGatewayForSimCardGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.SetPrivateWirelessGatewayForSimCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrivateWirelessGatewayForSimCardGroup`: GetSimCardGroupAction200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.SetPrivateWirelessGatewayForSimCardGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPrivateWirelessGatewayForSimCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setPrivateWirelessGatewayForSimCardGroupRequest** | [**SetPrivateWirelessGatewayForSimCardGroupRequest**](SetPrivateWirelessGatewayForSimCardGroupRequest.md) |  | 

### Return type

[**GetSimCardGroupAction200Response**](GetSimCardGroupAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSimCardGroup

> CreateSimCardGroup200Response UpdateSimCardGroup(ctx, id).SIMCardGroupPatch(sIMCardGroupPatch).Execute()

Update a SIM card group



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the SIM group.
	sIMCardGroupPatch := *openapiclient.NewSIMCardGroupPatch() // SIMCardGroupPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMCardGroupsAPI.UpdateSimCardGroup(context.Background(), id).SIMCardGroupPatch(sIMCardGroupPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMCardGroupsAPI.UpdateSimCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSimCardGroup`: CreateSimCardGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `SIMCardGroupsAPI.UpdateSimCardGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the SIM group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSimCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sIMCardGroupPatch** | [**SIMCardGroupPatch**](SIMCardGroupPatch.md) |  | 

### Return type

[**CreateSimCardGroup200Response**](CreateSimCardGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

