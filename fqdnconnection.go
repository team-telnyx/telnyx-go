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

// FQDN connection operations
//
// FqdnConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFqdnConnectionService] method instead.
type FqdnConnectionService struct {
	Options []option.RequestOption
}

// NewFqdnConnectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFqdnConnectionService(opts ...option.RequestOption) (r FqdnConnectionService) {
	r = FqdnConnectionService{}
	r.Options = opts
	return
}

// Creates a FQDN connection.
func (r *FqdnConnectionService) New(ctx context.Context, body FqdnConnectionNewParams, opts ...option.RequestOption) (res *FqdnConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fqdn_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing FQDN connection.
func (r *FqdnConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FqdnConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fqdn_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing FQDN connection.
func (r *FqdnConnectionService) Update(ctx context.Context, id string, body FqdnConnectionUpdateParams, opts ...option.RequestOption) (res *FqdnConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fqdn_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your FQDN connections.
func (r *FqdnConnectionService) List(ctx context.Context, query FqdnConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[FqdnConnection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "fqdn_connections"
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

// Returns a list of your FQDN connections.
func (r *FqdnConnectionService) ListAutoPaging(ctx context.Context, query FqdnConnectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[FqdnConnection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an FQDN connection.
func (r *FqdnConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FqdnConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fqdn_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FqdnConnection struct {
	// A user-assigned name to help manage the connection.
	ConnectionName string `json:"connection_name" api:"required"`
	// Identifies the resource.
	ID string `json:"id"`
	// Defaults to true
	Active bool `json:"active"`
	// Indicates whether DTMF timestamp adjustment is enabled.
	AdjustDtmfTimestamp bool `json:"adjust_dtmf_timestamp"`
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
	// Indicates whether call cost calculation is enabled.
	CallCostEnabled bool `json:"call_cost_enabled"`
	// Specifies if call cost webhooks should be sent for this connection.
	CallCostInWebhooks bool `json:"call_cost_in_webhooks"`
	// ISO 8601 formatted date indicating when the resource was created.
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
	// Indicates whether DTMF duration should be ignored.
	IgnoreDtmfDuration bool `json:"ignore_dtmf_duration"`
	// Indicates whether the mark bit should be ignored.
	IgnoreMarkBit bool        `json:"ignore_mark_bit"`
	Inbound       InboundFqdn `json:"inbound"`
	// The uuid of the push credential for Ios
	IosPushCredentialID string `json:"ios_push_credential_id" api:"nullable"`
	// Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams
	// of SIP Trunking calls. The feature is off unless enabled. You may define min and
	// max values in msec for customized buffering behaviors. Larger values add latency
	// but tolerate more jitter, while smaller values reduce latency but are more
	// sensitive to jitter and reordering.
	JitterBuffer shared.ConnectionJitterBuffer `json:"jitter_buffer"`
	// The connection is enabled for Microsoft Teams Direct Routing.
	MicrosoftTeamsSbc bool `json:"microsoft_teams_sbc"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression FqdnConnectionNoiseSuppression `json:"noise_suppression"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetails `json:"noise_suppression_details"`
	// Enable on-net T38 if you prefer that the sender and receiver negotiate T38
	// directly when both are on the Telnyx network. If this is disabled, Telnyx will
	// be able to use T38 on just one leg of the call according to each leg's settings.
	OnnetT38PassthroughEnabled bool         `json:"onnet_t38_passthrough_enabled"`
	Outbound                   OutboundFqdn `json:"outbound"`
	// The password for the FQDN connection.
	Password string `json:"password"`
	// Identifies the type of the resource.
	RecordType   string                 `json:"record_type"`
	RtcpSettings ConnectionRtcpSettings `json:"rtcp_settings"`
	// Defines if codecs should be passed on stream change.
	RtpPassCodecsOnStreamChange bool `json:"rtp_pass_codecs_on_stream_change"`
	// Indicates whether normalized timestamps should be sent.
	SendNormalizedTimestamps bool `json:"send_normalized_timestamps"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// Indicates whether third-party control is enabled.
	ThirdPartyControlEnabled bool `json:"third_party_control_enabled"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
	// FQDN authentication.
	//
	// Any of "UDP", "TCP", "TLS".
	TransportProtocol TransportProtocol `json:"transport_protocol"`
	// The name for the TXT record associated with the FQDN connection.
	TxtName string `json:"txt_name"`
	// The time to live for the TXT record associated with the FQDN connection.
	TxtTtl int64 `json:"txt_ttl"`
	// The value for the TXT record associated with the FQDN connection.
	TxtValue string `json:"txt_value"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// The username for the FQDN connection.
	UserName string `json:"user_name"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion WebhookAPIVersion `json:"webhook_api_version"`
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
		ConnectionName                   respjson.Field
		ID                               respjson.Field
		Active                           respjson.Field
		AdjustDtmfTimestamp              respjson.Field
		AnchorsiteOverride               respjson.Field
		AndroidPushCredentialID          respjson.Field
		CallCostEnabled                  respjson.Field
		CallCostInWebhooks               respjson.Field
		CreatedAt                        respjson.Field
		DefaultOnHoldComfortNoiseEnabled respjson.Field
		DtmfType                         respjson.Field
		EncodeContactHeaderEnabled       respjson.Field
		EncryptedMedia                   respjson.Field
		IgnoreDtmfDuration               respjson.Field
		IgnoreMarkBit                    respjson.Field
		Inbound                          respjson.Field
		IosPushCredentialID              respjson.Field
		JitterBuffer                     respjson.Field
		MicrosoftTeamsSbc                respjson.Field
		NoiseSuppression                 respjson.Field
		NoiseSuppressionDetails          respjson.Field
		OnnetT38PassthroughEnabled       respjson.Field
		Outbound                         respjson.Field
		Password                         respjson.Field
		RecordType                       respjson.Field
		RtcpSettings                     respjson.Field
		RtpPassCodecsOnStreamChange      respjson.Field
		SendNormalizedTimestamps         respjson.Field
		Tags                             respjson.Field
		ThirdPartyControlEnabled         respjson.Field
		TransportProtocol                respjson.Field
		TxtName                          respjson.Field
		TxtTtl                           respjson.Field
		TxtValue                         respjson.Field
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
func (r FqdnConnection) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type FqdnConnectionNoiseSuppression string

const (
	FqdnConnectionNoiseSuppressionInbound  FqdnConnectionNoiseSuppression = "inbound"
	FqdnConnectionNoiseSuppressionOutbound FqdnConnectionNoiseSuppression = "outbound"
	FqdnConnectionNoiseSuppressionBoth     FqdnConnectionNoiseSuppression = "both"
	FqdnConnectionNoiseSuppressionDisabled FqdnConnectionNoiseSuppression = "disabled"
)

type InboundFqdn struct {
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat InboundFqdnAniNumberFormat `json:"ani_number_format"`
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit" api:"nullable"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs"`
	// The default primary FQDN to use for the number. Only settable if the connection
	// is of FQDN type. Value must be the ID of an FQDN set on the connection.
	DefaultPrimaryFqdnID string `json:"default_primary_fqdn_id" api:"nullable"`
	// Default routing method to be used when a number is associated with the
	// connection. Must be one of the routing method types or null, other values are
	// not allowed.
	//
	// Any of "sequential", "round-robin".
	DefaultRoutingMethod InboundFqdnDefaultRoutingMethod `json:"default_routing_method" api:"nullable"`
	// The default secondary FQDN to use for the number. Only settable if the
	// connection is of FQDN type. Value must be the ID of an FQDN set on the
	// connection.
	DefaultSecondaryFqdnID string `json:"default_secondary_fqdn_id" api:"nullable"`
	// The default tertiary FQDN to use for the number. Only settable if the connection
	// is of FQDN type. Value must be the ID of an FQDN set on the connection.
	DefaultTertiaryFqdnID string `json:"default_tertiary_fqdn_id" api:"nullable"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat InboundFqdnDnisNumberFormat `json:"dnis_number_format"`
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
	// Defaults to true.
	SipCompactHeadersEnabled bool `json:"sip_compact_headers_enabled"`
	// Selects which `sip_region` to receive inbound calls from. If null, the default
	// region (US) will be used.
	//
	// Any of "US", "Europe", "Australia".
	SipRegion InboundFqdnSipRegion `json:"sip_region"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection,
	// in the same way a phone number is used, from a SIP endpoint. Example: the
	// subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using
	// the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any
	// alphanumeric value. Please note TLS encrypted calls are not allowed for
	// subdomain calls.
	SipSubdomain string `json:"sip_subdomain" api:"nullable"`
	// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
	// the public Internet) or "Only my connections" (any connection assigned to the
	// same Telnyx user).
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings InboundFqdnSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs int64 `json:"timeout_1xx_secs"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs int64 `json:"timeout_2xx_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniNumberFormat             respjson.Field
		ChannelLimit                respjson.Field
		Codecs                      respjson.Field
		DefaultPrimaryFqdnID        respjson.Field
		DefaultRoutingMethod        respjson.Field
		DefaultSecondaryFqdnID      respjson.Field
		DefaultTertiaryFqdnID       respjson.Field
		DnisNumberFormat            respjson.Field
		GenerateRingbackTone        respjson.Field
		IsupHeadersEnabled          respjson.Field
		PrackEnabled                respjson.Field
		ShakenStirEnabled           respjson.Field
		SipCompactHeadersEnabled    respjson.Field
		SipRegion                   respjson.Field
		SipSubdomain                respjson.Field
		SipSubdomainReceiveSettings respjson.Field
		Timeout1xxSecs              respjson.Field
		Timeout2xxSecs              respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundFqdn) RawJSON() string { return r.JSON.raw }
func (r *InboundFqdn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundFqdn to a InboundFqdnParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundFqdnParam.Overrides()
func (r InboundFqdn) ToParam() InboundFqdnParam {
	return param.Override[InboundFqdnParam](json.RawMessage(r.RawJSON()))
}

// This setting allows you to set the format with which the caller's number (ANI)
// is sent for inbound phone calls.
type InboundFqdnAniNumberFormat string

const (
	InboundFqdnAniNumberFormatPlusE164         InboundFqdnAniNumberFormat = "+E.164"
	InboundFqdnAniNumberFormatE164             InboundFqdnAniNumberFormat = "E.164"
	InboundFqdnAniNumberFormatPlusE164National InboundFqdnAniNumberFormat = "+E.164-national"
	InboundFqdnAniNumberFormatE164National     InboundFqdnAniNumberFormat = "E.164-national"
)

// Default routing method to be used when a number is associated with the
// connection. Must be one of the routing method types or null, other values are
// not allowed.
type InboundFqdnDefaultRoutingMethod string

const (
	InboundFqdnDefaultRoutingMethodSequential InboundFqdnDefaultRoutingMethod = "sequential"
	InboundFqdnDefaultRoutingMethodRoundRobin InboundFqdnDefaultRoutingMethod = "round-robin"
)

type InboundFqdnDnisNumberFormat string

const (
	InboundFqdnDnisNumberFormatPlusE164    InboundFqdnDnisNumberFormat = "+e164"
	InboundFqdnDnisNumberFormatE164        InboundFqdnDnisNumberFormat = "e164"
	InboundFqdnDnisNumberFormatNational    InboundFqdnDnisNumberFormat = "national"
	InboundFqdnDnisNumberFormatSipUsername InboundFqdnDnisNumberFormat = "sip_username"
)

// Selects which `sip_region` to receive inbound calls from. If null, the default
// region (US) will be used.
type InboundFqdnSipRegion string

const (
	InboundFqdnSipRegionUs        InboundFqdnSipRegion = "US"
	InboundFqdnSipRegionEurope    InboundFqdnSipRegion = "Europe"
	InboundFqdnSipRegionAustralia InboundFqdnSipRegion = "Australia"
)

// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
// the public Internet) or "Only my connections" (any connection assigned to the
// same Telnyx user).
type InboundFqdnSipSubdomainReceiveSettings string

const (
	InboundFqdnSipSubdomainReceiveSettingsOnlyMyConnections InboundFqdnSipSubdomainReceiveSettings = "only_my_connections"
	InboundFqdnSipSubdomainReceiveSettingsFromAnyone        InboundFqdnSipSubdomainReceiveSettings = "from_anyone"
)

type InboundFqdnParam struct {
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// The default primary FQDN to use for the number. Only settable if the connection
	// is of FQDN type. Value must be the ID of an FQDN set on the connection.
	DefaultPrimaryFqdnID param.Opt[string] `json:"default_primary_fqdn_id,omitzero"`
	// The default secondary FQDN to use for the number. Only settable if the
	// connection is of FQDN type. Value must be the ID of an FQDN set on the
	// connection.
	DefaultSecondaryFqdnID param.Opt[string] `json:"default_secondary_fqdn_id,omitzero"`
	// The default tertiary FQDN to use for the number. Only settable if the connection
	// is of FQDN type. Value must be the ID of an FQDN set on the connection.
	DefaultTertiaryFqdnID param.Opt[string] `json:"default_tertiary_fqdn_id,omitzero"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection,
	// in the same way a phone number is used, from a SIP endpoint. Example: the
	// subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using
	// the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any
	// alphanumeric value. Please note TLS encrypted calls are not allowed for
	// subdomain calls.
	SipSubdomain param.Opt[string] `json:"sip_subdomain,omitzero"`
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
	// Default routing method to be used when a number is associated with the
	// connection. Must be one of the routing method types or null, other values are
	// not allowed.
	//
	// Any of "sequential", "round-robin".
	DefaultRoutingMethod InboundFqdnDefaultRoutingMethod `json:"default_routing_method,omitzero"`
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat InboundFqdnAniNumberFormat `json:"ani_number_format,omitzero"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs,omitzero"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat InboundFqdnDnisNumberFormat `json:"dnis_number_format,omitzero"`
	// Selects which `sip_region` to receive inbound calls from. If null, the default
	// region (US) will be used.
	//
	// Any of "US", "Europe", "Australia".
	SipRegion InboundFqdnSipRegion `json:"sip_region,omitzero"`
	// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
	// the public Internet) or "Only my connections" (any connection assigned to the
	// same Telnyx user).
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings InboundFqdnSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings,omitzero"`
	paramObj
}

func (r InboundFqdnParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundFqdnParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundFqdnParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundFqdn struct {
	// Set a phone number as the ani_override value to override caller id number on
	// outbound calls.
	AniOverride string `json:"ani_override"`
	// Specifies when we should apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType OutboundFqdnAniOverrideType `json:"ani_override_type"`
	// Forces all SIP calls originated on this connection to be \"parked\" instead of
	// \"bridged\" to the destination specified on the URI. Parked calls will return
	// ringback to the caller and will await for a Call Control command to define which
	// action will be taken next.
	CallParkingEnabled bool `json:"call_parking_enabled" api:"nullable"`
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Enable use of SRTP for encryption. Cannot be set if the transport_portocol is
	// TLS.
	//
	// Any of "SRTP".
	EncryptedMedia EncryptedMedia `json:"encrypted_media" api:"nullable"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone bool `json:"generate_ringback_tone"`
	// When set, ringback will not wait for indication before sending ringback tone to
	// calling party.
	InstantRingbackEnabled bool `json:"instant_ringback_enabled"`
	// Any of "credential-authentication", "ip-authentication".
	IPAuthenticationMethod OutboundFqdnIPAuthenticationMethod `json:"ip_authentication_method"`
	IPAuthenticationToken  string                             `json:"ip_authentication_token"`
	// A 2-character country code specifying the country whose national dialing rules
	// should be used. For example, if set to `US` then any US number can be dialed
	// without preprending +1 to the number. When left blank, Telnyx will try US and GB
	// dialing rules, in that order, by default.",
	Localization string `json:"localization"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite. By default, Telnyx
	// will send the re-invite. If set to `customer`, the caller is expected to send
	// the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource OutboundFqdnT38ReinviteSource `json:"t38_reinvite_source"`
	// Numerical chars only, exactly 4 characters.
	TechPrefix string `json:"tech_prefix"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs int64 `json:"timeout_1xx_secs"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs int64 `json:"timeout_2xx_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniOverride            respjson.Field
		AniOverrideType        respjson.Field
		CallParkingEnabled     respjson.Field
		ChannelLimit           respjson.Field
		EncryptedMedia         respjson.Field
		GenerateRingbackTone   respjson.Field
		InstantRingbackEnabled respjson.Field
		IPAuthenticationMethod respjson.Field
		IPAuthenticationToken  respjson.Field
		Localization           respjson.Field
		OutboundVoiceProfileID respjson.Field
		T38ReinviteSource      respjson.Field
		TechPrefix             respjson.Field
		Timeout1xxSecs         respjson.Field
		Timeout2xxSecs         respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundFqdn) RawJSON() string { return r.JSON.raw }
func (r *OutboundFqdn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutboundFqdn to a OutboundFqdnParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutboundFqdnParam.Overrides()
func (r OutboundFqdn) ToParam() OutboundFqdnParam {
	return param.Override[OutboundFqdnParam](json.RawMessage(r.RawJSON()))
}

// Specifies when we should apply your ani_override setting. Only applies when
// ani_override is not blank.
type OutboundFqdnAniOverrideType string

const (
	OutboundFqdnAniOverrideTypeAlways    OutboundFqdnAniOverrideType = "always"
	OutboundFqdnAniOverrideTypeNormal    OutboundFqdnAniOverrideType = "normal"
	OutboundFqdnAniOverrideTypeEmergency OutboundFqdnAniOverrideType = "emergency"
)

type OutboundFqdnIPAuthenticationMethod string

const (
	OutboundFqdnIPAuthenticationMethodCredentialAuthentication OutboundFqdnIPAuthenticationMethod = "credential-authentication"
	OutboundFqdnIPAuthenticationMethodIPAuthentication         OutboundFqdnIPAuthenticationMethod = "ip-authentication"
)

// This setting only affects connections with Fax-type Outbound Voice Profiles. The
// setting dictates whether or not Telnyx sends a t.38 reinvite. By default, Telnyx
// will send the re-invite. If set to `customer`, the caller is expected to send
// the t.38 reinvite.
type OutboundFqdnT38ReinviteSource string

const (
	OutboundFqdnT38ReinviteSourceTelnyx         OutboundFqdnT38ReinviteSource = "telnyx"
	OutboundFqdnT38ReinviteSourceCustomer       OutboundFqdnT38ReinviteSource = "customer"
	OutboundFqdnT38ReinviteSourceDisabled       OutboundFqdnT38ReinviteSource = "disabled"
	OutboundFqdnT38ReinviteSourcePassthru       OutboundFqdnT38ReinviteSource = "passthru"
	OutboundFqdnT38ReinviteSourceCallerPassthru OutboundFqdnT38ReinviteSource = "caller-passthru"
	OutboundFqdnT38ReinviteSourceCalleePassthru OutboundFqdnT38ReinviteSource = "callee-passthru"
)

type OutboundFqdnParam struct {
	// Forces all SIP calls originated on this connection to be \"parked\" instead of
	// \"bridged\" to the destination specified on the URI. Parked calls will return
	// ringback to the caller and will await for a Call Control command to define which
	// action will be taken next.
	CallParkingEnabled param.Opt[bool] `json:"call_parking_enabled,omitzero"`
	// Set a phone number as the ani_override value to override caller id number on
	// outbound calls.
	AniOverride param.Opt[string] `json:"ani_override,omitzero"`
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone param.Opt[bool] `json:"generate_ringback_tone,omitzero"`
	// When set, ringback will not wait for indication before sending ringback tone to
	// calling party.
	InstantRingbackEnabled param.Opt[bool]   `json:"instant_ringback_enabled,omitzero"`
	IPAuthenticationToken  param.Opt[string] `json:"ip_authentication_token,omitzero"`
	// A 2-character country code specifying the country whose national dialing rules
	// should be used. For example, if set to `US` then any US number can be dialed
	// without preprending +1 to the number. When left blank, Telnyx will try US and GB
	// dialing rules, in that order, by default.",
	Localization param.Opt[string] `json:"localization,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	// Numerical chars only, exactly 4 characters.
	TechPrefix param.Opt[string] `json:"tech_prefix,omitzero"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs param.Opt[int64] `json:"timeout_1xx_secs,omitzero"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs param.Opt[int64] `json:"timeout_2xx_secs,omitzero"`
	// Enable use of SRTP for encryption. Cannot be set if the transport_portocol is
	// TLS.
	//
	// Any of "SRTP".
	EncryptedMedia EncryptedMedia `json:"encrypted_media,omitzero"`
	// Specifies when we should apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType OutboundFqdnAniOverrideType `json:"ani_override_type,omitzero"`
	// Any of "credential-authentication", "ip-authentication".
	IPAuthenticationMethod OutboundFqdnIPAuthenticationMethod `json:"ip_authentication_method,omitzero"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite. By default, Telnyx
	// will send the re-invite. If set to `customer`, the caller is expected to send
	// the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource OutboundFqdnT38ReinviteSource `json:"t38_reinvite_source,omitzero"`
	paramObj
}

func (r OutboundFqdnParam) MarshalJSON() (data []byte, err error) {
	type shadow OutboundFqdnParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutboundFqdnParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
// FQDN authentication.
type TransportProtocol string

const (
	TransportProtocolUdp TransportProtocol = "UDP"
	TransportProtocolTcp TransportProtocol = "TCP"
	TransportProtocolTls TransportProtocol = "TLS"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type WebhookAPIVersion string

const (
	WebhookAPIVersionV1 WebhookAPIVersion = "1"
	WebhookAPIVersionV2 WebhookAPIVersion = "2"
)

type FqdnConnectionNewResponse struct {
	Data FqdnConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnConnectionGetResponse struct {
	Data FqdnConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnConnectionUpdateResponse struct {
	Data FqdnConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnConnectionDeleteResponse struct {
	Data FqdnConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnConnectionNewParams struct {
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
	// When enabled, the connection will be created for Microsoft Teams Direct Routing.
	// A \*.mstsbc.telnyx.tech FQDN will be created for the connection automatically.
	MicrosoftTeamsSbc param.Opt[bool] `json:"microsoft_teams_sbc,omitzero"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly
	// if both are on the Telnyx network. If this is disabled, Telnyx will be able to
	// use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled param.Opt[bool] `json:"onnet_t38_passthrough_enabled,omitzero"`
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
	DtmfType DtmfType         `json:"dtmf_type,omitzero"`
	Inbound  InboundFqdnParam `json:"inbound,omitzero"`
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
	NoiseSuppression FqdnConnectionNewParamsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetailsParam `json:"noise_suppression_details,omitzero"`
	Outbound                OutboundFqdnParam                             `json:"outbound,omitzero"`
	RtcpSettings            ConnectionRtcpSettingsParam                   `json:"rtcp_settings,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
	// FQDN authentication.
	//
	// Any of "UDP", "TCP", "TLS".
	TransportProtocol TransportProtocol `json:"transport_protocol,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion WebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r FqdnConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FqdnConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FqdnConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type FqdnConnectionNewParamsNoiseSuppression string

const (
	FqdnConnectionNewParamsNoiseSuppressionInbound  FqdnConnectionNewParamsNoiseSuppression = "inbound"
	FqdnConnectionNewParamsNoiseSuppressionOutbound FqdnConnectionNewParamsNoiseSuppression = "outbound"
	FqdnConnectionNewParamsNoiseSuppressionBoth     FqdnConnectionNewParamsNoiseSuppression = "both"
	FqdnConnectionNewParamsNoiseSuppressionDisabled FqdnConnectionNewParamsNoiseSuppression = "disabled"
)

type FqdnConnectionUpdateParams struct {
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
	// Enable on-net T38 if you prefer that the sender and receiver negotiate T38
	// directly when both are on the Telnyx network. If this is disabled, Telnyx will
	// be able to use T38 on just one leg of the call according to each leg's settings.
	OnnetT38PassthroughEnabled param.Opt[bool] `json:"onnet_t38_passthrough_enabled,omitzero"`
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
	DtmfType DtmfType         `json:"dtmf_type,omitzero"`
	Inbound  InboundFqdnParam `json:"inbound,omitzero"`
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
	NoiseSuppression FqdnConnectionUpdateParamsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetailsParam `json:"noise_suppression_details,omitzero"`
	Outbound                OutboundFqdnParam                             `json:"outbound,omitzero"`
	RtcpSettings            ConnectionRtcpSettingsParam                   `json:"rtcp_settings,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
	// FQDN authentication.
	//
	// Any of "UDP", "TCP", "TLS".
	TransportProtocol TransportProtocol `json:"transport_protocol,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion WebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r FqdnConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FqdnConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FqdnConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type FqdnConnectionUpdateParamsNoiseSuppression string

const (
	FqdnConnectionUpdateParamsNoiseSuppressionInbound  FqdnConnectionUpdateParamsNoiseSuppression = "inbound"
	FqdnConnectionUpdateParamsNoiseSuppressionOutbound FqdnConnectionUpdateParamsNoiseSuppression = "outbound"
	FqdnConnectionUpdateParamsNoiseSuppressionBoth     FqdnConnectionUpdateParamsNoiseSuppression = "both"
	FqdnConnectionUpdateParamsNoiseSuppressionDisabled FqdnConnectionUpdateParamsNoiseSuppression = "disabled"
)

type FqdnConnectionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
	// filter[outbound.outbound_voice_profile_id]
	Filter FqdnConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
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
	Sort FqdnConnectionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FqdnConnectionListParams]'s query parameters as
// `url.Values`.
func (r FqdnConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
// filter[outbound.outbound_voice_profile_id]
type FqdnConnectionListParamsFilter struct {
	// If present, connections with an `fqdn` that equals the given value will be
	// returned. Matching is case-sensitive, and the full string must match.
	Fqdn param.Opt[string] `query:"fqdn,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" json:"-"`
	// Filter by connection_name using nested operations
	ConnectionName FqdnConnectionListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FqdnConnectionListParamsFilter]'s query parameters as
// `url.Values`.
func (r FqdnConnectionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by connection_name using nested operations
type FqdnConnectionListParamsFilterConnectionName struct {
	// If present, connections with <code>connection_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FqdnConnectionListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r FqdnConnectionListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
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
type FqdnConnectionListParamsSort string

const (
	FqdnConnectionListParamsSortCreatedAt      FqdnConnectionListParamsSort = "created_at"
	FqdnConnectionListParamsSortConnectionName FqdnConnectionListParamsSort = "connection_name"
	FqdnConnectionListParamsSortActive         FqdnConnectionListParamsSort = "active"
)
