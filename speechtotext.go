// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
)

// SpeechToTextService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeechToTextService] method instead.
type SpeechToTextService struct {
	Options []option.RequestOption
}

// NewSpeechToTextService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeechToTextService(opts ...option.RequestOption) (r SpeechToTextService) {
	r = SpeechToTextService{}
	r.Options = opts
	return
}

// Transcribe audio streams to text over WebSocket.
func (r *SpeechToTextService) Transcribe(ctx context.Context, query SpeechToTextTranscribeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "speech-to-text/transcription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type SpeechToTextTranscribeParams struct {
	// The format of input audio stream.
	//
	// Any of "mp3", "wav".
	InputFormat SpeechToTextTranscribeParamsInputFormat `query:"input_format,omitzero" api:"required" json:"-"`
	// The transcription engine to use for processing the audio stream.
	//
	// Any of "Azure", "Deepgram", "Google", "Telnyx".
	TranscriptionEngine SpeechToTextTranscribeParamsTranscriptionEngine `query:"transcription_engine,omitzero" api:"required" json:"-"`
	// Whether to receive interim transcription results.
	InterimResults param.Opt[bool] `query:"interim_results,omitzero" json:"-"`
	// The language spoken in the audio stream.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// The specific model to use within the selected transcription engine.
	//
	// Any of "fast", "deepgram/nova-2", "deepgram/nova-3", "latest_long",
	// "latest_short", "command_and_search", "phone_call", "video", "default",
	// "medical_conversation", "medical_dictation", "openai/whisper-tiny",
	// "openai/whisper-large-v3-turbo".
	Model SpeechToTextTranscribeParamsModel `query:"model,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpeechToTextTranscribeParams]'s query parameters as
// `url.Values`.
func (r SpeechToTextTranscribeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The format of input audio stream.
type SpeechToTextTranscribeParamsInputFormat string

const (
	SpeechToTextTranscribeParamsInputFormatMP3 SpeechToTextTranscribeParamsInputFormat = "mp3"
	SpeechToTextTranscribeParamsInputFormatWav SpeechToTextTranscribeParamsInputFormat = "wav"
)

// The transcription engine to use for processing the audio stream.
type SpeechToTextTranscribeParamsTranscriptionEngine string

const (
	SpeechToTextTranscribeParamsTranscriptionEngineAzure    SpeechToTextTranscribeParamsTranscriptionEngine = "Azure"
	SpeechToTextTranscribeParamsTranscriptionEngineDeepgram SpeechToTextTranscribeParamsTranscriptionEngine = "Deepgram"
	SpeechToTextTranscribeParamsTranscriptionEngineGoogle   SpeechToTextTranscribeParamsTranscriptionEngine = "Google"
	SpeechToTextTranscribeParamsTranscriptionEngineTelnyx   SpeechToTextTranscribeParamsTranscriptionEngine = "Telnyx"
)

// The specific model to use within the selected transcription engine.
type SpeechToTextTranscribeParamsModel string

const (
	SpeechToTextTranscribeParamsModelFast                      SpeechToTextTranscribeParamsModel = "fast"
	SpeechToTextTranscribeParamsModelDeepgramNova2             SpeechToTextTranscribeParamsModel = "deepgram/nova-2"
	SpeechToTextTranscribeParamsModelDeepgramNova3             SpeechToTextTranscribeParamsModel = "deepgram/nova-3"
	SpeechToTextTranscribeParamsModelLatestLong                SpeechToTextTranscribeParamsModel = "latest_long"
	SpeechToTextTranscribeParamsModelLatestShort               SpeechToTextTranscribeParamsModel = "latest_short"
	SpeechToTextTranscribeParamsModelCommandAndSearch          SpeechToTextTranscribeParamsModel = "command_and_search"
	SpeechToTextTranscribeParamsModelPhoneCall                 SpeechToTextTranscribeParamsModel = "phone_call"
	SpeechToTextTranscribeParamsModelVideo                     SpeechToTextTranscribeParamsModel = "video"
	SpeechToTextTranscribeParamsModelDefault                   SpeechToTextTranscribeParamsModel = "default"
	SpeechToTextTranscribeParamsModelMedicalConversation       SpeechToTextTranscribeParamsModel = "medical_conversation"
	SpeechToTextTranscribeParamsModelMedicalDictation          SpeechToTextTranscribeParamsModel = "medical_dictation"
	SpeechToTextTranscribeParamsModelOpenAIWhisperTiny         SpeechToTextTranscribeParamsModel = "openai/whisper-tiny"
	SpeechToTextTranscribeParamsModelOpenAIWhisperLargeV3Turbo SpeechToTextTranscribeParamsModel = "openai/whisper-large-v3-turbo"
)
