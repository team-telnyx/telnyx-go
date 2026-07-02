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

// Manage KV storage namespaces
//
// StorageKvService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageKvService] method instead.
type StorageKvService struct {
	Options []option.RequestOption
	// Read and write keys within a KV namespace
	Keys StorageKvKeyService
}

// NewStorageKvService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageKvService(opts ...option.RequestOption) (r StorageKvService) {
	r = StorageKvService{}
	r.Options = opts
	r.Keys = NewStorageKvKeyService(opts...)
	return
}

// Creates a new KV namespace. Provisioning is asynchronous: the namespace is
// returned with status `pending` and becomes usable once it reaches
// `provision_ok`.
func (r *StorageKvService) New(ctx context.Context, body StorageKvNewParams, opts ...option.RequestOption) (res *StorageKvNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/kvs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a KV namespace by its ID, including its provisioning status.
func (r *StorageKvService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *StorageKvGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/kvs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the KV namespaces for the authenticated user's organization. Results use
// page-based pagination (`page[number]`/`page[size]`).
func (r *StorageKvService) List(ctx context.Context, query StorageKvListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[StorageKvListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/kvs"
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

// Lists the KV namespaces for the authenticated user's organization. Results use
// page-based pagination (`page[number]`/`page[size]`).
func (r *StorageKvService) ListAutoPaging(ctx context.Context, query StorageKvListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[StorageKvListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a KV namespace and all of the keys it contains. Deletion is
// asynchronous: the namespace is returned with status `deleting`. Deleting a
// namespace whose deletion is already in progress returns a `409`.
func (r *StorageKvService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *StorageKvDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/kvs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type StorageKvNewResponse struct {
	Data StorageKvNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageKvNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNewResponseData struct {
	ID         string    `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	// Provisioning status. A namespace is usable once `status` is `provision_ok`. Once
	// deletion completes, the namespace no longer appears in the API.
	//
	// Any of "pending", "provision_ok", "provision_failed", "deleting",
	// "delete_failed".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageKvNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvGetResponse struct {
	Data StorageKvGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageKvGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvGetResponseData struct {
	ID         string    `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	// Provisioning status. A namespace is usable once `status` is `provision_ok`. Once
	// deletion completes, the namespace no longer appears in the API.
	//
	// Any of "pending", "provision_ok", "provision_failed", "deleting",
	// "delete_failed".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageKvGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvListResponse struct {
	ID         string    `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	// Provisioning status. A namespace is usable once `status` is `provision_ok`. Once
	// deletion completes, the namespace no longer appears in the API.
	//
	// Any of "pending", "provision_ok", "provision_failed", "deleting",
	// "delete_failed".
	Status    StorageKvListResponseStatus `json:"status"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageKvListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning status. A namespace is usable once `status` is `provision_ok`. Once
// deletion completes, the namespace no longer appears in the API.
type StorageKvListResponseStatus string

const (
	StorageKvListResponseStatusPending         StorageKvListResponseStatus = "pending"
	StorageKvListResponseStatusProvisionOk     StorageKvListResponseStatus = "provision_ok"
	StorageKvListResponseStatusProvisionFailed StorageKvListResponseStatus = "provision_failed"
	StorageKvListResponseStatusDeleting        StorageKvListResponseStatus = "deleting"
	StorageKvListResponseStatusDeleteFailed    StorageKvListResponseStatus = "delete_failed"
)

type StorageKvDeleteResponse struct {
	Data StorageKvDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageKvDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvDeleteResponseData struct {
	ID         string    `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	// Provisioning status. A namespace is usable once `status` is `provision_ok`. Once
	// deletion completes, the namespace no longer appears in the API.
	//
	// Any of "pending", "provision_ok", "provision_failed", "deleting",
	// "delete_failed".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageKvDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNewParams struct {
	// Namespace name. May contain lowercase letters, numbers, and hyphens only.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r StorageKvNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageKvNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageKvNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvListParams struct {
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page. Values above 250 are treated as 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageKvListParams]'s query parameters as `url.Values`.
func (r StorageKvListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
