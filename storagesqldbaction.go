// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage SQL databases and run SQL against them
//
// StorageSqldbActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageSqldbActionService] method instead.
type StorageSqldbActionService struct {
	Options []option.RequestOption
}

// NewStorageSqldbActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageSqldbActionService(opts ...option.RequestOption) (r StorageSqldbActionService) {
	r = StorageSqldbActionService{}
	r.Options = opts
	return
}

// Runs SQL against the database and returns the resulting rows — empty for
// statements that return none, such as DDL. Bind positional `?` placeholders with
// `params` rather than interpolating values into the SQL string.
func (r *StorageSqldbActionService) Query(ctx context.Context, id string, body StorageSqldbActionQueryParams, opts ...option.RequestOption) (res *StorageSqldbActionQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/sqldbs/%s/actions/query", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type StorageSqldbActionQueryResponse struct {
	Data StorageSqldbActionQueryResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageSqldbActionQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageSqldbActionQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSqldbActionQueryResponseData struct {
	// Number of rows returned.
	Count int64 `json:"count"`
	// Wall-clock duration of the request, in milliseconds.
	Duration float64                                 `json:"duration"`
	Meta     StorageSqldbActionQueryResponseDataMeta `json:"meta"`
	// The result rows, each a map of column name to value. Empty for statements that
	// return no rows.
	Results []map[string]any `json:"results"`
	Success bool             `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Duration    respjson.Field
		Meta        respjson.Field
		Results     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageSqldbActionQueryResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageSqldbActionQueryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSqldbActionQueryResponseDataMeta struct {
	// Number of rows added, changed, or removed by the statement.
	Changes int64 `json:"changes"`
	// Wall-clock duration of the statement, in milliseconds.
	Duration float64 `json:"duration"`
	// Rowid of the last inserted row, when applicable.
	LastRowID   int64 `json:"last_row_id"`
	RowsRead    int64 `json:"rows_read"`
	RowsWritten int64 `json:"rows_written"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		Duration    respjson.Field
		LastRowID   respjson.Field
		RowsRead    respjson.Field
		RowsWritten respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageSqldbActionQueryResponseDataMeta) RawJSON() string { return r.JSON.raw }
func (r *StorageSqldbActionQueryResponseDataMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSqldbActionQueryParams struct {
	// The SQL to run. Use positional `?` placeholders and supply the values in
	// `params` rather than interpolating them into this string.
	Sql string `json:"sql" api:"required"`
	// Positional bind parameters, in placeholder order. Each value is a string, a
	// number, a boolean, or null; booleans are cast to `1`/`0`. The count must match
	// the number of `?` placeholders exactly — a mismatch is rejected with 422 rather
	// than binding null for the ones you left out. (Not enforced for multi-statement
	// scripts or named parameters, where the placeholder count is not the number
	// bound.)
	Params []StorageSqldbActionQueryParamsParamUnion `json:"params,omitzero"`
	paramObj
}

func (r StorageSqldbActionQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageSqldbActionQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageSqldbActionQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type StorageSqldbActionQueryParamsParamUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u StorageSqldbActionQueryParamsParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *StorageSqldbActionQueryParamsParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *StorageSqldbActionQueryParamsParamUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
