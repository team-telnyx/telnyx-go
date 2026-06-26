// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
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

// Retrieve the canonical list of supported speech-to-text providers, models,
// accepted language codes, and the service types each model supports.
//
// Service types:
//
//   - `streaming` — standalone WebSocket transcription via
//     `/speech-to-text/transcription`.
//   - `file_based` — file-based transcription via `/ai/audio/transcriptions`.
//   - `in_call` — live call transcription via Call Control `transcription_start`.
//   - `ai_assistant` — STT configured on a Call Control AI Assistant via
//     voice-assistant `TranscriptionConfig` (covers both live-streaming and
//     non-streaming/batch models).
//
// Use this endpoint to discover which (provider, model) combinations are available
// for the surface you need, and which language codes each accepts. `auto` in a
// `languages` array indicates the provider performs language detection.
func (r *SpeechToTextService) ListProviders(ctx context.Context, query SpeechToTextListProvidersParams, opts ...option.RequestOption) (res *SpeechToTextListProvidersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.telnyx.com/v2/")}, opts...)
	path := "speech-to-text/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Open a WebSocket connection to stream audio and receive transcriptions in
// real-time. Authentication is provided via the standard
// `Authorization: Bearer <API_KEY>` header.
//
// Supported engines: `Azure`, `Deepgram`, `Google`, `Telnyx`, `xAI`,
// `Speechmatics`, `Soniox`.
//
// **Connection flow:**
//
//  1. Open WebSocket with query parameters specifying engine, input format, and
//     language.
//  2. Send binary audio frames (mp3/wav format).
//  3. Receive JSON transcript frames with `transcript`, `is_final`, and
//     `confidence` fields.
//  4. Close connection when done.
func (r *SpeechToTextService) GetTranscription(ctx context.Context, query SpeechToTextGetTranscriptionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("wss://api.telnyx.com/v2/")}, opts...)
	path := "speech-to-text/transcription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Service surface a model is available on. `ai_assistant` is the STT surface
// configured via Call Control voice-assistant transcription; it covers both
// live-streaming and non-streaming/batch models (matching the
// `TranscriptionConfig.model` enum on `call-control` voice assistants).
type SttServiceType string

const (
	SttServiceTypeStreaming   SttServiceType = "streaming"
	SttServiceTypeFileBased   SttServiceType = "file_based"
	SttServiceTypeInCall      SttServiceType = "in_call"
	SttServiceTypeAIAssistant SttServiceType = "ai_assistant"
)

// List of supported STT providers and models.
type SpeechToTextListProvidersResponse struct {
	Data []SpeechToTextListProvidersResponseData `json:"data" api:"required"`
	Meta SpeechToTextListProvidersResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextListProvidersResponse) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextListProvidersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A (provider, model) tuple along with the service surfaces it supports. Each
