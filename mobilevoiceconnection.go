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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
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
func (r *MobileVoiceConnectionService) List(ctx context.Context, query MobileVoiceConnectionListParams, opts ...option.RequestOption) (res *MobileVoiceConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/mobile_voice_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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

type MobileVoiceConnectionNewResponse struct {
	Data MobileVoiceConnectionNewResponseData `json:"data"`
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

type MobileVoiceConnectionNewResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Indicates if the connection is active.
	Active bool `json:"active"`
	// The name of the connection.
	ConnectionName string                                       `json:"connection_name"`
	CreatedAt      time.Time                                    `json:"created_at" format:"date-time"`
	Inbound        MobileVoiceConnectionNewResponseDataInbound  `json:"inbound"`
	Outbound       MobileVoiceConnectionNewResponseDataOutbound `json:"outbound"`
	// Identifies the type of the resource.
	//
	// Any of "mobile_voice_connection".
	RecordType string `json:"record_type"`
	// A list of tags associated with the connection.
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The API version for webhooks.
	//
	// Any of "1", "2".
	WebhookAPIVersion string `json:"webhook_api_version,nullable"`
	// The failover URL where webhooks are sent.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable"`
	// The URL where webhooks are sent.
	WebhookEventURL string `json:"webhook_event_url,nullable"`
	// The timeout for webhooks in seconds.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
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
func (r MobileVoiceConnectionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionNewResponseDataInbound struct {
	ChannelLimit int64 `json:"channel_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionNewResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionNewResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionNewResponseDataOutbound struct {
	ChannelLimit           int64  `json:"channel_limit,nullable"`
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionNewResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionNewResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionGetResponse struct {
	Data MobileVoiceConnectionGetResponseData `json:"data"`
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

type MobileVoiceConnectionGetResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Indicates if the connection is active.
	Active bool `json:"active"`
	// The name of the connection.
	ConnectionName string                                       `json:"connection_name"`
	CreatedAt      time.Time                                    `json:"created_at" format:"date-time"`
	Inbound        MobileVoiceConnectionGetResponseDataInbound  `json:"inbound"`
	Outbound       MobileVoiceConnectionGetResponseDataOutbound `json:"outbound"`
	// Identifies the type of the resource.
	//
	// Any of "mobile_voice_connection".
	RecordType string `json:"record_type"`
	// A list of tags associated with the connection.
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The API version for webhooks.
	//
	// Any of "1", "2".
	WebhookAPIVersion string `json:"webhook_api_version,nullable"`
	// The failover URL where webhooks are sent.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable"`
	// The URL where webhooks are sent.
	WebhookEventURL string `json:"webhook_event_url,nullable"`
	// The timeout for webhooks in seconds.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
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
func (r MobileVoiceConnectionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionGetResponseDataInbound struct {
	ChannelLimit int64 `json:"channel_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionGetResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionGetResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionGetResponseDataOutbound struct {
	ChannelLimit           int64  `json:"channel_limit,nullable"`
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionGetResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionGetResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateResponse struct {
	Data MobileVoiceConnectionUpdateResponseData `json:"data"`
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

type MobileVoiceConnectionUpdateResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Indicates if the connection is active.
	Active bool `json:"active"`
	// The name of the connection.
	ConnectionName string                                          `json:"connection_name"`
	CreatedAt      time.Time                                       `json:"created_at" format:"date-time"`
	Inbound        MobileVoiceConnectionUpdateResponseDataInbound  `json:"inbound"`
	Outbound       MobileVoiceConnectionUpdateResponseDataOutbound `json:"outbound"`
	// Identifies the type of the resource.
	//
	// Any of "mobile_voice_connection".
	RecordType string `json:"record_type"`
	// A list of tags associated with the connection.
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The API version for webhooks.
	//
	// Any of "1", "2".
	WebhookAPIVersion string `json:"webhook_api_version,nullable"`
	// The failover URL where webhooks are sent.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable"`
	// The URL where webhooks are sent.
	WebhookEventURL string `json:"webhook_event_url,nullable"`
	// The timeout for webhooks in seconds.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
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
func (r MobileVoiceConnectionUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateResponseDataInbound struct {
	ChannelLimit int64 `json:"channel_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionUpdateResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionUpdateResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionUpdateResponseDataOutbound struct {
	ChannelLimit           int64  `json:"channel_limit,nullable"`
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionUpdateResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionUpdateResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionListResponse struct {
	Data []MobileVoiceConnectionListResponseData `json:"data"`
	Meta MobileVoiceConnectionListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionListResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Indicates if the connection is active.
	Active bool `json:"active"`
	// The name of the connection.
	ConnectionName string                                        `json:"connection_name"`
	CreatedAt      time.Time                                     `json:"created_at" format:"date-time"`
	Inbound        MobileVoiceConnectionListResponseDataInbound  `json:"inbound"`
	Outbound       MobileVoiceConnectionListResponseDataOutbound `json:"outbound"`
	// Identifies the type of the resource.
	//
	// Any of "mobile_voice_connection".
	RecordType string `json:"record_type"`
	// A list of tags associated with the connection.
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The API version for webhooks.
	//
	// Any of "1", "2".
	WebhookAPIVersion string `json:"webhook_api_version,nullable"`
	// The failover URL where webhooks are sent.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable"`
	// The URL where webhooks are sent.
	WebhookEventURL string `json:"webhook_event_url,nullable"`
	// The timeout for webhooks in seconds.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
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
func (r MobileVoiceConnectionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionListResponseDataInbound struct {
	ChannelLimit int64 `json:"channel_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionListResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionListResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionListResponseDataOutbound struct {
	ChannelLimit           int64  `json:"channel_limit,nullable"`
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionListResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionListResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionListResponseMeta struct {
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
func (r MobileVoiceConnectionListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionDeleteResponse struct {
	Data MobileVoiceConnectionDeleteResponseData `json:"data"`
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

type MobileVoiceConnectionDeleteResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Indicates if the connection is active.
	Active bool `json:"active"`
	// The name of the connection.
	ConnectionName string                                          `json:"connection_name"`
	CreatedAt      time.Time                                       `json:"created_at" format:"date-time"`
	Inbound        MobileVoiceConnectionDeleteResponseDataInbound  `json:"inbound"`
	Outbound       MobileVoiceConnectionDeleteResponseDataOutbound `json:"outbound"`
	// Identifies the type of the resource.
	//
	// Any of "mobile_voice_connection".
	RecordType string `json:"record_type"`
	// A list of tags associated with the connection.
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The API version for webhooks.
	//
	// Any of "1", "2".
	WebhookAPIVersion string `json:"webhook_api_version,nullable"`
	// The failover URL where webhooks are sent.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable"`
	// The URL where webhooks are sent.
	WebhookEventURL string `json:"webhook_event_url,nullable"`
	// The timeout for webhooks in seconds.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
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
func (r MobileVoiceConnectionDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionDeleteResponseDataInbound struct {
	ChannelLimit int64 `json:"channel_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionDeleteResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionDeleteResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileVoiceConnectionDeleteResponseDataOutbound struct {
	ChannelLimit           int64  `json:"channel_limit,nullable"`
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileVoiceConnectionDeleteResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobileVoiceConnectionDeleteResponseDataOutbound) UnmarshalJSON(data []byte) error {
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
	MobileVoiceConnectionNewParamsWebhookAPIVersion1 MobileVoiceConnectionNewParamsWebhookAPIVersion = "1"
	MobileVoiceConnectionNewParamsWebhookAPIVersion2 MobileVoiceConnectionNewParamsWebhookAPIVersion = "2"
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
	MobileVoiceConnectionUpdateParamsWebhookAPIVersion1 MobileVoiceConnectionUpdateParamsWebhookAPIVersion = "1"
	MobileVoiceConnectionUpdateParamsWebhookAPIVersion2 MobileVoiceConnectionUpdateParamsWebhookAPIVersion = "2"
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
