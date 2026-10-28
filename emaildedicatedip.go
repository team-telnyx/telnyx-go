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

// Purchase and manage dedicated sending IP addresses for your Telnyx email
// traffic.
//
// EmailDedicatedIPService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailDedicatedIPService] method instead.
type EmailDedicatedIPService struct {
	Options []option.RequestOption
	// Purchase and manage dedicated sending IP addresses for your Telnyx email
	// traffic.
	Actions EmailDedicatedIPActionService
}

// NewEmailDedicatedIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailDedicatedIPService(opts ...option.RequestOption) (r EmailDedicatedIPService) {
	r = EmailDedicatedIPService{}
	r.Options = opts
	r.Actions = NewEmailDedicatedIPActionService(opts...)
	return
}

// Purchase a dedicated sending IP. Requires a verified, enterprise, or legacy
// account, and an `Idempotency-Key` header: a request without the header returns
// `400` (code `10015`) and allocates nothing. Telnyx allocates the next available
// IP from stock. The optional `region` and `ip_version` fields each restrict which
// stock IP is allocated; they never change the allocated IP's metadata. An omitted
// field applies no restriction for that field, and omitting the body applies
// neither restriction. An unsupported selector value returns `422`; a supported
// value that matches no available IP returns `429` and records the request on the
// waitlist. Billing: $14/month/IP, charged as a monthly recurring charge (MRC)
// from purchase. The IP is returned in `provisioning` status with a Telnyx-managed
// PTR hostname and carries no mail until you call `start_warmup` and add it to a
// pool that routes mail.
func (r *EmailDedicatedIPService) New(ctx context.Context, params EmailDedicatedIPNewParams, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "email/dedicated_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a single dedicated IP by its identifier, including warm-up state,
// billing state, and the DNS/PTR provisioning checklist.
func (r *EmailDedicatedIPService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Move the IP to another pool, or pause/resume its warm-up ramp. One change
// applies per request. `pool_id` takes precedence over warm-up changes. Otherwise,
// use `warmup_paused` or `warmup.paused`; a Boolean `warmup_paused` takes
// precedence. Moving the last IP out of a pool requires `force: true` and
// `drain: true`. The PTR hostname is managed by Telnyx: a request that includes
// `ptr_hostname` returns `422` unless `pool_id` or a warm-up field is present, in
// which case `ptr_hostname` is ignored.
func (r *EmailDedicatedIPService) Update(ctx context.Context, id string, body EmailDedicatedIPUpdateParams, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the dedicated sending IPs owned by the account, newest first. Narrow the
// list with `filter[status]`, `filter[pool_id]`, and `filter[region]`; supplied
// filters combine, so an IP must match every one. An empty value, or a nested or
// array value such as `filter[status][]=active`, applies no filter, and other
// `filter[...]` keys are ignored. A value that no IP can match (an unknown status,
// a `pool_id` that is not a UUID, or a value containing a NUL byte) returns an
// empty list, not an error.
func (r *EmailDedicatedIPService) List(ctx context.Context, query EmailDedicatedIPListParams, opts ...option.RequestOption) (res *EmailDedicatedIPListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email/dedicated_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Release the dedicated IP to Telnyx. The IP stops carrying mail immediately,
// leaves its pool, and moves to `released`. Telnyx sets `quarantine_until` to 30
// days after release. Release requests that the MRC stop asynchronously. Check
// `billing`: `stopping` means the stop is pending, `stopped` means it completed,
// and `stop_failed` means the charge remains. If the charge was never created or
// its creation has not completed, `billing` keeps showing the creation status
// (`not_started`, `pending`, `in_progress`, or `failed`). Any status except
// `released` can be released. Releasing the last IP in a pool requires
// `force=true` and `drain=true`.
func (r *EmailDedicatedIPService) Delete(ctx context.Context, id string, params EmailDedicatedIPDeleteParams, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// A dedicated sending IP address owned by the account.
type DedicatedIP struct {
	ID string `json:"id" format:"uuid"`
	// Monthly recurring charge (MRC) state for the IP's current allocation.
	Billing   DedicatedIPBilling `json:"billing"`
	CreatedAt time.Time          `json:"created_at" format:"date-time"`
	// When the PTR hostname was forward-confirmed (PTR → hostname → A record → this
	// IP). The IP carries mail only while this is set.
	DNSConfirmedAt time.Time `json:"dns_confirmed_at" api:"nullable" format:"date-time"`
	// The allocated IPv4/IPv6 address.
	IPAddress string `json:"ip_address"`
	// Any of "ipv4", "ipv6".
	IPVersion DedicatedIPIPVersion `json:"ip_version"`
	PoolID    string               `json:"pool_id" api:"nullable" format:"uuid"`
	// Read-only status of the Telnyx-managed PTR record. No customer action is
	// required to complete it.
	Provisioning DedicatedIPProvisioning `json:"provisioning"`
	// Same value as `provisioning.dns_status`.
	//
	// Any of "not_applicable", "pending_ptr", "pending_dns", "dns_confirmed",
	// "dns_lost".
	ProvisioningStatus DedicatedIPProvisioningStatus `json:"provisioning_status"`
	// Telnyx-managed reverse DNS (PTR) hostname of the IP, in the form
	// `mta-<a-b-c-d>.msgtelnyx.com`. Also used as the EHLO hostname when sending from
	// the IP. Managed by Telnyx and cannot be changed through the API.
	PtrHostname string `json:"ptr_hostname" api:"nullable"`
	// Set on release to 30 days after `released_at`. The address is not returned to
	// stock before this time.
	QuarantineUntil time.Time `json:"quarantine_until" api:"nullable" format:"date-time"`
	// Any of "dedicated_ip".
	RecordType DedicatedIPRecordType `json:"record_type"`
	// Region label of the allocated IP (for example, `ewr-nj1`).
	Region string `json:"region"`
	// When the IP was released. Null until release.
	ReleasedAt time.Time `json:"released_at" api:"nullable" format:"date-time"`
	// Lifecycle status. `provisioning`: purchased and billed, carries no mail until
	// `start_warmup`. `warming_up`: warm-up ramp in progress. `active`: ramp complete.
	// `suspended`: suspended by Telnyx; carries no mail and is still billed.
	// `released`: returned to Telnyx; carries no mail.
	//
	// Any of "provisioning", "warming_up", "active", "suspended", "released".
	Status DedicatedIPStatus `json:"status"`
	// Reason Telnyx recorded when suspending the IP. Null unless `status` is
	// `suspended`.
	SuspensionReason string    `json:"suspension_reason" api:"nullable"`
	UpdatedAt        time.Time `json:"updated_at" format:"date-time"`
	// Warmup ramp state for a dedicated IP.
	Warmup DedicatedIPWarmup `json:"warmup"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Billing            respjson.Field
		CreatedAt          respjson.Field
		DNSConfirmedAt     respjson.Field
		IPAddress          respjson.Field
		IPVersion          respjson.Field
		PoolID             respjson.Field
		Provisioning       respjson.Field
		ProvisioningStatus respjson.Field
		PtrHostname        respjson.Field
		QuarantineUntil    respjson.Field
		RecordType         respjson.Field
		Region             respjson.Field
		ReleasedAt         respjson.Field
		Status             respjson.Field
		SuspensionReason   respjson.Field
		UpdatedAt          respjson.Field
		Warmup             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIP) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly recurring charge (MRC) state for the IP's current allocation.
type DedicatedIPBilling struct {
	// True when `status` is `active`, `stopping`, or `stop_failed`. An IP carries mail
	// only after its charge has been created.
	Charged bool `json:"charged"`
	// `not_started`: no charge recorded for the current allocation. `pending`: the
	// charge is queued. `in_progress`: the charge is being created. `active`: the MRC
	// is active. `stopping`: release requested; the MRC stop is being processed.
	// `stopped`: the MRC has stopped. `stop_failed`: the MRC stop failed and the
	// charge remains. `failed`: the charge could not be created.
	//
	// Any of "not_started", "pending", "in_progress", "active", "stopping", "stopped",
	// "stop_failed", "failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charged     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIPBilling) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIPBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DedicatedIPIPVersion string

const (
	DedicatedIPIPVersionIpv4 DedicatedIPIPVersion = "ipv4"
	DedicatedIPIPVersionIpv6 DedicatedIPIPVersion = "ipv6"
)

// Read-only status of the Telnyx-managed PTR record. No customer action is
// required to complete it.
type DedicatedIPProvisioning struct {
	// Informational progress flags for the Telnyx-managed PTR. They do not gate
	// `start_warmup`; use `dns_status` to check whether the IP can carry mail.
	Checklist DedicatedIPProvisioningChecklist `json:"checklist"`
	// `dns_confirmed`: the PTR is forward-confirmed and `dns_confirmed_at` is set.
	// `pending_ptr`: the PTR record is not yet in place. `pending_dns`: the PTR record
	// is waiting for forward confirmation. `dns_lost`: a previously confirmed PTR no
	// longer confirms, or the PTR request was abandoned before it ever confirmed
	// (`ptr_request_state` is `abandoned`); the IP carries no mail until DNS is
	// confirmed. `not_applicable`: the IP is released or has no PTR record on file.
	//
	// Any of "not_applicable", "pending_ptr", "pending_dns", "dns_confirmed",
	// "dns_lost".
	DNSStatus string `json:"dns_status"`
	// Reference to the Telnyx PTR change request, when one exists. Informational.
	PtrRequestPrURL string `json:"ptr_request_pr_url" api:"nullable"`
	// State of the Telnyx PTR change: `pending`, `pr_open`, `merged`, `confirmed`,
	// `failed`, or `abandoned`. Informational.
	//
	// Any of "pending", "pr_open", "merged", "confirmed", "failed", "abandoned".
	PtrRequestState string `json:"ptr_request_state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checklist       respjson.Field
		DNSStatus       respjson.Field
		PtrRequestPrURL respjson.Field
		PtrRequestState respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIPProvisioning) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIPProvisioning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Informational progress flags for the Telnyx-managed PTR. They do not gate
// `start_warmup`; use `dns_status` to check whether the IP can carry mail.
type DedicatedIPProvisioningChecklist struct {
	ForwardConfirmed bool `json:"forward_confirmed"`
	PtrAuthored      bool `json:"ptr_authored"`
	PtrMerged        bool `json:"ptr_merged"`
	PtrRequested     bool `json:"ptr_requested"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ForwardConfirmed respjson.Field
		PtrAuthored      respjson.Field
		PtrMerged        respjson.Field
		PtrRequested     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIPProvisioningChecklist) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIPProvisioningChecklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Same value as `provisioning.dns_status`.
type DedicatedIPProvisioningStatus string

const (
	DedicatedIPProvisioningStatusNotApplicable DedicatedIPProvisioningStatus = "not_applicable"
	DedicatedIPProvisioningStatusPendingPtr    DedicatedIPProvisioningStatus = "pending_ptr"
	DedicatedIPProvisioningStatusPendingDNS    DedicatedIPProvisioningStatus = "pending_dns"
	DedicatedIPProvisioningStatusDNSConfirmed  DedicatedIPProvisioningStatus = "dns_confirmed"
	DedicatedIPProvisioningStatusDNSLost       DedicatedIPProvisioningStatus = "dns_lost"
)

type DedicatedIPRecordType string

const (
	DedicatedIPRecordTypeDedicatedIP DedicatedIPRecordType = "dedicated_ip"
)

// Lifecycle status. `provisioning`: purchased and billed, carries no mail until
// `start_warmup`. `warming_up`: warm-up ramp in progress. `active`: ramp complete.
// `suspended`: suspended by Telnyx; carries no mail and is still billed.
// `released`: returned to Telnyx; carries no mail.
type DedicatedIPStatus string

const (
	DedicatedIPStatusProvisioning DedicatedIPStatus = "provisioning"
	DedicatedIPStatusWarmingUp    DedicatedIPStatus = "warming_up"
	DedicatedIPStatusActive       DedicatedIPStatus = "active"
	DedicatedIPStatusSuspended    DedicatedIPStatus = "suspended"
	DedicatedIPStatusReleased     DedicatedIPStatus = "released"
)

// Warmup ramp state for a dedicated IP.
type DedicatedIPWarmup struct {
	// Daily cap of the current ramp stage. Null until the scheduler arms the first
	// stage, and cleared when the ramp completes, automatically or through
	// `complete_warmup`. On an `active` IP this value is informational and can keep an
	// earlier stage cap after a suspension is lifted; it does not limit sending.
	DailyCap int64 `json:"daily_cap" api:"nullable"`
	// True while the ramp is paused.
	Paused bool `json:"paused"`
	// When the ramp was paused. Null when not paused.
	PausedAt time.Time `json:"paused_at" api:"nullable" format:"date-time"`
	// Ramp progress (0-100): the share of the 28 send-day schedule completed. Only UTC
	// days on which the IP sent mail count. Set to 100 when the ramp completes,
	// automatically or through `complete_warmup`. On an `active` IP this value is
	// informational and can keep an earlier value after a suspension is lifted.
	Progress int64 `json:"progress"`
	// When `start_warmup` was called. Null before warm-up starts and after release.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Today's warmup allowance report card (current UTC day). Null outside the warmup
	// window (before start_warmup, after release/suspension). `active` IPs report the
	// ramp-complete allowance, except as described for `cap`.
	Today DedicatedIPWarmupToday `json:"today" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyCap    respjson.Field
		Paused      respjson.Field
		PausedAt    respjson.Field
		Progress    respjson.Field
		StartedAt   respjson.Field
		Today       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIPWarmup) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIPWarmup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Today's warmup allowance report card (current UTC day). Null outside the warmup
// window (before start_warmup, after release/suspension). `active` IPs report the
// ramp-complete allowance, except as described for `cap`.
type DedicatedIPWarmupToday struct {
	// The day's allowance cap. Null when the ramp has not opened yet (day 0: the
	// scheduler has not armed a stage cap) — distinct from a spent day, which reports
	// the cap with remaining 0. `active` IPs report `100000000`, the ramp-complete
	// cap; `complete_warmup` applies it to the current day immediately. A day's
	// allowance created before the IP became `active` can keep its stored `cap` and
	// `remaining` until the next first-attempt send refreshes them.
	Cap int64 `json:"cap" api:"nullable"`
	// The UTC day this snapshot describes (warmup caps reset at UTC midnight).
	Day time.Time `json:"day" format:"date"`
	// First-attempt sends released from warmup hold so far today. Delivery retries
	// never lease allowance and are not counted.
	Granted int64 `json:"granted"`
	// Headroom left today. A queued send held by the daily cap is explained by
	// remaining == 0; null mirrors cap (ramp not open yet). On an `active` IP whose
	// `cap` is below `100000000`, remaining == 0 is a stored allowance and does not
	// block the next eligible first-attempt send.
	Remaining int64 `json:"remaining" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cap         respjson.Field
		Day         respjson.Field
		Granted     respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIPWarmupToday) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIPWarmupToday) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DedicatedIPShowResponse struct {
	// A dedicated sending IP address owned by the account.
	Data DedicatedIP `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedIPShowResponse) RawJSON() string { return r.JSON.raw }
func (r *DedicatedIPShowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDedicatedIPListResponse struct {
	Data []DedicatedIP `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDedicatedIPListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDedicatedIPListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDedicatedIPNewParams struct {
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// Allocate an IP of this address family. Supported value: `ipv4`. Any other value,
	// including `ipv6` or `null`, returns `422`.
	//
	// Any of "ipv4".
	IPVersion EmailDedicatedIPNewParamsIPVersion `json:"ip_version,omitzero"`
	// Allocate an IP in this region. Supported value: `ewr-nj1`. Any other value,
	// including an empty string or `null`, returns `422`.
	//
	// Any of "ewr-nj1".
	Region EmailDedicatedIPNewParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r EmailDedicatedIPNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDedicatedIPNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDedicatedIPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allocate an IP of this address family. Supported value: `ipv4`. Any other value,
// including `ipv6` or `null`, returns `422`.
type EmailDedicatedIPNewParamsIPVersion string

const (
	EmailDedicatedIPNewParamsIPVersionIpv4 EmailDedicatedIPNewParamsIPVersion = "ipv4"
)

// Allocate an IP in this region. Supported value: `ewr-nj1`. Any other value,
// including an empty string or `null`, returns `422`.
type EmailDedicatedIPNewParamsRegion string

const (
	EmailDedicatedIPNewParamsRegionEwrNj1 EmailDedicatedIPNewParamsRegion = "ewr-nj1"
)

type EmailDedicatedIPUpdateParams struct {
	// Set together with `force: true` to move the last IP out of its current pool.
	// Acknowledges that mail routed to the emptied pool is held. Used only with
	// `pool_id`.
	Drain param.Opt[bool] `json:"drain,omitzero"`
	// Set together with `drain: true` to move the last IP out of its current pool.
	// Used only with `pool_id`.
	Force param.Opt[bool] `json:"force,omitzero"`
	// Move the IP to this pool, replacing its current membership. Use the IP's current
	// pool to change only its weight. To take an IP out of a pool, use
	// `DELETE /email/ip_pools/{id}/ips/{ip_id}`.
	PoolID param.Opt[string] `json:"pool_id,omitzero" format:"uuid"`
	// Pause (`true`) or resume (`false`) the warm-up ramp; equivalent to the
	// `pause_warmup` and `resume_warmup` actions. A Boolean value takes precedence
	// over `warmup.paused`. When a pause field is present and `pool_id` is absent,
	// supplying neither pause field as a Boolean invokes resume.
	WarmupPaused param.Opt[bool] `json:"warmup_paused,omitzero"`
	// Routing weight in the target pool. Used only with `pool_id`.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	// Alias for `warmup_paused`.
	Warmup EmailDedicatedIPUpdateParamsWarmup `json:"warmup,omitzero"`
	paramObj
}

func (r EmailDedicatedIPUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDedicatedIPUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDedicatedIPUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Alias for `warmup_paused`.
type EmailDedicatedIPUpdateParamsWarmup struct {
	// Pause (`true`) or resume (`false`) the warm-up ramp. Used only when
	// `warmup_paused` is not a Boolean.
	Paused param.Opt[bool] `json:"paused,omitzero"`
	paramObj
}

func (r EmailDedicatedIPUpdateParamsWarmup) MarshalJSON() (data []byte, err error) {
	type shadow EmailDedicatedIPUpdateParamsWarmup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDedicatedIPUpdateParamsWarmup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDedicatedIPListParams struct {
	// Return only IPs in this pool (exact match on `pool_id`). A value that is not a
	// UUID returns an empty list. An empty value applies no filter.
	FilterPoolID param.Opt[string] `query:"filter[pool_id],omitzero" format:"uuid" json:"-"`
	// Return only IPs in this region (exact, case-sensitive match on `region`). An
	// empty value applies no filter.
	FilterRegion param.Opt[string] `query:"filter[region],omitzero" json:"-"`
	// Return only IPs in this lifecycle status (exact match). Any other value returns
	// an empty list. An empty value applies no filter.
	//
	// Any of "provisioning", "warming_up", "active", "suspended", "released".
	FilterStatus EmailDedicatedIPListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailDedicatedIPListParams]'s query parameters as
// `url.Values`.
func (r EmailDedicatedIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Return only IPs in this lifecycle status (exact match). Any other value returns
// an empty list. An empty value applies no filter.
type EmailDedicatedIPListParamsFilterStatus string

const (
	EmailDedicatedIPListParamsFilterStatusProvisioning EmailDedicatedIPListParamsFilterStatus = "provisioning"
	EmailDedicatedIPListParamsFilterStatusWarmingUp    EmailDedicatedIPListParamsFilterStatus = "warming_up"
	EmailDedicatedIPListParamsFilterStatusActive       EmailDedicatedIPListParamsFilterStatus = "active"
	EmailDedicatedIPListParamsFilterStatusSuspended    EmailDedicatedIPListParamsFilterStatus = "suspended"
	EmailDedicatedIPListParamsFilterStatusReleased     EmailDedicatedIPListParamsFilterStatus = "released"
)

type EmailDedicatedIPDeleteParams struct {
	// Set together with `force=true` to release the last IP in a pool. Acknowledges
	// that mail routed to the emptied pool is held until the pool has a sendable IP
	// again. Also accepted in the JSON body.
	QueryDrain param.Opt[bool] `query:"drain,omitzero" json:"-"`
	// Set together with `drain=true` to release the last IP in a pool. Either flag
	// alone has no effect. Also accepted in the JSON body.
	QueryForce param.Opt[bool] `query:"force,omitzero" json:"-"`
	// Free-text reason recorded with the release. Also accepted in the JSON body.
	QueryReason param.Opt[string] `query:"reason,omitzero" json:"-"`
	// Set together with `force: true` to release the last IP in a pool. Acknowledges
	// that mail routed to the emptied pool is held until the pool has a sendable IP
	// again.
	BodyDrain param.Opt[bool] `json:"drain,omitzero"`
	// Set together with `drain: true` to release the last IP in a pool. Either flag
	// alone has no effect.
	BodyForce param.Opt[bool] `json:"force,omitzero"`
	// Free-text reason recorded with the release.
	BodyReason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r EmailDedicatedIPDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDedicatedIPDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDedicatedIPDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EmailDedicatedIPDeleteParams]'s query parameters as
// `url.Values`.
func (r EmailDedicatedIPDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
