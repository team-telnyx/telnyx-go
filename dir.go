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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// DirService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirService] method instead.
type DirService struct {
	Options []option.RequestOption
	// Read messages from the Telnyx vetting team and reply with clarifying
	// information.
	Comments DirCommentService
	// Phone numbers are submitted to Telnyx for vetting in batches. Batches group all
	// numbers added in a single request under the same Letter of Authorization.
	PhoneNumberBatches DirPhoneNumberBatchService
	// Associate phone numbers with a verified DIR so calls from those numbers carry
	// the DIR's display identity.
	PhoneNumbers DirPhoneNumberService
	// Submit and manage the two business references and one financial reference that
	// vouch for a DIR. References are contacted to confirm the business identity
	// during vetting.
	References DirReferenceService
	// Verify ownership of a DIR's authorizer email. A short code is emailed and
	// confirmed; the email must be verified before references can be submitted.
	VerifyEmail DirVerifyEmailService
}

// NewDirService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDirService(opts ...option.RequestOption) (r DirService) {
	r = DirService{}
	r.Options = opts
	r.Comments = NewDirCommentService(opts...)
	r.PhoneNumberBatches = NewDirPhoneNumberBatchService(opts...)
	r.PhoneNumbers = NewDirPhoneNumberService(opts...)
	r.References = NewDirReferenceService(opts...)
	r.VerifyEmail = NewDirVerifyEmailService(opts...)
	return
}

