// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Country Coverage
//
// CountryCoverageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCountryCoverageService] method instead.
type CountryCoverageService struct {
	Options []option.RequestOption
}

// NewCountryCoverageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCountryCoverageService(opts ...option.RequestOption) (r CountryCoverageService) {
	r = CountryCoverageService{}
	r.Options = opts
	return
}

// Get country coverage
func (r *CountryCoverageService) Get(ctx context.Context, opts ...option.RequestOption) (res *CountryCoverageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "country_coverage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get coverage for a specific country
func (r *CountryCoverageService) GetCountry(ctx context.Context, countryCode string, opts ...option.RequestOption) (res *CountryCoverageGetCountryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if countryCode == "" {
		err = errors.New("missing required country_code parameter")
		return
	}
	path := fmt.Sprintf("country_coverage/countries/%s", countryCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CountryCoverage struct {
	// Country ISO code
	Code string `json:"code"`
	// Set of features supported
	Features         []string `json:"features"`
	InternationalSMS bool     `json:"international_sms"`
	// Indicates whether country can be queried with inventory coverage endpoint
	InventoryCoverage bool                 `json:"inventory_coverage"`
	Local             CountryCoverageLocal `json:"local"`
	Mobile            map[string]any       `json:"mobile"`
	National          map[string]any       `json:"national"`
	Numbers           bool                 `json:"numbers"`
	P2p               bool                 `json:"p2p"`
	// Phone number type
	PhoneNumberType []string `json:"phone_number_type"`
	// Supports quickship
	Quickship bool `json:"quickship"`
	// Geographic region (e.g., AMER, EMEA, APAC)
	Region string `json:"region" api:"nullable"`
	// Supports reservable
	Reservable bool                    `json:"reservable"`
	SharedCost map[string]any          `json:"shared_cost"`
	TollFree   CountryCoverageTollFree `json:"toll_free"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code              respjson.Field
		Features          respjson.Field
		InternationalSMS  respjson.Field
		InventoryCoverage respjson.Field
		Local             respjson.Field
		Mobile            respjson.Field
		National          respjson.Field
		Numbers           respjson.Field
		P2p               respjson.Field
		PhoneNumberType   respjson.Field
		Quickship         respjson.Field
		Region            respjson.Field
		Reservable        respjson.Field
		SharedCost        respjson.Field
		TollFree          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryCoverage) RawJSON() string { return r.JSON.raw }
func (r *CountryCoverage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryCoverageLocal struct {
	Features            []string `json:"features"`
	FullPstnReplacement bool     `json:"full_pstn_replacement"`
	InternationalSMS    bool     `json:"international_sms"`
	P2p                 bool     `json:"p2p"`
	Quickship           bool     `json:"quickship"`
	Reservable          bool     `json:"reservable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features            respjson.Field
		FullPstnReplacement respjson.Field
		InternationalSMS    respjson.Field
		P2p                 respjson.Field
		Quickship           respjson.Field
		Reservable          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryCoverageLocal) RawJSON() string { return r.JSON.raw }
func (r *CountryCoverageLocal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryCoverageTollFree struct {
	Features            []string `json:"features"`
	FullPstnReplacement bool     `json:"full_pstn_replacement"`
	InternationalSMS    bool     `json:"international_sms"`
	P2p                 bool     `json:"p2p"`
	Quickship           bool     `json:"quickship"`
	Reservable          bool     `json:"reservable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features            respjson.Field
		FullPstnReplacement respjson.Field
		InternationalSMS    respjson.Field
		P2p                 respjson.Field
		Quickship           respjson.Field
		Reservable          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryCoverageTollFree) RawJSON() string { return r.JSON.raw }
func (r *CountryCoverageTollFree) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryCoverageGetResponse struct {
	Data map[string]CountryCoverage `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryCoverageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CountryCoverageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryCoverageGetCountryResponse struct {
	Data CountryCoverage `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryCoverageGetCountryResponse) RawJSON() string { return r.JSON.raw }
func (r *CountryCoverageGetCountryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
