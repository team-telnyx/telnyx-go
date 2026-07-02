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

// A Display Identity Record (DIR) is the verified calling identity (display name,
// logo, call reasons) shown to recipients on outbound calls.
//
// EnterpriseDirService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseDirService] method instead.
type EnterpriseDirService struct {
	Options []option.RequestOption
}

// NewEnterpriseDirService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnterpriseDirService(opts ...option.RequestOption) (r EnterpriseDirService) {
	r = EnterpriseDirService{}
	r.Options = opts
	return
}

// Create a new DIR under the given enterprise. The DIR starts in `draft` status;
// it must be submitted (`POST .../submit`) and approved by Telnyx before any phone
// number can be attached.
//
// **Field rules**
//
//   - `display_name`: 1–35 characters, no emoji or whitespace-only strings; this is
//     the name shown to recipients.
//   - `call_reasons`: 1–10 strings, each ≤64 characters; describe why your business
//     calls customers (e.g. 'Appointment reminders', 'Billing inquiries'). Validate
//     the wording against `POST /call_reasons/validate`.
//   - `logo_url`: HTTPS URL (max 128 chars) to a 256×256 BMP (max 1 MB). The image
//     is downloaded and hashed at submission time.
//   - `documents`: up to 20 entries; each `document_id` must be obtained by
//     uploading the file via the Telnyx Documents API first. Within one DIR a
//     `document_id` may only appear once.
//   - `certify_brand_is_accurate`, `certify_no_shaft_content`,
//     `certify_ip_ownership` MUST all be `true`.
//
// **Failure modes**
//
//   - `422` - validation error; `errors[].source.pointer` names the offending field.
//   - `403` - Branded Calling not activated on this enterprise (see
//     `POST /enterprises/{id}/branded_calling`).
//   - `404` - enterprise does not exist or does not belong to your account.
func (r *EnterpriseDirService) New(ctx context.Context, enterpriseID string, body EnterpriseDirNewParams, opts ...option.RequestOption) (res *EnterpriseDirNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/dir", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return the DIRs (Display Identity Records) belonging to a single enterprise.
// Pagination is JSON:API style (`page[number]`, `page[size]`, max 250). Supports
// `filter[]` query params: `filter[status]`, `filter[display_name][contains]`,
// `filter[call_reason][contains]`, plus the renewal-window filters
// `filter[expiring_at][gte]` / `filter[expiring_at][lte]` and the convenience
// `filter[expiring_within_days]` (mutually exclusive with the explicit gte/lte
// form). Sortable by `created_at`, `updated_at`, `display_name`, `status`,
// `submitted_at`, `verified_at`, `expiring_at` (prefix `-` for descending; default
// `-created_at`).
func (r *EnterpriseDirService) List(ctx context.Context, enterpriseID string, query EnterpriseDirListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EnterpriseDirListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/dir", enterpriseID)
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

// Return the DIRs (Display Identity Records) belonging to a single enterprise.
// Pagination is JSON:API style (`page[number]`, `page[size]`, max 250). Supports
// `filter[]` query params: `filter[status]`, `filter[display_name][contains]`,
// `filter[call_reason][contains]`, plus the renewal-window filters
// `filter[expiring_at][gte]` / `filter[expiring_at][lte]` and the convenience
// `filter[expiring_within_days]` (mutually exclusive with the explicit gte/lte
// form). Sortable by `created_at`, `updated_at`, `display_name`, `status`,
// `submitted_at`, `verified_at`, `expiring_at` (prefix `-` for descending; default
// `-created_at`).
func (r *EnterpriseDirService) ListAutoPaging(ctx context.Context, enterpriseID string, query EnterpriseDirListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EnterpriseDirListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, enterpriseID, query, opts...))
}

