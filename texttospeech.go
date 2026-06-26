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

// Generate synthesized speech audio from text input. Returns audio in the
// requested format (binary audio stream, base64-encoded JSON, or an audio URL for
// later retrieval).
//
// Authentication is provided via the standard `Authorization: Bearer <API_KEY>`
// header.
//
// The `voice` parameter provides a convenient shorthand to specify provider,
// model, and voice in a single string (e.g. `telnyx.NaturalHD.Alloy` or
// `Telnyx.Ultra.<voice_id>`). Alternatively, specify `provider` explicitly along
// with provider-specific parameters.
//
// Supported providers: `aws`, `telnyx`, `azure`, `elevenlabs`, `minimax`, `rime`,
// `resemble`, `xai`.
//
// The Telnyx `Ultra` model supports 44 languages with emotion control, speed
// adjustment, and volume control. Use the `telnyx` provider-specific parameters to
// configure these features.
func (r *TextToSpeechService) GenerateSpeech(ctx context.Context, body TextToSpeechGenerateSpeechParams, opts ...option.RequestOption) (res *TextToSpeechGenerateSpeechResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "text-to-speech/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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
	return res, err
}

// Open a WebSocket connection to stream text and receive synthesized audio in real
// time. Authentication is provided via the standard
// `Authorization: Bearer <API_KEY>` header. Send JSON frames with text to
// synthesize; receive JSON frames containing base64-encoded audio chunks.
//
// Supported providers: `aws`, `telnyx`, `azure`, `murfai`, `minimax`, `rime`,
// `resemble`, `elevenlabs`, `xai`.
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
//
// **Note:** The Telnyx `Ultra` model is not available over WebSocket. Use the HTTP
// POST `/text-to-speech/speech` endpoint instead.
func (r *TextToSpeechService) GetSpeech(ctx context.Context, query TextToSpeechGetSpeechParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "text-to-speech/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Response when `output_type` is `base64_output`.
type TextToSpeechGenerateSpeechResponse struct {
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
func (r TextToSpeechGenerateSpeechResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechGenerateSpeechResponse) UnmarshalJSON(data []byte) error {
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

// Client-to-server frame containing text to synthesize.
type StreamClientEvent struct {
	// Text to convert to speech. Send `" "` (single space) as an initial handshake
	// with optional `voice_settings`. Subsequent messages contain the actual text to
	// synthesize.
	Text string `json:"text" api:"required"`
	// When `true`, stops the current synthesis worker and starts a new one. Used to
	// interrupt speech mid-stream and begin synthesizing new text.
	Force bool `json:"force"`
	// Provider-specific voice settings sent with the initial handshake. Contents vary
	// by provider — e.g. `{"speed": 1.2}` for Minimax, `{"voice_speed": 1.5}` for
	// Telnyx.
	VoiceSettings map[string]any `json:"voice_settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text          respjson.Field
		Force         respjson.Field
		VoiceSettings respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamClientEvent) RawJSON() string { return r.JSON.raw }
func (r *StreamClientEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StreamClientEvent to a StreamClientEventParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StreamClientEventParam.Overrides()
func (r StreamClientEvent) ToParam() StreamClientEventParam {
	return param.Override[StreamClientEventParam](json.RawMessage(r.RawJSON()))
}

// Client-to-server frame containing text to synthesize.
//
// The property Text is required.
type StreamClientEventParam struct {
	// Text to convert to speech. Send `" "` (single space) as an initial handshake
	// with optional `voice_settings`. Subsequent messages contain the actual text to
	// synthesize.
	Text string `json:"text" api:"required"`
	// When `true`, stops the current synthesis worker and starts a new one. Used to
	// interrupt speech mid-stream and begin synthesizing new text.
	Force param.Opt[bool] `json:"force,omitzero"`
	// Provider-specific voice settings sent with the initial handshake. Contents vary
	// by provider — e.g. `{"speed": 1.2}` for Minimax, `{"voice_speed": 1.5}` for
	// Telnyx.
	VoiceSettings map[string]any `json:"voice_settings,omitzero"`
	paramObj
}

func (r StreamClientEventParam) MarshalJSON() (data []byte, err error) {
	type shadow StreamClientEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamClientEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StreamServerEventUnion contains all possible properties and values from
// [StreamServerEventAudioChunk], [StreamServerEventFinalFrameEvent],
// [StreamServerEventError].
//
// Use the [StreamServerEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StreamServerEventUnion struct {
	// This field is a union of [string], [any]
	Audio StreamServerEventUnionAudio `json:"audio"`
	// This field is from variant [StreamServerEventAudioChunk].
	Cached                  bool   `json:"cached"`
	IsFinal                 bool   `json:"isFinal"`
	Text                    string `json:"text"`
	TimeToFirstAudioFrameMs int64  `json:"timeToFirstAudioFrameMs"`
	// Any of "audio_chunk", "final", "error".
	Type string `json:"type"`
	// This field is from variant [StreamServerEventError].
	Error string `json:"error"`
	JSON  struct {
		Audio                   respjson.Field
		Cached                  respjson.Field
		IsFinal                 respjson.Field
		Text                    respjson.Field
		TimeToFirstAudioFrameMs respjson.Field
		Type                    respjson.Field
		Error                   respjson.Field
		raw                     string
	} `json:"-"`
}

// anyStreamServerEvent is implemented by each variant of [StreamServerEventUnion]
// to add type safety for the return type of [StreamServerEventUnion.AsAny]
type anyStreamServerEvent interface {
	implStreamServerEventUnion()
}

func (StreamServerEventAudioChunk) implStreamServerEventUnion()      {}
func (StreamServerEventFinalFrameEvent) implStreamServerEventUnion() {}
func (StreamServerEventError) implStreamServerEventUnion()           {}

// Use the following switch statement to find the correct variant
//
//	switch variant := StreamServerEventUnion.AsAny().(type) {
//	case telnyx.StreamServerEventAudioChunk:
//	case telnyx.StreamServerEventFinalFrameEvent:
//	case telnyx.StreamServerEventError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u StreamServerEventUnion) AsAny() anyStreamServerEvent {
	switch u.Type {
	case "audio_chunk":
		return u.AsAudioChunk()
	case "final":
		return u.AsFinalFrameEvent()
	case "error":
		return u.AsError()
	}
	return nil
}

func (u StreamServerEventUnion) AsAudioChunk() (v StreamServerEventAudioChunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StreamServerEventUnion) AsFinalFrameEvent() (v StreamServerEventFinalFrameEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StreamServerEventUnion) AsError() (v StreamServerEventError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StreamServerEventUnion) RawJSON() string { return u.JSON.raw }

func (r *StreamServerEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StreamServerEventUnionAudio is an implicit subunion of [StreamServerEventUnion].
// StreamServerEventUnionAudio provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [StreamServerEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStreamServerEventFinalFrameEventAudio]
type StreamServerEventUnionAudio struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfStreamServerEventFinalFrameEventAudio any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfStreamServerEventFinalFrameEventAudio respjson.Field
		raw                                     string
	} `json:"-"`
}

func (r *StreamServerEventUnionAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-to-client frame containing a base64-encoded audio chunk.
type StreamServerEventAudioChunk struct {
	// Base64-encoded audio data. May be `null` for providers that use
	// `drop_concatenated_audio` mode (Telnyx Natural/NaturalHD, Rime, Minimax, MurfAI,
	// Resemble) — in that case only streamed chunks carry audio.
	Audio string `json:"audio" api:"nullable"`
	// Whether this audio was served from cache.
	Cached bool `json:"cached"`
	// Always `false` for audio chunk frames.
	IsFinal bool `json:"isFinal"`
	// The text segment that this audio chunk corresponds to.
	Text string `json:"text" api:"nullable"`
	// Milliseconds from the start-of-speech request to the first audio frame. Only
	// present on the first audio chunk of a synthesis request.
	TimeToFirstAudioFrameMs int64 `json:"timeToFirstAudioFrameMs"`
	// Frame type identifier.
	//
	// Any of "audio_chunk".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio                   respjson.Field
		Cached                  respjson.Field
		IsFinal                 respjson.Field
		Text                    respjson.Field
		TimeToFirstAudioFrameMs respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamServerEventAudioChunk) RawJSON() string { return r.JSON.raw }
func (r *StreamServerEventAudioChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-to-client frame indicating synthesis is complete for the current text.
type StreamServerEventFinalFrameEvent struct {
	// Always `null` for the final frame.
	Audio any `json:"audio" api:"nullable"`
	// Always `true`.
	//
	// Any of true.
	IsFinal bool `json:"isFinal"`
	// Empty string.
	Text string `json:"text"`
	// Present if this was the first response frame.
	TimeToFirstAudioFrameMs int64 `json:"timeToFirstAudioFrameMs"`
	// Frame type identifier.
	//
	// Any of "final".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio                   respjson.Field
		IsFinal                 respjson.Field
		Text                    respjson.Field
		TimeToFirstAudioFrameMs respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamServerEventFinalFrameEvent) RawJSON() string { return r.JSON.raw }
func (r *StreamServerEventFinalFrameEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-to-client frame indicating an error during synthesis. The connection is
// closed shortly after.
type StreamServerEventError struct {
	// Error message describing what went wrong.
	Error string `json:"error"`
	// Frame type identifier.
	//
	// Any of "error".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamServerEventError) RawJSON() string { return r.JSON.raw }
func (r *StreamServerEventError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechGenerateSpeechParams struct {
	// When `true`, bypass the audio cache and generate fresh audio.
	DisableCache param.Opt[bool] `json:"disable_cache,omitzero"`
	// Language code (e.g. `en-US`). Usage varies by provider.
	Language param.Opt[string] `json:"language,omitzero"`
	// The text to convert to speech.
	Text param.Opt[string] `json:"text,omitzero"`
	// Voice identifier in the format `provider.model_id.voice_id` or
	// `provider.voice_id`. Examples: `telnyx.NaturalHD.Alloy`,
	// `Telnyx.Ultra.<voice_id>`, `azure.en-US-AvaMultilingualNeural`,
	// `aws.Polly.Generative.Lucia`. When provided, `provider`, `model_id`, and
	// `voice_id` are extracted automatically and take precedence over individual
	// parameters.
	Voice param.Opt[string] `json:"voice,omitzero"`
	// AWS Polly provider-specific parameters.
	Aws TextToSpeechGenerateSpeechParamsAws `json:"aws,omitzero"`
	// Azure Cognitive Services provider-specific parameters.
	Azure TextToSpeechGenerateSpeechParamsAzure `json:"azure,omitzero"`
	// ElevenLabs provider-specific parameters.
	Elevenlabs TextToSpeechGenerateSpeechParamsElevenlabs `json:"elevenlabs,omitzero"`
	// Minimax provider-specific parameters.
	Minimax TextToSpeechGenerateSpeechParamsMinimax `json:"minimax,omitzero"`
	// Determines the response format. `binary_output` returns raw audio bytes,
	// `base64_output` returns base64-encoded audio in JSON.
	//
	// Any of "binary_output", "base64_output".
	OutputType TextToSpeechGenerateSpeechParamsOutputType `json:"output_type,omitzero"`
	// TTS provider. Required unless `voice` is provided.
	//
	// Any of "aws", "telnyx", "azure", "elevenlabs", "minimax", "rime", "resemble",
	// "xai".
	Provider TextToSpeechGenerateSpeechParamsProvider `json:"provider,omitzero"`
	// Resemble AI provider-specific parameters.
	Resemble TextToSpeechGenerateSpeechParamsResemble `json:"resemble,omitzero"`
	// Rime provider-specific parameters.
	Rime TextToSpeechGenerateSpeechParamsRime `json:"rime,omitzero"`
	// Telnyx provider-specific parameters. Use `voice_speed` and `temperature` for
	// `Natural` and `NaturalHD` models. For the `Ultra` model, use `voice_speed`,
	// `volume`, and `emotion`.
	Telnyx TextToSpeechGenerateSpeechParamsTelnyx `json:"telnyx,omitzero"`
	// Text type. Use `ssml` for SSML-formatted input (supported by AWS and Azure).
	//
	// Any of "text", "ssml".
	TextType TextToSpeechGenerateSpeechParamsTextType `json:"text_type,omitzero"`
	// Provider-specific voice settings. Contents vary by provider — see
	// provider-specific parameter objects below.
	VoiceSettings map[string]any `json:"voice_settings,omitzero"`
	// xAI provider-specific parameters.
	Xai TextToSpeechGenerateSpeechParamsXai `json:"xai,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateSpeechParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS Polly provider-specific parameters.
type TextToSpeechGenerateSpeechParamsAws struct {
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

func (r TextToSpeechGenerateSpeechParamsAws) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsAws
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsAws) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TextToSpeechGenerateSpeechParamsAws](
		"text_type", "text", "ssml",
	)
}

// Azure Cognitive Services provider-specific parameters.
type TextToSpeechGenerateSpeechParamsAzure struct {
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

func (r TextToSpeechGenerateSpeechParamsAzure) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsAzure
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsAzure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TextToSpeechGenerateSpeechParamsAzure](
		"text_type", "text", "ssml",
	)
}

// ElevenLabs provider-specific parameters.
type TextToSpeechGenerateSpeechParamsElevenlabs struct {
	// Custom ElevenLabs API key. If not provided, the default Telnyx key is used.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// Language code.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// ElevenLabs voice settings (stability, similarity_boost, etc.).
	VoiceSettings map[string]any `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateSpeechParamsElevenlabs) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsElevenlabs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsElevenlabs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Minimax provider-specific parameters.
type TextToSpeechGenerateSpeechParamsMinimax struct {
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

func (r TextToSpeechGenerateSpeechParamsMinimax) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsMinimax
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsMinimax) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the response format. `binary_output` returns raw audio bytes,
// `base64_output` returns base64-encoded audio in JSON.
type TextToSpeechGenerateSpeechParamsOutputType string

const (
	TextToSpeechGenerateSpeechParamsOutputTypeBinaryOutput TextToSpeechGenerateSpeechParamsOutputType = "binary_output"
	TextToSpeechGenerateSpeechParamsOutputTypeBase64Output TextToSpeechGenerateSpeechParamsOutputType = "base64_output"
)

// TTS provider. Required unless `voice` is provided.
type TextToSpeechGenerateSpeechParamsProvider string

const (
	TextToSpeechGenerateSpeechParamsProviderAws        TextToSpeechGenerateSpeechParamsProvider = "aws"
	TextToSpeechGenerateSpeechParamsProviderTelnyx     TextToSpeechGenerateSpeechParamsProvider = "telnyx"
	TextToSpeechGenerateSpeechParamsProviderAzure      TextToSpeechGenerateSpeechParamsProvider = "azure"
	TextToSpeechGenerateSpeechParamsProviderElevenlabs TextToSpeechGenerateSpeechParamsProvider = "elevenlabs"
	TextToSpeechGenerateSpeechParamsProviderMinimax    TextToSpeechGenerateSpeechParamsProvider = "minimax"
	TextToSpeechGenerateSpeechParamsProviderRime       TextToSpeechGenerateSpeechParamsProvider = "rime"
	TextToSpeechGenerateSpeechParamsProviderResemble   TextToSpeechGenerateSpeechParamsProvider = "resemble"
	TextToSpeechGenerateSpeechParamsProviderXai        TextToSpeechGenerateSpeechParamsProvider = "xai"
)

// Resemble AI provider-specific parameters.
type TextToSpeechGenerateSpeechParamsResemble struct {
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

func (r TextToSpeechGenerateSpeechParamsResemble) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsResemble
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsResemble) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rime provider-specific parameters.
type TextToSpeechGenerateSpeechParamsRime struct {
	// Audio output format.
	ResponseFormat param.Opt[string] `json:"response_format,omitzero"`
	// Audio sampling rate in Hz.
	SamplingRate param.Opt[int64] `json:"sampling_rate,omitzero"`
	// Voice speed multiplier.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateSpeechParamsRime) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsRime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsRime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Telnyx provider-specific parameters. Use `voice_speed` and `temperature` for
// `Natural` and `NaturalHD` models. For the `Ultra` model, use `voice_speed`,
// `volume`, and `emotion`.
type TextToSpeechGenerateSpeechParamsTelnyx struct {
	// Audio response format.
	ResponseFormat param.Opt[string] `json:"response_format,omitzero"`
	// Audio sampling rate in Hz.
	SamplingRate param.Opt[int64] `json:"sampling_rate,omitzero"`
	// Sampling temperature. Applies to `Natural` and `NaturalHD` models only.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Voice speed multiplier. Applies to all models. Range: 0.5 to 2.0.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
	// Volume level for the Ultra model. Range: 0.0 to 2.0.
	Volume param.Opt[float64] `json:"volume,omitzero"`
	// Emotion control for the Ultra model. Adjusts the emotional tone of the
	// synthesized speech.
	//
	// Any of "neutral", "happy", "sad", "angry", "fearful", "disgusted", "surprised".
	Emotion string `json:"emotion,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateSpeechParamsTelnyx) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsTelnyx
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsTelnyx) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TextToSpeechGenerateSpeechParamsTelnyx](
		"emotion", "neutral", "happy", "sad", "angry", "fearful", "disgusted", "surprised",
	)
}

