// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage Whatsapp business accounts
//
// WhatsappUserDataService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappUserDataService] method instead.
type WhatsappUserDataService struct {
	Options []option.RequestOption
}

// NewWhatsappUserDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappUserDataService(opts ...option.RequestOption) (r WhatsappUserDataService) {
	r = WhatsappUserDataService{}
	r.Options = opts
	return
}

// Fetch Whatsapp user data
func (r *WhatsappUserDataService) Get(ctx context.Context, opts ...option.RequestOption) (res *WhatsappUserDataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/whatsapp/user_data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Whatsapp user data
func (r *WhatsappUserDataService) Update(ctx context.Context, body WhatsappUserDataUpdateParams, opts ...option.RequestOption) (res *WhatsappUserDataUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/whatsapp/user_data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type WhatsappUserDataGetResponse struct {
	Data WhatsappUserDataGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappUserDataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappUserDataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappUserDataGetResponseData struct {
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	RecordType string    `json:"record_type"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// Failover URL to receive Whatsapp signup events
	WebhookFailoverURL string `json:"webhook_failover_url" format:"url"`
	// URL to receive Whatsapp signup events
	WebhookURL string `json:"webhook_url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		RecordType         respjson.Field
		UpdatedAt          respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappUserDataGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappUserDataGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappUserDataUpdateResponse struct {
	Data WhatsappUserDataUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappUserDataUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappUserDataUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappUserDataUpdateResponseData struct {
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	RecordType string    `json:"record_type"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// Failover URL to receive Whatsapp signup events
	WebhookFailoverURL string `json:"webhook_failover_url" format:"url"`
	// URL to receive Whatsapp signup events
	WebhookURL string `json:"webhook_url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		RecordType         respjson.Field
		UpdatedAt          respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappUserDataUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappUserDataUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappUserDataUpdateParams struct {
	// Failover URL to send Whatsapp signup events
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// URL to send Whatsapp signup events
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	paramObj
}

func (r WhatsappUserDataUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappUserDataUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappUserDataUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
