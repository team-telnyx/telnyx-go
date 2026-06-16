// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Accept and review the Branded Calling and Phone Number Reputation terms of
// service.
//
// TermsOfServiceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTermsOfServiceService] method instead.
type TermsOfServiceService struct {
	Options []option.RequestOption
	// Accept and review the Branded Calling and Phone Number Reputation terms of
	// service.
	NumberReputation TermsOfServiceNumberReputationService
	// Accept and review the Branded Calling and Phone Number Reputation terms of
	// service.
	Agreements TermsOfServiceAgreementService
	// Accept and review the Branded Calling and Phone Number Reputation terms of
	// service.
	BrandedCalling TermsOfServiceBrandedCallingService
}

// NewTermsOfServiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTermsOfServiceService(opts ...option.RequestOption) (r TermsOfServiceService) {
	r = TermsOfServiceService{}
	r.Options = opts
	r.NumberReputation = NewTermsOfServiceNumberReputationService(opts...)
	r.Agreements = NewTermsOfServiceAgreementService(opts...)
	r.BrandedCalling = NewTermsOfServiceBrandedCallingService(opts...)
	return
}

// Returns the available Terms of Service agreements (product, current version,
// terms URL, effective date). Omit `product_type` to return all products; pass it
// to scope to one.
func (r *TermsOfServiceService) Info(ctx context.Context, query TermsOfServiceInfoParams, opts ...option.RequestOption) (res *TermsOfServiceInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "terms_of_service/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns whether the authenticated user has agreed to the current Terms of
// Service for the product given by `product_type`. Used during onboarding to
// decide whether to prompt the user with the ToS dialog before continuing.
//
// `agreement_required: true` means the user has not yet agreed (or has agreed to
// an outdated version) and must agree before using that product's endpoints.
func (r *TermsOfServiceService) Status(ctx context.Context, query TermsOfServiceStatusParams, opts ...option.RequestOption) (res *TermsOfServiceStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "terms_of_service/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TermsOfServiceInfoResponse struct {
	Agreements []TermsOfServiceInfoResponseAgreement `json:"agreements"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agreements  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TermsOfServiceInfoResponseAgreement struct {
	CurrentVersion string    `json:"current_version"`
	Description    string    `json:"description"`
	EffectiveDate  time.Time `json:"effective_date" format:"date"`
	// Telnyx product the Terms of Service apply to.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType string `json:"product_type"`
	TermsURL    string `json:"terms_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentVersion respjson.Field
		Description    respjson.Field
		EffectiveDate  respjson.Field
		ProductType    respjson.Field
		TermsURL       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceInfoResponseAgreement) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceInfoResponseAgreement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TermsOfServiceStatusResponse struct {
	// Whether the calling user has agreed to a product's current Terms of Service. The
	// `user_id` is intentionally omitted on this public surface.
	Data TermsOfServiceStatusResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the calling user has agreed to a product's current Terms of Service. The
// `user_id` is intentionally omitted on this public surface.
type TermsOfServiceStatusResponseData struct {
	// `true` when the user must agree to the latest version before using the product.
	// Equivalent to `!has_agreed`.
	AgreementRequired bool `json:"agreement_required" api:"required"`
	// Latest published version of the ToS for this product.
	CurrentTermsVersion string `json:"current_terms_version" api:"required"`
	// `true` if the user has agreed to the latest version.
	HasAgreed bool `json:"has_agreed" api:"required"`
	// Telnyx product the Terms of Service apply to.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType string    `json:"product_type" api:"required"`
	AgreedAt    time.Time `json:"agreed_at" api:"nullable" format:"date-time"`
	// Version the user previously agreed to (may be older than
	// `current_terms_version`). `null` if the user has never agreed.
	AgreedVersion string `json:"agreed_version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementRequired   respjson.Field
		CurrentTermsVersion respjson.Field
		HasAgreed           respjson.Field
		ProductType         respjson.Field
		AgreedAt            respjson.Field
		AgreedVersion       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TermsOfServiceInfoParams struct {
	// Optional product filter. Omit to return info for all products.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType TermsOfServiceInfoParamsProductType `query:"product_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TermsOfServiceInfoParams]'s query parameters as
// `url.Values`.
func (r TermsOfServiceInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional product filter. Omit to return info for all products.
type TermsOfServiceInfoParamsProductType string

const (
	TermsOfServiceInfoParamsProductTypeBrandedCalling   TermsOfServiceInfoParamsProductType = "branded_calling"
	TermsOfServiceInfoParamsProductTypeNumberReputation TermsOfServiceInfoParamsProductType = "number_reputation"
)

type TermsOfServiceStatusParams struct {
	// Which product's ToS to check. Defaults to `branded_calling`.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType TermsOfServiceStatusParamsProductType `query:"product_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TermsOfServiceStatusParams]'s query parameters as
// `url.Values`.
func (r TermsOfServiceStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which product's ToS to check. Defaults to `branded_calling`.
type TermsOfServiceStatusParamsProductType string

const (
	TermsOfServiceStatusParamsProductTypeBrandedCalling   TermsOfServiceStatusParamsProductType = "branded_calling"
	TermsOfServiceStatusParamsProductTypeNumberReputation TermsOfServiceStatusParamsProductType = "number_reputation"
)
