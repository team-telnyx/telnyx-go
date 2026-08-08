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

// Recipient suppression records (`/v2/email_blocks`).
//
// EmailBlockService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailBlockService] method instead.
type EmailBlockService struct {
	Options []option.RequestOption
	// Async CSV import of competitor suppression lists.
	Import EmailBlockImportService
}

// NewEmailBlockService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailBlockService(opts ...option.RequestOption) (r EmailBlockService) {
	r = EmailBlockService{}
	r.Options = opts
	r.Import = NewEmailBlockImportService(opts...)
	return
}

// Creates a suppression with `reason: manual_block` and `source: manual`.
// Caller-supplied `reason` / `source` are **ignored**; `scope` is **derived**
// server-side from `domain_id` / `from` and is never trusted. Idempotent: if a
// matching row already exists (NULL-safe dedupe key: account_id, scope, to,
// reason, domain_id, from), returns the existing record with `200` (no new audit
// event).
//
// `bounce_category`, `dsn_code`, `meta`, and `group_id` are **not accepted** on
// the public surface. Use the unsubscribe-group suppression endpoint or the
// internal create surface for those.
func (r *EmailBlockService) New(ctx context.Context, body EmailBlockNewParams, opts ...option.RequestOption) (res *EmailBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_blocks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the account-owned suppression identified by ID. Cross-account lookups
// and malformed IDs return `404` without exposing another account’s data.
func (r *EmailBlockService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_blocks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Account-scoped list. Two mutually exclusive pagination modes:
//
//   - **Offset**: `page[number]` (default 1) + `page[size]` (default 25, max 100).
//     `meta` contains `total_pages`.
//   - **Cursor**: `page[after]` and/or `page[before]` (opaque `Base.url_encode64` of
//     `{"created_at","id"}`). Cannot combine with `page[number]`; `after`+`before`
//     together is an error. `meta` contains `next_cursor` / `previous_cursor`
//     (omitted when their flag is false).
//
// Sort defaults to `-created_at` (desc); only `created_at` is sortable. A `--`
// prefix is an error. `nil`/empty filter values are silently dropped.
func (r *EmailBlockService) List(ctx context.Context, query EmailBlockListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EmailBlock], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "email_blocks"
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

// Account-scoped list. Two mutually exclusive pagination modes:
//
//   - **Offset**: `page[number]` (default 1) + `page[size]` (default 25, max 100).
//     `meta` contains `total_pages`.
//   - **Cursor**: `page[after]` and/or `page[before]` (opaque `Base.url_encode64` of
//     `{"created_at","id"}`). Cannot combine with `page[number]`; `after`+`before`
//     together is an error. `meta` contains `next_cursor` / `previous_cursor`
//     (omitted when their flag is false).
//
// Sort defaults to `-created_at` (desc); only `created_at` is sortable. A `--`
// prefix is an error. `nil`/empty filter values are silently dropped.
func (r *EmailBlockService) ListAutoPaging(ctx context.Context, query EmailBlockListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EmailBlock] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Soft-deletes (status → `removed`; tombstone retained). A `removed` audit event
// is appended unless the block was already `removed` (idempotent — returns the
// existing row with `200` and no new event). Mutates `updated_at`.
func (r *EmailBlockService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_blocks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Offset pagination only (`page[number]` default 1, `page[size]` default **50**,
// max 100). No `sort`, no `filter`, no cursor — ordering is fixed
// `desc occurred_at, desc id`. Verifies the block belongs to the account first
// (cross-account → 404).
func (r *EmailBlockService) GetEvents(ctx context.Context, id string, query EmailBlockGetEventsParams, opts ...option.RequestOption) (res *EmailBlockGetEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_blocks/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Streams the account's suppressions as a chunked CSV (server-side cursor; never
// materialized). Content-type `text/csv`, header
// `Content-Disposition: attachment; filename="email_blocks_export.csv"`.
//
// Filters (`filter[reason]`, `filter[domain_id]`, `filter[created_after]`,
// `filter[created_before]`) are the only params that affect output. `sort` and
// `page[*]` are **parsed** (bad values still produce `400`) but **ignored** — rows
// stream `ORDER BY created_at ASC, id ASC` with no pagination.
//
// CSV columns:
// `id,to,from,reason,source,scope,status,domain_id, created_at,updated_at,expires_at,group_id`.
// The CSV carries the `group_id` column so group-scoped suppressions' group link
// survives the export (empty for account-scope rows).
func (r *EmailBlockService) GetExport(ctx context.Context, query EmailBlockGetExportParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := "email_blocks/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Suppression record. Schema fields hidden by the view: `account_id`,
// `bounce_category`, `dsn_code`, `meta`.
type EmailBlock struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "hard_bounce", "spam_complaint", "unsubscribe", "invalid",
	// "manual_block".
	Reason EmailBlockReason `json:"reason" api:"required"`
	// View-only discriminator.
	//
	// Any of "email_block".
	RecordType EmailBlockRecordType `json:"record_type" api:"required"`
	// Derived server-side from `domain_id`/`from`; never trusted from the caller.
	//
	// Any of "account", "domain", "address".
	Scope EmailBlockScope `json:"scope" api:"required"`
	// Any of "feedback", "manual", "import", "system".
	Source EmailBlockSource `json:"source" api:"required"`
	// Any of "active", "expired", "removed".
	Status EmailBlockStatus `json:"status" api:"required"`
	// Normalized recipient. (schema: to_address)
	To        string    `json:"to" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// `null` ⇒ account scope. Stored on the row; exposed here.
	DomainID  string    `json:"domain_id" api:"nullable" format:"uuid"`
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// `null` ⇒ not address-scope. (schema: from_address)
	From string `json:"from" api:"nullable"`
	// `null` ⇒ global; set ⇒ group-scoped opt-out.
	GroupID string `json:"group_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Reason      respjson.Field
		RecordType  respjson.Field
		Scope       respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		To          respjson.Field
		UpdatedAt   respjson.Field
		DomainID    respjson.Field
		ExpiresAt   respjson.Field
		From        respjson.Field
		GroupID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailBlockReason string

const (
	EmailBlockReasonHardBounce    EmailBlockReason = "hard_bounce"
	EmailBlockReasonSpamComplaint EmailBlockReason = "spam_complaint"
	EmailBlockReasonUnsubscribe   EmailBlockReason = "unsubscribe"
	EmailBlockReasonInvalid       EmailBlockReason = "invalid"
	EmailBlockReasonManualBlock   EmailBlockReason = "manual_block"
)

// View-only discriminator.
type EmailBlockRecordType string

const (
	EmailBlockRecordTypeEmailBlock EmailBlockRecordType = "email_block"
)

// Derived server-side from `domain_id`/`from`; never trusted from the caller.
type EmailBlockScope string

const (
	EmailBlockScopeAccount EmailBlockScope = "account"
	EmailBlockScopeDomain  EmailBlockScope = "domain"
	EmailBlockScopeAddress EmailBlockScope = "address"
)

type EmailBlockSource string

const (
	EmailBlockSourceFeedback EmailBlockSource = "feedback"
	EmailBlockSourceManual   EmailBlockSource = "manual"
	EmailBlockSourceImport   EmailBlockSource = "import"
	EmailBlockSourceSystem   EmailBlockSource = "system"
)

type EmailBlockStatus string

const (
	EmailBlockStatusActive  EmailBlockStatus = "active"
	EmailBlockStatusExpired EmailBlockStatus = "expired"
	EmailBlockStatusRemoved EmailBlockStatus = "removed"
)

type EmailBlockResponse struct {
	// Suppression record. Schema fields hidden by the view: `account_id`,
	// `bounce_category`, `dsn_code`, `meta`.
	Data EmailBlock `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailBlockResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailBlockResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OffsetMeta struct {
	PageNumber   int64 `json:"page_number" api:"required"`
	PageSize     int64 `json:"page_size" api:"required"`
	TotalPages   int64 `json:"total_pages" api:"required"`
	TotalResults int64 `json:"total_results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *OffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailBlockGetEventsResponse struct {
	Data []EmailBlockGetEventsResponseData `json:"data" api:"required"`
	Meta OffsetMeta                        `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailBlockGetEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailBlockGetEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailBlockGetEventsResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Free-text (`user_id`/`org_id`/`api_key`/`dev_bypass`/`system`/`manual`).
	Actor string `json:"actor" api:"required"`
	// Any of "created", "removed", "expired", "override_used".
	EventType  string    `json:"event_type" api:"required"`
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// Free-text snapshot of the block's reason at event time.
	Reason string `json:"reason" api:"required"`
	// View-only.
	//
	// Any of "email_block_event".
	RecordType string `json:"record_type" api:"required"`
	// Free-text snapshot of the block's source at event time.
	Source string `json:"source" api:"required"`
	// `null` when the schema field is nil (the context usually sets it to `{}`).
	Meta map[string]any `json:"meta" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actor       respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Reason      respjson.Field
		RecordType  respjson.Field
		Source      respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailBlockGetEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailBlockGetEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailBlockNewParams struct {
	// Recipient address (normalized: trim + lower-case).
	To string `json:"to" api:"required"`
	// `null` ⇒ account scope.
	DomainID  param.Opt[string]    `json:"domain_id,omitzero" format:"uuid"`
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Sender address (normalized). `null` ⇒ account/domain scope.
	From param.Opt[string] `json:"from,omitzero"`
	paramObj
}

func (r EmailBlockNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailBlockNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailBlockNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailBlockListParams struct {
	// `created_at > value` (ISO 8601).
	FilterCreatedAfter param.Opt[time.Time] `query:"filter[created_after],omitzero" format:"date-time" json:"-"`
	// `created_at < value` (ISO 8601).
	FilterCreatedBefore param.Opt[time.Time] `query:"filter[created_before],omitzero" format:"date-time" json:"-"`
	// Exact-match filter on domain_id (UUID).
	FilterDomainID param.Opt[string] `query:"filter[domain_id],omitzero" format:"uuid" json:"-"`
	// Opaque cursor (`Base.url_encode64` of `{"created_at","id"}`). Cursor mode;
	// mutually exclusive with `page[number]` and `page[before]`.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Opaque cursor (see `page[after]`). Mutually exclusive with `page[after]` and
	// `page[number]`.
	PageBefore param.Opt[string] `query:"page[before],omitzero" json:"-"`
	// Offset page number (≥1, default 1).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Page size (1–100, default 25).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Exact-match filter on reason.
	//
	// Any of "hard_bounce", "spam_complaint", "unsubscribe", "invalid",
	// "manual_block".
	FilterReason EmailBlockListParamsFilterReason `query:"filter[reason],omitzero" json:"-"`
	// Sort field. Leading `-` = desc; only `created_at` is sortable. Default
	// `-created_at`. `--` is an error.
	//
	// Any of "created_at", "-created_at".
	Sort EmailBlockListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailBlockListParams]'s query parameters as `url.Values`.
func (r EmailBlockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact-match filter on reason.
type EmailBlockListParamsFilterReason string

const (
	EmailBlockListParamsFilterReasonHardBounce    EmailBlockListParamsFilterReason = "hard_bounce"
	EmailBlockListParamsFilterReasonSpamComplaint EmailBlockListParamsFilterReason = "spam_complaint"
	EmailBlockListParamsFilterReasonUnsubscribe   EmailBlockListParamsFilterReason = "unsubscribe"
	EmailBlockListParamsFilterReasonInvalid       EmailBlockListParamsFilterReason = "invalid"
	EmailBlockListParamsFilterReasonManualBlock   EmailBlockListParamsFilterReason = "manual_block"
)

// Sort field. Leading `-` = desc; only `created_at` is sortable. Default
// `-created_at`. `--` is an error.
type EmailBlockListParamsSort string

const (
	EmailBlockListParamsSortCreatedAt     EmailBlockListParamsSort = "created_at"
	EmailBlockListParamsSortCreatedAtDesc EmailBlockListParamsSort = "-created_at"
)

type EmailBlockGetEventsParams struct {
	// Offset page number (≥1, default 1).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Page size (default 50, max 100).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailBlockGetEventsParams]'s query parameters as
// `url.Values`.
func (r EmailBlockGetEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailBlockGetExportParams struct {
	// `created_at > value` (ISO 8601).
	FilterCreatedAfter param.Opt[time.Time] `query:"filter[created_after],omitzero" format:"date-time" json:"-"`
	// `created_at < value` (ISO 8601).
	FilterCreatedBefore param.Opt[time.Time] `query:"filter[created_before],omitzero" format:"date-time" json:"-"`
	// Exact-match filter on domain_id (UUID).
	FilterDomainID param.Opt[string] `query:"filter[domain_id],omitzero" format:"uuid" json:"-"`
	// Offset page number (≥1, default 1).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Page size (1–100, default 25).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Exact-match filter on reason.
	//
	// Any of "hard_bounce", "spam_complaint", "unsubscribe", "invalid",
	// "manual_block".
	FilterReason EmailBlockGetExportParamsFilterReason `query:"filter[reason],omitzero" json:"-"`
	// Sort field. Leading `-` = desc; only `created_at` is sortable. Default
	// `-created_at`. `--` is an error.
	//
	// Any of "created_at", "-created_at".
	Sort EmailBlockGetExportParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailBlockGetExportParams]'s query parameters as
// `url.Values`.
func (r EmailBlockGetExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exact-match filter on reason.
type EmailBlockGetExportParamsFilterReason string

const (
	EmailBlockGetExportParamsFilterReasonHardBounce    EmailBlockGetExportParamsFilterReason = "hard_bounce"
	EmailBlockGetExportParamsFilterReasonSpamComplaint EmailBlockGetExportParamsFilterReason = "spam_complaint"
	EmailBlockGetExportParamsFilterReasonUnsubscribe   EmailBlockGetExportParamsFilterReason = "unsubscribe"
	EmailBlockGetExportParamsFilterReasonInvalid       EmailBlockGetExportParamsFilterReason = "invalid"
	EmailBlockGetExportParamsFilterReasonManualBlock   EmailBlockGetExportParamsFilterReason = "manual_block"
)

// Sort field. Leading `-` = desc; only `created_at` is sortable. Default
// `-created_at`. `--` is an error.
type EmailBlockGetExportParamsSort string

const (
	EmailBlockGetExportParamsSortCreatedAt     EmailBlockGetExportParamsSort = "created_at"
	EmailBlockGetExportParamsSortCreatedAtDesc EmailBlockGetExportParamsSort = "-created_at"
)
