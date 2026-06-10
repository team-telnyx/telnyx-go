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

// Phone numbers are submitted to Telnyx for vetting in batches. Batches group all
// numbers added in a single request under the same Letter of Authorization.
//
// DirPhoneNumberBatchService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirPhoneNumberBatchService] method instead.
type DirPhoneNumberBatchService struct {
	Options []option.RequestOption
}

// NewDirPhoneNumberBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDirPhoneNumberBatchService(opts ...option.RequestOption) (r DirPhoneNumberBatchService) {
	r = DirPhoneNumberBatchService{}
	r.Options = opts
	return
}

// Get a single phone-number batch by id. The enterprise is resolved server-side
// from the DIR id.
func (r *DirPhoneNumberBatchService) Get(ctx context.Context, batchID string, query DirPhoneNumberBatchGetParams, opts ...option.RequestOption) (res *DirPhoneNumberBatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.DirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/phone_number_batches/%s", query.DirID, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the phone-number batches submitted under a DIR. The enterprise is resolved
// server-side from the DIR id.
func (r *DirPhoneNumberBatchService) List(ctx context.Context, dirID string, query DirPhoneNumberBatchListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DirPhoneNumberBatchListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/phone_number_batches", dirID)
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

// List the phone-number batches submitted under a DIR. The enterprise is resolved
// server-side from the DIR id.
func (r *DirPhoneNumberBatchService) ListAutoPaging(ctx context.Context, dirID string, query DirPhoneNumberBatchListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DirPhoneNumberBatchListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, dirID, query, opts...))
}

