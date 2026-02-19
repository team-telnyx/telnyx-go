// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// PaymentAutoRechargePrefService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentAutoRechargePrefService] method instead.
type PaymentAutoRechargePrefService struct {
	Options []option.RequestOption
}

// NewPaymentAutoRechargePrefService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentAutoRechargePrefService(opts ...option.RequestOption) (r PaymentAutoRechargePrefService) {
	r = PaymentAutoRechargePrefService{}
	r.Options = opts
	return
}

// Update payment auto recharge preferences.
func (r *PaymentAutoRechargePrefService) Update(ctx context.Context, body PaymentAutoRechargePrefUpdateParams, opts ...option.RequestOption) (res *PaymentAutoRechargePrefUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payment/auto_recharge_prefs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns the payment auto recharge preferences.
func (r *PaymentAutoRechargePrefService) List(ctx context.Context, opts ...option.RequestOption) (res *PaymentAutoRechargePrefListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payment/auto_recharge_prefs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaymentAutoRechargePrefUpdateResponse struct {
	Data PaymentAutoRechargePrefUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAutoRechargePrefUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentAutoRechargePrefUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAutoRechargePrefUpdateResponseData struct {
	// The unique identifier for the auto recharge preference.
	ID string `json:"id"`
	// Whether auto recharge is enabled.
	Enabled        bool `json:"enabled"`
	InvoiceEnabled bool `json:"invoice_enabled"`
	// The payment preference for auto recharge.
	//
	// Any of "credit_paypal", "ach".
	Preference string `json:"preference"`
	// The amount to recharge the account, the actual recharge amount will be the
	// amount necessary to reach the threshold amount plus the recharge amount.
	RechargeAmount string `json:"recharge_amount" format:"decimal"`
	// The record type.
	RecordType string `json:"record_type"`
	// The threshold amount at which the account will be recharged.
	ThresholdAmount string `json:"threshold_amount" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Enabled         respjson.Field
		InvoiceEnabled  respjson.Field
		Preference      respjson.Field
		RechargeAmount  respjson.Field
		RecordType      respjson.Field
		ThresholdAmount respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAutoRechargePrefUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaymentAutoRechargePrefUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAutoRechargePrefListResponse struct {
	Data PaymentAutoRechargePrefListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAutoRechargePrefListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentAutoRechargePrefListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAutoRechargePrefListResponseData struct {
	// The unique identifier for the auto recharge preference.
	ID string `json:"id"`
	// Whether auto recharge is enabled.
	Enabled        bool `json:"enabled"`
	InvoiceEnabled bool `json:"invoice_enabled"`
	// The payment preference for auto recharge.
	//
	// Any of "credit_paypal", "ach".
	Preference string `json:"preference"`
	// The amount to recharge the account, the actual recharge amount will be the
	// amount necessary to reach the threshold amount plus the recharge amount.
	RechargeAmount string `json:"recharge_amount" format:"decimal"`
	// The record type.
	RecordType string `json:"record_type"`
	// The threshold amount at which the account will be recharged.
	ThresholdAmount string `json:"threshold_amount" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Enabled         respjson.Field
		InvoiceEnabled  respjson.Field
		Preference      respjson.Field
		RechargeAmount  respjson.Field
		RecordType      respjson.Field
		ThresholdAmount respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAutoRechargePrefListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaymentAutoRechargePrefListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAutoRechargePrefUpdateParams struct {
	// Whether auto recharge is enabled.
	Enabled        param.Opt[bool] `json:"enabled,omitzero"`
	InvoiceEnabled param.Opt[bool] `json:"invoice_enabled,omitzero"`
	// The amount to recharge the account, the actual recharge amount will be the
	// amount necessary to reach the threshold amount plus the recharge amount.
	RechargeAmount param.Opt[string] `json:"recharge_amount,omitzero" format:"decimal"`
	// The threshold amount at which the account will be recharged.
	ThresholdAmount param.Opt[string] `json:"threshold_amount,omitzero" format:"decimal"`
	// The payment preference for auto recharge.
	//
	// Any of "credit_paypal", "ach".
	Preference PaymentAutoRechargePrefUpdateParamsPreference `json:"preference,omitzero"`
	paramObj
}

func (r PaymentAutoRechargePrefUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentAutoRechargePrefUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentAutoRechargePrefUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment preference for auto recharge.
type PaymentAutoRechargePrefUpdateParamsPreference string

const (
	PaymentAutoRechargePrefUpdateParamsPreferenceCreditPaypal PaymentAutoRechargePrefUpdateParamsPreference = "credit_paypal"
	PaymentAutoRechargePrefUpdateParamsPreferenceACH          PaymentAutoRechargePrefUpdateParamsPreference = "ach"
)
