// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
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
	opts = append(r.Options[:], opts...)
	path := "advanced_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Advanced Order
func (r *AdvancedOrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *AdvancedOrderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("advanced_orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Advanced Order
func (r *AdvancedOrderService) Update(ctx context.Context, orderID string, body AdvancedOrderUpdateParams, opts ...option.RequestOption) (res *AdvancedOrderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("advanced_orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Advanced Orders
func (r *AdvancedOrderService) List(ctx context.Context, opts ...option.RequestOption) (res *AdvancedOrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "advanced_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AdvancedOrderNewResponse = any

type AdvancedOrderGetResponse = any

type AdvancedOrderUpdateResponse = any

type AdvancedOrderListResponse = any

type AdvancedOrderNewParams struct {
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
	PhoneNumberType AdvancedOrderNewParamsPhoneNumberType `json:"phone_number_type,omitzero"`
	paramObj
}

func (r AdvancedOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderNewParamsPhoneNumberType string

const (
	AdvancedOrderNewParamsPhoneNumberTypeLocal      AdvancedOrderNewParamsPhoneNumberType = "local"
	AdvancedOrderNewParamsPhoneNumberTypeMobile     AdvancedOrderNewParamsPhoneNumberType = "mobile"
	AdvancedOrderNewParamsPhoneNumberTypeTollFree   AdvancedOrderNewParamsPhoneNumberType = "toll_free"
	AdvancedOrderNewParamsPhoneNumberTypeSharedCost AdvancedOrderNewParamsPhoneNumberType = "shared_cost"
	AdvancedOrderNewParamsPhoneNumberTypeNational   AdvancedOrderNewParamsPhoneNumberType = "national"
	AdvancedOrderNewParamsPhoneNumberTypeLandline   AdvancedOrderNewParamsPhoneNumberType = "landline"
)

type AdvancedOrderUpdateParams struct {
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
	PhoneNumberType AdvancedOrderUpdateParamsPhoneNumberType `json:"phone_number_type,omitzero"`
	paramObj
}

func (r AdvancedOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedOrderUpdateParamsPhoneNumberType string

const (
	AdvancedOrderUpdateParamsPhoneNumberTypeLocal      AdvancedOrderUpdateParamsPhoneNumberType = "local"
	AdvancedOrderUpdateParamsPhoneNumberTypeMobile     AdvancedOrderUpdateParamsPhoneNumberType = "mobile"
	AdvancedOrderUpdateParamsPhoneNumberTypeTollFree   AdvancedOrderUpdateParamsPhoneNumberType = "toll_free"
	AdvancedOrderUpdateParamsPhoneNumberTypeSharedCost AdvancedOrderUpdateParamsPhoneNumberType = "shared_cost"
	AdvancedOrderUpdateParamsPhoneNumberTypeNational   AdvancedOrderUpdateParamsPhoneNumberType = "national"
	AdvancedOrderUpdateParamsPhoneNumberTypeLandline   AdvancedOrderUpdateParamsPhoneNumberType = "landline"
)
