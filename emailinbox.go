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

// Create and manage agent inboxes, retrieve inbound messages and threads, and
// reply to or forward messages.
//
// EmailInboxService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxService] method instead.
type EmailInboxService struct {
	Options []option.RequestOption
	// Create, list, retrieve, update, delete, and send unsent draft messages belonging
	// to an agent inbox.
	Drafts EmailInboxDraftService
	// Create and manage agent inboxes, retrieve inbound messages and threads, and
	// reply to or forward messages.
	Filters  EmailInboxFilterService
	Messages EmailInboxMessageService
	// Create and manage agent inboxes, retrieve inbound messages and threads, and
	// reply to or forward messages.
	Threads EmailInboxThreadService
}

// NewEmailInboxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailInboxService(opts ...option.RequestOption) (r EmailInboxService) {
	r = EmailInboxService{}
	r.Options = opts
	r.Drafts = NewEmailInboxDraftService(opts...)
	r.Filters = NewEmailInboxFilterService(opts...)
	r.Messages = NewEmailInboxMessageService(opts...)
	r.Threads = NewEmailInboxThreadService(opts...)
	return
}

// Creates an inbox on an inbound-enabled domain. When `domain_id` is omitted,
// Telnyx allocates the account's shared inbound subdomain so the inbox is
// immediately usable without customer DNS setup. When `username` is omitted, a
// unique username is generated.
func (r *EmailInboxService) New(ctx context.Context, body EmailInboxNewParams, opts ...option.RequestOption) (res *EmailInboxResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns an account-scoped, non-deleted inbox. Missing and foreign inboxes are
// indistinguishable.
func (r *EmailInboxService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailInboxResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the account's non-deleted inboxes newest first using stable cursor
// pagination.
func (r *EmailInboxService) List(ctx context.Context, query EmailInboxListParams, opts ...option.RequestOption) (res *EmailInboxListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Soft-deletes an account-scoped inbox. Its address remains reserved and the inbox
// is no longer returned by list or get operations.
func (r *EmailInboxService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("email_inboxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type EmailInbox struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	Address   string    `json:"address" api:"required" format:"email"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Domain name used by the inbox address.
	Domain   string `json:"domain" api:"required"`
	DomainID string `json:"domain_id" api:"required" format:"uuid"`
	// Any of "email_inbox".
	RecordType EmailInboxRecordType `json:"record_type" api:"required"`
	Settings   map[string]any       `json:"settings" api:"required"`
	// Any of "active", "paused".
	Status    EmailInboxStatus `json:"status" api:"required"`
	UpdatedAt time.Time        `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		CreatedAt   respjson.Field
		Domain      respjson.Field
		DomainID    respjson.Field
		RecordType  respjson.Field
		Settings    respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInbox) RawJSON() string { return r.JSON.raw }
func (r *EmailInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxRecordType string

const (
	EmailInboxRecordTypeEmailInbox EmailInboxRecordType = "email_inbox"
)

type EmailInboxStatus string

const (
	EmailInboxStatusActive EmailInboxStatus = "active"
	EmailInboxStatusPaused EmailInboxStatus = "paused"
)

type EmailInboxResponse struct {
	Data EmailInbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxListResponse struct {
	Data []EmailInbox               `json:"data" api:"required"`
	Meta EmailInboxListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxListResponseMeta struct {
	PageSize int64 `json:"page_size" api:"required"`
	// Cursor for the next inbox page, when more results are available.
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
func (r EmailInboxListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxNewParams struct {
	// Account-owned, inbound-enabled domain UUID. The account's shared inbound
	// subdomain is allocated when omitted.
	DomainID param.Opt[string] `json:"domain_id,omitzero" format:"uuid"`
	// Inbox local part. Trimmed and lowercased before validation; the normalized value
	// must be 1-64 characters, start and end with a letter or digit, and contain only
	// letters, digits, dots, hyphens, and underscores. Generated when omitted.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r EmailInboxNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailInboxNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailInboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxListParams struct {
	// Opaque cursor returned by the previous inbox page.
	PageCursor param.Opt[string] `query:"page_cursor,omitzero" json:"-"`
	// Number of results to return. Defaults to 20; maximum is 250.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailInboxListParams]'s query parameters as `url.Values`.
func (r EmailInboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
