// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Call Recordings operations.
//
// RecordingTranscriptionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordingTranscriptionService] method instead.
type RecordingTranscriptionService struct {
	Options []option.RequestOption
}

// NewRecordingTranscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRecordingTranscriptionService(opts ...option.RequestOption) (r RecordingTranscriptionService) {
	r = RecordingTranscriptionService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing recording transcription.
func (r *RecordingTranscriptionService) Get(ctx context.Context, recordingTranscriptionID string, opts ...option.RequestOption) (res *RecordingTranscriptionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordingTranscriptionID == "" {
		err = errors.New("missing required recording_transcription_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("recording_transcriptions/%s", recordingTranscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a list of your recording transcriptions.
func (r *RecordingTranscriptionService) List(ctx context.Context, query RecordingTranscriptionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[RecordingTranscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "recording_transcriptions"
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

// Returns a list of your recording transcriptions.
func (r *RecordingTranscriptionService) ListAutoPaging(ctx context.Context, query RecordingTranscriptionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[RecordingTranscription] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a recording transcription.
func (r *RecordingTranscriptionService) Delete(ctx context.Context, recordingTranscriptionID string, opts ...option.RequestOption) (res *RecordingTranscriptionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordingTranscriptionID == "" {
		err = errors.New("missing required recording_transcription_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("recording_transcriptions/%s", recordingTranscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type RecordingTranscription struct {
	// Uniquely identifies the recording transcription.
	ID string `json:"id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The duration of the recording transcription in milliseconds.
	DurationMillis int64 `json:"duration_millis"`
	// Any of "recording_transcription".
	RecordType RecordingTranscriptionRecordType `json:"record_type"`
	// Uniquely identifies the recording associated with this transcription.
	RecordingID string `json:"recording_id"`
	// The status of the recording transcription. Only `completed` has transcription
	// text available.
	//
	// Any of "in-progress", "completed".
	Status RecordingTranscriptionStatus `json:"status"`
	// The recording's transcribed text.
	TranscriptionText string `json:"transcription_text"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		DurationMillis    respjson.Field
		RecordType        respjson.Field
		RecordingID       respjson.Field
		Status            respjson.Field
		TranscriptionText respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscription) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionRecordType string

const (
	RecordingTranscriptionRecordTypeRecordingTranscription RecordingTranscriptionRecordType = "recording_transcription"
)

// The status of the recording transcription. Only `completed` has transcription
// text available.
type RecordingTranscriptionStatus string

const (
	RecordingTranscriptionStatusInProgress RecordingTranscriptionStatus = "in-progress"
	RecordingTranscriptionStatusCompleted  RecordingTranscriptionStatus = "completed"
)

type RecordingTranscriptionGetResponse struct {
	Data RecordingTranscription `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionDeleteResponse struct {
	Data RecordingTranscription `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter recording transcriptions by various attributes.
	Filter RecordingTranscriptionListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingTranscriptionListParams]'s query parameters as
// `url.Values`.
func (r RecordingTranscriptionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter recording transcriptions by various attributes.
type RecordingTranscriptionListParamsFilter struct {
	// If present, transcriptions will be filtered to those associated with the given
	// recording.
	RecordingID param.Opt[string]                               `query:"recording_id,omitzero" format:"uuid" json:"-"`
	CreatedAt   RecordingTranscriptionListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingTranscriptionListParamsFilter]'s query parameters
// as `url.Values`.
func (r RecordingTranscriptionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RecordingTranscriptionListParamsFilterCreatedAt struct {
	// Returns only transcriptions created later than or at given ISO 8601 datetime.
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Returns only transcriptions created earlier than or at given ISO 8601 datetime.
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingTranscriptionListParamsFilterCreatedAt]'s query
// parameters as `url.Values`.
func (r RecordingTranscriptionListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
