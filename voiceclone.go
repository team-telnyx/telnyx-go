// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Capture and manage voice identities as clones for use in text-to-speech
// synthesis.
//
// VoiceCloneService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceCloneService] method instead.
type VoiceCloneService struct {
	Options []option.RequestOption
}

// NewVoiceCloneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceCloneService(opts ...option.RequestOption) (r VoiceCloneService) {
	r = VoiceCloneService{}
	r.Options = opts
	return
}

// Creates a new voice clone by capturing the voice identity of an existing voice
// design. The clone can then be used for text-to-speech synthesis.
func (r *VoiceCloneService) New(ctx context.Context, body VoiceCloneNewParams, opts ...option.RequestOption) (res *VoiceCloneNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice_clones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the name, language, or gender of a voice clone.
func (r *VoiceCloneService) Update(ctx context.Context, id string, body VoiceCloneUpdateParams, opts ...option.RequestOption) (res *VoiceCloneUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("voice_clones/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of voice clones belonging to the authenticated account.
func (r *VoiceCloneService) List(ctx context.Context, query VoiceCloneListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[VoiceCloneData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "voice_clones"
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

// Returns a paginated list of voice clones belonging to the authenticated account.
func (r *VoiceCloneService) ListAutoPaging(ctx context.Context, query VoiceCloneListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[VoiceCloneData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a voice clone. This action cannot be undone.
func (r *VoiceCloneService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("voice_clones/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates a new voice clone by uploading an audio file directly. Supported
// formats: WAV, MP3, FLAC, OGG, M4A. For best results, provide 5–10 seconds of
// clear speech. Maximum file size: 5MB for Telnyx, 20MB for Minimax.
func (r *VoiceCloneService) NewFromUpload(ctx context.Context, body VoiceCloneNewFromUploadParams, opts ...option.RequestOption) (res *VoiceCloneNewFromUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice_clones/from_upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Downloads the WAV audio sample that was used to create the voice clone.
func (r *VoiceCloneService) DownloadSample(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/wav")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("voice_clones/%s/sample", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A voice clone object.
type VoiceCloneData struct {
	// Unique identifier for the voice clone.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the voice clone was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Gender of the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender VoiceCloneDataGender `json:"gender" api:"nullable"`
	// Voice style description. If not explicitly set on upload, falls back to the
	// source design's prompt text.
	Label string `json:"label" api:"nullable"`
	// ISO 639-1 language code of the voice clone.
	Language string `json:"language" api:"nullable"`
	// TTS model identifier for the voice clone.
	//
	// Any of "Qwen3TTS", "Ultra", "speech-2.8-turbo".
	ModelID VoiceCloneDataModelID `json:"model_id"`
	// Name of the voice clone.
	Name string `json:"name"`
	// Voice synthesis provider used for this clone.
	//
	// Any of "telnyx", "minimax".
	Provider VoiceCloneDataProvider `json:"provider"`
	// List of TTS model identifiers supported by this clone's provider.
	ProviderSupportedModels []string `json:"provider_supported_models"`
	// Provider-specific voice identifier used for TTS synthesis. May differ from the
	// clone UUID depending on the provider and model.
	ProviderVoiceID string `json:"provider_voice_id" api:"nullable"`
	// Identifies the resource type.
	//
	// Any of "voice_clone".
	RecordType VoiceCloneDataRecordType `json:"record_type"`
	// UUID of the source voice design. `null` for upload-based clones.
	SourceVoiceDesignID string `json:"source_voice_design_id" api:"nullable" format:"uuid"`
	// Version of the source voice design used. `null` for upload-based clones.
	SourceVoiceDesignVersion int64 `json:"source_voice_design_version" api:"nullable"`
	// Clone status. pending for Ultra clones while on-prem import is in progress,
	// active once ready, failed if verification timed out, expired if not kept alive.
	//
	// Any of "active", "pending", "failed", "expired".
	Status VoiceCloneDataStatus `json:"status"`
	// Timestamp when the voice clone was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Gender                   respjson.Field
		Label                    respjson.Field
		Language                 respjson.Field
		ModelID                  respjson.Field
		Name                     respjson.Field
		Provider                 respjson.Field
		ProviderSupportedModels  respjson.Field
		ProviderVoiceID          respjson.Field
		RecordType               respjson.Field
		SourceVoiceDesignID      respjson.Field
		SourceVoiceDesignVersion respjson.Field
		Status                   respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCloneData) RawJSON() string { return r.JSON.raw }
func (r *VoiceCloneData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gender of the voice clone.
type VoiceCloneDataGender string

const (
	VoiceCloneDataGenderMale    VoiceCloneDataGender = "male"
	VoiceCloneDataGenderFemale  VoiceCloneDataGender = "female"
	VoiceCloneDataGenderNeutral VoiceCloneDataGender = "neutral"
)

// TTS model identifier for the voice clone.
type VoiceCloneDataModelID string

const (
	VoiceCloneDataModelIDQwen3Tts       VoiceCloneDataModelID = "Qwen3TTS"
	VoiceCloneDataModelIDUltra          VoiceCloneDataModelID = "Ultra"
	VoiceCloneDataModelIDSpeech2_8Turbo VoiceCloneDataModelID = "speech-2.8-turbo"
)

// Voice synthesis provider used for this clone.
type VoiceCloneDataProvider string

const (
	VoiceCloneDataProviderTelnyx  VoiceCloneDataProvider = "telnyx"
	VoiceCloneDataProviderMinimax VoiceCloneDataProvider = "minimax"
)

// Identifies the resource type.
type VoiceCloneDataRecordType string

const (
	VoiceCloneDataRecordTypeVoiceClone VoiceCloneDataRecordType = "voice_clone"
)

// Clone status. pending for Ultra clones while on-prem import is in progress,
// active once ready, failed if verification timed out, expired if not kept alive.
type VoiceCloneDataStatus string

const (
	VoiceCloneDataStatusActive  VoiceCloneDataStatus = "active"
	VoiceCloneDataStatusPending VoiceCloneDataStatus = "pending"
	VoiceCloneDataStatusFailed  VoiceCloneDataStatus = "failed"
	VoiceCloneDataStatusExpired VoiceCloneDataStatus = "expired"
)

// Response envelope for a single voice clone.
type VoiceCloneNewResponse struct {
	// A voice clone object.
	Data VoiceCloneData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCloneNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceCloneNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response envelope for a single voice clone.
type VoiceCloneUpdateResponse struct {
	// A voice clone object.
	Data VoiceCloneData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCloneUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceCloneUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response envelope for a single voice clone.
type VoiceCloneNewFromUploadResponse struct {
	// A voice clone object.
	Data VoiceCloneData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCloneNewFromUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceCloneNewFromUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceCloneNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Create
	// a voice clone from a voice design using the Telnyx provider.
	OfTelnyxDesignClone *VoiceCloneNewParamsBodyTelnyxDesignClone `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// a voice clone from a voice design using the Minimax provider.
	OfMinimaxDesignClone *VoiceCloneNewParamsBodyMinimaxDesignClone `json:",inline"`

	paramObj
}

func (u VoiceCloneNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTelnyxDesignClone, u.OfMinimaxDesignClone)
}
func (r *VoiceCloneNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create a voice clone from a voice design using the Telnyx provider.
//
// The properties Gender, Language, Name, VoiceDesignID are required.
type VoiceCloneNewParamsBodyTelnyxDesignClone struct {
	// Gender of the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender string `json:"gender,omitzero" api:"required"`
	// ISO 639-1 language code for the clone. Supports the combined Telnyx language
	// set.
	Language string `json:"language" api:"required"`
	// Name for the voice clone.
	Name string `json:"name" api:"required"`
	// UUID of the source voice design to clone.
	VoiceDesignID string `json:"voice_design_id" api:"required" format:"uuid"`
	// Voice synthesis provider. Defaults to `telnyx`.
	//
	// Any of "telnyx", "Telnyx".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r VoiceCloneNewParamsBodyTelnyxDesignClone) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCloneNewParamsBodyTelnyxDesignClone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCloneNewParamsBodyTelnyxDesignClone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceCloneNewParamsBodyTelnyxDesignClone](
		"gender", "male", "female", "neutral",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewParamsBodyTelnyxDesignClone](
		"provider", "telnyx", "Telnyx",
	)
}

// Create a voice clone from a voice design using the Minimax provider.
//
// The properties Gender, Language, Name, Provider, VoiceDesignID are required.
type VoiceCloneNewParamsBodyMinimaxDesignClone struct {
	// Gender of the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender string `json:"gender,omitzero" api:"required"`
	// ISO 639-1 language code for the clone. Supports the Minimax language set.
	Language string `json:"language" api:"required"`
	// Name for the voice clone.
	Name string `json:"name" api:"required"`
	// Voice synthesis provider. Must be `minimax`.
	//
	// Any of "minimax", "Minimax".
	Provider string `json:"provider,omitzero" api:"required"`
	// UUID of the source voice design to clone.
	VoiceDesignID string `json:"voice_design_id" api:"required" format:"uuid"`
	paramObj
}

func (r VoiceCloneNewParamsBodyMinimaxDesignClone) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCloneNewParamsBodyMinimaxDesignClone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCloneNewParamsBodyMinimaxDesignClone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceCloneNewParamsBodyMinimaxDesignClone](
		"gender", "male", "female", "neutral",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewParamsBodyMinimaxDesignClone](
		"provider", "minimax", "Minimax",
	)
}

type VoiceCloneUpdateParams struct {
	// New name for the voice clone.
	Name string `json:"name" api:"required"`
	// Updated ISO 639-1 language code or `auto`.
	Language param.Opt[string] `json:"language,omitzero"`
	// Updated gender for the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender VoiceCloneUpdateParamsGender `json:"gender,omitzero"`
	paramObj
}

func (r VoiceCloneUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCloneUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCloneUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Updated gender for the voice clone.
type VoiceCloneUpdateParamsGender string

const (
	VoiceCloneUpdateParamsGenderMale    VoiceCloneUpdateParamsGender = "male"
	VoiceCloneUpdateParamsGenderFemale  VoiceCloneUpdateParamsGender = "female"
	VoiceCloneUpdateParamsGenderNeutral VoiceCloneUpdateParamsGender = "neutral"
)

type VoiceCloneListParams struct {
	// Case-insensitive substring filter on the name field.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Page number for pagination (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by voice synthesis provider. Case-insensitive.
	//
	// Any of "telnyx", "minimax".
	FilterProvider VoiceCloneListParamsFilterProvider `query:"filter[provider],omitzero" json:"-"`
	// Sort order. Prefix with `-` for descending. Defaults to `-created_at`.
	//
	// Any of "name", "-name", "created_at", "-created_at".
	Sort VoiceCloneListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceCloneListParams]'s query parameters as `url.Values`.
func (r VoiceCloneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by voice synthesis provider. Case-insensitive.
type VoiceCloneListParamsFilterProvider string

const (
	VoiceCloneListParamsFilterProviderTelnyx  VoiceCloneListParamsFilterProvider = "telnyx"
	VoiceCloneListParamsFilterProviderMinimax VoiceCloneListParamsFilterProvider = "minimax"
)

// Sort order. Prefix with `-` for descending. Defaults to `-created_at`.
type VoiceCloneListParamsSort string

const (
	VoiceCloneListParamsSortName          VoiceCloneListParamsSort = "name"
	VoiceCloneListParamsSortNameDesc      VoiceCloneListParamsSort = "-name"
	VoiceCloneListParamsSortCreatedAt     VoiceCloneListParamsSort = "created_at"
	VoiceCloneListParamsSortCreatedAtDesc VoiceCloneListParamsSort = "-created_at"
)

type VoiceCloneNewFromUploadParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Upload-based voice clone using the Telnyx Qwen3TTS model (default).
	OfTelnyxQwen3TtsClone *VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Upload-based voice clone using the Telnyx Ultra model.
	OfTelnyxUltraClone *VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Upload-based voice clone using the Minimax provider.
	OfMinimaxClone *VoiceCloneNewFromUploadParamsBodyMinimaxClone `json:",inline"`

	paramObj
}

func (r VoiceCloneNewFromUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.Body, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Upload-based voice clone using the Telnyx Qwen3TTS model (default).
//
// The properties AudioFile, Gender, Language, Name, Provider are required.
type VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone struct {
	// Audio file to clone the voice from. Supported formats: WAV, MP3, FLAC, OGG, M4A.
	// For best quality, provide 5–10 seconds of clear, uninterrupted speech. Maximum
	// size: 5MB.
	AudioFile io.Reader `json:"audio_file,omitzero" api:"required" format:"binary"`
	// Gender of the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender string `json:"gender,omitzero" api:"required"`
	// ISO 639-1 language code from the Qwen language set.
	Language string `json:"language" api:"required"`
	// Name for the voice clone.
	Name string `json:"name" api:"required"`
	// Voice synthesis provider. Must be `telnyx`.
	//
	// Any of "telnyx", "Telnyx".
	Provider string `json:"provider,omitzero" api:"required"`
	// Optional custom label describing the voice style.
	Label param.Opt[string] `json:"label,omitzero"`
	// Optional transcript of the audio file. Providing this improves clone quality.
	RefText param.Opt[string] `json:"ref_text,omitzero"`
	// TTS model identifier. Nullable/omittable — defaults to Qwen3TTS.
	//
	// Any of "Qwen3TTS".
	ModelID string `json:"model_id,omitzero"`
	paramObj
}

func (r VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone](
		"gender", "male", "female", "neutral",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone](
		"provider", "telnyx", "Telnyx",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyTelnyxQwen3TtsClone](
		"model_id", "Qwen3TTS",
	)
}

// Upload-based voice clone using the Telnyx Ultra model.
//
// The properties AudioFile, Gender, Language, ModelID, Name, Provider are
// required.
type VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone struct {
	// Audio file to clone the voice from. Supported formats: WAV, MP3, FLAC, OGG, M4A.
	// For best quality, provide 5–10 seconds of clear, uninterrupted speech. Maximum
	// size: 5MB.
	AudioFile io.Reader `json:"audio_file,omitzero" api:"required" format:"binary"`
	// Gender of the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender string `json:"gender,omitzero" api:"required"`
	// ISO 639-1 language code from the Ultra language set (40 languages).
	Language string `json:"language" api:"required"`
	// TTS model identifier. Must be `Ultra`.
	//
	// Any of "Ultra".
	ModelID string `json:"model_id,omitzero" api:"required"`
	// Name for the voice clone.
	Name string `json:"name" api:"required"`
	// Voice synthesis provider. Must be `telnyx`.
	//
	// Any of "telnyx", "Telnyx".
	Provider string `json:"provider,omitzero" api:"required"`
	// Optional custom label describing the voice style.
	Label param.Opt[string] `json:"label,omitzero"`
	// Optional transcript of the audio file. Providing this improves clone quality.
	RefText param.Opt[string] `json:"ref_text,omitzero"`
	paramObj
}

func (r VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone](
		"gender", "male", "female", "neutral",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone](
		"model_id", "Ultra",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyTelnyxUltraClone](
		"provider", "telnyx", "Telnyx",
	)
}

// Upload-based voice clone using the Minimax provider.
//
// The properties AudioFile, Gender, Language, Name, Provider are required.
type VoiceCloneNewFromUploadParamsBodyMinimaxClone struct {
	// Audio file to clone the voice from. Supported formats: WAV, MP3, FLAC, OGG, M4A.
	// For best quality, provide 5–10 seconds of clear, uninterrupted speech. Maximum
	// size: 20MB.
	AudioFile io.Reader `json:"audio_file,omitzero" api:"required" format:"binary"`
	// Gender of the voice clone.
	//
	// Any of "male", "female", "neutral".
	Gender string `json:"gender,omitzero" api:"required"`
	// ISO 639-1 language code from the Minimax language set.
	Language string `json:"language" api:"required"`
	// Name for the voice clone.
	Name string `json:"name" api:"required"`
	// Voice synthesis provider. Must be `minimax`.
	//
	// Any of "minimax", "Minimax".
	Provider string `json:"provider,omitzero" api:"required"`
	// Optional custom label describing the voice style.
	Label param.Opt[string] `json:"label,omitzero"`
	// Optional transcript of the audio file. Providing this improves clone quality.
	RefText param.Opt[string] `json:"ref_text,omitzero"`
	// TTS model identifier. Nullable — defaults to speech-2.8-turbo.
	//
	// Any of "speech-2.8-turbo".
	ModelID string `json:"model_id,omitzero"`
	paramObj
}

func (r VoiceCloneNewFromUploadParamsBodyMinimaxClone) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCloneNewFromUploadParamsBodyMinimaxClone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCloneNewFromUploadParamsBodyMinimaxClone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyMinimaxClone](
		"gender", "male", "female", "neutral",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyMinimaxClone](
		"provider", "minimax", "Minimax",
	)
	apijson.RegisterFieldValidator[VoiceCloneNewFromUploadParamsBodyMinimaxClone](
		"model_id", "speech-2.8-turbo",
	)
}
