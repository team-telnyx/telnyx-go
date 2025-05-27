# \RoomRecordingsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoomRecording**](RoomRecordingsAPI.md#DeleteRoomRecording) | **Delete** /room_recordings/{room_recording_id} | Delete a room recording.
[**DeleteRoomRecordings**](RoomRecordingsAPI.md#DeleteRoomRecordings) | **Delete** /room_recordings | Delete several room recordings in a bulk.
[**ListRoomRecordings**](RoomRecordingsAPI.md#ListRoomRecordings) | **Get** /room_recordings | View a list of room recordings.
[**ViewRoomRecording**](RoomRecordingsAPI.md#ViewRoomRecording) | **Get** /room_recordings/{room_recording_id} | View a room recording.



## DeleteRoomRecording

> DeleteRoomRecording(ctx, roomRecordingId).Execute()

Delete a room recording.



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
	roomRecordingId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room recording.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomRecordingsAPI.DeleteRoomRecording(context.Background(), roomRecordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomRecordingsAPI.DeleteRoomRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomRecordingId** | **string** | The unique identifier of a room recording. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomRecordingRequest struct via the builder pattern


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


## DeleteRoomRecordings

> BulkRoomRecordingsDeleteResponse DeleteRoomRecordings(ctx).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterDateStartedAtEq(filterDateStartedAtEq).FilterDateStartedAtGte(filterDateStartedAtGte).FilterDateStartedAtLte(filterDateStartedAtLte).FilterRoomId(filterRoomId).FilterParticipantId(filterParticipantId).FilterSessionId(filterSessionId).FilterStatus(filterStatus).FilterType(filterType).FilterDurationSecsEq(filterDurationSecsEq).FilterDurationSecsLte(filterDurationSecsLte).FilterDurationSecsGte(filterDurationSecsGte).PageSize(pageSize).PageNumber(pageNumber).Execute()

Delete several room recordings in a bulk.



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
	filterDateEndedAtEq := time.Now() // string | ISO 8601 date for filtering room recordings ended on that date. (optional)
	filterDateEndedAtGte := time.Now() // string | ISO 8601 date for filtering room recordings ended after that date. (optional)
	filterDateEndedAtLte := time.Now() // string | ISO 8601 date for filtering room recordings ended before that date. (optional)
	filterDateStartedAtEq := time.Now() // string | ISO 8601 date for filtering room recordings started on that date. (optional)
	filterDateStartedAtGte := time.Now() // string | ISO 8601 date for filtering room recordings started after that date. (optional)
	filterDateStartedAtLte := time.Now() // string | ISO 8601 date for filtering room recordings started before that date. (optional)
	filterRoomId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | room_id for filtering room recordings. (optional)
	filterParticipantId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | participant_id for filtering room recordings. (optional)
	filterSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | session_id for filtering room recordings. (optional)
	filterStatus := "completed" // string | status for filtering room recordings. (optional)
	filterType := "audio" // string | type for filtering room recordings. (optional)
	filterDurationSecsEq := int32(20) // int32 | duration_secs equal for filtering room recordings. (optional)
	filterDurationSecsLte := int32(20) // int32 | duration_secs less or equal for filtering room recordings. (optional)
	filterDurationSecsGte := int32(20) // int32 | duration_secs greater or equal for filtering room recordings. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomRecordingsAPI.DeleteRoomRecordings(context.Background()).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterDateStartedAtEq(filterDateStartedAtEq).FilterDateStartedAtGte(filterDateStartedAtGte).FilterDateStartedAtLte(filterDateStartedAtLte).FilterRoomId(filterRoomId).FilterParticipantId(filterParticipantId).FilterSessionId(filterSessionId).FilterStatus(filterStatus).FilterType(filterType).FilterDurationSecsEq(filterDurationSecsEq).FilterDurationSecsLte(filterDurationSecsLte).FilterDurationSecsGte(filterDurationSecsGte).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomRecordingsAPI.DeleteRoomRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoomRecordings`: BulkRoomRecordingsDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoomRecordingsAPI.DeleteRoomRecordings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateEndedAtEq** | **string** | ISO 8601 date for filtering room recordings ended on that date. | 
 **filterDateEndedAtGte** | **string** | ISO 8601 date for filtering room recordings ended after that date. | 
 **filterDateEndedAtLte** | **string** | ISO 8601 date for filtering room recordings ended before that date. | 
 **filterDateStartedAtEq** | **string** | ISO 8601 date for filtering room recordings started on that date. | 
 **filterDateStartedAtGte** | **string** | ISO 8601 date for filtering room recordings started after that date. | 
 **filterDateStartedAtLte** | **string** | ISO 8601 date for filtering room recordings started before that date. | 
 **filterRoomId** | **string** | room_id for filtering room recordings. | 
 **filterParticipantId** | **string** | participant_id for filtering room recordings. | 
 **filterSessionId** | **string** | session_id for filtering room recordings. | 
 **filterStatus** | **string** | status for filtering room recordings. | 
 **filterType** | **string** | type for filtering room recordings. | 
 **filterDurationSecsEq** | **int32** | duration_secs equal for filtering room recordings. | 
 **filterDurationSecsLte** | **int32** | duration_secs less or equal for filtering room recordings. | 
 **filterDurationSecsGte** | **int32** | duration_secs greater or equal for filtering room recordings. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**BulkRoomRecordingsDeleteResponse**](BulkRoomRecordingsDeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomRecordings

> ListRoomRecordings200Response ListRoomRecordings(ctx).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterDateStartedAtEq(filterDateStartedAtEq).FilterDateStartedAtGte(filterDateStartedAtGte).FilterDateStartedAtLte(filterDateStartedAtLte).FilterRoomId(filterRoomId).FilterParticipantId(filterParticipantId).FilterSessionId(filterSessionId).FilterStatus(filterStatus).FilterType(filterType).FilterDurationSecsEq(filterDurationSecsEq).FilterDurationSecsLte(filterDurationSecsLte).FilterDurationSecsGte(filterDurationSecsGte).PageSize(pageSize).PageNumber(pageNumber).Execute()

View a list of room recordings.



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
	filterDateEndedAtEq := time.Now() // string | ISO 8601 date for filtering room recordings ended on that date. (optional)
	filterDateEndedAtGte := time.Now() // string | ISO 8601 date for filtering room recordings ended after that date. (optional)
	filterDateEndedAtLte := time.Now() // string | ISO 8601 date for filtering room recordings ended before that date. (optional)
	filterDateStartedAtEq := time.Now() // string | ISO 8601 date for filtering room recordings started on that date. (optional)
	filterDateStartedAtGte := time.Now() // string | ISO 8601 date for filtering room recordings started after that date. (optional)
	filterDateStartedAtLte := time.Now() // string | ISO 8601 date for filtering room recordings started before that date. (optional)
	filterRoomId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | room_id for filtering room recordings. (optional)
	filterParticipantId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | participant_id for filtering room recordings. (optional)
	filterSessionId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | session_id for filtering room recordings. (optional)
	filterStatus := "completed" // string | status for filtering room recordings. (optional)
	filterType := "audio" // string | type for filtering room recordings. (optional)
	filterDurationSecsEq := int32(20) // int32 | duration_secs equal for filtering room recordings. (optional)
	filterDurationSecsLte := int32(20) // int32 | duration_secs less or equal for filtering room recordings. (optional)
	filterDurationSecsGte := int32(20) // int32 | duration_secs greater or equal for filtering room recordings. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomRecordingsAPI.ListRoomRecordings(context.Background()).FilterDateEndedAtEq(filterDateEndedAtEq).FilterDateEndedAtGte(filterDateEndedAtGte).FilterDateEndedAtLte(filterDateEndedAtLte).FilterDateStartedAtEq(filterDateStartedAtEq).FilterDateStartedAtGte(filterDateStartedAtGte).FilterDateStartedAtLte(filterDateStartedAtLte).FilterRoomId(filterRoomId).FilterParticipantId(filterParticipantId).FilterSessionId(filterSessionId).FilterStatus(filterStatus).FilterType(filterType).FilterDurationSecsEq(filterDurationSecsEq).FilterDurationSecsLte(filterDurationSecsLte).FilterDurationSecsGte(filterDurationSecsGte).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomRecordingsAPI.ListRoomRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoomRecordings`: ListRoomRecordings200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomRecordingsAPI.ListRoomRecordings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoomRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateEndedAtEq** | **string** | ISO 8601 date for filtering room recordings ended on that date. | 
 **filterDateEndedAtGte** | **string** | ISO 8601 date for filtering room recordings ended after that date. | 
 **filterDateEndedAtLte** | **string** | ISO 8601 date for filtering room recordings ended before that date. | 
 **filterDateStartedAtEq** | **string** | ISO 8601 date for filtering room recordings started on that date. | 
 **filterDateStartedAtGte** | **string** | ISO 8601 date for filtering room recordings started after that date. | 
 **filterDateStartedAtLte** | **string** | ISO 8601 date for filtering room recordings started before that date. | 
 **filterRoomId** | **string** | room_id for filtering room recordings. | 
 **filterParticipantId** | **string** | participant_id for filtering room recordings. | 
 **filterSessionId** | **string** | session_id for filtering room recordings. | 
 **filterStatus** | **string** | status for filtering room recordings. | 
 **filterType** | **string** | type for filtering room recordings. | 
 **filterDurationSecsEq** | **int32** | duration_secs equal for filtering room recordings. | 
 **filterDurationSecsLte** | **int32** | duration_secs less or equal for filtering room recordings. | 
 **filterDurationSecsGte** | **int32** | duration_secs greater or equal for filtering room recordings. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageNumber** | **int32** | The page number to load. | [default to 1]

### Return type

[**ListRoomRecordings200Response**](ListRoomRecordings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRoomRecording

> ViewRoomRecording200Response ViewRoomRecording(ctx, roomRecordingId).Execute()

View a room recording.

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
	roomRecordingId := "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0" // string | The unique identifier of a room recording.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomRecordingsAPI.ViewRoomRecording(context.Background(), roomRecordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomRecordingsAPI.ViewRoomRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewRoomRecording`: ViewRoomRecording200Response
	fmt.Fprintf(os.Stdout, "Response from `RoomRecordingsAPI.ViewRoomRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomRecordingId** | **string** | The unique identifier of a room recording. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRoomRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ViewRoomRecording200Response**](ViewRoomRecording200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

