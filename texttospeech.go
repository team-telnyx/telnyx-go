// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
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

// Converts the provided text to speech using the specified voice and returns audio
// data
func (r *TextToSpeechService) GenerateSpeech(ctx context.Context, body TextToSpeechGenerateSpeechParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "text-to-speech/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of voices that can be used with the text to speech commands.
func (r *TextToSpeechService) ListVoices(ctx context.Context, query TextToSpeechListVoicesParams, opts ...option.RequestOption) (res *TextToSpeechListVoicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "text-to-speech/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

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

type TextToSpeechListVoicesResponseVoice struct {
	ID       string `json:"id"`
	Accent   string `json:"accent"`
	Age      string `json:"age"`
	Gender   string `json:"gender"`
	Label    string `json:"label"`
	Language string `json:"language"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Accent      respjson.Field
		Age         respjson.Field
		Gender      respjson.Field
		Label       respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToSpeechListVoicesResponseVoice) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechListVoicesResponseVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechGenerateSpeechParams struct {
	// The text to convert to speech
	Text string `json:"text,required"`
	// The voice ID in the format Provider.ModelId.VoiceId.
	//
	// Examples:
	//
	// - AWS.Polly.Joanna-Neural
	// - Azure.en-US-AvaMultilingualNeural
	// - ElevenLabs.eleven_multilingual_v2.Rachel
	// - Telnyx.KokoroTTS.af
	//
	// Use the `GET /text-to-speech/voices` endpoint to get a complete list of
	// available voices.
	Voice string `json:"voice,required"`
	paramObj
}

func (r TextToSpeechGenerateSpeechParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechGenerateSpeechParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechGenerateSpeechParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechListVoicesParams struct {
	// Reference to your ElevenLabs API key stored in the Telnyx Portal
	ElevenlabsAPIKeyRef param.Opt[string] `query:"elevenlabs_api_key_ref,omitzero" json:"-"`
	// Filter voices by provider
	//
	// Any of "aws", "azure", "elevenlabs", "telnyx".
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

// Filter voices by provider
type TextToSpeechListVoicesParamsProvider string

const (
	TextToSpeechListVoicesParamsProviderAws        TextToSpeechListVoicesParamsProvider = "aws"
	TextToSpeechListVoicesParamsProviderAzure      TextToSpeechListVoicesParamsProvider = "azure"
	TextToSpeechListVoicesParamsProviderElevenlabs TextToSpeechListVoicesParamsProvider = "elevenlabs"
	TextToSpeechListVoicesParamsProviderTelnyx     TextToSpeechListVoicesParamsProvider = "telnyx"
)
