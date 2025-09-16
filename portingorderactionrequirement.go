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

// PortingOrderActionRequirementService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderActionRequirementService] method instead.
type PortingOrderActionRequirementService struct {
	Options []option.RequestOption
}

// NewPortingOrderActionRequirementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderActionRequirementService(opts ...option.RequestOption) (r PortingOrderActionRequirementService) {
	r = PortingOrderActionRequirementService{}
	r.Options = opts
	return
}

// Returns a list of action requirements for a specific porting order.
func (r *PortingOrderActionRequirementService) List(ctx context.Context, portingOrderID string, query PortingOrderActionRequirementListParams, opts ...option.RequestOption) (res *PortingOrderActionRequirementListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/action_requirements", portingOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiates a specific action requirement for a porting order.
func (r *PortingOrderActionRequirementService) Initiate(ctx context.Context, id string, params PortingOrderActionRequirementInitiateParams, opts ...option.RequestOption) (res *PortingOrderActionRequirementInitiateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.PortingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/action_requirements/%s/initiate", params.PortingOrderID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PortingOrderActionRequirementListResponse struct {
	Data []PortingOrderActionRequirementListResponseData `json:"data"`
	Meta PaginationMeta                                  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionRequirementListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionRequirementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionRequirementListResponseData struct {
	// Identifies the action requirement
	ID string `json:"id"`
	// The type of action required
	ActionType string `json:"action_type"`
	// Optional URL for the action
	ActionURL string `json:"action_url,nullable"`
	// Reason for cancellation if status is 'cancelled'
	CancelReason string `json:"cancel_reason,nullable"`
	// ISO 8601 formatted date-time indicating when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The ID of the porting order this action requirement belongs to
	PortingOrderID string `json:"porting_order_id"`
	// Identifies the type of the resource
	//
	// Any of "porting_action_requirement".
	RecordType string `json:"record_type"`
	// The ID of the requirement type
	RequirementTypeID string `json:"requirement_type_id"`
	// Current status of the action requirement
	//
	// Any of "created", "pending", "completed", "cancelled", "failed".
	Status string `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ActionType        respjson.Field
		ActionURL         respjson.Field
		CancelReason      respjson.Field
		CreatedAt         respjson.Field
		PortingOrderID    respjson.Field
		RecordType        respjson.Field
		RequirementTypeID respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionRequirementListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionRequirementListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionRequirementInitiateResponse struct {
	Data PortingOrderActionRequirementInitiateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionRequirementInitiateResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionRequirementInitiateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionRequirementInitiateResponseData struct {
	// Identifies the action requirement
	ID string `json:"id"`
	// The type of action required
	ActionType string `json:"action_type"`
	// Optional URL for the action
	ActionURL string `json:"action_url,nullable"`
	// Reason for cancellation if status is 'cancelled'
	CancelReason string `json:"cancel_reason,nullable"`
	// ISO 8601 formatted date-time indicating when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The ID of the porting order this action requirement belongs to
	PortingOrderID string `json:"porting_order_id"`
	// Identifies the type of the resource
	//
	// Any of "porting_action_requirement".
	RecordType string `json:"record_type"`
	// The ID of the requirement type
	RequirementTypeID string `json:"requirement_type_id"`
	// Current status of the action requirement
	//
	// Any of "created", "pending", "completed", "cancelled", "failed".
	Status string `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ActionType        respjson.Field
		ActionURL         respjson.Field
		CancelReason      respjson.Field
		CreatedAt         respjson.Field
		PortingOrderID    respjson.Field
		RecordType        respjson.Field
		RequirementTypeID respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionRequirementInitiateResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionRequirementInitiateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionRequirementListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[id][in][],
	// filter[requirement_type_id], filter[action_type], filter[status]
	Filter PortingOrderActionRequirementListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderActionRequirementListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderActionRequirementListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderActionRequirementListParams]'s query parameters
// as `url.Values`.
func (r PortingOrderActionRequirementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[id][in][],
// filter[requirement_type_id], filter[action_type], filter[status]
type PortingOrderActionRequirementListParamsFilter struct {
	// Filter action requirements by requirement type ID
	RequirementTypeID param.Opt[string] `query:"requirement_type_id,omitzero" format:"uuid" json:"-"`
	// Filter action requirements by a list of IDs
	ID []string `query:"id,omitzero" format:"uuid" json:"-"`
	// Filter action requirements by action type
	//
	// Any of "au_id_verification".
	ActionType string `query:"action_type,omitzero" json:"-"`
	// Filter action requirements by status
	//
	// Any of "created", "pending", "completed", "cancelled", "failed".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderActionRequirementListParamsFilter]'s query
// parameters as `url.Values`.
func (r PortingOrderActionRequirementListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PortingOrderActionRequirementListParamsFilter](
		"action_type", "au_id_verification",
	)
	apijson.RegisterFieldValidator[PortingOrderActionRequirementListParamsFilter](
		"status", "created", "pending", "completed", "cancelled", "failed",
	)
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderActionRequirementListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderActionRequirementListParamsPage]'s query
// parameters as `url.Values`.
func (r PortingOrderActionRequirementListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderActionRequirementListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at", "-created_at", "updated_at", "-updated_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderActionRequirementListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderActionRequirementListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PortingOrderActionRequirementListParamsSort](
		"value", "created_at", "-created_at", "updated_at", "-updated_at",
	)
}

type PortingOrderActionRequirementInitiateParams struct {
	PortingOrderID string `path:"porting_order_id,required" json:"-"`
	// Required information for initiating the action requirement for AU ID
	// verification.
	Params PortingOrderActionRequirementInitiateParamsParams `json:"params,omitzero,required"`
	paramObj
}

func (r PortingOrderActionRequirementInitiateParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderActionRequirementInitiateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderActionRequirementInitiateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required information for initiating the action requirement for AU ID
// verification.
//
// The properties FirstName, LastName are required.
type PortingOrderActionRequirementInitiateParamsParams struct {
	// The first name of the person that will perform the verification flow.
	FirstName string `json:"first_name,required"`
	// The last name of the person that will perform the verification flow.
	LastName string `json:"last_name,required"`
	paramObj
}

func (r PortingOrderActionRequirementInitiateParamsParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderActionRequirementInitiateParamsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderActionRequirementInitiateParamsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
