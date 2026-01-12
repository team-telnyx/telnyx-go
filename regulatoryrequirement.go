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
)

// RegulatoryRequirementService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegulatoryRequirementService] method instead.
type RegulatoryRequirementService struct {
	Options []option.RequestOption
}

// NewRegulatoryRequirementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegulatoryRequirementService(opts ...option.RequestOption) (r RegulatoryRequirementService) {
	r = RegulatoryRequirementService{}
	r.Options = opts
	return
}

// Retrieve regulatory requirements
func (r *RegulatoryRequirementService) Get(ctx context.Context, query RegulatoryRequirementGetParams, opts ...option.RequestOption) (res *RegulatoryRequirementGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "regulatory_requirements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RegulatoryRequirementGetResponse struct {
	Data []RegulatoryRequirementGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryRequirementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RegulatoryRequirementGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryRequirementGetResponseData struct {
	Action                 string                                                      `json:"action"`
	CountryCode            string                                                      `json:"country_code"`
	PhoneNumberType        string                                                      `json:"phone_number_type"`
	RegulatoryRequirements []RegulatoryRequirementGetResponseDataRegulatoryRequirement `json:"regulatory_requirements"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action                 respjson.Field
		CountryCode            respjson.Field
		PhoneNumberType        respjson.Field
		RegulatoryRequirements respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryRequirementGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *RegulatoryRequirementGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryRequirementGetResponseDataRegulatoryRequirement struct {
	ID                 string                                                                      `json:"id" format:"uuid"`
	AcceptanceCriteria RegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria `json:"acceptance_criteria"`
	Description        string                                                                      `json:"description"`
	Example            string                                                                      `json:"example"`
	FieldType          string                                                                      `json:"field_type"`
	Name               string                                                                      `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AcceptanceCriteria respjson.Field
		Description        respjson.Field
		Example            respjson.Field
		FieldType          respjson.Field
		Name               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryRequirementGetResponseDataRegulatoryRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *RegulatoryRequirementGetResponseDataRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria struct {
	AcceptableCharacters string   `json:"acceptable_characters"`
	AcceptableValues     []string `json:"acceptable_values"`
	CaseSensitive        string   `json:"case_sensitive"`
	LocalityLimit        string   `json:"locality_limit"`
	MaxLength            string   `json:"max_length"`
	MinLength            string   `json:"min_length"`
	Regex                string   `json:"regex"`
	TimeLimit            string   `json:"time_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptableCharacters respjson.Field
		AcceptableValues     respjson.Field
		CaseSensitive        respjson.Field
		LocalityLimit        respjson.Field
		MaxLength            respjson.Field
		MinLength            respjson.Field
		Regex                respjson.Field
		TimeLimit            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria) RawJSON() string {
	return r.JSON.raw
}
func (r *RegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryRequirementGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[phone_number], filter[requirement_group_id], filter[country_code],
	// filter[phone_number_type], filter[action]
	Filter RegulatoryRequirementGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegulatoryRequirementGetParams]'s query parameters as
// `url.Values`.
func (r RegulatoryRequirementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[phone_number], filter[requirement_group_id], filter[country_code],
// filter[phone_number_type], filter[action]
type RegulatoryRequirementGetParamsFilter struct {
	// Country code(iso2) to check requirements for
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Phone number to check requirements for
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// ID of requirement group to check requirements for
	RequirementGroupID param.Opt[string] `query:"requirement_group_id,omitzero" json:"-"`
	// Action to check requirements for
	//
	// Any of "ordering", "porting", "action".
	Action string `query:"action,omitzero" json:"-"`
	// Phone number type to check requirements for
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost".
	PhoneNumberType string `query:"phone_number_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegulatoryRequirementGetParamsFilter]'s query parameters as
// `url.Values`.
func (r RegulatoryRequirementGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
