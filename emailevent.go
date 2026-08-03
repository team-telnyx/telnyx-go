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

// Retrieve account-level email events and event statistics.
//
// EmailEventService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailEventService] method instead.
type EmailEventService struct {
	Options []option.RequestOption
}

// NewEmailEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailEventService(opts ...option.RequestOption) (r EmailEventService) {
	r = EmailEventService{}
	r.Options = opts
	return
}

// Lists account-level email events sorted oldest first by
// `occurred_at asc, id asc`.
func (r *EmailEventService) List(ctx context.Context, query EmailEventListParams, opts ...option.RequestOption) (res *EmailEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns counts and rates for email events over a time range. The default start
// time is 30 days ago.
func (r *EmailEventService) GetStats(ctx context.Context, query EmailEventGetStatsParams, opts ...option.RequestOption) (res *EmailEventGetStatsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_events/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type EmailEventType string

const (
	EmailEventTypeQueued             EmailEventType = "queued"
	EmailEventTypeDeferred           EmailEventType = "deferred"
	EmailEventTypeScheduled          EmailEventType = "scheduled"
	EmailEventTypeCancelled          EmailEventType = "cancelled"
	EmailEventTypeSandbox            EmailEventType = "sandbox"
	EmailEventTypeSending            EmailEventType = "sending"
	EmailEventTypeSent               EmailEventType = "sent"
	EmailEventTypeFailed             EmailEventType = "failed"
	EmailEventTypeDelivered          EmailEventType = "delivered"
	EmailEventTypeBounced            EmailEventType = "bounced"
	EmailEventTypeComplained         EmailEventType = "complained"
	EmailEventTypeRejected           EmailEventType = "rejected"
	EmailEventTypeOpened             EmailEventType = "opened"
	EmailEventTypeClicked            EmailEventType = "clicked"
	EmailEventTypeUnsubscribed       EmailEventType = "unsubscribed"
	EmailEventTypeDailyLimitExceeded EmailEventType = "daily_limit_exceeded"
)

type TimeRange struct {
	From time.Time `json:"from" api:"required" format:"date-time"`
	To   time.Time `json:"to" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeRange) RawJSON() string { return r.JSON.raw }
func (r *TimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventListResponse struct {
	Data []EmailEventListResponseData `json:"data" api:"required"`
	Meta EmailEventListResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventListResponseData struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	EmailID    string    `json:"email_id" api:"required" format:"uuid"`
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// Any of "email_event".
	RecordType string `json:"record_type" api:"required"`
	// Any of "queued", "deferred", "scheduled", "cancelled", "sandbox", "sending",
	// "sent", "failed", "delivered", "bounced", "complained", "rejected", "opened",
	// "clicked", "unsubscribed", "daily_limit_exceeded".
	Type EmailEventType `json:"type" api:"required"`
	// Summary of the associated email message. Present when the email_message preload
	// is available.
	Email   EmailEventListResponseDataEmail `json:"email"`
	Payload map[string]any                  `json:"payload"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EmailID     respjson.Field
		OccurredAt  respjson.Field
		RecordType  respjson.Field
		Type        respjson.Field
		Email       respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of the associated email message. Present when the email_message preload
// is available.
type EmailEventListResponseDataEmail struct {
	Cc      []EmailAddress `json:"cc" api:"required"`
	From    EmailAddress   `json:"from" api:"required"`
	Subject string         `json:"subject" api:"required"`
	To      []EmailAddress `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cc          respjson.Field
		From        respjson.Field
		Subject     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponseDataEmail) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponseDataEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventListResponseMeta struct {
	PageSize  int64     `json:"page_size" api:"required"`
	TimeRange TimeRange `json:"time_range" api:"required"`
	// Cursor for the next page, when more results are available.
	PageCursor string `json:"page_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageSize    respjson.Field
		TimeRange   respjson.Field
		PageCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventGetStatsResponse struct {
	Data EmailEventGetStatsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventGetStatsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailEventGetStatsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventGetStatsResponseData struct {
	// Recipient-level outcome counts for the queried time range. Each to, cc, and bcc
	// recipient counts separately; repeated events of the same type for the same
	// message and recipient count once. Partial MTA injection results count successful
	// recipients as sent and unsuccessful recipients as failed. Only the ten listed
	// event types are counted; other valid event types (scheduled, cancelled, sandbox,
	// sending, rejected) are not included in stats.
	Counts EmailEventGetStatsResponseDataCounts `json:"counts" api:"required"`
	// Recipient-level event rates as percentages, rounded to 2 decimal places.
	Rates EmailEventGetStatsResponseDataRates `json:"rates" api:"required"`
	// Any of "email_event_stats".
	RecordType string    `json:"record_type" api:"required"`
	TimeRange  TimeRange `json:"time_range" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counts      respjson.Field
		Rates       respjson.Field
		RecordType  respjson.Field
		TimeRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventGetStatsResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailEventGetStatsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient-level outcome counts for the queried time range. Each to, cc, and bcc
// recipient counts separately; repeated events of the same type for the same
// message and recipient count once. Partial MTA injection results count successful
// recipients as sent and unsuccessful recipients as failed. Only the ten listed
// event types are counted; other valid event types (scheduled, cancelled, sandbox,
// sending, rejected) are not included in stats.
type EmailEventGetStatsResponseDataCounts struct {
	Bounced      int64 `json:"bounced" api:"required"`
	Clicked      int64 `json:"clicked" api:"required"`
	Complained   int64 `json:"complained" api:"required"`
	Deferred     int64 `json:"deferred" api:"required"`
	Delivered    int64 `json:"delivered" api:"required"`
	Failed       int64 `json:"failed" api:"required"`
	Opened       int64 `json:"opened" api:"required"`
	Queued       int64 `json:"queued" api:"required"`
	Sent         int64 `json:"sent" api:"required"`
	Unsubscribed int64 `json:"unsubscribed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced      respjson.Field
		Clicked      respjson.Field
		Complained   respjson.Field
		Deferred     respjson.Field
		Delivered    respjson.Field
		Failed       respjson.Field
		Opened       respjson.Field
		Queued       respjson.Field
		Sent         respjson.Field
		Unsubscribed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventGetStatsResponseDataCounts) RawJSON() string { return r.JSON.raw }
func (r *EmailEventGetStatsResponseDataCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient-level event rates as percentages, rounded to 2 decimal places.
type EmailEventGetStatsResponseDataRates struct {
	// Bounced recipients / queued recipients as a percentage.
	BounceRate float64 `json:"bounce_rate" api:"required"`
	// Recipients clicked / recipients opened as a percentage.
	ClickRate float64 `json:"click_rate" api:"required"`
	// Recipients with a complaint feedback report / delivered recipients as a
	// percentage.
	ComplaintRate float64 `json:"complaint_rate" api:"required"`
	// Deferred recipients / queued recipients as a percentage.
	DeferredRate float64 `json:"deferred_rate" api:"required"`
	// Delivered recipients / queued recipients as a percentage.
	DeliveryRate float64 `json:"delivery_rate" api:"required"`
	// Recipients opened / recipients delivered as a percentage.
	OpenRate float64 `json:"open_rate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BounceRate    respjson.Field
		ClickRate     respjson.Field
		ComplaintRate respjson.Field
		DeferredRate  respjson.Field
		DeliveryRate  respjson.Field
		OpenRate      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventGetStatsResponseDataRates) RawJSON() string { return r.JSON.raw }
func (r *EmailEventGetStatsResponseDataRates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventListParams struct {
	// Filter events for a specific email message UUID. Invalid UUID values are
	// silently ignored (no filter applied).
	EmailID param.Opt[string] `query:"email_id,omitzero" format:"uuid" json:"-"`
	// Inclusive ISO 8601 start timestamp. Defaults to 30 days ago when omitted.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Opaque URL-safe Base64 cursor returned by a previous list response.
	PageCursor param.Opt[string] `query:"page_cursor,omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100. Invalid values are
	// clamped to the valid range.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Inclusive ISO 8601 end timestamp. When `from` is provided without `to`, defaults
	// to `from + 30 days`.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Comma-separated list of event types to include. Also accepts repeated query
	// parameters (e.g. event_type=delivered&event_type=bounced). Unknown values return
	// no matches.
	EventType EmailEventListParamsEventTypeUnion `query:"event_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailEventListParams]'s query parameters as `url.Values`.
func (r EmailEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmailEventListParamsEventTypeUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *EmailEventListParamsEventTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type EmailEventGetStatsParams struct {
	// Inclusive ISO 8601 start timestamp. Defaults to 30 days ago when omitted.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Inclusive ISO 8601 end timestamp. When `from` is provided without `to`, defaults
	// to `from + 30 days`.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [EmailEventGetStatsParams]'s query parameters as
// `url.Values`.
func (r EmailEventGetStatsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