// Text type. Use `ssml` for SSML-formatted input (supported by AWS and Azure).
type TextToSpeechGenerateSpeechParamsTextType string

const (
	TextToSpeechGenerateSpeechParamsTextTypeText TextToSpeechGenerateSpeechParamsTextType = "text"
	TextToSpeechGenerateSpeechParamsTextTypeSsml TextToSpeechGenerateSpeechParamsTextType = "ssml"
)

// xAI provider-specific parameters.
//
// The property VoiceID is required.
type TextToSpeechGenerateSpeechParamsXai struct {
	// xAI voice identifier.
	//
	// Any of "eve", "ara", "rex", "sal", "leo".
	VoiceID string `json:"voice_id,omitzero" api:"required"`
	// Language code, or `auto` to detect.
	Language param.Opt[string] `json:"language,omitzero"`
	// Audio output format.
	//
	// Any of "mp3", "wav", "pcm", "mulaw", "alaw".
	OutputFormat string `json:"output_format,omitzero"`
	// Audio sample rate in Hz.
	//
	// Any of 8000, 16000, 22050, 24000, 44100, 48000.
	SampleRate int64 `json:"sample_rate,omitzero"`
	paramObj
}

func (r TextToSpeechGenerateSpeechParamsXai) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParamsXai
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParamsXai) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TextToSpeechGenerateSpeechParamsXai](
		"voice_id", "eve", "ara", "rex", "sal", "leo",
	)
	apijson.RegisterFieldValidator[TextToSpeechGenerateSpeechParamsXai](
		"output_format", "mp3", "wav", "pcm", "mulaw", "alaw",
	)
	apijson.RegisterFieldValidator[TextToSpeechGenerateSpeechParamsXai](
		"sample_rate", 8000, 16000, 22050, 24000, 44100, 48000,
	)
}

