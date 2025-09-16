// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// IPService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r IPService) {
	r = IPService{}
	r.Options = opts
	return
}

// Create a new IP object.
func (r *IPService) New(ctx context.Context, body IPNewParams, opts ...option.RequestOption) (res *IPNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the details regarding a specific IP.
func (r *IPService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the details of a specific IP.
func (r *IPService) Update(ctx context.Context, id string, body IPUpdateParams, opts ...option.RequestOption) (res *IPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all IPs belonging to the user that match the given filters.
func (r *IPService) List(ctx context.Context, query IPListParams, opts ...option.RequestOption) (res *IPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an IP.
func (r *IPService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *IPDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type IP struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"int64"`
	// ID of the IP Connection to which this IP should be attached.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// IP adddress represented by this resource.
	IPAddress string `json:"ip_address"`
	// Port to use when connecting to this IP.
	Port int64 `json:"port"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ConnectionID respjson.Field
		CreatedAt    respjson.Field
		IPAddress    respjson.Field
		Port         respjson.Field
		RecordType   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IP) RawJSON() string { return r.JSON.raw }
func (r *IP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPNewResponse struct {
	Data IP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPNewResponse) RawJSON() string { return r.JSON.raw }
func (r *IPNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPGetResponse struct {
	Data IP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IPGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPUpdateResponse struct {
	Data IP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *IPUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPListResponse struct {
	Data []IP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPListResponse) RawJSON() string { return r.JSON.raw }
func (r *IPListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPDeleteResponse struct {
	Data IP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *IPDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPNewParams struct {
	// IP adddress represented by this resource.
	IPAddress string `json:"ip_address,required"`
	// ID of the IP Connection to which this IP should be attached.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// Port to use when connecting to this IP.
	Port param.Opt[int64] `json:"port,omitzero"`
	paramObj
}

func (r IPNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IPNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPUpdateParams struct {
	// IP adddress represented by this resource.
	IPAddress string `json:"ip_address,required"`
	// ID of the IP Connection to which this IP should be attached.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// Port to use when connecting to this IP.
	Port param.Opt[int64] `json:"port,omitzero"`
	paramObj
}

func (r IPUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow IPUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_id], filter[ip_address], filter[port]
	Filter IPListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page IPListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_id], filter[ip_address], filter[port]
type IPListParamsFilter struct {
	// ID of the IP Connection to which this IP should be attached.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// IP adddress represented by this resource.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// Port to use when connecting to this IP.
	Port param.Opt[int64] `query:"port,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPListParamsFilter]'s query parameters as `url.Values`.
func (r IPListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type IPListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPListParamsPage]'s query parameters as `url.Values`.
func (r IPListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
