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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// UAC connection operations
//
// UacConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUacConnectionService] method instead.
type UacConnectionService struct {
	Options []option.RequestOption
	// UAC connection operations
	Actions UacConnectionActionService
}

// NewUacConnectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUacConnectionService(opts ...option.RequestOption) (r UacConnectionService) {
	r = UacConnectionService{}
	r.Options = opts
	r.Actions = NewUacConnectionActionService(opts...)
	return
}

// Creates a UAC connection. A UAC (User Agent Client) Connection registers Telnyx
// to your PBX — the opposite of a standard SIP trunk, where the PBX registers to
// Telnyx. Use UAC when your PBX doesn’t support outbound SIP registration or you
// need Telnyx to maintain the registration.
func (r *UacConnectionService) New(ctx context.Context, body UacConnectionNewParams, opts ...option.RequestOption) (res *UacConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "uac_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing UAC connection.
func (r *UacConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UacConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("uac_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates settings of an existing UAC connection.
func (r *UacConnectionService) Update(ctx context.Context, id string, body UacConnectionUpdateParams, opts ...option.RequestOption) (res *UacConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("uac_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of your UAC connections. A UAC (User Agent Client) Connection
// registers Telnyx to your PBX — the opposite of a standard SIP trunk, where the
// PBX registers to Telnyx. Use UAC when your PBX doesn’t support outbound SIP
// registration or you need Telnyx to maintain the registration.
func (r *UacConnectionService) List(ctx context.Context, query UacConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[UacConnection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "uac_connections"
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

// Returns a list of your UAC connections. A UAC (User Agent Client) Connection
// registers Telnyx to your PBX — the opposite of a standard SIP trunk, where the
// PBX registers to Telnyx. Use UAC when your PBX doesn’t support outbound SIP
// registration or you need Telnyx to maintain the registration.
func (r *UacConnectionService) ListAutoPaging(ctx context.Context, query UacConnectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[UacConnection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing UAC connection.
func (r *UacConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *UacConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("uac_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A UAC (User Agent Client) Connection registers Telnyx to your PBX — the opposite
// of a standard SIP trunk, where the PBX registers to Telnyx. Use UAC when your
// PBX doesn’t support outbound SIP registration or you need Telnyx to maintain the
// registration.
type UacConnection struct {
	// Identifies the type of resource.
	ID string `json:"id"`
	// Defaults to true
	Active bool `json:"active"`
	// `Latency` directs Telnyx to route media through the site with the lowest
	// round-trip time to the user's connection. Telnyx calculates this time using ICMP
	// ping messages. This can be disabled by specifying a site to handle all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "Sydney,
	// Australia", "Amsterdam, Netherlands", "London, UK", "Toronto, Canada",
	// "Vancouver, Canada", "Frankfurt, Germany".
	AnchorsiteOverride AnchorsiteOverride `json:"anchorsite_override"`
	// The uuid of the push credential for Android
	AndroidPushCredentialID string `json:"android_push_credential_id" api:"nullable"`
	// The authentication strategy used by this connection.
	//
	// Any of "uac-authentication".
	Authentication UacConnectionAuthentication `json:"authentication"`
	// Specifies if call cost webhooks should be sent for this connection.
	CallCostInWebhooks bool   `json:"call_cost_in_webhooks"`
	ConnectionName     string `json:"connection_name"`
	// ISO-8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// When enabled, Telnyx will generate comfort noise when you place the call on
	// hold. If disabled, you will need to generate comfort noise or on hold music to
	// avoid RTP timeout.
	DefaultOnHoldComfortNoiseEnabled bool `json:"default_on_hold_comfort_noise_enabled"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType DtmfType `json:"dtmf_type"`
	// Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG
	// scenarios.
	EncodeContactHeaderEnabled bool `json:"encode_contact_header_enabled"`
	// Enable use of SRTP for encryption. Cannot be set if the transport_portocol is
	// TLS.
	//
	// Any of "SRTP".
	EncryptedMedia EncryptedMedia `json:"encrypted_media" api:"nullable"`
	// External SIP peer settings used by Telnyx when registering to your PBX and
	// routing outbound calls.
	ExternalUacSettings UacExternalSettings `json:"external_uac_settings"`
	// The Telnyx-managed FQDN generated for the UAC connection.
	Fqdn string `json:"fqdn"`
	// The fixed outbound authentication mode used by UAC FQDN records.
	//
	// Any of "credential-authentication".
	FqdnOutboundAuthentication UacConnectionFqdnOutboundAuthentication `json:"fqdn_outbound_authentication"`
	// FQDN records managed automatically by the UAC connection lifecycle.
	Fqdns   []Fqdn     `json:"fqdns"`
	Inbound UacInbound `json:"inbound"`
	// Internal Telnyx-side settings for a UAC connection.
	InternalUacSettings UacInternalSettings `json:"internal_uac_settings"`
	// The uuid of the push credential for Ios
	IosPushCredentialID string `json:"ios_push_credential_id" api:"nullable"`
	// Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams
	// of SIP Trunking calls. The feature is off unless enabled. You may define min and
	// max values in msec for customized buffering behaviors. Larger values add latency
	// but tolerate more jitter, while smaller values reduce latency but are more
	// sensitive to jitter and reordering.
	JitterBuffer shared.ConnectionJitterBuffer `json:"jitter_buffer"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression UacConnectionNoiseSuppression `json:"noise_suppression"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetails `json:"noise_suppression_details"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly
	// if both are on the Telnyx network. If this is disabled, Telnyx will be able to
	// use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled bool        `json:"onnet_t38_passthrough_enabled"`
	Outbound                   UacOutbound `json:"outbound"`
	// The password to be used as part of the credentials. Must be 8 to 128 characters
	// long.
	Password string `json:"password"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The most recently known SIP registration status for this UAC connection.
	RegistrationStatus string `json:"registration_status" api:"nullable"`
	// ISO 8601 formatted date indicating when the registration status was last
	// updated.
	RegistrationStatusUpdatedAt string                 `json:"registration_status_updated_at" api:"nullable"`
	RtcpSettings                ConnectionRtcpSettings `json:"rtcp_settings"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection.
	// If enabled for all (unrestricted) then anyone who calls the SIP URI
	// <your-username>@telnyx.com will be connected to your Connection. You can also
	// choose to allow only calls that are originated on any Connections under your
	// account (internal).
	//
	// Any of "disabled", "unrestricted", "internal".
	SipUriCallingPreference UacConnectionSipUriCallingPreference `json:"sip_uri_calling_preference"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// ISO-8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// The user name to be used as part of the credentials. Must be 4-32 characters
	// long and alphanumeric values only (no spaces or special characters). At least
	// one of the first 5 characters must be a letter.
	UserName string `json:"user_name"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion UacConnectionWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url" api:"nullable" format:"uri"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		Active                           respjson.Field
		AnchorsiteOverride               respjson.Field
		AndroidPushCredentialID          respjson.Field
		Authentication                   respjson.Field
		CallCostInWebhooks               respjson.Field
		ConnectionName                   respjson.Field
		CreatedAt                        respjson.Field
		DefaultOnHoldComfortNoiseEnabled respjson.Field
		DtmfType                         respjson.Field
		EncodeContactHeaderEnabled       respjson.Field
		EncryptedMedia                   respjson.Field
		ExternalUacSettings              respjson.Field
		Fqdn                             respjson.Field
		FqdnOutboundAuthentication       respjson.Field
		Fqdns                            respjson.Field
		Inbound                          respjson.Field
		InternalUacSettings              respjson.Field
		IosPushCredentialID              respjson.Field
		JitterBuffer                     respjson.Field
		NoiseSuppression                 respjson.Field
		NoiseSuppressionDetails          respjson.Field
		OnnetT38PassthroughEnabled       respjson.Field
		Outbound                         respjson.Field
		Password                         respjson.Field
		RecordType                       respjson.Field
		RegistrationStatus               respjson.Field
		RegistrationStatusUpdatedAt      respjson.Field
		RtcpSettings                     respjson.Field
		SipUriCallingPreference          respjson.Field
		Tags                             respjson.Field
		UpdatedAt                        respjson.Field
		UserName                         respjson.Field
		WebhookAPIVersion                respjson.Field
		WebhookEventFailoverURL          respjson.Field
		WebhookEventURL                  respjson.Field
		WebhookTimeoutSecs               respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacConnection) RawJSON() string { return r.JSON.raw }
func (r *UacConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The authentication strategy used by this connection.
type UacConnectionAuthentication string

const (
	UacConnectionAuthenticationUacAuthentication UacConnectionAuthentication = "uac-authentication"
)

// The fixed outbound authentication mode used by UAC FQDN records.
type UacConnectionFqdnOutboundAuthentication string

const (
	UacConnectionFqdnOutboundAuthenticationCredentialAuthentication UacConnectionFqdnOutboundAuthentication = "credential-authentication"
)

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type UacConnectionNoiseSuppression string

const (
	UacConnectionNoiseSuppressionInbound  UacConnectionNoiseSuppression = "inbound"
	UacConnectionNoiseSuppressionOutbound UacConnectionNoiseSuppression = "outbound"
	UacConnectionNoiseSuppressionBoth     UacConnectionNoiseSuppression = "both"
	UacConnectionNoiseSuppressionDisabled UacConnectionNoiseSuppression = "disabled"
)

// This feature enables inbound SIP URI calls to your Credential Auth Connection.
// If enabled for all (unrestricted) then anyone who calls the SIP URI
// <your-username>@telnyx.com will be connected to your Connection. You can also
// choose to allow only calls that are originated on any Connections under your
// account (internal).
type UacConnectionSipUriCallingPreference string

const (
	UacConnectionSipUriCallingPreferenceDisabled     UacConnectionSipUriCallingPreference = "disabled"
	UacConnectionSipUriCallingPreferenceUnrestricted UacConnectionSipUriCallingPreference = "unrestricted"
	UacConnectionSipUriCallingPreferenceInternal     UacConnectionSipUriCallingPreference = "internal"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type UacConnectionWebhookAPIVersion string

const (
	UacConnectionWebhookAPIVersionV1 UacConnectionWebhookAPIVersion = "1"
	UacConnectionWebhookAPIVersionV2 UacConnectionWebhookAPIVersion = "2"
)

// External SIP peer settings used by Telnyx when registering to your PBX and
// routing outbound calls.
type UacExternalSettings struct {
	// The authentication username used in SIP digest authentication. If not set, the
	// Username value will be used.
	AuthUsername string `json:"auth_username" api:"nullable"`
	// The registration interval, in seconds, indicating how often the system refreshes
	// the SIP registration with the external SIP peer.
	ExpirationSec int64 `json:"expiration_sec" api:"nullable"`
	// The user portion of the SIP From header used in outbound requests. This controls
	// the caller identity presented to the external SIP peer.
	FromUser string `json:"from_user" api:"nullable"`
	// An optional SIP proxy used to route outbound requests before reaching the
	// external SIP peer.
	OutboundProxy string `json:"outbound_proxy" api:"nullable"`
	// The SIP password used for digest authentication with the external SIP peer.
	Password string `json:"password"`
	// The SIP proxy address of the external SIP peer used for registrations and
	// outbound call routing.
	Proxy string `json:"proxy"`
	// The transport protocol used for SIP signaling when communicating with the
	// external SIP peer. One of UDP, TLS, or TCP.
	//
	// Any of "UDP", "TLS", "TCP".
	Transport UacExternalSettingsTransport `json:"transport" api:"nullable"`
	// The SIP username used to authenticate with the external SIP peer for
	// registrations and outbound calls. Must start with a letter or number and contain
	// only letters, numbers, hyphens, and underscores.
	Username string `json:"username"`
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
func (r UacExternalSettings) RawJSON() string { return r.JSON.raw }
func (r *UacExternalSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UacExternalSettings to a UacExternalSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UacExternalSettingsParam.Overrides()
func (r UacExternalSettings) ToParam() UacExternalSettingsParam {
	return param.Override[UacExternalSettingsParam](json.RawMessage(r.RawJSON()))
}

// The transport protocol used for SIP signaling when communicating with the
// external SIP peer. One of UDP, TLS, or TCP.
type UacExternalSettingsTransport string

const (
	UacExternalSettingsTransportUdp UacExternalSettingsTransport = "UDP"
	UacExternalSettingsTransportTls UacExternalSettingsTransport = "TLS"
	UacExternalSettingsTransportTcp UacExternalSettingsTransport = "TCP"
)

// External SIP peer settings used by Telnyx when registering to your PBX and
// routing outbound calls.
type UacExternalSettingsParam struct {
	// The authentication username used in SIP digest authentication. If not set, the
	// Username value will be used.
	AuthUsername param.Opt[string] `json:"auth_username,omitzero"`
	// The registration interval, in seconds, indicating how often the system refreshes
	// the SIP registration with the external SIP peer.
	ExpirationSec param.Opt[int64] `json:"expiration_sec,omitzero"`
	// The user portion of the SIP From header used in outbound requests. This controls
	// the caller identity presented to the external SIP peer.
	FromUser param.Opt[string] `json:"from_user,omitzero"`
	// An optional SIP proxy used to route outbound requests before reaching the
	// external SIP peer.
	OutboundProxy param.Opt[string] `json:"outbound_proxy,omitzero"`
	// The SIP password used for digest authentication with the external SIP peer.
	Password param.Opt[string] `json:"password,omitzero"`
	// The SIP proxy address of the external SIP peer used for registrations and
	// outbound call routing.
	Proxy param.Opt[string] `json:"proxy,omitzero"`
	// The SIP username used to authenticate with the external SIP peer for
	// registrations and outbound calls. Must start with a letter or number and contain
	// only letters, numbers, hyphens, and underscores.
	Username param.Opt[string] `json:"username,omitzero"`
	// The transport protocol used for SIP signaling when communicating with the
	// external SIP peer. One of UDP, TLS, or TCP.
	//
	// Any of "UDP", "TLS", "TCP".
	Transport UacExternalSettingsTransport `json:"transport,omitzero"`
	paramObj
}

func (r UacExternalSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow UacExternalSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacExternalSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacInbound struct {
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat UacInboundAniNumberFormat `json:"ani_number_format"`
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs"`
	// Default routing method to be used when a number is associated with the
	// connection. Must be one of the routing method types or left blank, other values
	// are not allowed.
	//
	// Any of "sequential", "round-robin".
	DefaultRoutingMethod UacInboundDefaultRoutingMethod `json:"default_routing_method"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat UacInboundDnisNumberFormat `json:"dnis_number_format"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone bool `json:"generate_ringback_tone"`
	// When set, inbound phone calls will receive ISUP parameters via SIP headers.
	// (Only when available and only when using TCP or TLS transport.)
	IsupHeadersEnabled bool `json:"isup_headers_enabled"`
	// Enable PRACK messages as defined in RFC3262.
	PrackEnabled bool `json:"prack_enabled"`
	// When enabled the SIP Connection will receive the Identity header with
	// Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP
	// transport.
	ShakenStirEnabled bool `json:"shaken_stir_enabled"`
	// When enabled, allows multiple devices to ring simultaneously on incoming calls.
	//
	// Any of "disabled", "enabled".
	SimultaneousRinging UacInboundSimultaneousRinging `json:"simultaneous_ringing"`
	// Defaults to true.
	SipCompactHeadersEnabled bool `json:"sip_compact_headers_enabled"`
	// The Telnyx-generated SIP subdomain for this UAC connection.
	SipSubdomain string `json:"sip_subdomain"`
	// Controls which SIP URI callers may reach this connection.
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings UacInboundSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs int64 `json:"timeout_1xx_secs"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs int64 `json:"timeout_2xx_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniNumberFormat             respjson.Field
		ChannelLimit                respjson.Field
		Codecs                      respjson.Field
		DefaultRoutingMethod        respjson.Field
		DnisNumberFormat            respjson.Field
		GenerateRingbackTone        respjson.Field
		IsupHeadersEnabled          respjson.Field
		PrackEnabled                respjson.Field
		ShakenStirEnabled           respjson.Field
		SimultaneousRinging         respjson.Field
		SipCompactHeadersEnabled    respjson.Field
		SipSubdomain                respjson.Field
		SipSubdomainReceiveSettings respjson.Field
		Timeout1xxSecs              respjson.Field
		Timeout2xxSecs              respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacInbound) RawJSON() string { return r.JSON.raw }
func (r *UacInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This setting allows you to set the format with which the caller's number (ANI)
// is sent for inbound phone calls.
type UacInboundAniNumberFormat string

const (
	UacInboundAniNumberFormatPlusE164         UacInboundAniNumberFormat = "+E.164"
	UacInboundAniNumberFormatE164             UacInboundAniNumberFormat = "E.164"
	UacInboundAniNumberFormatPlusE164National UacInboundAniNumberFormat = "+E.164-national"
	UacInboundAniNumberFormatE164National     UacInboundAniNumberFormat = "E.164-national"
)

// Default routing method to be used when a number is associated with the
// connection. Must be one of the routing method types or left blank, other values
// are not allowed.
type UacInboundDefaultRoutingMethod string

const (
	UacInboundDefaultRoutingMethodSequential UacInboundDefaultRoutingMethod = "sequential"
	UacInboundDefaultRoutingMethodRoundRobin UacInboundDefaultRoutingMethod = "round-robin"
)

type UacInboundDnisNumberFormat string

const (
	UacInboundDnisNumberFormatPlusE164    UacInboundDnisNumberFormat = "+e164"
	UacInboundDnisNumberFormatE164        UacInboundDnisNumberFormat = "e164"
	UacInboundDnisNumberFormatNational    UacInboundDnisNumberFormat = "national"
	UacInboundDnisNumberFormatSipUsername UacInboundDnisNumberFormat = "sip_username"
)

// When enabled, allows multiple devices to ring simultaneously on incoming calls.
type UacInboundSimultaneousRinging string

const (
	UacInboundSimultaneousRingingDisabled UacInboundSimultaneousRinging = "disabled"
	UacInboundSimultaneousRingingEnabled  UacInboundSimultaneousRinging = "enabled"
)

// Controls which SIP URI callers may reach this connection.
type UacInboundSipSubdomainReceiveSettings string

const (
	UacInboundSipSubdomainReceiveSettingsOnlyMyConnections UacInboundSipSubdomainReceiveSettings = "only_my_connections"
	UacInboundSipSubdomainReceiveSettingsFromAnyone        UacInboundSipSubdomainReceiveSettings = "from_anyone"
)

// Internal Telnyx-side settings for a UAC connection.
type UacInternalSettings struct {
	// The SIP URI that Telnyx will call when handling an inbound request from the
	// external peer. Do not include a `sip:` prefix. The value must be in the format
	// `userinfo@<subdomain.>sip.telnyx.com` or
	// `userinfo@<subdomain.>sipdev.telnyx.com`; the userinfo portion may contain only
	// letters, digits, hyphens, and underscores.
	DestinationUri string `json:"destination_uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationUri respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacInternalSettings) RawJSON() string { return r.JSON.raw }
func (r *UacInternalSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UacInternalSettings to a UacInternalSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UacInternalSettingsParam.Overrides()
func (r UacInternalSettings) ToParam() UacInternalSettingsParam {
	return param.Override[UacInternalSettingsParam](json.RawMessage(r.RawJSON()))
}

// Internal Telnyx-side settings for a UAC connection.
type UacInternalSettingsParam struct {
	// The SIP URI that Telnyx will call when handling an inbound request from the
	// external peer. Do not include a `sip:` prefix. The value must be in the format
	// `userinfo@<subdomain.>sip.telnyx.com` or
	// `userinfo@<subdomain.>sipdev.telnyx.com`; the userinfo portion may contain only
	// letters, digits, hyphens, and underscores.
	DestinationUri param.Opt[string] `json:"destination_uri,omitzero"`
	paramObj
}

func (r UacInternalSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow UacInternalSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacInternalSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacOutbound struct {
	// Set a phone number as the ani_override value to override caller id number on
	// outbound calls.
	AniOverride string `json:"ani_override"`
	// Specifies when we apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType UacOutboundAniOverrideType `json:"ani_override_type"`
	// Forces all SIP calls originated on this connection to be "parked" instead of
	// "bridged" to the destination specified on the URI. Parked calls will return
	// ringback to the caller and will await for a Call Control command to define which
	// action will be taken next.
	CallParkingEnabled bool `json:"call_parking_enabled" api:"nullable"`
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone bool `json:"generate_ringback_tone"`
	// When set, ringback will not wait for indication before sending ringback tone to
	// calling party.
	InstantRingbackEnabled bool `json:"instant_ringback_enabled"`
	// A 2-character country code specifying the country whose national dialing rules
	// should be used. For example, if set to `US` then any US number can be dialed
	// without preprending +1 to the number. When left blank, Telnyx will try US and GB
	// dialing rules, in that order, by default.
	Localization string `json:"localization"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
	// default, Telnyx will send the re-invite. If set to `customer`, the caller is
	// expected to send the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource UacOutboundT38ReinviteSource `json:"t38_reinvite_source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniOverride            respjson.Field
		AniOverrideType        respjson.Field
		CallParkingEnabled     respjson.Field
		ChannelLimit           respjson.Field
		GenerateRingbackTone   respjson.Field
		InstantRingbackEnabled respjson.Field
		Localization           respjson.Field
		OutboundVoiceProfileID respjson.Field
		T38ReinviteSource      respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacOutbound) RawJSON() string { return r.JSON.raw }
func (r *UacOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UacOutbound to a UacOutboundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UacOutboundParam.Overrides()
func (r UacOutbound) ToParam() UacOutboundParam {
	return param.Override[UacOutboundParam](json.RawMessage(r.RawJSON()))
}

// Specifies when we apply your ani_override setting. Only applies when
// ani_override is not blank.
type UacOutboundAniOverrideType string

const (
	UacOutboundAniOverrideTypeAlways    UacOutboundAniOverrideType = "always"
	UacOutboundAniOverrideTypeNormal    UacOutboundAniOverrideType = "normal"
	UacOutboundAniOverrideTypeEmergency UacOutboundAniOverrideType = "emergency"
)

// This setting only affects connections with Fax-type Outbound Voice Profiles. The
// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
// default, Telnyx will send the re-invite. If set to `customer`, the caller is
// expected to send the t.38 reinvite.
type UacOutboundT38ReinviteSource string

const (
	UacOutboundT38ReinviteSourceTelnyx         UacOutboundT38ReinviteSource = "telnyx"
	UacOutboundT38ReinviteSourceCustomer       UacOutboundT38ReinviteSource = "customer"
	UacOutboundT38ReinviteSourceDisabled       UacOutboundT38ReinviteSource = "disabled"
	UacOutboundT38ReinviteSourcePassthru       UacOutboundT38ReinviteSource = "passthru"
	UacOutboundT38ReinviteSourceCallerPassthru UacOutboundT38ReinviteSource = "caller-passthru"
	UacOutboundT38ReinviteSourceCalleePassthru UacOutboundT38ReinviteSource = "callee-passthru"
)

type UacOutboundParam struct {
	// Forces all SIP calls originated on this connection to be "parked" instead of
	// "bridged" to the destination specified on the URI. Parked calls will return
	// ringback to the caller and will await for a Call Control command to define which
	// action will be taken next.
	CallParkingEnabled param.Opt[bool] `json:"call_parking_enabled,omitzero"`
	// Set a phone number as the ani_override value to override caller id number on
	// outbound calls.
	AniOverride param.Opt[string] `json:"ani_override,omitzero"`
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone param.Opt[bool] `json:"generate_ringback_tone,omitzero"`
	// When set, ringback will not wait for indication before sending ringback tone to
	// calling party.
	InstantRingbackEnabled param.Opt[bool] `json:"instant_ringback_enabled,omitzero"`
	// A 2-character country code specifying the country whose national dialing rules
	// should be used. For example, if set to `US` then any US number can be dialed
	// without preprending +1 to the number. When left blank, Telnyx will try US and GB
	// dialing rules, in that order, by default.
	Localization param.Opt[string] `json:"localization,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	// Specifies when we apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType UacOutboundAniOverrideType `json:"ani_override_type,omitzero"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
	// default, Telnyx will send the re-invite. If set to `customer`, the caller is
	// expected to send the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource UacOutboundT38ReinviteSource `json:"t38_reinvite_source,omitzero"`
	paramObj
}

func (r UacOutboundParam) MarshalJSON() (data []byte, err error) {
	type shadow UacOutboundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacOutboundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacConnectionNewResponse struct {
	// A UAC (User Agent Client) Connection registers Telnyx to your PBX — the opposite
	// of a standard SIP trunk, where the PBX registers to Telnyx. Use UAC when your
	// PBX doesn’t support outbound SIP registration or you need Telnyx to maintain the
	// registration.
	Data UacConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UacConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacConnectionGetResponse struct {
	// A UAC (User Agent Client) Connection registers Telnyx to your PBX — the opposite
	// of a standard SIP trunk, where the PBX registers to Telnyx. Use UAC when your
	// PBX doesn’t support outbound SIP registration or you need Telnyx to maintain the
	// registration.
	Data UacConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UacConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacConnectionUpdateResponse struct {
	// A UAC (User Agent Client) Connection registers Telnyx to your PBX — the opposite
	// of a standard SIP trunk, where the PBX registers to Telnyx. Use UAC when your
	// PBX doesn’t support outbound SIP registration or you need Telnyx to maintain the
	// registration.
	Data UacConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *UacConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacConnectionDeleteResponse struct {
	// A UAC (User Agent Client) Connection registers Telnyx to your PBX — the opposite
	// of a standard SIP trunk, where the PBX registers to Telnyx. Use UAC when your
	// PBX doesn’t support outbound SIP registration or you need Telnyx to maintain the
	// registration.
	Data UacConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UacConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *UacConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UacConnectionNewParams struct {
	// A user-assigned name to help manage the connection.
	ConnectionName string `json:"connection_name" api:"required"`
	// The uuid of the push credential for Android
	AndroidPushCredentialID param.Opt[string] `json:"android_push_credential_id,omitzero"`
	// The uuid of the push credential for Ios
	IosPushCredentialID param.Opt[string] `json:"ios_push_credential_id,omitzero"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Defaults to true
	Active param.Opt[bool] `json:"active,omitzero"`
	// Specifies if call cost webhooks should be sent for this connection.
	CallCostInWebhooks param.Opt[bool] `json:"call_cost_in_webhooks,omitzero"`
	// When enabled, Telnyx will generate comfort noise when you place the call on
	// hold. If disabled, you will need to generate comfort noise or on hold music to
	// avoid RTP timeout.
	DefaultOnHoldComfortNoiseEnabled param.Opt[bool] `json:"default_on_hold_comfort_noise_enabled,omitzero"`
	// Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG
	// scenarios.
	EncodeContactHeaderEnabled param.Opt[bool] `json:"encode_contact_header_enabled,omitzero"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly
	// if both are on the Telnyx network. If this is disabled, Telnyx will be able to
	// use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled param.Opt[bool] `json:"onnet_t38_passthrough_enabled,omitzero"`
	// The password to be used as part of the credentials. Must be 8 to 128 characters
	// long.
	Password param.Opt[string] `json:"password,omitzero"`
	// The user name to be used as part of the credentials. Must be 4-32 characters
	// long and alphanumeric values only (no spaces or special characters).
	UserName param.Opt[string] `json:"user_name,omitzero"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL param.Opt[string] `json:"webhook_event_url,omitzero" format:"uri"`
	// Enable use of SRTP for encryption. Cannot be set if the transport_portocol is
	// TLS.
	//
	// Any of "SRTP".
	EncryptedMedia EncryptedMedia `json:"encrypted_media,omitzero"`
	// `Latency` directs Telnyx to route media through the site with the lowest
	// round-trip time to the user's connection. Telnyx calculates this time using ICMP
	// ping messages. This can be disabled by specifying a site to handle all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "Sydney,
	// Australia", "Amsterdam, Netherlands", "London, UK", "Toronto, Canada",
	// "Vancouver, Canada", "Frankfurt, Germany".
	AnchorsiteOverride AnchorsiteOverride `json:"anchorsite_override,omitzero"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType DtmfType `json:"dtmf_type,omitzero"`
	// External SIP peer settings used by Telnyx when registering to your PBX and
	// routing outbound calls.
	ExternalUacSettings UacExternalSettingsParam `json:"external_uac_settings,omitzero"`
	// Inbound settings that can be supplied when creating or updating a UAC
	// connection. The SIP subdomain fields returned in UAC connection responses are
	// generated by Telnyx and are not accepted as request parameters.
	Inbound UacConnectionNewParamsInbound `json:"inbound,omitzero"`
	// Internal Telnyx-side settings for a UAC connection.
	InternalUacSettings UacInternalSettingsParam `json:"internal_uac_settings,omitzero"`
	// Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams
	// of SIP Trunking calls. The feature is off unless enabled. You may define min and
	// max values in msec for customized buffering behaviors. Larger values add latency
	// but tolerate more jitter, while smaller values reduce latency but are more
	// sensitive to jitter and reordering.
	JitterBuffer shared.ConnectionJitterBufferParam `json:"jitter_buffer,omitzero"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression UacConnectionNewParamsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetailsParam `json:"noise_suppression_details,omitzero"`
	Outbound                UacOutboundParam                              `json:"outbound,omitzero"`
	RtcpSettings            ConnectionRtcpSettingsParam                   `json:"rtcp_settings,omitzero"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection.
	// If enabled for all (unrestricted) then anyone who calls the SIP URI
	// <your-username>@telnyx.com will be connected to your Connection. You can also
	// choose to allow only calls that are originated on any Connections under your
	// account (internal).
	//
	// Any of "disabled", "unrestricted", "internal".
	SipUriCallingPreference UacConnectionNewParamsSipUriCallingPreference `json:"sip_uri_calling_preference,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1, v2 or texml. Note -
	// texml can only be set when the outbound object parameter call_parking_enabled is
	// included and set to true.
	//
	// Any of "1", "2", "texml".
	WebhookAPIVersion UacConnectionNewParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r UacConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UacConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inbound settings that can be supplied when creating or updating a UAC
// connection. The SIP subdomain fields returned in UAC connection responses are
// generated by Telnyx and are not accepted as request parameters.
type UacConnectionNewParamsInbound struct {
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone param.Opt[bool] `json:"generate_ringback_tone,omitzero"`
	// When set, inbound phone calls will receive ISUP parameters via SIP headers.
	// (Only when available and only when using TCP or TLS transport.)
	IsupHeadersEnabled param.Opt[bool] `json:"isup_headers_enabled,omitzero"`
	// Enable PRACK messages as defined in RFC3262.
	PrackEnabled param.Opt[bool] `json:"prack_enabled,omitzero"`
	// When enabled the SIP Connection will receive the Identity header with
	// Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP
	// transport.
	ShakenStirEnabled param.Opt[bool] `json:"shaken_stir_enabled,omitzero"`
	// Defaults to true.
	SipCompactHeadersEnabled param.Opt[bool] `json:"sip_compact_headers_enabled,omitzero"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs param.Opt[int64] `json:"timeout_1xx_secs,omitzero"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs param.Opt[int64] `json:"timeout_2xx_secs,omitzero"`
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat string `json:"ani_number_format,omitzero"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs,omitzero"`
	// Default routing method to be used when a number is associated with the
	// connection. Must be one of the routing method types or left blank, other values
	// are not allowed.
	//
	// Any of "sequential", "round-robin".
	DefaultRoutingMethod string `json:"default_routing_method,omitzero"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat string `json:"dnis_number_format,omitzero"`
	// When enabled, allows multiple devices to ring simultaneously on incoming calls.
	//
	// Any of "disabled", "enabled".
	SimultaneousRinging string `json:"simultaneous_ringing,omitzero"`
	paramObj
}

func (r UacConnectionNewParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow UacConnectionNewParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacConnectionNewParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UacConnectionNewParamsInbound](
		"ani_number_format", "+E.164", "E.164", "+E.164-national", "E.164-national",
	)
	apijson.RegisterFieldValidator[UacConnectionNewParamsInbound](
		"default_routing_method", "sequential", "round-robin",
	)
	apijson.RegisterFieldValidator[UacConnectionNewParamsInbound](
		"dnis_number_format", "+e164", "e164", "national", "sip_username",
	)
	apijson.RegisterFieldValidator[UacConnectionNewParamsInbound](
		"simultaneous_ringing", "disabled", "enabled",
	)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type UacConnectionNewParamsNoiseSuppression string

const (
	UacConnectionNewParamsNoiseSuppressionInbound  UacConnectionNewParamsNoiseSuppression = "inbound"
	UacConnectionNewParamsNoiseSuppressionOutbound UacConnectionNewParamsNoiseSuppression = "outbound"
	UacConnectionNewParamsNoiseSuppressionBoth     UacConnectionNewParamsNoiseSuppression = "both"
	UacConnectionNewParamsNoiseSuppressionDisabled UacConnectionNewParamsNoiseSuppression = "disabled"
)

// This feature enables inbound SIP URI calls to your Credential Auth Connection.
// If enabled for all (unrestricted) then anyone who calls the SIP URI
// <your-username>@telnyx.com will be connected to your Connection. You can also
// choose to allow only calls that are originated on any Connections under your
// account (internal).
type UacConnectionNewParamsSipUriCallingPreference string

const (
	UacConnectionNewParamsSipUriCallingPreferenceDisabled     UacConnectionNewParamsSipUriCallingPreference = "disabled"
	UacConnectionNewParamsSipUriCallingPreferenceUnrestricted UacConnectionNewParamsSipUriCallingPreference = "unrestricted"
	UacConnectionNewParamsSipUriCallingPreferenceInternal     UacConnectionNewParamsSipUriCallingPreference = "internal"
)

// Determines which webhook format will be used, Telnyx API v1, v2 or texml. Note -
// texml can only be set when the outbound object parameter call_parking_enabled is
// included and set to true.
type UacConnectionNewParamsWebhookAPIVersion string

const (
	UacConnectionNewParamsWebhookAPIVersionV1    UacConnectionNewParamsWebhookAPIVersion = "1"
	UacConnectionNewParamsWebhookAPIVersionV2    UacConnectionNewParamsWebhookAPIVersion = "2"
	UacConnectionNewParamsWebhookAPIVersionTexml UacConnectionNewParamsWebhookAPIVersion = "texml"
)

type UacConnectionUpdateParams struct {
	// The uuid of the push credential for Android
	AndroidPushCredentialID param.Opt[string] `json:"android_push_credential_id,omitzero"`
	// The uuid of the push credential for Ios
	IosPushCredentialID param.Opt[string] `json:"ios_push_credential_id,omitzero"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Defaults to true
	Active param.Opt[bool] `json:"active,omitzero"`
	// Specifies if call cost webhooks should be sent for this connection.
	CallCostInWebhooks param.Opt[bool] `json:"call_cost_in_webhooks,omitzero"`
	// A user-assigned name to help manage the connection.
	ConnectionName param.Opt[string] `json:"connection_name,omitzero"`
	// When enabled, Telnyx will generate comfort noise when you place the call on
	// hold. If disabled, you will need to generate comfort noise or on hold music to
	// avoid RTP timeout.
	DefaultOnHoldComfortNoiseEnabled param.Opt[bool] `json:"default_on_hold_comfort_noise_enabled,omitzero"`
	// Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG
	// scenarios.
	EncodeContactHeaderEnabled param.Opt[bool] `json:"encode_contact_header_enabled,omitzero"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly
	// if both are on the Telnyx network. If this is disabled, Telnyx will be able to
	// use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled param.Opt[bool] `json:"onnet_t38_passthrough_enabled,omitzero"`
	// The password to be used as part of the credentials. Must be 8 to 128 characters
	// long.
	Password param.Opt[string] `json:"password,omitzero"`
	// The user name to be used as part of the credentials. Must be 4-32 characters
	// long and alphanumeric values only (no spaces or special characters).
	UserName param.Opt[string] `json:"user_name,omitzero"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL param.Opt[string] `json:"webhook_event_url,omitzero" format:"uri"`
	// Enable use of SRTP for encryption. Cannot be set if the transport_portocol is
	// TLS.
	//
	// Any of "SRTP".
	EncryptedMedia EncryptedMedia `json:"encrypted_media,omitzero"`
	// `Latency` directs Telnyx to route media through the site with the lowest
	// round-trip time to the user's connection. Telnyx calculates this time using ICMP
	// ping messages. This can be disabled by specifying a site to handle all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "Sydney,
	// Australia", "Amsterdam, Netherlands", "London, UK", "Toronto, Canada",
	// "Vancouver, Canada", "Frankfurt, Germany".
	AnchorsiteOverride AnchorsiteOverride `json:"anchorsite_override,omitzero"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType DtmfType `json:"dtmf_type,omitzero"`
	// External SIP peer settings used by Telnyx when registering to your PBX and
	// routing outbound calls.
	ExternalUacSettings UacExternalSettingsParam `json:"external_uac_settings,omitzero"`
	// Inbound settings that can be supplied when creating or updating a UAC
	// connection. The SIP subdomain fields returned in UAC connection responses are
	// generated by Telnyx and are not accepted as request parameters.
	Inbound UacConnectionUpdateParamsInbound `json:"inbound,omitzero"`
	// Internal Telnyx-side settings for a UAC connection.
	InternalUacSettings UacInternalSettingsParam `json:"internal_uac_settings,omitzero"`
	// Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams
	// of SIP Trunking calls. The feature is off unless enabled. You may define min and
	// max values in msec for customized buffering behaviors. Larger values add latency
	// but tolerate more jitter, while smaller values reduce latency but are more
	// sensitive to jitter and reordering.
	JitterBuffer shared.ConnectionJitterBufferParam `json:"jitter_buffer,omitzero"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression UacConnectionUpdateParamsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetailsParam `json:"noise_suppression_details,omitzero"`
	Outbound                UacOutboundParam                              `json:"outbound,omitzero"`
	RtcpSettings            ConnectionRtcpSettingsParam                   `json:"rtcp_settings,omitzero"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection.
	// If enabled for all (unrestricted) then anyone who calls the SIP URI
	// <your-username>@telnyx.com will be connected to your Connection. You can also
	// choose to allow only calls that are originated on any Connections under your
	// account (internal).
	//
	// Any of "disabled", "unrestricted", "internal".
	SipUriCallingPreference UacConnectionUpdateParamsSipUriCallingPreference `json:"sip_uri_calling_preference,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion UacConnectionUpdateParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r UacConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UacConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inbound settings that can be supplied when creating or updating a UAC
// connection. The SIP subdomain fields returned in UAC connection responses are
// generated by Telnyx and are not accepted as request parameters.
type UacConnectionUpdateParamsInbound struct {
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone param.Opt[bool] `json:"generate_ringback_tone,omitzero"`
	// When set, inbound phone calls will receive ISUP parameters via SIP headers.
	// (Only when available and only when using TCP or TLS transport.)
	IsupHeadersEnabled param.Opt[bool] `json:"isup_headers_enabled,omitzero"`
	// Enable PRACK messages as defined in RFC3262.
	PrackEnabled param.Opt[bool] `json:"prack_enabled,omitzero"`
	// When enabled the SIP Connection will receive the Identity header with
	// Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP
	// transport.
	ShakenStirEnabled param.Opt[bool] `json:"shaken_stir_enabled,omitzero"`
	// Defaults to true.
	SipCompactHeadersEnabled param.Opt[bool] `json:"sip_compact_headers_enabled,omitzero"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs param.Opt[int64] `json:"timeout_1xx_secs,omitzero"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs param.Opt[int64] `json:"timeout_2xx_secs,omitzero"`
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat string `json:"ani_number_format,omitzero"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs,omitzero"`
	// Default routing method to be used when a number is associated with the
	// connection. Must be one of the routing method types or left blank, other values
	// are not allowed.
	//
	// Any of "sequential", "round-robin".
	DefaultRoutingMethod string `json:"default_routing_method,omitzero"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat string `json:"dnis_number_format,omitzero"`
	// When enabled, allows multiple devices to ring simultaneously on incoming calls.
	//
	// Any of "disabled", "enabled".
	SimultaneousRinging string `json:"simultaneous_ringing,omitzero"`
	paramObj
}

func (r UacConnectionUpdateParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow UacConnectionUpdateParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UacConnectionUpdateParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UacConnectionUpdateParamsInbound](
		"ani_number_format", "+E.164", "E.164", "+E.164-national", "E.164-national",
	)
	apijson.RegisterFieldValidator[UacConnectionUpdateParamsInbound](
		"default_routing_method", "sequential", "round-robin",
	)
	apijson.RegisterFieldValidator[UacConnectionUpdateParamsInbound](
		"dnis_number_format", "+e164", "e164", "national", "sip_username",
	)
	apijson.RegisterFieldValidator[UacConnectionUpdateParamsInbound](
		"simultaneous_ringing", "disabled", "enabled",
	)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type UacConnectionUpdateParamsNoiseSuppression string

const (
	UacConnectionUpdateParamsNoiseSuppressionInbound  UacConnectionUpdateParamsNoiseSuppression = "inbound"
	UacConnectionUpdateParamsNoiseSuppressionOutbound UacConnectionUpdateParamsNoiseSuppression = "outbound"
	UacConnectionUpdateParamsNoiseSuppressionBoth     UacConnectionUpdateParamsNoiseSuppression = "both"
	UacConnectionUpdateParamsNoiseSuppressionDisabled UacConnectionUpdateParamsNoiseSuppression = "disabled"
)

// This feature enables inbound SIP URI calls to your Credential Auth Connection.
// If enabled for all (unrestricted) then anyone who calls the SIP URI
// <your-username>@telnyx.com will be connected to your Connection. You can also
// choose to allow only calls that are originated on any Connections under your
// account (internal).
type UacConnectionUpdateParamsSipUriCallingPreference string

const (
	UacConnectionUpdateParamsSipUriCallingPreferenceDisabled     UacConnectionUpdateParamsSipUriCallingPreference = "disabled"
	UacConnectionUpdateParamsSipUriCallingPreferenceUnrestricted UacConnectionUpdateParamsSipUriCallingPreference = "unrestricted"
	UacConnectionUpdateParamsSipUriCallingPreferenceInternal     UacConnectionUpdateParamsSipUriCallingPreference = "internal"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type UacConnectionUpdateParamsWebhookAPIVersion string

const (
	UacConnectionUpdateParamsWebhookAPIVersionV1 UacConnectionUpdateParamsWebhookAPIVersion = "1"
	UacConnectionUpdateParamsWebhookAPIVersionV2 UacConnectionUpdateParamsWebhookAPIVersion = "2"
)

type UacConnectionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
	// filter[outbound.outbound_voice_profile_id]
	Filter UacConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code> -</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>connection_name</code>: sorts the result by the
	//	  <code>connection_name</code> field in ascending order.
	//	</li>
	//
	//	<li>
	//	  <code>-connection_name</code>: sorts the result by the
	//	  <code>connection_name</code> field in descending order.
	//	</li>
	//
	// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
	//
	// Any of "created_at", "connection_name", "active".
	Sort UacConnectionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UacConnectionListParams]'s query parameters as
// `url.Values`.
func (r UacConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
// filter[outbound.outbound_voice_profile_id]
type UacConnectionListParamsFilter struct {
	// If present, connections with an `fqdn` that equals the given value will be
	// returned. Matching is case-sensitive, and the full string must match.
	Fqdn param.Opt[string] `query:"fqdn,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" json:"-"`
	// Filter by connection_name using nested operations
	ConnectionName UacConnectionListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UacConnectionListParamsFilter]'s query parameters as
// `url.Values`.
func (r UacConnectionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by connection_name using nested operations
type UacConnectionListParamsFilterConnectionName struct {
	// If present, connections with <code>connection_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UacConnectionListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r UacConnectionListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. By default sorting direction is ascending.
// To have the results sorted in descending order add the <code> -</code>
// prefix.<br/><br/> That is: <ul>
//
//	<li>
//	  <code>connection_name</code>: sorts the result by the
//	  <code>connection_name</code> field in ascending order.
//	</li>
//
//	<li>
//	  <code>-connection_name</code>: sorts the result by the
//	  <code>connection_name</code> field in descending order.
//	</li>
//
// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
type UacConnectionListParamsSort string

const (
	UacConnectionListParamsSortCreatedAt      UacConnectionListParamsSort = "created_at"
	UacConnectionListParamsSortConnectionName UacConnectionListParamsSort = "connection_name"
	UacConnectionListParamsSortActive         UacConnectionListParamsSort = "active"
)
