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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
)

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
func (r *LegacyReportingUsageReportNumberLookupService) New(ctx context.Context, body LegacyReportingUsageReportNumberLookupNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "legacy/reporting/usage_reports/number_lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve a specific telco data usage report by its ID
func (r *LegacyReportingUsageReportNumberLookupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/number_lookup/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Retrieve a paginated list of telco data usage reports
func (r *LegacyReportingUsageReportNumberLookupService) List(ctx context.Context, query LegacyReportingUsageReportNumberLookupListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "legacy/reporting/usage_reports/number_lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Delete a specific telco data usage report by its ID
func (r *LegacyReportingUsageReportNumberLookupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/number_lookup/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
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

type LegacyReportingUsageReportNumberLookupListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LegacyReportingUsageReportNumberLookupListParams]'s query
// parameters as `url.Values`.
func (r LegacyReportingUsageReportNumberLookupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
