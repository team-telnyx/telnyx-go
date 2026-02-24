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

// SimCardGroupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardGroupService] method instead.
type SimCardGroupService struct {
	Options []option.RequestOption
	Actions SimCardGroupActionService
}

// NewSimCardGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimCardGroupService(opts ...option.RequestOption) (r SimCardGroupService) {
	r = SimCardGroupService{}
	r.Options = opts
	r.Actions = NewSimCardGroupActionService(opts...)
	return
}

// Creates a new SIM card group object
func (r *SimCardGroupService) New(ctx context.Context, body SimCardGroupNewParams, opts ...option.RequestOption) (res *SimCardGroupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_card_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the details regarding a specific SIM card group
func (r *SimCardGroupService) Get(ctx context.Context, id string, query SimCardGroupGetParams, opts ...option.RequestOption) (res *SimCardGroupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates a SIM card group
func (r *SimCardGroupService) Update(ctx context.Context, id string, body SimCardGroupUpdateParams, opts ...option.RequestOption) (res *SimCardGroupUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all SIM card groups belonging to the user that match the given filters.
func (r *SimCardGroupService) List(ctx context.Context, query SimCardGroupListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[SimCardGroupListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "sim_card_groups"
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

// Get all SIM card groups belonging to the user that match the given filters.
func (r *SimCardGroupService) ListAutoPaging(ctx context.Context, query SimCardGroupListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[SimCardGroupListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a SIM card group
func (r *SimCardGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGroupDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Represents the amount of data consumed.
type ConsumedData struct {
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsumedData) RawJSON() string { return r.JSON.raw }
func (r *ConsumedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroup struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// Represents the amount of data consumed.
	ConsumedData ConsumedData `json:"consumed_data"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Upper limit on the amount of data the SIM cards, within the group, can use.
	DataLimit SimCardGroupDataLimit `json:"data_limit"`
	// Indicates whether the SIM card group is the users default group.<br/>The default
	// group is created for the user and can not be removed.
	Default bool `json:"default"`
	// A user friendly name for the SIM card group.
	Name string `json:"name"`
	// The identification of the related Private Wireless Gateway resource.
	PrivateWirelessGatewayID string `json:"private_wireless_gateway_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// The identification of the related Wireless Blocklist resource.
	WirelessBlocklistID string `json:"wireless_blocklist_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ConsumedData             respjson.Field
		CreatedAt                respjson.Field
		DataLimit                respjson.Field
		Default                  respjson.Field
		Name                     respjson.Field
		PrivateWirelessGatewayID respjson.Field
		RecordType               respjson.Field
		UpdatedAt                respjson.Field
		WirelessBlocklistID      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroup) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upper limit on the amount of data the SIM cards, within the group, can use.
type SimCardGroupDataLimit struct {
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupDataLimit) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupDataLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupNewResponse struct {
	Data SimCardGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupGetResponse struct {
	Data SimCardGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupUpdateResponse struct {
	Data SimCardGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupListResponse struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// Represents the amount of data consumed.
	ConsumedData ConsumedData `json:"consumed_data"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Upper limit on the amount of data the SIM cards, within the group, can use.
	DataLimit SimCardGroupListResponseDataLimit `json:"data_limit"`
	// Indicates whether the SIM card group is the users default group.<br/>The default
	// group is created for the user and can not be removed.
	Default bool `json:"default"`
	// A user friendly name for the SIM card group.
	Name string `json:"name"`
	// The identification of the related Private Wireless Gateway resource.
	PrivateWirelessGatewayID string `json:"private_wireless_gateway_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The number of SIM cards associated with the group.
	SimCardCount int64 `json:"sim_card_count"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// The identification of the related Wireless Blocklist resource.
	WirelessBlocklistID string `json:"wireless_blocklist_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ConsumedData             respjson.Field
		CreatedAt                respjson.Field
		DataLimit                respjson.Field
		Default                  respjson.Field
		Name                     respjson.Field
		PrivateWirelessGatewayID respjson.Field
		RecordType               respjson.Field
		SimCardCount             respjson.Field
		UpdatedAt                respjson.Field
		WirelessBlocklistID      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upper limit on the amount of data the SIM cards, within the group, can use.
type SimCardGroupListResponseDataLimit struct {
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupListResponseDataLimit) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupListResponseDataLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupDeleteResponse struct {
	Data SimCardGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupNewParams struct {
	// A user friendly name for the SIM card group.
	Name string `json:"name" api:"required"`
	// Upper limit on the amount of data the SIM cards, within the group, can use.
	DataLimit SimCardGroupNewParamsDataLimit `json:"data_limit,omitzero"`
	paramObj
}

func (r SimCardGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upper limit on the amount of data the SIM cards, within the group, can use.
type SimCardGroupNewParamsDataLimit struct {
	Amount param.Opt[string] `json:"amount,omitzero"`
	Unit   param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r SimCardGroupNewParamsDataLimit) MarshalJSON() (data []byte, err error) {
	type shadow SimCardGroupNewParamsDataLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardGroupNewParamsDataLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupGetParams struct {
	// It includes a list of associated ICCIDs.
	IncludeIccids param.Opt[bool] `query:"include_iccids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardGroupGetParams]'s query parameters as `url.Values`.
func (r SimCardGroupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SimCardGroupUpdateParams struct {
	// A user friendly name for the SIM card group.
	Name param.Opt[string] `json:"name,omitzero"`
	// Upper limit on the amount of data the SIM cards, within the group, can use.
	DataLimit SimCardGroupUpdateParamsDataLimit `json:"data_limit,omitzero"`
	paramObj
}

func (r SimCardGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upper limit on the amount of data the SIM cards, within the group, can use.
type SimCardGroupUpdateParamsDataLimit struct {
	Amount param.Opt[string] `json:"amount,omitzero"`
	Unit   param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r SimCardGroupUpdateParamsDataLimit) MarshalJSON() (data []byte, err error) {
	type shadow SimCardGroupUpdateParamsDataLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardGroupUpdateParamsDataLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupListParams struct {
	// A valid SIM card group name.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// A Private Wireless Gateway ID associated with the group.
	FilterPrivateWirelessGatewayID param.Opt[string] `query:"filter[private_wireless_gateway_id],omitzero" format:"uuid" json:"-"`
	// A Wireless Blocklist ID associated with the group.
	FilterWirelessBlocklistID param.Opt[string] `query:"filter[wireless_blocklist_id],omitzero" format:"uuid" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardGroupListParams]'s query parameters as `url.Values`.
func (r SimCardGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
