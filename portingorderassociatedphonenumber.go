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

// PortingOrderAssociatedPhoneNumberService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderAssociatedPhoneNumberService] method instead.
type PortingOrderAssociatedPhoneNumberService struct {
	Options []option.RequestOption
}

// NewPortingOrderAssociatedPhoneNumberService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderAssociatedPhoneNumberService(opts ...option.RequestOption) (r PortingOrderAssociatedPhoneNumberService) {
	r = PortingOrderAssociatedPhoneNumberService{}
	r.Options = opts
	return
}

// Creates a new associated phone number for a porting order. This is used for
// partial porting in GB to specify which phone numbers should be kept or
// disconnected.
func (r *PortingOrderAssociatedPhoneNumberService) New(ctx context.Context, portingOrderID string, body PortingOrderAssociatedPhoneNumberNewParams, opts ...option.RequestOption) (res *PortingOrderAssociatedPhoneNumberNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/associated_phone_numbers", portingOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of all associated phone numbers for a porting order. Associated
// phone numbers are used for partial porting in GB to specify which phone numbers
// should be kept or disconnected.
func (r *PortingOrderAssociatedPhoneNumberService) List(ctx context.Context, portingOrderID string, query PortingOrderAssociatedPhoneNumberListParams, opts ...option.RequestOption) (res *PortingOrderAssociatedPhoneNumberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/associated_phone_numbers", portingOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an associated phone number from a porting order.
func (r *PortingOrderAssociatedPhoneNumberService) Delete(ctx context.Context, id string, body PortingOrderAssociatedPhoneNumberDeleteParams, opts ...option.RequestOption) (res *PortingOrderAssociatedPhoneNumberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.PortingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/associated_phone_numbers/%s", body.PortingOrderID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PortingAssociatedPhoneNumber struct {
	// Uniquely identifies this associated phone number.
	ID string `json:"id" format:"uuid"`
	// Specifies the action to take with this phone number during partial porting.
	//
	// Any of "keep", "disconnect".
	Action PortingAssociatedPhoneNumberAction `json:"action"`
	// Specifies the country code for this associated phone number. It is a two-letter
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies the phone number range for this associated phone number.
	PhoneNumberRange PortingAssociatedPhoneNumberPhoneNumberRange `json:"phone_number_range"`
	// Specifies the phone number type for this associated phone number.
	//
	// Any of "landline", "local", "mobile", "national", "shared_cost", "toll_free".
	PhoneNumberType PortingAssociatedPhoneNumberPhoneNumberType `json:"phone_number_type"`
	// Identifies the porting order associated with this phone number.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Action           respjson.Field
		CountryCode      respjson.Field
		CreatedAt        respjson.Field
		PhoneNumberRange respjson.Field
		PhoneNumberType  respjson.Field
		PortingOrderID   respjson.Field
		RecordType       respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingAssociatedPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PortingAssociatedPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the action to take with this phone number during partial porting.
type PortingAssociatedPhoneNumberAction string

const (
	PortingAssociatedPhoneNumberActionKeep       PortingAssociatedPhoneNumberAction = "keep"
	PortingAssociatedPhoneNumberActionDisconnect PortingAssociatedPhoneNumberAction = "disconnect"
)

// Specifies the phone number range for this associated phone number.
type PortingAssociatedPhoneNumberPhoneNumberRange struct {
	// Specifies the end of the phone number range for this associated phone number.
	EndAt string `json:"end_at"`
	// Specifies the start of the phone number range for this associated phone number.
	StartAt string `json:"start_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndAt       respjson.Field
		StartAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingAssociatedPhoneNumberPhoneNumberRange) RawJSON() string { return r.JSON.raw }
func (r *PortingAssociatedPhoneNumberPhoneNumberRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the phone number type for this associated phone number.
type PortingAssociatedPhoneNumberPhoneNumberType string

const (
	PortingAssociatedPhoneNumberPhoneNumberTypeLandline   PortingAssociatedPhoneNumberPhoneNumberType = "landline"
	PortingAssociatedPhoneNumberPhoneNumberTypeLocal      PortingAssociatedPhoneNumberPhoneNumberType = "local"
	PortingAssociatedPhoneNumberPhoneNumberTypeMobile     PortingAssociatedPhoneNumberPhoneNumberType = "mobile"
	PortingAssociatedPhoneNumberPhoneNumberTypeNational   PortingAssociatedPhoneNumberPhoneNumberType = "national"
	PortingAssociatedPhoneNumberPhoneNumberTypeSharedCost PortingAssociatedPhoneNumberPhoneNumberType = "shared_cost"
	PortingAssociatedPhoneNumberPhoneNumberTypeTollFree   PortingAssociatedPhoneNumberPhoneNumberType = "toll_free"
)

type PortingOrderAssociatedPhoneNumberNewResponse struct {
	Data PortingAssociatedPhoneNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAssociatedPhoneNumberNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAssociatedPhoneNumberNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAssociatedPhoneNumberListResponse struct {
	Data []PortingAssociatedPhoneNumber `json:"data"`
	Meta PaginationMeta                 `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAssociatedPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAssociatedPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAssociatedPhoneNumberDeleteResponse struct {
	Data PortingAssociatedPhoneNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAssociatedPhoneNumberDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAssociatedPhoneNumberDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAssociatedPhoneNumberNewParams struct {
	// Specifies the action to take with this phone number during partial porting.
	//
	// Any of "keep", "disconnect".
	Action           PortingOrderAssociatedPhoneNumberNewParamsAction           `json:"action,omitzero,required"`
	PhoneNumberRange PortingOrderAssociatedPhoneNumberNewParamsPhoneNumberRange `json:"phone_number_range,omitzero,required"`
	paramObj
}

func (r PortingOrderAssociatedPhoneNumberNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderAssociatedPhoneNumberNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderAssociatedPhoneNumberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the action to take with this phone number during partial porting.
type PortingOrderAssociatedPhoneNumberNewParamsAction string

const (
	PortingOrderAssociatedPhoneNumberNewParamsActionKeep       PortingOrderAssociatedPhoneNumberNewParamsAction = "keep"
	PortingOrderAssociatedPhoneNumberNewParamsActionDisconnect PortingOrderAssociatedPhoneNumberNewParamsAction = "disconnect"
)

type PortingOrderAssociatedPhoneNumberNewParamsPhoneNumberRange struct {
	// Specifies the end of the phone number range for this associated phone number.
	EndAt param.Opt[string] `json:"end_at,omitzero"`
	// Specifies the start of the phone number range for this associated phone number.
	StartAt param.Opt[string] `json:"start_at,omitzero"`
	paramObj
}

func (r PortingOrderAssociatedPhoneNumberNewParamsPhoneNumberRange) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderAssociatedPhoneNumberNewParamsPhoneNumberRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderAssociatedPhoneNumberNewParamsPhoneNumberRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAssociatedPhoneNumberListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[phone_number], filter[action]
	Filter PortingOrderAssociatedPhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderAssociatedPhoneNumberListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderAssociatedPhoneNumberListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAssociatedPhoneNumberListParams]'s query
// parameters as `url.Values`.
func (r PortingOrderAssociatedPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[phone_number], filter[action]
type PortingOrderAssociatedPhoneNumberListParamsFilter struct {
	// Filter results by a phone number. It should be in E.164 format.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter results by action type
	//
	// Any of "keep", "disconnect".
	Action string `query:"action,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAssociatedPhoneNumberListParamsFilter]'s query
// parameters as `url.Values`.
func (r PortingOrderAssociatedPhoneNumberListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PortingOrderAssociatedPhoneNumberListParamsFilter](
		"action", "keep", "disconnect",
	)
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderAssociatedPhoneNumberListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAssociatedPhoneNumberListParamsPage]'s query
// parameters as `url.Values`.
func (r PortingOrderAssociatedPhoneNumberListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderAssociatedPhoneNumberListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order
	//
	// Any of "-created_at", "created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAssociatedPhoneNumberListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderAssociatedPhoneNumberListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PortingOrderAssociatedPhoneNumberListParamsSort](
		"value", "-created_at", "created_at",
	)
}

type PortingOrderAssociatedPhoneNumberDeleteParams struct {
	PortingOrderID string `path:"porting_order_id,required" format:"uuid" json:"-"`
	paramObj
}
