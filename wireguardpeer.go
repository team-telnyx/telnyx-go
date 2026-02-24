// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// WireguardPeerService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWireguardPeerService] method instead.
type WireguardPeerService struct {
	Options []option.RequestOption
}

// NewWireguardPeerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWireguardPeerService(opts ...option.RequestOption) (r WireguardPeerService) {
	r = WireguardPeerService{}
	r.Options = opts
	return
}

// Create a new WireGuard Peer. Current limitation of 5 peers per interface can be
// created.
func (r *WireguardPeerService) New(ctx context.Context, body WireguardPeerNewParams, opts ...option.RequestOption) (res *WireguardPeerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireguard_peers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the WireGuard peer.
func (r *WireguardPeerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WireguardPeerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireguard_peers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the WireGuard peer.
func (r *WireguardPeerService) Update(ctx context.Context, id string, body WireguardPeerUpdateParams, opts ...option.RequestOption) (res *WireguardPeerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireguard_peers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all WireGuard peers.
func (r *WireguardPeerService) List(ctx context.Context, query WireguardPeerListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WireguardPeerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "wireguard_peers"
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

// List all WireGuard peers.
func (r *WireguardPeerService) ListAutoPaging(ctx context.Context, query WireguardPeerListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WireguardPeerListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete the WireGuard peer.
func (r *WireguardPeerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WireguardPeerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireguard_peers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve Wireguard config template for Peer
func (r *WireguardPeerService) GetConfig(ctx context.Context, id string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireguard_peers/%s/config", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WireguardPeerPatch struct {
	// The WireGuard `PublicKey`.<br /><br />If you do not provide a Public Key, a new
	// Public and Private key pair will be generated for you.
	PublicKey string `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerPatch) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerPatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WireguardPeerPatch to a WireguardPeerPatchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WireguardPeerPatchParam.Overrides()
func (r WireguardPeerPatch) ToParam() WireguardPeerPatchParam {
	return param.Override[WireguardPeerPatchParam](json.RawMessage(r.RawJSON()))
}

type WireguardPeerPatchParam struct {
	// The WireGuard `PublicKey`.<br /><br />If you do not provide a Public Key, a new
	// Public and Private key pair will be generated for you.
	PublicKey param.Opt[string] `json:"public_key,omitzero"`
	paramObj
}

func (r WireguardPeerPatchParam) MarshalJSON() (data []byte, err error) {
	type shadow WireguardPeerPatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WireguardPeerPatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerNewResponse struct {
	Data WireguardPeerNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerNewResponseData struct {
	// ISO 8601 formatted date-time indicating when peer sent traffic last time.
	LastSeen string `json:"last_seen"`
	// Your WireGuard `Interface.PrivateKey`.<br /><br />This attribute is only ever
	// utlised if, on POST, you do NOT provide your own `public_key`. In which case, a
	// new Public and Private key pair will be generated for you. When your
	// `private_key` is returned, you must save this immediately as we do not save it
	// within Telnyx. If you lose your Private Key, it can not be recovered.
	PrivateKey string `json:"private_key"`
	// The id of the wireguard interface associated with the peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastSeen             respjson.Field
		PrivateKey           respjson.Field
		WireguardInterfaceID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	Record
	WireguardPeerPatch
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerGetResponse struct {
	Data WireguardPeerGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerGetResponseData struct {
	// ISO 8601 formatted date-time indicating when peer sent traffic last time.
	LastSeen string `json:"last_seen"`
	// Your WireGuard `Interface.PrivateKey`.<br /><br />This attribute is only ever
	// utlised if, on POST, you do NOT provide your own `public_key`. In which case, a
	// new Public and Private key pair will be generated for you. When your
	// `private_key` is returned, you must save this immediately as we do not save it
	// within Telnyx. If you lose your Private Key, it can not be recovered.
	PrivateKey string `json:"private_key"`
	// The id of the wireguard interface associated with the peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastSeen             respjson.Field
		PrivateKey           respjson.Field
		WireguardInterfaceID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	Record
	WireguardPeerPatch
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerUpdateResponse struct {
	Data WireguardPeerUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerUpdateResponseData struct {
	// ISO 8601 formatted date-time indicating when peer sent traffic last time.
	LastSeen string `json:"last_seen"`
	// Your WireGuard `Interface.PrivateKey`.<br /><br />This attribute is only ever
	// utlised if, on POST, you do NOT provide your own `public_key`. In which case, a
	// new Public and Private key pair will be generated for you. When your
	// `private_key` is returned, you must save this immediately as we do not save it
	// within Telnyx. If you lose your Private Key, it can not be recovered.
	PrivateKey string `json:"private_key"`
	// The id of the wireguard interface associated with the peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastSeen             respjson.Field
		PrivateKey           respjson.Field
		WireguardInterfaceID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	Record
	WireguardPeerPatch
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerListResponse struct {
	// ISO 8601 formatted date-time indicating when peer sent traffic last time.
	LastSeen string `json:"last_seen"`
	// Your WireGuard `Interface.PrivateKey`.<br /><br />This attribute is only ever
	// utlised if, on POST, you do NOT provide your own `public_key`. In which case, a
	// new Public and Private key pair will be generated for you. When your
	// `private_key` is returned, you must save this immediately as we do not save it
	// within Telnyx. If you lose your Private Key, it can not be recovered.
	PrivateKey string `json:"private_key"`
	// The id of the wireguard interface associated with the peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastSeen             respjson.Field
		PrivateKey           respjson.Field
		WireguardInterfaceID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	Record
	WireguardPeerPatch
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerListResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerDeleteResponse struct {
	Data WireguardPeerDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerDeleteResponseData struct {
	// ISO 8601 formatted date-time indicating when peer sent traffic last time.
	LastSeen string `json:"last_seen"`
	// Your WireGuard `Interface.PrivateKey`.<br /><br />This attribute is only ever
	// utlised if, on POST, you do NOT provide your own `public_key`. In which case, a
	// new Public and Private key pair will be generated for you. When your
	// `private_key` is returned, you must save this immediately as we do not save it
	// within Telnyx. If you lose your Private Key, it can not be recovered.
	PrivateKey string `json:"private_key"`
	// The id of the wireguard interface associated with the peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastSeen             respjson.Field
		PrivateKey           respjson.Field
		WireguardInterfaceID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	Record
	WireguardPeerPatch
}

// Returns the unmodified JSON received from the API
func (r WireguardPeerDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *WireguardPeerDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerNewParams struct {
	// The id of the wireguard interface associated with the peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" api:"required" format:"uuid"`
	// The WireGuard `PublicKey`.<br /><br />If you do not provide a Public Key, a new
	// Public and Private key pair will be generated for you.
	PublicKey param.Opt[string] `json:"public_key,omitzero"`
	paramObj
}

func (r WireguardPeerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WireguardPeerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WireguardPeerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardPeerUpdateParams struct {
	WireguardPeerPatch WireguardPeerPatchParam
	paramObj
}

func (r WireguardPeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WireguardPeerPatch)
}
func (r *WireguardPeerUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WireguardPeerPatch)
}

type WireguardPeerListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[wireguard_interface_id]
	Filter WireguardPeerListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WireguardPeerListParams]'s query parameters as
// `url.Values`.
func (r WireguardPeerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[wireguard_interface_id]
type WireguardPeerListParamsFilter struct {
	// The id of the associated WireGuard interface to filter on.
	WireguardInterfaceID param.Opt[string] `query:"wireguard_interface_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WireguardPeerListParamsFilter]'s query parameters as
// `url.Values`.
func (r WireguardPeerListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
