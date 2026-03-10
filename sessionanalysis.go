// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Analyze voice AI sessions, costs, and event hierarchies across Telnyx products.
//
// SessionAnalysisService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionAnalysisService] method instead.
type SessionAnalysisService struct {
	Options []option.RequestOption
	// Analyze voice AI sessions, costs, and event hierarchies across Telnyx products.
	Metadata SessionAnalysisMetadataService
}

// NewSessionAnalysisService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionAnalysisService(opts ...option.RequestOption) (r SessionAnalysisService) {
	r = SessionAnalysisService{}
	r.Options = opts
	r.Metadata = NewSessionAnalysisMetadataService(opts...)
	return
}

// Retrieves a full session analysis tree for a given event, including costs, child
// events, and product linkages.
func (r *SessionAnalysisService) Get(ctx context.Context, eventID string, params SessionAnalysisGetParams, opts ...option.RequestOption) (res *SessionAnalysisGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.RecordType == "" {
		err = errors.New("missing required record_type parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("session_analysis/%s/%s", params.RecordType, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type EventNode struct {
	// Event identifier.
	ID string `json:"id" api:"required"`
	// Child events in the session tree.
	Children []EventNode   `json:"children" api:"required"`
	Cost     EventNodeCost `json:"cost" api:"required"`
	// Name of the event type.
	EventName string         `json:"event_name" api:"required"`
	Links     EventNodeLinks `json:"links" api:"required"`
	// Product that generated this event.
	Product string `json:"product" api:"required"`
	// The underlying detail record data. Contents vary by record type.
	Record map[string]any `json:"record" api:"required"`
	// Relationship to the parent node, null for root.
	Relationship EventNodeRelationship `json:"relationship" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Children     respjson.Field
		Cost         respjson.Field
		EventName    respjson.Field
		Links        respjson.Field
		Product      respjson.Field
		Record       respjson.Field
		Relationship respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventNode) RawJSON() string { return r.JSON.raw }
func (r *EventNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventNodeCost struct {
	// Cumulative cost including all descendants.
	CumulativeCost string `json:"cumulative_cost" api:"required"`
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Cost of this individual event.
	EventCost string `json:"event_cost" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CumulativeCost respjson.Field
		Currency       respjson.Field
		EventCost      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventNodeCost) RawJSON() string { return r.JSON.raw }
func (r *EventNodeCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventNodeLinks struct {
	// Link to the underlying detail records.
	Records string `json:"records" api:"required"`
	// Link to this session analysis node.
	Self string `json:"self" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Records     respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventNodeLinks) RawJSON() string { return r.JSON.raw }
func (r *EventNodeLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relationship to the parent node, null for root.
type EventNodeRelationship struct {
	// Identifier of the parent event.
	ParentID string `json:"parent_id" api:"required"`
	// Relationship type identifier.
	Type string                   `json:"type" api:"required"`
	Via  EventNodeRelationshipVia `json:"via" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParentID    respjson.Field
		Type        respjson.Field
		Via         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventNodeRelationship) RawJSON() string { return r.JSON.raw }
func (r *EventNodeRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventNodeRelationshipVia struct {
	// Field name on the child record.
	LocalField string `json:"local_field" api:"required"`
	// Field name on the parent record.
	ParentField string `json:"parent_field" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LocalField  respjson.Field
		ParentField respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventNodeRelationshipVia) RawJSON() string { return r.JSON.raw }
func (r *EventNodeRelationshipVia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisGetResponse struct {
	Cost SessionAnalysisGetResponseCost `json:"cost" api:"required"`
	// When the session started.
	CreatedAt time.Time                      `json:"created_at" api:"required" format:"date-time"`
	Meta      SessionAnalysisGetResponseMeta `json:"meta" api:"required"`
	Root      EventNode                      `json:"root" api:"required"`
	// Identifier for the analyzed session.
	SessionID string `json:"session_id" api:"required"`
	// Analysis status (e.g. "completed").
	Status string `json:"status" api:"required"`
	// When the session completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		CreatedAt   respjson.Field
		Meta        respjson.Field
		Root        respjson.Field
		SessionID   respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisGetResponseCost struct {
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Total session cost as a decimal string.
	Total string `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisGetResponseCost) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisGetResponseCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisGetResponseMeta struct {
	// Total number of events in the session tree.
	EventCount int64 `json:"event_count" api:"required"`
	// List of distinct products involved in the session.
	Products []string `json:"products" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventCount  respjson.Field
		Products    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisGetParams struct {
	RecordType string `path:"record_type" api:"required" json:"-"`
	// ISO 8601 timestamp to narrow index selection for faster lookups.
	DateTime param.Opt[time.Time] `query:"date_time,omitzero" format:"date-time" json:"-"`
	// Whether to include child events in the response.
	IncludeChildren param.Opt[bool] `query:"include_children,omitzero" json:"-"`
	// Maximum traversal depth for the event tree.
	MaxDepth param.Opt[int64] `query:"max_depth,omitzero" json:"-"`
	// Controls what data to expand on each event node.
	//
	// Any of "record", "none".
	Expand SessionAnalysisGetParamsExpand `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SessionAnalysisGetParams]'s query parameters as
// `url.Values`.
func (r SessionAnalysisGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls what data to expand on each event node.
type SessionAnalysisGetParamsExpand string

const (
	SessionAnalysisGetParamsExpandRecord SessionAnalysisGetParamsExpand = "record"
	SessionAnalysisGetParamsExpandNone   SessionAnalysisGetParamsExpand = "none"
)
