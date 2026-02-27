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

// Speech to text usage reports
//
// LegacyReportingUsageReportService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingUsageReportService] method instead.
type LegacyReportingUsageReportService struct {
	Options []option.RequestOption
	// Messaging usage reports
	Messaging LegacyReportingUsageReportMessagingService
	// Number lookup usage reports
	NumberLookup LegacyReportingUsageReportNumberLookupService
	// Voice usage reports
	Voice LegacyReportingUsageReportVoiceService
}

// NewLegacyReportingUsageReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLegacyReportingUsageReportService(opts ...option.RequestOption) (r LegacyReportingUsageReportService) {
	r = LegacyReportingUsageReportService{}
	r.Options = opts
	r.Messaging = NewLegacyReportingUsageReportMessagingService(opts...)
	r.NumberLookup = NewLegacyReportingUsageReportNumberLookupService(opts...)
	r.Voice = NewLegacyReportingUsageReportVoiceService(opts...)
	return
}

// Generate and fetch speech to text usage report synchronously. This endpoint will
// both generate and fetch the speech to text report over a specified time period.
func (r *LegacyReportingUsageReportService) GetSpeechToText(ctx context.Context, query LegacyReportingUsageReportGetSpeechToTextParams, opts ...option.RequestOption) (res *LegacyReportingUsageReportGetSpeechToTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/usage_reports/speech_to_text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LegacyReportingUsageReportGetSpeechToTextResponse struct {
	Data map[string]any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportGetSpeechToTextResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportGetSpeechToTextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportGetSpeechToTextParams struct {
	EndDate   param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [LegacyReportingUsageReportGetSpeechToTextParams]'s query
// parameters as `url.Values`.
func (r LegacyReportingUsageReportGetSpeechToTextParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