// entry in `service_types` describes one surface and the languages accepted on it.
type SpeechToTextListProvidersResponseData struct {
	// Provider-scoped model name.
	Model string `json:"model" api:"required"`
	// STT provider name.
	Provider string `json:"provider" api:"required"`
	// Service surfaces this (provider, model) supports. When the request filters by
	// `service_type`, only the matching nested entry is returned for each matching
	// model.
	ServiceTypes []SpeechToTextListProvidersResponseDataServiceType `json:"service_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model        respjson.Field
		Provider     respjson.Field
		ServiceTypes respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextListProvidersResponseData) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextListProvidersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A supported service surface for a given (provider, model), along with the
// language codes accepted on that surface. Language support can differ per surface
// — for example, a model may accept a narrower language set for streaming than for
// file transcription.
type SpeechToTextListProvidersResponseDataServiceType struct {
	// Languages accepted on this service surface, in the provider's native code
	// format. `auto` indicates the provider performs language detection.
	Languages []string `json:"languages" api:"required"`
	// Service surface a model is available on. `ai_assistant` is the STT surface
	// configured via Call Control voice-assistant transcription; it covers both
	// live-streaming and non-streaming/batch models (matching the
	// `TranscriptionConfig.model` enum on `call-control` voice assistants).
	//
	// Any of "streaming", "file_based", "in_call", "ai_assistant".
	Type SttServiceType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Languages   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextListProvidersResponseDataServiceType) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextListProvidersResponseDataServiceType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextListProvidersResponseMeta struct {
	// Total number of entries returned.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextListProvidersResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextListProvidersResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscribeClientEvent = string

type TranscribeClientEventParam = string

// TranscribeServerEventUnion contains all possible properties and values from
// [TranscribeServerEventTranscript], [TranscribeServerEventError].
//
// Use the [TranscribeServerEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TranscribeServerEventUnion struct {
	// This field is from variant [TranscribeServerEventTranscript].
	Transcript string `json:"transcript"`
	// Any of "transcript", "error".
	Type string `json:"type"`
	// This field is from variant [TranscribeServerEventTranscript].
	Confidence float64 `json:"confidence"`
	// This field is from variant [TranscribeServerEventTranscript].
	IsFinal bool `json:"is_final"`
	// This field is from variant [TranscribeServerEventError].
	Error string `json:"error"`
	JSON  struct {
		Transcript respjson.Field
		Type       respjson.Field
		Confidence respjson.Field
		IsFinal    respjson.Field
		Error      respjson.Field
		raw        string
	} `json:"-"`
}

// anyTranscribeServerEvent is implemented by each variant of
// [TranscribeServerEventUnion] to add type safety for the return type of
// [TranscribeServerEventUnion.AsAny]
type anyTranscribeServerEvent interface {
	implTranscribeServerEventUnion()
}

func (TranscribeServerEventTranscript) implTranscribeServerEventUnion() {}
func (TranscribeServerEventError) implTranscribeServerEventUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TranscribeServerEventUnion.AsAny().(type) {
//	case telnyx.TranscribeServerEventTranscript:
//	case telnyx.TranscribeServerEventError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TranscribeServerEventUnion) AsAny() anyTranscribeServerEvent {
	switch u.Type {
	case "transcript":
		return u.AsTranscript()
	case "error":
		return u.AsError()
	}
	return nil
}

func (u TranscribeServerEventUnion) AsTranscript() (v TranscribeServerEventTranscript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TranscribeServerEventUnion) AsError() (v TranscribeServerEventError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TranscribeServerEventUnion) RawJSON() string { return u.JSON.raw }

func (r *TranscribeServerEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-to-client frame containing a transcription result.
type TranscribeServerEventTranscript struct {
	// The transcribed text from the audio.
	Transcript string `json:"transcript" api:"required"`
	// Frame type identifier.
	Type constant.Transcript `json:"type" default:"transcript"`
	// Confidence score of the transcription, ranging from 0 to 1.
	Confidence float64 `json:"confidence"`
	// Whether this is a final transcription result. When `false`, this is an interim
	// result that may be refined.
	IsFinal bool `json:"is_final"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transcript  respjson.Field
		Type        respjson.Field
		Confidence  respjson.Field
		IsFinal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscribeServerEventTranscript) RawJSON() string { return r.JSON.raw }
