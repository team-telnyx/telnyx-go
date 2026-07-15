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

// Manage CloudFS filesystems — JuiceFS-compatible filesystems backed by Telnyx
// Cloud Storage
//
// StorageCloudfService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageCloudfService] method instead.
type StorageCloudfService struct {
	Options []option.RequestOption
	// Manage CloudFS filesystems — JuiceFS-compatible filesystems backed by Telnyx
	// Cloud Storage
	Actions StorageCloudfActionService
}

// NewStorageCloudfService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageCloudfService(opts ...option.RequestOption) (r StorageCloudfService) {
	r = StorageCloudfService{}
	r.Options = opts
	r.Actions = NewStorageCloudfActionService(opts...)
	return
}

// Creates a CloudFS filesystem. Provisioning is synchronous — typically a few
// seconds, up to a few minutes — and the filesystem is returned with status
// `ready`, together with its S3 bucket and metadata connection details. This
// response is the only time the filesystem's `meta_token` — and the
// credential-bearing `meta_url` — are returned; store them securely. If the token
// is lost, issue a new one with the rotate-meta-token action. Names are unique
// within your organization: creating with an existing name returns a `422`.
// Requests are idempotent: retrying with the same `Idempotency-Key` within 24
// hours replays the original response instead of creating another filesystem.
func (r *StorageCloudfService) New(ctx context.Context, params StorageCloudfNewParams, opts ...option.RequestOption) (res *CloudfsFilesystemResponseWrapper, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "storage/cloudfs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a CloudFS filesystem by its ID. The returned `meta_url` omits the
// credential — the metadata token is only ever returned by create and
// rotate-meta-token. A filesystem whose last lifecycle action failed includes a
// customer-safe `error` message.
func (r *StorageCloudfService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CloudfsFilesystemDetailResponseWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/cloudfs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a CloudFS filesystem. Only `name` can be changed; other fields are
// immutable and unknown fields are rejected with a `400`. Renaming to a name that
// already exists in your organization returns a `422`.
func (r *StorageCloudfService) Update(ctx context.Context, id string, body StorageCloudfUpdateParams, opts ...option.RequestOption) (res *CloudfsFilesystemDetailResponseWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/cloudfs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists the CloudFS filesystems for the authenticated user's organization. Results
// use cursor-based pagination: fetch the next page by passing `meta.cursors.after`
// as `page[after]`, or follow the `meta.next` URL.
func (r *StorageCloudfService) List(ctx context.Context, query StorageCloudfListParams, opts ...option.RequestOption) (res *StorageCloudfListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/cloudfs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a CloudFS filesystem, removing its S3 bucket and its
// metadata database. Deletion is synchronous: the response returns the
// filesystem's final state with status `deleted`. There is no restore. A
// filesystem that is still `provisioning` returns a `409`. If the filesystem still
// contains data, the request may be rejected with a `409` — drain the bucket and
// retry.
func (r *StorageCloudfService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CloudfsFilesystemDetailResponseWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/cloudfs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CloudfsFilesystemDetailResponseWrapper struct {
	// A CloudFS filesystem as returned by get, update, and delete. `meta_url` omits
	// the credential and there is no `meta_token` field — the token is only returned
	// by create and rotate-meta-token.
	Data CloudfsFilesystemDetailResponseWrapperData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudfsFilesystemDetailResponseWrapper) RawJSON() string { return r.JSON.raw }
func (r *CloudfsFilesystemDetailResponseWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A CloudFS filesystem as returned by get, update, and delete. `meta_url` omits
// the credential and there is no `meta_token` field — the token is only returned
// by create and rotate-meta-token.
type CloudfsFilesystemDetailResponseWrapperData struct {
	ID        string    `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Explanation of the most recent failed lifecycle action. Present only when the
	// filesystem is in a `failed` state.
	Error string `json:"error"`
	// PostgreSQL connection URL for the filesystem's metadata database, without the
	// credential. Combine it with your stored metadata token, or issue a new token
	// with rotate-meta-token.
	MetaURL    string `json:"meta_url"`
	Name       string `json:"name"`
	RecordType string `json:"record_type"`
	Region     string `json:"region"`
	// Name of the bucket that stores this filesystem's data. Created during
	// provisioning.
	S3Bucket string `json:"s3_bucket"`
	// URL of the Telnyx Cloud Storage endpoint backing this filesystem.
	S3Endpoint string `json:"s3_endpoint"`
	// Lifecycle status of the filesystem. `ready` means it is fully provisioned and
	// usable. `needs_format` means the storage bucket and metadata database were
	// provisioned but the filesystem has not yet been formatted — run `juicefs format`
	// with the filesystem's `meta_url` before mounting. `failed` means the last
	// lifecycle action failed — see the filesystem's `error` message. `deleted`
	// appears only in the delete response: deleted filesystems are excluded from list
	// results and return a `404` on retrieval.
	//
	// Any of "provisioning", "ready", "needs_format", "deleting", "failed", "deleted".
	Status    CloudfsFilesystemStatus `json:"status"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Error       respjson.Field
		MetaURL     respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Region      respjson.Field
		S3Bucket    respjson.Field
		S3Endpoint  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudfsFilesystemDetailResponseWrapperData) RawJSON() string { return r.JSON.raw }
func (r *CloudfsFilesystemDetailResponseWrapperData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloudfsFilesystemResponseWrapper struct {
	// A CloudFS filesystem, including its metadata credential. This shape is returned
	// only by create and rotate-meta-token.
	Data CloudfsFilesystemResponseWrapperData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudfsFilesystemResponseWrapper) RawJSON() string { return r.JSON.raw }
func (r *CloudfsFilesystemResponseWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A CloudFS filesystem, including its metadata credential. This shape is returned
// only by create and rotate-meta-token.
type CloudfsFilesystemResponseWrapperData struct {
	ID        string    `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Metadata access token, in cleartext. Returned only by create and
	// rotate-meta-token and not retrievable afterwards — store it securely.
	MetaToken string `json:"meta_token"`
	// PostgreSQL connection URL for the filesystem's metadata database, including the
	// metadata token as the password. Pass it to `juicefs mount`: the storage
	// configuration is baked in at provisioning, so the metadata URL is all a client
	// needs to mount the filesystem.
	MetaURL    string `json:"meta_url"`
	Name       string `json:"name"`
	RecordType string `json:"record_type"`
	Region     string `json:"region"`
	// Name of the bucket that stores this filesystem's data. Created during
	// provisioning.
	S3Bucket string `json:"s3_bucket"`
	// URL of the Telnyx Cloud Storage endpoint backing this filesystem.
	S3Endpoint string `json:"s3_endpoint"`
	// Lifecycle status of the filesystem. `ready` means it is fully provisioned and
	// usable. `needs_format` means the storage bucket and metadata database were
	// provisioned but the filesystem has not yet been formatted — run `juicefs format`
	// with the filesystem's `meta_url` before mounting. `failed` means the last
	// lifecycle action failed — see the filesystem's `error` message. `deleted`
	// appears only in the delete response: deleted filesystems are excluded from list
	// results and return a `404` on retrieval.
	//
	// Any of "provisioning", "ready", "needs_format", "deleting", "failed", "deleted".
	Status    CloudfsFilesystemStatus `json:"status"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		MetaToken   respjson.Field
		MetaURL     respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Region      respjson.Field
		S3Bucket    respjson.Field
		S3Endpoint  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudfsFilesystemResponseWrapperData) RawJSON() string { return r.JSON.raw }
func (r *CloudfsFilesystemResponseWrapperData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the filesystem. `ready` means it is fully provisioned and
// usable. `needs_format` means the storage bucket and metadata database were
// provisioned but the filesystem has not yet been formatted — run `juicefs format`
// with the filesystem's `meta_url` before mounting. `failed` means the last
// lifecycle action failed — see the filesystem's `error` message. `deleted`
// appears only in the delete response: deleted filesystems are excluded from list
// results and return a `404` on retrieval.
type CloudfsFilesystemStatus string

const (
	CloudfsFilesystemStatusProvisioning CloudfsFilesystemStatus = "provisioning"
	CloudfsFilesystemStatusReady        CloudfsFilesystemStatus = "ready"
	CloudfsFilesystemStatusNeedsFormat  CloudfsFilesystemStatus = "needs_format"
	CloudfsFilesystemStatusDeleting     CloudfsFilesystemStatus = "deleting"
	CloudfsFilesystemStatusFailed       CloudfsFilesystemStatus = "failed"
	CloudfsFilesystemStatusDeleted      CloudfsFilesystemStatus = "deleted"
)

type StorageCloudfListResponse struct {
	Data []StorageCloudfListResponseData `json:"data"`
	Meta StorageCloudfListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCloudfListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageCloudfListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A CloudFS filesystem as returned in list results. Connection details
// (`meta_url`, `meta_token`) are omitted — retrieve the filesystem by ID for its
// redacted `meta_url`.
type StorageCloudfListResponseData struct {
	ID         string    `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	Region     string    `json:"region"`
	// Name of the bucket that stores this filesystem's data. Created during
	// provisioning.
	S3Bucket string `json:"s3_bucket"`
	// URL of the Telnyx Cloud Storage endpoint backing this filesystem.
	S3Endpoint string `json:"s3_endpoint"`
	// Lifecycle status of the filesystem. `ready` means it is fully provisioned and
	// usable. `needs_format` means the storage bucket and metadata database were
	// provisioned but the filesystem has not yet been formatted — run `juicefs format`
	// with the filesystem's `meta_url` before mounting. `failed` means the last
	// lifecycle action failed — see the filesystem's `error` message. `deleted`
	// appears only in the delete response: deleted filesystems are excluded from list
	// results and return a `404` on retrieval.
	//
	// Any of "provisioning", "ready", "needs_format", "deleting", "failed", "deleted".
	Status    CloudfsFilesystemStatus `json:"status"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Region      respjson.Field
		S3Bucket    respjson.Field
		S3Endpoint  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCloudfListResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageCloudfListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCloudfListResponseMeta struct {
	// Opaque cursors for the adjacent pages. Empty when there are no adjacent pages.
	Cursors StorageCloudfListResponseMetaCursors `json:"cursors"`
	// Relative URL (path and query) of the next page. Omitted when there are no
	// further results.
	Next string `json:"next"`
	// Relative URL (path and query) of the previous page. Omitted on the first page.
	Previous string `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCloudfListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *StorageCloudfListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opaque cursors for the adjacent pages. Empty when there are no adjacent pages.
type StorageCloudfListResponseMetaCursors struct {
	// Cursor for the next page; pass it as `page[after]`. Omitted on the last page.
	After string `json:"after"`
	// Cursor for the previous page; pass it as `page[before]`. Omitted on the first
	// page.
	Before string `json:"before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCloudfListResponseMetaCursors) RawJSON() string { return r.JSON.raw }
func (r *StorageCloudfListResponseMetaCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCloudfNewParams struct {
	// Filesystem name, unique within your organization. Names are trimmed and
	// lowercased; after normalization they may contain lowercase letters, numbers,
	// `.`, `_`, and `-` only.
	Name string `json:"name" api:"required"`
	// Region where the filesystem's storage and metadata are provisioned.
	//
	// Any of "us-central-1", "us-east-1", "us-west-1".
	Region         StorageCloudfNewParamsRegion `json:"region,omitzero" api:"required"`
	IdempotencyKey string                       `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r StorageCloudfNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageCloudfNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageCloudfNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the filesystem's storage and metadata are provisioned.
type StorageCloudfNewParamsRegion string

const (
	StorageCloudfNewParamsRegionUsCentral1 StorageCloudfNewParamsRegion = "us-central-1"
	StorageCloudfNewParamsRegionUsEast1    StorageCloudfNewParamsRegion = "us-east-1"
	StorageCloudfNewParamsRegionUsWest1    StorageCloudfNewParamsRegion = "us-west-1"
)

type StorageCloudfUpdateParams struct {
	// New filesystem name, unique within your organization. Names are trimmed and
	// lowercased; after normalization they may contain lowercase letters, numbers,
	// `.`, `_`, and `-` only.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r StorageCloudfUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageCloudfUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageCloudfUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCloudfListParams struct {
	// Return only the filesystem whose name matches exactly.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Return only filesystems in this region.
	FilterRegion param.Opt[string] `query:"filter[region],omitzero" json:"-"`
	// Opaque cursor from a previous response's `meta.cursors.after`; returns the page
	// after it. Mutually exclusive with `page[before]`.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Opaque cursor from a previous response's `meta.cursors.before`; returns the page
	// before it. Mutually exclusive with `page[after]`.
	PageBefore param.Opt[string] `query:"page[before],omitzero" json:"-"`
	// The number of filesystems to return per page. Values above 250 are treated
	// as 250.
	PageLimit param.Opt[int64] `query:"page[limit],omitzero" json:"-"`
	// Return only filesystems with this status. Unrecognized values are ignored.
	//
	// Any of "provisioning", "ready", "needs_format", "deleting", "failed".
	FilterStatus StorageCloudfListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	// Sort order for the results: a field name for ascending, or the field name
	// prefixed with `-` for descending.
	//
	// Any of "created_at", "-created_at", "updated_at", "-updated_at", "name",
	// "-name".
	Sort StorageCloudfListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageCloudfListParams]'s query parameters as
// `url.Values`.
func (r StorageCloudfListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Return only filesystems with this status. Unrecognized values are ignored.
type StorageCloudfListParamsFilterStatus string

const (
	StorageCloudfListParamsFilterStatusProvisioning StorageCloudfListParamsFilterStatus = "provisioning"
	StorageCloudfListParamsFilterStatusReady        StorageCloudfListParamsFilterStatus = "ready"
	StorageCloudfListParamsFilterStatusNeedsFormat  StorageCloudfListParamsFilterStatus = "needs_format"
	StorageCloudfListParamsFilterStatusDeleting     StorageCloudfListParamsFilterStatus = "deleting"
	StorageCloudfListParamsFilterStatusFailed       StorageCloudfListParamsFilterStatus = "failed"
)

// Sort order for the results: a field name for ascending, or the field name
// prefixed with `-` for descending.
type StorageCloudfListParamsSort string

const (
	StorageCloudfListParamsSortCreatedAt     StorageCloudfListParamsSort = "created_at"
	StorageCloudfListParamsSortCreatedAtDesc StorageCloudfListParamsSort = "-created_at"
	StorageCloudfListParamsSortUpdatedAt     StorageCloudfListParamsSort = "updated_at"
	StorageCloudfListParamsSortUpdatedAtDesc StorageCloudfListParamsSort = "-updated_at"
	StorageCloudfListParamsSortName          StorageCloudfListParamsSort = "name"
	StorageCloudfListParamsSortNameDesc      StorageCloudfListParamsSort = "-name"
)
