// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

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
	opts = append(r.Options[:], opts...)
	path := "wireguard_interfaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a WireGuard Interfaces.
func (r *WireguardInterfaceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WireguardInterfaceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireguard_interfaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all WireGuard Interfaces.
func (r *WireguardInterfaceService) List(ctx context.Context, query WireguardInterfaceListParams, opts ...option.RequestOption) (res *WireguardInterfaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "wireguard_interfaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a WireGuard Interface.
func (r *WireguardInterfaceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WireguardInterfaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireguard_interfaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WireguardInterfaceNewResponse struct {
	Data WireguardInterfaceNewResponseData `json:"data"`
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

type WireguardInterfaceNewResponseData struct {
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking bool `json:"enable_sip_trunking"`
	// The Telnyx WireGuard peers `Peer.endpoint` value.
	Endpoint string `json:"endpoint"`
	// The Telnyx WireGuard peers `Peer.PublicKey`.
	PublicKey string `json:"public_key"`
	// Identifies the type of the resource.
	RecordType string                                  `json:"record_type"`
	Region     WireguardInterfaceNewResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableSipTrunking respjson.Field
		Endpoint          respjson.Field
		PublicKey         respjson.Field
		RecordType        respjson.Field
		Region            respjson.Field
		RegionCode        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
	Interface
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceNewResponseDataRegion struct {
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
func (r WireguardInterfaceNewResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceNewResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceGetResponse struct {
	Data WireguardInterfaceGetResponseData `json:"data"`
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

type WireguardInterfaceGetResponseData struct {
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking bool `json:"enable_sip_trunking"`
	// The Telnyx WireGuard peers `Peer.endpoint` value.
	Endpoint string `json:"endpoint"`
	// The Telnyx WireGuard peers `Peer.PublicKey`.
	PublicKey string `json:"public_key"`
	// Identifies the type of the resource.
	RecordType string                                  `json:"record_type"`
	Region     WireguardInterfaceGetResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableSipTrunking respjson.Field
		Endpoint          respjson.Field
		PublicKey         respjson.Field
		RecordType        respjson.Field
		Region            respjson.Field
		RegionCode        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
	Interface
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceGetResponseDataRegion struct {
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
func (r WireguardInterfaceGetResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceGetResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceListResponse struct {
	Data []WireguardInterfaceListResponseData `json:"data"`
	Meta PaginationMeta                       `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceListResponseData struct {
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking bool `json:"enable_sip_trunking"`
	// The Telnyx WireGuard peers `Peer.endpoint` value.
	Endpoint string `json:"endpoint"`
	// The Telnyx WireGuard peers `Peer.PublicKey`.
	PublicKey string `json:"public_key"`
	// Identifies the type of the resource.
	RecordType string                                   `json:"record_type"`
	Region     WireguardInterfaceListResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableSipTrunking respjson.Field
		Endpoint          respjson.Field
		PublicKey         respjson.Field
		RecordType        respjson.Field
		Region            respjson.Field
		RegionCode        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
	Interface
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceListResponseDataRegion struct {
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
func (r WireguardInterfaceListResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceListResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceDeleteResponse struct {
	Data WireguardInterfaceDeleteResponseData `json:"data"`
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

type WireguardInterfaceDeleteResponseData struct {
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking bool `json:"enable_sip_trunking"`
	// The Telnyx WireGuard peers `Peer.endpoint` value.
	Endpoint string `json:"endpoint"`
	// The Telnyx WireGuard peers `Peer.PublicKey`.
	PublicKey string `json:"public_key"`
	// Identifies the type of the resource.
	RecordType string                                     `json:"record_type"`
	Region     WireguardInterfaceDeleteResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableSipTrunking respjson.Field
		Endpoint          respjson.Field
		PublicKey         respjson.Field
		RecordType        respjson.Field
		Region            respjson.Field
		RegionCode        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
	Interface
}

// Returns the unmodified JSON received from the API
func (r WireguardInterfaceDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceDeleteResponseDataRegion struct {
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
func (r WireguardInterfaceDeleteResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *WireguardInterfaceDeleteResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardInterfaceNewParams struct {
	// The id of the network associated with the interface.
	NetworkID string `json:"network_id,required" format:"uuid"`
	// The region the interface should be deployed to.
	RegionCode string `json:"region_code,required"`
	// Enable SIP traffic forwarding over VPN interface.
	EnableSipTrunking param.Opt[bool] `json:"enable_sip_trunking,omitzero"`
	// A user specified name for the interface.
	Name param.Opt[string] `json:"name,omitzero"`
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
	// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
	Filter WireguardInterfaceListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page WireguardInterfaceListParamsPage `query:"page,omitzero" json:"-"`
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

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type WireguardInterfaceListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WireguardInterfaceListParamsPage]'s query parameters as
// `url.Values`.
func (r WireguardInterfaceListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
