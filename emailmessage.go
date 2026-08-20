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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Send and manage email messages. Legacy `/v2/emails` routes are aliases for these
// endpoints.
//
// EmailMessageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailMessageService] method instead.
type EmailMessageService struct {
	Options []option.RequestOption
	// Send and manage email messages. Legacy `/v2/emails` routes are aliases for these
	// endpoints.
	Recipients EmailMessageRecipientService
}

// NewEmailMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailMessageService(opts ...option.RequestOption) (r EmailMessageService) {
	r = EmailMessageService{}
	r.Options = opts
	r.Recipients = NewEmailMessageRecipientService(opts...)
	return
}

// Queues, schedules, or sandbox-sends an email message. The legacy `/v2/emails`
// POST route is a backward-compatible alias for this operation.
//
// `subject` is required unless `template_id` is supplied. When using
// `template_id`, do not also provide `subject`, `html_body`, or `text_body`; the
// template is rendered with `template_variables`.
//
// Note: template lookup failures (not found, wrong account) return 400, not 404.
func (r *EmailMessageService) New(ctx context.Context, params EmailMessageNewParams, opts ...option.RequestOption) (res *EmailMessageResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "email_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// The legacy `/v2/emails/{id}` GET route is a backward-compatible alias for this
// operation.
func (r *EmailMessageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailMessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists messages sorted newest first by `created_at desc, id desc`. No filters
// other than cursor pagination are implemented. The legacy `/v2/emails` GET route
// is a backward-compatible alias for this operation.
func (r *EmailMessageService) List(ctx context.Context, query EmailMessageListParams, opts ...option.RequestOption) (res *EmailMessageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes an account-scoped email message, its events, its durable
// recipients, and unshared attachment objects. Returns 404 when the message does
// not exist in the authenticated account. The legacy `/v2/emails/{id}` DELETE
// route is a backward-compatible alias.
func (r *EmailMessageService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("email_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates up to 50 email messages in a single request.
func (r *EmailMessageService) Batch(ctx context.Context, params EmailMessageBatchParams, opts ...option.RequestOption) (res *EmailMessageBatchResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "email_messages/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Permanently deletes every email in the authenticated account sent from or to the
// supplied address, including retained events whose parent message has expired.
// Events and durable recipients are deleted immediately with each message. The
// operation never searches or reports matches in another account. The legacy
// `/v2/emails` DELETE route is a backward-compatible alias.
func (r *EmailMessageService) DeleteAll(ctx context.Context, body EmailMessageDeleteAllParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "email_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Cancels a scheduled email and returns it with status `cancelled`. The legacy
// `/v2/emails/{id}/schedule` DELETE route is an alias.
func (r *EmailMessageService) DeleteSchedule(ctx context.Context, emailID string, opts ...option.RequestOption) (res *EmailMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailID == "" {
		err = errors.New("missing required email_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_messages/%s/schedule", emailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists events for a single message sorted oldest first by
// `occurred_at asc, id asc`. The legacy `/v2/emails/{id}/events` GET route is a
// backward-compatible alias.
func (r *EmailMessageService) GetEvents(ctx context.Context, emailID string, query EmailMessageGetEventsParams, opts ...option.RequestOption) (res *EmailMessageGetEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailID == "" {
		err = errors.New("missing required email_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_messages/%s/events", emailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AttachmentRequestParam struct {
	// MIME Content-ID used to reference an inline attachment.
	ContentID param.Opt[string] `json:"content_id,omitzero"`
	// Attachment content, typically Base64-encoded. Defaults to empty string when
	// omitted.
	Content param.Opt[string] `json:"content,omitzero"`
	// MIME content type. Defaults to "application/octet-stream" when omitted.
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// MIME disposition (`attachment` or `inline`).
	Disposition param.Opt[string] `json:"disposition,omitzero"`
	// Attachment filename. Defaults to "attachment" when omitted.
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r AttachmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AttachmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func EmailAddressInputParamOfEmailAddress(email string) EmailAddressInputUnionParam {
	var variant EmailAddressParam
	variant.Email = email
	return EmailAddressInputUnionParam{OfEmailAddress: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmailAddressInputUnionParam struct {
	OfString       param.Opt[string]  `json:",omitzero,inline"`
	OfEmailAddress *EmailAddressParam `json:",omitzero,inline"`
	paramUnion
}

func (u EmailAddressInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEmailAddress)
}
func (u *EmailAddressInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmailAddressInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEmailAddress) {
		return u.OfEmailAddress
	}
	return nil
}

type MessageEvent struct {
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// Any of "queued", "deferred", "scheduled", "cancelled", "sandbox", "sending",
	// "sent", "failed", "delivered", "bounced", "complained", "rejected", "opened",
	// "clicked", "unsubscribed", "daily_limit_exceeded".
	Type    EmailEventType `json:"type" api:"required"`
	Payload map[string]any `json:"payload"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OccurredAt  respjson.Field
		Type        respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressedRecipient struct {
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
func (r SuppressedRecipient) RawJSON() string { return r.JSON.raw }
func (r *SuppressedRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-send open and click tracking overrides. Omitted properties inherit the
// sender domain's tracking settings.
type TrackingSettingsParam struct {
	// Whether to rewrite links for click tracking in this message.
	ClickTracking param.Opt[bool] `json:"click_tracking,omitzero"`
	// Whether to inject an open-tracking pixel for this message.
	OpenTracking param.Opt[bool] `json:"open_tracking,omitzero"`
	paramObj
}

func (r TrackingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TrackingSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageGetResponse struct {
	Data EmailMessageGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageGetResponseData struct {
	// HTML body submitted for the message.
	HTMLBody string `json:"html_body" api:"required"`
	// Plain-text body submitted for the message.
	TextBody string `json:"text_body" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTMLBody    respjson.Field
		TextBody    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EmailMessage
}

// Returns the unmodified JSON received from the API
func (r EmailMessageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageListResponse struct {
	Data []EmailMessage      `json:"data" api:"required"`
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
func (r EmailMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageBatchResponse struct {
	Data   []EmailMessage                   `json:"data" api:"required"`
	Errors []EmailMessageBatchResponseError `json:"errors" api:"required"`
	Meta   EmailMessageBatchResponseMeta    `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageBatchResponseError struct {
	// Batch item errors use `message` (not `detail`) for the human-readable text.
	//
	// Any of "bad_request", "not_found", "forbidden", "service_unavailable",
	// "validation_error", "recipient_suppressed", "reputation_suspended".
	Code string `json:"code" api:"required"`
	// Zero-based index of the failed message in the request array.
	Index   int64  `json:"index" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Index       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageBatchResponseError) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageBatchResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageBatchResponseMeta struct {
	Failed    int64 `json:"failed" api:"required"`
	Succeeded int64 `json:"succeeded" api:"required"`
	Total     int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Succeeded   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageBatchResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageBatchResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageGetEventsResponse struct {
	Data []MessageEvent      `json:"data" api:"required"`
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
func (r EmailMessageGetEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageGetEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageNewParams struct {
	From EmailAddressInputUnionParam   `json:"from,omitzero" api:"required"`
	To   []EmailAddressInputUnionParam `json:"to,omitzero" api:"required"`
	// Telnyx message UUID of the message this send forwards. Forwarded messages start
	// a NEW thread per RFC 5322 — NO `In-Reply-To` or `References` headers are set on
	// the outbound MIME. The id is recorded in the message's metadata for EDR
	// provenance only.
	//
	// The id is validated as a UUID but is NOT looked up against the message store —
	// existence is the caller's responsibility (the forward is pure metadata; it does
	// not affect delivery). Cannot be combined with `in_reply_to_message_id` (422).
	ForwardOfMessageID param.Opt[string] `json:"forward_of_message_id,omitzero" format:"uuid"`
	// Optional unsubscribe-group UUID used for group-scoped suppression checks and
	// unsubscribe handling.
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Telnyx message UUID of the message this send replies to. When provided, the API
	// sets RFC 5322 `In-Reply-To` and `References` headers on the outbound MIME so the
	// recipient's mailbox (Gmail/Outlook) threads it correctly. The parent is looked
	// up under the caller's account scope; a UUID belonging to another account yields
	// a non-enumerating 404.
	//
	// Wire-only (Phase 1): the API sets the headers and does NOT resolve or mutate
	// `thread_id` on the server side. Messages sent without this parameter are
	// standalone (no threading headers injected).
	//
	// Cannot be combined with `forward_of_message_id` (422).
	InReplyToMessageID param.Opt[string] `json:"in_reply_to_message_id,omitzero" format:"uuid"`
	// Indicates a reply-all intent. In Phase 1 (wire-only) this does not change the
	// threading headers — recipient selection is customer- controlled (`to`/`cc`), and
	// a thread is not defined by its audience. When the referenced message has no
	// thread context, reply-all degrades to a plain reply (parent ID only in
	// `References`). The resolution engine (separate work) will expand the ancestor
	// chain at a later phase with no API change.
	//
	// Only meaningful alongside `in_reply_to_message_id`.
	ReplyToAll param.Opt[bool] `json:"reply_to_all,omitzero"`
	// Future ISO 8601 time to schedule sending. Invalid or past timestamps are
	// silently ignored and the email is sent immediately. The legacy alias `send_at`
	// is still accepted for backward compatibility; when both are provided,
	// `scheduled_at` wins.
	ScheduledAt param.Opt[time.Time] `json:"scheduled_at,omitzero" format:"date-time"`
	// Optional display name for string `from`; overrides `from.name` when provided.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// HTML email body. Returned only by `GET /email_messages/{id}`; omitted from
	// create and list responses.
	HTMLBody param.Opt[string] `json:"html_body,omitzero"`
	// When true, allows delivery to recipients whose suppressions explicitly permit an
	// override. Hard bounces, spam complaints, and invalid-address suppressions cannot
	// be overridden. Requires the `email:override` API scope.
	IgnoreSuppression param.Opt[bool] `json:"ignore_suppression,omitzero"`
	InlineCss         param.Opt[bool] `json:"inline_css,omitzero"`
	SandboxMode       param.Opt[bool] `json:"sandbox_mode,omitzero"`
	// Deprecated alias for `scheduled_at`.
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	// Required unless `template_id` is supplied. When using a template, the template's
	// subject is rendered; if the template has no subject or renders empty, the
	// request returns 400.
	Subject    param.Opt[string] `json:"subject,omitzero"`
	TemplateID param.Opt[string] `json:"template_id,omitzero" format:"uuid"`
	// Plain text email body. Returned only by `GET /email_messages/{id}`; omitted from
	// create and list responses.
	TextBody       param.Opt[string]             `json:"text_body,omitzero"`
	IdempotencyKey param.Opt[string]             `header:"Idempotency-Key,omitzero" json:"-"`
	Attachments    []AttachmentRequestParam      `json:"attachments,omitzero"`
	Bcc            []EmailAddressInputUnionParam `json:"bcc,omitzero"`
	Cc             []EmailAddressInputUnionParam `json:"cc,omitzero"`
	// Custom email headers. Write-only; not returned in responses.
	Headers map[string]string `json:"headers,omitzero"`
	// Custom metadata. Write-only; not returned in responses.
	Metadata map[string]any `json:"metadata,omitzero"`
	// Reply-to address. If provided as an object with a name, only the email is
	// stored; the name is ignored.
	ReplyTo EmailAddressInputUnionParam `json:"reply_to,omitzero"`
	// Tags for categorization and reporting. Stored on the message and propagated to
	// Email Detail Records. Not returned in API responses.
	Tags []string `json:"tags,omitzero"`
	// Variables for Liquid template rendering. Non-object values may cause a 422
	// validation error on message creation, but are silently treated as an empty
	// object for template rendering.
	TemplateVariables map[string]any `json:"template_variables,omitzero"`
	// Per-send open and click tracking overrides. Omitted properties inherit the
	// sender domain's tracking settings.
	TrackingSettings TrackingSettingsParam `json:"tracking_settings,omitzero"`
	paramObj
}

func (r EmailMessageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailMessageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailMessageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageListParams struct {
	// Opaque URL-safe Base64 cursor returned by a previous list response.
	PageCursor param.Opt[string] `query:"page_cursor,omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100. Invalid values are
	// clamped to the valid range.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailMessageListParams]'s query parameters as `url.Values`.
func (r EmailMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailMessageBatchParams struct {
	Messages []EmailMessageBatchParamsMessage `json:"messages,omitzero" api:"required"`
	// Applies sandbox mode to all messages in the batch. Overrides any per-message
	// sandbox_mode in the messages array.
	SandboxMode    param.Opt[bool]   `json:"sandbox_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r EmailMessageBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailMessageBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailMessageBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single message in a batch create request. This schema mirrors
// `CreateEmailRequest` EXCEPT it does not accept the reply/forward threading
// parameters (`in_reply_to_message_id`, `reply_to_all`, `forward_of_message_id`) —
// those are single-send-only in Phase 1 (MSG-1491) and are not yet implemented on
// the batch endpoint. Recipient email addresses must be unique across `to`, `cc`,
// and `bcc` after case-insensitive normalization. Duplicate recipients return
// `400`.
//
// The properties From, To are required.
type EmailMessageBatchParamsMessage struct {
	From EmailAddressInputUnionParam   `json:"from,omitzero" api:"required"`
	To   []EmailAddressInputUnionParam `json:"to,omitzero" api:"required"`
	// Optional unsubscribe-group UUID used for group-scoped suppression checks and
	// unsubscribe handling.
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Future ISO 8601 time to schedule sending. Invalid or past timestamps are
	// silently ignored and the email is sent immediately. The legacy alias `send_at`
	// is still accepted for backward compatibility; when both are provided,
	// `scheduled_at` wins.
	ScheduledAt param.Opt[time.Time] `json:"scheduled_at,omitzero" format:"date-time"`
	// Optional display name for string `from`; overrides `from.name` when provided.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// HTML email body. Returned only by `GET /email_messages/{id}`; omitted from
	// create and list responses.
	HTMLBody param.Opt[string] `json:"html_body,omitzero"`
	// When true, allows delivery to recipients whose suppressions explicitly permit an
	// override. Hard bounces, spam complaints, and invalid-address suppressions cannot
	// be overridden. Requires the `email:override` API scope.
	IgnoreSuppression param.Opt[bool] `json:"ignore_suppression,omitzero"`
	InlineCss         param.Opt[bool] `json:"inline_css,omitzero"`
	SandboxMode       param.Opt[bool] `json:"sandbox_mode,omitzero"`
	// Deprecated alias for `scheduled_at`.
	//
	// Deprecated: deprecated
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	// Required unless `template_id` is supplied. When using a template, the template's
	// subject is rendered; if the template has no subject or renders empty, the
	// request returns 400.
	Subject    param.Opt[string] `json:"subject,omitzero"`
	TemplateID param.Opt[string] `json:"template_id,omitzero" format:"uuid"`
	// Plain text email body. Returned only by `GET /email_messages/{id}`; omitted from
	// create and list responses.
	TextBody    param.Opt[string]             `json:"text_body,omitzero"`
	Attachments []AttachmentRequestParam      `json:"attachments,omitzero"`
	Bcc         []EmailAddressInputUnionParam `json:"bcc,omitzero"`
	Cc          []EmailAddressInputUnionParam `json:"cc,omitzero"`
	// Custom email headers. Write-only; not returned in responses.
	Headers map[string]string `json:"headers,omitzero"`
	// Custom metadata. Write-only; not returned in responses.
	Metadata map[string]any `json:"metadata,omitzero"`
	// Reply-to address. If provided as an object with a name, only the email is
	// stored; the name is ignored.
	ReplyTo EmailAddressInputUnionParam `json:"reply_to,omitzero"`
	// Tags for categorization and reporting. Stored on the message and propagated to
	// Email Detail Records. Not returned in API responses.
	Tags []string `json:"tags,omitzero"`
	// Variables for Liquid template rendering. Non-object values may cause a 422
	// validation error on message creation, but are silently treated as an empty
	// object for template rendering.
	TemplateVariables map[string]any `json:"template_variables,omitzero"`
	// Per-send open and click tracking overrides. Omitted properties inherit the
	// sender domain's tracking settings.
	TrackingSettings TrackingSettingsParam `json:"tracking_settings,omitzero"`
	paramObj
}

func (r EmailMessageBatchParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow EmailMessageBatchParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailMessageBatchParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageDeleteAllParams struct {
	// Sender or recipient address to delete. Matching is trimmed and case-insensitive.
	Address string `query:"address" api:"required" format:"email" json:"-"`
	paramObj
}

// URLQuery serializes [EmailMessageDeleteAllParams]'s query parameters as
// `url.Values`.
func (r EmailMessageDeleteAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailMessageGetEventsParams struct {
	// Opaque URL-safe Base64 cursor returned by a previous list response.
	PageCursor param.Opt[string] `query:"page_cursor,omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100. Invalid values are
	// clamped to the valid range.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailMessageGetEventsParams]'s query parameters as
// `url.Values`.
func (r EmailMessageGetEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
