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
	// Raw external-side registration block reported by the registrar.
	B2buaExternal map[string]any `json:"b2bua_external"`
	// Raw internal-side block reported by the registrar.
	B2buaInternal map[string]any `json:"b2bua_internal"`
	// Identifier of the UAC connection.
	ConnectionID string `json:"connection_id"`
	// Human-readable connection name.
	ConnectionName string `json:"connection_name"`
	// The credential type that was looked up.
	//
	// Any of "uac_external_credential".
	CredentialType SipRegistrationStatusGetResponseCredentialType `json:"credential_type"`
	// Registration state on the external (UAC / PBX) side, e.g. REGED.
	ExternalState string `json:"external_state"`
	// Outward-facing SIP settings used for registration. Password is redacted.
	ExternalUacSettings SipRegistrationStatusGetResponseExternalUacSettings `json:"external_uac_settings"`
	// Internal routing target the connection delivers calls to.
	InternalUacSettings SipRegistrationStatusGetResponseInternalUacSettings `json:"internal_uac_settings"`
	// SIP response from the last registration attempt.
	LastRegistrationResponse string `json:"last_registration_response"`
	// Internal pairing state, e.g. ACTIVE or INACTIVE.
	PairState string `json:"pair_state"`
	// True if the endpoint is currently registered.
	Registered bool `json:"registered"`
	// Owner of the connection.
	UserID string `json:"user_id"`
	// SIP username used for the registration.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		B2buaExternal            respjson.Field
		B2buaInternal            respjson.Field
		ConnectionID             respjson.Field
		ConnectionName           respjson.Field
		CredentialType           respjson.Field
		ExternalState            respjson.Field
		ExternalUacSettings      respjson.Field
		InternalUacSettings      respjson.Field
		LastRegistrationResponse respjson.Field
		PairState                respjson.Field
		Registered               respjson.Field
		UserID                   respjson.Field
		Username                 respjson.Field
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

// Outward-facing SIP settings used for registration. Password is redacted.
type SipRegistrationStatusGetResponseExternalUacSettings struct {
	AuthUsername  string `json:"auth_username"`
	ExpirationSec int64  `json:"expiration_sec"`
	FromUser      string `json:"from_user"`
	OutboundProxy string `json:"outbound_proxy"`
	// Always redacted.
	Password string `json:"password"`
	Proxy    string `json:"proxy"`
	// Any of "TCP", "UDP", "TLS".
	Transport string `json:"transport"`
	Username  string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthUsername  respjson.Field
		ExpirationSec respjson.Field
		FromUser      respjson.Field
		OutboundProxy respjson.Field
		Password      respjson.Field
		Proxy         respjson.Field
		Transport     respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipRegistrationStatusGetResponseExternalUacSettings) RawJSON() string { return r.JSON.raw }
func (r *SipRegistrationStatusGetResponseExternalUacSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal routing target the connection delivers calls to.
type SipRegistrationStatusGetResponseInternalUacSettings struct {
	DestinationUri string `json:"destination_uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationUri respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipRegistrationStatusGetResponseInternalUacSettings) RawJSON() string { return r.JSON.raw }
func (r *SipRegistrationStatusGetResponseInternalUacSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
