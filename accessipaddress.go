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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AccessIPAddressService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessIPAddressService] method instead.
type AccessIPAddressService struct {
	Options []option.RequestOption
}

// NewAccessIPAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessIPAddressService(opts ...option.RequestOption) (r AccessIPAddressService) {
	r = AccessIPAddressService{}
	r.Options = opts
	return
}

// Create new Access IP Address
func (r *AccessIPAddressService) New(ctx context.Context, body AccessIPAddressNewParams, opts ...option.RequestOption) (res *AccessIPAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "access_ip_address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an access IP address
func (r *AccessIPAddressService) Get(ctx context.Context, accessIPAddressID string, opts ...option.RequestOption) (res *AccessIPAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accessIPAddressID == "" {
		err = errors.New("missing required access_ip_address_id parameter")
		return
	}
	path := fmt.Sprintf("access_ip_address/%s", accessIPAddressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Access IP Addresses
func (r *AccessIPAddressService) List(ctx context.Context, query AccessIPAddressListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AccessIPAddressResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "access_ip_address"
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

// List all Access IP Addresses
func (r *AccessIPAddressService) ListAutoPaging(ctx context.Context, query AccessIPAddressListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AccessIPAddressResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete access IP address
func (r *AccessIPAddressService) Delete(ctx context.Context, accessIPAddressID string, opts ...option.RequestOption) (res *AccessIPAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accessIPAddressID == "" {
		err = errors.New("missing required access_ip_address_id parameter")
		return
	}
	path := fmt.Sprintf("access_ip_address/%s", accessIPAddressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessIPAddressResponse struct {
	ID        string `json:"id,required"`
	IPAddress string `json:"ip_address,required"`
	Source    string `json:"source,required"`
	// An enumeration.
	//
	// Any of "pending", "added".
	Status      CloudflareSyncStatus `json:"status,required"`
	UserID      string               `json:"user_id,required"`
	CreatedAt   time.Time            `json:"created_at" format:"date-time"`
	Description string               `json:"description"`
	UpdatedAt   time.Time            `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IPAddress   respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		UserID      respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessIPAddressResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessIPAddressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An enumeration.
type CloudflareSyncStatus string

const (
	CloudflareSyncStatusPending CloudflareSyncStatus = "pending"
	CloudflareSyncStatusAdded   CloudflareSyncStatus = "added"
)

type PaginationMetaCloudflareIPListSync struct {
	PageNumber   int64 `json:"page_number,required"`
	PageSize     int64 `json:"page_size,required"`
	TotalPages   int64 `json:"total_pages,required"`
	TotalResults int64 `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationMetaCloudflareIPListSync) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetaCloudflareIPListSync) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIPAddressNewParams struct {
	IPAddress   string            `json:"ip_address,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AccessIPAddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessIPAddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessIPAddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIPAddressListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[ip_source],
	// filter[ip_address], filter[created_at]. Supports complex bracket operations for
	// dynamic filtering.
	Filter AccessIPAddressListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccessIPAddressListParams]'s query parameters as
// `url.Values`.
func (r AccessIPAddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[ip_source],
// filter[ip_address], filter[created_at]. Supports complex bracket operations for
// dynamic filtering.
type AccessIPAddressListParamsFilter struct {
	// Filter by IP address
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// Filter by IP source
	IPSource param.Opt[string] `query:"ip_source,omitzero" json:"-"`
	// Filter by exact creation date-time
	CreatedAt   AccessIPAddressListParamsFilterCreatedAtUnion `query:"created_at,omitzero" format:"date-time" json:"-"`
	ExtraFields map[string]any                                `json:"-"`
	paramObj
}

// URLQuery serializes [AccessIPAddressListParamsFilter]'s query parameters as
// `url.Values`.
func (r AccessIPAddressListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccessIPAddressListParamsFilterCreatedAtUnion struct {
	OfTime            param.Opt[time.Time]                                     `query:",omitzero,inline"`
	OfDateRangeFilter *AccessIPAddressListParamsFilterCreatedAtDateRangeFilter `query:",omitzero,inline"`
	paramUnion
}

func (u *AccessIPAddressListParamsFilterCreatedAtUnion) asAny() any {
	if !param.IsOmitted(u.OfTime) {
		return &u.OfTime.Value
	} else if !param.IsOmitted(u.OfDateRangeFilter) {
		return u.OfDateRangeFilter
	}
	return nil
}

// Date range filtering operations
type AccessIPAddressListParamsFilterCreatedAtDateRangeFilter struct {
	// Filter for creation date-time greater than
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Filter for creation date-time greater than or equal to
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Filter for creation date-time less than
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// Filter for creation date-time less than or equal to
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [AccessIPAddressListParamsFilterCreatedAtDateRangeFilter]'s
// query parameters as `url.Values`.
func (r AccessIPAddressListParamsFilterCreatedAtDateRangeFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
