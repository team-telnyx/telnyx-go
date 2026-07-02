// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/stainless-sdks/telnyx-go/v4/internal/encoding/json"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/pagination"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
	"github.com/stainless-sdks/telnyx-go/v4/packages/respjson"
)

// Static reference values the API accepts: call reasons, document types, rejection
// types.
//
// CallReasonService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallReasonService] method instead.
type CallReasonService struct {
	Options []option.RequestOption
}

// NewCallReasonService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCallReasonService(opts ...option.RequestOption) (r CallReasonService) {
	r = CallReasonService{}
	r.Options = opts
	return
}

// Telnyx maintains a library of pre-vetted call-reason phrases (e.g. "Appointment
// reminders", "Billing inquiries") that carry through DIR vetting smoothly. You
// can use any string that fits your use case in `DirCreateRequest.call_reasons`,
// but matching one of these reduces the chance the vetting team flags the phrasing
// for clarification.
func (r *CallReasonService) List(ctx context.Context, query CallReasonListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[CallReasonListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "call_reasons"
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

// Telnyx maintains a library of pre-vetted call-reason phrases (e.g. "Appointment
// reminders", "Billing inquiries") that carry through DIR vetting smoothly. You
// can use any string that fits your use case in `DirCreateRequest.call_reasons`,
// but matching one of these reduces the chance the vetting team flags the phrasing
// for clarification.
func (r *CallReasonService) ListAutoPaging(ctx context.Context, query CallReasonListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[CallReasonListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Check up to 10 candidate `call_reasons` strings against Telnyx's vetting
// heuristics before sending them on a DIR create or update. The endpoint flags
// strings that are likely to be rejected during vetting (too generic, banned
// phrases, length issues, etc.) so you can fix them up front.
func (r *CallReasonService) Validate(ctx context.Context, body CallReasonValidateParams, opts ...option.RequestOption) (res *CallReasonValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "call_reasons/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Pre-vetted call-reason library entry.
type CallReasonListResponse struct {
	ID          string `json:"id" format:"uuid"`
	Description string `json:"description"`
	Reason      string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReasonListResponse) RawJSON() string { return r.JSON.raw }
func (r *CallReasonListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReasonValidateResponse struct {
	Data CallReasonValidateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReasonValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *CallReasonValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReasonValidateResponseData struct {
	// `true` when every supplied reason matches a pre-vetted entry in the call-reason
	// library. When `true`, the DIR will sail through the call-reasons portion of
	// vetting.
	AllPreApproved bool `json:"all_pre_approved" api:"required"`
	// Subset of the input that does NOT match the pre-vetted library. The DIR can
	// still be submitted with these - they will go through manual review.
	NonApprovedReasons []string `json:"non_approved_reasons" api:"required"`
	// `true` when at least one supplied reason is in `non_approved_reasons`.
	// Equivalent to `non_approved_reasons.length > 0` and the inverse of
	// `all_pre_approved`.
	RequiresManualVetting bool `json:"requires_manual_vetting" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllPreApproved        respjson.Field
		NonApprovedReasons    respjson.Field
		RequiresManualVetting respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReasonValidateResponseData) RawJSON() string { return r.JSON.raw }
func (r *CallReasonValidateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReasonListParams struct {
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Default `100` for this endpoint (the call-reason library is
	// small and most callers want the whole list in one call). Maximum 250; values
	// above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallReasonListParams]'s query parameters as `url.Values`.
func (r CallReasonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CallReasonValidateParams struct {
	// **Bare JSON array** of candidate call-reason strings (NOT an object - there is
	// no top-level `call_reasons` key on this endpoint). 1–10 strings, each ≤64
	// characters.
	Body []string
	paramObj
}

func (r CallReasonValidateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *CallReasonValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
