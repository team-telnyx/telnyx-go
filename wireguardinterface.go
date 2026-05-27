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

// WireGuard Interface operations
//
// WireguardInterfaceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWireguardInterfaceService] method instead.
type WireguardInterfaceService struct {
	Options []option.RequestOption
}

// NewWireguardInterfaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWireguardInterfaceService(opts ...option.RequestOption) (r WireguardInterfaceService) {
	r = WireguardInterfaceService{}
	r.Options = opts
	return
}

// Create a new WireGuard Interface. Current limitation of 10 interfaces per user
// can be created.
func (r *WireguardInterfaceService) New(ctx context.Context, body WireguardInterfaceNewParams, opts ...option.RequestOption) (res *WireguardInterfaceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireguard_interfaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a WireGuard Interfaces.
func (r *WireguardInterfaceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WireguardInterfaceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("wireguard_interfaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all WireGuard Interfaces.
func (r *WireguardInterfaceService) List(ctx context.Context, query WireguardInterfaceListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WireguardInterfaceRead], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "wireguard_interfaces"
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

// List all WireGuard Interfaces.
func (r *WireguardInterfaceService) ListAutoPaging(ctx context.Context, query WireguardInterfaceListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WireguardInterfaceRead] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a WireGuard Interface.
func (r *WireguardInterfaceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WireguardInterfaceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("wireguard_interfaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type WireguardInterfaceRead struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking bool `json:"enable_sip_trunking"`
	// The Telnyx WireGuard peers `Peer.endpoint` value.
	Endpoint string `json:"endpoint"`
	// A user specified name for the interface.
	Name string `json:"name"`
	// The id of the network associated with the interface.
	NetworkID string `json:"network_id" format:"uuid"`
	// The Telnyx WireGuard peers `Peer.PublicKey`.
	PublicKey string `json:"public_key"`
	// Identifies the type of the resource.
	RecordType string                       `json:"record_type"`
	Region     WireguardInterfaceReadRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		EnableSipTrunking respjson.Field
		Endpoint          respjson.Field
		Name              respjson.Field
		NetworkID         respjson.Field
		PublicKey         respjson.Field
		RecordType        respjson.Field
		Region            respjson.Field
		RegionCode        respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceRead) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceReadRegion struct {
	// Region code of the interface.
	Code string `json:"code"`
	// Region name of the interface.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceReadRegion) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceReadRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceNewResponse struct {
	Data WireguardInterfaceRead `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceGetResponse struct {
	Data WireguardInterfaceRead `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceDeleteResponse struct {
	Data WireguardInterfaceRead `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceNewParams struct {
	// The region the interface should be deployed to.
	RegionCode string `json:"region_code" api:"required"`
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking param.Opt[bool] `json:"enable_sip_trunking,omitzero"`
	// A user specified name for the interface.
	Name param.Opt[string] `json:"name,omitzero"`
	// The id of the network associated with the interface.
	NetworkID param.Opt[string] `json:"network_id,omitzero" format:"uuid"`
	paramObj
}

func (r WireguardInterfaceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WireguardInterfaceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WireguardInterfaceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
	Filter WireguardInterfaceListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WireguardInterfaceListParams]'s query parameters as
// `url.Values`.
func (r WireguardInterfaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
type WireguardInterfaceListParamsFilter struct {
	// The associated network id to filter on.
	NetworkID param.Opt[string] `query:"network_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WireguardInterfaceListParamsFilter]'s query parameters as
// `url.Values`.
func (r WireguardInterfaceListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
