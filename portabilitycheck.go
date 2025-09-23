// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortabilityCheckService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortabilityCheckService] method instead.
type PortabilityCheckService struct {
	Options []option.RequestOption
}

// NewPortabilityCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPortabilityCheckService(opts ...option.RequestOption) (r PortabilityCheckService) {
	r = PortabilityCheckService{}
	r.Options = opts
	return
}

// Runs a portability check, returning the results immediately.
func (r *PortabilityCheckService) Run(ctx context.Context, body PortabilityCheckRunParams, opts ...option.RequestOption) (res *PortabilityCheckRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "portability_checks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PortabilityCheckRunResponse struct {
	Data []PortabilityCheckRunResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortabilityCheckRunResponse) RawJSON() string { return r.JSON.raw }
func (r *PortabilityCheckRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortabilityCheckRunResponseData struct {
	// Indicates whether this phone number is FastPort eligible
	FastPortable bool `json:"fast_portable"`
	// If this phone number is not portable, explains why. Empty string if the number
	// is portable.
	NotPortableReason string `json:"not_portable_reason"`
	// The +E.164 formatted phone number this result is about
	PhoneNumber string `json:"phone_number"`
	// Indicates whether this phone number is portable
	Portable bool `json:"portable"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FastPortable      respjson.Field
		NotPortableReason respjson.Field
		PhoneNumber       respjson.Field
		Portable          respjson.Field
		RecordType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortabilityCheckRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortabilityCheckRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortabilityCheckRunParams struct {
	// The list of +E.164 formatted phone numbers to check for portability
	PhoneNumbers []string `json:"phone_numbers,omitzero"`
	paramObj
}

func (r PortabilityCheckRunParams) MarshalJSON() (data []byte, err error) {
	type shadow PortabilityCheckRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortabilityCheckRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
