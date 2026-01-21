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

// CredentialConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialConnectionService] method instead.
type CredentialConnectionService struct {
	Options []option.RequestOption
	Actions CredentialConnectionActionService
}

// NewCredentialConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCredentialConnectionService(opts ...option.RequestOption) (r CredentialConnectionService) {
	r = CredentialConnectionService{}
	r.Options = opts
	r.Actions = NewCredentialConnectionActionService(opts...)
	return
}

// Creates a credential connection.
func (r *CredentialConnectionService) New(ctx context.Context, body CredentialConnectionNewParams, opts ...option.RequestOption) (res *CredentialConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "credential_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing credential connection.
func (r *CredentialConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CredentialConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credential_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing credential connection.
func (r *CredentialConnectionService) Update(ctx context.Context, id string, body CredentialConnectionUpdateParams, opts ...option.RequestOption) (res *CredentialConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credential_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your credential connections.
func (r *CredentialConnectionService) List(ctx context.Context, query CredentialConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[CredentialConnection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credential_connections"
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

// Returns a list of your credential connections.
func (r *CredentialConnectionService) ListAutoPaging(ctx context.Context, query CredentialConnectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[CredentialConnection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing credential connection.
func (r *CredentialConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CredentialConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credential_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// `Latency` directs Telnyx to route media through the site with the lowest
// round-trip time to the user's connection. Telnyx calculates this time using ICMP
// ping messages. This can be disabled by specifying a site to handle all media.
type AnchorsiteOverride string

const (
	AnchorsiteOverrideLatency              AnchorsiteOverride = "Latency"
	AnchorsiteOverrideChicagoIl            AnchorsiteOverride = "Chicago, IL"
	AnchorsiteOverrideAshburnVa            AnchorsiteOverride = "Ashburn, VA"
	AnchorsiteOverrideSanJoseCa            AnchorsiteOverride = "San Jose, CA"
	AnchorsiteOverrideSydneyAustralia      AnchorsiteOverride = "Sydney, Australia"
	AnchorsiteOverrideAmsterdamNetherlands AnchorsiteOverride = "Amsterdam, Netherlands"
	AnchorsiteOverrideLondonUk             AnchorsiteOverride = "London, UK"
	AnchorsiteOverrideTorontoCanada        AnchorsiteOverride = "Toronto, Canada"
	AnchorsiteOverrideVancouverCanada      AnchorsiteOverride = "Vancouver, Canada"
	AnchorsiteOverrideFrankfurtGermany     AnchorsiteOverride = "Frankfurt, Germany"
)

type ConnectionRtcpSettings struct {
	// BETA - Enable the capture and storage of RTCP messages to create QoS reports on
	// the Telnyx Mission Control Portal.
	CaptureEnabled bool `json:"capture_enabled"`
	// RTCP port by default is rtp+1, it can also be set to rtcp-mux
	//
	// Any of "rtcp-mux", "rtp+1".
	Port ConnectionRtcpSettingsPort `json:"port"`
	// RTCP reports are sent to customers based on the frequency set. Frequency is in
	// seconds and it can be set to values from 5 to 3000 seconds.
	ReportFrequencySecs int64 `json:"report_frequency_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CaptureEnabled      respjson.Field
		Port                respjson.Field
		ReportFrequencySecs respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionRtcpSettings) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRtcpSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConnectionRtcpSettings to a ConnectionRtcpSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConnectionRtcpSettingsParam.Overrides()
func (r ConnectionRtcpSettings) ToParam() ConnectionRtcpSettingsParam {
	return param.Override[ConnectionRtcpSettingsParam](json.RawMessage(r.RawJSON()))
}

// RTCP port by default is rtp+1, it can also be set to rtcp-mux
type ConnectionRtcpSettingsPort string

const (
	ConnectionRtcpSettingsPortRtcpMux ConnectionRtcpSettingsPort = "rtcp-mux"
	ConnectionRtcpSettingsPortRtp1    ConnectionRtcpSettingsPort = "rtp+1"
)

type ConnectionRtcpSettingsParam struct {
	// BETA - Enable the capture and storage of RTCP messages to create QoS reports on
	// the Telnyx Mission Control Portal.
	CaptureEnabled param.Opt[bool] `json:"capture_enabled,omitzero"`
	// RTCP reports are sent to customers based on the frequency set. Frequency is in
	// seconds and it can be set to values from 5 to 3000 seconds.
	ReportFrequencySecs param.Opt[int64] `json:"report_frequency_secs,omitzero"`
	// RTCP port by default is rtp+1, it can also be set to rtcp-mux
	//
	// Any of "rtcp-mux", "rtp+1".
	Port ConnectionRtcpSettingsPort `json:"port,omitzero"`
	paramObj
}

func (r ConnectionRtcpSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectionRtcpSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectionRtcpSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialConnection struct {
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
	EncryptedMedia EncryptedMedia    `json:"encrypted_media,nullable"`
	Inbound        CredentialInbound `json:"inbound"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression CredentialConnectionNoiseSuppression `json:"noise_suppression"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetails `json:"noise_suppression_details"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly
	// if both are on the Telnyx network. If this is disabled, Telnyx will be able to
	// use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled bool               `json:"onnet_t38_passthrough_enabled"`
	Outbound                   CredentialOutbound `json:"outbound"`
	// The password to be used as part of the credentials. Must be 8 to 128 characters
	// long.
	Password string `json:"password"`
	// Identifies the type of the resource.
	RecordType   string                 `json:"record_type"`
	RtcpSettings ConnectionRtcpSettings `json:"rtcp_settings"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection.
	// If enabled for all (unrestricted) then anyone who calls the SIP URI
	// <your-username>@telnyx.com will be connected to your Connection. You can also
	// choose to allow only calls that are originated on any Connections under your
	// account (internal).
	//
	// Any of "disabled", "unrestricted", "internal".
	SipUriCallingPreference CredentialConnectionSipUriCallingPreference `json:"sip_uri_calling_preference"`
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
	WebhookAPIVersion CredentialConnectionWebhookAPIVersion `json:"webhook_api_version"`
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
		CallCostInWebhooks               respjson.Field
		ConnectionName                   respjson.Field
		CreatedAt                        respjson.Field
		DefaultOnHoldComfortNoiseEnabled respjson.Field
		DtmfType                         respjson.Field
		EncodeContactHeaderEnabled       respjson.Field
		EncryptedMedia                   respjson.Field
		Inbound                          respjson.Field
		NoiseSuppression                 respjson.Field
		NoiseSuppressionDetails          respjson.Field
		OnnetT38PassthroughEnabled       respjson.Field
		Outbound                         respjson.Field
		Password                         respjson.Field
		RecordType                       respjson.Field
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
func (r CredentialConnection) RawJSON() string { return r.JSON.raw }
func (r *CredentialConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type CredentialConnectionNoiseSuppression string

const (
	CredentialConnectionNoiseSuppressionInbound  CredentialConnectionNoiseSuppression = "inbound"
	CredentialConnectionNoiseSuppressionOutbound CredentialConnectionNoiseSuppression = "outbound"
	CredentialConnectionNoiseSuppressionBoth     CredentialConnectionNoiseSuppression = "both"
	CredentialConnectionNoiseSuppressionDisabled CredentialConnectionNoiseSuppression = "disabled"
)

// This feature enables inbound SIP URI calls to your Credential Auth Connection.
// If enabled for all (unrestricted) then anyone who calls the SIP URI
// <your-username>@telnyx.com will be connected to your Connection. You can also
// choose to allow only calls that are originated on any Connections under your
// account (internal).
type CredentialConnectionSipUriCallingPreference string

const (
	CredentialConnectionSipUriCallingPreferenceDisabled     CredentialConnectionSipUriCallingPreference = "disabled"
	CredentialConnectionSipUriCallingPreferenceUnrestricted CredentialConnectionSipUriCallingPreference = "unrestricted"
	CredentialConnectionSipUriCallingPreferenceInternal     CredentialConnectionSipUriCallingPreference = "internal"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type CredentialConnectionWebhookAPIVersion string

const (
	CredentialConnectionWebhookAPIVersionV1 CredentialConnectionWebhookAPIVersion = "1"
	CredentialConnectionWebhookAPIVersionV2 CredentialConnectionWebhookAPIVersion = "2"
)

type CredentialInbound struct {
	// This setting allows you to set the format with which the caller's number (ANI)
	// is sent for inbound phone calls.
	//
	// Any of "+E.164", "E.164", "+E.164-national", "E.164-national".
	AniNumberFormat CredentialInboundAniNumberFormat `json:"ani_number_format"`
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat CredentialInboundDnisNumberFormat `json:"dnis_number_format"`
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
	SimultaneousRinging CredentialInboundSimultaneousRinging `json:"simultaneous_ringing"`
	// Defaults to true.
	SipCompactHeadersEnabled bool `json:"sip_compact_headers_enabled"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs int64 `json:"timeout_1xx_secs"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs int64 `json:"timeout_2xx_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AniNumberFormat          respjson.Field
		ChannelLimit             respjson.Field
		Codecs                   respjson.Field
		DnisNumberFormat         respjson.Field
		GenerateRingbackTone     respjson.Field
		IsupHeadersEnabled       respjson.Field
		PrackEnabled             respjson.Field
		ShakenStirEnabled        respjson.Field
		SimultaneousRinging      respjson.Field
		SipCompactHeadersEnabled respjson.Field
		Timeout1xxSecs           respjson.Field
		Timeout2xxSecs           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialInbound) RawJSON() string { return r.JSON.raw }
func (r *CredentialInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CredentialInbound to a CredentialInboundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CredentialInboundParam.Overrides()
func (r CredentialInbound) ToParam() CredentialInboundParam {
	return param.Override[CredentialInboundParam](json.RawMessage(r.RawJSON()))
}

// This setting allows you to set the format with which the caller's number (ANI)
// is sent for inbound phone calls.
type CredentialInboundAniNumberFormat string

const (
	CredentialInboundAniNumberFormatPlusE164         CredentialInboundAniNumberFormat = "+E.164"
	CredentialInboundAniNumberFormatE164             CredentialInboundAniNumberFormat = "E.164"
	CredentialInboundAniNumberFormatPlusE164National CredentialInboundAniNumberFormat = "+E.164-national"
	CredentialInboundAniNumberFormatE164National     CredentialInboundAniNumberFormat = "E.164-national"
)

type CredentialInboundDnisNumberFormat string

const (
	CredentialInboundDnisNumberFormatPlusE164    CredentialInboundDnisNumberFormat = "+e164"
	CredentialInboundDnisNumberFormatE164        CredentialInboundDnisNumberFormat = "e164"
	CredentialInboundDnisNumberFormatNational    CredentialInboundDnisNumberFormat = "national"
	CredentialInboundDnisNumberFormatSipUsername CredentialInboundDnisNumberFormat = "sip_username"
)

// When enabled, allows multiple devices to ring simultaneously on incoming calls.
type CredentialInboundSimultaneousRinging string

const (
	CredentialInboundSimultaneousRingingDisabled CredentialInboundSimultaneousRinging = "disabled"
	CredentialInboundSimultaneousRingingEnabled  CredentialInboundSimultaneousRinging = "enabled"
)

type CredentialInboundParam struct {
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
	AniNumberFormat CredentialInboundAniNumberFormat `json:"ani_number_format,omitzero"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific
	// number on your portal account, in priority order. This only works when the
	// Connection the number is assigned to uses Media Handling mode: default. OPUS and
	// H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs,omitzero"`
	// Any of "+e164", "e164", "national", "sip_username".
	DnisNumberFormat CredentialInboundDnisNumberFormat `json:"dnis_number_format,omitzero"`
	// When enabled, allows multiple devices to ring simultaneously on incoming calls.
	//
	// Any of "disabled", "enabled".
	SimultaneousRinging CredentialInboundSimultaneousRinging `json:"simultaneous_ringing,omitzero"`
	paramObj
}

func (r CredentialInboundParam) MarshalJSON() (data []byte, err error) {
	type shadow CredentialInboundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialInboundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOutbound struct {
	// Set a phone number as the ani_override value to override caller id number on
	// outbound calls.
	AniOverride string `json:"ani_override"`
	// Specifies when we apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType CredentialOutboundAniOverrideType `json:"ani_override_type"`
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
	T38ReinviteSource CredentialOutboundT38ReinviteSource `json:"t38_reinvite_source"`
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
func (r CredentialOutbound) RawJSON() string { return r.JSON.raw }
func (r *CredentialOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CredentialOutbound to a CredentialOutboundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CredentialOutboundParam.Overrides()
func (r CredentialOutbound) ToParam() CredentialOutboundParam {
	return param.Override[CredentialOutboundParam](json.RawMessage(r.RawJSON()))
}

// Specifies when we apply your ani_override setting. Only applies when
// ani_override is not blank.
type CredentialOutboundAniOverrideType string

const (
	CredentialOutboundAniOverrideTypeAlways    CredentialOutboundAniOverrideType = "always"
	CredentialOutboundAniOverrideTypeNormal    CredentialOutboundAniOverrideType = "normal"
	CredentialOutboundAniOverrideTypeEmergency CredentialOutboundAniOverrideType = "emergency"
)

// This setting only affects connections with Fax-type Outbound Voice Profiles. The
// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
// default, Telnyx will send the re-invite. If set to `customer`, the caller is
// expected to send the t.38 reinvite.
type CredentialOutboundT38ReinviteSource string

const (
	CredentialOutboundT38ReinviteSourceTelnyx         CredentialOutboundT38ReinviteSource = "telnyx"
	CredentialOutboundT38ReinviteSourceCustomer       CredentialOutboundT38ReinviteSource = "customer"
	CredentialOutboundT38ReinviteSourceDisabled       CredentialOutboundT38ReinviteSource = "disabled"
	CredentialOutboundT38ReinviteSourcePassthru       CredentialOutboundT38ReinviteSource = "passthru"
	CredentialOutboundT38ReinviteSourceCallerPassthru CredentialOutboundT38ReinviteSource = "caller-passthru"
	CredentialOutboundT38ReinviteSourceCalleePassthru CredentialOutboundT38ReinviteSource = "callee-passthru"
)

type CredentialOutboundParam struct {
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
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero" format:"int64"`
	// Specifies when we apply your ani_override setting. Only applies when
	// ani_override is not blank.
	//
	// Any of "always", "normal", "emergency".
	AniOverrideType CredentialOutboundAniOverrideType `json:"ani_override_type,omitzero"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The
	// setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By
	// default, Telnyx will send the re-invite. If set to `customer`, the caller is
	// expected to send the t.38 reinvite.
	//
	// Any of "telnyx", "customer", "disabled", "passthru", "caller-passthru",
	// "callee-passthru".
	T38ReinviteSource CredentialOutboundT38ReinviteSource `json:"t38_reinvite_source,omitzero"`
	paramObj
}

func (r CredentialOutboundParam) MarshalJSON() (data []byte, err error) {
	type shadow CredentialOutboundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialOutboundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
// digits sent to Telnyx will be accepted in all formats.
type DtmfType string

const (
	DtmfTypeRfc2833 DtmfType = "RFC 2833"
	DtmfTypeInband  DtmfType = "Inband"
	DtmfTypeSipInfo DtmfType = "SIP INFO"
)

// Enable use of SRTP for encryption. Cannot be set if the transport_portocol is
// TLS.
type EncryptedMedia string

const (
	EncryptedMediaSrtp EncryptedMedia = "SRTP"
)

type CredentialConnectionNewResponse struct {
	Data CredentialConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialConnectionGetResponse struct {
	Data CredentialConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialConnectionUpdateResponse struct {
	Data CredentialConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialConnectionDeleteResponse struct {
	Data CredentialConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialConnectionNewParams struct {
	// A user-assigned name to help manage the connection.
	ConnectionName string `json:"connection_name,required"`
	// The password to be used as part of the credentials. Must be 8 to 128 characters
	// long.
	Password string `json:"password,required"`
	// The user name to be used as part of the credentials. Must be 4-32 characters
	// long and alphanumeric values only (no spaces or special characters).
	UserName string `json:"user_name,required"`
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
	DtmfType DtmfType               `json:"dtmf_type,omitzero"`
	Inbound  CredentialInboundParam `json:"inbound,omitzero"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression CredentialConnectionNewParamsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetailsParam `json:"noise_suppression_details,omitzero"`
	Outbound                CredentialOutboundParam                       `json:"outbound,omitzero"`
	RtcpSettings            ConnectionRtcpSettingsParam                   `json:"rtcp_settings,omitzero"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection.
	// If enabled for all (unrestricted) then anyone who calls the SIP URI
	// <your-username>@telnyx.com will be connected to your Connection. You can also
	// choose to allow only calls that are originated on any Connections under your
	// account (internal).
	//
	// Any of "disabled", "unrestricted", "internal".
	SipUriCallingPreference CredentialConnectionNewParamsSipUriCallingPreference `json:"sip_uri_calling_preference,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1, v2 or texml. Note -
	// texml can only be set when the outbound object parameter call_parking_enabled is
	// included and set to true.
	//
	// Any of "1", "2", "texml".
	WebhookAPIVersion CredentialConnectionNewParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r CredentialConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type CredentialConnectionNewParamsNoiseSuppression string

const (
	CredentialConnectionNewParamsNoiseSuppressionInbound  CredentialConnectionNewParamsNoiseSuppression = "inbound"
	CredentialConnectionNewParamsNoiseSuppressionOutbound CredentialConnectionNewParamsNoiseSuppression = "outbound"
	CredentialConnectionNewParamsNoiseSuppressionBoth     CredentialConnectionNewParamsNoiseSuppression = "both"
	CredentialConnectionNewParamsNoiseSuppressionDisabled CredentialConnectionNewParamsNoiseSuppression = "disabled"
)

// This feature enables inbound SIP URI calls to your Credential Auth Connection.
// If enabled for all (unrestricted) then anyone who calls the SIP URI
// <your-username>@telnyx.com will be connected to your Connection. You can also
// choose to allow only calls that are originated on any Connections under your
// account (internal).
type CredentialConnectionNewParamsSipUriCallingPreference string

const (
	CredentialConnectionNewParamsSipUriCallingPreferenceDisabled     CredentialConnectionNewParamsSipUriCallingPreference = "disabled"
	CredentialConnectionNewParamsSipUriCallingPreferenceUnrestricted CredentialConnectionNewParamsSipUriCallingPreference = "unrestricted"
	CredentialConnectionNewParamsSipUriCallingPreferenceInternal     CredentialConnectionNewParamsSipUriCallingPreference = "internal"
)

// Determines which webhook format will be used, Telnyx API v1, v2 or texml. Note -
// texml can only be set when the outbound object parameter call_parking_enabled is
// included and set to true.
type CredentialConnectionNewParamsWebhookAPIVersion string

const (
	CredentialConnectionNewParamsWebhookAPIVersionV1    CredentialConnectionNewParamsWebhookAPIVersion = "1"
	CredentialConnectionNewParamsWebhookAPIVersionV2    CredentialConnectionNewParamsWebhookAPIVersion = "2"
	CredentialConnectionNewParamsWebhookAPIVersionTexml CredentialConnectionNewParamsWebhookAPIVersion = "texml"
)

type CredentialConnectionUpdateParams struct {
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
	DtmfType DtmfType               `json:"dtmf_type,omitzero"`
	Inbound  CredentialInboundParam `json:"inbound,omitzero"`
	// Controls when noise suppression is applied to calls. When set to 'inbound',
	// noise suppression is applied to incoming audio. When set to 'outbound', it's
	// applied to outgoing audio. When set to 'both', it's applied in both directions.
	// When set to 'disabled', noise suppression is turned off.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression CredentialConnectionUpdateParamsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration options for noise suppression. These settings are stored
	// regardless of the noise_suppression value, but only take effect when
	// noise_suppression is not 'disabled'. If you disable noise suppression and later
	// re-enable it, the previously configured settings will be used.
	NoiseSuppressionDetails shared.ConnectionNoiseSuppressionDetailsParam `json:"noise_suppression_details,omitzero"`
	Outbound                CredentialOutboundParam                       `json:"outbound,omitzero"`
	RtcpSettings            ConnectionRtcpSettingsParam                   `json:"rtcp_settings,omitzero"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection.
	// If enabled for all (unrestricted) then anyone who calls the SIP URI
	// <your-username>@telnyx.com will be connected to your Connection. You can also
	// choose to allow only calls that are originated on any Connections under your
	// account (internal).
	//
	// Any of "disabled", "unrestricted", "internal".
	SipUriCallingPreference CredentialConnectionUpdateParamsSipUriCallingPreference `json:"sip_uri_calling_preference,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion CredentialConnectionUpdateParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r CredentialConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when noise suppression is applied to calls. When set to 'inbound',
// noise suppression is applied to incoming audio. When set to 'outbound', it's
// applied to outgoing audio. When set to 'both', it's applied in both directions.
// When set to 'disabled', noise suppression is turned off.
type CredentialConnectionUpdateParamsNoiseSuppression string

const (
	CredentialConnectionUpdateParamsNoiseSuppressionInbound  CredentialConnectionUpdateParamsNoiseSuppression = "inbound"
	CredentialConnectionUpdateParamsNoiseSuppressionOutbound CredentialConnectionUpdateParamsNoiseSuppression = "outbound"
	CredentialConnectionUpdateParamsNoiseSuppressionBoth     CredentialConnectionUpdateParamsNoiseSuppression = "both"
	CredentialConnectionUpdateParamsNoiseSuppressionDisabled CredentialConnectionUpdateParamsNoiseSuppression = "disabled"
)

// This feature enables inbound SIP URI calls to your Credential Auth Connection.
// If enabled for all (unrestricted) then anyone who calls the SIP URI
// <your-username>@telnyx.com will be connected to your Connection. You can also
// choose to allow only calls that are originated on any Connections under your
// account (internal).
type CredentialConnectionUpdateParamsSipUriCallingPreference string

const (
	CredentialConnectionUpdateParamsSipUriCallingPreferenceDisabled     CredentialConnectionUpdateParamsSipUriCallingPreference = "disabled"
	CredentialConnectionUpdateParamsSipUriCallingPreferenceUnrestricted CredentialConnectionUpdateParamsSipUriCallingPreference = "unrestricted"
	CredentialConnectionUpdateParamsSipUriCallingPreferenceInternal     CredentialConnectionUpdateParamsSipUriCallingPreference = "internal"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type CredentialConnectionUpdateParamsWebhookAPIVersion string

const (
	CredentialConnectionUpdateParamsWebhookAPIVersionV1 CredentialConnectionUpdateParamsWebhookAPIVersion = "1"
	CredentialConnectionUpdateParamsWebhookAPIVersionV2 CredentialConnectionUpdateParamsWebhookAPIVersion = "2"
)

type CredentialConnectionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
	// filter[outbound.outbound_voice_profile_id]
	Filter CredentialConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
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
	Sort CredentialConnectionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CredentialConnectionListParams]'s query parameters as
// `url.Values`.
func (r CredentialConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
// filter[outbound.outbound_voice_profile_id]
type CredentialConnectionListParamsFilter struct {
	// If present, connections with an `fqdn` that equals the given value will be
	// returned. Matching is case-sensitive, and the full string must match.
	Fqdn param.Opt[string] `query:"fqdn,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" format:"int64" json:"-"`
	// Filter by connection_name using nested operations
	ConnectionName CredentialConnectionListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CredentialConnectionListParamsFilter]'s query parameters as
// `url.Values`.
func (r CredentialConnectionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by connection_name using nested operations
type CredentialConnectionListParamsFilterConnectionName struct {
	// If present, connections with <code>connection_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CredentialConnectionListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r CredentialConnectionListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
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
type CredentialConnectionListParamsSort string

const (
	CredentialConnectionListParamsSortCreatedAt      CredentialConnectionListParamsSort = "created_at"
	CredentialConnectionListParamsSortConnectionName CredentialConnectionListParamsSort = "connection_name"
	CredentialConnectionListParamsSortActive         CredentialConnectionListParamsSort = "active"
)
