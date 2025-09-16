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

// PhoneNumbersRegulatoryRequirementService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumbersRegulatoryRequirementService] method instead.
type PhoneNumbersRegulatoryRequirementService struct {
	Options []option.RequestOption
}

// NewPhoneNumbersRegulatoryRequirementService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPhoneNumbersRegulatoryRequirementService(opts ...option.RequestOption) (r PhoneNumbersRegulatoryRequirementService) {
	r = PhoneNumbersRegulatoryRequirementService{}
	r.Options = opts
	return
}

// Retrieve regulatory requirements for a list of phone numbers
func (r *PhoneNumbersRegulatoryRequirementService) Get(ctx context.Context, query PhoneNumbersRegulatoryRequirementGetParams, opts ...option.RequestOption) (res *PhoneNumbersRegulatoryRequirementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "phone_numbers_regulatory_requirements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PhoneNumbersRegulatoryRequirementGetResponse struct {
	Data []PhoneNumbersRegulatoryRequirementGetResponseData `json:"data"`
	Meta PaginationMeta                                     `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersRegulatoryRequirementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersRegulatoryRequirementGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumbersRegulatoryRequirementGetResponseData struct {
	PhoneNumber            string                                                                  `json:"phone_number"`
	PhoneNumberType        string                                                                  `json:"phone_number_type"`
	RecordType             string                                                                  `json:"record_type"`
	RegionInformation      []PhoneNumbersRegulatoryRequirementGetResponseDataRegionInformation     `json:"region_information"`
	RegulatoryRequirements []PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirement `json:"regulatory_requirements"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber            respjson.Field
		PhoneNumberType        respjson.Field
		RecordType             respjson.Field
		RegionInformation      respjson.Field
		RegulatoryRequirements respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersRegulatoryRequirementGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersRegulatoryRequirementGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumbersRegulatoryRequirementGetResponseDataRegionInformation struct {
	RegionName string `json:"region_name"`
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
func (r PhoneNumbersRegulatoryRequirementGetResponseDataRegionInformation) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumbersRegulatoryRequirementGetResponseDataRegionInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirement struct {
	ID                 string                                                                                  `json:"id" format:"uuid"`
	AcceptanceCriteria PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria `json:"acceptance_criteria"`
	Description        string                                                                                  `json:"description"`
	Example            string                                                                                  `json:"example"`
	FieldType          string                                                                                  `json:"field_type"`
	Label              string                                                                                  `json:"label"`
	RecordType         string                                                                                  `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AcceptanceCriteria respjson.Field
		Description        respjson.Field
		Example            respjson.Field
		FieldType          respjson.Field
		Label              respjson.Field
		RecordType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria struct {
	FieldType     string `json:"field_type"`
	FieldValue    string `json:"field_value"`
	LocalityLimit string `json:"locality_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType     respjson.Field
		FieldValue    respjson.Field
		LocalityLimit respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumbersRegulatoryRequirementGetResponseDataRegulatoryRequirementAcceptanceCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumbersRegulatoryRequirementGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[phone_number]
	Filter PhoneNumbersRegulatoryRequirementGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumbersRegulatoryRequirementGetParams]'s query
// parameters as `url.Values`.
func (r PhoneNumbersRegulatoryRequirementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[phone_number]
type PhoneNumbersRegulatoryRequirementGetParamsFilter struct {
	// Record type phone number/ phone numbers
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumbersRegulatoryRequirementGetParamsFilter]'s query
// parameters as `url.Values`.
func (r PhoneNumbersRegulatoryRequirementGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
