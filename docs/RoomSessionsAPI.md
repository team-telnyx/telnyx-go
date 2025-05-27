# \RoomSessionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndSession**](RoomSessionsAPI.md#EndSession) | **Post** /room_sessions/{room_session_id}/actions/end | End a room session.
[**KickParticipantInSession**](RoomSessionsAPI.md#KickParticipantInSession) | **Post** /room_sessions/{room_session_id}/actions/kick | Kick participants from a room session.
[**ListRoomSessions**](RoomSessionsAPI.md#ListRoomSessions) | **Get** /room_sessions | View a list of room sessions.
[**MuteParticipantInSession**](RoomSessionsAPI.md#MuteParticipantInSession) | **Post** /room_sessions/{room_session_id}/actions/mute | Mute participants in room session.
[**RetrieveListRoomParticipants**](RoomSessionsAPI.md#RetrieveListRoomParticipants) | **Get** /room_sessions/{room_session_id}/participants | View a list of room participants.
[**UnmuteParticipantInSession**](RoomSessionsAPI.md#UnmuteParticipantInSession) | **Post** /room_sessions/{room_session_id}/actions/unmute | Unmute participants in room session.
[**ViewRoomSession**](RoomSessionsAPI.md#ViewRoomSession) | **Get** /room_sessions/{room_session_id} | View a room session.



## EndSession

> EndSession200Response EndSession(ctx, roomSessionId).Execute()

End a room session.



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
	roomSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room session.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.EndSession(context.Background(), roomSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.EndSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndSession`: EndSession200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.EndSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSessionId** | **string** | The unique identifier of a room session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndSession200Response**](EndSession200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KickParticipantInSession

> EndSession200Response KickParticipantInSession(ctx, roomSessionId).ActionsParticipantsRequest(actionsParticipantsRequest).Execute()

Kick participants from a room session.



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
	roomSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room session.
	actionsParticipantsRequest := *openapiclient.NewActionsParticipantsRequest() // ActionsParticipantsRequest | Parameters that can be defined during Kick action.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.KickParticipantInSession(context.Background(), roomSessionId).ActionsParticipantsRequest(actionsParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.KickParticipantInSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KickParticipantInSession`: EndSession200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.KickParticipantInSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSessionId** | **string** | The unique identifier of a room session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKickParticipantInSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsParticipantsRequest** | [**ActionsParticipantsRequest**](ActionsParticipantsRequest.md) | Parameters that can be defined during Kick action. | 

### Return type

[**EndSession200Response**](EndSession200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomSessions

> ListRoomSessions200Response ListRoomSessions(ctx).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterRoomId(filterRoomId).FilterActive(filterActive).IncludeParticipants(includeParticipants).PageSize(pageSize).PageNumber(pageNumber).Execute()

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
	filterDateCreatedAtEq := time.Now() // string | ISO 8601 date for filtering room sessions created on that date. (optional)
	filterDateCreatedAtGte := time.Now() // string | ISO 8601 date for filtering room sessions created after that date. (optional)
	filterDateCreatedAtLte := time.Now() // string | ISO 8601 date for filtering room sessions created before that date. (optional)
	filterDateUpdatedAtEq := time.Now() // string | ISO 8601 date for filtering room sessions updated on that date. (optional)
	filterDateUpdatedAtGte := time.Now() // string | ISO 8601 date for filtering room sessions updated after that date. (optional)
	filterDateUpdatedAtLte := time.Now() // string | ISO 8601 date for filtering room sessions updated before that date. (optional)
	filterDateEndedAtEq := time.Now() // string | ISO 8601 date for filtering room sessions ended on that date. (optional)
	filterDateEndedAtGte := time.Now() // string | ISO 8601 date for filtering room sessions ended after that date. (optional)
	filterDateEndedAtLte := time.Now() // string | ISO 8601 date for filtering room sessions ended before that date. (optional)
	filterRoomId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | Room_id for filtering room sessions. (optional)
	filterActive := true // bool | Filter active or inactive room sessions. (optional)
	includeParticipants := true // bool | To decide if room participants should be included in the response. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.ListRoomSessions(context.Background()).FilterDateCreatedAtEq(filterDateCreatedAtEq).FilterDateCreatedAtGte(filterDateCreatedAtGte).FilterDateCreatedAtLte(filterDateCreatedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterRoomId(filterRoomId).FilterActive(filterActive).IncludeParticipants(includeParticipants).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.ListRoomSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoomSessions`: ListRoomSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.ListRoomSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoomSessionsRequest struct via the builder pattern


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
 **filterRoomId** | **string** | Room_id for filtering room sessions. | 
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


## MuteParticipantInSession

> EndSession200Response MuteParticipantInSession(ctx, roomSessionId).ActionsParticipantsRequest(actionsParticipantsRequest).Execute()

Mute participants in room session.



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
	roomSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room session.
	actionsParticipantsRequest := *openapiclient.NewActionsParticipantsRequest() // ActionsParticipantsRequest | Parameters that can be defined during Mute action.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.MuteParticipantInSession(context.Background(), roomSessionId).ActionsParticipantsRequest(actionsParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.MuteParticipantInSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteParticipantInSession`: EndSession200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.MuteParticipantInSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSessionId** | **string** | The unique identifier of a room session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteParticipantInSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsParticipantsRequest** | [**ActionsParticipantsRequest**](ActionsParticipantsRequest.md) | Parameters that can be defined during Mute action. | 

### Return type

[**EndSession200Response**](EndSession200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListRoomParticipants

> ListRoomParticipants200Response RetrieveListRoomParticipants(ctx, roomSessionId).FilterDateJoinedAtEq(filterDateJoinedAtEq).FilterDateJoinedAtGte(filterDateJoinedAtGte).FilterDateJoinedAtLte(filterDateJoinedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateLeftAtEq(filterDateLeftAtEq).FilterDateLeftAtGte(filterDateLeftAtGte).FilterDateLeftAtLte(filterDateLeftAtLte).FilterContext(filterContext).PageSize(pageSize).PageNumber(pageNumber).Execute()

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
	roomSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room session.
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
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.RetrieveListRoomParticipants(context.Background(), roomSessionId).FilterDateJoinedAtEq(filterDateJoinedAtEq).FilterDateJoinedAtGte(filterDateJoinedAtGte).FilterDateJoinedAtLte(filterDateJoinedAtLte).FilterDateUpdatedAtEq(filterDateUpdatedAtEq).FilterDateUpdatedAtGte(filterDateUpdatedAtGte).FilterDateUpdatedAtLte(filterDateUpdatedAtLte).FilterDateLeftAtEq(filterDateLeftAtEq).FilterDateLeftAtGte(filterDateLeftAtGte).FilterDateLeftAtLte(filterDateLeftAtLte).FilterContext(filterContext).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.RetrieveListRoomParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveListRoomParticipants`: ListRoomParticipants200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.RetrieveListRoomParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSessionId** | **string** | The unique identifier of a room session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListRoomParticipantsRequest struct via the builder pattern


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


## UnmuteParticipantInSession

> EndSession200Response UnmuteParticipantInSession(ctx, roomSessionId).ActionsParticipantsRequest(actionsParticipantsRequest).Execute()

Unmute participants in room session.



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
	roomSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room session.
	actionsParticipantsRequest := *openapiclient.NewActionsParticipantsRequest() // ActionsParticipantsRequest | Parameters that can be defined during Unmute action.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.UnmuteParticipantInSession(context.Background(), roomSessionId).ActionsParticipantsRequest(actionsParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.UnmuteParticipantInSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmuteParticipantInSession`: EndSession200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.UnmuteParticipantInSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSessionId** | **string** | The unique identifier of a room session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteParticipantInSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsParticipantsRequest** | [**ActionsParticipantsRequest**](ActionsParticipantsRequest.md) | Parameters that can be defined during Unmute action. | 

### Return type

[**EndSession200Response**](EndSession200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRoomSession

> ViewRoomSession200Response ViewRoomSession(ctx, roomSessionId).IncludeParticipants(includeParticipants).Execute()

View a room session.

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
	roomSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room session.
	includeParticipants := true // bool | To decide if room participants should be included in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomSessionsAPI.ViewRoomSession(context.Background(), roomSessionId).IncludeParticipants(includeParticipants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomSessionsAPI.ViewRoomSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewRoomSession`: ViewRoomSession200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomSessionsAPI.ViewRoomSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSessionId** | **string** | The unique identifier of a room session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRoomSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeParticipants** | **bool** | To decide if room participants should be included in the response. | 

### Return type

[**ViewRoomSession200Response**](ViewRoomSession200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

