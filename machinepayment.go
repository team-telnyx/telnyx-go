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

// Machine payment (MPP) account-credit operations. Fund your Telnyx account
// programmatically from a machine or agent using the Machine Payment Protocol, an
// HTTP-402 flow settled via Stripe or Tempo.
//
// MachinePaymentService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMachinePaymentService] method instead.
type MachinePaymentService struct {
	Options []option.RequestOption
}

// NewMachinePaymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMachinePaymentService(opts ...option.RequestOption) (r MachinePaymentService) {
	r = MachinePaymentService{}
	r.Options = opts
	return
}

// Creates an account credit using the Machine Payment Protocol (MPP), an HTTP-402
// payment flow for machines and agents.
//
// The flow has two steps. First, send an authenticated request with the
// `amount_usd` to credit; the response is `402 Payment Required` with one or more
// payment challenges (for example separate Tempo and Stripe challenges) in the
// `WWW-Authenticate` header. Second, retry the request with an
// `Authorization: Payment ...` credential constructed from the challenge; on
// success the response includes the credited transaction and a `Payment-Receipt`
// header.
//
// The credited account is never chosen by the request body: the initial request
// credits the account of the authenticated user, and a paid retry credits the
// account bound to the verified payment credential. The amount must be within the
// configured bounds (by default between 5.00 and 500.00 USD).
//
// Successful paid retries are idempotent — when Rails reaches its
// duplicate-transaction lookup for an already-recorded payment, it returns the
// existing transaction with `created: false` instead of crediting the account
// again. This deduplication applies to successful fulfillment: re-sending the same
// Stripe credential may instead be rejected by the upstream provider as an
// idempotent replay and return `402 Payment Required` rather than the existing
// transaction.
//
// > **Warning: the payment credential is bound to a specific Telnyx account ID.**
// > A payment is captured before the bound account is validated. If the credential
// > names an account that is missing, suspended, blocked, cancelled, dormant, or
// > ineligible for the tier, the payment is captured but **no account is
// > credited**. If the credential names a different but eligible account, that
// > account is credited — the service does not compare it against the payer's
// > account. There is **no automatic refund**: if the captured payment does not
// > credit the intended account, contact Telnyx support for remediation.
func (r *MachinePaymentService) AccountCredit(ctx context.Context, body MachinePaymentAccountCreditParams, opts ...option.RequestOption) (res *MachinePaymentAccountCreditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "machine-payments/account-credit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MachinePaymentAccountCreditResponse struct {
	// An account-credit transaction settled through the Machine Payment Protocol.
	Data MachinePaymentAccountCreditResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachinePaymentAccountCreditResponse) RawJSON() string { return r.JSON.raw }
func (r *MachinePaymentAccountCreditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An account-credit transaction settled through the Machine Payment Protocol.
type MachinePaymentAccountCreditResponseData struct {
	// Unique identifier of the account-credit transaction.
	ID string `json:"id" api:"required"`
	// Identifier of the credited Telnyx account. Derived from the authenticated user
	// on the initial request and from the verified payment credential on a paid retry
	// — never from the request body.
	AccountID string `json:"account_id" api:"required"`
	// Credited amount as a decimal string with two fractional digits.
	Amount string `json:"amount" api:"required"`
	// ISO 4217 currency code of the credited amount (currently always USD).
	Currency string `json:"currency" api:"required"`
	// Payment source identifier distinguishing machine payments from other
	// account-credit sources.
	//
	// Any of "machine_payment".
	PaymentSource string `json:"payment_source" api:"required"`
	// Record type identifier.
	//
	// Any of "machine_payment_account_credit".
	RecordType string `json:"record_type" api:"required"`
	// True when this response created a new account credit, false when an existing
	// transaction was returned for a duplicate paid retry.
	Created bool `json:"created"`
	// ISO 8601 timestamp when the transaction was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Machine Payment Protocol resource identifier the payment credential was bound
	// to.
	MppResource string `json:"mpp_resource" api:"nullable"`
	// Stripe PaymentIntent identifier for Stripe settlements. Absent for Tempo
	// settlements.
	PaymentIntentID string `json:"payment_intent_id" api:"nullable"`
	// Payment method used by the provider: `stripe_spt` for Stripe Shared Payment
	// Token payments, `tempo_usdc` for Tempo USDC payments.
	//
	// Any of "stripe_spt", "tempo_usdc".
	PaymentMethod string `json:"payment_method" api:"nullable"`
	// Upstream payment provider that settled the payment.
	//
	// Any of "stripe", "tempo".
	Provider string `json:"provider" api:"nullable"`
	// Provider receipt reference: the Stripe PaymentIntent identifier for Stripe
	// settlements, or the on-chain transaction hash for Tempo settlements.
	ReceiptReference string `json:"receipt_reference" api:"nullable"`
	// Status of the transaction. Successful machine payment credits are recorded as
	// `settled`.
	//
	// Any of "new", "processing", "settled", "expired", "invalid".
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		Amount           respjson.Field
		Currency         respjson.Field
		PaymentSource    respjson.Field
		RecordType       respjson.Field
		Created          respjson.Field
		CreatedAt        respjson.Field
		MppResource      respjson.Field
		PaymentIntentID  respjson.Field
		PaymentMethod    respjson.Field
		Provider         respjson.Field
		ReceiptReference respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachinePaymentAccountCreditResponseData) RawJSON() string { return r.JSON.raw }
func (r *MachinePaymentAccountCreditResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachinePaymentAccountCreditParams struct {
	// Amount to credit in USD, as a decimal string with up to two fractional digits
	// (by default between 5.00 and 500.00). The request body is required on the
	// initial challenge request and remains required on a paid retry, where you
	// re-send the identical body plus the payment credential — the credential, not the
	// body, selects the payment, and the retried body is not re-validated.
	AmountUsd string `json:"amount_usd" api:"required"`
	paramObj
}

func (r MachinePaymentAccountCreditParams) MarshalJSON() (data []byte, err error) {
	type shadow MachinePaymentAccountCreditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MachinePaymentAccountCreditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
