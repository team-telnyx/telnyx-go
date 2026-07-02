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

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/pagination"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
	"github.com/stainless-sdks/telnyx-go/v4/packages/respjson"
)

// Associate phone numbers with a verified DIR so calls from those numbers carry
// the DIR's display identity.
//
// DirPhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirPhoneNumberService] method instead.
type DirPhoneNumberService struct {
	Options []option.RequestOption
}

// NewDirPhoneNumberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirPhoneNumberService(opts ...option.RequestOption) (r DirPhoneNumberService) {
	r = DirPhoneNumberService{}
	r.Options = opts
	return
}

// List the phone numbers registered under a DIR. The enterprise is resolved
// server-side from the DIR id.
func (r *DirPhoneNumberService) List(ctx context.Context, dirID string, query DirPhoneNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DirPhoneNumberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/phone_numbers", dirID)
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

// List the phone numbers registered under a DIR. The enterprise is resolved
// server-side from the DIR id.
func (r *DirPhoneNumberService) ListAutoPaging(ctx context.Context, dirID string, query DirPhoneNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DirPhoneNumberListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, dirID, query, opts...))
}

// Register phone numbers under a DIR. The enterprise is resolved server-side from
// the DIR id. Same body, failure modes, and batch semantics whichever path form
// you use.
//
// **Pricing:** This is a billable action. See https://telnyx.com/pricing/numbers
// for current pricing.
func (r *DirPhoneNumberService) Add(ctx context.Context, dirID string, body DirPhoneNumberAddParams, opts ...option.RequestOption) (res *DirPhoneNumberAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/phone_numbers", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deregister phone numbers from a DIR. The enterprise is resolved server-side from
// the DIR id. Returns a partial-success envelope.
func (r *DirPhoneNumberService) Remove(ctx context.Context, dirID string, body DirPhoneNumberRemoveParams, opts ...option.RequestOption) (res *DirPhoneNumberRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/phone_numbers", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type DirPhoneNumberListResponse struct {
	ID string `json:"id" format:"uuid"`
	// Id of the batch this number was vetted as part of.
	BatchID      string    `json:"batch_id" api:"nullable" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	DirID        string    `json:"dir_id" format:"uuid"`
	EnterpriseID string    `json:"enterprise_id" format:"uuid"`
	// Id of the Letter of Authorization document attached to this number's batch.
	LoaDocumentID string `json:"loa_document_id" api:"nullable" format:"uuid"`
	// E.164 with leading `+`.
	PhoneNumber string `json:"phone_number"`
	// Populated when `status` is `unsuccessful` or `permanently_rejected`.
	RejectionReason DirPhoneNumberListResponseRejectionReason `json:"rejection_reason" api:"nullable"`
	// Phone-number lifecycle status.
	//
	//   - `submitted` / `in_review` - Telnyx is reviewing the batch this number belongs
	//     to.
	//   - `verified` - approved; the DIR's display identity will be shown on outbound
	//     calls from this number.
	//   - `unsuccessful` - Telnyx rejected this submission; the customer may re-add to
	//     retry.
	//   - `suspended` - temporarily disabled (e.g. by an active infringement claim on
	//     the DIR).
	//   - `expired` - verification expired; re-add to renew.
	//   - `permanently_rejected` - terminal; cannot be re-added on this or any other DIR
	//     you own.
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	Status     DirPhoneNumberListResponseStatus `json:"status"`
	UpdatedAt  time.Time                        `json:"updated_at" format:"date-time"`
	VerifiedAt time.Time                        `json:"verified_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BatchID         respjson.Field
		CreatedAt       respjson.Field
		DirID           respjson.Field
		EnterpriseID    respjson.Field
		LoaDocumentID   respjson.Field
		PhoneNumber     respjson.Field
		RejectionReason respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		VerifiedAt      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Populated when `status` is `unsuccessful` or `permanently_rejected`.
type DirPhoneNumberListResponseRejectionReason struct {
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
func (r DirPhoneNumberListResponseRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberListResponseRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Phone-number lifecycle status.
//
//   - `submitted` / `in_review` - Telnyx is reviewing the batch this number belongs
//     to.
//   - `verified` - approved; the DIR's display identity will be shown on outbound
//     calls from this number.
//   - `unsuccessful` - Telnyx rejected this submission; the customer may re-add to
//     retry.
//   - `suspended` - temporarily disabled (e.g. by an active infringement claim on
//     the DIR).
//   - `expired` - verification expired; re-add to renew.
//   - `permanently_rejected` - terminal; cannot be re-added on this or any other DIR
//     you own.
type DirPhoneNumberListResponseStatus string

const (
	DirPhoneNumberListResponseStatusSubmitted           DirPhoneNumberListResponseStatus = "submitted"
	DirPhoneNumberListResponseStatusInReview            DirPhoneNumberListResponseStatus = "in_review"
	DirPhoneNumberListResponseStatusVerified            DirPhoneNumberListResponseStatus = "verified"
	DirPhoneNumberListResponseStatusUnsuccessful        DirPhoneNumberListResponseStatus = "unsuccessful"
	DirPhoneNumberListResponseStatusSuspended           DirPhoneNumberListResponseStatus = "suspended"
	DirPhoneNumberListResponseStatusExpired             DirPhoneNumberListResponseStatus = "expired"
	DirPhoneNumberListResponseStatusPermanentlyRejected DirPhoneNumberListResponseStatus = "permanently_rejected"
)

// Bulk-add success response (HTTP 201). All numbers in the request were accepted
// into a single new batch. Every entry in `data` shares the same `batch_id` - read
// it from any element to obtain the batch id for subsequent
// `GET .../phone_number_batches/{batch_id}` calls. If any number in the request
// fails (schema-invalid, not in inventory, already attached to another DIR, etc.)
// the entire request is rejected with HTTP 400 and the canonical Telnyx error
// envelope; the success body described here is therefore an all-or-nothing
// payload.
type DirPhoneNumberAddResponse struct {
	// Phone numbers accepted into the new batch. List order mirrors the request order.
	// Each element shares the same `batch_id`.
	Data []DirPhoneNumberAddResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberAddResponse) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberAddResponseData struct {
	ID string `json:"id" format:"uuid"`
	// Id of the batch this number was vetted as part of.
	BatchID      string    `json:"batch_id" api:"nullable" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	DirID        string    `json:"dir_id" format:"uuid"`
	EnterpriseID string    `json:"enterprise_id" format:"uuid"`
	// Id of the Letter of Authorization document attached to this number's batch.
	LoaDocumentID string `json:"loa_document_id" api:"nullable" format:"uuid"`
	// E.164 with leading `+`.
	PhoneNumber string `json:"phone_number"`
	// Populated when `status` is `unsuccessful` or `permanently_rejected`.
	RejectionReason DirPhoneNumberAddResponseDataRejectionReason `json:"rejection_reason" api:"nullable"`
	// Phone-number lifecycle status.
	//
	//   - `submitted` / `in_review` - Telnyx is reviewing the batch this number belongs
	//     to.
	//   - `verified` - approved; the DIR's display identity will be shown on outbound
	//     calls from this number.
	//   - `unsuccessful` - Telnyx rejected this submission; the customer may re-add to
	//     retry.
	//   - `suspended` - temporarily disabled (e.g. by an active infringement claim on
	//     the DIR).
	//   - `expired` - verification expired; re-add to renew.
	//   - `permanently_rejected` - terminal; cannot be re-added on this or any other DIR
	//     you own.
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	Status     string    `json:"status"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	VerifiedAt time.Time `json:"verified_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BatchID         respjson.Field
		CreatedAt       respjson.Field
		DirID           respjson.Field
		EnterpriseID    respjson.Field
		LoaDocumentID   respjson.Field
		PhoneNumber     respjson.Field
		RejectionReason respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		VerifiedAt      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberAddResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberAddResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Populated when `status` is `unsuccessful` or `permanently_rejected`.
type DirPhoneNumberAddResponseDataRejectionReason struct {
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
func (r DirPhoneNumberAddResponseDataRejectionReason) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberAddResponseDataRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bulk-delete partial-success response. `data` is the list of phone numbers that
// were soft-deleted. `meta.errors` holds per-number failures (e.g. number not
// associated with this DIR). When EVERY number in the request fails, the endpoint
// instead returns 400 with the canonical Telnyx error envelope and `data`/`meta`
// are absent.
type DirPhoneNumberRemoveResponse struct {
	// Phone numbers that were successfully soft-deleted. Bare E.164 strings.
	Data []string                         `json:"data" api:"required"`
	Meta DirPhoneNumberRemoveResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberRemoveResponseMeta struct {
	// Per-number failures that did not block the call. Each entry has `phone_number`,
	// `code`, `title`, `detail`.
	Errors []DirPhoneNumberRemoveResponseMetaError `json:"errors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberRemoveResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberRemoveResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-number error returned by the bulk-delete endpoint. Bulk-add does not use
// this shape - it returns a 400 with the canonical envelope grouping numbers by
// failure category.
type DirPhoneNumberRemoveResponseMetaError struct {
	// Stable per-number error code. Currently only `not_associated` is emitted, when
	// the number is not attached to this DIR.
	//
	// Any of "not_associated".
	Code        string `json:"code" api:"required"`
	Detail      string `json:"detail" api:"required"`
	PhoneNumber string `json:"phone_number" api:"required"`
	Title       string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		PhoneNumber respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberRemoveResponseMetaError) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberRemoveResponseMetaError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberListParams struct {
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by phone-number status.
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	Status DirPhoneNumberListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirPhoneNumberListParams]'s query parameters as
// `url.Values`.
func (r DirPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by phone-number status.
type DirPhoneNumberListParamsStatus string

const (
	DirPhoneNumberListParamsStatusSubmitted           DirPhoneNumberListParamsStatus = "submitted"
	DirPhoneNumberListParamsStatusInReview            DirPhoneNumberListParamsStatus = "in_review"
	DirPhoneNumberListParamsStatusVerified            DirPhoneNumberListParamsStatus = "verified"
	DirPhoneNumberListParamsStatusUnsuccessful        DirPhoneNumberListParamsStatus = "unsuccessful"
	DirPhoneNumberListParamsStatusSuspended           DirPhoneNumberListParamsStatus = "suspended"
	DirPhoneNumberListParamsStatusExpired             DirPhoneNumberListParamsStatus = "expired"
	DirPhoneNumberListParamsStatusPermanentlyRejected DirPhoneNumberListParamsStatus = "permanently_rejected"
)

type DirPhoneNumberAddParams struct {
	// Supporting documents covering this batch. At least one entry with
	// `document_type: letter_of_authorization` is required - the LOA authorises Telnyx
	// to register these numbers under the DIR. Each `document_id` must come from the
	// Telnyx Documents API. Additional document types (e.g. business registration) may
	// be included alongside the LOA.
	Documents []DirPhoneNumberAddParamsDocument `json:"documents,omitzero" api:"required"`
	// 1–15 phone numbers in E.164 format. 10-digit US numbers are auto-prefixed with
	// `1`.
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r DirPhoneNumberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow DirPhoneNumberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirPhoneNumberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DocumentID, DocumentType are required.
type DirPhoneNumberAddParamsDocument struct {
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

func (r DirPhoneNumberAddParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow DirPhoneNumberAddParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirPhoneNumberAddParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DirPhoneNumberAddParamsDocument](
		"document_type", "letter_of_authorization", "business_registration", "articles_of_incorporation", "tax_document", "ein_letter", "trademark_registration", "website_ownership", "business_license", "professional_license", "government_id", "utility_bill", "bank_statement", "other",
	)
}

type DirPhoneNumberRemoveParams struct {
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r DirPhoneNumberRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow DirPhoneNumberRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirPhoneNumberRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
