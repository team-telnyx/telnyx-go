// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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

// EmailDomainService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailDomainService] method instead.
type EmailDomainService struct {
	Options []option.RequestOption
	// Per-domain webhook endpoints with event subscriptions
	Webhooks EmailDomainWebhookService
}

// NewEmailDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailDomainService(opts ...option.RequestOption) (r EmailDomainService) {
	r = EmailDomainService{}
	r.Options = opts
	r.Webhooks = NewEmailDomainWebhookService(opts...)
	return
}

// Registers a domain for email sending and optional inbound delivery. The response
// includes the domain configuration and current verification state.
func (r *EmailDomainService) New(ctx context.Context, body EmailDomainNewParams, opts ...option.RequestOption) (res *EmailDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Shared (`type: shared`) Telnyx-managed domains are included/readable for every
// account, in addition to the account's own custom domains.
func (r *EmailDomainService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable settings for an existing email domain, including inbound
// delivery and tracking configuration. Shared domains are read-only for non-owner
// accounts.
func (r *EmailDomainService) Update(ctx context.Context, id string, body EmailDomainUpdateParams, opts ...option.RequestOption) (res *EmailDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Shared (`type: shared`) Telnyx-managed domains are included/readable for every
// account, in addition to the account's own custom domains.
func (r *EmailDomainService) List(ctx context.Context, query EmailDomainListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EmailDomain], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "email_domains"
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

// Shared (`type: shared`) Telnyx-managed domains are included/readable for every
// account, in addition to the account's own custom domains.
func (r *EmailDomainService) ListAutoPaging(ctx context.Context, query EmailDomainListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EmailDomain] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an email domain configuration. Verified domains require `force=true`,
// and shared domains are read-only for non-owner accounts.
func (r *EmailDomainService) Delete(ctx context.Context, id string, body EmailDomainDeleteParams, opts ...option.RequestOption) (res *EmailDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Returns the DNS records Telnyx generated for domain ownership and DKIM
// verification, plus MX records when inbound delivery is enabled.
func (r *EmailDomainService) GetDNSRecords(ctx context.Context, domainID string, opts ...option.RequestOption) (res *EmailDomainGetDNSRecordsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/dns_records", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a summary of domain health including verification status and usability.
func (r *EmailDomainService) GetHealth(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailDomainGetHealthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/health", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Checks the published DNS records against the records required for the email
// domain and returns the latest verification results.
func (r *EmailDomainService) Verify(ctx context.Context, domainID string, opts ...option.RequestOption) (res *EmailDomainResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/verify", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type DNSRecord struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Host string `json:"host" api:"required"`
	// Any of "ownership", "spf", "dkim", "dmarc", "mx".
	Purpose DNSRecordPurpose `json:"purpose" api:"required"`
	// Any of "TXT", "MX".
	RecordType DNSRecordRecordType `json:"record_type" api:"required"`
	Required   bool                `json:"required" api:"required"`
	// Any of "pending", "verified", "failed", "not_required".
	Status      DNSRecordStatus `json:"status" api:"required"`
	Value       string          `json:"value" api:"required"`
	ActualValue string          `json:"actual_value" api:"nullable"`
	Priority    int64           `json:"priority" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Host        respjson.Field
		Purpose     respjson.Field
		RecordType  respjson.Field
		Required    respjson.Field
		Status      respjson.Field
		Value       respjson.Field
		ActualValue respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSRecord) RawJSON() string { return r.JSON.raw }
func (r *DNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordPurpose string

const (
	DNSRecordPurposeOwnership DNSRecordPurpose = "ownership"
	DNSRecordPurposeSpf       DNSRecordPurpose = "spf"
	DNSRecordPurposeDkim      DNSRecordPurpose = "dkim"
	DNSRecordPurposeDmarc     DNSRecordPurpose = "dmarc"
	DNSRecordPurposeMx        DNSRecordPurpose = "mx"
)

type DNSRecordRecordType string

const (
	DNSRecordRecordTypeTxt DNSRecordRecordType = "TXT"
	DNSRecordRecordTypeMx  DNSRecordRecordType = "MX"
)

type DNSRecordStatus string

const (
	DNSRecordStatusPending     DNSRecordStatus = "pending"
	DNSRecordStatusVerified    DNSRecordStatus = "verified"
	DNSRecordStatusFailed      DNSRecordStatus = "failed"
	DNSRecordStatusNotRequired DNSRecordStatus = "not_required"
)

type DomainsTrackingSettings struct {
	// Rewrite HTML links through a tracking redirect to record click events.
	ClickTracking bool `json:"click_tracking"`
	// Inject a tracking pixel into HTML messages to record open events.
	OpenTracking bool `json:"open_tracking"`
	// Add RFC 8058 List-Unsubscribe headers with a signed one-click unsubscribe URL.
	// Enabled by default; Gmail/Yahoo bulk-sender rules require one-click unsubscribe
	// support.
	UnsubscribeTracking bool `json:"unsubscribe_tracking"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClickTracking       respjson.Field
		OpenTracking        respjson.Field
		UnsubscribeTracking respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainsTrackingSettings) RawJSON() string { return r.JSON.raw }
func (r *DomainsTrackingSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DomainsTrackingSettings to a DomainsTrackingSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DomainsTrackingSettingsParam.Overrides()
func (r DomainsTrackingSettings) ToParam() DomainsTrackingSettingsParam {
	return param.Override[DomainsTrackingSettingsParam](json.RawMessage(r.RawJSON()))
}

type DomainsTrackingSettingsParam struct {
	// Rewrite HTML links through a tracking redirect to record click events.
	ClickTracking param.Opt[bool] `json:"click_tracking,omitzero"`
	// Inject a tracking pixel into HTML messages to record open events.
	OpenTracking param.Opt[bool] `json:"open_tracking,omitzero"`
	// Add RFC 8058 List-Unsubscribe headers with a signed one-click unsubscribe URL.
	// Enabled by default; Gmail/Yahoo bulk-sender rules require one-click unsubscribe
	// support.
	UnsubscribeTracking param.Opt[bool] `json:"unsubscribe_tracking,omitzero"`
	paramObj
}

func (r DomainsTrackingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow DomainsTrackingSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainsTrackingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DMARC policy for a sending domain. Drives the recommended \_dmarc.<domain> TXT
// record. DMARC is advisory and never blocks sending. When omitted or null, the
// domain uses the advisory default (v=DMARC1; p=none;
// rua=mailto:dmarc@telnyx.com).
type EmailDmarcPolicy struct {
	// Policy applied to messages that fail alignment.
	//
	// Any of "none", "quarantine", "reject".
	P EmailDmarcPolicyP `json:"p"`
	// Percentage of messages the policy applies to. Omitted from the record when 100.
	Pct int64 `json:"pct"`
	// URI for aggregate reports. Defaults to the Telnyx address when absent; null
	// omits it.
	Rua string `json:"rua" api:"nullable"`
	// Policy for subdomains. Omitted from the record when null.
	//
	// Any of "none", "quarantine", "reject".
	Sp EmailDmarcPolicySp `json:"sp" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		P           respjson.Field
		Pct         respjson.Field
		Rua         respjson.Field
		Sp          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDmarcPolicy) RawJSON() string { return r.JSON.raw }
func (r *EmailDmarcPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailDmarcPolicy to a EmailDmarcPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailDmarcPolicyParam.Overrides()
func (r EmailDmarcPolicy) ToParam() EmailDmarcPolicyParam {
	return param.Override[EmailDmarcPolicyParam](json.RawMessage(r.RawJSON()))
}

// Policy applied to messages that fail alignment.
type EmailDmarcPolicyP string

const (
	EmailDmarcPolicyPNone       EmailDmarcPolicyP = "none"
	EmailDmarcPolicyPQuarantine EmailDmarcPolicyP = "quarantine"
	EmailDmarcPolicyPReject     EmailDmarcPolicyP = "reject"
)

// Policy for subdomains. Omitted from the record when null.
type EmailDmarcPolicySp string

const (
	EmailDmarcPolicySpNone       EmailDmarcPolicySp = "none"
	EmailDmarcPolicySpQuarantine EmailDmarcPolicySp = "quarantine"
	EmailDmarcPolicySpReject     EmailDmarcPolicySp = "reject"
)

// DMARC policy for a sending domain. Drives the recommended \_dmarc.<domain> TXT
// record. DMARC is advisory and never blocks sending. When omitted or null, the
// domain uses the advisory default (v=DMARC1; p=none;
// rua=mailto:dmarc@telnyx.com).
type EmailDmarcPolicyParam struct {
	// URI for aggregate reports. Defaults to the Telnyx address when absent; null
	// omits it.
	Rua param.Opt[string] `json:"rua,omitzero"`
	// Percentage of messages the policy applies to. Omitted from the record when 100.
	Pct param.Opt[int64] `json:"pct,omitzero"`
	// Policy for subdomains. Omitted from the record when null.
	//
	// Any of "none", "quarantine", "reject".
	Sp EmailDmarcPolicySp `json:"sp,omitzero"`
	// Policy applied to messages that fail alignment.
	//
	// Any of "none", "quarantine", "reject".
	P EmailDmarcPolicyP `json:"p,omitzero"`
	paramObj
}

func (r EmailDmarcPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailDmarcPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDmarcPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomain struct {
	ID        string          `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time       `json:"created_at" api:"required" format:"date-time"`
	Dkim      EmailDomainDkim `json:"dkim" api:"required"`
	// DMARC policy for a sending domain. Drives the recommended \_dmarc.<domain> TXT
	// record. DMARC is advisory and never blocks sending. When omitted or null, the
	// domain uses the advisory default (v=DMARC1; p=none;
	// rua=mailto:dmarc@telnyx.com).
	DmarcPolicy EmailDmarcPolicy   `json:"dmarc_policy" api:"required"`
	DNSRecords  []DNSRecord        `json:"dns_records" api:"required"`
	Domain      string             `json:"domain" api:"required"`
	Inbound     EmailDomainInbound `json:"inbound" api:"required"`
	// Any of "email_domain".
	RecordType EmailDomainRecordType `json:"record_type" api:"required"`
	// Any of "pending", "verifying", "verified", "failed", "degraded", "suspended".
	Status   EmailDomainStatus       `json:"status" api:"required"`
	Tracking DomainsTrackingSettings `json:"tracking" api:"required"`
	// Domain type. `custom` domains are account-owned (BYOD). `shared` domains are
	// Telnyx-managed, visible to and usable by ALL accounts for sending, but
	// read-only: only the owning (system) account may modify, verify, or delete them;
	// other accounts receive 403 (code 10008).
	//
	// Any of "custom", "shared", "shared_inbound".
	Type             EmailDomainType         `json:"type" api:"required"`
	UpdatedAt        time.Time               `json:"updated_at" api:"required" format:"date-time"`
	UsableForInbound bool                    `json:"usable_for_inbound" api:"required"`
	UsableForSending bool                    `json:"usable_for_sending" api:"required"`
	Verification     EmailDomainVerification `json:"verification" api:"required"`
	// Sender reputation for this domain (present on all domain responses).
	Reputation EmailDomainReputation `json:"reputation"`
	VerifiedAt time.Time             `json:"verified_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Dkim             respjson.Field
		DmarcPolicy      respjson.Field
		DNSRecords       respjson.Field
		Domain           respjson.Field
		Inbound          respjson.Field
		RecordType       respjson.Field
		Status           respjson.Field
		Tracking         respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		UsableForInbound respjson.Field
		UsableForSending respjson.Field
		Verification     respjson.Field
		Reputation       respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomain) RawJSON() string { return r.JSON.raw }
func (r *EmailDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainDkim struct {
	Active bool `json:"active" api:"required"`
	// Any of "rsa-sha256".
	Algorithm string `json:"algorithm" api:"required"`
	// Any of 2048.
	KeyLength int64     `json:"key_length" api:"required"`
	RotatedAt time.Time `json:"rotated_at" api:"required" format:"date-time"`
	Selector  string    `json:"selector" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Algorithm   respjson.Field
		KeyLength   respjson.Field
		RotatedAt   respjson.Field
		Selector    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainDkim) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainDkim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainInbound struct {
	CatchAll   bool `json:"catch_all" api:"required"`
	Enabled    bool `json:"enabled" api:"required"`
	MxRequired bool `json:"mx_required" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CatchAll    respjson.Field
		Enabled     respjson.Field
		MxRequired  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainInbound) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainRecordType string

const (
	EmailDomainRecordTypeEmailDomain EmailDomainRecordType = "email_domain"
)

// Sender reputation for this domain (present on all domain responses).
type EmailDomainReputation struct {
	// Reputation band, e.g. good/warn/poor.
	Band       string         `json:"band"`
	Breakdown  map[string]any `json:"breakdown"`
	ComputedAt time.Time      `json:"computed_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Band        respjson.Field
		Breakdown   respjson.Field
		ComputedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainReputation) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainReputation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainResponse struct {
	Data EmailDomain `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainStatus string

const (
	EmailDomainStatusPending   EmailDomainStatus = "pending"
	EmailDomainStatusVerifying EmailDomainStatus = "verifying"
	EmailDomainStatusVerified  EmailDomainStatus = "verified"
	EmailDomainStatusFailed    EmailDomainStatus = "failed"
	EmailDomainStatusDegraded  EmailDomainStatus = "degraded"
	EmailDomainStatusSuspended EmailDomainStatus = "suspended"
)

type EmailDomainType string

const (
	EmailDomainTypeCustom        EmailDomainType = "custom"
	EmailDomainTypeShared        EmailDomainType = "shared"
	EmailDomainTypeSharedInbound EmailDomainType = "shared_inbound"
)

type EmailDomainVerification struct {
	// Any of "pending", "verified", "failed".
	Dkim EmailDomainVerificationDkim `json:"dkim" api:"required"`
	// Any of "missing_optional", "verified", "failed".
	Dmarc EmailDomainVerificationDmarc `json:"dmarc" api:"required"`
	// Any of "not_required", "pending", "verified", "failed".
	Mx EmailDomainVerificationMx `json:"mx" api:"required"`
	// Any of "pending", "verified", "not_required".
	Ownership EmailDomainVerificationOwnership `json:"ownership" api:"required"`
	// Any of "missing_optional", "verified", "failed", "not_required".
	Spf EmailDomainVerificationSpf `json:"spf" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		Dmarc       respjson.Field
		Mx          respjson.Field
		Ownership   respjson.Field
		Spf         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainVerification) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainVerificationDkim string

const (
	EmailDomainVerificationDkimPending  EmailDomainVerificationDkim = "pending"
	EmailDomainVerificationDkimVerified EmailDomainVerificationDkim = "verified"
	EmailDomainVerificationDkimFailed   EmailDomainVerificationDkim = "failed"
)

type EmailDomainVerificationDmarc string

const (
	EmailDomainVerificationDmarcMissingOptional EmailDomainVerificationDmarc = "missing_optional"
	EmailDomainVerificationDmarcVerified        EmailDomainVerificationDmarc = "verified"
	EmailDomainVerificationDmarcFailed          EmailDomainVerificationDmarc = "failed"
)

type EmailDomainVerificationMx string

const (
	EmailDomainVerificationMxNotRequired EmailDomainVerificationMx = "not_required"
	EmailDomainVerificationMxPending     EmailDomainVerificationMx = "pending"
	EmailDomainVerificationMxVerified    EmailDomainVerificationMx = "verified"
	EmailDomainVerificationMxFailed      EmailDomainVerificationMx = "failed"
)

type EmailDomainVerificationOwnership string

const (
	EmailDomainVerificationOwnershipPending     EmailDomainVerificationOwnership = "pending"
	EmailDomainVerificationOwnershipVerified    EmailDomainVerificationOwnership = "verified"
	EmailDomainVerificationOwnershipNotRequired EmailDomainVerificationOwnership = "not_required"
)

type EmailDomainVerificationSpf string

const (
	EmailDomainVerificationSpfMissingOptional EmailDomainVerificationSpf = "missing_optional"
	EmailDomainVerificationSpfVerified        EmailDomainVerificationSpf = "verified"
	EmailDomainVerificationSpfFailed          EmailDomainVerificationSpf = "failed"
	EmailDomainVerificationSpfNotRequired     EmailDomainVerificationSpf = "not_required"
)

type EmailDomainGetDNSRecordsResponse struct {
	Data []DNSRecord `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainGetDNSRecordsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainGetDNSRecordsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainGetHealthResponse struct {
	Data EmailDomainGetHealthResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainGetHealthResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainGetHealthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainGetHealthResponseData struct {
	// Unique identifier for the email domain
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of the last health check
	CheckedAt time.Time `json:"checked_at" api:"required" format:"date-time"`
	// Record type discriminator
	//
	// Any of "email_domain_health".
	RecordType string `json:"record_type" api:"required"`
	// Current domain status
	//
	// Any of "pending", "verifying", "verified", "failed", "degraded", "suspended".
	Status string `json:"status" api:"required"`
	// Whether the domain is usable for receiving inbound email
	UsableForInbound bool `json:"usable_for_inbound" api:"required"`
	// Whether the domain is usable for sending email
	UsableForSending bool                    `json:"usable_for_sending" api:"required"`
	Verification     EmailDomainVerification `json:"verification" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CheckedAt        respjson.Field
		RecordType       respjson.Field
		Status           respjson.Field
		UsableForInbound respjson.Field
		UsableForSending respjson.Field
		Verification     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainGetHealthResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainGetHealthResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainNewParams struct {
	Domain string `json:"domain" api:"required"`
	// Enable inbound routing for this domain
	InboundEnabled param.Opt[bool] `json:"inbound_enabled,omitzero"`
	// DMARC policy for a sending domain. Drives the recommended \_dmarc.<domain> TXT
	// record. DMARC is advisory and never blocks sending. When omitted or null, the
	// domain uses the advisory default (v=DMARC1; p=none;
	// rua=mailto:dmarc@telnyx.com).
	DmarcPolicy EmailDmarcPolicyParam        `json:"dmarc_policy,omitzero"`
	Tracking    DomainsTrackingSettingsParam `json:"tracking,omitzero"`
	paramObj
}

func (r EmailDomainNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDomainNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainUpdateParams struct {
	// Enable or disable inbound routing for this domain
	InboundEnabled param.Opt[bool] `json:"inbound_enabled,omitzero"`
	// DMARC policy for a sending domain. Drives the recommended \_dmarc.<domain> TXT
	// record. DMARC is advisory and never blocks sending. When omitted or null, the
	// domain uses the advisory default (v=DMARC1; p=none;
	// rua=mailto:dmarc@telnyx.com).
	DmarcPolicy EmailDmarcPolicyParam        `json:"dmarc_policy,omitzero"`
	Tracking    DomainsTrackingSettingsParam `json:"tracking,omitzero"`
	paramObj
}

func (r EmailDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDomainUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDomainUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainListParams struct {
	// Partial match on domain name (case-insensitive)
	FilterDomain param.Opt[string] `query:"filter[domain],omitzero" json:"-"`
	// Filter by profile UUID
	FilterProfileID param.Opt[string] `query:"filter[profile_id],omitzero" format:"uuid" json:"-"`
	// Filter domains by whether they can currently receive inbound email.
	FilterUsableForInbound param.Opt[bool] `query:"filter[usable_for_inbound],omitzero" json:"-"`
	// Filter domains by whether they can currently be used to send email.
	FilterUsableForSending param.Opt[bool] `query:"filter[usable_for_sending],omitzero" json:"-"`
	// Cursor for records after the provided value (cursor pagination)
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Cursor for records before the provided value (cursor pagination)
	PageBefore param.Opt[string] `query:"page[before],omitzero" json:"-"`
	// Page number to return (offset pagination)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of records per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter domains by verification status: pending, verifying, verified, failed,
	// degraded, or suspended.
	//
	// Any of "pending", "verifying", "verified", "failed", "degraded", "suspended".
	FilterStatus EmailDomainStatus `query:"filter[status],omitzero" json:"-"`
	// Filter domains by type: custom, shared, or shared_inbound.
	//
	// Any of "custom", "shared", "shared_inbound".
	FilterType EmailDomainType `query:"filter[type],omitzero" json:"-"`
	// Field to sort by. Prefix with `-` for descending order.
	//
	// Any of "created_at", "-created_at", "domain", "-domain".
	Sort EmailDomainListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailDomainListParams]'s query parameters as `url.Values`.
func (r EmailDomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by. Prefix with `-` for descending order.
type EmailDomainListParamsSort string

const (
	EmailDomainListParamsSortCreatedAt     EmailDomainListParamsSort = "created_at"
	EmailDomainListParamsSortCreatedAtDesc EmailDomainListParamsSort = "-created_at"
	EmailDomainListParamsSortDomain        EmailDomainListParamsSort = "domain"
	EmailDomainListParamsSortMinusDomain   EmailDomainListParamsSort = "-domain"
)

type EmailDomainDeleteParams struct {
	// Required as true when deleting verified domains
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailDomainDeleteParams]'s query parameters as
// `url.Values`.
func (r EmailDomainDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
