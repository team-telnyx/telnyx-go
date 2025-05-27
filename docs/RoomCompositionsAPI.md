# \RoomCompositionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoomComposition**](RoomCompositionsAPI.md#CreateRoomComposition) | **Post** /room_compositions | Create a room composition.
[**DeleteRoomComposition**](RoomCompositionsAPI.md#DeleteRoomComposition) | **Delete** /room_compositions/{room_composition_id} | Delete a room composition.
[**ListRoomCompositions**](RoomCompositionsAPI.md#ListRoomCompositions) | **Get** /room_compositions | View a list of room compositions.
[**ViewRoomComposition**](RoomCompositionsAPI.md#ViewRoomComposition) | **Get** /room_compositions/{room_composition_id} | View a room composition.



## CreateRoomComposition

> CreateRoomComposition202Response CreateRoomComposition(ctx).CreateRoomCompositionRequest(createRoomCompositionRequest).Execute()

Create a room composition.



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
	createRoomCompositionRequest := *openapiclient.NewCreateRoomCompositionRequest() // CreateRoomCompositionRequest | Parameters that can be defined during room composition creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomCompositionsAPI.CreateRoomComposition(context.Background()).CreateRoomCompositionRequest(createRoomCompositionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomCompositionsAPI.CreateRoomComposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomComposition`: CreateRoomComposition202Response
	fmt.Fprintf(os.Stdout, "Response from `RoomCompositionsAPI.CreateRoomComposition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoomCompositionRequest** | [**CreateRoomCompositionRequest**](CreateRoomCompositionRequest.md) | Parameters that can be defined during room composition creation. | 

### Return type

[**CreateRoomComposition202Response**](CreateRoomComposition202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomComposition

> DeleteRoomComposition(ctx, roomCompositionId).Execute()

Delete a room composition.



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
	roomCompositionId := "5219b3af-87c6-4c08-9b58-5a533d893e21" // string | The unique identifier of a room composition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomCompositionsAPI.DeleteRoomComposition(context.Background(), roomCompositionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomCompositionsAPI.DeleteRoomComposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomCompositionId** | **string** | The unique identifier of a room composition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomCompositionRequest struct via the builder pattern


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


## ListRoomCompositions

> ListRoomCompositions200Response ListRoomCompositions(ctx).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterSessionId(filterSessionId).FilterStatus(filterStatus).PageSize(pageSize).PageNumber(pageNumber).Execute()

View a list of room compositions.



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
	filterDateCreatedAtEq := time.Now() // string | ISO 8601 date for filtering room compositions created on that date. (optional)
	filterDateCreatedAtGte := time.Now() // string | ISO 8601 date for filtering room compositions created after that date. (optional)
	filterDateCreatedAtLte := time.Now() // string | ISO 8601 date for filtering room compositions created before that date. (optional)
	filterSessionId := "92e7d459-bcc5-4386-9f5f-6dd14a82588d" // string | The session_id for filtering room compositions. (optional)
	filterStatus := "completed" // string | The status for filtering room compositions. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomCompositionsAPI.ListRoomCompositions(context.Background()).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterSessionId(filterSessionId).FilterStatus(filterStatus).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomCompositionsAPI.ListRoomCompositions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoomCompositions`: ListRoomCompositions200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomCompositionsAPI.ListRoomCompositions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoomCompositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateCreatedAtEq** | **string** | ISO 8601 date for filtering room compositions created on that date. | 
 **filterDateCreatedAtGte** | **string** | ISO 8601 date for filtering room compositions created after that date. | 
 **filterDateCreatedAtLte** | **string** | ISO 8601 date for filtering room compositions created before that date. | 
 **filterSessionId** | **string** | The session_id for filtering room compositions. | 
 **filterStatus** | **string** | The status for filtering room compositions. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**ListRoomCompositions200Response**](ListRoomCompositions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRoomComposition

> CreateRoomComposition202Response ViewRoomComposition(ctx, roomCompositionId).Execute()

View a room composition.

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
	roomCompositionId := "5219b3af-87c6-4c08-9b58-5a533d893e21" // string | The unique identifier of a room composition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomCompositionsAPI.ViewRoomComposition(context.Background(), roomCompositionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomCompositionsAPI.ViewRoomComposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewRoomComposition`: CreateRoomComposition202Response
	fmt.Fprintf(os.Stdout, "Response from `RoomCompositionsAPI.ViewRoomComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomCompositionId** | **string** | The unique identifier of a room composition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRoomCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateRoomComposition202Response**](CreateRoomComposition202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

