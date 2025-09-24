// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// DocumentService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r DocumentService) {
	r = DocumentService{}
	r.Options = opts
	return
}

// Retrieve a document.
func (r *DocumentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DocumentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a document.
func (r *DocumentService) Update(ctx context.Context, id string, body DocumentUpdateParams, opts ...option.RequestOption) (res *DocumentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all documents ordered by created_at descending.
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *DocumentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a document.<br /><br />A document can only be deleted if it's not linked
// to a service. If it is linked to a service, it must be unlinked prior to
// deleting.
func (r *DocumentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DocumentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Download a document.
func (r *DocumentService) Download(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("documents/%s/download", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generates a temporary pre-signed URL that can be used to download the document
// directly from the storage backend without authentication.
func (r *DocumentService) GenerateDownloadLink(ctx context.Context, id string, opts ...option.RequestOption) (res *DocumentGenerateDownloadLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("documents/%s/download_link", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a document.<br /><br />Uploaded files must be linked to a service within
// 30 minutes or they will be automatically deleted.
func (r *DocumentService) Upload(ctx context.Context, body DocumentUploadParams, opts ...option.RequestOption) (res *DocumentUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DocServiceDocument struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// The document's content_type.
	ContentType string `json:"content_type"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Optional reference string for customer tracking.
	CustomerReference string `json:"customer_reference"`
	// The filename of the document.
	Filename string `json:"filename"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The document's SHA256 hash provided for optional verification purposes.
	Sha256 string `json:"sha256"`
	// Indicates the document's filesize
	Size DocServiceDocumentSize `json:"size"`
	// Indicates the current document reviewing status
	//
	// Any of "pending", "verified", "denied".
	Status DocServiceDocumentStatus `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ContentType       respjson.Field
		CreatedAt         respjson.Field
		CustomerReference respjson.Field
		Filename          respjson.Field
		RecordType        respjson.Field
		Sha256            respjson.Field
		Size              respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocServiceDocument) RawJSON() string { return r.JSON.raw }
func (r *DocServiceDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DocServiceDocument to a DocServiceDocumentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DocServiceDocumentParam.Overrides()
func (r DocServiceDocument) ToParam() DocServiceDocumentParam {
	return param.Override[DocServiceDocumentParam](json.RawMessage(r.RawJSON()))
}

// Indicates the document's filesize
type DocServiceDocumentSize struct {
	// The number of bytes
	Amount int64 `json:"amount"`
	// Identifies the unit
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocServiceDocumentSize) RawJSON() string { return r.JSON.raw }
func (r *DocServiceDocumentSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the current document reviewing status
type DocServiceDocumentStatus string

const (
	DocServiceDocumentStatusPending  DocServiceDocumentStatus = "pending"
	DocServiceDocumentStatusVerified DocServiceDocumentStatus = "verified"
	DocServiceDocumentStatusDenied   DocServiceDocumentStatus = "denied"
)

type DocServiceDocumentParam struct {
	// Identifies the resource.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The document's content_type.
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt param.Opt[string] `json:"created_at,omitzero"`
	// Optional reference string for customer tracking.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// The filename of the document.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Identifies the type of the resource.
	RecordType param.Opt[string] `json:"record_type,omitzero"`
	// The document's SHA256 hash provided for optional verification purposes.
	Sha256 param.Opt[string] `json:"sha256,omitzero"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt param.Opt[string] `json:"updated_at,omitzero"`
	// Indicates the document's filesize
	Size DocServiceDocumentSizeParam `json:"size,omitzero"`
	// Indicates the current document reviewing status
	//
	// Any of "pending", "verified", "denied".
	Status DocServiceDocumentStatus `json:"status,omitzero"`
	paramObj
}

func (r DocServiceDocumentParam) MarshalJSON() (data []byte, err error) {
	type shadow DocServiceDocumentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocServiceDocumentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the document's filesize
type DocServiceDocumentSizeParam struct {
	// The number of bytes
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// Identifies the unit
	Unit param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r DocServiceDocumentSizeParam) MarshalJSON() (data []byte, err error) {
	type shadow DocServiceDocumentSizeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocServiceDocumentSizeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentGetResponse struct {
	Data DocServiceDocument `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUpdateResponse struct {
	Data DocServiceDocument `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListResponse struct {
	Data []DocServiceDocument `json:"data"`
	Meta PaginationMeta       `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentListResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentDeleteResponse struct {
	Data DocServiceDocument `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentGenerateDownloadLinkResponse struct {
	Data DocumentGenerateDownloadLinkResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentGenerateDownloadLinkResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentGenerateDownloadLinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentGenerateDownloadLinkResponseData struct {
	// Pre-signed temporary URL for downloading the document
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentGenerateDownloadLinkResponseData) RawJSON() string { return r.JSON.raw }
func (r *DocumentGenerateDownloadLinkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUploadResponse struct {
	Data DocServiceDocument `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUpdateParams struct {
	DocServiceDocument DocServiceDocumentParam
	paramObj
}

func (r DocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DocServiceDocument)
}
func (r *DocumentUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DocServiceDocument)
}

type DocumentListParams struct {
	// Consolidated filter parameter for documents (deepObject style). Originally:
	// filter[filename][contains], filter[customer_reference][eq],
	// filter[customer_reference][in][], filter[created_at][gt], filter[created_at][lt]
	Filter DocumentListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page DocumentListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter for documents (deepObject style). Originally: sort[]
	//
	// Any of "filename", "created_at", "updated_at", "-filename", "-created_at",
	// "-updated_at".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for documents (deepObject style). Originally:
// filter[filename][contains], filter[customer_reference][eq],
// filter[customer_reference][in][], filter[created_at][gt], filter[created_at][lt]
type DocumentListParamsFilter struct {
	CreatedAt         DocumentListParamsFilterCreatedAt         `query:"created_at,omitzero" json:"-"`
	CustomerReference DocumentListParamsFilterCustomerReference `query:"customer_reference,omitzero" json:"-"`
	Filename          DocumentListParamsFilterFilename          `query:"filename,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParamsFilter]'s query parameters as
// `url.Values`.
func (r DocumentListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentListParamsFilterCreatedAt struct {
	// Filter by created at greater than provided value.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Filter by created at less than provided value.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParamsFilterCreatedAt]'s query parameters as
// `url.Values`.
func (r DocumentListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentListParamsFilterCustomerReference struct {
	// Filter documents by a customer reference.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filter documents by a list of customer references.
	In []string `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParamsFilterCustomerReference]'s query
// parameters as `url.Values`.
func (r DocumentListParamsFilterCustomerReference) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentListParamsFilterFilename struct {
	// Filter by string matching part of filename.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParamsFilterFilename]'s query parameters as
// `url.Values`.
func (r DocumentListParamsFilterFilename) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type DocumentListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParamsPage]'s query parameters as `url.Values`.
func (r DocumentListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentUploadParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfDocServiceDocumentUploadURL *DocumentUploadParamsBodyDocServiceDocumentUploadURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfDocServiceDocumentUploadInline *DocumentUploadParamsBodyDocServiceDocumentUploadInline `json:",inline"`

	paramObj
}

func (u DocumentUploadParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDocServiceDocumentUploadURL, u.OfDocServiceDocumentUploadInline)
}
func (r *DocumentUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type DocumentUploadParamsBodyDocServiceDocumentUploadURL struct {
	// If the file is already hosted publicly, you can provide a URL and have the
	// documents service fetch it for you.
	URL string `json:"url,required"`
	// Optional reference string for customer tracking.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// The filename of the document.
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r DocumentUploadParamsBodyDocServiceDocumentUploadURL) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUploadParamsBodyDocServiceDocumentUploadURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUploadParamsBodyDocServiceDocumentUploadURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property File is required.
type DocumentUploadParamsBodyDocServiceDocumentUploadInline struct {
	// The Base64 encoded contents of the file you are uploading.
	File string `json:"file,required" format:"byte"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// The filename of the document.
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r DocumentUploadParamsBodyDocServiceDocumentUploadInline) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUploadParamsBodyDocServiceDocumentUploadInline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUploadParamsBodyDocServiceDocumentUploadInline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
