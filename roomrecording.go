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

// RoomRecordingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomRecordingService] method instead.
type RoomRecordingService struct {
	Options []option.RequestOption
}

// NewRoomRecordingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoomRecordingService(opts ...option.RequestOption) (r RoomRecordingService) {
	r = RoomRecordingService{}
	r.Options = opts
	return
}

// View a room recording.
func (r *RoomRecordingService) Get(ctx context.Context, roomRecordingID string, opts ...option.RequestOption) (res *RoomRecordingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomRecordingID == "" {
		err = errors.New("missing required room_recording_id parameter")
		return
	}
	path := fmt.Sprintf("room_recordings/%s", roomRecordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View a list of room recordings.
func (r *RoomRecordingService) List(ctx context.Context, query RoomRecordingListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[RoomRecordingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "room_recordings"
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

// View a list of room recordings.
func (r *RoomRecordingService) ListAutoPaging(ctx context.Context, query RoomRecordingListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[RoomRecordingListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Synchronously delete a Room Recording.
func (r *RoomRecordingService) Delete(ctx context.Context, roomRecordingID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if roomRecordingID == "" {
		err = errors.New("missing required room_recording_id parameter")
		return
	}
	path := fmt.Sprintf("room_recordings/%s", roomRecordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete several room recordings in a bulk.
func (r *RoomRecordingService) DeleteBulk(ctx context.Context, body RoomRecordingDeleteBulkParams, opts ...option.RequestOption) (res *RoomRecordingDeleteBulkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "room_recordings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type RoomRecordingGetResponse struct {
	Data RoomRecordingGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomRecordingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomRecordingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomRecordingGetResponseData struct {
	// A unique identifier for the room recording.
	ID string `json:"id" format:"uuid"`
	// Shows the codec used for the room recording.
	Codec string `json:"codec"`
	// ISO 8601 timestamp when the room recording has completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// ISO 8601 timestamp when the room recording was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Url to download the recording.
	DownloadURL string `json:"download_url"`
	// Shows the room recording duration in seconds.
	DurationSecs int64 `json:"duration_secs"`
	// ISO 8601 timestamp when the room recording has ended.
	EndedAt time.Time `json:"ended_at" format:"date-time"`
	// Identify the room participant associated with the room recording.
	ParticipantID string `json:"participant_id" format:"uuid"`
	RecordType    string `json:"record_type"`
	// Identify the room associated with the room recording.
	RoomID string `json:"room_id" format:"uuid"`
	// Identify the room session associated with the room recording.
	SessionID string `json:"session_id" format:"uuid"`
	// Shows the room recording size in MB.
	SizeMB float64 `json:"size_mb"`
	// ISO 8601 timestamp when the room recording has stated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Shows the room recording status.
	//
	// Any of "completed", "processing".
	Status string `json:"status"`
	// Shows the room recording type.
	//
	// Any of "audio", "video".
	Type string `json:"type"`
	// ISO 8601 timestamp when the room recording was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Codec         respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		DownloadURL   respjson.Field
		DurationSecs  respjson.Field
		EndedAt       respjson.Field
		ParticipantID respjson.Field
		RecordType    respjson.Field
		RoomID        respjson.Field
		SessionID     respjson.Field
		SizeMB        respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomRecordingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomRecordingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomRecordingListResponse struct {
	// A unique identifier for the room recording.
	ID string `json:"id" format:"uuid"`
	// Shows the codec used for the room recording.
	Codec string `json:"codec"`
	// ISO 8601 timestamp when the room recording has completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// ISO 8601 timestamp when the room recording was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Url to download the recording.
	DownloadURL string `json:"download_url"`
	// Shows the room recording duration in seconds.
	DurationSecs int64 `json:"duration_secs"`
	// ISO 8601 timestamp when the room recording has ended.
	EndedAt time.Time `json:"ended_at" format:"date-time"`
	// Identify the room participant associated with the room recording.
	ParticipantID string `json:"participant_id" format:"uuid"`
	RecordType    string `json:"record_type"`
	// Identify the room associated with the room recording.
	RoomID string `json:"room_id" format:"uuid"`
	// Identify the room session associated with the room recording.
	SessionID string `json:"session_id" format:"uuid"`
	// Shows the room recording size in MB.
	SizeMB float64 `json:"size_mb"`
	// ISO 8601 timestamp when the room recording has stated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Shows the room recording status.
	//
	// Any of "completed", "processing".
	Status RoomRecordingListResponseStatus `json:"status"`
	// Shows the room recording type.
	//
	// Any of "audio", "video".
	Type RoomRecordingListResponseType `json:"type"`
	// ISO 8601 timestamp when the room recording was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Codec         respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		DownloadURL   respjson.Field
		DurationSecs  respjson.Field
		EndedAt       respjson.Field
		ParticipantID respjson.Field
		RecordType    respjson.Field
		RoomID        respjson.Field
		SessionID     respjson.Field
		SizeMB        respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomRecordingListResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomRecordingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shows the room recording status.
type RoomRecordingListResponseStatus string

const (
	RoomRecordingListResponseStatusCompleted  RoomRecordingListResponseStatus = "completed"
	RoomRecordingListResponseStatusProcessing RoomRecordingListResponseStatus = "processing"
)

// Shows the room recording type.
type RoomRecordingListResponseType string

const (
	RoomRecordingListResponseTypeAudio RoomRecordingListResponseType = "audio"
	RoomRecordingListResponseTypeVideo RoomRecordingListResponseType = "video"
)

type RoomRecordingDeleteBulkResponse struct {
	Data RoomRecordingDeleteBulkResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomRecordingDeleteBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomRecordingDeleteBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomRecordingDeleteBulkResponseData struct {
	// Amount of room recordings affected
	RoomRecordings int64 `json:"room_recordings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RoomRecordings respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomRecordingDeleteBulkResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomRecordingDeleteBulkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomRecordingListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_ended_at][eq], filter[date_ended_at][gte],
	// filter[date_ended_at][lte], filter[date_started_at][eq],
	// filter[date_started_at][gte], filter[date_started_at][lte], filter[room_id],
	// filter[participant_id], filter[session_id], filter[status], filter[type],
	// filter[duration_secs]
	Filter RoomRecordingListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingListParams]'s query parameters as
// `url.Values`.
func (r RoomRecordingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_ended_at][eq], filter[date_ended_at][gte],
// filter[date_ended_at][lte], filter[date_started_at][eq],
// filter[date_started_at][gte], filter[date_started_at][lte], filter[room_id],
// filter[participant_id], filter[session_id], filter[status], filter[type],
// filter[duration_secs]
type RoomRecordingListParamsFilter struct {
	// duration_secs greater or equal for filtering room recordings.
	DurationSecs param.Opt[int64] `query:"duration_secs,omitzero" json:"-"`
	// participant_id for filtering room recordings.
	ParticipantID param.Opt[string] `query:"participant_id,omitzero" format:"uuid" json:"-"`
	// room_id for filtering room recordings.
	RoomID param.Opt[string] `query:"room_id,omitzero" format:"uuid" json:"-"`
	// session_id for filtering room recordings.
	SessionID param.Opt[string] `query:"session_id,omitzero" format:"uuid" json:"-"`
	// status for filtering room recordings.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// type for filtering room recordings.
	Type          param.Opt[string]                          `query:"type,omitzero" json:"-"`
	DateEndedAt   RoomRecordingListParamsFilterDateEndedAt   `query:"date_ended_at,omitzero" json:"-"`
	DateStartedAt RoomRecordingListParamsFilterDateStartedAt `query:"date_started_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingListParamsFilter]'s query parameters as
// `url.Values`.
func (r RoomRecordingListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomRecordingListParamsFilterDateEndedAt struct {
	// ISO 8601 date for filtering room recordings ended on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings ended on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings ended on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingListParamsFilterDateEndedAt]'s query
// parameters as `url.Values`.
func (r RoomRecordingListParamsFilterDateEndedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomRecordingListParamsFilterDateStartedAt struct {
	// ISO 8601 date for filtering room recordings started on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings started on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings started on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingListParamsFilterDateStartedAt]'s query
// parameters as `url.Values`.
func (r RoomRecordingListParamsFilterDateStartedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomRecordingDeleteBulkParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_ended_at][eq], filter[date_ended_at][gte],
	// filter[date_ended_at][lte], filter[date_started_at][eq],
	// filter[date_started_at][gte], filter[date_started_at][lte], filter[room_id],
	// filter[participant_id], filter[session_id], filter[status], filter[type],
	// filter[duration_secs]
	Filter RoomRecordingDeleteBulkParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingDeleteBulkParams]'s query parameters as
// `url.Values`.
func (r RoomRecordingDeleteBulkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_ended_at][eq], filter[date_ended_at][gte],
// filter[date_ended_at][lte], filter[date_started_at][eq],
// filter[date_started_at][gte], filter[date_started_at][lte], filter[room_id],
// filter[participant_id], filter[session_id], filter[status], filter[type],
// filter[duration_secs]
type RoomRecordingDeleteBulkParamsFilter struct {
	// duration_secs greater or equal for filtering room recordings.
	DurationSecs param.Opt[int64] `query:"duration_secs,omitzero" json:"-"`
	// participant_id for filtering room recordings.
	ParticipantID param.Opt[string] `query:"participant_id,omitzero" format:"uuid" json:"-"`
	// room_id for filtering room recordings.
	RoomID param.Opt[string] `query:"room_id,omitzero" format:"uuid" json:"-"`
	// session_id for filtering room recordings.
	SessionID param.Opt[string] `query:"session_id,omitzero" format:"uuid" json:"-"`
	// status for filtering room recordings.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// type for filtering room recordings.
	Type          param.Opt[string]                                `query:"type,omitzero" json:"-"`
	DateEndedAt   RoomRecordingDeleteBulkParamsFilterDateEndedAt   `query:"date_ended_at,omitzero" json:"-"`
	DateStartedAt RoomRecordingDeleteBulkParamsFilterDateStartedAt `query:"date_started_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingDeleteBulkParamsFilter]'s query parameters as
// `url.Values`.
func (r RoomRecordingDeleteBulkParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomRecordingDeleteBulkParamsFilterDateEndedAt struct {
	// ISO 8601 date for filtering room recordings ended on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings ended on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings ended on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingDeleteBulkParamsFilterDateEndedAt]'s query
// parameters as `url.Values`.
func (r RoomRecordingDeleteBulkParamsFilterDateEndedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomRecordingDeleteBulkParamsFilterDateStartedAt struct {
	// ISO 8601 date for filtering room recordings started on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings started on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room recordings started on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomRecordingDeleteBulkParamsFilterDateStartedAt]'s query
// parameters as `url.Values`.
func (r RoomRecordingDeleteBulkParamsFilterDateStartedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
