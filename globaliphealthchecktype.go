// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Global IPs
//
// GlobalIPHealthCheckTypeService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPHealthCheckTypeService] method instead.
type GlobalIPHealthCheckTypeService struct {
	Options []option.RequestOption
}

// NewGlobalIPHealthCheckTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalIPHealthCheckTypeService(opts ...option.RequestOption) (r GlobalIPHealthCheckTypeService) {
	r = GlobalIPHealthCheckTypeService{}
	r.Options = opts
	return
}

// List all Global IP Health check types.
func (r *GlobalIPHealthCheckTypeService) List(ctx context.Context, opts ...option.RequestOption) (res *GlobalIPHealthCheckTypeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ip_health_check_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GlobalIPHealthCheckTypeListResponse struct {
	Data []GlobalIPHealthCheckTypeListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckTypeListResponseData struct {
	// Global IP Health check params.
	HealthCheckParams map[string]any `json:"health_check_params"`
	// Global IP Health check type.
	HealthCheckType string `json:"health_check_type"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HealthCheckParams respjson.Field
		HealthCheckType   respjson.Field
		RecordType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckTypeListResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckTypeListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
