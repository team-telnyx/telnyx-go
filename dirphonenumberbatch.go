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
func (r *DirPhoneNumberBatchService) List(ctx context.Context, dirID string, query DirPhoneNumberBatchListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PhoneNumberBatch], err error) {
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
func (r *DirPhoneNumberBatchService) ListAutoPaging(ctx context.Context, dirID string, query DirPhoneNumberBatchListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PhoneNumberBatch] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, dirID, query, opts...))
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
type DirPhoneNumberStatus string

const (
	DirPhoneNumberStatusSubmitted           DirPhoneNumberStatus = "submitted"
	DirPhoneNumberStatusInReview            DirPhoneNumberStatus = "in_review"
	DirPhoneNumberStatusVerified            DirPhoneNumberStatus = "verified"
	DirPhoneNumberStatusUnsuccessful        DirPhoneNumberStatus = "unsuccessful"
	DirPhoneNumberStatusSuspended           DirPhoneNumberStatus = "suspended"
	DirPhoneNumberStatusExpired             DirPhoneNumberStatus = "expired"
	DirPhoneNumberStatusPermanentlyRejected DirPhoneNumberStatus = "permanently_rejected"
)

// A phone-number batch groups all numbers added in a single bulk-add request.
// Telnyx vets the batch as a unit. The response embeds the full `phone_numbers`
// array so you can read per-number status without a separate call, plus a
// batch-level `status` summarising the unit's progress.
type PhoneNumberBatch struct {
	BatchID string `json:"batch_id" format:"uuid"`
	// The DIR's display name at the time the batch was read.
	DirDisplayName string `json:"dir_display_name"`
	DirID          string `json:"dir_id" format:"uuid"`
	// Documents attached to this batch (e.g. a Letter of Authorization). Empty when
	// none were supplied at add time.
	Documents    []Document `json:"documents"`
	EnterpriseID string     `json:"enterprise_id" format:"uuid"`
	// All phone numbers in this batch, with per-number status.
	PhoneNumbers []DirPhoneNumber `json:"phone_numbers"`
	// Aggregate batch status. Mirrors the values used on individual phone numbers
	// (`submitted`, `in_review`, `verified`, `unsuccessful`, `permanently_rejected`,
	// etc.).
	//
	// Any of "submitted", "in_review", "verified", "unsuccessful", "suspended",
	// "expired", "permanently_rejected".
	Status DirPhoneNumberStatus `json:"status"`
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
func (r PhoneNumberBatch) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirPhoneNumberBatchGetResponse struct {
	// A phone-number batch groups all numbers added in a single bulk-add request.
	// Telnyx vets the batch as a unit. The response embeds the full `phone_numbers`
	// array so you can read per-number status without a separate call, plus a
	// batch-level `status` summarising the unit's progress.
	Data PhoneNumberBatch `json:"data" api:"required"`
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
	FilterStatus DirPhoneNumberStatus `query:"filter[status],omitzero" json:"-"`
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
