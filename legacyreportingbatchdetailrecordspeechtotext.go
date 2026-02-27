// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Speech to text batch detail records
//
// LegacyReportingBatchDetailRecordSpeechToTextService contains methods and other
// services that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingBatchDetailRecordSpeechToTextService] method instead.
type LegacyReportingBatchDetailRecordSpeechToTextService struct {
	Options []option.RequestOption
}

// NewLegacyReportingBatchDetailRecordSpeechToTextService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewLegacyReportingBatchDetailRecordSpeechToTextService(opts ...option.RequestOption) (r LegacyReportingBatchDetailRecordSpeechToTextService) {
	r = LegacyReportingBatchDetailRecordSpeechToTextService{}
	r.Options = opts
	return
}

// Creates a new Speech to Text batch report request with the specified filters
func (r *LegacyReportingBatchDetailRecordSpeechToTextService) New(ctx context.Context, body LegacyReportingBatchDetailRecordSpeechToTextNewParams, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordSpeechToTextNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/speech_to_text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific Speech to Text batch report request by ID
func (r *LegacyReportingBatchDetailRecordSpeechToTextService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordSpeechToTextGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/batch_detail_records/speech_to_text/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all Speech to Text batch report requests for the authenticated user
func (r *LegacyReportingBatchDetailRecordSpeechToTextService) List(ctx context.Context, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordSpeechToTextListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/speech_to_text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific Speech to Text batch report request by ID
func (r *LegacyReportingBatchDetailRecordSpeechToTextService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordSpeechToTextDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/batch_detail_records/speech_to_text/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SttDetailReportResponse struct {
	// Identifies the resource
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// URL to download the report
	DownloadLink string    `json:"download_link"`
	EndDate      time.Time `json:"end_date" format:"date-time"`
	RecordType   string    `json:"record_type"`
	StartDate    time.Time `json:"start_date" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status SttDetailReportResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		DownloadLink respjson.Field
		EndDate      respjson.Field
		RecordType   respjson.Field
		StartDate    respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SttDetailReportResponse) RawJSON() string { return r.JSON.raw }
func (r *SttDetailReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SttDetailReportResponseStatus string

const (
	SttDetailReportResponseStatusPending  SttDetailReportResponseStatus = "PENDING"
	SttDetailReportResponseStatusComplete SttDetailReportResponseStatus = "COMPLETE"
	SttDetailReportResponseStatusFailed   SttDetailReportResponseStatus = "FAILED"
	SttDetailReportResponseStatusExpired  SttDetailReportResponseStatus = "EXPIRED"
)

type LegacyReportingBatchDetailRecordSpeechToTextNewResponse struct {
	Data SttDetailReportResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordSpeechToTextNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordSpeechToTextNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordSpeechToTextGetResponse struct {
	Data SttDetailReportResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordSpeechToTextGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordSpeechToTextGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordSpeechToTextListResponse struct {
	Data []SttDetailReportResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordSpeechToTextListResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordSpeechToTextListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordSpeechToTextDeleteResponse struct {
	Data SttDetailReportResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordSpeechToTextDeleteResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordSpeechToTextDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordSpeechToTextNewParams struct {
	// End date in ISO format with timezone (date range must be up to one month)
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// Start date in ISO format with timezone
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	paramObj
}

func (r LegacyReportingBatchDetailRecordSpeechToTextNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingBatchDetailRecordSpeechToTextNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingBatchDetailRecordSpeechToTextNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
