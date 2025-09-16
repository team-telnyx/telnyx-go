// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// CommentService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommentService] method instead.
type CommentService struct {
	Options []option.RequestOption
}

// NewCommentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommentService(opts ...option.RequestOption) (r CommentService) {
	r = CommentService{}
	r.Options = opts
	return
}

// Create a comment
func (r *CommentService) New(ctx context.Context, body CommentNewParams, opts ...option.RequestOption) (res *CommentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "comments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a comment
func (r *CommentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CommentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("comments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all comments
func (r *CommentService) List(ctx context.Context, query CommentListParams, opts ...option.RequestOption) (res *CommentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "comments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Mark a comment as read
func (r *CommentService) MarkAsRead(ctx context.Context, id string, opts ...option.RequestOption) (res *CommentMarkAsReadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("comments/%s/read", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type CommentNewResponse struct {
	Data CommentNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CommentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentNewResponseData struct {
	ID              string `json:"id" format:"uuid"`
	Body            string `json:"body"`
	CommentRecordID string `json:"comment_record_id" format:"uuid"`
	// Any of "sub_number_order", "requirement_group".
	CommentRecordType string `json:"comment_record_type"`
	Commenter         string `json:"commenter"`
	// Any of "admin", "user".
	CommenterType string `json:"commenter_type"`
	// An ISO 8901 datetime string denoting when the comment was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was read.
	ReadAt time.Time `json:"read_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Body              respjson.Field
		CommentRecordID   respjson.Field
		CommentRecordType respjson.Field
		Commenter         respjson.Field
		CommenterType     respjson.Field
		CreatedAt         respjson.Field
		ReadAt            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *CommentNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentGetResponse struct {
	Data CommentGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CommentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentGetResponseData struct {
	ID              string `json:"id" format:"uuid"`
	Body            string `json:"body"`
	CommentRecordID string `json:"comment_record_id" format:"uuid"`
	// Any of "sub_number_order", "requirement_group".
	CommentRecordType string `json:"comment_record_type"`
	Commenter         string `json:"commenter"`
	// Any of "admin", "user".
	CommenterType string `json:"commenter_type"`
	// An ISO 8901 datetime string denoting when the comment was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was read.
	ReadAt time.Time `json:"read_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Body              respjson.Field
		CommentRecordID   respjson.Field
		CommentRecordType respjson.Field
		Commenter         respjson.Field
		CommenterType     respjson.Field
		CreatedAt         respjson.Field
		ReadAt            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *CommentGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentListResponse struct {
	Data []CommentListResponseData `json:"data"`
	Meta PaginationMeta            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentListResponse) RawJSON() string { return r.JSON.raw }
func (r *CommentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentListResponseData struct {
	ID              string `json:"id" format:"uuid"`
	Body            string `json:"body"`
	CommentRecordID string `json:"comment_record_id" format:"uuid"`
	// Any of "sub_number_order", "requirement_group".
	CommentRecordType string `json:"comment_record_type"`
	Commenter         string `json:"commenter"`
	// Any of "admin", "user".
	CommenterType string `json:"commenter_type"`
	// An ISO 8901 datetime string denoting when the comment was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was read.
	ReadAt time.Time `json:"read_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Body              respjson.Field
		CommentRecordID   respjson.Field
		CommentRecordType respjson.Field
		Commenter         respjson.Field
		CommenterType     respjson.Field
		CreatedAt         respjson.Field
		ReadAt            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CommentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentMarkAsReadResponse struct {
	Data CommentMarkAsReadResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentMarkAsReadResponse) RawJSON() string { return r.JSON.raw }
func (r *CommentMarkAsReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentMarkAsReadResponseData struct {
	ID              string `json:"id" format:"uuid"`
	Body            string `json:"body"`
	CommentRecordID string `json:"comment_record_id" format:"uuid"`
	// Any of "sub_number_order", "requirement_group".
	CommentRecordType string `json:"comment_record_type"`
	Commenter         string `json:"commenter"`
	// Any of "admin", "user".
	CommenterType string `json:"commenter_type"`
	// An ISO 8901 datetime string denoting when the comment was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was read.
	ReadAt time.Time `json:"read_at" format:"date-time"`
	// An ISO 8901 datetime string for when the comment was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Body              respjson.Field
		CommentRecordID   respjson.Field
		CommentRecordType respjson.Field
		Commenter         respjson.Field
		CommenterType     respjson.Field
		CreatedAt         respjson.Field
		ReadAt            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentMarkAsReadResponseData) RawJSON() string { return r.JSON.raw }
func (r *CommentMarkAsReadResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentNewParams struct {
	Body            param.Opt[string] `json:"body,omitzero"`
	CommentRecordID param.Opt[string] `json:"comment_record_id,omitzero" format:"uuid"`
	// Any of "sub_number_order", "requirement_group".
	CommentRecordType CommentNewParamsCommentRecordType `json:"comment_record_type,omitzero"`
	paramObj
}

func (r CommentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CommentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentNewParamsCommentRecordType string

const (
	CommentNewParamsCommentRecordTypeSubNumberOrder   CommentNewParamsCommentRecordType = "sub_number_order"
	CommentNewParamsCommentRecordTypeRequirementGroup CommentNewParamsCommentRecordType = "requirement_group"
)

type CommentListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[comment_record_type], filter[comment_record_id]
	Filter CommentListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommentListParams]'s query parameters as `url.Values`.
func (r CommentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[comment_record_type], filter[comment_record_id]
type CommentListParamsFilter struct {
	// ID of the record the comments relate to
	CommentRecordID param.Opt[string] `query:"comment_record_id,omitzero" json:"-"`
	// Record type that the comment relates to
	//
	// Any of "sub_number_order", "requirement_group".
	CommentRecordType string `query:"comment_record_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommentListParamsFilter]'s query parameters as
// `url.Values`.
func (r CommentListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[CommentListParamsFilter](
		"comment_record_type", "sub_number_order", "requirement_group",
	)
}