type DirPhoneNumberBatchGetResponse struct {
	// A phone-number batch groups all numbers added in a single bulk-add request.
	// Telnyx vets the batch as a unit. The response embeds the full `phone_numbers`
	// array so you can read per-number status without a separate call, plus a
	// batch-level `status` summarising the unit's progress.
	Data DirPhoneNumberBatchGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberBatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phone-number batch groups all numbers added in a single bulk-add request.
// Telnyx vets the batch as a unit. The response embeds the full `phone_numbers`
// array so you can read per-number status without a separate call, plus a
// batch-level `status` summarising the unit's progress.
type DirPhoneNumberBatchGetResponseData struct {
	BatchID string `json:"batch_id" format:"uuid"`
	// The DIR's display name at the time the batch was read.
	DirDisplayName string `json:"dir_display_name"`
	DirID          string `json:"dir_id" format:"uuid"`
	// Documents attached to this batch (e.g. a Letter of Authorization). Empty when
	// none were supplied at add time.
	Documents    []DirPhoneNumberBatchGetResponseDataDocument `json:"documents"`
	EnterpriseID string                                       `json:"enterprise_id" format:"uuid"`
	// All phone numbers in this batch, with per-number status.
	PhoneNumbers []DirPhoneNumberBatchGetResponseDataPhoneNumber `json:"phone_numbers"`
	// Aggregate batch status. Mirrors the values used on individual phone numbers
	// (`submitted`, `in_review`, `verified`, `unsuccessful`, `permanently_rejected`,
	// etc.).
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	Status string `json:"status"`
	// When the batch was created (and implicitly submitted for vetting).
	SubmittedAt time.Time `json:"submitted_at" format:"date-time"`
	// Number of phone numbers in this batch (length of `phone_numbers`).
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchID        respjson.Field
		DirDisplayName respjson.Field
		DirID          respjson.Field
		Documents      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumbers   respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		TotalCount     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberBatchGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberBatchGetResponseDataDocument struct {
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
func (r DirPhoneNumberBatchGetResponseDataDocument) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchGetResponseDataDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberBatchGetResponseDataPhoneNumber struct {
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
	RejectionReason DirPhoneNumberBatchGetResponseDataPhoneNumberRejectionReason `json:"rejection_reason" api:"nullable"`
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
func (r DirPhoneNumberBatchGetResponseDataPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchGetResponseDataPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Populated when `status` is `unsuccessful` or `permanently_rejected`.
type DirPhoneNumberBatchGetResponseDataPhoneNumberRejectionReason struct {
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
func (r DirPhoneNumberBatchGetResponseDataPhoneNumberRejectionReason) RawJSON() string {
	return r.JSON.raw
}
func (r *DirPhoneNumberBatchGetResponseDataPhoneNumberRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phone-number batch groups all numbers added in a single bulk-add request.
// Telnyx vets the batch as a unit. The response embeds the full `phone_numbers`
// array so you can read per-number status without a separate call, plus a
// batch-level `status` summarising the unit's progress.
type DirPhoneNumberBatchListResponse struct {
	BatchID string `json:"batch_id" format:"uuid"`
	// The DIR's display name at the time the batch was read.
	DirDisplayName string `json:"dir_display_name"`
	DirID          string `json:"dir_id" format:"uuid"`
	// Documents attached to this batch (e.g. a Letter of Authorization). Empty when
	// none were supplied at add time.
	Documents    []DirPhoneNumberBatchListResponseDocument `json:"documents"`
	EnterpriseID string                                    `json:"enterprise_id" format:"uuid"`
	// All phone numbers in this batch, with per-number status.
	PhoneNumbers []DirPhoneNumberBatchListResponsePhoneNumber `json:"phone_numbers"`
	// Aggregate batch status. Mirrors the values used on individual phone numbers
	// (`submitted`, `in_review`, `verified`, `unsuccessful`, `permanently_rejected`,
	// etc.).
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	Status DirPhoneNumberBatchListResponseStatus `json:"status"`
	// When the batch was created (and implicitly submitted for vetting).
	SubmittedAt time.Time `json:"submitted_at" format:"date-time"`
	// Number of phone numbers in this batch (length of `phone_numbers`).
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchID        respjson.Field
		DirDisplayName respjson.Field
		DirID          respjson.Field
		Documents      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumbers   respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		TotalCount     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirPhoneNumberBatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberBatchListResponseDocument struct {
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
func (r DirPhoneNumberBatchListResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchListResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberBatchListResponsePhoneNumber struct {
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
	RejectionReason DirPhoneNumberBatchListResponsePhoneNumberRejectionReason `json:"rejection_reason" api:"nullable"`
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
func (r DirPhoneNumberBatchListResponsePhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *DirPhoneNumberBatchListResponsePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Populated when `status` is `unsuccessful` or `permanently_rejected`.
type DirPhoneNumberBatchListResponsePhoneNumberRejectionReason struct {
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
func (r DirPhoneNumberBatchListResponsePhoneNumberRejectionReason) RawJSON() string {
	return r.JSON.raw
}
func (r *DirPhoneNumberBatchListResponsePhoneNumberRejectionReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregate batch status. Mirrors the values used on individual phone numbers
// (`submitted`, `in_review`, `verified`, `unsuccessful`, `permanently_rejected`,
// etc.).
type DirPhoneNumberBatchListResponseStatus string

const (
	DirPhoneNumberBatchListResponseStatusSubmitted           DirPhoneNumberBatchListResponseStatus = "submitted"
	DirPhoneNumberBatchListResponseStatusInReview            DirPhoneNumberBatchListResponseStatus = "in_review"
	DirPhoneNumberBatchListResponseStatusVerified            DirPhoneNumberBatchListResponseStatus = "verified"
	DirPhoneNumberBatchListResponseStatusUnsuccessful        DirPhoneNumberBatchListResponseStatus = "unsuccessful"
	DirPhoneNumberBatchListResponseStatusSuspended           DirPhoneNumberBatchListResponseStatus = "suspended"
	DirPhoneNumberBatchListResponseStatusExpired             DirPhoneNumberBatchListResponseStatus = "expired"
	DirPhoneNumberBatchListResponseStatusPermanentlyRejected DirPhoneNumberBatchListResponseStatus = "permanently_rejected"
)

type DirPhoneNumberBatchGetParams struct {
	DirID string `path:"dir_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type DirPhoneNumberBatchListParams struct {
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Restrict to batches whose aggregate status equals this value.
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	FilterStatus DirPhoneNumberBatchListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirPhoneNumberBatchListParams]'s query parameters as
// `url.Values`.
func (r DirPhoneNumberBatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restrict to batches whose aggregate status equals this value.
type DirPhoneNumberBatchListParamsFilterStatus string

const (
	DirPhoneNumberBatchListParamsFilterStatusSubmitted           DirPhoneNumberBatchListParamsFilterStatus = "submitted"
	DirPhoneNumberBatchListParamsFilterStatusInReview            DirPhoneNumberBatchListParamsFilterStatus = "in_review"
	DirPhoneNumberBatchListParamsFilterStatusVerified            DirPhoneNumberBatchListParamsFilterStatus = "verified"
	DirPhoneNumberBatchListParamsFilterStatusUnsuccessful        DirPhoneNumberBatchListParamsFilterStatus = "unsuccessful"
	DirPhoneNumberBatchListParamsFilterStatusSuspended           DirPhoneNumberBatchListParamsFilterStatus = "suspended"
	DirPhoneNumberBatchListParamsFilterStatusExpired             DirPhoneNumberBatchListParamsFilterStatus = "expired"
	DirPhoneNumberBatchListParamsFilterStatusPermanentlyRejected DirPhoneNumberBatchListParamsFilterStatus = "permanently_rejected"
)
