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
)

// CallControlApplicationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallControlApplicationService] method instead.
type CallControlApplicationService struct {
	Options []option.RequestOption
}

// NewCallControlApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCallControlApplicationService(opts ...option.RequestOption) (r CallControlApplicationService) {
	r = CallControlApplicationService{}
	r.Options = opts
	return
}

// Create a call control application.
func (r *CallControlApplicationService) New(ctx context.Context, body CallControlApplicationNewParams, opts ...option.RequestOption) (res *CallControlApplicationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "call_control_applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing call control application.
func (r *CallControlApplicationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CallControlApplicationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("call_control_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing call control application.
func (r *CallControlApplicationService) Update(ctx context.Context, id string, body CallControlApplicationUpdateParams, opts ...option.RequestOption) (res *CallControlApplicationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("call_control_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Return a list of call control applications.
func (r *CallControlApplicationService) List(ctx context.Context, query CallControlApplicationListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[CallControlApplication], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "call_control_applications"
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

// Return a list of call control applications.
func (r *CallControlApplicationService) ListAutoPaging(ctx context.Context, query CallControlApplicationListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[CallControlApplication] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a call control application.
func (r *CallControlApplicationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CallControlApplicationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("call_control_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CallControlApplication struct {
	ID string `json:"id"`
	// Specifies whether the connection can be used.
	Active bool `json:"active"`
	// <code>Latency</code> directs Telnyx to route media through the site with the
	// lowest round-trip time to the user's connection. Telnyx calculates this time
	// using ICMP ping messages. This can be disabled by specifying a site to handle
	// all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "London, UK",
	// "Chennai, IN", "Amsterdam, Netherlands", "Toronto, Canada", "Sydney, Australia".
	AnchorsiteOverride CallControlApplicationAnchorsiteOverride `json:"anchorsite_override"`
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name"`
	// Specifies if call cost webhooks should be sent for this Call Control
	// Application.
	CallCostInWebhooks bool `json:"call_cost_in_webhooks"`
	// ISO 8601 formatted date of when the resource was created
	CreatedAt string `json:"created_at"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType CallControlApplicationDtmfType `json:"dtmf_type"`
	// Specifies whether calls to phone numbers associated with this connection should
	// hangup after timing out.
	FirstCommandTimeout bool `json:"first_command_timeout"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs int64                          `json:"first_command_timeout_secs"`
	Inbound                 CallControlApplicationInbound  `json:"inbound"`
	Outbound                CallControlApplicationOutbound `json:"outbound"`
	// Any of "call_control_application".
	RecordType CallControlApplicationRecordType `json:"record_type"`
	// When enabled, DTMF digits entered by users will be redacted in debug logs to
	// protect PII data entered through IVR interactions.
	RedactDtmfDebugLogging bool `json:"redact_dtmf_debug_logging"`
	// Tags assigned to the Call Control Application.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date of when the resource was last updated
	UpdatedAt string `json:"updated_at"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion CallControlApplicationWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as `https`.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url" api:"nullable" format:"url"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as `https`.
	WebhookEventURL    string `json:"webhook_event_url" format:"url"`
	WebhookTimeoutSecs int64  `json:"webhook_timeout_secs" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Active                  respjson.Field
		AnchorsiteOverride      respjson.Field
		ApplicationName         respjson.Field
		CallCostInWebhooks      respjson.Field
		CreatedAt               respjson.Field
		DtmfType                respjson.Field
		FirstCommandTimeout     respjson.Field
		FirstCommandTimeoutSecs respjson.Field
		Inbound                 respjson.Field
		Outbound                respjson.Field
		RecordType              respjson.Field
		RedactDtmfDebugLogging  respjson.Field
		Tags                    respjson.Field
		UpdatedAt               respjson.Field
		WebhookAPIVersion       respjson.Field
		WebhookEventFailoverURL respjson.Field
		WebhookEventURL         respjson.Field
		WebhookTimeoutSecs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplication) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// <code>Latency</code> directs Telnyx to route media through the site with the
// lowest round-trip time to the user's connection. Telnyx calculates this time
// using ICMP ping messages. This can be disabled by specifying a site to handle
// all media.
type CallControlApplicationAnchorsiteOverride string

const (
	CallControlApplicationAnchorsiteOverrideLatency              CallControlApplicationAnchorsiteOverride = "Latency"
	CallControlApplicationAnchorsiteOverrideChicagoIl            CallControlApplicationAnchorsiteOverride = "Chicago, IL"
	CallControlApplicationAnchorsiteOverrideAshburnVa            CallControlApplicationAnchorsiteOverride = "Ashburn, VA"
	CallControlApplicationAnchorsiteOverrideSanJoseCa            CallControlApplicationAnchorsiteOverride = "San Jose, CA"
	CallControlApplicationAnchorsiteOverrideLondonUk             CallControlApplicationAnchorsiteOverride = "London, UK"
	CallControlApplicationAnchorsiteOverrideChennaiIn            CallControlApplicationAnchorsiteOverride = "Chennai, IN"
	CallControlApplicationAnchorsiteOverrideAmsterdamNetherlands CallControlApplicationAnchorsiteOverride = "Amsterdam, Netherlands"
	CallControlApplicationAnchorsiteOverrideTorontoCanada        CallControlApplicationAnchorsiteOverride = "Toronto, Canada"
	CallControlApplicationAnchorsiteOverrideSydneyAustralia      CallControlApplicationAnchorsiteOverride = "Sydney, Australia"
)

// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
// digits sent to Telnyx will be accepted in all formats.
type CallControlApplicationDtmfType string

const (
	CallControlApplicationDtmfTypeRfc2833 CallControlApplicationDtmfType = "RFC 2833"
	CallControlApplicationDtmfTypeInband  CallControlApplicationDtmfType = "Inband"
	CallControlApplicationDtmfTypeSipInfo CallControlApplicationDtmfType = "SIP INFO"
)

type CallControlApplicationRecordType string

const (
	CallControlApplicationRecordTypeCallControlApplication CallControlApplicationRecordType = "call_control_application"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type CallControlApplicationWebhookAPIVersion string

const (
	CallControlApplicationWebhookAPIVersionV1 CallControlApplicationWebhookAPIVersion = "1"
	CallControlApplicationWebhookAPIVersionV2 CallControlApplicationWebhookAPIVersion = "2"
)

type CallControlApplicationInbound struct {
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// When enabled Telnyx will include Shaken/Stir data in the Webhook for new inbound
	// calls.
	ShakenStirEnabled bool `json:"shaken_stir_enabled"`
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
	SipSubdomainReceiveSettings CallControlApplicationInboundSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit                respjson.Field
		ShakenStirEnabled           respjson.Field
		SipSubdomain                respjson.Field
		SipSubdomainReceiveSettings respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplicationInbound) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplicationInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CallControlApplicationInbound to a
// CallControlApplicationInboundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CallControlApplicationInboundParam.Overrides()
func (r CallControlApplicationInbound) ToParam() CallControlApplicationInboundParam {
	return param.Override[CallControlApplicationInboundParam](json.RawMessage(r.RawJSON()))
}

// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
// the public Internet) or "Only my connections" (any connection assigned to the
// same Telnyx user).
type CallControlApplicationInboundSipSubdomainReceiveSettings string

const (
	CallControlApplicationInboundSipSubdomainReceiveSettingsOnlyMyConnections CallControlApplicationInboundSipSubdomainReceiveSettings = "only_my_connections"
	CallControlApplicationInboundSipSubdomainReceiveSettingsFromAnyone        CallControlApplicationInboundSipSubdomainReceiveSettings = "from_anyone"
)

type CallControlApplicationInboundParam struct {
	// When set, this will limit the total number of inbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// When enabled Telnyx will include Shaken/Stir data in the Webhook for new inbound
	// calls.
	ShakenStirEnabled param.Opt[bool] `json:"shaken_stir_enabled,omitzero"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection,
	// in the same way a phone number is used, from a SIP endpoint. Example: the
	// subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using
	// the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any
	// alphanumeric value. Please note TLS encrypted calls are not allowed for
	// subdomain calls.
	SipSubdomain param.Opt[string] `json:"sip_subdomain,omitzero"`
	// This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in
	// the public Internet) or "Only my connections" (any connection assigned to the
	// same Telnyx user).
	//
	// Any of "only_my_connections", "from_anyone".
	SipSubdomainReceiveSettings CallControlApplicationInboundSipSubdomainReceiveSettings `json:"sip_subdomain_receive_settings,omitzero"`
	paramObj
}

func (r CallControlApplicationInboundParam) MarshalJSON() (data []byte, err error) {
	type shadow CallControlApplicationInboundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallControlApplicationInboundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallControlApplicationOutbound struct {
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplicationOutbound) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplicationOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CallControlApplicationOutbound to a
// CallControlApplicationOutboundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CallControlApplicationOutboundParam.Overrides()
func (r CallControlApplicationOutbound) ToParam() CallControlApplicationOutboundParam {
	return param.Override[CallControlApplicationOutboundParam](json.RawMessage(r.RawJSON()))
}

type CallControlApplicationOutboundParam struct {
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r CallControlApplicationOutboundParam) MarshalJSON() (data []byte, err error) {
	type shadow CallControlApplicationOutboundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallControlApplicationOutboundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallControlApplicationNewResponse struct {
	Data CallControlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplicationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplicationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallControlApplicationGetResponse struct {
	Data CallControlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplicationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplicationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallControlApplicationUpdateResponse struct {
	Data CallControlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplicationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplicationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallControlApplicationDeleteResponse struct {
	Data CallControlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallControlApplicationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CallControlApplicationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallControlApplicationNewParams struct {
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name" api:"required"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" api:"required" format:"url"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"url"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Specifies if call cost webhooks should be sent for this Call Control
	// Application.
	CallCostInWebhooks param.Opt[bool] `json:"call_cost_in_webhooks,omitzero"`
	// Specifies whether calls to phone numbers associated with this connection should
	// hangup after timing out.
	FirstCommandTimeout param.Opt[bool] `json:"first_command_timeout,omitzero"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs param.Opt[int64] `json:"first_command_timeout_secs,omitzero"`
	// When enabled, DTMF digits entered by users will be redacted in debug logs to
	// protect PII data entered through IVR interactions.
	RedactDtmfDebugLogging param.Opt[bool] `json:"redact_dtmf_debug_logging,omitzero"`
	// <code>Latency</code> directs Telnyx to route media through the site with the
	// lowest round-trip time to the user's connection. Telnyx calculates this time
	// using ICMP ping messages. This can be disabled by specifying a site to handle
	// all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "London, UK",
	// "Chennai, IN", "Amsterdam, Netherlands", "Toronto, Canada", "Sydney, Australia".
	AnchorsiteOverride CallControlApplicationNewParamsAnchorsiteOverride `json:"anchorsite_override,omitzero"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType CallControlApplicationNewParamsDtmfType `json:"dtmf_type,omitzero"`
	Inbound  CallControlApplicationInboundParam      `json:"inbound,omitzero"`
	Outbound CallControlApplicationOutboundParam     `json:"outbound,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion CallControlApplicationNewParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r CallControlApplicationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CallControlApplicationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallControlApplicationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// <code>Latency</code> directs Telnyx to route media through the site with the
// lowest round-trip time to the user's connection. Telnyx calculates this time
// using ICMP ping messages. This can be disabled by specifying a site to handle
// all media.
type CallControlApplicationNewParamsAnchorsiteOverride string

const (
	CallControlApplicationNewParamsAnchorsiteOverrideLatency              CallControlApplicationNewParamsAnchorsiteOverride = "Latency"
	CallControlApplicationNewParamsAnchorsiteOverrideChicagoIl            CallControlApplicationNewParamsAnchorsiteOverride = "Chicago, IL"
	CallControlApplicationNewParamsAnchorsiteOverrideAshburnVa            CallControlApplicationNewParamsAnchorsiteOverride = "Ashburn, VA"
	CallControlApplicationNewParamsAnchorsiteOverrideSanJoseCa            CallControlApplicationNewParamsAnchorsiteOverride = "San Jose, CA"
	CallControlApplicationNewParamsAnchorsiteOverrideLondonUk             CallControlApplicationNewParamsAnchorsiteOverride = "London, UK"
	CallControlApplicationNewParamsAnchorsiteOverrideChennaiIn            CallControlApplicationNewParamsAnchorsiteOverride = "Chennai, IN"
	CallControlApplicationNewParamsAnchorsiteOverrideAmsterdamNetherlands CallControlApplicationNewParamsAnchorsiteOverride = "Amsterdam, Netherlands"
	CallControlApplicationNewParamsAnchorsiteOverrideTorontoCanada        CallControlApplicationNewParamsAnchorsiteOverride = "Toronto, Canada"
	CallControlApplicationNewParamsAnchorsiteOverrideSydneyAustralia      CallControlApplicationNewParamsAnchorsiteOverride = "Sydney, Australia"
)

// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
// digits sent to Telnyx will be accepted in all formats.
type CallControlApplicationNewParamsDtmfType string

const (
	CallControlApplicationNewParamsDtmfTypeRfc2833 CallControlApplicationNewParamsDtmfType = "RFC 2833"
	CallControlApplicationNewParamsDtmfTypeInband  CallControlApplicationNewParamsDtmfType = "Inband"
	CallControlApplicationNewParamsDtmfTypeSipInfo CallControlApplicationNewParamsDtmfType = "SIP INFO"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type CallControlApplicationNewParamsWebhookAPIVersion string

const (
	CallControlApplicationNewParamsWebhookAPIVersionV1 CallControlApplicationNewParamsWebhookAPIVersion = "1"
	CallControlApplicationNewParamsWebhookAPIVersionV2 CallControlApplicationNewParamsWebhookAPIVersion = "2"
)

type CallControlApplicationUpdateParams struct {
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name" api:"required"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" api:"required" format:"url"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"url"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Specifies if call cost webhooks should be sent for this Call Control
	// Application.
	CallCostInWebhooks param.Opt[bool] `json:"call_cost_in_webhooks,omitzero"`
	// Specifies whether calls to phone numbers associated with this connection should
	// hangup after timing out.
	FirstCommandTimeout param.Opt[bool] `json:"first_command_timeout,omitzero"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs param.Opt[int64] `json:"first_command_timeout_secs,omitzero"`
	// When enabled, DTMF digits entered by users will be redacted in debug logs to
	// protect PII data entered through IVR interactions.
	RedactDtmfDebugLogging param.Opt[bool] `json:"redact_dtmf_debug_logging,omitzero"`
	// <code>Latency</code> directs Telnyx to route media through the site with the
	// lowest round-trip time to the user's connection. Telnyx calculates this time
	// using ICMP ping messages. This can be disabled by specifying a site to handle
	// all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "London, UK",
	// "Chennai, IN", "Amsterdam, Netherlands", "Toronto, Canada", "Sydney, Australia".
	AnchorsiteOverride CallControlApplicationUpdateParamsAnchorsiteOverride `json:"anchorsite_override,omitzero"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType CallControlApplicationUpdateParamsDtmfType `json:"dtmf_type,omitzero"`
	Inbound  CallControlApplicationInboundParam         `json:"inbound,omitzero"`
	Outbound CallControlApplicationOutboundParam        `json:"outbound,omitzero"`
	// Tags assigned to the Call Control Application.
	Tags []string `json:"tags,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion CallControlApplicationUpdateParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r CallControlApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CallControlApplicationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallControlApplicationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// <code>Latency</code> directs Telnyx to route media through the site with the
// lowest round-trip time to the user's connection. Telnyx calculates this time
// using ICMP ping messages. This can be disabled by specifying a site to handle
// all media.
type CallControlApplicationUpdateParamsAnchorsiteOverride string

const (
	CallControlApplicationUpdateParamsAnchorsiteOverrideLatency              CallControlApplicationUpdateParamsAnchorsiteOverride = "Latency"
	CallControlApplicationUpdateParamsAnchorsiteOverrideChicagoIl            CallControlApplicationUpdateParamsAnchorsiteOverride = "Chicago, IL"
	CallControlApplicationUpdateParamsAnchorsiteOverrideAshburnVa            CallControlApplicationUpdateParamsAnchorsiteOverride = "Ashburn, VA"
	CallControlApplicationUpdateParamsAnchorsiteOverrideSanJoseCa            CallControlApplicationUpdateParamsAnchorsiteOverride = "San Jose, CA"
	CallControlApplicationUpdateParamsAnchorsiteOverrideLondonUk             CallControlApplicationUpdateParamsAnchorsiteOverride = "London, UK"
	CallControlApplicationUpdateParamsAnchorsiteOverrideChennaiIn            CallControlApplicationUpdateParamsAnchorsiteOverride = "Chennai, IN"
	CallControlApplicationUpdateParamsAnchorsiteOverrideAmsterdamNetherlands CallControlApplicationUpdateParamsAnchorsiteOverride = "Amsterdam, Netherlands"
	CallControlApplicationUpdateParamsAnchorsiteOverrideTorontoCanada        CallControlApplicationUpdateParamsAnchorsiteOverride = "Toronto, Canada"
	CallControlApplicationUpdateParamsAnchorsiteOverrideSydneyAustralia      CallControlApplicationUpdateParamsAnchorsiteOverride = "Sydney, Australia"
)

// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
// digits sent to Telnyx will be accepted in all formats.
type CallControlApplicationUpdateParamsDtmfType string

const (
	CallControlApplicationUpdateParamsDtmfTypeRfc2833 CallControlApplicationUpdateParamsDtmfType = "RFC 2833"
	CallControlApplicationUpdateParamsDtmfTypeInband  CallControlApplicationUpdateParamsDtmfType = "Inband"
	CallControlApplicationUpdateParamsDtmfTypeSipInfo CallControlApplicationUpdateParamsDtmfType = "SIP INFO"
)

// Determines which webhook format will be used, Telnyx API v1 or v2.
type CallControlApplicationUpdateParamsWebhookAPIVersion string

const (
	CallControlApplicationUpdateParamsWebhookAPIVersionV1 CallControlApplicationUpdateParamsWebhookAPIVersion = "1"
	CallControlApplicationUpdateParamsWebhookAPIVersionV2 CallControlApplicationUpdateParamsWebhookAPIVersion = "2"
)

type CallControlApplicationListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[application_name][contains], filter[outbound.outbound_voice_profile_id],
	// filter[leg_id], filter[application_session_id], filter[connection_id],
	// filter[product], filter[failed], filter[from], filter[to], filter[name],
	// filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]
	Filter CallControlApplicationListParamsFilter `query:"filter,omitzero" json:"-"`
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
	Sort CallControlApplicationListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallControlApplicationListParams]'s query parameters as
// `url.Values`.
func (r CallControlApplicationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[application_name][contains], filter[outbound.outbound_voice_profile_id],
// filter[leg_id], filter[application_session_id], filter[connection_id],
// filter[product], filter[failed], filter[from], filter[to], filter[name],
// filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]
type CallControlApplicationListParamsFilter struct {
	// The unique identifier of the call session. A session may include multiple call
	// leg events.
	ApplicationSessionID param.Opt[string] `query:"application_session_id,omitzero" format:"uuid" json:"-"`
	// The unique identifier of the conection.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Delivery failed or not.
	Failed param.Opt[bool] `query:"failed,omitzero" json:"-"`
	// Filter by From number.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// The unique identifier of an individual call leg.
	LegID param.Opt[string] `query:"leg_id,omitzero" format:"uuid" json:"-"`
	// If present, conferences will be filtered to those with a matching `name`
	// attribute. Matching is case-sensitive
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundOutboundVoiceProfileID param.Opt[string] `query:"outbound.outbound_voice_profile_id,omitzero" json:"-"`
	// Filter by To number.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Application name filters
	ApplicationName CallControlApplicationListParamsFilterApplicationName `query:"application_name,omitzero" json:"-"`
	// Event occurred_at filters
	OccurredAt CallControlApplicationListParamsFilterOccurredAt `query:"occurred_at,omitzero" json:"-"`
	// Filter by product.
	//
	// Any of "call_control", "fax", "texml".
	Product string `query:"product,omitzero" json:"-"`
	// If present, conferences will be filtered by status.
	//
	// Any of "init", "in_progress", "completed".
	Status string `query:"status,omitzero" json:"-"`
	// Event type
	//
	// Any of "command", "webhook".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallControlApplicationListParamsFilter]'s query parameters
// as `url.Values`.
func (r CallControlApplicationListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Application name filters
type CallControlApplicationListParamsFilterApplicationName struct {
	// If present, applications with <code>application_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallControlApplicationListParamsFilterApplicationName]'s
// query parameters as `url.Values`.
func (r CallControlApplicationListParamsFilterApplicationName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Event occurred_at filters
type CallControlApplicationListParamsFilterOccurredAt struct {
	// Event occurred_at: equal
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Event occurred_at: greater than
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Event occurred_at: greater than or equal
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Event occurred_at: lower than
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	// Event occurred_at: lower than or equal
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallControlApplicationListParamsFilterOccurredAt]'s query
// parameters as `url.Values`.
func (r CallControlApplicationListParamsFilterOccurredAt) URLQuery() (v url.Values, err error) {
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
type CallControlApplicationListParamsSort string

const (
	CallControlApplicationListParamsSortCreatedAt      CallControlApplicationListParamsSort = "created_at"
	CallControlApplicationListParamsSortConnectionName CallControlApplicationListParamsSort = "connection_name"
	CallControlApplicationListParamsSortActive         CallControlApplicationListParamsSort = "active"
)
