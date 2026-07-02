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

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/pagination"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
	"github.com/stainless-sdks/telnyx-go/v4/packages/respjson"
)

// Accept and review the Branded Calling and Phone Number Reputation terms of
// service.
//
// TermsOfServiceAgreementService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTermsOfServiceAgreementService] method instead.
type TermsOfServiceAgreementService struct {
	Options []option.RequestOption
}

// NewTermsOfServiceAgreementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTermsOfServiceAgreementService(opts ...option.RequestOption) (r TermsOfServiceAgreementService) {
	r = TermsOfServiceAgreementService{}
	r.Options = opts
	return
}

// Retrieve a single ToS agreement record. Returns `404` if the agreement does not
// exist or does not belong to the authenticated user.
func (r *TermsOfServiceAgreementService) Get(ctx context.Context, agreementID string, opts ...option.RequestOption) (res *TermsOfServiceAgreementGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agreementID == "" {
		err = errors.New("missing required agreement_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("terms_of_service/agreements/%s", agreementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the Terms of Service agreements the authenticated user has on file. Each
// entry records the version agreed to and when. Most users only have one agreement
// per product, but if the ToS is updated and the user re-agrees a new entry is
// added.
//
// Results are paginated with the standard `page[number]` / `page[size]`
// parameters; the response uses the standard `{data, meta}` JSON:API envelope.
//
// By default this returns agreements for **all** products the user has agreed to.
// Pass the `product_type` query parameter to scope the result to a single product.
func (r *TermsOfServiceAgreementService) List(ctx context.Context, query TermsOfServiceAgreementListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[TermsOfServiceAgreementListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "terms_of_service/agreements"
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

// Returns the Terms of Service agreements the authenticated user has on file. Each
// entry records the version agreed to and when. Most users only have one agreement
// per product, but if the ToS is updated and the user re-agrees a new entry is
// added.
//
// Results are paginated with the standard `page[number]` / `page[size]`
// parameters; the response uses the standard `{data, meta}` JSON:API envelope.
//
// By default this returns agreements for **all** products the user has agreed to.
// Pass the `product_type` query parameter to scope the result to a single product.
func (r *TermsOfServiceAgreementService) ListAutoPaging(ctx context.Context, query TermsOfServiceAgreementListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[TermsOfServiceAgreementListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type TermsOfServiceAgreementGetResponse struct {
	// A recorded user agreement to a product's Terms of Service. The `user_id` is
	// intentionally NOT echoed back on this public surface - the caller already knows
	// their own identity.
	Data TermsOfServiceAgreementGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TermsOfServiceAgreementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceAgreementGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A recorded user agreement to a product's Terms of Service. The `user_id` is
// intentionally NOT echoed back on this public surface - the caller already knows
// their own identity.
type TermsOfServiceAgreementGetResponseData struct {
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
func (r TermsOfServiceAgreementGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceAgreementGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A recorded user agreement to a product's Terms of Service. The `user_id` is
// intentionally NOT echoed back on this public surface - the caller already knows
// their own identity.
type TermsOfServiceAgreementListResponse struct {
	ID        string    `json:"id" format:"uuid"`
	AgreedAt  time.Time `json:"agreed_at" format:"date-time"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Telnyx product the Terms of Service apply to.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType  TermsOfServiceAgreementListResponseProductType `json:"product_type"`
	TermsVersion string                                         `json:"terms_version"`
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
func (r TermsOfServiceAgreementListResponse) RawJSON() string { return r.JSON.raw }
func (r *TermsOfServiceAgreementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Telnyx product the Terms of Service apply to.
type TermsOfServiceAgreementListResponseProductType string

const (
	TermsOfServiceAgreementListResponseProductTypeBrandedCalling   TermsOfServiceAgreementListResponseProductType = "branded_calling"
	TermsOfServiceAgreementListResponseProductTypeNumberReputation TermsOfServiceAgreementListResponseProductType = "number_reputation"
)

type TermsOfServiceAgreementListParams struct {
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Optional filter. Omit to list the user's agreements for **every** product
	// (branded_calling and number_reputation); pass a value to return only that
	// product's agreements.
	//
	// Any of "branded_calling", "number_reputation".
	ProductType TermsOfServiceAgreementListParamsProductType `query:"product_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TermsOfServiceAgreementListParams]'s query parameters as
// `url.Values`.
func (r TermsOfServiceAgreementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional filter. Omit to list the user's agreements for **every** product
// (branded_calling and number_reputation); pass a value to return only that
// product's agreements.
type TermsOfServiceAgreementListParamsProductType string

const (
	TermsOfServiceAgreementListParamsProductTypeBrandedCalling   TermsOfServiceAgreementListParamsProductType = "branded_calling"
	TermsOfServiceAgreementListParamsProductTypeNumberReputation TermsOfServiceAgreementListParamsProductType = "number_reputation"
)
