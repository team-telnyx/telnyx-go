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
)

// PortingOrderPhoneNumberBlockService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderPhoneNumberBlockService] method instead.
type PortingOrderPhoneNumberBlockService struct {
	Options []option.RequestOption
}

// NewPortingOrderPhoneNumberBlockService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderPhoneNumberBlockService(opts ...option.RequestOption) (r PortingOrderPhoneNumberBlockService) {
	r = PortingOrderPhoneNumberBlockService{}
	r.Options = opts
	return
}

// Creates a new phone number block.
func (r *PortingOrderPhoneNumberBlockService) New(ctx context.Context, portingOrderID string, body PortingOrderPhoneNumberBlockNewParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberBlockNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/phone_number_blocks", portingOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of all phone number blocks of a porting order.
func (r *PortingOrderPhoneNumberBlockService) List(ctx context.Context, portingOrderID string, query PortingOrderPhoneNumberBlockListParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberBlockListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/phone_number_blocks", portingOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a phone number block.
func (r *PortingOrderPhoneNumberBlockService) Delete(ctx context.Context, id string, body PortingOrderPhoneNumberBlockDeleteParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberBlockDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PortingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/phone_number_blocks/%s", body.PortingOrderID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PortingPhoneNumberBlock struct {
	// Uniquely identifies this porting phone number block.
	ID string `json:"id" format:"uuid"`
	// Specifies the activation ranges for this porting phone number block. The
	// activation range must be within the phone number range and should not overlap
	// with other activation ranges.
	ActivationRanges []PortingPhoneNumberBlockActivationRange `json:"activation_ranges"`
	// Specifies the country code for this porting phone number block. It is a
	// two-letter ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies the phone number range for this porting phone number block.
	PhoneNumberRange PortingPhoneNumberBlockPhoneNumberRange `json:"phone_number_range"`
	// Specifies the phone number type for this porting phone number block.
	//
	// Any of "landline", "local", "mobile", "national", "shared_cost", "toll_free".
	PhoneNumberType PortingPhoneNumberBlockPhoneNumberType `json:"phone_number_type"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ActivationRanges respjson.Field
		CountryCode      respjson.Field
		CreatedAt        respjson.Field
		PhoneNumberRange respjson.Field
		PhoneNumberType  respjson.Field
		RecordType       respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingPhoneNumberBlock) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingPhoneNumberBlockActivationRange struct {
	// Specifies the end of the activation range. It must be no more than the end of
	// the phone number range.
	EndAt string `json:"end_at"`
	// Specifies the start of the activation range. Must be greater or equal the start
	// of the phone number range.
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
func (r PortingPhoneNumberBlockActivationRange) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberBlockActivationRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the phone number range for this porting phone number block.
type PortingPhoneNumberBlockPhoneNumberRange struct {
	// Specifies the end of the phone number range for this porting phone number block.
	EndAt string `json:"end_at"`
	// Specifies the start of the phone number range for this porting phone number
	// block.
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
func (r PortingPhoneNumberBlockPhoneNumberRange) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberBlockPhoneNumberRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the phone number type for this porting phone number block.
type PortingPhoneNumberBlockPhoneNumberType string

const (
	PortingPhoneNumberBlockPhoneNumberTypeLandline   PortingPhoneNumberBlockPhoneNumberType = "landline"
	PortingPhoneNumberBlockPhoneNumberTypeLocal      PortingPhoneNumberBlockPhoneNumberType = "local"
	PortingPhoneNumberBlockPhoneNumberTypeMobile     PortingPhoneNumberBlockPhoneNumberType = "mobile"
	PortingPhoneNumberBlockPhoneNumberTypeNational   PortingPhoneNumberBlockPhoneNumberType = "national"
	PortingPhoneNumberBlockPhoneNumberTypeSharedCost PortingPhoneNumberBlockPhoneNumberType = "shared_cost"
	PortingPhoneNumberBlockPhoneNumberTypeTollFree   PortingPhoneNumberBlockPhoneNumberType = "toll_free"
)

type PortingOrderPhoneNumberBlockNewResponse struct {
	Data PortingPhoneNumberBlock `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberBlockNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberBlockNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberBlockListResponse struct {
	Data []PortingPhoneNumberBlock `json:"data"`
	Meta PaginationMeta            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberBlockListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberBlockListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberBlockDeleteResponse struct {
	Data PortingPhoneNumberBlock `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberBlockDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberBlockDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberBlockNewParams struct {
	// Specifies the activation ranges for this porting phone number block. The
	// activation range must be within the block range and should not overlap with
	// other activation ranges.
	ActivationRanges []PortingOrderPhoneNumberBlockNewParamsActivationRange `json:"activation_ranges,omitzero,required"`
	PhoneNumberRange PortingOrderPhoneNumberBlockNewParamsPhoneNumberRange  `json:"phone_number_range,omitzero,required"`
	paramObj
}

func (r PortingOrderPhoneNumberBlockNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberBlockNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberBlockNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndAt, StartAt are required.
type PortingOrderPhoneNumberBlockNewParamsActivationRange struct {
	// Specifies the end of the activation range. It must be no more than the end of
	// the extension range.
	EndAt string `json:"end_at,required"`
	// Specifies the start of the activation range. Must be greater or equal the start
	// of the extension range.
	StartAt string `json:"start_at,required"`
	paramObj
}

func (r PortingOrderPhoneNumberBlockNewParamsActivationRange) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberBlockNewParamsActivationRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberBlockNewParamsActivationRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndAt, StartAt are required.
type PortingOrderPhoneNumberBlockNewParamsPhoneNumberRange struct {
	// Specifies the end of the phone number range for this porting phone number block.
	EndAt string `json:"end_at,required"`
	// Specifies the start of the phone number range for this porting phone number
	// block.
	StartAt string `json:"start_at,required"`
	paramObj
}

func (r PortingOrderPhoneNumberBlockNewParamsPhoneNumberRange) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberBlockNewParamsPhoneNumberRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberBlockNewParamsPhoneNumberRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberBlockListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[porting_order_id], filter[support_key], filter[status],
	// filter[phone_number], filter[activation_status], filter[portability_status]
	Filter PortingOrderPhoneNumberBlockListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderPhoneNumberBlockListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderPhoneNumberBlockListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberBlockListParams]'s query parameters
// as `url.Values`.
func (r PortingOrderPhoneNumberBlockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[porting_order_id], filter[support_key], filter[status],
// filter[phone_number], filter[activation_status], filter[portability_status]
type PortingOrderPhoneNumberBlockListParamsFilter struct {
	// Filter results by activation status
	//
	// Any of "New", "Pending", "Conflict", "Cancel Pending", "Failed", "Concurred",
	// "Activate RDY", "Disconnect Pending", "Concurrence Sent", "Old", "Sending",
	// "Active", "Cancelled".
	ActivationStatus string `query:"activation_status,omitzero" json:"-"`
	// Filter results by a list of phone numbers
	PhoneNumber []string `query:"phone_number,omitzero" json:"-"`
	// Filter results by portability status
	//
	// Any of "pending", "confirmed", "provisional".
	PortabilityStatus string `query:"portability_status,omitzero" json:"-"`
	// Filter results by a list of porting order ids
	PortingOrderID []string `query:"porting_order_id,omitzero" format:"uuid" json:"-"`
	// Filter porting orders by status(es). Originally: filter[status],
	// filter[status][in][]
	Status PortingOrderPhoneNumberBlockListParamsFilterStatusUnion `query:"status,omitzero" json:"-"`
	// Filter results by support key(s). Originally: filter[support_key][eq],
	// filter[support_key][in][]
	SupportKey PortingOrderPhoneNumberBlockListParamsFilterSupportKeyUnion `query:"support_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberBlockListParamsFilter]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberBlockListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PortingOrderPhoneNumberBlockListParamsFilterStatusUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfPortingOrderPhoneNumberBlockListsFilterStatusString)
	OfPortingOrderPhoneNumberBlockListsFilterStatusString         param.Opt[string] `query:",omitzero,inline"`
	OfPortingOrderPhoneNumberBlockListsFilterStatusArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *PortingOrderPhoneNumberBlockListParamsFilterStatusUnion) asAny() any {
	if !param.IsOmitted(u.OfPortingOrderPhoneNumberBlockListsFilterStatusString) {
		return &u.OfPortingOrderPhoneNumberBlockListsFilterStatusString
	} else if !param.IsOmitted(u.OfPortingOrderPhoneNumberBlockListsFilterStatusArrayItemArray) {
		return &u.OfPortingOrderPhoneNumberBlockListsFilterStatusArrayItemArray
	}
	return nil
}

// Filter by single status
type PortingOrderPhoneNumberBlockListParamsFilterStatusString string

const (
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringDraft            PortingOrderPhoneNumberBlockListParamsFilterStatusString = "draft"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringInProcess        PortingOrderPhoneNumberBlockListParamsFilterStatusString = "in-process"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringSubmitted        PortingOrderPhoneNumberBlockListParamsFilterStatusString = "submitted"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringException        PortingOrderPhoneNumberBlockListParamsFilterStatusString = "exception"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringFocDateConfirmed PortingOrderPhoneNumberBlockListParamsFilterStatusString = "foc-date-confirmed"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringCancelPending    PortingOrderPhoneNumberBlockListParamsFilterStatusString = "cancel-pending"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringPorted           PortingOrderPhoneNumberBlockListParamsFilterStatusString = "ported"
	PortingOrderPhoneNumberBlockListParamsFilterStatusStringCancelled        PortingOrderPhoneNumberBlockListParamsFilterStatusString = "cancelled"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PortingOrderPhoneNumberBlockListParamsFilterSupportKeyUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *PortingOrderPhoneNumberBlockListParamsFilterSupportKeyUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderPhoneNumberBlockListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberBlockListParamsPage]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberBlockListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderPhoneNumberBlockListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order
	//
	// Any of "-created_at", "created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberBlockListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberBlockListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderPhoneNumberBlockDeleteParams struct {
	PortingOrderID string `path:"porting_order_id,required" format:"uuid" json:"-"`
	paramObj
}
