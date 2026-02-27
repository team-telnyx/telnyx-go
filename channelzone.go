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

// Voice Channels
//
// ChannelZoneService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelZoneService] method instead.
type ChannelZoneService struct {
	Options []option.RequestOption
}

// NewChannelZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelZoneService(opts ...option.RequestOption) (r ChannelZoneService) {
	r = ChannelZoneService{}
	r.Options = opts
	return
}

// Update the number of Voice Channels for the Non-US Zones. This allows your
// account to handle multiple simultaneous inbound calls to Non-US numbers. Use
// this endpoint to increase or decrease your capacity based on expected call
// volume.
func (r *ChannelZoneService) Update(ctx context.Context, channelZoneID string, body ChannelZoneUpdateParams, opts ...option.RequestOption) (res *ChannelZoneUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelZoneID == "" {
		err = errors.New("missing required channel_zone_id parameter")
		return
	}
	path := fmt.Sprintf("channel_zones/%s", channelZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the non-US voice channels for your account. voice channels allow you to
// use Channel Billing for calls to your Telnyx phone numbers. Please check the
// <a href="https://support.telnyx.com/en/articles/8428806-global-channel-billing">Telnyx
// Support Articles</a> section for full information and examples of how to utilize
// Channel Billing.
func (r *ChannelZoneService) List(ctx context.Context, query ChannelZoneListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ChannelZoneListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "channel_zones"
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

// Returns the non-US voice channels for your account. voice channels allow you to
// use Channel Billing for calls to your Telnyx phone numbers. Please check the
// <a href="https://support.telnyx.com/en/articles/8428806-global-channel-billing">Telnyx
// Support Articles</a> section for full information and examples of how to utilize
// Channel Billing.
func (r *ChannelZoneService) ListAutoPaging(ctx context.Context, query ChannelZoneListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ChannelZoneListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type ChannelZoneUpdateResponse struct {
	ID       string `json:"id" api:"required"`
	Channels int64  `json:"channels" api:"required"`
	// List of countries (in ISO 3166-2, capitalized) members of the billing channel
	// zone
	Countries []string `json:"countries" api:"required"`
	Name      string   `json:"name" api:"required"`
	// Any of "channel_zone".
	RecordType ChannelZoneUpdateResponseRecordType `json:"record_type" api:"required"`
	// ISO 8601 formatted date of when the channel zone was created
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date of when the channel zone was updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Channels    respjson.Field
		Countries   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelZoneUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ChannelZoneUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelZoneUpdateResponseRecordType string

const (
	ChannelZoneUpdateResponseRecordTypeChannelZone ChannelZoneUpdateResponseRecordType = "channel_zone"
)

type ChannelZoneListResponse struct {
	ID       string `json:"id" api:"required"`
	Channels int64  `json:"channels" api:"required"`
	// List of countries (in ISO 3166-2, capitalized) members of the billing channel
	// zone
	Countries []string `json:"countries" api:"required"`
	Name      string   `json:"name" api:"required"`
	// Any of "channel_zone".
	RecordType ChannelZoneListResponseRecordType `json:"record_type" api:"required"`
	// ISO 8601 formatted date of when the channel zone was created
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date of when the channel zone was updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Channels    respjson.Field
		Countries   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelZoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChannelZoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelZoneListResponseRecordType string

const (
	ChannelZoneListResponseRecordTypeChannelZone ChannelZoneListResponseRecordType = "channel_zone"
)

type ChannelZoneUpdateParams struct {
	// The number of reserved channels
	Channels int64 `json:"channels" api:"required"`
	paramObj
}

func (r ChannelZoneUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelZoneUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelZoneUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelZoneListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelZoneListParams]'s query parameters as `url.Values`.
func (r ChannelZoneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
