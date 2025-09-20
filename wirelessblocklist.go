// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// WirelessBlocklistService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWirelessBlocklistService] method instead.
type WirelessBlocklistService struct {
	Options []option.RequestOption
}

// NewWirelessBlocklistService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWirelessBlocklistService(opts ...option.RequestOption) (r WirelessBlocklistService) {
	r = WirelessBlocklistService{}
	r.Options = opts
	return
}

// Create a Wireless Blocklist to prevent SIMs from connecting to certain networks.
func (r *WirelessBlocklistService) New(ctx context.Context, body WirelessBlocklistNewParams, opts ...option.RequestOption) (res *WirelessBlocklistNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireless_blocklists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve information about a Wireless Blocklist.
func (r *WirelessBlocklistService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WirelessBlocklistGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireless_blocklists/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Wireless Blocklist.
func (r *WirelessBlocklistService) Update(ctx context.Context, body WirelessBlocklistUpdateParams, opts ...option.RequestOption) (res *WirelessBlocklistUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireless_blocklists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all Wireless Blocklists belonging to the user.
func (r *WirelessBlocklistService) List(ctx context.Context, query WirelessBlocklistListParams, opts ...option.RequestOption) (res *WirelessBlocklistListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireless_blocklists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the Wireless Blocklist.
func (r *WirelessBlocklistService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WirelessBlocklistDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireless_blocklists/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WirelessBlocklist struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The wireless blocklist name.
	Name       string `json:"name"`
	RecordType string `json:"record_type"`
	// The type of the wireless blocklist.
	//
	// Any of "country", "mcc", "plmn".
	Type WirelessBlocklistType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Values to block. The values here depend on the `type` of Wireless Blocklist.
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklist) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the wireless blocklist.
type WirelessBlocklistType string

const (
	WirelessBlocklistTypeCountry WirelessBlocklistType = "country"
	WirelessBlocklistTypeMcc     WirelessBlocklistType = "mcc"
	WirelessBlocklistTypePlmn    WirelessBlocklistType = "plmn"
)

type WirelessBlocklistNewResponse struct {
	Data WirelessBlocklist `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistGetResponse struct {
	Data WirelessBlocklist `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistUpdateResponse struct {
	Data WirelessBlocklist `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistListResponse struct {
	Data []WirelessBlocklist `json:"data"`
	Meta PaginationMeta      `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistListResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistDeleteResponse struct {
	Data WirelessBlocklist `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistNewParams struct {
	// The name of the Wireless Blocklist.
	Name string `json:"name,required"`
	// The type of wireless blocklist.
	//
	// Any of "country", "mcc", "plmn".
	Type WirelessBlocklistNewParamsType `json:"type,omitzero,required"`
	// Values to block. The values here depend on the `type` of Wireless Blocklist.
	Values []string `json:"values,omitzero,required"`
	paramObj
}

func (r WirelessBlocklistNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WirelessBlocklistNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WirelessBlocklistNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of wireless blocklist.
type WirelessBlocklistNewParamsType string

const (
	WirelessBlocklistNewParamsTypeCountry WirelessBlocklistNewParamsType = "country"
	WirelessBlocklistNewParamsTypeMcc     WirelessBlocklistNewParamsType = "mcc"
	WirelessBlocklistNewParamsTypePlmn    WirelessBlocklistNewParamsType = "plmn"
)

type WirelessBlocklistUpdateParams struct {
	// The name of the Wireless Blocklist.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of wireless blocklist.
	//
	// Any of "country", "mcc", "plmn".
	Type WirelessBlocklistUpdateParamsType `json:"type,omitzero"`
	// Values to block. The values here depend on the `type` of Wireless Blocklist.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r WirelessBlocklistUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WirelessBlocklistUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WirelessBlocklistUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of wireless blocklist.
type WirelessBlocklistUpdateParamsType string

const (
	WirelessBlocklistUpdateParamsTypeCountry WirelessBlocklistUpdateParamsType = "country"
	WirelessBlocklistUpdateParamsTypeMcc     WirelessBlocklistUpdateParamsType = "mcc"
	WirelessBlocklistUpdateParamsTypePlmn    WirelessBlocklistUpdateParamsType = "plmn"
)

type WirelessBlocklistListParams struct {
	// The name of the Wireless Blocklist.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// When the Private Wireless Gateway was last updated.
	FilterType param.Opt[string] `query:"filter[type],omitzero" json:"-"`
	// Values to filter on (inclusive).
	FilterValues param.Opt[string] `query:"filter[values],omitzero" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WirelessBlocklistListParams]'s query parameters as
// `url.Values`.
func (r WirelessBlocklistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
