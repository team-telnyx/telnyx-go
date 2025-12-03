// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// RoomSessionActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomSessionActionService] method instead.
type RoomSessionActionService struct {
	Options []option.RequestOption
}

// NewRoomSessionActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRoomSessionActionService(opts ...option.RequestOption) (r RoomSessionActionService) {
	r = RoomSessionActionService{}
	r.Options = opts
	return
}

// Note: this will also kick all participants currently present in the room
func (r *RoomSessionActionService) End(ctx context.Context, roomSessionID string, opts ...option.RequestOption) (res *RoomSessionActionEndResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomSessionID == "" {
		err = errors.New("missing required room_session_id parameter")
		return
	}
	path := fmt.Sprintf("room_sessions/%s/actions/end", roomSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Kick participants from a room session.
func (r *RoomSessionActionService) Kick(ctx context.Context, roomSessionID string, body RoomSessionActionKickParams, opts ...option.RequestOption) (res *RoomSessionActionKickResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomSessionID == "" {
		err = errors.New("missing required room_session_id parameter")
		return
	}
	path := fmt.Sprintf("room_sessions/%s/actions/kick", roomSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Mute participants in room session.
func (r *RoomSessionActionService) Mute(ctx context.Context, roomSessionID string, body RoomSessionActionMuteParams, opts ...option.RequestOption) (res *RoomSessionActionMuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomSessionID == "" {
		err = errors.New("missing required room_session_id parameter")
		return
	}
	path := fmt.Sprintf("room_sessions/%s/actions/mute", roomSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unmute participants in room session.
func (r *RoomSessionActionService) Unmute(ctx context.Context, roomSessionID string, body RoomSessionActionUnmuteParams, opts ...option.RequestOption) (res *RoomSessionActionUnmuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomSessionID == "" {
		err = errors.New("missing required room_session_id parameter")
		return
	}
	path := fmt.Sprintf("room_sessions/%s/actions/unmute", roomSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ActionsParticipantsRequestParam struct {
	// List of participant id to exclude from the action.
	Exclude []string `json:"exclude,omitzero" format:"uuid"`
	// Either a list of participant id to perform the action on, or the keyword "all"
	// to perform the action on all participant.
	Participants ActionsParticipantsRequestParticipantsUnionParam `json:"participants,omitzero" format:"uuid"`
	paramObj
}

func (r ActionsParticipantsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ActionsParticipantsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionsParticipantsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ActionsParticipantsRequestParticipantsUnionParam struct {
	// Check if union is this variant with !param.IsOmitted(union.OfAllParticipants)
	OfAllParticipants param.Opt[string] `json:",omitzero,inline"`
	OfStringArray     []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ActionsParticipantsRequestParticipantsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAllParticipants, u.OfStringArray)
}
func (u *ActionsParticipantsRequestParticipantsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ActionsParticipantsRequestParticipantsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAllParticipants) {
		return &u.OfAllParticipants
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type ActionsParticipantsRequestParticipantsAllParticipants string

const (
	ActionsParticipantsRequestParticipantsAllParticipantsAll ActionsParticipantsRequestParticipantsAllParticipants = "all"
)

type RoomSessionActionEndResponse struct {
	Data RoomSessionActionEndResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionEndResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionEndResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionEndResponseData struct {
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionEndResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionEndResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionKickResponse struct {
	Data RoomSessionActionKickResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionKickResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionKickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionKickResponseData struct {
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionKickResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionKickResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionMuteResponse struct {
	Data RoomSessionActionMuteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionMuteResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionMuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionMuteResponseData struct {
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionMuteResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionMuteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionUnmuteResponse struct {
	Data RoomSessionActionUnmuteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionUnmuteResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionUnmuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionUnmuteResponseData struct {
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomSessionActionUnmuteResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomSessionActionUnmuteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomSessionActionKickParams struct {
	ActionsParticipantsRequest ActionsParticipantsRequestParam
	paramObj
}

func (r RoomSessionActionKickParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ActionsParticipantsRequest)
}
func (r *RoomSessionActionKickParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ActionsParticipantsRequest)
}

type RoomSessionActionMuteParams struct {
	ActionsParticipantsRequest ActionsParticipantsRequestParam
	paramObj
}

func (r RoomSessionActionMuteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ActionsParticipantsRequest)
}
func (r *RoomSessionActionMuteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ActionsParticipantsRequest)
}

type RoomSessionActionUnmuteParams struct {
	ActionsParticipantsRequest ActionsParticipantsRequestParam
	paramObj
}

func (r RoomSessionActionUnmuteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ActionsParticipantsRequest)
}
func (r *RoomSessionActionUnmuteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ActionsParticipantsRequest)
}
