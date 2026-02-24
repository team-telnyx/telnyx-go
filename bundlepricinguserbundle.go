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

// BundlePricingUserBundleService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBundlePricingUserBundleService] method instead.
type BundlePricingUserBundleService struct {
	Options []option.RequestOption
}

// NewBundlePricingUserBundleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBundlePricingUserBundleService(opts ...option.RequestOption) (r BundlePricingUserBundleService) {
	r = BundlePricingUserBundleService{}
	r.Options = opts
	return
}

// Creates multiple user bundles for the user.
func (r *BundlePricingUserBundleService) New(ctx context.Context, params BundlePricingUserBundleNewParams, opts ...option.RequestOption) (res *BundlePricingUserBundleNewResponse, err error) {
	if !param.IsOmitted(params.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%v", params.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "bundle_pricing/user_bundles/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves a user bundle by its ID.
func (r *BundlePricingUserBundleService) Get(ctx context.Context, userBundleID string, query BundlePricingUserBundleGetParams, opts ...option.RequestOption) (res *BundlePricingUserBundleGetResponse, err error) {
	if !param.IsOmitted(query.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%v", query.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userBundleID == "" {
		err = errors.New("missing required user_bundle_id parameter")
		return
	}
	path := fmt.Sprintf("bundle_pricing/user_bundles/%s", userBundleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a paginated list of user bundles.
func (r *BundlePricingUserBundleService) List(ctx context.Context, params BundlePricingUserBundleListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[UserBundle], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%v", params.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bundle_pricing/user_bundles"
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

// Get a paginated list of user bundles.
func (r *BundlePricingUserBundleService) ListAutoPaging(ctx context.Context, params BundlePricingUserBundleListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[UserBundle] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deactivates a user bundle by its ID.
func (r *BundlePricingUserBundleService) Deactivate(ctx context.Context, userBundleID string, body BundlePricingUserBundleDeactivateParams, opts ...option.RequestOption) (res *BundlePricingUserBundleDeactivateResponse, err error) {
	if !param.IsOmitted(body.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%v", body.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userBundleID == "" {
		err = errors.New("missing required user_bundle_id parameter")
		return
	}
	path := fmt.Sprintf("bundle_pricing/user_bundles/%s", userBundleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieves the resources of a user bundle by its ID.
func (r *BundlePricingUserBundleService) ListResources(ctx context.Context, userBundleID string, query BundlePricingUserBundleListResourcesParams, opts ...option.RequestOption) (res *BundlePricingUserBundleListResourcesResponse, err error) {
	if !param.IsOmitted(query.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%v", query.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userBundleID == "" {
		err = errors.New("missing required user_bundle_id parameter")
		return
	}
	path := fmt.Sprintf("bundle_pricing/user_bundles/%s/resources", userBundleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all user bundles that aren't in use.
func (r *BundlePricingUserBundleService) ListUnused(ctx context.Context, params BundlePricingUserBundleListUnusedParams, opts ...option.RequestOption) (res *BundlePricingUserBundleListUnusedResponse, err error) {
	if !param.IsOmitted(params.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%v", params.AuthorizationBearer.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "bundle_pricing/user_bundles/unused"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type UserBundle struct {
	// User bundle's ID, this is used to identify the user bundle in the API.
	ID string `json:"id" api:"required" format:"uuid"`
	// Status of the user bundle.
	Active        bool                 `json:"active" api:"required"`
	BillingBundle BillingBundleSummary `json:"billing_bundle" api:"required"`
	// Date the user bundle was created.
	CreatedAt time.Time            `json:"created_at" api:"required" format:"date"`
	Resources []UserBundleResource `json:"resources" api:"required"`
	// The customer's ID that owns this user bundle.
	UserID string `json:"user_id" api:"required" format:"uuid"`
	// Date the user bundle was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Active        respjson.Field
		BillingBundle respjson.Field
		CreatedAt     respjson.Field
		Resources     respjson.Field
		UserID        respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserBundle) RawJSON() string { return r.JSON.raw }
func (r *UserBundle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserBundleResource struct {
	// Resource's ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// Date the resource was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date"`
	// The resource itself (usually a phone number).
	Resource string `json:"resource" api:"required"`
	// The type of the resource (usually 'number').
	ResourceType string `json:"resource_type" api:"required"`
	// Date the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Resource     respjson.Field
		ResourceType respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserBundleResource) RawJSON() string { return r.JSON.raw }
func (r *UserBundleResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleNewResponse struct {
	Data []UserBundle `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingUserBundleNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingUserBundleNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleGetResponse struct {
	Data UserBundle `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingUserBundleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingUserBundleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleDeactivateResponse struct {
	Data UserBundle `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingUserBundleDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingUserBundleDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleListResourcesResponse struct {
	Data []UserBundleResource `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingUserBundleListResourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingUserBundleListResourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleListUnusedResponse struct {
	Data []BundlePricingUserBundleListUnusedResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingUserBundleListUnusedResponse) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingUserBundleListUnusedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleListUnusedResponseData struct {
	BillingBundle BillingBundleSummary `json:"billing_bundle" api:"required"`
	// List of user bundle IDs for given bundle.
	UserBundleIDs []string `json:"user_bundle_ids" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingBundle respjson.Field
		UserBundleIDs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BundlePricingUserBundleListUnusedResponseData) RawJSON() string { return r.JSON.raw }
func (r *BundlePricingUserBundleListUnusedResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleNewParams struct {
	// Idempotency key for the request. Can be any UUID, but should always be unique
	// for each request.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero" format:"uuid"`
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string]                      `header:"authorization_bearer,omitzero" json:"-"`
	Items               []BundlePricingUserBundleNewParamsItem `json:"items,omitzero"`
	paramObj
}

func (r BundlePricingUserBundleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BundlePricingUserBundleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BundlePricingUserBundleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BillingBundleID, Quantity are required.
type BundlePricingUserBundleNewParamsItem struct {
	// Quantity of user bundles to order.
	BillingBundleID string `json:"billing_bundle_id" api:"required" format:"uuid"`
	// Quantity of user bundles to order.
	Quantity int64 `json:"quantity" api:"required"`
	paramObj
}

func (r BundlePricingUserBundleNewParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow BundlePricingUserBundleNewParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BundlePricingUserBundleNewParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BundlePricingUserBundleGetParams struct {
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	paramObj
}

type BundlePricingUserBundleListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Supports filtering by
	// country_iso and resource. Examples: filter[country_iso]=US or
	// filter[resource]=+15617819942
	Filter BundlePricingUserBundleListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingUserBundleListParams]'s query parameters as
// `url.Values`.
func (r BundlePricingUserBundleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Supports filtering by
// country_iso and resource. Examples: filter[country_iso]=US or
// filter[resource]=+15617819942
type BundlePricingUserBundleListParamsFilter struct {
	// Filter by country code.
	CountryISO []string `query:"country_iso,omitzero" json:"-"`
	// Filter by resource.
	Resource []string `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingUserBundleListParamsFilter]'s query parameters
// as `url.Values`.
func (r BundlePricingUserBundleListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BundlePricingUserBundleDeactivateParams struct {
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	paramObj
}

type BundlePricingUserBundleListResourcesParams struct {
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	paramObj
}

type BundlePricingUserBundleListUnusedParams struct {
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Supports filtering by
	// country_iso and resource. Examples: filter[country_iso]=US or
	// filter[resource]=+15617819942
	Filter BundlePricingUserBundleListUnusedParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingUserBundleListUnusedParams]'s query parameters
// as `url.Values`.
func (r BundlePricingUserBundleListUnusedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Supports filtering by
// country_iso and resource. Examples: filter[country_iso]=US or
// filter[resource]=+15617819942
type BundlePricingUserBundleListUnusedParamsFilter struct {
	// Filter by country code.
	CountryISO []string `query:"country_iso,omitzero" json:"-"`
	// Filter by resource.
	Resource []string `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BundlePricingUserBundleListUnusedParamsFilter]'s query
// parameters as `url.Values`.
func (r BundlePricingUserBundleListUnusedParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
