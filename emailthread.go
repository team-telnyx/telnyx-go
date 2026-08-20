// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Account-wide conversation threads across every inbox, for agents operating many
// inboxes at once.
//
// EmailThreadService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailThreadService] method instead.
type EmailThreadService struct {
	Options []option.RequestOption
}

// NewEmailThreadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailThreadService(opts ...option.RequestOption) (r EmailThreadService) {
	r = EmailThreadService{}
	r.Options = opts
	return
}

// Returns a thread and a bounded page of its inbound and outbound messages,
// interleaved in chronological order. The `inbox_id` returned by the list endpoint
// is required because a thread ID can occur in multiple inboxes. Only messages
// matching that `(inbox_id, thread_id)` pair are returned. Threads outside the
// account return an opaque 404.
func (r *EmailThreadService) Get(ctx context.Context, threadID string, query EmailThreadGetParams, opts ...option.RequestOption) (res *EmailThreadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists thread summaries for the whole account, newest first, using stable cursor
// pagination. An agent operating many inboxes gets every conversation in one call
// instead of one call per inbox. Each thread carries its own `inbox_id` so a reply
// can be routed back to the right inbox. Use `filter[inbox_id]` (repeatable) to
// narrow the result to specific inboxes. Because a thread ID can be delivered to
// multiple inboxes, each result is identified by its `(inbox_id, id)` pair.
func (r *EmailThreadService) List(ctx context.Context, query EmailThreadListParams, opts ...option.RequestOption) (res *pagination.EmailBracketCursorPagination[InboundThread], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "email_threads"
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

// Lists thread summaries for the whole account, newest first, using stable cursor
// pagination. An agent operating many inboxes gets every conversation in one call
// instead of one call per inbox. Each thread carries its own `inbox_id` so a reply
// can be routed back to the right inbox. Use `filter[inbox_id]` (repeatable) to
// narrow the result to specific inboxes. Because a thread ID can be delivered to
// multiple inboxes, each result is identified by its `(inbox_id, id)` pair.
func (r *EmailThreadService) ListAutoPaging(ctx context.Context, query EmailThreadListParams, opts ...option.RequestOption) *pagination.EmailBracketCursorPaginationAutoPager[InboundThread] {
	return pagination.NewEmailBracketCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

type EmailThreadGetResponse struct {
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
func (r EmailThreadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailThreadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailThreadGetParams struct {
	// Inbox UUID that, together with `thread_id`, identifies the thread.
	InboxID string `query:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Opaque message cursor returned by the previous thread-detail page.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Number of thread messages to return. Defaults to 25; maximum is 100.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailThreadGetParams]'s query parameters as `url.Values`.
func (r EmailThreadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailThreadListParams struct {
	// Returns only threads carrying this label. Matching is exact and case-sensitive.
	// Thread labels are independent of the labels on the thread's messages.
	FilterLabel param.Opt[string] `query:"filter[label],omitzero" json:"-"`
	// Opaque cursor returned by the previous page.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Restrict results to one or more inboxes. Repeat the parameter
	// (`filter[inbox_id][]=...&filter[inbox_id][]=...`) or pass a comma-separated
	// list. Omit to list every inbox in the account. Inboxes outside the account are
	// silently excluded. If the filter is present, it must contain at least one
	// non-empty UUID.
	FilterInboxID []string `query:"filter[inbox_id],omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [EmailThreadListParams]'s query parameters as `url.Values`.
func (r EmailThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
