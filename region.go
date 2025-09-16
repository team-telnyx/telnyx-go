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

// RegionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionService] method instead.
type RegionService struct {
	Options []option.RequestOption
}

// NewRegionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRegionService(opts ...option.RequestOption) (r RegionService) {
	r = RegionService{}
	r.Options = opts
	return
}

// List all regions and the interfaces that region supports
func (r *RegionService) List(ctx context.Context, opts ...option.RequestOption) (res *RegionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RegionListResponse struct {
	Data []RegionListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegionListResponseData struct {
	// A code for the region.
	Code string `json:"code"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A name for the region.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// List of interface types supported in this region.
	SupportedInterfaces []string `json:"supported_interfaces"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code                respjson.Field
		CreatedAt           respjson.Field
		Name                respjson.Field
		RecordType          respjson.Field
		SupportedInterfaces respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *RegionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
