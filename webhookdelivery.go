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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
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
func (r *WebhookDeliveryService) List(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WebhookDeliveryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "webhook_deliveries"
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

// Lists webhook_deliveries for the authenticated user
func (r *WebhookDeliveryService) ListAutoPaging(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WebhookDeliveryListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Webhook delivery attempt details.
type Attempt struct {
	// Webhook delivery error codes.
	Errors []int64 `json:"errors"`
	// ISO 8601 timestamp indicating when the attempt has finished.
	FinishedAt time.Time `json:"finished_at" format:"date-time"`
	// HTTP request and response information.
	HTTP HTTP `json:"http"`
	// ISO 8601 timestamp indicating when the attempt was initiated.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Any of "delivered", "failed".
	Status AttemptStatus `json:"status"`
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
func (r Attempt) RawJSON() string { return r.JSON.raw }
func (r *Attempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttemptStatus string

const (
	AttemptStatusDelivered AttemptStatus = "delivered"
	AttemptStatusFailed    AttemptStatus = "failed"
)

// HTTP request and response information.
type HTTP struct {
	// Request details.
	Request HTTPRequest `json:"request"`
	// Response details, optional.
	Response HTTPResponse `json:"response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HTTP) RawJSON() string { return r.JSON.raw }
func (r *HTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request details.
type HTTPRequest struct {
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
func (r HTTPRequest) RawJSON() string { return r.JSON.raw }
func (r *HTTPRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response details, optional.
type HTTPResponse struct {
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
func (r HTTPResponse) RawJSON() string { return r.JSON.raw }
func (r *HTTPResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Attempts []Attempt `json:"attempts"`
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

// Original webhook JSON data. Payload fields vary according to event type.
type WebhookDeliveryGetResponseDataWebhook struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time      `json:"occurred_at" format:"date-time"`
	Payload    map[string]any `json:"payload"`
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

// Record of all attempts to deliver a webhook.
type WebhookDeliveryListResponse struct {
	// Uniquely identifies the webhook_delivery record.
	ID string `json:"id" format:"uuid"`
	// Detailed delivery attempts, ordered by most recent.
	Attempts []Attempt `json:"attempts"`
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
	Status WebhookDeliveryListResponseStatus `json:"status"`
	// Uniquely identifies the user that owns the webhook_delivery record.
	UserID string `json:"user_id" format:"uuid"`
	// Original webhook JSON data. Payload fields vary according to event type.
	Webhook WebhookDeliveryListResponseWebhook `json:"webhook"`
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
func (r WebhookDeliveryListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delivery status: 'delivered' when successfuly delivered or 'failed' if all
// attempts have failed.
type WebhookDeliveryListResponseStatus string

const (
	WebhookDeliveryListResponseStatusDelivered WebhookDeliveryListResponseStatus = "delivered"
	WebhookDeliveryListResponseStatusFailed    WebhookDeliveryListResponseStatus = "failed"
)

// Original webhook JSON data. Payload fields vary according to event type.
type WebhookDeliveryListResponseWebhook struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time      `json:"occurred_at" format:"date-time"`
	Payload    map[string]any `json:"payload"`
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
func (r WebhookDeliveryListResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[status][eq], filter[event_type], filter[webhook][contains],
	// filter[attempts][contains], filter[started_at][gte], filter[started_at][lte],
	// filter[finished_at][gte], filter[finished_at][lte]
	Filter WebhookDeliveryListParamsFilter `query:"filter,omitzero" json:"-"`
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
