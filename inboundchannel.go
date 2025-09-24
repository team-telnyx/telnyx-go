// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// InboundChannelService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundChannelService] method instead.
type InboundChannelService struct {
	Options []option.RequestOption
}

// NewInboundChannelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboundChannelService(opts ...option.RequestOption) (r InboundChannelService) {
	r = InboundChannelService{}
	r.Options = opts
	return
}

// Update the number of Voice Channels for the US Zone. This allows your account to
// handle multiple simultaneous inbound calls to US numbers. Use this endpoint to
// increase or decrease your capacity based on expected call volume.
func (r *InboundChannelService) Update(ctx context.Context, body InboundChannelUpdateParams, opts ...option.RequestOption) (res *InboundChannelUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inbound_channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns the US Zone voice channels for your account. voice channels allows you
// to use Channel Billing for calls to your Telnyx phone numbers. Please check the
// <a href="https://support.telnyx.com/en/articles/8428806-global-channel-billing">Telnyx
// Support Articles</a> section for full information and examples of how to utilize
// Channel Billing.
func (r *InboundChannelService) List(ctx context.Context, opts ...option.RequestOption) (res *InboundChannelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inbound_channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InboundChannelUpdateResponse struct {
	Data InboundChannelUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundChannelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundChannelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundChannelUpdateResponseData struct {
	// The number of channels set for the account
	Channels int64 `json:"channels"`
	// Identifies the type of the response
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundChannelUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *InboundChannelUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundChannelListResponse struct {
	Data InboundChannelListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundChannelListResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundChannelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundChannelListResponseData struct {
	// The current number of concurrent channels set for the account
	Channels int64 `json:"channels"`
	// Identifies the type of the response
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundChannelListResponseData) RawJSON() string { return r.JSON.raw }
func (r *InboundChannelListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundChannelUpdateParams struct {
	// The new number of concurrent channels for the account
	Channels int64 `json:"channels,required"`
	paramObj
}

func (r InboundChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboundChannelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
