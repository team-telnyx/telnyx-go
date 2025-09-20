// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// GlobalIPAllowedPortService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPAllowedPortService] method instead.
type GlobalIPAllowedPortService struct {
	Options []option.RequestOption
}

// NewGlobalIPAllowedPortService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalIPAllowedPortService(opts ...option.RequestOption) (r GlobalIPAllowedPortService) {
	r = GlobalIPAllowedPortService{}
	r.Options = opts
	return
}

// List all Global IP Allowed Ports
func (r *GlobalIPAllowedPortService) List(ctx context.Context, opts ...option.RequestOption) (res *GlobalIPAllowedPortListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ip_allowed_ports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GlobalIPAllowedPortListResponse struct {
	Data []GlobalIPAllowedPortListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAllowedPortListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAllowedPortListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAllowedPortListResponseData struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// First port of a range.
	FirstPort int64 `json:"first_port"`
	// Last port of a range.
	LastPort int64 `json:"last_port"`
	// A name for the Global IP ports range.
	Name string `json:"name"`
	// The Global IP Protocol code.
	ProtocolCode string `json:"protocol_code"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		FirstPort    respjson.Field
		LastPort     respjson.Field
		Name         respjson.Field
		ProtocolCode respjson.Field
		RecordType   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAllowedPortListResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAllowedPortListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
