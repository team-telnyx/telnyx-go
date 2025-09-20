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
	"github.com/team-telnyx/telnyx-go/shared"
)

// FqdnService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFqdnService] method instead.
type FqdnService struct {
	Options []option.RequestOption
}

// NewFqdnService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFqdnService(opts ...option.RequestOption) (r FqdnService) {
	r = FqdnService{}
	r.Options = opts
	return
}

// Create a new FQDN object.
func (r *FqdnService) New(ctx context.Context, body FqdnNewParams, opts ...option.RequestOption) (res *FqdnNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fqdns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the details regarding a specific FQDN.
func (r *FqdnService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FqdnGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fqdns/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the details of a specific FQDN.
func (r *FqdnService) Update(ctx context.Context, id string, body FqdnUpdateParams, opts ...option.RequestOption) (res *FqdnUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fqdns/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all FQDNs belonging to the user that match the given filters.
func (r *FqdnService) List(ctx context.Context, query FqdnListParams, opts ...option.RequestOption) (res *FqdnListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fqdns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an FQDN.
func (r *FqdnService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FqdnDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fqdns/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Fqdn struct {
	// Identifies the resource.
	ID string `json:"id" format:"int64"`
	// ID of the FQDN connection to which this FQDN is attached.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The DNS record type for the FQDN. For cases where a port is not set, the DNS
	// record type must be 'srv'. For cases where a port is set, the DNS record type
	// must be 'a'. If the DNS record type is 'a' and a port is not specified, 5060
	// will be used.
	DNSRecordType string `json:"dns_record_type"`
	// FQDN represented by this resource.
	Fqdn string `json:"fqdn"`
	// Port to use when connecting to this FQDN.
	Port int64 `json:"port"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ConnectionID  respjson.Field
		CreatedAt     respjson.Field
		DNSRecordType respjson.Field
		Fqdn          respjson.Field
		Port          respjson.Field
		RecordType    respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Fqdn) RawJSON() string { return r.JSON.raw }
func (r *Fqdn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnNewResponse struct {
	Data Fqdn `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnGetResponse struct {
	Data Fqdn `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnUpdateResponse struct {
	Data Fqdn `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnListResponse struct {
	Data []Fqdn                           `json:"data"`
	Meta shared.ConnectionsPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnListResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnDeleteResponse struct {
	Data Fqdn `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnNewParams struct {
	// ID of the FQDN connection to which this IP should be attached.
	ConnectionID string `json:"connection_id,required"`
	// The DNS record type for the FQDN. For cases where a port is not set, the DNS
	// record type must be 'srv'. For cases where a port is set, the DNS record type
	// must be 'a'. If the DNS record type is 'a' and a port is not specified, 5060
	// will be used.
	DNSRecordType string `json:"dns_record_type,required"`
	// FQDN represented by this resource.
	Fqdn string `json:"fqdn,required"`
	// Port to use when connecting to this FQDN.
	Port param.Opt[int64] `json:"port,omitzero"`
	paramObj
}

func (r FqdnNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FqdnNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FqdnNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnUpdateParams struct {
	// Port to use when connecting to this FQDN.
	Port param.Opt[int64] `json:"port,omitzero"`
	// ID of the FQDN connection to which this IP should be attached.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// The DNS record type for the FQDN. For cases where a port is not set, the DNS
	// record type must be 'srv'. For cases where a port is set, the DNS record type
	// must be 'a'. If the DNS record type is 'a' and a port is not specified, 5060
	// will be used.
	DNSRecordType param.Opt[string] `json:"dns_record_type,omitzero"`
	// FQDN represented by this resource.
	Fqdn param.Opt[string] `json:"fqdn,omitzero"`
	paramObj
}

func (r FqdnUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FqdnUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FqdnUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[connection_id], filter[fqdn], filter[port], filter[dns_record_type]
	Filter FqdnListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page FqdnListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FqdnListParams]'s query parameters as `url.Values`.
func (r FqdnListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[connection_id], filter[fqdn], filter[port], filter[dns_record_type]
type FqdnListParamsFilter struct {
	// ID of the FQDN connection to which the FQDN belongs.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// DNS record type used by the FQDN.
	DNSRecordType param.Opt[string] `query:"dns_record_type,omitzero" json:"-"`
	// FQDN represented by the resource.
	Fqdn param.Opt[string] `query:"fqdn,omitzero" json:"-"`
	// Port to use when connecting to the FQDN.
	Port param.Opt[int64] `query:"port,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FqdnListParamsFilter]'s query parameters as `url.Values`.
func (r FqdnListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type FqdnListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FqdnListParamsPage]'s query parameters as `url.Values`.
func (r FqdnListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
