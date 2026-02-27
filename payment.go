// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Operations for managing stored payment transactions.
//
// PaymentService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
	// V2 Auto Recharge Preferences API
	AutoRechargePrefs PaymentAutoRechargePrefService
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.Options = opts
	r.AutoRechargePrefs = NewPaymentAutoRechargePrefService(opts...)
	return
}

// Create a stored payment transaction
func (r *PaymentService) NewStoredPaymentTransaction(ctx context.Context, body PaymentNewStoredPaymentTransactionParams, opts ...option.RequestOption) (res *PaymentNewStoredPaymentTransactionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/payment/stored_payment_transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PaymentNewStoredPaymentTransactionResponse struct {
	Data PaymentNewStoredPaymentTransactionResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentNewStoredPaymentTransactionResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentNewStoredPaymentTransactionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewStoredPaymentTransactionResponseData struct {
	ID              string    `json:"id"`
	AmountCents     int64     `json:"amount_cents"`
	AmountCurrency  string    `json:"amount_currency"`
	AutoRecharge    bool      `json:"auto_recharge"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	ProcessorStatus string    `json:"processor_status"`
	// Any of "transaction".
	RecordType string `json:"record_type"`
	// Any of "stored_payment".
	TransactionProcessingType string `json:"transaction_processing_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		AmountCents               respjson.Field
		AmountCurrency            respjson.Field
		AutoRecharge              respjson.Field
		CreatedAt                 respjson.Field
		ProcessorStatus           respjson.Field
		RecordType                respjson.Field
		TransactionProcessingType respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentNewStoredPaymentTransactionResponseData) RawJSON() string { return r.JSON.raw }
func (r *PaymentNewStoredPaymentTransactionResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewStoredPaymentTransactionParams struct {
	// Amount in dollars and cents, e.g. "120.00"
	Amount string `json:"amount" api:"required"`
	paramObj
}

func (r PaymentNewStoredPaymentTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewStoredPaymentTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewStoredPaymentTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
