// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// CallActionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallActionService] method instead.
type CallActionService struct {
	Options []option.RequestOption
}

// NewCallActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCallActionService(opts ...option.RequestOption) (r CallActionService) {
	r = CallActionService{}
	r.Options = opts
	return
}

// Answer an incoming call. You must issue this command before executing subsequent
// commands on an incoming call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/answer-call#callbacks)
// below):**
//
//   - `call.answered`
//   - `streaming.started`, `streaming.stopped` or `streaming.failed` if `stream_url`
//     was set
//
// When the `record` parameter is set to `record-from-answer`, the response will
// include a `recording_id` field.
func (r *CallActionService) Answer(ctx context.Context, callControlID string, body CallActionAnswerParams, opts ...option.RequestOption) (res *CallActionAnswerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/answer", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bridge two call control calls.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/bridge-call#callbacks)
// below):**
//
// - `call.bridged` for Leg A
// - `call.bridged` for Leg B
func (r *CallActionService) Bridge(ctx context.Context, callControlID string, body CallActionBridgeParams, opts ...option.RequestOption) (res *CallActionBridgeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/bridge", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Put the call in a queue.
func (r *CallActionService) Enqueue(ctx context.Context, callControlID string, body CallActionEnqueueParams, opts ...option.RequestOption) (res *CallActionEnqueueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/enqueue", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gather DTMF signals to build interactive menus.
//
// You can pass a list of valid digits. The `Answer` command must be issued before
// the `gather` command.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/gather-call#callbacks)
// below):**
//
// - `call.dtmf.received` (you may receive many of these webhooks)
// - `call.gather.ended`
func (r *CallActionService) Gather(ctx context.Context, callControlID string, body CallActionGatherParams, opts ...option.RequestOption) (res *CallActionGatherResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/gather", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gather parameters defined in the request payload using a voice assistant.
//
// You can pass parameters described as a JSON Schema object and the voice
// assistant will attempt to gather these informations.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/call-gather-using-ai#callbacks)
// below):**
//
//   - `call.ai_gather.ended`
//   - `call.conversation.ended`
//   - `call.ai_gather.partial_results` (if `send_partial_results` is set to `true`)
//   - `call.ai_gather.message_history_updated` (if `send_message_history_updates` is
//     set to `true`)
func (r *CallActionService) GatherUsingAI(ctx context.Context, callControlID string, body CallActionGatherUsingAIParams, opts ...option.RequestOption) (res *CallActionGatherUsingAIResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/gather_using_ai", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Play an audio file on the call until the required DTMF signals are gathered to
// build interactive menus.
//
// You can pass a list of valid digits along with an 'invalid_audio_url', which
// will be played back at the beginning of each prompt. Playback will be
// interrupted when a DTMF signal is received. The
// `Answer command must be issued before the `gather_using_audio` command.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/gather-using-audio#callbacks)
// below):**
//
// - `call.playback.started`
// - `call.playback.ended`
// - `call.dtmf.received` (you may receive many of these webhooks)
// - `call.gather.ended`
func (r *CallActionService) GatherUsingAudio(ctx context.Context, callControlID string, body CallActionGatherUsingAudioParams, opts ...option.RequestOption) (res *CallActionGatherUsingAudioResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/gather_using_audio", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Convert text to speech and play it on the call until the required DTMF signals
// are gathered to build interactive menus.
//
// You can pass a list of valid digits along with an 'invalid_payload', which will
// be played back at the beginning of each prompt. Speech will be interrupted when
// a DTMF signal is received. The `Answer` command must be issued before the
// `gather_using_speak` command.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/gather-using-speak#callbacks)
// below):**
//
// - `call.dtmf.received` (you may receive many of these webhooks)
// - `call.gather.ended`
func (r *CallActionService) GatherUsingSpeak(ctx context.Context, callControlID string, body CallActionGatherUsingSpeakParams, opts ...option.RequestOption) (res *CallActionGatherUsingSpeakResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/gather_using_speak", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Hang up the call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/hangup-call#callbacks)
// below):**
//
// - `call.hangup`
// - `call.recording.saved`
func (r *CallActionService) Hangup(ctx context.Context, callControlID string, body CallActionHangupParams, opts ...option.RequestOption) (res *CallActionHangupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/hangup", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes the call from a queue.
func (r *CallActionService) LeaveQueue(ctx context.Context, callControlID string, body CallActionLeaveQueueParams, opts ...option.RequestOption) (res *CallActionLeaveQueueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/leave_queue", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pause recording the call. Recording can be resumed via Resume recording command.
//
// **Expected Webhooks:**
//
// There are no webhooks associated with this command.
func (r *CallActionService) PauseRecording(ctx context.Context, callControlID string, body CallActionPauseRecordingParams, opts ...option.RequestOption) (res *CallActionPauseRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/record_pause", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate a SIP Refer on a Call Control call. You can initiate a SIP Refer at any
// point in the duration of a call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/refer-call#callbacks)
// below):**
//
// - `call.refer.started`
// - `call.refer.completed`
// - `call.refer.failed`
func (r *CallActionService) Refer(ctx context.Context, callControlID string, body CallActionReferParams, opts ...option.RequestOption) (res *CallActionReferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/refer", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reject an incoming call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/reject-call#callbacks)
// below):**
//
// - `call.hangup`
func (r *CallActionService) Reject(ctx context.Context, callControlID string, body CallActionRejectParams, opts ...option.RequestOption) (res *CallActionRejectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/reject", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Resume recording the call.
//
// **Expected Webhooks:**
//
// There are no webhooks associated with this command.
func (r *CallActionService) ResumeRecording(ctx context.Context, callControlID string, body CallActionResumeRecordingParams, opts ...option.RequestOption) (res *CallActionResumeRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/record_resume", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sends DTMF tones from this leg. DTMF tones will be heard by the other end of the
// call.
//
// **Expected Webhooks:**
//
// There are no webhooks associated with this command.
func (r *CallActionService) SendDtmf(ctx context.Context, callControlID string, body CallActionSendDtmfParams, opts ...option.RequestOption) (res *CallActionSendDtmfResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/send_dtmf", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sends SIP info from this leg.
//
// **Expected Webhooks:**
//
// - `call.sip_info.received` (to be received on the target call leg)
func (r *CallActionService) SendSipInfo(ctx context.Context, callControlID string, body CallActionSendSipInfoParams, opts ...option.RequestOption) (res *CallActionSendSipInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/send_sip_info", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Convert text to speech and play it back on the call. If multiple speak text
// commands are issued consecutively, the audio files will be placed in a queue
// awaiting playback.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/speak-call#callbacks)
// below):**
//
// - `call.speak.started`
// - `call.speak.ended`
func (r *CallActionService) Speak(ctx context.Context, callControlID string, body CallActionSpeakParams, opts ...option.RequestOption) (res *CallActionSpeakResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/speak", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start an AI assistant on the call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/call-start-ai-assistant#callbacks)
// below):**
//
// - `call.conversation.ended`
// - `call.conversation_insights.generated`
func (r *CallActionService) StartAIAssistant(ctx context.Context, callControlID string, body CallActionStartAIAssistantParams, opts ...option.RequestOption) (res *CallActionStartAIAssistantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/ai_assistant_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Call forking allows you to stream the media from a call to a specific target in
// realtime. This stream can be used to enable realtime audio analysis to support a
// variety of use cases, including fraud detection, or the creation of AI-generated
// audio responses. Requests must specify either the `target` attribute or the `rx`
// and `tx` attributes.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/start-call-fork#callbacks)
// below):**
//
// - `call.fork.started`
// - `call.fork.stopped`
func (r *CallActionService) StartForking(ctx context.Context, callControlID string, body CallActionStartForkingParams, opts ...option.RequestOption) (res *CallActionStartForkingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/fork_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Noise Suppression Start (BETA)
func (r *CallActionService) StartNoiseSuppression(ctx context.Context, callControlID string, body CallActionStartNoiseSuppressionParams, opts ...option.RequestOption) (res *CallActionStartNoiseSuppressionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/suppression_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Play an audio file on the call. If multiple play audio commands are issued
// consecutively, the audio files will be placed in a queue awaiting playback.
//
// _Notes:_
//
//   - When `overlay` is enabled, `target_legs` is limited to `self`.
//   - A customer cannot Play Audio with `overlay=true` unless there is a Play Audio
//     with `overlay=false` actively playing.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/start-call-playback#callbacks)
// below):**
//
// - `call.playback.started`
// - `call.playback.ended`
func (r *CallActionService) StartPlayback(ctx context.Context, callControlID string, body CallActionStartPlaybackParams, opts ...option.RequestOption) (res *CallActionStartPlaybackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/playback_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start recording the call. Recording will stop on call hang-up, or can be
// initiated via the Stop Recording command.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/start-call-record#callbacks)
// below):**
//
// - `call.recording.saved`
// - `call.recording.transcription.saved`
// - `call.recording.error`
func (r *CallActionService) StartRecording(ctx context.Context, callControlID string, body CallActionStartRecordingParams, opts ...option.RequestOption) (res *CallActionStartRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/record_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start siprec session to configured in SIPREC connector SRS.
//
// **Expected Webhooks:**
//
// - `siprec.started`
// - `siprec.stopped`
// - `siprec.failed`
func (r *CallActionService) StartSiprec(ctx context.Context, callControlID string, body CallActionStartSiprecParams, opts ...option.RequestOption) (res *CallActionStartSiprecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/siprec_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start streaming the media from a call to a specific WebSocket address or
// Dialogflow connection in near-realtime. Audio will be delivered as
// base64-encoded RTP payload (raw audio), wrapped in JSON payloads.
//
// Please find more details about media streaming messages specification under the
// [link](https://developers.telnyx.com/docs/voice/programmable-voice/media-streaming).
func (r *CallActionService) StartStreaming(ctx context.Context, callControlID string, body CallActionStartStreamingParams, opts ...option.RequestOption) (res *CallActionStartStreamingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/streaming_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start real-time transcription. Transcription will stop on call hang-up, or can
// be initiated via the Transcription stop command.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/start-call-transcription#callbacks)
// below):**
//
// - `call.transcription`
func (r *CallActionService) StartTranscription(ctx context.Context, callControlID string, body CallActionStartTranscriptionParams, opts ...option.RequestOption) (res *CallActionStartTranscriptionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/transcription_start", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop an AI assistant on the call.
func (r *CallActionService) StopAIAssistant(ctx context.Context, callControlID string, body CallActionStopAIAssistantParams, opts ...option.RequestOption) (res *CallActionStopAIAssistantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/ai_assistant_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop forking a call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/stop-call-fork#callbacks)
// below):**
//
// - `call.fork.stopped`
func (r *CallActionService) StopForking(ctx context.Context, callControlID string, body CallActionStopForkingParams, opts ...option.RequestOption) (res *CallActionStopForkingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/fork_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop current gather.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/stop-call-gather#callbacks)
// below):**
//
// - `call.gather.ended`
func (r *CallActionService) StopGather(ctx context.Context, callControlID string, body CallActionStopGatherParams, opts ...option.RequestOption) (res *CallActionStopGatherResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/gather_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Noise Suppression Stop (BETA)
func (r *CallActionService) StopNoiseSuppression(ctx context.Context, callControlID string, body CallActionStopNoiseSuppressionParams, opts ...option.RequestOption) (res *CallActionStopNoiseSuppressionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/suppression_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop audio being played on the call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/stop-call-playback#callbacks)
// below):**
//
// - `call.playback.ended` or `call.speak.ended`
func (r *CallActionService) StopPlayback(ctx context.Context, callControlID string, body CallActionStopPlaybackParams, opts ...option.RequestOption) (res *CallActionStopPlaybackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/playback_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop recording the call.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/stop-call-recording#callbacks)
// below):**
//
// - `call.recording.saved`
func (r *CallActionService) StopRecording(ctx context.Context, callControlID string, body CallActionStopRecordingParams, opts ...option.RequestOption) (res *CallActionStopRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/record_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop SIPREC session.
//
// **Expected Webhooks:**
//
// - `siprec.stopped`
func (r *CallActionService) StopSiprec(ctx context.Context, callControlID string, body CallActionStopSiprecParams, opts ...option.RequestOption) (res *CallActionStopSiprecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/siprec_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop streaming a call to a WebSocket.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/stop-call-streaming#callbacks)
// below):**
//
// - `streaming.stopped`
func (r *CallActionService) StopStreaming(ctx context.Context, callControlID string, body CallActionStopStreamingParams, opts ...option.RequestOption) (res *CallActionStopStreamingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/streaming_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop real-time transcription.
func (r *CallActionService) StopTranscription(ctx context.Context, callControlID string, body CallActionStopTranscriptionParams, opts ...option.RequestOption) (res *CallActionStopTranscriptionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/transcription_stop", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Switch the supervisor role for a bridged call. This allows switching between
// different supervisor modes during an active call
func (r *CallActionService) SwitchSupervisorRole(ctx context.Context, callControlID string, body CallActionSwitchSupervisorRoleParams, opts ...option.RequestOption) (res *CallActionSwitchSupervisorRoleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/switch_supervisor_role", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Transfer a call to a new destination. If the transfer is unsuccessful, a
// `call.hangup` webhook for the other call (Leg B) will be sent indicating that
// the transfer could not be completed. The original call will remain active and
// may be issued additional commands, potentially transfering the call to an
// alternate destination.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/transfer-call#callbacks)
// below):**
//
//   - `call.initiated`
//   - `call.bridged` to Leg B
//   - `call.answered` or `call.hangup`
//   - `call.machine.detection.ended` if `answering_machine_detection` was requested
//   - `call.machine.greeting.ended` if `answering_machine_detection` was requested
//     to detect the end of machine greeting
//   - `call.machine.premium.detection.ended` if
//     `answering_machine_detection=premium` was requested
//   - `call.machine.premium.greeting.ended` if `answering_machine_detection=premium`
//     was requested and a beep was detected
func (r *CallActionService) Transfer(ctx context.Context, callControlID string, body CallActionTransferParams, opts ...option.RequestOption) (res *CallActionTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/transfer", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates client state
func (r *CallActionService) UpdateClientState(ctx context.Context, callControlID string, body CallActionUpdateClientStateParams, opts ...option.RequestOption) (res *CallActionUpdateClientStateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s/actions/client_state_update", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CallControlCommandResult struct {
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlCommandResult) RawJSON() string { return r.JSON.raw }
func (r *CallControlCommandResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElevenLabsVoiceSettingsParam struct {
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	paramObj
}

func (r ElevenLabsVoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ElevenLabsVoiceSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElevenLabsVoiceSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Language to use for speech recognition
type GoogleTranscriptionLanguage string

const (
	GoogleTranscriptionLanguageAf  GoogleTranscriptionLanguage = "af"
	GoogleTranscriptionLanguageSq  GoogleTranscriptionLanguage = "sq"
	GoogleTranscriptionLanguageAm  GoogleTranscriptionLanguage = "am"
	GoogleTranscriptionLanguageAr  GoogleTranscriptionLanguage = "ar"
	GoogleTranscriptionLanguageHy  GoogleTranscriptionLanguage = "hy"
	GoogleTranscriptionLanguageAz  GoogleTranscriptionLanguage = "az"
	GoogleTranscriptionLanguageEu  GoogleTranscriptionLanguage = "eu"
	GoogleTranscriptionLanguageBn  GoogleTranscriptionLanguage = "bn"
	GoogleTranscriptionLanguageBs  GoogleTranscriptionLanguage = "bs"
	GoogleTranscriptionLanguageBg  GoogleTranscriptionLanguage = "bg"
	GoogleTranscriptionLanguageMy  GoogleTranscriptionLanguage = "my"
	GoogleTranscriptionLanguageCa  GoogleTranscriptionLanguage = "ca"
	GoogleTranscriptionLanguageYue GoogleTranscriptionLanguage = "yue"
	GoogleTranscriptionLanguageZh  GoogleTranscriptionLanguage = "zh"
	GoogleTranscriptionLanguageHr  GoogleTranscriptionLanguage = "hr"
	GoogleTranscriptionLanguageCs  GoogleTranscriptionLanguage = "cs"
	GoogleTranscriptionLanguageDa  GoogleTranscriptionLanguage = "da"
	GoogleTranscriptionLanguageNl  GoogleTranscriptionLanguage = "nl"
	GoogleTranscriptionLanguageEn  GoogleTranscriptionLanguage = "en"
	GoogleTranscriptionLanguageEt  GoogleTranscriptionLanguage = "et"
	GoogleTranscriptionLanguageFil GoogleTranscriptionLanguage = "fil"
	GoogleTranscriptionLanguageFi  GoogleTranscriptionLanguage = "fi"
	GoogleTranscriptionLanguageFr  GoogleTranscriptionLanguage = "fr"
	GoogleTranscriptionLanguageGl  GoogleTranscriptionLanguage = "gl"
	GoogleTranscriptionLanguageKa  GoogleTranscriptionLanguage = "ka"
	GoogleTranscriptionLanguageDe  GoogleTranscriptionLanguage = "de"
	GoogleTranscriptionLanguageEl  GoogleTranscriptionLanguage = "el"
	GoogleTranscriptionLanguageGu  GoogleTranscriptionLanguage = "gu"
	GoogleTranscriptionLanguageIw  GoogleTranscriptionLanguage = "iw"
	GoogleTranscriptionLanguageHi  GoogleTranscriptionLanguage = "hi"
	GoogleTranscriptionLanguageHu  GoogleTranscriptionLanguage = "hu"
	GoogleTranscriptionLanguageIs  GoogleTranscriptionLanguage = "is"
	GoogleTranscriptionLanguageID  GoogleTranscriptionLanguage = "id"
	GoogleTranscriptionLanguageIt  GoogleTranscriptionLanguage = "it"
	GoogleTranscriptionLanguageJa  GoogleTranscriptionLanguage = "ja"
	GoogleTranscriptionLanguageJv  GoogleTranscriptionLanguage = "jv"
	GoogleTranscriptionLanguageKn  GoogleTranscriptionLanguage = "kn"
	GoogleTranscriptionLanguageKk  GoogleTranscriptionLanguage = "kk"
	GoogleTranscriptionLanguageKm  GoogleTranscriptionLanguage = "km"
	GoogleTranscriptionLanguageKo  GoogleTranscriptionLanguage = "ko"
	GoogleTranscriptionLanguageLo  GoogleTranscriptionLanguage = "lo"
	GoogleTranscriptionLanguageLv  GoogleTranscriptionLanguage = "lv"
	GoogleTranscriptionLanguageLt  GoogleTranscriptionLanguage = "lt"
	GoogleTranscriptionLanguageMk  GoogleTranscriptionLanguage = "mk"
	GoogleTranscriptionLanguageMs  GoogleTranscriptionLanguage = "ms"
	GoogleTranscriptionLanguageMl  GoogleTranscriptionLanguage = "ml"
	GoogleTranscriptionLanguageMr  GoogleTranscriptionLanguage = "mr"
	GoogleTranscriptionLanguageMn  GoogleTranscriptionLanguage = "mn"
	GoogleTranscriptionLanguageNe  GoogleTranscriptionLanguage = "ne"
	GoogleTranscriptionLanguageNo  GoogleTranscriptionLanguage = "no"
	GoogleTranscriptionLanguageFa  GoogleTranscriptionLanguage = "fa"
	GoogleTranscriptionLanguagePl  GoogleTranscriptionLanguage = "pl"
	GoogleTranscriptionLanguagePt  GoogleTranscriptionLanguage = "pt"
	GoogleTranscriptionLanguagePa  GoogleTranscriptionLanguage = "pa"
	GoogleTranscriptionLanguageRo  GoogleTranscriptionLanguage = "ro"
	GoogleTranscriptionLanguageRu  GoogleTranscriptionLanguage = "ru"
	GoogleTranscriptionLanguageRw  GoogleTranscriptionLanguage = "rw"
	GoogleTranscriptionLanguageSr  GoogleTranscriptionLanguage = "sr"
	GoogleTranscriptionLanguageSi  GoogleTranscriptionLanguage = "si"
	GoogleTranscriptionLanguageSk  GoogleTranscriptionLanguage = "sk"
	GoogleTranscriptionLanguageSl  GoogleTranscriptionLanguage = "sl"
	GoogleTranscriptionLanguageSS  GoogleTranscriptionLanguage = "ss"
	GoogleTranscriptionLanguageSt  GoogleTranscriptionLanguage = "st"
	GoogleTranscriptionLanguageEs  GoogleTranscriptionLanguage = "es"
	GoogleTranscriptionLanguageSu  GoogleTranscriptionLanguage = "su"
	GoogleTranscriptionLanguageSw  GoogleTranscriptionLanguage = "sw"
	GoogleTranscriptionLanguageSv  GoogleTranscriptionLanguage = "sv"
	GoogleTranscriptionLanguageTa  GoogleTranscriptionLanguage = "ta"
	GoogleTranscriptionLanguageTe  GoogleTranscriptionLanguage = "te"
	GoogleTranscriptionLanguageTh  GoogleTranscriptionLanguage = "th"
	GoogleTranscriptionLanguageTn  GoogleTranscriptionLanguage = "tn"
	GoogleTranscriptionLanguageTr  GoogleTranscriptionLanguage = "tr"
	GoogleTranscriptionLanguageTs  GoogleTranscriptionLanguage = "ts"
	GoogleTranscriptionLanguageUk  GoogleTranscriptionLanguage = "uk"
	GoogleTranscriptionLanguageUr  GoogleTranscriptionLanguage = "ur"
	GoogleTranscriptionLanguageUz  GoogleTranscriptionLanguage = "uz"
	GoogleTranscriptionLanguageVe  GoogleTranscriptionLanguage = "ve"
	GoogleTranscriptionLanguageVi  GoogleTranscriptionLanguage = "vi"
	GoogleTranscriptionLanguageXh  GoogleTranscriptionLanguage = "xh"
	GoogleTranscriptionLanguageZu  GoogleTranscriptionLanguage = "zu"
)

// Settings for handling user interruptions during assistant speech
type InterruptionSettingsParam struct {
	// When true, allows users to interrupt the assistant while speaking
	Enable param.Opt[bool] `json:"enable,omitzero"`
	paramObj
}

func (r InterruptionSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow InterruptionSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterruptionSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LoopcountUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u LoopcountUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LoopcountUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LoopcountUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type StopRecordingRequestParam struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Uniquely identifies the resource.
	RecordingID param.Opt[string] `json:"recording_id,omitzero" format:"uuid"`
	paramObj
}

func (r StopRecordingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow StopRecordingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StopRecordingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelnyxVoiceSettingsParam struct {
	// The voice speed to be used for the voice. The voice speed must be between 0.1
	// and 2.0. Default value is 1.0.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
	paramObj
}

func (r TelnyxVoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TelnyxVoiceSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelnyxVoiceSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings associated with speech to text for the voice assistant. This is
// only relevant if the assistant uses a text-to-text language model. Any assistant
// using a model with native audio support (e.g. `fixie-ai/ultravox-v0_4`) will
// ignore this field.
type TranscriptionConfigParam struct {
	// The speech to text model to be used by the voice assistant.
	//
	//   - `distil-whisper/distil-large-v2` is lower latency but English-only.
	//   - `openai/whisper-large-v3-turbo` is multi-lingual with automatic language
	//     detection but slightly higher latency.
	//   - `google` is a multi-lingual option, please describe the language in the
	//     `language` field.
	Model param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r TranscriptionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionEngineAConfigParam struct {
	// Enables speaker diarization.
	EnableSpeakerDiarization param.Opt[bool] `json:"enable_speaker_diarization,omitzero"`
	// Whether to send also interim results. If set to false, only final results will
	// be sent.
	InterimResults param.Opt[bool] `json:"interim_results,omitzero"`
	// Defines maximum number of speakers in the conversation.
	MaxSpeakerCount param.Opt[int64] `json:"max_speaker_count,omitzero"`
	// Defines minimum number of speakers in the conversation.
	MinSpeakerCount param.Opt[int64] `json:"min_speaker_count,omitzero"`
	// Enables profanity_filter.
	ProfanityFilter param.Opt[bool] `json:"profanity_filter,omitzero"`
	// Enables enhanced transcription, this works for models `phone_call` and `video`.
	UseEnhanced param.Opt[bool] `json:"use_enhanced,omitzero"`
	// Hints to improve transcription accuracy.
	Hints []string `json:"hints,omitzero"`
	// Language to use for speech recognition
	//
	// Any of "af", "sq", "am", "ar", "hy", "az", "eu", "bn", "bs", "bg", "my", "ca",
	// "yue", "zh", "hr", "cs", "da", "nl", "en", "et", "fil", "fi", "fr", "gl", "ka",
	// "de", "el", "gu", "iw", "hi", "hu", "is", "id", "it", "ja", "jv", "kn", "kk",
	// "km", "ko", "lo", "lv", "lt", "mk", "ms", "ml", "mr", "mn", "ne", "no", "fa",
	// "pl", "pt", "pa", "ro", "ru", "rw", "sr", "si", "sk", "sl", "ss", "st", "es",
	// "su", "sw", "sv", "ta", "te", "th", "tn", "tr", "ts", "uk", "ur", "uz", "ve",
	// "vi", "xh", "zu".
	Language GoogleTranscriptionLanguage `json:"language,omitzero"`
	// The model to use for transcription.
	//
	// Any of "latest_long", "latest_short", "command_and_search", "phone_call",
	// "video", "default", "medical_conversation", "medical_dictation".
	Model TranscriptionEngineAConfigModel `json:"model,omitzero"`
	// Speech context to improve transcription accuracy.
	SpeechContext []TranscriptionEngineAConfigSpeechContextParam `json:"speech_context,omitzero"`
	// Engine identifier for Google transcription service
	//
	// Any of "A".
	TranscriptionEngine TranscriptionEngineAConfigTranscriptionEngine `json:"transcription_engine,omitzero"`
	paramObj
}

func (r TranscriptionEngineAConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionEngineAConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionEngineAConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The model to use for transcription.
type TranscriptionEngineAConfigModel string

const (
	TranscriptionEngineAConfigModelLatestLong          TranscriptionEngineAConfigModel = "latest_long"
	TranscriptionEngineAConfigModelLatestShort         TranscriptionEngineAConfigModel = "latest_short"
	TranscriptionEngineAConfigModelCommandAndSearch    TranscriptionEngineAConfigModel = "command_and_search"
	TranscriptionEngineAConfigModelPhoneCall           TranscriptionEngineAConfigModel = "phone_call"
	TranscriptionEngineAConfigModelVideo               TranscriptionEngineAConfigModel = "video"
	TranscriptionEngineAConfigModelDefault             TranscriptionEngineAConfigModel = "default"
	TranscriptionEngineAConfigModelMedicalConversation TranscriptionEngineAConfigModel = "medical_conversation"
	TranscriptionEngineAConfigModelMedicalDictation    TranscriptionEngineAConfigModel = "medical_dictation"
)

type TranscriptionEngineAConfigSpeechContextParam struct {
	// Boost factor for the speech context.
	Boost   param.Opt[float64] `json:"boost,omitzero"`
	Phrases []string           `json:"phrases,omitzero"`
	paramObj
}

func (r TranscriptionEngineAConfigSpeechContextParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionEngineAConfigSpeechContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionEngineAConfigSpeechContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engine identifier for Google transcription service
type TranscriptionEngineAConfigTranscriptionEngine string

const (
	TranscriptionEngineAConfigTranscriptionEngineA TranscriptionEngineAConfigTranscriptionEngine = "A"
)

type TranscriptionEngineBConfigParam struct {
	// Language to use for speech recognition
	//
	// Any of "en", "zh", "de", "es", "ru", "ko", "fr", "ja", "pt", "tr", "pl", "ca",
	// "nl", "ar", "sv", "it", "id", "hi", "fi", "vi", "he", "uk", "el", "ms", "cs",
	// "ro", "da", "hu", "ta", "no", "th", "ur", "hr", "bg", "lt", "la", "mi", "ml",
	// "cy", "sk", "te", "fa", "lv", "bn", "sr", "az", "sl", "kn", "et", "mk", "br",
	// "eu", "is", "hy", "ne", "mn", "bs", "kk", "sq", "sw", "gl", "mr", "pa", "si",
	// "km", "sn", "yo", "so", "af", "oc", "ka", "be", "tg", "sd", "gu", "am", "yi",
	// "lo", "uz", "fo", "ht", "ps", "tk", "nn", "mt", "sa", "lb", "my", "bo", "tl",
	// "mg", "as", "tt", "haw", "ln", "ha", "ba", "jw", "su", "auto_detect".
	Language TranscriptionEngineBConfigLanguage `json:"language,omitzero"`
	// Engine identifier for Telnyx transcription service
	//
	// Any of "B".
	TranscriptionEngine TranscriptionEngineBConfigTranscriptionEngine `json:"transcription_engine,omitzero"`
	// The model to use for transcription.
	//
	// Any of "openai/whisper-tiny", "openai/whisper-large-v3-turbo".
	TranscriptionModel TranscriptionEngineBConfigTranscriptionModel `json:"transcription_model,omitzero"`
	paramObj
}

func (r TranscriptionEngineBConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionEngineBConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionEngineBConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Language to use for speech recognition
type TranscriptionEngineBConfigLanguage string

const (
	TranscriptionEngineBConfigLanguageEn         TranscriptionEngineBConfigLanguage = "en"
	TranscriptionEngineBConfigLanguageZh         TranscriptionEngineBConfigLanguage = "zh"
	TranscriptionEngineBConfigLanguageDe         TranscriptionEngineBConfigLanguage = "de"
	TranscriptionEngineBConfigLanguageEs         TranscriptionEngineBConfigLanguage = "es"
	TranscriptionEngineBConfigLanguageRu         TranscriptionEngineBConfigLanguage = "ru"
	TranscriptionEngineBConfigLanguageKo         TranscriptionEngineBConfigLanguage = "ko"
	TranscriptionEngineBConfigLanguageFr         TranscriptionEngineBConfigLanguage = "fr"
	TranscriptionEngineBConfigLanguageJa         TranscriptionEngineBConfigLanguage = "ja"
	TranscriptionEngineBConfigLanguagePt         TranscriptionEngineBConfigLanguage = "pt"
	TranscriptionEngineBConfigLanguageTr         TranscriptionEngineBConfigLanguage = "tr"
	TranscriptionEngineBConfigLanguagePl         TranscriptionEngineBConfigLanguage = "pl"
	TranscriptionEngineBConfigLanguageCa         TranscriptionEngineBConfigLanguage = "ca"
	TranscriptionEngineBConfigLanguageNl         TranscriptionEngineBConfigLanguage = "nl"
	TranscriptionEngineBConfigLanguageAr         TranscriptionEngineBConfigLanguage = "ar"
	TranscriptionEngineBConfigLanguageSv         TranscriptionEngineBConfigLanguage = "sv"
	TranscriptionEngineBConfigLanguageIt         TranscriptionEngineBConfigLanguage = "it"
	TranscriptionEngineBConfigLanguageID         TranscriptionEngineBConfigLanguage = "id"
	TranscriptionEngineBConfigLanguageHi         TranscriptionEngineBConfigLanguage = "hi"
	TranscriptionEngineBConfigLanguageFi         TranscriptionEngineBConfigLanguage = "fi"
	TranscriptionEngineBConfigLanguageVi         TranscriptionEngineBConfigLanguage = "vi"
	TranscriptionEngineBConfigLanguageHe         TranscriptionEngineBConfigLanguage = "he"
	TranscriptionEngineBConfigLanguageUk         TranscriptionEngineBConfigLanguage = "uk"
	TranscriptionEngineBConfigLanguageEl         TranscriptionEngineBConfigLanguage = "el"
	TranscriptionEngineBConfigLanguageMs         TranscriptionEngineBConfigLanguage = "ms"
	TranscriptionEngineBConfigLanguageCs         TranscriptionEngineBConfigLanguage = "cs"
	TranscriptionEngineBConfigLanguageRo         TranscriptionEngineBConfigLanguage = "ro"
	TranscriptionEngineBConfigLanguageDa         TranscriptionEngineBConfigLanguage = "da"
	TranscriptionEngineBConfigLanguageHu         TranscriptionEngineBConfigLanguage = "hu"
	TranscriptionEngineBConfigLanguageTa         TranscriptionEngineBConfigLanguage = "ta"
	TranscriptionEngineBConfigLanguageNo         TranscriptionEngineBConfigLanguage = "no"
	TranscriptionEngineBConfigLanguageTh         TranscriptionEngineBConfigLanguage = "th"
	TranscriptionEngineBConfigLanguageUr         TranscriptionEngineBConfigLanguage = "ur"
	TranscriptionEngineBConfigLanguageHr         TranscriptionEngineBConfigLanguage = "hr"
	TranscriptionEngineBConfigLanguageBg         TranscriptionEngineBConfigLanguage = "bg"
	TranscriptionEngineBConfigLanguageLt         TranscriptionEngineBConfigLanguage = "lt"
	TranscriptionEngineBConfigLanguageLa         TranscriptionEngineBConfigLanguage = "la"
	TranscriptionEngineBConfigLanguageMi         TranscriptionEngineBConfigLanguage = "mi"
	TranscriptionEngineBConfigLanguageMl         TranscriptionEngineBConfigLanguage = "ml"
	TranscriptionEngineBConfigLanguageCy         TranscriptionEngineBConfigLanguage = "cy"
	TranscriptionEngineBConfigLanguageSk         TranscriptionEngineBConfigLanguage = "sk"
	TranscriptionEngineBConfigLanguageTe         TranscriptionEngineBConfigLanguage = "te"
	TranscriptionEngineBConfigLanguageFa         TranscriptionEngineBConfigLanguage = "fa"
	TranscriptionEngineBConfigLanguageLv         TranscriptionEngineBConfigLanguage = "lv"
	TranscriptionEngineBConfigLanguageBn         TranscriptionEngineBConfigLanguage = "bn"
	TranscriptionEngineBConfigLanguageSr         TranscriptionEngineBConfigLanguage = "sr"
	TranscriptionEngineBConfigLanguageAz         TranscriptionEngineBConfigLanguage = "az"
	TranscriptionEngineBConfigLanguageSl         TranscriptionEngineBConfigLanguage = "sl"
	TranscriptionEngineBConfigLanguageKn         TranscriptionEngineBConfigLanguage = "kn"
	TranscriptionEngineBConfigLanguageEt         TranscriptionEngineBConfigLanguage = "et"
	TranscriptionEngineBConfigLanguageMk         TranscriptionEngineBConfigLanguage = "mk"
	TranscriptionEngineBConfigLanguageBr         TranscriptionEngineBConfigLanguage = "br"
	TranscriptionEngineBConfigLanguageEu         TranscriptionEngineBConfigLanguage = "eu"
	TranscriptionEngineBConfigLanguageIs         TranscriptionEngineBConfigLanguage = "is"
	TranscriptionEngineBConfigLanguageHy         TranscriptionEngineBConfigLanguage = "hy"
	TranscriptionEngineBConfigLanguageNe         TranscriptionEngineBConfigLanguage = "ne"
	TranscriptionEngineBConfigLanguageMn         TranscriptionEngineBConfigLanguage = "mn"
	TranscriptionEngineBConfigLanguageBs         TranscriptionEngineBConfigLanguage = "bs"
	TranscriptionEngineBConfigLanguageKk         TranscriptionEngineBConfigLanguage = "kk"
	TranscriptionEngineBConfigLanguageSq         TranscriptionEngineBConfigLanguage = "sq"
	TranscriptionEngineBConfigLanguageSw         TranscriptionEngineBConfigLanguage = "sw"
	TranscriptionEngineBConfigLanguageGl         TranscriptionEngineBConfigLanguage = "gl"
	TranscriptionEngineBConfigLanguageMr         TranscriptionEngineBConfigLanguage = "mr"
	TranscriptionEngineBConfigLanguagePa         TranscriptionEngineBConfigLanguage = "pa"
	TranscriptionEngineBConfigLanguageSi         TranscriptionEngineBConfigLanguage = "si"
	TranscriptionEngineBConfigLanguageKm         TranscriptionEngineBConfigLanguage = "km"
	TranscriptionEngineBConfigLanguageSn         TranscriptionEngineBConfigLanguage = "sn"
	TranscriptionEngineBConfigLanguageYo         TranscriptionEngineBConfigLanguage = "yo"
	TranscriptionEngineBConfigLanguageSo         TranscriptionEngineBConfigLanguage = "so"
	TranscriptionEngineBConfigLanguageAf         TranscriptionEngineBConfigLanguage = "af"
	TranscriptionEngineBConfigLanguageOc         TranscriptionEngineBConfigLanguage = "oc"
	TranscriptionEngineBConfigLanguageKa         TranscriptionEngineBConfigLanguage = "ka"
	TranscriptionEngineBConfigLanguageBe         TranscriptionEngineBConfigLanguage = "be"
	TranscriptionEngineBConfigLanguageTg         TranscriptionEngineBConfigLanguage = "tg"
	TranscriptionEngineBConfigLanguageSd         TranscriptionEngineBConfigLanguage = "sd"
	TranscriptionEngineBConfigLanguageGu         TranscriptionEngineBConfigLanguage = "gu"
	TranscriptionEngineBConfigLanguageAm         TranscriptionEngineBConfigLanguage = "am"
	TranscriptionEngineBConfigLanguageYi         TranscriptionEngineBConfigLanguage = "yi"
	TranscriptionEngineBConfigLanguageLo         TranscriptionEngineBConfigLanguage = "lo"
	TranscriptionEngineBConfigLanguageUz         TranscriptionEngineBConfigLanguage = "uz"
	TranscriptionEngineBConfigLanguageFo         TranscriptionEngineBConfigLanguage = "fo"
	TranscriptionEngineBConfigLanguageHt         TranscriptionEngineBConfigLanguage = "ht"
	TranscriptionEngineBConfigLanguagePs         TranscriptionEngineBConfigLanguage = "ps"
	TranscriptionEngineBConfigLanguageTk         TranscriptionEngineBConfigLanguage = "tk"
	TranscriptionEngineBConfigLanguageNn         TranscriptionEngineBConfigLanguage = "nn"
	TranscriptionEngineBConfigLanguageMt         TranscriptionEngineBConfigLanguage = "mt"
	TranscriptionEngineBConfigLanguageSa         TranscriptionEngineBConfigLanguage = "sa"
	TranscriptionEngineBConfigLanguageLb         TranscriptionEngineBConfigLanguage = "lb"
	TranscriptionEngineBConfigLanguageMy         TranscriptionEngineBConfigLanguage = "my"
	TranscriptionEngineBConfigLanguageBo         TranscriptionEngineBConfigLanguage = "bo"
	TranscriptionEngineBConfigLanguageTl         TranscriptionEngineBConfigLanguage = "tl"
	TranscriptionEngineBConfigLanguageMg         TranscriptionEngineBConfigLanguage = "mg"
	TranscriptionEngineBConfigLanguageAs         TranscriptionEngineBConfigLanguage = "as"
	TranscriptionEngineBConfigLanguageTt         TranscriptionEngineBConfigLanguage = "tt"
	TranscriptionEngineBConfigLanguageHaw        TranscriptionEngineBConfigLanguage = "haw"
	TranscriptionEngineBConfigLanguageLn         TranscriptionEngineBConfigLanguage = "ln"
	TranscriptionEngineBConfigLanguageHa         TranscriptionEngineBConfigLanguage = "ha"
	TranscriptionEngineBConfigLanguageBa         TranscriptionEngineBConfigLanguage = "ba"
	TranscriptionEngineBConfigLanguageJw         TranscriptionEngineBConfigLanguage = "jw"
	TranscriptionEngineBConfigLanguageSu         TranscriptionEngineBConfigLanguage = "su"
	TranscriptionEngineBConfigLanguageAutoDetect TranscriptionEngineBConfigLanguage = "auto_detect"
)

// Engine identifier for Telnyx transcription service
type TranscriptionEngineBConfigTranscriptionEngine string

const (
	TranscriptionEngineBConfigTranscriptionEngineB TranscriptionEngineBConfigTranscriptionEngine = "B"
)

// The model to use for transcription.
type TranscriptionEngineBConfigTranscriptionModel string

const (
	TranscriptionEngineBConfigTranscriptionModelOpenAIWhisperTiny         TranscriptionEngineBConfigTranscriptionModel = "openai/whisper-tiny"
	TranscriptionEngineBConfigTranscriptionModelOpenAIWhisperLargeV3Turbo TranscriptionEngineBConfigTranscriptionModel = "openai/whisper-large-v3-turbo"
)

type TranscriptionStartRequestParam struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Indicates which leg of the call will be transcribed. Use `inbound` for the leg
	// that requested the transcription, `outbound` for the other leg, and `both` for
	// both legs of the call. Will default to `inbound`.
	TranscriptionTracks param.Opt[string] `json:"transcription_tracks,omitzero"`
	// Engine to use for speech recognition. `A` - `Google`, `B` - `Telnyx`.
	//
	// Any of "A", "B".
	TranscriptionEngine       TranscriptionStartRequestTranscriptionEngine                 `json:"transcription_engine,omitzero"`
	TranscriptionEngineConfig TranscriptionStartRequestTranscriptionEngineConfigUnionParam `json:"transcription_engine_config,omitzero"`
	paramObj
}

func (r TranscriptionStartRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionStartRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionStartRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Engine to use for speech recognition. `A` - `Google`, `B` - `Telnyx`.
type TranscriptionStartRequestTranscriptionEngine string

const (
	TranscriptionStartRequestTranscriptionEngineA TranscriptionStartRequestTranscriptionEngine = "A"
	TranscriptionStartRequestTranscriptionEngineB TranscriptionStartRequestTranscriptionEngine = "B"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TranscriptionStartRequestTranscriptionEngineConfigUnionParam struct {
	OfA *TranscriptionEngineAConfigParam `json:",omitzero,inline"`
	OfB *TranscriptionEngineBConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfA, u.OfB)
}
func (u *TranscriptionStartRequestTranscriptionEngineConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TranscriptionStartRequestTranscriptionEngineConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfA) {
		return u.OfA
	} else if !param.IsOmitted(u.OfB) {
		return u.OfB
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetEnableSpeakerDiarization() *bool {
	if vt := u.OfA; vt != nil && vt.EnableSpeakerDiarization.Valid() {
		return &vt.EnableSpeakerDiarization.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetHints() []string {
	if vt := u.OfA; vt != nil {
		return vt.Hints
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetInterimResults() *bool {
	if vt := u.OfA; vt != nil && vt.InterimResults.Valid() {
		return &vt.InterimResults.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetMaxSpeakerCount() *int64 {
	if vt := u.OfA; vt != nil && vt.MaxSpeakerCount.Valid() {
		return &vt.MaxSpeakerCount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetMinSpeakerCount() *int64 {
	if vt := u.OfA; vt != nil && vt.MinSpeakerCount.Valid() {
		return &vt.MinSpeakerCount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetModel() *string {
	if vt := u.OfA; vt != nil {
		return (*string)(&vt.Model)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetProfanityFilter() *bool {
	if vt := u.OfA; vt != nil && vt.ProfanityFilter.Valid() {
		return &vt.ProfanityFilter.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetSpeechContext() []TranscriptionEngineAConfigSpeechContextParam {
	if vt := u.OfA; vt != nil {
		return vt.SpeechContext
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetUseEnhanced() *bool {
	if vt := u.OfA; vt != nil && vt.UseEnhanced.Valid() {
		return &vt.UseEnhanced.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetTranscriptionModel() *string {
	if vt := u.OfB; vt != nil {
		return (*string)(&vt.TranscriptionModel)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetLanguage() *string {
	if vt := u.OfA; vt != nil {
		return (*string)(&vt.Language)
	} else if vt := u.OfB; vt != nil {
		return (*string)(&vt.Language)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscriptionStartRequestTranscriptionEngineConfigUnionParam) GetTranscriptionEngine() *string {
	if vt := u.OfA; vt != nil {
		return (*string)(&vt.TranscriptionEngine)
	} else if vt := u.OfB; vt != nil {
		return (*string)(&vt.TranscriptionEngine)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[TranscriptionStartRequestTranscriptionEngineConfigUnionParam](
		"transcription_engine",
		apijson.Discriminator[TranscriptionEngineAConfigParam]("A"),
		apijson.Discriminator[TranscriptionEngineBConfigParam]("B"),
	)
}

type CallActionAnswerResponse struct {
	Data CallActionAnswerResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionAnswerResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionAnswerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionAnswerResponseData struct {
	// The ID of the recording. Only present when the record parameter is set to
	// record-from-answer.
	RecordingID string `json:"recording_id" format:"uuid"`
	Result      string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordingID respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionAnswerResponseData) RawJSON() string { return r.JSON.raw }
func (r *CallActionAnswerResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionBridgeResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionBridgeResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionBridgeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionEnqueueResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionEnqueueResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionEnqueueResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionGatherResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionGatherResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingAIResponse struct {
	Data CallActionGatherUsingAIResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionGatherUsingAIResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionGatherUsingAIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingAIResponseData struct {
	// The ID of the conversation created by the command.
	ConversationID string `json:"conversation_id" format:"uuid"`
	Result         string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID respjson.Field
		Result         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionGatherUsingAIResponseData) RawJSON() string { return r.JSON.raw }
func (r *CallActionGatherUsingAIResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingAudioResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionGatherUsingAudioResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionGatherUsingAudioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingSpeakResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionGatherUsingSpeakResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionGatherUsingSpeakResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionHangupResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionHangupResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionHangupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionLeaveQueueResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionLeaveQueueResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionLeaveQueueResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionPauseRecordingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionPauseRecordingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionPauseRecordingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionReferResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionReferResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionReferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionRejectResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionRejectResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionRejectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionResumeRecordingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionResumeRecordingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionResumeRecordingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSendDtmfResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionSendDtmfResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionSendDtmfResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSendSipInfoResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionSendSipInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionSendSipInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSpeakResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionSpeakResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionSpeakResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartAIAssistantResponse struct {
	Data CallActionStartAIAssistantResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartAIAssistantResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartAIAssistantResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartAIAssistantResponseData struct {
	// The ID of the conversation created by the command.
	ConversationID string `json:"conversation_id" format:"uuid"`
	Result         string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID respjson.Field
		Result         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartAIAssistantResponseData) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartAIAssistantResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartForkingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartForkingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartForkingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartNoiseSuppressionResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartNoiseSuppressionResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartNoiseSuppressionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartPlaybackResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartPlaybackResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartPlaybackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartRecordingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartRecordingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartRecordingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartSiprecResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartSiprecResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartSiprecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartStreamingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartStreamingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartStreamingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStartTranscriptionResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStartTranscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStartTranscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopAIAssistantResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopAIAssistantResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopAIAssistantResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopForkingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopForkingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopForkingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopGatherResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopGatherResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopGatherResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopNoiseSuppressionResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopNoiseSuppressionResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopNoiseSuppressionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopPlaybackResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopPlaybackResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopPlaybackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopRecordingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopRecordingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopRecordingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopSiprecResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopSiprecResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopSiprecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopStreamingResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopStreamingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopStreamingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopTranscriptionResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionStopTranscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionStopTranscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSwitchSupervisorRoleResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionSwitchSupervisorRoleResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionSwitchSupervisorRoleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionTransferResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionUpdateClientStateResponse struct {
	Data CallControlCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallActionUpdateClientStateResponse) RawJSON() string { return r.JSON.raw }
func (r *CallActionUpdateClientStateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionAnswerParams struct {
	// Use this field to set the Billing Group ID for the call. Must be a valid and
	// existing Billing Group ID.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero" format:"uuid"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The custom recording file name to be used instead of the default `call_leg_id`.
	// Telnyx will still add a Unix timestamp suffix.
	RecordCustomFileName param.Opt[string] `json:"record_custom_file_name,omitzero"`
	// Defines the maximum length for the recording in seconds when `record` is
	// specified. The minimum value is 0. The maximum value is 43200. The default value
	// is 0 (infinite).
	RecordMaxLength param.Opt[int64] `json:"record_max_length,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected when `record` is specified. The timer only starts when the
	// speech is detected. Please note that call transcription is used to detect
	// silence and the related charge will be applied. The minimum value is 0. The
	// default value is 0 (infinite).
	RecordTimeoutSecs param.Opt[int64] `json:"record_timeout_secs,omitzero"`
	// Generate silence RTP packets when no transmission available.
	SendSilenceWhenIdle param.Opt[bool] `json:"send_silence_when_idle,omitzero"`
	// The destination WebSocket address where the stream is going to be delivered.
	StreamURL param.Opt[string] `json:"stream_url,omitzero"`
	// Enable transcription upon call answer. The default value is false.
	Transcription param.Opt[bool] `json:"transcription,omitzero"`
	// Use this field to override the URL for which Telnyx will send subsequent
	// webhooks to for this call.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Custom headers to be added to the SIP INVITE response.
	CustomHeaders []CustomSipHeaderParam `json:"custom_headers,omitzero"`
	// The list of comma-separated codecs in a preferred order for the forked media to
	// be received.
	//
	// Any of "G722,PCMU,PCMA,G729,OPUS,VP8,H264".
	PreferredCodecs CallActionAnswerParamsPreferredCodecs `json:"preferred_codecs,omitzero"`
	// Start recording automatically after an event. Disabled by default.
	//
	// Any of "record-from-answer".
	Record CallActionAnswerParamsRecord `json:"record,omitzero"`
	// Defines which channel should be recorded ('single' or 'dual') when `record` is
	// specified.
	//
	// Any of "single", "dual".
	RecordChannels CallActionAnswerParamsRecordChannels `json:"record_channels,omitzero"`
	// Defines the format of the recording ('wav' or 'mp3') when `record` is specified.
	//
	// Any of "wav", "mp3".
	RecordFormat CallActionAnswerParamsRecordFormat `json:"record_format,omitzero"`
	// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
	// If only single track is specified (`inbound`, `outbound`), `channels`
	// configuration is ignored and it will be recorded as mono (single channel).
	//
	// Any of "both", "inbound", "outbound".
	RecordTrack CallActionAnswerParamsRecordTrack `json:"record_track,omitzero"`
	// When set to `trim-silence`, silence will be removed from the beginning and end
	// of the recording.
	//
	// Any of "trim-silence".
	RecordTrim CallActionAnswerParamsRecordTrim `json:"record_trim,omitzero"`
	// SIP headers to be added to the SIP INVITE response. Currently only User-to-User
	// header is supported.
	SipHeaders []SipHeaderParam `json:"sip_headers,omitzero"`
	// Use this field to modify sound effects, for example adjust the pitch.
	SoundModifications SoundModificationsParam `json:"sound_modifications,omitzero"`
	// Indicates codec for bidirectional streaming RTP payloads. Used only with
	// stream_bidirectional_mode=rtp. Case sensitive.
	//
	// Any of "PCMU", "PCMA", "G722", "OPUS", "AMR-WB".
	StreamBidirectionalCodec StreamBidirectionalCodec `json:"stream_bidirectional_codec,omitzero"`
	// Configures method of bidirectional streaming (mp3, rtp).
	//
	// Any of "mp3", "rtp".
	StreamBidirectionalMode StreamBidirectionalMode `json:"stream_bidirectional_mode,omitzero"`
	// Specifies which call legs should receive the bidirectional stream audio.
	//
	// Any of "both", "self", "opposite".
	StreamBidirectionalTargetLegs StreamBidirectionalTargetLegs `json:"stream_bidirectional_target_legs,omitzero"`
	// Specifies the codec to be used for the streamed audio. When set to 'default' or
	// when transcoding is not possible, the codec from the call will be used.
	// Currently, transcoding is only supported between PCMU and PCMA codecs.
	//
	// Any of "PCMA", "PCMU", "default".
	StreamCodec StreamCodec `json:"stream_codec,omitzero"`
	// Specifies which track should be streamed.
	//
	// Any of "inbound_track", "outbound_track", "both_tracks".
	StreamTrack         CallActionAnswerParamsStreamTrack `json:"stream_track,omitzero"`
	TranscriptionConfig TranscriptionStartRequestParam    `json:"transcription_config,omitzero"`
	// HTTP request type used for `webhook_url`.
	//
	// Any of "POST", "GET".
	WebhookURLMethod CallActionAnswerParamsWebhookURLMethod `json:"webhook_url_method,omitzero"`
	paramObj
}

func (r CallActionAnswerParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionAnswerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionAnswerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The list of comma-separated codecs in a preferred order for the forked media to
// be received.
type CallActionAnswerParamsPreferredCodecs string

const (
	CallActionAnswerParamsPreferredCodecsG722PcmuPcmaG729OpusVp8H264 CallActionAnswerParamsPreferredCodecs = "G722,PCMU,PCMA,G729,OPUS,VP8,H264"
)

// Start recording automatically after an event. Disabled by default.
type CallActionAnswerParamsRecord string

const (
	CallActionAnswerParamsRecordRecordFromAnswer CallActionAnswerParamsRecord = "record-from-answer"
)

// Defines which channel should be recorded ('single' or 'dual') when `record` is
// specified.
type CallActionAnswerParamsRecordChannels string

const (
	CallActionAnswerParamsRecordChannelsSingle CallActionAnswerParamsRecordChannels = "single"
	CallActionAnswerParamsRecordChannelsDual   CallActionAnswerParamsRecordChannels = "dual"
)

// Defines the format of the recording ('wav' or 'mp3') when `record` is specified.
type CallActionAnswerParamsRecordFormat string

const (
	CallActionAnswerParamsRecordFormatWav CallActionAnswerParamsRecordFormat = "wav"
	CallActionAnswerParamsRecordFormatMP3 CallActionAnswerParamsRecordFormat = "mp3"
)

// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
// If only single track is specified (`inbound`, `outbound`), `channels`
// configuration is ignored and it will be recorded as mono (single channel).
type CallActionAnswerParamsRecordTrack string

const (
	CallActionAnswerParamsRecordTrackBoth     CallActionAnswerParamsRecordTrack = "both"
	CallActionAnswerParamsRecordTrackInbound  CallActionAnswerParamsRecordTrack = "inbound"
	CallActionAnswerParamsRecordTrackOutbound CallActionAnswerParamsRecordTrack = "outbound"
)

// When set to `trim-silence`, silence will be removed from the beginning and end
// of the recording.
type CallActionAnswerParamsRecordTrim string

const (
	CallActionAnswerParamsRecordTrimTrimSilence CallActionAnswerParamsRecordTrim = "trim-silence"
)

// Specifies which track should be streamed.
type CallActionAnswerParamsStreamTrack string

const (
	CallActionAnswerParamsStreamTrackInboundTrack  CallActionAnswerParamsStreamTrack = "inbound_track"
	CallActionAnswerParamsStreamTrackOutboundTrack CallActionAnswerParamsStreamTrack = "outbound_track"
	CallActionAnswerParamsStreamTrackBothTracks    CallActionAnswerParamsStreamTrack = "both_tracks"
)

// HTTP request type used for `webhook_url`.
type CallActionAnswerParamsWebhookURLMethod string

const (
	CallActionAnswerParamsWebhookURLMethodPost CallActionAnswerParamsWebhookURLMethod = "POST"
	CallActionAnswerParamsWebhookURLMethodGet  CallActionAnswerParamsWebhookURLMethod = "GET"
)

type CallActionBridgeParams struct {
	// The Call Control ID of the call you want to bridge with, can't be used together
	// with queue parameter or video_room_id parameter.
	CallControlID string `json:"call_control_id,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up
	// or is transferred). If supplied with the value `self`, the current leg will be
	// parked after unbridge. If not set, the default behavior is to hang up the leg.
	ParkAfterUnbridge param.Opt[string] `json:"park_after_unbridge,omitzero"`
	// Specifies whether to play a ringtone if the call you want to bridge with has not
	// yet been answered.
	PlayRingtone param.Opt[bool] `json:"play_ringtone,omitzero"`
	// The name of the queue you want to bridge with, can't be used together with
	// call_control_id parameter or video_room_id parameter. Bridging with a queue
	// means bridging with the first call in the queue. The call will always be removed
	// from the queue regardless of whether bridging succeeds. Returns an error when
	// the queue is empty.
	Queue param.Opt[string] `json:"queue,omitzero"`
	// The custom recording file name to be used instead of the default `call_leg_id`.
	// Telnyx will still add a Unix timestamp suffix.
	RecordCustomFileName param.Opt[string] `json:"record_custom_file_name,omitzero"`
	// Defines the maximum length for the recording in seconds when `record` is
	// specified. The minimum value is 0. The maximum value is 43200. The default value
	// is 0 (infinite).
	RecordMaxLength param.Opt[int64] `json:"record_max_length,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected when `record` is specified. The timer only starts when the
	// speech is detected. Please note that call transcription is used to detect
	// silence and the related charge will be applied. The minimum value is 0. The
	// default value is 0 (infinite).
	RecordTimeoutSecs param.Opt[int64] `json:"record_timeout_secs,omitzero"`
	// The additional parameter that will be passed to the video conference. It is a
	// text field and the user can decide how to use it. For example, you can set the
	// participant name or pass JSON text. It can be used only with video_room_id
	// parameter.
	VideoRoomContext param.Opt[string] `json:"video_room_context,omitzero"`
	// The ID of the video room you want to bridge with, can't be used together with
	// call_control_id parameter or queue parameter.
	VideoRoomID param.Opt[string] `json:"video_room_id,omitzero" format:"uuid"`
	// When enabled, DTMF tones are not passed to the call participant. The webhooks
	// containing the DTMF information will be sent.
	//
	// Any of "none", "both", "self", "opposite".
	MuteDtmf CallActionBridgeParamsMuteDtmf `json:"mute_dtmf,omitzero"`
	// Start recording automatically after an event. Disabled by default.
	//
	// Any of "record-from-answer".
	Record CallActionBridgeParamsRecord `json:"record,omitzero"`
	// Defines which channel should be recorded ('single' or 'dual') when `record` is
	// specified.
	//
	// Any of "single", "dual".
	RecordChannels CallActionBridgeParamsRecordChannels `json:"record_channels,omitzero"`
	// Defines the format of the recording ('wav' or 'mp3') when `record` is specified.
	//
	// Any of "wav", "mp3".
	RecordFormat CallActionBridgeParamsRecordFormat `json:"record_format,omitzero"`
	// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
	// If only single track is specified (`inbound`, `outbound`), `channels`
	// configuration is ignored and it will be recorded as mono (single channel).
	//
	// Any of "both", "inbound", "outbound".
	RecordTrack CallActionBridgeParamsRecordTrack `json:"record_track,omitzero"`
	// When set to `trim-silence`, silence will be removed from the beginning and end
	// of the recording.
	//
	// Any of "trim-silence".
	RecordTrim CallActionBridgeParamsRecordTrim `json:"record_trim,omitzero"`
	// Specifies which country ringtone to play when `play_ringtone` is set to `true`.
	// If not set, the US ringtone will be played.
	//
	// Any of "at", "au", "be", "bg", "br", "ch", "cl", "cn", "cz", "de", "dk", "ee",
	// "es", "fi", "fr", "gr", "hu", "il", "in", "it", "jp", "lt", "mx", "my", "nl",
	// "no", "nz", "ph", "pl", "pt", "ru", "se", "sg", "th", "tw", "uk", "us-old",
	// "us", "ve", "za".
	Ringtone CallActionBridgeParamsRingtone `json:"ringtone,omitzero"`
	paramObj
}

func (r CallActionBridgeParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionBridgeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionBridgeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, DTMF tones are not passed to the call participant. The webhooks
// containing the DTMF information will be sent.
type CallActionBridgeParamsMuteDtmf string

const (
	CallActionBridgeParamsMuteDtmfNone     CallActionBridgeParamsMuteDtmf = "none"
	CallActionBridgeParamsMuteDtmfBoth     CallActionBridgeParamsMuteDtmf = "both"
	CallActionBridgeParamsMuteDtmfSelf     CallActionBridgeParamsMuteDtmf = "self"
	CallActionBridgeParamsMuteDtmfOpposite CallActionBridgeParamsMuteDtmf = "opposite"
)

// Start recording automatically after an event. Disabled by default.
type CallActionBridgeParamsRecord string

const (
	CallActionBridgeParamsRecordRecordFromAnswer CallActionBridgeParamsRecord = "record-from-answer"
)

// Defines which channel should be recorded ('single' or 'dual') when `record` is
// specified.
type CallActionBridgeParamsRecordChannels string

const (
	CallActionBridgeParamsRecordChannelsSingle CallActionBridgeParamsRecordChannels = "single"
	CallActionBridgeParamsRecordChannelsDual   CallActionBridgeParamsRecordChannels = "dual"
)

// Defines the format of the recording ('wav' or 'mp3') when `record` is specified.
type CallActionBridgeParamsRecordFormat string

const (
	CallActionBridgeParamsRecordFormatWav CallActionBridgeParamsRecordFormat = "wav"
	CallActionBridgeParamsRecordFormatMP3 CallActionBridgeParamsRecordFormat = "mp3"
)

// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
// If only single track is specified (`inbound`, `outbound`), `channels`
// configuration is ignored and it will be recorded as mono (single channel).
type CallActionBridgeParamsRecordTrack string

const (
	CallActionBridgeParamsRecordTrackBoth     CallActionBridgeParamsRecordTrack = "both"
	CallActionBridgeParamsRecordTrackInbound  CallActionBridgeParamsRecordTrack = "inbound"
	CallActionBridgeParamsRecordTrackOutbound CallActionBridgeParamsRecordTrack = "outbound"
)

// When set to `trim-silence`, silence will be removed from the beginning and end
// of the recording.
type CallActionBridgeParamsRecordTrim string

const (
	CallActionBridgeParamsRecordTrimTrimSilence CallActionBridgeParamsRecordTrim = "trim-silence"
)

// Specifies which country ringtone to play when `play_ringtone` is set to `true`.
// If not set, the US ringtone will be played.
type CallActionBridgeParamsRingtone string

const (
	CallActionBridgeParamsRingtoneAt    CallActionBridgeParamsRingtone = "at"
	CallActionBridgeParamsRingtoneAu    CallActionBridgeParamsRingtone = "au"
	CallActionBridgeParamsRingtoneBe    CallActionBridgeParamsRingtone = "be"
	CallActionBridgeParamsRingtoneBg    CallActionBridgeParamsRingtone = "bg"
	CallActionBridgeParamsRingtoneBr    CallActionBridgeParamsRingtone = "br"
	CallActionBridgeParamsRingtoneCh    CallActionBridgeParamsRingtone = "ch"
	CallActionBridgeParamsRingtoneCl    CallActionBridgeParamsRingtone = "cl"
	CallActionBridgeParamsRingtoneCn    CallActionBridgeParamsRingtone = "cn"
	CallActionBridgeParamsRingtoneCz    CallActionBridgeParamsRingtone = "cz"
	CallActionBridgeParamsRingtoneDe    CallActionBridgeParamsRingtone = "de"
	CallActionBridgeParamsRingtoneDk    CallActionBridgeParamsRingtone = "dk"
	CallActionBridgeParamsRingtoneEe    CallActionBridgeParamsRingtone = "ee"
	CallActionBridgeParamsRingtoneEs    CallActionBridgeParamsRingtone = "es"
	CallActionBridgeParamsRingtoneFi    CallActionBridgeParamsRingtone = "fi"
	CallActionBridgeParamsRingtoneFr    CallActionBridgeParamsRingtone = "fr"
	CallActionBridgeParamsRingtoneGr    CallActionBridgeParamsRingtone = "gr"
	CallActionBridgeParamsRingtoneHu    CallActionBridgeParamsRingtone = "hu"
	CallActionBridgeParamsRingtoneIl    CallActionBridgeParamsRingtone = "il"
	CallActionBridgeParamsRingtoneIn    CallActionBridgeParamsRingtone = "in"
	CallActionBridgeParamsRingtoneIt    CallActionBridgeParamsRingtone = "it"
	CallActionBridgeParamsRingtoneJp    CallActionBridgeParamsRingtone = "jp"
	CallActionBridgeParamsRingtoneLt    CallActionBridgeParamsRingtone = "lt"
	CallActionBridgeParamsRingtoneMx    CallActionBridgeParamsRingtone = "mx"
	CallActionBridgeParamsRingtoneMy    CallActionBridgeParamsRingtone = "my"
	CallActionBridgeParamsRingtoneNl    CallActionBridgeParamsRingtone = "nl"
	CallActionBridgeParamsRingtoneNo    CallActionBridgeParamsRingtone = "no"
	CallActionBridgeParamsRingtoneNz    CallActionBridgeParamsRingtone = "nz"
	CallActionBridgeParamsRingtonePh    CallActionBridgeParamsRingtone = "ph"
	CallActionBridgeParamsRingtonePl    CallActionBridgeParamsRingtone = "pl"
	CallActionBridgeParamsRingtonePt    CallActionBridgeParamsRingtone = "pt"
	CallActionBridgeParamsRingtoneRu    CallActionBridgeParamsRingtone = "ru"
	CallActionBridgeParamsRingtoneSe    CallActionBridgeParamsRingtone = "se"
	CallActionBridgeParamsRingtoneSg    CallActionBridgeParamsRingtone = "sg"
	CallActionBridgeParamsRingtoneTh    CallActionBridgeParamsRingtone = "th"
	CallActionBridgeParamsRingtoneTw    CallActionBridgeParamsRingtone = "tw"
	CallActionBridgeParamsRingtoneUk    CallActionBridgeParamsRingtone = "uk"
	CallActionBridgeParamsRingtoneUsOld CallActionBridgeParamsRingtone = "us-old"
	CallActionBridgeParamsRingtoneUs    CallActionBridgeParamsRingtone = "us"
	CallActionBridgeParamsRingtoneVe    CallActionBridgeParamsRingtone = "ve"
	CallActionBridgeParamsRingtoneZa    CallActionBridgeParamsRingtone = "za"
)

type CallActionEnqueueParams struct {
	// The name of the queue the call should be put in. If a queue with a given name
	// doesn't exist yet it will be created.
	QueueName string `json:"queue_name,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The maximum number of calls allowed in the queue at a given time. Can't be
	// modified for an existing queue.
	MaxSize param.Opt[int64] `json:"max_size,omitzero"`
	// The number of seconds after which the call will be removed from the queue.
	MaxWaitTimeSecs param.Opt[int64] `json:"max_wait_time_secs,omitzero"`
	paramObj
}

func (r CallActionEnqueueParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionEnqueueParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionEnqueueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// An id that will be sent back in the corresponding `call.gather.ended` webhook.
	// Will be randomly generated if not specified.
	GatherID param.Opt[string] `json:"gather_id,omitzero"`
	// The number of milliseconds to wait for the first DTMF.
	InitialTimeoutMillis param.Opt[int64] `json:"initial_timeout_millis,omitzero"`
	// The number of milliseconds to wait for input between digits.
	InterDigitTimeoutMillis param.Opt[int64] `json:"inter_digit_timeout_millis,omitzero"`
	// The maximum number of digits to fetch. This parameter has a maximum value
	// of 128.
	MaximumDigits param.Opt[int64] `json:"maximum_digits,omitzero"`
	// The minimum number of digits to fetch. This parameter has a minimum value of 1.
	MinimumDigits param.Opt[int64] `json:"minimum_digits,omitzero"`
	// The digit used to terminate input if fewer than `maximum_digits` digits have
	// been gathered.
	TerminatingDigit param.Opt[string] `json:"terminating_digit,omitzero"`
	// The number of milliseconds to wait to complete the request.
	TimeoutMillis param.Opt[int64] `json:"timeout_millis,omitzero"`
	// A list of all digits accepted as valid.
	ValidDigits param.Opt[string] `json:"valid_digits,omitzero"`
	paramObj
}

func (r CallActionGatherParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionGatherParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionGatherParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingAIParams struct {
	// The parameters described as a JSON Schema object that needs to be gathered by
	// the voice assistant. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	Parameters any `json:"parameters,omitzero,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Text that will be played when the gathering starts, if none then nothing will be
	// played when the gathering starts. The greeting can be text for any voice or SSML
	// for `AWS.Polly.<voice_id>` voices. There is a 3,000 character limit.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// Default is `false`. If set to `true`, the voice assistant will send updates to
	// the message history via the `call.ai_gather.message_history_updated`
	// [callback](https://developers.telnyx.com/api/call-control/call-gather-using-ai#callbacks)
	// in real time as the message history is updated.
	SendMessageHistoryUpdates param.Opt[bool] `json:"send_message_history_updates,omitzero"`
	// Default is `false`. If set to `true`, the voice assistant will send partial
	// results via the `call.ai_gather.partial_results`
	// [callback](https://developers.telnyx.com/api/call-control/call-gather-using-ai#callbacks)
	// in real time as individual fields are gathered. If set to `false`, the voice
	// assistant will only send the final result via the `call.ai_gather.ended`
	// callback.
	SendPartialResults param.Opt[bool] `json:"send_partial_results,omitzero"`
	// The number of milliseconds to wait for a user response before the voice
	// assistant times out and check if the user is still there.
	UserResponseTimeoutMs param.Opt[int64] `json:"user_response_timeout_ms,omitzero"`
	// The voice to be used by the voice assistant. Currently we support ElevenLabs,
	// Telnyx and AWS voices.
	//
	// **Supported Providers:**
	//
	//   - **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural
	//     voices, which provide more realistic, human-like speech, append `-Neural` to
	//     the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the
	//     [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html)
	//     for compatibility.
	//   - **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural,
	//     Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural,
	//     Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to
	//     [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)
	//   - **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g.,
	//     `ElevenLabs.BaseModel.John`). The `ModelId` part is optional. To use
	//     ElevenLabs, you must provide your ElevenLabs API key as an integration secret
	//     under `"voice_settings": {"api_key_ref": "<secret_id>"}`. See
	//     [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	//     for details. Check
	//     [available voices](https://elevenlabs.io/docs/api-reference/get-voices).
	//   - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`
	Voice param.Opt[string] `json:"voice,omitzero"`
	// Assistant configuration including choice of LLM, custom instructions, and tools.
	Assistant AssistantParam `json:"assistant,omitzero"`
	// Settings for handling user interruptions during assistant speech
	InterruptionSettings InterruptionSettingsParam `json:"interruption_settings,omitzero"`
	// Language to use for speech recognition
	//
	// Any of "af", "sq", "am", "ar", "hy", "az", "eu", "bn", "bs", "bg", "my", "ca",
	// "yue", "zh", "hr", "cs", "da", "nl", "en", "et", "fil", "fi", "fr", "gl", "ka",
	// "de", "el", "gu", "iw", "hi", "hu", "is", "id", "it", "ja", "jv", "kn", "kk",
	// "km", "ko", "lo", "lv", "lt", "mk", "ms", "ml", "mr", "mn", "ne", "no", "fa",
	// "pl", "pt", "pa", "ro", "ru", "rw", "sr", "si", "sk", "sl", "ss", "st", "es",
	// "su", "sw", "sv", "ta", "te", "th", "tn", "tr", "ts", "uk", "ur", "uz", "ve",
	// "vi", "xh", "zu".
	Language GoogleTranscriptionLanguage `json:"language,omitzero"`
	// The message history you want the voice assistant to be aware of, this can be
	// useful to keep the context of the conversation, or to pass additional
	// information to the voice assistant.
	MessageHistory []CallActionGatherUsingAIParamsMessageHistory `json:"message_history,omitzero"`
	// The settings associated with speech to text for the voice assistant. This is
	// only relevant if the assistant uses a text-to-text language model. Any assistant
	// using a model with native audio support (e.g. `fixie-ai/ultravox-v0_4`) will
	// ignore this field.
	Transcription TranscriptionConfigParam `json:"transcription,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings CallActionGatherUsingAIParamsVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r CallActionGatherUsingAIParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionGatherUsingAIParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionGatherUsingAIParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingAIParamsMessageHistory struct {
	// The content of the message
	Content param.Opt[string] `json:"content,omitzero"`
	// The role of the message sender
	//
	// Any of "assistant", "user".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r CallActionGatherUsingAIParamsMessageHistory) MarshalJSON() (data []byte, err error) {
	type shadow CallActionGatherUsingAIParamsMessageHistory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionGatherUsingAIParamsMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallActionGatherUsingAIParamsMessageHistory](
		"role", "assistant", "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallActionGatherUsingAIParamsVoiceSettingsUnion struct {
	OfElevenLabsVoiceSettings *ElevenLabsVoiceSettingsParam `json:",omitzero,inline"`
	OfTelnyxVoiceSettings     *TelnyxVoiceSettingsParam     `json:",omitzero,inline"`
	paramUnion
}

func (u CallActionGatherUsingAIParamsVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenLabsVoiceSettings, u.OfTelnyxVoiceSettings)
}
func (u *CallActionGatherUsingAIParamsVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallActionGatherUsingAIParamsVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenLabsVoiceSettings) {
		return u.OfElevenLabsVoiceSettings
	} else if !param.IsOmitted(u.OfTelnyxVoiceSettings) {
		return u.OfTelnyxVoiceSettings
	}
	return nil
}

type CallActionGatherUsingAudioParams struct {
	// The URL of a file to be played back at the beginning of each prompt. The URL can
	// point to either a WAV or MP3 file. media_name and audio_url cannot be used
	// together in one request.
	AudioURL param.Opt[string] `json:"audio_url,omitzero"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The number of milliseconds to wait for input between digits.
	InterDigitTimeoutMillis param.Opt[int64] `json:"inter_digit_timeout_millis,omitzero"`
	// The URL of a file to play when digits don't match the `valid_digits` parameter
	// or the number of digits is not between `min` and `max`. The URL can point to
	// either a WAV or MP3 file. invalid_media_name and invalid_audio_url cannot be
	// used together in one request.
	InvalidAudioURL param.Opt[string] `json:"invalid_audio_url,omitzero"`
	// The media_name of a file to be played back when digits don't match the
	// `valid_digits` parameter or the number of digits is not between `min` and `max`.
	// The media_name must point to a file previously uploaded to
	// api.telnyx.com/v2/media by the same user/organization. The file must either be a
	// WAV or MP3 file.
	InvalidMediaName param.Opt[string] `json:"invalid_media_name,omitzero"`
	// The maximum number of digits to fetch. This parameter has a maximum value
	// of 128.
	MaximumDigits param.Opt[int64] `json:"maximum_digits,omitzero"`
	// The maximum number of times the file should be played if there is no input from
	// the user on the call.
	MaximumTries param.Opt[int64] `json:"maximum_tries,omitzero"`
	// The media_name of a file to be played back at the beginning of each prompt. The
	// media_name must point to a file previously uploaded to api.telnyx.com/v2/media
	// by the same user/organization. The file must either be a WAV or MP3 file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// The minimum number of digits to fetch. This parameter has a minimum value of 1.
	MinimumDigits param.Opt[int64] `json:"minimum_digits,omitzero"`
	// The digit used to terminate input if fewer than `maximum_digits` digits have
	// been gathered.
	TerminatingDigit param.Opt[string] `json:"terminating_digit,omitzero"`
	// The number of milliseconds to wait for a DTMF response after file playback ends
	// before a replaying the sound file.
	TimeoutMillis param.Opt[int64] `json:"timeout_millis,omitzero"`
	// A list of all digits accepted as valid.
	ValidDigits param.Opt[string] `json:"valid_digits,omitzero"`
	paramObj
}

func (r CallActionGatherUsingAudioParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionGatherUsingAudioParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionGatherUsingAudioParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionGatherUsingSpeakParams struct {
	// The text or SSML to be converted into speech. There is a 3,000 character limit.
	Payload string `json:"payload,required"`
	// Specifies the voice used in speech synthesis.
	//
	//   - Define voices using the format `<Provider>.<Model>.<VoiceId>`. Specifying only
	//     the provider will give default values for voice_id and model_id.
	//
	//     **Supported Providers:**
	//
	//   - **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural
	//     voices, which provide more realistic, human-like speech, append `-Neural` to
	//     the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the
	//     [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html)
	//     for compatibility.
	//   - **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural,
	//     Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural,
	//     Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to
	//     [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)
	//   - **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g.,
	//     `ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM`). The `ModelId` part
	//     is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an
	//     integration identifier secret in
	//     `"voice_settings": {"api_key_ref": "<secret_identifier>"}`. See
	//     [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	//     for details. Check
	//     [available voices](https://elevenlabs.io/docs/api-reference/get-voices).
	//   - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`
	//
	// For service_level basic, you may define the gender of the speaker (male or
	// female).
	Voice string `json:"voice,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The number of milliseconds to wait for input between digits.
	InterDigitTimeoutMillis param.Opt[int64] `json:"inter_digit_timeout_millis,omitzero"`
	// The text or SSML to be converted into speech when digits don't match the
	// `valid_digits` parameter or the number of digits is not between `min` and `max`.
	// There is a 3,000 character limit.
	InvalidPayload param.Opt[string] `json:"invalid_payload,omitzero"`
	// The maximum number of digits to fetch. This parameter has a maximum value
	// of 128.
	MaximumDigits param.Opt[int64] `json:"maximum_digits,omitzero"`
	// The maximum number of times that a file should be played back if there is no
	// input from the user on the call.
	MaximumTries param.Opt[int64] `json:"maximum_tries,omitzero"`
	// The minimum number of digits to fetch. This parameter has a minimum value of 1.
	MinimumDigits param.Opt[int64] `json:"minimum_digits,omitzero"`
	// The digit used to terminate input if fewer than `maximum_digits` digits have
	// been gathered.
	TerminatingDigit param.Opt[string] `json:"terminating_digit,omitzero"`
	// The number of milliseconds to wait for a DTMF response after speak ends before a
	// replaying the sound file.
	TimeoutMillis param.Opt[int64] `json:"timeout_millis,omitzero"`
	// A list of all digits accepted as valid.
	ValidDigits param.Opt[string] `json:"valid_digits,omitzero"`
	// The language you want spoken. This parameter is ignored when a `Polly.*` voice
	// is specified.
	//
	// Any of "arb", "cmn-CN", "cy-GB", "da-DK", "de-DE", "en-AU", "en-GB",
	// "en-GB-WLS", "en-IN", "en-US", "es-ES", "es-MX", "es-US", "fr-CA", "fr-FR",
	// "hi-IN", "is-IS", "it-IT", "ja-JP", "ko-KR", "nb-NO", "nl-NL", "pl-PL", "pt-BR",
	// "pt-PT", "ro-RO", "ru-RU", "sv-SE", "tr-TR".
	Language CallActionGatherUsingSpeakParamsLanguage `json:"language,omitzero"`
	// The type of the provided payload. The payload can either be plain text, or
	// Speech Synthesis Markup Language (SSML).
	//
	// Any of "text", "ssml".
	PayloadType CallActionGatherUsingSpeakParamsPayloadType `json:"payload_type,omitzero"`
	// This parameter impacts speech quality, language options and payload types. When
	// using `basic`, only the `en-US` language and payload type `text` are allowed.
	//
	// Any of "basic", "premium".
	ServiceLevel CallActionGatherUsingSpeakParamsServiceLevel `json:"service_level,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings CallActionGatherUsingSpeakParamsVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r CallActionGatherUsingSpeakParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionGatherUsingSpeakParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionGatherUsingSpeakParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language you want spoken. This parameter is ignored when a `Polly.*` voice
// is specified.
type CallActionGatherUsingSpeakParamsLanguage string

const (
	CallActionGatherUsingSpeakParamsLanguageArb     CallActionGatherUsingSpeakParamsLanguage = "arb"
	CallActionGatherUsingSpeakParamsLanguageCmnCn   CallActionGatherUsingSpeakParamsLanguage = "cmn-CN"
	CallActionGatherUsingSpeakParamsLanguageCyGB    CallActionGatherUsingSpeakParamsLanguage = "cy-GB"
	CallActionGatherUsingSpeakParamsLanguageDaDk    CallActionGatherUsingSpeakParamsLanguage = "da-DK"
	CallActionGatherUsingSpeakParamsLanguageDeDe    CallActionGatherUsingSpeakParamsLanguage = "de-DE"
	CallActionGatherUsingSpeakParamsLanguageEnAu    CallActionGatherUsingSpeakParamsLanguage = "en-AU"
	CallActionGatherUsingSpeakParamsLanguageEnGB    CallActionGatherUsingSpeakParamsLanguage = "en-GB"
	CallActionGatherUsingSpeakParamsLanguageEnGBWls CallActionGatherUsingSpeakParamsLanguage = "en-GB-WLS"
	CallActionGatherUsingSpeakParamsLanguageEnIn    CallActionGatherUsingSpeakParamsLanguage = "en-IN"
	CallActionGatherUsingSpeakParamsLanguageEnUs    CallActionGatherUsingSpeakParamsLanguage = "en-US"
	CallActionGatherUsingSpeakParamsLanguageEsEs    CallActionGatherUsingSpeakParamsLanguage = "es-ES"
	CallActionGatherUsingSpeakParamsLanguageEsMx    CallActionGatherUsingSpeakParamsLanguage = "es-MX"
	CallActionGatherUsingSpeakParamsLanguageEsUs    CallActionGatherUsingSpeakParamsLanguage = "es-US"
	CallActionGatherUsingSpeakParamsLanguageFrCa    CallActionGatherUsingSpeakParamsLanguage = "fr-CA"
	CallActionGatherUsingSpeakParamsLanguageFrFr    CallActionGatherUsingSpeakParamsLanguage = "fr-FR"
	CallActionGatherUsingSpeakParamsLanguageHiIn    CallActionGatherUsingSpeakParamsLanguage = "hi-IN"
	CallActionGatherUsingSpeakParamsLanguageIsIs    CallActionGatherUsingSpeakParamsLanguage = "is-IS"
	CallActionGatherUsingSpeakParamsLanguageItIt    CallActionGatherUsingSpeakParamsLanguage = "it-IT"
	CallActionGatherUsingSpeakParamsLanguageJaJp    CallActionGatherUsingSpeakParamsLanguage = "ja-JP"
	CallActionGatherUsingSpeakParamsLanguageKoKr    CallActionGatherUsingSpeakParamsLanguage = "ko-KR"
	CallActionGatherUsingSpeakParamsLanguageNbNo    CallActionGatherUsingSpeakParamsLanguage = "nb-NO"
	CallActionGatherUsingSpeakParamsLanguageNlNl    CallActionGatherUsingSpeakParamsLanguage = "nl-NL"
	CallActionGatherUsingSpeakParamsLanguagePlPl    CallActionGatherUsingSpeakParamsLanguage = "pl-PL"
	CallActionGatherUsingSpeakParamsLanguagePtBr    CallActionGatherUsingSpeakParamsLanguage = "pt-BR"
	CallActionGatherUsingSpeakParamsLanguagePtPt    CallActionGatherUsingSpeakParamsLanguage = "pt-PT"
	CallActionGatherUsingSpeakParamsLanguageRoRo    CallActionGatherUsingSpeakParamsLanguage = "ro-RO"
	CallActionGatherUsingSpeakParamsLanguageRuRu    CallActionGatherUsingSpeakParamsLanguage = "ru-RU"
	CallActionGatherUsingSpeakParamsLanguageSvSe    CallActionGatherUsingSpeakParamsLanguage = "sv-SE"
	CallActionGatherUsingSpeakParamsLanguageTrTr    CallActionGatherUsingSpeakParamsLanguage = "tr-TR"
)

// The type of the provided payload. The payload can either be plain text, or
// Speech Synthesis Markup Language (SSML).
type CallActionGatherUsingSpeakParamsPayloadType string

const (
	CallActionGatherUsingSpeakParamsPayloadTypeText CallActionGatherUsingSpeakParamsPayloadType = "text"
	CallActionGatherUsingSpeakParamsPayloadTypeSsml CallActionGatherUsingSpeakParamsPayloadType = "ssml"
)

// This parameter impacts speech quality, language options and payload types. When
// using `basic`, only the `en-US` language and payload type `text` are allowed.
type CallActionGatherUsingSpeakParamsServiceLevel string

const (
	CallActionGatherUsingSpeakParamsServiceLevelBasic   CallActionGatherUsingSpeakParamsServiceLevel = "basic"
	CallActionGatherUsingSpeakParamsServiceLevelPremium CallActionGatherUsingSpeakParamsServiceLevel = "premium"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallActionGatherUsingSpeakParamsVoiceSettingsUnion struct {
	OfElevenLabsVoiceSettings *ElevenLabsVoiceSettingsParam `json:",omitzero,inline"`
	OfTelnyxVoiceSettings     *TelnyxVoiceSettingsParam     `json:",omitzero,inline"`
	paramUnion
}

func (u CallActionGatherUsingSpeakParamsVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenLabsVoiceSettings, u.OfTelnyxVoiceSettings)
}
func (u *CallActionGatherUsingSpeakParamsVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallActionGatherUsingSpeakParamsVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenLabsVoiceSettings) {
		return u.OfElevenLabsVoiceSettings
	} else if !param.IsOmitted(u.OfTelnyxVoiceSettings) {
		return u.OfTelnyxVoiceSettings
	}
	return nil
}

type CallActionHangupParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionHangupParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionHangupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionHangupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionLeaveQueueParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionLeaveQueueParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionLeaveQueueParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionLeaveQueueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionPauseRecordingParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Uniquely identifies the resource.
	RecordingID param.Opt[string] `json:"recording_id,omitzero" format:"uuid"`
	paramObj
}

func (r CallActionPauseRecordingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionPauseRecordingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionPauseRecordingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionReferParams struct {
	// The SIP URI to which the call will be referred to.
	SipAddress string `json:"sip_address,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore
	// subsequent commands with the same `command_id` as one that has already been
	// executed.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword param.Opt[string] `json:"sip_auth_password,omitzero"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername param.Opt[string] `json:"sip_auth_username,omitzero"`
	// Custom headers to be added to the SIP INVITE.
	CustomHeaders []CustomSipHeaderParam `json:"custom_headers,omitzero"`
	// SIP headers to be added to the request. Currently only User-to-User header is
	// supported.
	SipHeaders []SipHeaderParam `json:"sip_headers,omitzero"`
	paramObj
}

func (r CallActionReferParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionReferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionReferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionRejectParams struct {
	// Cause for call rejection.
	//
	// Any of "CALL_REJECTED", "USER_BUSY".
	Cause CallActionRejectParamsCause `json:"cause,omitzero,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionRejectParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionRejectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionRejectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cause for call rejection.
type CallActionRejectParamsCause string

const (
	CallActionRejectParamsCauseCallRejected CallActionRejectParamsCause = "CALL_REJECTED"
	CallActionRejectParamsCauseUserBusy     CallActionRejectParamsCause = "USER_BUSY"
)

type CallActionResumeRecordingParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Uniquely identifies the resource.
	RecordingID param.Opt[string] `json:"recording_id,omitzero" format:"uuid"`
	paramObj
}

func (r CallActionResumeRecordingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionResumeRecordingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionResumeRecordingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSendDtmfParams struct {
	// DTMF digits to send. Valid digits are 0-9, A-D, \*, and #. Pauses can be added
	// using w (0.5s) and W (1s).
	Digits string `json:"digits,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Specifies for how many milliseconds each digit will be played in the audio
	// stream. Ranges from 100 to 500ms
	DurationMillis param.Opt[int64] `json:"duration_millis,omitzero"`
	paramObj
}

func (r CallActionSendDtmfParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionSendDtmfParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionSendDtmfParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSendSipInfoParams struct {
	// Content of the SIP INFO
	Body string `json:"body,required"`
	// Content type of the INFO body. Must be MIME type compliant. There is a 1,400
	// bytes limit
	ContentType string `json:"content_type,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionSendSipInfoParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionSendSipInfoParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionSendSipInfoParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSpeakParams struct {
	// The text or SSML to be converted into speech. There is a 3,000 character limit.
	Payload string `json:"payload,required"`
	// Specifies the voice used in speech synthesis.
	//
	//   - Define voices using the format `<Provider>.<Model>.<VoiceId>`. Specifying only
	//     the provider will give default values for voice_id and model_id.
	//
	//     **Supported Providers:**
	//
	//   - **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural
	//     voices, which provide more realistic, human-like speech, append `-Neural` to
	//     the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the
	//     [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html)
	//     for compatibility.
	//   - **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural,
	//     Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural,
	//     Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to
	//     [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)
	//   - **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g.,
	//     `ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM`). The `ModelId` part
	//     is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an
	//     integration identifier secret in
	//     `"voice_settings": {"api_key_ref": "<secret_identifier>"}`. See
	//     [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	//     for details. Check
	//     [available voices](https://elevenlabs.io/docs/api-reference/get-voices).
	//   - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`
	//
	// For service_level basic, you may define the gender of the speaker (male or
	// female).
	Voice string `json:"voice,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// When specified, it stops the current audio being played. Specify `current` to
	// stop the current audio being played, and to play the next file in the queue.
	// Specify `all` to stop the current audio file being played and to also clear all
	// audio files from the queue.
	Stop param.Opt[string] `json:"stop,omitzero"`
	// The language you want spoken. This parameter is ignored when a `Polly.*` voice
	// is specified.
	//
	// Any of "arb", "cmn-CN", "cy-GB", "da-DK", "de-DE", "en-AU", "en-GB",
	// "en-GB-WLS", "en-IN", "en-US", "es-ES", "es-MX", "es-US", "fr-CA", "fr-FR",
	// "hi-IN", "is-IS", "it-IT", "ja-JP", "ko-KR", "nb-NO", "nl-NL", "pl-PL", "pt-BR",
	// "pt-PT", "ro-RO", "ru-RU", "sv-SE", "tr-TR".
	Language CallActionSpeakParamsLanguage `json:"language,omitzero"`
	// The type of the provided payload. The payload can either be plain text, or
	// Speech Synthesis Markup Language (SSML).
	//
	// Any of "text", "ssml".
	PayloadType CallActionSpeakParamsPayloadType `json:"payload_type,omitzero"`
	// This parameter impacts speech quality, language options and payload types. When
	// using `basic`, only the `en-US` language and payload type `text` are allowed.
	//
	// Any of "basic", "premium".
	ServiceLevel CallActionSpeakParamsServiceLevel `json:"service_level,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings CallActionSpeakParamsVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r CallActionSpeakParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionSpeakParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionSpeakParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language you want spoken. This parameter is ignored when a `Polly.*` voice
// is specified.
type CallActionSpeakParamsLanguage string

const (
	CallActionSpeakParamsLanguageArb     CallActionSpeakParamsLanguage = "arb"
	CallActionSpeakParamsLanguageCmnCn   CallActionSpeakParamsLanguage = "cmn-CN"
	CallActionSpeakParamsLanguageCyGB    CallActionSpeakParamsLanguage = "cy-GB"
	CallActionSpeakParamsLanguageDaDk    CallActionSpeakParamsLanguage = "da-DK"
	CallActionSpeakParamsLanguageDeDe    CallActionSpeakParamsLanguage = "de-DE"
	CallActionSpeakParamsLanguageEnAu    CallActionSpeakParamsLanguage = "en-AU"
	CallActionSpeakParamsLanguageEnGB    CallActionSpeakParamsLanguage = "en-GB"
	CallActionSpeakParamsLanguageEnGBWls CallActionSpeakParamsLanguage = "en-GB-WLS"
	CallActionSpeakParamsLanguageEnIn    CallActionSpeakParamsLanguage = "en-IN"
	CallActionSpeakParamsLanguageEnUs    CallActionSpeakParamsLanguage = "en-US"
	CallActionSpeakParamsLanguageEsEs    CallActionSpeakParamsLanguage = "es-ES"
	CallActionSpeakParamsLanguageEsMx    CallActionSpeakParamsLanguage = "es-MX"
	CallActionSpeakParamsLanguageEsUs    CallActionSpeakParamsLanguage = "es-US"
	CallActionSpeakParamsLanguageFrCa    CallActionSpeakParamsLanguage = "fr-CA"
	CallActionSpeakParamsLanguageFrFr    CallActionSpeakParamsLanguage = "fr-FR"
	CallActionSpeakParamsLanguageHiIn    CallActionSpeakParamsLanguage = "hi-IN"
	CallActionSpeakParamsLanguageIsIs    CallActionSpeakParamsLanguage = "is-IS"
	CallActionSpeakParamsLanguageItIt    CallActionSpeakParamsLanguage = "it-IT"
	CallActionSpeakParamsLanguageJaJp    CallActionSpeakParamsLanguage = "ja-JP"
	CallActionSpeakParamsLanguageKoKr    CallActionSpeakParamsLanguage = "ko-KR"
	CallActionSpeakParamsLanguageNbNo    CallActionSpeakParamsLanguage = "nb-NO"
	CallActionSpeakParamsLanguageNlNl    CallActionSpeakParamsLanguage = "nl-NL"
	CallActionSpeakParamsLanguagePlPl    CallActionSpeakParamsLanguage = "pl-PL"
	CallActionSpeakParamsLanguagePtBr    CallActionSpeakParamsLanguage = "pt-BR"
	CallActionSpeakParamsLanguagePtPt    CallActionSpeakParamsLanguage = "pt-PT"
	CallActionSpeakParamsLanguageRoRo    CallActionSpeakParamsLanguage = "ro-RO"
	CallActionSpeakParamsLanguageRuRu    CallActionSpeakParamsLanguage = "ru-RU"
	CallActionSpeakParamsLanguageSvSe    CallActionSpeakParamsLanguage = "sv-SE"
	CallActionSpeakParamsLanguageTrTr    CallActionSpeakParamsLanguage = "tr-TR"
)

// The type of the provided payload. The payload can either be plain text, or
// Speech Synthesis Markup Language (SSML).
type CallActionSpeakParamsPayloadType string

const (
	CallActionSpeakParamsPayloadTypeText CallActionSpeakParamsPayloadType = "text"
	CallActionSpeakParamsPayloadTypeSsml CallActionSpeakParamsPayloadType = "ssml"
)

// This parameter impacts speech quality, language options and payload types. When
// using `basic`, only the `en-US` language and payload type `text` are allowed.
type CallActionSpeakParamsServiceLevel string

const (
	CallActionSpeakParamsServiceLevelBasic   CallActionSpeakParamsServiceLevel = "basic"
	CallActionSpeakParamsServiceLevelPremium CallActionSpeakParamsServiceLevel = "premium"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallActionSpeakParamsVoiceSettingsUnion struct {
	OfElevenLabsVoiceSettings *ElevenLabsVoiceSettingsParam `json:",omitzero,inline"`
	OfTelnyxVoiceSettings     *TelnyxVoiceSettingsParam     `json:",omitzero,inline"`
	paramUnion
}

func (u CallActionSpeakParamsVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenLabsVoiceSettings, u.OfTelnyxVoiceSettings)
}
func (u *CallActionSpeakParamsVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallActionSpeakParamsVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenLabsVoiceSettings) {
		return u.OfElevenLabsVoiceSettings
	} else if !param.IsOmitted(u.OfTelnyxVoiceSettings) {
		return u.OfTelnyxVoiceSettings
	}
	return nil
}

type CallActionStartAIAssistantParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Text that will be played when the assistant starts, if none then nothing will be
	// played when the assistant starts. The greeting can be text for any voice or SSML
	// for `AWS.Polly.<voice_id>` voices. There is a 3,000 character limit.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// The voice to be used by the voice assistant. Currently we support ElevenLabs,
	// Telnyx and AWS voices.
	//
	// **Supported Providers:**
	//
	//   - **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural
	//     voices, which provide more realistic, human-like speech, append `-Neural` to
	//     the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the
	//     [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html)
	//     for compatibility.
	//   - **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural,
	//     Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural,
	//     Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to
	//     [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)
	//   - **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g.,
	//     `ElevenLabs.BaseModel.John`). The `ModelId` part is optional. To use
	//     ElevenLabs, you must provide your ElevenLabs API key as an integration secret
	//     under `"voice_settings": {"api_key_ref": "<secret_id>"}`. See
	//     [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	//     for details. Check
	//     [available voices](https://elevenlabs.io/docs/api-reference/get-voices).
	//   - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`
	Voice param.Opt[string] `json:"voice,omitzero"`
	// AI Assistant configuration
	Assistant CallActionStartAIAssistantParamsAssistant `json:"assistant,omitzero"`
	// Settings for handling user interruptions during assistant speech
	InterruptionSettings InterruptionSettingsParam `json:"interruption_settings,omitzero"`
	// The settings associated with speech to text for the voice assistant. This is
	// only relevant if the assistant uses a text-to-text language model. Any assistant
	// using a model with native audio support (e.g. `fixie-ai/ultravox-v0_4`) will
	// ignore this field.
	Transcription TranscriptionConfigParam `json:"transcription,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings CallActionStartAIAssistantParamsVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r CallActionStartAIAssistantParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartAIAssistantParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartAIAssistantParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AI Assistant configuration
type CallActionStartAIAssistantParamsAssistant struct {
	// The identifier of the AI assistant to use
	ID param.Opt[string] `json:"id,omitzero"`
	// The system instructions that the voice assistant uses during the start assistant
	// command. This will overwrite the instructions set in the assistant
	// configuration.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Reference to the OpenAI API key. Required only when using OpenAI models
	OpenAIAPIKeyRef param.Opt[string] `json:"openai_api_key_ref,omitzero"`
	paramObj
}

func (r CallActionStartAIAssistantParamsAssistant) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartAIAssistantParamsAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartAIAssistantParamsAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallActionStartAIAssistantParamsVoiceSettingsUnion struct {
	OfElevenLabsVoiceSettings *ElevenLabsVoiceSettingsParam `json:",omitzero,inline"`
	OfTelnyxVoiceSettings     *TelnyxVoiceSettingsParam     `json:",omitzero,inline"`
	paramUnion
}

func (u CallActionStartAIAssistantParamsVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenLabsVoiceSettings, u.OfTelnyxVoiceSettings)
}
func (u *CallActionStartAIAssistantParamsVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallActionStartAIAssistantParamsVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenLabsVoiceSettings) {
		return u.OfElevenLabsVoiceSettings
	} else if !param.IsOmitted(u.OfTelnyxVoiceSettings) {
		return u.OfTelnyxVoiceSettings
	}
	return nil
}

type CallActionStartForkingParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The network target, <udp:ip_address:port>, where the call's incoming RTP media
	// packets should be forwarded.
	Rx param.Opt[string] `json:"rx,omitzero"`
	// The network target, <udp:ip_address:port>, where the call's outgoing RTP media
	// packets should be forwarded.
	Tx param.Opt[string] `json:"tx,omitzero"`
	// Optionally specify a media type to stream. If `decrypted` selected, Telnyx will
	// decrypt incoming SIP media before forking to the target. `rx` and `tx` are
	// required fields if `decrypted` selected.
	//
	// Any of "decrypted".
	StreamType CallActionStartForkingParamsStreamType `json:"stream_type,omitzero"`
	paramObj
}

func (r CallActionStartForkingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartForkingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartForkingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optionally specify a media type to stream. If `decrypted` selected, Telnyx will
// decrypt incoming SIP media before forking to the target. `rx` and `tx` are
// required fields if `decrypted` selected.
type CallActionStartForkingParamsStreamType string

const (
	CallActionStartForkingParamsStreamTypeDecrypted CallActionStartForkingParamsStreamType = "decrypted"
)

type CallActionStartNoiseSuppressionParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The direction of the audio stream to be noise suppressed.
	//
	// Any of "inbound", "outbound", "both".
	Direction CallActionStartNoiseSuppressionParamsDirection `json:"direction,omitzero"`
	// The engine to use for noise suppression. A - rnnoise engine B - deepfilter
	// engine.
	//
	// Any of "A", "B".
	NoiseSuppressionEngine CallActionStartNoiseSuppressionParamsNoiseSuppressionEngine `json:"noise_suppression_engine,omitzero"`
	paramObj
}

func (r CallActionStartNoiseSuppressionParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartNoiseSuppressionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartNoiseSuppressionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the audio stream to be noise suppressed.
type CallActionStartNoiseSuppressionParamsDirection string

const (
	CallActionStartNoiseSuppressionParamsDirectionInbound  CallActionStartNoiseSuppressionParamsDirection = "inbound"
	CallActionStartNoiseSuppressionParamsDirectionOutbound CallActionStartNoiseSuppressionParamsDirection = "outbound"
	CallActionStartNoiseSuppressionParamsDirectionBoth     CallActionStartNoiseSuppressionParamsDirection = "both"
)

// The engine to use for noise suppression. A - rnnoise engine B - deepfilter
// engine.
type CallActionStartNoiseSuppressionParamsNoiseSuppressionEngine string

const (
	CallActionStartNoiseSuppressionParamsNoiseSuppressionEngineA CallActionStartNoiseSuppressionParamsNoiseSuppressionEngine = "A"
	CallActionStartNoiseSuppressionParamsNoiseSuppressionEngineB CallActionStartNoiseSuppressionParamsNoiseSuppressionEngine = "B"
)

type CallActionStartPlaybackParams struct {
	// The URL of a file to be played back on the call. The URL can point to either a
	// WAV or MP3 file. media_name and audio_url cannot be used together in one
	// request.
	AudioURL param.Opt[string] `json:"audio_url,omitzero"`
	// Caches the audio file. Useful when playing the same audio file multiple times
	// during the call.
	CacheAudio param.Opt[bool] `json:"cache_audio,omitzero"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The media_name of a file to be played back on the call. The media_name must
	// point to a file previously uploaded to api.telnyx.com/v2/media by the same
	// user/organization. The file must either be a WAV or MP3 file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// When enabled, audio will be mixed on top of any other audio that is actively
	// being played back. Note that `overlay: true` will only work if there is another
	// audio file already being played on the call.
	Overlay param.Opt[bool] `json:"overlay,omitzero"`
	// Allows a user to provide base64 encoded mp3 or wav. Note: when using this
	// parameter, `media_url` and `media_name` in the `playback_started` and
	// `playback_ended` webhooks will be empty
	PlaybackContent param.Opt[string] `json:"playback_content,omitzero"`
	// When specified, it stops the current audio being played. Specify `current` to
	// stop the current audio being played, and to play the next file in the queue.
	// Specify `all` to stop the current audio file being played and to also clear all
	// audio files from the queue.
	Stop param.Opt[string] `json:"stop,omitzero"`
	// Specifies the leg or legs on which audio will be played. If supplied, the value
	// must be either `self`, `opposite` or `both`.
	TargetLegs param.Opt[string] `json:"target_legs,omitzero"`
	// Specifies the type of audio provided in `audio_url` or `playback_content`.
	//
	// Any of "mp3", "wav".
	AudioType CallActionStartPlaybackParamsAudioType `json:"audio_type,omitzero"`
	// The number of times the audio file should be played. If supplied, the value must
	// be an integer between 1 and 100, or the special string `infinity` for an endless
	// loop.
	Loop LoopcountUnionParam `json:"loop,omitzero"`
	paramObj
}

func (r CallActionStartPlaybackParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartPlaybackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartPlaybackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of audio provided in `audio_url` or `playback_content`.
type CallActionStartPlaybackParamsAudioType string

const (
	CallActionStartPlaybackParamsAudioTypeMP3 CallActionStartPlaybackParamsAudioType = "mp3"
	CallActionStartPlaybackParamsAudioTypeWav CallActionStartPlaybackParamsAudioType = "wav"
)

type CallActionStartRecordingParams struct {
	// When `dual`, final audio file will be stereo recorded with the first leg on
	// channel A, and the rest on channel B.
	//
	// Any of "single", "dual".
	Channels CallActionStartRecordingParamsChannels `json:"channels,omitzero,required"`
	// The audio file format used when storing the call recording. Can be either `mp3`
	// or `wav`.
	//
	// Any of "wav", "mp3".
	Format CallActionStartRecordingParamsFormat `json:"format,omitzero,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The custom recording file name to be used instead of the default `call_leg_id`.
	// Telnyx will still add a Unix timestamp suffix.
	CustomFileName param.Opt[string] `json:"custom_file_name,omitzero"`
	// Defines the maximum length for the recording in seconds. The minimum value is 0.
	// The maximum value is 14400. The default value is 0 (infinite)
	MaxLength param.Opt[int64] `json:"max_length,omitzero"`
	// If enabled, a beep sound will be played at the start of a recording.
	PlayBeep param.Opt[bool] `json:"play_beep,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected. The timer only starts when the speech is detected. Please
	// note that call transcription is used to detect silence and the related charge
	// will be applied. The minimum value is 0. The default value is 0 (infinite)
	TimeoutSecs param.Opt[int64] `json:"timeout_secs,omitzero"`
	// Enable post recording transcription. The default value is false.
	Transcription param.Opt[bool] `json:"transcription,omitzero"`
	// Engine to use for speech recognition. `A` - `Google`
	TranscriptionEngine param.Opt[string] `json:"transcription_engine,omitzero"`
	// Defines maximum number of speakers in the conversation. Applies to `google`
	// engine only.
	TranscriptionMaxSpeakerCount param.Opt[int64] `json:"transcription_max_speaker_count,omitzero"`
	// Defines minimum number of speakers in the conversation. Applies to `google`
	// engine only.
	TranscriptionMinSpeakerCount param.Opt[int64] `json:"transcription_min_speaker_count,omitzero"`
	// Enables profanity_filter. Applies to `google` engine only.
	TranscriptionProfanityFilter param.Opt[bool] `json:"transcription_profanity_filter,omitzero"`
	// Enables speaker diarization. Applies to `google` engine only.
	TranscriptionSpeakerDiarization param.Opt[bool] `json:"transcription_speaker_diarization,omitzero"`
	// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
	// If only single track is specified (`inbound`, `outbound`), `channels`
	// configuration is ignored and it will be recorded as mono (single channel).
	//
	// Any of "both", "inbound", "outbound".
	RecordingTrack CallActionStartRecordingParamsRecordingTrack `json:"recording_track,omitzero"`
	// Language to use for speech recognition
	//
	// Any of "af-ZA", "am-ET", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IL", "ar-IQ",
	// "ar-JO", "ar-KW", "ar-LB", "ar-MA", "ar-MR", "ar-OM", "ar-PS", "ar-QA", "ar-SA",
	// "ar-TN", "ar-YE", "az-AZ", "bg-BG", "bn-BD", "bn-IN", "bs-BA", "ca-ES", "cs-CZ",
	// "da-DK", "de-AT", "de-CH", "de-DE", "el-GR", "en-AU", "en-CA", "en-GB", "en-GH",
	// "en-HK", "en-IE", "en-IN", "en-KE", "en-NG", "en-NZ", "en-PH", "en-PK", "en-SG",
	// "en-TZ", "en-US", "en-ZA", "es-AR", "es-BO", "es-CL", "es-CO", "es-CR", "es-DO",
	// "es-EC", "es-ES", "es-GT", "es-HN", "es-MX", "es-NI", "es-PA", "es-PE", "es-PR",
	// "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et-EE", "eu-ES", "fa-IR", "fi-FI",
	// "fil-PH", "fr-BE", "fr-CA", "fr-CH", "fr-FR", "gl-ES", "gu-IN", "hi-IN",
	// "hr-HR", "hu-HU", "hy-AM", "id-ID", "is-IS", "it-CH", "it-IT", "iw-IL", "ja-JP",
	// "jv-ID", "ka-GE", "kk-KZ", "km-KH", "kn-IN", "ko-KR", "lo-LA", "lt-LT", "lv-LV",
	// "mk-MK", "ml-IN", "mn-MN", "mr-IN", "ms-MY", "my-MM", "ne-NP", "nl-BE", "nl-NL",
	// "no-NO", "pa-Guru-IN", "pl-PL", "pt-BR", "pt-PT", "ro-RO", "ru-RU", "rw-RW",
	// "si-LK", "sk-SK", "sl-SI", "sq-AL", "sr-RS", "ss-latn-za", "st-ZA", "su-ID",
	// "sv-SE", "sw-KE", "sw-TZ", "ta-IN", "ta-LK", "ta-MY", "ta-SG", "te-IN", "th-TH",
	// "tn-latn-za", "tr-TR", "ts-ZA", "uk-UA", "ur-IN", "ur-PK", "uz-UZ", "ve-ZA",
	// "vi-VN", "xh-ZA", "yue-Hant-HK", "zh", "zh-TW", "zu-ZA".
	TranscriptionLanguage CallActionStartRecordingParamsTranscriptionLanguage `json:"transcription_language,omitzero"`
	// When set to `trim-silence`, silence will be removed from the beginning and end
	// of the recording.
	//
	// Any of "trim-silence".
	Trim CallActionStartRecordingParamsTrim `json:"trim,omitzero"`
	paramObj
}

func (r CallActionStartRecordingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartRecordingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartRecordingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When `dual`, final audio file will be stereo recorded with the first leg on
// channel A, and the rest on channel B.
type CallActionStartRecordingParamsChannels string

const (
	CallActionStartRecordingParamsChannelsSingle CallActionStartRecordingParamsChannels = "single"
	CallActionStartRecordingParamsChannelsDual   CallActionStartRecordingParamsChannels = "dual"
)

// The audio file format used when storing the call recording. Can be either `mp3`
// or `wav`.
type CallActionStartRecordingParamsFormat string

const (
	CallActionStartRecordingParamsFormatWav CallActionStartRecordingParamsFormat = "wav"
	CallActionStartRecordingParamsFormatMP3 CallActionStartRecordingParamsFormat = "mp3"
)

// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
// If only single track is specified (`inbound`, `outbound`), `channels`
// configuration is ignored and it will be recorded as mono (single channel).
type CallActionStartRecordingParamsRecordingTrack string

const (
	CallActionStartRecordingParamsRecordingTrackBoth     CallActionStartRecordingParamsRecordingTrack = "both"
	CallActionStartRecordingParamsRecordingTrackInbound  CallActionStartRecordingParamsRecordingTrack = "inbound"
	CallActionStartRecordingParamsRecordingTrackOutbound CallActionStartRecordingParamsRecordingTrack = "outbound"
)

// Language to use for speech recognition
type CallActionStartRecordingParamsTranscriptionLanguage string

const (
	CallActionStartRecordingParamsTranscriptionLanguageAfZa      CallActionStartRecordingParamsTranscriptionLanguage = "af-ZA"
	CallActionStartRecordingParamsTranscriptionLanguageAmEt      CallActionStartRecordingParamsTranscriptionLanguage = "am-ET"
	CallActionStartRecordingParamsTranscriptionLanguageArAe      CallActionStartRecordingParamsTranscriptionLanguage = "ar-AE"
	CallActionStartRecordingParamsTranscriptionLanguageArBh      CallActionStartRecordingParamsTranscriptionLanguage = "ar-BH"
	CallActionStartRecordingParamsTranscriptionLanguageArDz      CallActionStartRecordingParamsTranscriptionLanguage = "ar-DZ"
	CallActionStartRecordingParamsTranscriptionLanguageArEg      CallActionStartRecordingParamsTranscriptionLanguage = "ar-EG"
	CallActionStartRecordingParamsTranscriptionLanguageArIl      CallActionStartRecordingParamsTranscriptionLanguage = "ar-IL"
	CallActionStartRecordingParamsTranscriptionLanguageArIq      CallActionStartRecordingParamsTranscriptionLanguage = "ar-IQ"
	CallActionStartRecordingParamsTranscriptionLanguageArJo      CallActionStartRecordingParamsTranscriptionLanguage = "ar-JO"
	CallActionStartRecordingParamsTranscriptionLanguageArKw      CallActionStartRecordingParamsTranscriptionLanguage = "ar-KW"
	CallActionStartRecordingParamsTranscriptionLanguageArLb      CallActionStartRecordingParamsTranscriptionLanguage = "ar-LB"
	CallActionStartRecordingParamsTranscriptionLanguageArMa      CallActionStartRecordingParamsTranscriptionLanguage = "ar-MA"
	CallActionStartRecordingParamsTranscriptionLanguageArMr      CallActionStartRecordingParamsTranscriptionLanguage = "ar-MR"
	CallActionStartRecordingParamsTranscriptionLanguageArOm      CallActionStartRecordingParamsTranscriptionLanguage = "ar-OM"
	CallActionStartRecordingParamsTranscriptionLanguageArPs      CallActionStartRecordingParamsTranscriptionLanguage = "ar-PS"
	CallActionStartRecordingParamsTranscriptionLanguageArQa      CallActionStartRecordingParamsTranscriptionLanguage = "ar-QA"
	CallActionStartRecordingParamsTranscriptionLanguageArSa      CallActionStartRecordingParamsTranscriptionLanguage = "ar-SA"
	CallActionStartRecordingParamsTranscriptionLanguageArTn      CallActionStartRecordingParamsTranscriptionLanguage = "ar-TN"
	CallActionStartRecordingParamsTranscriptionLanguageArYe      CallActionStartRecordingParamsTranscriptionLanguage = "ar-YE"
	CallActionStartRecordingParamsTranscriptionLanguageAzAz      CallActionStartRecordingParamsTranscriptionLanguage = "az-AZ"
	CallActionStartRecordingParamsTranscriptionLanguageBgBg      CallActionStartRecordingParamsTranscriptionLanguage = "bg-BG"
	CallActionStartRecordingParamsTranscriptionLanguageBnBd      CallActionStartRecordingParamsTranscriptionLanguage = "bn-BD"
	CallActionStartRecordingParamsTranscriptionLanguageBnIn      CallActionStartRecordingParamsTranscriptionLanguage = "bn-IN"
	CallActionStartRecordingParamsTranscriptionLanguageBsBa      CallActionStartRecordingParamsTranscriptionLanguage = "bs-BA"
	CallActionStartRecordingParamsTranscriptionLanguageCaEs      CallActionStartRecordingParamsTranscriptionLanguage = "ca-ES"
	CallActionStartRecordingParamsTranscriptionLanguageCsCz      CallActionStartRecordingParamsTranscriptionLanguage = "cs-CZ"
	CallActionStartRecordingParamsTranscriptionLanguageDaDk      CallActionStartRecordingParamsTranscriptionLanguage = "da-DK"
	CallActionStartRecordingParamsTranscriptionLanguageDeAt      CallActionStartRecordingParamsTranscriptionLanguage = "de-AT"
	CallActionStartRecordingParamsTranscriptionLanguageDeCh      CallActionStartRecordingParamsTranscriptionLanguage = "de-CH"
	CallActionStartRecordingParamsTranscriptionLanguageDeDe      CallActionStartRecordingParamsTranscriptionLanguage = "de-DE"
	CallActionStartRecordingParamsTranscriptionLanguageElGr      CallActionStartRecordingParamsTranscriptionLanguage = "el-GR"
	CallActionStartRecordingParamsTranscriptionLanguageEnAu      CallActionStartRecordingParamsTranscriptionLanguage = "en-AU"
	CallActionStartRecordingParamsTranscriptionLanguageEnCa      CallActionStartRecordingParamsTranscriptionLanguage = "en-CA"
	CallActionStartRecordingParamsTranscriptionLanguageEnGB      CallActionStartRecordingParamsTranscriptionLanguage = "en-GB"
	CallActionStartRecordingParamsTranscriptionLanguageEnGh      CallActionStartRecordingParamsTranscriptionLanguage = "en-GH"
	CallActionStartRecordingParamsTranscriptionLanguageEnHk      CallActionStartRecordingParamsTranscriptionLanguage = "en-HK"
	CallActionStartRecordingParamsTranscriptionLanguageEnIe      CallActionStartRecordingParamsTranscriptionLanguage = "en-IE"
	CallActionStartRecordingParamsTranscriptionLanguageEnIn      CallActionStartRecordingParamsTranscriptionLanguage = "en-IN"
	CallActionStartRecordingParamsTranscriptionLanguageEnKe      CallActionStartRecordingParamsTranscriptionLanguage = "en-KE"
	CallActionStartRecordingParamsTranscriptionLanguageEnNg      CallActionStartRecordingParamsTranscriptionLanguage = "en-NG"
	CallActionStartRecordingParamsTranscriptionLanguageEnNz      CallActionStartRecordingParamsTranscriptionLanguage = "en-NZ"
	CallActionStartRecordingParamsTranscriptionLanguageEnPh      CallActionStartRecordingParamsTranscriptionLanguage = "en-PH"
	CallActionStartRecordingParamsTranscriptionLanguageEnPk      CallActionStartRecordingParamsTranscriptionLanguage = "en-PK"
	CallActionStartRecordingParamsTranscriptionLanguageEnSg      CallActionStartRecordingParamsTranscriptionLanguage = "en-SG"
	CallActionStartRecordingParamsTranscriptionLanguageEnTz      CallActionStartRecordingParamsTranscriptionLanguage = "en-TZ"
	CallActionStartRecordingParamsTranscriptionLanguageEnUs      CallActionStartRecordingParamsTranscriptionLanguage = "en-US"
	CallActionStartRecordingParamsTranscriptionLanguageEnZa      CallActionStartRecordingParamsTranscriptionLanguage = "en-ZA"
	CallActionStartRecordingParamsTranscriptionLanguageEsAr      CallActionStartRecordingParamsTranscriptionLanguage = "es-AR"
	CallActionStartRecordingParamsTranscriptionLanguageEsBo      CallActionStartRecordingParamsTranscriptionLanguage = "es-BO"
	CallActionStartRecordingParamsTranscriptionLanguageEsCl      CallActionStartRecordingParamsTranscriptionLanguage = "es-CL"
	CallActionStartRecordingParamsTranscriptionLanguageEsCo      CallActionStartRecordingParamsTranscriptionLanguage = "es-CO"
	CallActionStartRecordingParamsTranscriptionLanguageEsCr      CallActionStartRecordingParamsTranscriptionLanguage = "es-CR"
	CallActionStartRecordingParamsTranscriptionLanguageEsDo      CallActionStartRecordingParamsTranscriptionLanguage = "es-DO"
	CallActionStartRecordingParamsTranscriptionLanguageEsEc      CallActionStartRecordingParamsTranscriptionLanguage = "es-EC"
	CallActionStartRecordingParamsTranscriptionLanguageEsEs      CallActionStartRecordingParamsTranscriptionLanguage = "es-ES"
	CallActionStartRecordingParamsTranscriptionLanguageEsGt      CallActionStartRecordingParamsTranscriptionLanguage = "es-GT"
	CallActionStartRecordingParamsTranscriptionLanguageEsHn      CallActionStartRecordingParamsTranscriptionLanguage = "es-HN"
	CallActionStartRecordingParamsTranscriptionLanguageEsMx      CallActionStartRecordingParamsTranscriptionLanguage = "es-MX"
	CallActionStartRecordingParamsTranscriptionLanguageEsNi      CallActionStartRecordingParamsTranscriptionLanguage = "es-NI"
	CallActionStartRecordingParamsTranscriptionLanguageEsPa      CallActionStartRecordingParamsTranscriptionLanguage = "es-PA"
	CallActionStartRecordingParamsTranscriptionLanguageEsPe      CallActionStartRecordingParamsTranscriptionLanguage = "es-PE"
	CallActionStartRecordingParamsTranscriptionLanguageEsPr      CallActionStartRecordingParamsTranscriptionLanguage = "es-PR"
	CallActionStartRecordingParamsTranscriptionLanguageEsPy      CallActionStartRecordingParamsTranscriptionLanguage = "es-PY"
	CallActionStartRecordingParamsTranscriptionLanguageEsSv      CallActionStartRecordingParamsTranscriptionLanguage = "es-SV"
	CallActionStartRecordingParamsTranscriptionLanguageEsUs      CallActionStartRecordingParamsTranscriptionLanguage = "es-US"
	CallActionStartRecordingParamsTranscriptionLanguageEsUy      CallActionStartRecordingParamsTranscriptionLanguage = "es-UY"
	CallActionStartRecordingParamsTranscriptionLanguageEsVe      CallActionStartRecordingParamsTranscriptionLanguage = "es-VE"
	CallActionStartRecordingParamsTranscriptionLanguageEtEe      CallActionStartRecordingParamsTranscriptionLanguage = "et-EE"
	CallActionStartRecordingParamsTranscriptionLanguageEuEs      CallActionStartRecordingParamsTranscriptionLanguage = "eu-ES"
	CallActionStartRecordingParamsTranscriptionLanguageFaIr      CallActionStartRecordingParamsTranscriptionLanguage = "fa-IR"
	CallActionStartRecordingParamsTranscriptionLanguageFiFi      CallActionStartRecordingParamsTranscriptionLanguage = "fi-FI"
	CallActionStartRecordingParamsTranscriptionLanguageFilPh     CallActionStartRecordingParamsTranscriptionLanguage = "fil-PH"
	CallActionStartRecordingParamsTranscriptionLanguageFrBe      CallActionStartRecordingParamsTranscriptionLanguage = "fr-BE"
	CallActionStartRecordingParamsTranscriptionLanguageFrCa      CallActionStartRecordingParamsTranscriptionLanguage = "fr-CA"
	CallActionStartRecordingParamsTranscriptionLanguageFrCh      CallActionStartRecordingParamsTranscriptionLanguage = "fr-CH"
	CallActionStartRecordingParamsTranscriptionLanguageFrFr      CallActionStartRecordingParamsTranscriptionLanguage = "fr-FR"
	CallActionStartRecordingParamsTranscriptionLanguageGlEs      CallActionStartRecordingParamsTranscriptionLanguage = "gl-ES"
	CallActionStartRecordingParamsTranscriptionLanguageGuIn      CallActionStartRecordingParamsTranscriptionLanguage = "gu-IN"
	CallActionStartRecordingParamsTranscriptionLanguageHiIn      CallActionStartRecordingParamsTranscriptionLanguage = "hi-IN"
	CallActionStartRecordingParamsTranscriptionLanguageHrHr      CallActionStartRecordingParamsTranscriptionLanguage = "hr-HR"
	CallActionStartRecordingParamsTranscriptionLanguageHuHu      CallActionStartRecordingParamsTranscriptionLanguage = "hu-HU"
	CallActionStartRecordingParamsTranscriptionLanguageHyAm      CallActionStartRecordingParamsTranscriptionLanguage = "hy-AM"
	CallActionStartRecordingParamsTranscriptionLanguageIDID      CallActionStartRecordingParamsTranscriptionLanguage = "id-ID"
	CallActionStartRecordingParamsTranscriptionLanguageIsIs      CallActionStartRecordingParamsTranscriptionLanguage = "is-IS"
	CallActionStartRecordingParamsTranscriptionLanguageItCh      CallActionStartRecordingParamsTranscriptionLanguage = "it-CH"
	CallActionStartRecordingParamsTranscriptionLanguageItIt      CallActionStartRecordingParamsTranscriptionLanguage = "it-IT"
	CallActionStartRecordingParamsTranscriptionLanguageIwIl      CallActionStartRecordingParamsTranscriptionLanguage = "iw-IL"
	CallActionStartRecordingParamsTranscriptionLanguageJaJp      CallActionStartRecordingParamsTranscriptionLanguage = "ja-JP"
	CallActionStartRecordingParamsTranscriptionLanguageJvID      CallActionStartRecordingParamsTranscriptionLanguage = "jv-ID"
	CallActionStartRecordingParamsTranscriptionLanguageKaGe      CallActionStartRecordingParamsTranscriptionLanguage = "ka-GE"
	CallActionStartRecordingParamsTranscriptionLanguageKkKz      CallActionStartRecordingParamsTranscriptionLanguage = "kk-KZ"
	CallActionStartRecordingParamsTranscriptionLanguageKmKh      CallActionStartRecordingParamsTranscriptionLanguage = "km-KH"
	CallActionStartRecordingParamsTranscriptionLanguageKnIn      CallActionStartRecordingParamsTranscriptionLanguage = "kn-IN"
	CallActionStartRecordingParamsTranscriptionLanguageKoKr      CallActionStartRecordingParamsTranscriptionLanguage = "ko-KR"
	CallActionStartRecordingParamsTranscriptionLanguageLoLa      CallActionStartRecordingParamsTranscriptionLanguage = "lo-LA"
	CallActionStartRecordingParamsTranscriptionLanguageLtLt      CallActionStartRecordingParamsTranscriptionLanguage = "lt-LT"
	CallActionStartRecordingParamsTranscriptionLanguageLvLv      CallActionStartRecordingParamsTranscriptionLanguage = "lv-LV"
	CallActionStartRecordingParamsTranscriptionLanguageMkMk      CallActionStartRecordingParamsTranscriptionLanguage = "mk-MK"
	CallActionStartRecordingParamsTranscriptionLanguageMlIn      CallActionStartRecordingParamsTranscriptionLanguage = "ml-IN"
	CallActionStartRecordingParamsTranscriptionLanguageMnMn      CallActionStartRecordingParamsTranscriptionLanguage = "mn-MN"
	CallActionStartRecordingParamsTranscriptionLanguageMrIn      CallActionStartRecordingParamsTranscriptionLanguage = "mr-IN"
	CallActionStartRecordingParamsTranscriptionLanguageMsMy      CallActionStartRecordingParamsTranscriptionLanguage = "ms-MY"
	CallActionStartRecordingParamsTranscriptionLanguageMyMm      CallActionStartRecordingParamsTranscriptionLanguage = "my-MM"
	CallActionStartRecordingParamsTranscriptionLanguageNeNp      CallActionStartRecordingParamsTranscriptionLanguage = "ne-NP"
	CallActionStartRecordingParamsTranscriptionLanguageNlBe      CallActionStartRecordingParamsTranscriptionLanguage = "nl-BE"
	CallActionStartRecordingParamsTranscriptionLanguageNlNl      CallActionStartRecordingParamsTranscriptionLanguage = "nl-NL"
	CallActionStartRecordingParamsTranscriptionLanguageNoNo      CallActionStartRecordingParamsTranscriptionLanguage = "no-NO"
	CallActionStartRecordingParamsTranscriptionLanguagePaGuruIn  CallActionStartRecordingParamsTranscriptionLanguage = "pa-Guru-IN"
	CallActionStartRecordingParamsTranscriptionLanguagePlPl      CallActionStartRecordingParamsTranscriptionLanguage = "pl-PL"
	CallActionStartRecordingParamsTranscriptionLanguagePtBr      CallActionStartRecordingParamsTranscriptionLanguage = "pt-BR"
	CallActionStartRecordingParamsTranscriptionLanguagePtPt      CallActionStartRecordingParamsTranscriptionLanguage = "pt-PT"
	CallActionStartRecordingParamsTranscriptionLanguageRoRo      CallActionStartRecordingParamsTranscriptionLanguage = "ro-RO"
	CallActionStartRecordingParamsTranscriptionLanguageRuRu      CallActionStartRecordingParamsTranscriptionLanguage = "ru-RU"
	CallActionStartRecordingParamsTranscriptionLanguageRwRw      CallActionStartRecordingParamsTranscriptionLanguage = "rw-RW"
	CallActionStartRecordingParamsTranscriptionLanguageSiLk      CallActionStartRecordingParamsTranscriptionLanguage = "si-LK"
	CallActionStartRecordingParamsTranscriptionLanguageSkSk      CallActionStartRecordingParamsTranscriptionLanguage = "sk-SK"
	CallActionStartRecordingParamsTranscriptionLanguageSlSi      CallActionStartRecordingParamsTranscriptionLanguage = "sl-SI"
	CallActionStartRecordingParamsTranscriptionLanguageSqAl      CallActionStartRecordingParamsTranscriptionLanguage = "sq-AL"
	CallActionStartRecordingParamsTranscriptionLanguageSrRs      CallActionStartRecordingParamsTranscriptionLanguage = "sr-RS"
	CallActionStartRecordingParamsTranscriptionLanguageSSLatnZa  CallActionStartRecordingParamsTranscriptionLanguage = "ss-latn-za"
	CallActionStartRecordingParamsTranscriptionLanguageStZa      CallActionStartRecordingParamsTranscriptionLanguage = "st-ZA"
	CallActionStartRecordingParamsTranscriptionLanguageSuID      CallActionStartRecordingParamsTranscriptionLanguage = "su-ID"
	CallActionStartRecordingParamsTranscriptionLanguageSvSe      CallActionStartRecordingParamsTranscriptionLanguage = "sv-SE"
	CallActionStartRecordingParamsTranscriptionLanguageSwKe      CallActionStartRecordingParamsTranscriptionLanguage = "sw-KE"
	CallActionStartRecordingParamsTranscriptionLanguageSwTz      CallActionStartRecordingParamsTranscriptionLanguage = "sw-TZ"
	CallActionStartRecordingParamsTranscriptionLanguageTaIn      CallActionStartRecordingParamsTranscriptionLanguage = "ta-IN"
	CallActionStartRecordingParamsTranscriptionLanguageTaLk      CallActionStartRecordingParamsTranscriptionLanguage = "ta-LK"
	CallActionStartRecordingParamsTranscriptionLanguageTaMy      CallActionStartRecordingParamsTranscriptionLanguage = "ta-MY"
	CallActionStartRecordingParamsTranscriptionLanguageTaSg      CallActionStartRecordingParamsTranscriptionLanguage = "ta-SG"
	CallActionStartRecordingParamsTranscriptionLanguageTeIn      CallActionStartRecordingParamsTranscriptionLanguage = "te-IN"
	CallActionStartRecordingParamsTranscriptionLanguageThTh      CallActionStartRecordingParamsTranscriptionLanguage = "th-TH"
	CallActionStartRecordingParamsTranscriptionLanguageTnLatnZa  CallActionStartRecordingParamsTranscriptionLanguage = "tn-latn-za"
	CallActionStartRecordingParamsTranscriptionLanguageTrTr      CallActionStartRecordingParamsTranscriptionLanguage = "tr-TR"
	CallActionStartRecordingParamsTranscriptionLanguageTsZa      CallActionStartRecordingParamsTranscriptionLanguage = "ts-ZA"
	CallActionStartRecordingParamsTranscriptionLanguageUkUa      CallActionStartRecordingParamsTranscriptionLanguage = "uk-UA"
	CallActionStartRecordingParamsTranscriptionLanguageUrIn      CallActionStartRecordingParamsTranscriptionLanguage = "ur-IN"
	CallActionStartRecordingParamsTranscriptionLanguageUrPk      CallActionStartRecordingParamsTranscriptionLanguage = "ur-PK"
	CallActionStartRecordingParamsTranscriptionLanguageUzUz      CallActionStartRecordingParamsTranscriptionLanguage = "uz-UZ"
	CallActionStartRecordingParamsTranscriptionLanguageVeZa      CallActionStartRecordingParamsTranscriptionLanguage = "ve-ZA"
	CallActionStartRecordingParamsTranscriptionLanguageViVn      CallActionStartRecordingParamsTranscriptionLanguage = "vi-VN"
	CallActionStartRecordingParamsTranscriptionLanguageXhZa      CallActionStartRecordingParamsTranscriptionLanguage = "xh-ZA"
	CallActionStartRecordingParamsTranscriptionLanguageYueHantHk CallActionStartRecordingParamsTranscriptionLanguage = "yue-Hant-HK"
	CallActionStartRecordingParamsTranscriptionLanguageZh        CallActionStartRecordingParamsTranscriptionLanguage = "zh"
	CallActionStartRecordingParamsTranscriptionLanguageZhTw      CallActionStartRecordingParamsTranscriptionLanguage = "zh-TW"
	CallActionStartRecordingParamsTranscriptionLanguageZuZa      CallActionStartRecordingParamsTranscriptionLanguage = "zu-ZA"
)

// When set to `trim-silence`, silence will be removed from the beginning and end
// of the recording.
type CallActionStartRecordingParamsTrim string

const (
	CallActionStartRecordingParamsTrimTrimSilence CallActionStartRecordingParamsTrim = "trim-silence"
)

type CallActionStartSiprecParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Name of configured SIPREC connector to be used.
	ConnectorName param.Opt[string] `json:"connector_name,omitzero"`
	// Sets `Session-Expires` header to the INVITE. A reinvite is sent every half the
	// value set. Usefull for session keep alive. Minimum value is 90, set to 0 to
	// disable.
	SessionTimeoutSecs param.Opt[int64] `json:"session_timeout_secs,omitzero"`
	// When set, custom parameters will be added as metadata
	// (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip
	// headers.
	//
	// Any of true, false.
	IncludeMetadataCustomHeaders bool `json:"include_metadata_custom_headers,omitzero"`
	// Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set
	// you need to configure SRS port in your connector to 5061.
	//
	// Any of true, false.
	Secure bool `json:"secure,omitzero"`
	// Specifies SIP transport protocol.
	//
	// Any of "udp", "tcp", "tls".
	SipTransport CallActionStartSiprecParamsSipTransport `json:"sip_transport,omitzero"`
	// Specifies which track should be sent on siprec session.
	//
	// Any of "inbound_track", "outbound_track", "both_tracks".
	SiprecTrack CallActionStartSiprecParamsSiprecTrack `json:"siprec_track,omitzero"`
	paramObj
}

func (r CallActionStartSiprecParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartSiprecParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartSiprecParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies SIP transport protocol.
type CallActionStartSiprecParamsSipTransport string

const (
	CallActionStartSiprecParamsSipTransportUdp CallActionStartSiprecParamsSipTransport = "udp"
	CallActionStartSiprecParamsSipTransportTcp CallActionStartSiprecParamsSipTransport = "tcp"
	CallActionStartSiprecParamsSipTransportTls CallActionStartSiprecParamsSipTransport = "tls"
)

// Specifies which track should be sent on siprec session.
type CallActionStartSiprecParamsSiprecTrack string

const (
	CallActionStartSiprecParamsSiprecTrackInboundTrack  CallActionStartSiprecParamsSiprecTrack = "inbound_track"
	CallActionStartSiprecParamsSiprecTrackOutboundTrack CallActionStartSiprecParamsSiprecTrack = "outbound_track"
	CallActionStartSiprecParamsSiprecTrackBothTracks    CallActionStartSiprecParamsSiprecTrack = "both_tracks"
)

type CallActionStartStreamingParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Enables Dialogflow for the current call. The default value is false.
	EnableDialogflow param.Opt[bool] `json:"enable_dialogflow,omitzero"`
	// The destination WebSocket address where the stream is going to be delivered.
	StreamURL        param.Opt[string]     `json:"stream_url,omitzero"`
	DialogflowConfig DialogflowConfigParam `json:"dialogflow_config,omitzero"`
	// Indicates codec for bidirectional streaming RTP payloads. Used only with
	// stream_bidirectional_mode=rtp. Case sensitive.
	//
	// Any of "PCMU", "PCMA", "G722", "OPUS", "AMR-WB".
	StreamBidirectionalCodec StreamBidirectionalCodec `json:"stream_bidirectional_codec,omitzero"`
	// Configures method of bidirectional streaming (mp3, rtp).
	//
	// Any of "mp3", "rtp".
	StreamBidirectionalMode StreamBidirectionalMode `json:"stream_bidirectional_mode,omitzero"`
	// Specifies which call legs should receive the bidirectional stream audio.
	//
	// Any of "both", "self", "opposite".
	StreamBidirectionalTargetLegs StreamBidirectionalTargetLegs `json:"stream_bidirectional_target_legs,omitzero"`
	// Specifies the codec to be used for the streamed audio. When set to 'default' or
	// when transcoding is not possible, the codec from the call will be used.
	// Currently, transcoding is only supported between PCMU and PCMA codecs.
	//
	// Any of "PCMA", "PCMU", "default".
	StreamCodec StreamCodec `json:"stream_codec,omitzero"`
	// Specifies which track should be streamed.
	//
	// Any of "inbound_track", "outbound_track", "both_tracks".
	StreamTrack CallActionStartStreamingParamsStreamTrack `json:"stream_track,omitzero"`
	paramObj
}

func (r CallActionStartStreamingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStartStreamingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStartStreamingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies which track should be streamed.
type CallActionStartStreamingParamsStreamTrack string

const (
	CallActionStartStreamingParamsStreamTrackInboundTrack  CallActionStartStreamingParamsStreamTrack = "inbound_track"
	CallActionStartStreamingParamsStreamTrackOutboundTrack CallActionStartStreamingParamsStreamTrack = "outbound_track"
	CallActionStartStreamingParamsStreamTrackBothTracks    CallActionStartStreamingParamsStreamTrack = "both_tracks"
)

type CallActionStartTranscriptionParams struct {
	TranscriptionStartRequest TranscriptionStartRequestParam
	paramObj
}

func (r CallActionStartTranscriptionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TranscriptionStartRequest)
}
func (r *CallActionStartTranscriptionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TranscriptionStartRequest)
}

type CallActionStopAIAssistantParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionStopAIAssistantParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopAIAssistantParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopAIAssistantParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopForkingParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Optionally specify a `stream_type`. This should match the `stream_type` that was
	// used in `fork_start` command to properly stop the fork.
	//
	// Any of "raw", "decrypted".
	StreamType CallActionStopForkingParamsStreamType `json:"stream_type,omitzero"`
	paramObj
}

func (r CallActionStopForkingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopForkingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopForkingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optionally specify a `stream_type`. This should match the `stream_type` that was
// used in `fork_start` command to properly stop the fork.
type CallActionStopForkingParamsStreamType string

const (
	CallActionStopForkingParamsStreamTypeRaw       CallActionStopForkingParamsStreamType = "raw"
	CallActionStopForkingParamsStreamTypeDecrypted CallActionStopForkingParamsStreamType = "decrypted"
)

type CallActionStopGatherParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionStopGatherParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopGatherParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopGatherParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopNoiseSuppressionParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionStopNoiseSuppressionParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopNoiseSuppressionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopNoiseSuppressionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopPlaybackParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// When enabled, it stops the audio being played in the overlay queue.
	Overlay param.Opt[bool] `json:"overlay,omitzero"`
	// Use `current` to stop the current audio being played. Use `all` to stop the
	// current audio file being played and clear all audio files from the queue.
	Stop param.Opt[string] `json:"stop,omitzero"`
	paramObj
}

func (r CallActionStopPlaybackParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopPlaybackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopPlaybackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopRecordingParams struct {
	StopRecordingRequest StopRecordingRequestParam
	paramObj
}

func (r CallActionStopRecordingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.StopRecordingRequest)
}
func (r *CallActionStopRecordingParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.StopRecordingRequest)
}

type CallActionStopSiprecParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionStopSiprecParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopSiprecParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopSiprecParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopStreamingParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Identifies the stream. If the `stream_id` is not provided the command stops all
	// streams associated with a given `call_control_id`.
	StreamID param.Opt[string] `json:"stream_id,omitzero" format:"uuid"`
	paramObj
}

func (r CallActionStopStreamingParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopStreamingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopStreamingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionStopTranscriptionParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	paramObj
}

func (r CallActionStopTranscriptionParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionStopTranscriptionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionStopTranscriptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallActionSwitchSupervisorRoleParams struct {
	// The supervisor role to switch to. 'barge' allows speaking to both parties,
	// 'whisper' allows speaking to caller only, 'monitor' allows listening only.
	//
	// Any of "barge", "whisper", "monitor".
	Role CallActionSwitchSupervisorRoleParamsRole `json:"role,omitzero,required"`
	paramObj
}

func (r CallActionSwitchSupervisorRoleParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionSwitchSupervisorRoleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionSwitchSupervisorRoleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The supervisor role to switch to. 'barge' allows speaking to both parties,
// 'whisper' allows speaking to caller only, 'monitor' allows listening only.
type CallActionSwitchSupervisorRoleParamsRole string

const (
	CallActionSwitchSupervisorRoleParamsRoleBarge   CallActionSwitchSupervisorRoleParamsRole = "barge"
	CallActionSwitchSupervisorRoleParamsRoleWhisper CallActionSwitchSupervisorRoleParamsRole = "whisper"
	CallActionSwitchSupervisorRoleParamsRoleMonitor CallActionSwitchSupervisorRoleParamsRole = "monitor"
)

type CallActionTransferParams struct {
	// The DID or SIP URI to dial out to.
	To string `json:"to,required"`
	// The URL of a file to be played back when the transfer destination answers before
	// bridging the call. The URL can point to either a WAV or MP3 file. media_name and
	// audio_url cannot be used together in one request.
	AudioURL param.Opt[string] `json:"audio_url,omitzero"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// If set to false, early media will not be passed to the originating leg.
	EarlyMedia param.Opt[bool] `json:"early_media,omitzero"`
	// The `from` number to be used as the caller id presented to the destination (`to`
	// number). The number should be in +E164 format. This attribute will default to
	// the `to` number of the original call if omitted.
	From param.Opt[string] `json:"from,omitzero"`
	// The `from_display_name` string to be used as the caller id name (SIP From
	// Display Name) presented to the destination (`to` number). The string should have
	// a maximum of 128 characters, containing only letters, numbers, spaces, and
	// -\_~!.+ special characters. If ommited, the display name will be the same as the
	// number in the `from` field.
	FromDisplayName param.Opt[string] `json:"from_display_name,omitzero"`
	// The media_name of a file to be played back when the transfer destination answers
	// before bridging the call. The media_name must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. The file must
	// either be a WAV or MP3 file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up
	// or is transferred). If supplied with the value `self`, the current leg will be
	// parked after unbridge. If not set, the default behavior is to hang up the leg.
	ParkAfterUnbridge param.Opt[string] `json:"park_after_unbridge,omitzero"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword param.Opt[string] `json:"sip_auth_password,omitzero"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername param.Opt[string] `json:"sip_auth_username,omitzero"`
	// Use this field to add state to every subsequent webhook for the new leg. It must
	// be a valid Base-64 encoded string.
	TargetLegClientState param.Opt[string] `json:"target_leg_client_state,omitzero"`
	// Sets the maximum duration of a Call Control Leg in seconds. If the time limit is
	// reached, the call will hangup and a `call.hangup` webhook with a `hangup_cause`
	// of `time_limit` will be sent. For example, by setting a time limit of 120
	// seconds, a Call Leg will be automatically terminated two minutes after being
	// answered. The default time limit is 14400 seconds or 4 hours and this is also
	// the maximum allowed call length.
	TimeLimitSecs param.Opt[int64] `json:"time_limit_secs,omitzero"`
	// The number of seconds that Telnyx will wait for the call to be answered by the
	// destination to which it is being transferred. If the timeout is reached before
	// an answer is received, the call will hangup and a `call.hangup` webhook with a
	// `hangup_cause` of `timeout` will be sent. Minimum value is 5 seconds. Maximum
	// value is 600 seconds.
	TimeoutSecs param.Opt[int64] `json:"timeout_secs,omitzero"`
	// Use this field to override the URL for which Telnyx will send subsequent
	// webhooks to for this call.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Enables Answering Machine Detection. When a call is answered, Telnyx runs
	// real-time detection to determine if it was picked up by a human or a machine and
	// sends an `call.machine.detection.ended` webhook with the analysis result. If
	// 'greeting_end' or 'detect_words' is used and a 'machine' is detected, you will
	// receive another 'call.machine.greeting.ended' webhook when the answering machine
	// greeting ends with a beep or silence. If `detect_beep` is used, you will only
	// receive 'call.machine.greeting.ended' if a beep is detected.
	//
	// Any of "premium", "detect", "detect_beep", "detect_words", "greeting_end",
	// "disabled".
	AnsweringMachineDetection CallActionTransferParamsAnsweringMachineDetection `json:"answering_machine_detection,omitzero"`
	// Optional configuration parameters to modify 'answering_machine_detection'
	// performance.
	AnsweringMachineDetectionConfig CallActionTransferParamsAnsweringMachineDetectionConfig `json:"answering_machine_detection_config,omitzero"`
	// Custom headers to be added to the SIP INVITE.
	CustomHeaders []CustomSipHeaderParam `json:"custom_headers,omitzero"`
	// Defines whether media should be encrypted on the new call leg.
	//
	// Any of "disabled", "SRTP", "DTLS".
	MediaEncryption CallActionTransferParamsMediaEncryption `json:"media_encryption,omitzero"`
	// When enabled, DTMF tones are not passed to the call participant. The webhooks
	// containing the DTMF information will be sent.
	//
	// Any of "none", "both", "self", "opposite".
	MuteDtmf CallActionTransferParamsMuteDtmf `json:"mute_dtmf,omitzero"`
	// SIP headers to be added to the SIP INVITE. Currently only User-to-User header is
	// supported.
	SipHeaders []SipHeaderParam `json:"sip_headers,omitzero"`
	// Defines SIP transport protocol to be used on the call.
	//
	// Any of "UDP", "TCP", "TLS".
	SipTransportProtocol CallActionTransferParamsSipTransportProtocol `json:"sip_transport_protocol,omitzero"`
	// Use this field to modify sound effects, for example adjust the pitch.
	SoundModifications SoundModificationsParam `json:"sound_modifications,omitzero"`
	// HTTP request type used for `webhook_url`.
	//
	// Any of "POST", "GET".
	WebhookURLMethod CallActionTransferParamsWebhookURLMethod `json:"webhook_url_method,omitzero"`
	paramObj
}

func (r CallActionTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables Answering Machine Detection. When a call is answered, Telnyx runs
// real-time detection to determine if it was picked up by a human or a machine and
// sends an `call.machine.detection.ended` webhook with the analysis result. If
// 'greeting_end' or 'detect_words' is used and a 'machine' is detected, you will
// receive another 'call.machine.greeting.ended' webhook when the answering machine
// greeting ends with a beep or silence. If `detect_beep` is used, you will only
// receive 'call.machine.greeting.ended' if a beep is detected.
type CallActionTransferParamsAnsweringMachineDetection string

const (
	CallActionTransferParamsAnsweringMachineDetectionPremium     CallActionTransferParamsAnsweringMachineDetection = "premium"
	CallActionTransferParamsAnsweringMachineDetectionDetect      CallActionTransferParamsAnsweringMachineDetection = "detect"
	CallActionTransferParamsAnsweringMachineDetectionDetectBeep  CallActionTransferParamsAnsweringMachineDetection = "detect_beep"
	CallActionTransferParamsAnsweringMachineDetectionDetectWords CallActionTransferParamsAnsweringMachineDetection = "detect_words"
	CallActionTransferParamsAnsweringMachineDetectionGreetingEnd CallActionTransferParamsAnsweringMachineDetection = "greeting_end"
	CallActionTransferParamsAnsweringMachineDetectionDisabled    CallActionTransferParamsAnsweringMachineDetection = "disabled"
)

// Optional configuration parameters to modify 'answering_machine_detection'
// performance.
type CallActionTransferParamsAnsweringMachineDetectionConfig struct {
	// Silence duration threshold after a greeting message or voice for it be
	// considered human.
	AfterGreetingSilenceMillis param.Opt[int64] `json:"after_greeting_silence_millis,omitzero"`
	// Maximum threshold for silence between words.
	BetweenWordsSilenceMillis param.Opt[int64] `json:"between_words_silence_millis,omitzero"`
	// Maximum threshold of a human greeting. If greeting longer than this value,
	// considered machine.
	GreetingDurationMillis param.Opt[int64] `json:"greeting_duration_millis,omitzero"`
	// If machine already detected, maximum threshold for silence between words. If
	// exceeded, the greeting is considered ended.
	GreetingSilenceDurationMillis param.Opt[int64] `json:"greeting_silence_duration_millis,omitzero"`
	// If machine already detected, maximum timeout threshold to determine the end of
	// the machine greeting.
	GreetingTotalAnalysisTimeMillis param.Opt[int64] `json:"greeting_total_analysis_time_millis,omitzero"`
	// If initial silence duration is greater than this value, consider it a machine.
	InitialSilenceMillis param.Opt[int64] `json:"initial_silence_millis,omitzero"`
	// If number of detected words is greater than this value, consder it a machine.
	MaximumNumberOfWords param.Opt[int64] `json:"maximum_number_of_words,omitzero"`
	// If a single word lasts longer than this threshold, consider it a machine.
	MaximumWordLengthMillis param.Opt[int64] `json:"maximum_word_length_millis,omitzero"`
	// Minimum noise threshold for any analysis.
	SilenceThreshold param.Opt[int64] `json:"silence_threshold,omitzero"`
	// Maximum timeout threshold for overall detection.
	TotalAnalysisTimeMillis param.Opt[int64] `json:"total_analysis_time_millis,omitzero"`
	paramObj
}

func (r CallActionTransferParamsAnsweringMachineDetectionConfig) MarshalJSON() (data []byte, err error) {
	type shadow CallActionTransferParamsAnsweringMachineDetectionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionTransferParamsAnsweringMachineDetectionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether media should be encrypted on the new call leg.
type CallActionTransferParamsMediaEncryption string

const (
	CallActionTransferParamsMediaEncryptionDisabled CallActionTransferParamsMediaEncryption = "disabled"
	CallActionTransferParamsMediaEncryptionSrtp     CallActionTransferParamsMediaEncryption = "SRTP"
	CallActionTransferParamsMediaEncryptionDtls     CallActionTransferParamsMediaEncryption = "DTLS"
)

// When enabled, DTMF tones are not passed to the call participant. The webhooks
// containing the DTMF information will be sent.
type CallActionTransferParamsMuteDtmf string

const (
	CallActionTransferParamsMuteDtmfNone     CallActionTransferParamsMuteDtmf = "none"
	CallActionTransferParamsMuteDtmfBoth     CallActionTransferParamsMuteDtmf = "both"
	CallActionTransferParamsMuteDtmfSelf     CallActionTransferParamsMuteDtmf = "self"
	CallActionTransferParamsMuteDtmfOpposite CallActionTransferParamsMuteDtmf = "opposite"
)

// Defines SIP transport protocol to be used on the call.
type CallActionTransferParamsSipTransportProtocol string

const (
	CallActionTransferParamsSipTransportProtocolUdp CallActionTransferParamsSipTransportProtocol = "UDP"
	CallActionTransferParamsSipTransportProtocolTcp CallActionTransferParamsSipTransportProtocol = "TCP"
	CallActionTransferParamsSipTransportProtocolTls CallActionTransferParamsSipTransportProtocol = "TLS"
)

// HTTP request type used for `webhook_url`.
type CallActionTransferParamsWebhookURLMethod string

const (
	CallActionTransferParamsWebhookURLMethodPost CallActionTransferParamsWebhookURLMethod = "POST"
	CallActionTransferParamsWebhookURLMethodGet  CallActionTransferParamsWebhookURLMethod = "GET"
)

type CallActionUpdateClientStateParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState string `json:"client_state,required"`
	paramObj
}

func (r CallActionUpdateClientStateParams) MarshalJSON() (data []byte, err error) {
	type shadow CallActionUpdateClientStateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallActionUpdateClientStateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
