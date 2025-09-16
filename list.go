// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// ListService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	Options []option.RequestOption
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.Options = opts
	return
}

// Retrieve a list of all phone numbers using Channel Billing, grouped by Zone.
func (r *ListService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *ListGetAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of phone numbers using Channel Billing for a specific Zone.
func (r *ListService) GetByZone(ctx context.Context, channelZoneID string, opts ...option.RequestOption) (res *ListGetByZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelZoneID == "" {
		err = errors.New("missing required channel_zone_id parameter")
		return
	}
	path := fmt.Sprintf("list/%s", channelZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ListGetAllResponse struct {
	Data []ListGetAllResponseData `json:"data"`
	Meta PaginationMeta           `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetAllResponse) RawJSON() string { return r.JSON.raw }
func (r *ListGetAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetAllResponseData struct {
	NumberOfChannels int64                          `json:"number_of_channels"`
	Numbers          []ListGetAllResponseDataNumber `json:"numbers"`
	ZoneID           string                         `json:"zone_id"`
	ZoneName         string                         `json:"zone_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberOfChannels respjson.Field
		Numbers          respjson.Field
		ZoneID           respjson.Field
		ZoneName         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetAllResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListGetAllResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetAllResponseDataNumber struct {
	Country string `json:"country"`
	Number  string `json:"number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		Number      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetAllResponseDataNumber) RawJSON() string { return r.JSON.raw }
func (r *ListGetAllResponseDataNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetByZoneResponse struct {
	Data []ListGetByZoneResponseData `json:"data"`
	Meta PaginationMeta              `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetByZoneResponse) RawJSON() string { return r.JSON.raw }
func (r *ListGetByZoneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetByZoneResponseData struct {
	NumberOfChannels int64                             `json:"number_of_channels"`
	Numbers          []ListGetByZoneResponseDataNumber `json:"numbers"`
	ZoneID           string                            `json:"zone_id"`
	ZoneName         string                            `json:"zone_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberOfChannels respjson.Field
		Numbers          respjson.Field
		ZoneID           respjson.Field
		ZoneName         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetByZoneResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListGetByZoneResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetByZoneResponseDataNumber struct {
	Country string `json:"country"`
	Number  string `json:"number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		Number      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetByZoneResponseDataNumber) RawJSON() string { return r.JSON.raw }
func (r *ListGetByZoneResponseDataNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
