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
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Phone-number reputation monitoring (spam-score lookup and tracking).
//
// EnterpriseReputationNumberService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationNumberService] method instead.
type EnterpriseReputationNumberService struct {
	Options []option.RequestOption
}

// NewEnterpriseReputationNumberService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnterpriseReputationNumberService(opts ...option.RequestOption) (r EnterpriseReputationNumberService) {
	r = EnterpriseReputationNumberService{}
	r.Options = opts
	return
}

// Retrieve one registered number with its latest reputation snapshot. The
// `phone_number` path parameter is in E.164 format and must be URL-encoded (e.g.
// `%2B19493253498`).
func (r *EnterpriseReputationNumberService) Get(ctx context.Context, phoneNumber string, params EnterpriseReputationNumberGetParams, opts ...option.RequestOption) (res *ReputationPhoneNumberWithReputation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s", params.EnterpriseID, url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Paginated list of phone numbers registered for reputation monitoring under this
// enterprise. The response includes the latest reputation snapshot per number
// where one has been collected.
func (r *EnterpriseReputationNumberService) List(ctx context.Context, enterpriseID string, query EnterpriseReputationNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ReputationPhoneNumber], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers", enterpriseID)
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

// Paginated list of phone numbers registered for reputation monitoring under this
// enterprise. The response includes the latest reputation snapshot per number
// where one has been collected.
func (r *EnterpriseReputationNumberService) ListAutoPaging(ctx context.Context, enterpriseID string, query EnterpriseReputationNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ReputationPhoneNumber] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, enterpriseID, query, opts...))
}

