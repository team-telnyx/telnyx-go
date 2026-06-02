// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Look up the live SIP registration status of a UAC connection.
//
// SipRegistrationStatusService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSipRegistrationStatusService] method instead.
type SipRegistrationStatusService struct {
	Options []option.RequestOption
}

// NewSipRegistrationStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSipRegistrationStatusService(opts ...option.RequestOption) (r SipRegistrationStatusService) {
	r = SipRegistrationStatusService{}
	r.Options = opts
	return
}

// Returns the live SIP registration state of a UAC connection: whether it is
// currently registered, when it last registered, and the last response Telnyx
// received from the registrar. Only `uac_external_credential` is supported today.
func (r *SipRegistrationStatusService) Get(ctx context.Context, query SipRegistrationStatusGetParams, opts ...option.RequestOption) (res *SipRegistrationStatusGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sip_registration_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SipRegistrationStatusGetResponse struct {
	// Identifier of the UAC connection.
	ConnectionID string `json:"connection_id"`
	// Human-readable connection name.
	ConnectionName string `json:"connection_name"`
	// The credential type that was looked up.
	//
	// Any of "uac_external_credential".
	CredentialType SipRegistrationStatusGetResponseCredentialType `json:"credential_type"`
	// SIP username used for the registration.
	CredentialUsername string `json:"credential_username"`
	// SIP response from the last registration attempt.
	LastRegistrationResponse string `json:"last_registration_response"`
	// True if the endpoint is currently registered.
	Registered bool `json:"registered"`
	// Detailed registration information reported by the registrar.
	SipRegistrationDetails SipRegistrationStatusGetResponseSipRegistrationDetails `json:"sip_registration_details"`
	// Human-readable registration status derived from the registrar state.
	//
	// Any of "unregistering", "connection_disabled", "standby", "failed", "trying",
	// "registered", "unknown".
	SipRegistrationStatus SipRegistrationStatusGetResponseSipRegistrationStatus `json:"sip_registration_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID             respjson.Field
		ConnectionName           respjson.Field
		CredentialType           respjson.Field
		CredentialUsername       respjson.Field
		LastRegistrationResponse respjson.Field
		Registered               respjson.Field
		SipRegistrationDetails   respjson.Field
		SipRegistrationStatus    respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipRegistrationStatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SipRegistrationStatusGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The credential type that was looked up.
type SipRegistrationStatusGetResponseCredentialType string

const (
	SipRegistrationStatusGetResponseCredentialTypeUacExternalCredential SipRegistrationStatusGetResponseCredentialType = "uac_external_credential"
)

// Detailed registration information reported by the registrar.
type SipRegistrationStatusGetResponseSipRegistrationDetails struct {
	// Number of authentication retries on the last attempt.
	AuthRetries int64 `json:"auth_retries"`
	// Unix timestamp when the current registration expires.
	Expires int64 `json:"expires"`
	// Count of consecutive registration failures.
	Failures int64 `json:"failures"`
	// Unix timestamp of the next scheduled registration action.
	NextActionAt int64 `json:"next_action_at"`
	// SIP URI user@host of the registered contact.
	SipUriUserHost string `json:"sip_uri_user_host"`
	// Registration uptime reported by the registrar.
	Uptime int64 `json:"uptime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthRetries    respjson.Field
		Expires        respjson.Field
		Failures       respjson.Field
		NextActionAt   respjson.Field
		SipUriUserHost respjson.Field
		Uptime         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipRegistrationStatusGetResponseSipRegistrationDetails) RawJSON() string { return r.JSON.raw }
func (r *SipRegistrationStatusGetResponseSipRegistrationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Human-readable registration status derived from the registrar state.
type SipRegistrationStatusGetResponseSipRegistrationStatus string

const (
	SipRegistrationStatusGetResponseSipRegistrationStatusUnregistering      SipRegistrationStatusGetResponseSipRegistrationStatus = "unregistering"
	SipRegistrationStatusGetResponseSipRegistrationStatusConnectionDisabled SipRegistrationStatusGetResponseSipRegistrationStatus = "connection_disabled"
	SipRegistrationStatusGetResponseSipRegistrationStatusStandby            SipRegistrationStatusGetResponseSipRegistrationStatus = "standby"
	SipRegistrationStatusGetResponseSipRegistrationStatusFailed             SipRegistrationStatusGetResponseSipRegistrationStatus = "failed"
	SipRegistrationStatusGetResponseSipRegistrationStatusTrying             SipRegistrationStatusGetResponseSipRegistrationStatus = "trying"
	SipRegistrationStatusGetResponseSipRegistrationStatusRegistered         SipRegistrationStatusGetResponseSipRegistrationStatus = "registered"
	SipRegistrationStatusGetResponseSipRegistrationStatusUnknown            SipRegistrationStatusGetResponseSipRegistrationStatus = "unknown"
)

type SipRegistrationStatusGetParams struct {
	// Identifier of the UAC connection to look up.
	ConnectionID string `query:"connection_id" api:"required" json:"-"`
	// The kind of credential to look up. Only `uac_external_credential` is supported
	// today.
	//
	// Any of "uac_external_credential".
	CredentialType SipRegistrationStatusGetParamsCredentialType `query:"credential_type,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SipRegistrationStatusGetParams]'s query parameters as
// `url.Values`.
func (r SipRegistrationStatusGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The kind of credential to look up. Only `uac_external_credential` is supported
// today.
type SipRegistrationStatusGetParamsCredentialType string

const (
	SipRegistrationStatusGetParamsCredentialTypeUacExternalCredential SipRegistrationStatusGetParamsCredentialType = "uac_external_credential"
)
