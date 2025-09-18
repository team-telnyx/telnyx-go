// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
	"github.com/team-telnyx/telnyx-go/shared"
)

// RoomSessionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomSessionService] method instead.
type RoomSessionService struct {
	Options []option.RequestOption
	Actions RoomSessionActionService
}

// NewRoomSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoomSessionService(opts ...option.RequestOption) (r RoomSessionService) {
	r = RoomSessionService{}
	r.Options = opts
	r.Actions = NewRoomSessionActionService(opts...)
	return
}

// View a room session.
func (r *RoomSessionService) Get(ctx context.Context, roomSessionID string, query RoomSessionGetParams, opts ...option.RequestOption) (res *RoomSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if roomSessionID == "" {
		err = errors.New("missing required room_session_id parameter")
		return
	}
	path := fmt.Sprintf("room_sessions/%s", roomSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// View a list of room sessions.
func (r *RoomSessionService) List0(ctx context.Context, query RoomSessionList0Params, opts ...option.RequestOption) (res *RoomSessionList0Response, err error) {
	opts = append(r.Options[:], opts...)
	path := "room_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// View a list of room sessions.
func (r *RoomSessionService) List1(ctx context.Context, roomID string, query RoomSessionList1Params, opts ...option.RequestOption) (res *RoomSessionList1Response, err error) {
	opts = append(r.Options[:], opts...)
	if roomID == "" {
		err = errors.New("missing required room_id parameter")
		return
	}
	path := fmt.Sprintf("rooms/%s/sessions", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// View a list of room participants.
func (r *RoomSessionService) GetParticipants(ctx context.Context, roomSessionID string, query RoomSessionGetParticipantsParams, opts ...option.RequestOption) (res *RoomSessionGetParticipantsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if roomSessionID == "" {
		err = errors.New("missing required room_session_id parameter")
		return
	}
	path := fmt.Sprintf("room_sessions/%s/participants", roomSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RoomSessionGetResponse struct {
	Data RoomSession `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionList0Response struct {
	Data []RoomSession  `json:"data"`
	Meta PaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionList0Response) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionList0Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionList1Response struct {
	Data []RoomSession  `json:"data"`
	Meta PaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionList1Response) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionList1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionGetParticipantsResponse struct {
	Data []shared.RoomParticipant `json:"data"`
	Meta PaginationMeta           `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionGetParticipantsResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionGetParticipantsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionGetParams struct {
	// To decide if room participants should be included in the response.
	IncludeParticipants param.Opt[bool] `query:"include_participants,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParams]'s query parameters as `url.Values`.
func (r RoomSessionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList0Params struct {
	// To decide if room participants should be included in the response.
	IncludeParticipants param.Opt[bool] `query:"include_participants,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_created_at][eq], filter[date_created_at][gte],
	// filter[date_created_at][lte], filter[date_updated_at][eq],
	// filter[date_updated_at][gte], filter[date_updated_at][lte],
	// filter[date_ended_at][eq], filter[date_ended_at][gte],
	// filter[date_ended_at][lte], filter[room_id], filter[active]
	Filter RoomSessionList0ParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RoomSessionList0ParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList0Params]'s query parameters as `url.Values`.
func (r RoomSessionList0Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_created_at][eq], filter[date_created_at][gte],
// filter[date_created_at][lte], filter[date_updated_at][eq],
// filter[date_updated_at][gte], filter[date_updated_at][lte],
// filter[date_ended_at][eq], filter[date_ended_at][gte],
// filter[date_ended_at][lte], filter[room_id], filter[active]
type RoomSessionList0ParamsFilter struct {
	// Filter active or inactive room sessions.
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
	// Room_id for filtering room sessions.
	RoomID        param.Opt[string]                         `query:"room_id,omitzero" json:"-"`
	DateCreatedAt RoomSessionList0ParamsFilterDateCreatedAt `query:"date_created_at,omitzero" json:"-"`
	DateEndedAt   RoomSessionList0ParamsFilterDateEndedAt   `query:"date_ended_at,omitzero" json:"-"`
	DateUpdatedAt RoomSessionList0ParamsFilterDateUpdatedAt `query:"date_updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList0ParamsFilter]'s query parameters as
// `url.Values`.
func (r RoomSessionList0ParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList0ParamsFilterDateCreatedAt struct {
	// ISO 8601 date for filtering room sessions created on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions created on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions created on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList0ParamsFilterDateCreatedAt]'s query
// parameters as `url.Values`.
func (r RoomSessionList0ParamsFilterDateCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList0ParamsFilterDateEndedAt struct {
	// ISO 8601 date for filtering room sessions ended on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions ended on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions ended on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList0ParamsFilterDateEndedAt]'s query parameters
// as `url.Values`.
func (r RoomSessionList0ParamsFilterDateEndedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList0ParamsFilterDateUpdatedAt struct {
	// ISO 8601 date for filtering room sessions updated on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions updated on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions updated on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList0ParamsFilterDateUpdatedAt]'s query
// parameters as `url.Values`.
func (r RoomSessionList0ParamsFilterDateUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RoomSessionList0ParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList0ParamsPage]'s query parameters as
// `url.Values`.
func (r RoomSessionList0ParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList1Params struct {
	// To decide if room participants should be included in the response.
	IncludeParticipants param.Opt[bool] `query:"include_participants,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_created_at][eq], filter[date_created_at][gte],
	// filter[date_created_at][lte], filter[date_updated_at][eq],
	// filter[date_updated_at][gte], filter[date_updated_at][lte],
	// filter[date_ended_at][eq], filter[date_ended_at][gte],
	// filter[date_ended_at][lte], filter[active]
	Filter RoomSessionList1ParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RoomSessionList1ParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList1Params]'s query parameters as `url.Values`.
func (r RoomSessionList1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_created_at][eq], filter[date_created_at][gte],
// filter[date_created_at][lte], filter[date_updated_at][eq],
// filter[date_updated_at][gte], filter[date_updated_at][lte],
// filter[date_ended_at][eq], filter[date_ended_at][gte],
// filter[date_ended_at][lte], filter[active]
type RoomSessionList1ParamsFilter struct {
	// Filter active or inactive room sessions.
	Active        param.Opt[bool]                           `query:"active,omitzero" json:"-"`
	DateCreatedAt RoomSessionList1ParamsFilterDateCreatedAt `query:"date_created_at,omitzero" json:"-"`
	DateEndedAt   RoomSessionList1ParamsFilterDateEndedAt   `query:"date_ended_at,omitzero" json:"-"`
	DateUpdatedAt RoomSessionList1ParamsFilterDateUpdatedAt `query:"date_updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList1ParamsFilter]'s query parameters as
// `url.Values`.
func (r RoomSessionList1ParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList1ParamsFilterDateCreatedAt struct {
	// ISO 8601 date for filtering room sessions created on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions created on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions created on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList1ParamsFilterDateCreatedAt]'s query
// parameters as `url.Values`.
func (r RoomSessionList1ParamsFilterDateCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList1ParamsFilterDateEndedAt struct {
	// ISO 8601 date for filtering room sessions ended on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions ended on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions ended on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList1ParamsFilterDateEndedAt]'s query parameters
// as `url.Values`.
func (r RoomSessionList1ParamsFilterDateEndedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionList1ParamsFilterDateUpdatedAt struct {
	// ISO 8601 date for filtering room sessions updated on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions updated on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room sessions updated on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList1ParamsFilterDateUpdatedAt]'s query
// parameters as `url.Values`.
func (r RoomSessionList1ParamsFilterDateUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RoomSessionList1ParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionList1ParamsPage]'s query parameters as
// `url.Values`.
func (r RoomSessionList1ParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionGetParticipantsParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_joined_at][eq], filter[date_joined_at][gte],
	// filter[date_joined_at][lte], filter[date_updated_at][eq],
	// filter[date_updated_at][gte], filter[date_updated_at][lte],
	// filter[date_left_at][eq], filter[date_left_at][gte], filter[date_left_at][lte],
	// filter[context]
	Filter RoomSessionGetParticipantsParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RoomSessionGetParticipantsParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParticipantsParams]'s query parameters as
// `url.Values`.
func (r RoomSessionGetParticipantsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_joined_at][eq], filter[date_joined_at][gte],
// filter[date_joined_at][lte], filter[date_updated_at][eq],
// filter[date_updated_at][gte], filter[date_updated_at][lte],
// filter[date_left_at][eq], filter[date_left_at][gte], filter[date_left_at][lte],
// filter[context]
type RoomSessionGetParticipantsParamsFilter struct {
	// Filter room participants based on the context.
	Context       param.Opt[string]                                   `query:"context,omitzero" json:"-"`
	DateJoinedAt  RoomSessionGetParticipantsParamsFilterDateJoinedAt  `query:"date_joined_at,omitzero" json:"-"`
	DateLeftAt    RoomSessionGetParticipantsParamsFilterDateLeftAt    `query:"date_left_at,omitzero" json:"-"`
	DateUpdatedAt RoomSessionGetParticipantsParamsFilterDateUpdatedAt `query:"date_updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParticipantsParamsFilter]'s query parameters
// as `url.Values`.
func (r RoomSessionGetParticipantsParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionGetParticipantsParamsFilterDateJoinedAt struct {
	// ISO 8601 date for filtering room participants that joined on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that joined on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that joined on or before that
	// date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParticipantsParamsFilterDateJoinedAt]'s query
// parameters as `url.Values`.
func (r RoomSessionGetParticipantsParamsFilterDateJoinedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionGetParticipantsParamsFilterDateLeftAt struct {
	// ISO 8601 date for filtering room participants that left on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that left on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that left on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParticipantsParamsFilterDateLeftAt]'s query
// parameters as `url.Values`.
func (r RoomSessionGetParticipantsParamsFilterDateLeftAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomSessionGetParticipantsParamsFilterDateUpdatedAt struct {
	// ISO 8601 date for filtering room participants updated on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants updated on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants updated on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParticipantsParamsFilterDateUpdatedAt]'s
// query parameters as `url.Values`.
func (r RoomSessionGetParticipantsParamsFilterDateUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RoomSessionGetParticipantsParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomSessionGetParticipantsParamsPage]'s query parameters as
// `url.Values`.
func (r RoomSessionGetParticipantsParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
