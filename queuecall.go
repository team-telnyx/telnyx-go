// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// QueueCallService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueCallService] method instead.
type QueueCallService struct {
	Options []option.RequestOption
}

// NewQueueCallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueueCallService(opts ...option.RequestOption) (r QueueCallService) {
	r = QueueCallService{}
	r.Options = opts
	return
}

// Retrieve an existing call from an existing queue
func (r *QueueCallService) Get(ctx context.Context, callControlID string, query QueueCallGetParams, opts ...option.RequestOption) (res *QueueCallGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.QueueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("queues/%s/calls/%s", query.QueueName, callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the list of calls in an existing queue
func (r *QueueCallService) List(ctx context.Context, queueName string, query QueueCallListParams, opts ...option.RequestOption) (res *QueueCallListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	path := fmt.Sprintf("queues/%s/calls", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type QueueCallGetResponse struct {
	Data QueueCallGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueCallGetResponse) RawJSON() string { return r.JSON.raw }
func (r *QueueCallGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueCallGetResponseData struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id,required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id,required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id,required"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id,required"`
	// ISO 8601 formatted date of when the call was put in the queue
	EnqueuedAt string `json:"enqueued_at,required"`
	// Number or SIP URI placing the call.
	From string `json:"from,required"`
	// Unique identifier of the queue the call is in.
	QueueID string `json:"queue_id,required"`
	// Current position of the call in the queue
	QueuePosition int64 `json:"queue_position,required"`
	// Any of "queue_call".
	RecordType string `json:"record_type,required"`
	// Destination number or SIP URI of the call.
	To string `json:"to,required"`
	// The time the call has been waiting in the queue, given in seconds
	WaitTimeSecs int64 `json:"wait_time_secs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ConnectionID  respjson.Field
		EnqueuedAt    respjson.Field
		From          respjson.Field
		QueueID       respjson.Field
		QueuePosition respjson.Field
		RecordType    respjson.Field
		To            respjson.Field
		WaitTimeSecs  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueCallGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *QueueCallGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueCallListResponse struct {
	Data []QueueCallListResponseData `json:"data"`
	Meta PaginationMeta              `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueCallListResponse) RawJSON() string { return r.JSON.raw }
func (r *QueueCallListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueCallListResponseData struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id,required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id,required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id,required"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id,required"`
	// ISO 8601 formatted date of when the call was put in the queue
	EnqueuedAt string `json:"enqueued_at,required"`
	// Number or SIP URI placing the call.
	From string `json:"from,required"`
	// Unique identifier of the queue the call is in.
	QueueID string `json:"queue_id,required"`
	// Current position of the call in the queue
	QueuePosition int64 `json:"queue_position,required"`
	// Any of "queue_call".
	RecordType string `json:"record_type,required"`
	// Destination number or SIP URI of the call.
	To string `json:"to,required"`
	// The time the call has been waiting in the queue, given in seconds
	WaitTimeSecs int64 `json:"wait_time_secs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ConnectionID  respjson.Field
		EnqueuedAt    respjson.Field
		From          respjson.Field
		QueueID       respjson.Field
		QueuePosition respjson.Field
		RecordType    respjson.Field
		To            respjson.Field
		WaitTimeSecs  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueCallListResponseData) RawJSON() string { return r.JSON.raw }
func (r *QueueCallListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueCallGetParams struct {
	QueueName string `path:"queue_name,required" json:"-"`
	paramObj
}

type QueueCallListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[after],
	// page[before], page[limit], page[size], page[number]
	Page QueueCallListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QueueCallListParams]'s query parameters as `url.Values`.
func (r QueueCallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[after],
// page[before], page[limit], page[size], page[number]
type QueueCallListParamsPage struct {
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

// URLQuery serializes [QueueCallListParamsPage]'s query parameters as
// `url.Values`.
func (r QueueCallListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
