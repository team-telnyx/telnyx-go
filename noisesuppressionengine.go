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

// Noise suppression engines that can be selected when configuring noise
// suppression on voice connections.
//
// NoiseSuppressionEngineService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNoiseSuppressionEngineService] method instead.
type NoiseSuppressionEngineService struct {
	Options []option.RequestOption
}

// NewNoiseSuppressionEngineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNoiseSuppressionEngineService(opts ...option.RequestOption) (r NoiseSuppressionEngineService) {
	r = NoiseSuppressionEngineService{}
	r.Options = opts
	return
}

// Returns all noise suppression engines available to the authenticated user.
// Engines gated behind a feature flag are included only when the flag is enabled
// for the user's account. Results are not paginated; the number of engines is
// expected to remain small.
func (r *NoiseSuppressionEngineService) List(ctx context.Context, opts ...option.RequestOption) (res *NoiseSuppressionEngineListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "noise_suppression_engines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type NoiseSuppressionEngineListResponse struct {
	Data []NoiseSuppressionEngineListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoiseSuppressionEngineListResponse) RawJSON() string { return r.JSON.raw }
func (r *NoiseSuppressionEngineListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A noise suppression engine available to the authenticated user.
type NoiseSuppressionEngineListResponseData struct {
	// Default attenuation level of the engine (0-100, in multiples of ten).
	DefaultAttenuationLevel int64 `json:"default_attenuation_level" api:"required"`
	// Human-readable name of the engine.
	Label string `json:"label" api:"required"`
	// Machine-readable identifier of the engine, used when configuring noise
	// suppression.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultAttenuationLevel respjson.Field
		Label                   respjson.Field
		Value                   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NoiseSuppressionEngineListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NoiseSuppressionEngineListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
