// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
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

// Number lookup usage reports
//
// LegacyReportingUsageReportNumberLookupService contains methods and other
// services that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingUsageReportNumberLookupService] method instead.
type LegacyReportingUsageReportNumberLookupService struct {
	Options []option.RequestOption
}

// NewLegacyReportingUsageReportNumberLookupService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLegacyReportingUsageReportNumberLookupService(opts ...option.RequestOption) (r LegacyReportingUsageReportNumberLookupService) {
	r = LegacyReportingUsageReportNumberLookupService{}
	r.Options = opts
	return
}

// Submit a new telco data usage report
func (r *LegacyReportingUsageReportNumberLookupService) New(ctx context.Context, body LegacyReportingUsageReportNumberLookupNewParams, opts ...option.RequestOption) (res *LegacyReportingUsageReportNumberLookupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/usage_reports/number_lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific telco data usage report by its ID
func (r *LegacyReportingUsageReportNumberLookupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingUsageReportNumberLookupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/number_lookup/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of telco data usage reports
func (r *LegacyReportingUsageReportNumberLookupService) List(ctx context.Context, opts ...option.RequestOption) (res *LegacyReportingUsageReportNumberLookupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/usage_reports/number_lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific telco data usage report by its ID
func (r *LegacyReportingUsageReportNumberLookupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/number_lookup/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type TelcoDataAggregation struct {
	// Currency code
	Currency string `json:"currency"`
	// Total cost for this aggregation
	TotalCost float64 `json:"total_cost"`
	// Total number of lookups performed
	TotalDips int64 `json:"total_dips"`
	// Type of telco data lookup
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		TotalCost   respjson.Field
		TotalDips   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelcoDataAggregation) RawJSON() string { return r.JSON.raw }
func (r *TelcoDataAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelcoDataUsageRecord struct {
	// List of aggregations by lookup type
	Aggregations []TelcoDataAggregation `json:"aggregations"`
	// Record type identifier
	RecordType string `json:"record_type"`
	// User ID
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregations respjson.Field
		RecordType   respjson.Field
		UserID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelcoDataUsageRecord) RawJSON() string { return r.JSON.raw }
func (r *TelcoDataUsageRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Telco data usage report response
type TelcoDataUsageReportResponse struct {
	// Unique identifier for the report
	ID string `json:"id" format:"uuid"`
	// Type of aggregation used in the report
	AggregationType string `json:"aggregation_type"`
	// Timestamp when the report was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// End date of the report period
	EndDate time.Time `json:"end_date" format:"date"`
	// List of managed account IDs included in the report
	ManagedAccounts []string `json:"managed_accounts"`
	// Record type identifier
	RecordType string `json:"record_type"`
	// URL to download the complete report
	ReportURL string `json:"report_url"`
	// Array of usage records
	Result []TelcoDataUsageRecord `json:"result"`
	// Start date of the report period
	StartDate time.Time `json:"start_date" format:"date"`
	// Current status of the report
	Status string `json:"status"`
	// Timestamp when the report was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AggregationType respjson.Field
		CreatedAt       respjson.Field
		EndDate         respjson.Field
		ManagedAccounts respjson.Field
		RecordType      respjson.Field
		ReportURL       respjson.Field
		Result          respjson.Field
		StartDate       respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelcoDataUsageReportResponse) RawJSON() string { return r.JSON.raw }
func (r *TelcoDataUsageReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportNumberLookupNewResponse struct {
	// Telco data usage report response
	Data TelcoDataUsageReportResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportNumberLookupNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportNumberLookupNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportNumberLookupGetResponse struct {
	// Telco data usage report response
	Data TelcoDataUsageReportResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportNumberLookupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportNumberLookupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportNumberLookupListResponse struct {
	Data []TelcoDataUsageReportResponse `json:"data"`
	Meta StandardPaginationMeta         `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportNumberLookupListResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportNumberLookupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportNumberLookupNewParams struct {
	// End date for the usage report
	EndDate param.Opt[time.Time] `json:"endDate,omitzero" format:"date"`
	// Start date for the usage report
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date"`
	// Type of aggregation for the report
	//
	// Any of "ALL", "BY_ORGANIZATION_MEMBER".
	AggregationType LegacyReportingUsageReportNumberLookupNewParamsAggregationType `json:"aggregationType,omitzero"`
	// List of managed accounts to include in the report
	ManagedAccounts []string `json:"managedAccounts,omitzero"`
	paramObj
}

func (r LegacyReportingUsageReportNumberLookupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingUsageReportNumberLookupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingUsageReportNumberLookupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of aggregation for the report
type LegacyReportingUsageReportNumberLookupNewParamsAggregationType string

const (
	LegacyReportingUsageReportNumberLookupNewParamsAggregationTypeAll                  LegacyReportingUsageReportNumberLookupNewParamsAggregationType = "ALL"
	LegacyReportingUsageReportNumberLookupNewParamsAggregationTypeByOrganizationMember LegacyReportingUsageReportNumberLookupNewParamsAggregationType = "BY_ORGANIZATION_MEMBER"
)