// Add up to 100 phone numbers to reputation monitoring on this enterprise. Each
// must be in E.164 format (`+1NPANXXXXXX` for US/CA). A number in your Telnyx
// phone-number inventory is accepted as `source: telnyx`; a number that is not is
// treated as hosted with another provider and must meet the requirements below.
//
// **Prerequisite**: reputation must already be enabled on this enterprise (see
// `POST .../reputation`).
//
// **Pricing:** This is a billable action. See https://telnyx.com/pricing/numbers
// for current pricing.
//
// A phone number that is not in your Telnyx inventory is treated as a number
// hosted with another provider (`source: external`). It must be a US local number,
// and the enterprise's ownership proof must be approved first
// (`POST /enterprises/{enterprise_id}/reputation/external_verification`);
// otherwise the request fails with a 400 whose `detail` contains
// `external_verification_required`. Ownership of such numbers is checked every
// day.
func (r *EnterpriseReputationNumberService) Associate(ctx context.Context, enterpriseID string, body EnterpriseReputationNumberAssociateParams, opts ...option.RequestOption) (res *ReputationPhoneNumberList, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop tracking the reputation of this phone number. The number itself remains in
// your inventory; only the reputation registration is removed.
func (r *EnterpriseReputationNumberService) Disassociate(ctx context.Context, phoneNumber string, body EnterpriseReputationNumberDisassociateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return err
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s", body.EnterpriseID, url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Answer an ownership question about a phone number hosted with another provider.
//
// Send `still_own: true` to confirm you still own a number whose
// `ownership_status` is `pending_confirmation` or `flagged`; its status returns to
// `verified` and, after a carrier change, the new carrier becomes the one on
// record.
//
// Send `still_own: false` to report that you no longer own the number. It is
// removed from Number Reputation immediately, whatever its current status, and the
// response shows it with `ownership_status: removed`.
func (r *EnterpriseReputationNumberService) OwnershipConfirm(ctx context.Context, phoneNumber string, params EnterpriseReputationNumberOwnershipConfirmParams, opts ...option.RequestOption) (res *ReputationPhoneNumberWithReputation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s/ownership_confirm", params.EnterpriseID, url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Immediately refresh the stored reputation data for the listed numbers. This is
// in addition to the periodic refresh determined by `check_frequency`. Up to 100
// numbers per call. The response carries the kicked-off jobs; the actual refresh
// runs asynchronously.
//
// **Pricing:** Forcing a refresh performs live reputation lookups, which are
// billable. See https://telnyx.com/pricing/numbers for current pricing.
func (r *EnterpriseReputationNumberService) Refresh(ctx context.Context, enterpriseID string, body EnterpriseReputationNumberRefreshParams, opts ...option.RequestOption) (res *EnterpriseReputationNumberRefreshResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/refresh", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the ownership status of a phone number, when it was last checked, the
// carrier recorded as its baseline, and the history of ownership events.
func (r *EnterpriseReputationNumberService) GetOwnership(ctx context.Context, phoneNumber string, query EnterpriseReputationNumberGetOwnershipParams, opts ...option.RequestOption) (res *EnterpriseReputationNumberGetOwnershipResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s/ownership", query.EnterpriseID, url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Where the phone number is hosted. `telnyx`: in your Telnyx inventory.
// `external`: a number you own that is hosted with another provider.
type ReputationNumberSource string

const (
	ReputationNumberSourceTelnyx   ReputationNumberSource = "telnyx"
	ReputationNumberSourceExternal ReputationNumberSource = "external"
)

// Result of ongoing ownership checks. `verified`: the number is with the carrier
// on record. `pending_confirmation`: the number's carrier changed, or a periodic
// ownership re-verification is due; confirm you still own it before `confirm_by`.
// `flagged`: a periodic re-verification went unanswered; the number stays
// monitored and is not removed, but confirm ownership
// (`POST .../ownership_confirm`) to return it to `verified`. `removed`: the number
// is no longer monitored; returned by `POST .../ownership_confirm` with
// `still_own: false`, and when listing numbers with
// `filter[ownership_status]=removed`.
type ReputationOwnershipStatus string

const (
	ReputationOwnershipStatusVerified            ReputationOwnershipStatus = "verified"
	ReputationOwnershipStatusFlagged             ReputationOwnershipStatus = "flagged"
	ReputationOwnershipStatusPendingConfirmation ReputationOwnershipStatus = "pending_confirmation"
	ReputationOwnershipStatusRemoved             ReputationOwnershipStatus = "removed"
)

type ReputationPhoneNumber struct {
	ID string `json:"id" format:"uuid"`
	// Deadline to confirm you still own the number. Set only while `ownership_status`
	// is `pending_confirmation`.
	ConfirmBy    time.Time `json:"confirm_by" api:"nullable" format:"date-time"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	EnterpriseID string    `json:"enterprise_id" format:"uuid"`
	// Why the number's ownership was flagged. `carrier_changed` when an ownership
	// check found the number with a different carrier.
	FlagReason string `json:"flag_reason" api:"nullable"`
	// Result of ongoing ownership checks. `verified`: the number is with the carrier
	// on record. `pending_confirmation`: the number's carrier changed, or a periodic
	// ownership re-verification is due; confirm you still own it before `confirm_by`.
	// `flagged`: a periodic re-verification went unanswered; the number stays
	// monitored and is not removed, but confirm ownership
	// (`POST .../ownership_confirm`) to return it to `verified`. `removed`: the number
	// is no longer monitored; returned by `POST .../ownership_confirm` with
	// `still_own: false`, and when listing numbers with
	// `filter[ownership_status]=removed`.
	//
	// Any of "verified", "flagged", "pending_confirmation", "removed".
	OwnershipStatus ReputationOwnershipStatus `json:"ownership_status"`
	// E.164 with leading `+`.
	PhoneNumber string `json:"phone_number"`
	// Why a number was removed. `ported_away`: you confirmed the number left your
	// control after a carrier change. `unconfirmed_ownership`: a carrier change was
	// not confirmed before the deadline. `reassigned`: the number was reported as
	// permanently disconnected. `customer_reported`: you reported that you no longer
	// own the number. `customer_removed`: you removed the number.
	//
	// Any of "ported_away", "unconfirmed_ownership", "reassigned",
	// "customer_reported", "customer_removed".
	RemovalReason ReputationPhoneNumberRemovalReason `json:"removal_reason" api:"nullable"`
	// When the number was removed. Set only for removed numbers.
	RemovedAt time.Time `json:"removed_at" api:"nullable" format:"date-time"`
	// Reputation snapshot for a phone number. Each metric is a 0–100 score;
	// `spam_risk` is a coarse bucket. Field set may grow over time - read by key.
	ReputationData shared.ReputationData `json:"reputation_data" api:"nullable"`
	// Where the phone number is hosted. `telnyx`: in your Telnyx inventory.
	// `external`: a number you own that is hosted with another provider.
	//
	// Any of "telnyx", "external".
	Source    ReputationNumberSource `json:"source"`
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ConfirmBy       respjson.Field
		CreatedAt       respjson.Field
		EnterpriseID    respjson.Field
		FlagReason      respjson.Field
		OwnershipStatus respjson.Field
		PhoneNumber     respjson.Field
		RemovalReason   respjson.Field
		RemovedAt       respjson.Field
		ReputationData  respjson.Field
		Source          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ReputationPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why a number was removed. `ported_away`: you confirmed the number left your
// control after a carrier change. `unconfirmed_ownership`: a carrier change was
// not confirmed before the deadline. `reassigned`: the number was reported as
// permanently disconnected. `customer_reported`: you reported that you no longer
// own the number. `customer_removed`: you removed the number.
type ReputationPhoneNumberRemovalReason string

const (
	ReputationPhoneNumberRemovalReasonPortedAway           ReputationPhoneNumberRemovalReason = "ported_away"
	ReputationPhoneNumberRemovalReasonUnconfirmedOwnership ReputationPhoneNumberRemovalReason = "unconfirmed_ownership"
	ReputationPhoneNumberRemovalReasonReassigned           ReputationPhoneNumberRemovalReason = "reassigned"
	ReputationPhoneNumberRemovalReasonCustomerReported     ReputationPhoneNumberRemovalReason = "customer_reported"
	ReputationPhoneNumberRemovalReasonCustomerRemoved      ReputationPhoneNumberRemovalReason = "customer_removed"
)

type ReputationPhoneNumberList struct {
	Data []ReputationPhoneNumber `json:"data" api:"required"`
	// JSON:API pagination metadata returned with every paginated list response. Page
	// numbering is 1-based. `page_size` reports the number of items actually returned
	// in `data` for this page; the requested size is taken from the `page[size]` query
	// parameter.
	Meta NumberReputationPaginationMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationPhoneNumberList) RawJSON() string { return r.JSON.raw }
func (r *ReputationPhoneNumberList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationPhoneNumberWithReputation struct {
	Data ReputationPhoneNumber `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationPhoneNumberWithReputation) RawJSON() string { return r.JSON.raw }
func (r *ReputationPhoneNumberWithReputation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberRefreshResponse struct {
	Data EnterpriseReputationNumberRefreshResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberRefreshResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberRefreshResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberRefreshResponseData struct {
	// Per-number outcome of the refresh.
	Results         []EnterpriseReputationNumberRefreshResponseDataResult `json:"results" api:"required"`
	TotalFailed     int64                                                 `json:"total_failed" api:"required"`
	TotalRequested  int64                                                 `json:"total_requested" api:"required"`
	TotalSuccessful int64                                                 `json:"total_successful" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results         respjson.Field
		TotalFailed     respjson.Field
		TotalRequested  respjson.Field
		TotalSuccessful respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberRefreshResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberRefreshResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberRefreshResponseDataResult struct {
	PhoneNumber string `json:"phone_number" api:"required"`
	Success     bool   `json:"success" api:"required"`
	// `null` when `success` is `true`; carries an error message otherwise.
	Error string `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		Success     respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberRefreshResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberRefreshResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetOwnershipResponse struct {
	Data EnterpriseReputationNumberGetOwnershipResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberGetOwnershipResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberGetOwnershipResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetOwnershipResponseData struct {
	// Service provider ID (SPID) of the carrier that held the number when its
	// ownership was last established.
	BaselineCarrier    string    `json:"baseline_carrier" api:"nullable"`
	BaselineRecordedAt time.Time `json:"baseline_recorded_at" api:"nullable" format:"date-time"`
	ConfirmBy          time.Time `json:"confirm_by" api:"nullable" format:"date-time"`
	FlagReason         string    `json:"flag_reason" api:"nullable"`
	// Ownership events, most recent first.
	History []EnterpriseReputationNumberGetOwnershipResponseDataHistory `json:"history"`
	// When ownership of the number was last checked.
	LastCheckedAt time.Time `json:"last_checked_at" api:"nullable" format:"date-time"`
	// Result of ongoing ownership checks. `verified`: the number is with the carrier
	// on record. `pending_confirmation`: the number's carrier changed, or a periodic
	// ownership re-verification is due; confirm you still own it before `confirm_by`.
	// `flagged`: a periodic re-verification went unanswered; the number stays
	// monitored and is not removed, but confirm ownership
	// (`POST .../ownership_confirm`) to return it to `verified`. `removed`: the number
	// is no longer monitored; returned by `POST .../ownership_confirm` with
	// `still_own: false`, and when listing numbers with
	// `filter[ownership_status]=removed`.
	//
	// Any of "verified", "flagged", "pending_confirmation", "removed".
	OwnershipStatus ReputationOwnershipStatus `json:"ownership_status"`
	PhoneNumber     string                    `json:"phone_number"`
	// Where the phone number is hosted. `telnyx`: in your Telnyx inventory.
	// `external`: a number you own that is hosted with another provider.
	//
	// Any of "telnyx", "external".
	Source ReputationNumberSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineCarrier    respjson.Field
		BaselineRecordedAt respjson.Field
		ConfirmBy          respjson.Field
		FlagReason         respjson.Field
		History            respjson.Field
		LastCheckedAt      respjson.Field
		OwnershipStatus    respjson.Field
		PhoneNumber        respjson.Field
		Source             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberGetOwnershipResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberGetOwnershipResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetOwnershipResponseDataHistory struct {
	// `added`, `check`, `flag`, `confirm` or `removed`.
	EventType  string    `json:"event_type"`
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// What the event found or did, for example `unchanged` or `carrier_changed`.
	Outcome string `json:"outcome" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Outcome     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberGetOwnershipResponseDataHistory) RawJSON() string {
	return r.JSON.raw
}
func (r *EnterpriseReputationNumberGetOwnershipResponseDataHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	// When true, fetches fresh reputation data (incurs API cost). When false
	// (default), returns cached data.
	Fresh param.Opt[bool] `query:"fresh,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseReputationNumberGetParams]'s query parameters as
// `url.Values`.
func (r EnterpriseReputationNumberGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnterpriseReputationNumberListParams struct {
	// Partial match on phone number. Must contain at least 5 digits.
	FilterPhoneNumberContains param.Opt[string] `query:"filter[phone_number][contains],omitzero" json:"-"`
	// Exact phone-number match (E.164).
	FilterPhoneNumberEq param.Opt[string] `query:"filter[phone_number][eq],omitzero" json:"-"`
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Default 10. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Only numbers with this ownership status. `removed` lists the numbers that were
	// removed, each with its `removal_reason` and `removed_at`; removed numbers are
	// not returned otherwise.
	//
	// Any of "verified", "flagged", "pending_confirmation", "removed".
	FilterOwnershipStatus ReputationOwnershipStatus `query:"filter[ownership_status],omitzero" json:"-"`
	// Only numbers with this source.
	//
	// Any of "telnyx", "external".
	FilterSource ReputationNumberSource `query:"filter[source],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseReputationNumberListParams]'s query parameters as
// `url.Values`.
func (r EnterpriseReputationNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnterpriseReputationNumberAssociateParams struct {
	// 1–100 phone numbers in E.164 format with a leading `+`.
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r EnterpriseReputationNumberAssociateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationNumberAssociateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationNumberAssociateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberDisassociateParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type EnterpriseReputationNumberOwnershipConfirmParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	// `true` confirms you still own the number. `false` reports that you no longer own
	// it and removes it from Number Reputation.
	StillOwn bool `json:"still_own" api:"required"`
	// Optional id of an uploaded document supporting the confirmation.
	EvidenceDocumentID param.Opt[string] `json:"evidence_document_id,omitzero"`
	paramObj
}

func (r EnterpriseReputationNumberOwnershipConfirmParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationNumberOwnershipConfirmParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationNumberOwnershipConfirmParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberRefreshParams struct {
	// Phone numbers to refresh reputation data for. 1–100 numbers per request, each in
	// E.164 format. Reputation refreshes are subject to per-enterprise rate limits.
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r EnterpriseReputationNumberRefreshParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationNumberRefreshParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationNumberRefreshParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetOwnershipParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
