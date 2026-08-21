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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create and manage agent inboxes, retrieve inbound messages and threads, and
// reply to or forward messages.
//
// EmailInboxThreadService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxThreadService] method instead.
type EmailInboxThreadService struct {
	Options []option.RequestOption
	// Create and manage agent inboxes, retrieve inbound messages and threads, and
	// reply to or forward messages.
	Labels EmailInboxThreadLabelService
}

// NewEmailInboxThreadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailInboxThreadService(opts ...option.RequestOption) (r EmailInboxThreadService) {
	r = EmailInboxThreadService{}
	r.Options = opts
	r.Labels = NewEmailInboxThreadLabelService(opts...)
	return
}

// Returns a bounded page of inbound and outbound thread messages interleaved in
// chronological order using stable cursor pagination.
func (r *EmailInboxThreadService) Get(ctx context.Context, threadID string, params EmailInboxThreadGetParams, opts ...option.RequestOption) (res *EmailInboxThreadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/threads/%s", params.InboxID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Lists thread summaries newest first using stable cursor pagination.
func (r *EmailInboxThreadService) List(ctx context.Context, inboxID string, query EmailInboxThreadListParams, opts ...option.RequestOption) (res *pagination.EmailBracketCursorPagination[InboundThread], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/threads", inboxID)
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

// Lists thread summaries newest first using stable cursor pagination.
func (r *EmailInboxThreadService) ListAutoPaging(ctx context.Context, inboxID string, query EmailInboxThreadListParams, opts ...option.RequestOption) *pagination.EmailBracketCursorPaginationAutoPager[InboundThread] {
	return pagination.NewEmailBracketCursorPaginationAutoPager(r.List(ctx, inboxID, query, opts...))
}

type EmailPaginationMeta struct {
	PageSize int64 `json:"page_size" api:"required"`
	// Cursor for the next page, when more results are available.
	PageCursor string `json:"page_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageSize    respjson.Field
		PageCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *EmailPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundEmailAddress struct {
	Email string `json:"email" api:"required" format:"email"`
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
func (r InboundEmailAddress) RawJSON() string { return r.JSON.raw }
func (r *InboundEmailAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundThread struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	InboxID   string    `json:"inbox_id" api:"required" format:"uuid"`
	// Mutable thread labels used for agent workflow state. Independent of the labels
	// on the thread's messages, and distinct from the send-time `tags` on outbound
	// messages.
	Labels        []string  `json:"labels" api:"required"`
	LastMessageAt time.Time `json:"last_message_at" api:"required" format:"date-time"`
	LastMessageID string    `json:"last_message_id" api:"required" format:"uuid"`
	// Total inbound and outbound messages in the thread.
	MessageCount int64  `json:"message_count" api:"required"`
	Preview      string `json:"preview" api:"required"`
	// Any of "email_thread".
	RecordType InboundThreadRecordType `json:"record_type" api:"required"`
	Subject    string                  `json:"subject" api:"required"`
	// Unread inbound messages; outbound messages never increment this count.
	UnreadCount int64     `json:"unread_count" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		InboxID       respjson.Field
		Labels        respjson.Field
		LastMessageAt respjson.Field
		LastMessageID respjson.Field
		MessageCount  respjson.Field
		Preview       respjson.Field
		RecordType    respjson.Field
		Subject       respjson.Field
		UnreadCount   respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundThread) RawJSON() string { return r.JSON.raw }
func (r *InboundThread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundThreadRecordType string

const (
	InboundThreadRecordTypeEmailThread InboundThreadRecordType = "email_thread"
)

type InboundThreadDetail struct {
	Messages []ThreadMessage `json:"messages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	InboundThread
}

// Returns the unmodified JSON received from the API
func (r InboundThreadDetail) RawJSON() string { return r.JSON.raw }
func (r *InboundThreadDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundThreadListResponse struct {
	Data []InboundThread     `json:"data" api:"required"`
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
func (r InboundThreadListResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundThreadListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThreadMessage struct {
	ID          string                `json:"id" api:"required" format:"uuid"`
	Attachments []map[string]any      `json:"attachments" api:"required"`
	Bcc         []InboundEmailAddress `json:"bcc" api:"required"`
	Cc          []InboundEmailAddress `json:"cc" api:"required"`
	CreatedAt   time.Time             `json:"created_at" api:"required" format:"date-time"`
	// Any of "inbound", "outbound".
	Direction ThreadMessageDirection `json:"direction" api:"required"`
	From      InboundEmailAddress    `json:"from" api:"required"`
	// Whether conservative plain-text extraction detected a quoted tail. False does
	// not prove that the source contains no quoted content.
	HasQuotedText bool           `json:"has_quoted_text" api:"required"`
	Headers       map[string]any `json:"headers" api:"required"`
	// URL for an offloaded HTML body. Null means the body is not offloaded to a URL;
	// an inline HTML body may still exist but is not returned on list reads. Reply
	// extraction uses only the plain-text body during ingest.
	HTMLBodyURL string           `json:"html_body_url" api:"required" format:"uri"`
	InReplyTo   string           `json:"in_reply_to" api:"required"`
	InboxID     string           `json:"inbox_id" api:"required" format:"uuid"`
	InlineFiles []map[string]any `json:"inline_files" api:"required"`
	// Mutable message labels used for agent workflow state (for example `spam`,
	// `needs_review`, `processed`). Distinct from the immutable send-time `tags` on
	// outbound messages: labels are never propagated to Email Detail Records or
	// Mission Control reporting. Always empty for outbound messages. Labels on a
	// message are independent of the labels on its thread.
	Labels []string `json:"labels" api:"required"`
	// RFC Message-ID header. Null is possible for legacy outbound messages.
	MessageID string `json:"message_id" api:"required"`
	// Time the inbound message was marked read. Null means unread.
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// Receipt time for inbound messages; null for outbound messages.
	ReceivedAt time.Time `json:"received_at" api:"required" format:"date-time"`
	// Any of "email_message".
	RecordType ThreadMessageRecordType `json:"record_type" api:"required"`
	// Ordered RFC Message-ID values from the References header.
	References []string `json:"references" api:"required"`
	// Conservatively extracted new-reply content persisted from the plain-text body
	// during ingest. Null means no plain-text extraction input was available or
	// extraction was skipped or failed; HTML bodies are not parsed.
	ReplyText string                `json:"reply_text" api:"required"`
	ReplyTo   []InboundEmailAddress `json:"reply_to" api:"required"`
	// Creation/send-acceptance time for outbound messages; null for inbound messages.
	SentAt time.Time `json:"sent_at" api:"required" format:"date-time"`
	// Received for inbound messages; the current send status for outbound messages.
	Status  string `json:"status" api:"required"`
	Subject string `json:"subject" api:"required"`
	// URL for an offloaded plain-text body. Null means the body is not offloaded to a
	// URL; an inline plain-text body may still exist but is not returned on list
	// reads. `reply_text` and `has_quoted_text` are persisted during ingest before any
	// body offload.
	TextBodyURL string                `json:"text_body_url" api:"required" format:"uri"`
	ThreadID    string                `json:"thread_id" api:"required" format:"uuid"`
	To          []InboundEmailAddress `json:"to" api:"required"`
	UpdatedAt   time.Time             `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Attachments   respjson.Field
		Bcc           respjson.Field
		Cc            respjson.Field
		CreatedAt     respjson.Field
		Direction     respjson.Field
		From          respjson.Field
		HasQuotedText respjson.Field
		Headers       respjson.Field
		HTMLBodyURL   respjson.Field
		InReplyTo     respjson.Field
		InboxID       respjson.Field
		InlineFiles   respjson.Field
		Labels        respjson.Field
		MessageID     respjson.Field
		ReadAt        respjson.Field
		ReceivedAt    respjson.Field
		RecordType    respjson.Field
		References    respjson.Field
		ReplyText     respjson.Field
		ReplyTo       respjson.Field
		SentAt        respjson.Field
		Status        respjson.Field
		Subject       respjson.Field
		TextBodyURL   respjson.Field
		ThreadID      respjson.Field
		To            respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThreadMessage) RawJSON() string { return r.JSON.raw }
func (r *ThreadMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThreadMessageDirection string

const (
	ThreadMessageDirectionInbound  ThreadMessageDirection = "inbound"
	ThreadMessageDirectionOutbound ThreadMessageDirection = "outbound"
)

type ThreadMessageRecordType string

const (
	ThreadMessageRecordTypeEmailMessage ThreadMessageRecordType = "email_message"
)

type EmailInboxThreadGetResponse struct {
	Data InboundThreadDetail `json:"data" api:"required"`
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
func (r EmailInboxThreadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxThreadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxThreadGetParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Opaque message cursor returned by the previous thread-detail page.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Number of thread messages to return. Defaults to 25; maximum is 100.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailInboxThreadGetParams]'s query parameters as
// `url.Values`.
func (r EmailInboxThreadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailInboxThreadListParams struct {
	// Returns only threads carrying this label. Thread labels are independent of the
	// labels on the thread's messages.
	FilterLabel param.Opt[string] `query:"filter[label],omitzero" json:"-"`
	// Opaque cursor returned by the previous page.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailInboxThreadListParams]'s query parameters as
// `url.Values`.
func (r EmailInboxThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
