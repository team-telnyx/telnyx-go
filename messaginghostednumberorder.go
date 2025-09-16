// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
	"github.com/stainless-sdks/telnyx-go/shared"
)

// MessagingHostedNumberOrderService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingHostedNumberOrderService] method instead.
type MessagingHostedNumberOrderService struct {
	Options []option.RequestOption
	Actions MessagingHostedNumberOrderActionService
}

// NewMessagingHostedNumberOrderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingHostedNumberOrderService(opts ...option.RequestOption) (r MessagingHostedNumberOrderService) {
	r = MessagingHostedNumberOrderService{}
	r.Options = opts
	r.Actions = NewMessagingHostedNumberOrderActionService(opts...)
	return
}

// Create a messaging hosted number order
func (r *MessagingHostedNumberOrderService) New(ctx context.Context, body MessagingHostedNumberOrderNewParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "messaging_hosted_number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a messaging hosted number order
func (r *MessagingHostedNumberOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberOrderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List messaging hosted number orders
func (r *MessagingHostedNumberOrderService) List(ctx context.Context, query MessagingHostedNumberOrderListParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "messaging_hosted_number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a messaging hosted number order and all associated phone numbers.
func (r *MessagingHostedNumberOrderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberOrderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check eligibility of phone numbers for hosted messaging
func (r *MessagingHostedNumberOrderService) CheckEligibility(ctx context.Context, body MessagingHostedNumberOrderCheckEligibilityParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderCheckEligibilityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "messaging_hosted_number_orders/eligibility_numbers_check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create verification codes to validate numbers of the hosted order. The
// verification codes will be sent to the numbers of the hosted order.
func (r *MessagingHostedNumberOrderService) NewVerificationCodes(ctx context.Context, id string, body MessagingHostedNumberOrderNewVerificationCodesParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderNewVerificationCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s/verification_codes", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validate the verification codes sent to the numbers of the hosted order. The
// verification codes must be created in the verification codes endpoint.
func (r *MessagingHostedNumberOrderService) ValidateCodes(ctx context.Context, id string, body MessagingHostedNumberOrderValidateCodesParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderValidateCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s/validation_codes", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessagingHostedNumberOrderNewResponse struct {
	Data shared.MessagingHostedNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderGetResponse struct {
	Data shared.MessagingHostedNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderListResponse struct {
	Data []shared.MessagingHostedNumberOrder `json:"data"`
	Meta PaginationMeta                      `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderDeleteResponse struct {
	Data shared.MessagingHostedNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderCheckEligibilityResponse struct {
	// List of phone numbers with their eligibility status.
	PhoneNumbers []MessagingHostedNumberOrderCheckEligibilityResponsePhoneNumber `json:"phone_numbers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderCheckEligibilityResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderCheckEligibilityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderCheckEligibilityResponsePhoneNumber struct {
	// Detailed information about the eligibility status.
	Detail string `json:"detail"`
	// Whether the phone number is eligible for hosted messaging.
	Eligible bool `json:"eligible"`
	// The eligibility status of the phone number.
	//
	// Any of "NUMBER_CAN_NOT_BE_REPEATED", "NUMBER_CAN_NOT_BE_VALIDATED",
	// "NUMBER_CAN_NOT_BE_WIRELESS", "NUMBER_CAN_NOT_BE_ACTIVE_IN_YOUR_ACCOUNT",
	// "NUMBER_CAN_NOT_HOSTED_WITH_A_TELNYX_SUBSCRIBER", "NUMBER_CAN_NOT_BE_IN_TELNYX",
	// "NUMBER_IS_NOT_A_US_NUMBER", "NUMBER_IS_NOT_A_VALID_ROUTING_NUMBER",
	// "NUMBER_IS_NOT_IN_E164_FORMAT", "BILLING_ACCOUNT_CHECK_FAILED",
	// "BILLING_ACCOUNT_IS_ABOLISHED", "ELIGIBLE".
	EligibleStatus string `json:"eligible_status"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail         respjson.Field
		Eligible       respjson.Field
		EligibleStatus respjson.Field
		PhoneNumber    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderCheckEligibilityResponsePhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MessagingHostedNumberOrderCheckEligibilityResponsePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderNewVerificationCodesResponse struct {
	Data []MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderNewVerificationCodesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderNewVerificationCodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion contains all
// possible properties and values from
// [MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess],
// [MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeError].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion struct {
	PhoneNumber string `json:"phone_number"`
	// This field is from variant
	// [MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess].
	Type string `json:"type"`
	// This field is from variant
	// [MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess].
	VerificationCodeID string `json:"verification_code_id"`
	// This field is from variant
	// [MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeError].
	Error string `json:"error"`
	JSON  struct {
		PhoneNumber        respjson.Field
		Type               respjson.Field
		VerificationCodeID respjson.Field
		Error              respjson.Field
		raw                string
	} `json:"-"`
}

func (u MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion) AsVerificationCodeSuccess() (v MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion) AsVerificationCodeError() (v MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MessagingHostedNumberOrderNewVerificationCodesResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful verification code creation response
type MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess struct {
	// Phone number for which the verification code was created
	PhoneNumber string `json:"phone_number,required" format:"+E.164"`
	// Type of verification method used
	//
	// Any of "sms", "call", "flashcall".
	Type string `json:"type,required"`
	// Unique identifier for the verification code
	VerificationCodeID string `json:"verification_code_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber        respjson.Field
		Type               respjson.Field
		VerificationCodeID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Failed verification code creation response
type MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeError struct {
	// Error message describing why the verification code creation failed
	Error string `json:"error,required"`
	// Phone number for which the verification code creation failed
	PhoneNumber string `json:"phone_number,required" format:"+E.164"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeError) RawJSON() string {
	return r.JSON.raw
}
func (r *MessagingHostedNumberOrderNewVerificationCodesResponseDataVerificationCodeError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderValidateCodesResponse struct {
	Data MessagingHostedNumberOrderValidateCodesResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderValidateCodesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderValidateCodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderValidateCodesResponseData struct {
	OrderID      string                                                           `json:"order_id,required" format:"uuid"`
	PhoneNumbers []MessagingHostedNumberOrderValidateCodesResponseDataPhoneNumber `json:"phone_numbers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderID      respjson.Field
		PhoneNumbers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderValidateCodesResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderValidateCodesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderValidateCodesResponseDataPhoneNumber struct {
	PhoneNumber string `json:"phone_number,required" format:"+E.164"`
	// Any of "verified", "rejected", "already_verified".
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderValidateCodesResponseDataPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *MessagingHostedNumberOrderValidateCodesResponseDataPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderNewParams struct {
	// Automatically associate the number with this messaging profile ID when the order
	// is complete.
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	// Phone numbers to be used for hosted messaging.
	PhoneNumbers []string `json:"phone_numbers,omitzero" format:"+E.164"`
	paramObj
}

func (r MessagingHostedNumberOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingHostedNumberOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingHostedNumberOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MessagingHostedNumberOrderListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingHostedNumberOrderListParams]'s query parameters as
// `url.Values`.
func (r MessagingHostedNumberOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MessagingHostedNumberOrderListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingHostedNumberOrderListParamsPage]'s query
// parameters as `url.Values`.
func (r MessagingHostedNumberOrderListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingHostedNumberOrderCheckEligibilityParams struct {
	// List of phone numbers to check eligibility
	PhoneNumbers []string `json:"phone_numbers,omitzero,required" format:"+E.164"`
	paramObj
}

func (r MessagingHostedNumberOrderCheckEligibilityParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingHostedNumberOrderCheckEligibilityParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingHostedNumberOrderCheckEligibilityParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderNewVerificationCodesParams struct {
	PhoneNumbers []string `json:"phone_numbers,omitzero,required" format:"+E.164"`
	// Any of "sms", "call", "flashcall".
	VerificationMethod MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethod `json:"verification_method,omitzero,required"`
	paramObj
}

func (r MessagingHostedNumberOrderNewVerificationCodesParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingHostedNumberOrderNewVerificationCodesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingHostedNumberOrderNewVerificationCodesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethod string

const (
	MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethodSMS       MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethod = "sms"
	MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethodCall      MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethod = "call"
	MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethodFlashcall MessagingHostedNumberOrderNewVerificationCodesParamsVerificationMethod = "flashcall"
)

type MessagingHostedNumberOrderValidateCodesParams struct {
	VerificationCodes []MessagingHostedNumberOrderValidateCodesParamsVerificationCode `json:"verification_codes,omitzero,required"`
	paramObj
}

func (r MessagingHostedNumberOrderValidateCodesParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingHostedNumberOrderValidateCodesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingHostedNumberOrderValidateCodesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Code, PhoneNumber are required.
type MessagingHostedNumberOrderValidateCodesParamsVerificationCode struct {
	Code        string `json:"code,required"`
	PhoneNumber string `json:"phone_number,required" format:"+E.164"`
	paramObj
}

func (r MessagingHostedNumberOrderValidateCodesParamsVerificationCode) MarshalJSON() (data []byte, err error) {
	type shadow MessagingHostedNumberOrderValidateCodesParamsVerificationCode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingHostedNumberOrderValidateCodesParamsVerificationCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
