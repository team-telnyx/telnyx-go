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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Daily and monthly spend limits per product. A limit applies to the organization
// of the authenticated user, or to the user's own account when they belong to no
// organization; every user of the organization sees and changes the same limits.
//
//   - **Periods.** `daily` covers the current UTC day and `monthly` the current UTC
//     calendar month. The two limits are independent: you can set either, both or
//     neither.
//   - **Blocking.** When spend in a period goes above the limit (strictly greater),
//     the product is blocked until the period ends: 00:00 UTC the next day for
//     `daily`, 00:00 UTC on the 1st of the next month for `monthly`. A block appears
//     within about 2 minutes (daily) or 10 minutes (monthly) of the spend being
//     recorded.
//   - **Changes apply immediately.** Creating, updating or deleting a limit checks
//     the period's spend in the same request: raising the limit above the spend, or
//     removing it, lifts that period's block, and lowering it below the spend blocks
//     the product at once. The `evaluation` object in the response says what
//     happened.
//   - **Supported products.** Today only `inference` supports spend limits. A
//     blocked account gets HTTP 403 with the error title
//     `Inference spend limit reached` (code `10039`) on new billable chat
//     completions, Responses, Anthropic Messages and classification requests;
//     requests already running finish normally. Take the list of products from the
//     list operation.
//   - **Limits set by Telnyx.** Telnyx support can also set a limit on your account.
//     It is listed with `origin: operator` and you can update or delete it like your
//     own.
//
// SpendLimitService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendLimitService] method instead.
type SpendLimitService struct {
	Options []option.RequestOption
}

// NewSpendLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpendLimitService(opts ...option.RequestOption) (r SpendLimitService) {
	r = SpendLimitService{}
	r.Options = opts
	return
}

