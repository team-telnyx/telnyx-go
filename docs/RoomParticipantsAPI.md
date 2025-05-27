# \RoomParticipantsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRoomParticipants**](RoomParticipantsAPI.md#ListRoomParticipants) | **Get** /room_participants | View a list of room participants.
[**ViewRoomParticipant**](RoomParticipantsAPI.md#ViewRoomParticipant) | **Get** /room_participants/{room_participant_id} | View a room participant.



## ListRoomParticipants

> ListRoomParticipants200Response ListRoomParticipants(ctx).FilterDateJoinedAtEq(filterDateJoinedAtEq).FilterDateJoinedAtGte(filterDateJoinedAtGte).FilterDateJoinedAtLte(filterDateJoinedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateLeftAtEq(filterDateLeftAtEq).FilterDateLeftAtGte(filterDateLeftAtGte).FilterDateLeftAtLte(filterDateLeftAtLte).FilterContext(filterContext).FilterSessionId(filterSessionId).PageSize(pageSize).PageNumber(pageNumber).Execute()

View a list of room participants.



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
	filterDateJoinedAtEq := time.Now() // string | ISO 8601 date for filtering room participants that joined on that date. (optional)
	filterDateJoinedAtGte := time.Now() // string | ISO 8601 date for filtering room participants that joined after that date. (optional)
	filterDateJoinedAtLte := time.Now() // string | ISO 8601 date for filtering room participants that joined before that date. (optional)
	filterDateUpdatedAtEq := time.Now() // string | ISO 8601 date for filtering room participants updated on that date. (optional)
	filterDateUpdatedAtGte := time.Now() // string | ISO 8601 date for filtering room participants updated after that date. (optional)
	filterDateUpdatedAtLte := time.Now() // string | ISO 8601 date for filtering room participants updated before that date. (optional)
	filterDateLeftAtEq := time.Now() // string | ISO 8601 date for filtering room participants that left on that date. (optional)
	filterDateLeftAtGte := time.Now() // string | ISO 8601 date for filtering room participants that left after that date. (optional)
	filterDateLeftAtLte := time.Now() // string | ISO 8601 date for filtering room participants that left before that date. (optional)
	filterContext := "Alice" // string | Filter room participants based on the context. (optional)
	filterSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | Session_id for filtering room participants. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomParticipantsAPI.ListRoomParticipants(context.Background()).FilterDateJoinedAtEq(filterDateJoinedAtEq).FilterDateJoinedAtGte(filterDateJoinedAtGte).FilterDateJoinedAtLte(filterDateJoinedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateLeftAtEq(filterDateLeftAtEq).FilterDateLeftAtGte(filterDateLeftAtGte).FilterDateLeftAtLte(filterDateLeftAtLte).FilterContext(filterContext).FilterSessionId(filterSessionId).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomParticipantsAPI.ListRoomParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoomParticipants`: ListRoomParticipants200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomParticipantsAPI.ListRoomParticipants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoomParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateJoinedAtEq** | **string** | ISO 8601 date for filtering room participants that joined on that date. | 
 **filterDateJoinedAtGte** | **string** | ISO 8601 date for filtering room participants that joined after that date. | 
 **filterDateJoinedAtLte** | **string** | ISO 8601 date for filtering room participants that joined before that date. | 
 **filterDateUpdatedAtEq** | **string** | ISO 8601 date for filtering room participants updated on that date. | 
 **filterDateUpdatedAtGte** | **string** | ISO 8601 date for filtering room participants updated after that date. | 
 **filterDateUpdatedAtLte** | **string** | ISO 8601 date for filtering room participants updated before that date. | 
 **filterDateLeftAtEq** | **string** | ISO 8601 date for filtering room participants that left on that date. | 
 **filterDateLeftAtGte** | **string** | ISO 8601 date for filtering room participants that left after that date. | 
 **filterDateLeftAtLte** | **string** | ISO 8601 date for filtering room participants that left before that date. | 
 **filterContext** | **string** | Filter room participants based on the context. | 
 **filterSessionId** | **string** | Session_id for filtering room participants. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**ListRoomParticipants200Response**](ListRoomParticipants200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRoomParticipant

> ViewRoomParticipant200Response ViewRoomParticipant(ctx, roomParticipantId).Execute()

View a room participant.

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
	roomParticipantId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room participant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomParticipantsAPI.ViewRoomParticipant(context.Background(), roomParticipantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomParticipantsAPI.ViewRoomParticipant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewRoomParticipant`: ViewRoomParticipant200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomParticipantsAPI.ViewRoomParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomParticipantId** | **string** | The unique identifier of a room participant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRoomParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ViewRoomParticipant200Response**](ViewRoomParticipant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

