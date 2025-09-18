// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// AvailablePhoneNumberBlockService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailablePhoneNumberBlockService] method instead.
type AvailablePhoneNumberBlockService struct {
	Options []option.RequestOption
}

// NewAvailablePhoneNumberBlockService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAvailablePhoneNumberBlockService(opts ...option.RequestOption) (r AvailablePhoneNumberBlockService) {
	r = AvailablePhoneNumberBlockService{}
	r.Options = opts
	return
}

// List available phone number blocks
func (r *AvailablePhoneNumberBlockService) List(ctx context.Context, query AvailablePhoneNumberBlockListParams, opts ...option.RequestOption) (res *AvailablePhoneNumberBlockListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "available_phone_number_blocks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AvailablePhoneNumberBlockListResponse struct {
	Data []AvailablePhoneNumberBlockListResponseData `json:"data"`
	Meta AvailablePhoneNumberBlockListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberBlockListResponse) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberBlockListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberBlockListResponseData struct {
	CostInformation AvailablePhoneNumberBlockListResponseDataCostInformation `json:"cost_information"`
	Features        []AvailablePhoneNumberBlockListResponseDataFeature       `json:"features"`
	Range           int64                                                    `json:"range"`
	// Any of "available_phone_number_block".
	RecordType        string                                                       `json:"record_type"`
	RegionInformation []AvailablePhoneNumberBlockListResponseDataRegionInformation `json:"region_information"`
	StartingNumber    string                                                       `json:"starting_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CostInformation   respjson.Field
		Features          respjson.Field
		Range             respjson.Field
		RecordType        respjson.Field
		RegionInformation respjson.Field
		StartingNumber    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberBlockListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberBlockListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberBlockListResponseDataCostInformation struct {
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
func (r AvailablePhoneNumberBlockListResponseDataCostInformation) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberBlockListResponseDataCostInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberBlockListResponseDataFeature struct {
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberBlockListResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberBlockListResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberBlockListResponseDataRegionInformation struct {
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
func (r AvailablePhoneNumberBlockListResponseDataRegionInformation) RawJSON() string {
	return r.JSON.raw
}
func (r *AvailablePhoneNumberBlockListResponseDataRegionInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberBlockListResponseMeta struct {
	BestEffortResults int64 `json:"best_effort_results"`
	TotalResults      int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BestEffortResults respjson.Field
		TotalResults      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumberBlockListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumberBlockListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumberBlockListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[locality],
	// filter[country_code], filter[national_destination_code],
	// filter[phone_number_type]
	Filter AvailablePhoneNumberBlockListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AvailablePhoneNumberBlockListParams]'s query parameters as
// `url.Values`.
func (r AvailablePhoneNumberBlockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[locality],
// filter[country_code], filter[national_destination_code],
// filter[phone_number_type]
type AvailablePhoneNumberBlockListParamsFilter struct {
	// Filter phone numbers by country.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter phone numbers by city.
	Locality param.Opt[string] `query:"locality,omitzero" json:"-"`
	// Filter by the national destination code of the number.
	NationalDestinationCode param.Opt[string] `query:"national_destination_code,omitzero" json:"-"`
	// Filter phone numbers by number type.
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost".
	PhoneNumberType string `query:"phone_number_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AvailablePhoneNumberBlockListParamsFilter]'s query
// parameters as `url.Values`.
func (r AvailablePhoneNumberBlockListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[AvailablePhoneNumberBlockListParamsFilter](
		"phone_number_type", "local", "toll_free", "mobile", "national", "shared_cost",
	)
}
