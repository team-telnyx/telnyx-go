// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	shimjson "github.com/stainless-sdks/telnyx-go/internal/encoding/json"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// StorageMigrationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageMigrationService] method instead.
type StorageMigrationService struct {
	Options []option.RequestOption
	Actions StorageMigrationActionService
}

// NewStorageMigrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageMigrationService(opts ...option.RequestOption) (r StorageMigrationService) {
	r = StorageMigrationService{}
	r.Options = opts
	r.Actions = NewStorageMigrationActionService(opts...)
	return
}

// Initiate a migration of data from an external provider into Telnyx Cloud
// Storage. Currently, only S3 is supported.
func (r *StorageMigrationService) New(ctx context.Context, body StorageMigrationNewParams, opts ...option.RequestOption) (res *StorageMigrationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "storage/migrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Migration
func (r *StorageMigrationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *StorageMigrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/migrations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Migrations
func (r *StorageMigrationService) List(ctx context.Context, opts ...option.RequestOption) (res *StorageMigrationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "storage/migrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MigrationParamsResp struct {
	// ID of the Migration Source from which to migrate data.
	SourceID string `json:"source_id,required"`
	// Bucket name to migrate the data into. Will default to the same name as the
	// `source_bucket_name`.
	TargetBucketName string `json:"target_bucket_name,required"`
	// Telnyx Cloud Storage region to migrate the data to.
	TargetRegion string `json:"target_region,required"`
	// Unique identifier for the data migration.
	ID string `json:"id"`
	// Total amount of data that has been succesfully migrated.
	BytesMigrated int64 `json:"bytes_migrated"`
	// Total amount of data found in source bucket to migrate.
	BytesToMigrate int64 `json:"bytes_to_migrate"`
	// Time when data migration was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Estimated time the migration will complete.
	Eta time.Time `json:"eta" format:"date-time"`
	// Time when data migration was last copied from the source.
	LastCopy time.Time `json:"last_copy" format:"date-time"`
	// If true, will continue to poll the source bucket to ensure new data is
	// continually migrated over.
	Refresh bool `json:"refresh"`
	// Current speed of the migration.
	Speed int64 `json:"speed"`
	// Status of the migration.
	//
	// Any of "pending", "checking", "migrating", "complete", "error", "stopped".
	Status MigrationParamsStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceID         respjson.Field
		TargetBucketName respjson.Field
		TargetRegion     respjson.Field
		ID               respjson.Field
		BytesMigrated    respjson.Field
		BytesToMigrate   respjson.Field
		CreatedAt        respjson.Field
		Eta              respjson.Field
		LastCopy         respjson.Field
		Refresh          respjson.Field
		Speed            respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MigrationParamsResp) RawJSON() string { return r.JSON.raw }
func (r *MigrationParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MigrationParamsResp to a MigrationParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MigrationParams.Overrides()
func (r MigrationParamsResp) ToParam() MigrationParams {
	return param.Override[MigrationParams](json.RawMessage(r.RawJSON()))
}

// Status of the migration.
type MigrationParamsStatus string

const (
	MigrationParamsStatusPending   MigrationParamsStatus = "pending"
	MigrationParamsStatusChecking  MigrationParamsStatus = "checking"
	MigrationParamsStatusMigrating MigrationParamsStatus = "migrating"
	MigrationParamsStatusComplete  MigrationParamsStatus = "complete"
	MigrationParamsStatusError     MigrationParamsStatus = "error"
	MigrationParamsStatusStopped   MigrationParamsStatus = "stopped"
)

// The properties SourceID, TargetBucketName, TargetRegion are required.
type MigrationParams struct {
	// ID of the Migration Source from which to migrate data.
	SourceID string `json:"source_id,required"`
	// Bucket name to migrate the data into. Will default to the same name as the
	// `source_bucket_name`.
	TargetBucketName string `json:"target_bucket_name,required"`
	// Telnyx Cloud Storage region to migrate the data to.
	TargetRegion string `json:"target_region,required"`
	// Unique identifier for the data migration.
	ID param.Opt[string] `json:"id,omitzero"`
	// Total amount of data that has been succesfully migrated.
	BytesMigrated param.Opt[int64] `json:"bytes_migrated,omitzero"`
	// Total amount of data found in source bucket to migrate.
	BytesToMigrate param.Opt[int64] `json:"bytes_to_migrate,omitzero"`
	// Time when data migration was created
	CreatedAt param.Opt[time.Time] `json:"created_at,omitzero" format:"date-time"`
	// Estimated time the migration will complete.
	Eta param.Opt[time.Time] `json:"eta,omitzero" format:"date-time"`
	// Time when data migration was last copied from the source.
	LastCopy param.Opt[time.Time] `json:"last_copy,omitzero" format:"date-time"`
	// If true, will continue to poll the source bucket to ensure new data is
	// continually migrated over.
	Refresh param.Opt[bool] `json:"refresh,omitzero"`
	// Current speed of the migration.
	Speed param.Opt[int64] `json:"speed,omitzero"`
	// Status of the migration.
	//
	// Any of "pending", "checking", "migrating", "complete", "error", "stopped".
	Status MigrationParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r MigrationParams) MarshalJSON() (data []byte, err error) {
	type shadow MigrationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MigrationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationNewResponse struct {
	Data MigrationParamsResp `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationGetResponse struct {
	Data MigrationParamsResp `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationListResponse struct {
	Data []MigrationParamsResp `json:"data"`
	Meta PaginationMetaSimple  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationNewParams struct {
	MigrationParams MigrationParams
	paramObj
}

func (r StorageMigrationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MigrationParams)
}
func (r *StorageMigrationNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MigrationParams)
}
