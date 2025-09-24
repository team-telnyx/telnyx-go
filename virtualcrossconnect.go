// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// VirtualCrossConnectService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVirtualCrossConnectService] method instead.
type VirtualCrossConnectService struct {
	Options []option.RequestOption
}

// NewVirtualCrossConnectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVirtualCrossConnectService(opts ...option.RequestOption) (r VirtualCrossConnectService) {
	r = VirtualCrossConnectService{}
	r.Options = opts
	return
}

// Create a new Virtual Cross Connect.<br /><br />For AWS and GCE, you have the
// option of creating the primary connection first and the secondary connection
// later. You also have the option of disabling the primary and/or secondary
// connections at any time and later re-enabling them. With Azure, you do not have
// this option. Azure requires both the primary and secondary connections to be
// created at the same time and they can not be independantly disabled.
func (r *VirtualCrossConnectService) New(ctx context.Context, body VirtualCrossConnectNewParams, opts ...option.RequestOption) (res *VirtualCrossConnectNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "virtual_cross_connects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Virtual Cross Connect.
func (r *VirtualCrossConnectService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VirtualCrossConnectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("virtual_cross_connects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Virtual Cross Connect.<br /><br />Cloud IPs can only be patched
// during the `created` state, as GCE will only inform you of your generated IP
// once the pending connection requested has been accepted. Once the Virtual Cross
// Connect has moved to `provisioning`, the IPs can no longer be
// patched.<br /><br />Once the Virtual Cross Connect has moved to `provisioned`
// and you are ready to enable routing, you can toggle the routing announcements to
// `true`.
func (r *VirtualCrossConnectService) Update(ctx context.Context, id string, body VirtualCrossConnectUpdateParams, opts ...option.RequestOption) (res *VirtualCrossConnectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("virtual_cross_connects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Virtual Cross Connects.
func (r *VirtualCrossConnectService) List(ctx context.Context, query VirtualCrossConnectListParams, opts ...option.RequestOption) (res *VirtualCrossConnectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "virtual_cross_connects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Virtual Cross Connect.
func (r *VirtualCrossConnectService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *VirtualCrossConnectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("virtual_cross_connects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type VirtualCrossConnectNewResponse struct {
	Data VirtualCrossConnectNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectNewResponseData struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value
	// will be assigned by Telnyx.
	BgpAsn float64 `json:"bgp_asn,required"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `json:"cloud_provider,required"`
	// The region where your Virtual Private Cloud hosts are located.<br /><br />The
	// available regions can be found using the /virtual_cross_connect_regions
	// endpoint.
	CloudProviderRegion string `json:"cloud_provider_region,required"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.
	PrimaryCloudAccountID string `json:"primary_cloud_account_id,required"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code,required"`
	// The desired throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.<br /><br />The available bandwidths can be found using the
	// /virtual_cross_connect_regions endpoint.
	BandwidthMbps float64 `json:"bandwidth_mbps"`
	// The authentication key for BGP peer configuration.
	PrimaryBgpKey string `json:"primary_bgp_key"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	PrimaryCloudIP string `json:"primary_cloud_ip"`
	// Indicates whether the primary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	PrimaryEnabled bool `json:"primary_enabled"`
	// Whether the primary BGP route is being announced.
	PrimaryRoutingAnnouncement bool `json:"primary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryTelnyxIP string `json:"primary_telnyx_ip"`
	// Identifies the type of the resource.
	RecordType string                                   `json:"record_type"`
	Region     VirtualCrossConnectNewResponseDataRegion `json:"region"`
	// The authentication key for BGP peer configuration.
	SecondaryBgpKey string `json:"secondary_bgp_key"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.<br /><br />This attribute is only necessary for
	// GCE.
	SecondaryCloudAccountID string `json:"secondary_cloud_account_id"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	SecondaryCloudIP string `json:"secondary_cloud_ip"`
	// Indicates whether the secondary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	SecondaryEnabled bool `json:"secondary_enabled"`
	// Whether the secondary BGP route is being announced.
	SecondaryRoutingAnnouncement bool `json:"secondary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryTelnyxIP string `json:"secondary_telnyx_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgpAsn                       respjson.Field
		CloudProvider                respjson.Field
		CloudProviderRegion          respjson.Field
		PrimaryCloudAccountID        respjson.Field
		RegionCode                   respjson.Field
		BandwidthMbps                respjson.Field
		PrimaryBgpKey                respjson.Field
		PrimaryCloudIP               respjson.Field
		PrimaryEnabled               respjson.Field
		PrimaryRoutingAnnouncement   respjson.Field
		PrimaryTelnyxIP              respjson.Field
		RecordType                   respjson.Field
		Region                       respjson.Field
		SecondaryBgpKey              respjson.Field
		SecondaryCloudAccountID      respjson.Field
		SecondaryCloudIP             respjson.Field
		SecondaryEnabled             respjson.Field
		SecondaryRoutingAnnouncement respjson.Field
		SecondaryTelnyxIP            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
	Record
	Interface
	RegionIn
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectNewResponseDataRegion struct {
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
func (r VirtualCrossConnectNewResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectNewResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectGetResponse struct {
	Data VirtualCrossConnectGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectGetResponseData struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value
	// will be assigned by Telnyx.
	BgpAsn float64 `json:"bgp_asn,required"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `json:"cloud_provider,required"`
	// The region where your Virtual Private Cloud hosts are located.<br /><br />The
	// available regions can be found using the /virtual_cross_connect_regions
	// endpoint.
	CloudProviderRegion string `json:"cloud_provider_region,required"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.
	PrimaryCloudAccountID string `json:"primary_cloud_account_id,required"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code,required"`
	// The desired throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.<br /><br />The available bandwidths can be found using the
	// /virtual_cross_connect_regions endpoint.
	BandwidthMbps float64 `json:"bandwidth_mbps"`
	// The authentication key for BGP peer configuration.
	PrimaryBgpKey string `json:"primary_bgp_key"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	PrimaryCloudIP string `json:"primary_cloud_ip"`
	// Indicates whether the primary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	PrimaryEnabled bool `json:"primary_enabled"`
	// Whether the primary BGP route is being announced.
	PrimaryRoutingAnnouncement bool `json:"primary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryTelnyxIP string `json:"primary_telnyx_ip"`
	// Identifies the type of the resource.
	RecordType string                                   `json:"record_type"`
	Region     VirtualCrossConnectGetResponseDataRegion `json:"region"`
	// The authentication key for BGP peer configuration.
	SecondaryBgpKey string `json:"secondary_bgp_key"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.<br /><br />This attribute is only necessary for
	// GCE.
	SecondaryCloudAccountID string `json:"secondary_cloud_account_id"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	SecondaryCloudIP string `json:"secondary_cloud_ip"`
	// Indicates whether the secondary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	SecondaryEnabled bool `json:"secondary_enabled"`
	// Whether the secondary BGP route is being announced.
	SecondaryRoutingAnnouncement bool `json:"secondary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryTelnyxIP string `json:"secondary_telnyx_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgpAsn                       respjson.Field
		CloudProvider                respjson.Field
		CloudProviderRegion          respjson.Field
		PrimaryCloudAccountID        respjson.Field
		RegionCode                   respjson.Field
		BandwidthMbps                respjson.Field
		PrimaryBgpKey                respjson.Field
		PrimaryCloudIP               respjson.Field
		PrimaryEnabled               respjson.Field
		PrimaryRoutingAnnouncement   respjson.Field
		PrimaryTelnyxIP              respjson.Field
		RecordType                   respjson.Field
		Region                       respjson.Field
		SecondaryBgpKey              respjson.Field
		SecondaryCloudAccountID      respjson.Field
		SecondaryCloudIP             respjson.Field
		SecondaryEnabled             respjson.Field
		SecondaryRoutingAnnouncement respjson.Field
		SecondaryTelnyxIP            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
	Record
	Interface
	RegionIn
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectGetResponseDataRegion struct {
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
func (r VirtualCrossConnectGetResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectGetResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectUpdateResponse struct {
	Data VirtualCrossConnectUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectUpdateResponseData struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value
	// will be assigned by Telnyx.
	BgpAsn float64 `json:"bgp_asn,required"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `json:"cloud_provider,required"`
	// The region where your Virtual Private Cloud hosts are located.<br /><br />The
	// available regions can be found using the /virtual_cross_connect_regions
	// endpoint.
	CloudProviderRegion string `json:"cloud_provider_region,required"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.
	PrimaryCloudAccountID string `json:"primary_cloud_account_id,required"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code,required"`
	// The desired throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.<br /><br />The available bandwidths can be found using the
	// /virtual_cross_connect_regions endpoint.
	BandwidthMbps float64 `json:"bandwidth_mbps"`
	// The authentication key for BGP peer configuration.
	PrimaryBgpKey string `json:"primary_bgp_key"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	PrimaryCloudIP string `json:"primary_cloud_ip"`
	// Indicates whether the primary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	PrimaryEnabled bool `json:"primary_enabled"`
	// Whether the primary BGP route is being announced.
	PrimaryRoutingAnnouncement bool `json:"primary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryTelnyxIP string `json:"primary_telnyx_ip"`
	// Identifies the type of the resource.
	RecordType string                                      `json:"record_type"`
	Region     VirtualCrossConnectUpdateResponseDataRegion `json:"region"`
	// The authentication key for BGP peer configuration.
	SecondaryBgpKey string `json:"secondary_bgp_key"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.<br /><br />This attribute is only necessary for
	// GCE.
	SecondaryCloudAccountID string `json:"secondary_cloud_account_id"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	SecondaryCloudIP string `json:"secondary_cloud_ip"`
	// Indicates whether the secondary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	SecondaryEnabled bool `json:"secondary_enabled"`
	// Whether the secondary BGP route is being announced.
	SecondaryRoutingAnnouncement bool `json:"secondary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryTelnyxIP string `json:"secondary_telnyx_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgpAsn                       respjson.Field
		CloudProvider                respjson.Field
		CloudProviderRegion          respjson.Field
		PrimaryCloudAccountID        respjson.Field
		RegionCode                   respjson.Field
		BandwidthMbps                respjson.Field
		PrimaryBgpKey                respjson.Field
		PrimaryCloudIP               respjson.Field
		PrimaryEnabled               respjson.Field
		PrimaryRoutingAnnouncement   respjson.Field
		PrimaryTelnyxIP              respjson.Field
		RecordType                   respjson.Field
		Region                       respjson.Field
		SecondaryBgpKey              respjson.Field
		SecondaryCloudAccountID      respjson.Field
		SecondaryCloudIP             respjson.Field
		SecondaryEnabled             respjson.Field
		SecondaryRoutingAnnouncement respjson.Field
		SecondaryTelnyxIP            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
	Record
	Interface
	RegionIn
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectUpdateResponseDataRegion struct {
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
func (r VirtualCrossConnectUpdateResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectUpdateResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectListResponse struct {
	Data []VirtualCrossConnectListResponseData `json:"data"`
	Meta PaginationMeta                        `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectListResponse) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectListResponseData struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value
	// will be assigned by Telnyx.
	BgpAsn float64 `json:"bgp_asn,required"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `json:"cloud_provider,required"`
	// The region where your Virtual Private Cloud hosts are located.<br /><br />The
	// available regions can be found using the /virtual_cross_connect_regions
	// endpoint.
	CloudProviderRegion string `json:"cloud_provider_region,required"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.
	PrimaryCloudAccountID string `json:"primary_cloud_account_id,required"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code,required"`
	// The desired throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.<br /><br />The available bandwidths can be found using the
	// /virtual_cross_connect_regions endpoint.
	BandwidthMbps float64 `json:"bandwidth_mbps"`
	// The authentication key for BGP peer configuration.
	PrimaryBgpKey string `json:"primary_bgp_key"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	PrimaryCloudIP string `json:"primary_cloud_ip"`
	// Indicates whether the primary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	PrimaryEnabled bool `json:"primary_enabled"`
	// Whether the primary BGP route is being announced.
	PrimaryRoutingAnnouncement bool `json:"primary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryTelnyxIP string `json:"primary_telnyx_ip"`
	// Identifies the type of the resource.
	RecordType string                                    `json:"record_type"`
	Region     VirtualCrossConnectListResponseDataRegion `json:"region"`
	// The authentication key for BGP peer configuration.
	SecondaryBgpKey string `json:"secondary_bgp_key"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.<br /><br />This attribute is only necessary for
	// GCE.
	SecondaryCloudAccountID string `json:"secondary_cloud_account_id"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	SecondaryCloudIP string `json:"secondary_cloud_ip"`
	// Indicates whether the secondary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	SecondaryEnabled bool `json:"secondary_enabled"`
	// Whether the secondary BGP route is being announced.
	SecondaryRoutingAnnouncement bool `json:"secondary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryTelnyxIP string `json:"secondary_telnyx_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgpAsn                       respjson.Field
		CloudProvider                respjson.Field
		CloudProviderRegion          respjson.Field
		PrimaryCloudAccountID        respjson.Field
		RegionCode                   respjson.Field
		BandwidthMbps                respjson.Field
		PrimaryBgpKey                respjson.Field
		PrimaryCloudIP               respjson.Field
		PrimaryEnabled               respjson.Field
		PrimaryRoutingAnnouncement   respjson.Field
		PrimaryTelnyxIP              respjson.Field
		RecordType                   respjson.Field
		Region                       respjson.Field
		SecondaryBgpKey              respjson.Field
		SecondaryCloudAccountID      respjson.Field
		SecondaryCloudIP             respjson.Field
		SecondaryEnabled             respjson.Field
		SecondaryRoutingAnnouncement respjson.Field
		SecondaryTelnyxIP            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
	Record
	Interface
	RegionIn
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectListResponseData) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectListResponseDataRegion struct {
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
func (r VirtualCrossConnectListResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectListResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectDeleteResponse struct {
	Data VirtualCrossConnectDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectDeleteResponseData struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value
	// will be assigned by Telnyx.
	BgpAsn float64 `json:"bgp_asn,required"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `json:"cloud_provider,required"`
	// The region where your Virtual Private Cloud hosts are located.<br /><br />The
	// available regions can be found using the /virtual_cross_connect_regions
	// endpoint.
	CloudProviderRegion string `json:"cloud_provider_region,required"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.
	PrimaryCloudAccountID string `json:"primary_cloud_account_id,required"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code,required"`
	// The desired throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.<br /><br />The available bandwidths can be found using the
	// /virtual_cross_connect_regions endpoint.
	BandwidthMbps float64 `json:"bandwidth_mbps"`
	// The authentication key for BGP peer configuration.
	PrimaryBgpKey string `json:"primary_bgp_key"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	PrimaryCloudIP string `json:"primary_cloud_ip"`
	// Indicates whether the primary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	PrimaryEnabled bool `json:"primary_enabled"`
	// Whether the primary BGP route is being announced.
	PrimaryRoutingAnnouncement bool `json:"primary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryTelnyxIP string `json:"primary_telnyx_ip"`
	// Identifies the type of the resource.
	RecordType string                                      `json:"record_type"`
	Region     VirtualCrossConnectDeleteResponseDataRegion `json:"region"`
	// The authentication key for BGP peer configuration.
	SecondaryBgpKey string `json:"secondary_bgp_key"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.<br /><br />This attribute is only necessary for
	// GCE.
	SecondaryCloudAccountID string `json:"secondary_cloud_account_id"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	SecondaryCloudIP string `json:"secondary_cloud_ip"`
	// Indicates whether the secondary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	SecondaryEnabled bool `json:"secondary_enabled"`
	// Whether the secondary BGP route is being announced.
	SecondaryRoutingAnnouncement bool `json:"secondary_routing_announcement"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryTelnyxIP string `json:"secondary_telnyx_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgpAsn                       respjson.Field
		CloudProvider                respjson.Field
		CloudProviderRegion          respjson.Field
		PrimaryCloudAccountID        respjson.Field
		RegionCode                   respjson.Field
		BandwidthMbps                respjson.Field
		PrimaryBgpKey                respjson.Field
		PrimaryCloudIP               respjson.Field
		PrimaryEnabled               respjson.Field
		PrimaryRoutingAnnouncement   respjson.Field
		PrimaryTelnyxIP              respjson.Field
		RecordType                   respjson.Field
		Region                       respjson.Field
		SecondaryBgpKey              respjson.Field
		SecondaryCloudAccountID      respjson.Field
		SecondaryCloudIP             respjson.Field
		SecondaryEnabled             respjson.Field
		SecondaryRoutingAnnouncement respjson.Field
		SecondaryTelnyxIP            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
	Record
	Interface
	RegionIn
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectDeleteResponseDataRegion struct {
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
func (r VirtualCrossConnectDeleteResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectDeleteResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectNewParams struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value
	// will be assigned by Telnyx.
	BgpAsn float64 `json:"bgp_asn,required"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider VirtualCrossConnectNewParamsCloudProvider `json:"cloud_provider,omitzero,required"`
	// The region where your Virtual Private Cloud hosts are located.<br /><br />The
	// available regions can be found using the /virtual_cross_connect_regions
	// endpoint.
	CloudProviderRegion string `json:"cloud_provider_region,required"`
	// The id of the network associated with the interface.
	NetworkID string `json:"network_id,required" format:"uuid"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.
	PrimaryCloudAccountID string `json:"primary_cloud_account_id,required"`
	// The region the interface should be deployed to.
	RegionCode string `json:"region_code,required"`
	// The desired throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.<br /><br />The available bandwidths can be found using the
	// /virtual_cross_connect_regions endpoint.
	BandwidthMbps param.Opt[float64] `json:"bandwidth_mbps,omitzero"`
	// A user specified name for the interface.
	Name param.Opt[string] `json:"name,omitzero"`
	// The authentication key for BGP peer configuration.
	PrimaryBgpKey param.Opt[string] `json:"primary_bgp_key,omitzero"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryCloudIP param.Opt[string] `json:"primary_cloud_ip,omitzero"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	PrimaryTelnyxIP param.Opt[string] `json:"primary_telnyx_ip,omitzero"`
	// The authentication key for BGP peer configuration.
	SecondaryBgpKey param.Opt[string] `json:"secondary_bgp_key,omitzero"`
	// The identifier for your Virtual Private Cloud. The number will be different
	// based upon your Cloud provider.<br /><br />This attribute is only necessary for
	// GCE.
	SecondaryCloudAccountID param.Opt[string] `json:"secondary_cloud_account_id,omitzero"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryCloudIP param.Opt[string] `json:"secondary_cloud_ip,omitzero"`
	// The IP address assigned to the Telnyx side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value should be null for GCE as Google will only inform you
	// of your assigned IP once the connection has been accepted.
	SecondaryTelnyxIP param.Opt[string] `json:"secondary_telnyx_ip,omitzero"`
	paramObj
}

func (r VirtualCrossConnectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VirtualCrossConnectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VirtualCrossConnectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Virtual Private Cloud with which you would like to establish a cross
// connect.
type VirtualCrossConnectNewParamsCloudProvider string

const (
	VirtualCrossConnectNewParamsCloudProviderAws   VirtualCrossConnectNewParamsCloudProvider = "aws"
	VirtualCrossConnectNewParamsCloudProviderAzure VirtualCrossConnectNewParamsCloudProvider = "azure"
	VirtualCrossConnectNewParamsCloudProviderGce   VirtualCrossConnectNewParamsCloudProvider = "gce"
)

type VirtualCrossConnectUpdateParams struct {
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	PrimaryCloudIP param.Opt[string] `json:"primary_cloud_ip,omitzero"`
	// Indicates whether the primary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	PrimaryEnabled param.Opt[bool] `json:"primary_enabled,omitzero"`
	// Whether the primary BGP route is being announced.
	PrimaryRoutingAnnouncement param.Opt[bool] `json:"primary_routing_announcement,omitzero"`
	// The IP address assigned for your side of the Virtual Cross
	// Connect.<br /><br />If none is provided, one will be generated for
	// you.<br /><br />This value can not be patched once the VXC has bene provisioned.
	SecondaryCloudIP param.Opt[string] `json:"secondary_cloud_ip,omitzero"`
	// Indicates whether the secondary circuit is enabled. Setting this to `false` will
	// disable the circuit.
	SecondaryEnabled param.Opt[bool] `json:"secondary_enabled,omitzero"`
	// Whether the secondary BGP route is being announced.
	SecondaryRoutingAnnouncement param.Opt[bool] `json:"secondary_routing_announcement,omitzero"`
	paramObj
}

func (r VirtualCrossConnectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VirtualCrossConnectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VirtualCrossConnectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
	Filter VirtualCrossConnectListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page VirtualCrossConnectListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectListParams]'s query parameters as
// `url.Values`.
func (r VirtualCrossConnectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
type VirtualCrossConnectListParamsFilter struct {
	// The associated network id to filter on.
	NetworkID param.Opt[string] `query:"network_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectListParamsFilter]'s query parameters as
// `url.Values`.
func (r VirtualCrossConnectListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type VirtualCrossConnectListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectListParamsPage]'s query parameters as
// `url.Values`.
func (r VirtualCrossConnectListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
