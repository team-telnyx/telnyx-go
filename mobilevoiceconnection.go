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

// MobileVoiceConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMobileVoiceConnectionService] method instead.
type MobileVoiceConnectionService struct {
	Options []option.RequestOption
}

// NewMobileVoiceConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMobileVoiceConnectionService(opts ...option.RequestOption) (r MobileVoiceConnectionService) {
	r = MobileVoiceConnectionService{}
	r.Options = opts
	return
}

// Create a Mobile Voice Connection
func (r *MobileVoiceConnectionService) New(ctx context.Context, body MobileVoiceConnectionNewParams, opts ...option.RequestOption) (res *MobileVoiceConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/mobile_voice_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Mobile Voice Connection
func (r *MobileVoiceConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MobileVoiceConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/mobile_voice_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Mobile Voice Connection
func (r *MobileVoiceConnectionService) Update(ctx context.Context, id string, body MobileVoiceConnectionUpdateParams, opts ...option.RequestOption) (res *MobileVoiceConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/mobile_voice_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Mobile Voice Connections
func (r *MobileVoiceConnectionService) List(ctx context.Context, query MobileVoiceConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[MobileVoiceConnection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/mobile_voice_connections"
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

// List Mobile Voice Connections
func (r *MobileVoiceConnectionService) ListAutoPaging(ctx context.Context, query MobileVoiceConnectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[MobileVoiceConnection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Mobile Voice Connection
func (r *MobileVoiceConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MobileVoiceConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/mobile_voice_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MobileVoiceConnection struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Indicates if the connection is active.
	Active bool `json:"active"`
	// The name of the connection.
	ConnectionName string                        `json:"connection_name"`
	CreatedAt      time.Time                     `json:"created_at" format:"date-time"`
	Inbound        MobileVoiceConnectionInbound  `json:"inbound"`
	Outbound       MobileVoiceConnectionOutbound `json:"outbound"`
	// Identifies the type of the resource.
	//
	// Any of "mobile_voice_connection".
	RecordType MobileVoiceConnectionRecordType `json:"record_type"`
	// A list of tags associated with the connection.
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The API version for webhooks.
	//
	// Any of "1", "2".
	WebhookAPIVersion MobileVoiceConnectionWebhookAPIVersion `json:"webhook_api_version" api:"nullable"`
	// The failover URL where webhooks are sent.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url" api:"nullable"`
	// The URL where webhooks are sent.
	WebhookEventURL string `json:"webhook_event_url" api:"nullable"`
	// The timeout for webhooks in seconds.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Active                  respjson.Field
		ConnectionName          respjson.Field
		CreatedAt               respjson.Field
		Inbound                 respjson.Field
		Outbound                respjson.Field
		RecordType              respjson.Field
		Tags                    respjson.Field
		UpdatedAt               respjson.Field
		WebhookAPIVersion       respjson.Field
		WebhookEventFailoverURL respjson.Field
		WebhookEventURL         respjson.Field
		WebhookTimeoutSecs      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnection) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionInbound struct {
	ChannelLimit int64 `json:"channel_limit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionInbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionOutbound struct {
	ChannelLimit           int64  `json:"channel_limit" api:"nullable"`
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type MobileVoiceConnectionRecordType string

const (
	MobileVoiceConnectionRecordTypeMobileVoiceConnection MobileVoiceConnectionRecordType = "mobile_voice_connection"
)

// The API version for webhooks.
type MobileVoiceConnectionWebhookAPIVersion string

const (
	MobileVoiceConnectionWebhookAPIVersionV1 MobileVoiceConnectionWebhookAPIVersion = "1"
	MobileVoiceConnectionWebhookAPIVersionV2 MobileVoiceConnectionWebhookAPIVersion = "2"
)

type MobileVoiceConnectionNewResponse struct {
	Data MobileVoiceConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionGetResponse struct {
	Data MobileVoiceConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateResponse struct {
	Data MobileVoiceConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionDeleteResponse struct {
	Data MobileVoiceConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionNewParams struct {
	WebhookEventFailoverURL param.Opt[string]                      `json:"webhook_event_failover_url,omitzero"`
	WebhookEventURL         param.Opt[string]                      `json:"webhook_event_url,omitzero"`
	WebhookTimeoutSecs      param.Opt[int64]                       `json:"webhook_timeout_secs,omitzero"`
	Active                  param.Opt[bool]                        `json:"active,omitzero"`
	ConnectionName          param.Opt[string]                      `json:"connection_name,omitzero"`
	Inbound                 MobileVoiceConnectionNewParamsInbound  `json:"inbound,omitzero"`
	Outbound                MobileVoiceConnectionNewParamsOutbound `json:"outbound,omitzero"`
	Tags                    []string                               `json:"tags,omitzero"`
	// Any of "1", "2".
	WebhookAPIVersion MobileVoiceConnectionNewParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r MobileVoiceConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MobileVoiceConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobileVoiceConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionNewParamsInbound struct {
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	paramObj
}

func (r MobileVoiceConnectionNewParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow MobileVoiceConnectionNewParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobileVoiceConnectionNewParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionNewParamsOutbound struct {
	ChannelLimit           param.Opt[int64]  `json:"channel_limit,omitzero"`
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r MobileVoiceConnectionNewParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow MobileVoiceConnectionNewParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobileVoiceConnectionNewParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionNewParamsWebhookAPIVersion string

const (
	MobileVoiceConnectionNewParamsWebhookAPIVersionV1 MobileVoiceConnectionNewParamsWebhookAPIVersion = "1"
	MobileVoiceConnectionNewParamsWebhookAPIVersionV2 MobileVoiceConnectionNewParamsWebhookAPIVersion = "2"
)

type MobileVoiceConnectionUpdateParams struct {
	WebhookEventFailoverURL param.Opt[string]                         `json:"webhook_event_failover_url,omitzero"`
	WebhookEventURL         param.Opt[string]                         `json:"webhook_event_url,omitzero"`
	Active                  param.Opt[bool]                           `json:"active,omitzero"`
	ConnectionName          param.Opt[string]                         `json:"connection_name,omitzero"`
	WebhookTimeoutSecs      param.Opt[int64]                          `json:"webhook_timeout_secs,omitzero"`
	Inbound                 MobileVoiceConnectionUpdateParamsInbound  `json:"inbound,omitzero"`
	Outbound                MobileVoiceConnectionUpdateParamsOutbound `json:"outbound,omitzero"`
	Tags                    []string                                  `json:"tags,omitzero"`
	// Any of "1", "2".
	WebhookAPIVersion MobileVoiceConnectionUpdateParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r MobileVoiceConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MobileVoiceConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobileVoiceConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateParamsInbound struct {
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	paramObj
}

func (r MobileVoiceConnectionUpdateParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow MobileVoiceConnectionUpdateParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobileVoiceConnectionUpdateParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateParamsOutbound struct {
	ChannelLimit           param.Opt[int64]  `json:"channel_limit,omitzero"`
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero"`
	paramObj
}

func (r MobileVoiceConnectionUpdateParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow MobileVoiceConnectionUpdateParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobileVoiceConnectionUpdateParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateParamsWebhookAPIVersion string

const (
	MobileVoiceConnectionUpdateParamsWebhookAPIVersionV1 MobileVoiceConnectionUpdateParamsWebhookAPIVersion = "1"
	MobileVoiceConnectionUpdateParamsWebhookAPIVersionV2 MobileVoiceConnectionUpdateParamsWebhookAPIVersion = "2"
)

type MobileVoiceConnectionListParams struct {
	// Filter by connection name containing the given string
	FilterConnectionNameContains param.Opt[string] `query:"filter[connection_name][contains],omitzero" json:"-"`
	// The page number to load
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Sort by field (e.g., created_at, connection_name, active). Prefix with - for
	// descending order.
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobileVoiceConnectionListParams]'s query parameters as
// `url.Values`.
func (r MobileVoiceConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
