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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// WebhookDeliveryService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookDeliveryService] method instead.
type WebhookDeliveryService struct {
	Options []option.RequestOption
}

// NewWebhookDeliveryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookDeliveryService(opts ...option.RequestOption) (r WebhookDeliveryService) {
	r = WebhookDeliveryService{}
	r.Options = opts
	return
}

// Provides webhook_delivery debug data, such as timestamps, delivery status and
// attempts.
func (r *WebhookDeliveryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookDeliveryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("webhook_deliveries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists webhook_deliveries for the authenticated user
func (r *WebhookDeliveryService) List(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) (res *WebhookDeliveryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhook_deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WebhookDeliveryGetResponse struct {
	// Record of all attempts to deliver a webhook.
	Data WebhookDeliveryGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Record of all attempts to deliver a webhook.
type WebhookDeliveryGetResponseData struct {
	// Uniquely identifies the webhook_delivery record.
	ID string `json:"id" format:"uuid"`
	// Detailed delivery attempts, ordered by most recent.
	Attempts []WebhookDeliveryGetResponseDataAttempt `json:"attempts"`
	// ISO 8601 timestamp indicating when the last webhook response has been received.
	FinishedAt time.Time `json:"finished_at" format:"date-time"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 timestamp indicating when the first request attempt was initiated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Delivery status: 'delivered' when successfuly delivered or 'failed' if all
	// attempts have failed.
	//
	// Any of "delivered", "failed".
	Status string `json:"status"`
	// Uniquely identifies the user that owns the webhook_delivery record.
	UserID string `json:"user_id" format:"uuid"`
	// Original webhook JSON data. Payload fields vary according to event type.
	Webhook WebhookDeliveryGetResponseDataWebhook `json:"webhook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attempts    respjson.Field
		FinishedAt  respjson.Field
		RecordType  respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		UserID      respjson.Field
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook delivery attempt details.
type WebhookDeliveryGetResponseDataAttempt struct {
	// Webhook delivery error codes.
	Errors []int64 `json:"errors"`
	// ISO 8601 timestamp indicating when the attempt has finished.
	FinishedAt time.Time `json:"finished_at" format:"date-time"`
	// HTTP request and response information.
	HTTP WebhookDeliveryGetResponseDataAttemptHTTP `json:"http"`
	// ISO 8601 timestamp indicating when the attempt was initiated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Any of "delivered", "failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		FinishedAt  respjson.Field
		HTTP        respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponseDataAttempt) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseDataAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request and response information.
type WebhookDeliveryGetResponseDataAttemptHTTP struct {
	// Request details.
	Request WebhookDeliveryGetResponseDataAttemptHTTPRequest `json:"request"`
	// Response details, optional.
	Response WebhookDeliveryGetResponseDataAttemptHTTPResponse `json:"response,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponseDataAttemptHTTP) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseDataAttemptHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request details.
type WebhookDeliveryGetResponseDataAttemptHTTPRequest struct {
	// List of headers, limited to 10kB.
	Headers [][]string `json:"headers"`
	URL     string     `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponseDataAttemptHTTPRequest) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseDataAttemptHTTPRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response details, optional.
type WebhookDeliveryGetResponseDataAttemptHTTPResponse struct {
	// Raw response body, limited to 10kB.
	Body string `json:"body"`
	// List of headers, limited to 10kB.
	Headers [][]string `json:"headers"`
	Status  int64      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponseDataAttemptHTTPResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseDataAttemptHTTPResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Original webhook JSON data. Payload fields vary according to event type.
type WebhookDeliveryGetResponseDataWebhook struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	Payload    any       `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetResponseDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetResponseDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponse struct {
	Data []WebhookDeliveryListResponseData `json:"data"`
	Meta WebhookDeliveryListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Record of all attempts to deliver a webhook.
type WebhookDeliveryListResponseData struct {
	// Uniquely identifies the webhook_delivery record.
	ID string `json:"id" format:"uuid"`
	// Detailed delivery attempts, ordered by most recent.
	Attempts []WebhookDeliveryListResponseDataAttempt `json:"attempts"`
	// ISO 8601 timestamp indicating when the last webhook response has been received.
	FinishedAt time.Time `json:"finished_at" format:"date-time"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 timestamp indicating when the first request attempt was initiated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Delivery status: 'delivered' when successfuly delivered or 'failed' if all
	// attempts have failed.
	//
	// Any of "delivered", "failed".
	Status string `json:"status"`
	// Uniquely identifies the user that owns the webhook_delivery record.
	UserID string `json:"user_id" format:"uuid"`
	// Original webhook JSON data. Payload fields vary according to event type.
	Webhook WebhookDeliveryListResponseDataWebhook `json:"webhook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attempts    respjson.Field
		FinishedAt  respjson.Field
		RecordType  respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		UserID      respjson.Field
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook delivery attempt details.
type WebhookDeliveryListResponseDataAttempt struct {
	// Webhook delivery error codes.
	Errors []int64 `json:"errors"`
	// ISO 8601 timestamp indicating when the attempt has finished.
	FinishedAt time.Time `json:"finished_at" format:"date-time"`
	// HTTP request and response information.
	HTTP WebhookDeliveryListResponseDataAttemptHTTP `json:"http"`
	// ISO 8601 timestamp indicating when the attempt was initiated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Any of "delivered", "failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		FinishedAt  respjson.Field
		HTTP        respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseDataAttempt) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseDataAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request and response information.
type WebhookDeliveryListResponseDataAttemptHTTP struct {
	// Request details.
	Request WebhookDeliveryListResponseDataAttemptHTTPRequest `json:"request"`
	// Response details, optional.
	Response WebhookDeliveryListResponseDataAttemptHTTPResponse `json:"response,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseDataAttemptHTTP) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseDataAttemptHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request details.
type WebhookDeliveryListResponseDataAttemptHTTPRequest struct {
	// List of headers, limited to 10kB.
	Headers [][]string `json:"headers"`
	URL     string     `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseDataAttemptHTTPRequest) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseDataAttemptHTTPRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response details, optional.
type WebhookDeliveryListResponseDataAttemptHTTPResponse struct {
	// Raw response body, limited to 10kB.
	Body string `json:"body"`
	// List of headers, limited to 10kB.
	Headers [][]string `json:"headers"`
	Status  int64      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseDataAttemptHTTPResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseDataAttemptHTTPResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Original webhook JSON data. Payload fields vary according to event type.
type WebhookDeliveryListResponseDataWebhook struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	Payload    any       `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[status][eq], filter[event_type], filter[webhook][contains],
	// filter[attempts][contains], filter[started_at][gte], filter[started_at][lte],
	// filter[finished_at][gte], filter[finished_at][lte]
	Filter WebhookDeliveryListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page WebhookDeliveryListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[status][eq], filter[event_type], filter[webhook][contains],
// filter[attempts][contains], filter[started_at][gte], filter[started_at][lte],
// filter[finished_at][gte], filter[finished_at][lte]
type WebhookDeliveryListParamsFilter struct {
	// Return only webhook_deliveries matching the given value of `event_type`. Accepts
	// multiple values separated by a `,`.
	EventType  param.Opt[string]                         `query:"event_type,omitzero" json:"-"`
	Attempts   WebhookDeliveryListParamsFilterAttempts   `query:"attempts,omitzero" json:"-"`
	FinishedAt WebhookDeliveryListParamsFilterFinishedAt `query:"finished_at,omitzero" json:"-"`
	StartedAt  WebhookDeliveryListParamsFilterStartedAt  `query:"started_at,omitzero" json:"-"`
	Status     WebhookDeliveryListParamsFilterStatus     `query:"status,omitzero" json:"-"`
	Webhook    WebhookDeliveryListParamsFilterWebhook    `query:"webhook,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsFilter]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParamsFilterAttempts struct {
	// Return only webhook_deliveries whose `attempts` component contains the given
	// text
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsFilterAttempts]'s query parameters
// as `url.Values`.
func (r WebhookDeliveryListParamsFilterAttempts) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParamsFilterFinishedAt struct {
	// Return only webhook_deliveries whose delivery finished later than or at given
	// ISO 8601 datetime
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Return only webhook_deliveries whose delivery finished earlier than or at given
	// ISO 8601 datetime
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsFilterFinishedAt]'s query
// parameters as `url.Values`.
func (r WebhookDeliveryListParamsFilterFinishedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParamsFilterStartedAt struct {
	// Return only webhook_deliveries whose delivery started later than or at given ISO
	// 8601 datetime
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Return only webhook_deliveries whose delivery started earlier than or at given
	// ISO 8601 datetime
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsFilterStartedAt]'s query
// parameters as `url.Values`.
func (r WebhookDeliveryListParamsFilterStartedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParamsFilterStatus struct {
	// Return only webhook_deliveries matching the given `status`
	//
	// Any of "delivered", "failed".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsFilterStatus]'s query parameters
// as `url.Values`.
func (r WebhookDeliveryListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryListParamsFilterWebhook struct {
	// Return only webhook deliveries whose `webhook` component contains the given text
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsFilterWebhook]'s query parameters
// as `url.Values`.
func (r WebhookDeliveryListParamsFilterWebhook) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type WebhookDeliveryListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParamsPage]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
