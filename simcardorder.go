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

// SimCardOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardOrderService] method instead.
type SimCardOrderService struct {
	Options []option.RequestOption
}

// NewSimCardOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimCardOrderService(opts ...option.RequestOption) (r SimCardOrderService) {
	r = SimCardOrderService{}
	r.Options = opts
	return
}

// Creates a new order for SIM cards.
func (r *SimCardOrderService) New(ctx context.Context, body SimCardOrderNewParams, opts ...option.RequestOption) (res *SimCardOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_card_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single SIM card order by its ID.
func (r *SimCardOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all SIM card orders according to filters.
func (r *SimCardOrderService) List(ctx context.Context, query SimCardOrderListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[SimCardOrder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "sim_card_orders"
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

// Get all SIM card orders according to filters.
func (r *SimCardOrderService) ListAutoPaging(ctx context.Context, query SimCardOrderListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[SimCardOrder] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type SimCardOrder struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// An object representing the total cost of the order.
	Cost SimCardOrderCost `json:"cost"`
	// ISO 8601 formatted date-time indicating when the resource was last created.
	CreatedAt string `json:"created_at"`
	// An object representing the address information from when the order was
	// submitted.
	OrderAddress SimCardOrderOrderAddress `json:"order_address"`
	// The amount of SIM cards requested in the SIM card order.
	Quantity int64 `json:"quantity"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the SIM Card order.<ul> <li><code>pending</code> - the
	// order is waiting to be processed.</li> <li><code>processing</code> - the order
	// is currently being processed.</li> <li><code>ready_to_ship</code> - the order is
	// ready to be shipped to the specified <b>address</b>.</li>
	// <li><code>shipped</code> - the order was shipped and is on its way to be
	// delivered to the specified <b>address</b>.</li> <li><code>delivered</code> - the
	// order was delivered to the specified <b>address</b>.</li>
	// <li><code>canceled</code> - the order was canceled.</li> </ul>
	//
	// Any of "pending", "processing", "ready_to_ship", "shipped", "delivered",
	// "canceled".
	Status SimCardOrderStatus `json:"status"`
	// The URL used to get tracking information about the order.
	TrackingURL string `json:"tracking_url" format:"uri"`
	// ISO 8601 formatted date-time indicating when the resource was last updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Cost         respjson.Field
		CreatedAt    respjson.Field
		OrderAddress respjson.Field
		Quantity     respjson.Field
		RecordType   respjson.Field
		Status       respjson.Field
		TrackingURL  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrder) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object representing the total cost of the order.
type SimCardOrderCost struct {
	// A string representing the cost amount.
	Amount string `json:"amount"`
	// Filter by ISO 4217 currency string.
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
func (r SimCardOrderCost) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object representing the address information from when the order was
// submitted.
type SimCardOrderOrderAddress struct {
	// Uniquely identifies the address for the order.
	ID string `json:"id"`
	// State or province where the address is located.
	AdministrativeArea string `json:"administrative_area"`
	// The name of the business where the address is located.
	BusinessName string `json:"business_name"`
	// The mobile operator two-character (ISO 3166-1 alpha-2) origin country code.
	CountryCode string `json:"country_code"`
	// Supplemental field for address information.
	ExtendedAddress string `json:"extended_address"`
	// The first name of the shipping recipient.
	FirstName string `json:"first_name"`
	// The last name of the shipping recipient.
	LastName string `json:"last_name"`
	// The name of the city where the address is located.
	Locality string `json:"locality"`
	// Postal code for the address.
	PostalCode string `json:"postal_code"`
	// The name of the street where the address is located.
	StreetAddress string `json:"street_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AdministrativeArea respjson.Field
		BusinessName       respjson.Field
		CountryCode        respjson.Field
		ExtendedAddress    respjson.Field
		FirstName          respjson.Field
		LastName           respjson.Field
		Locality           respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderOrderAddress) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderOrderAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the SIM Card order.<ul> <li><code>pending</code> - the
// order is waiting to be processed.</li> <li><code>processing</code> - the order
// is currently being processed.</li> <li><code>ready_to_ship</code> - the order is
// ready to be shipped to the specified <b>address</b>.</li>
// <li><code>shipped</code> - the order was shipped and is on its way to be
// delivered to the specified <b>address</b>.</li> <li><code>delivered</code> - the
// order was delivered to the specified <b>address</b>.</li>
// <li><code>canceled</code> - the order was canceled.</li> </ul>
type SimCardOrderStatus string

const (
	SimCardOrderStatusPending     SimCardOrderStatus = "pending"
	SimCardOrderStatusProcessing  SimCardOrderStatus = "processing"
	SimCardOrderStatusReadyToShip SimCardOrderStatus = "ready_to_ship"
	SimCardOrderStatusShipped     SimCardOrderStatus = "shipped"
	SimCardOrderStatusDelivered   SimCardOrderStatus = "delivered"
	SimCardOrderStatusCanceled    SimCardOrderStatus = "canceled"
)

type SimCardOrderNewResponse struct {
	Data SimCardOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderGetResponse struct {
	Data SimCardOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderNewParams struct {
	// Uniquely identifies the address for the order.
	AddressID string `json:"address_id,required"`
	// The amount of SIM cards to order.
	Quantity int64 `json:"quantity,required"`
	paramObj
}

func (r SimCardOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardOrderListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter for SIM card orders (deepObject style).
	// Originally: filter[created_at], filter[updated_at], filter[quantity],
	// filter[cost.amount], filter[cost.currency], filter[address.id],
	// filter[address.street_address], filter[address.extended_address],
	// filter[address.locality], filter[address.administrative_area],
	// filter[address.country_code], filter[address.postal_code]
	Filter SimCardOrderListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardOrderListParams]'s query parameters as `url.Values`.
func (r SimCardOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for SIM card orders (deepObject style).
// Originally: filter[created_at], filter[updated_at], filter[quantity],
// filter[cost.amount], filter[cost.currency], filter[address.id],
// filter[address.street_address], filter[address.extended_address],
// filter[address.locality], filter[address.administrative_area],
// filter[address.country_code], filter[address.postal_code]
type SimCardOrderListParamsFilter struct {
	// Filter by state or province where the address is located.
	AddressAdministrativeArea param.Opt[string] `query:"address.administrative_area,omitzero" json:"-"`
	// Filter by the mobile operator two-character (ISO 3166-1 alpha-2) origin country
	// code.
	AddressCountryCode param.Opt[string] `query:"address.country_code,omitzero" json:"-"`
	// Returns entries with matching name of the supplemental field for address
	// information.
	AddressExtendedAddress param.Opt[string] `query:"address.extended_address,omitzero" json:"-"`
	// Uniquely identifies the address for the order.
	AddressID param.Opt[string] `query:"address.id,omitzero" json:"-"`
	// Filter by the name of the city where the address is located.
	AddressLocality param.Opt[string] `query:"address.locality,omitzero" json:"-"`
	// Filter by postal code for the address.
	AddressPostalCode param.Opt[string] `query:"address.postal_code,omitzero" json:"-"`
	// Returns entries with matching name of the street where the address is located.
	AddressStreetAddress param.Opt[string] `query:"address.street_address,omitzero" json:"-"`
	// The total monetary amount of the order.
	CostAmount param.Opt[string] `query:"cost.amount,omitzero" json:"-"`
	// Filter by ISO 4217 currency string.
	CostCurrency param.Opt[string] `query:"cost.currency,omitzero" json:"-"`
	// Filter by ISO 8601 formatted date-time string matching resource creation
	// date-time.
	CreatedAt param.Opt[time.Time] `query:"created_at,omitzero" format:"date-time" json:"-"`
	// Filter orders by how many SIM cards were ordered.
	Quantity param.Opt[int64] `query:"quantity,omitzero" json:"-"`
	// Filter by ISO 8601 formatted date-time string matching resource last update
	// date-time.
	UpdatedAt param.Opt[time.Time] `query:"updated_at,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardOrderListParamsFilter]'s query parameters as
// `url.Values`.
func (r SimCardOrderListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
