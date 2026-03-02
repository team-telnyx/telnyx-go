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

// Generate synthesized speech audio from text input. Returns audio in the
// requested format (binary audio stream, base64-encoded JSON, or an audio URL for
// later retrieval).
//
// Authentication is provided via the standard `Authorization: Bearer <API_KEY>`
// header.
//
// The `voice` parameter provides a convenient shorthand to specify provider,
// model, and voice in a single string (e.g. `telnyx.NaturalHD.Alloy`).
// Alternatively, specify `provider` explicitly along with provider-specific
// parameters.
//
// Supported providers: `aws`, `telnyx`, `azure`, `elevenlabs`, `minimax`, `rime`,
// `resemble`.
func (r *TextToSpeechService) Generate(ctx context.Context, body TextToSpeechGenerateParams, opts ...option.RequestOption) (res *TextToSpeechGenerateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "text-to-speech/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// Response when `output_type` is `base64_output`.
type TextToSpeechGenerateResponse struct {
	// Base64-encoded audio data.
	Base64Audio string `json:"base64_audio"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base64Audio respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToSpeechGenerateResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechGenerateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type TextToSpeechGenerateParams struct {
	// When `true`, bypass the audio cache and generate fresh audio.
	DisableCache param.Opt[bool] `json:"disable_cache,omitzero"`
	// Language code (e.g. `en-US`). Usage varies by provider.
	Language param.Opt[string] `json:"language,omitzero"`
	// The text to convert to speech.
	Text param.Opt[string] `json:"text,omitzero"`
	// Voice identifier in the format `provider.model_id.voice_id` or
	// `provider.voice_id`. Examples: `telnyx.NaturalHD.Alloy`,
	// `azure.en-US-AvaMultilingualNeural`, `aws.Polly.Generative.Lucia`. When
	// provided, `provider`, `model_id`, and `voice_id` are extracted automatically and
	// take precedence over individual parameters.
	Voice param.Opt[string] `json:"voice,omitzero"`
	// AWS Polly provider-specific parameters.
	Aws TextToSpeechGenerateParamsAws `json:"aws,omitzero"`
	// Azure Cognitive Services provider-specific parameters.
	Azure TextToSpeechGenerateParamsAzure `json:"azure,omitzero"`
	// ElevenLabs provider-specific parameters.
	Elevenlabs TextToSpeechGenerateParamsElevenlabs `json:"elevenlabs,omitzero"`
	// Minimax provider-specific parameters.
	Minimax TextToSpeechGenerateParamsMinimax `json:"minimax,omitzero"`
	// Determines the response format. `binary_output` returns raw audio bytes,
	// `base64_output` returns base64-encoded audio in JSON.
	//
	// Any of "binary_output", "base64_output".
	OutputType TextToSpeechGenerateParamsOutputType `json:"output_type,omitzero"`
	// TTS provider. Required unless `voice` is provided.
	//
	// Any of "aws", "telnyx", "azure", "elevenlabs", "minimax", "rime", "resemble".
	Provider TextToSpeechGenerateParamsProvider `json:"provider,omitzero"`
	// Resemble AI provider-specific parameters.
	Resemble TextToSpeechGenerateParamsResemble `json:"resemble,omitzero"`
	// Rime provider-specific parameters.
	Rime TextToSpeechGenerateParamsRime `json:"rime,omitzero"`
	// Telnyx provider-specific parameters.
	Telnyx TextToSpeechGenerateParamsTelnyx `json:"telnyx,omitzero"`
	// Text type. Use `ssml` for SSML-formatted input (supported by AWS and Azure).
	//
	// Any of "text", "ssml".
	TextType TextToSpeechGenerateParamsTextType `json:"text_type,omitzero"`
	// Provider-specific voice settings. Contents vary by provider — see
	// provider-specific parameter objects below.
	VoiceSettings map[string]any `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS Polly provider-specific parameters.
type TextToSpeechGenerateParamsAws struct {
	// Language code (e.g. `en-US`, `es-ES`).
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// Audio output format.
	OutputFormat param.Opt[string] `json:"output_format,omitzero"`
	// Audio sample rate.
	SampleRate param.Opt[string] `json:"sample_rate,omitzero"`
	// List of lexicon names to apply.
	LexiconNames []string `json:"lexicon_names,omitzero"`
	// Input text type.
	//
	// Any of "text", "ssml".
	TextType string `json:"text_type,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsAws) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsAws
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsAws) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TextToSpeechGenerateParamsAws](
		"text_type", "text", "ssml",
	)
}

// Azure Cognitive Services provider-specific parameters.
type TextToSpeechGenerateParamsAzure struct {
	// Custom Azure API key. If not provided, the default Telnyx key is used.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// Custom Azure deployment ID.
	DeploymentID param.Opt[string] `json:"deployment_id,omitzero"`
	// Azure audio effect to apply.
	Effect param.Opt[string] `json:"effect,omitzero"`
	// Voice gender preference.
	Gender param.Opt[string] `json:"gender,omitzero"`
	// Language code (e.g. `en-US`).
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// Azure audio output format.
	OutputFormat param.Opt[string] `json:"output_format,omitzero"`
	// Azure region (e.g. `eastus`, `westeurope`).
	Region param.Opt[string] `json:"region,omitzero"`
	// Input text type. Use `ssml` for SSML-formatted input.
	//
	// Any of "text", "ssml".
	TextType string `json:"text_type,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsAzure) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsAzure
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsAzure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TextToSpeechGenerateParamsAzure](
		"text_type", "text", "ssml",
	)
}

// ElevenLabs provider-specific parameters.
type TextToSpeechGenerateParamsElevenlabs struct {
	// Custom ElevenLabs API key. If not provided, the default Telnyx key is used.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// Language code.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// ElevenLabs voice settings (stability, similarity_boost, etc.).
	VoiceSettings map[string]any `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsElevenlabs) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsElevenlabs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsElevenlabs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Minimax provider-specific parameters.
type TextToSpeechGenerateParamsMinimax struct {
	// Language code to boost pronunciation for.
	LanguageBoost param.Opt[string] `json:"language_boost,omitzero"`
	// Pitch adjustment.
	Pitch param.Opt[int64] `json:"pitch,omitzero"`
	// Audio output format.
	ResponseFormat param.Opt[string] `json:"response_format,omitzero"`
	// Speech speed multiplier.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Volume level.
	Vol param.Opt[float64] `json:"vol,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsMinimax) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsMinimax
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsMinimax) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the response format. `binary_output` returns raw audio bytes,
// `base64_output` returns base64-encoded audio in JSON.
type TextToSpeechGenerateParamsOutputType string

const (
	TextToSpeechGenerateParamsOutputTypeBinaryOutput TextToSpeechGenerateParamsOutputType = "binary_output"
	TextToSpeechGenerateParamsOutputTypeBase64Output TextToSpeechGenerateParamsOutputType = "base64_output"
)

// TTS provider. Required unless `voice` is provided.
type TextToSpeechGenerateParamsProvider string

const (
	TextToSpeechGenerateParamsProviderAws        TextToSpeechGenerateParamsProvider = "aws"
	TextToSpeechGenerateParamsProviderTelnyx     TextToSpeechGenerateParamsProvider = "telnyx"
	TextToSpeechGenerateParamsProviderAzure      TextToSpeechGenerateParamsProvider = "azure"
	TextToSpeechGenerateParamsProviderElevenlabs TextToSpeechGenerateParamsProvider = "elevenlabs"
	TextToSpeechGenerateParamsProviderMinimax    TextToSpeechGenerateParamsProvider = "minimax"
	TextToSpeechGenerateParamsProviderRime       TextToSpeechGenerateParamsProvider = "rime"
	TextToSpeechGenerateParamsProviderResemble   TextToSpeechGenerateParamsProvider = "resemble"
)

// Resemble AI provider-specific parameters.
type TextToSpeechGenerateParamsResemble struct {
	// Custom Resemble API key.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// Audio output format.
	Format param.Opt[string] `json:"format,omitzero"`
	// Synthesis precision.
	Precision param.Opt[string] `json:"precision,omitzero"`
	// Audio sample rate.
	SampleRate param.Opt[string] `json:"sample_rate,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsResemble) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsResemble
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsResemble) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rime provider-specific parameters.
type TextToSpeechGenerateParamsRime struct {
	// Audio output format.
	ResponseFormat param.Opt[string] `json:"response_format,omitzero"`
	// Audio sampling rate in Hz.
	SamplingRate param.Opt[int64] `json:"sampling_rate,omitzero"`
	// Voice speed multiplier.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsRime) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsRime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsRime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Telnyx provider-specific parameters.
type TextToSpeechGenerateParamsTelnyx struct {
	// Audio response format.
	ResponseFormat param.Opt[string] `json:"response_format,omitzero"`
	// Audio sampling rate in Hz.
	SamplingRate param.Opt[int64] `json:"sampling_rate,omitzero"`
	// Sampling temperature.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Voice speed multiplier.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateParamsTelnyx) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateParamsTelnyx
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateParamsTelnyx) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text type. Use `ssml` for SSML-formatted input (supported by AWS and Azure).
type TextToSpeechGenerateParamsTextType string

const (
	TextToSpeechGenerateParamsTextTypeText TextToSpeechGenerateParamsTextType = "text"
	TextToSpeechGenerateParamsTextTypeSsml TextToSpeechGenerateParamsTextType = "ssml"
)

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
