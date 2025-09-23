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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
	"github.com/team-telnyx/telnyx-go/shared"
)

// NumberOrderPhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberOrderPhoneNumberService] method instead.
type NumberOrderPhoneNumberService struct {
	Options []option.RequestOption
}

// NewNumberOrderPhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumberOrderPhoneNumberService(opts ...option.RequestOption) (r NumberOrderPhoneNumberService) {
	r = NumberOrderPhoneNumberService{}
	r.Options = opts
	return
}

// Get an existing phone number in number order.
func (r *NumberOrderPhoneNumberService) Get(ctx context.Context, numberOrderPhoneNumberID string, opts ...option.RequestOption) (res *NumberOrderPhoneNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if numberOrderPhoneNumberID == "" {
		err = errors.New("missing required number_order_phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("number_order_phone_numbers/%s", numberOrderPhoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of phone numbers associated to orders.
func (r *NumberOrderPhoneNumberService) List(ctx context.Context, query NumberOrderPhoneNumberListParams, opts ...option.RequestOption) (res *NumberOrderPhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "number_order_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update requirement group for a phone number order
func (r *NumberOrderPhoneNumberService) UpdateRequirementGroup(ctx context.Context, id string, body NumberOrderPhoneNumberUpdateRequirementGroupParams, opts ...option.RequestOption) (res *NumberOrderPhoneNumberUpdateRequirementGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("number_order_phone_numbers/%s/requirement_group", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates requirements for a single phone number within a number order.
func (r *NumberOrderPhoneNumberService) UpdateRequirements(ctx context.Context, numberOrderPhoneNumberID string, body NumberOrderPhoneNumberUpdateRequirementsParams, opts ...option.RequestOption) (res *NumberOrderPhoneNumberUpdateRequirementsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if numberOrderPhoneNumberID == "" {
		err = errors.New("missing required number_order_phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("number_order_phone_numbers/%s", numberOrderPhoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type NumberOrderPhoneNumber struct {
	ID             string    `json:"id" format:"uuid"`
	BundleID       string    `json:"bundle_id" format:"uuid"`
	CountryCode    string    `json:"country_code"`
	Deadline       time.Time `json:"deadline" format:"date-time"`
	IsBlockNumber  bool      `json:"is_block_number"`
	Locality       string    `json:"locality"`
	OrderRequestID string    `json:"order_request_id" format:"uuid"`
	PhoneNumber    string    `json:"phone_number"`
	// Any of "local", "toll_free", "mobile", "national", "shared_cost", "landline".
	PhoneNumberType        NumberOrderPhoneNumberPhoneNumberType                 `json:"phone_number_type"`
	RecordType             string                                                `json:"record_type"`
	RegulatoryRequirements []shared.SubNumberOrderRegulatoryRequirementWithValue `json:"regulatory_requirements"`
	// True if all requirements are met for a phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met"`
	// Status of requirements (if applicable)
	//
	// Any of "pending", "approved", "cancelled", "deleted",
	// "requirement-info-exception", "requirement-info-pending",
	// "requirement-info-under-review".
	RequirementsStatus NumberOrderPhoneNumberRequirementsStatus `json:"requirements_status"`
	// The status of the phone number in the order.
	//
	// Any of "pending", "success", "failure".
	Status           NumberOrderPhoneNumberStatus `json:"status"`
	SubNumberOrderID string                       `json:"sub_number_order_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BundleID               respjson.Field
		CountryCode            respjson.Field
		Deadline               respjson.Field
		IsBlockNumber          respjson.Field
		Locality               respjson.Field
		OrderRequestID         respjson.Field
		PhoneNumber            respjson.Field
		PhoneNumberType        respjson.Field
		RecordType             respjson.Field
		RegulatoryRequirements respjson.Field
		RequirementsMet        respjson.Field
		RequirementsStatus     respjson.Field
		Status                 respjson.Field
		SubNumberOrderID       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberPhoneNumberType string

const (
	NumberOrderPhoneNumberPhoneNumberTypeLocal      NumberOrderPhoneNumberPhoneNumberType = "local"
	NumberOrderPhoneNumberPhoneNumberTypeTollFree   NumberOrderPhoneNumberPhoneNumberType = "toll_free"
	NumberOrderPhoneNumberPhoneNumberTypeMobile     NumberOrderPhoneNumberPhoneNumberType = "mobile"
	NumberOrderPhoneNumberPhoneNumberTypeNational   NumberOrderPhoneNumberPhoneNumberType = "national"
	NumberOrderPhoneNumberPhoneNumberTypeSharedCost NumberOrderPhoneNumberPhoneNumberType = "shared_cost"
	NumberOrderPhoneNumberPhoneNumberTypeLandline   NumberOrderPhoneNumberPhoneNumberType = "landline"
)

// Status of requirements (if applicable)
type NumberOrderPhoneNumberRequirementsStatus string

const (
	NumberOrderPhoneNumberRequirementsStatusPending                    NumberOrderPhoneNumberRequirementsStatus = "pending"
	NumberOrderPhoneNumberRequirementsStatusApproved                   NumberOrderPhoneNumberRequirementsStatus = "approved"
	NumberOrderPhoneNumberRequirementsStatusCancelled                  NumberOrderPhoneNumberRequirementsStatus = "cancelled"
	NumberOrderPhoneNumberRequirementsStatusDeleted                    NumberOrderPhoneNumberRequirementsStatus = "deleted"
	NumberOrderPhoneNumberRequirementsStatusRequirementInfoException   NumberOrderPhoneNumberRequirementsStatus = "requirement-info-exception"
	NumberOrderPhoneNumberRequirementsStatusRequirementInfoPending     NumberOrderPhoneNumberRequirementsStatus = "requirement-info-pending"
	NumberOrderPhoneNumberRequirementsStatusRequirementInfoUnderReview NumberOrderPhoneNumberRequirementsStatus = "requirement-info-under-review"
)

// The status of the phone number in the order.
type NumberOrderPhoneNumberStatus string

const (
	NumberOrderPhoneNumberStatusPending NumberOrderPhoneNumberStatus = "pending"
	NumberOrderPhoneNumberStatusSuccess NumberOrderPhoneNumberStatus = "success"
	NumberOrderPhoneNumberStatusFailure NumberOrderPhoneNumberStatus = "failure"
)

type UpdateRegulatoryRequirementParam struct {
	// The value of the requirement. For address and document requirements, this should
	// be the ID of the resource. For textual, this should be the value of the
	// requirement.
	FieldValue param.Opt[string] `json:"field_value,omitzero"`
	// Unique id for a requirement.
	RequirementID param.Opt[string] `json:"requirement_id,omitzero" format:"uuid"`
	paramObj
}

func (r UpdateRegulatoryRequirementParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateRegulatoryRequirementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateRegulatoryRequirementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberGetResponse struct {
	Data NumberOrderPhoneNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderPhoneNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberListResponse struct {
	Data []NumberOrderPhoneNumber `json:"data"`
	Meta PaginationMeta           `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberUpdateRequirementGroupResponse struct {
	Data NumberOrderPhoneNumberUpdateRequirementGroupResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumberUpdateRequirementGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderPhoneNumberUpdateRequirementGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberUpdateRequirementGroupResponseData struct {
	ID                     string                                                                          `json:"id" format:"uuid"`
	BundleID               string                                                                          `json:"bundle_id,nullable" format:"uuid"`
	CountryCode            string                                                                          `json:"country_code"`
	Deadline               time.Time                                                                       `json:"deadline" format:"date-time"`
	IsBlockNumber          bool                                                                            `json:"is_block_number"`
	Locality               string                                                                          `json:"locality"`
	OrderRequestID         string                                                                          `json:"order_request_id" format:"uuid"`
	PhoneNumber            string                                                                          `json:"phone_number"`
	PhoneNumberType        string                                                                          `json:"phone_number_type"`
	RecordType             string                                                                          `json:"record_type"`
	RegulatoryRequirements []NumberOrderPhoneNumberUpdateRequirementGroupResponseDataRegulatoryRequirement `json:"regulatory_requirements"`
	RequirementsMet        bool                                                                            `json:"requirements_met"`
	RequirementsStatus     string                                                                          `json:"requirements_status"`
	Status                 string                                                                          `json:"status"`
	SubNumberOrderID       string                                                                          `json:"sub_number_order_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BundleID               respjson.Field
		CountryCode            respjson.Field
		Deadline               respjson.Field
		IsBlockNumber          respjson.Field
		Locality               respjson.Field
		OrderRequestID         respjson.Field
		PhoneNumber            respjson.Field
		PhoneNumberType        respjson.Field
		RecordType             respjson.Field
		RegulatoryRequirements respjson.Field
		RequirementsMet        respjson.Field
		RequirementsStatus     respjson.Field
		Status                 respjson.Field
		SubNumberOrderID       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumberUpdateRequirementGroupResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderPhoneNumberUpdateRequirementGroupResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberUpdateRequirementGroupResponseDataRegulatoryRequirement struct {
	FieldType     string `json:"field_type"`
	FieldValue    string `json:"field_value"`
	RequirementID string `json:"requirement_id" format:"uuid"`
	Status        string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType     respjson.Field
		FieldValue    respjson.Field
		RequirementID respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumberUpdateRequirementGroupResponseDataRegulatoryRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *NumberOrderPhoneNumberUpdateRequirementGroupResponseDataRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberUpdateRequirementsResponse struct {
	Data NumberOrderPhoneNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderPhoneNumberUpdateRequirementsResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderPhoneNumberUpdateRequirementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[country_code]
	Filter NumberOrderPhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberOrderPhoneNumberListParams]'s query parameters as
// `url.Values`.
func (r NumberOrderPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[country_code]
type NumberOrderPhoneNumberListParamsFilter struct {
	// Country code of the order phone number.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberOrderPhoneNumberListParamsFilter]'s query parameters
// as `url.Values`.
func (r NumberOrderPhoneNumberListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NumberOrderPhoneNumberUpdateRequirementGroupParams struct {
	// The ID of the requirement group to associate
	RequirementGroupID string `json:"requirement_group_id,required" format:"uuid"`
	paramObj
}

func (r NumberOrderPhoneNumberUpdateRequirementGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberOrderPhoneNumberUpdateRequirementGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberOrderPhoneNumberUpdateRequirementGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderPhoneNumberUpdateRequirementsParams struct {
	RegulatoryRequirements []UpdateRegulatoryRequirementParam `json:"regulatory_requirements,omitzero"`
	paramObj
}

func (r NumberOrderPhoneNumberUpdateRequirementsParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberOrderPhoneNumberUpdateRequirementsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberOrderPhoneNumberUpdateRequirementsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
