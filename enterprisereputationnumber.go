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
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s", params.EnterpriseID, phoneNumber)
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
// must be in E.164 format (`+1NPANXXXXXX` for US/CA) and belong to your Telnyx
// phone-number inventory.
//
// **Prerequisite**: reputation must already be enabled on this enterprise (see
// `POST .../reputation`).
//
// **Pricing:** This is a billable action. See https://telnyx.com/pricing/numbers
// for current pricing.
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
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s", body.EnterpriseID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
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

type ReputationPhoneNumber struct {
	ID           string    `json:"id" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	EnterpriseID string    `json:"enterprise_id" format:"uuid"`
	// E.164 with leading `+`.
	PhoneNumber string `json:"phone_number"`
	// `null` until the first refresh has been collected for this number.
	ReputationData shared.ReputationData `json:"reputation_data" api:"nullable"`
	UpdatedAt      time.Time             `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumber    respjson.Field
		ReputationData respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ReputationPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
