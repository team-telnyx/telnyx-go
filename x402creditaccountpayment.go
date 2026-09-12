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

// Operations for x402 cryptocurrency payment transactions. Fund your Telnyx
// account using USDC stablecoin payments via the x402 protocol.
//
// X402CreditAccountPaymentService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewX402CreditAccountPaymentService] method instead.
type X402CreditAccountPaymentService struct {
	Options []option.RequestOption
}

// NewX402CreditAccountPaymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewX402CreditAccountPaymentService(opts ...option.RequestOption) (r X402CreditAccountPaymentService) {
	r = X402CreditAccountPaymentService{}
	r.Options = opts
	return
}

// Returns a single x402 payment transaction by ID. The transaction must belong to
// the authenticated user; organization sub-users must have read permission on
// transactions. Returns 404 if the transaction does not exist or belongs to
// another user.
func (r *X402CreditAccountPaymentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *X402CreditAccountPaymentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/x402/credit_account/payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of the authenticated user's x402 payment transactions,
// newest first. Organization sub-users must have read permission on transactions;
// without it the list is empty.
func (r *X402CreditAccountPaymentService) List(ctx context.Context, query X402CreditAccountPaymentListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[X402TransactionRecord], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/x402/credit_account/payments"
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

// Returns a paginated list of the authenticated user's x402 payment transactions,
// newest first. Organization sub-users must have read permission on transactions;
// without it the list is empty.
func (r *X402CreditAccountPaymentService) ListAutoPaging(ctx context.Context, query X402CreditAccountPaymentListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[X402TransactionRecord] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// An x402 payment transaction.
type X402TransactionRecord struct {
	// Unique transaction identifier.
	ID string `json:"id" format:"uuid"`
	// The transaction amount in the specified currency.
	Amount string `json:"amount"`
	// ISO 8601 timestamp when the transaction was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The currency of the transaction amount (e.g. USD).
	Currency string `json:"currency"`
	// The original quote ID associated with this transaction.
	QuoteID string `json:"quote_id"`
	// Any of "x402_transaction".
	RecordType X402TransactionRecordRecordType `json:"record_type"`
	// The settlement status of the transaction. x402 transactions are created after
	// successful on-chain settlement, so the status is `settled`.
	//
	// Any of "settled".
	Status X402TransactionRecordStatus `json:"status"`
	// The on-chain transaction hash, if available.
	TxHash string `json:"tx_hash" api:"nullable"`
	// ISO 8601 timestamp when the transaction was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
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
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402TransactionRecord) RawJSON() string { return r.JSON.raw }
func (r *X402TransactionRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402TransactionRecordRecordType string

const (
	X402TransactionRecordRecordTypeX402Transaction X402TransactionRecordRecordType = "x402_transaction"
)

// The settlement status of the transaction. x402 transactions are created after
// successful on-chain settlement, so the status is `settled`.
type X402TransactionRecordStatus string

const (
	X402TransactionRecordStatusSettled X402TransactionRecordStatus = "settled"
)

type X402CreditAccountPaymentGetResponse struct {
	// An x402 payment transaction.
	Data X402TransactionRecord `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402CreditAccountPaymentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *X402CreditAccountPaymentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type X402CreditAccountPaymentListParams struct {
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [X402CreditAccountPaymentListParams]'s query parameters as
// `url.Values`.
func (r X402CreditAccountPaymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
