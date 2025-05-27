# \CallCommandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerCall**](CallCommandsAPI.md#AnswerCall) | **Post** /calls/{call_control_id}/actions/answer | Answer call
[**BridgeCall**](CallCommandsAPI.md#BridgeCall) | **Post** /calls/{call_control_id}/actions/bridge | Bridge calls
[**CallGatherUsingAI**](CallCommandsAPI.md#CallGatherUsingAI) | **Post** /calls/{call_control_id}/actions/gather_using_ai | Gather using AI
[**CallStartAIAssistant**](CallCommandsAPI.md#CallStartAIAssistant) | **Post** /calls/{call_control_id}/actions/ai_assistant_start | Start AI Assistant
[**CallStopAIAssistant**](CallCommandsAPI.md#CallStopAIAssistant) | **Post** /calls/{call_control_id}/actions/ai_assistant_stop | Stop AI Assistant
[**DialCall**](CallCommandsAPI.md#DialCall) | **Post** /calls | Dial
[**EnqueueCall**](CallCommandsAPI.md#EnqueueCall) | **Post** /calls/{call_control_id}/actions/enqueue | Enqueue call
[**GatherCall**](CallCommandsAPI.md#GatherCall) | **Post** /calls/{call_control_id}/actions/gather | Gather
[**GatherUsingAudio**](CallCommandsAPI.md#GatherUsingAudio) | **Post** /calls/{call_control_id}/actions/gather_using_audio | Gather using audio
[**GatherUsingSpeak**](CallCommandsAPI.md#GatherUsingSpeak) | **Post** /calls/{call_control_id}/actions/gather_using_speak | Gather using speak
[**HangupCall**](CallCommandsAPI.md#HangupCall) | **Post** /calls/{call_control_id}/actions/hangup | Hangup call
[**LeaveQueue**](CallCommandsAPI.md#LeaveQueue) | **Post** /calls/{call_control_id}/actions/leave_queue | Remove call from a queue
[**NoiseSuppressionStart**](CallCommandsAPI.md#NoiseSuppressionStart) | **Post** /calls/{call_control_id}/actions/suppression_start | Noise Suppression Start (BETA)
[**NoiseSuppressionStop**](CallCommandsAPI.md#NoiseSuppressionStop) | **Post** /calls/{call_control_id}/actions/suppression_stop | Noise Suppression Stop (BETA)
[**PauseCallRecording**](CallCommandsAPI.md#PauseCallRecording) | **Post** /calls/{call_control_id}/actions/record_pause | Record pause
[**ReferCall**](CallCommandsAPI.md#ReferCall) | **Post** /calls/{call_control_id}/actions/refer | SIP Refer a call
[**RejectCall**](CallCommandsAPI.md#RejectCall) | **Post** /calls/{call_control_id}/actions/reject | Reject a call
[**ResumeCallRecording**](CallCommandsAPI.md#ResumeCallRecording) | **Post** /calls/{call_control_id}/actions/record_resume | Record resume
[**SendDTMF**](CallCommandsAPI.md#SendDTMF) | **Post** /calls/{call_control_id}/actions/send_dtmf | Send DTMF
[**SendSIPInfo**](CallCommandsAPI.md#SendSIPInfo) | **Post** /calls/{call_control_id}/actions/send_sip_info | Send SIP info
[**SpeakCall**](CallCommandsAPI.md#SpeakCall) | **Post** /calls/{call_control_id}/actions/speak | Speak text
[**StartCallFork**](CallCommandsAPI.md#StartCallFork) | **Post** /calls/{call_control_id}/actions/fork_start | Forking start
[**StartCallPlayback**](CallCommandsAPI.md#StartCallPlayback) | **Post** /calls/{call_control_id}/actions/playback_start | Play audio URL
[**StartCallRecord**](CallCommandsAPI.md#StartCallRecord) | **Post** /calls/{call_control_id}/actions/record_start | Recording start
[**StartCallStreaming**](CallCommandsAPI.md#StartCallStreaming) | **Post** /calls/{call_control_id}/actions/streaming_start | Streaming start
[**StartCallTranscription**](CallCommandsAPI.md#StartCallTranscription) | **Post** /calls/{call_control_id}/actions/transcription_start | Transcription start
[**StartSiprecSession**](CallCommandsAPI.md#StartSiprecSession) | **Post** /calls/{call_control_id}/actions/siprec_start | SIPREC start
[**StopCallFork**](CallCommandsAPI.md#StopCallFork) | **Post** /calls/{call_control_id}/actions/fork_stop | Forking stop
[**StopCallGather**](CallCommandsAPI.md#StopCallGather) | **Post** /calls/{call_control_id}/actions/gather_stop | Gather stop
[**StopCallPlayback**](CallCommandsAPI.md#StopCallPlayback) | **Post** /calls/{call_control_id}/actions/playback_stop | Stop audio playback
[**StopCallRecording**](CallCommandsAPI.md#StopCallRecording) | **Post** /calls/{call_control_id}/actions/record_stop | Recording stop
[**StopCallStreaming**](CallCommandsAPI.md#StopCallStreaming) | **Post** /calls/{call_control_id}/actions/streaming_stop | Streaming stop
[**StopCallTranscription**](CallCommandsAPI.md#StopCallTranscription) | **Post** /calls/{call_control_id}/actions/transcription_stop | Transcription stop
[**StopSiprecSession**](CallCommandsAPI.md#StopSiprecSession) | **Post** /calls/{call_control_id}/actions/siprec_stop | SIPREC stop
[**TransferCall**](CallCommandsAPI.md#TransferCall) | **Post** /calls/{call_control_id}/actions/transfer | Transfer call
[**UpdateClientState**](CallCommandsAPI.md#UpdateClientState) | **Put** /calls/{call_control_id}/actions/client_state_update | Update client state



## AnswerCall

> CallControlCommandResponseWithRecordingID AnswerCall(ctx, callControlId).AnswerRequest(answerRequest).Execute()

Answer call



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	answerRequest := *openapiclient.NewAnswerRequest() // AnswerRequest | Answer call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.AnswerCall(context.Background(), callControlId).AnswerRequest(answerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.AnswerCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnswerCall`: CallControlCommandResponseWithRecordingID
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.AnswerCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **answerRequest** | [**AnswerRequest**](AnswerRequest.md) | Answer call request | 

### Return type

[**CallControlCommandResponseWithRecordingID**](CallControlCommandResponseWithRecordingID.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BridgeCall

> CallControlCommandResponse BridgeCall(ctx, callControlId).BridgeRequest(bridgeRequest).Execute()

Bridge calls



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	bridgeRequest := *openapiclient.NewBridgeRequest("v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg") // BridgeRequest | Bridge call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.BridgeCall(context.Background(), callControlId).BridgeRequest(bridgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.BridgeCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BridgeCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.BridgeCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiBridgeCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bridgeRequest** | [**BridgeRequest**](BridgeRequest.md) | Bridge call request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallGatherUsingAI

> CallControlCommandResponse CallGatherUsingAI(ctx, callControlId).GatherUsingAIRequest(gatherUsingAIRequest).Execute()

Gather using AI



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	gatherUsingAIRequest := *openapiclient.NewGatherUsingAIRequest(map[string]interface{}({"properties":{"age":{"description":"The age of the customer.","type":"integer"},"location":{"description":"The location of the customer.","type":"string"}},"required":["age","location"],"type":"object"})) // GatherUsingAIRequest | Gather using AI request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.CallGatherUsingAI(context.Background(), callControlId).GatherUsingAIRequest(gatherUsingAIRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.CallGatherUsingAI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallGatherUsingAI`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.CallGatherUsingAI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallGatherUsingAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatherUsingAIRequest** | [**GatherUsingAIRequest**](GatherUsingAIRequest.md) | Gather using AI request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallStartAIAssistant

> CallControlCommandResponse CallStartAIAssistant(ctx, callControlId).AIAssistantStartRequest(aIAssistantStartRequest).Execute()

Start AI Assistant



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	aIAssistantStartRequest := *openapiclient.NewAIAssistantStartRequest() // AIAssistantStartRequest | AI Assistant request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.CallStartAIAssistant(context.Background(), callControlId).AIAssistantStartRequest(aIAssistantStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.CallStartAIAssistant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallStartAIAssistant`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.CallStartAIAssistant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallStartAIAssistantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aIAssistantStartRequest** | [**AIAssistantStartRequest**](AIAssistantStartRequest.md) | AI Assistant request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallStopAIAssistant

> CallControlCommandResponse CallStopAIAssistant(ctx, callControlId).AIAssistantStopRequest(aIAssistantStopRequest).Execute()

Stop AI Assistant



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	aIAssistantStopRequest := *openapiclient.NewAIAssistantStopRequest() // AIAssistantStopRequest | AI Assistant request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.CallStopAIAssistant(context.Background(), callControlId).AIAssistantStopRequest(aIAssistantStopRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.CallStopAIAssistant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallStopAIAssistant`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.CallStopAIAssistant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallStopAIAssistantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aIAssistantStopRequest** | [**AIAssistantStopRequest**](AIAssistantStopRequest.md) | AI Assistant request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DialCall

> RetrieveCallStatusResponseWithRecordingID DialCall(ctx).CallRequest(callRequest).Execute()

Dial



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
	callRequest := *openapiclient.NewCallRequest(openapiclient.CallRequest_to{ArrayOfString: new([]string)}, "+18005550101", "ConnectionId_example") // CallRequest | Call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.DialCall(context.Background()).CallRequest(callRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.DialCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DialCall`: RetrieveCallStatusResponseWithRecordingID
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.DialCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDialCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callRequest** | [**CallRequest**](CallRequest.md) | Call request | 

### Return type

[**RetrieveCallStatusResponseWithRecordingID**](RetrieveCallStatusResponseWithRecordingID.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnqueueCall

> CallControlCommandResponse EnqueueCall(ctx, callControlId).EnqueueRequest(enqueueRequest).Execute()

Enqueue call



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	enqueueRequest := *openapiclient.NewEnqueueRequest("tier_1_support") // EnqueueRequest | Enqueue call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.EnqueueCall(context.Background(), callControlId).EnqueueRequest(enqueueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.EnqueueCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnqueueCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.EnqueueCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnqueueCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enqueueRequest** | [**EnqueueRequest**](EnqueueRequest.md) | Enqueue call request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatherCall

> CallControlCommandResponse GatherCall(ctx, callControlId).GatherRequest(gatherRequest).Execute()

Gather



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	gatherRequest := *openapiclient.NewGatherRequest() // GatherRequest | Gather

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.GatherCall(context.Background(), callControlId).GatherRequest(gatherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.GatherCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatherCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.GatherCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatherCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatherRequest** | [**GatherRequest**](GatherRequest.md) | Gather | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatherUsingAudio

> CallControlCommandResponse GatherUsingAudio(ctx, callControlId).GatherUsingAudioRequest(gatherUsingAudioRequest).Execute()

Gather using audio



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	gatherUsingAudioRequest := *openapiclient.NewGatherUsingAudioRequest() // GatherUsingAudioRequest | Gather using audio request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.GatherUsingAudio(context.Background(), callControlId).GatherUsingAudioRequest(gatherUsingAudioRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.GatherUsingAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatherUsingAudio`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.GatherUsingAudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatherUsingAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatherUsingAudioRequest** | [**GatherUsingAudioRequest**](GatherUsingAudioRequest.md) | Gather using audio request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatherUsingSpeak

> CallControlCommandResponse GatherUsingSpeak(ctx, callControlId).GatherUsingSpeakRequest(gatherUsingSpeakRequest).Execute()

Gather using speak



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	gatherUsingSpeakRequest := *openapiclient.NewGatherUsingSpeakRequest("Say this on the call", "Telnyx.KokoroTTS.af") // GatherUsingSpeakRequest | Gather using speak request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.GatherUsingSpeak(context.Background(), callControlId).GatherUsingSpeakRequest(gatherUsingSpeakRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.GatherUsingSpeak``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatherUsingSpeak`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.GatherUsingSpeak`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatherUsingSpeakRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatherUsingSpeakRequest** | [**GatherUsingSpeakRequest**](GatherUsingSpeakRequest.md) | Gather using speak request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HangupCall

> CallControlCommandResponse HangupCall(ctx, callControlId).HangupRequest(hangupRequest).Execute()

Hangup call



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	hangupRequest := *openapiclient.NewHangupRequest() // HangupRequest | Hangup request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.HangupCall(context.Background(), callControlId).HangupRequest(hangupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.HangupCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HangupCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.HangupCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiHangupCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hangupRequest** | [**HangupRequest**](HangupRequest.md) | Hangup request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveQueue

> CallControlCommandResponse LeaveQueue(ctx, callControlId).LeaveQueueRequest(leaveQueueRequest).Execute()

Remove call from a queue



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	leaveQueueRequest := *openapiclient.NewLeaveQueueRequest() // LeaveQueueRequest | Removes the call from the queue, the call currently is enqueued in.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.LeaveQueue(context.Background(), callControlId).LeaveQueueRequest(leaveQueueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.LeaveQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaveQueue`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.LeaveQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leaveQueueRequest** | [**LeaveQueueRequest**](LeaveQueueRequest.md) | Removes the call from the queue, the call currently is enqueued in. | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoiseSuppressionStart

> CallControlCommandResponse NoiseSuppressionStart(ctx, callControlId).NoiseSuppressionStart(noiseSuppressionStart).Execute()

Noise Suppression Start (BETA)

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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	noiseSuppressionStart := *openapiclient.NewNoiseSuppressionStart() // NoiseSuppressionStart | Start streaming media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.NoiseSuppressionStart(context.Background(), callControlId).NoiseSuppressionStart(noiseSuppressionStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.NoiseSuppressionStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoiseSuppressionStart`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.NoiseSuppressionStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoiseSuppressionStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noiseSuppressionStart** | [**NoiseSuppressionStart**](NoiseSuppressionStart.md) | Start streaming media request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoiseSuppressionStop

> CallControlCommandResponse NoiseSuppressionStop(ctx, callControlId).NoiseSuppressionStop(noiseSuppressionStop).Execute()

Noise Suppression Stop (BETA)

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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	noiseSuppressionStop := *openapiclient.NewNoiseSuppressionStop() // NoiseSuppressionStop | Start streaming media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.NoiseSuppressionStop(context.Background(), callControlId).NoiseSuppressionStop(noiseSuppressionStop).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.NoiseSuppressionStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoiseSuppressionStop`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.NoiseSuppressionStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoiseSuppressionStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noiseSuppressionStop** | [**NoiseSuppressionStop**](NoiseSuppressionStop.md) | Start streaming media request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseCallRecording

> CallControlCommandResponse PauseCallRecording(ctx, callControlId).PauseRecordingRequest(pauseRecordingRequest).Execute()

Record pause



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	pauseRecordingRequest := *openapiclient.NewPauseRecordingRequest() // PauseRecordingRequest | Pause recording call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.PauseCallRecording(context.Background(), callControlId).PauseRecordingRequest(pauseRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.PauseCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseCallRecording`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.PauseCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pauseRecordingRequest** | [**PauseRecordingRequest**](PauseRecordingRequest.md) | Pause recording call request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferCall

> CallControlCommandResponse ReferCall(ctx, callControlId).ReferRequest(referRequest).Execute()

SIP Refer a call



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	referRequest := *openapiclient.NewReferRequest("sip:username@sip.non-telnyx-address.com") // ReferRequest | Refer request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.ReferCall(context.Background(), callControlId).ReferRequest(referRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.ReferCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.ReferCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **referRequest** | [**ReferRequest**](ReferRequest.md) | Refer request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectCall

> CallControlCommandResponse RejectCall(ctx, callControlId).RejectRequest(rejectRequest).Execute()

Reject a call



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	rejectRequest := *openapiclient.NewRejectRequest("USER_BUSY") // RejectRequest | Reject request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.RejectCall(context.Background(), callControlId).RejectRequest(rejectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.RejectCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.RejectCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectRequest** | [**RejectRequest**](RejectRequest.md) | Reject request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeCallRecording

> CallControlCommandResponse ResumeCallRecording(ctx, callControlId).ResumeRecordingRequest(resumeRecordingRequest).Execute()

Record resume



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	resumeRecordingRequest := *openapiclient.NewResumeRecordingRequest() // ResumeRecordingRequest | Resume recording call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.ResumeCallRecording(context.Background(), callControlId).ResumeRecordingRequest(resumeRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.ResumeCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeCallRecording`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.ResumeCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resumeRecordingRequest** | [**ResumeRecordingRequest**](ResumeRecordingRequest.md) | Resume recording call request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDTMF

> CallControlCommandResponse SendDTMF(ctx, callControlId).SendDTMFRequest(sendDTMFRequest).Execute()

Send DTMF



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	sendDTMFRequest := *openapiclient.NewSendDTMFRequest("1www2WABCDw9") // SendDTMFRequest | Send DTMF request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.SendDTMF(context.Background(), callControlId).SendDTMFRequest(sendDTMFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.SendDTMF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendDTMF`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.SendDTMF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendDTMFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendDTMFRequest** | [**SendDTMFRequest**](SendDTMFRequest.md) | Send DTMF request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSIPInfo

> CallControlCommandResponse SendSIPInfo(ctx, callControlId).SendSIPInfoRequest(sendSIPInfoRequest).Execute()

Send SIP info



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	sendSIPInfoRequest := *openapiclient.NewSendSIPInfoRequest("application/json", "{"key": "value", "numValue": 100}") // SendSIPInfoRequest | Send SIP INFO request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.SendSIPInfo(context.Background(), callControlId).SendSIPInfoRequest(sendSIPInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.SendSIPInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSIPInfo`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.SendSIPInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSIPInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendSIPInfoRequest** | [**SendSIPInfoRequest**](SendSIPInfoRequest.md) | Send SIP INFO request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeakCall

> CallControlCommandResponse SpeakCall(ctx, callControlId).SpeakRequest(speakRequest).Execute()

Speak text



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	speakRequest := *openapiclient.NewSpeakRequest("Say this on the call", "Telnyx.KokoroTTS.af") // SpeakRequest | Speak request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.SpeakCall(context.Background(), callControlId).SpeakRequest(speakRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.SpeakCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeakCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.SpeakCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeakCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **speakRequest** | [**SpeakRequest**](SpeakRequest.md) | Speak request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCallFork

> CallControlCommandResponse StartCallFork(ctx, callControlId).StartForkingRequest(startForkingRequest).Execute()

Forking start



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	startForkingRequest := *openapiclient.NewStartForkingRequest() // StartForkingRequest | Fork media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StartCallFork(context.Background(), callControlId).StartForkingRequest(startForkingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StartCallFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCallFork`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StartCallFork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCallForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startForkingRequest** | [**StartForkingRequest**](StartForkingRequest.md) | Fork media request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCallPlayback

> CallControlCommandResponse StartCallPlayback(ctx, callControlId).PlayAudioUrlRequest(playAudioUrlRequest).Execute()

Play audio URL



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	playAudioUrlRequest := *openapiclient.NewPlayAudioUrlRequest() // PlayAudioUrlRequest | Play audio URL request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StartCallPlayback(context.Background(), callControlId).PlayAudioUrlRequest(playAudioUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StartCallPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCallPlayback`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StartCallPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCallPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **playAudioUrlRequest** | [**PlayAudioUrlRequest**](PlayAudioUrlRequest.md) | Play audio URL request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCallRecord

> CallControlCommandResponse StartCallRecord(ctx, callControlId).StartRecordingRequest(startRecordingRequest).Execute()

Recording start



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	startRecordingRequest := *openapiclient.NewStartRecordingRequest("mp3", "single") // StartRecordingRequest | Start recording audio request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StartCallRecord(context.Background(), callControlId).StartRecordingRequest(startRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StartCallRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCallRecord`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StartCallRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCallRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startRecordingRequest** | [**StartRecordingRequest**](StartRecordingRequest.md) | Start recording audio request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCallStreaming

> CallControlCommandResponse StartCallStreaming(ctx, callControlId).StartStreamingRequest(startStreamingRequest).Execute()

Streaming start



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	startStreamingRequest := *openapiclient.NewStartStreamingRequest() // StartStreamingRequest | Start streaming media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StartCallStreaming(context.Background(), callControlId).StartStreamingRequest(startStreamingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StartCallStreaming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCallStreaming`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StartCallStreaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCallStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startStreamingRequest** | [**StartStreamingRequest**](StartStreamingRequest.md) | Start streaming media request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCallTranscription

> CallControlCommandResponse StartCallTranscription(ctx, callControlId).TranscriptionStartRequest(transcriptionStartRequest).Execute()

Transcription start



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	transcriptionStartRequest := *openapiclient.NewTranscriptionStartRequest() // TranscriptionStartRequest | Transcription start request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StartCallTranscription(context.Background(), callControlId).TranscriptionStartRequest(transcriptionStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StartCallTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCallTranscription`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StartCallTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCallTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transcriptionStartRequest** | [**TranscriptionStartRequest**](TranscriptionStartRequest.md) | Transcription start request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSiprecSession

> CallControlCommandResponse StartSiprecSession(ctx, callControlId).StartSiprecRequest(startSiprecRequest).Execute()

SIPREC start



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	startSiprecRequest := *openapiclient.NewStartSiprecRequest() // StartSiprecRequest | Start siprec session to configured in SIPREC connector SRS.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StartSiprecSession(context.Background(), callControlId).StartSiprecRequest(startSiprecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StartSiprecSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSiprecSession`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StartSiprecSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSiprecSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startSiprecRequest** | [**StartSiprecRequest**](StartSiprecRequest.md) | Start siprec session to configured in SIPREC connector SRS. | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCallFork

> CallControlCommandResponse StopCallFork(ctx, callControlId).StopForkingRequest(stopForkingRequest).Execute()

Forking stop



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	stopForkingRequest := *openapiclient.NewStopForkingRequest() // StopForkingRequest | Stop forking media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopCallFork(context.Background(), callControlId).StopForkingRequest(stopForkingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopCallFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCallFork`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopCallFork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCallForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopForkingRequest** | [**StopForkingRequest**](StopForkingRequest.md) | Stop forking media request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCallGather

> CallControlCommandResponse StopCallGather(ctx, callControlId).StopGatherRequest(stopGatherRequest).Execute()

Gather stop



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	stopGatherRequest := *openapiclient.NewStopGatherRequest() // StopGatherRequest | Stop current gather

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopCallGather(context.Background(), callControlId).StopGatherRequest(stopGatherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopCallGather``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCallGather`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopCallGather`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCallGatherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopGatherRequest** | [**StopGatherRequest**](StopGatherRequest.md) | Stop current gather | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCallPlayback

> CallControlCommandResponse StopCallPlayback(ctx, callControlId).PlaybackStopRequest(playbackStopRequest).Execute()

Stop audio playback



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	playbackStopRequest := *openapiclient.NewPlaybackStopRequest() // PlaybackStopRequest | Stop audio playback request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopCallPlayback(context.Background(), callControlId).PlaybackStopRequest(playbackStopRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopCallPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCallPlayback`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopCallPlayback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCallPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **playbackStopRequest** | [**PlaybackStopRequest**](PlaybackStopRequest.md) | Stop audio playback request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCallRecording

> CallControlCommandResponse StopCallRecording(ctx, callControlId).StopRecordingRequest(stopRecordingRequest).Execute()

Recording stop



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	stopRecordingRequest := *openapiclient.NewStopRecordingRequest() // StopRecordingRequest | Stop recording call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopCallRecording(context.Background(), callControlId).StopRecordingRequest(stopRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCallRecording`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopRecordingRequest** | [**StopRecordingRequest**](StopRecordingRequest.md) | Stop recording call request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCallStreaming

> CallControlCommandResponse StopCallStreaming(ctx, callControlId).StopStreamingRequest(stopStreamingRequest).Execute()

Streaming stop



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	stopStreamingRequest := *openapiclient.NewStopStreamingRequest() // StopStreamingRequest | Stop streaming media request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopCallStreaming(context.Background(), callControlId).StopStreamingRequest(stopStreamingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopCallStreaming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCallStreaming`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopCallStreaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCallStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopStreamingRequest** | [**StopStreamingRequest**](StopStreamingRequest.md) | Stop streaming media request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCallTranscription

> CallControlCommandResponse StopCallTranscription(ctx, callControlId).TranscriptionStopRequest(transcriptionStopRequest).Execute()

Transcription stop



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	transcriptionStopRequest := *openapiclient.NewTranscriptionStopRequest() // TranscriptionStopRequest | Transcription stop request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopCallTranscription(context.Background(), callControlId).TranscriptionStopRequest(transcriptionStopRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopCallTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCallTranscription`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopCallTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCallTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transcriptionStopRequest** | [**TranscriptionStopRequest**](TranscriptionStopRequest.md) | Transcription stop request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSiprecSession

> CallControlCommandResponse StopSiprecSession(ctx, callControlId).StopSiprecRequest(stopSiprecRequest).Execute()

SIPREC stop



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	stopSiprecRequest := *openapiclient.NewStopSiprecRequest() // StopSiprecRequest | Stop siprec session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.StopSiprecSession(context.Background(), callControlId).StopSiprecRequest(stopSiprecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.StopSiprecSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopSiprecSession`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.StopSiprecSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSiprecSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopSiprecRequest** | [**StopSiprecRequest**](StopSiprecRequest.md) | Stop siprec session | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferCall

> CallControlCommandResponse TransferCall(ctx, callControlId).TransferCallRequest(transferCallRequest).Execute()

Transfer call



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	transferCallRequest := *openapiclient.NewTransferCallRequest("+18005550100 or sip:username@sip.telnyx.com") // TransferCallRequest | Transfer call request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.TransferCall(context.Background(), callControlId).TransferCallRequest(transferCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.TransferCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferCall`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.TransferCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferCallRequest** | [**TransferCallRequest**](TransferCallRequest.md) | Transfer call request | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientState

> CallControlCommandResponse UpdateClientState(ctx, callControlId).ClientStateUpdateRequest(clientStateUpdateRequest).Execute()

Update client state



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
	callControlId := "callControlId_example" // string | Unique identifier and token for controlling the call
	clientStateUpdateRequest := *openapiclient.NewClientStateUpdateRequest("aGF2ZSBhIG5pY2UgZGF5ID1d") // ClientStateUpdateRequest | Updates client state for every subsequent webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallCommandsAPI.UpdateClientState(context.Background(), callControlId).ClientStateUpdateRequest(clientStateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallCommandsAPI.UpdateClientState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientState`: CallControlCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `CallCommandsAPI.UpdateClientState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callControlId** | **string** | Unique identifier and token for controlling the call | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientStateUpdateRequest** | [**ClientStateUpdateRequest**](ClientStateUpdateRequest.md) | Updates client state for every subsequent webhook | 

### Return type

[**CallControlCommandResponse**](CallControlCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

