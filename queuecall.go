// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// Update queued call's keep_after_hangup flag
func (r *QueueCallService) Update(ctx context.Context, callControlID string, params QueueCallUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.QueueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("queues/%s/calls/%s", params.QueueName, callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Retrieve the list of calls in an existing queue
func (r *QueueCallService) List(ctx context.Context, queueName string, query QueueCallListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[QueueCallListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if queueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	path := fmt.Sprintf("queues/%s/calls", queueName)
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

// Retrieve the list of calls in an existing queue
func (r *QueueCallService) ListAutoPaging(ctx context.Context, queueName string, query QueueCallListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[QueueCallListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, queueName, query, opts...))
}

// Removes an inactive call from a queue. If the call is no longer active, use this
// command to remove it from the queue.
func (r *QueueCallService) Remove(ctx context.Context, callControlID string, body QueueCallRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.QueueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("queues/%s/calls/%s", body.QueueName, callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
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
	CallControlID string `json:"call_control_id" api:"required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id" api:"required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id" api:"required"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id" api:"required"`
	// ISO 8601 formatted date of when the call was put in the queue
	EnqueuedAt string `json:"enqueued_at" api:"required"`
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// Unique identifier of the queue the call is in.
	QueueID string `json:"queue_id" api:"required"`
	// Current position of the call in the queue
	QueuePosition int64 `json:"queue_position" api:"required"`
	// Any of "queue_call".
	RecordType string `json:"record_type" api:"required"`
	// Destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The time the call has been waiting in the queue, given in seconds
	WaitTimeSecs int64 `json:"wait_time_secs" api:"required"`
	// Indicates whether the call is still active in the queue.
	IsAlive bool `json:"is_alive"`
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
		IsAlive       respjson.Field
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
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id" api:"required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id" api:"required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id" api:"required"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id" api:"required"`
	// ISO 8601 formatted date of when the call was put in the queue
	EnqueuedAt string `json:"enqueued_at" api:"required"`
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// Unique identifier of the queue the call is in.
	QueueID string `json:"queue_id" api:"required"`
	// Current position of the call in the queue
	QueuePosition int64 `json:"queue_position" api:"required"`
	// Any of "queue_call".
	RecordType QueueCallListResponseRecordType `json:"record_type" api:"required"`
	// Destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The time the call has been waiting in the queue, given in seconds
	WaitTimeSecs int64 `json:"wait_time_secs" api:"required"`
	// Indicates whether the call is still active in the queue.
	IsAlive bool `json:"is_alive"`
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
		IsAlive       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueCallListResponse) RawJSON() string { return r.JSON.raw }
func (r *QueueCallListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueCallListResponseRecordType string

const (
	QueueCallListResponseRecordTypeQueueCall QueueCallListResponseRecordType = "queue_call"
)

type QueueCallGetParams struct {
	QueueName string `path:"queue_name" api:"required" json:"-"`
	paramObj
}

type QueueCallUpdateParams struct {
	QueueName string `path:"queue_name" api:"required" json:"-"`
	// Whether the call should remain in queue after hangup.
	KeepAfterHangup param.Opt[bool] `json:"keep_after_hangup,omitzero"`
	paramObj
}

func (r QueueCallUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow QueueCallUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueueCallUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueCallListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QueueCallListParams]'s query parameters as `url.Values`.
func (r QueueCallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QueueCallRemoveParams struct {
	QueueName string `path:"queue_name" api:"required" json:"-"`
	paramObj
}
