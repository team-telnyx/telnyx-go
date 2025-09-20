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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// GlobalIPAssignmentService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPAssignmentService] method instead.
type GlobalIPAssignmentService struct {
	Options []option.RequestOption
}

// NewGlobalIPAssignmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalIPAssignmentService(opts ...option.RequestOption) (r GlobalIPAssignmentService) {
	r = GlobalIPAssignmentService{}
	r.Options = opts
	return
}

// Create a Global IP assignment.
func (r *GlobalIPAssignmentService) New(ctx context.Context, body GlobalIPAssignmentNewParams, opts ...option.RequestOption) (res *GlobalIPAssignmentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ip_assignments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Global IP assignment.
func (r *GlobalIPAssignmentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPAssignmentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ip_assignments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Global IP assignment.
func (r *GlobalIPAssignmentService) Update(ctx context.Context, id string, body GlobalIPAssignmentUpdateParams, opts ...option.RequestOption) (res *GlobalIPAssignmentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ip_assignments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Global IP assignments.
func (r *GlobalIPAssignmentService) List(ctx context.Context, query GlobalIPAssignmentListParams, opts ...option.RequestOption) (res *GlobalIPAssignmentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ip_assignments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Global IP assignment.
func (r *GlobalIPAssignmentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPAssignmentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ip_assignments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type GlobalIPAssignment struct {
	// Global IP ID.
	GlobalIPID string `json:"global_ip_id" format:"uuid"`
	// Status of BGP announcement.
	IsAnnounced bool `json:"is_announced"`
	// Wireguard peer is connected.
	IsConnected bool `json:"is_connected"`
	// Enable/disable BGP announcement.
	IsInMaintenance bool `json:"is_in_maintenance"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the interface deployment.
	//
	// Any of "created", "provisioning", "provisioned", "deleting".
	Status InterfaceStatus `json:"status"`
	// Wireguard peer ID.
	WireguardPeerID string `json:"wireguard_peer_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIPID      respjson.Field
		IsAnnounced     respjson.Field
		IsConnected     respjson.Field
		IsInMaintenance respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		WireguardPeerID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GlobalIPAssignment to a GlobalIPAssignmentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GlobalIPAssignmentParam.Overrides()
func (r GlobalIPAssignment) ToParam() GlobalIPAssignmentParam {
	return param.Override[GlobalIPAssignmentParam](json.RawMessage(r.RawJSON()))
}

type GlobalIPAssignmentParam struct {
	// Global IP ID.
	GlobalIPID param.Opt[string] `json:"global_ip_id,omitzero" format:"uuid"`
	// Status of BGP announcement.
	IsAnnounced param.Opt[bool] `json:"is_announced,omitzero"`
	// Wireguard peer is connected.
	IsConnected param.Opt[bool] `json:"is_connected,omitzero"`
	// Enable/disable BGP announcement.
	IsInMaintenance param.Opt[bool] `json:"is_in_maintenance,omitzero"`
	// Identifies the type of the resource.
	RecordType param.Opt[string] `json:"record_type,omitzero"`
	// The current status of the interface deployment.
	Status InterfaceStatus `json:"status,omitzero"`
	// Wireguard peer ID.
	WireguardPeerID param.Opt[string] `json:"wireguard_peer_id,omitzero" format:"uuid"`
	RecordParam
}

func (r GlobalIPAssignmentParam) MarshalJSON() (data []byte, err error) {
	type shadow GlobalIPAssignmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type Record struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Record) RawJSON() string { return r.JSON.raw }
func (r *Record) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Record to a RecordParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RecordParam.Overrides()
func (r Record) ToParam() RecordParam {
	return param.Override[RecordParam](json.RawMessage(r.RawJSON()))
}

type RecordParam struct {
	// Identifies the resource.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt param.Opt[string] `json:"created_at,omitzero"`
	// Identifies the type of the resource.
	RecordType param.Opt[string] `json:"record_type,omitzero"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt param.Opt[string] `json:"updated_at,omitzero"`
	paramObj
}

func (r RecordParam) MarshalJSON() (data []byte, err error) {
	type shadow RecordParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentNewResponse struct {
	Data GlobalIPAssignment `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentGetResponse struct {
	Data GlobalIPAssignment `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentUpdateResponse struct {
	Data GlobalIPAssignment `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentListResponse struct {
	Data []GlobalIPAssignment `json:"data"`
	Meta PaginationMeta       `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentDeleteResponse struct {
	Data GlobalIPAssignment `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentNewParams struct {
	GlobalIPAssignment GlobalIPAssignmentParam
	paramObj
}

func (r GlobalIPAssignmentNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GlobalIPAssignment)
}
func (r *GlobalIPAssignmentNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.GlobalIPAssignment)
}

type GlobalIPAssignmentUpdateParams struct {
	Body GlobalIPAssignmentUpdateParamsBody
	paramObj
}

func (r GlobalIPAssignmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *GlobalIPAssignmentUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type GlobalIPAssignmentUpdateParamsBody struct {
	GlobalIPID      param.Opt[string] `json:"global_ip_id,omitzero"`
	WireguardPeerID param.Opt[string] `json:"wireguard_peer_id,omitzero"`
	GlobalIPAssignmentParam
}

func (r GlobalIPAssignmentUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow GlobalIPAssignmentUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

type GlobalIPAssignmentListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page GlobalIPAssignmentListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentListParams]'s query parameters as
// `url.Values`.
func (r GlobalIPAssignmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type GlobalIPAssignmentListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentListParamsPage]'s query parameters as
// `url.Values`.
func (r GlobalIPAssignmentListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
