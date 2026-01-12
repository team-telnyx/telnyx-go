// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// DocumentLinkService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentLinkService] method instead.
type DocumentLinkService struct {
	Options []option.RequestOption
}

// NewDocumentLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentLinkService(opts ...option.RequestOption) (r DocumentLinkService) {
	r = DocumentLinkService{}
	r.Options = opts
	return
}

// List all documents links ordered by created_at descending.
func (r *DocumentLinkService) List(ctx context.Context, query DocumentLinkListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[DocumentLinkListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "document_links"
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

// List all documents links ordered by created_at descending.
func (r *DocumentLinkService) ListAutoPaging(ctx context.Context, query DocumentLinkListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[DocumentLinkListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

type DocumentLinkListResponse struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Identifies the associated document.
	DocumentID string `json:"document_id" format:"uuid"`
	// The linked resource's record type.
	LinkedRecordType string `json:"linked_record_type"`
	// Identifies the linked resource.
	LinkedResourceID string `json:"linked_resource_id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DocumentID       respjson.Field
		LinkedRecordType respjson.Field
		LinkedResourceID respjson.Field
		RecordType       respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentLinkListResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentLinkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentLinkListParams struct {
	// Consolidated filter parameter for document links (deepObject style). Originally:
	// filter[linked_record_type], filter[linked_resource_id]
	Filter DocumentLinkListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page DocumentLinkListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentLinkListParams]'s query parameters as `url.Values`.
func (r DocumentLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for document links (deepObject style). Originally:
// filter[linked_record_type], filter[linked_resource_id]
type DocumentLinkListParamsFilter struct {
	// The linked_record_type of the document to filter on.
	LinkedRecordType param.Opt[string] `query:"linked_record_type,omitzero" json:"-"`
	// The linked_resource_id of the document to filter on.
	LinkedResourceID param.Opt[string] `query:"linked_resource_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentLinkListParamsFilter]'s query parameters as
// `url.Values`.
func (r DocumentLinkListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type DocumentLinkListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentLinkListParamsPage]'s query parameters as
// `url.Values`.
func (r DocumentLinkListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
