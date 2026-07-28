// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create, list, retrieve, update, delete, and send unsent draft messages belonging
// to an agent inbox.
//
// EmailInboxDraftService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxDraftService] method instead.
type EmailInboxDraftService struct {
	Options []option.RequestOption
}

// NewEmailInboxDraftService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailInboxDraftService(opts ...option.RequestOption) (r EmailInboxDraftService) {
	r = EmailInboxDraftService{}
	r.Options = opts
	return
}

// Creates an unsent draft in the inbox. Every field is optional — a draft is a
// work-in-progress and may be saved incomplete. Send-time requirements (sender,
// subject, at least one recipient) are enforced when the draft is sent, not when
// it is created.
//
// Drafts are unbillable and emit no Email Detail Records until they are sent.
func (r *EmailInboxDraftService) New(ctx context.Context, inboxID string, body EmailInboxDraftNewParams, opts ...option.RequestOption) (res *EmailDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/drafts", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single draft. Drafts that have been sent remain retrievable, so the
// exact content that was sent stays auditable.
func (r *EmailInboxDraftService) Get(ctx context.Context, draftID string, query EmailInboxDraftGetParams, opts ...option.RequestOption) (res *EmailDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/drafts/%s", query.InboxID, draftID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the supplied fields on a draft. `account_id` and `inbox_id` are
// server-owned and ignored if present in the body, so a draft can never be moved
// between accounts or inboxes.
//
// A draft that is being sent or has already been sent is immutable and returns 422
// — modifying it would race with delivery or rewrite the record of what was
// actually sent.
func (r *EmailInboxDraftService) Update(ctx context.Context, draftID string, params EmailInboxDraftUpdateParams, opts ...option.RequestOption) (res *EmailDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/drafts/%s", params.InboxID, draftID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Lists drafts newest first using stable cursor pagination. All access is scoped
// to the authenticated account and the given inbox.
func (r *EmailInboxDraftService) List(ctx context.Context, inboxID string, query EmailInboxDraftListParams, opts ...option.RequestOption) (res *EmailInboxDraftListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/drafts", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes an unsent draft. Drafts that are being sent or have been
// sent cannot be deleted; sent drafts are retained for audit.
func (r *EmailInboxDraftService) Delete(ctx context.Context, draftID string, body EmailInboxDraftDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return err
	}
	path := fmt.Sprintf("email_inboxes/%s/drafts/%s", body.InboxID, draftID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Sends the draft through the standard send pipeline — the same domain resolution,
// suppression, reputation, daily-quota, persistence and Detail Record behaviour as
// `POST /v2/email_messages`. The response body is the created email message.
//
// If the draft has no explicit `from_email`, the inbox address is used.
//
// The draft is marked `sent` only after the send is accepted; a send rejected for
// suppression, quota or reputation leaves the draft editable so it can be fixed
// and retried. A draft that is already `sent` returns 422 rather than sending
// twice.
func (r *EmailInboxDraftService) Send(ctx context.Context, draftID string, body EmailInboxDraftSendParams, opts ...option.RequestOption) (res *EmailMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/drafts/%s/send", body.InboxID, draftID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type EmailAddress struct {
	Email string `json:"email" api:"required"`
	Name  string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailAddress) RawJSON() string { return r.JSON.raw }
func (r *EmailAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailAddress to a EmailAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailAddressParam.Overrides()
func (r EmailAddress) ToParam() EmailAddressParam {
	return param.Override[EmailAddressParam](json.RawMessage(r.RawJSON()))
}

// The property Email is required.
type EmailAddressParam struct {
	Email string            `json:"email" api:"required"`
	Name  param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r EmailAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An unsent, mutable draft message belonging to an inbox.
type EmailDraft struct {
	ID      string `json:"id" api:"required" format:"uuid"`
	InboxID string `json:"inbox_id" api:"required" format:"uuid"`
	// Any of "email_draft".
	RecordType EmailDraftRecordType `json:"record_type" api:"required"`
	// `draft` until the draft is sent. A sent draft is retained for audit and becomes
	// immutable.
	//
	// Any of "draft", "sending", "sent".
	Status      EmailDraftStatus `json:"status" api:"required"`
	Attachments []any            `json:"attachments"`
	Bcc         []EmailAddress   `json:"bcc"`
	Cc          []EmailAddress   `json:"cc"`
	CreatedAt   time.Time        `json:"created_at" format:"date-time"`
	// Sender address. Defaults to the inbox address at send time when null.
	From     string `json:"from" api:"nullable"`
	FromName string `json:"from_name" api:"nullable"`
	// Custom headers. Reply drafts carry `In-Reply-To` and `References`.
	Headers  map[string]string `json:"headers"`
	HTMLBody string            `json:"html_body" api:"nullable"`
	// Mutable mailbox-state labels. Not propagated to Email Detail Records.
	Labels []string `json:"labels"`
	// Arbitrary customer-defined metadata.
	Metadata any    `json:"metadata"`
	ReplyTo  string `json:"reply_to" api:"nullable"`
	// Inbound message this draft replies to. Server-owned; set only on reply drafts.
	ReplyToMessageID string    `json:"reply_to_message_id" api:"nullable" format:"uuid"`
	SentAt           time.Time `json:"sent_at" api:"nullable" format:"date-time"`
	// The email message created when this draft was sent.
	SentMessageID string `json:"sent_message_id" api:"nullable" format:"uuid"`
	Subject       string `json:"subject" api:"nullable"`
	// Transport/reporting attribution tags, propagated to Email Detail Records at send
	// time.
	Tags     []string `json:"tags"`
	TextBody string   `json:"text_body" api:"nullable"`
	// Conversation thread inherited from the parent message.
	ThreadID  string         `json:"thread_id" api:"nullable" format:"uuid"`
	To        []EmailAddress `json:"to"`
	UpdatedAt time.Time      `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		InboxID          respjson.Field
		RecordType       respjson.Field
		Status           respjson.Field
		Attachments      respjson.Field
		Bcc              respjson.Field
		Cc               respjson.Field
		CreatedAt        respjson.Field
		From             respjson.Field
		FromName         respjson.Field
		Headers          respjson.Field
		HTMLBody         respjson.Field
		Labels           respjson.Field
		Metadata         respjson.Field
		ReplyTo          respjson.Field
		ReplyToMessageID respjson.Field
		SentAt           respjson.Field
		SentMessageID    respjson.Field
		Subject          respjson.Field
		Tags             respjson.Field
		TextBody         respjson.Field
		ThreadID         respjson.Field
		To               respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDraft) RawJSON() string { return r.JSON.raw }
func (r *EmailDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDraftRecordType string

const (
	EmailDraftRecordTypeEmailDraft EmailDraftRecordType = "email_draft"
)

// `draft` until the draft is sent. A sent draft is retained for audit and becomes
// immutable.
type EmailDraftStatus string

const (
	EmailDraftStatusDraft   EmailDraftStatus = "draft"
	EmailDraftStatusSending EmailDraftStatus = "sending"
	EmailDraftStatusSent    EmailDraftStatus = "sent"
)

// All fields are optional — a draft may be saved incomplete. `account_id`,
// `inbox_id`, `status`, `sent_at`, `sent_message_id`, `reply_to_message_id` and
// `thread_id` are server-owned and ignored if supplied.
type EmailDraftRequestParam struct {
	FromEmail param.Opt[string] `json:"from_email,omitzero"`
	FromName  param.Opt[string] `json:"from_name,omitzero"`
	// Alias for `html_body`, matching the send endpoint.
	HTML     param.Opt[string] `json:"html,omitzero"`
	HTMLBody param.Opt[string] `json:"html_body,omitzero"`
	ReplyTo  param.Opt[string] `json:"reply_to,omitzero"`
	Subject  param.Opt[string] `json:"subject,omitzero"`
	// Alias for `text_body`, matching the send endpoint.
	Text        param.Opt[string]             `json:"text,omitzero"`
	TextBody    param.Opt[string]             `json:"text_body,omitzero"`
	Attachments []any                         `json:"attachments,omitzero"`
	Bcc         []EmailAddressInputUnionParam `json:"bcc,omitzero"`
	Cc          []EmailAddressInputUnionParam `json:"cc,omitzero"`
	Headers     map[string]string             `json:"headers,omitzero"`
	Labels      []string                      `json:"labels,omitzero"`
	Metadata    any                           `json:"metadata,omitzero"`
	Tags        []string                      `json:"tags,omitzero"`
	To          []EmailAddressInputUnionParam `json:"to,omitzero"`
	paramObj
}

func (r EmailDraftRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailDraftRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDraftRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDraftResponse struct {
	// An unsent, mutable draft message belonging to an inbox.
	Data EmailDraft `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessage struct {
	ID          string                   `json:"id" api:"required" format:"uuid"`
	Attachments []EmailMessageAttachment `json:"attachments" api:"required"`
	Bcc         []EmailAddress           `json:"bcc" api:"required"`
	Cc          []EmailAddress           `json:"cc" api:"required"`
	CreatedAt   time.Time                `json:"created_at" api:"required" format:"date-time"`
	Events      []MessageEvent           `json:"events" api:"required"`
	From        EmailAddress             `json:"from" api:"required"`
	// Any of "email_message".
	RecordType EmailMessageRecordType `json:"record_type" api:"required"`
	ReplyTo    string                 `json:"reply_to" api:"required"`
	// Current status of an email message. Lifecycle statuses (queued, scheduled, etc.)
	// are set on creation. Delivery statuses (delivered, bounced, etc.) are updated by
	// delivery event consumers.
	//
	// Any of "queued", "scheduled", "cancelled", "sandbox", "sending", "sent",
	// "failed", "deferred", "delivered", "bounced", "complained", "rejected",
	// "opened", "clicked", "unsubscribed".
	Status            EmailMessageStatus `json:"status" api:"required"`
	Subject           string             `json:"subject" api:"required"`
	TemplateID        string             `json:"template_id" api:"required" format:"uuid"`
	TemplateVariables map[string]any     `json:"template_variables" api:"required"`
	To                []EmailAddress     `json:"to" api:"required"`
	// Present when true in the immediate create response. Not persisted; absent on
	// subsequent GET requests.
	InlineCss bool `json:"inline_css"`
	// Per-status recipient counts for the message. Present only for outbound messages
	// with recipient rows. Keys are recipient statuses, values are counts. Example:
	// `{"delivered": 998, "bounced": 2}`.
	RecipientStatuses map[string]int64 `json:"recipient_statuses"`
	// Present when sandbox mode was used.
	Sandbox bool `json:"sandbox"`
	// Present when a scheduled_at value was stored. Persists even after the scheduled
	// send has been processed or cancelled.
	ScheduledAt time.Time `json:"scheduled_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Attachments       respjson.Field
		Bcc               respjson.Field
		Cc                respjson.Field
		CreatedAt         respjson.Field
		Events            respjson.Field
		From              respjson.Field
		RecordType        respjson.Field
		ReplyTo           respjson.Field
		Status            respjson.Field
		Subject           respjson.Field
		TemplateID        respjson.Field
		TemplateVariables respjson.Field
		To                respjson.Field
		InlineCss         respjson.Field
		RecipientStatuses respjson.Field
		Sandbox           respjson.Field
		ScheduledAt       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessage) RawJSON() string { return r.JSON.raw }
func (r *EmailMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EDR-aligned attachment metadata. The base64 `content` is never returned.
type EmailMessageAttachment struct {
	// MIME Content-ID for inline references.
	ContentID   string `json:"content_id" api:"required"`
	ContentType string `json:"content_type" api:"required"`
	// MIME disposition (e.g. `attachment` or `inline`). Runtime passes through the
	// stored value without enforcing an enum.
	Disposition string `json:"disposition" api:"required"`
	Filename    string `json:"filename" api:"required"`
	// SHA-256 hex digest of the attachment content.
	Sha256 string `json:"sha256" api:"required"`
	// Attachment size in bytes.
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Telnyx-hosted public URL for the attachment content.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentID   respjson.Field
		ContentType respjson.Field
		Disposition respjson.Field
		Filename    respjson.Field
		Sha256      respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageRecordType string

const (
	EmailMessageRecordTypeEmailMessage EmailMessageRecordType = "email_message"
)

// Current status of an email message. Lifecycle statuses (queued, scheduled, etc.)
// are set on creation. Delivery statuses (delivered, bounced, etc.) are updated by
// delivery event consumers.
type EmailMessageStatus string

const (
	EmailMessageStatusQueued       EmailMessageStatus = "queued"
	EmailMessageStatusScheduled    EmailMessageStatus = "scheduled"
	EmailMessageStatusCancelled    EmailMessageStatus = "cancelled"
	EmailMessageStatusSandbox      EmailMessageStatus = "sandbox"
	EmailMessageStatusSending      EmailMessageStatus = "sending"
	EmailMessageStatusSent         EmailMessageStatus = "sent"
	EmailMessageStatusFailed       EmailMessageStatus = "failed"
	EmailMessageStatusDeferred     EmailMessageStatus = "deferred"
	EmailMessageStatusDelivered    EmailMessageStatus = "delivered"
	EmailMessageStatusBounced      EmailMessageStatus = "bounced"
	EmailMessageStatusComplained   EmailMessageStatus = "complained"
	EmailMessageStatusRejected     EmailMessageStatus = "rejected"
	EmailMessageStatusOpened       EmailMessageStatus = "opened"
	EmailMessageStatusClicked      EmailMessageStatus = "clicked"
	EmailMessageStatusUnsubscribed EmailMessageStatus = "unsubscribed"
)

type EmailMessageResponse struct {
	Data EmailMessage `json:"data" api:"required"`
	// Recipients removed by suppression checks when at least one recipient remains and
	// the message is accepted.
	Suppressed []EmailMessageResponseSuppressed `json:"suppressed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Suppressed  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageResponseSuppressed struct {
	// Whether an authorized send may override this suppression.
	OverrideAllowed bool `json:"override_allowed" api:"required"`
	// Suppression reason returned by the recipient suppression service.
	Reason string `json:"reason" api:"required"`
	// Scope at which the suppression applies.
	Scope string `json:"scope" api:"required"`
	// Suppressed recipient email address.
	To string `json:"to" api:"required" format:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OverrideAllowed respjson.Field
		Reason          respjson.Field
		Scope           respjson.Field
		To              respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageResponseSuppressed) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageResponseSuppressed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxDraftListResponse struct {
	Data []EmailDraft        `json:"data" api:"required"`
	Meta EmailPaginationMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxDraftListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxDraftListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxDraftNewParams struct {
	// All fields are optional — a draft may be saved incomplete. `account_id`,
	// `inbox_id`, `status`, `sent_at`, `sent_message_id`, `reply_to_message_id` and
	// `thread_id` are server-owned and ignored if supplied.
	EmailDraftRequest EmailDraftRequestParam
	paramObj
}

func (r EmailInboxDraftNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailDraftRequest)
}
func (r *EmailInboxDraftNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxDraftGetParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type EmailInboxDraftUpdateParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// All fields are optional — a draft may be saved incomplete. `account_id`,
	// `inbox_id`, `status`, `sent_at`, `sent_message_id`, `reply_to_message_id` and
	// `thread_id` are server-owned and ignored if supplied.
	EmailDraftRequest EmailDraftRequestParam
	paramObj
}

func (r EmailInboxDraftUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailDraftRequest)
}
func (r *EmailInboxDraftUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxDraftListParams struct {
	// Opaque cursor returned by the previous page.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Restrict results to drafts in this state.
	//
	// Any of "draft", "sending", "sent".
	FilterStatus EmailInboxDraftListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailInboxDraftListParams]'s query parameters as
// `url.Values`.
func (r EmailInboxDraftListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restrict results to drafts in this state.
type EmailInboxDraftListParamsFilterStatus string

const (
	EmailInboxDraftListParamsFilterStatusDraft   EmailInboxDraftListParamsFilterStatus = "draft"
	EmailInboxDraftListParamsFilterStatusSending EmailInboxDraftListParamsFilterStatus = "sending"
	EmailInboxDraftListParamsFilterStatusSent    EmailInboxDraftListParamsFilterStatus = "sent"
)

type EmailInboxDraftDeleteParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type EmailInboxDraftSendParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
