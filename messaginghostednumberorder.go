// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
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
	opts = slices.Concat(r.Options, opts)
	path := "messaging_hosted_number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a messaging hosted number order
func (r *MessagingHostedNumberOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List messaging hosted number orders
func (r *MessagingHostedNumberOrderService) List(ctx context.Context, query MessagingHostedNumberOrderListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.MessagingHostedNumberOrder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "messaging_hosted_number_orders"
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

// List messaging hosted number orders
func (r *MessagingHostedNumberOrderService) ListAutoPaging(ctx context.Context, query MessagingHostedNumberOrderListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.MessagingHostedNumberOrder] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a messaging hosted number order and all associated phone numbers.
func (r *MessagingHostedNumberOrderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberOrderDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check hosted messaging eligibility
func (r *MessagingHostedNumberOrderService) CheckEligibility(ctx context.Context, body MessagingHostedNumberOrderCheckEligibilityParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderCheckEligibilityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messaging_hosted_number_orders/eligibility_numbers_check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create verification codes to validate numbers of the hosted order. The
// verification codes will be sent to the numbers of the hosted order.
func (r *MessagingHostedNumberOrderService) NewVerificationCodes(ctx context.Context, id string, body MessagingHostedNumberOrderNewVerificationCodesParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderNewVerificationCodesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	Data []MessagingHostedNumberOrderNewVerificationCodesResponseData `json:"data,required"`
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

// Verification code result response
type MessagingHostedNumberOrderNewVerificationCodesResponseData struct {
	// Phone number for which the verification code was created
	PhoneNumber string `json:"phone_number,required" format:"+E.164"`
	// Error message describing why the verification code creation failed
	Error string `json:"error"`
	// Type of verification method used
	//
	// Any of "sms", "call", "flashcall".
	Type string `json:"type"`
	// Unique identifier for the verification code
	VerificationCodeID string `json:"verification_code_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber        respjson.Field
		Error              respjson.Field
		Type               respjson.Field
		VerificationCodeID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderNewVerificationCodesResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *MessagingHostedNumberOrderNewVerificationCodesResponseData) UnmarshalJSON(data []byte) error {
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
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
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
