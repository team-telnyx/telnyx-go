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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// TexmlAccountService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountService] method instead.
type TexmlAccountService struct {
	Options        []option.RequestOption
	Calls          TexmlAccountCallService
	Conferences    TexmlAccountConferenceService
	Recordings     TexmlAccountRecordingService
	Transcriptions TexmlAccountTranscriptionService
}

// NewTexmlAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTexmlAccountService(opts ...option.RequestOption) (r TexmlAccountService) {
	r = TexmlAccountService{}
	r.Options = opts
	r.Calls = NewTexmlAccountCallService(opts...)
	r.Conferences = NewTexmlAccountConferenceService(opts...)
	r.Recordings = NewTexmlAccountRecordingService(opts...)
	r.Transcriptions = NewTexmlAccountTranscriptionService(opts...)
	return
}

// Returns multiple recording resources for an account.
func (r *TexmlAccountService) GetRecordingsJson(ctx context.Context, accountSid string, query TexmlAccountGetRecordingsJsonParams, opts ...option.RequestOption) (res *TexmlAccountGetRecordingsJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Recordings.json", accountSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns multiple recording transcription resources for an account.
func (r *TexmlAccountService) GetTranscriptionsJson(ctx context.Context, accountSid string, query TexmlAccountGetTranscriptionsJsonParams, opts ...option.RequestOption) (res *TexmlAccountGetTranscriptionsJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Transcriptions.json", accountSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TexmlGetCallRecordingResponseBody struct {
	AccountSid string `json:"account_sid"`
	CallSid    string `json:"call_sid"`
	// Any of 1, 2.
	Channels      int64     `json:"channels"`
	ConferenceSid string    `json:"conference_sid,nullable" format:"uuid"`
	DateCreated   time.Time `json:"date_created" format:"date-time"`
	DateUpdated   time.Time `json:"date_updated" format:"date-time"`
	// The duration of this recording, given in seconds.
	Duration  string `json:"duration,nullable"`
	ErrorCode string `json:"error_code,nullable"`
	MediaURL  string `json:"media_url" format:"uri"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// Defines how the recording was created.
	//
	// Any of "StartCallRecordingAPI", "StartConferenceRecordingAPI", "OutboundAPI",
	// "DialVerb", "Conference", "RecordVerb", "Trunking".
	Source    TexmlGetCallRecordingResponseBodySource `json:"source"`
	StartTime time.Time                               `json:"start_time" format:"date-time"`
	// Any of "in-progress", "completed", "paused", "stopped".
	Status TexmlGetCallRecordingResponseBodyStatus `json:"status"`
	// Subresources details for a recording if available.
	SubresourcesUris TexmlRecordingSubresourcesUris `json:"subresources_uris"`
	// The relative URI for this recording resource.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid       respjson.Field
		CallSid          respjson.Field
		Channels         respjson.Field
		ConferenceSid    respjson.Field
		DateCreated      respjson.Field
		DateUpdated      respjson.Field
		Duration         respjson.Field
		ErrorCode        respjson.Field
		MediaURL         respjson.Field
		Sid              respjson.Field
		Source           respjson.Field
		StartTime        respjson.Field
		Status           respjson.Field
		SubresourcesUris respjson.Field
		Uri              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlGetCallRecordingResponseBody) RawJSON() string { return r.JSON.raw }
func (r *TexmlGetCallRecordingResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines how the recording was created.
type TexmlGetCallRecordingResponseBodySource string

const (
	TexmlGetCallRecordingResponseBodySourceStartCallRecordingAPI       TexmlGetCallRecordingResponseBodySource = "StartCallRecordingAPI"
	TexmlGetCallRecordingResponseBodySourceStartConferenceRecordingAPI TexmlGetCallRecordingResponseBodySource = "StartConferenceRecordingAPI"
	TexmlGetCallRecordingResponseBodySourceOutboundAPI                 TexmlGetCallRecordingResponseBodySource = "OutboundAPI"
	TexmlGetCallRecordingResponseBodySourceDialVerb                    TexmlGetCallRecordingResponseBodySource = "DialVerb"
	TexmlGetCallRecordingResponseBodySourceConference                  TexmlGetCallRecordingResponseBodySource = "Conference"
	TexmlGetCallRecordingResponseBodySourceRecordVerb                  TexmlGetCallRecordingResponseBodySource = "RecordVerb"
	TexmlGetCallRecordingResponseBodySourceTrunking                    TexmlGetCallRecordingResponseBodySource = "Trunking"
)

type TexmlGetCallRecordingResponseBodyStatus string

const (
	TexmlGetCallRecordingResponseBodyStatusInProgress TexmlGetCallRecordingResponseBodyStatus = "in-progress"
	TexmlGetCallRecordingResponseBodyStatusCompleted  TexmlGetCallRecordingResponseBodyStatus = "completed"
	TexmlGetCallRecordingResponseBodyStatusPaused     TexmlGetCallRecordingResponseBodyStatus = "paused"
	TexmlGetCallRecordingResponseBodyStatusStopped    TexmlGetCallRecordingResponseBodyStatus = "stopped"
)

// Subresources details for a recording if available.
type TexmlRecordingSubresourcesUris struct {
	Transcriptions string `json:"transcriptions,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transcriptions respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlRecordingSubresourcesUris) RawJSON() string { return r.JSON.raw }
func (r *TexmlRecordingSubresourcesUris) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountGetRecordingsJsonResponse struct {
	// The number of the last element on the page, zero-indexed.
	End int64 `json:"end"`
	// Relative uri to the first page of the query results
	FirstPageUri string `json:"first_page_uri" format:"uri"`
	// Relative uri to the next page of the query results
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize int64 `json:"page_size"`
	// Relative uri to the previous page of the query results
	PreviousPageUri string                              `json:"previous_page_uri" format:"uri"`
	Recordings      []TexmlGetCallRecordingResponseBody `json:"recordings"`
	// The number of the first element on the page, zero-indexed.
	Start int64 `json:"start"`
	// The URI of the current page.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End             respjson.Field
		FirstPageUri    respjson.Field
		NextPageUri     respjson.Field
		Page            respjson.Field
		PageSize        respjson.Field
		PreviousPageUri respjson.Field
		Recordings      respjson.Field
		Start           respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountGetRecordingsJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountGetRecordingsJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountGetTranscriptionsJsonResponse struct {
	// The number of the last element on the page, zero-indexed
	End int64 `json:"end"`
	// Relative uri to the first page of the query results
	FirstPageUri string `json:"first_page_uri" format:"uri"`
	// Relative uri to the next page of the query results
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize int64 `json:"page_size"`
	// Relative uri to the previous page of the query results
	PreviousPageUri string `json:"previous_page_uri" format:"uri"`
	// The number of the first element on the page, zero-indexed.
	Start          int64                                                    `json:"start"`
	Transcriptions []TexmlAccountGetTranscriptionsJsonResponseTranscription `json:"transcriptions"`
	// The URI of the current page.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End             respjson.Field
		FirstPageUri    respjson.Field
		NextPageUri     respjson.Field
		Page            respjson.Field
		PageSize        respjson.Field
		PreviousPageUri respjson.Field
		Start           respjson.Field
		Transcriptions  respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountGetTranscriptionsJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountGetTranscriptionsJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountGetTranscriptionsJsonResponseTranscription struct {
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion  string    `json:"api_version"`
	CallSid     string    `json:"call_sid"`
	DateCreated time.Time `json:"date_created" format:"date-time"`
	DateUpdated time.Time `json:"date_updated" format:"date-time"`
	// The duration of this recording, given in seconds.
	Duration string `json:"duration,nullable"`
	// Identifier of a resource.
	RecordingSid string `json:"recording_sid"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// The status of the recording transcriptions. The transcription text will be
	// available only when the status is completed.
	//
	// Any of "in-progress", "completed".
	Status string `json:"status"`
	// The recording's transcribed text
	TranscriptionText string `json:"transcription_text"`
	// The relative URI for the recording transcription resource.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid        respjson.Field
		APIVersion        respjson.Field
		CallSid           respjson.Field
		DateCreated       respjson.Field
		DateUpdated       respjson.Field
		Duration          respjson.Field
		RecordingSid      respjson.Field
		Sid               respjson.Field
		Status            respjson.Field
		TranscriptionText respjson.Field
		Uri               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountGetTranscriptionsJsonResponseTranscription) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountGetTranscriptionsJsonResponseTranscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountGetRecordingsJsonParams struct {
	// Filters recording by the creation date. Expected format is ISO8601 date or
	// date-time, ie. {YYYY}-{MM}-{DD} or {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}Z. Also
	// accepts inequality operators, e.g. DateCreated>=2023-05-22.
	DateCreated param.Opt[time.Time] `query:"DateCreated,omitzero" format:"date-time" json:"-"`
	// The number of the page to be displayed, zero-indexed, should be used in
	// conjuction with PageToken.
	Page param.Opt[int64] `query:"Page,omitzero" json:"-"`
	// The number of records to be displayed on a page
	PageSize param.Opt[int64] `query:"PageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlAccountGetRecordingsJsonParams]'s query parameters as
// `url.Values`.
func (r TexmlAccountGetRecordingsJsonParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TexmlAccountGetTranscriptionsJsonParams struct {
	// The number of records to be displayed on a page
	PageSize param.Opt[int64] `query:"PageSize,omitzero" json:"-"`
	// Used to request the next page of results.
	PageToken param.Opt[string] `query:"PageToken,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlAccountGetTranscriptionsJsonParams]'s query parameters
// as `url.Values`.
func (r TexmlAccountGetTranscriptionsJsonParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
