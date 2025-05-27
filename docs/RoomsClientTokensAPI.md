# \RoomsClientTokensAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoomClientToken**](RoomsClientTokensAPI.md#CreateRoomClientToken) | **Post** /rooms/{room_id}/actions/generate_join_client_token | Create Client Token to join a room.
[**RefreshRoomClientToken**](RoomsClientTokensAPI.md#RefreshRoomClientToken) | **Post** /rooms/{room_id}/actions/refresh_client_token | Refresh Client Token to join a room.



## CreateRoomClientToken

> CreateRoomClientToken201Response CreateRoomClientToken(ctx, roomId).CreateRoomClientTokenRequest(createRoomClientTokenRequest).Execute()

Create Client Token to join a room.



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
	roomId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room.
	createRoomClientTokenRequest := *openapiclient.NewCreateRoomClientTokenRequest() // CreateRoomClientTokenRequest | Parameters that can be defined during Room Client Token creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsClientTokensAPI.CreateRoomClientToken(context.Background(), roomId).CreateRoomClientTokenRequest(createRoomClientTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsClientTokensAPI.CreateRoomClientToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomClientToken`: CreateRoomClientToken201Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsClientTokensAPI.CreateRoomClientToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | The unique identifier of a room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomClientTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoomClientTokenRequest** | [**CreateRoomClientTokenRequest**](CreateRoomClientTokenRequest.md) | Parameters that can be defined during Room Client Token creation. | 

### Return type

[**CreateRoomClientToken201Response**](CreateRoomClientToken201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshRoomClientToken

> RefreshRoomClientToken201Response RefreshRoomClientToken(ctx, roomId).RefreshRoomClientTokenRequest(refreshRoomClientTokenRequest).Execute()

Refresh Client Token to join a room.



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
	roomId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room.
	refreshRoomClientTokenRequest := *openapiclient.NewRefreshRoomClientTokenRequest("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZWxueXhfdGVsZXBob255IiwiZXhwIjoxNTkwMDEwMTQzLCJpYXQiOjE1ODc1OTA5NDMsImlzcyI6InRlbG55eF90ZWxlcGhvbnkiLCJqdGkiOiJiOGM3NDgzNy1kODllLTRhNjUtOWNmMi0zNGM3YTZmYTYwYzgiLCJuYmYiOjE1ODc1OTA5NDIsInN1YiI6IjVjN2FjN2QwLWRiNjUtNGYxMS05OGUxLWVlYzBkMWQ1YzZhZSIsInRlbF90b2tlbiI6InJqX1pra1pVT1pNeFpPZk9tTHBFVUIzc2lVN3U2UmpaRmVNOXMtZ2JfeENSNTZXRktGQUppTXlGMlQ2Q0JSbWxoX1N5MGlfbGZ5VDlBSThzRWlmOE1USUlzenl6U2xfYURuRzQ4YU81MHlhSEd1UlNZYlViU1ltOVdJaVEwZz09IiwidHlwIjoiYWNjZXNzIn0.gNEwzTow5MLLPLQENytca7pUN79PmPj6FyqZWW06ZeEmesxYpwKh0xRtA0TzLh6CDYIRHrI8seofOO0YFGDhpQ") // RefreshRoomClientTokenRequest | Parameters that can be defined during Room Client Token refresh.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsClientTokensAPI.RefreshRoomClientToken(context.Background(), roomId).RefreshRoomClientTokenRequest(refreshRoomClientTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsClientTokensAPI.RefreshRoomClientToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshRoomClientToken`: RefreshRoomClientToken201Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsClientTokensAPI.RefreshRoomClientToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | The unique identifier of a room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRoomClientTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshRoomClientTokenRequest** | [**RefreshRoomClientTokenRequest**](RefreshRoomClientTokenRequest.md) | Parameters that can be defined during Room Client Token refresh. | 

### Return type

[**RefreshRoomClientToken201Response**](RefreshRoomClientToken201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

