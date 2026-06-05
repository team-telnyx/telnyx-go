// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Phone-number reputation monitoring (spam-score lookup and tracking).
//
// EnterpriseReputationLoaService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationLoaService] method instead.
type EnterpriseReputationLoaService struct {
	Options []option.RequestOption
}

// NewEnterpriseReputationLoaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnterpriseReputationLoaService(opts ...option.RequestOption) (r EnterpriseReputationLoaService) {
	r = EnterpriseReputationLoaService{}
	r.Options = opts
	return
}

// Point the enterprise's reputation settings at a new signed LOA document. This
// resets LOA approval to `pending`; the new document must be approved before
// additional phone numbers can be added.
func (r *EnterpriseReputationLoaService) Update(ctx context.Context, enterpriseID string, body EnterpriseReputationLoaUpdateParams, opts ...option.RequestOption) (res *EnterpriseReputationLoaUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/loa", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Render the LOA for this enterprise as a PDF. The enterprise identity, address,
// and authorized-representative contact are taken from the enterprise record; the
// optional `agent` block is supplied only when a third-party partner manages the
// numbers. The response is the PDF itself (unsigned unless a `signature` is
// provided). Sign it and upload it to the Telnyx Documents API
// (`POST /v2/documents`, see https://developers.telnyx.com/api/documents) to
// obtain the `loa_document_id` required by `POST .../reputation`.
func (r *EnterpriseReputationLoaService) Render(ctx context.Context, enterpriseID string, body EnterpriseReputationLoaRenderParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/loa", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EnterpriseReputationLoaUpdateResponse struct {
	Data EnterpriseReputationPublic `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationLoaUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationLoaUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationLoaUpdateParams struct {
	// Id of the new signed LOA document (from the Telnyx Documents API). Changing it
	// resets LOA approval; the new document must be approved before more numbers can
	// be added.
	LoaDocumentID string `json:"loa_document_id" api:"required"`
	paramObj
}

func (r EnterpriseReputationLoaUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationLoaUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationLoaUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationLoaRenderParams struct {
	// Third-party reseller / partner managing the enterprise's phone numbers. Omit
	// when the enterprise works directly with Telnyx.
	Agent EnterpriseReputationLoaRenderParamsAgent `json:"agent,omitzero"`
	// Optional signature embedded in the rendered PDF. When omitted the PDF is
	// returned unsigned for the customer to sign and upload.
	Signature EnterpriseReputationLoaRenderParamsSignature `json:"signature,omitzero"`
	paramObj
}

func (r EnterpriseReputationLoaRenderParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationLoaRenderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationLoaRenderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Third-party reseller / partner managing the enterprise's phone numbers. Omit
// when the enterprise works directly with Telnyx.
//
// The properties AdministrativeArea, City, ContactEmail, ContactName,
// ContactPhone, ContactTitle, Country, LegalName, PostalCode, StreetAddress are
// required.
type EnterpriseReputationLoaRenderParamsAgent struct {
	AdministrativeArea string            `json:"administrative_area" api:"required"`
	City               string            `json:"city" api:"required"`
	ContactEmail       string            `json:"contact_email" api:"required" format:"email"`
	ContactName        string            `json:"contact_name" api:"required"`
	ContactPhone       string            `json:"contact_phone" api:"required"`
	ContactTitle       string            `json:"contact_title" api:"required"`
	Country            string            `json:"country" api:"required"`
	LegalName          string            `json:"legal_name" api:"required"`
	PostalCode         string            `json:"postal_code" api:"required"`
	StreetAddress      string            `json:"street_address" api:"required"`
	Dba                param.Opt[string] `json:"dba,omitzero"`
	ExtendedAddress    param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r EnterpriseReputationLoaRenderParamsAgent) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationLoaRenderParamsAgent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationLoaRenderParamsAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional signature embedded in the rendered PDF. When omitted the PDF is
// returned unsigned for the customer to sign and upload.
//
// The property ImageBase64 is required.
type EnterpriseReputationLoaRenderParamsSignature struct {
	// Base64-encoded signature image.
	ImageBase64 string            `json:"image_base64" api:"required"`
	SignerName  param.Opt[string] `json:"signer_name,omitzero"`
	paramObj
}

func (r EnterpriseReputationLoaRenderParamsSignature) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationLoaRenderParamsSignature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationLoaRenderParamsSignature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
