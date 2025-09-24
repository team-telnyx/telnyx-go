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
	"github.com/team-telnyx/telnyx-go/shared"
)

// DynamicEmergencyEndpointService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDynamicEmergencyEndpointService] method instead.
type DynamicEmergencyEndpointService struct {
	Options []option.RequestOption
}

// NewDynamicEmergencyEndpointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDynamicEmergencyEndpointService(opts ...option.RequestOption) (r DynamicEmergencyEndpointService) {
	r = DynamicEmergencyEndpointService{}
	r.Options = opts
	return
}

// Creates a dynamic emergency endpoints.
func (r *DynamicEmergencyEndpointService) New(ctx context.Context, body DynamicEmergencyEndpointNewParams, opts ...option.RequestOption) (res *DynamicEmergencyEndpointNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dynamic_emergency_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the dynamic emergency endpoint based on the ID provided
func (r *DynamicEmergencyEndpointService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DynamicEmergencyEndpointGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("dynamic_emergency_endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the dynamic emergency endpoints according to filters
func (r *DynamicEmergencyEndpointService) List(ctx context.Context, query DynamicEmergencyEndpointListParams, opts ...option.RequestOption) (res *DynamicEmergencyEndpointListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dynamic_emergency_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the dynamic emergency endpoint based on the ID provided
func (r *DynamicEmergencyEndpointService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DynamicEmergencyEndpointDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("dynamic_emergency_endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DynamicEmergencyEndpoint struct {
	CallbackNumber string `json:"callback_number,required"`
	CallerName     string `json:"caller_name,required"`
	// An id of a currently active dynamic emergency location.
	DynamicEmergencyAddressID string `json:"dynamic_emergency_address_id,required"`
	ID                        string `json:"id"`
	// ISO 8601 formatted date of when the resource was created
	CreatedAt string `json:"created_at"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	SipFromID  string `json:"sip_from_id"`
	// Status of dynamic emergency address
	//
	// Any of "pending", "activated", "rejected".
	Status DynamicEmergencyEndpointStatus `json:"status"`
	// ISO 8601 formatted date of when the resource was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallbackNumber            respjson.Field
		CallerName                respjson.Field
		DynamicEmergencyAddressID respjson.Field
		ID                        respjson.Field
		CreatedAt                 respjson.Field
		RecordType                respjson.Field
		SipFromID                 respjson.Field
		Status                    respjson.Field
		UpdatedAt                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyEndpoint) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DynamicEmergencyEndpoint to a
// DynamicEmergencyEndpointParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DynamicEmergencyEndpointParam.Overrides()
func (r DynamicEmergencyEndpoint) ToParam() DynamicEmergencyEndpointParam {
	return param.Override[DynamicEmergencyEndpointParam](json.RawMessage(r.RawJSON()))
}

// Status of dynamic emergency address
type DynamicEmergencyEndpointStatus string

const (
	DynamicEmergencyEndpointStatusPending   DynamicEmergencyEndpointStatus = "pending"
	DynamicEmergencyEndpointStatusActivated DynamicEmergencyEndpointStatus = "activated"
	DynamicEmergencyEndpointStatusRejected  DynamicEmergencyEndpointStatus = "rejected"
)

// The properties CallbackNumber, CallerName, DynamicEmergencyAddressID are
// required.
type DynamicEmergencyEndpointParam struct {
	CallbackNumber string `json:"callback_number,required"`
	CallerName     string `json:"caller_name,required"`
	// An id of a currently active dynamic emergency location.
	DynamicEmergencyAddressID string            `json:"dynamic_emergency_address_id,required"`
	ID                        param.Opt[string] `json:"id,omitzero"`
	// ISO 8601 formatted date of when the resource was created
	CreatedAt param.Opt[string] `json:"created_at,omitzero"`
	// Identifies the type of the resource.
	RecordType param.Opt[string] `json:"record_type,omitzero"`
	SipFromID  param.Opt[string] `json:"sip_from_id,omitzero"`
	// ISO 8601 formatted date of when the resource was last updated
	UpdatedAt param.Opt[string] `json:"updated_at,omitzero"`
	// Status of dynamic emergency address
	//
	// Any of "pending", "activated", "rejected".
	Status DynamicEmergencyEndpointStatus `json:"status,omitzero"`
	paramObj
}

func (r DynamicEmergencyEndpointParam) MarshalJSON() (data []byte, err error) {
	type shadow DynamicEmergencyEndpointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DynamicEmergencyEndpointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyEndpointNewResponse struct {
	Data DynamicEmergencyEndpoint `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyEndpointNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyEndpointNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyEndpointGetResponse struct {
	Data DynamicEmergencyEndpoint `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyEndpointGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyEndpointGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyEndpointListResponse struct {
	Data []DynamicEmergencyEndpoint `json:"data"`
	Meta shared.Metadata            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyEndpointListResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyEndpointListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyEndpointDeleteResponse struct {
	Data DynamicEmergencyEndpoint `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyEndpointDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyEndpointDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyEndpointNewParams struct {
	DynamicEmergencyEndpoint DynamicEmergencyEndpointParam
	paramObj
}

func (r DynamicEmergencyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DynamicEmergencyEndpoint)
}
func (r *DynamicEmergencyEndpointNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DynamicEmergencyEndpoint)
}

type DynamicEmergencyEndpointListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[status],
	// filter[country_code]
	Filter DynamicEmergencyEndpointListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page DynamicEmergencyEndpointListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DynamicEmergencyEndpointListParams]'s query parameters as
// `url.Values`.
func (r DynamicEmergencyEndpointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[status],
// filter[country_code]
type DynamicEmergencyEndpointListParamsFilter struct {
	// Filter by country code.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "pending", "activated", "rejected".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DynamicEmergencyEndpointListParamsFilter]'s query
// parameters as `url.Values`.
func (r DynamicEmergencyEndpointListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type DynamicEmergencyEndpointListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DynamicEmergencyEndpointListParamsPage]'s query parameters
// as `url.Values`.
func (r DynamicEmergencyEndpointListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
