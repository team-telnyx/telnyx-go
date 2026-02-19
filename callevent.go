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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// CallEventService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallEventService] method instead.
type CallEventService struct {
	Options []option.RequestOption
}

// NewCallEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCallEventService(opts ...option.RequestOption) (r CallEventService) {
	r = CallEventService{}
	r.Options = opts
	return
}

// Filters call events by given filter parameters. Events are ordered by
// `occurred_at`. If filter for `leg_id` or `application_session_id` is not
// present, it only filters events from the last 24 hours.
//
// **Note**: Only one `filter[occurred_at]` can be passed.
func (r *CallEventService) List(ctx context.Context, query CallEventListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[CallEventListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "call_events"
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

// Filters call events by given filter parameters. Events are ordered by
// `occurred_at`. If filter for `leg_id` or `application_session_id` is not
// present, it only filters events from the last 24 hours.
//
// **Note**: Only one `filter[occurred_at]` can be passed.
func (r *CallEventService) ListAutoPaging(ctx context.Context, query CallEventListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[CallEventListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type CallEventListResponse struct {
	// Uniquely identifies an individual call leg.
	CallLegID string `json:"call_leg_id,required"`
	// Uniquely identifies the call control session. A session may include multiple
	// call leg events.
	CallSessionID string `json:"call_session_id,required"`
	// Event timestamp
	EventTimestamp string `json:"event_timestamp,required"`
	// Event metadata, which includes raw event, and extra information based on event
	// type
	Metadata map[string]any `json:"metadata,required"`
	// Event name
	Name string `json:"name,required"`
	// Any of "call_event".
	RecordType CallEventListResponseRecordType `json:"record_type,required"`
	// Event type
	//
	// Any of "command", "webhook".
	Type CallEventListResponseType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallLegID      respjson.Field
		CallSessionID  respjson.Field
		EventTimestamp respjson.Field
		Metadata       respjson.Field
		Name           respjson.Field
		RecordType     respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *CallEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallEventListResponseRecordType string

const (
	CallEventListResponseRecordTypeCallEvent CallEventListResponseRecordType = "call_event"
)

// Event type
type CallEventListResponseType string

const (
	CallEventListResponseTypeCommand CallEventListResponseType = "command"
	CallEventListResponseTypeWebhook CallEventListResponseType = "webhook"
)

type CallEventListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[application_name][contains], filter[outbound.outbound_voice_profile_id],
	// filter[leg_id], filter[application_session_id], filter[connection_id],
	// filter[product], filter[failed], filter[from], filter[to], filter[name],
	// filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]
	Filter CallEventListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallEventListParams]'s query parameters as `url.Values`.
func (r CallEventListParams) URLQuery() (v url.Values, err error) {
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
type CallEventListParamsFilter struct {
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
	ApplicationName CallEventListParamsFilterApplicationName `query:"application_name,omitzero" json:"-"`
	// Event occurred_at filters
	OccurredAt CallEventListParamsFilterOccurredAt `query:"occurred_at,omitzero" json:"-"`
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

// URLQuery serializes [CallEventListParamsFilter]'s query parameters as
// `url.Values`.
func (r CallEventListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Application name filters
type CallEventListParamsFilterApplicationName struct {
	// If present, applications with <code>application_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallEventListParamsFilterApplicationName]'s query
// parameters as `url.Values`.
func (r CallEventListParamsFilterApplicationName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Event occurred_at filters
type CallEventListParamsFilterOccurredAt struct {
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

// URLQuery serializes [CallEventListParamsFilterOccurredAt]'s query parameters as
// `url.Values`.
func (r CallEventListParamsFilterOccurredAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
