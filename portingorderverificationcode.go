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
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// PortingOrderVerificationCodeService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderVerificationCodeService] method instead.
type PortingOrderVerificationCodeService struct {
	Options []option.RequestOption
}

// NewPortingOrderVerificationCodeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderVerificationCodeService(opts ...option.RequestOption) (r PortingOrderVerificationCodeService) {
	r = PortingOrderVerificationCodeService{}
	r.Options = opts
	return
}

// Returns a list of verification codes for a porting order.
func (r *PortingOrderVerificationCodeService) List(ctx context.Context, id string, query PortingOrderVerificationCodeListParams, opts ...option.RequestOption) (res *PortingOrderVerificationCodeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/verification_codes", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Send the verification code for all porting phone numbers.
func (r *PortingOrderVerificationCodeService) Send(ctx context.Context, id string, body PortingOrderVerificationCodeSendParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/verification_codes/send", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Verifies the verification code for a list of phone numbers.
func (r *PortingOrderVerificationCodeService) Verify(ctx context.Context, id string, body PortingOrderVerificationCodeVerifyParams, opts ...option.RequestOption) (res *PortingOrderVerificationCodeVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/verification_codes/verify", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PortingOrderVerificationCodeListResponse struct {
	Data []PortingOrderVerificationCodeListResponseData `json:"data"`
	Meta PaginationMeta                                 `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderVerificationCodeListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderVerificationCodeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderVerificationCodeListResponseData struct {
	// Uniquely identifies this porting verification code
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// E164 formatted phone number
	PhoneNumber string `json:"phone_number"`
	// Identifies the associated porting order
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Indicates whether the verification code has been verified
	Verified bool `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		PhoneNumber    respjson.Field
		PortingOrderID respjson.Field
		RecordType     respjson.Field
		UpdatedAt      respjson.Field
		Verified       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderVerificationCodeListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderVerificationCodeListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderVerificationCodeVerifyResponse struct {
	Data []PortingOrderVerificationCodeVerifyResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderVerificationCodeVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderVerificationCodeVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderVerificationCodeVerifyResponseData struct {
	// Uniquely identifies this porting verification code
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// E164 formatted phone number
	PhoneNumber string `json:"phone_number"`
	// Identifies the associated porting order
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Indicates whether the verification code has been verified
	Verified bool `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		PhoneNumber    respjson.Field
		PortingOrderID respjson.Field
		RecordType     respjson.Field
		UpdatedAt      respjson.Field
		Verified       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderVerificationCodeVerifyResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderVerificationCodeVerifyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderVerificationCodeListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[verified]
	Filter PortingOrderVerificationCodeListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderVerificationCodeListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderVerificationCodeListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderVerificationCodeListParams]'s query parameters
// as `url.Values`.
func (r PortingOrderVerificationCodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[verified]
type PortingOrderVerificationCodeListParamsFilter struct {
	// Filter verification codes that have been verified or not
	Verified param.Opt[bool] `query:"verified,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderVerificationCodeListParamsFilter]'s query
// parameters as `url.Values`.
func (r PortingOrderVerificationCodeListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderVerificationCodeListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderVerificationCodeListParamsPage]'s query
// parameters as `url.Values`.
func (r PortingOrderVerificationCodeListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderVerificationCodeListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at", "-created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderVerificationCodeListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderVerificationCodeListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderVerificationCodeSendParams struct {
	PhoneNumbers []string `json:"phone_numbers,omitzero"`
	// Any of "sms", "call".
	VerificationMethod PortingOrderVerificationCodeSendParamsVerificationMethod `json:"verification_method,omitzero"`
	paramObj
}

func (r PortingOrderVerificationCodeSendParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderVerificationCodeSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderVerificationCodeSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderVerificationCodeSendParamsVerificationMethod string

const (
	PortingOrderVerificationCodeSendParamsVerificationMethodSMS  PortingOrderVerificationCodeSendParamsVerificationMethod = "sms"
	PortingOrderVerificationCodeSendParamsVerificationMethodCall PortingOrderVerificationCodeSendParamsVerificationMethod = "call"
)

type PortingOrderVerificationCodeVerifyParams struct {
	VerificationCodes []PortingOrderVerificationCodeVerifyParamsVerificationCode `json:"verification_codes,omitzero"`
	paramObj
}

func (r PortingOrderVerificationCodeVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderVerificationCodeVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderVerificationCodeVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderVerificationCodeVerifyParamsVerificationCode struct {
	Code        param.Opt[string] `json:"code,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	paramObj
}

func (r PortingOrderVerificationCodeVerifyParamsVerificationCode) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderVerificationCodeVerifyParamsVerificationCode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderVerificationCodeVerifyParamsVerificationCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
