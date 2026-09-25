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
// `occurred_at asc, id asc`. Each row contains a legacy email.-prefixed event_type
// and an additive canonical_event_type. Gateway rejection renders email.failed
// with canonical email.gw_reject; ambiguous injection timeout renders
// email.injection_timeout in both; MTA expiration renders email.bounced with
// canonical email.expired. Message-scoped queued, sending, sandbox, cancelled, and
// daily_limit_exceeded rows fan out per durable recipient with stable derived IDs
// matching webhook delivery. Scheduled is the cardinality exception: account
// polling retains one message-scoped scheduled row with its stored event ID, while
// scheduled webhook publication fans out per recipient with derived IDs; reconcile
// scheduled events by message ID, event type, and occurrence time rather than
// event UUID. Recipient-scoped stored rows retain their stored UUIDs across
// polling and webhook delivery. Legacy names are derived from stored rows; an
// AdminBounce row stored as failed renders email.failed in polling while its
// webhook retains email.bounced, both with canonical email.failed.
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

// Bare stored event names returned by message history. In addition to the normal
// send and delivery lifecycle, polling can expose suppression, scan, and
// quarantine lifecycle rows. Sharp canonical names gw_reject, injection_timeout,
// and expired distinguish gateway rejection, ambiguous injection timeout, and MTA
// expiration. The failed and bounced names remain valid for system/admin failures
// and hard bounces respectively. Existing stored rows retain their original names.
type EmailEventType string

const (
	EmailEventTypeQueued                      EmailEventType = "queued"
	EmailEventTypeDeferred                    EmailEventType = "deferred"
	EmailEventTypeScheduled                   EmailEventType = "scheduled"
	EmailEventTypeCancelled                   EmailEventType = "cancelled"
	EmailEventTypeSandbox                     EmailEventType = "sandbox"
	EmailEventTypeSending                     EmailEventType = "sending"
	EmailEventTypeSent                        EmailEventType = "sent"
	EmailEventTypeFailed                      EmailEventType = "failed"
	EmailEventTypeDelivered                   EmailEventType = "delivered"
	EmailEventTypeBounced                     EmailEventType = "bounced"
	EmailEventTypeComplained                  EmailEventType = "complained"
	EmailEventTypeSuppressed                  EmailEventType = "suppressed"
	EmailEventTypeRejected                    EmailEventType = "rejected"
	EmailEventTypeOpened                      EmailEventType = "opened"
	EmailEventTypeClicked                     EmailEventType = "clicked"
	EmailEventTypeUnsubscribed                EmailEventType = "unsubscribed"
	EmailEventTypeDailyLimitExceeded          EmailEventType = "daily_limit_exceeded"
	EmailEventTypeScanDeferred                EmailEventType = "scan_deferred"
	EmailEventTypeQuarantined                 EmailEventType = "quarantined"
	EmailEventTypeQuarantineReleased          EmailEventType = "quarantine_released"
	EmailEventTypeQuarantineReleaseDispatched EmailEventType = "quarantine_release_dispatched"
	EmailEventTypeQuarantineRejected          EmailEventType = "quarantine_rejected"
	EmailEventTypeQuarantineExpired           EmailEventType = "quarantine_expired"
	EmailEventTypeGwReject                    EmailEventType = "gw_reject"
	EmailEventTypeInjectionTimeout            EmailEventType = "injection_timeout"
	EmailEventTypeExpired                     EmailEventType = "expired"
)

