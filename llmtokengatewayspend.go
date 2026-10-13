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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage and report AI Gateway usage.
//
// LlmTokenGatewaySpendService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmTokenGatewaySpendService] method instead.
type LlmTokenGatewaySpendService struct {
	Options []option.RequestOption
}

// NewLlmTokenGatewaySpendService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLlmTokenGatewaySpendService(opts ...option.RequestOption) (r LlmTokenGatewaySpendService) {
	r = LlmTokenGatewaySpendService{}
	r.Options = opts
	return
}

// Return request-level usage history for the authenticated account. Supply both
// UTC dates or neither; omitting both returns the previous full UTC day. Dates use
// a half-open range [start_date, end_date) of at most 31 days. Cost reflects
// budget reference rates, not invoice charges. Unknown spend is excluded and
// reported separately. The report snapshot supports at most 10,000 matching
// requests, independently of page size. Exceeding this bound returns HTTP 503
// without partial results. Narrow the date range or resource filters before
// retrying. A failed report must not be interpreted as zero usage. Request tags
// can filter either report and metadata:<key> can group the summary.
func (r *LlmTokenGatewaySpendService) GetEvents(ctx context.Context, query LlmTokenGatewaySpendGetEventsParams, opts ...option.RequestOption) (res *LlmTokenGatewaySpendGetEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/spend/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return request counts, token counts and reference spend grouped by one
// attribution dimension. Both UTC dates and group_by are required. Dates use a
// half-open range [start_date, end_date) of at most 31 days. Cost reflects budget
// reference rates, not invoice charges. Unknown spend is excluded and reported
// separately. The report snapshot supports at most 10,000 matching requests,
// independently of page size. Exceeding this bound returns HTTP 503 without
// partial results. Narrow the date range or resource filters before retrying. A
// failed report must not be interpreted as zero usage. Request tags can filter
// either report and metadata:<key> can group the summary.
func (r *LlmTokenGatewaySpendService) GetSummary(ctx context.Context, query LlmTokenGatewaySpendGetSummaryParams, opts ...option.RequestOption) (res *LlmTokenGatewaySpendGetSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/spend/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SpendGroupBy string

const (
	SpendGroupByTokenGroup SpendGroupBy = "token_group"
	SpendGroupByTokenUser  SpendGroupBy = "token_user"
	SpendGroupByTokenKey   SpendGroupBy = "token_key"
	SpendGroupByEndUser    SpendGroupBy = "end_user"
)

type LlmTokenGatewaySpendGetEventsResponse struct {
	Data []LlmTokenGatewaySpendGetEventsResponseData `json:"data" api:"required"`
	Meta LlmTokenGatewaySpendGetEventsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewaySpendGetEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewaySpendGetEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewaySpendGetEventsResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Provider attempts begun for the request (fallbacks included); 0 for a cache hit.
	// Token and cost totals sum every attempt.
	Attempts int64 `json:"attempts" api:"required"`
	// Any of "telnyx", "provider", "none".
	Biller               string `json:"biller" api:"required"`
	CacheHit             bool   `json:"cache_hit" api:"required"`
	ConfigurationVersion int64  `json:"configuration_version" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Cost        float64   `json:"cost" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	EndUserID   string    `json:"end_user_id" api:"required"`
	InputTokens int64     `json:"input_tokens" api:"required"`
	// Customer request tags from the x-ltg-metadata header: up to 5 entries. Keys are
	// case-sensitive; keys starting with ltg*, telnyx*, litellm* or cf* (any case) are
	// reserved. Values are strings (1-256 characters, no control or format
	// characters), integers within +/-9007199254740991, or booleans; floats, null,
	// objects and arrays are rejected. Values that look like credentials are rejected.
	// Tags are attribution only: never forwarded upstream, never part of the response
	// cache key, never used for budgets or routing. An untagged request reports {}.
	Metadata map[string]LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion `json:"metadata" api:"required"`
	// Public alias that served (the last attempt); rate_version, biller and
	// provider_key_id follow it.
	Model         string `json:"model" api:"required"`
	OutputTokens  int64  `json:"output_tokens" api:"required"`
	ProviderKeyID string `json:"provider_key_id" api:"required" format:"uuid"`
	RateVersion   string `json:"rate_version" api:"required"`
	RequestID     string `json:"request_id" api:"required" format:"uuid"`
	// Public alias the caller requested; differs from model only when a group fallback
	// served.
	RequestedModel      string `json:"requested_model" api:"required"`
	ReservationMicroUsd int64  `json:"reservation_micro_usd" api:"required"`
	// Any of "succeeded", "failed", "partial", "unknown".
	Status       string `json:"status" api:"required"`
	TokenGroupID string `json:"token_group_id" api:"required" format:"uuid"`
	TokenKeyID   string `json:"token_key_id" api:"required" format:"uuid"`
	TokenUserID  string `json:"token_user_id" api:"required" format:"uuid"`
	// Any of "known", "unknown", "reconciled".
	UsageStatus string `json:"usage_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Attempts             respjson.Field
		Biller               respjson.Field
		CacheHit             respjson.Field
		ConfigurationVersion respjson.Field
		Cost                 respjson.Field
		CreatedAt            respjson.Field
		EndUserID            respjson.Field
		InputTokens          respjson.Field
		Metadata             respjson.Field
		Model                respjson.Field
		OutputTokens         respjson.Field
		ProviderKeyID        respjson.Field
		RateVersion          respjson.Field
		RequestID            respjson.Field
		RequestedModel       respjson.Field
		ReservationMicroUsd  respjson.Field
		Status               respjson.Field
		TokenGroupID         respjson.Field
		TokenKeyID           respjson.Field
		TokenUserID          respjson.Field
		UsageStatus          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewaySpendGetEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewaySpendGetEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion contains all possible
// properties and values from [string], [int64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *LlmTokenGatewaySpendGetEventsResponseDataMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewaySpendGetEventsResponseMeta struct {
	HasMore    bool   `json:"has_more" api:"required"`
	PageNumber int64  `json:"page_number" api:"required"`
	PageSize   int64  `json:"page_size" api:"required"`
	Snapshot   string `json:"snapshot" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		PageNumber  respjson.Field
		PageSize    respjson.Field
		Snapshot    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewaySpendGetEventsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewaySpendGetEventsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewaySpendGetSummaryResponse struct {
	Data []LlmTokenGatewaySpendGetSummaryResponseData `json:"data" api:"required"`
	Meta LlmTokenGatewaySpendGetSummaryResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewaySpendGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewaySpendGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewaySpendGetSummaryResponseData struct {
	DimensionID  string `json:"dimension_id" api:"required"`
	InputTokens  int64  `json:"input_tokens" api:"required"`
	Name         string `json:"name" api:"required"`
	OutputTokens int64  `json:"output_tokens" api:"required"`
	Requests     int64  `json:"requests" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Spend           float64 `json:"spend" api:"required"`
	UnknownRequests int64   `json:"unknown_requests" api:"required"`
	EndUserID       string  `json:"end_user_id" api:"nullable"`
	// Present only for group_by=metadata:<key>: the tag key. dimension_id is then the
	// tag value in text form (booleans as true/false, integers in decimal), or null
	// for requests without the key; name is null.
	MetadataKey  string `json:"metadata_key"`
	TokenGroupID string `json:"token_group_id" api:"nullable" format:"uuid"`
	TokenKeyID   string `json:"token_key_id" api:"nullable" format:"uuid"`
	TokenUserID  string `json:"token_user_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DimensionID     respjson.Field
		InputTokens     respjson.Field
		Name            respjson.Field
		OutputTokens    respjson.Field
		Requests        respjson.Field
		ReservedSpend   respjson.Field
		Spend           respjson.Field
		UnknownRequests respjson.Field
		EndUserID       respjson.Field
		MetadataKey     respjson.Field
		TokenGroupID    respjson.Field
		TokenKeyID      respjson.Field
		TokenUserID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewaySpendGetSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewaySpendGetSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewaySpendGetSummaryResponseMeta struct {
	EndDate time.Time `json:"end_date" api:"required" format:"date"`
	// Group by one request tag key. Keys starting with ltg*, telnyx*, litellm* or cf*
	// (any case) are rejected.
	GroupBy    SpendGroupBy `json:"group_by" api:"required"`
	HasMore    bool         `json:"has_more" api:"required"`
	PageNumber int64        `json:"page_number" api:"required"`
	PageSize   int64        `json:"page_size" api:"required"`
	Snapshot   string       `json:"snapshot" api:"required"`
	StartDate  time.Time    `json:"start_date" api:"required" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		GroupBy     respjson.Field
		HasMore     respjson.Field
		PageNumber  respjson.Field
		PageSize    respjson.Field
		Snapshot    respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewaySpendGetSummaryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewaySpendGetSummaryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewaySpendGetEventsParams struct {
	EndDate    param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	EndUserID  param.Opt[string]    `query:"end_user_id,omitzero" json:"-"`
	PageNumber param.Opt[int64]     `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64]     `query:"page[size],omitzero" json:"-"`
	// Required after page 1; account/filter-bound 15-minute stable snapshot;
	// expired/mismatch 409.
	PageSnapshot param.Opt[string]    `query:"page[snapshot],omitzero" json:"-"`
	StartDate    param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	TokenGroupID param.Opt[string]    `query:"token_group_id,omitzero" format:"uuid" json:"-"`
	TokenKeyID   param.Opt[string]    `query:"token_key_id,omitzero" format:"uuid" json:"-"`
	TokenUserID  param.Opt[string]    `query:"token_user_id,omitzero" format:"uuid" json:"-"`
	// Request tag filters as metadata[<key>]=<value>, at most 5, combined with AND. A
	// value matches the tag's text form exactly (booleans as true/false, integers in
	// decimal), so metadata[flag]=true matches both boolean true and the string
	// "true". Keys are case-sensitive; reserved-prefix keys are rejected. Unindexed:
	// bounded by the account, the date range and the report timeout.
	Metadata map[string]string `query:"metadata,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LlmTokenGatewaySpendGetEventsParams]'s query parameters as
// `url.Values`.
func (r LlmTokenGatewaySpendGetEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LlmTokenGatewaySpendGetSummaryParams struct {
	EndDate time.Time `query:"end_date" api:"required" format:"date" json:"-"`
	// A scope, or metadata:<key> to group by one request tag's text-form value;
	// requests without the key form one group with a null dimension_id.
	GroupBy    SpendGroupBy      `query:"group_by,omitzero" api:"required" json:"-"`
	StartDate  time.Time         `query:"start_date" api:"required" format:"date" json:"-"`
	EndUserID  param.Opt[string] `query:"end_user_id,omitzero" json:"-"`
	PageNumber param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	// Required after page 1; account/filter-bound 15-minute stable snapshot;
	// expired/mismatch 409.
	PageSnapshot param.Opt[string] `query:"page[snapshot],omitzero" json:"-"`
	TokenGroupID param.Opt[string] `query:"token_group_id,omitzero" format:"uuid" json:"-"`
	TokenKeyID   param.Opt[string] `query:"token_key_id,omitzero" format:"uuid" json:"-"`
	TokenUserID  param.Opt[string] `query:"token_user_id,omitzero" format:"uuid" json:"-"`
	// Request tag filters as metadata[<key>]=<value>, at most 5, combined with AND. A
	// value matches the tag's text form exactly (booleans as true/false, integers in
	// decimal), so metadata[flag]=true matches both boolean true and the string
	// "true". Keys are case-sensitive; reserved-prefix keys are rejected. Unindexed:
	// bounded by the account, the date range and the report timeout.
	Metadata map[string]string `query:"metadata,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LlmTokenGatewaySpendGetSummaryParams]'s query parameters as
// `url.Values`.
func (r LlmTokenGatewaySpendGetSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
