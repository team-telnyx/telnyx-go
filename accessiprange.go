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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// AccessIPRangeService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessIPRangeService] method instead.
type AccessIPRangeService struct {
	Options []option.RequestOption
}

// NewAccessIPRangeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessIPRangeService(opts ...option.RequestOption) (r AccessIPRangeService) {
	r = AccessIPRangeService{}
	r.Options = opts
	return
}

// Create new Access IP Range
func (r *AccessIPRangeService) New(ctx context.Context, body AccessIPRangeNewParams, opts ...option.RequestOption) (res *AccessIPRange, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "access_ip_ranges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all Access IP Ranges
func (r *AccessIPRangeService) List(ctx context.Context, query AccessIPRangeListParams, opts ...option.RequestOption) (res *AccessIPRangeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "access_ip_ranges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete access IP ranges
func (r *AccessIPRangeService) Delete(ctx context.Context, accessIPRangeID string, opts ...option.RequestOption) (res *AccessIPRange, err error) {
	opts = slices.Concat(r.Options, opts)
	if accessIPRangeID == "" {
		err = errors.New("missing required access_ip_range_id parameter")
		return
	}
	path := fmt.Sprintf("access_ip_ranges/%s", accessIPRangeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessIPRange struct {
	ID        string `json:"id,required"`
	CidrBlock string `json:"cidr_block,required"`
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
		CidrBlock   respjson.Field
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
func (r AccessIPRange) RawJSON() string { return r.JSON.raw }
func (r *AccessIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIPRangeListResponse struct {
	Data []AccessIPRange               `json:"data,required"`
	Meta AccessIPRangeListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessIPRangeListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessIPRangeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIPRangeListResponseMeta struct {
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
func (r AccessIPRangeListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AccessIPRangeListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIPRangeNewParams struct {
	CidrBlock   string            `json:"cidr_block,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AccessIPRangeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessIPRangeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessIPRangeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessIPRangeListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[cidr_block], filter[cidr_block][startswith],
	// filter[cidr_block][endswith], filter[cidr_block][contains], filter[created_at].
	// Supports complex bracket operations for dynamic filtering.
	Filter AccessIPRangeListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page AccessIPRangeListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccessIPRangeListParams]'s query parameters as
// `url.Values`.
func (r AccessIPRangeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[cidr_block], filter[cidr_block][startswith],
// filter[cidr_block][endswith], filter[cidr_block][contains], filter[created_at].
// Supports complex bracket operations for dynamic filtering.
type AccessIPRangeListParamsFilter struct {
	// Filter by exact CIDR block match
	CidrBlock AccessIPRangeListParamsFilterCidrBlockUnion `query:"cidr_block,omitzero" json:"-"`
	// Filter by exact creation date-time
	CreatedAt   AccessIPRangeListParamsFilterCreatedAtUnion `query:"created_at,omitzero" format:"date-time" json:"-"`
	ExtraFields map[string]any                              `json:"-"`
	paramObj
}

// URLQuery serializes [AccessIPRangeListParamsFilter]'s query parameters as
// `url.Values`.
func (r AccessIPRangeListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccessIPRangeListParamsFilterCidrBlockUnion struct {
	OfString                 param.Opt[string]                                             `query:",omitzero,inline"`
	OfCidrBlockPatternFilter *AccessIPRangeListParamsFilterCidrBlockCidrBlockPatternFilter `query:",omitzero,inline"`
	paramUnion
}

func (u *AccessIPRangeListParamsFilterCidrBlockUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCidrBlockPatternFilter) {
		return u.OfCidrBlockPatternFilter
	}
	return nil
}

// CIDR block pattern matching operations
type AccessIPRangeListParamsFilterCidrBlockCidrBlockPatternFilter struct {
	// Filter CIDR blocks containing the specified string
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter CIDR blocks ending with the specified string
	Endswith param.Opt[string] `query:"endswith,omitzero" json:"-"`
	// Filter CIDR blocks starting with the specified string
	Startswith param.Opt[string] `query:"startswith,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [AccessIPRangeListParamsFilterCidrBlockCidrBlockPatternFilter]'s query
// parameters as `url.Values`.
func (r AccessIPRangeListParamsFilterCidrBlockCidrBlockPatternFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccessIPRangeListParamsFilterCreatedAtUnion struct {
	OfTime            param.Opt[time.Time]                                   `query:",omitzero,inline"`
	OfDateRangeFilter *AccessIPRangeListParamsFilterCreatedAtDateRangeFilter `query:",omitzero,inline"`
	paramUnion
}

func (u *AccessIPRangeListParamsFilterCreatedAtUnion) asAny() any {
	if !param.IsOmitted(u.OfTime) {
		return &u.OfTime.Value
	} else if !param.IsOmitted(u.OfDateRangeFilter) {
		return u.OfDateRangeFilter
	}
	return nil
}

// Date range filtering operations
type AccessIPRangeListParamsFilterCreatedAtDateRangeFilter struct {
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

// URLQuery serializes [AccessIPRangeListParamsFilterCreatedAtDateRangeFilter]'s
// query parameters as `url.Values`.
func (r AccessIPRangeListParamsFilterCreatedAtDateRangeFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type AccessIPRangeListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	Size   param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccessIPRangeListParamsPage]'s query parameters as
// `url.Values`.
func (r AccessIPRangeListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
