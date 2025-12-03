// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// NetworkDefaultGatewayService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkDefaultGatewayService] method instead.
type NetworkDefaultGatewayService struct {
	Options []option.RequestOption
}

// NewNetworkDefaultGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkDefaultGatewayService(opts ...option.RequestOption) (r NetworkDefaultGatewayService) {
	r = NetworkDefaultGatewayService{}
	r.Options = opts
	return
}

// Create Default Gateway.
func (r *NetworkDefaultGatewayService) New(ctx context.Context, id string, body NetworkDefaultGatewayNewParams, opts ...option.RequestOption) (res *NetworkDefaultGatewayNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s/default_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Default Gateway status.
func (r *NetworkDefaultGatewayService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkDefaultGatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s/default_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Default Gateway.
func (r *NetworkDefaultGatewayService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkDefaultGatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s/default_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type NetworkDefaultGatewayNewResponse struct {
	Data []NetworkDefaultGatewayNewResponseData `json:"data"`
	Meta PaginationMeta                         `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkDefaultGatewayNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkDefaultGatewayNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDefaultGatewayNewResponseData struct {
	// Network ID.
	NetworkID string `json:"network_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// Wireguard peer ID.
	WireguardPeerID string `json:"wireguard_peer_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetworkID       respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		WireguardPeerID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkDefaultGatewayNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkDefaultGatewayNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDefaultGatewayGetResponse struct {
	Data []NetworkDefaultGatewayGetResponseData `json:"data"`
	Meta PaginationMeta                         `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkDefaultGatewayGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkDefaultGatewayGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDefaultGatewayGetResponseData struct {
	// Network ID.
	NetworkID string `json:"network_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// Wireguard peer ID.
	WireguardPeerID string `json:"wireguard_peer_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetworkID       respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		WireguardPeerID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkDefaultGatewayGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkDefaultGatewayGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDefaultGatewayDeleteResponse struct {
	Data []NetworkDefaultGatewayDeleteResponseData `json:"data"`
	Meta PaginationMeta                            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkDefaultGatewayDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkDefaultGatewayDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDefaultGatewayDeleteResponseData struct {
	// Network ID.
	NetworkID string `json:"network_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// Wireguard peer ID.
	WireguardPeerID string `json:"wireguard_peer_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetworkID       respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		WireguardPeerID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkDefaultGatewayDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkDefaultGatewayDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDefaultGatewayNewParams struct {
	// Wireguard peer ID.
	WireguardPeerID param.Opt[string] `json:"wireguard_peer_id,omitzero" format:"uuid"`
	paramObj
}

func (r NetworkDefaultGatewayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkDefaultGatewayNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkDefaultGatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
