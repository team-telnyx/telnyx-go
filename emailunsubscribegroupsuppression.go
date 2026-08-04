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
)

// Named groups and group-scoped suppressions.
//
// EmailUnsubscribeGroupSuppressionService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailUnsubscribeGroupSuppressionService] method instead.
type EmailUnsubscribeGroupSuppressionService struct {
	Options []option.RequestOption
}

// NewEmailUnsubscribeGroupSuppressionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailUnsubscribeGroupSuppressionService(opts ...option.RequestOption) (r EmailUnsubscribeGroupSuppressionService) {
	r = EmailUnsubscribeGroupSuppressionService{}
	r.Options = opts
	return
}

// Creates a suppression with `reason: unsubscribe`, `source: manual`,
// `group_id: <this group>`. All other body fields are ignored; only `to` is read.
// Idempotent (same dedupe key → `200`, no new event).
func (r *EmailUnsubscribeGroupSuppressionService) New(ctx context.Context, id string, body EmailUnsubscribeGroupSuppressionNewParams, opts ...option.RequestOption) (res *EmailBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_unsubscribe_groups/%s/suppressions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Account + group scoped. Offset pagination only (`page[number]` default 1,
// `page[size]` default 25, max 100). No `sort`/`filter`/ cursor — ordering fixed
// `desc created_at, desc id`. Uses the shared `QueryParser.parse_offset/1` — a
// malformed `page` returns `400` (code `10015`), consistent with
// `GET /v2/email_blocks`. `meta` includes `total_pages`. Rows reuse the standard
// suppression shape (`group_id` set to this group).
func (r *EmailUnsubscribeGroupSuppressionService) List(ctx context.Context, id string, query EmailUnsubscribeGroupSuppressionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EmailBlock], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_unsubscribe_groups/%s/suppressions", id)
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

// Account + group scoped. Offset pagination only (`page[number]` default 1,
// `page[size]` default 25, max 100). No `sort`/`filter`/ cursor — ordering fixed
// `desc created_at, desc id`. Uses the shared `QueryParser.parse_offset/1` — a
// malformed `page` returns `400` (code `10015`), consistent with
// `GET /v2/email_blocks`. `meta` includes `total_pages`. Rows reuse the standard
// suppression shape (`group_id` set to this group).
func (r *EmailUnsubscribeGroupSuppressionService) ListAutoPaging(ctx context.Context, id string, query EmailUnsubscribeGroupSuppressionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EmailBlock] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, id, query, opts...))
}

// Soft-deletes all active blocks for (account, group, normalized email) — one
// `removed` audit event per block (`actor: manual`). The `email` path segment is
// normalized (trim + lower-case) before matching. Idempotent on already-removed
// rows (returns `404` since they're no longer `active`).
//
// Two distinct `404` cases: a missing/cross-account **group** returns
// `10001 "The requested unsubscribe group was not found"`; a group that exists but
// has **no active suppression** for that email returns
// `10001 "The requested group suppression was not found"`.
func (r *EmailUnsubscribeGroupSuppressionService) Delete(ctx context.Context, email string, body EmailUnsubscribeGroupSuppressionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if email == "" {
		err = errors.New("missing required email parameter")
		return err
	}
	path := fmt.Sprintf("email_unsubscribe_groups/%s/suppressions/%s", body.ID, email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type EmailUnsubscribeGroupSuppressionNewParams struct {
	To string `json:"to" api:"required"`
	paramObj
}

func (r EmailUnsubscribeGroupSuppressionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailUnsubscribeGroupSuppressionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailUnsubscribeGroupSuppressionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailUnsubscribeGroupSuppressionListParams struct {
	// Offset page number (≥1, default 1).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Page size (1–100, default 25).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailUnsubscribeGroupSuppressionListParams]'s query
// parameters as `url.Values`.
func (r EmailUnsubscribeGroupSuppressionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailUnsubscribeGroupSuppressionDeleteParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	paramObj
}
