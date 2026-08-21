// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Public pricing operations
//
// PricingProductService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingProductService] method instead.
type PricingProductService struct {
	Options []option.RequestOption
}

// NewPricingProductService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPricingProductService(opts ...option.RequestOption) (r PricingProductService) {
	r = PricingProductService{}
	r.Options = opts
	return
}

// Returns pricing entries for a single product. Most products return standard rate
// entries with fields like rate, unit, country_iso, direction, and tiers.
// Inference products return model-specific fields (model, input_rate, output_rate,
// cached_input_rate) with tiered pricing. Some products use rate decks
// (pricing_type: rate_deck) where rates are determined dynamically.
func (r *PricingProductService) Get(ctx context.Context, slug string, query PricingProductGetParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PricingProductGetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("pricing/products/%s", slug)
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

// Returns pricing entries for a single product. Most products return standard rate
// entries with fields like rate, unit, country_iso, direction, and tiers.
// Inference products return model-specific fields (model, input_rate, output_rate,
// cached_input_rate) with tiered pricing. Some products use rate decks
// (pricing_type: rate_deck) where rates are determined dynamically.
func (r *PricingProductService) GetAutoPaging(ctx context.Context, slug string, query PricingProductGetParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PricingProductGetResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.Get(ctx, slug, query, opts...))
}

// Returns the full product catalog with pagination. Each entry contains a slug,
// display name, and description. Use the slug to fetch per-product pricing via GET
// /pricing/products/{slug}.
func (r *PricingProductService) List(ctx context.Context, query PricingProductListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PricingProductListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pricing/products"
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

// Returns the full product catalog with pagination. Each entry contains a slug,
// display name, and description. Use the slug to fetch per-product pricing via GET
// /pricing/products/{slug}.
func (r *PricingProductService) ListAutoPaging(ctx context.Context, query PricingProductListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PricingProductListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type PricingPaginationMeta struct {
	PageNumber   int64 `json:"page_number" api:"required"`
	PageSize     int64 `json:"page_size" api:"required"`
	TotalPages   int64 `json:"total_pages" api:"required"`
	TotalResults int64 `json:"total_results" api:"required"`
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
func (r PricingPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *PricingPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PricingTier struct {
	// Upper bound of the tier (exclusive). Null means no upper limit.
	Max int64 `json:"max" api:"required"`
	// Lower bound of the tier (inclusive).
	Min int64 `json:"min" api:"required"`
	// Rate for this tier. Numeric for standard products, string for inference
	// products.
	Rate PricingTierRateUnion `json:"rate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PricingTier) RawJSON() string { return r.JSON.raw }
func (r *PricingTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PricingTierRateUnion contains all possible properties and values from [float64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type PricingTierRateUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u PricingTierRateUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PricingTierRateUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PricingTierRateUnion) RawJSON() string { return u.JSON.raw }

func (r *PricingTierRateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single pricing entry. Standard products include rate, unit, currency, type,
// country_iso, direction, and tiers. Inference products include model, input_rate,
// output_rate, cached_input_rate, and their respective tier arrays. Rate-deck
// products include pricing_type and note fields with null rate and empty tiers.
type PricingProductGetResponse struct {
	// Cached input token rate. Present only on inference product entries.
	CachedInputRate string `json:"cached_input_rate"`
	// Cached input token tiered pricing. Present only on inference product entries.
	CachedInputTiers []PricingTier `json:"cached_input_tiers"`
	// ISO country code. Null for non-geographic products.
	CountryISO string `json:"country_iso" api:"nullable"`
	// ISO currency code (e.g., USD).
	Currency string `json:"currency"`
	// Direction (e.g., termination). Null for non-directional products.
	Direction string `json:"direction" api:"nullable"`
	// Input token rate. Present only on inference product entries.
	InputRate string `json:"input_rate"`
	// Input token tiered pricing. Present only on inference product entries.
	InputTiers []PricingTier `json:"input_tiers"`
	// Model identifier. Present only on inference product entries.
	Model string `json:"model"`
	// Human-readable name describing the pricing entry.
	Name string `json:"name"`
	// Additional note for rate-deck products (e.g., "Pricing is determined by the
	// WhatsApp rate deck.").
	Note string `json:"note" api:"nullable"`
	// Output token rate. Present only on inference product entries.
	OutputRate string `json:"output_rate"`
	// Output token tiered pricing. Present only on inference product entries.
	OutputTiers []PricingTier `json:"output_tiers"`
	// Pricing type for non-standard products (e.g., rate_deck). Absent on standard
	// products.
	PricingType string `json:"pricing_type" api:"nullable"`
	// Per-unit rate. Numeric for standard products, string for inference products.
	// Null for rate-deck products.
	Rate PricingProductGetResponseRateUnion `json:"rate" api:"nullable"`
	// Volume-based tiered pricing. Empty for rate-deck products.
	Tiers []PricingTier `json:"tiers"`
	// Pricing type (e.g., usage).
	Type string `json:"type"`
	// Unit of measurement (e.g., part, message, GB, per_1k_tokens).
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedInputRate  respjson.Field
		CachedInputTiers respjson.Field
		CountryISO       respjson.Field
		Currency         respjson.Field
		Direction        respjson.Field
		InputRate        respjson.Field
		InputTiers       respjson.Field
		Model            respjson.Field
		Name             respjson.Field
		Note             respjson.Field
		OutputRate       respjson.Field
		OutputTiers      respjson.Field
		PricingType      respjson.Field
		Rate             respjson.Field
		Tiers            respjson.Field
		Type             respjson.Field
		Unit             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PricingProductGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PricingProductGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PricingProductGetResponseRateUnion contains all possible properties and values
// from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type PricingProductGetResponseRateUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u PricingProductGetResponseRateUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PricingProductGetResponseRateUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PricingProductGetResponseRateUnion) RawJSON() string { return u.JSON.raw }

func (r *PricingProductGetResponseRateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PricingProductListResponse struct {
	// Human-readable description of the product.
	Description string `json:"description" api:"required"`
	// Display name of the product.
	Name string `json:"name" api:"required"`
	// Product identifier used in the per-product pricing endpoint.
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PricingProductListResponse) RawJSON() string { return r.JSON.raw }
func (r *PricingProductListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PricingProductGetParams struct {
	// Two-letter ISO 3166-1 alpha-2 country code (uppercase, e.g. US) to filter
	// pricing to a single country.
	FilterCountryISO param.Opt[string] `query:"filter[country_iso],omitzero" json:"-"`
	// Page number (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page (max 100).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PricingProductGetParams]'s query parameters as
// `url.Values`.
func (r PricingProductGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PricingProductListParams struct {
	// Page number (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page (max 100).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PricingProductListParams]'s query parameters as
// `url.Values`.
func (r PricingProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
