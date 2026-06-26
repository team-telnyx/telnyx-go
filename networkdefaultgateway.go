// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Network operations
//
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
func (r *NetworkDefaultGatewayService) New(ctx context.Context, networkIdentifier string, body NetworkDefaultGatewayNewParams, opts ...option.RequestOption) (res *NetworkDefaultGatewayNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if networkIdentifier == "" {
		err = errors.New("missing required network_identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("networks/%s/default_gateway", networkIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get Default Gateway status.
func (r *NetworkDefaultGatewayService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkDefaultGatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("networks/%s/default_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete Default Gateway.
func (r *NetworkDefaultGatewayService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkDefaultGatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("networks/%s/default_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type DefaultGateway struct {
	// Network ID.
	NetworkID string `json:"network_id" format:"uuid"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// Wireguard peer ID.
	WireguardPeerID string `json:"wireguard_peer_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NetworkID       respjson.Field
		Status          respjson.Field
		WireguardPeerID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r DefaultGateway) RawJSON() string { return r.JSON.raw }
func (r *DefaultGateway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DefaultGateway to a DefaultGatewayParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DefaultGatewayParam.Overrides()
func (r DefaultGateway) ToParam() DefaultGatewayParam {
	return param.Override[DefaultGatewayParam](json.RawMessage(r.RawJSON()))
}

type DefaultGatewayParam struct {
	// Wireguard peer ID.
	WireguardPeerID param.Opt[string] `json:"wireguard_peer_id,omitzero" format:"uuid"`
	RecordParam
}

func (r DefaultGatewayParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*DefaultGatewayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type NetworkDefaultGatewayNewResponse struct {
	Data []DefaultGateway `json:"data"`
	Meta PaginationMeta   `json:"meta"`
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

type NetworkDefaultGatewayGetResponse struct {
	Data []DefaultGateway `json:"data"`
	Meta PaginationMeta   `json:"meta"`
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

type NetworkDefaultGatewayDeleteResponse struct {
	Data []DefaultGateway `json:"data"`
	Meta PaginationMeta   `json:"meta"`
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

type NetworkDefaultGatewayNewParams struct {
	DefaultGateway DefaultGatewayParam
	paramObj
}

func (r NetworkDefaultGatewayNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DefaultGateway)
}
func (r *NetworkDefaultGatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
