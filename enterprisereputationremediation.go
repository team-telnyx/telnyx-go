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

// Phone-number reputation monitoring (spam-score lookup and tracking).
//
// EnterpriseReputationRemediationService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationRemediationService] method instead.
type EnterpriseReputationRemediationService struct {
	Options []option.RequestOption
}

// NewEnterpriseReputationRemediationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnterpriseReputationRemediationService(opts ...option.RequestOption) (r EnterpriseReputationRemediationService) {
	r = EnterpriseReputationRemediationService{}
	r.Options = opts
	return
}

// Retrieve the full detail of a remediation request, including current status,
// per-number results (once available), and submission metadata.
func (r *EnterpriseReputationRemediationService) Get(ctx context.Context, remediationID string, query EnterpriseReputationRemediationGetParams, opts ...option.RequestOption) (res *EnterpriseReputationRemediationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	if remediationID == "" {
		err = errors.New("missing required remediation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/remediation/%s", query.EnterpriseID, remediationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Paginated list of remediation requests for this enterprise. List items omit
// per-number results and webhook URLs to keep the response small; call GET by id
// for full detail. Supports JSON:API pagination and optional filters on status and
// created-at range.
func (r *EnterpriseReputationRemediationService) List(ctx context.Context, enterpriseID string, query EnterpriseReputationRemediationListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EnterpriseReputationRemediationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/remediation", enterpriseID)
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

// Paginated list of remediation requests for this enterprise. List items omit
// per-number results and webhook URLs to keep the response small; call GET by id
// for full detail. Supports JSON:API pagination and optional filters on status and
// created-at range.
func (r *EnterpriseReputationRemediationService) ListAutoPaging(ctx context.Context, enterpriseID string, query EnterpriseReputationRemediationListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EnterpriseReputationRemediationListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, enterpriseID, query, opts...))
}

// Submit a batch of phone numbers belonging to this enterprise for reputation
// remediation. The request is accepted asynchronously: this endpoint returns `202`
// with the persisted request id, then the request transitions through processing
// states until completion. Use the GET endpoints to poll status and per-number
// results.
//
// Each phone number must be in E.164 format and belong to this enterprise. A
// number that already has an in-flight remediation request is rejected.
func (r *EnterpriseReputationRemediationService) Submit(ctx context.Context, enterpriseID string, body EnterpriseReputationRemediationSubmitParams, opts ...option.RequestOption) (res *EnterpriseReputationRemediationSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/remediation", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EnterpriseReputationRemediationGetResponse struct {
	// Full detail of a remediation request, returned on submit and GET by id.
	Data EnterpriseReputationRemediationGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full detail of a remediation request, returned on submit and GET by id.
type EnterpriseReputationRemediationGetResponseData struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CallPurpose string    `json:"call_purpose" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// Total phone numbers in this batch, including any later cancelled. May exceed the
	// sum of the per-category result buckets, which omit cancelled numbers.
	PhoneNumbersCount int64 `json:"phone_numbers_count" api:"required"`
	// Numbers rejected before submission (e.g. cooldown).
	PhoneNumbersIneligible int64 `json:"phone_numbers_ineligible" api:"required"`
	// Numbers accepted for remediation, i.e. not rejected as ineligible. Counts
	// numbers still queued (pending) as well as processed ones.
	PhoneNumbersSubmitted int64 `json:"phone_numbers_submitted" api:"required"`
	// Customer-facing status of a remediation request.
	//
	// Any of "pending", "in_progress", "completed", "failed", "cancelled".
	Status       string    `json:"status" api:"required"`
	UpdatedAt    time.Time `json:"updated_at" api:"required" format:"date-time"`
	ContactEmail string    `json:"contact_email" api:"nullable" format:"email"`
	// Per-category buckets. Populated once results are available. Null while the
	// request is still pending.
	Results          EnterpriseReputationRemediationGetResponseDataResults `json:"results" api:"nullable"`
	Tier1CompletedAt time.Time                                             `json:"tier1_completed_at" api:"nullable" format:"date-time"`
	Tier2CompletedAt time.Time                                             `json:"tier2_completed_at" api:"nullable" format:"date-time"`
	WebhookURL       string                                                `json:"webhook_url" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CallPurpose            respjson.Field
		CreatedAt              respjson.Field
		PhoneNumbersCount      respjson.Field
		PhoneNumbersIneligible respjson.Field
		PhoneNumbersSubmitted  respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		ContactEmail           respjson.Field
		Results                respjson.Field
		Tier1CompletedAt       respjson.Field
		Tier2CompletedAt       respjson.Field
		WebhookURL             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-category buckets. Populated once results are available. Null while the
// request is still pending.
type EnterpriseReputationRemediationGetResponseDataResults struct {
	Ineligible     []string `json:"ineligible"`
	NotFlagged     []string `json:"not_flagged"`
	Refused        []string `json:"refused"`
	Remediated     []string `json:"remediated"`
	RequiresReview []string `json:"requires_review"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ineligible     respjson.Field
		NotFlagged     respjson.Field
		Refused        respjson.Field
		Remediated     respjson.Field
		RequiresReview respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationGetResponseDataResults) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationGetResponseDataResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Slim list-endpoint shape. Omits per-number results and webhook URLs to keep
// responses small.
type EnterpriseReputationRemediationListResponse struct {
	ID                string    `json:"id" api:"required" format:"uuid"`
	CallPurpose       string    `json:"call_purpose" api:"required"`
	CreatedAt         time.Time `json:"created_at" api:"required" format:"date-time"`
	PhoneNumbersCount int64     `json:"phone_numbers_count" api:"required"`
	// Customer-facing status of a remediation request.
	//
	// Any of "pending", "in_progress", "completed", "failed", "cancelled".
	Status           EnterpriseReputationRemediationListResponseStatus `json:"status" api:"required"`
	UpdatedAt        time.Time                                         `json:"updated_at" api:"required" format:"date-time"`
	Tier1CompletedAt time.Time                                         `json:"tier1_completed_at" api:"nullable" format:"date-time"`
	Tier2CompletedAt time.Time                                         `json:"tier2_completed_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CallPurpose       respjson.Field
		CreatedAt         respjson.Field
		PhoneNumbersCount respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		Tier1CompletedAt  respjson.Field
		Tier2CompletedAt  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationListResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer-facing status of a remediation request.
type EnterpriseReputationRemediationListResponseStatus string

const (
	EnterpriseReputationRemediationListResponseStatusPending    EnterpriseReputationRemediationListResponseStatus = "pending"
	EnterpriseReputationRemediationListResponseStatusInProgress EnterpriseReputationRemediationListResponseStatus = "in_progress"
	EnterpriseReputationRemediationListResponseStatusCompleted  EnterpriseReputationRemediationListResponseStatus = "completed"
	EnterpriseReputationRemediationListResponseStatusFailed     EnterpriseReputationRemediationListResponseStatus = "failed"
	EnterpriseReputationRemediationListResponseStatusCancelled  EnterpriseReputationRemediationListResponseStatus = "cancelled"
)

type EnterpriseReputationRemediationSubmitResponse struct {
	// Full detail of a remediation request, returned on submit and GET by id.
	Data EnterpriseReputationRemediationSubmitResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full detail of a remediation request, returned on submit and GET by id.
type EnterpriseReputationRemediationSubmitResponseData struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CallPurpose string    `json:"call_purpose" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// Total phone numbers in this batch, including any later cancelled. May exceed the
	// sum of the per-category result buckets, which omit cancelled numbers.
	PhoneNumbersCount int64 `json:"phone_numbers_count" api:"required"`
	// Numbers rejected before submission (e.g. cooldown).
	PhoneNumbersIneligible int64 `json:"phone_numbers_ineligible" api:"required"`
	// Numbers accepted for remediation, i.e. not rejected as ineligible. Counts
	// numbers still queued (pending) as well as processed ones.
	PhoneNumbersSubmitted int64 `json:"phone_numbers_submitted" api:"required"`
	// Customer-facing status of a remediation request.
	//
	// Any of "pending", "in_progress", "completed", "failed", "cancelled".
	Status       string    `json:"status" api:"required"`
	UpdatedAt    time.Time `json:"updated_at" api:"required" format:"date-time"`
	ContactEmail string    `json:"contact_email" api:"nullable" format:"email"`
	// Per-category buckets. Populated once results are available. Null while the
	// request is still pending.
	Results          EnterpriseReputationRemediationSubmitResponseDataResults `json:"results" api:"nullable"`
	Tier1CompletedAt time.Time                                                `json:"tier1_completed_at" api:"nullable" format:"date-time"`
	Tier2CompletedAt time.Time                                                `json:"tier2_completed_at" api:"nullable" format:"date-time"`
	WebhookURL       string                                                   `json:"webhook_url" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CallPurpose            respjson.Field
		CreatedAt              respjson.Field
		PhoneNumbersCount      respjson.Field
		PhoneNumbersIneligible respjson.Field
		PhoneNumbersSubmitted  respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		ContactEmail           respjson.Field
		Results                respjson.Field
		Tier1CompletedAt       respjson.Field
		Tier2CompletedAt       respjson.Field
		WebhookURL             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationSubmitResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationSubmitResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-category buckets. Populated once results are available. Null while the
// request is still pending.
type EnterpriseReputationRemediationSubmitResponseDataResults struct {
	Ineligible     []string `json:"ineligible"`
	NotFlagged     []string `json:"not_flagged"`
	Refused        []string `json:"refused"`
	Remediated     []string `json:"remediated"`
	RequiresReview []string `json:"requires_review"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ineligible     respjson.Field
		NotFlagged     respjson.Field
		Refused        respjson.Field
		Remediated     respjson.Field
		RequiresReview respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationRemediationSubmitResponseDataResults) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationRemediationSubmitResponseDataResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationRemediationGetParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type EnterpriseReputationRemediationListParams struct {
	// Only requests created on or after this timestamp (ISO 8601).
	FilterCreatedAtGte param.Opt[time.Time] `query:"filter[created_at][gte],omitzero" format:"date-time" json:"-"`
	// Only requests created on or before this timestamp (ISO 8601).
	FilterCreatedAtLte param.Opt[time.Time] `query:"filter[created_at][lte],omitzero" format:"date-time" json:"-"`
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by customer-facing status.
	//
	// Any of "pending", "in_progress", "completed", "failed", "cancelled".
	FilterStatus EnterpriseReputationRemediationListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseReputationRemediationListParams]'s query
// parameters as `url.Values`.
func (r EnterpriseReputationRemediationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by customer-facing status.
type EnterpriseReputationRemediationListParamsFilterStatus string

const (
	EnterpriseReputationRemediationListParamsFilterStatusPending    EnterpriseReputationRemediationListParamsFilterStatus = "pending"
	EnterpriseReputationRemediationListParamsFilterStatusInProgress EnterpriseReputationRemediationListParamsFilterStatus = "in_progress"
	EnterpriseReputationRemediationListParamsFilterStatusCompleted  EnterpriseReputationRemediationListParamsFilterStatus = "completed"
	EnterpriseReputationRemediationListParamsFilterStatusFailed     EnterpriseReputationRemediationListParamsFilterStatus = "failed"
	EnterpriseReputationRemediationListParamsFilterStatusCancelled  EnterpriseReputationRemediationListParamsFilterStatus = "cancelled"
)

type EnterpriseReputationRemediationSubmitParams struct {
	// How the numbers are used (free text).
	CallPurpose string `json:"call_purpose" api:"required"`
	// Phone numbers in E.164 format. Each must belong to this enterprise. Maximum
	// 2,000 per request.
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	// Optional contact email for this remediation request.
	ContactEmail param.Opt[string] `json:"contact_email,omitzero" format:"email"`
	// Optional https:// URL for status notifications.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"uri"`
	paramObj
}

func (r EnterpriseReputationRemediationSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationRemediationSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationRemediationSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
