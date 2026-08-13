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

// Manage SQL databases and run SQL against them
//
// StorageSqldbService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageSqldbService] method instead.
type StorageSqldbService struct {
	Options []option.RequestOption
	// Manage SQL databases and run SQL against them
	Actions StorageSqldbActionService
}

// NewStorageSqldbService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageSqldbService(opts ...option.RequestOption) (r StorageSqldbService) {
	r = StorageSqldbService{}
	r.Options = opts
	r.Actions = NewStorageSqldbActionService(opts...)
	return
}

// Creates a new SQL database. Provisioning is asynchronous: the database is
// returned with status `pending` and becomes usable once it reaches
// `provision_ok`.
func (r *StorageSqldbService) New(ctx context.Context, body StorageSqldbNewParams, opts ...option.RequestOption) (res *SqlDatabaseResponseWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/sqldbs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a SQL database by its ID, including its provisioning status.
func (r *StorageSqldbService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SqlDatabaseResponseWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/sqldbs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the SQL databases for the authenticated user's organization. Results use
// page-based pagination (`page[number]`/`page[size]`) and can be filtered and
// sorted.
func (r *StorageSqldbService) List(ctx context.Context, query StorageSqldbListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[SqlDatabase], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/sqldbs"
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

// Lists the SQL databases for the authenticated user's organization. Results use
// page-based pagination (`page[number]`/`page[size]`) and can be filtered and
// sorted.
func (r *StorageSqldbService) ListAutoPaging(ctx context.Context, query StorageSqldbListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[SqlDatabase] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a SQL database and all of the data it holds. Deletion is asynchronous
// and returns `202` with an empty body — the record is not removed synchronously.
// Poll `GET /storage/sqldbs/{id}`, which returns `404` once the database has been
// purged; there is no durable `deleted` state. A database still bound by a
// function is refused with `409` unless `force=true`.
func (r *StorageSqldbService) Delete(ctx context.Context, id string, body StorageSqldbDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("storage/sqldbs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

type SqlDatabase struct {
	ID         string    `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	// Provisioning status. A database is usable once `status` is `provision_ok`. Once
	// deletion completes, the database no longer appears in the API.
	//
	// Any of "pending", "provision_ok", "provision_failed", "deleting",
	// "delete_failed".
	Status    SqlDatabaseStatus `json:"status"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
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
func (r SqlDatabase) RawJSON() string { return r.JSON.raw }
func (r *SqlDatabase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning status. A database is usable once `status` is `provision_ok`. Once
// deletion completes, the database no longer appears in the API.
type SqlDatabaseStatus string

const (
	SqlDatabaseStatusPending         SqlDatabaseStatus = "pending"
	SqlDatabaseStatusProvisionOk     SqlDatabaseStatus = "provision_ok"
	SqlDatabaseStatusProvisionFailed SqlDatabaseStatus = "provision_failed"
	SqlDatabaseStatusDeleting        SqlDatabaseStatus = "deleting"
	SqlDatabaseStatusDeleteFailed    SqlDatabaseStatus = "delete_failed"
)

type SqlDatabaseResponseWrapper struct {
	Data SqlDatabase `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SqlDatabaseResponseWrapper) RawJSON() string { return r.JSON.raw }
func (r *SqlDatabaseResponseWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSqldbNewParams struct {
	// Database name. Lowercase letters, numbers, and hyphens only; must start and end
	// with a letter or number.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r StorageSqldbNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageSqldbNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageSqldbNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSqldbListParams struct {
	// Filter by exact name match.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page. Values above 250 are treated as 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by provisioning status.
	//
	// Any of "pending", "provision_ok", "provision_failed", "deleting",
	// "delete_failed".
	FilterStatus StorageSqldbListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	// Sort field; prefix with `-` for descending order.
	//
	// Any of "name", "-name", "status", "-status", "created_at", "-created_at".
	Sort StorageSqldbListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageSqldbListParams]'s query parameters as `url.Values`.
func (r StorageSqldbListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by provisioning status.
type StorageSqldbListParamsFilterStatus string

const (
	StorageSqldbListParamsFilterStatusPending         StorageSqldbListParamsFilterStatus = "pending"
	StorageSqldbListParamsFilterStatusProvisionOk     StorageSqldbListParamsFilterStatus = "provision_ok"
	StorageSqldbListParamsFilterStatusProvisionFailed StorageSqldbListParamsFilterStatus = "provision_failed"
	StorageSqldbListParamsFilterStatusDeleting        StorageSqldbListParamsFilterStatus = "deleting"
	StorageSqldbListParamsFilterStatusDeleteFailed    StorageSqldbListParamsFilterStatus = "delete_failed"
)

// Sort field; prefix with `-` for descending order.
type StorageSqldbListParamsSort string

const (
	StorageSqldbListParamsSortName          StorageSqldbListParamsSort = "name"
	StorageSqldbListParamsSortNameDesc      StorageSqldbListParamsSort = "-name"
	StorageSqldbListParamsSortStatus        StorageSqldbListParamsSort = "status"
	StorageSqldbListParamsSortStatusDesc    StorageSqldbListParamsSort = "-status"
	StorageSqldbListParamsSortCreatedAt     StorageSqldbListParamsSort = "created_at"
	StorageSqldbListParamsSortCreatedAtDesc StorageSqldbListParamsSort = "-created_at"
)

type StorageSqldbDeleteParams struct {
	// Delete the database even when functions still bind it. Their bindings stop
	// resolving.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageSqldbDeleteParams]'s query parameters as
// `url.Values`.
func (r StorageSqldbDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
