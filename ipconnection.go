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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// IPConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPConnectionService] method instead.
type IPConnectionService struct {
	Options []option.RequestOption
}

// NewIPConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIPConnectionService(opts ...option.RequestOption) (r IPConnectionService) {
	r = IPConnectionService{}
	r.Options = opts
	return
}

// Creates an IP connection.
func (r *IPConnectionService) New(ctx context.Context, body IPConnectionNewParams, opts ...option.RequestOption) (res *IPConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ip_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing ip connection.
func (r *IPConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IPConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ip_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing IP connection.
func (r *IPConnectionService) Update(ctx context.Context, id string, body IPConnectionUpdateParams, opts ...option.RequestOption) (res *IPConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ip_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your IP connections.
func (r *IPConnectionService) List(ctx context.Context, query IPConnectionListParams, opts ...option.RequestOption) (res *IPConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ip_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing IP connection.
func (r *IPConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *IPConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ip_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type InboundIP struct {
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat InboundIPAniNumberFormat `json:"ani_number_format"`
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs"`
	// The default primary IP to use for the number. Only settable if the connection is
	// of IP authentication type. Value must be the ID of an authorized IP set on the
	// connection.
	DefaultPrimaryIPID string `json:"default_primary_ip_id"`
	// Default routing method to be used when a number is associated with the
	// connection. Must be one of the routing method types or left blank, other values
	// are not allowed.
	//
	// Any of "sequential", "round-robin".
	DefaultRoutingMethod InboundIPDefaultRoutingMethod `json:"default_routing_method"`
	// The default secondary IP to use for the number. Only settable if the connection
	// is of IP authentication type. Value must be the ID of an authorized IP set on
	// the connection.
	DefaultSecondaryIPID string `json:"default_secondary_ip_id"`
	// The default tertiary IP to use for the number. Only settable if the connection
	// is of IP authentication type. Value must be the ID of an authorized IP set on
	// the connection.
	DefaultTertiaryIPID string `json:"default_tertiary_ip_id"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat InboundIPDnisNumberFormat `json:"dnis_number_format"`
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
	SipRegion InboundIPSipRegion `json:"sip_region"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection,
	// in the same way a phone number is used, from a SIP endpoint. Example: the
	// subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using
	// the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any
	// alphanumeric value. Please note TLS encrypted calls are not allowed for
	// subdomain calls.
	SipSubdomain string `json:"sip_subdomain"`
	// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
	// the public Internet) or "Only my connections" (any connection assigned to the
	// same Telnyx user).
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings InboundIPSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs int64 `json:"timeout_1xx_secs"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs int64 `json:"timeout_2xx_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniNumberFormat             respjson.Field
		ChannelLimit                respjson.Field
		Codecs                      respjson.Field
		DefaultPrimaryIPID          respjson.Field
		DefaultRoutingMethod        respjson.Field
		DefaultSecondaryIPID        respjson.Field
		DefaultTertiaryIPID         respjson.Field
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
func (r InboundIP) RawJSON() string { return r.JSON.raw }
func (r *InboundIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundIP to a InboundIPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundIPParam.Overrides()
func (r InboundIP) ToParam() InboundIPParam {
	return param.Override[InboundIPParam](json.RawMessage(r.RawJSON()))
}

// This setting allows you to set the format with which the caller's number (ANI)
// is sent for inbound phone calls.
type InboundIPAniNumberFormat string

const (
	InboundIPAniNumberFormatPlusE164         InboundIPAniNumberFormat = "+E.164"
	InboundIPAniNumberFormatE164             InboundIPAniNumberFormat = "E.164"
	InboundIPAniNumberFormatPlusE164National InboundIPAniNumberFormat = "+E.164-national"
	InboundIPAniNumberFormatE164National     InboundIPAniNumberFormat = "E.164-national"
)

// Default routing method to be used when a number is associated with the
// connection. Must be one of the routing method types or left blank, other values
// are not allowed.
type InboundIPDefaultRoutingMethod string

const (
	InboundIPDefaultRoutingMethodSequential InboundIPDefaultRoutingMethod = "sequential"
	InboundIPDefaultRoutingMethodRoundRobin InboundIPDefaultRoutingMethod = "round-robin"
)

type InboundIPDnisNumberFormat string

const (
	InboundIPDnisNumberFormatPlusE164    InboundIPDnisNumberFormat = "+e164"
	InboundIPDnisNumberFormatE164        InboundIPDnisNumberFormat = "e164"
	InboundIPDnisNumberFormatNational    InboundIPDnisNumberFormat = "national"
	InboundIPDnisNumberFormatSipUsername InboundIPDnisNumberFormat = "sip_username"
)

// Selects which `sip_region` to receive inbound calls from. If null, the default
// region (US) will be used.
type InboundIPSipRegion string

const (
	InboundIPSipRegionUs        InboundIPSipRegion = "US"
	InboundIPSipRegionEurope    InboundIPSipRegion = "Europe"
	InboundIPSipRegionAustralia InboundIPSipRegion = "Australia"
)

// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
// the public Internet) or "Only my connections" (any connection assigned to the
// same Telnyx user).
type InboundIPSipSubdomainReceiveSettings string

const (
	InboundIPSipSubdomainReceiveSettingsOnlyMyConnections InboundIPSipSubdomainReceiveSettings = "only_my_connections"
	InboundIPSipSubdomainReceiveSettingsFromAnyone        InboundIPSipSubdomainReceiveSettings = "from_anyone"
)

type InboundIPParam struct {
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// The default primary IP to use for the number. Only settable if the connection is
	// of IP authentication type. Value must be the ID of an authorized IP set on the
	// connection.
	DefaultPrimaryIPID param.Opt[string] `json:"default_primary_ip_id,omitzero"`
	// The default secondary IP to use for the number. Only settable if the connection
	// is of IP authentication type. Value must be the ID of an authorized IP set on
	// the connection.
	DefaultSecondaryIPID param.Opt[string] `json:"default_secondary_ip_id,omitzero"`
	// The default tertiary IP to use for the number. Only settable if the connection
	// is of IP authentication type. Value must be the ID of an authorized IP set on
	// the connection.
	DefaultTertiaryIPID param.Opt[string] `json:"default_tertiary_ip_id,omitzero"`
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
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection,
	// in the same way a phone number is used, from a SIP endpoint. Example: the
	// subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using
	// the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any
	// alphanumeric value. Please note TLS encrypted calls are not allowed for
	// subdomain calls.
	SipSubdomain param.Opt[string] `json:"sip_subdomain,omitzero"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs param.Opt[int64] `json:"timeout_1xx_secs,omitzero"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs param.Opt[int64] `json:"timeout_2xx_secs,omitzero"`
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat InboundIPAniNumberFormat `json:"ani_number_format,omitzero"`
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
	DefaultRoutingMethod InboundIPDefaultRoutingMethod `json:"default_routing_method,omitzero"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat InboundIPDnisNumberFormat `json:"dnis_number_format,omitzero"`
	// Selects which `sip_region` to receive inbound calls from. If null, the default
	// region (US) will be used.
	//
	// Any of "US", "Europe", "Australia".
	SipRegion InboundIPSipRegion `json:"sip_region,omitzero"`
	// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
	// the public Internet) or "Only my connections" (any connection assigned to the
	// same Telnyx user).
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings InboundIPSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings,omitzero"`
	paramObj
}

func (r InboundIPParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundIPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundIPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnection struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"int64"`
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
	ConnectionName     string             `json:"connection_name"`
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
	EncryptedMedia EncryptedMedia `json:"encrypted_media,nullable"`
	Inbound        InboundIP      `json:"inbound"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly
	// if both are on the Telnyx network. If this is disabled, Telnyx will be able to
	// use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled bool       `json:"onnet_t38_passthrough_enabled"`
	Outbound                   OutboundIP `json:"outbound"`
	// Identifies the type of the resource.
	RecordType   string                 `json:"record_type"`
	RtcpSettings ConnectionRtcpSettings `json:"rtcp_settings"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
	// FQDN authentication.
	//
	// Any of "UDP", "TCP", "TLS".
	TransportProtocol IPConnectionTransportProtocol `json:"transport_protocol"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion IPConnectionWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable" format:"uri"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		Active                           respjson.Field
		AnchorsiteOverride               respjson.Field
		ConnectionName                   respjson.Field
		CreatedAt                        respjson.Field
		DefaultOnHoldComfortNoiseEnabled respjson.Field
		DtmfType                         respjson.Field
		EncodeContactHeaderEnabled       respjson.Field
		EncryptedMedia                   respjson.Field
		Inbound                          respjson.Field
		OnnetT38PassthroughEnabled       respjson.Field
		Outbound                         respjson.Field
		RecordType                       respjson.Field
		RtcpSettings                     respjson.Field
		Tags                             respjson.Field
		TransportProtocol                respjson.Field
		UpdatedAt                        respjson.Field
		WebhookAPIVersion                respjson.Field
		WebhookEventFailoverURL          respjson.Field
		WebhookEventURL                  respjson.Field
		WebhookTimeoutSecs               respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPConnection) RawJSON() string { return r.JSON.raw }
func (r *IPConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
// FQDN authentication.
type IPConnectionTransportProtocol string

const (
	IPConnectionTransportProtocolUdp IPConnectionTransportProtocol = "UDP"
	IPConnectionTransportProtocolTcp IPConnectionTransportProtocol = "TCP"
	IPConnectionTransportProtocolTls IPConnectionTransportProtocol = "TLS"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type IPConnectionWebhookAPIVersion string

const (
	IPConnectionWebhookAPIVersion1 IPConnectionWebhookAPIVersion = "1"
	IPConnectionWebhookAPIVersion2 IPConnectionWebhookAPIVersion = "2"
)

type OutboundIP struct {
	// Set a phone number as the ani_override value to override caller id number on
	// outbound calls.
	AniOverride string `json:"ani_override"`
	// Specifies when we apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType OutboundIPAniOverrideType `json:"ani_override_type"`
	// Forces all SIP calls originated on this connection to be "parked" instead of
	// "bridged" to the destination specified on the URI. Parked calls will return
	// ringback to the caller and will await for a Call Control command to define which
	// action will be taken next.
	CallParkingEnabled bool `json:"call_parking_enabled,nullable"`
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone bool `json:"generate_ringback_tone"`
	// When set, ringback will not wait for indication before sending ringback tone to
	// calling party.
	InstantRingbackEnabled bool `json:"instant_ringback_enabled"`
	// Any of "tech-prefixp-charge-info", "token".
	IPAuthenticationMethod OutboundIPIPAuthenticationMethod `json:"ip_authentication_method"`
	IPAuthenticationToken  string                           `json:"ip_authentication_token"`
	// A 2-character country code specifying the country whose national dialing rules
	// should be used. For example, if set to `US` then any US number can be dialed
	// without preprending +1 to the number. When left blank, Telnyx will try US and GB
	// dialing rules, in that order, by default.
	Localization string `json:"localization"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id" format:"int64"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
	// default, Telnyx will send the re-invite. If set to `customer`, the caller is
	// expected to send the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource OutboundIPT38ReinviteSource `json:"t38_reinvite_source"`
	// Numerical chars only, exactly 4 characters.
	TechPrefix string `json:"tech_prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniOverride            respjson.Field
		AniOverrideType        respjson.Field
		CallParkingEnabled     respjson.Field
		ChannelLimit           respjson.Field
		GenerateRingbackTone   respjson.Field
		InstantRingbackEnabled respjson.Field
		IPAuthenticationMethod respjson.Field
		IPAuthenticationToken  respjson.Field
		Localization           respjson.Field
		OutboundVoiceProfileID respjson.Field
		T38ReinviteSource      respjson.Field
		TechPrefix             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundIP) RawJSON() string { return r.JSON.raw }
func (r *OutboundIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutboundIP to a OutboundIPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutboundIPParam.Overrides()
func (r OutboundIP) ToParam() OutboundIPParam {
	return param.Override[OutboundIPParam](json.RawMessage(r.RawJSON()))
}

// Specifies when we apply your ani_override setting. Only applies when
// ani_override is not blank.
type OutboundIPAniOverrideType string

const (
	OutboundIPAniOverrideTypeAlways    OutboundIPAniOverrideType = "always"
	OutboundIPAniOverrideTypeNormal    OutboundIPAniOverrideType = "normal"
	OutboundIPAniOverrideTypeEmergency OutboundIPAniOverrideType = "emergency"
)

type OutboundIPIPAuthenticationMethod string

const (
	OutboundIPIPAuthenticationMethodTechPrefixpChargeInfo OutboundIPIPAuthenticationMethod = "tech-prefixp-charge-info"
	OutboundIPIPAuthenticationMethodToken                 OutboundIPIPAuthenticationMethod = "token"
)

// This setting only affects connections with Fax-type Outbound Voice Profiles. The
// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
// default, Telnyx will send the re-invite. If set to `customer`, the caller is
// expected to send the t.38 reinvite.
type OutboundIPT38ReinviteSource string

const (
	OutboundIPT38ReinviteSourceTelnyx         OutboundIPT38ReinviteSource = "telnyx"
	OutboundIPT38ReinviteSourceCustomer       OutboundIPT38ReinviteSource = "customer"
	OutboundIPT38ReinviteSourceDisabled       OutboundIPT38ReinviteSource = "disabled"
	OutboundIPT38ReinviteSourcePassthru       OutboundIPT38ReinviteSource = "passthru"
	OutboundIPT38ReinviteSourceCallerPassthru OutboundIPT38ReinviteSource = "caller-passthru"
	OutboundIPT38ReinviteSourceCalleePassthru OutboundIPT38ReinviteSource = "callee-passthru"
)

type OutboundIPParam struct {
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
	InstantRingbackEnabled param.Opt[bool]   `json:"instant_ringback_enabled,omitzero"`
	IPAuthenticationToken  param.Opt[string] `json:"ip_authentication_token,omitzero"`
	// A 2-character country code specifying the country whose national dialing rules
	// should be used. For example, if set to `US` then any US number can be dialed
	// without preprending +1 to the number. When left blank, Telnyx will try US and GB
	// dialing rules, in that order, by default.
	Localization param.Opt[string] `json:"localization,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero" format:"int64"`
	// Numerical chars only, exactly 4 characters.
	TechPrefix param.Opt[string] `json:"tech_prefix,omitzero"`
	// Specifies when we apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType OutboundIPAniOverrideType `json:"ani_override_type,omitzero"`
	// Any of "tech-prefixp-charge-info", "token".
	IPAuthenticationMethod OutboundIPIPAuthenticationMethod `json:"ip_authentication_method,omitzero"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
	// default, Telnyx will send the re-invite. If set to `customer`, the caller is
	// expected to send the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource OutboundIPT38ReinviteSource `json:"t38_reinvite_source,omitzero"`
	paramObj
}

func (r OutboundIPParam) MarshalJSON() (data []byte, err error) {
	type shadow OutboundIPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutboundIPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionNewResponse struct {
	Data IPConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *IPConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionGetResponse struct {
	Data IPConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IPConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionUpdateResponse struct {
	Data IPConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *IPConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionListResponse struct {
	Data []IPConnection                   `json:"data"`
	Meta shared.ConnectionsPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *IPConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionDeleteResponse struct {
	Data IPConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *IPConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionNewParams struct {
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
	Active         param.Opt[bool]   `json:"active,omitzero"`
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
	DtmfType     DtmfType                     `json:"dtmf_type,omitzero"`
	Inbound      IPConnectionNewParamsInbound `json:"inbound,omitzero"`
	Outbound     OutboundIPParam              `json:"outbound,omitzero"`
	RtcpSettings ConnectionRtcpSettingsParam  `json:"rtcp_settings,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
	// FQDN authentication.
	//
	// Any of "UDP", "TCP", "TLS".
	TransportProtocol IPConnectionNewParamsTransportProtocol `json:"transport_protocol,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion IPConnectionNewParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r IPConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IPConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPConnectionNewParamsInbound struct {
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
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection,
	// in the same way a phone number is used, from a SIP endpoint. Example: the
	// subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using
	// the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any
	// alphanumeric value. Please note TLS encrypted calls are not allowed for
	// subdomain calls.
	SipSubdomain param.Opt[string] `json:"sip_subdomain,omitzero"`
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
	// Selects which `sip_region` to receive inbound calls from. If null, the default
	// region (US) will be used.
	//
	// Any of "US", "Europe", "Australia".
	SipRegion string `json:"sip_region,omitzero"`
	// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
	// the public Internet) or "Only my connections" (any connection assigned to the
	// same Telnyx user).
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings string `json:"sip_subdomain_receive_settings,omitzero"`
	paramObj
}

func (r IPConnectionNewParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow IPConnectionNewParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPConnectionNewParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IPConnectionNewParamsInbound](
		"ani_number_format", "+E.164", "E.164", "+E.164-national", "E.164-national",
	)
	apijson.RegisterFieldValidator[IPConnectionNewParamsInbound](
		"default_routing_method", "sequential", "round-robin",
	)
	apijson.RegisterFieldValidator[IPConnectionNewParamsInbound](
		"dnis_number_format", "+e164", "e164", "national", "sip_username",
	)
	apijson.RegisterFieldValidator[IPConnectionNewParamsInbound](
		"sip_region", "US", "Europe", "Australia",
	)
	apijson.RegisterFieldValidator[IPConnectionNewParamsInbound](
		"sip_subdomain_receive_settings", "only_my_connections", "from_anyone",
	)
}

// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
// FQDN authentication.
type IPConnectionNewParamsTransportProtocol string

const (
	IPConnectionNewParamsTransportProtocolUdp IPConnectionNewParamsTransportProtocol = "UDP"
	IPConnectionNewParamsTransportProtocolTcp IPConnectionNewParamsTransportProtocol = "TCP"
	IPConnectionNewParamsTransportProtocolTls IPConnectionNewParamsTransportProtocol = "TLS"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type IPConnectionNewParamsWebhookAPIVersion string

const (
	IPConnectionNewParamsWebhookAPIVersion1 IPConnectionNewParamsWebhookAPIVersion = "1"
	IPConnectionNewParamsWebhookAPIVersion2 IPConnectionNewParamsWebhookAPIVersion = "2"
)

type IPConnectionUpdateParams struct {
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
	Active         param.Opt[bool]   `json:"active,omitzero"`
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
	DtmfType     DtmfType                    `json:"dtmf_type,omitzero"`
	Inbound      InboundIPParam              `json:"inbound,omitzero"`
	Outbound     OutboundIPParam             `json:"outbound,omitzero"`
	RtcpSettings ConnectionRtcpSettingsParam `json:"rtcp_settings,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
	// FQDN authentication.
	//
	// Any of "UDP", "TCP", "TLS".
	TransportProtocol IPConnectionUpdateParamsTransportProtocol `json:"transport_protocol,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion IPConnectionUpdateParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r IPConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow IPConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or
// FQDN authentication.
type IPConnectionUpdateParamsTransportProtocol string

const (
	IPConnectionUpdateParamsTransportProtocolUdp IPConnectionUpdateParamsTransportProtocol = "UDP"
	IPConnectionUpdateParamsTransportProtocolTcp IPConnectionUpdateParamsTransportProtocol = "TCP"
	IPConnectionUpdateParamsTransportProtocolTls IPConnectionUpdateParamsTransportProtocol = "TLS"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type IPConnectionUpdateParamsWebhookAPIVersion string

const (
	IPConnectionUpdateParamsWebhookAPIVersion1 IPConnectionUpdateParamsWebhookAPIVersion = "1"
	IPConnectionUpdateParamsWebhookAPIVersion2 IPConnectionUpdateParamsWebhookAPIVersion = "2"
)

type IPConnectionListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
	// filter[outbound.outbound_voice_profile_id]
	Filter IPConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page IPConnectionListParamsPage `query:"page,omitzero" json:"-"`
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
	Sort IPConnectionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPConnectionListParams]'s query parameters as `url.Values`.
func (r IPConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
// filter[outbound.outbound_voice_profile_id]
type IPConnectionListParamsFilter struct {
	// If present, connections with an `fqdn` that equals the given value will be
	// returned. Matching is case-sensitive, and the full string must match.
	Fqdn param.Opt[string] `query:"fqdn,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" format:"int64" json:"-"`
	// Filter by connection_name using nested operations
	ConnectionName IPConnectionListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPConnectionListParamsFilter]'s query parameters as
// `url.Values`.
func (r IPConnectionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by connection_name using nested operations
type IPConnectionListParamsFilterConnectionName struct {
	// If present, connections with <code>connection_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPConnectionListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r IPConnectionListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type IPConnectionListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPConnectionListParamsPage]'s query parameters as
// `url.Values`.
func (r IPConnectionListParamsPage) URLQuery() (v url.Values, err error) {
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
type IPConnectionListParamsSort string

const (
	IPConnectionListParamsSortCreatedAt      IPConnectionListParamsSort = "created_at"
	IPConnectionListParamsSortConnectionName IPConnectionListParamsSort = "connection_name"
	IPConnectionListParamsSortActive         IPConnectionListParamsSort = "active"
)
