// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// PortoutCommentService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortoutCommentService] method instead.
type PortoutCommentService struct {
	Options []option.RequestOption
}

// NewPortoutCommentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortoutCommentService(opts ...option.RequestOption) (r PortoutCommentService) {
	r = PortoutCommentService{}
	r.Options = opts
	return
}

// Creates a comment on a portout request.
func (r *PortoutCommentService) New(ctx context.Context, id string, body PortoutCommentNewParams, opts ...option.RequestOption) (res *PortoutCommentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/%s/comments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of comments for a portout request.
func (r *PortoutCommentService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *PortoutCommentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/%s/comments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PortoutCommentNewResponse struct {
	Data PortoutCommentNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutCommentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutCommentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutCommentNewResponseData struct {
	ID string `json:"id,required"`
	// Comment body
	Body string `json:"body,required"`
	// Comment creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at,required"`
	// Identifies the user who created the comment. Will be null if created by Telnyx
	// Admin
	UserID string `json:"user_id,required"`
	// Identifies the associated port request
	PortoutID string `json:"portout_id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		CreatedAt   respjson.Field
		UserID      respjson.Field
		PortoutID   respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutCommentNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutCommentNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutCommentListResponse struct {
	Data []PortoutCommentListResponseData `json:"data"`
	Meta shared.Metadata                  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutCommentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutCommentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutCommentListResponseData struct {
	ID string `json:"id,required"`
	// Comment body
	Body string `json:"body,required"`
	// Comment creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at,required"`
	// Identifies the user who created the comment. Will be null if created by Telnyx
	// Admin
	UserID string `json:"user_id,required"`
	// Identifies the associated port request
	PortoutID string `json:"portout_id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		CreatedAt   respjson.Field
		UserID      respjson.Field
		PortoutID   respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutCommentListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutCommentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutCommentNewParams struct {
	// Comment to post on this portout request
	Body param.Opt[string] `json:"body,omitzero"`
	paramObj
}

func (r PortoutCommentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortoutCommentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortoutCommentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
