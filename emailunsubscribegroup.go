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

// Named groups and group-scoped suppressions.
//
// EmailUnsubscribeGroupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailUnsubscribeGroupService] method instead.
type EmailUnsubscribeGroupService struct {
	Options []option.RequestOption
	// Named groups and group-scoped suppressions.
	Suppressions EmailUnsubscribeGroupSuppressionService
}

// NewEmailUnsubscribeGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailUnsubscribeGroupService(opts ...option.RequestOption) (r EmailUnsubscribeGroupService) {
	r = EmailUnsubscribeGroupService{}
	r.Options = opts
	r.Suppressions = NewEmailUnsubscribeGroupSuppressionService(opts...)
	return
}

// Creates an account-owned unsubscribe group for associating email categories with
// separate recipient suppression lists.
func (r *EmailUnsubscribeGroupService) New(ctx context.Context, body EmailUnsubscribeGroupNewParams, opts ...option.RequestOption) (res *UnsubscribeGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_unsubscribe_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the account-owned unsubscribe group identified by ID.
func (r *EmailUnsubscribeGroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UnsubscribeGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_unsubscribe_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partial update (only `name` / `description`). `PUT` is not routed.
func (r *EmailUnsubscribeGroupService) Update(ctx context.Context, id string, body EmailUnsubscribeGroupUpdateParams, opts ...option.RequestOption) (res *UnsubscribeGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_unsubscribe_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Offset pagination only (`page[number]` default 1, `page[size]` default 25, max
// 100). No `sort`/`filter`/cursor — ordering fixed `desc created_at, desc id`.
// Uses the shared `QueryParser.parse_offset/1` — a malformed `page` (e.g. flat
// `?page=1` instead of `?page[number]=1`) returns `400` (code `10015`), consistent
// with `GET /v2/email_blocks`. `meta` includes `total_pages`.
func (r *EmailUnsubscribeGroupService) List(ctx context.Context, query EmailUnsubscribeGroupListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[UnsubscribeGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "email_unsubscribe_groups"
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

// Offset pagination only (`page[number]` default 1, `page[size]` default 25, max
// 100). No `sort`/`filter`/cursor — ordering fixed `desc created_at, desc id`.
// Uses the shared `QueryParser.parse_offset/1` — a malformed `page` (e.g. flat
// `?page=1` instead of `?page[number]=1`) returns `400` (code `10015`), consistent
// with `GET /v2/email_blocks`. `meta` includes `total_pages`.
func (r *EmailUnsubscribeGroupService) ListAutoPaging(ctx context.Context, query EmailUnsubscribeGroupListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[UnsubscribeGroup] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// If the group has 0 active suppressions, hard-deletes the row. With `force=true`,
// soft-deletes all active suppressions first (status → `removed`, `group_id`
// cleared, `removed` audit event per block) in a single transaction, then
// hard-deletes the group. Without `force` and active suppressions present → `409`.
// Audit trail is preserved. `force` only accepts the string `"true"` or boolean
// `true`; all other values are false.
func (r *EmailUnsubscribeGroupService) Delete(ctx context.Context, id string, body EmailUnsubscribeGroupDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("email_unsubscribe_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Group list `meta` (consistent with `GET /v2/email_blocks`).
type GroupListMeta struct {
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
func (r GroupListMeta) RawJSON() string { return r.JSON.raw }
func (r *GroupListMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnsubscribeGroup struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Always present (not omit-nullable); `null` when unset.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	// View-only.
	//
	// Any of "email_unsubscribe_group".
	RecordType UnsubscribeGroupRecordType `json:"record_type" api:"required"`
	UpdatedAt  time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnsubscribeGroup) RawJSON() string { return r.JSON.raw }
func (r *UnsubscribeGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// View-only.
type UnsubscribeGroupRecordType string

const (
	UnsubscribeGroupRecordTypeEmailUnsubscribeGroup UnsubscribeGroupRecordType = "email_unsubscribe_group"
)

type UnsubscribeGroupResponse struct {
	Data UnsubscribeGroup `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnsubscribeGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *UnsubscribeGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailUnsubscribeGroupNewParams struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r EmailUnsubscribeGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailUnsubscribeGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailUnsubscribeGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailUnsubscribeGroupUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r EmailUnsubscribeGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailUnsubscribeGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailUnsubscribeGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailUnsubscribeGroupListParams struct {
	// Offset page number (≥1, default 1).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Page size (1–100, default 25).
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailUnsubscribeGroupListParams]'s query parameters as
// `url.Values`.
func (r EmailUnsubscribeGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailUnsubscribeGroupDeleteParams struct {
	// Force-delete a group with active suppressions. Only `"true"` (string) or `true`
	// (bool) are truthy; all other values are false.
	Force EmailUnsubscribeGroupDeleteParamsForceUnion `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailUnsubscribeGroupDeleteParams]'s query parameters as
// `url.Values`.
func (r EmailUnsubscribeGroupDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmailUnsubscribeGroupDeleteParamsForceUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfEmailUnsubscribeGroupDeletesForceString)
	OfEmailUnsubscribeGroupDeletesForceString param.Opt[string] `query:",omitzero,inline"`
	OfBool                                    param.Opt[bool]   `query:",omitzero,inline"`
	paramUnion
}

func (u *EmailUnsubscribeGroupDeleteParamsForceUnion) asAny() any {
	if !param.IsOmitted(u.OfEmailUnsubscribeGroupDeletesForceString) {
		return &u.OfEmailUnsubscribeGroupDeletesForceString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type EmailUnsubscribeGroupDeleteParamsForceString string

const (
	EmailUnsubscribeGroupDeleteParamsForceStringTrue  EmailUnsubscribeGroupDeleteParamsForceString = "true"
	EmailUnsubscribeGroupDeleteParamsForceStringFalse EmailUnsubscribeGroupDeleteParamsForceString = "false"
)
