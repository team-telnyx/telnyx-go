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

// FaxApplicationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFaxApplicationService] method instead.
type FaxApplicationService struct {
	Options []option.RequestOption
}

// NewFaxApplicationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFaxApplicationService(opts ...option.RequestOption) (r FaxApplicationService) {
	r = FaxApplicationService{}
	r.Options = opts
	return
}

// Creates a new Fax Application based on the parameters sent in the request. The
// application name and webhook URL are required. Once created, you can assign
// phone numbers to your application using the `/phone_numbers` endpoint.
func (r *FaxApplicationService) New(ctx context.Context, body FaxApplicationNewParams, opts ...option.RequestOption) (res *FaxApplicationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fax_applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the details of an existing Fax Application inside the 'data' attribute of
// the response.
func (r *FaxApplicationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FaxApplicationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fax_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing Fax Application based on the parameters of the
// request.
func (r *FaxApplicationService) Update(ctx context.Context, id string, body FaxApplicationUpdateParams, opts ...option.RequestOption) (res *FaxApplicationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fax_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// This endpoint returns a list of your Fax Applications inside the 'data'
// attribute of the response. You can adjust which applications are listed by using
// filters. Fax Applications are used to configure how you send and receive faxes
// using the Programmable Fax API with Telnyx.
func (r *FaxApplicationService) List(ctx context.Context, query FaxApplicationListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[FaxApplication], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "fax_applications"
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

// This endpoint returns a list of your Fax Applications inside the 'data'
// attribute of the response. You can adjust which applications are listed by using
// filters. Fax Applications are used to configure how you send and receive faxes
// using the Programmable Fax API with Telnyx.
func (r *FaxApplicationService) ListAutoPaging(ctx context.Context, query FaxApplicationListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[FaxApplication] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a Fax Application. Deletion may be prevented if the
// application is in use by phone numbers.
func (r *FaxApplicationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FaxApplicationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fax_applications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FaxApplication struct {
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
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string                 `json:"created_at"`
	Inbound   FaxApplicationInbound  `json:"inbound"`
	Outbound  FaxApplicationOutbound `json:"outbound"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Tags associated with the Fax Application.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
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
		ID                      respjson.Field
		Active                  respjson.Field
		AnchorsiteOverride      respjson.Field
		ApplicationName         respjson.Field
		CreatedAt               respjson.Field
		Inbound                 respjson.Field
		Outbound                respjson.Field
		RecordType              respjson.Field
		Tags                    respjson.Field
		UpdatedAt               respjson.Field
		WebhookEventFailoverURL respjson.Field
		WebhookEventURL         respjson.Field
		WebhookTimeoutSecs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxApplication) RawJSON() string { return r.JSON.raw }
func (r *FaxApplication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationInbound struct {
	// When set, this will limit the number of concurrent inbound calls to phone
	// numbers associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
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
		SipSubdomain                respjson.Field
		SipSubdomainReceiveSettings respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxApplicationInbound) RawJSON() string { return r.JSON.raw }
func (r *FaxApplicationInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationOutbound struct {
	// When set, this will limit the number of concurrent outbound calls to phone
	// numbers associated with this connection.
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
func (r FaxApplicationOutbound) RawJSON() string { return r.JSON.raw }
func (r *FaxApplicationOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationNewResponse struct {
	Data FaxApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxApplicationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxApplicationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationGetResponse struct {
	Data FaxApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxApplicationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxApplicationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationUpdateResponse struct {
	Data FaxApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxApplicationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxApplicationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationDeleteResponse struct {
	Data FaxApplication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxApplicationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxApplicationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationNewParams struct {
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name,required"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url,required" format:"uri"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// `Latency` directs Telnyx to route media through the site with the lowest
	// round-trip time to the user's connection. Telnyx calculates this time using ICMP
	// ping messages. This can be disabled by specifying a site to handle all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "Sydney,
	// Australia", "Amsterdam, Netherlands", "London, UK", "Toronto, Canada",
	// "Vancouver, Canada", "Frankfurt, Germany".
	AnchorsiteOverride AnchorsiteOverride              `json:"anchorsite_override,omitzero"`
	Inbound            FaxApplicationNewParamsInbound  `json:"inbound,omitzero"`
	Outbound           FaxApplicationNewParamsOutbound `json:"outbound,omitzero"`
	// Tags associated with the Fax Application.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FaxApplicationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FaxApplicationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxApplicationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationNewParamsInbound struct {
	// When set, this will limit the number of concurrent inbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
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

func (r FaxApplicationNewParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow FaxApplicationNewParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxApplicationNewParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FaxApplicationNewParamsInbound](
		"sip_subdomain_receive_settings", "only_my_connections", "from_anyone",
	)
}

type FaxApplicationNewParamsOutbound struct {
	// When set, this will limit the number of concurrent outbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r FaxApplicationNewParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow FaxApplicationNewParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxApplicationNewParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationUpdateParams struct {
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name,required"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url,required" format:"uri"`
	// Specifies an email address where faxes sent to this application will be
	// forwarded to (as pdf or tiff attachments)
	FaxEmailRecipient param.Opt[string] `json:"fax_email_recipient,omitzero"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// `Latency` directs Telnyx to route media through the site with the lowest
	// round-trip time to the user's connection. Telnyx calculates this time using ICMP
	// ping messages. This can be disabled by specifying a site to handle all media.
	//
	// Any of "Latency", "Chicago, IL", "Ashburn, VA", "San Jose, CA", "Sydney,
	// Australia", "Amsterdam, Netherlands", "London, UK", "Toronto, Canada",
	// "Vancouver, Canada", "Frankfurt, Germany".
	AnchorsiteOverride AnchorsiteOverride                 `json:"anchorsite_override,omitzero"`
	Inbound            FaxApplicationUpdateParamsInbound  `json:"inbound,omitzero"`
	Outbound           FaxApplicationUpdateParamsOutbound `json:"outbound,omitzero"`
	// Tags associated with the Fax Application.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FaxApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FaxApplicationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxApplicationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationUpdateParamsInbound struct {
	// When set, this will limit the number of concurrent inbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
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

func (r FaxApplicationUpdateParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow FaxApplicationUpdateParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxApplicationUpdateParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FaxApplicationUpdateParamsInbound](
		"sip_subdomain_receive_settings", "only_my_connections", "from_anyone",
	)
}

type FaxApplicationUpdateParamsOutbound struct {
	// When set, this will limit the number of concurrent outbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r FaxApplicationUpdateParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow FaxApplicationUpdateParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxApplicationUpdateParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxApplicationListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[application_name][contains], filter[outbound_voice_profile_id]
	Filter FaxApplicationListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page FaxApplicationListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code> -</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>application_name</code>: sorts the result by the
	//	  <code>application_name</code> field in ascending order.
	//	</li>
	//
	//	<li>
	//	  <code>-application_name</code>: sorts the result by the
	//	  <code>application_name</code> field in descending order.
	//	</li>
	//
	// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
	//
	// Any of "created_at", "application_name", "active".
	Sort FaxApplicationListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxApplicationListParams]'s query parameters as
// `url.Values`.
func (r FaxApplicationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[application_name][contains], filter[outbound_voice_profile_id]
type FaxApplicationListParamsFilter struct {
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" json:"-"`
	// Application name filtering operations
	ApplicationName FaxApplicationListParamsFilterApplicationName `query:"application_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxApplicationListParamsFilter]'s query parameters as
// `url.Values`.
func (r FaxApplicationListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Application name filtering operations
type FaxApplicationListParamsFilterApplicationName struct {
	// If present, applications with <code>application_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxApplicationListParamsFilterApplicationName]'s query
// parameters as `url.Values`.
func (r FaxApplicationListParamsFilterApplicationName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type FaxApplicationListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxApplicationListParamsPage]'s query parameters as
// `url.Values`.
func (r FaxApplicationListParamsPage) URLQuery() (v url.Values, err error) {
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
//	  <code>application_name</code>: sorts the result by the
//	  <code>application_name</code> field in ascending order.
//	</li>
//
//	<li>
//	  <code>-application_name</code>: sorts the result by the
//	  <code>application_name</code> field in descending order.
//	</li>
//
// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
type FaxApplicationListParamsSort string

const (
	FaxApplicationListParamsSortCreatedAt       FaxApplicationListParamsSort = "created_at"
	FaxApplicationListParamsSortApplicationName FaxApplicationListParamsSort = "application_name"
	FaxApplicationListParamsSortActive          FaxApplicationListParamsSort = "active"
)
