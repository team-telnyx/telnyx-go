// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// PublicInternetGatewayService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPublicInternetGatewayService] method instead.
type PublicInternetGatewayService struct {
	Options []option.RequestOption
}

// NewPublicInternetGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPublicInternetGatewayService(opts ...option.RequestOption) (r PublicInternetGatewayService) {
	r = PublicInternetGatewayService{}
	r.Options = opts
	return
}

// Create a new Public Internet Gateway.
func (r *PublicInternetGatewayService) New(ctx context.Context, body PublicInternetGatewayNewParams, opts ...option.RequestOption) (res *PublicInternetGatewayNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "public_internet_gateways"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Public Internet Gateway.
func (r *PublicInternetGatewayService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PublicInternetGatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("public_internet_gateways/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Public Internet Gateways.
func (r *PublicInternetGatewayService) List(ctx context.Context, query PublicInternetGatewayListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PublicInternetGatewayListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "public_internet_gateways"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all Public Internet Gateways.
func (r *PublicInternetGatewayService) ListAutoPaging(ctx context.Context, query PublicInternetGatewayListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PublicInternetGatewayListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Public Internet Gateway.
func (r *PublicInternetGatewayService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *PublicInternetGatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("public_internet_gateways/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type NetworkInterface struct {
	// A user specified name for the interface.
	Name string `json:"name"`
	// The id of the network associated with the interface.
	NetworkID string `json:"network_id" format:"uuid"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		NetworkID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterface) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkInterfaceRegion struct {
	// The region the interface should be deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceRegion) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayNewResponse struct {
	Data PublicInternetGatewayNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayNewResponseData struct {
	// The publically accessible ip for this interface.
	PublicIP string                                     `json:"public_ip"`
	Region   PublicInternetGatewayNewResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicIP    respjson.Field
		Region      respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
	NetworkInterface
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayNewResponseDataRegion struct {
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
func (r PublicInternetGatewayNewResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayNewResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayGetResponse struct {
	Data PublicInternetGatewayGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayGetResponseData struct {
	// The publically accessible ip for this interface.
	PublicIP string                                     `json:"public_ip"`
	Region   PublicInternetGatewayGetResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicIP    respjson.Field
		Region      respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
	NetworkInterface
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayGetResponseDataRegion struct {
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
func (r PublicInternetGatewayGetResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayGetResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayListResponse struct {
	// The publically accessible ip for this interface.
	PublicIP string                                  `json:"public_ip"`
	Region   PublicInternetGatewayListResponseRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicIP    respjson.Field
		Region      respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
	NetworkInterface
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayListResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayListResponseRegion struct {
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
func (r PublicInternetGatewayListResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayListResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayDeleteResponse struct {
	Data PublicInternetGatewayDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayDeleteResponseData struct {
	// The publically accessible ip for this interface.
	PublicIP string                                        `json:"public_ip"`
	Region   PublicInternetGatewayDeleteResponseDataRegion `json:"region"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicIP    respjson.Field
		Region      respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
	NetworkInterface
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayDeleteResponseDataRegion struct {
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
func (r PublicInternetGatewayDeleteResponseDataRegion) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayDeleteResponseDataRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayNewParams struct {
	// A user specified name for the interface.
	Name param.Opt[string] `json:"name,omitzero"`
	// The id of the network associated with the interface.
	NetworkID param.Opt[string] `json:"network_id,omitzero" format:"uuid"`
	// The region the interface should be deployed to.
	RegionCode param.Opt[string] `json:"region_code,omitzero"`
	paramObj
}

func (r PublicInternetGatewayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PublicInternetGatewayNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicInternetGatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
	Filter PublicInternetGatewayListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page PublicInternetGatewayListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PublicInternetGatewayListParams]'s query parameters as
// `url.Values`.
func (r PublicInternetGatewayListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
type PublicInternetGatewayListParamsFilter struct {
	// The associated network id to filter on.
	NetworkID param.Opt[string] `query:"network_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PublicInternetGatewayListParamsFilter]'s query parameters
// as `url.Values`.
func (r PublicInternetGatewayListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type PublicInternetGatewayListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PublicInternetGatewayListParamsPage]'s query parameters as
// `url.Values`.
func (r PublicInternetGatewayListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