type EnterpriseDirNewResponse struct {
	Data EnterpriseDirNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseDirNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirNewResponseData struct {
	ID                     string                                   `json:"id" format:"uuid"`
	AuthorizerEmail        string                                   `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                                   `json:"authorizer_name" api:"nullable"`
	CallReasons            []EnterpriseDirNewResponseDataCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                                     `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                                     `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                                     `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                                `json:"created_at" format:"date-time"`
	DisplayName            string                                   `json:"display_name"`
	Documents              []EnterpriseDirNewResponseDataDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                                   `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                                `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                                   `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                                `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []EnterpriseDirNewResponseDataRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                                          `json:"reselling"`
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
func (r EnterpriseDirNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirNewResponseDataCallReason struct {
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
func (r EnterpriseDirNewResponseDataCallReason) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirNewResponseDataCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirNewResponseDataDocument struct {
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
func (r EnterpriseDirNewResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirNewResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirNewResponseDataRejectionReason struct {
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
func (r EnterpriseDirNewResponseDataRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirNewResponseDataRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirListResponse struct {
	ID                     string                                `json:"id" format:"uuid"`
	AuthorizerEmail        string                                `json:"authorizer_email" api:"nullable" format:"email"`
	AuthorizerName         string                                `json:"authorizer_name" api:"nullable"`
	CallReasons            []EnterpriseDirListResponseCallReason `json:"call_reasons"`
	CertifyBrandIsAccurate bool                                  `json:"certify_brand_is_accurate"`
	CertifyIPOwnership     bool                                  `json:"certify_ip_ownership"`
	CertifyNoShaftContent  bool                                  `json:"certify_no_shaft_content"`
	CreatedAt              time.Time                             `json:"created_at" format:"date-time"`
	DisplayName            string                                `json:"display_name"`
	Documents              []EnterpriseDirListResponseDocument   `json:"documents" api:"nullable"`
	EnterpriseID           string                                `json:"enterprise_id" format:"uuid"`
	ExpiringAt             time.Time                             `json:"expiring_at" api:"nullable" format:"date-time"`
	LogoURL                string                                `json:"logo_url" api:"nullable" format:"uri"`
	RejectedAt             time.Time                             `json:"rejected_at" api:"nullable" format:"date-time"`
	// Populated when `status` is `rejected`; cleared on `/submit` or successful
	// approval.
	RejectionReasons []EnterpriseDirListResponseRejectionReason `json:"rejection_reasons" api:"nullable"`
	Reselling        bool                                       `json:"reselling"`
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
	Status      EnterpriseDirListResponseStatus `json:"status"`
	SubmittedAt time.Time                       `json:"submitted_at" api:"nullable" format:"date-time"`
	UpdatedAt   time.Time                       `json:"updated_at" format:"date-time"`
	VerifiedAt  time.Time                       `json:"verified_at" api:"nullable" format:"date-time"`
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
func (r EnterpriseDirListResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirListResponseCallReason struct {
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
func (r EnterpriseDirListResponseCallReason) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirListResponseCallReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirListResponseDocument struct {
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
func (r EnterpriseDirListResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirListResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseDirListResponseRejectionReason struct {
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
func (r EnterpriseDirListResponseRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseDirListResponseRejectionReason) UnmarshalJSON(data []byte) error {
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
type EnterpriseDirListResponseStatus string

const (
	EnterpriseDirListResponseStatusDraft               EnterpriseDirListResponseStatus = "draft"
	EnterpriseDirListResponseStatusSubmitted           EnterpriseDirListResponseStatus = "submitted"
	EnterpriseDirListResponseStatusInReview            EnterpriseDirListResponseStatus = "in_review"
	EnterpriseDirListResponseStatusVerified            EnterpriseDirListResponseStatus = "verified"
	EnterpriseDirListResponseStatusRejected            EnterpriseDirListResponseStatus = "rejected"
	EnterpriseDirListResponseStatusUnsuccessful        EnterpriseDirListResponseStatus = "unsuccessful"
	EnterpriseDirListResponseStatusSuspended           EnterpriseDirListResponseStatus = "suspended"
	EnterpriseDirListResponseStatusExpired             EnterpriseDirListResponseStatus = "expired"
	EnterpriseDirListResponseStatusInfringementClaimed EnterpriseDirListResponseStatus = "infringement_claimed"
	EnterpriseDirListResponseStatusPermanentlyRejected EnterpriseDirListResponseStatus = "permanently_rejected"
)

type EnterpriseDirNewParams struct {
	// Contact email of the authorizer. Telnyx may send verification or
	// infringement-notice email here; use a monitored mailbox.
	AuthorizerEmail string `json:"authorizer_email" api:"required" format:"email"`
	// Name of the person at your enterprise who is authorizing this DIR registration.
	// Must be a real individual (used for audit and trademark-claim contests).
	AuthorizerName string `json:"authorizer_name" api:"required"`
	// 1–10 reasons your business calls customers. Validate phrasing against
	// `POST /call_reasons/validate`.
	CallReasons []string `json:"call_reasons,omitzero" api:"required"`
	// Must be `true`.
	//
	// Any of true.
	CertifyBrandIsAccurate bool `json:"certify_brand_is_accurate,omitzero" api:"required"`
	// Must be `true`. Confirms ownership of any logos/trademarks shown.
	//
	// Any of true.
	CertifyIPOwnership bool `json:"certify_ip_ownership,omitzero" api:"required"`
	// Must be `true`. Confirms this DIR is not used for SHAFT content (Sex, Hate,
	// Alcohol, Firearms, Tobacco) where prohibited.
	//
	// Any of true.
	CertifyNoShaftContent bool `json:"certify_no_shaft_content,omitzero" api:"required"`
	// Name shown to call recipients. No emoji; not whitespace-only.
	DisplayName string `json:"display_name" api:"required"`
	// Publicly accessible HTTPS URL (max 128 chars) to a 256x256 BMP logo (max 1 MB).
	LogoURL param.Opt[string] `json:"logo_url,omitzero" format:"uri"`
	// Set to true if your organization places calls on behalf of other enterprises
	// (BPO/reseller).
	Reselling param.Opt[bool] `json:"reselling,omitzero"`
	// Supporting documents. Each `document_id` may appear at most once on a DIR.
	Documents []EnterpriseDirNewParamsDocument `json:"documents,omitzero"`
	paramObj
}

func (r EnterpriseDirNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseDirNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseDirNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DocumentID, DocumentType are required.
type EnterpriseDirNewParamsDocument struct {
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

func (r EnterpriseDirNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseDirNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseDirNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EnterpriseDirNewParamsDocument](
		"document_type", "letter_of_authorization", "business_registration", "articles_of_incorporation", "tax_document", "ein_letter", "trademark_registration", "website_ownership", "business_license", "professional_license", "government_id", "utility_bill", "bank_statement", "other",
	)
}

type EnterpriseDirListParams struct {
	// Case-insensitive partial match on call reason.
	FilterCallReasonContains param.Opt[string] `query:"filter[call_reason][contains],omitzero" json:"-"`
	// Case-insensitive partial match on display name.
	FilterDisplayNameContains param.Opt[string] `query:"filter[display_name][contains],omitzero" json:"-"`
	// Return only DIRs whose `expiring_at` is at or after this ISO-8601 timestamp.
	FilterExpiringAtGte param.Opt[time.Time] `query:"filter[expiring_at][gte],omitzero" format:"date-time" json:"-"`
	// Return only DIRs whose `expiring_at` is at or before this ISO-8601 timestamp.
	FilterExpiringAtLte param.Opt[time.Time] `query:"filter[expiring_at][lte],omitzero" format:"date-time" json:"-"`
	// Convenience: returns DIRs whose `expiring_at` falls within the next N days
	// (1–365). Equivalent to setting `filter[expiring_at][gte]=<now>` +
	// `filter[expiring_at][lte]=<now+N>`. Mutually exclusive with the explicit
	// `[gte]`/`[lte]` filters - combining returns 400.
	FilterExpiringWithinDays param.Opt[int64] `query:"filter[expiring_within_days],omitzero" json:"-"`
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by DIR status.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	FilterStatus EnterpriseDirListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	// Sort field. Allowed: `created_at`, `updated_at`, `display_name`, `status`,
	// `submitted_at`, `verified_at`, `expiring_at`. Prefix with `-` for descending.
	// Default `-created_at`.
	//
	// Any of "created_at", "-created_at", "updated_at", "-updated_at", "display_name",
	// "-display_name", "status", "-status", "submitted_at", "-submitted_at",
	// "verified_at", "-verified_at", "expiring_at", "-expiring_at".
	Sort EnterpriseDirListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseDirListParams]'s query parameters as
// `url.Values`.
func (r EnterpriseDirListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by DIR status.
type EnterpriseDirListParamsFilterStatus string

const (
	EnterpriseDirListParamsFilterStatusDraft               EnterpriseDirListParamsFilterStatus = "draft"
	EnterpriseDirListParamsFilterStatusSubmitted           EnterpriseDirListParamsFilterStatus = "submitted"
	EnterpriseDirListParamsFilterStatusInReview            EnterpriseDirListParamsFilterStatus = "in_review"
	EnterpriseDirListParamsFilterStatusVerified            EnterpriseDirListParamsFilterStatus = "verified"
	EnterpriseDirListParamsFilterStatusRejected            EnterpriseDirListParamsFilterStatus = "rejected"
	EnterpriseDirListParamsFilterStatusUnsuccessful        EnterpriseDirListParamsFilterStatus = "unsuccessful"
	EnterpriseDirListParamsFilterStatusSuspended           EnterpriseDirListParamsFilterStatus = "suspended"
	EnterpriseDirListParamsFilterStatusExpired             EnterpriseDirListParamsFilterStatus = "expired"
	EnterpriseDirListParamsFilterStatusInfringementClaimed EnterpriseDirListParamsFilterStatus = "infringement_claimed"
	EnterpriseDirListParamsFilterStatusPermanentlyRejected EnterpriseDirListParamsFilterStatus = "permanently_rejected"
)

// Sort field. Allowed: `created_at`, `updated_at`, `display_name`, `status`,
// `submitted_at`, `verified_at`, `expiring_at`. Prefix with `-` for descending.
// Default `-created_at`.
type EnterpriseDirListParamsSort string

const (
	EnterpriseDirListParamsSortCreatedAt        EnterpriseDirListParamsSort = "created_at"
	EnterpriseDirListParamsSortCreatedAtDesc    EnterpriseDirListParamsSort = "-created_at"
	EnterpriseDirListParamsSortUpdatedAt        EnterpriseDirListParamsSort = "updated_at"
	EnterpriseDirListParamsSortUpdatedAtDesc    EnterpriseDirListParamsSort = "-updated_at"
	EnterpriseDirListParamsSortDisplayName      EnterpriseDirListParamsSort = "display_name"
	EnterpriseDirListParamsSortMinusDisplayName EnterpriseDirListParamsSort = "-display_name"
	EnterpriseDirListParamsSortStatus           EnterpriseDirListParamsSort = "status"
	EnterpriseDirListParamsSortStatusDesc       EnterpriseDirListParamsSort = "-status"
	EnterpriseDirListParamsSortSubmittedAt      EnterpriseDirListParamsSort = "submitted_at"
	EnterpriseDirListParamsSortMinusSubmittedAt EnterpriseDirListParamsSort = "-submitted_at"
	EnterpriseDirListParamsSortVerifiedAt       EnterpriseDirListParamsSort = "verified_at"
	EnterpriseDirListParamsSortMinusVerifiedAt  EnterpriseDirListParamsSort = "-verified_at"
	EnterpriseDirListParamsSortExpiringAt       EnterpriseDirListParamsSort = "expiring_at"
	EnterpriseDirListParamsSortMinusExpiringAt  EnterpriseDirListParamsSort = "-expiring_at"
)
