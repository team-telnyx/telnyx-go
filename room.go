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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
	"github.com/team-telnyx/telnyx-go/shared"
)

// RoomService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomService] method instead.
type RoomService struct {
	Options  []option.RequestOption
	Actions  RoomActionService
	Sessions RoomSessionService
}

// NewRoomService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoomService(opts ...option.RequestOption) (r RoomService) {
	r = RoomService{}
	r.Options = opts
	r.Actions = NewRoomActionService(opts...)
	r.Sessions = NewRoomSessionService(opts...)
	return
}

// Synchronously create a Room.
func (r *RoomService) New(ctx context.Context, body RoomNewParams, opts ...option.RequestOption) (res *RoomNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rooms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// View a room.
func (r *RoomService) Get(ctx context.Context, roomID string, query RoomGetParams, opts ...option.RequestOption) (res *RoomGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomID == "" {
		err = errors.New("missing required room_id parameter")
		return
	}
	path := fmt.Sprintf("rooms/%s", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Synchronously update a Room.
func (r *RoomService) Update(ctx context.Context, roomID string, body RoomUpdateParams, opts ...option.RequestOption) (res *RoomUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomID == "" {
		err = errors.New("missing required room_id parameter")
		return
	}
	path := fmt.Sprintf("rooms/%s", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// View a list of rooms.
func (r *RoomService) List(ctx context.Context, query RoomListParams, opts ...option.RequestOption) (res *RoomListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rooms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Synchronously delete a Room. Participants from that room will be kicked out,
// they won't be able to join that room anymore, and you won't be charged anymore
// for that room.
func (r *RoomService) Delete(ctx context.Context, roomID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if roomID == "" {
		err = errors.New("missing required room_id parameter")
		return
	}
	path := fmt.Sprintf("rooms/%s", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Room struct {
	// A unique identifier for the room.
	ID string `json:"id" format:"uuid"`
	// The identifier of the active room session if any.
	ActiveSessionID string `json:"active_session_id" format:"uuid"`
	// ISO 8601 timestamp when the room was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Enable or disable recording for that room.
	EnableRecording bool `json:"enable_recording"`
	// Maximum participants allowed in the room.
	MaxParticipants int64         `json:"max_participants"`
	RecordType      string        `json:"record_type"`
	Sessions        []RoomSession `json:"sessions"`
	// The unique (within the Telnyx account scope) name of the room.
	UniqueName string `json:"unique_name"`
	// ISO 8601 timestamp when the room was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The failover URL where webhooks related to this room will be sent if sending to
	// the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable" format:"uri"`
	// The URL where webhooks related to this room will be sent. Must include a scheme,
	// such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		ActiveSessionID         respjson.Field
		CreatedAt               respjson.Field
		EnableRecording         respjson.Field
		MaxParticipants         respjson.Field
		RecordType              respjson.Field
		Sessions                respjson.Field
		UniqueName              respjson.Field
		UpdatedAt               respjson.Field
		WebhookEventFailoverURL respjson.Field
		WebhookEventURL         respjson.Field
		WebhookTimeoutSecs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Room) RawJSON() string { return r.JSON.raw }
func (r *Room) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSession struct {
	// A unique identifier for the room session.
	ID string `json:"id" format:"uuid"`
	// Shows if the room session is active or not.
	Active bool `json:"active"`
	// ISO 8601 timestamp when the room session was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ISO 8601 timestamp when the room session has ended.
	EndedAt      time.Time                `json:"ended_at" format:"date-time"`
	Participants []shared.RoomParticipant `json:"participants"`
	RecordType   string                   `json:"record_type"`
	// Identify the room hosting that room session.
	RoomID string `json:"room_id" format:"uuid"`
	// ISO 8601 timestamp when the room session was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Active       respjson.Field
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Participants respjson.Field
		RecordType   respjson.Field
		RoomID       respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSession) RawJSON() string { return r.JSON.raw }
func (r *RoomSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomNewResponse struct {
	Data Room `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomGetResponse struct {
	Data Room `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomUpdateResponse struct {
	Data Room `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomListResponse struct {
	Data []Room         `json:"data"`
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
func (r RoomListResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomNewParams struct {
	// The failover URL where webhooks related to this room will be sent if sending to
	// the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Enable or disable recording for that room.
	EnableRecording param.Opt[bool] `json:"enable_recording,omitzero"`
	// The maximum amount of participants allowed in a room. If new participants try to
	// join after that limit is reached, their request will be rejected.
	MaxParticipants param.Opt[int64] `json:"max_participants,omitzero"`
	// The unique (within the Telnyx account scope) name of the room.
	UniqueName param.Opt[string] `json:"unique_name,omitzero"`
	// The URL where webhooks related to this room will be sent. Must include a scheme,
	// such as 'https'.
	WebhookEventURL param.Opt[string] `json:"webhook_event_url,omitzero" format:"uri"`
	paramObj
}

func (r RoomNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RoomNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoomNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomGetParams struct {
	// To decide if room sessions should be included in the response.
	IncludeSessions param.Opt[bool] `query:"include_sessions,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomGetParams]'s query parameters as `url.Values`.
func (r RoomGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomUpdateParams struct {
	// The failover URL where webhooks related to this room will be sent if sending to
	// the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Enable or disable recording for that room.
	EnableRecording param.Opt[bool] `json:"enable_recording,omitzero"`
	// The maximum amount of participants allowed in a room. If new participants try to
	// join after that limit is reached, their request will be rejected.
	MaxParticipants param.Opt[int64] `json:"max_participants,omitzero"`
	// The unique (within the Telnyx account scope) name of the room.
	UniqueName param.Opt[string] `json:"unique_name,omitzero"`
	// The URL where webhooks related to this room will be sent. Must include a scheme,
	// such as 'https'.
	WebhookEventURL param.Opt[string] `json:"webhook_event_url,omitzero" format:"uri"`
	paramObj
}

func (r RoomUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RoomUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoomUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomListParams struct {
	// To decide if room sessions should be included in the response.
	IncludeSessions param.Opt[bool] `query:"include_sessions,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[date_created_at][eq], filter[date_created_at][gte],
	// filter[date_created_at][lte], filter[date_updated_at][eq],
	// filter[date_updated_at][gte], filter[date_updated_at][lte], filter[unique_name]
	Filter RoomListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RoomListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomListParams]'s query parameters as `url.Values`.
func (r RoomListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[date_created_at][eq], filter[date_created_at][gte],
// filter[date_created_at][lte], filter[date_updated_at][eq],
// filter[date_updated_at][gte], filter[date_updated_at][lte], filter[unique_name]
type RoomListParamsFilter struct {
	// Unique_name for filtering rooms.
	UniqueName    param.Opt[string]                 `query:"unique_name,omitzero" json:"-"`
	DateCreatedAt RoomListParamsFilterDateCreatedAt `query:"date_created_at,omitzero" json:"-"`
	DateUpdatedAt RoomListParamsFilterDateUpdatedAt `query:"date_updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomListParamsFilter]'s query parameters as `url.Values`.
func (r RoomListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomListParamsFilterDateCreatedAt struct {
	// ISO 8601 date for filtering rooms created on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering rooms created on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering rooms created on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomListParamsFilterDateCreatedAt]'s query parameters as
// `url.Values`.
func (r RoomListParamsFilterDateCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoomListParamsFilterDateUpdatedAt struct {
	// ISO 8601 date for filtering rooms updated on that date.
	Eq param.Opt[time.Time] `query:"eq,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering rooms updated on or after that date.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date" json:"-"`
	// ISO 8601 date for filtering rooms updated on or before that date.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [RoomListParamsFilterDateUpdatedAt]'s query parameters as
// `url.Values`.
func (r RoomListParamsFilterDateUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RoomListParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoomListParamsPage]'s query parameters as `url.Values`.
func (r RoomListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
