// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortoutSupportingDocumentService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortoutSupportingDocumentService] method instead.
type PortoutSupportingDocumentService struct {
	Options []option.RequestOption
}

// NewPortoutSupportingDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortoutSupportingDocumentService(opts ...option.RequestOption) (r PortoutSupportingDocumentService) {
	r = PortoutSupportingDocumentService{}
	r.Options = opts
	return
}

// Creates a list of supporting documents on a portout request.
func (r *PortoutSupportingDocumentService) New(ctx context.Context, id string, body PortoutSupportingDocumentNewParams, opts ...option.RequestOption) (res *PortoutSupportingDocumentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/%s/supporting_documents", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List every supporting documents for a portout request.
func (r *PortoutSupportingDocumentService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *PortoutSupportingDocumentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/%s/supporting_documents", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PortoutSupportingDocumentNewResponse struct {
	Data []PortoutSupportingDocumentNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutSupportingDocumentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutSupportingDocumentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutSupportingDocumentNewResponseData struct {
	ID string `json:"id,required" format:"uuid"`
	// Supporting document creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at,required"`
	// Identifies the associated document
	DocumentID string `json:"document_id,required" format:"uuid"`
	// Identifies the associated port request
	PortoutID string `json:"portout_id,required" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,required"`
	// Identifies the type of the document
	//
	// Any of "loa", "invoice".
	Type string `json:"type,required"`
	// Supporting document last changed timestamp in ISO 8601 format
	UpdatedAt string `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DocumentID  respjson.Field
		PortoutID   respjson.Field
		RecordType  respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutSupportingDocumentNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutSupportingDocumentNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutSupportingDocumentListResponse struct {
	Data []PortoutSupportingDocumentListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutSupportingDocumentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutSupportingDocumentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutSupportingDocumentListResponseData struct {
	ID string `json:"id,required" format:"uuid"`
	// Supporting document creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at,required"`
	// Identifies the associated document
	DocumentID string `json:"document_id,required" format:"uuid"`
	// Identifies the associated port request
	PortoutID string `json:"portout_id,required" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,required"`
	// Identifies the type of the document
	//
	// Any of "loa", "invoice".
	Type string `json:"type,required"`
	// Supporting document last changed timestamp in ISO 8601 format
	UpdatedAt string `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DocumentID  respjson.Field
		PortoutID   respjson.Field
		RecordType  respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutSupportingDocumentListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutSupportingDocumentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutSupportingDocumentNewParams struct {
	// List of supporting documents parameters
	Documents []PortoutSupportingDocumentNewParamsDocument `json:"documents,omitzero"`
	paramObj
}

func (r PortoutSupportingDocumentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortoutSupportingDocumentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortoutSupportingDocumentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DocumentID, Type are required.
type PortoutSupportingDocumentNewParamsDocument struct {
	// Identifies the associated document
	DocumentID string `json:"document_id,required" format:"uuid"`
	// Identifies the type of the document
	//
	// Any of "loa", "invoice".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r PortoutSupportingDocumentNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow PortoutSupportingDocumentNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortoutSupportingDocumentNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PortoutSupportingDocumentNewParamsDocument](
		"type", "loa", "invoice",
	)
}
