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
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// RoomParticipantService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomParticipantService] method instead.
type RoomParticipantService struct {
	Options []option.RequestOption
}

// NewRoomParticipantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoomParticipantService(opts ...option.RequestOption) (r RoomParticipantService) {
	r = RoomParticipantService{}
	r.Options = opts
	return
}

// View a room participant.
func (r *RoomParticipantService) Get(ctx context.Context, roomParticipantID string, opts ...option.RequestOption) (res *RoomParticipantGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomParticipantID == "" {
		err = errors.New("missing required room_participant_id parameter")
		return
	}
	path := fmt.Sprintf("room_participants/%s", roomParticipantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View a list of room participants.
func (r *RoomParticipantService) List(ctx context.Context, query RoomParticipantListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[shared.RoomParticipant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "room_participants"
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

// View a list of room participants.
func (r *RoomParticipantService) ListAutoPaging(ctx context.Context, query RoomParticipantListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[shared.RoomParticipant] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

type RoomParticipantGetResponse struct {
	Data shared.RoomParticipant `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomParticipantGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomParticipantGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomParticipantListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_joined_at][eq], filter[date_joined_at][gte],
	// filter[date_joined_at][lte], filter[date_updated_at][eq],
	// filter[date_updated_at][gte], filter[date_updated_at][lte],
	// filter[date_left_at][eq], filter[date_left_at][gte], filter[date_left_at][lte],
	// filter[context], filter[session_id]
	Filter RoomParticipantListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RoomParticipantListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomParticipantListParams]'s query parameters as
// `url.Values`.
func (r RoomParticipantListParams) URLQuery() (v url.Values, err error) {
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
// filter[context], filter[session_id]
type RoomParticipantListParamsFilter struct {
	// Filter room participants based on the context.
	Context param.Opt[string] `query:"context,omitzero" json:"-"`
	// Session_id for filtering room participants.
	SessionID     param.Opt[string]                            `query:"session_id,omitzero" json:"-"`
	DateJoinedAt  RoomParticipantListParamsFilterDateJoinedAt  `query:"date_joined_at,omitzero" json:"-"`
	DateLeftAt    RoomParticipantListParamsFilterDateLeftAt    `query:"date_left_at,omitzero" json:"-"`
	DateUpdatedAt RoomParticipantListParamsFilterDateUpdatedAt `query:"date_updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomParticipantListParamsFilter]'s query parameters as
// `url.Values`.
func (r RoomParticipantListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomParticipantListParamsFilterDateJoinedAt struct {
	// ISO 8601 date for filtering room participants that joined on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that joined on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that joined on or before that
	// date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomParticipantListParamsFilterDateJoinedAt]'s query
// parameters as `url.Values`.
func (r RoomParticipantListParamsFilterDateJoinedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomParticipantListParamsFilterDateLeftAt struct {
	// ISO 8601 date for filtering room participants that left on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that left on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants that left on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomParticipantListParamsFilterDateLeftAt]'s query
// parameters as `url.Values`.
func (r RoomParticipantListParamsFilterDateLeftAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomParticipantListParamsFilterDateUpdatedAt struct {
	// ISO 8601 date for filtering room participants updated on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants updated on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering room participants updated on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomParticipantListParamsFilterDateUpdatedAt]'s query
// parameters as `url.Values`.
func (r RoomParticipantListParamsFilterDateUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RoomParticipantListParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomParticipantListParamsPage]'s query parameters as
// `url.Values`.
func (r RoomParticipantListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
