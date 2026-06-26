// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Public Internet Gateway operations
//
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
	return res, err
}

// Retrieve a Public Internet Gateway.
func (r *PublicInternetGatewayService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PublicInternetGatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("public_internet_gateways/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all Public Internet Gateways.
func (r *PublicInternetGatewayService) List(ctx context.Context, query PublicInternetGatewayListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PublicInternetGatewayRead], err error) {
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
func (r *PublicInternetGatewayService) ListAutoPaging(ctx context.Context, query PublicInternetGatewayListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PublicInternetGatewayRead] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Public Internet Gateway.
func (r *PublicInternetGatewayService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *PublicInternetGatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("public_internet_gateways/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type NetworkInterfaceParam struct {
	// A user specified name for the interface.
	Name param.Opt[string] `json:"name,omitzero"`
	// The id of the network associated with the interface.
	NetworkID param.Opt[string] `json:"network_id,omitzero" format:"uuid"`
	paramObj
}

func (r NetworkInterfaceParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkInterfaceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkInterfaceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkInterfaceRegionParam struct {
	// The region the interface should be deployed to.
	RegionCode param.Opt[string] `json:"region_code,omitzero"`
	paramObj
}

func (r NetworkInterfaceRegionParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkInterfaceRegionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkInterfaceRegionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayParam struct {
	// The region interface is deployed to.
	RegionCode param.Opt[string] `json:"region_code,omitzero"`
	RecordParam
	NetworkInterfaceParam
	paramObj
}

func (r PublicInternetGatewayParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*PublicInternetGatewayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type PublicInternetGatewayRead struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A user specified name for the interface.
	Name string `json:"name"`
	// The id of the network associated with the interface.
	NetworkID string `json:"network_id" format:"uuid"`
	// The publically accessible ip for this interface.
	PublicIP string `json:"public_ip"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The region interface is deployed to.
	RegionCode string `json:"region_code"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		NetworkID   respjson.Field
		PublicIP    respjson.Field
		RecordType  respjson.Field
		RegionCode  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicInternetGatewayRead) RawJSON() string { return r.JSON.raw }
func (r *PublicInternetGatewayRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayNewResponse struct {
	Data PublicInternetGatewayRead `json:"data"`
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

type PublicInternetGatewayGetResponse struct {
	Data PublicInternetGatewayRead `json:"data"`
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

type PublicInternetGatewayDeleteResponse struct {
	Data PublicInternetGatewayRead `json:"data"`
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

type PublicInternetGatewayNewParams struct {
	Body PublicInternetGatewayNewParamsBody
	paramObj
}

func (r PublicInternetGatewayNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PublicInternetGatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicInternetGatewayNewParamsBody struct {
	PublicInternetGatewayParam
}

func (r PublicInternetGatewayNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*PublicInternetGatewayNewParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type PublicInternetGatewayListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[network_id]
	Filter PublicInternetGatewayListParamsFilter `query:"filter,omitzero" json:"-"`
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
