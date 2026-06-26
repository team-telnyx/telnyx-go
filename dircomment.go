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

// Read messages from the Telnyx vetting team and reply with clarifying
// information.
//
// DirCommentService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirCommentService] method instead.
type DirCommentService struct {
	Options []option.RequestOption
}

// NewDirCommentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirCommentService(opts ...option.RequestOption) (r DirCommentService) {
	r = DirCommentService{}
	r.Options = opts
	return
}

// Post a customer comment on a DIR (for example, to respond to reviewer notes).
// Send only `content` (1–5000 chars) and an optional `parent_comment_id`; the
// server sets the comment type, visibility, and author automatically. The
// enterprise is resolved server-side from the DIR id.
func (r *DirCommentService) New(ctx context.Context, dirID string, body DirCommentNewParams, opts ...option.RequestOption) (res *DirCommentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/comments", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the comments on a DIR. The enterprise is resolved server-side from the DIR
// id.
func (r *DirCommentService) List(ctx context.Context, dirID string, query DirCommentListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DirComment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/comments", dirID)
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

// List the comments on a DIR. The enterprise is resolved server-side from the DIR
// id.
func (r *DirCommentService) ListAutoPaging(ctx context.Context, dirID string, query DirCommentListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DirComment] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, dirID, query, opts...))
}

// Comment categorisation. Customers post `customer_inquiry`. The Telnyx team posts
// `vetting_comment`, `rejection_reason`, `notification`, `status_update`, or
// `admin_response`. `internal_note` is filtered out of customer-visible responses.
type CommentType string

const (
	CommentTypeVettingComment  CommentType = "vetting_comment"
	CommentTypeRejectionReason CommentType = "rejection_reason"
	CommentTypeInternalNote    CommentType = "internal_note"
	CommentTypeNotification    CommentType = "notification"
	CommentTypeStatusUpdate    CommentType = "status_update"
	CommentTypeCustomerInquiry CommentType = "customer_inquiry"
	CommentTypeAdminResponse   CommentType = "admin_response"
)

type DirComment struct {
	ID string `json:"id" format:"uuid"`
	// Display name of the author. May be `null`.
	AuthorName string `json:"author_name" api:"nullable"`
	// Who wrote the comment. `admin` covers the Telnyx vetting team.
	//
	// Any of "customer", "admin".
	AuthorRole DirCommentAuthorRole `json:"author_role"`
	// Comment categorisation. Customers post `customer_inquiry`. The Telnyx team posts
	// `vetting_comment`, `rejection_reason`, `notification`, `status_update`, or
	// `admin_response`. `internal_note` is filtered out of customer-visible responses.
	//
	// Any of "vetting_comment", "rejection_reason", "internal_note", "notification",
	// "status_update", "customer_inquiry", "admin_response".
	CommentType CommentType `json:"comment_type"`
	Content     string      `json:"content"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	// Resource the comment is attached to. Always `dir` on this endpoint.
	//
	// Any of "dir".
	EntityType DirCommentEntityType `json:"entity_type"`
	// Always `customer` on this endpoint - internal-only comments are filtered out.
	//
	// Any of "customer".
	Visibility DirCommentVisibility `json:"visibility"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AuthorName  respjson.Field
		AuthorRole  respjson.Field
		CommentType respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		EntityType  respjson.Field
		Visibility  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirComment) RawJSON() string { return r.JSON.raw }
func (r *DirComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who wrote the comment. `admin` covers the Telnyx vetting team.
type DirCommentAuthorRole string

const (
	DirCommentAuthorRoleCustomer DirCommentAuthorRole = "customer"
	DirCommentAuthorRoleAdmin    DirCommentAuthorRole = "admin"
)

// Resource the comment is attached to. Always `dir` on this endpoint.
type DirCommentEntityType string

const (
	DirCommentEntityTypeDir DirCommentEntityType = "dir"
)

// Always `customer` on this endpoint - internal-only comments are filtered out.
type DirCommentVisibility string

const (
	DirCommentVisibilityCustomer DirCommentVisibility = "customer"
)

type DirCommentNewResponse struct {
	Data DirComment `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirCommentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DirCommentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirCommentNewParams struct {
	// Comment body. 1–5000 characters.
	Content string `json:"content" api:"required"`
	// Optional parent comment id to thread this reply under.
	ParentCommentID param.Opt[string] `json:"parent_comment_id,omitzero" format:"uuid"`
	paramObj
}

func (r DirCommentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DirCommentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirCommentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirCommentListParams struct {
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Restrict to comments of this category. Customer-visible categories only:
	// internal-only comments are filtered out regardless of this filter.
	//
	// Any of "vetting_comment", "rejection_reason", "internal_note", "notification",
	// "status_update", "customer_inquiry", "admin_response".
	CommentType CommentType `query:"comment_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirCommentListParams]'s query parameters as `url.Values`.
func (r DirCommentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
