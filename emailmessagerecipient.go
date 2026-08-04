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

// Send and manage email messages. Legacy `/v2/emails` routes are aliases for these
// endpoints.
//
// EmailMessageRecipientService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailMessageRecipientService] method instead.
type EmailMessageRecipientService struct {
	Options []option.RequestOption
}

// NewEmailMessageRecipientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailMessageRecipientService(opts ...option.RequestOption) (r EmailMessageRecipientService) {
	r = EmailMessageRecipientService{}
	r.Options = opts
	return
}

// Returns the current delivery state of a single recipient, including status,
// billable flag, SMTP detail, and lifecycle timestamps. BCC recipient addresses
// are redacted (returned as null).
func (r *EmailMessageRecipientService) Get(ctx context.Context, recipientID string, query EmailMessageRecipientGetParams, opts ...option.RequestOption) (res *EmailMessageRecipientGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.EmailID == "" {
		err = errors.New("missing required email_id parameter")
		return nil, err
	}
	if recipientID == "" {
		err = errors.New("missing required recipient_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_messages/%s/recipients/%s", query.EmailID, recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists per-recipient delivery states for a single message with cursor pagination.
// Each recipient has an independent status, billable flag, and lifecycle
// timestamps. BCC recipient addresses are redacted (returned as null) to protect
// BCC privacy. Default page size is 25, maximum is 100.
func (r *EmailMessageRecipientService) List(ctx context.Context, emailID string, query EmailMessageRecipientListParams, opts ...option.RequestOption) (res *EmailMessageRecipientListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailID == "" {
		err = errors.New("missing required email_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_messages/%s/recipients", emailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type EmailRecipient struct {
	// Recipient UUID.
	ID string `json:"id" api:"required" format:"uuid"`
	// Recipient email address. Null for BCC recipients (redacted for privacy).
	Address string `json:"address" api:"required" format:"email"`
	// Whether this recipient's delivery is billable (set on queue acceptance).
	Billable bool `json:"billable" api:"required"`
	// Any of "to", "cc", "bcc".
	Kind EmailRecipientKind `json:"kind" api:"required"`
	// Parent email message UUID.
	MessageID string `json:"message_id" api:"required" format:"uuid"`
	// Any of "email_recipient".
	RecordType EmailRecipientRecordType `json:"record_type" api:"required"`
	// Current per-recipient delivery status.
	//
	// Any of "queued", "sending", "sent", "deferred", "delivered", "bounced",
	// "failed", "gw_reject", "cancelled".
	Status      EmailRecipientStatus `json:"status" api:"required"`
	DeliveredAt time.Time            `json:"delivered_at" api:"nullable" format:"date-time"`
	FailedAt    time.Time            `json:"failed_at" api:"nullable" format:"date-time"`
	SentAt      time.Time            `json:"sent_at" api:"nullable" format:"date-time"`
	// SMTP response code when available (e.g. 550 for bounces).
	SmtpCode int64 `json:"smtp_code" api:"nullable"`
	// SMTP response message when available.
	SmtpResponse string `json:"smtp_response" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Address      respjson.Field
		Billable     respjson.Field
		Kind         respjson.Field
		MessageID    respjson.Field
		RecordType   respjson.Field
		Status       respjson.Field
		DeliveredAt  respjson.Field
		FailedAt     respjson.Field
		SentAt       respjson.Field
		SmtpCode     respjson.Field
		SmtpResponse respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailRecipient) RawJSON() string { return r.JSON.raw }
func (r *EmailRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRecipientKind string

const (
	EmailRecipientKindTo  EmailRecipientKind = "to"
	EmailRecipientKindCc  EmailRecipientKind = "cc"
	EmailRecipientKindBcc EmailRecipientKind = "bcc"
)

type EmailRecipientRecordType string

const (
	EmailRecipientRecordTypeEmailRecipient EmailRecipientRecordType = "email_recipient"
)

// Current per-recipient delivery status.
type EmailRecipientStatus string

const (
	EmailRecipientStatusQueued    EmailRecipientStatus = "queued"
	EmailRecipientStatusSending   EmailRecipientStatus = "sending"
	EmailRecipientStatusSent      EmailRecipientStatus = "sent"
	EmailRecipientStatusDeferred  EmailRecipientStatus = "deferred"
	EmailRecipientStatusDelivered EmailRecipientStatus = "delivered"
	EmailRecipientStatusBounced   EmailRecipientStatus = "bounced"
	EmailRecipientStatusFailed    EmailRecipientStatus = "failed"
	EmailRecipientStatusGwReject  EmailRecipientStatus = "gw_reject"
	EmailRecipientStatusCancelled EmailRecipientStatus = "cancelled"
)

type EmailMessageRecipientGetResponse struct {
	Data EmailRecipient `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageRecipientGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageRecipientGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageRecipientListResponse struct {
	Data []EmailRecipient                      `json:"data" api:"required"`
	Meta EmailMessageRecipientListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageRecipientListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageRecipientListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageRecipientListResponseMeta struct {
	PageSize int64 `json:"page_size" api:"required"`
	// Cursor for the next page. Absent when there are no more results.
	PageCursor string `json:"page_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageSize    respjson.Field
		PageCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailMessageRecipientListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *EmailMessageRecipientListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailMessageRecipientGetParams struct {
	EmailID string `path:"email_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type EmailMessageRecipientListParams struct {
	// Opaque URL-safe Base64 cursor returned by a previous list response.
	PageCursor param.Opt[string] `query:"page_cursor,omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100. Invalid values are
	// clamped to the valid range.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter recipients by address kind.
	//
	// Any of "to", "cc", "bcc".
	Kind EmailMessageRecipientListParamsKind `query:"kind,omitzero" json:"-"`
	// Filter recipients by status.
	//
	// Any of "queued", "sending", "sent", "deferred", "delivered", "bounced",
	// "failed", "gw_reject", "cancelled".
	Status EmailMessageRecipientListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailMessageRecipientListParams]'s query parameters as
// `url.Values`.
func (r EmailMessageRecipientListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter recipients by address kind.
type EmailMessageRecipientListParamsKind string

const (
	EmailMessageRecipientListParamsKindTo  EmailMessageRecipientListParamsKind = "to"
	EmailMessageRecipientListParamsKindCc  EmailMessageRecipientListParamsKind = "cc"
	EmailMessageRecipientListParamsKindBcc EmailMessageRecipientListParamsKind = "bcc"
)

// Filter recipients by status.
type EmailMessageRecipientListParamsStatus string

const (
	EmailMessageRecipientListParamsStatusQueued    EmailMessageRecipientListParamsStatus = "queued"
	EmailMessageRecipientListParamsStatusSending   EmailMessageRecipientListParamsStatus = "sending"
	EmailMessageRecipientListParamsStatusSent      EmailMessageRecipientListParamsStatus = "sent"
	EmailMessageRecipientListParamsStatusDeferred  EmailMessageRecipientListParamsStatus = "deferred"
	EmailMessageRecipientListParamsStatusDelivered EmailMessageRecipientListParamsStatus = "delivered"
	EmailMessageRecipientListParamsStatusBounced   EmailMessageRecipientListParamsStatus = "bounced"
	EmailMessageRecipientListParamsStatusFailed    EmailMessageRecipientListParamsStatus = "failed"
	EmailMessageRecipientListParamsStatusGwReject  EmailMessageRecipientListParamsStatus = "gw_reject"
	EmailMessageRecipientListParamsStatusCancelled EmailMessageRecipientListParamsStatus = "cancelled"
)
