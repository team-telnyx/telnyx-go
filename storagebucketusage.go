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
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Bucket Usage operations
//
// StorageBucketUsageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageBucketUsageService] method instead.
type StorageBucketUsageService struct {
	Options []option.RequestOption
}

// NewStorageBucketUsageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageBucketUsageService(opts ...option.RequestOption) (r StorageBucketUsageService) {
	r = StorageBucketUsageService{}
	r.Options = opts
	return
}

// Returns the detail on API usage on a bucket of a particular time period, group
// by method category.
func (r *StorageBucketUsageService) GetAPIUsage(ctx context.Context, bucketName string, query StorageBucketUsageGetAPIUsageParams, opts ...option.RequestOption) (res *StorageBucketUsageGetAPIUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketName == "" {
		err = errors.New("missing required bucketName parameter")
		return
	}
	path := fmt.Sprintf("storage/buckets/%s/usage/api", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns the amount of storage space and number of files a bucket takes up.
func (r *StorageBucketUsageService) GetBucketUsage(ctx context.Context, bucketName string, opts ...option.RequestOption) (res *StorageBucketUsageGetBucketUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketName == "" {
		err = errors.New("missing required bucketName parameter")
		return
	}
	path := fmt.Sprintf("storage/buckets/%s/usage/storage", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaginationMetaSimple struct {
	PageNumber   int64 `json:"page_number" api:"required"`
	TotalPages   int64 `json:"total_pages" api:"required"`
	PageSize     int64 `json:"page_size"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		TotalPages   respjson.Field
		PageSize     respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationMetaSimple) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetaSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetAPIUsageResponse struct {
	Data []StorageBucketUsageGetAPIUsageResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketUsageGetAPIUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketUsageGetAPIUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetAPIUsageResponseData struct {
	Categories []StorageBucketUsageGetAPIUsageResponseDataCategory `json:"categories"`
	// The time the usage was recorded
	Timestamp time.Time                                      `json:"timestamp" format:"date-time"`
	Total     StorageBucketUsageGetAPIUsageResponseDataTotal `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		Timestamp   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketUsageGetAPIUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketUsageGetAPIUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetAPIUsageResponseDataCategory struct {
	// The number of bytes received
	BytesReceived int64 `json:"bytes_received"`
	// The number of bytes sent
	BytesSent int64 `json:"bytes_sent"`
	// The category of the bucket operation
	//
	// Any of "list_bucket", "list_buckets", "get-bucket_location", "create_bucket",
	// "stat_bucket", "get_bucket_versioning", "set_bucket_versioning", "get_obj",
	// "put_obj", "delete_obj".
	Category string `json:"category"`
	// The number of operations
	Ops int64 `json:"ops"`
	// The number of successful operations
	SuccessfulOps int64 `json:"successful_ops"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BytesReceived respjson.Field
		BytesSent     respjson.Field
		Category      respjson.Field
		Ops           respjson.Field
		SuccessfulOps respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketUsageGetAPIUsageResponseDataCategory) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketUsageGetAPIUsageResponseDataCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetAPIUsageResponseDataTotal struct {
	// The number of bytes received
	BytesReceived int64 `json:"bytes_received"`
	// The number of bytes sent
	BytesSent int64 `json:"bytes_sent"`
	// The number of operations
	Ops int64 `json:"ops"`
	// The number of successful operations
	SuccessfulOps int64 `json:"successful_ops"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BytesReceived respjson.Field
		BytesSent     respjson.Field
		Ops           respjson.Field
		SuccessfulOps respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketUsageGetAPIUsageResponseDataTotal) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketUsageGetAPIUsageResponseDataTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetBucketUsageResponse struct {
	Data []StorageBucketUsageGetBucketUsageResponseData `json:"data"`
	Meta PaginationMetaSimple                           `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketUsageGetBucketUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketUsageGetBucketUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetBucketUsageResponseData struct {
	// The number of objects in the bucket
	NumObjects int64 `json:"num_objects"`
	// The size of the bucket in bytes
	Size int64 `json:"size"`
	// The size of the bucket in kilobytes
	SizeKB int64 `json:"size_kb"`
	// The time the snapshot was taken
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumObjects  respjson.Field
		Size        respjson.Field
		SizeKB      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketUsageGetBucketUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketUsageGetBucketUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketUsageGetAPIUsageParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[start_time], filter[end_time]
	Filter StorageBucketUsageGetAPIUsageParamsFilter `query:"filter,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [StorageBucketUsageGetAPIUsageParams]'s query parameters as
// `url.Values`.
func (r StorageBucketUsageGetAPIUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[start_time], filter[end_time]
//
// The properties EndTime, StartTime are required.
type StorageBucketUsageGetAPIUsageParamsFilter struct {
	// The end time of the period to filter the usage (ISO microsecond format)
	EndTime time.Time `query:"end_time" api:"required" format:"date-time" json:"-"`
	// The start time of the period to filter the usage (ISO microsecond format)
	StartTime time.Time `query:"start_time" api:"required" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [StorageBucketUsageGetAPIUsageParamsFilter]'s query
// parameters as `url.Values`.
func (r StorageBucketUsageGetAPIUsageParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
