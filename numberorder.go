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
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// NumberOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberOrderService] method instead.
type NumberOrderService struct {
	Options []option.RequestOption
}

// NewNumberOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumberOrderService(opts ...option.RequestOption) (r NumberOrderService) {
	r = NumberOrderService{}
	r.Options = opts
	return
}

// Creates a phone number order.
func (r *NumberOrderService) New(ctx context.Context, body NumberOrderNewParams, opts ...option.RequestOption) (res *NumberOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing phone number order.
func (r *NumberOrderService) Get(ctx context.Context, numberOrderID string, opts ...option.RequestOption) (res *NumberOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if numberOrderID == "" {
		err = errors.New("missing required number_order_id parameter")
		return
	}
	path := fmt.Sprintf("number_orders/%s", numberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a phone number order.
func (r *NumberOrderService) Update(ctx context.Context, numberOrderID string, body NumberOrderUpdateParams, opts ...option.RequestOption) (res *NumberOrderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if numberOrderID == "" {
		err = errors.New("missing required number_order_id parameter")
		return
	}
	path := fmt.Sprintf("number_orders/%s", numberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of number orders.
func (r *NumberOrderService) List(ctx context.Context, query NumberOrderListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[NumberOrderListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "number_orders"
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

// Get a paginated list of number orders.
func (r *NumberOrderService) ListAutoPaging(ctx context.Context, query NumberOrderListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[NumberOrderListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type NumberOrderWithPhoneNumbers struct {
	ID string `json:"id" format:"uuid"`
	// Identifies the messaging profile associated with the phone number.
	BillingGroupID string `json:"billing_group_id"`
	// Identifies the connection associated with this phone number.
	ConnectionID string `json:"connection_id"`
	// An ISO 8901 datetime string denoting when the number order was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileID string        `json:"messaging_profile_id"`
	PhoneNumbers       []PhoneNumber `json:"phone_numbers"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount int64  `json:"phone_numbers_count"`
	RecordType        string `json:"record_type"`
	// True if all requirements are met for every phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met"`
	// The status of the order.
	//
	// Any of "pending", "success", "failure".
	Status             NumberOrderWithPhoneNumbersStatus `json:"status"`
	SubNumberOrdersIDs []string                          `json:"sub_number_orders_ids"`
	// An ISO 8901 datetime string for when the number order was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingGroupID     respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		MessagingProfileID respjson.Field
		PhoneNumbers       respjson.Field
		PhoneNumbersCount  respjson.Field
		RecordType         respjson.Field
		RequirementsMet    respjson.Field
		Status             respjson.Field
		SubNumberOrdersIDs respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderWithPhoneNumbers) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderWithPhoneNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the order.
type NumberOrderWithPhoneNumbersStatus string

const (
	NumberOrderWithPhoneNumbersStatusPending NumberOrderWithPhoneNumbersStatus = "pending"
	NumberOrderWithPhoneNumbersStatusSuccess NumberOrderWithPhoneNumbersStatus = "success"
	NumberOrderWithPhoneNumbersStatusFailure NumberOrderWithPhoneNumbersStatus = "failure"
)

type PhoneNumber struct {
	ID       string `json:"id" format:"uuid"`
	BundleID string `json:"bundle_id" format:"uuid"`
	// Country code of the phone number
	CountryCode string `json:"country_code"`
	// The ISO 3166-1 alpha-2 country code of the phone number.
	CountryISOAlpha2 string `json:"country_iso_alpha2"`
	PhoneNumber      string `json:"phone_number"`
	// Phone number type
	//
	// Any of "local", "mobile", "national", "shared_cost", "toll_free".
	PhoneNumberType        PhoneNumberPhoneNumberType                            `json:"phone_number_type"`
	RecordType             string                                                `json:"record_type"`
	RegulatoryRequirements []shared.SubNumberOrderRegulatoryRequirementWithValue `json:"regulatory_requirements"`
	// True if all requirements are met for a phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met"`
	// Status of document requirements (if applicable)
	//
	// Any of "pending", "approved", "cancelled", "deleted",
	// "requirement-info-exception", "requirement-info-pending",
	// "requirement-info-under-review".
	RequirementsStatus PhoneNumberRequirementsStatus `json:"requirements_status"`
	// The status of the phone number in the order.
	//
	// Any of "pending", "success", "failure".
	Status PhoneNumberStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BundleID               respjson.Field
		CountryCode            respjson.Field
		CountryISOAlpha2       respjson.Field
		PhoneNumber            respjson.Field
		PhoneNumberType        respjson.Field
		RecordType             respjson.Field
		RegulatoryRequirements respjson.Field
		RequirementsMet        respjson.Field
		RequirementsStatus     respjson.Field
		Status                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Phone number type
type PhoneNumberPhoneNumberType string

const (
	PhoneNumberPhoneNumberTypeLocal      PhoneNumberPhoneNumberType = "local"
	PhoneNumberPhoneNumberTypeMobile     PhoneNumberPhoneNumberType = "mobile"
	PhoneNumberPhoneNumberTypeNational   PhoneNumberPhoneNumberType = "national"
	PhoneNumberPhoneNumberTypeSharedCost PhoneNumberPhoneNumberType = "shared_cost"
	PhoneNumberPhoneNumberTypeTollFree   PhoneNumberPhoneNumberType = "toll_free"
)

// Status of document requirements (if applicable)
type PhoneNumberRequirementsStatus string

const (
	PhoneNumberRequirementsStatusPending                    PhoneNumberRequirementsStatus = "pending"
	PhoneNumberRequirementsStatusApproved                   PhoneNumberRequirementsStatus = "approved"
	PhoneNumberRequirementsStatusCancelled                  PhoneNumberRequirementsStatus = "cancelled"
	PhoneNumberRequirementsStatusDeleted                    PhoneNumberRequirementsStatus = "deleted"
	PhoneNumberRequirementsStatusRequirementInfoException   PhoneNumberRequirementsStatus = "requirement-info-exception"
	PhoneNumberRequirementsStatusRequirementInfoPending     PhoneNumberRequirementsStatus = "requirement-info-pending"
	PhoneNumberRequirementsStatusRequirementInfoUnderReview PhoneNumberRequirementsStatus = "requirement-info-under-review"
)

// The status of the phone number in the order.
type PhoneNumberStatus string

const (
	PhoneNumberStatusPending PhoneNumberStatus = "pending"
	PhoneNumberStatusSuccess PhoneNumberStatus = "success"
	PhoneNumberStatusFailure PhoneNumberStatus = "failure"
)

type NumberOrderNewResponse struct {
	Data NumberOrderWithPhoneNumbers `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderGetResponse struct {
	Data NumberOrderWithPhoneNumbers `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderUpdateResponse struct {
	Data NumberOrderWithPhoneNumbers `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderListResponse struct {
	ID string `json:"id" format:"uuid"`
	// Identifies the messaging profile associated with the phone number.
	BillingGroupID string `json:"billing_group_id"`
	// Identifies the connection associated with this phone number.
	ConnectionID string `json:"connection_id"`
	// An ISO 8901 datetime string denoting when the number order was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileID string                              `json:"messaging_profile_id"`
	PhoneNumbers       []shared.PhoneNumbersJobPhoneNumber `json:"phone_numbers"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount int64  `json:"phone_numbers_count"`
	RecordType        string `json:"record_type"`
	// True if all requirements are met for every phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met"`
	// The status of the order.
	//
	// Any of "pending", "success", "failure".
	Status             NumberOrderListResponseStatus `json:"status"`
	SubNumberOrdersIDs []string                      `json:"sub_number_orders_ids"`
	// An ISO 8901 datetime string for when the number order was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingGroupID     respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		MessagingProfileID respjson.Field
		PhoneNumbers       respjson.Field
		PhoneNumbersCount  respjson.Field
		RecordType         respjson.Field
		RequirementsMet    respjson.Field
		Status             respjson.Field
		SubNumberOrdersIDs respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the order.
type NumberOrderListResponseStatus string

const (
	NumberOrderListResponseStatusPending NumberOrderListResponseStatus = "pending"
	NumberOrderListResponseStatusSuccess NumberOrderListResponseStatus = "success"
	NumberOrderListResponseStatusFailure NumberOrderListResponseStatus = "failure"
)

type NumberOrderNewParams struct {
	// Identifies the billing group associated with the phone number.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero"`
	// Identifies the connection associated with this phone number.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileID param.Opt[string]                 `json:"messaging_profile_id,omitzero"`
	PhoneNumbers       []NumberOrderNewParamsPhoneNumber `json:"phone_numbers,omitzero"`
	paramObj
}

func (r NumberOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PhoneNumber is required.
type NumberOrderNewParamsPhoneNumber struct {
	// e164_phone_number
	PhoneNumber string `json:"phone_number" api:"required"`
	// ID of bundle to associate the number to
	BundleID param.Opt[string] `json:"bundle_id,omitzero"`
	// ID of requirement group to use to satisfy number requirements
	RequirementGroupID param.Opt[string] `json:"requirement_group_id,omitzero"`
	paramObj
}

func (r NumberOrderNewParamsPhoneNumber) MarshalJSON() (data []byte, err error) {
	type shadow NumberOrderNewParamsPhoneNumber
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberOrderNewParamsPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderUpdateParams struct {
	// A customer reference string for customer look ups.
	CustomerReference      param.Opt[string]                  `json:"customer_reference,omitzero"`
	RegulatoryRequirements []UpdateRegulatoryRequirementParam `json:"regulatory_requirements,omitzero"`
	paramObj
}

func (r NumberOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[status],
	// filter[created_at], filter[phone_numbers_count], filter[customer_reference],
	// filter[requirements_met]
	Filter NumberOrderListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberOrderListParams]'s query parameters as `url.Values`.
func (r NumberOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[status],
// filter[created_at], filter[phone_numbers_count], filter[customer_reference],
// filter[requirements_met]
type NumberOrderListParamsFilter struct {
	// Filter number orders via the customer reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter number order with this amount of numbers
	PhoneNumbersCount param.Opt[string] `query:"phone_numbers_count,omitzero" json:"-"`
	// Filter number orders by requirements met.
	RequirementsMet param.Opt[bool] `query:"requirements_met,omitzero" json:"-"`
	// Filter number orders by status.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter number orders by date range.
	CreatedAt NumberOrderListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberOrderListParamsFilter]'s query parameters as
// `url.Values`.
func (r NumberOrderListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter number orders by date range.
type NumberOrderListParamsFilterCreatedAt struct {
	// Filter number orders later than this value.
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Filter number orders earlier than this value.
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberOrderListParamsFilterCreatedAt]'s query parameters as
// `url.Values`.
func (r NumberOrderListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
