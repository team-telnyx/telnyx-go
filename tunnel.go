// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Tunnels connect a Network to hosts behind a tunnel process within the user's
// network.
//
// TunnelService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelService] method instead.
type TunnelService struct {
	Options []option.RequestOption
	// Tunnels connect a Network to hosts behind a tunnel process within the user's
	// network.
	Actions TunnelActionService
}

// NewTunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTunnelService(opts ...option.RequestOption) (r TunnelService) {
	r = TunnelService{}
	r.Options = opts
	r.Actions = NewTunnelActionService(opts...)
	return
}

// Creates a tunnel and returns its `tunnel_token`. The token is the only
// credential the tunnel process needs, and this response and
// `POST /tunnels/{id}/actions/rotate_token` are the only ones that return it. A
// new tunnel is `inactive` until a tunnel process connects with the token.
func (r *TunnelService) New(ctx context.Context, body TunnelNewParams, opts ...option.RequestOption) (res *TunnelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tunnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a tunnel with its status and the tunnel processes connected to it. The
// response never includes `tunnel_token`.
func (r *TunnelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TunnelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tunnels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the tunnels on the account with their status and connected tunnel
// processes. Responses never include `tunnel_token`.
func (r *TunnelService) List(ctx context.Context, query TunnelListParams, opts ...option.RequestOption) (res *TunnelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tunnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a tunnel. A tunnel that a route names as its next hop cannot be deleted;
// delete those routes first. Connected tunnel processes are disconnected and
// cannot reconnect with the tunnel's token.
func (r *TunnelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("tunnels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Tunnel struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// The number of tunnel processes connected with the tunnel's token.
	ConnectionCount int64 `json:"connection_count"`
	// The tunnel processes connected with the tunnel's token.
	Connections []TunnelConnection `json:"connections"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A user specified name for the tunnel.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// `inactive` when no tunnel process is connected, and `healthy` when one or more
	// are. `degraded` (Preview, not reported yet) when a process is connected but
	// fails to reach its destinations.
	//
	// Any of "inactive", "healthy", "degraded".
	Status TunnelStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ConnectionCount respjson.Field
		Connections     respjson.Field
		CreatedAt       respjson.Field
		Name            respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tunnel) RawJSON() string { return r.JSON.raw }
func (r *Tunnel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tunnel process connected to the tunnel.
type TunnelConnection struct {
	// When the process connected.
	ConnectedAt time.Time `json:"connected_at" format:"date-time"`
	// The name the host reports for itself. Preview: empty until the CLI reports it.
	Hostname string `json:"hostname"`
	// The Telnyx region the process is connected to. Preview: empty until the CLI
	// reports it.
	RegionCode string `json:"region_code"`
	// The `telnyx-edge` CLI version of the process. Preview: empty until the CLI
	// reports it.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectedAt respjson.Field
		Hostname    respjson.Field
		RegionCode  respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TunnelConnection) RawJSON() string { return r.JSON.raw }
func (r *TunnelConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `inactive` when no tunnel process is connected, and `healthy` when one or more
// are. `degraded` (Preview, not reported yet) when a process is connected but
// fails to reach its destinations.
type TunnelStatus string

const (
	TunnelStatusInactive TunnelStatus = "inactive"
	TunnelStatusHealthy  TunnelStatus = "healthy"
	TunnelStatusDegraded TunnelStatus = "degraded"
)

type TunnelWithToken struct {
	// The credential that a tunnel process presents with
	// `telnyx-edge tunnel run --token`. Returned only by create and rotate_token.
	// Telnyx stores only a hash of it.
	TunnelToken string `json:"tunnel_token"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TunnelToken respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Tunnel
}

// Returns the unmodified JSON received from the API
func (r TunnelWithToken) RawJSON() string { return r.JSON.raw }
func (r *TunnelWithToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelNewResponse struct {
	Data TunnelWithToken `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TunnelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TunnelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponse struct {
	Data Tunnel `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TunnelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TunnelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelListResponse struct {
	Data []Tunnel                    `json:"data"`
	Meta NetappsCursorPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TunnelListResponse) RawJSON() string { return r.JSON.raw }
func (r *TunnelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelNewParams struct {
	// A user specified name for the tunnel.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r TunnelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TunnelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TunnelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[name],
	// filter[status]
	Filter TunnelListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[after],
	// page[limit]
	Page TunnelListParamsPage `query:"page,omitzero" json:"-"`
	// Sorts the results by a field, ascending, or descending with a leading `-`. The
	// default is `-created_at`.
	//
	// Any of "name", "-name", "status", "-status", "created_at", "-created_at",
	// "updated_at", "-updated_at".
	Sort TunnelListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TunnelListParams]'s query parameters as `url.Values`.
func (r TunnelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[name],
// filter[status]
type TunnelListParamsFilter struct {
	// The name of the tunnel.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The status of the tunnel.
	//
	// Any of "inactive", "healthy", "degraded".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TunnelListParamsFilter]'s query parameters as `url.Values`.
func (r TunnelListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[after],
// page[limit]
type TunnelListParamsPage struct {
	// The cursor from `meta.cursors.after` of the previous page.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The number of results per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TunnelListParamsPage]'s query parameters as `url.Values`.
func (r TunnelListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts the results by a field, ascending, or descending with a leading `-`. The
// default is `-created_at`.
type TunnelListParamsSort string

const (
	TunnelListParamsSortName          TunnelListParamsSort = "name"
	TunnelListParamsSortNameDesc      TunnelListParamsSort = "-name"
	TunnelListParamsSortStatus        TunnelListParamsSort = "status"
	TunnelListParamsSortStatusDesc    TunnelListParamsSort = "-status"
	TunnelListParamsSortCreatedAt     TunnelListParamsSort = "created_at"
	TunnelListParamsSortCreatedAtDesc TunnelListParamsSort = "-created_at"
	TunnelListParamsSortUpdatedAt     TunnelListParamsSort = "updated_at"
	TunnelListParamsSortUpdatedAtDesc TunnelListParamsSort = "-updated_at"
)
