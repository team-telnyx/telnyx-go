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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// PortingOrderCommentService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderCommentService] method instead.
type PortingOrderCommentService struct {
	Options []option.RequestOption
}

// NewPortingOrderCommentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPortingOrderCommentService(opts ...option.RequestOption) (r PortingOrderCommentService) {
	r = PortingOrderCommentService{}
	r.Options = opts
	return
}

// Creates a new comment for a porting order.
func (r *PortingOrderCommentService) New(ctx context.Context, id string, body PortingOrderCommentNewParams, opts ...option.RequestOption) (res *PortingOrderCommentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/comments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of all comments of a porting order.
func (r *PortingOrderCommentService) List(ctx context.Context, id string, query PortingOrderCommentListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PortingOrderCommentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/comments", id)
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

// Returns a list of all comments of a porting order.
func (r *PortingOrderCommentService) ListAutoPaging(ctx context.Context, id string, query PortingOrderCommentListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PortingOrderCommentListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, id, query, opts...))
}

type PortingOrderCommentNewResponse struct {
	Data PortingOrderCommentNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderCommentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderCommentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderCommentNewResponseData struct {
	ID string `json:"id" format:"uuid"`
	// Body of comment
	Body string `json:"body"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt      time.Time `json:"created_at" format:"date-time"`
	PortingOrderID string    `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates whether this comment was created by a Telnyx Admin, user, or system
	//
	// Any of "admin", "user", "system".
	UserType string `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Body           respjson.Field
		CreatedAt      respjson.Field
		PortingOrderID respjson.Field
		RecordType     respjson.Field
		UserType       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderCommentNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderCommentNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderCommentListResponse struct {
	ID string `json:"id" format:"uuid"`
	// Body of comment
	Body string `json:"body"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt      time.Time `json:"created_at" format:"date-time"`
	PortingOrderID string    `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates whether this comment was created by a Telnyx Admin, user, or system
	//
	// Any of "admin", "user", "system".
	UserType PortingOrderCommentListResponseUserType `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Body           respjson.Field
		CreatedAt      respjson.Field
		PortingOrderID respjson.Field
		RecordType     respjson.Field
		UserType       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderCommentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderCommentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether this comment was created by a Telnyx Admin, user, or system
type PortingOrderCommentListResponseUserType string

const (
	PortingOrderCommentListResponseUserTypeAdmin  PortingOrderCommentListResponseUserType = "admin"
	PortingOrderCommentListResponseUserTypeUser   PortingOrderCommentListResponseUserType = "user"
	PortingOrderCommentListResponseUserTypeSystem PortingOrderCommentListResponseUserType = "system"
)

type PortingOrderCommentNewParams struct {
	Body param.Opt[string] `json:"body,omitzero"`
	paramObj
}

func (r PortingOrderCommentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderCommentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderCommentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderCommentListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderCommentListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderCommentListParams]'s query parameters as
// `url.Values`.
func (r PortingOrderCommentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderCommentListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderCommentListParamsPage]'s query parameters as
// `url.Values`.
func (r PortingOrderCommentListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
