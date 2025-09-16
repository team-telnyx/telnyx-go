// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// InventoryCoverageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInventoryCoverageService] method instead.
type InventoryCoverageService struct {
	Options []option.RequestOption
}

// NewInventoryCoverageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInventoryCoverageService(opts ...option.RequestOption) (r InventoryCoverageService) {
	r = InventoryCoverageService{}
	r.Options = opts
	return
}

// Creates an inventory coverage request. If locality, npa or
// national_destination_code is used in groupBy, and no region or locality filters
// are used, the whole paginated set is returned.
func (r *InventoryCoverageService) List(ctx context.Context, query InventoryCoverageListParams, opts ...option.RequestOption) (res *InventoryCoverageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "inventory_coverage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InventoryCoverageListResponse struct {
	Data []InventoryCoverageListResponseData `json:"data"`
	Meta InventoryCoverageListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InventoryCoverageListResponse) RawJSON() string { return r.JSON.raw }
func (r *InventoryCoverageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InventoryCoverageListResponseData struct {
	AdministrativeArea string `json:"administrative_area"`
	// Indicates if the phone number requires advance requirements.
	AdvanceRequirements bool  `json:"advance_requirements"`
	Count               int64 `json:"count"`
	// Any of "number", "block".
	CoverageType string `json:"coverage_type"`
	Group        string `json:"group"`
	GroupType    string `json:"group_type"`
	NumberRange  int64  `json:"number_range"`
	// Any of "did", "toll-free".
	NumberType string `json:"number_type"`
	// Any of "local", "toll_free", "national", "landline", "shared_cost", "mobile".
	PhoneNumberType string `json:"phone_number_type"`
	RecordType      string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea  respjson.Field
		AdvanceRequirements respjson.Field
		Count               respjson.Field
		CoverageType        respjson.Field
		Group               respjson.Field
		GroupType           respjson.Field
		NumberRange         respjson.Field
		NumberType          respjson.Field
		PhoneNumberType     respjson.Field
		RecordType          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InventoryCoverageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *InventoryCoverageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InventoryCoverageListResponseMeta struct {
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InventoryCoverageListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *InventoryCoverageListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InventoryCoverageListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[npa],
	// filter[nxx], filter[administrative_area], filter[phone_number_type],
	// filter[country_code], filter[count], filter[features], filter[groupBy]
	Filter InventoryCoverageListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InventoryCoverageListParams]'s query parameters as
// `url.Values`.
func (r InventoryCoverageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[npa],
// filter[nxx], filter[administrative_area], filter[phone_number_type],
// filter[country_code], filter[count], filter[features], filter[groupBy]
type InventoryCoverageListParamsFilter struct {
	// Filter by administrative area
	AdministrativeArea param.Opt[string] `query:"administrative_area,omitzero" json:"-"`
	// Include count in the result
	Count param.Opt[bool] `query:"count,omitzero" json:"-"`
	// Filter by npa
	Npa param.Opt[int64] `query:"npa,omitzero" json:"-"`
	// Filter by nxx
	Nxx param.Opt[int64] `query:"nxx,omitzero" json:"-"`
	// Filter by country. Defaults to US
	//
	// Any of "AT", "AU", "BE", "BG", "CA", "CH", "CN", "CY", "CZ", "DE", "DK", "EE",
	// "ES", "FI", "FR", "GB", "GR", "HU", "HR", "IE", "IT", "LT", "LU", "LV", "NL",
	// "NZ", "MX", "NO", "PL", "PT", "RO", "SE", "SG", "SI", "SK", "US".
	CountryCode string `query:"country_code,omitzero" json:"-"`
	// Filter if the phone number should be used for voice, fax, mms, sms, emergency.
	// Returns features in the response when used.
	//
	// Any of "sms", "mms", "voice", "fax", "emergency".
	Features []string `query:"features,omitzero" json:"-"`
	// Filter to group results
	//
	// Any of "locality", "npa", "national_destination_code".
	GroupBy string `query:"groupBy,omitzero" json:"-"`
	// Filter by phone number type
	//
	// Any of "local", "toll_free", "national", "mobile", "landline", "shared_cost".
	PhoneNumberType string `query:"phone_number_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InventoryCoverageListParamsFilter]'s query parameters as
// `url.Values`.
func (r InventoryCoverageListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[InventoryCoverageListParamsFilter](
		"country_code", "AT", "AU", "BE", "BG", "CA", "CH", "CN", "CY", "CZ", "DE", "DK", "EE", "ES", "FI", "FR", "GB", "GR", "HU", "HR", "IE", "IT", "LT", "LU", "LV", "NL", "NZ", "MX", "NO", "PL", "PT", "RO", "SE", "SG", "SI", "SK", "US",
	)
	apijson.RegisterFieldValidator[InventoryCoverageListParamsFilter](
		"groupBy", "locality", "npa", "national_destination_code",
	)
	apijson.RegisterFieldValidator[InventoryCoverageListParamsFilter](
		"phone_number_type", "local", "toll_free", "national", "mobile", "landline", "shared_cost",
	)
}
