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
// RecordingService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordingService] method instead.
type RecordingService struct {
	Options []option.RequestOption
	// Call Recordings operations.
	Actions RecordingActionService
}

// NewRecordingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecordingService(opts ...option.RequestOption) (r RecordingService) {
	r = RecordingService{}
	r.Options = opts
	r.Actions = NewRecordingActionService(opts...)
	return
}

// Retrieves the details of an existing call recording.
func (r *RecordingService) Get(ctx context.Context, recordingID string, opts ...option.RequestOption) (res *RecordingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordingID == "" {
		err = errors.New("missing required recording_id parameter")
		return
	}
	path := fmt.Sprintf("recordings/%s", recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your call recordings.
func (r *RecordingService) List(ctx context.Context, query RecordingListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[RecordingResponseData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "recordings"
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

// Returns a list of your call recordings.
func (r *RecordingService) ListAutoPaging(ctx context.Context, query RecordingListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[RecordingResponseData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a call recording.
func (r *RecordingService) Delete(ctx context.Context, recordingID string, opts ...option.RequestOption) (res *RecordingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordingID == "" {
		err = errors.New("missing required recording_id parameter")
		return
	}
	path := fmt.Sprintf("recordings/%s", recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RecordingResponseData struct {
	// Uniquely identifies the recording.
	ID string `json:"id"`
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID unique to the call leg (used to correlate webhook events).
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// When `dual`, the final audio file has the first leg on channel A, and the rest
	// on channel B.
	//
	// Any of "single", "dual".
	Channels RecordingResponseDataChannels `json:"channels"`
	// Uniquely identifies the conference.
	ConferenceID string `json:"conference_id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Links to download the recording files.
	DownloadURLs RecordingResponseDataDownloadURLs `json:"download_urls"`
	// The duration of the recording in milliseconds.
	DurationMillis int64 `json:"duration_millis"`
	// Any of "recording".
	RecordType RecordingResponseDataRecordType `json:"record_type"`
	// ISO 8601 formatted date of when the recording ended.
	RecordingEndedAt string `json:"recording_ended_at"`
	// ISO 8601 formatted date of when the recording started.
	RecordingStartedAt string `json:"recording_started_at"`
	// The kind of event that led to this recording being created.
	//
	// Any of "conference", "call".
	Source RecordingResponseDataSource `json:"source"`
	// The status of the recording. Only `completed` recordings are currently
	// supported.
	//
	// Any of "completed".
	Status RecordingResponseDataStatus `json:"status"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CallControlID      respjson.Field
		CallLegID          respjson.Field
		CallSessionID      respjson.Field
		Channels           respjson.Field
		ConferenceID       respjson.Field
		CreatedAt          respjson.Field
		DownloadURLs       respjson.Field
		DurationMillis     respjson.Field
		RecordType         respjson.Field
		RecordingEndedAt   respjson.Field
		RecordingStartedAt respjson.Field
		Source             respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingResponseData) RawJSON() string { return r.JSON.raw }
func (r *RecordingResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When `dual`, the final audio file has the first leg on channel A, and the rest
// on channel B.
type RecordingResponseDataChannels string

const (
	RecordingResponseDataChannelsSingle RecordingResponseDataChannels = "single"
	RecordingResponseDataChannelsDual   RecordingResponseDataChannels = "dual"
)

// Links to download the recording files.
type RecordingResponseDataDownloadURLs struct {
	// Link to download the recording in mp3 format.
	MP3 string `json:"mp3"`
	// Link to download the recording in wav format.
	Wav string `json:"wav"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingResponseDataDownloadURLs) RawJSON() string { return r.JSON.raw }
func (r *RecordingResponseDataDownloadURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingResponseDataRecordType string

const (
	RecordingResponseDataRecordTypeRecording RecordingResponseDataRecordType = "recording"
)

// The kind of event that led to this recording being created.
type RecordingResponseDataSource string

const (
	RecordingResponseDataSourceConference RecordingResponseDataSource = "conference"
	RecordingResponseDataSourceCall       RecordingResponseDataSource = "call"
)

// The status of the recording. Only `completed` recordings are currently
// supported.
type RecordingResponseDataStatus string

const (
	RecordingResponseDataStatusCompleted RecordingResponseDataStatus = "completed"
)

type RecordingGetResponse struct {
	Data RecordingResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingDeleteResponse struct {
	Data RecordingResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[conference_id], filter[created_at][gte], filter[created_at][lte],
	// filter[call_leg_id], filter[call_session_id], filter[from], filter[to],
	// filter[connection_id], filter[sip_call_id]
	Filter RecordingListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingListParams]'s query parameters as `url.Values`.
func (r RecordingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[conference_id], filter[created_at][gte], filter[created_at][lte],
// filter[call_leg_id], filter[call_session_id], filter[from], filter[to],
// filter[connection_id], filter[sip_call_id]
type RecordingListParamsFilter struct {
	// If present, recordings will be filtered to those with a matching call_leg_id.
	CallLegID param.Opt[string] `query:"call_leg_id,omitzero" format:"uuid" json:"-"`
	// If present, recordings will be filtered to those with a matching
	// call_session_id.
	CallSessionID param.Opt[string] `query:"call_session_id,omitzero" format:"uuid" json:"-"`
	// Returns only recordings associated with a given conference.
	ConferenceID param.Opt[string] `query:"conference_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `connection_id`
	// attribute (case-sensitive).
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `from`
	// attribute (case-sensitive).
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `sip_call_id`
	// attribute. Matching is case-sensitive
	SipCallID param.Opt[string] `query:"sip_call_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `to` attribute
	// (case-sensitive).
	To        param.Opt[string]                  `query:"to,omitzero" json:"-"`
	CreatedAt RecordingListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingListParamsFilter]'s query parameters as
// `url.Values`.
func (r RecordingListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RecordingListParamsFilterCreatedAt struct {
	// Returns only recordings created later than or at given ISO 8601 datetime.
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Returns only recordings created earlier than or at given ISO 8601 datetime.
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingListParamsFilterCreatedAt]'s query parameters as
// `url.Values`.
func (r RecordingListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
