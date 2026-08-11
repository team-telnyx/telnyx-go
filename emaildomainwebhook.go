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

// Per-domain webhook endpoints with event subscriptions
//
// EmailDomainWebhookService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailDomainWebhookService] method instead.
type EmailDomainWebhookService struct {
	Options []option.RequestOption
}

// NewEmailDomainWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailDomainWebhookService(opts ...option.RequestOption) (r EmailDomainWebhookService) {
	r = EmailDomainWebhookService{}
	r.Options = opts
	return
}

// Creates a webhook endpoint subscribed to a specific allowlist of event types.
// Both `email.*` events (published by email-api) and `email_domain.*` events
// (published by this service) flow through the same webhooks.
func (r *EmailDomainWebhookService) New(ctx context.Context, domainID string, body EmailDomainWebhookNewParams, opts ...option.RequestOption) (res *EmailWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/webhooks", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the webhook subscription identified by ID within the specified email
// domain.
func (r *EmailDomainWebhookService) Get(ctx context.Context, id string, query EmailDomainWebhookGetParams, opts ...option.RequestOption) (res *EmailWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.DomainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/webhooks/%s", query.DomainID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a webhook's URL and/or event subscription. A webhook is bound to its
// domain — `domain_id` is not mutable.
func (r *EmailDomainWebhookService) Update(ctx context.Context, id string, params EmailDomainWebhookUpdateParams, opts ...option.RequestOption) (res *EmailWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.DomainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/webhooks/%s", params.DomainID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of webhook subscriptions scoped to the email domain.
// Results can be sorted by creation time.
func (r *EmailDomainWebhookService) List(ctx context.Context, domainID string, query EmailDomainWebhookListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EmailWebhook], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/webhooks", domainID)
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

// Returns a paginated list of webhook subscriptions scoped to the email domain.
// Results can be sorted by creation time.
func (r *EmailDomainWebhookService) ListAutoPaging(ctx context.Context, domainID string, query EmailDomainWebhookListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EmailWebhook] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, domainID, query, opts...))
}

// Deletes the webhook subscription identified by ID within the specified email
// domain and returns the deleted subscription.
func (r *EmailDomainWebhookService) Delete(ctx context.Context, id string, body EmailDomainWebhookDeleteParams, opts ...option.RequestOption) (res *EmailWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.DomainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_domains/%s/webhooks/%s", body.DomainID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type EmailWebhook struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	DomainID  string    `json:"domain_id" api:"required" format:"uuid"`
	// Allowlist of event types delivered to this webhook. At least one event is
	// required — there is no default-to-all.
	Events []EmailWebhookEvent `json:"events" api:"required"`
	// Any of "email_webhook".
	RecordType EmailWebhookRecordType `json:"record_type" api:"required"`
	UpdatedAt  time.Time              `json:"updated_at" api:"required" format:"date-time"`
	// HTTPS endpoint to deliver subscribed events to.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DomainID    respjson.Field
		Events      respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailWebhook) RawJSON() string { return r.JSON.raw }
func (r *EmailWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailWebhookRecordType string

const (
	EmailWebhookRecordTypeEmailWebhook EmailWebhookRecordType = "email_webhook"
)

// Event types a webhook may subscribe to. The union of email._ events (published
// by email-api) and email_domain._ lifecycle events (published by this service).
// An event not listed here can never be subscribed to and is silently dropped.
type EmailWebhookEvent string

const (
	EmailWebhookEventEmailScheduled       EmailWebhookEvent = "email.scheduled"
	EmailWebhookEventEmailSandbox         EmailWebhookEvent = "email.sandbox"
	EmailWebhookEventEmailQueued          EmailWebhookEvent = "email.queued"
	EmailWebhookEventEmailSending         EmailWebhookEvent = "email.sending"
	EmailWebhookEventEmailSent            EmailWebhookEvent = "email.sent"
	EmailWebhookEventEmailDelivered       EmailWebhookEvent = "email.delivered"
	EmailWebhookEventEmailDeferred        EmailWebhookEvent = "email.deferred"
	EmailWebhookEventEmailBounced         EmailWebhookEvent = "email.bounced"
	EmailWebhookEventEmailFailed          EmailWebhookEvent = "email.failed"
	EmailWebhookEventEmailComplained      EmailWebhookEvent = "email.complained"
	EmailWebhookEventEmailOpened          EmailWebhookEvent = "email.opened"
	EmailWebhookEventEmailClicked         EmailWebhookEvent = "email.clicked"
	EmailWebhookEventEmailUnsubscribed    EmailWebhookEvent = "email.unsubscribed"
	EmailWebhookEventEmailReceived        EmailWebhookEvent = "email.received"
	EmailWebhookEventEmailDomainCreated   EmailWebhookEvent = "email_domain.created"
	EmailWebhookEventEmailDomainVerified  EmailWebhookEvent = "email_domain.verified"
	EmailWebhookEventEmailDomainDegraded  EmailWebhookEvent = "email_domain.degraded"
	EmailWebhookEventEmailDomainSuspended EmailWebhookEvent = "email_domain.suspended"
	EmailWebhookEventEmailDomainDeleted   EmailWebhookEvent = "email_domain.deleted"
)

type EmailWebhookResponse struct {
	Data EmailWebhook `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailWebhookResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailWebhookResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OffsetPaginationMeta struct {
	PageNumber   int64 `json:"page_number" api:"required"`
	PageSize     int64 `json:"page_size" api:"required"`
	TotalPages   int64 `json:"total_pages" api:"required"`
	TotalResults int64 `json:"total_results" api:"required"`
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
func (r OffsetPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *OffsetPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainWebhookNewParams struct {
	// At least one event type is required.
	Events []EmailWebhookEvent `json:"events,omitzero" api:"required"`
	// HTTPS endpoint to deliver subscribed events to.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r EmailDomainWebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDomainWebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDomainWebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainWebhookGetParams struct {
	DomainID string `path:"domain_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type EmailDomainWebhookUpdateParams struct {
	DomainID string              `path:"domain_id" api:"required" format:"uuid" json:"-"`
	URL      param.Opt[string]   `json:"url,omitzero" format:"uri"`
	Events   []EmailWebhookEvent `json:"events,omitzero"`
	paramObj
}

func (r EmailDomainWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDomainWebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDomainWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainWebhookListParams struct {
	// Page number to return (offset pagination)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of records per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Field to sort by. Prefix with `-` for descending order.
	//
	// Any of "created_at", "-created_at".
	Sort EmailDomainWebhookListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailDomainWebhookListParams]'s query parameters as
// `url.Values`.
func (r EmailDomainWebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by. Prefix with `-` for descending order.
type EmailDomainWebhookListParamsSort string

const (
	EmailDomainWebhookListParamsSortCreatedAt     EmailDomainWebhookListParamsSort = "created_at"
	EmailDomainWebhookListParamsSortCreatedAtDesc EmailDomainWebhookListParamsSort = "-created_at"
)

type EmailDomainWebhookDeleteParams struct {
	DomainID string `path:"domain_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
