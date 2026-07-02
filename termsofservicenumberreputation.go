// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/respjson"
)

// Accept and review the Branded Calling and Phone Number Reputation terms of
// service.
//
// TermsOfServiceNumberReputationService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTermsOfServiceNumberReputationService] method instead.
type TermsOfServiceNumberReputationService struct {
	Options []option.RequestOption
}

// NewTermsOfServiceNumberReputationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTermsOfServiceNumberReputationService(opts ...option.RequestOption) (r TermsOfServiceNumberReputationService) {
	r = TermsOfServiceNumberReputationService{}
	r.Options = opts
	return
}

// Records the authenticated user's agreement to the current Phone Number
// Reputation ToS. No body required. Idempotent.
//
// Prerequisite for using any of the `/v2/.../reputation/*` endpoints.
func (r *TermsOfServiceNumberReputationService) Agree(ctx context.Context, opts ...option.RequestOption) (res *TermsOfServiceNumberReputationAgreeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "terms_of_service/number_reputation/agree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TermsOfServiceNumberReputationAgreeResponse struct {
	// A recorded user agreement to a product's Terms of Service. The `user_id` is
	// intentionally NOT echoed back on this public surface - the caller already knows
	// their own identity.
	Data TermsOfServiceNumberReputationAgreeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceNumberReputationAgreeResponse) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceNumberReputationAgreeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A recorded user agreement to a product's Terms of Service. The `user_id` is
// intentionally NOT echoed back on this public surface - the caller already knows
// their own identity.
type TermsOfServiceNumberReputationAgreeResponseData struct {
	ID        string    `json:"id" format:"uuid"`
	AgreedAt  time.Time `json:"agreed_at" format:"date-time"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Telnyx product the Terms of Service apply to.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType  string `json:"product_type"`
	TermsVersion string `json:"terms_version"`
	// Convenience alias of `terms_version`. Both keys are present on every response.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AgreedAt     respjson.Field
		CreatedAt    respjson.Field
		ProductType  respjson.Field
		TermsVersion respjson.Field
		Version      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceNumberReputationAgreeResponseData) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceNumberReputationAgreeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