type TextToSpeechListVoicesParams struct {
	// API key for providers that require one to list voices (e.g. ElevenLabs).
	APIKey param.Opt[string] `query:"api_key,omitzero" json:"-"`
	// Filter voices by provider. If omitted, voices from all providers are returned.
	//
	// Any of "aws", "telnyx", "azure", "elevenlabs", "minimax", "rime", "resemble",
	// "xai".
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
	TextToSpeechListVoicesParamsProviderXai        TextToSpeechListVoicesParamsProvider = "xai"
)

type TextToSpeechGetSpeechParams struct {
	// When `true`, bypass the audio cache and generate fresh audio.
	DisableCache param.Opt[bool] `query:"disable_cache,omitzero" json:"-"`
	// Model identifier for the chosen provider. Examples: `Natural`, `NaturalHD`,
	// `Ultra` (Telnyx); `Polly.Generative` (AWS).
	ModelID param.Opt[string] `query:"model_id,omitzero" json:"-"`
	// Client-provided socket identifier for tracking. If not provided, one is
	// generated server-side.
	SocketID param.Opt[string] `query:"socket_id,omitzero" json:"-"`
	// Voice identifier in the format `provider.model_id.voice_id` or
	// `provider.voice_id` (e.g. `telnyx.NaturalHD.Telnyx_Alloy`,
	// `Telnyx.Ultra.<voice_id>`, or `azure.en-US-AvaMultilingualNeural`). When
	// provided, the `provider`, `model_id`, and `voice_id` are extracted
	// automatically. Takes precedence over individual `provider`/`model_id`/`voice_id`
	// parameters.
	Voice param.Opt[string] `query:"voice,omitzero" json:"-"`
	// Voice identifier for the chosen provider.
	VoiceID param.Opt[string] `query:"voice_id,omitzero" json:"-"`
	// Audio output format override. Supported for Telnyx models. `pcm` and `wav` are
	// available for `Natural`/`NaturalHD` models. The `Ultra` model outputs PCM at
	// 24kHz s16le or MP3 at 128kbps 24kHz.
	//
	// Any of "pcm", "wav", "mp3".
	AudioFormat TextToSpeechGetSpeechParamsAudioFormat `query:"audio_format,omitzero" json:"-"`
	// TTS provider. Defaults to `telnyx` if not specified. Ignored when `voice` is
	// provided.
	//
	// Any of "aws", "telnyx", "azure", "elevenlabs", "minimax", "murfai", "rime",
	// "resemble", "xai".
	Provider TextToSpeechGetSpeechParamsProvider `query:"provider,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TextToSpeechGetSpeechParams]'s query parameters as
// `url.Values`.
func (r TextToSpeechGetSpeechParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Audio output format override. Supported for Telnyx models. `pcm` and `wav` are
// available for `Natural`/`NaturalHD` models. The `Ultra` model outputs PCM at
// 24kHz s16le or MP3 at 128kbps 24kHz.
type TextToSpeechGetSpeechParamsAudioFormat string

const (
	TextToSpeechGetSpeechParamsAudioFormatPcm TextToSpeechGetSpeechParamsAudioFormat = "pcm"
	TextToSpeechGetSpeechParamsAudioFormatWav TextToSpeechGetSpeechParamsAudioFormat = "wav"
	TextToSpeechGetSpeechParamsAudioFormatMP3 TextToSpeechGetSpeechParamsAudioFormat = "mp3"
)

// TTS provider. Defaults to `telnyx` if not specified. Ignored when `voice` is
// provided.
type TextToSpeechGetSpeechParamsProvider string

const (
	TextToSpeechGetSpeechParamsProviderAws        TextToSpeechGetSpeechParamsProvider = "aws"
	TextToSpeechGetSpeechParamsProviderTelnyx     TextToSpeechGetSpeechParamsProvider = "telnyx"
	TextToSpeechGetSpeechParamsProviderAzure      TextToSpeechGetSpeechParamsProvider = "azure"
	TextToSpeechGetSpeechParamsProviderElevenlabs TextToSpeechGetSpeechParamsProvider = "elevenlabs"
	TextToSpeechGetSpeechParamsProviderMinimax    TextToSpeechGetSpeechParamsProvider = "minimax"
	TextToSpeechGetSpeechParamsProviderMurfai     TextToSpeechGetSpeechParamsProvider = "murfai"
	TextToSpeechGetSpeechParamsProviderRime       TextToSpeechGetSpeechParamsProvider = "rime"
	TextToSpeechGetSpeechParamsProviderResemble   TextToSpeechGetSpeechParamsProvider = "resemble"
	TextToSpeechGetSpeechParamsProviderXai        TextToSpeechGetSpeechParamsProvider = "xai"
)
