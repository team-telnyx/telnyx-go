// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// NumbersFeatureService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumbersFeatureService] method instead.
type NumbersFeatureService struct {
	Options []option.RequestOption
}

// NewNumbersFeatureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumbersFeatureService(opts ...option.RequestOption) (r NumbersFeatureService) {
	r = NumbersFeatureService{}
	r.Options = opts
	return
}

// Retrieve the features for a list of numbers
func (r *NumbersFeatureService) New(ctx context.Context, body NumbersFeatureNewParams, opts ...option.RequestOption) (res *NumbersFeatureNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "numbers_features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type NumbersFeatureNewResponse struct {
	Data []NumbersFeatureNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersFeatureNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NumbersFeatureNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersFeatureNewResponseData struct {
	Features    []string `json:"features" api:"required"`
	PhoneNumber string   `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersFeatureNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumbersFeatureNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersFeatureNewParams struct {
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r NumbersFeatureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NumbersFeatureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumbersFeatureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
