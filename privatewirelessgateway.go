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
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Private Wireless Gateways operations
//
// PrivateWirelessGatewayService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrivateWirelessGatewayService] method instead.
type PrivateWirelessGatewayService struct {
	Options []option.RequestOption
}

// NewPrivateWirelessGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrivateWirelessGatewayService(opts ...option.RequestOption) (r PrivateWirelessGatewayService) {
	r = PrivateWirelessGatewayService{}
	r.Options = opts
	return
}

// Asynchronously create a Private Wireless Gateway for SIM cards for a previously
// created network. This operation may take several minutes so you can check the
// Private Wireless Gateway status at the section Get a Private Wireless Gateway.
func (r *PrivateWirelessGatewayService) New(ctx context.Context, body PrivateWirelessGatewayNewParams, opts ...option.RequestOption) (res *PrivateWirelessGatewayNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "private_wireless_gateways"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve information about a Private Wireless Gateway.
func (r *PrivateWirelessGatewayService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PrivateWirelessGatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("private_wireless_gateways/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all Private Wireless Gateways belonging to the user.
func (r *PrivateWirelessGatewayService) List(ctx context.Context, query PrivateWirelessGatewayListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PrivateWirelessGateway], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "private_wireless_gateways"
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

// Get all Private Wireless Gateways belonging to the user.
func (r *PrivateWirelessGatewayService) ListAutoPaging(ctx context.Context, query PrivateWirelessGatewayListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PrivateWirelessGateway] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes the Private Wireless Gateway.
func (r *PrivateWirelessGatewayService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *PrivateWirelessGatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("private_wireless_gateways/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PrivateWirelessGateway struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// A list of the resources that have been assigned to the Private Wireless Gateway.
	AssignedResources []PwgAssignedResourcesSummary `json:"assigned_resources"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// IP block used to assign IPs to the SIM cards in the Private Wireless Gateway.
	IPRange string `json:"ip_range"`
	// The private wireless gateway name.
	Name string `json:"name"`
	// The identification of the related network resource.
	NetworkID  string `json:"network_id" format:"uuid"`
	RecordType string `json:"record_type"`
	// The name of the region where the Private Wireless Gateway is deployed.
	RegionCode string `json:"region_code"`
	// The current status or failure details of the Private Wireless Gateway.
	Status PrivateWirelessGatewayStatus `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AssignedResources respjson.Field
		CreatedAt         respjson.Field
		IPRange           respjson.Field
		Name              respjson.Field
		NetworkID         respjson.Field
		RecordType        respjson.Field
		RegionCode        respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateWirelessGateway) RawJSON() string { return r.JSON.raw }
func (r *PrivateWirelessGateway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status or failure details of the Private Wireless Gateway.
type PrivateWirelessGatewayStatus struct {
	// This attribute is an
	// [error code](https://developers.telnyx.com/development/api-fundamentals/api-errors)
	// related to the failure reason.
	ErrorCode string `json:"error_code" api:"nullable"`
	// This attribute provides a human-readable explanation of why a failure happened.
	ErrorDescription string `json:"error_description" api:"nullable"`
	// The current status or failure details of the Private Wireless Gateway. <ul>
	//
	//	<li><code>provisioning</code> - the Private Wireless Gateway is being provisioned.</li>
	//	<li><code>provisioned</code> - the Private Wireless Gateway was provisioned and able to receive connections.</li>
	//	<li><code>failed</code> - the provisioning had failed for a reason and it requires an intervention.</li>
	//	<li><code>decommissioning</code> - the Private Wireless Gateway is being removed from the network.</li>
	//	</ul>
	//	Transitioning between the provisioning and provisioned states may take some time.
	//
	// Any of "provisioning", "provisioned", "failed", "decommissioning".
	Value PrivateWirelessGatewayStatusValue `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorCode        respjson.Field
		ErrorDescription respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateWirelessGatewayStatus) RawJSON() string { return r.JSON.raw }
func (r *PrivateWirelessGatewayStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status or failure details of the Private Wireless Gateway. <ul>
//
//	<li><code>provisioning</code> - the Private Wireless Gateway is being provisioned.</li>
//	<li><code>provisioned</code> - the Private Wireless Gateway was provisioned and able to receive connections.</li>
//	<li><code>failed</code> - the provisioning had failed for a reason and it requires an intervention.</li>
//	<li><code>decommissioning</code> - the Private Wireless Gateway is being removed from the network.</li>
//	</ul>
//	Transitioning between the provisioning and provisioned states may take some time.
type PrivateWirelessGatewayStatusValue string

const (
	PrivateWirelessGatewayStatusValueProvisioning    PrivateWirelessGatewayStatusValue = "provisioning"
	PrivateWirelessGatewayStatusValueProvisioned     PrivateWirelessGatewayStatusValue = "provisioned"
	PrivateWirelessGatewayStatusValueFailed          PrivateWirelessGatewayStatusValue = "failed"
	PrivateWirelessGatewayStatusValueDecommissioning PrivateWirelessGatewayStatusValue = "decommissioning"
)

// The summary of the resource that have been assigned to the Private Wireless
// Gateway.
type PwgAssignedResourcesSummary struct {
	// The current count of a resource type assigned to the Private Wireless Gateway.
	Count int64 `json:"count"`
	// The type of the resource assigned to the Private Wireless Gateway.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PwgAssignedResourcesSummary) RawJSON() string { return r.JSON.raw }
func (r *PwgAssignedResourcesSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateWirelessGatewayNewResponse struct {
	Data PrivateWirelessGateway `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateWirelessGatewayNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PrivateWirelessGatewayNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateWirelessGatewayGetResponse struct {
	Data PrivateWirelessGateway `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateWirelessGatewayGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PrivateWirelessGatewayGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateWirelessGatewayDeleteResponse struct {
	Data PrivateWirelessGateway `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateWirelessGatewayDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PrivateWirelessGatewayDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateWirelessGatewayNewParams struct {
	// The private wireless gateway name.
	Name string `json:"name" api:"required"`
	// The identification of the related network resource.
	NetworkID string `json:"network_id" api:"required" format:"uuid"`
	// The code of the region where the private wireless gateway will be assigned. A
	// list of available regions can be found at the regions endpoint
	RegionCode param.Opt[string] `json:"region_code,omitzero"`
	paramObj
}

func (r PrivateWirelessGatewayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PrivateWirelessGatewayNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrivateWirelessGatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateWirelessGatewayListParams struct {
	// Private Wireless Gateway resource creation date.
	FilterCreatedAt param.Opt[string] `query:"filter[created_at],omitzero" json:"-"`
	// The IP address range of the Private Wireless Gateway.
	FilterIPRange param.Opt[string] `query:"filter[ip_range],omitzero" json:"-"`
	// The name of the Private Wireless Gateway.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// The name of the region where the Private Wireless Gateway is deployed.
	FilterRegionCode param.Opt[string] `query:"filter[region_code],omitzero" json:"-"`
	// When the Private Wireless Gateway was last updated.
	FilterUpdatedAt param.Opt[string] `query:"filter[updated_at],omitzero" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PrivateWirelessGatewayListParams]'s query parameters as
// `url.Values`.
func (r PrivateWirelessGatewayListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
