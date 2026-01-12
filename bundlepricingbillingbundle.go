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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// BundlePricingBillingBundleService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBundlePricingBillingBundleService] method instead.
type BundlePricingBillingBundleService struct {
	Options []option.RequestOption
}

// NewBundlePricingBillingBundleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBundlePricingBillingBundleService(opts ...option.RequestOption) (r BundlePricingBillingBundleService) {
	r = BundlePricingBillingBundleService{}
	r.Options = opts
	return
}

// Get a single bundle by ID.
func (r *BundlePricingBillingBundleService) Get(ctx context.Context, bundleID string, query BundlePricingBillingBundleGetParams, opts ...option.RequestOption) (res *BundlePricingBillingBundleGetResponse, err error) {
	if !param.IsOmitted(query.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%s", query.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if bundleID == "" {
		err = errors.New("missing required bundle_id parameter")
		return
	}
	path := fmt.Sprintf("bundle_pricing/billing_bundles/%s", bundleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all allowed bundles.
func (r *BundlePricingBillingBundleService) List(ctx context.Context, params BundlePricingBillingBundleListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[BillingBundleSummary], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%s", params.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bundle_pricing/billing_bundles"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Get all allowed bundles.
func (r *BundlePricingBillingBundleService) ListAutoPaging(ctx context.Context, params BundlePricingBillingBundleListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[BillingBundleSummary] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, params, opts...))
}

type BillingBundleSummary struct {
	// Bundle's ID, this is used to identify the bundle in the API.
	ID string `json:"id,required" format:"uuid"`
	// Bundle's cost code, this is used to identify the bundle in the billing system.
	CostCode string `json:"cost_code,required"`
	// Date the bundle was created.
	CreatedAt time.Time `json:"created_at,required" format:"date"`
	// Available to all customers or only to specific customers.
	IsPublic bool `json:"is_public,required"`
	// Bundle's name, this is used to identify the bundle in the UI.
	Name string `json:"name,required"`
	// Bundle's currency code.
	Currency string `json:"currency"`
	// Monthly recurring charge price.
	MrcPrice float64 `json:"mrc_price"`
	// Slugified version of the bundle's name.
	Slug  string   `json:"slug"`
	Specs []string `json:"specs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CostCode    respjson.Field
		CreatedAt   respjson.Field
		IsPublic    respjson.Field
		Name        respjson.Field
		Currency    respjson.Field
		MrcPrice    respjson.Field
		Slug        respjson.Field
		Specs       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingBundleSummary) RawJSON() string { return r.JSON.raw }
func (r *BillingBundleSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaginationResponse struct {
	// The current page number.
	PageNumber int64 `json:"page_number,required"`
	// The number of results per page.
	PageSize int64 `json:"page_size,required"`
	// Total number of pages from the results.
	TotalPages int64 `json:"total_pages,required"`
	// Total number of results returned.
	TotalResults int64 `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationResponse) RawJSON() string { return r.JSON.raw }
func (r *PaginationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingBillingBundleGetResponse struct {
	Data BundlePricingBillingBundleGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingBillingBundleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingBillingBundleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingBillingBundleGetResponseData struct {
	// Bundle's ID, this is used to identify the bundle in the API.
	ID string `json:"id,required" format:"uuid"`
	// If that bundle is active or not.
	Active       bool                                                   `json:"active,required"`
	BundleLimits []BundlePricingBillingBundleGetResponseDataBundleLimit `json:"bundle_limits,required"`
	// Bundle's cost code, this is used to identify the bundle in the billing system.
	CostCode string `json:"cost_code,required"`
	// Date the bundle was created.
	CreatedAt time.Time `json:"created_at,required" format:"date"`
	// Available to all customers or only to specific customers.
	IsPublic bool `json:"is_public,required"`
	// Bundle's name, this is used to identify the bundle in the UI.
	Name string `json:"name,required"`
	// Slugified version of the bundle's name.
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Active       respjson.Field
		BundleLimits respjson.Field
		CostCode     respjson.Field
		CreatedAt    respjson.Field
		IsPublic     respjson.Field
		Name         respjson.Field
		Slug         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingBillingBundleGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingBillingBundleGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingBillingBundleGetResponseDataBundleLimit struct {
	ID             string    `json:"id,required" format:"uuid"`
	CreatedAt      time.Time `json:"created_at,required" format:"date"`
	Metric         string    `json:"metric,required"`
	Service        string    `json:"service,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date"`
	BillingService string    `json:"billing_service"`
	// Use country_iso instead
	//
	// Deprecated: Use country_iso instead
	Country     string `json:"country"`
	CountryCode int64  `json:"country_code"`
	CountryISO  string `json:"country_iso"`
	// An enumeration.
	//
	// Any of "inbound", "outbound".
	Direction string   `json:"direction"`
	Limit     int64    `json:"limit"`
	Rate      string   `json:"rate"`
	Types     []string `json:"types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Metric         respjson.Field
		Service        respjson.Field
		UpdatedAt      respjson.Field
		BillingService respjson.Field
		Country        respjson.Field
		CountryCode    respjson.Field
		CountryISO     respjson.Field
		Direction      respjson.Field
		Limit          respjson.Field
		Rate           respjson.Field
		Types          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingBillingBundleGetResponseDataBundleLimit) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingBillingBundleGetResponseDataBundleLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingBillingBundleGetParams struct {
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	paramObj
}

type BundlePricingBillingBundleListParams struct {
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Supports filtering by
	// country_iso and resource. Examples: filter[country_iso]=US or
	// filter[resource]=+15617819942
	Filter BundlePricingBillingBundleListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page BundlePricingBillingBundleListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingBillingBundleListParams]'s query parameters as
// `url.Values`.
func (r BundlePricingBillingBundleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Supports filtering by
// country_iso and resource. Examples: filter[country_iso]=US or
// filter[resource]=+15617819942
type BundlePricingBillingBundleListParamsFilter struct {
	// Filter by country code.
	CountryISO []string `query:"country_iso,omitzero" json:"-"`
	// Filter by resource.
	Resource []string `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingBillingBundleListParamsFilter]'s query
// parameters as `url.Values`.
func (r BundlePricingBillingBundleListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type BundlePricingBillingBundleListParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingBillingBundleListParamsPage]'s query
// parameters as `url.Values`.
func (r BundlePricingBillingBundleListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
