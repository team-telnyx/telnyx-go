// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Agentic (bot) signup for Telnyx accounts. An AI agent solves a reverse-CAPTCHA
// challenge designed to be easy for LLMs and hard for humans, registers an
// account, and signs in by consuming a magic link emailed to the account owner.
// All endpoints are public and unauthenticated; signup endpoints are additionally
// gated by the freemium feature flags and per-country availability.
//
// BotSessionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotSessionService] method instead.
type BotSessionService struct {
	Options []option.RequestOption
}

// NewBotSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotSessionService(opts ...option.RequestOption) (r BotSessionService) {
	r = BotSessionService{}
	r.Options = opts
	return
}

// Consumes the one-time portal redirect (magic link) token emailed during bot
// signup and returns an API session. The token is a UUIDv7 that encodes its
// creation time; it expires after a configurable validity window (15 minutes by
// default) and is cleared on first use. Although the action creates a session, the
// route uses the GET verb because it is opened from an email link. On first use
// the account is also initialized. For bot signup (freemium) accounts the response
// is a minimal envelope containing only the `api_v2_token`; accounts that are
// permitted to use magic links but are not freemium accounts may instead receive
// an extended session payload when additional steps (such as two-factor
// authentication or identity verification) are required. This endpoint is public;
// the magic link token in the query string is the credential.
func (r *BotSessionService) List(ctx context.Context, query BotSessionListParams, opts ...option.RequestOption) (res *BotSessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/bot_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BotSessionListResponse struct {
	Data BotSessionListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotSessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *BotSessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotSessionListResponseData struct {
	// API v2 session token for the signed-in user. Use it as a bearer token on
	// authenticated endpoints.
	APIV2Token string `json:"api_v2_token" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIV2Token  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotSessionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BotSessionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotSessionListParams struct {
	// Email address associated with the magic link token.
	Email string `query:"email" api:"required" format:"email" json:"-"`
	// Single-use portal redirect (magic link) token, a UUIDv7 sent to the account
	// owner's email.
	PortalRedirectToken string `query:"portal_redirect_token" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BotSessionListParams]'s query parameters as `url.Values`.
func (r BotSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
