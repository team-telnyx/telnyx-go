// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// SubNumberOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubNumberOrderService] method instead.
type SubNumberOrderService struct {
	Options []option.RequestOption
}

// NewSubNumberOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubNumberOrderService(opts ...option.RequestOption) (r SubNumberOrderService) {
	r = SubNumberOrderService{}
	r.Options = opts
	return
}

// Get an existing sub number order.
func (r *SubNumberOrderService) Get(ctx context.Context, subNumberOrderID string, query SubNumberOrderGetParams, opts ...option.RequestOption) (res *SubNumberOrderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subNumberOrderID == "" {
		err = errors.New("missing required sub_number_order_id parameter")
		return
	}
	path := fmt.Sprintf("sub_number_orders/%s", subNumberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates a sub number order.
func (r *SubNumberOrderService) Update(ctx context.Context, subNumberOrderID string, body SubNumberOrderUpdateParams, opts ...option.RequestOption) (res *SubNumberOrderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subNumberOrderID == "" {
		err = errors.New("missing required sub_number_order_id parameter")
		return
	}
	path := fmt.Sprintf("sub_number_orders/%s", subNumberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of sub number orders.
func (r *SubNumberOrderService) List(ctx context.Context, query SubNumberOrderListParams, opts ...option.RequestOption) (res *SubNumberOrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sub_number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allows you to cancel a sub number order in 'pending' status.
func (r *SubNumberOrderService) Cancel(ctx context.Context, subNumberOrderID string, opts ...option.RequestOption) (res *SubNumberOrderCancelResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subNumberOrderID == "" {
		err = errors.New("missing required sub_number_order_id parameter")
		return
	}
	path := fmt.Sprintf("sub_number_orders/%s/cancel", subNumberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Update requirement group for a sub number order
func (r *SubNumberOrderService) UpdateRequirementGroup(ctx context.Context, id string, body SubNumberOrderUpdateRequirementGroupParams, opts ...option.RequestOption) (res *SubNumberOrderUpdateRequirementGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sub_number_orders/%s/requirement_group", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SubNumberOrder struct {
	ID          string `json:"id" format:"uuid"`
	CountryCode string `json:"country_code"`
	// An ISO 8901 datetime string denoting when the number order was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// True if the sub number order is a block sub number order
	IsBlockSubNumberOrder bool   `json:"is_block_sub_number_order"`
	OrderRequestID        string `json:"order_request_id" format:"uuid"`
	// Any of "local", "toll_free", "mobile", "national", "shared_cost", "landline".
	PhoneNumberType SubNumberOrderPhoneNumberType `json:"phone_number_type"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount      int64                                 `json:"phone_numbers_count"`
	RecordType             string                                `json:"record_type"`
	RegulatoryRequirements []SubNumberOrderRegulatoryRequirement `json:"regulatory_requirements"`
	// True if all requirements are met for every phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met"`
	// The status of the order.
	//
	// Any of "pending", "success", "failure".
	Status SubNumberOrderStatus `json:"status"`
	// An ISO 8901 datetime string for when the number order was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	UserID    string    `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CountryCode            respjson.Field
		CreatedAt              respjson.Field
		CustomerReference      respjson.Field
		IsBlockSubNumberOrder  respjson.Field
		OrderRequestID         respjson.Field
		PhoneNumberType        respjson.Field
		PhoneNumbersCount      respjson.Field
		RecordType             respjson.Field
		RegulatoryRequirements respjson.Field
		RequirementsMet        respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		UserID                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrder) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderPhoneNumberType string

const (
	SubNumberOrderPhoneNumberTypeLocal      SubNumberOrderPhoneNumberType = "local"
	SubNumberOrderPhoneNumberTypeTollFree   SubNumberOrderPhoneNumberType = "toll_free"
	SubNumberOrderPhoneNumberTypeMobile     SubNumberOrderPhoneNumberType = "mobile"
	SubNumberOrderPhoneNumberTypeNational   SubNumberOrderPhoneNumberType = "national"
	SubNumberOrderPhoneNumberTypeSharedCost SubNumberOrderPhoneNumberType = "shared_cost"
	SubNumberOrderPhoneNumberTypeLandline   SubNumberOrderPhoneNumberType = "landline"
)

// The status of the order.
type SubNumberOrderStatus string

const (
	SubNumberOrderStatusPending SubNumberOrderStatus = "pending"
	SubNumberOrderStatusSuccess SubNumberOrderStatus = "success"
	SubNumberOrderStatusFailure SubNumberOrderStatus = "failure"
)

type SubNumberOrderRegulatoryRequirement struct {
	// Any of "textual", "datetime", "address", "document".
	FieldType  SubNumberOrderRegulatoryRequirementFieldType `json:"field_type"`
	RecordType string                                       `json:"record_type"`
	// Unique id for a requirement.
	RequirementID string `json:"requirement_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType     respjson.Field
		RecordType    respjson.Field
		RequirementID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderRegulatoryRequirement) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderRegulatoryRequirementFieldType string

const (
	SubNumberOrderRegulatoryRequirementFieldTypeTextual  SubNumberOrderRegulatoryRequirementFieldType = "textual"
	SubNumberOrderRegulatoryRequirementFieldTypeDatetime SubNumberOrderRegulatoryRequirementFieldType = "datetime"
	SubNumberOrderRegulatoryRequirementFieldTypeAddress  SubNumberOrderRegulatoryRequirementFieldType = "address"
	SubNumberOrderRegulatoryRequirementFieldTypeDocument SubNumberOrderRegulatoryRequirementFieldType = "document"
)

type SubNumberOrderGetResponse struct {
	Data SubNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderUpdateResponse struct {
	Data SubNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderListResponse struct {
	Data []SubNumberOrder `json:"data"`
	Meta PaginationMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderCancelResponse struct {
	Data SubNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderUpdateRequirementGroupResponse struct {
	Data SubNumberOrderUpdateRequirementGroupResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderUpdateRequirementGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderUpdateRequirementGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderUpdateRequirementGroupResponseData struct {
	ID                     string                                                                  `json:"id" format:"uuid"`
	CountryCode            string                                                                  `json:"country_code"`
	CreatedAt              time.Time                                                               `json:"created_at" format:"date-time"`
	CustomerReference      string                                                                  `json:"customer_reference"`
	IsBlockSubNumberOrder  bool                                                                    `json:"is_block_sub_number_order"`
	OrderRequestID         string                                                                  `json:"order_request_id" format:"uuid"`
	PhoneNumberType        string                                                                  `json:"phone_number_type"`
	PhoneNumbers           []SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumber           `json:"phone_numbers"`
	PhoneNumbersCount      int64                                                                   `json:"phone_numbers_count"`
	RecordType             string                                                                  `json:"record_type"`
	RegulatoryRequirements []SubNumberOrderUpdateRequirementGroupResponseDataRegulatoryRequirement `json:"regulatory_requirements"`
	RequirementsMet        bool                                                                    `json:"requirements_met"`
	Status                 string                                                                  `json:"status"`
	UpdatedAt              time.Time                                                               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CountryCode            respjson.Field
		CreatedAt              respjson.Field
		CustomerReference      respjson.Field
		IsBlockSubNumberOrder  respjson.Field
		OrderRequestID         respjson.Field
		PhoneNumberType        respjson.Field
		PhoneNumbers           respjson.Field
		PhoneNumbersCount      respjson.Field
		RecordType             respjson.Field
		RegulatoryRequirements respjson.Field
		RequirementsMet        respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderUpdateRequirementGroupResponseData) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderUpdateRequirementGroupResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumber struct {
	ID                     string                                                                             `json:"id" format:"uuid"`
	BundleID               string                                                                             `json:"bundle_id" format:"uuid"`
	CountryCode            string                                                                             `json:"country_code"`
	PhoneNumber            string                                                                             `json:"phone_number"`
	PhoneNumberType        string                                                                             `json:"phone_number_type"`
	RecordType             string                                                                             `json:"record_type"`
	RegulatoryRequirements []SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumberRegulatoryRequirement `json:"regulatory_requirements"`
	RequirementsMet        bool                                                                               `json:"requirements_met"`
	RequirementsStatus     string                                                                             `json:"requirements_status"`
	Status                 string                                                                             `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BundleID               respjson.Field
		CountryCode            respjson.Field
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
func (r SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumberRegulatoryRequirement struct {
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
func (r SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumberRegulatoryRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *SubNumberOrderUpdateRequirementGroupResponseDataPhoneNumberRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderUpdateRequirementGroupResponseDataRegulatoryRequirement struct {
	FieldType     string `json:"field_type"`
	RecordType    string `json:"record_type"`
	RequirementID string `json:"requirement_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType     respjson.Field
		RecordType    respjson.Field
		RequirementID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderUpdateRequirementGroupResponseDataRegulatoryRequirement) RawJSON() string {
	return r.JSON.raw
}
func (r *SubNumberOrderUpdateRequirementGroupResponseDataRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[include_phone_numbers]
	Filter SubNumberOrderGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubNumberOrderGetParams]'s query parameters as
// `url.Values`.
func (r SubNumberOrderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[include_phone_numbers]
type SubNumberOrderGetParamsFilter struct {
	// Include the first 50 phone number objects in the results
	IncludePhoneNumbers param.Opt[bool] `query:"include_phone_numbers,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubNumberOrderGetParamsFilter]'s query parameters as
// `url.Values`.
func (r SubNumberOrderGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubNumberOrderUpdateParams struct {
	RegulatoryRequirements []UpdateRegulatoryRequirementParam `json:"regulatory_requirements,omitzero"`
	paramObj
}

func (r SubNumberOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SubNumberOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubNumberOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[status],
	// filter[order_request_id], filter[country_code], filter[phone_number_type],
	// filter[phone_numbers_count]
	Filter SubNumberOrderListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubNumberOrderListParams]'s query parameters as
// `url.Values`.
func (r SubNumberOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[status],
// filter[order_request_id], filter[country_code], filter[phone_number_type],
// filter[phone_numbers_count]
type SubNumberOrderListParamsFilter struct {
	// ISO alpha-2 country code.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// ID of the number order the sub number order belongs to
	OrderRequestID param.Opt[string] `query:"order_request_id,omitzero" format:"uuid" json:"-"`
	// Phone Number Type
	PhoneNumberType param.Opt[string] `query:"phone_number_type,omitzero" json:"-"`
	// Amount of numbers in the sub number order
	PhoneNumbersCount param.Opt[int64] `query:"phone_numbers_count,omitzero" json:"-"`
	// Filter sub number orders by status.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubNumberOrderListParamsFilter]'s query parameters as
// `url.Values`.
func (r SubNumberOrderListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubNumberOrderUpdateRequirementGroupParams struct {
	// The ID of the requirement group to associate
	RequirementGroupID string `json:"requirement_group_id,required" format:"uuid"`
	paramObj
}

func (r SubNumberOrderUpdateRequirementGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow SubNumberOrderUpdateRequirementGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubNumberOrderUpdateRequirementGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
