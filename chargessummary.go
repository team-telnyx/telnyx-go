// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// ChargesSummaryService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChargesSummaryService] method instead.
type ChargesSummaryService struct {
	Options []option.RequestOption
}

// NewChargesSummaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChargesSummaryService(opts ...option.RequestOption) (r ChargesSummaryService) {
	r = ChargesSummaryService{}
	r.Options = opts
	return
}

// Retrieve a summary of monthly charges for a specified date range. The date range
// cannot exceed 31 days.
func (r *ChargesSummaryService) Get(ctx context.Context, query ChargesSummaryGetParams, opts ...option.RequestOption) (res *ChargesSummaryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "charges_summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MonthDetail struct {
	// Monthly recurring charge amount as decimal string
	Mrc string `json:"mrc,required"`
	// Number of items
	Quantity int64 `json:"quantity,required"`
	// One-time charge amount as decimal string
	Otc string `json:"otc,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mrc         respjson.Field
		Quantity    respjson.Field
		Otc         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonthDetail) RawJSON() string { return r.JSON.raw }
func (r *MonthDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponse struct {
	Data ChargesSummaryGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponseData struct {
	// Currency code
	Currency string `json:"currency,required"`
	// End date of the summary period
	EndDate time.Time `json:"end_date,required" format:"date"`
	// Start date of the summary period
	StartDate time.Time                            `json:"start_date,required" format:"date"`
	Summary   ChargesSummaryGetResponseDataSummary `json:"summary,required"`
	Total     ChargesSummaryGetResponseDataTotal   `json:"total,required"`
	// User email address
	UserEmail string `json:"user_email,required" format:"email"`
	// User identifier
	UserID string `json:"user_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		EndDate     respjson.Field
		StartDate   respjson.Field
		Summary     respjson.Field
		Total       respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponseDataSummary struct {
	// List of billing adjustments
	Adjustments []ChargesSummaryGetResponseDataSummaryAdjustment `json:"adjustments,required"`
	// List of charge summary lines
	Lines []ChargesSummaryGetResponseDataSummaryLineUnion `json:"lines,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adjustments respjson.Field
		Lines       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponseDataSummary) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponseDataSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponseDataSummaryAdjustment struct {
	// Adjustment amount as decimal string
	Amount string `json:"amount,required"`
	// Description of the adjustment
	Description string `json:"description,required"`
	// Date when the adjustment occurred
	EventDate time.Time `json:"event_date,required" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Description respjson.Field
		EventDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponseDataSummaryAdjustment) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponseDataSummaryAdjustment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChargesSummaryGetResponseDataSummaryLineUnion contains all possible properties
// and values from [ChargesSummaryGetResponseDataSummaryLineComparative],
// [ChargesSummaryGetResponseDataSummaryLineSimple].
//
// Use the [ChargesSummaryGetResponseDataSummaryLineUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChargesSummaryGetResponseDataSummaryLineUnion struct {
	Alias string `json:"alias"`
	// This field is from variant
	// [ChargesSummaryGetResponseDataSummaryLineComparative].
	ExistingThisMonth MonthDetail `json:"existing_this_month"`
	Name              string      `json:"name"`
	// This field is from variant
	// [ChargesSummaryGetResponseDataSummaryLineComparative].
	NewThisMonth MonthDetail `json:"new_this_month"`
	// Any of "comparative", "simple".
	Type string `json:"type"`
	// This field is from variant [ChargesSummaryGetResponseDataSummaryLineSimple].
	Amount string `json:"amount"`
	// This field is from variant [ChargesSummaryGetResponseDataSummaryLineSimple].
	Quantity int64 `json:"quantity"`
	JSON     struct {
		Alias             respjson.Field
		ExistingThisMonth respjson.Field
		Name              respjson.Field
		NewThisMonth      respjson.Field
		Type              respjson.Field
		Amount            respjson.Field
		Quantity          respjson.Field
		raw               string
	} `json:"-"`
}

