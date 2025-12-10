// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// NetworkService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options        []option.RequestOption
	DefaultGateway NetworkDefaultGatewayService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r NetworkService) {
	r = NetworkService{}
	r.Options = opts
	r.DefaultGateway = NewNetworkDefaultGatewayService(opts...)
	return
}

// Create a new Network.
func (r *NetworkService) New(ctx context.Context, body NetworkNewParams, opts ...option.RequestOption) (res *NetworkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "networks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Network.
func (r *NetworkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Network.
func (r *NetworkService) Update(ctx context.Context, id string, body NetworkUpdateParams, opts ...option.RequestOption) (res *NetworkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Networks.
func (r *NetworkService) List(ctx context.Context, query NetworkListParams, opts ...option.RequestOption) (res *NetworkListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "networks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Network.
func (r *NetworkService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NetworkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all Interfaces for a Network.
func (r *NetworkService) ListInterfaces(ctx context.Context, id string, query NetworkListInterfacesParams, opts ...option.RequestOption) (res *NetworkListInterfacesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("networks/%s/network_interfaces", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The current status of the interface deployment.
type InterfaceStatus string

const (
	InterfaceStatusCreated      InterfaceStatus = "created"
	InterfaceStatusProvisioning InterfaceStatus = "provisioning"
	InterfaceStatusProvisioned  InterfaceStatus = "provisioned"
	InterfaceStatusDeleting     InterfaceStatus = "deleting"
)

type NetworkCreateParam struct {
	// A user specified name for the network.
	Name string `json:"name,required"`
	RecordParam
}

func (r NetworkCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type NetworkNewResponse struct {
	Data NetworkNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkNewResponseData struct {
	// A user specified name for the network.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkGetResponse struct {
	Data NetworkGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkGetResponseData struct {
	// A user specified name for the network.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkUpdateResponse struct {
	Data NetworkUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkUpdateResponseData struct {
	// A user specified name for the network.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponse struct {
	Data []NetworkListResponseData `json:"data"`
	Meta PaginationMeta            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponseData struct {
	// A user specified name for the network.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDeleteResponse struct {
	Data NetworkDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDeleteResponseData struct {
	// A user specified name for the network.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r NetworkDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListInterfacesResponse struct {
	Data []NetworkListInterfacesResponseData `json:"data"`
	Meta PaginationMeta                      `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkListInterfacesResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkListInterfacesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListInterfacesResponseData struct {
	// Identifies the type of the resource.
	RecordType string                                  `json:"record_type"`
	Region     NetworkListInterfacesResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// Identifies the type of the interface.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType  respjson.Field
		Region      respjson.Field
		RegionCode  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
	Interface
}

// Returns the unmodified JSON received from the API
func (r NetworkListInterfacesResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkListInterfacesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListInterfacesResponseDataRegion struct {
	// Region code of the interface.
	Code string `json:"code"`
	// Region name of the interface.
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
func (r NetworkListInterfacesResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *NetworkListInterfacesResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkNewParams struct {
	NetworkCreate NetworkCreateParam
	paramObj
}

func (r NetworkNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NetworkCreate)
}
func (r *NetworkNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NetworkCreate)
}

type NetworkUpdateParams struct {
	NetworkCreate NetworkCreateParam
	paramObj
}

func (r NetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NetworkCreate)
}
func (r *NetworkUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NetworkCreate)
}

type NetworkListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[name]
	Filter NetworkListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NetworkListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListParams]'s query parameters as `url.Values`.
func (r NetworkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[name]
type NetworkListParamsFilter struct {
	// The network name to filter on.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListParamsFilter]'s query parameters as
// `url.Values`.
func (r NetworkListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NetworkListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListParamsPage]'s query parameters as `url.Values`.
func (r NetworkListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetworkListInterfacesParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[name],
	// filter[type], filter[status]
	Filter NetworkListInterfacesParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NetworkListInterfacesParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListInterfacesParams]'s query parameters as
// `url.Values`.
func (r NetworkListInterfacesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[name],
// filter[type], filter[status]
type NetworkListInterfacesParamsFilter struct {
	// The interface name to filter on.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The interface type to filter on.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListInterfacesParamsFilter]'s query parameters as
// `url.Values`.
func (r NetworkListInterfacesParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NetworkListInterfacesParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListInterfacesParamsPage]'s query parameters as
// `url.Values`.
func (r NetworkListInterfacesParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
