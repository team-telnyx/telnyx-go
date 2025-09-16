// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// GlobalIPProtocolService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPProtocolService] method instead.
type GlobalIPProtocolService struct {
	Options []option.RequestOption
}

// NewGlobalIPProtocolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalIPProtocolService(opts ...option.RequestOption) (r GlobalIPProtocolService) {
	r = GlobalIPProtocolService{}
	r.Options = opts
	return
}

// List all Global IP Protocols
func (r *GlobalIPProtocolService) List(ctx context.Context, opts ...option.RequestOption) (res *GlobalIPProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global_ip_protocols"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GlobalIPProtocolListResponse struct {
	Data []GlobalIPProtocolListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPProtocolListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPProtocolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPProtocolListResponseData struct {
	// The Global IP Protocol code.
	Code string `json:"code"`
	// A name for Global IP Protocol.
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
func (r GlobalIPProtocolListResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPProtocolListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
