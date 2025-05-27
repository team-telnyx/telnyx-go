# \RoomsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoom**](RoomsAPI.md#CreateRoom) | **Post** /rooms | Create a room.
[**DeleteRoom**](RoomsAPI.md#DeleteRoom) | **Delete** /rooms/{room_id} | Delete a room.
[**ListRooms**](RoomsAPI.md#ListRooms) | **Get** /rooms | View a list of rooms.
[**RetrieveListRoomSessions**](RoomsAPI.md#RetrieveListRoomSessions) | **Get** /rooms/{room_id}/sessions | View a list of room sessions.
[**UpdateRoom**](RoomsAPI.md#UpdateRoom) | **Patch** /rooms/{room_id} | Update a room.
[**ViewRoom**](RoomsAPI.md#ViewRoom) | **Get** /rooms/{room_id} | View a room.



## CreateRoom

> CreateRoom201Response CreateRoom(ctx).CreateRoomRequest(createRoomRequest).Execute()

Create a room.



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
	createRoomRequest := *openapiclient.NewCreateRoomRequest() // CreateRoomRequest | Parameters that can be defined during room creation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoom(context.Background()).CreateRoomRequest(createRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoom`: CreateRoom201Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoomRequest** | [**CreateRoomRequest**](CreateRoomRequest.md) | Parameters that can be defined during room creation. | 

### Return type

[**CreateRoom201Response**](CreateRoom201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoom

> DeleteRoom(ctx, roomId).Execute()

Delete a room.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAPI.DeleteRoom(context.Background(), roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | The unique identifier of a room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomRequest struct via the builder pattern


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


## ListRooms

> ListRooms200Response ListRooms(ctx).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterUniqueName(filterUniqueName).IncludeSessions(includeSessions).PageSize(pageSize).PageNumber(pageNumber).Execute()

View a list of rooms.



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
	filterDateCreatedAtEq := time.Now() // string | ISO 8601 date for filtering rooms created on that date. (optional)
	filterDateCreatedAtGte := time.Now() // string | ISO 8601 date for filtering rooms created after that date. (optional)
	filterDateCreatedAtLte := time.Now() // string | ISO 8601 date for filtering rooms created before that date. (optional)
	filterDateUpdatedAtEq := time.Now() // string | ISO 8601 date for filtering rooms updated on that date. (optional)
	filterDateUpdatedAtGte := time.Now() // string | ISO 8601 date for filtering rooms updated after that date. (optional)
	filterDateUpdatedAtLte := time.Now() // string | ISO 8601 date for filtering rooms updated before that date. (optional)
	filterUniqueName := "my_video_room" // string | Unique_name for filtering rooms. (optional)
	includeSessions := true // bool | To decide if room sessions should be included in the response. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ListRooms(context.Background()).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterUniqueName(filterUniqueName).IncludeSessions(includeSessions).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ListRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRooms`: ListRooms200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ListRooms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateCreatedAtEq** | **string** | ISO 8601 date for filtering rooms created on that date. | 
 **filterDateCreatedAtGte** | **string** | ISO 8601 date for filtering rooms created after that date. | 
 **filterDateCreatedAtLte** | **string** | ISO 8601 date for filtering rooms created before that date. | 
 **filterDateUpdatedAtEq** | **string** | ISO 8601 date for filtering rooms updated on that date. | 
 **filterDateUpdatedAtGte** | **string** | ISO 8601 date for filtering rooms updated after that date. | 
 **filterDateUpdatedAtLte** | **string** | ISO 8601 date for filtering rooms updated before that date. | 
 **filterUniqueName** | **string** | Unique_name for filtering rooms. | 
 **includeSessions** | **bool** | To decide if room sessions should be included in the response. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**ListRooms200Response**](ListRooms200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListRoomSessions

> ListRoomSessions200Response RetrieveListRoomSessions(ctx, roomId).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterActive(filterActive).IncludeParticipants(includeParticipants).PageSize(pageSize).PageNumber(pageNumber).Execute()

View a list of room sessions.



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
	roomId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room.
	filterDateCreatedAtEq := time.Now() // string | ISO 8601 date for filtering room sessions created on that date. (optional)
	filterDateCreatedAtGte := time.Now() // string | ISO 8601 date for filtering room sessions created after that date. (optional)
	filterDateCreatedAtLte := time.Now() // string | ISO 8601 date for filtering room sessions created before that date. (optional)
	filterDateUpdatedAtEq := time.Now() // string | ISO 8601 date for filtering room sessions updated on that date. (optional)
	filterDateUpdatedAtGte := time.Now() // string | ISO 8601 date for filtering room sessions updated after that date. (optional)
	filterDateUpdatedAtLte := time.Now() // string | ISO 8601 date for filtering room sessions updated before that date. (optional)
	filterDateEndedAtEq := time.Now() // string | ISO 8601 date for filtering room sessions ended on that date. (optional)
	filterDateEndedAtGte := time.Now() // string | ISO 8601 date for filtering room sessions ended after that date. (optional)
	filterDateEndedAtLte := time.Now() // string | ISO 8601 date for filtering room sessions ended before that date. (optional)
	filterActive := true // bool | Filter active or inactive room sessions. (optional)
	includeParticipants := true // bool | To decide if room participants should be included in the response. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.RetrieveListRoomSessions(context.Background(), roomId).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterActive(filterActive).IncludeParticipants(includeParticipants).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.RetrieveListRoomSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveListRoomSessions`: ListRoomSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.RetrieveListRoomSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | The unique identifier of a room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListRoomSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDateCreatedAtEq** | **string** | ISO 8601 date for filtering room sessions created on that date. | 
 **filterDateCreatedAtGte** | **string** | ISO 8601 date for filtering room sessions created after that date. | 
 **filterDateCreatedAtLte** | **string** | ISO 8601 date for filtering room sessions created before that date. | 
 **filterDateUpdatedAtEq** | **string** | ISO 8601 date for filtering room sessions updated on that date. | 
 **filterDateUpdatedAtGte** | **string** | ISO 8601 date for filtering room sessions updated after that date. | 
 **filterDateUpdatedAtLte** | **string** | ISO 8601 date for filtering room sessions updated before that date. | 
 **filterDateEndedAtEq** | **string** | ISO 8601 date for filtering room sessions ended on that date. | 
 **filterDateEndedAtGte** | **string** | ISO 8601 date for filtering room sessions ended after that date. | 
 **filterDateEndedAtLte** | **string** | ISO 8601 date for filtering room sessions ended before that date. | 
 **filterActive** | **bool** | Filter active or inactive room sessions. | 
 **includeParticipants** | **bool** | To decide if room participants should be included in the response. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**ListRoomSessions200Response**](ListRoomSessions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> CreateRoom201Response UpdateRoom(ctx, roomId).PatchRoomRequest(patchRoomRequest).Execute()

Update a room.



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
	patchRoomRequest := *openapiclient.NewPatchRoomRequest() // PatchRoomRequest | Parameters that can be defined during room update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UpdateRoom(context.Background(), roomId).PatchRoomRequest(patchRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UpdateRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoom`: CreateRoom201Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UpdateRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | The unique identifier of a room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchRoomRequest** | [**PatchRoomRequest**](PatchRoomRequest.md) | Parameters that can be defined during room update. | 

### Return type

[**CreateRoom201Response**](CreateRoom201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRoom

> CreateRoom201Response ViewRoom(ctx, roomId).IncludeSessions(includeSessions).Execute()

View a room.

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
	includeSessions := true // bool | To decide if room sessions should be included in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ViewRoom(context.Background(), roomId).IncludeSessions(includeSessions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ViewRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewRoom`: CreateRoom201Response
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ViewRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | The unique identifier of a room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSessions** | **bool** | To decide if room sessions should be included in the response. | 

### Return type

[**CreateRoom201Response**](CreateRoom201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

