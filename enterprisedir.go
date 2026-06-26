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
func (r *EnterpriseDirService) New(ctx context.Context, enterpriseID string, body EnterpriseDirNewParams, opts ...option.RequestOption) (res *DirWrapped, err error) {
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
func (r *EnterpriseDirService) List(ctx context.Context, enterpriseID string, query EnterpriseDirListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Dir], err error) {
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
func (r *EnterpriseDirService) ListAutoPaging(ctx context.Context, enterpriseID string, query EnterpriseDirListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Dir] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, enterpriseID, query, opts...))
}

type EnterpriseDirNewParams struct {
	// Contact email of the authorizer. Telnyx may send verification or
	// infringement-notice email here; use a monitored mailbox.
	AuthorizerEmail string `json:"authorizer_email" api:"required" format:"email"`
	// Name of the person at your enterprise who is authorizing this DIR registration.
	// Must be a real individual (used for audit and trademark-claim contests).
	AuthorizerName string `json:"authorizer_name" api:"required"`
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
	// 1–10 reasons your business calls customers. Validate phrasing against
	// `POST /call_reasons/validate`.
	CallReasons []string `json:"call_reasons,omitzero"`
	// Supporting documents. Each `document_id` may appear at most once on a DIR.
	Documents []DocumentParam `json:"documents,omitzero"`
	paramObj
}

func (r EnterpriseDirNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseDirNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseDirNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	FilterStatus DirStatus `query:"filter[status],omitzero" json:"-"`
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
