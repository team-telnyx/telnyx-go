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
		return
	}
	path := fmt.Sprintf("connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your connections irrespective of type.
func (r *ConnectionService) List(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[ConnectionListResponse], err error) {
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
func (r *ConnectionService) ListAutoPaging(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[ConnectionListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Lists all active calls for given connection. Acceptable connections are either
// SIP connections with webhook_url or xml_request_url, call control or texml.
// Returned results are cursor paginated.
func (r *ConnectionService) ListActiveCalls(ctx context.Context, connectionID string, query ConnectionListActiveCallsParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[ConnectionListActiveCallsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("connections/%s/active_calls", connectionID)
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
func (r *ConnectionService) ListActiveCallsAutoPaging(ctx context.Context, connectionID string, query ConnectionListActiveCallsParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[ConnectionListActiveCallsResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.ListActiveCalls(ctx, connectionID, query, opts...))
}

type ConnectionGetResponse struct {
	Data ConnectionGetResponseData `json:"data"`
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

type ConnectionGetResponseData struct {
	// Identifies the specific resource.
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
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id" format:"int64"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion string `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable" format:"uri"`
	// The URL where webhooks related to this connection will be sent.
	WebhookEventURL string `json:"webhook_event_url,nullable" format:"uri"`
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
func (r ConnectionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ConnectionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListResponse struct {
	// Identifies the specific resource.
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
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id" format:"int64"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion ConnectionListResponseWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable" format:"uri"`
	// The URL where webhooks related to this connection will be sent.
	WebhookEventURL string `json:"webhook_event_url,nullable" format:"uri"`
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
func (r ConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which webhook format will be used, Telnyx API v1 or v2.
type ConnectionListResponseWebhookAPIVersion string

const (
	ConnectionListResponseWebhookAPIVersionV1 ConnectionListResponseWebhookAPIVersion = "1"
	ConnectionListResponseWebhookAPIVersionV2 ConnectionListResponseWebhookAPIVersion = "2"
)

type ConnectionListActiveCallsResponse struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id,required"`
	// Indicates the duration of the call in seconds
	CallDuration int64 `json:"call_duration,required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id,required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id,required"`
	// State received from a command.
	ClientState string `json:"client_state,required"`
	// Any of "call".
	RecordType ConnectionListActiveCallsResponseRecordType `json:"record_type,required"`
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

type ConnectionListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id],
	// filter[outbound.outbound_voice_profile_id]
	Filter ConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page ConnectionListParamsPage `query:"page,omitzero" json:"-"`
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
	OutboundVoiceProfileID param.Opt[string] `query:"outbound_voice_profile_id,omitzero" format:"int64" json:"-"`
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

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type ConnectionListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectionListParamsPage]'s query parameters as
// `url.Values`.
func (r ConnectionListParamsPage) URLQuery() (v url.Values, err error) {
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
	// Consolidated page parameter (deepObject style). Originally: page[after],
	// page[before], page[limit], page[size], page[number]
	Page ConnectionListActiveCallsParamsPage `query:"page,omitzero" json:"-"`
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

// Consolidated page parameter (deepObject style). Originally: page[after],
// page[before], page[limit], page[size], page[number]
type ConnectionListActiveCallsParamsPage struct {
	// Opaque identifier of next page
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Opaque identifier of previous page
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Limit of records per single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectionListActiveCallsParamsPage]'s query parameters as
// `url.Values`.
func (r ConnectionListActiveCallsParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
