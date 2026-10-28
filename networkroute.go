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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Routes on a Network that send a CIDR to a next hop, such as a tunnel.
//
// NetworkRouteService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkRouteService] method instead.
type NetworkRouteService struct {
	Options []option.RequestOption
}

// NewNetworkRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkRouteService(opts ...option.RequestOption) (r NetworkRouteService) {
	r = NetworkRouteService{}
	r.Options = opts
	return
}

// Adds a route to a Network. Traffic on the Network to an address in `cidr` goes
// via `next_hop`. With a tunnel as the next hop, functions on the Network reach
// the CIDR through the tunnel without a redeploy. A Network holds each CIDR once.
func (r *NetworkRouteService) New(ctx context.Context, body NetworkRouteNewParams, opts ...option.RequestOption) (res *NetworkRouteNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network_routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a route and its status.
func (r *NetworkRouteService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkRouteGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("network_routes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the routes of the account's Networks. A Network's list is its effective
// route table: the routes added with `POST /network_routes` (`origin` `static`)
// and the routes that its attachments install (`origin` `attachment`). Lookup is
// longest prefix within the Network.
func (r *NetworkRouteService) List(ctx context.Context, query NetworkRouteListParams, opts ...option.RequestOption) (res *NetworkRouteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network_routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a route that was added with `POST /network_routes`. Traffic to its CIDR
// stops going via the next hop at once. Routes that an attachment installs
// (`origin` `attachment`) leave with the attachment and cannot be deleted here.
func (r *NetworkRouteService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("network_routes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type NetappsCursorPaginationMeta struct {
	// Opaque cursors. Absent keys mean there is no such page.
	Cursors NetappsCursorPaginationMetaCursors `json:"cursors"`
	// The path of the next page, with the same query parameters. Absent on the last
	// page.
	Next string `json:"next"`
	// The path of the previous page. Present when the request used `page[after]`.
	Previous string `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetappsCursorPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *NetappsCursorPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opaque cursors. Absent keys mean there is no such page.
type NetappsCursorPaginationMetaCursors struct {
	// Pass as `page[after]` to fetch the next page. Absent on the last page.
	After string `json:"after"`
	// The cursor of the previous page. Present when the request used `page[after]`.
	Before string `json:"before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetappsCursorPaginationMetaCursors) RawJSON() string { return r.JSON.raw }
func (r *NetappsCursorPaginationMetaCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRoute struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// The IPv4 CIDR the route covers.
	Cidr string `json:"cidr"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The id of the Network the route belongs to.
	NetworkID string `json:"network_id" format:"uuid"`
	// Where the Network sends traffic for the route's CIDR.
	NextHop NetworkRouteNextHop `json:"next_hop"`
	// `static` for a route added with `POST /network_routes`, `attachment` for a route
	// that another resource installed and removes.
	//
	// Any of "static", "attachment".
	Origin NetworkRouteOrigin `json:"origin"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// `pending` until Telnyx installs the route on the Network, then `active`.
	// `inactive` while a tunnel next hop has no connected process. `failed` (Preview,
	// not reported yet) when Telnyx could not install the route.
	//
	// Any of "pending", "active", "inactive", "failed".
	Status NetworkRouteStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cidr        respjson.Field
		CreatedAt   respjson.Field
		NetworkID   respjson.Field
		NextHop     respjson.Field
		Origin      respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRoute) RawJSON() string { return r.JSON.raw }
func (r *NetworkRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the Network sends traffic for the route's CIDR.
type NetworkRouteNextHop struct {
	// The id of the tunnel or attachment.
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of the next hop. Routes added with `POST /network_routes` use `tunnel`.
	// Routes that attachments install carry the attachment's type.
	//
	// Any of "tunnel", "wireguard_peer", "wireguard_interface",
	// "private_wireless_gateway", "virtual_cross_connect", "ipsec_tunnel",
	// "public_internet_gateway".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouteNextHop) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouteNextHop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `static` for a route added with `POST /network_routes`, `attachment` for a route
// that another resource installed and removes.
type NetworkRouteOrigin string

const (
	NetworkRouteOriginStatic     NetworkRouteOrigin = "static"
	NetworkRouteOriginAttachment NetworkRouteOrigin = "attachment"
)

// `pending` until Telnyx installs the route on the Network, then `active`.
// `inactive` while a tunnel next hop has no connected process. `failed` (Preview,
// not reported yet) when Telnyx could not install the route.
type NetworkRouteStatus string

const (
	NetworkRouteStatusPending  NetworkRouteStatus = "pending"
	NetworkRouteStatusActive   NetworkRouteStatus = "active"
	NetworkRouteStatusInactive NetworkRouteStatus = "inactive"
	NetworkRouteStatusFailed   NetworkRouteStatus = "failed"
)

type NetworkRouteNewResponse struct {
	Data NetworkRoute `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouteNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouteNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouteGetResponse struct {
	Data NetworkRoute `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouteGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouteGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouteListResponse struct {
	Data []NetworkRoute              `json:"data"`
	Meta NetappsCursorPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouteListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouteNewParams struct {
	// The IPv4 subnet to route, such as a subnet behind the tunnel. It must be a
	// canonical CIDR. Default, loopback, link-local, multicast, and shared address
	// space (`100.64.0.0/10`) ranges are rejected. So are ranges that overlap the
	// Tailscale Subnet Router ranges, `172.27.254.0/24` and `198.18.0.0/20`, or a
	// subnet already on the Network.
	Cidr string `json:"cidr" api:"required"`
	// The id of the Network to add the route to.
	NetworkID string `json:"network_id" api:"required" format:"uuid"`
	// Where the Network sends traffic for `cidr`.
	NextHop NetworkRouteNewParamsNextHop `json:"next_hop,omitzero" api:"required"`
	paramObj
}

func (r NetworkRouteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the Network sends traffic for `cidr`.
//
// The properties ID, Type are required.
type NetworkRouteNewParamsNextHop struct {
	// The id of the tunnel.
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of the next hop.
	//
	// Any of "tunnel".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r NetworkRouteNewParamsNextHop) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouteNewParamsNextHop
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouteNewParamsNextHop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[NetworkRouteNewParamsNextHop](
		"type", "tunnel",
	)
}

type NetworkRouteListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[network_id], filter[origin], filter[next_hop.type], filter[cidr],
	// filter[status]
	Filter NetworkRouteListParamsFilter `query:"filter,omitzero" api:"required" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[after],
	// page[limit]
	Page NetworkRouteListParamsPage `query:"page,omitzero" json:"-"`
	// Sorts the results by a field, ascending, or descending with a leading `-`. The
	// default is `-created_at`.
	//
	// Any of "cidr", "-cidr", "status", "-status", "created_at", "-created_at".
	Sort NetworkRouteListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkRouteListParams]'s query parameters as `url.Values`.
func (r NetworkRouteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[network_id], filter[origin], filter[next_hop.type], filter[cidr],
// filter[status]
//
// The property NetworkID is required.
type NetworkRouteListParamsFilter struct {
	// The id of the Network whose route table is returned. Required.
	NetworkID string `query:"network_id" api:"required" format:"uuid" json:"-"`
	// Only routes with this CIDR.
	Cidr param.Opt[string] `query:"cidr,omitzero" json:"-"`
	// Only routes whose next hop has this type.
	//
	// Any of "tunnel", "wireguard_peer", "wireguard_interface",
	// "private_wireless_gateway", "virtual_cross_connect", "ipsec_tunnel",
	// "public_internet_gateway".
	NextHopType string `query:"next_hop.type,omitzero" json:"-"`
	// Only routes of this origin.
	//
	// Any of "static", "attachment".
	Origin string `query:"origin,omitzero" json:"-"`
	// Only routes with this status.
	//
	// Any of "pending", "active", "inactive", "failed".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkRouteListParamsFilter]'s query parameters as
// `url.Values`.
func (r NetworkRouteListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[after],
// page[limit]
type NetworkRouteListParamsPage struct {
	// The cursor from `meta.cursors.after` of the previous page.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The number of results per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkRouteListParamsPage]'s query parameters as
// `url.Values`.
func (r NetworkRouteListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts the results by a field, ascending, or descending with a leading `-`. The
// default is `-created_at`.
type NetworkRouteListParamsSort string

const (
	NetworkRouteListParamsSortCidr          NetworkRouteListParamsSort = "cidr"
	NetworkRouteListParamsSortMinusCidr     NetworkRouteListParamsSort = "-cidr"
	NetworkRouteListParamsSortStatus        NetworkRouteListParamsSort = "status"
	NetworkRouteListParamsSortStatusDesc    NetworkRouteListParamsSort = "-status"
	NetworkRouteListParamsSortCreatedAt     NetworkRouteListParamsSort = "created_at"
	NetworkRouteListParamsSortCreatedAtDesc NetworkRouteListParamsSort = "-created_at"
)
