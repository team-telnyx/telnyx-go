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
)

// Number portout operations
//
// PortoutService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortoutService] method instead.
type PortoutService struct {
	Options []option.RequestOption
	// Number portout operations
	Events PortoutEventService
	// Number portout operations
	Reports PortoutReportService
	// Number portout operations
	Comments PortoutCommentService
	// Number portout operations
	SupportingDocuments PortoutSupportingDocumentService
}

// NewPortoutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPortoutService(opts ...option.RequestOption) (r PortoutService) {
	r = PortoutService{}
	r.Options = opts
	r.Events = NewPortoutEventService(opts...)
	r.Reports = NewPortoutReportService(opts...)
	r.Comments = NewPortoutCommentService(opts...)
	r.SupportingDocuments = NewPortoutSupportingDocumentService(opts...)
	return
}

// Returns the portout request based on the ID provided
func (r *PortoutService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortoutGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the portout requests according to filters
func (r *PortoutService) List(ctx context.Context, query PortoutListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortoutDetails], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "portouts"
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

// Returns the portout requests according to filters
func (r *PortoutService) ListAutoPaging(ctx context.Context, query PortoutListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortoutDetails] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Given a port-out ID, list rejection codes that are eligible for that port-out
func (r *PortoutService) ListRejectionCodes(ctx context.Context, portoutID string, query PortoutListRejectionCodesParams, opts ...option.RequestOption) (res *PortoutListRejectionCodesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portoutID == "" {
		err = errors.New("missing required portout_id parameter")
		return
	}
	path := fmt.Sprintf("portouts/rejections/%s", portoutID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Authorize or reject portout request
func (r *PortoutService) UpdateStatus(ctx context.Context, status PortoutUpdateStatusParamsStatus, params PortoutUpdateStatusParams, opts ...option.RequestOption) (res *PortoutUpdateStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/%s/%v", params.ID, status)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type PortoutDetails struct {
	ID string `json:"id"`
	// Is true when the number is already ported
	AlreadyPorted bool `json:"already_ported"`
	// Name of person authorizing the porting order
	AuthorizedName string `json:"authorized_name"`
	// Carrier the number will be ported out to
	CarrierName string `json:"carrier_name"`
	// City or municipality of billing address
	City string `json:"city"`
	// ISO 8601 formatted date of when the portout was created
	CreatedAt string `json:"created_at"`
	// The current carrier
	CurrentCarrier string `json:"current_carrier"`
	// Person name or company name requesting the port
	EndUserName string `json:"end_user_name"`
	// ISO 8601 formatted Date/Time of the FOC date
	FocDate string `json:"foc_date"`
	// Indicates whether messaging services should be maintained with Telnyx after the
	// port out completes
	HostMessaging bool `json:"host_messaging"`
	// ISO 8601 formatted date of when the portout was created
	InsertedAt string `json:"inserted_at"`
	// The Local Service Request
	Lsr []string `json:"lsr" format:"uri"`
	// Phone numbers associated with this portout
	PhoneNumbers []string `json:"phone_numbers"`
	// Port order number assigned by the carrier the number will be ported out to
	Pon string `json:"pon"`
	// The reason why the order is being rejected by the user. If the order is
	// authorized, this field can be left null
	Reason string `json:"reason" api:"nullable"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The rejection code for one of the valid rejections to reject a port out order
	RejectionCode int64 `json:"rejection_code"`
	// ISO 8601 formatted Date/Time of the user requested FOC date
	RequestedFocDate string `json:"requested_foc_date"`
	// First line of billing address (street address)
	ServiceAddress string `json:"service_address"`
	// New service provider spid
	Spid string `json:"spid"`
	// State, province, or similar of billing address
	State string `json:"state"`
	// Status of portout request
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	Status PortoutDetailsStatus `json:"status"`
	// A key to reference this port out request when contacting Telnyx customer support
	SupportKey string `json:"support_key"`
	// ISO 8601 formatted date of when the portout was last updated
	UpdatedAt string `json:"updated_at"`
	// Identifies the user (or organization) who requested the port out
	UserID string `json:"user_id" format:"uuid"`
	// Telnyx partner providing network coverage
	Vendor string `json:"vendor" format:"uuid"`
	// Postal Code of billing address
	Zip string `json:"zip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AlreadyPorted    respjson.Field
		AuthorizedName   respjson.Field
		CarrierName      respjson.Field
		City             respjson.Field
		CreatedAt        respjson.Field
		CurrentCarrier   respjson.Field
		EndUserName      respjson.Field
		FocDate          respjson.Field
		HostMessaging    respjson.Field
		InsertedAt       respjson.Field
		Lsr              respjson.Field
		PhoneNumbers     respjson.Field
		Pon              respjson.Field
		Reason           respjson.Field
		RecordType       respjson.Field
		RejectionCode    respjson.Field
		RequestedFocDate respjson.Field
		ServiceAddress   respjson.Field
		Spid             respjson.Field
		State            respjson.Field
		Status           respjson.Field
		SupportKey       respjson.Field
		UpdatedAt        respjson.Field
		UserID           respjson.Field
		Vendor           respjson.Field
		Zip              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutDetails) RawJSON() string { return r.JSON.raw }
func (r *PortoutDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of portout request
type PortoutDetailsStatus string

const (
	PortoutDetailsStatusPending         PortoutDetailsStatus = "pending"
	PortoutDetailsStatusAuthorized      PortoutDetailsStatus = "authorized"
	PortoutDetailsStatusPorted          PortoutDetailsStatus = "ported"
	PortoutDetailsStatusRejected        PortoutDetailsStatus = "rejected"
	PortoutDetailsStatusRejectedPending PortoutDetailsStatus = "rejected-pending"
	PortoutDetailsStatusCanceled        PortoutDetailsStatus = "canceled"
)

type PortoutGetResponse struct {
	Data PortoutDetails `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutListRejectionCodesResponse struct {
	Data []PortoutListRejectionCodesResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutListRejectionCodesResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutListRejectionCodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutListRejectionCodesResponseData struct {
	Code           int64  `json:"code"`
	Description    string `json:"description"`
	ReasonRequired bool   `json:"reason_required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code           respjson.Field
		Description    respjson.Field
		ReasonRequired respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutListRejectionCodesResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutListRejectionCodesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutUpdateStatusResponse struct {
	Data PortoutDetails `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutUpdateStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutUpdateStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[carrier_name], filter[country_code], filter[country_code_in],
	// filter[foc_date], filter[inserted_at], filter[phone_number], filter[pon],
	// filter[ported_out_at], filter[spid], filter[status], filter[status_in],
	// filter[support_key]
	Filter PortoutListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutListParams]'s query parameters as `url.Values`.
func (r PortoutListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[carrier_name], filter[country_code], filter[country_code_in],
// filter[foc_date], filter[inserted_at], filter[phone_number], filter[pon],
// filter[ported_out_at], filter[spid], filter[status], filter[status_in],
// filter[support_key]
type PortoutListParamsFilter struct {
	// Filter by new carrier name.
	CarrierName param.Opt[string] `query:"carrier_name,omitzero" json:"-"`
	// Filter by 2-letter country code
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter by foc_date. Matches all portouts with the same date
	FocDate param.Opt[time.Time] `query:"foc_date,omitzero" format:"date-time" json:"-"`
	// Filter by a phone number on the portout. Matches all portouts with the phone
	// number
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter by Port Order Number (PON).
	Pon param.Opt[string] `query:"pon,omitzero" json:"-"`
	// Filter by new carrier spid.
	Spid param.Opt[string] `query:"spid,omitzero" json:"-"`
	// Filter by the portout's support_key
	SupportKey param.Opt[string] `query:"support_key,omitzero" json:"-"`
	// Filter by a list of 2-letter country codes
	CountryCodeIn []string `query:"country_code_in,omitzero" json:"-"`
	// Filter by inserted_at date range using nested operations
	InsertedAt PortoutListParamsFilterInsertedAt `query:"inserted_at,omitzero" json:"-"`
	// Filter by ported_out_at date range using nested operations
	PortedOutAt PortoutListParamsFilterPortedOutAt `query:"ported_out_at,omitzero" json:"-"`
	// Filter by portout status.
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	Status string `query:"status,omitzero" json:"-"`
	// Filter by a list of portout statuses
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	StatusIn []string `query:"status_in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortoutListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by inserted_at date range using nested operations
type PortoutListParamsFilterInsertedAt struct {
	// Filter by inserted_at date greater than or equal.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Filter by inserted_at date less than or equal.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutListParamsFilterInsertedAt]'s query parameters as
// `url.Values`.
func (r PortoutListParamsFilterInsertedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by ported_out_at date range using nested operations
type PortoutListParamsFilterPortedOutAt struct {
	// Filter by ported_out_at date greater than or equal.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Filter by ported_out_at date less than or equal.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutListParamsFilterPortedOutAt]'s query parameters as
// `url.Values`.
func (r PortoutListParamsFilterPortedOutAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortoutListRejectionCodesParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[code],
	// filter[code][in]
	Filter PortoutListRejectionCodesParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutListRejectionCodesParams]'s query parameters as
// `url.Values`.
func (r PortoutListRejectionCodesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[code],
// filter[code][in]
type PortoutListRejectionCodesParamsFilter struct {
	// Filter rejections of a specific code
	Code PortoutListRejectionCodesParamsFilterCodeUnion `query:"code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutListRejectionCodesParamsFilter]'s query parameters
// as `url.Values`.
func (r PortoutListRejectionCodesParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PortoutListRejectionCodesParamsFilterCodeUnion struct {
	OfInt         param.Opt[int64] `query:",omitzero,inline"`
	OfListOfCodes []int64          `query:",omitzero,inline"`
	paramUnion
}

func (u *PortoutListRejectionCodesParamsFilterCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfListOfCodes) {
		return &u.OfListOfCodes
	}
	return nil
}

type PortoutUpdateStatusParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	// Provide a reason if rejecting the port out request
	Reason string `json:"reason" api:"required"`
	// Indicates whether messaging services should be maintained with Telnyx after the
	// port out completes
	HostMessaging param.Opt[bool] `json:"host_messaging,omitzero"`
	paramObj
}

func (r PortoutUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow PortoutUpdateStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortoutUpdateStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutUpdateStatusParamsStatus string

const (
	PortoutUpdateStatusParamsStatusAuthorized      PortoutUpdateStatusParamsStatus = "authorized"
	PortoutUpdateStatusParamsStatusRejectedPending PortoutUpdateStatusParamsStatus = "rejected-pending"
)