// Returns a single DIR by id. The enterprise is resolved server-side from the DIR
// id. Returns `404` if the DIR does not exist or is not yours.
func (r *DirService) Get(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Edit a DIR. DIRs in `draft`, `rejected`, `unsuccessful`, or `suspended` can be
// edited freely: PATCH is a pure edit, `status` is never changed, and you re-vet
// by calling `POST /v2/dir/{dir_id}/submit` explicitly. A `verified` DIR can also
// be edited in place: a PATCH that changes any value returns the DIR to `draft`
// and branded delivery stops until you re-submit and the DIR is approved again,
// while a PATCH that changes nothing (an empty body or values identical to the
// current ones) leaves the DIR `verified`, so idempotent retries are safe. DIRs in
// any other status (`submitted`, `in_review`, `expired`, `infringement_claimed`,
// `permanently_rejected`) cannot be edited.
func (r *DirService) Update(ctx context.Context, dirID string, body DirUpdateParams, opts ...option.RequestOption) (res *DirWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns every DIR (Display Identity Record) you own, across all of your
// enterprises, as a single list. Pagination is JSON:API style (`page[number]`,
// `page[size]`, max 250). Supports `filter[]` query params:
// `filter[enterprise_id]`, `filter[status]`, `filter[display_name][contains]`,
// `filter[call_reason][contains]`, plus the renewal-window filters
// `filter[expiring_at][gte]` / `filter[expiring_at][lte]`. Sortable by
// `created_at`, `updated_at`, `display_name`, `status` (prefix `-` for descending;
// default `-created_at`).
func (r *DirService) List(ctx context.Context, query DirListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Dir], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "dir"
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

// Returns every DIR (Display Identity Record) you own, across all of your
// enterprises, as a single list. Pagination is JSON:API style (`page[number]`,
// `page[size]`, max 250). Supports `filter[]` query params:
// `filter[enterprise_id]`, `filter[status]`, `filter[display_name][contains]`,
// `filter[call_reason][contains]`, plus the renewal-window filters
// `filter[expiring_at][gte]` / `filter[expiring_at][lte]`. Sortable by
// `created_at`, `updated_at`, `display_name`, `status` (prefix `-` for descending;
// default `-created_at`).
func (r *DirService) ListAutoPaging(ctx context.Context, query DirListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Dir] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a DIR. Failure modes: `400` if a child phone number is in a non-deletable
// status, `409` if the DIR has an unresolved infringement claim, `404` if the DIR
// is not yours.
func (r *DirService) Delete(ctx context.Context, dirID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return err
	}
	path := fmt.Sprintf("dir/%s", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Reference list of `document_type` values accepted by
// `DirCreateRequest.documents[].document_type` and the infringement-contest
// endpoint. Each entry has a stable `short_name` (used in API calls) and a
// customer-facing description.
func (r *DirService) ListDocumentTypes(ctx context.Context, opts ...option.RequestOption) (res *DirListDocumentTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dir/document_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Return the trademark or copyright claims filed against this DIR. Each claim's
// `status` is `pending` (newly filed; DIR auto-suspended), `contested` (you have
// submitted contest evidence; awaiting resolution), or `resolved` (final).
// Resolution outcomes: `upheld` (claim accepted; DIR stays
// suspended/permanently_rejected), `rejected` (claim dismissed; DIR restored to
// `verified`), `modified` (partial outcome).
func (r *DirService) ListInfringementClaims(ctx context.Context, dirID string, query DirListInfringementClaimsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[InfringementClaim], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/infringement_claims", dirID)
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

// Return the trademark or copyright claims filed against this DIR. Each claim's
// `status` is `pending` (newly filed; DIR auto-suspended), `contested` (you have
// submitted contest evidence; awaiting resolution), or `resolved` (final).
// Resolution outcomes: `upheld` (claim accepted; DIR stays
// suspended/permanently_rejected), `rejected` (claim dismissed; DIR restored to
// `verified`), `modified` (partial outcome).
func (r *DirService) ListInfringementClaimsAutoPaging(ctx context.Context, dirID string, query DirListInfringementClaimsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[InfringementClaim] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListInfringementClaims(ctx, dirID, query, opts...))
}

// Generate a pre-filled Letter of Authorization (LOA) PDF for a DIR. Enterprise
// identity (legal name, DBA, address, contact, website, tax id) and the DIR
// display name are read server-side; the caller supplies the telephone numbers to
// authorize, an optional Authorized Agent block, and an optional drawn signature.
//
// When `signature` is omitted the PDF is returned unsigned so the customer can
// sign it externally and upload it via the Documents API. When `signature` is
// present the PDF embeds the supplied image, printed name, and signed-at date.
//
// Returns `application/pdf`.
func (r *DirService) NewLoa(ctx context.Context, dirID string, body DirNewLoaParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/loa", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submit a DIR for vetting. Sends the DIR back through the vetting cycle from any
// non-terminal status. When re-submitting from `suspended` or `expired`, the DIR's
// previous Branded Calling registration is torn down transactionally and its phone
// numbers flip back to `submitted`. When re-submitting from `verified`, the
// existing registration stays live throughout the new vetting cycle.
//
// Returns `400` from `submitted`/`in_review`/`permanently_rejected`. Returns `409`
// if the DIR has an unresolved infringement claim.
func (r *DirService) Submit(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/submit", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Push a fix for a DIR that is `suspended` with an open infringement claim back
// into vetting. `POST /dir/{dir_id}/submit` is blocked while a claim is open, so
// this is the customer-callable path to update the DIR's content and re-certify
// before Telnyx adjudicates the claim. All four certification booleans must be
// `true`. Optional content fields (`display_name`, `logo_url`, `call_reasons`,
// `documents`) update the DIR; documents are append-only.
func (r *DirService) UpdateInfringement(ctx context.Context, dirID string, body DirUpdateInfringementParams, opts ...option.RequestOption) (res *DirWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/infringement_update", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type Dir struct {
	ID                     string          `json:"id" format:"uuid"`
	AuthorizerEmail        string          `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string          `json:"authorizer_name" api:"nullable"`
	CallReasons            []DirCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool            `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool            `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool            `json:"certify_no_shaft_content"`
	CreatedAt              time.Time       `json:"created_at" format:"date-time"`
	DisplayName            string          `json:"display_name"`
	Documents              []Document      `json:"documents" api:"nullable"`
	EnterpriseID           string          `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time       `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string          `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time       `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []RejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool              `json:"reselling"`
	// DIR lifecycle status.
	//
	//   - `draft` - newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` - Telnyx is reviewing.
	//   - `verified` - approved; phone numbers may be attached.
	//   - `rejected` - Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` - system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` - temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` - verification expired; customer must resubmit.
	//   - `infringement_claimed` - a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` - terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status      DirStatus `json:"status"`
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	VerifiedAt  time.Time `json:"verified_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AuthorizerEmail        respjson.Field
		AuthorizerName         respjson.Field
		CallReasons            respjson.Field
		CertifyBrandIsAccurate respjson.Field
		CertifyIPOwnership     respjson.Field
		CertifyNoShaftContent  respjson.Field
		CreatedAt              respjson.Field
		DisplayName            respjson.Field
		Documents              respjson.Field
		EnterpriseID           respjson.Field
		ExpiringAt             respjson.Field
		LogoURL                respjson.Field
		RejectedAt             respjson.Field
		RejectionReasons       respjson.Field
		Reselling              respjson.Field
		Status                 respjson.Field
		SubmittedAt            respjson.Field
		UpdatedAt              respjson.Field
		VerifiedAt             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Dir) RawJSON() string { return r.JSON.raw }
func (r *Dir) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirCallReason struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Reason    string    `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirCallReason) RawJSON() string { return r.JSON.raw }
func (r *DirCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirList struct {
	Data []Dir `json:"data" api:"required"`
	// JSON:API pagination metadata returned with every paginated list response. Page
	// numbering is 1-based. `page_size` reports the number of items actually returned
	// in `data` for this page; the requested size is taken from the `page[size]` query
	// parameter.
	Meta BrandedCallingPaginationMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirList) RawJSON() string { return r.JSON.raw }
func (r *DirList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DIR lifecycle status.
//
//   - `draft` - newly created; editable; not yet submitted.
//   - `submitted` / `in_review` - Telnyx is reviewing.
//   - `verified` - approved; phone numbers may be attached.
//   - `rejected` - Telnyx rejected this submission; `rejection_reasons` is
//     populated; customer can edit and resubmit.
//   - `unsuccessful` - system-side error during processing; customer can edit and
//     resubmit.
//   - `suspended` - temporarily disabled (e.g. by an active infringement claim).
//   - `expired` - verification expired; customer must resubmit.
//   - `infringement_claimed` - a trademark/impersonation claim is open against this
//     DIR.
//   - `permanently_rejected` - terminal; cannot be resubmitted.
type DirStatus string

const (
	DirStatusDraft               DirStatus = "draft"
	DirStatusSubmitted           DirStatus = "submitted"
	DirStatusInReview            DirStatus = "in_review"
	DirStatusVerified            DirStatus = "verified"
	DirStatusRejected            DirStatus = "rejected"
	DirStatusUnsuccessful        DirStatus = "unsuccessful"
	DirStatusSuspended           DirStatus = "suspended"
	DirStatusExpired             DirStatus = "expired"
	DirStatusInfringementClaimed DirStatus = "infringement_claimed"
	DirStatusPermanentlyRejected DirStatus = "permanently_rejected"
)

type DirWrapped struct {
	Data Dir `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirWrapped) RawJSON() string { return r.JSON.raw }
func (r *DirWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Document struct {
	// Id returned by the Telnyx Documents API after you upload the file (upload via
	// `POST /v2/documents`; see https://developers.telnyx.com/api/documents).
	DocumentID string `json:"document_id" api:"required" format:"uuid"`
	// Type of supporting document. Pick the closest match to what the file actually
	// contains; `other` triggers manual vetting and may slow approval. The matching
	// short_name reference list is at `GET /v2/dir/document_types`.
	//
	// Any of "letter_of_authorization", "business_registration",
	// "articles_of_incorporation", "tax_document", "ein_letter",
	// "trademark_registration", "website_ownership", "business_license",
	// "professional_license", "government_id", "utility_bill", "bank_statement",
	// "other".
	DocumentType DocumentDocumentType `json:"document_type" api:"required"`
	Description  string               `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID   respjson.Field
		DocumentType respjson.Field
		Description  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Document) RawJSON() string { return r.JSON.raw }
func (r *Document) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Document to a DocumentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DocumentParam.Overrides()
func (r Document) ToParam() DocumentParam {
	return param.Override[DocumentParam](json.RawMessage(r.RawJSON()))
}

// Type of supporting document. Pick the closest match to what the file actually
// contains; `other` triggers manual vetting and may slow approval. The matching
// short_name reference list is at `GET /v2/dir/document_types`.
type DocumentDocumentType string

const (
	DocumentDocumentTypeLetterOfAuthorization   DocumentDocumentType = "letter_of_authorization"
	DocumentDocumentTypeBusinessRegistration    DocumentDocumentType = "business_registration"
	DocumentDocumentTypeArticlesOfIncorporation DocumentDocumentType = "articles_of_incorporation"
	DocumentDocumentTypeTaxDocument             DocumentDocumentType = "tax_document"
	DocumentDocumentTypeEinLetter               DocumentDocumentType = "ein_letter"
	DocumentDocumentTypeTrademarkRegistration   DocumentDocumentType = "trademark_registration"
	DocumentDocumentTypeWebsiteOwnership        DocumentDocumentType = "website_ownership"
	DocumentDocumentTypeBusinessLicense         DocumentDocumentType = "business_license"
	DocumentDocumentTypeProfessionalLicense     DocumentDocumentType = "professional_license"
	DocumentDocumentTypeGovernmentID            DocumentDocumentType = "government_id"
	DocumentDocumentTypeUtilityBill             DocumentDocumentType = "utility_bill"
	DocumentDocumentTypeBankStatement           DocumentDocumentType = "bank_statement"
	DocumentDocumentTypeOther                   DocumentDocumentType = "other"
)

// The properties DocumentID, DocumentType are required.
type DocumentParam struct {
	// Id returned by the Telnyx Documents API after you upload the file (upload via
	// `POST /v2/documents`; see https://developers.telnyx.com/api/documents).
	DocumentID string `json:"document_id" api:"required" format:"uuid"`
	// Type of supporting document. Pick the closest match to what the file actually
	// contains; `other` triggers manual vetting and may slow approval. The matching
	// short_name reference list is at `GET /v2/dir/document_types`.
	//
	// Any of "letter_of_authorization", "business_registration",
	// "articles_of_incorporation", "tax_document", "ein_letter",
	// "trademark_registration", "website_ownership", "business_license",
	// "professional_license", "government_id", "utility_bill", "bank_statement",
	// "other".
	DocumentType DocumentDocumentType `json:"document_type,omitzero" api:"required"`
	Description  param.Opt[string]    `json:"description,omitzero"`
	paramObj
}

func (r DocumentParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListDocumentTypesResponse struct {
	Data []DirListDocumentTypesResponseData `json:"data" api:"required"`
	// JSON:API pagination metadata returned with every paginated list response. Page
	// numbering is 1-based. `page_size` reports the number of items actually returned
	// in `data` for this page; the requested size is taken from the `page[size]` query
	// parameter.
	Meta BrandedCallingPaginationMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListDocumentTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *DirListDocumentTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single supported document type.
type DirListDocumentTypesResponseData struct {
	Description string `json:"description"`
	// Stable identifier passed to `Document.document_type`.
	ShortName string `json:"short_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ShortName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListDocumentTypesResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirListDocumentTypesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateParams struct {
	// Contact email of the authorizer. Telnyx may send verification or infringement
	// notices here.
	AuthorizerEmail param.Opt[string] `json:"authorizer_email,omitzero" format:"email"`
	// Name of the person at your enterprise authorizing this DIR. Must be a real
	// individual.
	AuthorizerName param.Opt[string] `json:"authorizer_name,omitzero"`
	// Certification that the DIR information is accurate. Must be `true` for the DIR
	// to be submitted for vetting.
	CertifyBrandIsAccurate param.Opt[bool] `json:"certify_brand_is_accurate,omitzero"`
	// Certification of ownership of any logos/trademarks shown. Must be `true` for the
	// DIR to be submitted for vetting.
	CertifyIPOwnership param.Opt[bool] `json:"certify_ip_ownership,omitzero"`
	// Certification that this DIR is not used for SHAFT content (Sex, Hate, Alcohol,
	// Firearms, Tobacco) where prohibited. Must be `true` for the DIR to be submitted
	// for vetting.
	CertifyNoShaftContent param.Opt[bool] `json:"certify_no_shaft_content,omitzero"`
	// Name shown to call recipients. 1–35 characters, no emoji, not whitespace-only.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Publicly accessible HTTPS URL (max 128 chars) to a 256x256 BMP logo (max 1 MB).
	LogoURL param.Opt[string] `json:"logo_url,omitzero" format:"uri"`
	// Set to true if your organization places calls on behalf of other enterprises
	// (BPO/reseller). Updating this triggers re-vetting on next submit.
	Reselling param.Opt[bool] `json:"reselling,omitzero"`
	// 1–10 reasons your business calls customers. Validate phrasing against
	// `POST /call_reasons/validate`.
	CallReasons []string `json:"call_reasons,omitzero"`
	// Additional supporting documents to attach. Append-only: existing documents are
	// never removed or replaced, and an empty or omitted list is a no-op. Each
	// `document_id` may appear at most once on a DIR.
	Documents []DocumentParam `json:"documents,omitzero"`
	paramObj
}

func (r DirUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DirUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListParams struct {
	// Case-insensitive partial match on call reason.
	FilterCallReasonContains param.Opt[string] `query:"filter[call_reason][contains],omitzero" json:"-"`
	// Case-insensitive partial match on display name.
	FilterDisplayNameContains param.Opt[string] `query:"filter[display_name][contains],omitzero" json:"-"`
	// Filter by enterprise ID.
	FilterEnterpriseID param.Opt[string] `query:"filter[enterprise_id],omitzero" format:"uuid" json:"-"`
	// Return only DIRs whose `expiring_at` is at or after this ISO-8601 timestamp.
	// Pairs with the `[lte]` variant to build renewal-window dashboards.
	FilterExpiringAtGte param.Opt[time.Time] `query:"filter[expiring_at][gte],omitzero" format:"date-time" json:"-"`
	// Return only DIRs whose `expiring_at` is at or before this ISO-8601 timestamp.
	FilterExpiringAtLte param.Opt[time.Time] `query:"filter[expiring_at][lte],omitzero" format:"date-time" json:"-"`
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by DIR status.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	FilterStatus DirStatus `query:"filter[status],omitzero" json:"-"`
	// Sort field. Allowed values: `created_at`, `updated_at`, `display_name`,
	// `status`. Prefix with `-` for descending. Default `-created_at`.
	//
	// Any of "created_at", "-created_at", "updated_at", "-updated_at", "display_name",
	// "-display_name", "status", "-status".
	Sort DirListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirListParams]'s query parameters as `url.Values`.
func (r DirListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field. Allowed values: `created_at`, `updated_at`, `display_name`,
// `status`. Prefix with `-` for descending. Default `-created_at`.
type DirListParamsSort string

const (
	DirListParamsSortCreatedAt        DirListParamsSort = "created_at"
	DirListParamsSortCreatedAtDesc    DirListParamsSort = "-created_at"
	DirListParamsSortUpdatedAt        DirListParamsSort = "updated_at"
	DirListParamsSortUpdatedAtDesc    DirListParamsSort = "-updated_at"
	DirListParamsSortDisplayName      DirListParamsSort = "display_name"
	DirListParamsSortMinusDisplayName DirListParamsSort = "-display_name"
	DirListParamsSortStatus           DirListParamsSort = "status"
	DirListParamsSortStatusDesc       DirListParamsSort = "-status"
)

type DirListInfringementClaimsParams struct {
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirListInfringementClaimsParams]'s query parameters as
// `url.Values`.
func (r DirListInfringementClaimsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DirNewLoaParams struct {
	// Telephone numbers to authorize on the DIR, in `+E164` format (`+` followed by
	// 10-15 digits). Max 15 per request.
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	// Third-party reseller / partner managing the enterprise's phone numbers. Omit
	// when the enterprise works directly with Telnyx.
	Agent AgentInputParam `json:"agent,omitzero"`
	// Optional. When provided the rendered PDF embeds the signature image, printed
	// name, and signed-at date. When absent the PDF is returned unsigned so the
	// customer can sign externally and upload it via the Documents API.
	Signature DirNewLoaParamsSignature `json:"signature,omitzero"`
	paramObj
}

func (r DirNewLoaParams) MarshalJSON() (data []byte, err error) {
	type shadow DirNewLoaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirNewLoaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. When provided the rendered PDF embeds the signature image, printed
// name, and signed-at date. When absent the PDF is returned unsigned so the
// customer can sign externally and upload it via the Documents API.
//
// The property ImageBase64 is required.
type DirNewLoaParamsSignature struct {
	// PNG image, base64-encoded.
	ImageBase64 string `json:"image_base64" api:"required"`
	// Optional. When absent the rendered PDF falls back to the enterprise contact's
	// legal name.
	SignerName param.Opt[string] `json:"signer_name,omitzero"`
	paramObj
}

func (r DirNewLoaParamsSignature) MarshalJSON() (data []byte, err error) {
	type shadow DirNewLoaParamsSignature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirNewLoaParamsSignature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateInfringementParams struct {
	// Must be `true`.
	//
	// Any of true.
	CertifyBrandIsAccurate bool `json:"certify_brand_is_accurate,omitzero" api:"required"`
	// Must be `true`.
	//
	// Any of true.
	CertifyIPOwnership bool `json:"certify_ip_ownership,omitzero" api:"required"`
	// Must be `true`.
	//
	// Any of true.
	CertifyNoInfringement bool `json:"certify_no_infringement,omitzero" api:"required"`
	// Must be `true`.
	//
	// Any of true.
	CertifyNoShaftContent bool `json:"certify_no_shaft_content,omitzero" api:"required"`
	// Explanation of how the infringement concern was addressed.
	InfringementResolutionNotes string            `json:"infringement_resolution_notes" api:"required"`
	DisplayName                 param.Opt[string] `json:"display_name,omitzero"`
	// Publicly accessible HTTPS URL (max 128 chars) to a 256x256 BMP logo (max 1 MB).
	LogoURL     param.Opt[string] `json:"logo_url,omitzero"`
	CallReasons []string          `json:"call_reasons,omitzero"`
	// Append-only supporting documents.
	Documents []DocumentParam `json:"documents,omitzero"`
	paramObj
}

func (r DirUpdateInfringementParams) MarshalJSON() (data []byte, err error) {
	type shadow DirUpdateInfringementParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirUpdateInfringementParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
