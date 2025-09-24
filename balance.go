// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// BalanceService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r BalanceService) {
	r = BalanceService{}
	r.Options = opts
	return
}

// Get user balance details
func (r *BalanceService) Get(ctx context.Context, opts ...option.RequestOption) (res *BalanceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BalanceGetResponse struct {
	Data BalanceGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BalanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceGetResponseData struct {
	// Available amount to spend (balance + credit limit)
	AvailableCredit string `json:"available_credit"`
	// The account's current balance.
	Balance string `json:"balance"`
	// The account's credit limit.
	CreditLimit string `json:"credit_limit"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
	// The account’s pending amount.
	Pending string `json:"pending"`
	// Identifies the type of the resource.
	//
	// Any of "balance".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableCredit respjson.Field
		Balance         respjson.Field
		CreditLimit     respjson.Field
		Currency        respjson.Field
		Pending         respjson.Field
		RecordType      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BalanceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
