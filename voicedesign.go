// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create and manage AI-generated voice designs using natural language prompts.
//
// VoiceDesignService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceDesignService] method instead.
type VoiceDesignService struct {
	Options []option.RequestOption
}

// NewVoiceDesignService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceDesignService(opts ...option.RequestOption) (r VoiceDesignService) {
	r = VoiceDesignService{}
	r.Options = opts
	return
}

// Creates a new voice design (version 1) when `voice_design_id` is omitted. When
// `voice_design_id` is provided, adds a new version to the existing design
// instead. A design can have at most 50 versions.
func (r *VoiceDesignService) New(ctx context.Context, body VoiceDesignNewParams, opts ...option.RequestOption) (res *VoiceDesignNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice_designs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the latest version of a voice design, or a specific version when
// `?version=N` is provided. The `id` parameter accepts either a UUID or the design
// name.
func (r *VoiceDesignService) Get(ctx context.Context, id string, query VoiceDesignGetParams, opts ...option.RequestOption) (res *VoiceDesignGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("voice_designs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of voice designs belonging to the authenticated
// account.
func (r *VoiceDesignService) List(ctx context.Context, query VoiceDesignListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[VoiceDesignListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "voice_designs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a paginated list of voice designs belonging to the authenticated
// account.
func (r *VoiceDesignService) ListAutoPaging(ctx context.Context, query VoiceDesignListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[VoiceDesignListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a voice design and all of its versions. This action cannot
// be undone.
func (r *VoiceDesignService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("voice_designs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Permanently deletes a specific version of a voice design. The version number
// must be a positive integer.
func (r *VoiceDesignService) DeleteVersion(ctx context.Context, version int64, body VoiceDesignDeleteVersionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("voice_designs/%s/versions/%v", body.ID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Downloads the WAV audio sample for the voice design. Returns the latest
// version's sample by default, or a specific version when `?version=N` is
// provided. The `id` parameter accepts either a UUID or the design name.
func (r *VoiceDesignService) DownloadSample(ctx context.Context, id string, query VoiceDesignDownloadSampleParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/wav")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("voice_designs/%s/sample", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates the name of a voice design. All versions retain their other properties.
func (r *VoiceDesignService) Rename(ctx context.Context, id string, body VoiceDesignRenameParams, opts ...option.RequestOption) (res *VoiceDesignRenameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("voice_designs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Response envelope for a single voice design with full version detail.
type VoiceDesignNewResponse struct {
	// A voice design object with full version detail.
	Data VoiceDesignNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A voice design object with full version detail.
type VoiceDesignNewResponseData struct {
	// Unique identifier for the voice design.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the voice design was first created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of the voice design.
	Name string `json:"name"`
	// Natural language prompt used to define the voice style for this version.
	Prompt string `json:"prompt"`
	// Voice synthesis provider used for this design.
	//
	// Any of "telnyx", "minimax", "Telnyx", "Minimax".
	Provider string `json:"provider" api:"nullable"`
	// List of TTS model identifiers supported by this design's provider (e.g.
	// `Qwen3TTS`, `speech-02-turbo`).
	ProviderSupportedModels []string `json:"provider_supported_models"`
	// Provider-specific voice identifier. For Telnyx designs this is the design
	// version ID; for Minimax it is the Minimax-assigned voice ID.
	ProviderVoiceID string `json:"provider_voice_id" api:"nullable"`
	// Identifies the resource type.
	//
	// Any of "voice_design".
	RecordType string `json:"record_type"`
	// Sample text used to synthesize this version.
	Text string `json:"text"`
	// Timestamp when the voice design was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Version number of this voice design.
	Version int64 `json:"version"`
	// Timestamp when this specific version was created.
	VersionCreatedAt time.Time `json:"version_created_at" format:"date-time"`
	// Size of the voice sample audio in bytes.
	VoiceSampleSize int64 `json:"voice_sample_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Prompt                  respjson.Field
		Provider                respjson.Field
		ProviderSupportedModels respjson.Field
		ProviderVoiceID         respjson.Field
		RecordType              respjson.Field
		Text                    respjson.Field
		UpdatedAt               respjson.Field
		Version                 respjson.Field
		VersionCreatedAt        respjson.Field
		VoiceSampleSize         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response envelope for a single voice design with full version detail.
type VoiceDesignGetResponse struct {
	// A voice design object with full version detail.
	Data VoiceDesignGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A voice design object with full version detail.
type VoiceDesignGetResponseData struct {
	// Unique identifier for the voice design.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the voice design was first created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of the voice design.
	Name string `json:"name"`
	// Natural language prompt used to define the voice style for this version.
	Prompt string `json:"prompt"`
	// Voice synthesis provider used for this design.
	//
	// Any of "telnyx", "minimax", "Telnyx", "Minimax".
	Provider string `json:"provider" api:"nullable"`
	// List of TTS model identifiers supported by this design's provider (e.g.
	// `Qwen3TTS`, `speech-02-turbo`).
	ProviderSupportedModels []string `json:"provider_supported_models"`
	// Provider-specific voice identifier. For Telnyx designs this is the design
	// version ID; for Minimax it is the Minimax-assigned voice ID.
	ProviderVoiceID string `json:"provider_voice_id" api:"nullable"`
	// Identifies the resource type.
	//
	// Any of "voice_design".
	RecordType string `json:"record_type"`
	// Sample text used to synthesize this version.
	Text string `json:"text"`
	// Timestamp when the voice design was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Version number of this voice design.
	Version int64 `json:"version"`
	// Timestamp when this specific version was created.
	VersionCreatedAt time.Time `json:"version_created_at" format:"date-time"`
	// Size of the voice sample audio in bytes.
	VoiceSampleSize int64 `json:"voice_sample_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Prompt                  respjson.Field
		Provider                respjson.Field
		ProviderSupportedModels respjson.Field
		ProviderVoiceID         respjson.Field
		RecordType              respjson.Field
		Text                    respjson.Field
		UpdatedAt               respjson.Field
		Version                 respjson.Field
		VersionCreatedAt        respjson.Field
		VoiceSampleSize         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summarized voice design object (without version-specific fields).
type VoiceDesignListResponse struct {
	// Unique identifier for the voice design.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the voice design was first created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of the voice design.
	Name string `json:"name"`
	// Voice synthesis provider used for this design.
	//
	// Any of "telnyx", "minimax", "Telnyx", "Minimax".
	Provider VoiceDesignListResponseProvider `json:"provider" api:"nullable"`
	// List of TTS model identifiers supported by this design's provider.
	ProviderSupportedModels []string `json:"provider_supported_models"`
	// Identifies the resource type.
	//
	// Any of "voice_design".
	RecordType VoiceDesignListResponseRecordType `json:"record_type"`
	// Timestamp when the voice design was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Provider                respjson.Field
		ProviderSupportedModels respjson.Field
		RecordType              respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignListResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Voice synthesis provider used for this design.
type VoiceDesignListResponseProvider string

const (
	VoiceDesignListResponseProviderTelnyx           VoiceDesignListResponseProvider = "telnyx"
	VoiceDesignListResponseProviderMinimax          VoiceDesignListResponseProvider = "minimax"
	VoiceDesignListResponseProviderTelnyxMixedCase  VoiceDesignListResponseProvider = "Telnyx"
	VoiceDesignListResponseProviderMinimaxMixedCase VoiceDesignListResponseProvider = "Minimax"
)

// Identifies the resource type.
type VoiceDesignListResponseRecordType string

const (
	VoiceDesignListResponseRecordTypeVoiceDesign VoiceDesignListResponseRecordType = "voice_design"
)

// Response envelope for a voice design after a rename operation (no
// version-specific fields).
type VoiceDesignRenameResponse struct {
	// A summarized voice design object (without version-specific fields).
	Data VoiceDesignRenameResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignRenameResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignRenameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summarized voice design object (without version-specific fields).
type VoiceDesignRenameResponseData struct {
	// Unique identifier for the voice design.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the voice design was first created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of the voice design.
	Name string `json:"name"`
	// Voice synthesis provider used for this design.
	//
	// Any of "telnyx", "minimax", "Telnyx", "Minimax".
	Provider string `json:"provider" api:"nullable"`
	// List of TTS model identifiers supported by this design's provider.
	ProviderSupportedModels []string `json:"provider_supported_models"`
	// Identifies the resource type.
	//
	// Any of "voice_design".
	RecordType string `json:"record_type"`
	// Timestamp when the voice design was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Provider                respjson.Field
		ProviderSupportedModels respjson.Field
		RecordType              respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceDesignRenameResponseData) RawJSON() string { return r.JSON.raw }
func (r *VoiceDesignRenameResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceDesignNewParams struct {
	// Natural language description of the voice style, e.g. 'Speak in a warm, friendly
	// tone with a slight British accent'.
	Prompt string `json:"prompt" api:"required"`
	// Sample text to synthesize for this voice design.
	Text string `json:"text" api:"required"`
	// Language for synthesis. Supported values: Auto, Chinese, English, Japanese,
	// Korean, German, French, Russian, Portuguese, Spanish, Italian. Defaults to Auto.
	Language param.Opt[string] `json:"language,omitzero"`
	// Maximum number of tokens to generate. Default: 2048.
	MaxNewTokens param.Opt[int64] `json:"max_new_tokens,omitzero"`
	// Name for the voice design. Required when creating a new design
	// (`voice_design_id` is not provided); ignored when adding a version. Cannot be a
	// UUID.
	Name param.Opt[string] `json:"name,omitzero"`
	// Repetition penalty to reduce repeated patterns in generated audio. Default:
	// 1.05.
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	// Sampling temperature controlling randomness. Higher values produce more varied
	// output. Default: 0.9.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Top-k sampling parameter — limits the token vocabulary considered at each step.
	// Default: 50.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Top-p (nucleus) sampling parameter — cumulative probability cutoff for token
	// selection. Default: 1.0.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// ID of an existing voice design to add a new version to. When provided, a new
	// version is created instead of a new design.
	VoiceDesignID param.Opt[string] `json:"voice_design_id,omitzero" format:"uuid"`
	// Voice synthesis provider. `telnyx` uses the Qwen3TTS model; `minimax` uses the
	// Minimax speech models. Case-insensitive. Defaults to `telnyx`.
	//
	// Any of "telnyx", "minimax", "Telnyx", "Minimax".
	Provider VoiceDesignNewParamsProvider `json:"provider,omitzero"`
	paramObj
}

func (r VoiceDesignNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceDesignNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceDesignNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Voice synthesis provider. `telnyx` uses the Qwen3TTS model; `minimax` uses the
// Minimax speech models. Case-insensitive. Defaults to `telnyx`.
type VoiceDesignNewParamsProvider string

const (
	VoiceDesignNewParamsProviderTelnyx           VoiceDesignNewParamsProvider = "telnyx"
	VoiceDesignNewParamsProviderMinimax          VoiceDesignNewParamsProvider = "minimax"
	VoiceDesignNewParamsProviderTelnyxMixedCase  VoiceDesignNewParamsProvider = "Telnyx"
	VoiceDesignNewParamsProviderMinimaxMixedCase VoiceDesignNewParamsProvider = "Minimax"
)

type VoiceDesignGetParams struct {
	// Specific version number to retrieve. Defaults to the latest version.
	Version param.Opt[int64] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceDesignGetParams]'s query parameters as `url.Values`.
func (r VoiceDesignGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoiceDesignListParams struct {
	// Case-insensitive substring filter on the name field.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Page number for pagination (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Sort order. Prefix with `-` for descending. Defaults to `-created_at`.
	//
	// Any of "name", "-name", "created_at", "-created_at".
	Sort VoiceDesignListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceDesignListParams]'s query parameters as `url.Values`.
func (r VoiceDesignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Prefix with `-` for descending. Defaults to `-created_at`.
type VoiceDesignListParamsSort string

const (
	VoiceDesignListParamsSortName          VoiceDesignListParamsSort = "name"
	VoiceDesignListParamsSortNameDesc      VoiceDesignListParamsSort = "-name"
	VoiceDesignListParamsSortCreatedAt     VoiceDesignListParamsSort = "created_at"
	VoiceDesignListParamsSortCreatedAtDesc VoiceDesignListParamsSort = "-created_at"
)

type VoiceDesignDeleteVersionParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type VoiceDesignDownloadSampleParams struct {
	// Specific version number to download the sample for. Defaults to the latest
	// version.
	Version param.Opt[int64] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceDesignDownloadSampleParams]'s query parameters as
// `url.Values`.
func (r VoiceDesignDownloadSampleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoiceDesignRenameParams struct {
	// New name for the voice design.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r VoiceDesignRenameParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceDesignRenameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceDesignRenameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
