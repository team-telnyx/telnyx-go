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
		return nil, err
	}
	path := fmt.Sprintf("recordings/%s", recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
		return nil, err
	}
	path := fmt.Sprintf("recordings/%s", recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
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
	// Identifies the Telnyx application (Call Control, TeXML) or SIP connection
	// resource associated with this recording.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Links to download the recording files.
	DownloadURLs RecordingResponseDataDownloadURLs `json:"download_urls"`
	// The duration of the recording in milliseconds.
	DurationMillis int64 `json:"duration_millis"`
	// The `from` (caller) number for the call that generated this recording.
	From string `json:"from"`
	// Indicates what triggered the recording. Possible values include `DialVerb`,
	// `Conference`, `OutboundAPI`, `Trunking`, `RecordVerb`, `StartCallRecordingAPI`,
	// `StartConferenceRecordingAPI`.
	InitiatedBy string `json:"initiated_by"`
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
	// The `to` (callee) number for the call that generated this recording.
	To string `json:"to"`
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
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		DownloadURLs       respjson.Field
		DurationMillis     respjson.Field
		From               respjson.Field
		InitiatedBy        respjson.Field
		RecordType         respjson.Field
		RecordingEndedAt   respjson.Field
		RecordingStartedAt respjson.Field
		Source             respjson.Field
		Status             respjson.Field
		To                 respjson.Field
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
	// Filter recordings by various attributes.
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

// Filter recordings by various attributes.
type RecordingListParamsFilter struct {
	// If present, recordings will be filtered to those with a matching
	// `call_control_id`.
	CallControlID param.Opt[string] `query:"call_control_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching call_leg_id.
	CallLegID param.Opt[string] `query:"call_leg_id,omitzero" format:"uuid" json:"-"`
	// If present, recordings will be filtered to those with a matching
	// call_session_id.
	CallSessionID param.Opt[string] `query:"call_session_id,omitzero" format:"uuid" json:"-"`
	// Returns only recordings associated with a given conference.
	ConferenceID param.Opt[string] `query:"conference_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching
	// `conference_region`.
	ConferenceRegion param.Opt[string] `query:"conference_region,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `connection_id`
	// attribute (case-sensitive).
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `from`
	// attribute (case-sensitive).
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `sip_call_id`
	// attribute. Matching is case-sensitive.
	SipCallID param.Opt[string] `query:"sip_call_id,omitzero" json:"-"`
	// If present, recordings will be filtered to those with a matching `to` attribute
	// (case-sensitive).
	To        param.Opt[string]                  `query:"to,omitzero" json:"-"`
	CreatedAt RecordingListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	EndTime   RecordingListParamsFilterEndTime   `query:"end_time,omitzero" json:"-"`
	StartTime RecordingListParamsFilterStartTime `query:"start_time,omitzero" json:"-"`
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

type RecordingListParamsFilterEndTime struct {
	// Returns only recordings with an end time later than or equal to the given ISO
	// 8601 datetime.
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Returns only recordings with an end time earlier than or equal to the given ISO
	// 8601 datetime.
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingListParamsFilterEndTime]'s query parameters as
// `url.Values`.
func (r RecordingListParamsFilterEndTime) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RecordingListParamsFilterStartTime struct {
	// Returns only recordings with a start time later than or equal to the given ISO
	// 8601 datetime.
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Returns only recordings with a start time earlier than or equal to the given ISO
	// 8601 datetime.
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecordingListParamsFilterStartTime]'s query parameters as
// `url.Values`.
func (r RecordingListParamsFilterStartTime) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
