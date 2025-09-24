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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortingOrderAdditionalDocumentService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderAdditionalDocumentService] method instead.
type PortingOrderAdditionalDocumentService struct {
	Options []option.RequestOption
}

// NewPortingOrderAdditionalDocumentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderAdditionalDocumentService(opts ...option.RequestOption) (r PortingOrderAdditionalDocumentService) {
	r = PortingOrderAdditionalDocumentService{}
	r.Options = opts
	return
}

// Creates a list of additional documents for a porting order.
func (r *PortingOrderAdditionalDocumentService) New(ctx context.Context, id string, body PortingOrderAdditionalDocumentNewParams, opts ...option.RequestOption) (res *PortingOrderAdditionalDocumentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/additional_documents", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of additional documents for a porting order.
func (r *PortingOrderAdditionalDocumentService) List(ctx context.Context, id string, query PortingOrderAdditionalDocumentListParams, opts ...option.RequestOption) (res *PortingOrderAdditionalDocumentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/additional_documents", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an additional document for a porting order.
func (r *PortingOrderAdditionalDocumentService) Delete(ctx context.Context, additionalDocumentID string, body PortingOrderAdditionalDocumentDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if additionalDocumentID == "" {
		err = errors.New("missing required additional_document_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/additional_documents/%s", body.ID, additionalDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PortingOrderAdditionalDocumentNewResponse struct {
	Data []PortingOrderAdditionalDocumentNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAdditionalDocumentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAdditionalDocumentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAdditionalDocumentNewResponseData struct {
	// Uniquely identifies this additional document
	ID string `json:"id" format:"uuid"`
	// The content type of the related document.
	ContentType string `json:"content_type"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the associated document
	DocumentID string `json:"document_id" format:"uuid"`
	// Identifies the type of additional document
	//
	// Any of "loa", "invoice", "csr", "other".
	DocumentType string `json:"document_type"`
	// The filename of the related document.
	Filename string `json:"filename"`
	// Identifies the associated porting order
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ContentType    respjson.Field
		CreatedAt      respjson.Field
		DocumentID     respjson.Field
		DocumentType   respjson.Field
		Filename       respjson.Field
		PortingOrderID respjson.Field
		RecordType     respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAdditionalDocumentNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAdditionalDocumentNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAdditionalDocumentListResponse struct {
	Data []PortingOrderAdditionalDocumentListResponseData `json:"data"`
	Meta PaginationMeta                                   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAdditionalDocumentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAdditionalDocumentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAdditionalDocumentListResponseData struct {
	// Uniquely identifies this additional document
	ID string `json:"id" format:"uuid"`
	// The content type of the related document.
	ContentType string `json:"content_type"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the associated document
	DocumentID string `json:"document_id" format:"uuid"`
	// Identifies the type of additional document
	//
	// Any of "loa", "invoice", "csr", "other".
	DocumentType string `json:"document_type"`
	// The filename of the related document.
	Filename string `json:"filename"`
	// Identifies the associated porting order
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ContentType    respjson.Field
		CreatedAt      respjson.Field
		DocumentID     respjson.Field
		DocumentType   respjson.Field
		Filename       respjson.Field
		PortingOrderID respjson.Field
		RecordType     respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderAdditionalDocumentListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderAdditionalDocumentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAdditionalDocumentNewParams struct {
	AdditionalDocuments []PortingOrderAdditionalDocumentNewParamsAdditionalDocument `json:"additional_documents,omitzero"`
	paramObj
}

func (r PortingOrderAdditionalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderAdditionalDocumentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderAdditionalDocumentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderAdditionalDocumentNewParamsAdditionalDocument struct {
	// The document identification
	DocumentID param.Opt[string] `json:"document_id,omitzero" format:"uuid"`
	// The type of document being created.
	//
	// Any of "loa", "invoice", "csr", "other".
	DocumentType string `json:"document_type,omitzero"`
	paramObj
}

func (r PortingOrderAdditionalDocumentNewParamsAdditionalDocument) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderAdditionalDocumentNewParamsAdditionalDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderAdditionalDocumentNewParamsAdditionalDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PortingOrderAdditionalDocumentNewParamsAdditionalDocument](
		"document_type", "loa", "invoice", "csr", "other",
	)
}

type PortingOrderAdditionalDocumentListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[document_type]
	Filter PortingOrderAdditionalDocumentListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderAdditionalDocumentListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderAdditionalDocumentListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAdditionalDocumentListParams]'s query
// parameters as `url.Values`.
func (r PortingOrderAdditionalDocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[document_type]
type PortingOrderAdditionalDocumentListParamsFilter struct {
	// Filter additional documents by a list of document types
	//
	// Any of "loa", "invoice", "csr", "other".
	DocumentType []string `query:"document_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAdditionalDocumentListParamsFilter]'s query
// parameters as `url.Values`.
func (r PortingOrderAdditionalDocumentListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderAdditionalDocumentListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAdditionalDocumentListParamsPage]'s query
// parameters as `url.Values`.
func (r PortingOrderAdditionalDocumentListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderAdditionalDocumentListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at", "-created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderAdditionalDocumentListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderAdditionalDocumentListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderAdditionalDocumentDeleteParams struct {
	ID string `path:"id,required" format:"uuid" json:"-"`
	paramObj
}
