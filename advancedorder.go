// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AdvancedOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedOrderService] method instead.
type AdvancedOrderService struct {
	Options []option.RequestOption
}

// NewAdvancedOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdvancedOrderService(opts ...option.RequestOption) (r AdvancedOrderService) {
	r = AdvancedOrderService{}
	r.Options = opts
	return
}

// Create Advanced Order
func (r *AdvancedOrderService) New(ctx context.Context, body AdvancedOrderNewParams, opts ...option.RequestOption) (res *AdvancedOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "advanced_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get Advanced Order
func (r *AdvancedOrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *AdvancedOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("advanced_orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Advanced Orders
func (r *AdvancedOrderService) List(ctx context.Context, opts ...option.RequestOption) (res *AdvancedOrderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "advanced_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Advanced Order
func (r *AdvancedOrderService) UpdateRequirementGroup(ctx context.Context, advancedOrderID string, body AdvancedOrderUpdateRequirementGroupParams, opts ...option.RequestOption) (res *AdvancedOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	if advancedOrderID == "" {
		err = errors.New("missing required advanced-order-id parameter")
		return nil, err
	}
	path := fmt.Sprintf("advanced_orders/%s/requirement_group", advancedOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type AdvancedOrder struct {
	ID                string `json:"id" format:"uuid"`
	AreaCode          string `json:"area_code"`
	Comments          string `json:"comments"`
	CountryCode       string `json:"country_code"`
	CustomerReference string `json:"customer_reference"`
	// Any of "sms", "mms", "voice", "fax", "emergency".
	Features []string `json:"features"`
	Orders   []string `json:"orders" format:"uuid"`
	// Any of "local", "mobile", "toll_free", "shared_cost", "national", "landline".
	PhoneNumberType []string `json:"phone_number_type"`
	Quantity        int64    `json:"quantity"`
	// The ID of the requirement group associated with this advanced order
	RequirementGroupID string `json:"requirement_group_id" format:"uuid"`
	// Any of "pending", "processing", "ordered".
	Status []string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AreaCode           respjson.Field
		Comments           respjson.Field
		CountryCode        respjson.Field
		CustomerReference  respjson.Field
		Features           respjson.Field
		Orders             respjson.Field
		PhoneNumberType    respjson.Field
		Quantity           respjson.Field
		RequirementGroupID respjson.Field
		Status             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedOrder) RawJSON() string { return r.JSON.raw }
func (r *AdvancedOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderRequestParam struct {
	AreaCode          param.Opt[string] `json:"area_code,omitzero"`
	Comments          param.Opt[string] `json:"comments,omitzero"`
	CountryCode       param.Opt[string] `json:"country_code,omitzero"`
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	Quantity          param.Opt[int64]  `json:"quantity,omitzero"`
	// The ID of the requirement group to associate with this advanced order
	RequirementGroupID param.Opt[string] `json:"requirement_group_id,omitzero" format:"uuid"`
	// Any of "sms", "mms", "voice", "fax", "emergency".
	Features []string `json:"features,omitzero"`
	// Any of "local", "mobile", "toll_free", "shared_cost", "national", "landline".
	PhoneNumberType AdvancedOrderRequestPhoneNumberType `json:"phone_number_type,omitzero"`
	paramObj
}

func (r AdvancedOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderRequestPhoneNumberType string

const (
	AdvancedOrderRequestPhoneNumberTypeLocal      AdvancedOrderRequestPhoneNumberType = "local"
	AdvancedOrderRequestPhoneNumberTypeMobile     AdvancedOrderRequestPhoneNumberType = "mobile"
	AdvancedOrderRequestPhoneNumberTypeTollFree   AdvancedOrderRequestPhoneNumberType = "toll_free"
	AdvancedOrderRequestPhoneNumberTypeSharedCost AdvancedOrderRequestPhoneNumberType = "shared_cost"
	AdvancedOrderRequestPhoneNumberTypeNational   AdvancedOrderRequestPhoneNumberType = "national"
	AdvancedOrderRequestPhoneNumberTypeLandline   AdvancedOrderRequestPhoneNumberType = "landline"
)

type AdvancedOrderListResponse struct {
	Data []AdvancedOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *AdvancedOrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderNewParams struct {
	AdvancedOrderRequest AdvancedOrderRequestParam
	paramObj
}

func (r AdvancedOrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AdvancedOrderRequest)
}
func (r *AdvancedOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderUpdateRequirementGroupParams struct {
	AdvancedOrderRequest AdvancedOrderRequestParam
	paramObj
}

func (r AdvancedOrderUpdateRequirementGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AdvancedOrderRequest)
}
func (r *AdvancedOrderUpdateRequirementGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
