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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// AvailablePhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailablePhoneNumberService] method instead.
type AvailablePhoneNumberService struct {
	Options []option.RequestOption
}

// NewAvailablePhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAvailablePhoneNumberService(opts ...option.RequestOption) (r AvailablePhoneNumberService) {
	r = AvailablePhoneNumberService{}
	r.Options = opts
	return
}

// List available phone numbers
func (r *AvailablePhoneNumberService) List(ctx context.Context, query AvailablePhoneNumberListParams, opts ...option.RequestOption) (res *AvailablePhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "available_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AvailablePhoneNumberListResponse struct {
	Data     []AvailablePhoneNumberListResponseData `json:"data"`
	Meta     shared.AvailablePhoneNumbersMetadata   `json:"meta"`
	Metadata shared.AvailablePhoneNumbersMetadata   `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberListResponseData struct {
	// Specifies whether the phone number is an exact match based on the search
	// criteria or not.
	BestEffort      bool                                                `json:"best_effort"`
	CostInformation AvailablePhoneNumberListResponseDataCostInformation `json:"cost_information"`
	Features        []AvailablePhoneNumberListResponseDataFeature       `json:"features"`
	PhoneNumber     string                                              `json:"phone_number"`
	// Specifies whether the phone number can receive calls immediately after purchase
	// or not.
	Quickship bool `json:"quickship"`
	// Any of "available_phone_number".
	RecordType        string                                                  `json:"record_type"`
	RegionInformation []AvailablePhoneNumberListResponseDataRegionInformation `json:"region_information"`
	// Specifies whether the phone number can be reserved before purchase or not.
	Reservable   bool   `json:"reservable"`
	VanityFormat string `json:"vanity_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BestEffort        respjson.Field
		CostInformation   respjson.Field
		Features          respjson.Field
		PhoneNumber       respjson.Field
		Quickship         respjson.Field
		RecordType        respjson.Field
		RegionInformation respjson.Field
		Reservable        respjson.Field
		VanityFormat      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberListResponseDataCostInformation struct {
	// The ISO 4217 code for the currency.
	Currency    string `json:"currency"`
	MonthlyCost string `json:"monthly_cost"`
	UpfrontCost string `json:"upfront_cost"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		MonthlyCost respjson.Field
		UpfrontCost respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberListResponseDataCostInformation) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberListResponseDataCostInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberListResponseDataFeature struct {
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberListResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberListResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberListResponseDataRegionInformation struct {
	RegionName string `json:"region_name"`
	// Any of "country_code", "rate_center", "state", "location".
	RegionType string `json:"region_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionName  respjson.Field
		RegionType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberListResponseDataRegionInformation) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberListResponseDataRegionInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[phone_number], filter[locality], filter[administrative_area],
	// filter[country_code], filter[national_destination_code], filter[rate_center],
	// filter[phone_number_type], filter[features], filter[limit], filter[best_effort],
	// filter[quickship], filter[reservable], filter[exclude_held_numbers]
	Filter AvailablePhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AvailablePhoneNumberListParams]'s query parameters as
// `url.Values`.
func (r AvailablePhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[phone_number], filter[locality], filter[administrative_area],
// filter[country_code], filter[national_destination_code], filter[rate_center],
// filter[phone_number_type], filter[features], filter[limit], filter[best_effort],
// filter[quickship], filter[reservable], filter[exclude_held_numbers]
type AvailablePhoneNumberListParamsFilter struct {
	// Find numbers in a particular US state or CA province.
	AdministrativeArea param.Opt[string] `query:"administrative_area,omitzero" json:"-"`
	// Filter to determine if best effort results should be included. Only available in
	// USA/CANADA.
	BestEffort param.Opt[bool] `query:"best_effort,omitzero" json:"-"`
	// Filter phone numbers by country.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter to exclude phone numbers that are currently on hold/reserved for your
	// account.
	ExcludeHeldNumbers param.Opt[bool] `query:"exclude_held_numbers,omitzero" json:"-"`
	// Limits the number of results.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter phone numbers by city.
	Locality param.Opt[string] `query:"locality,omitzero" json:"-"`
	// Filter by the national destination code of the number.
	NationalDestinationCode param.Opt[string] `query:"national_destination_code,omitzero" json:"-"`
	// Filter to exclude phone numbers that need additional time after to purchase to
	// activate. Only applicable for +1 toll_free numbers.
	Quickship param.Opt[bool] `query:"quickship,omitzero" json:"-"`
	// Filter phone numbers by rate center. This filter is only applicable to USA and
	// Canada numbers.
	RateCenter param.Opt[string] `query:"rate_center,omitzero" json:"-"`
	// Filter to ensure only numbers that can be reserved are included in the results.
	Reservable param.Opt[bool] `query:"reservable,omitzero" json:"-"`
	// Filter phone numbers with specific features.
	//
	// Any of "sms", "mms", "voice", "fax", "emergency", "hd_voice",
	// "international_sms", "local_calling".
	Features []string `query:"features,omitzero" json:"-"`
	// Filter phone numbers by pattern matching.
	PhoneNumber AvailablePhoneNumberListParamsFilterPhoneNumber `query:"phone_number,omitzero" json:"-"`
	// Filter phone numbers by number type.
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost".
	PhoneNumberType string `query:"phone_number_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AvailablePhoneNumberListParamsFilter]'s query parameters as
// `url.Values`.
func (r AvailablePhoneNumberListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter phone numbers by pattern matching.
type AvailablePhoneNumberListParamsFilterPhoneNumber struct {
	// Filter numbers containing a pattern (excludes NDC if used with
	// `national_destination_code` filter).
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter numbers ending with a pattern (excludes NDC if used with
	// `national_destination_code` filter).
	EndsWith param.Opt[string] `query:"ends_with,omitzero" json:"-"`
	// Filter numbers starting with a pattern (excludes NDC if used with
	// `national_destination_code` filter).
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AvailablePhoneNumberListParamsFilterPhoneNumber]'s query
// parameters as `url.Values`.
func (r AvailablePhoneNumberListParamsFilterPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
