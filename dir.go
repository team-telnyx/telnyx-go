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
	return
}

// Returns a single DIR by id. The enterprise is resolved server-side from the DIR
// id. Returns `404` if the DIR does not exist or is not yours.
func (r *DirService) Get(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Edit a DIR. Only DIRs in `draft`, `rejected`, `unsuccessful`, or `suspended` are
// editable. PATCH is a pure edit — `status` is never changed by this endpoint. To
// re-vet after editing, call `POST /v2/dir/{dir_id}/submit` explicitly.
func (r *DirService) Update(ctx context.Context, dirID string, body DirUpdateParams, opts ...option.RequestOption) (res *DirUpdateResponse, err error) {
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
func (r *DirService) List(ctx context.Context, query DirListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DirListResponse], err error) {
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
func (r *DirService) ListAutoPaging(ctx context.Context, query DirListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DirListResponse] {
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
func (r *DirService) ListInfringementClaims(ctx context.Context, dirID string, query DirListInfringementClaimsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DirListInfringementClaimsResponse], err error) {
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
func (r *DirService) ListInfringementClaimsAutoPaging(ctx context.Context, dirID string, query DirListInfringementClaimsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DirListInfringementClaimsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListInfringementClaims(ctx, dirID, query, opts...))
}

// Submit a DIR for vetting. Sends the DIR back through the vetting cycle from any
// non-terminal status. When re-submitting from `suspended` or `expired`, the DIR's
// previous Branded Calling registration is torn down transactionally and its phone
// numbers flip back to `submitted`. When re-submitting from `verified`, the
// existing registration stays live throughout the new vetting cycle.
//
// Returns `400` from `submitted`/`in_review`/`permanently_rejected`. Returns `409`
// if the DIR has an unresolved infringement claim.
func (r *DirService) Submit(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirSubmitResponse, err error) {
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
func (r *DirService) UpdateInfringement(ctx context.Context, dirID string, body DirUpdateInfringementParams, opts ...option.RequestOption) (res *DirUpdateInfringementResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/infringement_update", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type DirGetResponse struct {
	Data DirGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DirGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirGetResponseData struct {
	ID                     string                         `json:"id" format:"uuid"`
	AuthorizerEmail        string                         `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                         `json:"authorizer_name" api:"nullable"`
	CallReasons            []DirGetResponseDataCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                           `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                           `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                           `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                      `json:"created_at" format:"date-time"`
	DisplayName            string                         `json:"display_name"`
	Documents              []DirGetResponseDataDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                         `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                      `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                         `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                      `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []DirGetResponseDataRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                                `json:"reselling"`
	// DIR lifecycle status.
	//
	//   - `draft` — newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` — Telnyx is reviewing.
	//   - `verified` — approved; phone numbers may be attached.
	//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` — system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` — verification expired; customer must resubmit.
	//   - `infringement_claimed` — a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` — terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status      string    `json:"status"`
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
func (r DirGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirGetResponseDataCallReason struct {
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
func (r DirGetResponseDataCallReason) RawJSON() string { return r.JSON.raw }
func (r *DirGetResponseDataCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirGetResponseDataDocument struct {
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
	DocumentType string `json:"document_type" api:"required"`
	Description  string `json:"description"`
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
func (r DirGetResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *DirGetResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirGetResponseDataRejectionReason struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	// Customer-visible free-text comment from the Telnyx vetting team. Only the first
	// entry of `rejection_reasons` carries this; the rest are `null`.
	Message string `json:"message" api:"nullable"`
	Title   string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Message     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirGetResponseDataRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirGetResponseDataRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateResponse struct {
	Data DirUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateResponseData struct {
	ID                     string                            `json:"id" format:"uuid"`
	AuthorizerEmail        string                            `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                            `json:"authorizer_name" api:"nullable"`
	CallReasons            []DirUpdateResponseDataCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                              `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                              `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                              `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                         `json:"created_at" format:"date-time"`
	DisplayName            string                            `json:"display_name"`
	Documents              []DirUpdateResponseDataDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                            `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                         `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                            `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                         `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []DirUpdateResponseDataRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                                   `json:"reselling"`
	// DIR lifecycle status.
	//
	//   - `draft` — newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` — Telnyx is reviewing.
	//   - `verified` — approved; phone numbers may be attached.
	//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` — system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` — verification expired; customer must resubmit.
	//   - `infringement_claimed` — a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` — terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status      string    `json:"status"`
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
func (r DirUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateResponseDataCallReason struct {
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
func (r DirUpdateResponseDataCallReason) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateResponseDataCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateResponseDataDocument struct {
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
	DocumentType string `json:"document_type" api:"required"`
	Description  string `json:"description"`
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
func (r DirUpdateResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateResponseDataRejectionReason struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	// Customer-visible free-text comment from the Telnyx vetting team. Only the first
	// entry of `rejection_reasons` carries this; the rest are `null`.
	Message string `json:"message" api:"nullable"`
	Title   string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Message     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirUpdateResponseDataRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateResponseDataRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListResponse struct {
	ID                     string                      `json:"id" format:"uuid"`
	AuthorizerEmail        string                      `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                      `json:"authorizer_name" api:"nullable"`
	CallReasons            []DirListResponseCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                        `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                        `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                        `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                   `json:"created_at" format:"date-time"`
	DisplayName            string                      `json:"display_name"`
	Documents              []DirListResponseDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                      `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                   `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                      `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                   `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []DirListResponseRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                             `json:"reselling"`
	// DIR lifecycle status.
	//
	//   - `draft` — newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` — Telnyx is reviewing.
	//   - `verified` — approved; phone numbers may be attached.
	//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` — system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` — verification expired; customer must resubmit.
	//   - `infringement_claimed` — a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` — terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status      DirListResponseStatus `json:"status"`
	SubmittedAt time.Time             `json:"submitted_at" api:"nullable" format:"date-time"`
	UpdatedAt   time.Time             `json:"updated_at" format:"date-time"`
	VerifiedAt  time.Time             `json:"verified_at" api:"nullable" format:"date-time"`
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
func (r DirListResponse) RawJSON() string { return r.JSON.raw }
func (r *DirListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListResponseCallReason struct {
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
func (r DirListResponseCallReason) RawJSON() string { return r.JSON.raw }
func (r *DirListResponseCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListResponseDocument struct {
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
	DocumentType string `json:"document_type" api:"required"`
	Description  string `json:"description"`
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
func (r DirListResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *DirListResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListResponseRejectionReason struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	// Customer-visible free-text comment from the Telnyx vetting team. Only the first
	// entry of `rejection_reasons` carries this; the rest are `null`.
	Message string `json:"message" api:"nullable"`
	Title   string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Message     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListResponseRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirListResponseRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DIR lifecycle status.
//
//   - `draft` — newly created; editable; not yet submitted.
//   - `submitted` / `in_review` — Telnyx is reviewing.
//   - `verified` — approved; phone numbers may be attached.
//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
//     populated; customer can edit and resubmit.
//   - `unsuccessful` — system-side error during processing; customer can edit and
//     resubmit.
//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
//   - `expired` — verification expired; customer must resubmit.
//   - `infringement_claimed` — a trademark/impersonation claim is open against this
//     DIR.
//   - `permanently_rejected` — terminal; cannot be resubmitted.
type DirListResponseStatus string

const (
	DirListResponseStatusDraft               DirListResponseStatus = "draft"
	DirListResponseStatusSubmitted           DirListResponseStatus = "submitted"
	DirListResponseStatusInReview            DirListResponseStatus = "in_review"
	DirListResponseStatusVerified            DirListResponseStatus = "verified"
	DirListResponseStatusRejected            DirListResponseStatus = "rejected"
	DirListResponseStatusUnsuccessful        DirListResponseStatus = "unsuccessful"
	DirListResponseStatusSuspended           DirListResponseStatus = "suspended"
	DirListResponseStatusExpired             DirListResponseStatus = "expired"
	DirListResponseStatusInfringementClaimed DirListResponseStatus = "infringement_claimed"
	DirListResponseStatusPermanentlyRejected DirListResponseStatus = "permanently_rejected"
)

type DirListDocumentTypesResponse struct {
	Data []DirListDocumentTypesResponseData `json:"data" api:"required"`
	// JSON:API pagination metadata returned with every paginated list response. Page
	// numbering is 1-based. `page_size` reports the number of items actually returned
	// in `data` for this page; the requested size is taken from the `page[size]` query
	// parameter.
	Meta DirListDocumentTypesResponseMeta `json:"meta" api:"required"`
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

// JSON:API pagination metadata returned with every paginated list response. Page
// numbering is 1-based. `page_size` reports the number of items actually returned
// in `data` for this page; the requested size is taken from the `page[size]` query
// parameter.
type DirListDocumentTypesResponseMeta struct {
	// 1-based index of this page. Echoes the `page[number]` query parameter (default
	// `1`).
	PageNumber int64 `json:"page_number" api:"required"`
	// Number of items returned in this page's `data` array. Capped at 250.
	PageSize int64 `json:"page_size" api:"required"`
	// Total number of pages available given the current `page_size`.
	TotalPages int64 `json:"total_pages" api:"required"`
	// Total number of items across all pages (excludes soft-deleted rows).
	TotalResults int64 `json:"total_results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListDocumentTypesResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *DirListDocumentTypesResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirListInfringementClaimsResponse struct {
	ID string `json:"id" format:"uuid"`
	// When the claim was filed (set by the claimant's representative at file time).
	ClaimDate        time.Time `json:"claim_date" format:"date-time"`
	ClaimDescription string    `json:"claim_description"`
	// Category of infringement being claimed.
	//
	// Any of "trademark", "copyright".
	ClaimType       DirListInfringementClaimsResponseClaimType `json:"claim_type"`
	ClaimantContact string                                     `json:"claimant_contact"`
	ClaimantName    string                                     `json:"claimant_name"`
	// Aggregated across all customer contest submissions on this claim.
	ContestDocuments []DirListInfringementClaimsResponseContestDocument `json:"contest_documents"`
	// Per-round submission audit trail. Each entry records one
	// `POST /infringement_claims/{id}/contest` call (notes, timestamp, document
	// count). Aggregated documents live on `contest_documents`.
	ContestHistory []DirListInfringementClaimsResponseContestHistory `json:"contest_history"`
	CreatedAt      time.Time                                         `json:"created_at" format:"date-time"`
	// Snapshot of the DIR the claim is filed against, embedded for convenience.
	Dir          DirListInfringementClaimsResponseDir `json:"dir"`
	DirID        string                               `json:"dir_id" format:"uuid"`
	EnterpriseID string                               `json:"enterprise_id" format:"uuid"`
	// Set only when `status` is `resolved`.
	//
	// Any of "upheld", "rejected", "modified".
	Resolution      DirListInfringementClaimsResponseResolution `json:"resolution" api:"nullable"`
	ResolutionDate  time.Time                                   `json:"resolution_date" api:"nullable" format:"date-time"`
	ResolutionNotes string                                      `json:"resolution_notes" api:"nullable"`
	// Lifecycle status. `pending` — newly filed; the DIR is auto-suspended.
	// `contested` — you have submitted contest evidence; awaiting Telnyx review.
	// `resolved` — final.
	//
	// Any of "pending", "contested", "resolved".
	Status    DirListInfringementClaimsResponseStatus `json:"status"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ClaimDate        respjson.Field
		ClaimDescription respjson.Field
		ClaimType        respjson.Field
		ClaimantContact  respjson.Field
		ClaimantName     respjson.Field
		ContestDocuments respjson.Field
		ContestHistory   respjson.Field
		CreatedAt        respjson.Field
		Dir              respjson.Field
		DirID            respjson.Field
		EnterpriseID     respjson.Field
		Resolution       respjson.Field
		ResolutionDate   respjson.Field
		ResolutionNotes  respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListInfringementClaimsResponse) RawJSON() string { return r.JSON.raw }
func (r *DirListInfringementClaimsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category of infringement being claimed.
type DirListInfringementClaimsResponseClaimType string

const (
	DirListInfringementClaimsResponseClaimTypeTrademark DirListInfringementClaimsResponseClaimType = "trademark"
	DirListInfringementClaimsResponseClaimTypeCopyright DirListInfringementClaimsResponseClaimType = "copyright"
)

type DirListInfringementClaimsResponseContestDocument struct {
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
	DocumentType string `json:"document_type" api:"required"`
	Description  string `json:"description"`
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
func (r DirListInfringementClaimsResponseContestDocument) RawJSON() string { return r.JSON.raw }
func (r *DirListInfringementClaimsResponseContestDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One round of customer contest evidence on an infringement claim. The aggregated
// documents across rounds live on the parent claim's `contest_documents`; this
// submission record carries only the per-round notes and document count.
type DirListInfringementClaimsResponseContestHistory struct {
	DocumentCount int64     `json:"document_count"`
	Notes         string    `json:"notes"`
	SubmittedAt   time.Time `json:"submitted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentCount respjson.Field
		Notes         respjson.Field
		SubmittedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListInfringementClaimsResponseContestHistory) RawJSON() string { return r.JSON.raw }
func (r *DirListInfringementClaimsResponseContestHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot of the DIR the claim is filed against, embedded for convenience.
type DirListInfringementClaimsResponseDir struct {
	ID           string `json:"id" format:"uuid"`
	DisplayName  string `json:"display_name"`
	EnterpriseID string `json:"enterprise_id" format:"uuid"`
	// DIR lifecycle status.
	//
	//   - `draft` — newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` — Telnyx is reviewing.
	//   - `verified` — approved; phone numbers may be attached.
	//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` — system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` — verification expired; customer must resubmit.
	//   - `infringement_claimed` — a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` — terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayName  respjson.Field
		EnterpriseID respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirListInfringementClaimsResponseDir) RawJSON() string { return r.JSON.raw }
func (r *DirListInfringementClaimsResponseDir) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set only when `status` is `resolved`.
type DirListInfringementClaimsResponseResolution string

const (
	DirListInfringementClaimsResponseResolutionUpheld   DirListInfringementClaimsResponseResolution = "upheld"
	DirListInfringementClaimsResponseResolutionRejected DirListInfringementClaimsResponseResolution = "rejected"
	DirListInfringementClaimsResponseResolutionModified DirListInfringementClaimsResponseResolution = "modified"
)

// Lifecycle status. `pending` — newly filed; the DIR is auto-suspended.
// `contested` — you have submitted contest evidence; awaiting Telnyx review.
// `resolved` — final.
type DirListInfringementClaimsResponseStatus string

const (
	DirListInfringementClaimsResponseStatusPending   DirListInfringementClaimsResponseStatus = "pending"
	DirListInfringementClaimsResponseStatusContested DirListInfringementClaimsResponseStatus = "contested"
	DirListInfringementClaimsResponseStatusResolved  DirListInfringementClaimsResponseStatus = "resolved"
)

type DirSubmitResponse struct {
	Data DirSubmitResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *DirSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirSubmitResponseData struct {
	ID                     string                            `json:"id" format:"uuid"`
	AuthorizerEmail        string                            `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                            `json:"authorizer_name" api:"nullable"`
	CallReasons            []DirSubmitResponseDataCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                              `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                              `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                              `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                         `json:"created_at" format:"date-time"`
	DisplayName            string                            `json:"display_name"`
	Documents              []DirSubmitResponseDataDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                            `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                         `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                            `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                         `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []DirSubmitResponseDataRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                                   `json:"reselling"`
	// DIR lifecycle status.
	//
	//   - `draft` — newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` — Telnyx is reviewing.
	//   - `verified` — approved; phone numbers may be attached.
	//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` — system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` — verification expired; customer must resubmit.
	//   - `infringement_claimed` — a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` — terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status      string    `json:"status"`
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
func (r DirSubmitResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirSubmitResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirSubmitResponseDataCallReason struct {
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
func (r DirSubmitResponseDataCallReason) RawJSON() string { return r.JSON.raw }
func (r *DirSubmitResponseDataCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirSubmitResponseDataDocument struct {
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
	DocumentType string `json:"document_type" api:"required"`
	Description  string `json:"description"`
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
func (r DirSubmitResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *DirSubmitResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirSubmitResponseDataRejectionReason struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	// Customer-visible free-text comment from the Telnyx vetting team. Only the first
	// entry of `rejection_reasons` carries this; the rest are `null`.
	Message string `json:"message" api:"nullable"`
	Title   string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Message     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirSubmitResponseDataRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirSubmitResponseDataRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateInfringementResponse struct {
	Data DirUpdateInfringementResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirUpdateInfringementResponse) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateInfringementResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateInfringementResponseData struct {
	ID                     string                                        `json:"id" format:"uuid"`
	AuthorizerEmail        string                                        `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                                        `json:"authorizer_name" api:"nullable"`
	CallReasons            []DirUpdateInfringementResponseDataCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                                          `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                                          `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                                          `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                                     `json:"created_at" format:"date-time"`
	DisplayName            string                                        `json:"display_name"`
	Documents              []DirUpdateInfringementResponseDataDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                                        `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                                     `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                                        `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                                     `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []DirUpdateInfringementResponseDataRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                                               `json:"reselling"`
	// DIR lifecycle status.
	//
	//   - `draft` — newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` — Telnyx is reviewing.
	//   - `verified` — approved; phone numbers may be attached.
	//   - `rejected` — Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` — system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` — temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` — verification expired; customer must resubmit.
	//   - `infringement_claimed` — a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` — terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status      string    `json:"status"`
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
func (r DirUpdateInfringementResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateInfringementResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateInfringementResponseDataCallReason struct {
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
func (r DirUpdateInfringementResponseDataCallReason) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateInfringementResponseDataCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateInfringementResponseDataDocument struct {
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
	DocumentType string `json:"document_type" api:"required"`
	Description  string `json:"description"`
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
func (r DirUpdateInfringementResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateInfringementResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateInfringementResponseDataRejectionReason struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	// Customer-visible free-text comment from the Telnyx vetting team. Only the first
	// entry of `rejection_reasons` carries this; the rest are `null`.
	Message string `json:"message" api:"nullable"`
	Title   string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Message     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirUpdateInfringementResponseDataRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirUpdateInfringementResponseDataRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirUpdateParams struct {
	// Contact email of the authorizer. Telnyx may send verification or infringement
	// notices here.
	AuthorizerEmail param.Opt[string] `json:"authorizer_email,omitzero" format:"email"`
	// Name of the person at your enterprise authorizing this DIR. Must be a real
	// individual.
	AuthorizerName param.Opt[string] `json:"authorizer_name,omitzero"`
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
	// Restrict results to a single enterprise.
	EnterpriseID param.Opt[string] `query:"enterprise_id,omitzero" format:"uuid" json:"-"`
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
	// Case-insensitive partial match on `display_name` or call reason.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by DIR status.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	FilterStatus DirListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	// Sort field. Allowed values: `created_at`, `updated_at`, `display_name`,
	// `status`. Prefix with `-` for descending. Default `-created_at`.
	//
	// Any of "created_at", "-created_at", "updated_at", "-updated_at", "display_name",
	// "-display_name", "status", "-status".
	Sort DirListParamsSort `query:"sort,omitzero" json:"-"`
	// Filter by DIR status.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status DirListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirListParams]'s query parameters as `url.Values`.
func (r DirListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by DIR status.
type DirListParamsFilterStatus string

const (
	DirListParamsFilterStatusDraft               DirListParamsFilterStatus = "draft"
	DirListParamsFilterStatusSubmitted           DirListParamsFilterStatus = "submitted"
	DirListParamsFilterStatusInReview            DirListParamsFilterStatus = "in_review"
	DirListParamsFilterStatusVerified            DirListParamsFilterStatus = "verified"
	DirListParamsFilterStatusRejected            DirListParamsFilterStatus = "rejected"
	DirListParamsFilterStatusUnsuccessful        DirListParamsFilterStatus = "unsuccessful"
	DirListParamsFilterStatusSuspended           DirListParamsFilterStatus = "suspended"
	DirListParamsFilterStatusExpired             DirListParamsFilterStatus = "expired"
	DirListParamsFilterStatusInfringementClaimed DirListParamsFilterStatus = "infringement_claimed"
	DirListParamsFilterStatusPermanentlyRejected DirListParamsFilterStatus = "permanently_rejected"
)

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

// Filter by DIR status.
type DirListParamsStatus string

const (
	DirListParamsStatusDraft               DirListParamsStatus = "draft"
	DirListParamsStatusSubmitted           DirListParamsStatus = "submitted"
	DirListParamsStatusInReview            DirListParamsStatus = "in_review"
	DirListParamsStatusVerified            DirListParamsStatus = "verified"
	DirListParamsStatusRejected            DirListParamsStatus = "rejected"
	DirListParamsStatusUnsuccessful        DirListParamsStatus = "unsuccessful"
	DirListParamsStatusSuspended           DirListParamsStatus = "suspended"
	DirListParamsStatusExpired             DirListParamsStatus = "expired"
	DirListParamsStatusInfringementClaimed DirListParamsStatus = "infringement_claimed"
	DirListParamsStatusPermanentlyRejected DirListParamsStatus = "permanently_rejected"
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
	Documents []DirUpdateInfringementParamsDocument `json:"documents,omitzero"`
	paramObj
}

func (r DirUpdateInfringementParams) MarshalJSON() (data []byte, err error) {
	type shadow DirUpdateInfringementParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirUpdateInfringementParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DocumentID, DocumentType are required.
type DirUpdateInfringementParamsDocument struct {
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
	DocumentType string            `json:"document_type,omitzero" api:"required"`
	Description  param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r DirUpdateInfringementParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow DirUpdateInfringementParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirUpdateInfringementParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DirUpdateInfringementParamsDocument](
		"document_type", "letter_of_authorization", "business_registration", "articles_of_incorporation", "tax_document", "ein_letter", "trademark_registration", "website_ownership", "business_license", "professional_license", "government_id", "utility_bill", "bank_statement", "other",
	)
}
