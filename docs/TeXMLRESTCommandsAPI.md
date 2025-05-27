# \TeXMLRESTCommandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTexmlSecret**](TeXMLRESTCommandsAPI.md#CreateTexmlSecret) | **Post** /texml/secrets | Create a TeXML secret
[**DeleteTeXMLCallRecording**](TeXMLRESTCommandsAPI.md#DeleteTeXMLCallRecording) | **Delete** /texml/Accounts/{account_sid}/Recordings/{recording_sid}.json | Delete recording resource
[**DeleteTeXMLRecordingTranscription**](TeXMLRESTCommandsAPI.md#DeleteTeXMLRecordingTranscription) | **Delete** /texml/Accounts/{account_sid}/Transcriptions/{recording_transcription_sid}.json | Delete a recording transcription
[**DeleteTexmlConferenceParticipant**](TeXMLRESTCommandsAPI.md#DeleteTexmlConferenceParticipant) | **Delete** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants/{call_sid} | Delete a conference participant
[**DeprecatedInitiateTexmlCall**](TeXMLRESTCommandsAPI.md#DeprecatedInitiateTexmlCall) | **Post** /texml/calls/{application_id} | (Deprecated) Initiate an outbound call
[**DeprecatedUpdateTexmlCall**](TeXMLRESTCommandsAPI.md#DeprecatedUpdateTexmlCall) | **Post** /texml/calls/{call_sid}/update | (Deprecated) Update call
[**DialTexmlConferenceParticipant**](TeXMLRESTCommandsAPI.md#DialTexmlConferenceParticipant) | **Post** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants | Dial a new conference participant
[**FetchTeXMLCallRecordings**](TeXMLRESTCommandsAPI.md#FetchTeXMLCallRecordings) | **Get** /texml/Accounts/{account_sid}/Calls/{call_sid}/Recordings.json | Fetch recordings for a call
[**FetchTeXMLConferenceRecordings**](TeXMLRESTCommandsAPI.md#FetchTeXMLConferenceRecordings) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Recordings.json | Fetch recordings for a conference
[**GetTeXMLCallRecording**](TeXMLRESTCommandsAPI.md#GetTeXMLCallRecording) | **Get** /texml/Accounts/{account_sid}/Recordings/{recording_sid}.json | Fetch recording resource
[**GetTeXMLCallRecordings**](TeXMLRESTCommandsAPI.md#GetTeXMLCallRecordings) | **Get** /texml/Accounts/{account_sid}/Recordings.json | Fetch multiple recording resources
[**GetTeXMLRecordingTranscription**](TeXMLRESTCommandsAPI.md#GetTeXMLRecordingTranscription) | **Get** /texml/Accounts/{account_sid}/Transcriptions/{recording_transcription_sid}.json | Fetch a recording transcription resource
[**GetTeXMLRecordingTranscriptions**](TeXMLRESTCommandsAPI.md#GetTeXMLRecordingTranscriptions) | **Get** /texml/Accounts/{account_sid}/Transcriptions.json | List recording transcriptions
[**GetTexmlCall**](TeXMLRESTCommandsAPI.md#GetTexmlCall) | **Get** /texml/Accounts/{account_sid}/Calls/{call_sid} | Fetch a call
[**GetTexmlCalls**](TeXMLRESTCommandsAPI.md#GetTexmlCalls) | **Get** /texml/Accounts/{account_sid}/Calls | Fetch multiple call resources
[**GetTexmlConference**](TeXMLRESTCommandsAPI.md#GetTexmlConference) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid} | Fetch a conference resource
[**GetTexmlConferenceParticipant**](TeXMLRESTCommandsAPI.md#GetTexmlConferenceParticipant) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants/{call_sid} | Get conference participant resource
[**GetTexmlConferenceParticipants**](TeXMLRESTCommandsAPI.md#GetTexmlConferenceParticipants) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants | List conference participants
[**GetTexmlConferenceRecordings**](TeXMLRESTCommandsAPI.md#GetTexmlConferenceRecordings) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Recordings | List conference recordings
[**GetTexmlConferences**](TeXMLRESTCommandsAPI.md#GetTexmlConferences) | **Get** /texml/Accounts/{account_sid}/Conferences | List conference resources
[**InitiateTexmlCall**](TeXMLRESTCommandsAPI.md#InitiateTexmlCall) | **Post** /texml/Accounts/{account_sid}/Calls | Initiate an outbound call
[**StartTeXMLCallRecording**](TeXMLRESTCommandsAPI.md#StartTeXMLCallRecording) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Recordings.json | Request recording for a call
[**StartTeXMLCallStreaming**](TeXMLRESTCommandsAPI.md#StartTeXMLCallStreaming) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Streams.json | Start streaming media from a call.
[**StartTeXMLSiprecSession**](TeXMLRESTCommandsAPI.md#StartTeXMLSiprecSession) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Siprec.json | Request siprec session for a call
[**UpdateTeXMLCallRecording**](TeXMLRESTCommandsAPI.md#UpdateTeXMLCallRecording) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Recordings/{recording_sid}.json | Update recording on a call
[**UpdateTeXMLCallStreaming**](TeXMLRESTCommandsAPI.md#UpdateTeXMLCallStreaming) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Streams/{streaming_sid}.json | Update streaming on a call
[**UpdateTeXMLSiprecSession**](TeXMLRESTCommandsAPI.md#UpdateTeXMLSiprecSession) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Siprec/{siprec_sid}.json | Updates siprec session for a call
[**UpdateTexmlCall**](TeXMLRESTCommandsAPI.md#UpdateTexmlCall) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid} | Update call
[**UpdateTexmlConference**](TeXMLRESTCommandsAPI.md#UpdateTexmlConference) | **Post** /texml/Accounts/{account_sid}/Conferences/{conference_sid} | Update a conference resource
[**UpdateTexmlConferenceParticipant**](TeXMLRESTCommandsAPI.md#UpdateTexmlConferenceParticipant) | **Post** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants/{call_sid} | Update a conference participant



## CreateTexmlSecret

> CreateTeXMLSecretRequest CreateTexmlSecret(ctx).CreateTeXMLSecretRequest(createTeXMLSecretRequest).Execute()

Create a TeXML secret



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
	createTeXMLSecretRequest := *openapiclient.NewCreateTeXMLSecretRequest() // CreateTeXMLSecretRequest | Create TeXML secret request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.CreateTexmlSecret(context.Background()).CreateTeXMLSecretRequest(createTeXMLSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.CreateTexmlSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTexmlSecret`: CreateTeXMLSecretRequest
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.CreateTexmlSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTexmlSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeXMLSecretRequest** | [**CreateTeXMLSecretRequest**](CreateTeXMLSecretRequest.md) | Create TeXML secret request object | 

### Return type

[**CreateTeXMLSecretRequest**](CreateTeXMLSecretRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeXMLCallRecording

> DeleteTeXMLCallRecording(ctx, accountSid, recordingSid).Execute()

Delete recording resource



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	recordingSid := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeXMLRESTCommandsAPI.DeleteTeXMLCallRecording(context.Background(), accountSid, recordingSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.DeleteTeXMLCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**recordingSid** | **string** | Uniquely identifies the recording by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeXMLCallRecordingRequest struct via the builder pattern


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


## DeleteTeXMLRecordingTranscription

> DeleteTeXMLRecordingTranscription(ctx, accountSid, recordingTranscriptionSid).Execute()

Delete a recording transcription



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	recordingTranscriptionSid := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording transcription by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeXMLRESTCommandsAPI.DeleteTeXMLRecordingTranscription(context.Background(), accountSid, recordingTranscriptionSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.DeleteTeXMLRecordingTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**recordingTranscriptionSid** | **string** | Uniquely identifies the recording transcription by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeXMLRecordingTranscriptionRequest struct via the builder pattern


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


## DeleteTexmlConferenceParticipant

> DeleteTexmlConferenceParticipant(ctx, accountSid, conferenceSid, callSid).Execute()

Delete a conference participant



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeXMLRESTCommandsAPI.DeleteTexmlConferenceParticipant(context.Background(), accountSid, conferenceSid, callSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.DeleteTexmlConferenceParticipant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTexmlConferenceParticipantRequest struct via the builder pattern


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


## DeprecatedInitiateTexmlCall

> InitiateTeXMLCallResponse DeprecatedInitiateTexmlCall(ctx, applicationId).DeprecatedInitiateCallRequest(deprecatedInitiateCallRequest).Execute()

(Deprecated) Initiate an outbound call



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
	applicationId := "applicationId_example" // string | The ID of the TeXML application used for the call.
	deprecatedInitiateCallRequest := *openapiclient.NewDeprecatedInitiateCallRequest("+16175551212", "+16175551212") // DeprecatedInitiateCallRequest | Iniatiate Call request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.DeprecatedInitiateTexmlCall(context.Background(), applicationId).DeprecatedInitiateCallRequest(deprecatedInitiateCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.DeprecatedInitiateTexmlCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeprecatedInitiateTexmlCall`: InitiateTeXMLCallResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.DeprecatedInitiateTexmlCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the TeXML application used for the call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedInitiateTexmlCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deprecatedInitiateCallRequest** | [**DeprecatedInitiateCallRequest**](DeprecatedInitiateCallRequest.md) | Iniatiate Call request object | 

### Return type

[**InitiateTeXMLCallResponse**](InitiateTeXMLCallResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedUpdateTexmlCall

> TeXMLRESTCommandResponse DeprecatedUpdateTexmlCall(ctx, callSid).UpdateCallRequest(updateCallRequest).Execute()

(Deprecated) Update call



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
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	updateCallRequest := *openapiclient.NewUpdateCallRequest() // UpdateCallRequest | Update Call request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.DeprecatedUpdateTexmlCall(context.Background(), callSid).UpdateCallRequest(updateCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.DeprecatedUpdateTexmlCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeprecatedUpdateTexmlCall`: TeXMLRESTCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.DeprecatedUpdateTexmlCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedUpdateTexmlCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCallRequest** | [**UpdateCallRequest**](UpdateCallRequest.md) | Update Call request object | 

### Return type

[**TeXMLRESTCommandResponse**](TeXMLRESTCommandResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DialTexmlConferenceParticipant

> NewParticipantResource DialTexmlConferenceParticipant(ctx, accountSid, conferenceSid).Beep(beep).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).StatusCallbackEvent(statusCallbackEvent).To(to).From(from).Timeout(timeout).Muted(muted).StartConferenceOnEnter(startConferenceOnEnter).EndConferenceOnExit(endConferenceOnExit).EarlyMedia(earlyMedia).ConferenceStatusCallback(conferenceStatusCallback).ConferenceStatusCallbackMethod(conferenceStatusCallbackMethod).ConferenceStatusCallbackEvent(conferenceStatusCallbackEvent).WaitUrl(waitUrl).MaxParticipants(maxParticipants).Coaching(coaching).CallSidToCoach(callSidToCoach).CallerId(callerId).TimeLimit(timeLimit).MachineDetection(machineDetection).MachineDetectionTimeout(machineDetectionTimeout).MachineDetectionSpeechThreshold(machineDetectionSpeechThreshold).MachineDetectionSpeechEndThreshold(machineDetectionSpeechEndThreshold).MachineDetectionSilenceTimeout(machineDetectionSilenceTimeout).AmdStatusCallback(amdStatusCallback).AmdStatusCallbackMethod(amdStatusCallbackMethod).CancelPlaybackOnMachineDetection(cancelPlaybackOnMachineDetection).CancelPlaybackOnDetectMessageEnd(cancelPlaybackOnDetectMessageEnd).PreferredCodecs(preferredCodecs).Record(record).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingTrack(recordingTrack).SipAuthPassword(sipAuthPassword).SipAuthUsername(sipAuthUsername).Trim(trim).ConferenceRecord(conferenceRecord).ConferenceRecordingStatusCallback(conferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackMethod(conferenceRecordingStatusCallbackMethod).ConferenceRecordingStatusCallbackEvent(conferenceRecordingStatusCallbackEvent).ConferenceRecordingTimeout(conferenceRecordingTimeout).ConferenceTrim(conferenceTrim).Execute()

Dial a new conference participant



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.
	beep := "beep_example" // string | Whether to play a notification beep to the conference when the participant enters and exits. (optional)
	statusCallback := "statusCallback_example" // string | URL destination for Telnyx to send status callback events to for the call. (optional)
	statusCallbackMethod := "statusCallbackMethod_example" // string | HTTP request type used for `StatusCallback`. (optional)
	statusCallbackEvent := "statusCallbackEvent_example" // string | The changes to the call's state that should generate a call to `StatusCallback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`. (optional)
	to := "to_example" // string | The phone number of the called party. Phone numbers are formatted with a `+` and country code. (optional)
	from := "from_example" // string | The phone number of the party that initiated the call. Phone numbers are formatted with a `+` and country code. (optional)
	timeout := int32(56) // int32 | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between 5 and 120, inclusive. The default value is 30. (optional)
	muted := true // bool | Whether the participant should be muted. (optional)
	startConferenceOnEnter := true // bool | Whether to start the conference when the participant enters. Defaults to `true`. (optional)
	endConferenceOnExit := true // bool | Whether to end the conference when the participant leaves. Defaults to `false`. (optional)
	earlyMedia := true // bool | Whether participant shall be bridged to conference before the participant answers (from early media if available). Defaults to `false`. (optional) (default to false)
	conferenceStatusCallback := "conferenceStatusCallback_example" // string | The URL the conference callbacks will be sent to. (optional)
	conferenceStatusCallbackMethod := "conferenceStatusCallbackMethod_example" // string | HTTP request type used for `ConferenceStatusCallback`. Defaults to `POST`. (optional)
	conferenceStatusCallbackEvent := "conferenceStatusCallbackEvent_example" // string | The changes to the conference's state that should generate a call to `ConferenceStatusCallback`. Can be: `start`, `end`, `join` and `leave`. Separate multiple values with a space. By default no callbacks are sent. (optional)
	waitUrl := "waitUrl_example" // string | The URL to call for an audio file to play while the participant is waiting for the conference to start. (optional)
	maxParticipants := int32(56) // int32 | The maximum number of participants in the conference. Can be a positive integer from 2 to 800. The default value is 250. (optional)
	coaching := true // bool | Whether the participant is coaching another call. When `true`, `CallSidToCoach` has to be given. (optional)
	callSidToCoach := "callSidToCoach_example" // string | The SID of the participant who is being coached. The participant being coached is the only participant who can hear the participant who is coaching. (optional)
	callerId := "callerId_example" // string | To be used as the caller id name (SIP From Display Name) presented to the destination (`To` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and `-_~!.+` special characters. If ommited, the display name will be the same as the number in the `From` field. (optional)
	timeLimit := int32(56) // int32 | The maximum duration of the call in seconds. (optional)
	machineDetection := "machineDetection_example" // string | Whether to detect if a human or an answering machine picked up the call. Use `Enable` if you would like to ne notified as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. (optional)
	machineDetectionTimeout := int32(56) // int32 | How long answering machine detection should go on for before sending an `Unknown` result. Given in milliseconds. (optional)
	machineDetectionSpeechThreshold := int32(56) // int32 | Maximum threshold of a human greeting. If greeting longer than this value, considered machine. Ignored when `premium` detection is used. (optional) (default to 3500)
	machineDetectionSpeechEndThreshold := int32(56) // int32 | Silence duration threshold after a greeting message or voice for it be considered human. Ignored when `premium` detection is used. (optional) (default to 800)
	machineDetectionSilenceTimeout := int32(56) // int32 | If initial silence duration is greater than this value, consider it a machine. Ignored when `premium` detection is used. (optional) (default to 3500)
	amdStatusCallback := "amdStatusCallback_example" // string | The URL the result of answering machine detection will be sent to. (optional)
	amdStatusCallbackMethod := "amdStatusCallbackMethod_example" // string | HTTP request type used for `AmdStatusCallback`. Defaults to `POST`. (optional)
	cancelPlaybackOnMachineDetection := true // bool | Whether to cancel ongoing playback on `machine` detection. Defaults to `true`. (optional) (default to true)
	cancelPlaybackOnDetectMessageEnd := true // bool | Whether to cancel ongoing playback on `greeting ended` detection. Defaults to `true`. (optional) (default to true)
	preferredCodecs := "preferredCodecs_example" // string | The list of comma-separated codecs to be offered on a call. (optional)
	record := true // bool | Whether to record the entire participant's call leg. Defaults to `false`. (optional)
	recordingChannels := "recordingChannels_example" // string | The number of channels in the final recording. Defaults to `mono`. (optional)
	recordingStatusCallback := "recordingStatusCallback_example" // string | The URL the recording callbacks will be sent to. (optional)
	recordingStatusCallbackMethod := "recordingStatusCallbackMethod_example" // string | HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`. (optional)
	recordingStatusCallbackEvent := "recordingStatusCallbackEvent_example" // string | The changes to the recording's state that should generate a call to `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`. (optional)
	recordingTrack := "recordingTrack_example" // string | The audio track to record for the call. The default is `both`. (optional)
	sipAuthPassword := "sipAuthPassword_example" // string | The password to use for SIP authentication. (optional)
	sipAuthUsername := "sipAuthUsername_example" // string | The username to use for SIP authentication. (optional)
	trim := "trim_example" // string | Whether to trim any leading and trailing silence from the recording. Defaults to `trim-silence`. (optional)
	conferenceRecord := "conferenceRecord_example" // string | Whether to record the conference the participant is joining. Defualts to `do-not-record`. The boolean values `true` and `false` are synonymous with `record-from-start` and `do-not-record` respectively. (optional)
	conferenceRecordingStatusCallback := "conferenceRecordingStatusCallback_example" // string | The URL the conference recording callbacks will be sent to. (optional)
	conferenceRecordingStatusCallbackMethod := "conferenceRecordingStatusCallbackMethod_example" // string | HTTP request type used for `ConferenceRecordingStatusCallback`. Defaults to `POST`. (optional)
	conferenceRecordingStatusCallbackEvent := "conferenceRecordingStatusCallbackEvent_example" // string | The changes to the conference recording's state that should generate a call to `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`. `failed` and `absent` are synonymous. (optional)
	conferenceRecordingTimeout := int32(56) // int32 | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that the transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite) (optional) (default to 0)
	conferenceTrim := "conferenceTrim_example" // string | Whether to trim any leading and trailing silence from the conference recording. Defaults to `trim-silence`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.DialTexmlConferenceParticipant(context.Background(), accountSid, conferenceSid).Beep(beep).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).StatusCallbackEvent(statusCallbackEvent).To(to).From(from).Timeout(timeout).Muted(muted).StartConferenceOnEnter(startConferenceOnEnter).EndConferenceOnExit(endConferenceOnExit).EarlyMedia(earlyMedia).ConferenceStatusCallback(conferenceStatusCallback).ConferenceStatusCallbackMethod(conferenceStatusCallbackMethod).ConferenceStatusCallbackEvent(conferenceStatusCallbackEvent).WaitUrl(waitUrl).MaxParticipants(maxParticipants).Coaching(coaching).CallSidToCoach(callSidToCoach).CallerId(callerId).TimeLimit(timeLimit).MachineDetection(machineDetection).MachineDetectionTimeout(machineDetectionTimeout).MachineDetectionSpeechThreshold(machineDetectionSpeechThreshold).MachineDetectionSpeechEndThreshold(machineDetectionSpeechEndThreshold).MachineDetectionSilenceTimeout(machineDetectionSilenceTimeout).AmdStatusCallback(amdStatusCallback).AmdStatusCallbackMethod(amdStatusCallbackMethod).CancelPlaybackOnMachineDetection(cancelPlaybackOnMachineDetection).CancelPlaybackOnDetectMessageEnd(cancelPlaybackOnDetectMessageEnd).PreferredCodecs(preferredCodecs).Record(record).RecordingChannels(recordingChannels).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingTrack(recordingTrack).SipAuthPassword(sipAuthPassword).SipAuthUsername(sipAuthUsername).Trim(trim).ConferenceRecord(conferenceRecord).ConferenceRecordingStatusCallback(conferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackMethod(conferenceRecordingStatusCallbackMethod).ConferenceRecordingStatusCallbackEvent(conferenceRecordingStatusCallbackEvent).ConferenceRecordingTimeout(conferenceRecordingTimeout).ConferenceTrim(conferenceTrim).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.DialTexmlConferenceParticipant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DialTexmlConferenceParticipant`: NewParticipantResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.DialTexmlConferenceParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDialTexmlConferenceParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **beep** | **string** | Whether to play a notification beep to the conference when the participant enters and exits. | 
 **statusCallback** | **string** | URL destination for Telnyx to send status callback events to for the call. | 
 **statusCallbackMethod** | **string** | HTTP request type used for &#x60;StatusCallback&#x60;. | 
 **statusCallbackEvent** | **string** | The changes to the call&#39;s state that should generate a call to &#x60;StatusCallback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. Separate multiple values with a space. The default value is &#x60;completed&#x60;. | 
 **to** | **string** | The phone number of the called party. Phone numbers are formatted with a &#x60;+&#x60; and country code. | 
 **from** | **string** | The phone number of the party that initiated the call. Phone numbers are formatted with a &#x60;+&#x60; and country code. | 
 **timeout** | **int32** | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between 5 and 120, inclusive. The default value is 30. | 
 **muted** | **bool** | Whether the participant should be muted. | 
 **startConferenceOnEnter** | **bool** | Whether to start the conference when the participant enters. Defaults to &#x60;true&#x60;. | 
 **endConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Defaults to &#x60;false&#x60;. | 
 **earlyMedia** | **bool** | Whether participant shall be bridged to conference before the participant answers (from early media if available). Defaults to &#x60;false&#x60;. | [default to false]
 **conferenceStatusCallback** | **string** | The URL the conference callbacks will be sent to. | 
 **conferenceStatusCallbackMethod** | **string** | HTTP request type used for &#x60;ConferenceStatusCallback&#x60;. Defaults to &#x60;POST&#x60;. | 
 **conferenceStatusCallbackEvent** | **string** | The changes to the conference&#39;s state that should generate a call to &#x60;ConferenceStatusCallback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60; and &#x60;leave&#x60;. Separate multiple values with a space. By default no callbacks are sent. | 
 **waitUrl** | **string** | The URL to call for an audio file to play while the participant is waiting for the conference to start. | 
 **maxParticipants** | **int32** | The maximum number of participants in the conference. Can be a positive integer from 2 to 800. The default value is 250. | 
 **coaching** | **bool** | Whether the participant is coaching another call. When &#x60;true&#x60;, &#x60;CallSidToCoach&#x60; has to be given. | 
 **callSidToCoach** | **string** | The SID of the participant who is being coached. The participant being coached is the only participant who can hear the participant who is coaching. | 
 **callerId** | **string** | To be used as the caller id name (SIP From Display Name) presented to the destination (&#x60;To&#x60; number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and &#x60;-_~!.+&#x60; special characters. If ommited, the display name will be the same as the number in the &#x60;From&#x60; field. | 
 **timeLimit** | **int32** | The maximum duration of the call in seconds. | 
 **machineDetection** | **string** | Whether to detect if a human or an answering machine picked up the call. Use &#x60;Enable&#x60; if you would like to ne notified as soon as the called party is identified. Use &#x60;DetectMessageEnd&#x60;, if you would like to leave a message on an answering machine. | 
 **machineDetectionTimeout** | **int32** | How long answering machine detection should go on for before sending an &#x60;Unknown&#x60; result. Given in milliseconds. | 
 **machineDetectionSpeechThreshold** | **int32** | Maximum threshold of a human greeting. If greeting longer than this value, considered machine. Ignored when &#x60;premium&#x60; detection is used. | [default to 3500]
 **machineDetectionSpeechEndThreshold** | **int32** | Silence duration threshold after a greeting message or voice for it be considered human. Ignored when &#x60;premium&#x60; detection is used. | [default to 800]
 **machineDetectionSilenceTimeout** | **int32** | If initial silence duration is greater than this value, consider it a machine. Ignored when &#x60;premium&#x60; detection is used. | [default to 3500]
 **amdStatusCallback** | **string** | The URL the result of answering machine detection will be sent to. | 
 **amdStatusCallbackMethod** | **string** | HTTP request type used for &#x60;AmdStatusCallback&#x60;. Defaults to &#x60;POST&#x60;. | 
 **cancelPlaybackOnMachineDetection** | **bool** | Whether to cancel ongoing playback on &#x60;machine&#x60; detection. Defaults to &#x60;true&#x60;. | [default to true]
 **cancelPlaybackOnDetectMessageEnd** | **bool** | Whether to cancel ongoing playback on &#x60;greeting ended&#x60; detection. Defaults to &#x60;true&#x60;. | [default to true]
 **preferredCodecs** | **string** | The list of comma-separated codecs to be offered on a call. | 
 **record** | **bool** | Whether to record the entire participant&#39;s call leg. Defaults to &#x60;false&#x60;. | 
 **recordingChannels** | **string** | The number of channels in the final recording. Defaults to &#x60;mono&#x60;. | 
 **recordingStatusCallback** | **string** | The URL the recording callbacks will be sent to. | 
 **recordingStatusCallbackMethod** | **string** | HTTP request type used for &#x60;RecordingStatusCallback&#x60;. Defaults to &#x60;POST&#x60;. | 
 **recordingStatusCallbackEvent** | **string** | The changes to the recording&#39;s state that should generate a call to &#x60;RecoridngStatusCallback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Separate multiple values with a space. Defaults to &#x60;completed&#x60;. | 
 **recordingTrack** | **string** | The audio track to record for the call. The default is &#x60;both&#x60;. | 
 **sipAuthPassword** | **string** | The password to use for SIP authentication. | 
 **sipAuthUsername** | **string** | The username to use for SIP authentication. | 
 **trim** | **string** | Whether to trim any leading and trailing silence from the recording. Defaults to &#x60;trim-silence&#x60;. | 
 **conferenceRecord** | **string** | Whether to record the conference the participant is joining. Defualts to &#x60;do-not-record&#x60;. The boolean values &#x60;true&#x60; and &#x60;false&#x60; are synonymous with &#x60;record-from-start&#x60; and &#x60;do-not-record&#x60; respectively. | 
 **conferenceRecordingStatusCallback** | **string** | The URL the conference recording callbacks will be sent to. | 
 **conferenceRecordingStatusCallbackMethod** | **string** | HTTP request type used for &#x60;ConferenceRecordingStatusCallback&#x60;. Defaults to &#x60;POST&#x60;. | 
 **conferenceRecordingStatusCallbackEvent** | **string** | The changes to the conference recording&#39;s state that should generate a call to &#x60;RecoridngStatusCallback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Separate multiple values with a space. Defaults to &#x60;completed&#x60;. &#x60;failed&#x60; and &#x60;absent&#x60; are synonymous. | 
 **conferenceRecordingTimeout** | **int32** | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that the transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite) | [default to 0]
 **conferenceTrim** | **string** | Whether to trim any leading and trailing silence from the conference recording. Defaults to &#x60;trim-silence&#x60;. | 

### Return type

[**NewParticipantResource**](NewParticipantResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTeXMLCallRecordings

> TexmlGetCallRecordingsResponseBody FetchTeXMLCallRecordings(ctx, accountSid, callSid).Execute()

Fetch recordings for a call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.FetchTeXMLCallRecordings(context.Background(), accountSid, callSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.FetchTeXMLCallRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchTeXMLCallRecordings`: TexmlGetCallRecordingsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.FetchTeXMLCallRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTeXMLCallRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TexmlGetCallRecordingsResponseBody**](TexmlGetCallRecordingsResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTeXMLConferenceRecordings

> TexmlGetCallRecordingsResponseBody FetchTeXMLConferenceRecordings(ctx, accountSid, conferenceSid).Execute()

Fetch recordings for a conference



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.FetchTeXMLConferenceRecordings(context.Background(), accountSid, conferenceSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.FetchTeXMLConferenceRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchTeXMLConferenceRecordings`: TexmlGetCallRecordingsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.FetchTeXMLConferenceRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTeXMLConferenceRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TexmlGetCallRecordingsResponseBody**](TexmlGetCallRecordingsResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeXMLCallRecording

> TexmlGetCallRecordingResponseBody GetTeXMLCallRecording(ctx, accountSid, recordingSid).Execute()

Fetch recording resource



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	recordingSid := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTeXMLCallRecording(context.Background(), accountSid, recordingSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTeXMLCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeXMLCallRecording`: TexmlGetCallRecordingResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTeXMLCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**recordingSid** | **string** | Uniquely identifies the recording by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeXMLCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TexmlGetCallRecordingResponseBody**](TexmlGetCallRecordingResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeXMLCallRecordings

> TexmlGetCallRecordingsResponseBody GetTeXMLCallRecordings(ctx, accountSid).Page(page).PageSize(pageSize).DateCreated(dateCreated).Execute()

Fetch multiple recording resources



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	page := int32(1) // int32 | The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken. (optional)
	pageSize := int32(10) // int32 | The number of records to be displayed on a page (optional)
	dateCreated := "2023-05-22" // string | Filters recording by the creation date. Expected format is ISO8601 date or date-time, ie. {YYYY}-{MM}-{DD} or {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}Z. Also accepts inequality operators, e.g. DateCreated>=2023-05-22. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTeXMLCallRecordings(context.Background(), accountSid).Page(page).PageSize(pageSize).DateCreated(dateCreated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTeXMLCallRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeXMLCallRecordings`: TexmlGetCallRecordingsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTeXMLCallRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeXMLCallRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken. | 
 **pageSize** | **int32** | The number of records to be displayed on a page | 
 **dateCreated** | **string** | Filters recording by the creation date. Expected format is ISO8601 date or date-time, ie. {YYYY}-{MM}-{DD} or {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}Z. Also accepts inequality operators, e.g. DateCreated&gt;&#x3D;2023-05-22. | 

### Return type

[**TexmlGetCallRecordingsResponseBody**](TexmlGetCallRecordingsResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeXMLRecordingTranscription

> TexmlRecordingTranscription GetTeXMLRecordingTranscription(ctx, accountSid, recordingTranscriptionSid).Execute()

Fetch a recording transcription resource



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	recordingTranscriptionSid := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording transcription by id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTeXMLRecordingTranscription(context.Background(), accountSid, recordingTranscriptionSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTeXMLRecordingTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeXMLRecordingTranscription`: TexmlRecordingTranscription
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTeXMLRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**recordingTranscriptionSid** | **string** | Uniquely identifies the recording transcription by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeXMLRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TexmlRecordingTranscription**](TexmlRecordingTranscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeXMLRecordingTranscriptions

> ListRecordingTranscriptionsResponse GetTeXMLRecordingTranscriptions(ctx, accountSid).PageToken(pageToken).PageSize(pageSize).Execute()

List recording transcriptions



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	pageToken := "pageToken_example" // string | Used to request the next page of results. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTeXMLRecordingTranscriptions(context.Background(), accountSid).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTeXMLRecordingTranscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeXMLRecordingTranscriptions`: ListRecordingTranscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTeXMLRecordingTranscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeXMLRecordingTranscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Used to request the next page of results. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**ListRecordingTranscriptionsResponse**](ListRecordingTranscriptionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlCall

> CallResource GetTexmlCall(ctx, callSid, accountSid).Execute()

Fetch a call



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
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlCall(context.Background(), callSid, accountSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlCall`: CallResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSid** | **string** | The CallSid that identifies the call to update. | 
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CallResource**](CallResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlCalls

> CallResourceIndex GetTexmlCalls(ctx, accountSid).Page(page).PageSize(pageSize).PageToken(pageToken).To(to).From(from).Status(status).StartTime(startTime).StartTimeGt(startTimeGt).StartTimeLt(startTimeLt).EndTime(endTime).EndTimeGt(endTimeGt).EndTimeLt(endTimeLt).Execute()

Fetch multiple call resources



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	page := int32(1) // int32 | The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageToken := "pageToken_example" // string | Used to request the next page of results. (optional)
	to := "+1312345678" // string | Filters calls by the to number. (optional)
	from := "+1312345678" // string | Filters calls by the from number. (optional)
	status := "no-answer" // string | Filters calls by status. (optional)
	startTime := "2023-05-22" // string | Filters calls by their start date. Expected format is YYYY-MM-DD. (optional)
	startTimeGt := "2023-05-22" // string | Filters calls by their start date (after). Expected format is YYYY-MM-DD (optional)
	startTimeLt := "2023-05-22" // string | Filters calls by their start date (before). Expected format is YYYY-MM-DD (optional)
	endTime := "2023-05-22" // string | Filters calls by their end date. Expected format is YYYY-MM-DD (optional)
	endTimeGt := "2023-05-22" // string | Filters calls by their end date (after). Expected format is YYYY-MM-DD (optional)
	endTimeLt := "2023-05-22" // string | Filters calls by their end date (before). Expected format is YYYY-MM-DD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlCalls(context.Background(), accountSid).Page(page).PageSize(pageSize).PageToken(pageToken).To(to).From(from).Status(status).StartTime(startTime).StartTimeGt(startTimeGt).StartTimeLt(startTimeLt).EndTime(endTime).EndTimeGt(endTimeGt).EndTimeLt(endTimeLt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlCalls`: CallResourceIndex
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageToken** | **string** | Used to request the next page of results. | 
 **to** | **string** | Filters calls by the to number. | 
 **from** | **string** | Filters calls by the from number. | 
 **status** | **string** | Filters calls by status. | 
 **startTime** | **string** | Filters calls by their start date. Expected format is YYYY-MM-DD. | 
 **startTimeGt** | **string** | Filters calls by their start date (after). Expected format is YYYY-MM-DD | 
 **startTimeLt** | **string** | Filters calls by their start date (before). Expected format is YYYY-MM-DD | 
 **endTime** | **string** | Filters calls by their end date. Expected format is YYYY-MM-DD | 
 **endTimeGt** | **string** | Filters calls by their end date (after). Expected format is YYYY-MM-DD | 
 **endTimeLt** | **string** | Filters calls by their end date (before). Expected format is YYYY-MM-DD | 

### Return type

[**CallResourceIndex**](CallResourceIndex.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlConference

> ConferenceResource GetTexmlConference(ctx, accountSid, conferenceSid).Execute()

Fetch a conference resource



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlConference(context.Background(), accountSid, conferenceSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlConference`: ConferenceResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConferenceResource**](ConferenceResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlConferenceParticipant

> ParticipantResource GetTexmlConferenceParticipant(ctx, accountSid, conferenceSid, callSid).Execute()

Get conference participant resource



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlConferenceParticipant(context.Background(), accountSid, conferenceSid, callSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlConferenceParticipant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlConferenceParticipant`: ParticipantResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlConferenceParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlConferenceParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ParticipantResource**](ParticipantResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlConferenceParticipants

> ParticipantResourceIndex GetTexmlConferenceParticipants(ctx, accountSid, conferenceSid).Execute()

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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlConferenceParticipants(context.Background(), accountSid, conferenceSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlConferenceParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlConferenceParticipants`: ParticipantResourceIndex
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlConferenceParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParticipantResourceIndex**](ParticipantResourceIndex.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlConferenceRecordings

> ConferenceRecordingResourceIndex GetTexmlConferenceRecordings(ctx, accountSid, conferenceSid).Execute()

List conference recordings



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlConferenceRecordings(context.Background(), accountSid, conferenceSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlConferenceRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlConferenceRecordings`: ConferenceRecordingResourceIndex
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlConferenceRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlConferenceRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConferenceRecordingResourceIndex**](ConferenceRecordingResourceIndex.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTexmlConferences

> ConferenceResourceIndex GetTexmlConferences(ctx, accountSid).Page(page).PageSize(pageSize).PageToken(pageToken).FriendlyName(friendlyName).Status(status).DateCreated(dateCreated).DateUpdated(dateUpdated).Execute()

List conference resources



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	page := int32(1) // int32 | The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken. (optional)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	pageToken := "pageToken_example" // string | Used to request the next page of results. (optional)
	friendlyName := "weekly_review_call" // string | Filters conferences by their friendly name. (optional)
	status := "in-progress" // string | Filters conferences by status. (optional)
	dateCreated := ">=2023-05-22" // string | Filters conferences by the creation date. Expected format is YYYY-MM-DD. Also accepts inequality operators, e.g. DateCreated>=2023-05-22. (optional)
	dateUpdated := ">=2023-05-22" // string | Filters conferences by the time they were last updated. Expected format is YYYY-MM-DD. Also accepts inequality operators, e.g. DateUpdated>=2023-05-22. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.GetTexmlConferences(context.Background(), accountSid).Page(page).PageSize(pageSize).PageToken(pageToken).FriendlyName(friendlyName).Status(status).DateCreated(dateCreated).DateUpdated(dateUpdated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.GetTexmlConferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTexmlConferences`: ConferenceResourceIndex
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.GetTexmlConferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTexmlConferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken. | 
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **pageToken** | **string** | Used to request the next page of results. | 
 **friendlyName** | **string** | Filters conferences by their friendly name. | 
 **status** | **string** | Filters conferences by status. | 
 **dateCreated** | **string** | Filters conferences by the creation date. Expected format is YYYY-MM-DD. Also accepts inequality operators, e.g. DateCreated&gt;&#x3D;2023-05-22. | 
 **dateUpdated** | **string** | Filters conferences by the time they were last updated. Expected format is YYYY-MM-DD. Also accepts inequality operators, e.g. DateUpdated&gt;&#x3D;2023-05-22. | 

### Return type

[**ConferenceResourceIndex**](ConferenceResourceIndex.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateTexmlCall

> InitiateCallResult InitiateTexmlCall(ctx, accountSid).InitiateCallRequest(initiateCallRequest).Execute()

Initiate an outbound call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	initiateCallRequest := *openapiclient.NewInitiateCallRequest("ApplicationSid_example", "+16175551212", "+16175551212") // InitiateCallRequest | Iniatiate Call request object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.InitiateTexmlCall(context.Background(), accountSid).InitiateCallRequest(initiateCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.InitiateTexmlCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateTexmlCall`: InitiateCallResult
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.InitiateTexmlCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateTexmlCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initiateCallRequest** | [**InitiateCallRequest**](InitiateCallRequest.md) | Iniatiate Call request object | 

### Return type

[**InitiateCallResult**](InitiateCallResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTeXMLCallRecording

> TexmlCreateCallRecordingResponseBody StartTeXMLCallRecording(ctx, accountSid, callSid).PlayBeep(playBeep).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingChannels(recordingChannels).RecordingTrack(recordingTrack).Execute()

Request recording for a call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	playBeep := true // bool | Whether to play a beep when recording is started. (optional) (default to true)
	recordingStatusCallbackEvent := "recordingStatusCallbackEvent_example" // string | The changes to the recording's state that should generate a call to `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`. (optional)
	recordingStatusCallback := "recordingStatusCallback_example" // string | Url where status callbacks will be sent. (optional)
	recordingStatusCallbackMethod := openapiclient.TexmlStatusCallbackMethod("GET") // TexmlStatusCallbackMethod |  (optional) (default to "POST")
	recordingChannels := openapiclient.TexmlRecordingChannels("single") // TexmlRecordingChannels |  (optional) (default to "dual")
	recordingTrack := openapiclient.RecordingTrack("inbound") // RecordingTrack |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.StartTeXMLCallRecording(context.Background(), accountSid, callSid).PlayBeep(playBeep).RecordingStatusCallbackEvent(recordingStatusCallbackEvent).RecordingStatusCallback(recordingStatusCallback).RecordingStatusCallbackMethod(recordingStatusCallbackMethod).RecordingChannels(recordingChannels).RecordingTrack(recordingTrack).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.StartTeXMLCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartTeXMLCallRecording`: TexmlCreateCallRecordingResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.StartTeXMLCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTeXMLCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playBeep** | **bool** | Whether to play a beep when recording is started. | [default to true]
 **recordingStatusCallbackEvent** | **string** | The changes to the recording&#39;s state that should generate a call to &#x60;RecoridngStatusCallback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Separate multiple values with a space. Defaults to &#x60;completed&#x60;. | 
 **recordingStatusCallback** | **string** | Url where status callbacks will be sent. | 
 **recordingStatusCallbackMethod** | [**TexmlStatusCallbackMethod**](TexmlStatusCallbackMethod.md) |  | [default to &quot;POST&quot;]
 **recordingChannels** | [**TexmlRecordingChannels**](TexmlRecordingChannels.md) |  | [default to &quot;dual&quot;]
 **recordingTrack** | [**RecordingTrack**](RecordingTrack.md) |  | 

### Return type

[**TexmlCreateCallRecordingResponseBody**](TexmlCreateCallRecordingResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTeXMLCallStreaming

> TexmlCreateCallStreamingResponseBody StartTeXMLCallStreaming(ctx, accountSid, callSid).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Track(track).Name(name).BidirectionalMode(bidirectionalMode).BidirectionalCodec(bidirectionalCodec).Url(url).Execute()

Start streaming media from a call.



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	statusCallback := "statusCallback_example" // string | Url where status callbacks will be sent. (optional)
	statusCallbackMethod := openapiclient.TexmlStatusCallbackMethod("GET") // TexmlStatusCallbackMethod |  (optional) (default to "POST")
	track := openapiclient.StreamTrack("inbound_track") // StreamTrack |  (optional) (default to "inbound_track")
	name := "name_example" // string | The user specified name of Stream. (optional)
	bidirectionalMode := openapiclient.TexmlBidirectionalStreamMode("mp3") // TexmlBidirectionalStreamMode |  (optional) (default to "mp3")
	bidirectionalCodec := openapiclient.TexmlBidirectionalStreamCodec("PCMU") // TexmlBidirectionalStreamCodec |  (optional) (default to "PCMU")
	url := "url_example" // string | The destination WebSocket address where the stream is going to be delivered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.StartTeXMLCallStreaming(context.Background(), accountSid, callSid).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Track(track).Name(name).BidirectionalMode(bidirectionalMode).BidirectionalCodec(bidirectionalCodec).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.StartTeXMLCallStreaming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartTeXMLCallStreaming`: TexmlCreateCallStreamingResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.StartTeXMLCallStreaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTeXMLCallStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **statusCallback** | **string** | Url where status callbacks will be sent. | 
 **statusCallbackMethod** | [**TexmlStatusCallbackMethod**](TexmlStatusCallbackMethod.md) |  | [default to &quot;POST&quot;]
 **track** | [**StreamTrack**](StreamTrack.md) |  | [default to &quot;inbound_track&quot;]
 **name** | **string** | The user specified name of Stream. | 
 **bidirectionalMode** | [**TexmlBidirectionalStreamMode**](TexmlBidirectionalStreamMode.md) |  | [default to &quot;mp3&quot;]
 **bidirectionalCodec** | [**TexmlBidirectionalStreamCodec**](TexmlBidirectionalStreamCodec.md) |  | [default to &quot;PCMU&quot;]
 **url** | **string** | The destination WebSocket address where the stream is going to be delivered. | 

### Return type

[**TexmlCreateCallStreamingResponseBody**](TexmlCreateCallStreamingResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTeXMLSiprecSession

> TexmlCreateSiprecSessionResponseBody StartTeXMLSiprecSession(ctx, accountSid, callSid).ConnectorName(connectorName).Name(name).Track(track).IncludeMetadataCustomHeaders(includeMetadataCustomHeaders).Secure(secure).SessionTimeoutSecs(sessionTimeoutSecs).SipTransport(sipTransport).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Execute()

Request siprec session for a call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	connectorName := "connectorName_example" // string | The name of the connector to use for the SIPREC session. (optional)
	name := "name_example" // string | Name of the SIPREC session. May be used to stop the SIPREC session from TeXML instruction. (optional)
	track := "track_example" // string | The track to be used for siprec session. Can be `both_tracks`, `inbound_track` or `outbound_track`. Defaults to `both_tracks`. (optional)
	includeMetadataCustomHeaders := true // bool | When set, custom parameters will be added as metadata (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip headers. (optional)
	secure := true // bool | Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set you need to configure SRS port in your connector to 5061. (optional)
	sessionTimeoutSecs := int32(56) // int32 | Sets `Session-Expires` header to the INVITE. A reinvite is sent every half the value set. Usefull for session keep alive. Minimum value is 90, set to 0 to disable. (optional) (default to 1800)
	sipTransport := "sipTransport_example" // string | Specifies SIP transport protocol. (optional) (default to "udp")
	statusCallback := "statusCallback_example" // string | URL destination for Telnyx to send status callback events to for the siprec session. (optional)
	statusCallbackMethod := "statusCallbackMethod_example" // string | HTTP request type used for `StatusCallback`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.StartTeXMLSiprecSession(context.Background(), accountSid, callSid).ConnectorName(connectorName).Name(name).Track(track).IncludeMetadataCustomHeaders(includeMetadataCustomHeaders).Secure(secure).SessionTimeoutSecs(sessionTimeoutSecs).SipTransport(sipTransport).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.StartTeXMLSiprecSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartTeXMLSiprecSession`: TexmlCreateSiprecSessionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.StartTeXMLSiprecSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTeXMLSiprecSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectorName** | **string** | The name of the connector to use for the SIPREC session. | 
 **name** | **string** | Name of the SIPREC session. May be used to stop the SIPREC session from TeXML instruction. | 
 **track** | **string** | The track to be used for siprec session. Can be &#x60;both_tracks&#x60;, &#x60;inbound_track&#x60; or &#x60;outbound_track&#x60;. Defaults to &#x60;both_tracks&#x60;. | 
 **includeMetadataCustomHeaders** | **bool** | When set, custom parameters will be added as metadata (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip headers. | 
 **secure** | **bool** | Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set you need to configure SRS port in your connector to 5061. | 
 **sessionTimeoutSecs** | **int32** | Sets &#x60;Session-Expires&#x60; header to the INVITE. A reinvite is sent every half the value set. Usefull for session keep alive. Minimum value is 90, set to 0 to disable. | [default to 1800]
 **sipTransport** | **string** | Specifies SIP transport protocol. | [default to &quot;udp&quot;]
 **statusCallback** | **string** | URL destination for Telnyx to send status callback events to for the siprec session. | 
 **statusCallbackMethod** | **string** | HTTP request type used for &#x60;StatusCallback&#x60;. | 

### Return type

[**TexmlCreateSiprecSessionResponseBody**](TexmlCreateSiprecSessionResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeXMLCallRecording

> TexmlCreateCallRecordingResponseBody UpdateTeXMLCallRecording(ctx, accountSid, callSid, recordingSid).Status(status).Execute()

Update recording on a call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	recordingSid := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the recording by id.
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.UpdateTeXMLCallRecording(context.Background(), accountSid, callSid, recordingSid).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.UpdateTeXMLCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeXMLCallRecording`: TexmlCreateCallRecordingResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.UpdateTeXMLCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 
**recordingSid** | **string** | Uniquely identifies the recording by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeXMLCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** |  | 

### Return type

[**TexmlCreateCallRecordingResponseBody**](TexmlCreateCallRecordingResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeXMLCallStreaming

> TexmlUpdateCallStreamingResponseBody UpdateTeXMLCallStreaming(ctx, accountSid, callSid, streamingSid).Status(status).Execute()

Update streaming on a call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	streamingSid := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Uniquely identifies the streaming by id.
	status := openapiclient.StreamStatus("stopped") // StreamStatus |  (optional) (default to "stopped")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.UpdateTeXMLCallStreaming(context.Background(), accountSid, callSid, streamingSid).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.UpdateTeXMLCallStreaming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeXMLCallStreaming`: TexmlUpdateCallStreamingResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.UpdateTeXMLCallStreaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 
**streamingSid** | **string** | Uniquely identifies the streaming by id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeXMLCallStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | [**StreamStatus**](StreamStatus.md) |  | [default to &quot;stopped&quot;]

### Return type

[**TexmlUpdateCallStreamingResponseBody**](TexmlUpdateCallStreamingResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeXMLSiprecSession

> TexmlUpdateSiprecSessionResponseBody UpdateTeXMLSiprecSession(ctx, accountSid, callSid, siprecSid).Status(status).Execute()

Updates siprec session for a call



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	siprecSid := "siprecSid_example" // string | The SiprecSid that uniquely identifies the Sip Recording.
	status := "status_example" // string | The new status of the resource. Specifying `stopped` will end the siprec session. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.UpdateTeXMLSiprecSession(context.Background(), accountSid, callSid, siprecSid).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.UpdateTeXMLSiprecSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeXMLSiprecSession`: TexmlUpdateSiprecSessionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.UpdateTeXMLSiprecSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 
**siprecSid** | **string** | The SiprecSid that uniquely identifies the Sip Recording. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeXMLSiprecSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** | The new status of the resource. Specifying &#x60;stopped&#x60; will end the siprec session. | 

### Return type

[**TexmlUpdateSiprecSessionResponseBody**](TexmlUpdateSiprecSessionResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTexmlCall

> CallResource UpdateTexmlCall(ctx, callSid, accountSid).Status(status).Url(url).Method(method).FallbackUrl(fallbackUrl).FallbackMethod(fallbackMethod).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Texml(texml).Execute()

Update call



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
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	status := "status_example" // string | The value to set the call status to. Setting the status to completed ends the call. (optional)
	url := "url_example" // string | The URL where TeXML will make a request to retrieve a new set of TeXML instructions to continue the call flow. (optional)
	method := "method_example" // string | HTTP request type used for `Url`. (optional)
	fallbackUrl := "fallbackUrl_example" // string | A failover URL for which Telnyx will retrieve the TeXML call instructions if the Url is not responding. (optional)
	fallbackMethod := "fallbackMethod_example" // string | HTTP request type used for `FallbackUrl`. (optional)
	statusCallback := "statusCallback_example" // string | URL destination for Telnyx to send status callback events to for the call. (optional)
	statusCallbackMethod := "statusCallbackMethod_example" // string | HTTP request type used for `StatusCallback`. (optional)
	texml := "texml_example" // string | TeXML to replace the current one with. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.UpdateTexmlCall(context.Background(), callSid, accountSid).Status(status).Url(url).Method(method).FallbackUrl(fallbackUrl).FallbackMethod(fallbackMethod).StatusCallback(statusCallback).StatusCallbackMethod(statusCallbackMethod).Texml(texml).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.UpdateTexmlCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTexmlCall`: CallResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.UpdateTexmlCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSid** | **string** | The CallSid that identifies the call to update. | 
**accountSid** | **string** | The id of the account the resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTexmlCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | The value to set the call status to. Setting the status to completed ends the call. | 
 **url** | **string** | The URL where TeXML will make a request to retrieve a new set of TeXML instructions to continue the call flow. | 
 **method** | **string** | HTTP request type used for &#x60;Url&#x60;. | 
 **fallbackUrl** | **string** | A failover URL for which Telnyx will retrieve the TeXML call instructions if the Url is not responding. | 
 **fallbackMethod** | **string** | HTTP request type used for &#x60;FallbackUrl&#x60;. | 
 **statusCallback** | **string** | URL destination for Telnyx to send status callback events to for the call. | 
 **statusCallbackMethod** | **string** | HTTP request type used for &#x60;StatusCallback&#x60;. | 
 **texml** | **string** | TeXML to replace the current one with. | 

### Return type

[**CallResource**](CallResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTexmlConference

> ConferenceResource UpdateTexmlConference(ctx, accountSid, conferenceSid).Status(status).AnnounceUrl(announceUrl).AnnounceMethod(announceMethod).Execute()

Update a conference resource



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.
	status := "status_example" // string | The new status of the resource. Specifying `completed` will end the conference and hang up all participants. (optional)
	announceUrl := "announceUrl_example" // string | The URL we should call to announce something into the conference. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs. (optional)
	announceMethod := "announceMethod_example" // string | The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.UpdateTexmlConference(context.Background(), accountSid, conferenceSid).Status(status).AnnounceUrl(announceUrl).AnnounceMethod(announceMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.UpdateTexmlConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTexmlConference`: ConferenceResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.UpdateTexmlConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTexmlConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | The new status of the resource. Specifying &#x60;completed&#x60; will end the conference and hang up all participants. | 
 **announceUrl** | **string** | The URL we should call to announce something into the conference. The URL may return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60;, &#x60;&lt;Pause&gt;&#x60;, or &#x60;&lt;Redirect&gt;&#x60; verbs. | 
 **announceMethod** | **string** | The HTTP method used to call the &#x60;AnnounceUrl&#x60;. Defaults to &#x60;POST&#x60;. | 

### Return type

[**ConferenceResource**](ConferenceResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTexmlConferenceParticipant

> ParticipantResource UpdateTexmlConferenceParticipant(ctx, accountSid, conferenceSid, callSid).Muted(muted).Hold(hold).HoldUrl(holdUrl).HoldMethod(holdMethod).AnnounceUrl(announceUrl).AnnounceMethod(announceMethod).WaitUrl(waitUrl).BeepOnExit(beepOnExit).EndConferenceOnExit(endConferenceOnExit).Coaching(coaching).CallSidToCoach(callSidToCoach).Execute()

Update a conference participant



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
	accountSid := "accountSid_example" // string | The id of the account the resource belongs to.
	conferenceSid := "conferenceSid_example" // string | The ConferenceSid that uniquely identifies a conference.
	callSid := "callSid_example" // string | The CallSid that identifies the call to update.
	muted := true // bool | Whether the participant should be muted. (optional)
	hold := true // bool | Whether the participant should be on hold. (optional)
	holdUrl := "holdUrl_example" // string | The URL to be called using the `HoldMethod` for music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs. (optional)
	holdMethod := "holdMethod_example" // string | The HTTP method to use when calling the `HoldUrl`. (optional)
	announceUrl := "announceUrl_example" // string | The URL to call to announce something to the participant. The URL may return an MP3 fileo a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs. (optional)
	announceMethod := "announceMethod_example" // string | The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`. (optional)
	waitUrl := "waitUrl_example" // string | The URL to call for an audio file to play while the participant is waiting for the conference to start. (optional)
	beepOnExit := true // bool | Whether to play a notification beep to the conference when the participant exits. (optional)
	endConferenceOnExit := true // bool | Whether to end the conference when the participant leaves. (optional)
	coaching := true // bool | Whether the participant is coaching another call. When `true`, `CallSidToCoach` has to be given. (optional)
	callSidToCoach := "callSidToCoach_example" // string | The SID of the participant who is being coached. The participant being coached is the only participant who can hear the participant who is coaching. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeXMLRESTCommandsAPI.UpdateTexmlConferenceParticipant(context.Background(), accountSid, conferenceSid, callSid).Muted(muted).Hold(hold).HoldUrl(holdUrl).HoldMethod(holdMethod).AnnounceUrl(announceUrl).AnnounceMethod(announceMethod).WaitUrl(waitUrl).BeepOnExit(beepOnExit).EndConferenceOnExit(endConferenceOnExit).Coaching(coaching).CallSidToCoach(callSidToCoach).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeXMLRESTCommandsAPI.UpdateTexmlConferenceParticipant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTexmlConferenceParticipant`: ParticipantResource
	fmt.Fprintf(os.Stdout, "Response from `TeXMLRESTCommandsAPI.UpdateTexmlConferenceParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountSid** | **string** | The id of the account the resource belongs to. | 
**conferenceSid** | **string** | The ConferenceSid that uniquely identifies a conference. | 
**callSid** | **string** | The CallSid that identifies the call to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTexmlConferenceParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **muted** | **bool** | Whether the participant should be muted. | 
 **hold** | **bool** | Whether the participant should be on hold. | 
 **holdUrl** | **string** | The URL to be called using the &#x60;HoldMethod&#x60; for music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60;, &#x60;&lt;Pause&gt;&#x60;, or &#x60;&lt;Redirect&gt;&#x60; verbs. | 
 **holdMethod** | **string** | The HTTP method to use when calling the &#x60;HoldUrl&#x60;. | 
 **announceUrl** | **string** | The URL to call to announce something to the participant. The URL may return an MP3 fileo a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60;, &#x60;&lt;Pause&gt;&#x60;, or &#x60;&lt;Redirect&gt;&#x60; verbs. | 
 **announceMethod** | **string** | The HTTP method used to call the &#x60;AnnounceUrl&#x60;. Defaults to &#x60;POST&#x60;. | 
 **waitUrl** | **string** | The URL to call for an audio file to play while the participant is waiting for the conference to start. | 
 **beepOnExit** | **bool** | Whether to play a notification beep to the conference when the participant exits. | 
 **endConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. | 
 **coaching** | **bool** | Whether the participant is coaching another call. When &#x60;true&#x60;, &#x60;CallSidToCoach&#x60; has to be given. | 
 **callSidToCoach** | **string** | The SID of the participant who is being coached. The participant being coached is the only participant who can hear the participant who is coaching. | 

### Return type

[**ParticipantResource**](ParticipantResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

