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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// NumberBlockOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberBlockOrderService] method instead.
type NumberBlockOrderService struct {
	Options []option.RequestOption
}

// NewNumberBlockOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumberBlockOrderService(opts ...option.RequestOption) (r NumberBlockOrderService) {
	r = NumberBlockOrderService{}
	r.Options = opts
	return
}

// Creates a phone number block order.
func (r *NumberBlockOrderService) New(ctx context.Context, body NumberBlockOrderNewParams, opts ...option.RequestOption) (res *NumberBlockOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "number_block_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing phone number block order.
func (r *NumberBlockOrderService) Get(ctx context.Context, numberBlockOrderID string, opts ...option.RequestOption) (res *NumberBlockOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if numberBlockOrderID == "" {
		err = errors.New("missing required number_block_order_id parameter")
		return
	}
	path := fmt.Sprintf("number_block_orders/%s", numberBlockOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a paginated list of number block orders.
func (r *NumberBlockOrderService) List(ctx context.Context, query NumberBlockOrderListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[NumberBlockOrder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "number_block_orders"
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

// Get a paginated list of number block orders.
func (r *NumberBlockOrderService) ListAutoPaging(ctx context.Context, query NumberBlockOrderListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[NumberBlockOrder] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

type NumberBlockOrder struct {
	ID string `json:"id" format:"uuid"`
	// Identifies the connection associated to all numbers in the phone number block.
	ConnectionID string `json:"connection_id"`
	// An ISO 8901 datetime string denoting when the number order was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Identifies the messaging profile associated to all numbers in the phone number
	// block.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount int64 `json:"phone_numbers_count"`
	// The phone number range included in the block.
	Range      int64  `json:"range"`
	RecordType string `json:"record_type"`
	// True if all requirements are met for every phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met"`
	// Starting phone number block
	StartingNumber string `json:"starting_number"`
	// The status of the order.
	//
	// Any of "pending", "success", "failure".
	Status NumberBlockOrderStatus `json:"status"`
	// An ISO 8901 datetime string for when the number order was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		MessagingProfileID respjson.Field
		PhoneNumbersCount  respjson.Field
		Range              respjson.Field
		RecordType         respjson.Field
		RequirementsMet    respjson.Field
		StartingNumber     respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberBlockOrder) RawJSON() string { return r.JSON.raw }
func (r *NumberBlockOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the order.
type NumberBlockOrderStatus string

const (
	NumberBlockOrderStatusPending NumberBlockOrderStatus = "pending"
	NumberBlockOrderStatusSuccess NumberBlockOrderStatus = "success"
	NumberBlockOrderStatusFailure NumberBlockOrderStatus = "failure"
)

type NumberBlockOrderNewResponse struct {
	Data NumberBlockOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberBlockOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberBlockOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberBlockOrderGetResponse struct {
	Data NumberBlockOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberBlockOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberBlockOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberBlockOrderNewParams struct {
	// The phone number range included in the block.
	Range int64 `json:"range,required"`
	// Starting phone number block
	StartingNumber string `json:"starting_number,required"`
	// Identifies the connection associated with this phone number.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	paramObj
}

func (r NumberBlockOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberBlockOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberBlockOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberBlockOrderListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[status],
	// filter[created_at], filter[phone_numbers.starting_number]
	Filter NumberBlockOrderListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page NumberBlockOrderListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberBlockOrderListParams]'s query parameters as
// `url.Values`.
func (r NumberBlockOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[status],
// filter[created_at], filter[phone_numbers.starting_number]
type NumberBlockOrderListParamsFilter struct {
	// Filter number block orders having these phone numbers.
	PhoneNumbersStartingNumber param.Opt[string] `query:"phone_numbers.starting_number,omitzero" json:"-"`
	// Filter number block orders by status.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter number block orders by date range.
	CreatedAt NumberBlockOrderListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberBlockOrderListParamsFilter]'s query parameters as
// `url.Values`.
func (r NumberBlockOrderListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter number block orders by date range.
type NumberBlockOrderListParamsFilterCreatedAt struct {
	// Filter number block orders later than this value.
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Filter number block orders earlier than this value.
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberBlockOrderListParamsFilterCreatedAt]'s query
// parameters as `url.Values`.
func (r NumberBlockOrderListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type NumberBlockOrderListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberBlockOrderListParamsPage]'s query parameters as
// `url.Values`.
func (r NumberBlockOrderListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
