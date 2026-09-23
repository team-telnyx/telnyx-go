// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Agentic (bot) signup for Telnyx accounts. An AI agent solves a reverse-CAPTCHA
// challenge designed to be easy for LLMs and hard for humans, registers an
// account, and signs in by consuming a magic link emailed to the account owner.
// All endpoints are public and unauthenticated; signup endpoints are additionally
// gated by the freemium feature flags and per-country availability.
//
// BotSignupService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotSignupService] method instead.
type BotSignupService struct {
	Options []option.RequestOption
}

// NewBotSignupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotSignupService(opts ...option.RequestOption) (r BotSignupService) {
	r = BotSignupService{}
	r.Options = opts
	return
}

// Creates a freemium Telnyx account through the agentic signup flow. The request
// must carry a valid answer to a previously issued bot challenge
// (`bot_challenge_nonce` and `bot_challenge_answer`), accept the terms of service,
// and echo the exact terms-and-conditions and privacy-policy URLs returned by the
// challenge endpoint. When EU consent enforcement is enabled,
// `terms_of_service_eu` and `terms_and_conditions_eu_url` are also required. On
// success a one-time sign-in (magic) link is emailed to the address provided; if
// the email address belongs to an existing account, a sign-in link is sent instead
// of creating a duplicate account. `email` may only be omitted when
// placeholder-email registration is enabled server-side. This endpoint is public
// and unauthenticated, gated by the freemium feature flags and per-country
// availability, and subject to per-IP and per-domain registration limits.
func (r *BotSignupService) New(ctx context.Context, body BotSignupNewParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/bot_signup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Resends the one-time sign-in (magic) link for an eligible bot signup account.
// Eligibility (account exists, was registered through bot signup, is active, and
// has not exceeded the resend limit or rate window) is evaluated server-side; the
// response is intentionally uniform and does not reveal whether the account exists
// or whether a link was actually sent. This endpoint is public and
// unauthenticated, gated by the freemium feature flags and per-country
// availability.
func (r *BotSignupService) ResendMagicLink(ctx context.Context, body BotSignupResendMagicLinkParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/bot_signup/resend_magic_link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Status envelope used by the signup and magic-link flows.
type SuccessResponse struct {
	// Human-readable status message.
	Message string `json:"message" api:"required"`
	// Whether the request was accepted.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *SuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotSignupNewParams struct {
	// Answer to the issued bot challenge.
	BotChallengeAnswer string `json:"bot_challenge_answer" api:"required"`
	// Nonce from a previously issued bot challenge.
	BotChallengeNonce string `json:"bot_challenge_nonce" api:"required" format:"uuid"`
	// Must exactly match the privacy-policy URL returned by the challenge endpoint.
	PrivacyPolicyURL string `json:"privacy_policy_url" api:"required"`
	// Must exactly match the terms-and-conditions URL returned by the challenge
	// endpoint.
	TermsAndConditionsURL string `json:"terms_and_conditions_url" api:"required"`
	// Must be true to accept the terms of service.
	//
	// Any of true.
	TermsOfService bool `json:"terms_of_service,omitzero" api:"required"`
	// Email address for the new account. The magic link is sent here. May only be
	// omitted when placeholder-email registration is enabled server-side.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// EU terms-and-conditions URL. Required when EU consent enforcement is enabled.
	TermsAndConditionsEuURL param.Opt[string] `json:"terms_and_conditions_eu_url,omitzero"`
	// EU terms-of-service acceptance. Required when EU consent enforcement is enabled.
	//
	// Any of true.
	TermsOfServiceEu bool `json:"terms_of_service_eu,omitzero"`
	paramObj
}

func (r BotSignupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BotSignupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BotSignupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotSignupResendMagicLinkParams struct {
	// Email address of the bot signup account to resend the magic link to.
	Email string `json:"email" api:"required" format:"email"`
	paramObj
}

func (r BotSignupResendMagicLinkParams) MarshalJSON() (data []byte, err error) {
	type shadow BotSignupResendMagicLinkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BotSignupResendMagicLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