type EmailWebhookRecipient struct {
	Email string `json:"email" api:"required" format:"email"`
	// Any of "to", "cc", "bcc".
	Kind EmailWebhookRecipientKind `json:"kind"`
	Name string                    `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailWebhookRecipient) RawJSON() string { return r.JSON.raw }
func (r *EmailWebhookRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailWebhookRecipientKind string

const (
	EmailWebhookRecipientKindTo  EmailWebhookRecipientKind = "to"
	EmailWebhookRecipientKindCc  EmailWebhookRecipientKind = "cc"
	EmailWebhookRecipientKindBcc EmailWebhookRecipientKind = "bcc"
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

// An account-polling event. The envelope is webhook-shaped, but polling preserves
// stored-event cardinality: queued, sending, sandbox, cancelled, and
// daily_limit_exceeded message events fan out per recipient; scheduled remains one
// message-scoped row. Payload fields vary among recipient-scoped, message-scoped,
// and minimal fallback rows.
type EmailEventListResponseData struct {
	// Event UUID.
	ID string `json:"id" api:"required" format:"uuid"`
	// Additive canonical outcome name, prefixed with `email.`. Gateway rejection is
	// `email.gw_reject`, ambiguous injection timeout is `email.injection_timeout`, and
	// MTA expiration is `email.expired`. Unchanged outcomes retain their names.
	// Existing stored rows are translated only when recorded payload evidence proves
	// the outcome; a legacy failed row is not guessed or sharpened.
	CanonicalEventType string `json:"canonical_event_type" api:"required"`
	// Legacy customer-visible event name, prefixed with `email.`. Gateway rejections
	// render `email.failed`; MTA expirations render `email.bounced`. Webhook
	// subscription allowlists match the legacy name.
	EventType  string    `json:"event_type" api:"required"`
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// Payload returned by GET /email_events. Every row includes id, status, and
	// occurred_at. Recipient-scoped rows also include recipient_id, from, subject, and
	// exactly one object-valued to, cc, or bcc field. Legacy or message-scoped rows
	// can omit recipient_id and use object-valued or string-valued to/cc fields,
	// including an empty string when no address exists; bcc is redacted. If the
	// related message or recipient cannot be loaded, the minimal fallback can omit
	// from, subject, and recipient fields. Additional persisted public evidence can be
	// present.
	Payload EmailEventListResponseDataPayload `json:"payload" api:"required"`
	// Durable email recipient UUID. Present for recipient-scoped events, including
	// each queued, sending, sandbox, cancelled, and daily_limit_exceeded fan-out
	// event.
	RecipientID string `json:"recipient_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CanonicalEventType respjson.Field
		EventType          respjson.Field
		OccurredAt         respjson.Field
		Payload            respjson.Field
		RecipientID        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload returned by GET /email_events. Every row includes id, status, and
// occurred_at. Recipient-scoped rows also include recipient_id, from, subject, and
// exactly one object-valued to, cc, or bcc field. Legacy or message-scoped rows
// can omit recipient_id and use object-valued or string-valued to/cc fields,
// including an empty string when no address exists; bcc is redacted. If the
// related message or recipient cannot be loaded, the minimal fallback can omit
// from, subject, and recipient fields. Additional persisted public evidence can be
// present.
type EmailEventListResponseDataPayload struct {
	// Email message UUID.
	ID         string    `json:"id" api:"required" format:"uuid"`
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// Stored event outcome slug, not the authoritative recipient status. Account
	// polling returns the stored name, including suppression, scan, and quarantine
	// lifecycle names. Webhooks retain legacy payload names: gateway rejections use
	// failed and MTA expirations use bounced. New sharp stored rows can expose
	// gw_reject, injection_timeout, or expired. Use the envelope canonical_event_type
	// to identify the outcome across surfaces.
	//
	// Any of "queued", "deferred", "scheduled", "cancelled", "sandbox", "sending",
	// "sent", "failed", "delivered", "bounced", "complained", "suppressed",
	// "rejected", "opened", "clicked", "unsubscribed", "daily_limit_exceeded",
	// "scan_deferred", "quarantined", "quarantine_released",
	// "quarantine_release_dispatched", "quarantine_rejected", "quarantine_expired",
	// "gw_reject", "injection_timeout", "expired".
	Status string                                    `json:"status" api:"required"`
	Bcc    EmailEventListResponseDataPayloadBccUnion `json:"bcc"`
	// Legacy message-scoped address, or an empty string when absent.
	Cc EmailEventListResponseDataPayloadCcUnion `json:"cc"`
	// Sender projection in account event polling. The display name is explicitly null
	// when the message has no sender name.
	From EmailEventListResponseDataPayloadFrom `json:"from"`
	// Durable email recipient UUID. Present for recipient-scoped events.
	RecipientID string `json:"recipient_id" format:"uuid"`
	Subject     string `json:"subject"`
	// Legacy message-scoped address, or an empty string when absent.
	To          EmailEventListResponseDataPayloadToUnion `json:"to"`
	ExtraFields map[string]any                           `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		OccurredAt  respjson.Field
		Status      respjson.Field
		Bcc         respjson.Field
		Cc          respjson.Field
		From        respjson.Field
		RecipientID respjson.Field
		Subject     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponseDataPayload) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponseDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EmailEventListResponseDataPayloadBccUnion contains all possible properties and
// values from [EmailWebhookRecipient], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEmailEventListResponseDataPayloadBccString]
type EmailEventListResponseDataPayloadBccUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfEmailEventListResponseDataPayloadBccString string `json:",inline"`
	// This field is from variant [EmailWebhookRecipient].
	Email string `json:"email"`
	// This field is from variant [EmailWebhookRecipient].
	Kind EmailWebhookRecipientKind `json:"kind"`
	// This field is from variant [EmailWebhookRecipient].
	Name string `json:"name"`
	JSON struct {
		OfEmailEventListResponseDataPayloadBccString respjson.Field
		Email                                        respjson.Field
		Kind                                         respjson.Field
		Name                                         respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u EmailEventListResponseDataPayloadBccUnion) AsEmailWebhookRecipient() (v EmailWebhookRecipient) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailEventListResponseDataPayloadBccUnion) AsEmailEventListResponseDataPayloadBccString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EmailEventListResponseDataPayloadBccUnion) RawJSON() string { return u.JSON.raw }

func (r *EmailEventListResponseDataPayloadBccUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailEventListResponseDataPayloadBccString string

const (
	EmailEventListResponseDataPayloadBccStringRedacted EmailEventListResponseDataPayloadBccString = "redacted"
)

// EmailEventListResponseDataPayloadCcUnion contains all possible properties and
// values from [EmailWebhookRecipient], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type EmailEventListResponseDataPayloadCcUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [EmailWebhookRecipient].
	Email string `json:"email"`
	// This field is from variant [EmailWebhookRecipient].
	Kind EmailWebhookRecipientKind `json:"kind"`
	// This field is from variant [EmailWebhookRecipient].
	Name string `json:"name"`
	JSON struct {
		OfString respjson.Field
		Email    respjson.Field
		Kind     respjson.Field
		Name     respjson.Field
		raw      string
	} `json:"-"`
}

func (u EmailEventListResponseDataPayloadCcUnion) AsEmailWebhookRecipient() (v EmailWebhookRecipient) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailEventListResponseDataPayloadCcUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EmailEventListResponseDataPayloadCcUnion) RawJSON() string { return u.JSON.raw }

func (r *EmailEventListResponseDataPayloadCcUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sender projection in account event polling. The display name is explicitly null
// when the message has no sender name.
type EmailEventListResponseDataPayloadFrom struct {
	Email string `json:"email" api:"required" format:"email"`
	Name  string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailEventListResponseDataPayloadFrom) RawJSON() string { return r.JSON.raw }
func (r *EmailEventListResponseDataPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EmailEventListResponseDataPayloadToUnion contains all possible properties and
// values from [EmailWebhookRecipient], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type EmailEventListResponseDataPayloadToUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [EmailWebhookRecipient].
	Email string `json:"email"`
	// This field is from variant [EmailWebhookRecipient].
	Kind EmailWebhookRecipientKind `json:"kind"`
	// This field is from variant [EmailWebhookRecipient].
	Name string `json:"name"`
	JSON struct {
		OfString respjson.Field
		Email    respjson.Field
		Kind     respjson.Field
		Name     respjson.Field
		raw      string
	} `json:"-"`
}

func (u EmailEventListResponseDataPayloadToUnion) AsEmailWebhookRecipient() (v EmailWebhookRecipient) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailEventListResponseDataPayloadToUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EmailEventListResponseDataPayloadToUnion) RawJSON() string { return u.JSON.raw }

func (r *EmailEventListResponseDataPayloadToUnion) UnmarshalJSON(data []byte) error {
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
	// Number of results to return. Defaults to 25; maximum is 100. Invalid values are
	// clamped to the valid range.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Opaque URL-safe Base64 cursor returned by a previous event list response. The
	// legacy `page[after]` and flat `page_cursor` forms are also accepted.
	PageCursor param.Opt[string] `query:"page[cursor],omitzero" json:"-"`
	// Inclusive ISO 8601 end timestamp. When `from` is provided without `to`, defaults
	// to `from + 30 days`.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Comma-separated list of event types to include. Also accepts repeated query
	// parameters (e.g. event_type=delivered&event_type=bounced). Unknown values return
	// no matches.
	//
	// Dual-name compatibility: values are accepted bare or `email.`-prefixed. A legacy
	// value keeps matching the rows it matched pre-rename — no widening: `failed` also
	// matches the rows that now store the canonical names of the outcomes it covered
	// (`gw_reject`, `injection_timeout`, `expired`); `bounced` matches stored
	// `bounced` rows only (recipient-scoped Expirations stored `failed` pre-rename and
	// never matched `bounced`, so `expired` is deliberately not a `bounced`
	// expansion). A canonical value matches its own rows plus legacy rows whose
	// recorded payload evidence proves that outcome (`expired` also surfaces legacy
	// `bounced` rows with `bounce_category: transient`). The additive
	// `canonical_event_type` field in each response row names the canonical outcome.
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
