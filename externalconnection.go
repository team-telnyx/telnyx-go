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

// ExternalConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalConnectionService] method instead.
type ExternalConnectionService struct {
	Options        []option.RequestOption
	LogMessages    ExternalConnectionLogMessageService
	CivicAddresses ExternalConnectionCivicAddressService
	PhoneNumbers   ExternalConnectionPhoneNumberService
	Releases       ExternalConnectionReleaseService
	Uploads        ExternalConnectionUploadService
}

// NewExternalConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExternalConnectionService(opts ...option.RequestOption) (r ExternalConnectionService) {
	r = ExternalConnectionService{}
	r.Options = opts
	r.LogMessages = NewExternalConnectionLogMessageService(opts...)
	r.CivicAddresses = NewExternalConnectionCivicAddressService(opts...)
	r.PhoneNumbers = NewExternalConnectionPhoneNumberService(opts...)
	r.Releases = NewExternalConnectionReleaseService(opts...)
	r.Uploads = NewExternalConnectionUploadService(opts...)
	return
}

// Creates a new External Connection based on the parameters sent in the request.
// The external_sip_connection and outbound voice profile id are required. Once
// created, you can assign phone numbers to your application using the
// `/phone_numbers` endpoint.
func (r *ExternalConnectionService) New(ctx context.Context, body ExternalConnectionNewParams, opts ...option.RequestOption) (res *ExternalConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "external_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the details of an existing External Connection inside the 'data'
// attribute of the response.
func (r *ExternalConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing External Connection based on the parameters of
// the request.
func (r *ExternalConnectionService) Update(ctx context.Context, id string, body ExternalConnectionUpdateParams, opts ...option.RequestOption) (res *ExternalConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// This endpoint returns a list of your External Connections inside the 'data'
// attribute of the response. External Connections are used by Telnyx customers to
// seamless configure SIP trunking integrations with Telnyx Partners, through
// External Voice Integrations in Mission Control Portal.
func (r *ExternalConnectionService) List(ctx context.Context, query ExternalConnectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ExternalConnection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "external_connections"
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

// This endpoint returns a list of your External Connections inside the 'data'
// attribute of the response. External Connections are used by Telnyx customers to
// seamless configure SIP trunking integrations with Telnyx Partners, through
// External Voice Integrations in Mission Control Portal.
func (r *ExternalConnectionService) ListAutoPaging(ctx context.Context, query ExternalConnectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ExternalConnection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes an External Connection. Deletion may be prevented if the
// application is in use by phone numbers, is active, or if it is an Operator
// Connect connection. To remove an Operator Connect integration please contact
// Telnyx support.
func (r *ExternalConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update a location's static emergency address
func (r *ExternalConnectionService) UpdateLocation(ctx context.Context, locationID string, params ExternalConnectionUpdateLocationParams, opts ...option.RequestOption) (res *ExternalConnectionUpdateLocationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/locations/%s", params.ID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type ExternalConnection struct {
	// Uniquely identifies the resource.
	ID string `json:"id" format:"int64"`
	// Specifies whether the connection can be used.
	Active bool `json:"active"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// If the credential associated with this service is active.
	CredentialActive bool `json:"credential_active"`
	// The service that will be consuming this connection.
	//
	// Any of "zoom", "operator_connect".
	ExternalSipConnection ExternalConnectionExternalSipConnection `json:"external_sip_connection"`
	Inbound               ExternalConnectionInbound               `json:"inbound"`
	Outbound              ExternalConnectionOutbound              `json:"outbound"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Tags associated with the connection.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	//
	// Any of "1", "2".
	WebhookAPIVersion ExternalConnectionWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL string `json:"webhook_event_failover_url,nullable" format:"uri"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL string `json:"webhook_event_url" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs int64 `json:"webhook_timeout_secs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Active                  respjson.Field
		CreatedAt               respjson.Field
		CredentialActive        respjson.Field
		ExternalSipConnection   respjson.Field
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
func (r ExternalConnection) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The service that will be consuming this connection.
type ExternalConnectionExternalSipConnection string

const (
	ExternalConnectionExternalSipConnectionZoom            ExternalConnectionExternalSipConnection = "zoom"
	ExternalConnectionExternalSipConnectionOperatorConnect ExternalConnectionExternalSipConnection = "operator_connect"
)

type ExternalConnectionInbound struct {
	// When set, this will limit the number of concurrent inbound calls to phone
	// numbers associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionInbound) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionOutbound struct {
	// When set, this will limit the number of concurrent outbound calls to phone
	// numbers associated with this connection.
	ChannelLimit int64 `json:"channel_limit"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id" format:"int64"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelLimit           respjson.Field
		OutboundVoiceProfileID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionOutbound) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which webhook format will be used, Telnyx API v1 or v2.
type ExternalConnectionWebhookAPIVersion string

const (
	ExternalConnectionWebhookAPIVersionV1 ExternalConnectionWebhookAPIVersion = "1"
	ExternalConnectionWebhookAPIVersionV2 ExternalConnectionWebhookAPIVersion = "2"
)

type ExternalVoiceIntegrationsPaginationMeta struct {
	PageNumber   int64 `json:"page_number,required"`
	TotalPages   int64 `json:"total_pages,required"`
	PageSize     int64 `json:"page_size"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		TotalPages   respjson.Field
		PageSize     respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalVoiceIntegrationsPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *ExternalVoiceIntegrationsPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionNewResponse struct {
	Data ExternalConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionGetResponse struct {
	Data ExternalConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUpdateResponse struct {
	Data ExternalConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionDeleteResponse struct {
	Data ExternalConnection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUpdateLocationResponse struct {
	Data ExternalConnectionUpdateLocationResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUpdateLocationResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUpdateLocationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUpdateLocationResponseData struct {
	AcceptedAddressSuggestions bool   `json:"accepted_address_suggestions"`
	LocationID                 string `json:"location_id" format:"uuid"`
	StaticEmergencyAddressID   string `json:"static_emergency_address_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedAddressSuggestions respjson.Field
		LocationID                 respjson.Field
		StaticEmergencyAddressID   respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUpdateLocationResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUpdateLocationResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionNewParams struct {
	// The service that will be consuming this connection.
	//
	// Any of "zoom".
	ExternalSipConnection ExternalConnectionNewParamsExternalSipConnection `json:"external_sip_connection,omitzero,required"`
	Outbound              ExternalConnectionNewParamsOutbound              `json:"outbound,omitzero,required"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL param.Opt[string]                  `json:"webhook_event_url,omitzero" format:"uri"`
	Inbound         ExternalConnectionNewParamsInbound `json:"inbound,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ExternalConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The service that will be consuming this connection.
type ExternalConnectionNewParamsExternalSipConnection string

const (
	ExternalConnectionNewParamsExternalSipConnectionZoom ExternalConnectionNewParamsExternalSipConnection = "zoom"
)

type ExternalConnectionNewParamsOutbound struct {
	// When set, this will limit the number of concurrent outbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID param.Opt[string] `json:"outbound_voice_profile_id,omitzero" format:"int64"`
	paramObj
}

func (r ExternalConnectionNewParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionNewParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionNewParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property OutboundVoiceProfileID is required.
type ExternalConnectionNewParamsInbound struct {
	// The ID of the outbound voice profile to use for inbound calls.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,required"`
	// When set, this will limit the number of concurrent inbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	paramObj
}

func (r ExternalConnectionNewParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionNewParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionNewParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUpdateParams struct {
	Outbound ExternalConnectionUpdateParamsOutbound `json:"outbound,omitzero,required"`
	// The failover URL where webhooks related to this connection will be sent if
	// sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverURL param.Opt[string] `json:"webhook_event_failover_url,omitzero" format:"uri"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs param.Opt[int64] `json:"webhook_timeout_secs,omitzero"`
	// Specifies whether the connection can be used.
	Active param.Opt[bool] `json:"active,omitzero"`
	// The URL where webhooks related to this connection will be sent. Must include a
	// scheme, such as 'https'.
	WebhookEventURL param.Opt[string]                     `json:"webhook_event_url,omitzero" format:"uri"`
	Inbound         ExternalConnectionUpdateParamsInbound `json:"inbound,omitzero"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ExternalConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property OutboundVoiceProfileID is required.
type ExternalConnectionUpdateParamsOutbound struct {
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileID string `json:"outbound_voice_profile_id,required" format:"int64"`
	// When set, this will limit the number of concurrent outbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	paramObj
}

func (r ExternalConnectionUpdateParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionUpdateParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionUpdateParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUpdateParamsInbound struct {
	// When set, this will limit the number of concurrent inbound calls to phone
	// numbers associated with this connection.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	paramObj
}

func (r ExternalConnectionUpdateParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionUpdateParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionUpdateParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter parameter for external connections (deepObject style). Supports filtering
	// by connection_name, external_sip_connection, id, created_at, and phone_number.
	Filter ExternalConnectionListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionListParams]'s query parameters as
// `url.Values`.
func (r ExternalConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameter for external connections (deepObject style). Supports filtering
// by connection_name, external_sip_connection, id, created_at, and phone_number.
type ExternalConnectionListParamsFilter struct {
	// If present, connections with <code>id</code> matching the given value will be
	// returned.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// If present, connections with <code>created_at</code> date matching the given
	// YYYY-MM-DD date will be returned.
	CreatedAt      param.Opt[string]                                `query:"created_at,omitzero" json:"-"`
	ConnectionName ExternalConnectionListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	// If present, connections with <code>external_sip_connection</code> matching the
	// given value will be returned.
	//
	// Any of "zoom", "operator_connect".
	ExternalSipConnection string `query:"external_sip_connection,omitzero" json:"-"`
	// Phone number filter for connections. Note: Despite the 'contains' name, this
	// requires a full E164 match per the original specification.
	PhoneNumber ExternalConnectionListParamsFilterPhoneNumber `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionListParamsFilter]'s query parameters as
// `url.Values`.
func (r ExternalConnectionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionListParamsFilterConnectionName struct {
	// If present, connections with <code>connection_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r ExternalConnectionListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Phone number filter for connections. Note: Despite the 'contains' name, this
// requires a full E164 match per the original specification.
type ExternalConnectionListParamsFilterPhoneNumber struct {
	// If present, connections associated with the given phone_number will be returned.
	// A full match is necessary with a e164 format.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionListParamsFilterPhoneNumber]'s query
// parameters as `url.Values`.
func (r ExternalConnectionListParamsFilterPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionUpdateLocationParams struct {
	ID string `path:"id,required" format:"uuid" json:"-"`
	// A new static emergency address ID to update the location with
	StaticEmergencyAddressID string `json:"static_emergency_address_id,required" format:"uuid"`
	paramObj
}

func (r ExternalConnectionUpdateLocationParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionUpdateLocationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionUpdateLocationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
