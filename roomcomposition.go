// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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

// RoomCompositionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomCompositionService] method instead.
type RoomCompositionService struct {
	Options []option.RequestOption
}

// NewRoomCompositionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoomCompositionService(opts ...option.RequestOption) (r RoomCompositionService) {
	r = RoomCompositionService{}
	r.Options = opts
	return
}

// Asynchronously create a room composition.
func (r *RoomCompositionService) New(ctx context.Context, body RoomCompositionNewParams, opts ...option.RequestOption) (res *RoomCompositionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "room_compositions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// View a room composition.
func (r *RoomCompositionService) Get(ctx context.Context, roomCompositionID string, opts ...option.RequestOption) (res *RoomCompositionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomCompositionID == "" {
		err = errors.New("missing required room_composition_id parameter")
		return
	}
	path := fmt.Sprintf("room_compositions/%s", roomCompositionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View a list of room compositions.
func (r *RoomCompositionService) List(ctx context.Context, query RoomCompositionListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[RoomComposition], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "room_compositions"
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

// View a list of room compositions.
func (r *RoomCompositionService) ListAutoPaging(ctx context.Context, query RoomCompositionListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[RoomComposition] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Synchronously delete a room composition.
func (r *RoomCompositionService) Delete(ctx context.Context, roomCompositionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if roomCompositionID == "" {
		err = errors.New("missing required room_composition_id parameter")
		return
	}
	path := fmt.Sprintf("room_compositions/%s", roomCompositionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type RoomComposition struct {
	// A unique identifier for the room composition.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 timestamp when the room composition has completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// ISO 8601 timestamp when the room composition was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Url to download the composition.
	DownloadURL string `json:"download_url"`
	// Shows the room composition duration in seconds.
	DurationSecs int64 `json:"duration_secs"`
	// ISO 8601 timestamp when the room composition has ended.
	EndedAt time.Time `json:"ended_at" format:"date-time"`
	// Shows format of the room composition.
	//
	// Any of "mp4".
	Format     RoomCompositionFormat `json:"format"`
	RecordType string                `json:"record_type"`
	// Identify the room associated with the room composition.
	RoomID string `json:"room_id" format:"uuid"`
	// Identify the room session associated with the room composition.
	SessionID string `json:"session_id" format:"uuid"`
	// Shows the room composition size in MB.
	SizeMB float64 `json:"size_mb"`
	// ISO 8601 timestamp when the room composition has stated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Shows the room composition status.
	//
	// Any of "completed", "enqueued", "processing".
	Status RoomCompositionStatus `json:"status"`
	// ISO 8601 timestamp when the room composition was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Identify the user associated with the room composition.
	UserID string `json:"user_id" format:"uuid"`
	// Describes the video layout of the room composition in terms of regions. Limited
	// to 2 regions.
	VideoLayout map[string]VideoRegion `json:"video_layout"`
	// The failover URL where webhooks related to this room composition will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url" format:"uri"`
	// The URL where webhooks related to this room composition will be sent. Must
	// include a scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CompletedAt             respjson.Field
		CreatedAt               respjson.Field
		DownloadURL             respjson.Field
		DurationSecs            respjson.Field
		EndedAt                 respjson.Field
		Format                  respjson.Field
		RecordType              respjson.Field
		RoomID                  respjson.Field
		SessionID               respjson.Field
		SizeMB                  respjson.Field
		StartedAt               respjson.Field
		Status                  respjson.Field
		UpdatedAt               respjson.Field
		UserID                  respjson.Field
		VideoLayout             respjson.Field
		WebhookEventFailoverURL respjson.Field
		WebhookEventURL         respjson.Field
		WebhookTimeoutSecs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomComposition) RawJSON() string { return r.JSON.raw }
func (r *RoomComposition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shows format of the room composition.
type RoomCompositionFormat string

const (
	RoomCompositionFormatMP4 RoomCompositionFormat = "mp4"
)

// Shows the room composition status.
type RoomCompositionStatus string

const (
	RoomCompositionStatusCompleted  RoomCompositionStatus = "completed"
	RoomCompositionStatusEnqueued   RoomCompositionStatus = "enqueued"
	RoomCompositionStatusProcessing RoomCompositionStatus = "processing"
)

type VideoRegion struct {
	// Height of the video region
	Height int64 `json:"height"`
	// Maximum number of columns of the region's placement grid. By default, the region
	// has as many columns as needed to layout all the specified video sources.
	MaxColumns int64 `json:"max_columns"`
	// Maximum number of rows of the region's placement grid. By default, the region
	// has as many rows as needed to layout all the specified video sources.
	MaxRows int64 `json:"max_rows"`
	// Array of video recording ids to be composed in the region. Can be "\*" to
	// specify all video recordings in the session
	VideoSources []string `json:"video_sources" format:"uuid"`
	// Width of the video region
	Width int64 `json:"width"`
	// X axis value (in pixels) of the region's upper left corner relative to the upper
	// left corner of the whole room composition viewport.
	XPos int64 `json:"x_pos"`
	// Y axis value (in pixels) of the region's upper left corner relative to the upper
	// left corner of the whole room composition viewport.
	YPos int64 `json:"y_pos"`
	// Regions with higher z_pos values are stacked on top of regions with lower z_pos
	// values
	ZPos int64 `json:"z_pos"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height       respjson.Field
		MaxColumns   respjson.Field
		MaxRows      respjson.Field
		VideoSources respjson.Field
		Width        respjson.Field
		XPos         respjson.Field
		YPos         respjson.Field
		ZPos         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoRegion) RawJSON() string { return r.JSON.raw }
func (r *VideoRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VideoRegion to a VideoRegionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VideoRegionParam.Overrides()
func (r VideoRegion) ToParam() VideoRegionParam {
	return param.Override[VideoRegionParam](json.RawMessage(r.RawJSON()))
}

type VideoRegionParam struct {
	// Height of the video region
	Height param.Opt[int64] `json:"height,omitzero"`
	// Maximum number of columns of the region's placement grid. By default, the region
	// has as many columns as needed to layout all the specified video sources.
	MaxColumns param.Opt[int64] `json:"max_columns,omitzero"`
	// Maximum number of rows of the region's placement grid. By default, the region
	// has as many rows as needed to layout all the specified video sources.
	MaxRows param.Opt[int64] `json:"max_rows,omitzero"`
	// Width of the video region
	Width param.Opt[int64] `json:"width,omitzero"`
	// X axis value (in pixels) of the region's upper left corner relative to the upper
	// left corner of the whole room composition viewport.
	XPos param.Opt[int64] `json:"x_pos,omitzero"`
	// Y axis value (in pixels) of the region's upper left corner relative to the upper
	// left corner of the whole room composition viewport.
	YPos param.Opt[int64] `json:"y_pos,omitzero"`
	// Regions with higher z_pos values are stacked on top of regions with lower z_pos
	// values
	ZPos param.Opt[int64] `json:"z_pos,omitzero"`
	// Array of video recording ids to be composed in the region. Can be "\*" to
	// specify all video recordings in the session
	VideoSources []string `json:"video_sources,omitzero" format:"uuid"`
	paramObj
}

func (r VideoRegionParam) MarshalJSON() (data []byte, err error) {
	type shadow VideoRegionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoRegionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomCompositionNewResponse struct {
	Data RoomComposition `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomCompositionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomCompositionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomCompositionGetResponse struct {
	Data RoomComposition `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomCompositionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomCompositionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomCompositionNewParams struct {
	// The desired format of the room composition.
	Format param.Opt[string] `json:"format,omitzero"`
	// The desired resolution (width/height in pixels) of the resulting video of the
	// room composition. Both width and height are required to be between 16 and 1280;
	// and width _ height should not exceed 1280 _ 720
	Resolution param.Opt[string] `json:"resolution,omitzero"`
	// id of the room session associated with the room composition.
	SessionID param.Opt[string] `json:"session_id,omitzero" format:"uuid"`
	// The failover URL where webhooks related to this room composition will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// The URL where webhooks related to this room composition will be sent. Must
	// include a scheme, such as 'https'.
	WebhookEventURL param.Opt[string] `json:"webhook_event_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Describes the video layout of the room composition in terms of regions.
	VideoLayout map[string]VideoRegionParam `json:"video_layout,omitzero"`
	paramObj
}

func (r RoomCompositionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RoomCompositionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoomCompositionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomCompositionListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_created_at][eq], filter[date_created_at][gte],
	// filter[date_created_at][lte], filter[session_id], filter[status]
	Filter RoomCompositionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RoomCompositionListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomCompositionListParams]'s query parameters as
// `url.Values`.
func (r RoomCompositionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_created_at][eq], filter[date_created_at][gte],
// filter[date_created_at][lte], filter[session_id], filter[status]
type RoomCompositionListParamsFilter struct {
	// The session_id for filtering room compositions.
	SessionID     param.Opt[string]                            `query:"session_id,omitzero" format:"uuid" json:"-"`
	DateCreatedAt RoomCompositionListParamsFilterDateCreatedAt `query:"date_created_at,omitzero" json:"-"`
	// The status for filtering room compositions.
	//
	// Any of "completed", "processing", "enqueued".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomCompositionListParamsFilter]'s query parameters as
// `url.Values`.
func (r RoomCompositionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomCompositionListParamsFilterDateCreatedAt struct {
	// ISO 8601 date for filtering room compositions created on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room compositions created on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room compositions created on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomCompositionListParamsFilterDateCreatedAt]'s query
// parameters as `url.Values`.
func (r RoomCompositionListParamsFilterDateCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RoomCompositionListParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomCompositionListParamsPage]'s query parameters as
// `url.Values`.
func (r RoomCompositionListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