// Sets a limit for a product and period that has none. Send exactly one of
// `amount` and `unlimited: true`. The period's spend is checked at once: if it is
// already above the new limit, the product is blocked immediately
// (`evaluation.blocked_now`). Returns 409 when a limit already exists for the
// product and period; update it instead.
func (r *SpendLimitService) New(ctx context.Context, body SpendLimitNewParams, opts ...option.RequestOption) (res *SpendLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "spend_limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replaces the value of the existing limit for the product and period. Send
// exactly one of `amount` and `unlimited: true`. The period's spend is checked at
// once: raising the limit above the spend lifts the period's block
// (`evaluation.released`), and lowering it below the spend blocks the product
// (`evaluation.blocked_now`). Returns 404 when no limit is set; create it instead.
func (r *SpendLimitService) Update(ctx context.Context, product string, params SpendLimitUpdateParams, opts ...option.RequestOption) (res *SpendLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if product == "" {
		err = errors.New("missing required product parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend_limits/%s", url.PathEscape(product))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns one entry per product and period you can set a limit on, with the limit,
// the spend so far in the period and whether the product is blocked. An entry
// without a limit is still listed (`limit: null`). When the spend cannot be read,
// the entry is returned with `spend_usd: null` and `spend_error` set. The list is
// not paginated.
func (r *SpendLimitService) List(ctx context.Context, opts ...option.RequestOption) (res *SpendLimitListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "spend_limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes the limit for the product and period. For `inference`, which has no
// default limit, the product becomes unlimited for the period and the period's
// block is lifted (`evaluation.released`). The response carries `limit: null` and
// the `effective_limit_usd` that applies after the removal. Returns 404 when no
// limit is set.
func (r *SpendLimitService) Delete(ctx context.Context, product string, body SpendLimitDeleteParams, opts ...option.RequestOption) (res *SpendLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if product == "" {
		err = errors.New("missing required product parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend_limits/%s", url.PathEscape(product))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// The spend limit, spend and block state of one product and period.
type SpendLimit struct {
	// The active block of the period. `null` when the period is not blocked.
	Block SpendLimitBlock `json:"block" api:"required"`
	// The product is blocked for this period. Always `false` in write responses; list
	// the limits to read the block state.
	Blocked bool `json:"blocked" api:"required"`
	// The limit in USD that is enforced, as a decimal string. `null` means unlimited.
	EffectiveLimitUsd string `json:"effective_limit_usd" api:"required"`
	// The limit set on the account for the product and period, whoever set it. `null`
	// when none is set.
	Limit SpendLimitLimit `json:"limit" api:"required"`
	// `daily` is the current UTC day; `monthly` is the current UTC calendar month.
	//
	// Any of "daily", "monthly".
	Period SpendLimitPeriod `json:"period" api:"required"`
	// Exclusive end of the current period, a UTC date.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date"`
	// First UTC day of the current period.
	PeriodStart time.Time `json:"period_start" api:"required" format:"date"`
	// Product the entry applies to.
	Product string `json:"product" api:"required"`
	// Display name of the product.
	ProductName string `json:"product_name" api:"required"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type" api:"required"`
	// Set when `spend_usd` is `null`.
	SpendError string `json:"spend_error" api:"required"`
	// Spend in USD so far in the period, as a decimal string. It can lag actual usage
	// by about a minute. `null` when it could not be read.
	SpendUsd string `json:"spend_usd" api:"required"`
	// What a create, update or delete did to the period at once. Only present in write
	// responses.
	Evaluation SpendLimitEvaluation `json:"evaluation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Block             respjson.Field
		Blocked           respjson.Field
		EffectiveLimitUsd respjson.Field
		Limit             respjson.Field
		Period            respjson.Field
		PeriodEnd         respjson.Field
		PeriodStart       respjson.Field
		Product           respjson.Field
		ProductName       respjson.Field
		RecordType        respjson.Field
		SpendError        respjson.Field
		SpendUsd          respjson.Field
		Evaluation        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendLimit) RawJSON() string { return r.JSON.raw }
func (r *SpendLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The active block of the period. `null` when the period is not blocked.
type SpendLimitBlock struct {
	// Exclusive end of the block: it is lifted at 00:00 UTC on this date at the
	// latest.
	BlockedUntil time.Time `json:"blocked_until" api:"required" format:"date"`
	// When the block started.
	DetectedAt time.Time `json:"detected_at" api:"required" format:"date-time"`
	// The limit in USD that the spend went above, as a decimal string.
	LimitUsd string `json:"limit_usd" api:"required"`
	// Spend in USD when the block started, as a decimal string.
	SpendUsd string `json:"spend_usd" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedUntil respjson.Field
		DetectedAt   respjson.Field
		LimitUsd     respjson.Field
		SpendUsd     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendLimitBlock) RawJSON() string { return r.JSON.raw }
func (r *SpendLimitBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The limit set on the account for the product and period, whoever set it. `null`
// when none is set.
type SpendLimitLimit struct {
	// Limit in USD, as a decimal string. `null` when `unlimited` is true.
	Amount string `json:"amount" api:"required"`
	// `self_service` when a user of the account set it, `operator` when Telnyx support
	// did.
	//
	// Any of "self_service", "operator".
	Origin string `json:"origin" api:"required"`
	// True when the limit was set to explicitly no cap.
	Unlimited bool `json:"unlimited" api:"required"`
	// When the limit was last set or changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Origin      respjson.Field
		Unlimited   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendLimitLimit) RawJSON() string { return r.JSON.raw }
func (r *SpendLimitLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a create, update or delete did to the period at once. Only present in write
// responses.
type SpendLimitEvaluation struct {
	// The change blocked the product: the spend was already above the new limit.
	BlockedNow bool `json:"blocked_now" api:"required"`
	// The spend could not be checked now. The change is saved and applied within a few
	// minutes.
	EvaluationDeferred bool `json:"evaluation_deferred" api:"required"`
	// The change lifted a block of this period.
	Released bool `json:"released" api:"required"`
	// Spend in USD used for the check, as a decimal string. `null` when the spend was
	// not checked.
	SpendUsd string `json:"spend_usd" api:"required"`
	// The other period has an active block, so the product stays blocked whatever this
	// period's result.
	StillBlockedOtherPeriod bool `json:"still_blocked_other_period" api:"required"`
	// A block of this period remains because the spend is still above the new limit.
	StillOverLimit bool `json:"still_over_limit" api:"required"`
	// Additional information about the result, when there is any.
	Note string `json:"note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedNow              respjson.Field
		EvaluationDeferred      respjson.Field
		Released                respjson.Field
		SpendUsd                respjson.Field
		StillBlockedOtherPeriod respjson.Field
		StillOverLimit          respjson.Field
		Note                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendLimitEvaluation) RawJSON() string { return r.JSON.raw }
func (r *SpendLimitEvaluation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `daily` is the current UTC day; `monthly` is the current UTC calendar month.
type SpendLimitPeriod string

const (
	SpendLimitPeriodDaily   SpendLimitPeriod = "daily"
	SpendLimitPeriodMonthly SpendLimitPeriod = "monthly"
)

type SpendLimitResponse struct {
	// The spend limit, spend and block state of one product and period.
	Data SpendLimit `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *SpendLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendLimitListResponse struct {
	Data []SpendLimit               `json:"data" api:"required"`
	Meta SpendLimitListResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendLimitListResponse) RawJSON() string { return r.JSON.raw }
func (r *SpendLimitListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendLimitListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
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
func (r SpendLimitListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SpendLimitListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendLimitNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. A limit
	// in USD.
	OfAmount *SpendLimitNewParamsBodyCreateSpendLimitWithAmount `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Explicitly no cap.
	OfUnlimited *SpendLimitNewParamsBodyCreateSpendLimitUnlimited `json:",inline"`

	paramObj
}

func (u SpendLimitNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAmount, u.OfUnlimited)
}
func (r *SpendLimitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A limit in USD.
//
// The properties Amount, Product are required.
type SpendLimitNewParamsBodyCreateSpendLimitWithAmount struct {
	// Limit in USD. `0` blocks at the first cent of spend.
	Amount float64 `json:"amount" api:"required"`
	// Product to limit, as returned in `product` by the list operation.
	Product string `json:"product" api:"required"`
	// Why the limit is set or changed, kept for audit.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// `daily` is the current UTC day; `monthly` is the current UTC calendar month.
	//
	// Any of "daily", "monthly".
	Period SpendLimitPeriod `json:"period,omitzero"`
	// Optional; only `false` is allowed together with `amount`.
	//
	// Any of false.
	Unlimited bool `json:"unlimited,omitzero"`
	paramObj
}

func (r SpendLimitNewParamsBodyCreateSpendLimitWithAmount) MarshalJSON() (data []byte, err error) {
	type shadow SpendLimitNewParamsBodyCreateSpendLimitWithAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendLimitNewParamsBodyCreateSpendLimitWithAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SpendLimitNewParamsBodyCreateSpendLimitWithAmount](
		"unlimited", false,
	)
}

// Explicitly no cap.
//
// The properties Product, Unlimited are required.
type SpendLimitNewParamsBodyCreateSpendLimitUnlimited struct {
	// Product to limit, as returned in `product` by the list operation.
	Product string `json:"product" api:"required"`
	// `true`: explicitly no cap.
	//
	// Any of true.
	Unlimited bool `json:"unlimited,omitzero" api:"required"`
	// Why the limit is set or changed, kept for audit.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// `daily` is the current UTC day; `monthly` is the current UTC calendar month.
	//
	// Any of "daily", "monthly".
	Period SpendLimitPeriod `json:"period,omitzero"`
	paramObj
}

func (r SpendLimitNewParamsBodyCreateSpendLimitUnlimited) MarshalJSON() (data []byte, err error) {
	type shadow SpendLimitNewParamsBodyCreateSpendLimitUnlimited
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendLimitNewParamsBodyCreateSpendLimitUnlimited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SpendLimitNewParamsBodyCreateSpendLimitUnlimited](
		"unlimited", true,
	)
}

type SpendLimitUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. A new
	// limit in USD.
	OfAmount *SpendLimitUpdateParamsBodyUpdateSpendLimitWithAmount `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Explicitly no cap.
	OfUnlimited *SpendLimitUpdateParamsBodyUpdateSpendLimitUnlimited `json:",inline"`

	// Limit period. Defaults to `daily`; send it explicitly.
	//
	// Any of "daily", "monthly".
	Period SpendLimitPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

func (u SpendLimitUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAmount, u.OfUnlimited)
}
func (r *SpendLimitUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SpendLimitUpdateParams]'s query parameters as `url.Values`.
func (r SpendLimitUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A new limit in USD.
//
// The property Amount is required.
type SpendLimitUpdateParamsBodyUpdateSpendLimitWithAmount struct {
	// Limit in USD. `0` blocks at the first cent of spend.
	Amount float64 `json:"amount" api:"required"`
	// Why the limit is set or changed, kept for audit.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Optional; only `false` is allowed together with `amount`.
	//
	// Any of false.
	Unlimited bool `json:"unlimited,omitzero"`
	paramObj
}

func (r SpendLimitUpdateParamsBodyUpdateSpendLimitWithAmount) MarshalJSON() (data []byte, err error) {
	type shadow SpendLimitUpdateParamsBodyUpdateSpendLimitWithAmount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendLimitUpdateParamsBodyUpdateSpendLimitWithAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SpendLimitUpdateParamsBodyUpdateSpendLimitWithAmount](
		"unlimited", false,
	)
}

// Explicitly no cap.
//
// The property Unlimited is required.
type SpendLimitUpdateParamsBodyUpdateSpendLimitUnlimited struct {
	// `true`: explicitly no cap.
	//
	// Any of true.
	Unlimited bool `json:"unlimited,omitzero" api:"required"`
	// Why the limit is set or changed, kept for audit.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r SpendLimitUpdateParamsBodyUpdateSpendLimitUnlimited) MarshalJSON() (data []byte, err error) {
	type shadow SpendLimitUpdateParamsBodyUpdateSpendLimitUnlimited
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendLimitUpdateParamsBodyUpdateSpendLimitUnlimited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SpendLimitUpdateParamsBodyUpdateSpendLimitUnlimited](
		"unlimited", true,
	)
}

type SpendLimitDeleteParams struct {
	// Why the limit is removed, kept for audit. At most 500 characters.
	Reason param.Opt[string] `query:"reason,omitzero" json:"-"`
	// Limit period. Defaults to `daily`; send it explicitly.
	//
	// Any of "daily", "monthly".
	Period SpendLimitPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendLimitDeleteParams]'s query parameters as `url.Values`.
func (r SpendLimitDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
