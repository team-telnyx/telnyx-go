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

// SimCardOrderPreviewService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardOrderPreviewService] method instead.
type SimCardOrderPreviewService struct {
	Options []option.RequestOption
}

// NewSimCardOrderPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimCardOrderPreviewService(opts ...option.RequestOption) (r SimCardOrderPreviewService) {
	r = SimCardOrderPreviewService{}
	r.Options = opts
	return
}

// Preview SIM card order purchases.
func (r *SimCardOrderPreviewService) Preview(ctx context.Context, body SimCardOrderPreviewPreviewParams, opts ...option.RequestOption) (res *SimCardOrderPreviewPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_card_order_preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimCardOrderPreviewPreviewResponse struct {
	Data SimCardOrderPreviewPreviewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderPreviewPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderPreviewPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderPreviewPreviewResponseData struct {
	// The amount of SIM cards requested in the SIM card order.
	Quantity int64 `json:"quantity"`
	// Identifies the type of the resource.
	RecordType   string                                             `json:"record_type"`
	ShippingCost SimCardOrderPreviewPreviewResponseDataShippingCost `json:"shipping_cost"`
	SimCardsCost SimCardOrderPreviewPreviewResponseDataSimCardsCost `json:"sim_cards_cost"`
	TotalCost    SimCardOrderPreviewPreviewResponseDataTotalCost    `json:"total_cost"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity     respjson.Field
		RecordType   respjson.Field
		ShippingCost respjson.Field
		SimCardsCost respjson.Field
		TotalCost    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderPreviewPreviewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderPreviewPreviewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderPreviewPreviewResponseDataShippingCost struct {
	// A string representing the cost amount.
	Amount string `json:"amount"`
	// ISO 4217 currency string.
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderPreviewPreviewResponseDataShippingCost) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderPreviewPreviewResponseDataShippingCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderPreviewPreviewResponseDataSimCardsCost struct {
	// A string representing the cost amount.
	Amount string `json:"amount"`
	// ISO 4217 currency string.
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderPreviewPreviewResponseDataSimCardsCost) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderPreviewPreviewResponseDataSimCardsCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderPreviewPreviewResponseDataTotalCost struct {
	// A string representing the cost amount.
	Amount string `json:"amount"`
	// ISO 4217 currency string.
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderPreviewPreviewResponseDataTotalCost) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderPreviewPreviewResponseDataTotalCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderPreviewPreviewParams struct {
	// Uniquely identifies the address for the order.
	AddressID string `json:"address_id,required"`
	// The amount of SIM cards that the user would like to purchase in the SIM card
	// order.
	Quantity int64 `json:"quantity,required"`
	paramObj
}

func (r SimCardOrderPreviewPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardOrderPreviewPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardOrderPreviewPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
