// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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

// NumberReservationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberReservationService] method instead.
type NumberReservationService struct {
	Options []option.RequestOption
	Actions NumberReservationActionService
}

// NewNumberReservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumberReservationService(opts ...option.RequestOption) (r NumberReservationService) {
	r = NumberReservationService{}
	r.Options = opts
	r.Actions = NewNumberReservationActionService(opts...)
	return
}

// Creates a Phone Number Reservation for multiple numbers.
func (r *NumberReservationService) New(ctx context.Context, body NumberReservationNewParams, opts ...option.RequestOption) (res *NumberReservationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "number_reservations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a single phone number reservation.
func (r *NumberReservationService) Get(ctx context.Context, numberReservationID string, opts ...option.RequestOption) (res *NumberReservationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if numberReservationID == "" {
		err = errors.New("missing required number_reservation_id parameter")
		return
	}
	path := fmt.Sprintf("number_reservations/%s", numberReservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a paginated list of phone number reservations.
func (r *NumberReservationService) List(ctx context.Context, query NumberReservationListParams, opts ...option.RequestOption) (res *NumberReservationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "number_reservations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NumberReservation struct {
	ID string `json:"id" format:"uuid"`
	// An ISO 8901 datetime string denoting when the numbers reservation was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string for customer look ups.
	CustomerReference string                `json:"customer_reference"`
	PhoneNumbers      []ReservedPhoneNumber `json:"phone_numbers"`
	RecordType        string                `json:"record_type"`
	// The status of the entire reservation.
	//
	// Any of "pending", "success", "failure".
	Status NumberReservationStatus `json:"status"`
	// An ISO 8901 datetime string for when the number reservation was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		CustomerReference respjson.Field
		PhoneNumbers      respjson.Field
		RecordType        respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberReservation) RawJSON() string { return r.JSON.raw }
func (r *NumberReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the entire reservation.
type NumberReservationStatus string

const (
	NumberReservationStatusPending NumberReservationStatus = "pending"
	NumberReservationStatusSuccess NumberReservationStatus = "success"
	NumberReservationStatusFailure NumberReservationStatus = "failure"
)

type ReservedPhoneNumber struct {
	ID string `json:"id" format:"uuid"`
	// An ISO 8901 datetime string denoting when the individual number reservation was
	// created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An ISO 8901 datetime string for when the individual number reservation is going
	// to expire
	ExpiredAt   time.Time `json:"expired_at" format:"date-time"`
	PhoneNumber string    `json:"phone_number"`
	RecordType  string    `json:"record_type"`
	// The status of the phone number's reservation.
	//
	// Any of "pending", "success", "failure".
	Status ReservedPhoneNumberStatus `json:"status"`
	// An ISO 8901 datetime string for when the the individual number reservation was
	// updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiredAt   respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReservedPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ReservedPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ReservedPhoneNumber to a ReservedPhoneNumberParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ReservedPhoneNumberParam.Overrides()
func (r ReservedPhoneNumber) ToParam() ReservedPhoneNumberParam {
	return param.Override[ReservedPhoneNumberParam](json.RawMessage(r.RawJSON()))
}

// The status of the phone number's reservation.
type ReservedPhoneNumberStatus string

const (
	ReservedPhoneNumberStatusPending ReservedPhoneNumberStatus = "pending"
	ReservedPhoneNumberStatusSuccess ReservedPhoneNumberStatus = "success"
	ReservedPhoneNumberStatusFailure ReservedPhoneNumberStatus = "failure"
)

type ReservedPhoneNumberParam struct {
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// An ISO 8901 datetime string denoting when the individual number reservation was
	// created.
	CreatedAt param.Opt[time.Time] `json:"created_at,omitzero" format:"date-time"`
	// An ISO 8901 datetime string for when the individual number reservation is going
	// to expire
	ExpiredAt   param.Opt[time.Time] `json:"expired_at,omitzero" format:"date-time"`
	PhoneNumber param.Opt[string]    `json:"phone_number,omitzero"`
	RecordType  param.Opt[string]    `json:"record_type,omitzero"`
	// An ISO 8901 datetime string for when the the individual number reservation was
	// updated.
	UpdatedAt param.Opt[time.Time] `json:"updated_at,omitzero" format:"date-time"`
	// The status of the phone number's reservation.
	//
	// Any of "pending", "success", "failure".
	Status ReservedPhoneNumberStatus `json:"status,omitzero"`
	paramObj
}

func (r ReservedPhoneNumberParam) MarshalJSON() (data []byte, err error) {
	type shadow ReservedPhoneNumberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReservedPhoneNumberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberReservationNewResponse struct {
	Data NumberReservation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberReservationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberReservationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberReservationGetResponse struct {
	Data NumberReservation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberReservationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberReservationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberReservationListResponse struct {
	Data []NumberReservation `json:"data"`
	Meta PaginationMeta      `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberReservationListResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberReservationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberReservationNewParams struct {
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string]          `json:"customer_reference,omitzero"`
	PhoneNumbers      []ReservedPhoneNumberParam `json:"phone_numbers,omitzero"`
	paramObj
}

func (r NumberReservationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberReservationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberReservationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberReservationListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[status],
	// filter[created_at], filter[phone_numbers.phone_number],
	// filter[customer_reference]
	Filter NumberReservationListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page NumberReservationListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberReservationListParams]'s query parameters as
// `url.Values`.
func (r NumberReservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[status],
// filter[created_at], filter[phone_numbers.phone_number],
// filter[customer_reference]
type NumberReservationListParamsFilter struct {
	// Filter number reservations via the customer reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter number reservations having these phone numbers.
	PhoneNumbersPhoneNumber param.Opt[string] `query:"phone_numbers.phone_number,omitzero" json:"-"`
	// Filter number reservations by status.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter number reservations by date range.
	CreatedAt NumberReservationListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberReservationListParamsFilter]'s query parameters as
// `url.Values`.
func (r NumberReservationListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter number reservations by date range.
type NumberReservationListParamsFilterCreatedAt struct {
	// Filter number reservations later than this value.
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Filter number reservations earlier than this value.
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberReservationListParamsFilterCreatedAt]'s query
// parameters as `url.Values`.
func (r NumberReservationListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type NumberReservationListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberReservationListParamsPage]'s query parameters as
// `url.Values`.
func (r NumberReservationListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
