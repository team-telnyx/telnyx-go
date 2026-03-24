// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Operations for x402 cryptocurrency payment transactions. Fund your Telnyx
// account using USDC stablecoin payments via the x402 protocol.
//
// X402CreditAccountService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewX402CreditAccountService] method instead.
type X402CreditAccountService struct {
	Options []option.RequestOption
}

// NewX402CreditAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewX402CreditAccountService(opts ...option.RequestOption) (r X402CreditAccountService) {
	r = X402CreditAccountService{}
	r.Options = opts
	return
}

// Creates a payment quote for the specified USD amount. Returns payment details
// including the x402 payment requirements, network, and expiration time. The quote
// must be settled before it expires.
func (r *X402CreditAccountService) NewQuote(ctx context.Context, body X402CreditAccountNewQuoteParams, opts ...option.RequestOption) (res *X402CreditAccountNewQuoteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/x402/credit_account/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Settles an x402 payment using the quote ID and a signed payment authorization.
// The payment signature can be provided via the `PAYMENT-SIGNATURE` header or the
// `payment_signature` body parameter. Settlement is idempotent — submitting the
// same quote ID multiple times returns the existing transaction.
func (r *X402CreditAccountService) Settle(ctx context.Context, params X402CreditAccountSettleParams, opts ...option.RequestOption) (res *X402CreditAccountSettleResponse, err error) {
	if !param.IsOmitted(params.HeaderPaymentSignature) {
		opts = append(opts, option.WithHeader("PAYMENT-SIGNATURE", fmt.Sprintf("%v", params.HeaderPaymentSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v2/x402/credit_account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type X402CreditAccountNewQuoteResponse struct {
	Data X402CreditAccountNewQuoteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountNewQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *X402CreditAccountNewQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountNewQuoteResponseData struct {
	// Unique quote identifier. Use this to settle the payment.
	ID string `json:"id"`
	// The equivalent amount in the payment cryptocurrency's smallest unit (e.g. USDC
	// has 6 decimals, so $50.00 = "50000000").
	AmountCrypto string `json:"amount_crypto"`
	// The quoted amount in USD.
	AmountUsd string `json:"amount_usd"`
	// ISO 8601 timestamp when the quote expires.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The blockchain network for the payment in CAIP-2 format (e.g. eip155:8453 for
	// Base).
	Network string `json:"network"`
	// x402 protocol v2 payment requirements. Contains all information needed to
	// construct and sign a payment authorization.
	PaymentRequirements X402CreditAccountNewQuoteResponseDataPaymentRequirements `json:"payment_requirements"`
	// Any of "quote".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AmountCrypto        respjson.Field
		AmountUsd           respjson.Field
		ExpiresAt           respjson.Field
		Network             respjson.Field
		PaymentRequirements respjson.Field
		RecordType          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountNewQuoteResponseData) RawJSON() string { return r.JSON.raw }
func (r *X402CreditAccountNewQuoteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// x402 protocol v2 payment requirements. Contains all information needed to
// construct and sign a payment authorization.
type X402CreditAccountNewQuoteResponseDataPaymentRequirements struct {
	// Accepted payment schemes. Currently only the `exact` EVM scheme is supported.
	Accepts []X402CreditAccountNewQuoteResponseDataPaymentRequirementsAccept `json:"accepts"`
	// The resource being paid for. Included in the payment signature.
	Resource X402CreditAccountNewQuoteResponseDataPaymentRequirementsResource `json:"resource"`
	// x402 protocol version. Currently always 2.
	//
	// Any of 2.
	X402Version int64 `json:"x402Version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepts     respjson.Field
		Resource    respjson.Field
		X402Version respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountNewQuoteResponseDataPaymentRequirements) RawJSON() string { return r.JSON.raw }
func (r *X402CreditAccountNewQuoteResponseDataPaymentRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountNewQuoteResponseDataPaymentRequirementsAccept struct {
	// Amount in the token's smallest unit.
	Amount string `json:"amount"`
	// Token contract address (e.g. USDC on Base).
	Asset string `json:"asset"`
	// Additional scheme-specific parameters including EIP-712 domain info and the
	// facilitator URL.
	Extra X402CreditAccountNewQuoteResponseDataPaymentRequirementsAcceptExtra `json:"extra"`
	// Maximum time in seconds before the payment authorization expires.
	MaxTimeoutSeconds int64 `json:"maxTimeoutSeconds"`
	// Blockchain network identifier in CAIP-2 format (e.g. "eip155:8453" for Base).
	Network string `json:"network"`
	// Recipient wallet address.
	PayTo string `json:"payTo"`
	// Payment scheme (e.g. "exact").
	Scheme string `json:"scheme"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		Asset             respjson.Field
		Extra             respjson.Field
		MaxTimeoutSeconds respjson.Field
		Network           respjson.Field
		PayTo             respjson.Field
		Scheme            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountNewQuoteResponseDataPaymentRequirementsAccept) RawJSON() string {
	return r.JSON.raw
}
func (r *X402CreditAccountNewQuoteResponseDataPaymentRequirementsAccept) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional scheme-specific parameters including EIP-712 domain info and the
// facilitator URL.
type X402CreditAccountNewQuoteResponseDataPaymentRequirementsAcceptExtra struct {
	FacilitatorURL string `json:"facilitatorUrl"`
	// EIP-712 domain name (e.g. "USD Coin").
	Name    string `json:"name"`
	QuoteID string `json:"quoteId"`
	// EIP-712 domain version.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FacilitatorURL respjson.Field
		Name           respjson.Field
		QuoteID        respjson.Field
		Version        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountNewQuoteResponseDataPaymentRequirementsAcceptExtra) RawJSON() string {
	return r.JSON.raw
}
func (r *X402CreditAccountNewQuoteResponseDataPaymentRequirementsAcceptExtra) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The resource being paid for. Included in the payment signature.
type X402CreditAccountNewQuoteResponseDataPaymentRequirementsResource struct {
	// Human-readable description of the payment.
	Description string `json:"description"`
	// MIME type of the resource.
	MimeType string `json:"mimeType"`
	// Canonical URL of the payment resource.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		MimeType    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountNewQuoteResponseDataPaymentRequirementsResource) RawJSON() string {
	return r.JSON.raw
}
func (r *X402CreditAccountNewQuoteResponseDataPaymentRequirementsResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountSettleResponse struct {
	Data X402CreditAccountSettleResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountSettleResponse) RawJSON() string { return r.JSON.raw }
func (r *X402CreditAccountSettleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountSettleResponseData struct {
	// Unique transaction identifier.
	ID string `json:"id"`
	// The transaction amount in the specified currency.
	Amount string `json:"amount"`
	// ISO 8601 timestamp when the transaction was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The currency of the transaction amount (e.g. USD).
	Currency string `json:"currency"`
	// The original quote ID associated with this transaction.
	QuoteID string `json:"quote_id"`
	// Any of "x402_transaction".
	RecordType string `json:"record_type"`
	// The settlement status of the transaction.
	//
	// Any of "settled".
	Status string `json:"status"`
	// The on-chain transaction hash, if available.
	TxHash string `json:"tx_hash" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		QuoteID     respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		TxHash      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountSettleResponseData) RawJSON() string { return r.JSON.raw }
func (r *X402CreditAccountSettleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountNewQuoteParams struct {
	// Amount in USD to fund (minimum 5.00, maximum 10000.00).
	AmountUsd string `json:"amount_usd" api:"required"`
	paramObj
}

func (r X402CreditAccountNewQuoteParams) MarshalJSON() (data []byte, err error) {
	type shadow X402CreditAccountNewQuoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *X402CreditAccountNewQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountSettleParams struct {
	// The quote ID to settle.
	ID string `json:"id" api:"required"`
	// Base64-encoded signed payment authorization (x402 PaymentPayload). Can
	// alternatively be provided via the PAYMENT-SIGNATURE header.
	PaymentSignature       param.Opt[string] `json:"payment_signature,omitzero"`
	HeaderPaymentSignature param.Opt[string] `header:"PAYMENT-SIGNATURE,omitzero" json:"-"`
	paramObj
}

func (r X402CreditAccountSettleParams) MarshalJSON() (data []byte, err error) {
	type shadow X402CreditAccountSettleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *X402CreditAccountSettleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
