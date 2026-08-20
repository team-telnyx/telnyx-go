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
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// EmailInboxMessageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxMessageService] method instead.
type EmailInboxMessageService struct {
	Options []option.RequestOption
	// Create and manage agent inboxes, retrieve inbound messages and threads, and
	// reply to or forward messages.
	Actions EmailInboxMessageActionService
	// Create and manage agent inboxes, retrieve inbound messages and threads, and
	// reply to or forward messages.
	Labels EmailInboxMessageLabelService
}

// NewEmailInboxMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailInboxMessageService(opts ...option.RequestOption) (r EmailInboxMessageService) {
	r = EmailInboxMessageService{}
	r.Options = opts
	r.Actions = NewEmailInboxMessageActionService(opts...)
	r.Labels = NewEmailInboxMessageLabelService(opts...)
	return
}

// Updates the explicit read state of an account-scoped inbound message. Set
// `read_at` to `true` to mark the message read at the server's current time, to an
// ISO 8601 timestamp to use that timestamp, or to `null` to mark the message
// unread. Repeating the same update is idempotent.
func (r *EmailInboxMessageService) Update(ctx context.Context, messageID string, params EmailInboxMessageUpdateParams, opts ...option.RequestOption) (res *EmailInboxMessageUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists inbound messages newest first. All access is scoped to the authenticated
// account. `filter[search]` performs PostgreSQL full-text search over the subject,
// plain-text body, and HTML body. Filters compose with stable cursor pagination.
func (r *EmailInboxMessageService) List(ctx context.Context, inboxID string, query EmailInboxMessageListParams, opts ...option.RequestOption) (res *pagination.EmailBracketCursorPagination[InboundMessage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages", inboxID)
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

// Lists inbound messages newest first. All access is scoped to the authenticated
// account. `filter[search]` performs PostgreSQL full-text search over the subject,
// plain-text body, and HTML body. Filters compose with stable cursor pagination.
func (r *EmailInboxMessageService) ListAutoPaging(ctx context.Context, inboxID string, query EmailInboxMessageListParams, opts ...option.RequestOption) *pagination.EmailBracketCursorPaginationAutoPager[InboundMessage] {
	return pagination.NewEmailBracketCursorPaginationAutoPager(r.List(ctx, inboxID, query, opts...))
}

// Creates an unsent reply draft for an inbound message. Unlike the
// `/actions/reply` endpoint, which sends immediately, this stores a draft that can
// be reviewed and edited before sending.
//
// `reply_to_message_id` and `thread_id` are inherited from the parent message and
// cannot be set by the caller. The recipient, `Re:` subject and
// `In-Reply-To`/`References` headers are pre-filled from the parent using the same
// rules as a live reply, so sending the draft threads identically. Supplying `to`
// or `subject` explicitly overrides the pre-filled value.
func (r *EmailInboxMessageService) Drafts(ctx context.Context, messageID string, params EmailInboxMessageDraftsParams, opts ...option.RequestOption) (res *EmailDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s/drafts", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type EmailInboxMessageUpdateResponse struct {
	Data InboundMessage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxMessageUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxMessageUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageUpdateParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Set to `true` for server time, an ISO 8601 timestamp for an explicit read time,
	// or `null` to mark unread.
	ReadAt EmailInboxMessageUpdateParamsReadAtUnion `json:"read_at,omitzero" api:"required" format:"date-time"`
	paramObj
}

func (r EmailInboxMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailInboxMessageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailInboxMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmailInboxMessageUpdateParamsReadAtUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfEmailInboxMessageUpdatesReadAtBoolean)
	OfEmailInboxMessageUpdatesReadAtBoolean param.Opt[bool]      `json:",omitzero,inline"`
	OfTime                                  param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u EmailInboxMessageUpdateParamsReadAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmailInboxMessageUpdatesReadAtBoolean, u.OfTime)
}
func (u *EmailInboxMessageUpdateParamsReadAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmailInboxMessageUpdateParamsReadAtUnion) asAny() any {
	if !param.IsOmitted(u.OfEmailInboxMessageUpdatesReadAtBoolean) {
		return &u.OfEmailInboxMessageUpdatesReadAtBoolean
	} else if !param.IsOmitted(u.OfTime) {
		return &u.OfTime.Value
	}
	return nil
}

type EmailInboxMessageUpdateParamsReadAtBoolean bool

const (
	EmailInboxMessageUpdateParamsReadAtBooleanTrue EmailInboxMessageUpdateParamsReadAtBoolean = true
)

type EmailInboxMessageListParams struct {
	// Case-insensitive literal substring of the sender address.
	FilterFrom param.Opt[string] `query:"filter[from],omitzero" json:"-"`
	// Returns only messages carrying this label. Matching is exact and case-sensitive.
	// Reserved `telnyx:` labels can be filtered on even though they cannot be written
	// by customers.
	FilterLabel param.Opt[string] `query:"filter[label],omitzero" json:"-"`
	// Whether the message has a read timestamp.
	FilterRead param.Opt[bool] `query:"filter[read],omitzero" json:"-"`
	// Inclusive ISO 8601 lower bound for the received timestamp.
	FilterReceivedAfter param.Opt[time.Time] `query:"filter[received_after],omitzero" format:"date-time" json:"-"`
	// Inclusive ISO 8601 upper bound for the received timestamp.
	FilterReceivedBefore param.Opt[time.Time] `query:"filter[received_before],omitzero" format:"date-time" json:"-"`
	// Full-text query over subject and body, up to 500 characters.
	FilterSearch param.Opt[string] `query:"filter[search],omitzero" json:"-"`
	// Case-insensitive literal substring of the subject.
	FilterSubject param.Opt[string] `query:"filter[subject],omitzero" json:"-"`
	// Whether the message has no read timestamp. Set to `true` to return only unread
	// messages.
	FilterUnread param.Opt[bool] `query:"filter[unread],omitzero" json:"-"`
	// Opaque cursor returned by the previous page.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailInboxMessageListParams]'s query parameters as
// `url.Values`.
func (r EmailInboxMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailInboxMessageDraftsParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// All fields are optional — a draft may be saved incomplete. `account_id`,
	// `inbox_id`, `status`, `sent_at`, `sent_message_id`, `reply_to_message_id` and
	// `thread_id` are server-owned and ignored if supplied.
	EmailDraftRequest EmailDraftRequestParam
	paramObj
}

func (r EmailInboxMessageDraftsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailDraftRequest)
}
func (r *EmailInboxMessageDraftsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
