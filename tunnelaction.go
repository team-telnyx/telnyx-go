// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Tunnels connect a Network to hosts behind a tunnel process within the user's
// network.
//
// TunnelActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelActionService] method instead.
type TunnelActionService struct {
	Options []option.RequestOption
}

// NewTunnelActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTunnelActionService(opts ...option.RequestOption) (r TunnelActionService) {
	r = TunnelActionService{}
	r.Options = opts
	return
}

// Issues a new `tunnel_token` and invalidates the previous one. Tunnel processes
// already connected stay connected until they reconnect, and are then refused
// until they use the new token.
func (r *TunnelActionService) RotateToken(ctx context.Context, id string, opts ...option.RequestOption) (res *TunnelActionRotateTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tunnels/%s/actions/rotate_token", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TunnelActionRotateTokenResponse struct {
	Data TunnelWithToken `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TunnelActionRotateTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *TunnelActionRotateTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
