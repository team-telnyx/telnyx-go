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
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Discover available speech-to-text providers, models, and supported languages.
//
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
//   - `file_transcription` — file-based transcription via
//     `/ai/audio/transcriptions`.
//   - `in_call_transcription` — live call transcription via Call Control
//     `transcription_start`.
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

// A (provider, model) tuple along with its supported service types and languages.
type SpeechToTextListProvidersResponseData struct {
	// Languages this (provider, model) accepts, in the provider's native code format.
	// `auto` indicates the provider performs language detection.
	Languages []string `json:"languages" api:"required"`
	// Provider-scoped model name.
	Model string `json:"model" api:"required"`
	// STT provider name.
	Provider string `json:"provider" api:"required"`
	// Service surfaces this (provider, model) supports.
	//
	// Any of "streaming", "file_transcription", "in_call_transcription".
	ServiceTypes []string `json:"service_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Languages    respjson.Field
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
	// Filter to entries that support the given service type.
	//
	// Any of "streaming", "file_transcription", "in_call_transcription".
	ServiceType SpeechToTextListProvidersParamsServiceType `query:"service_type,omitzero" json:"-"`
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

// Filter to entries that support the given service type.
type SpeechToTextListProvidersParamsServiceType string

const (
	SpeechToTextListProvidersParamsServiceTypeStreaming           SpeechToTextListProvidersParamsServiceType = "streaming"
	SpeechToTextListProvidersParamsServiceTypeFileTranscription   SpeechToTextListProvidersParamsServiceType = "file_transcription"
	SpeechToTextListProvidersParamsServiceTypeInCallTranscription SpeechToTextListProvidersParamsServiceType = "in_call_transcription"
)
