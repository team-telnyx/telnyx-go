// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// RoomActionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoomActionService] method instead.
type RoomActionService struct {
	Options []option.RequestOption
}

// NewRoomActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoomActionService(opts ...option.RequestOption) (r RoomActionService) {
	r = RoomActionService{}
	r.Options = opts
	return
}

// Synchronously create an Client Token to join a Room. Client Token is necessary
// to join a Telnyx Room. Client Token will expire after `token_ttl_secs`, a
// Refresh Token is also provided to refresh a Client Token, the Refresh Token
// expires after `refresh_token_ttl_secs`.
func (r *RoomActionService) GenerateJoinClientToken(ctx context.Context, roomID string, body RoomActionGenerateJoinClientTokenParams, opts ...option.RequestOption) (res *RoomActionGenerateJoinClientTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomID == "" {
		err = errors.New("missing required room_id parameter")
		return
	}
	path := fmt.Sprintf("rooms/%s/actions/generate_join_client_token", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Synchronously refresh an Client Token to join a Room. Client Token is necessary
// to join a Telnyx Room. Client Token will expire after `token_ttl_secs`.
func (r *RoomActionService) RefreshClientToken(ctx context.Context, roomID string, body RoomActionRefreshClientTokenParams, opts ...option.RequestOption) (res *RoomActionRefreshClientTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomID == "" {
		err = errors.New("missing required room_id parameter")
		return
	}
	path := fmt.Sprintf("rooms/%s/actions/refresh_client_token", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RoomActionGenerateJoinClientTokenResponse struct {
	Data RoomActionGenerateJoinClientTokenResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomActionGenerateJoinClientTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomActionGenerateJoinClientTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomActionGenerateJoinClientTokenResponseData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	// ISO 8601 timestamp when the refresh token expires.
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at" format:"date-time"`
	// ISO 8601 timestamp when the token expires.
	TokenExpiresAt time.Time `json:"token_expires_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token                 respjson.Field
		RefreshToken          respjson.Field
		RefreshTokenExpiresAt respjson.Field
		TokenExpiresAt        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomActionGenerateJoinClientTokenResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomActionGenerateJoinClientTokenResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomActionRefreshClientTokenResponse struct {
	Data RoomActionRefreshClientTokenResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomActionRefreshClientTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *RoomActionRefreshClientTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomActionRefreshClientTokenResponseData struct {
	Token string `json:"token"`
	// ISO 8601 timestamp when the token expires.
	TokenExpiresAt time.Time `json:"token_expires_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token          respjson.Field
		TokenExpiresAt respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomActionRefreshClientTokenResponseData) RawJSON() string { return r.JSON.raw }
func (r *RoomActionRefreshClientTokenResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomActionGenerateJoinClientTokenParams struct {
	// The time to live in seconds of the Refresh Token, after that time the Refresh
	// Token is invalid and can't be used to refresh Client Token.
	RefreshTokenTtlSecs param.Opt[int64] `json:"refresh_token_ttl_secs,omitzero"`
	// The time to live in seconds of the Client Token, after that time the Client
	// Token is invalid and can't be used to join a Room.
	TokenTtlSecs param.Opt[int64] `json:"token_ttl_secs,omitzero"`
	paramObj
}

func (r RoomActionGenerateJoinClientTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow RoomActionGenerateJoinClientTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoomActionGenerateJoinClientTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoomActionRefreshClientTokenParams struct {
	RefreshToken string `json:"refresh_token,required"`
	// The time to live in seconds of the Client Token, after that time the Client
	// Token is invalid and can't be used to join a Room.
	TokenTtlSecs param.Opt[int64] `json:"token_ttl_secs,omitzero"`
	paramObj
}

func (r RoomActionRefreshClientTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow RoomActionRefreshClientTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoomActionRefreshClientTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
