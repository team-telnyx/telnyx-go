// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// WirelessService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWirelessService] method instead.
type WirelessService struct {
	Options              []option.RequestOption
	DetailRecordsReports WirelessDetailRecordsReportService
}

// NewWirelessService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWirelessService(opts ...option.RequestOption) (r WirelessService) {
	r = WirelessService{}
	r.Options = opts
	r.DetailRecordsReports = NewWirelessDetailRecordsReportService(opts...)
	return
}

// Retrieve all wireless regions for the given product.
func (r *WirelessService) GetRegions(ctx context.Context, query WirelessGetRegionsParams, opts ...option.RequestOption) (res *WirelessGetRegionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireless/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WirelessGetRegionsResponse struct {
	Data []WirelessGetRegionsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessGetRegionsResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessGetRegionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessGetRegionsResponseData struct {
	// The unique code of the region.
	Code string `json:"code,required"`
	// The name of the region.
	Name string `json:"name,required"`
	// Timestamp when the region was inserted.
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Timestamp when the region was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		InsertedAt  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessGetRegionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WirelessGetRegionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessGetRegionsParams struct {
	// The product for which to list regions (e.g., 'public_ips',
	// 'private_wireless_gateways').
	Product string `query:"product,required" json:"-"`
	paramObj
}

// URLQuery serializes [WirelessGetRegionsParams]'s query parameters as
// `url.Values`.
func (r WirelessGetRegionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
