// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Text to speech streaming command operations
//
// TextToSpeechService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTextToSpeechService] method instead.
type TextToSpeechService struct {
	Options []option.RequestOption
}

// NewTextToSpeechService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTextToSpeechService(opts ...option.RequestOption) (r TextToSpeechService) {
	r = TextToSpeechService{}
	r.Options = opts
	return
}

// Retrieve a list of available voices from one or all TTS providers. When
// `provider` is specified, returns voices for that provider only. Otherwise,
// returns voices from all providers.
//
// Some providers (ElevenLabs, Resemble) require an API key to list voices.
func (r *TextToSpeechService) ListVoices(ctx context.Context, query TextToSpeechListVoicesParams, opts ...option.RequestOption) (res *TextToSpeechListVoicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "text-to-speech/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Open a WebSocket connection to stream text and receive synthesized audio in real
// time. Authentication is provided via the standard
// `Authorization: Bearer <API_KEY>` header. Send JSON frames with text to
// synthesize; receive JSON frames containing base64-encoded audio chunks.
//
// Supported providers: `aws`, `telnyx`, `azure`, `murfai`, `minimax`, `rime`,
// `resemble`, `elevenlabs`.
//
// **Connection flow:**
//
//  1. Open WebSocket with query parameters specifying provider, voice, and model.
//  2. Send an initial handshake message `{"text": " "}` (single space) with
//     optional `voice_settings` to initialize the session.
//  3. Send text messages as `{"text": "Hello world"}`.
//  4. Receive audio chunks as JSON frames with base64-encoded audio.
//  5. A final frame with `isFinal: true` indicates the end of audio for the current
//     text.
//
// To interrupt and restart synthesis mid-stream, send `{"force": true}` — the
// current worker is stopped and a new one is started.
func (r *TextToSpeechService) Stream(ctx context.Context, query TextToSpeechStreamParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "text-to-speech/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// List of available voices.
type TextToSpeechListVoicesResponse struct {
	Voices []TextToSpeechListVoicesResponseVoice `json:"voices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToSpeechListVoicesResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechListVoicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A voice available for text-to-speech synthesis.
type TextToSpeechListVoicesResponseVoice struct {
	// Voice gender.
	Gender string `json:"gender"`
	// Language code.
	Language string `json:"language"`
	// Voice name.
	Name string `json:"name"`
	// The TTS provider.
	Provider string `json:"provider"`
	// Voice identifier.
	VoiceID string `json:"voice_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gender      respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Provider    respjson.Field
		VoiceID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToSpeechListVoicesResponseVoice) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechListVoicesResponseVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechListVoicesParams struct {
	// API key for providers that require one to list voices (e.g. ElevenLabs).
	APIKey param.Opt[string] `query:"api_key,omitzero" json:"-"`
	// Filter voices by provider. If omitted, voices from all providers are returned.
	//
	// Any of "aws", "telnyx", "azure", "elevenlabs", "minimax", "rime", "resemble".
	Provider TextToSpeechListVoicesParamsProvider `query:"provider,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TextToSpeechListVoicesParams]'s query parameters as
// `url.Values`.
func (r TextToSpeechListVoicesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter voices by provider. If omitted, voices from all providers are returned.
type TextToSpeechListVoicesParamsProvider string

const (
	TextToSpeechListVoicesParamsProviderAws        TextToSpeechListVoicesParamsProvider = "aws"
	TextToSpeechListVoicesParamsProviderTelnyx     TextToSpeechListVoicesParamsProvider = "telnyx"
	TextToSpeechListVoicesParamsProviderAzure      TextToSpeechListVoicesParamsProvider = "azure"
	TextToSpeechListVoicesParamsProviderElevenlabs TextToSpeechListVoicesParamsProvider = "elevenlabs"
	TextToSpeechListVoicesParamsProviderMinimax    TextToSpeechListVoicesParamsProvider = "minimax"
	TextToSpeechListVoicesParamsProviderRime       TextToSpeechListVoicesParamsProvider = "rime"
	TextToSpeechListVoicesParamsProviderResemble   TextToSpeechListVoicesParamsProvider = "resemble"
)

type TextToSpeechStreamParams struct {
	// When `true`, bypass the audio cache and generate fresh audio.
	DisableCache param.Opt[bool] `query:"disable_cache,omitzero" json:"-"`
	// Model identifier for the chosen provider. Examples: `Natural`, `NaturalHD`
	// (Telnyx); `Polly.Generative` (AWS).
	ModelID param.Opt[string] `query:"model_id,omitzero" json:"-"`
	// Client-provided socket identifier for tracking. If not provided, one is
	// generated server-side.
	SocketID param.Opt[string] `query:"socket_id,omitzero" json:"-"`
	// Voice identifier in the format `provider.model_id.voice_id` or
	// `provider.voice_id` (e.g. `telnyx.NaturalHD.Telnyx_Alloy` or
	// `azure.en-US-AvaMultilingualNeural`). When provided, the `provider`, `model_id`,
	// and `voice_id` are extracted automatically. Takes precedence over individual
	// `provider`/`model_id`/`voice_id` parameters.
	Voice param.Opt[string] `query:"voice,omitzero" json:"-"`
	// Voice identifier for the chosen provider.
	VoiceID param.Opt[string] `query:"voice_id,omitzero" json:"-"`
	// Audio output format override. Supported for Telnyx `Natural`/`NaturalHD` models
	// only. Accepted values: `pcm`, `wav`.
	//
	// Any of "pcm", "wav".
	AudioFormat TextToSpeechStreamParamsAudioFormat `query:"audio_format,omitzero" json:"-"`
	// TTS provider. Defaults to `telnyx` if not specified. Ignored when `voice` is
	// provided.
	//
	// Any of "aws", "telnyx", "azure", "elevenlabs", "minimax", "murfai", "rime",
	// "resemble".
	Provider TextToSpeechStreamParamsProvider `query:"provider,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TextToSpeechStreamParams]'s query parameters as
// `url.Values`.
func (r TextToSpeechStreamParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Audio output format override. Supported for Telnyx `Natural`/`NaturalHD` models
// only. Accepted values: `pcm`, `wav`.
type TextToSpeechStreamParamsAudioFormat string

const (
	TextToSpeechStreamParamsAudioFormatPcm TextToSpeechStreamParamsAudioFormat = "pcm"
	TextToSpeechStreamParamsAudioFormatWav TextToSpeechStreamParamsAudioFormat = "wav"
)

// TTS provider. Defaults to `telnyx` if not specified. Ignored when `voice` is
// provided.
type TextToSpeechStreamParamsProvider string

const (
	TextToSpeechStreamParamsProviderAws        TextToSpeechStreamParamsProvider = "aws"
	TextToSpeechStreamParamsProviderTelnyx     TextToSpeechStreamParamsProvider = "telnyx"
	TextToSpeechStreamParamsProviderAzure      TextToSpeechStreamParamsProvider = "azure"
	TextToSpeechStreamParamsProviderElevenlabs TextToSpeechStreamParamsProvider = "elevenlabs"
	TextToSpeechStreamParamsProviderMinimax    TextToSpeechStreamParamsProvider = "minimax"
	TextToSpeechStreamParamsProviderMurfai     TextToSpeechStreamParamsProvider = "murfai"
	TextToSpeechStreamParamsProviderRime       TextToSpeechStreamParamsProvider = "rime"
	TextToSpeechStreamParamsProviderResemble   TextToSpeechStreamParamsProvider = "resemble"
)