func (r *TranscribeServerEventTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-to-client frame indicating an error during transcription. The connection
// may be closed shortly after.
type TranscribeServerEventError struct {
	// Error message describing what went wrong.
	Error string `json:"error" api:"required"`
	// Frame type identifier.
	Type constant.Error `json:"type" default:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscribeServerEventError) RawJSON() string { return r.JSON.raw }
func (r *TranscribeServerEventError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextListProvidersParams struct {
	// Filter to entries for a specific STT provider. The enum mirrors the providers
	// advertised across the speech-to-text spec (including `google` and `telnyx`,
	// which are accepted as WebSocket transcription engines). A provider that has no
	// models currently registered for any service type will return an empty `data`
	// array rather than an error.
	//
	// Any of "deepgram", "speechmatics", "assemblyai", "xai", "soniox", "azure",
	// "openai", "google", "telnyx".
	Provider SpeechToTextListProvidersParamsProvider `query:"provider,omitzero" json:"-"`
	// Filter to entries that support the given service type. For backward
	// compatibility with the values that briefly shipped before the product-aligned
	// rename, the legacy aliases `file_transcription`, `in_call_transcription`, and
	// `ai_assistant_transcription` are silently accepted and normalized to
	// `file_based`, `in_call`, and `ai_assistant` respectively. The response always
	// emits the canonical (post-rename) values.
	//
	// Any of "streaming", "file_based", "in_call", "ai_assistant".
	ServiceType SttServiceType `query:"service_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpeechToTextListProvidersParams]'s query parameters as
// `url.Values`.
func (r SpeechToTextListProvidersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to entries for a specific STT provider. The enum mirrors the providers
// advertised across the speech-to-text spec (including `google` and `telnyx`,
// which are accepted as WebSocket transcription engines). A provider that has no
// models currently registered for any service type will return an empty `data`
// array rather than an error.
type SpeechToTextListProvidersParamsProvider string

const (
	SpeechToTextListProvidersParamsProviderDeepgram     SpeechToTextListProvidersParamsProvider = "deepgram"
	SpeechToTextListProvidersParamsProviderSpeechmatics SpeechToTextListProvidersParamsProvider = "speechmatics"
	SpeechToTextListProvidersParamsProviderAssemblyai   SpeechToTextListProvidersParamsProvider = "assemblyai"
	SpeechToTextListProvidersParamsProviderXai          SpeechToTextListProvidersParamsProvider = "xai"
	SpeechToTextListProvidersParamsProviderSoniox       SpeechToTextListProvidersParamsProvider = "soniox"
	SpeechToTextListProvidersParamsProviderAzure        SpeechToTextListProvidersParamsProvider = "azure"
	SpeechToTextListProvidersParamsProviderOpenAI       SpeechToTextListProvidersParamsProvider = "openai"
	SpeechToTextListProvidersParamsProviderGoogle       SpeechToTextListProvidersParamsProvider = "google"
	SpeechToTextListProvidersParamsProviderTelnyx       SpeechToTextListProvidersParamsProvider = "telnyx"
)

type SpeechToTextGetTranscriptionParams struct {
	// The format of input audio stream.
	//
	// Any of "mp3", "wav".
	InputFormat SpeechToTextGetTranscriptionParamsInputFormat `query:"input_format,omitzero" api:"required" json:"-"`
	// The transcription engine to use for processing the audio stream.
	//
	// Any of "Azure", "Deepgram", "Google", "Telnyx", "xAI", "Speechmatics", "Soniox".
	TranscriptionEngine SpeechToTextGetTranscriptionParamsTranscriptionEngine `query:"transcription_engine,omitzero" api:"required" json:"-"`
	// Silence duration (in milliseconds) that triggers end-of-speech detection. When
	// set, the engine uses this value to determine when a speaker has stopped talking.
	// Supported by `xAI`, `Deepgram`, `Google`, `Speechmatics`, and `Soniox`. `Soniox`
	// accepts values between 500 and 3000. Other engines may not support this
	// parameter.
	Endpointing param.Opt[int64] `query:"endpointing,omitzero" json:"-"`
	// Whether to receive interim transcription results.
	InterimResults param.Opt[bool] `query:"interim_results,omitzero" json:"-"`
	// A key term to boost in the transcription. The engine will be more likely to
	// recognize this term. Can be specified multiple times for multiple terms.
	Keyterm param.Opt[string] `query:"keyterm,omitzero" json:"-"`
	// Comma-separated list of keywords to boost in the transcription. The engine will
	// prioritize recognition of these words.
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// The language spoken in the audio stream.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Enable redaction of sensitive information (e.g., PCI data, SSN) from
	// transcription results. Supported values depend on the transcription engine.
	Redact param.Opt[string] `query:"redact,omitzero" json:"-"`
	// The specific model to use within the selected transcription engine.
	//
	// Any of "fast", "deepgram/nova-2", "deepgram/nova-3", "latest_long",
	// "latest_short", "command_and_search", "phone_call", "video", "default",
	// "medical_conversation", "medical_dictation", "openai/whisper-tiny",
	// "openai/whisper-large-v3-turbo", "xai/grok-stt", "speechmatics/standard",
	// "soniox/stt-rt-v4".
	Model SpeechToTextGetTranscriptionParamsModel `query:"model,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpeechToTextGetTranscriptionParams]'s query parameters as
// `url.Values`.
func (r SpeechToTextGetTranscriptionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The format of input audio stream.
type SpeechToTextGetTranscriptionParamsInputFormat string

const (
	SpeechToTextGetTranscriptionParamsInputFormatMP3 SpeechToTextGetTranscriptionParamsInputFormat = "mp3"
	SpeechToTextGetTranscriptionParamsInputFormatWav SpeechToTextGetTranscriptionParamsInputFormat = "wav"
)

// The transcription engine to use for processing the audio stream.
type SpeechToTextGetTranscriptionParamsTranscriptionEngine string

const (
	SpeechToTextGetTranscriptionParamsTranscriptionEngineAzure        SpeechToTextGetTranscriptionParamsTranscriptionEngine = "Azure"
	SpeechToTextGetTranscriptionParamsTranscriptionEngineDeepgram     SpeechToTextGetTranscriptionParamsTranscriptionEngine = "Deepgram"
	SpeechToTextGetTranscriptionParamsTranscriptionEngineGoogle       SpeechToTextGetTranscriptionParamsTranscriptionEngine = "Google"
	SpeechToTextGetTranscriptionParamsTranscriptionEngineTelnyx       SpeechToTextGetTranscriptionParamsTranscriptionEngine = "Telnyx"
	SpeechToTextGetTranscriptionParamsTranscriptionEngineXAI          SpeechToTextGetTranscriptionParamsTranscriptionEngine = "xAI"
	SpeechToTextGetTranscriptionParamsTranscriptionEngineSpeechmatics SpeechToTextGetTranscriptionParamsTranscriptionEngine = "Speechmatics"
	SpeechToTextGetTranscriptionParamsTranscriptionEngineSoniox       SpeechToTextGetTranscriptionParamsTranscriptionEngine = "Soniox"
)

// The specific model to use within the selected transcription engine.
type SpeechToTextGetTranscriptionParamsModel string

const (
	SpeechToTextGetTranscriptionParamsModelFast                      SpeechToTextGetTranscriptionParamsModel = "fast"
	SpeechToTextGetTranscriptionParamsModelDeepgramNova2             SpeechToTextGetTranscriptionParamsModel = "deepgram/nova-2"
	SpeechToTextGetTranscriptionParamsModelDeepgramNova3             SpeechToTextGetTranscriptionParamsModel = "deepgram/nova-3"
	SpeechToTextGetTranscriptionParamsModelLatestLong                SpeechToTextGetTranscriptionParamsModel = "latest_long"
	SpeechToTextGetTranscriptionParamsModelLatestShort               SpeechToTextGetTranscriptionParamsModel = "latest_short"
	SpeechToTextGetTranscriptionParamsModelCommandAndSearch          SpeechToTextGetTranscriptionParamsModel = "command_and_search"
	SpeechToTextGetTranscriptionParamsModelPhoneCall                 SpeechToTextGetTranscriptionParamsModel = "phone_call"
	SpeechToTextGetTranscriptionParamsModelVideo                     SpeechToTextGetTranscriptionParamsModel = "video"
	SpeechToTextGetTranscriptionParamsModelDefault                   SpeechToTextGetTranscriptionParamsModel = "default"
	SpeechToTextGetTranscriptionParamsModelMedicalConversation       SpeechToTextGetTranscriptionParamsModel = "medical_conversation"
	SpeechToTextGetTranscriptionParamsModelMedicalDictation          SpeechToTextGetTranscriptionParamsModel = "medical_dictation"
	SpeechToTextGetTranscriptionParamsModelOpenAIWhisperTiny         SpeechToTextGetTranscriptionParamsModel = "openai/whisper-tiny"
	SpeechToTextGetTranscriptionParamsModelOpenAIWhisperLargeV3Turbo SpeechToTextGetTranscriptionParamsModel = "openai/whisper-large-v3-turbo"
	SpeechToTextGetTranscriptionParamsModelXaiGrokStt                SpeechToTextGetTranscriptionParamsModel = "xai/grok-stt"
	SpeechToTextGetTranscriptionParamsModelSpeechmaticsStandard      SpeechToTextGetTranscriptionParamsModel = "speechmatics/standard"
	SpeechToTextGetTranscriptionParamsModelSonioxSttRtV4             SpeechToTextGetTranscriptionParamsModel = "soniox/stt-rt-v4"
)