// anyChargesSummaryGetResponseDataSummaryLine is implemented by each variant of
// [ChargesSummaryGetResponseDataSummaryLineUnion] to add type safety for the
// return type of [ChargesSummaryGetResponseDataSummaryLineUnion.AsAny]
type anyChargesSummaryGetResponseDataSummaryLine interface {
	implChargesSummaryGetResponseDataSummaryLineUnion()
}

func (ChargesSummaryGetResponseDataSummaryLineComparative) implChargesSummaryGetResponseDataSummaryLineUnion() {
}
func (ChargesSummaryGetResponseDataSummaryLineSimple) implChargesSummaryGetResponseDataSummaryLineUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChargesSummaryGetResponseDataSummaryLineUnion.AsAny().(type) {
//	case telnyx.ChargesSummaryGetResponseDataSummaryLineComparative:
//	case telnyx.ChargesSummaryGetResponseDataSummaryLineSimple:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChargesSummaryGetResponseDataSummaryLineUnion) AsAny() anyChargesSummaryGetResponseDataSummaryLine {
	switch u.Type {
	case "comparative":
		return u.AsComparative()
	case "simple":
		return u.AsSimple()
	}
	return nil
}

func (u ChargesSummaryGetResponseDataSummaryLineUnion) AsComparative() (v ChargesSummaryGetResponseDataSummaryLineComparative) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChargesSummaryGetResponseDataSummaryLineUnion) AsSimple() (v ChargesSummaryGetResponseDataSummaryLineSimple) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChargesSummaryGetResponseDataSummaryLineUnion) RawJSON() string { return u.JSON.raw }

func (r *ChargesSummaryGetResponseDataSummaryLineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponseDataSummaryLineComparative struct {
	// Service alias
	Alias             string      `json:"alias,required"`
	ExistingThisMonth MonthDetail `json:"existing_this_month,required"`
	// Service name
	Name         string               `json:"name,required"`
	NewThisMonth MonthDetail          `json:"new_this_month,required"`
	Type         constant.Comparative `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias             respjson.Field
		ExistingThisMonth respjson.Field
		Name              respjson.Field
		NewThisMonth      respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponseDataSummaryLineComparative) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponseDataSummaryLineComparative) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponseDataSummaryLineSimple struct {
	// Service alias
	Alias string `json:"alias,required"`
	// Total amount as decimal string
	Amount string `json:"amount,required"`
	// Service name
	Name string `json:"name,required"`
	// Number of items
	Quantity int64           `json:"quantity,required"`
	Type     constant.Simple `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias       respjson.Field
		Amount      respjson.Field
		Name        respjson.Field
		Quantity    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponseDataSummaryLineSimple) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponseDataSummaryLineSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetResponseDataTotal struct {
	// Total credits as decimal string
	Credits string `json:"credits,required"`
	// Total existing monthly recurring charges as decimal string
	ExistingMrc string `json:"existing_mrc,required"`
	// Grand total of all charges as decimal string
	GrandTotal string `json:"grand_total,required"`
	// Ledger adjustments as decimal string
	LedgerAdjustments string `json:"ledger_adjustments,required"`
	// Total new monthly recurring charges as decimal string
	NewMrc string `json:"new_mrc,required"`
	// Total new one-time charges as decimal string
	NewOtc string `json:"new_otc,required"`
	// Other charges as decimal string
	Other string `json:"other,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits           respjson.Field
		ExistingMrc       respjson.Field
		GrandTotal        respjson.Field
		LedgerAdjustments respjson.Field
		NewMrc            respjson.Field
		NewOtc            respjson.Field
		Other             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChargesSummaryGetResponseDataTotal) RawJSON() string { return r.JSON.raw }
func (r *ChargesSummaryGetResponseDataTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChargesSummaryGetParams struct {
	// End date for the charges summary in ISO date format (YYYY-MM-DD). The date is
	// exclusive, data for the end_date itself is not included in the report. The
	// interval between start_date and end_date cannot exceed 31 days.
	EndDate time.Time `query:"end_date,required" format:"date" json:"-"`
	// Start date for the charges summary in ISO date format (YYYY-MM-DD)
	StartDate time.Time `query:"start_date,required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ChargesSummaryGetParams]'s query parameters as
// `url.Values`.
func (r ChargesSummaryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
