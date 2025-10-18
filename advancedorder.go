// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
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
func (r *AdvancedOrderService) New(ctx context.Context, body AdvancedOrderNewParams, opts ...option.RequestOption) (res *AdvancedOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "advanced_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Advanced Order
func (r *AdvancedOrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *AdvancedOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("advanced_orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Advanced Orders
func (r *AdvancedOrderService) List(ctx context.Context, opts ...option.RequestOption) (res *AdvancedOrderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "advanced_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Advanced Order
func (r *AdvancedOrderService) UpdateRequirementGroup(ctx context.Context, advancedOrderID string, body AdvancedOrderUpdateRequirementGroupParams, opts ...option.RequestOption) (res *AdvancedOrderUpdateRequirementGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if advancedOrderID == "" {
		err = errors.New("missing required advanced-order-id parameter")
		return
	}
	path := fmt.Sprintf("advanced_orders/%s/requirement_group", advancedOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AdvancedOrderParam struct {
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
	PhoneNumberType AdvancedOrderPhoneNumberType `json:"phone_number_type,omitzero"`
	paramObj
}

func (r AdvancedOrderParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedOrderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedOrderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderPhoneNumberType string

const (
	AdvancedOrderPhoneNumberTypeLocal      AdvancedOrderPhoneNumberType = "local"
	AdvancedOrderPhoneNumberTypeMobile     AdvancedOrderPhoneNumberType = "mobile"
	AdvancedOrderPhoneNumberTypeTollFree   AdvancedOrderPhoneNumberType = "toll_free"
	AdvancedOrderPhoneNumberTypeSharedCost AdvancedOrderPhoneNumberType = "shared_cost"
	AdvancedOrderPhoneNumberTypeNational   AdvancedOrderPhoneNumberType = "national"
	AdvancedOrderPhoneNumberTypeLandline   AdvancedOrderPhoneNumberType = "landline"
)

type AdvancedOrderNewResponse = any

type AdvancedOrderGetResponse = any

type AdvancedOrderListResponse = any

type AdvancedOrderUpdateRequirementGroupResponse = any

type AdvancedOrderNewParams struct {
	AdvancedOrder AdvancedOrderParam
	paramObj
}

func (r AdvancedOrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AdvancedOrder)
}
func (r *AdvancedOrderNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AdvancedOrder)
}

type AdvancedOrderUpdateRequirementGroupParams struct {
	AdvancedOrder AdvancedOrderParam
	paramObj
}

func (r AdvancedOrderUpdateRequirementGroupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AdvancedOrder)
}
func (r *AdvancedOrderUpdateRequirementGroupParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AdvancedOrder)
}
