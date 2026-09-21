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

// ConnectionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r ConnectionService) {
	r = ConnectionService{}
	r.Options = opts
	return
}

// Retrieves the high-level details of an existing connection. To retrieve specific
// authentication information, use the endpoint for the specific connection type.
func (r *ConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a list of your connections irrespective of type.
func (r *ConnectionService) List(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Connection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "connections"
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

// Returns a list of your connections irrespective of type.
func (r *ConnectionService) ListAutoPaging(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Connection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Lists all active calls for given connection. Acceptable connections are either
// SIP connections with webhook_url or xml_request_url, call control or texml.
// Returned results are cursor paginated.
func (r *ConnectionService) ListActiveCalls(ctx context.Context, connectionID string, query ConnectionListActiveCallsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ConnectionListActiveCallsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s/active_calls", url.PathEscape(connectionID))
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

// Lists all active calls for given connection. Acceptable connections are either
// SIP connections with webhook_url or xml_request_url, call control or texml.
// Returned results are cursor paginated.
func (r *ConnectionService) ListActiveCallsAutoPaging(ctx context.Context, connectionID string, query ConnectionListActiveCallsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ConnectionListActiveCallsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListActiveCalls(ctx, connectionID, query, opts...))
}

// Returns the number of connections associated with the authenticated user,
// grouped by connection type, together with the connection limits that apply to
// the user. Forward-only connections are excluded from the counts.
func (r *ConnectionService) GetCount(ctx context.Context, opts ...option.RequestOption) (res *ConnectionGetCountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connections/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Connection struct {
	// Identifies the specific resource.
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
	ConnectionName     string             `json:"connection_name"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion ConnectionWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url" api:"nullable" format:"uri"`
	// The URL where webhooks related to this connection will be sent.
	WebhookEventURL string `json:"webhook_event_url" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Active                  respjson.Field
		AnchorsiteOverride      respjson.Field
		ConnectionName          respjson.Field
		CreatedAt               respjson.Field
		OutboundVoiceProfileID  respjson.Field
		RecordType              respjson.Field
		Tags                    respjson.Field
		UpdatedAt               respjson.Field
		WebhookAPIVersion       respjson.Field
		WebhookEventFailoverURL respjson.Field
		WebhookEventURL         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Connection) RawJSON() string { return r.JSON.raw }
func (r *Connection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which webhook format will be used, Telnyx API v1 or v2.
type ConnectionWebhookAPIVersion string

const (
	ConnectionWebhookAPIVersionV1 ConnectionWebhookAPIVersion = "1"
	ConnectionWebhookAPIVersionV2 ConnectionWebhookAPIVersion = "2"
)

type ConnectionGetResponse struct {
	Data Connection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListActiveCallsResponse struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id" api:"required"`
	// Indicates the duration of the call in seconds
	CallDuration int64 `json:"call_duration" api:"required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id" api:"required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id" api:"required"`
	// State received from a command.
	ClientState string `json:"client_state" api:"required"`
	// Any of "call".
	RecordType ConnectionListActiveCallsResponseRecordType `json:"record_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallDuration  respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		RecordType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionListActiveCallsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectionListActiveCallsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListActiveCallsResponseRecordType string

const (
	ConnectionListActiveCallsResponseRecordTypeCall ConnectionListActiveCallsResponseRecordType = "call"
)

type ConnectionGetCountResponse struct {
	Data ConnectionGetCountResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionGetCountResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectionGetCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionGetCountResponseData struct {
	// Counts of the authenticated user's connections, grouped by connection type.
	// Forward-only connections are excluded.
	Counts ConnectionGetCountResponseDataCounts `json:"counts" api:"required"`
	// Connection limits that apply to the user. Contains a single global_limit when a
	// global connection limit applies, or per-type limits (standard_limit, texml_limit
	// and uac_limit) when the user has per-type connection count capabilities.
	Limits ConnectionGetCountResponseDataLimitsUnion `json:"limits" api:"required"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counts      respjson.Field
		Limits      respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionGetCountResponseData) RawJSON() string { return r.JSON.raw }
func (r *ConnectionGetCountResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Counts of the authenticated user's connections, grouped by connection type.
// Forward-only connections are excluded.
type ConnectionGetCountResponseDataCounts struct {
	// Number of Call Control applications.
	CallControlApplications int64 `json:"call_control_applications" api:"required"`
	// Number of credential connections.
	CredentialConnections int64 `json:"credential_connections" api:"required"`
	// Number of external connections.
	ExternalConnections int64 `json:"external_connections" api:"required"`
	// Number of Fax applications.
	FaxConnections int64 `json:"fax_connections" api:"required"`
	// Number of FQDN connections.
	FqdnConnections int64 `json:"fqdn_connections" api:"required"`
	// Number of IP connections.
	IPConnections int64 `json:"ip_connections" api:"required"`
	// Number of Microsoft Teams SBC (direct routing) connections.
	MicrosoftTeamsSbcConnections int64 `json:"microsoft_teams_sbc_connections" api:"required"`
	// Number of mobile voice (IMS) connections.
	MobileVoiceConnections int64 `json:"mobile_voice_connections" api:"required"`
	// Number of Microsoft Operator Connect connections.
	OperatorConnectConnections int64 `json:"operator_connect_connections" api:"required"`
	// Number of TeXML applications.
	TexmlApplications int64 `json:"texml_applications" api:"required"`
	// Number of third-party provider connections.
	ThirdPartyProviderConnections int64 `json:"third_party_provider_connections" api:"required"`
	// Number of UAC connections.
	UacConnections int64 `json:"uac_connections" api:"required"`
	// Number of Zoom SBC connections.
	ZoomSbcConnections int64 `json:"zoom_sbc_connections" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlApplications       respjson.Field
		CredentialConnections         respjson.Field
		ExternalConnections           respjson.Field
		FaxConnections                respjson.Field
		FqdnConnections               respjson.Field
		IPConnections                 respjson.Field
		MicrosoftTeamsSbcConnections  respjson.Field
		MobileVoiceConnections        respjson.Field
		OperatorConnectConnections    respjson.Field
		TexmlApplications             respjson.Field
		ThirdPartyProviderConnections respjson.Field
		UacConnections                respjson.Field
		ZoomSbcConnections            respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionGetCountResponseDataCounts) RawJSON() string { return r.JSON.raw }
func (r *ConnectionGetCountResponseDataCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConnectionGetCountResponseDataLimitsUnion contains all possible properties and
// values from [ConnectionGetCountResponseDataLimitsGlobalConnectionLimit],
// [ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConnectionGetCountResponseDataLimitsUnion struct {
	// This field is from variant
	// [ConnectionGetCountResponseDataLimitsGlobalConnectionLimit].
	GlobalLimit int64 `json:"global_limit"`
	// This field is from variant
	// [ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits].
	StandardLimit int64 `json:"standard_limit"`
	// This field is from variant
	// [ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits].
	TexmlLimit int64 `json:"texml_limit"`
	// This field is from variant
	// [ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits].
	UacLimit int64 `json:"uac_limit"`
	JSON     struct {
		GlobalLimit   respjson.Field
		StandardLimit respjson.Field
		TexmlLimit    respjson.Field
		UacLimit      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ConnectionGetCountResponseDataLimitsUnion) AsGlobalConnectionLimit() (v ConnectionGetCountResponseDataLimitsGlobalConnectionLimit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConnectionGetCountResponseDataLimitsUnion) AsPerTypeConnectionLimits() (v ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConnectionGetCountResponseDataLimitsUnion) RawJSON() string { return u.JSON.raw }

func (r *ConnectionGetCountResponseDataLimitsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionGetCountResponseDataLimitsGlobalConnectionLimit struct {
	// Maximum total number of connections allowed, when a global limit applies.
	GlobalLimit int64 `json:"global_limit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalLimit respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionGetCountResponseDataLimitsGlobalConnectionLimit) RawJSON() string {
	return r.JSON.raw
}
func (r *ConnectionGetCountResponseDataLimitsGlobalConnectionLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits struct {
	// Maximum number of standard connections allowed, when per-type limits apply.
	StandardLimit int64 `json:"standard_limit" api:"required"`
	// Maximum number of TeXML applications allowed, when per-type limits apply.
	TexmlLimit int64 `json:"texml_limit" api:"required"`
	// Maximum number of UAC connections allowed, when per-type limits apply.
	UacLimit int64 `json:"uac_limit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StandardLimit respjson.Field
		TexmlLimit    respjson.Field
		UacLimit      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits) RawJSON() string {
	return r.JSON.raw
}
func (r *ConnectionGetCountResponseDataLimitsPerTypeConnectionLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
	// filter[outbound.outbound_voice_profile_id]
	Filter ConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
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
	Sort ConnectionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectionListParams]'s query parameters as `url.Values`.
func (r ConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
// filter[outbound.outbound_voice_profile_id]
type ConnectionListParamsFilter struct {
	// If present, connections with an `fqdn` that equals the given value will be
	// returned. Matching is case-sensitive, and the full string must match.
	Fqdn param.Opt[string] `query:"fqdn,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" json:"-"`
	// Filter by connection_name using nested operations
	ConnectionName ConnectionListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectionListParamsFilter]'s query parameters as
// `url.Values`.
func (r ConnectionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by connection_name using nested operations
type ConnectionListParamsFilterConnectionName struct {
	// If present, connections with <code>connection_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectionListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r ConnectionListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
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
type ConnectionListParamsSort string

const (
	ConnectionListParamsSortCreatedAt      ConnectionListParamsSort = "created_at"
	ConnectionListParamsSortConnectionName ConnectionListParamsSort = "connection_name"
	ConnectionListParamsSortActive         ConnectionListParamsSort = "active"
)

type ConnectionListActiveCallsParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectionListActiveCallsParams]'s query parameters as
// `url.Values`.
func (r ConnectionListActiveCallsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
