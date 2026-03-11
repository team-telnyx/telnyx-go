// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
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
// WhatsappBusinessAccountSettingService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappBusinessAccountSettingService] method instead.
type WhatsappBusinessAccountSettingService struct {
	Options []option.RequestOption
}

// NewWhatsappBusinessAccountSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhatsappBusinessAccountSettingService(opts ...option.RequestOption) (r WhatsappBusinessAccountSettingService) {
	r = WhatsappBusinessAccountSettingService{}
	r.Options = opts
	return
}

// Get WABA settings
func (r *WhatsappBusinessAccountSettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappBusinessAccountSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/business_accounts/%s/settings", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update WABA settings
func (r *WhatsappBusinessAccountSettingService) Update(ctx context.Context, id string, body WhatsappBusinessAccountSettingUpdateParams, opts ...option.RequestOption) (res *WhatsappBusinessAccountSettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/business_accounts/%s/settings", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type WhatsappBusinessAccountSettingGetResponse struct {
	Data WhatsappBusinessAccountSettingGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountSettingGetResponseData struct {
	// Internal ID of Whatsapp business account
	ID         string    `json:"id" format:"uuid"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	Timezone   string    `json:"timezone"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// Enable/disable receiving Whatsapp events
	WebhookEnabled bool     `json:"webhook_enabled"`
	WebhookEvents  []string `json:"webhook_events"`
	// Failover URL to receive Whatsapp events
	WebhookFailoverURL string `json:"webhook_failover_url" format:"url"`
	// URL to receive Whatsapp events
	WebhookURL string `json:"webhook_url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		RecordType         respjson.Field
		Timezone           respjson.Field
		UpdatedAt          respjson.Field
		WebhookEnabled     respjson.Field
		WebhookEvents      respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountSettingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountSettingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountSettingUpdateResponse struct {
	Data WhatsappBusinessAccountSettingUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountSettingUpdateResponseData struct {
	// Internal ID of Whatsapp business account
	ID         string    `json:"id" format:"uuid"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	Timezone   string    `json:"timezone"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// Enable/disable receiving Whatsapp events
	WebhookEnabled bool     `json:"webhook_enabled"`
	WebhookEvents  []string `json:"webhook_events"`
	// Failover URL to receive Whatsapp events
	WebhookFailoverURL string `json:"webhook_failover_url" format:"url"`
	// URL to receive Whatsapp events
	WebhookURL string `json:"webhook_url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		RecordType         respjson.Field
		Timezone           respjson.Field
		UpdatedAt          respjson.Field
		WebhookEnabled     respjson.Field
		WebhookEvents      respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountSettingUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountSettingUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountSettingUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// IANA timezone identifier
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Enable/disable receiving Whatsapp events
	WebhookEnabled param.Opt[bool] `json:"webhook_enabled,omitzero"`
	// Failover URL to send Whatsapp events
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// URL to send Whatsapp events
	WebhookURL    param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	WebhookEvents []string          `json:"webhook_events,omitzero"`
	paramObj
}

func (r WhatsappBusinessAccountSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappBusinessAccountSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappBusinessAccountSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
