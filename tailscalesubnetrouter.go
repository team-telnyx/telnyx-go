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

// Tailscale Subnet Router operations
//
// TailscaleSubnetRouterService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTailscaleSubnetRouterService] method instead.
type TailscaleSubnetRouterService struct {
	Options []option.RequestOption
}

// NewTailscaleSubnetRouterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTailscaleSubnetRouterService(opts ...option.RequestOption) (r TailscaleSubnetRouterService) {
	r = TailscaleSubnetRouterService{}
	r.Options = opts
	return
}

// Creates a Tailscale Subnet Router on a Network and enrolls it in the user's
// tailnet with the auth key. The response returns the router with `status`
// `provisioning`. Poll `GET /tailscale_subnet_routers/{id}` until it is
// `provisioned`. Then approve its routes in the tailnet, unless an auto-approver
// covers them.
func (r *TailscaleSubnetRouterService) New(ctx context.Context, body TailscaleSubnetRouterNewParams, opts ...option.RequestOption) (res *TailscaleSubnetRouterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tailscale_subnet_routers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the details of a single Tailscale Subnet Router by its identifier,
// including its enrollment state.
func (r *TailscaleSubnetRouterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TailscaleSubnetRouterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tailscale_subnet_routers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Renames the router, changes its advertised routes, or retries a failed
// enrollment. Any change queues a new provisioning run and returns the router with
// `status` `provisioning`. The router keeps its machine identity on the tailnet.
// Send an empty object to retry a `failed` router with its stored auth key.
// Include a new `auth_key` only while the router has never enrolled. The router's
// Network and region cannot change: create a new router instead.
func (r *TailscaleSubnetRouterService) Update(ctx context.Context, id string, body TailscaleSubnetRouterUpdateParams, opts ...option.RequestOption) (res *TailscaleSubnetRouterUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tailscale_subnet_routers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of the Tailscale Subnet Routers on the account, with
// support for filtering by Network.
func (r *TailscaleSubnetRouterService) List(ctx context.Context, query TailscaleSubnetRouterListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[TailscaleSubnetRouter], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tailscale_subnet_routers"
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

// Returns a paginated list of the Tailscale Subnet Routers on the account, with
// support for filtering by Network.
func (r *TailscaleSubnetRouterService) ListAutoPaging(ctx context.Context, query TailscaleSubnetRouterListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[TailscaleSubnetRouter] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Removes the router from the Network and logs it out of the tailnet. The response
// returns the router with `status` `deleting`. The record disappears once removal
// completes. A machine entry can remain in the tailnet's admin console until it is
// removed there.
func (r *TailscaleSubnetRouterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TailscaleSubnetRouterDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tailscale_subnet_routers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type TailscaleSubnetRouter struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// IPv4 subnets of the Network that the router advertises to the tailnet. Each
	// entry is a canonical CIDR that the Network routes. Default, loopback,
	// link-local, multicast, and Tailscale (`100.64.0.0/10`) ranges are rejected. So
	// are ranges that overlap the router's own, `172.27.254.0/24` and `198.18.0.0/20`.
	AdvertisedRoutes []string `json:"advertised_routes"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The router's address on the Network, from `172.27.254.0/24`. Tailnet traffic
	// enters the Network from this address.
	CustomerAddress string `json:"customer_address"`
	// The router's node id on the tailnet. Set once the router has enrolled.
	DeviceID string `json:"device_id" api:"nullable"`
	// A user specified name for the router.
	Name string `json:"name"`
	// The id of the Network the router is attached to.
	NetworkID string `json:"network_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The code of the region the router is created in. `GET /network_coverage` lists
	// the regions whose `available_services` include `tailscale_subnet_router`.
	RegionCode string `json:"region_code"`
	// `provisioning`: Telnyx is creating the router or applying a change.
	// `provisioned`: the router is on the tailnet and advertises its routes; the
	// tailnet still approves them. `failed`: enrollment did not complete, and `PATCH`
	// retries it. `deleting`: the router is being removed.
	//
	// Any of "provisioning", "provisioned", "failed", "deleting".
	Status TailscaleSubnetRouterStatus `json:"status"`
	// The name of the tailnet the router is enrolled in. Set once the router has
	// enrolled.
	TailnetName string `json:"tailnet_name" api:"nullable"`
	// The router's addresses on the tailnet. Empty until the router has enrolled.
	TailscaleIPs []string `json:"tailscale_ips"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AdvertisedRoutes respjson.Field
		CreatedAt        respjson.Field
		CustomerAddress  respjson.Field
		DeviceID         respjson.Field
		Name             respjson.Field
		NetworkID        respjson.Field
		RecordType       respjson.Field
		RegionCode       respjson.Field
		Status           respjson.Field
		TailnetName      respjson.Field
		TailscaleIPs     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TailscaleSubnetRouter) RawJSON() string { return r.JSON.raw }
func (r *TailscaleSubnetRouter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `provisioning`: Telnyx is creating the router or applying a change.
// `provisioned`: the router is on the tailnet and advertises its routes; the
// tailnet still approves them. `failed`: enrollment did not complete, and `PATCH`
// retries it. `deleting`: the router is being removed.
type TailscaleSubnetRouterStatus string

const (
	TailscaleSubnetRouterStatusProvisioning TailscaleSubnetRouterStatus = "provisioning"
	TailscaleSubnetRouterStatusProvisioned  TailscaleSubnetRouterStatus = "provisioned"
	TailscaleSubnetRouterStatusFailed       TailscaleSubnetRouterStatus = "failed"
	TailscaleSubnetRouterStatusDeleting     TailscaleSubnetRouterStatus = "deleting"
)

type TailscaleSubnetRouterNewResponse struct {
	Data TailscaleSubnetRouter `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TailscaleSubnetRouterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TailscaleSubnetRouterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TailscaleSubnetRouterGetResponse struct {
	Data TailscaleSubnetRouter `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TailscaleSubnetRouterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TailscaleSubnetRouterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TailscaleSubnetRouterUpdateResponse struct {
	Data TailscaleSubnetRouter `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TailscaleSubnetRouterUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TailscaleSubnetRouterUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TailscaleSubnetRouterDeleteResponse struct {
	Data TailscaleSubnetRouter `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TailscaleSubnetRouterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TailscaleSubnetRouterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TailscaleSubnetRouterNewParams struct {
	// IPv4 subnets of the Network that the router advertises to the tailnet. Each
	// entry is a canonical CIDR that the Network routes. Default, loopback,
	// link-local, multicast, and Tailscale (`100.64.0.0/10`) ranges are rejected. So
	// are ranges that overlap the router's own, `172.27.254.0/24` and `198.18.0.0/20`.
	AdvertisedRoutes []string `json:"advertised_routes,omitzero" api:"required"`
	// A Tailscale auth key that enrolls the router in the user's tailnet: reusable,
	// pre-approved, tagged, and not ephemeral. Telnyx stores it encrypted until
	// enrollment succeeds, then discards it. It is never returned.
	AuthKey string `json:"auth_key" api:"required" format:"password"`
	// A user specified name for the router.
	Name string `json:"name" api:"required"`
	// The id of the Network to attach the router to.
	NetworkID string `json:"network_id" api:"required" format:"uuid"`
	// The code of the region to create the router in. `GET /network_coverage` lists
	// the regions whose `available_services` include `tailscale_subnet_router`.
	RegionCode string `json:"region_code" api:"required"`
	paramObj
}

func (r TailscaleSubnetRouterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TailscaleSubnetRouterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TailscaleSubnetRouterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TailscaleSubnetRouterUpdateParams struct {
	// A replacement Tailscale auth key. Accepted only while the router has never
	// enrolled, for example after a `failed` enrollment. Never returned.
	AuthKey param.Opt[string] `json:"auth_key,omitzero" format:"password"`
	// A user specified name for the router.
	Name param.Opt[string] `json:"name,omitzero"`
	// IPv4 subnets of the Network that the router advertises to the tailnet. Each
	// entry is a canonical CIDR that the Network routes. Default, loopback,
	// link-local, multicast, and Tailscale (`100.64.0.0/10`) ranges are rejected. So
	// are ranges that overlap the router's own, `172.27.254.0/24` and `198.18.0.0/20`.
	AdvertisedRoutes []string `json:"advertised_routes,omitzero"`
	paramObj
}

func (r TailscaleSubnetRouterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TailscaleSubnetRouterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TailscaleSubnetRouterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TailscaleSubnetRouterListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
	Filter TailscaleSubnetRouterListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TailscaleSubnetRouterListParams]'s query parameters as
// `url.Values`.
func (r TailscaleSubnetRouterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
type TailscaleSubnetRouterListParamsFilter struct {
	// The id of the Network to filter on.
	NetworkID param.Opt[string] `query:"network_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TailscaleSubnetRouterListParamsFilter]'s query parameters
// as `url.Values`.
func (r TailscaleSubnetRouterListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
