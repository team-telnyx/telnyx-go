// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// TexmlApplicationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlApplicationService] method instead.
type TexmlApplicationService struct {
	Options []option.RequestOption
}

// NewTexmlApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTexmlApplicationService(opts ...option.RequestOption) (r TexmlApplicationService) {
	r = TexmlApplicationService{}
	r.Options = opts
	return
}

// Creates a TeXML Application.
func (r *TexmlApplicationService) New(ctx context.Context, body TexmlApplicationNewParams, opts ...option.RequestOption) (res *TexmlApplicationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "texml_applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing TeXML Application.
func (r *TexmlApplicationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TexmlApplicationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("texml_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing TeXML Application.
func (r *TexmlApplicationService) Update(ctx context.Context, id string, body TexmlApplicationUpdateParams, opts ...option.RequestOption) (res *TexmlApplicationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("texml_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your TeXML Applications.
func (r *TexmlApplicationService) List(ctx context.Context, query TexmlApplicationListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[TexmlApplication], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "texml_applications"
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

// Returns a list of your TeXML Applications.
func (r *TexmlApplicationService) ListAutoPaging(ctx context.Context, query TexmlApplicationListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[TexmlApplication] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a TeXML Application.
func (r *TexmlApplicationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TexmlApplicationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("texml_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TexmlApplication struct {
	// Uniquely identifies the resource.
	ID string `json:"id"`
	// Specifies whether the connection can be used.
	Active bool `json:"active"`
	// `Latency` directs Telnyx to route media through the site with the lowest
	// round-trip time to the user's connection. Telnyx calculates this time using ICMP
	// ping messages. This can be disabled by specifying a site to handle all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "Sydney,
	// Australia", "Amsterdam, Netherlands", "London, UK", "Toronto, Canada",
	// "Vancouver, Canada", "Frankfurt, Germany".
	AnchorsiteOverride AnchorsiteOverride `json:"anchorsite_override"`
	// Specifies if call cost webhooks should be sent for this TeXML Application.
	CallCostInWebhooks bool `json:"call_cost_in_webhooks"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF
	// digits sent to Telnyx will be accepted in all formats.
	//
	// Any of "RFC 2833", "Inband", "SIP INFO".
	DtmfType DtmfType `json:"dtmf_type"`
	// Specifies whether calls to phone numbers associated with this connection should
	// hangup after timing out.
	FirstCommandTimeout bool `json:"first_command_timeout"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs int64 `json:"first_command_timeout_secs"`
	// A user-assigned name to help manage the application.
	FriendlyName string                   `json:"friendly_name"`
	Inbound      TexmlApplicationInbound  `json:"inbound"`
	Outbound     TexmlApplicationOutbound `json:"outbound"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// URL for Telnyx to send requests to containing information about call progress
	// events.
	StatusCallback string `json:"status_callback" format:"uri"`
	// HTTP request method Telnyx should use when requesting the status_callback URL.
	//
	// Any of "get", "post".
	StatusCallbackMethod TexmlApplicationStatusCallbackMethod `json:"status_callback_method"`
	// Tags associated with the Texml Application.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// URL to which Telnyx will deliver your XML Translator webhooks if we get an error
	// response from your voice_url.
	VoiceFallbackURL string `json:"voice_fallback_url" format:"uri"`
	// HTTP request method Telnyx will use to interact with your XML Translator
	// webhooks. Either 'get' or 'post'.
	//
	// Any of "get", "post".
	VoiceMethod TexmlApplicationVoiceMethod `json:"voice_method"`
	// URL to which Telnyx will deliver your XML Translator webhooks.
	VoiceURL string `json:"voice_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Active                  respjson.Field
		AnchorsiteOverride      respjson.Field
		CallCostInWebhooks      respjson.Field
		CreatedAt               respjson.Field
		DtmfType                respjson.Field
		FirstCommandTimeout     respjson.Field
		FirstCommandTimeoutSecs respjson.Field
		FriendlyName            respjson.Field
		Inbound                 respjson.Field
		Outbound                respjson.Field
		RecordType              respjson.Field
		StatusCallback          respjson.Field
		StatusCallbackMethod    respjson.Field
		Tags                    respjson.Field
		UpdatedAt               respjson.Field
		VoiceFallbackURL        respjson.Field
		VoiceMethod             respjson.Field
		VoiceURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlApplication) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationInbound struct {
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
	SipSubdomainReceiveSettings string `json:"sip_subdomain_receive_settings"`
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
func (r TexmlApplicationInbound) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplicationInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationOutbound struct {
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
func (r TexmlApplicationOutbound) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplicationOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request method Telnyx should use when requesting the status_callback URL.
type TexmlApplicationStatusCallbackMethod string

const (
	TexmlApplicationStatusCallbackMethodGet  TexmlApplicationStatusCallbackMethod = "get"
	TexmlApplicationStatusCallbackMethodPost TexmlApplicationStatusCallbackMethod = "post"
)

// HTTP request method Telnyx will use to interact with your XML Translator
// webhooks. Either 'get' or 'post'.
type TexmlApplicationVoiceMethod string

const (
	TexmlApplicationVoiceMethodGet  TexmlApplicationVoiceMethod = "get"
	TexmlApplicationVoiceMethodPost TexmlApplicationVoiceMethod = "post"
)

type TexmlApplicationNewResponse struct {
	Data TexmlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlApplicationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplicationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationGetResponse struct {
	Data TexmlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlApplicationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplicationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationUpdateResponse struct {
	Data TexmlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlApplicationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplicationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationDeleteResponse struct {
	Data TexmlApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlApplicationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlApplicationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationNewParams struct {
	// A user-assigned name to help manage the application.
	FriendlyName string `json:"friendly_name,required"`
	// URL to which Telnyx will deliver your XML Translator webhooks.
	VoiceURL string `json:"voice_url,required" format:"uri"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Specifies if call cost webhooks should be sent for this TeXML Application.
	CallCostInWebhooks param.Opt[bool] `json:"call_cost_in_webhooks,omitzero"`
	// Specifies whether calls to phone numbers associated with this connection should
	// hangup after timing out.
	FirstCommandTimeout param.Opt[bool] `json:"first_command_timeout,omitzero"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs param.Opt[int64] `json:"first_command_timeout_secs,omitzero"`
	// URL for Telnyx to send requests to containing information about call progress
	// events.
	StatusCallback param.Opt[string] `json:"status_callback,omitzero" format:"uri"`
	// URL to which Telnyx will deliver your XML Translator webhooks if we get an error
	// response from your voice_url.
	VoiceFallbackURL param.Opt[string] `json:"voice_fallback_url,omitzero" format:"uri"`
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
	DtmfType DtmfType                          `json:"dtmf_type,omitzero"`
	Inbound  TexmlApplicationNewParamsInbound  `json:"inbound,omitzero"`
	Outbound TexmlApplicationNewParamsOutbound `json:"outbound,omitzero"`
	// HTTP request method Telnyx should use when requesting the status_callback URL.
	//
	// Any of "get", "post".
	StatusCallbackMethod TexmlApplicationNewParamsStatusCallbackMethod `json:"status_callback_method,omitzero"`
	// Tags associated with the Texml Application.
	Tags []string `json:"tags,omitzero"`
	// HTTP request method Telnyx will use to interact with your XML Translator
	// webhooks. Either 'get' or 'post'.
	//
	// Any of "get", "post".
	VoiceMethod TexmlApplicationNewParamsVoiceMethod `json:"voice_method,omitzero"`
	paramObj
}

func (r TexmlApplicationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlApplicationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlApplicationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationNewParamsInbound struct {
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
	SipSubdomainReceiveSettings string `json:"sip_subdomain_receive_settings,omitzero"`
	paramObj
}

func (r TexmlApplicationNewParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow TexmlApplicationNewParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlApplicationNewParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TexmlApplicationNewParamsInbound](
		"sip_subdomain_receive_settings", "only_my_connections", "from_anyone",
	)
}

type TexmlApplicationNewParamsOutbound struct {
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r TexmlApplicationNewParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow TexmlApplicationNewParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlApplicationNewParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request method Telnyx should use when requesting the status_callback URL.
type TexmlApplicationNewParamsStatusCallbackMethod string

const (
	TexmlApplicationNewParamsStatusCallbackMethodGet  TexmlApplicationNewParamsStatusCallbackMethod = "get"
	TexmlApplicationNewParamsStatusCallbackMethodPost TexmlApplicationNewParamsStatusCallbackMethod = "post"
)

// HTTP request method Telnyx will use to interact with your XML Translator
// webhooks. Either 'get' or 'post'.
type TexmlApplicationNewParamsVoiceMethod string

const (
	TexmlApplicationNewParamsVoiceMethodGet  TexmlApplicationNewParamsVoiceMethod = "get"
	TexmlApplicationNewParamsVoiceMethodPost TexmlApplicationNewParamsVoiceMethod = "post"
)

type TexmlApplicationUpdateParams struct {
	// A user-assigned name to help manage the application.
	FriendlyName string `json:"friendly_name,required"`
	// URL to which Telnyx will deliver your XML Translator webhooks.
	VoiceURL string `json:"voice_url,required" format:"uri"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Specifies if call cost webhooks should be sent for this TeXML Application.
	CallCostInWebhooks param.Opt[bool] `json:"call_cost_in_webhooks,omitzero"`
	// Specifies whether calls to phone numbers associated with this connection should
	// hangup after timing out.
	FirstCommandTimeout param.Opt[bool] `json:"first_command_timeout,omitzero"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs param.Opt[int64] `json:"first_command_timeout_secs,omitzero"`
	// URL for Telnyx to send requests to containing information about call progress
	// events.
	StatusCallback param.Opt[string] `json:"status_callback,omitzero" format:"uri"`
	// URL to which Telnyx will deliver your XML Translator webhooks if we get an error
	// response from your voice_url.
	VoiceFallbackURL param.Opt[string] `json:"voice_fallback_url,omitzero" format:"uri"`
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
	DtmfType DtmfType                             `json:"dtmf_type,omitzero"`
	Inbound  TexmlApplicationUpdateParamsInbound  `json:"inbound,omitzero"`
	Outbound TexmlApplicationUpdateParamsOutbound `json:"outbound,omitzero"`
	// HTTP request method Telnyx should use when requesting the status_callback URL.
	//
	// Any of "get", "post".
	StatusCallbackMethod TexmlApplicationUpdateParamsStatusCallbackMethod `json:"status_callback_method,omitzero"`
	// Tags associated with the Texml Application.
	Tags []string `json:"tags,omitzero"`
	// HTTP request method Telnyx will use to interact with your XML Translator
	// webhooks. Either 'get' or 'post'.
	//
	// Any of "get", "post".
	VoiceMethod TexmlApplicationUpdateParamsVoiceMethod `json:"voice_method,omitzero"`
	paramObj
}

func (r TexmlApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlApplicationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlApplicationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlApplicationUpdateParamsInbound struct {
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
	SipSubdomainReceiveSettings string `json:"sip_subdomain_receive_settings,omitzero"`
	paramObj
}

func (r TexmlApplicationUpdateParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow TexmlApplicationUpdateParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlApplicationUpdateParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TexmlApplicationUpdateParamsInbound](
		"sip_subdomain_receive_settings", "only_my_connections", "from_anyone",
	)
}

type TexmlApplicationUpdateParamsOutbound struct {
	// When set, this will limit the total number of outbound calls to phone numbers
	// associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r TexmlApplicationUpdateParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow TexmlApplicationUpdateParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlApplicationUpdateParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request method Telnyx should use when requesting the status_callback URL.
type TexmlApplicationUpdateParamsStatusCallbackMethod string

const (
	TexmlApplicationUpdateParamsStatusCallbackMethodGet  TexmlApplicationUpdateParamsStatusCallbackMethod = "get"
	TexmlApplicationUpdateParamsStatusCallbackMethodPost TexmlApplicationUpdateParamsStatusCallbackMethod = "post"
)

// HTTP request method Telnyx will use to interact with your XML Translator
// webhooks. Either 'get' or 'post'.
type TexmlApplicationUpdateParamsVoiceMethod string

const (
	TexmlApplicationUpdateParamsVoiceMethodGet  TexmlApplicationUpdateParamsVoiceMethod = "get"
	TexmlApplicationUpdateParamsVoiceMethodPost TexmlApplicationUpdateParamsVoiceMethod = "post"
)

type TexmlApplicationListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[outbound_voice_profile_id], filter[friendly_name]
	Filter TexmlApplicationListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page TexmlApplicationListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code> -</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>friendly_name</code>: sorts the result by the
	//	  <code>friendly_name</code> field in ascending order.
	//	</li>
	//
	//	<li>
	//	  <code>-friendly_name</code>: sorts the result by the
	//	  <code>friendly_name</code> field in descending order.
	//	</li>
	//
	// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
	//
	// Any of "created_at", "friendly_name", "active".
	Sort TexmlApplicationListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlApplicationListParams]'s query parameters as
// `url.Values`.
func (r TexmlApplicationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[outbound_voice_profile_id], filter[friendly_name]
type TexmlApplicationListParamsFilter struct {
	// If present, applications with <code>friendly_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	FriendlyName param.Opt[string] `query:"friendly_name,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlApplicationListParamsFilter]'s query parameters as
// `url.Values`.
func (r TexmlApplicationListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type TexmlApplicationListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlApplicationListParamsPage]'s query parameters as
// `url.Values`.
func (r TexmlApplicationListParamsPage) URLQuery() (v url.Values, err error) {
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
//	  <code>friendly_name</code>: sorts the result by the
//	  <code>friendly_name</code> field in ascending order.
//	</li>
//
//	<li>
//	  <code>-friendly_name</code>: sorts the result by the
//	  <code>friendly_name</code> field in descending order.
//	</li>
//
// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
type TexmlApplicationListParamsSort string

const (
	TexmlApplicationListParamsSortCreatedAt    TexmlApplicationListParamsSort = "created_at"
	TexmlApplicationListParamsSortFriendlyName TexmlApplicationListParamsSort = "friendly_name"
	TexmlApplicationListParamsSortActive       TexmlApplicationListParamsSort = "active"
)
