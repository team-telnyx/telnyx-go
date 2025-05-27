# \ConferenceCommandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConference**](ConferenceCommandsAPI.md#CreateConference) | **Post** /conferences | Create conference
[**HoldConferenceParticipants**](ConferenceCommandsAPI.md#HoldConferenceParticipants) | **Post** /conferences/{id}/actions/hold | Hold conference participants
[**JoinConference**](ConferenceCommandsAPI.md#JoinConference) | **Post** /conferences/{id}/actions/join | Join a conference
[**LeaveConference**](ConferenceCommandsAPI.md#LeaveConference) | **Post** /conferences/{id}/actions/leave | Leave a conference
[**ListConferenceParticipants**](ConferenceCommandsAPI.md#ListConferenceParticipants) | **Get** /conferences/{conference_id}/participants | List conference participants
[**ListConferences**](ConferenceCommandsAPI.md#ListConferences) | **Get** /conferences | List conferences
[**MuteConferenceParticipants**](ConferenceCommandsAPI.md#MuteConferenceParticipants) | **Post** /conferences/{id}/actions/mute | Mute conference participants
[**PauseConferenceRecording**](ConferenceCommandsAPI.md#PauseConferenceRecording) | **Post** /conferences/{id}/actions/record_pause | Conference recording pause
[**PlayConferenceAudio**](ConferenceCommandsAPI.md#PlayConferenceAudio) | **Post** /conferences/{id}/actions/play | Play audio to conference participants
[**ResumeConferenceRecording**](ConferenceCommandsAPI.md#ResumeConferenceRecording) | **Post** /conferences/{id}/actions/record_resume | Conference recording resume
[**RetrieveConference**](ConferenceCommandsAPI.md#RetrieveConference) | **Get** /conferences/{id} | Retrieve a conference
[**SpeakTextToConference**](ConferenceCommandsAPI.md#SpeakTextToConference) | **Post** /conferences/{id}/actions/speak | Speak text to conference participants
[**StartConferenceRecording**](ConferenceCommandsAPI.md#StartConferenceRecording) | **Post** /conferences/{id}/actions/record_start | Conference recording start
[**StopConferenceAudio**](ConferenceCommandsAPI.md#StopConferenceAudio) | **Post** /conferences/{id}/actions/stop | Stop audio being played on the conference
[**StopConferenceRecording**](ConferenceCommandsAPI.md#StopConferenceRecording) | **Post** /conferences/{id}/actions/record_stop | Conference recording stop
[**UnholdConferenceParticipants**](ConferenceCommandsAPI.md#UnholdConferenceParticipants) | **Post** /conferences/{id}/actions/unhold | Unhold conference participants
[**UnmuteConferenceParticipants**](ConferenceCommandsAPI.md#UnmuteConferenceParticipants) | **Post** /conferences/{id}/actions/unmute | Unmute conference participants
[**UpdateConference**](ConferenceCommandsAPI.md#UpdateConference) | **Post** /conferences/{id}/actions/update | Update conference participant



## CreateConference

> ConferenceResponse CreateConference(ctx).CreateConferenceRequest(createConferenceRequest).Execute()

Create conference



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
	createConferenceRequest := *openapiclient.NewCreateConferenceRequest("v2:T02llQxIyaRkhfRKxgAP8nY511EhFLizdvdUKJiSw8d6A9BborherQczRrZvZakpWxBlpw48KyZQ==", "Business") // CreateConferenceRequest | Create a conference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.CreateConference(context.Background()).CreateConferenceRequest(createConferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.CreateConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConference`: ConferenceResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.CreateConference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConferenceRequest** | [**CreateConferenceRequest**](CreateConferenceRequest.md) | Create a conference | 

### Return type

[**ConferenceResponse**](ConferenceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HoldConferenceParticipants

> ConferenceCommandResponse HoldConferenceParticipants(ctx, id).ConferenceHoldRequest(conferenceHoldRequest).Execute()

Hold conference participants



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	conferenceHoldRequest := *openapiclient.NewConferenceHoldRequest() // ConferenceHoldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.HoldConferenceParticipants(context.Background(), id).ConferenceHoldRequest(conferenceHoldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.HoldConferenceParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HoldConferenceParticipants`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.HoldConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiHoldConferenceParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferenceHoldRequest** | [**ConferenceHoldRequest**](ConferenceHoldRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinConference

> ConferenceCommandResponse JoinConference(ctx, id).JoinConferenceRequest(joinConferenceRequest).Execute()

Join a conference



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	joinConferenceRequest := *openapiclient.NewJoinConferenceRequest("v2:T02llQxIyaRkhfRKxgAP8nY511EhFLizdvdUKJiSw8d6A9BborherQczRrZvZakpWxBlpw48KyZQ==") // JoinConferenceRequest | Join Conference request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.JoinConference(context.Background(), id).JoinConferenceRequest(joinConferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.JoinConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JoinConference`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.JoinConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **joinConferenceRequest** | [**JoinConferenceRequest**](JoinConferenceRequest.md) | Join Conference request object | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveConference

> ConferenceCommandResponse LeaveConference(ctx, id).LeaveConferenceRequest(leaveConferenceRequest).Execute()

Leave a conference



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	leaveConferenceRequest := *openapiclient.NewLeaveConferenceRequest("f91269aa-61d1-417f-97b3-10e020e8bc47") // LeaveConferenceRequest | Leave Conference request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.LeaveConference(context.Background(), id).LeaveConferenceRequest(leaveConferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.LeaveConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaveConference`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.LeaveConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leaveConferenceRequest** | [**LeaveConferenceRequest**](LeaveConferenceRequest.md) | Leave Conference request object | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferenceParticipants

> ListParticipantsResponse ListConferenceParticipants(ctx, conferenceId).FilterMuted(filterMuted).FilterOnHold(filterOnHold).FilterWhispering(filterWhispering).PageNumber(pageNumber).PageSize(pageSize).Execute()

List conference participants



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
	conferenceId := "conferenceId_example" // string | Uniquely identifies the conference by id
	filterMuted := true // bool | If present, participants will be filtered to those who are/are not muted (optional)
	filterOnHold := true // bool | If present, participants will be filtered to those who are/are not put on hold (optional)
	filterWhispering := true // bool | If present, participants will be filtered to those who are whispering or are not (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.ListConferenceParticipants(context.Background(), conferenceId).FilterMuted(filterMuted).FilterOnHold(filterOnHold).FilterWhispering(filterWhispering).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.ListConferenceParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConferenceParticipants`: ListParticipantsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.ListConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Uniquely identifies the conference by id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConferenceParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterMuted** | **bool** | If present, participants will be filtered to those who are/are not muted | 
 **filterOnHold** | **bool** | If present, participants will be filtered to those who are/are not put on hold | 
 **filterWhispering** | **bool** | If present, participants will be filtered to those who are whispering or are not | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListParticipantsResponse**](ListParticipantsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferences

> ListConferencesResponse ListConferences(ctx).FilterName(filterName).FilterStatus(filterStatus).PageNumber(pageNumber).PageSize(pageSize).Execute()

List conferences



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
	filterName := "filterName_example" // string | If present, conferences will be filtered to those with a matching `name` attribute. Matching is case-sensitive (optional)
	filterStatus := "filterStatus_example" // string | If present, conferences will be filtered by status. (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.ListConferences(context.Background()).FilterName(filterName).FilterStatus(filterStatus).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.ListConferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConferences`: ListConferencesResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.ListConferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string** | If present, conferences will be filtered to those with a matching &#x60;name&#x60; attribute. Matching is case-sensitive | 
 **filterStatus** | **string** | If present, conferences will be filtered by status. | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListConferencesResponse**](ListConferencesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteConferenceParticipants

> ConferenceCommandResponse MuteConferenceParticipants(ctx, id).ConferenceMuteRequest(conferenceMuteRequest).Execute()

Mute conference participants



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	conferenceMuteRequest := *openapiclient.NewConferenceMuteRequest() // ConferenceMuteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.MuteConferenceParticipants(context.Background(), id).ConferenceMuteRequest(conferenceMuteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.MuteConferenceParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteConferenceParticipants`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.MuteConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteConferenceParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferenceMuteRequest** | [**ConferenceMuteRequest**](ConferenceMuteRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseConferenceRecording

> ConferenceCommandResponse PauseConferenceRecording(ctx, id).PauseConferenceRecordingRequest(pauseConferenceRecordingRequest).Execute()

Conference recording pause



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
	id := "id_example" // string | Specifies the conference by id or name
	pauseConferenceRecordingRequest := *openapiclient.NewPauseConferenceRecordingRequest() // PauseConferenceRecordingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.PauseConferenceRecording(context.Background(), id).PauseConferenceRecordingRequest(pauseConferenceRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.PauseConferenceRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseConferenceRecording`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.PauseConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pauseConferenceRecordingRequest** | [**PauseConferenceRecordingRequest**](PauseConferenceRecordingRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayConferenceAudio

> ConferenceCommandResponse PlayConferenceAudio(ctx, id).ConferencePlayRequest(conferencePlayRequest).Execute()

Play audio to conference participants



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	conferencePlayRequest := *openapiclient.NewConferencePlayRequest() // ConferencePlayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.PlayConferenceAudio(context.Background(), id).ConferencePlayRequest(conferencePlayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.PlayConferenceAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlayConferenceAudio`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.PlayConferenceAudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayConferenceAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferencePlayRequest** | [**ConferencePlayRequest**](ConferencePlayRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeConferenceRecording

> ConferenceCommandResponse ResumeConferenceRecording(ctx, id).ResumeConferenceRecordingRequest(resumeConferenceRecordingRequest).Execute()

Conference recording resume



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
	id := "id_example" // string | Specifies the conference by id or name
	resumeConferenceRecordingRequest := *openapiclient.NewResumeConferenceRecordingRequest() // ResumeConferenceRecordingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.ResumeConferenceRecording(context.Background(), id).ResumeConferenceRecordingRequest(resumeConferenceRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.ResumeConferenceRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeConferenceRecording`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.ResumeConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resumeConferenceRecordingRequest** | [**ResumeConferenceRecordingRequest**](ResumeConferenceRecordingRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConference

> ConferenceResponse RetrieveConference(ctx, id).Execute()

Retrieve a conference



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
	id := "id_example" // string | Uniquely identifies the conference by id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.RetrieveConference(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.RetrieveConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveConference`: ConferenceResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.RetrieveConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConferenceResponse**](ConferenceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeakTextToConference

> ConferenceCommandResponse SpeakTextToConference(ctx, id).ConferenceSpeakRequest(conferenceSpeakRequest).Execute()

Speak text to conference participants



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
	id := "id_example" // string | Specifies the conference by id or name
	conferenceSpeakRequest := *openapiclient.NewConferenceSpeakRequest("Say this to participants", "Telnyx.KokoroTTS.af") // ConferenceSpeakRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.SpeakTextToConference(context.Background(), id).ConferenceSpeakRequest(conferenceSpeakRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.SpeakTextToConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeakTextToConference`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.SpeakTextToConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeakTextToConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferenceSpeakRequest** | [**ConferenceSpeakRequest**](ConferenceSpeakRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartConferenceRecording

> ConferenceCommandResponse StartConferenceRecording(ctx, id).StartConferenceRecordingRequest(startConferenceRecordingRequest).Execute()

Conference recording start



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
	id := "id_example" // string | Specifies the conference to record by id or name
	startConferenceRecordingRequest := *openapiclient.NewStartConferenceRecordingRequest("mp3") // StartConferenceRecordingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.StartConferenceRecording(context.Background(), id).StartConferenceRecordingRequest(startConferenceRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.StartConferenceRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartConferenceRecording`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.StartConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the conference to record by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startConferenceRecordingRequest** | [**StartConferenceRecordingRequest**](StartConferenceRecordingRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopConferenceAudio

> ConferenceCommandResponse StopConferenceAudio(ctx, id).ConferenceStopRequest(conferenceStopRequest).Execute()

Stop audio being played on the conference



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	conferenceStopRequest := *openapiclient.NewConferenceStopRequest() // ConferenceStopRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.StopConferenceAudio(context.Background(), id).ConferenceStopRequest(conferenceStopRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.StopConferenceAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopConferenceAudio`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.StopConferenceAudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopConferenceAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferenceStopRequest** | [**ConferenceStopRequest**](ConferenceStopRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopConferenceRecording

> ConferenceCommandResponse StopConferenceRecording(ctx, id).StopRecordingRequest(stopRecordingRequest).Execute()

Conference recording stop



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
	id := "id_example" // string | Specifies the conference to stop the recording for by id or name
	stopRecordingRequest := *openapiclient.NewStopRecordingRequest() // StopRecordingRequest | Stop recording conference request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.StopConferenceRecording(context.Background(), id).StopRecordingRequest(stopRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.StopConferenceRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopConferenceRecording`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.StopConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the conference to stop the recording for by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopRecordingRequest** | [**StopRecordingRequest**](StopRecordingRequest.md) | Stop recording conference request | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnholdConferenceParticipants

> ConferenceCommandResponse UnholdConferenceParticipants(ctx, id).ConferenceUnholdRequest(conferenceUnholdRequest).Execute()

Unhold conference participants



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	conferenceUnholdRequest := *openapiclient.NewConferenceUnholdRequest([]string{"CallControlIds_example"}) // ConferenceUnholdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.UnholdConferenceParticipants(context.Background(), id).ConferenceUnholdRequest(conferenceUnholdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.UnholdConferenceParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnholdConferenceParticipants`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.UnholdConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnholdConferenceParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferenceUnholdRequest** | [**ConferenceUnholdRequest**](ConferenceUnholdRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmuteConferenceParticipants

> ConferenceCommandResponse UnmuteConferenceParticipants(ctx, id).ConferenceUnmuteRequest(conferenceUnmuteRequest).Execute()

Unmute conference participants



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	conferenceUnmuteRequest := *openapiclient.NewConferenceUnmuteRequest() // ConferenceUnmuteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.UnmuteConferenceParticipants(context.Background(), id).ConferenceUnmuteRequest(conferenceUnmuteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.UnmuteConferenceParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmuteConferenceParticipants`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.UnmuteConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteConferenceParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conferenceUnmuteRequest** | [**ConferenceUnmuteRequest**](ConferenceUnmuteRequest.md) |  | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConference

> ConferenceCommandResponse UpdateConference(ctx, id).UpdateConferenceRequest(updateConferenceRequest).Execute()

Update conference participant



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
	id := "id_example" // string | Uniquely identifies the conference by id or name
	updateConferenceRequest := *openapiclient.NewUpdateConferenceRequest() // UpdateConferenceRequest | Update Conference request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferenceCommandsAPI.UpdateConference(context.Background(), id).UpdateConferenceRequest(updateConferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferenceCommandsAPI.UpdateConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConference`: ConferenceCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `ConferenceCommandsAPI.UpdateConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Uniquely identifies the conference by id or name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConferenceRequest** | [**UpdateConferenceRequest**](UpdateConferenceRequest.md) | Update Conference request object | 

### Return type

[**ConferenceCommandResponse**](ConferenceCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

