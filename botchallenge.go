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
// BotChallengeService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotChallengeService] method instead.
type BotChallengeService struct {
	Options []option.RequestOption
}

// NewBotChallengeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotChallengeService(opts ...option.RequestOption) (r BotChallengeService) {
	r = BotChallengeService{}
	r.Options = opts
	return
}

// Generates a reverse-CAPTCHA challenge used to gate the bot signup flow. A random
// active problem is selected from the pool; math problems are returned obfuscated
// (case randomization, symbol injection, spacing noise) with an unobfuscated
// rounding instruction appended, while string and binary problems are returned
// as-is. The response contains a single-use nonce, the problem text, and the
// current terms-and-conditions and privacy-policy URLs, which must be echoed back
// on the signup request. Challenges expire after a short window (10 minutes by
// default) and can only be answered once. This endpoint is public and
// unauthenticated.
func (r *BotChallengeService) New(ctx context.Context, body BotChallengeNewParams, opts ...option.RequestOption) (res *BotChallengeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/bot_challenge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BotChallengeNewResponse struct {
	Data BotChallengeNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotChallengeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BotChallengeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotChallengeNewResponseData struct {
	// Type of challenge.
	//
	// Any of "math", "string", "binary".
	ChallengeType string `json:"challenge_type" api:"required"`
	// Single-use challenge identifier. Submit it as `bot_challenge_nonce` on the
	// signup request.
	Nonce string `json:"nonce" api:"required" format:"uuid"`
	// Current privacy-policy URL. Echo this back on the signup request.
	PrivacyPolicyURL string `json:"privacy_policy_url" api:"required"`
	// Problem text to solve. Math problems are obfuscated and end with an unobfuscated
	// rounding instruction; string and binary problems are returned as-is.
	Problem string `json:"problem" api:"required"`
	// Current terms-and-conditions URL. Echo this back on the signup request.
	TermsAndConditionsURL string `json:"terms_and_conditions_url" api:"required"`
	// Decimal places expected in the answer. Present only for math challenges.
	Precision int64 `json:"precision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChallengeType         respjson.Field
		Nonce                 respjson.Field
		PrivacyPolicyURL      respjson.Field
		Problem               respjson.Field
		TermsAndConditionsURL respjson.Field
		Precision             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BotChallengeNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *BotChallengeNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BotChallengeNewParams struct {
	// Name of the LLM the client is using.
	LlmModelName param.Opt[string] `json:"llm_model_name,omitzero"`
	// Parameter count of the client LLM.
	LlmParameterCount param.Opt[string] `json:"llm_parameter_count,omitzero"`
	// Quantization of the client LLM.
	LlmQuantization param.Opt[string] `json:"llm_quantization,omitzero"`
	paramObj
}

func (r BotChallengeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BotChallengeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BotChallengeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
