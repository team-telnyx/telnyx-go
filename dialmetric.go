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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Call information
//
// DialMetricService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDialMetricService] method instead.
type DialMetricService struct {
	Options []option.RequestOption
}

// NewDialMetricService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDialMetricService(opts ...option.RequestOption) (r DialMetricService) {
	r = DialMetricService{}
	r.Options = opts
	return
}

// Returns dials-per-second (DPS) metrics for your account's outbound dials made
// with the Call Control `Dial` command and TeXML, as time buckets with dial
// counts, accepted and rejected dials, average and peak DPS, and the DPS limit in
// force.
//
// The result has one record per region per bucket, ordered by `start_time` and
// then `region`, newest first unless `sort=start_time`. Regions are never summed:
// the DPS limit is enforced per region, so a cross-region total cannot be compared
// with it.
//
// Only buckets that saw dials are returned. Treat a missing bucket as zero dials.
//
// Only complete 5-second windows are counted, and a window is counted about 15
// seconds after it closes. The newest `granularity=second` record is therefore
// your live DPS, about 20 seconds old. Request the last minute or two with
// `sort=-start_time`; the first record of each `region` is that region's live rate
// (`page[limit]=1` would hide every region but one). If no recent record exists,
// the live rate is zero.
//
// Results are cursor paginated. Follow `meta.next`, or pass `meta.cursors.after`
// as `page[after]`, until `meta.next` is absent. No total count is returned.
//
// This endpoint is in development and can change before general availability.
func (r *DialMetricService) List(ctx context.Context, query DialMetricListParams, opts ...option.RequestOption) (res *pagination.CloudfsCursorPagination[DialMetricListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "dial_metrics"
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

// Returns dials-per-second (DPS) metrics for your account's outbound dials made
// with the Call Control `Dial` command and TeXML, as time buckets with dial
// counts, accepted and rejected dials, average and peak DPS, and the DPS limit in
// force.
//
// The result has one record per region per bucket, ordered by `start_time` and
// then `region`, newest first unless `sort=start_time`. Regions are never summed:
// the DPS limit is enforced per region, so a cross-region total cannot be compared
// with it.
//
// Only buckets that saw dials are returned. Treat a missing bucket as zero dials.
//
// Only complete 5-second windows are counted, and a window is counted about 15
// seconds after it closes. The newest `granularity=second` record is therefore
// your live DPS, about 20 seconds old. Request the last minute or two with
// `sort=-start_time`; the first record of each `region` is that region's live rate
// (`page[limit]=1` would hide every region but one). If no recent record exists,
// the live rate is zero.
//
// Results are cursor paginated. Follow `meta.next`, or pass `meta.cursors.after`
// as `page[after]`, until `meta.next` is absent. No total count is returned.
//
// This endpoint is in development and can change before general availability.
func (r *DialMetricService) ListAutoPaging(ctx context.Context, query DialMetricListParams, opts ...option.RequestOption) *pagination.CloudfsCursorPaginationAutoPager[DialMetricListResponse] {
	return pagination.NewCloudfsCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Dial counts and rates for one region in one time bucket. One record per region
// per bucket.
type DialMetricListResponse struct {
	// `dials.total` divided by the bucket duration in seconds, rounded to 4 decimal
	// places. Counts every dial attempt, rejected ones included.
	AverageDps float64 `json:"average_dps"`
	// Dial attempt counts in the bucket.
	Dials DialMetricListResponseDials `json:"dials"`
	// DPS limit in force during the busiest 5-second window, so it compares directly
	// with `peak_dps`. `null` means unlimited. Enforcement near the limit is
	// approximate; `dials.rejected.dps_limit` is the authoritative signal that the
	// limit was reached.
	DpsLimit int64 `json:"dps_limit" api:"nullable"`
	// End of the bucket (exclusive), UTC. Five seconds, one minute or one hour after
	// `start_time`, depending on `granularity`.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Dials per second, rounded to 4 decimal places, in the busiest 5-second window of
	// the bucket, summed across all call-processing nodes of the region. Counts every
	// dial attempt, rejected ones included, so `peak_dps` above `dps_limit` is
	// expected when dials were rejected. For `granularity=second` it equals
	// `average_dps`.
	PeakDps float64 `json:"peak_dps"`
	// Any of "dial_metric".
	RecordType DialMetricListResponseRecordType `json:"record_type"`
	// Region where the dials were handled. Current values are `us`, `eu`, `australia`,
	// `middle_east`, `canada` and `india`. New values can be added, so clients must
	// tolerate values not listed here. Buckets of different regions are separate
	// records.
	Region string `json:"region"`
	// Start of the bucket (inclusive), UTC.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageDps  respjson.Field
		Dials       respjson.Field
		DpsLimit    respjson.Field
		EndTime     respjson.Field
		PeakDps     respjson.Field
		RecordType  respjson.Field
		Region      respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialMetricListResponse) RawJSON() string { return r.JSON.raw }
func (r *DialMetricListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dial attempt counts in the bucket.
type DialMetricListResponseDials struct {
	// Dial attempts that were not rejected by a limit.
	Accepted int64 `json:"accepted"`
	// Dial attempts rejected by a limit.
	Rejected DialMetricListResponseDialsRejected `json:"rejected"`
	// All dial attempts, rejected ones included: accepted plus every rejected count.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Rejected    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialMetricListResponseDials) RawJSON() string { return r.JSON.raw }
func (r *DialMetricListResponseDials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dial attempts rejected by a limit.
type DialMetricListResponseDialsRejected struct {
	// Rejected by a daily limit.
	DailyLimit int64 `json:"daily_limit"`
	// Rejected by the dials-per-second limit.
	DpsLimit int64 `json:"dps_limit"`
	// Rejected by an hourly limit.
	HourlyLimit int64 `json:"hourly_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyLimit  respjson.Field
		DpsLimit    respjson.Field
		HourlyLimit respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialMetricListResponseDialsRejected) RawJSON() string { return r.JSON.raw }
func (r *DialMetricListResponseDialsRejected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialMetricListResponseRecordType string

const (
	DialMetricListResponseRecordTypeDialMetric DialMetricListResponseRecordType = "dial_metric"
)

type DialMetricListParams struct {
	// Exclusive end of the range, ISO 8601. Must be after `start_time`. Buckets that
	// start before this time are returned.
	EndTime time.Time `query:"end_time" api:"required" format:"date-time" json:"-"`
	// Bucket size. `second` returns 5-second buckets (the finest resolution) and is
	// kept for the last 48 hours. `minute` and `hour` are kept for the last 30 days.
	//
	// Any of "second", "minute", "hour".
	Granularity DialMetricListParamsGranularity `query:"granularity,omitzero" api:"required" json:"-"`
	// Inclusive start of the range, ISO 8601 (for example `2026-09-29T12:00:00.000Z`).
	// Buckets that start at or after this time are returned. A `start_time` older than
	// the retention window of the chosen `granularity` is moved forward to the start
	// of that window, and `meta.start_time` returns the start actually used.
	StartTime time.Time `query:"start_time" api:"required" format:"date-time" json:"-"`
	// Opaque cursor from `meta.cursors.after` of the previous page. Do not parse or
	// construct it. A cursor that matches no records returns an empty page whose
	// `meta.cursors.after` and `meta.cursors.before` both echo the given cursor.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Maximum number of buckets per page. Values outside the range are clamped to
	// 1-10,000.
	PageLimit param.Opt[int64] `query:"page[limit],omitzero" json:"-"`
	// Sort order of the buckets by `start_time`. `-start_time` (default) returns the
	// newest first, `start_time` the oldest first. `region` follows the same direction
	// for buckets with the same `start_time`. To read the live rate, use
	// `granularity=second` and `sort=-start_time`, then take the first record of each
	// `region`.
	//
	// Any of "-start_time", "start_time".
	Sort DialMetricListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DialMetricListParams]'s query parameters as `url.Values`.
func (r DialMetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bucket size. `second` returns 5-second buckets (the finest resolution) and is
// kept for the last 48 hours. `minute` and `hour` are kept for the last 30 days.
type DialMetricListParamsGranularity string

const (
	DialMetricListParamsGranularitySecond DialMetricListParamsGranularity = "second"
	DialMetricListParamsGranularityMinute DialMetricListParamsGranularity = "minute"
	DialMetricListParamsGranularityHour   DialMetricListParamsGranularity = "hour"
)

// Sort order of the buckets by `start_time`. `-start_time` (default) returns the
// newest first, `start_time` the oldest first. `region` follows the same direction
// for buckets with the same `start_time`. To read the live rate, use
// `granularity=second` and `sort=-start_time`, then take the first record of each
// `region`.
type DialMetricListParamsSort string

const (
	DialMetricListParamsSortMinusStartTime DialMetricListParamsSort = "-start_time"
	DialMetricListParamsSortStartTime      DialMetricListParamsSort = "start_time"
)
