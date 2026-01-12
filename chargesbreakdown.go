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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// ChargesBreakdownService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChargesBreakdownService] method instead.
type ChargesBreakdownService struct {
	Options []option.RequestOption
}

// NewChargesBreakdownService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewChargesBreakdownService(opts ...option.RequestOption) (r ChargesBreakdownService) {
	r = ChargesBreakdownService{}
	r.Options = opts
	return
}

// Retrieve a detailed breakdown of monthly charges for phone numbers in a
// specified date range. The date range cannot exceed 31 days.
func (r *ChargesBreakdownService) Get(ctx context.Context, query ChargesBreakdownGetParams, opts ...option.RequestOption) (res *ChargesBreakdownGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "charges_breakdown"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ChargesBreakdownGetResponse struct {
	Data ChargesBreakdownGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesBreakdownGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChargesBreakdownGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesBreakdownGetResponseData struct {
	// Currency code
	Currency string `json:"currency,required"`
	// End date of the breakdown period
	EndDate time.Time `json:"end_date,required" format:"date"`
	// List of phone number charge breakdowns
	Results []ChargesBreakdownGetResponseDataResult `json:"results,required"`
	// Start date of the breakdown period
	StartDate time.Time `json:"start_date,required" format:"date"`
	// User email address
	UserEmail string `json:"user_email,required" format:"email"`
	// User identifier
	UserID string `json:"user_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		EndDate     respjson.Field
		Results     respjson.Field
		StartDate   respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesBreakdownGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ChargesBreakdownGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesBreakdownGetResponseDataResult struct {
	// Type of charge for the number
	ChargeType string `json:"charge_type,required"`
	// Email address of the service owner
	ServiceOwnerEmail string `json:"service_owner_email,required" format:"email"`
	// User ID of the service owner
	ServiceOwnerUserID string `json:"service_owner_user_id,required"`
	// List of services associated with this number
	Services []ChargesBreakdownGetResponseDataResultService `json:"services,required"`
	// Phone number
	Tn string `json:"tn,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargeType         respjson.Field
		ServiceOwnerEmail  respjson.Field
		ServiceOwnerUserID respjson.Field
		Services           respjson.Field
		Tn                 respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesBreakdownGetResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *ChargesBreakdownGetResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesBreakdownGetResponseDataResultService struct {
	// Cost per unit as decimal string
	Cost string `json:"cost,required"`
	// Type of cost (MRC or OTC)
	CostType string `json:"cost_type,required"`
	// Service name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		CostType    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesBreakdownGetResponseDataResultService) RawJSON() string { return r.JSON.raw }
func (r *ChargesBreakdownGetResponseDataResultService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesBreakdownGetParams struct {
	// Start date for the charges breakdown in ISO date format (YYYY-MM-DD)
	StartDate time.Time `query:"start_date,required" format:"date" json:"-"`
	// End date for the charges breakdown in ISO date format (YYYY-MM-DD). If not
	// provided, defaults to start_date + 1 month. The date is exclusive, data for the
	// end_date itself is not included in the report. The interval between start_date
	// and end_date cannot exceed 31 days.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Response format
	//
	// Any of "json", "csv".
	Format ChargesBreakdownGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChargesBreakdownGetParams]'s query parameters as
// `url.Values`.
func (r ChargesBreakdownGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response format
type ChargesBreakdownGetParamsFormat string

const (
	ChargesBreakdownGetParamsFormatJson ChargesBreakdownGetParamsFormat = "json"
	ChargesBreakdownGetParamsFormatCsv  ChargesBreakdownGetParamsFormat = "csv"
)
