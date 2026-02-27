// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Migrate data from an external provider into Telnyx Cloud Storage
//
// StorageMigrationSourceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageMigrationSourceService] method instead.
type StorageMigrationSourceService struct {
	Options []option.RequestOption
}

// NewStorageMigrationSourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageMigrationSourceService(opts ...option.RequestOption) (r StorageMigrationSourceService) {
	r = StorageMigrationSourceService{}
	r.Options = opts
	return
}

// Create a source from which data can be migrated from.
func (r *StorageMigrationSourceService) New(ctx context.Context, body StorageMigrationSourceNewParams, opts ...option.RequestOption) (res *StorageMigrationSourceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/migration_sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Migration Source
func (r *StorageMigrationSourceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *StorageMigrationSourceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/migration_sources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Migration Sources
func (r *StorageMigrationSourceService) List(ctx context.Context, opts ...option.RequestOption) (res *StorageMigrationSourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/migration_sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Migration Source
func (r *StorageMigrationSourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *StorageMigrationSourceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/migration_sources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MigrationSourceParamsResp struct {
	// Bucket name to migrate the data from.
	BucketName string `json:"bucket_name" api:"required"`
	// Cloud provider from which to migrate data. Use 'telnyx' if you want to migrate
	// data from one Telnyx bucket to another.
	//
	// Any of "aws", "telnyx".
	Provider     MigrationSourceParamsProvider         `json:"provider" api:"required"`
	ProviderAuth MigrationSourceParamsProviderAuthResp `json:"provider_auth" api:"required"`
	// Unique identifier for the data migration source.
	ID string `json:"id"`
	// For intra-Telnyx buckets migration, specify the source bucket region in this
	// field.
	SourceRegion string `json:"source_region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketName   respjson.Field
		Provider     respjson.Field
		ProviderAuth respjson.Field
		ID           respjson.Field
		SourceRegion respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MigrationSourceParamsResp) RawJSON() string { return r.JSON.raw }
func (r *MigrationSourceParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MigrationSourceParamsResp to a MigrationSourceParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MigrationSourceParams.Overrides()
func (r MigrationSourceParamsResp) ToParam() MigrationSourceParams {
	return param.Override[MigrationSourceParams](json.RawMessage(r.RawJSON()))
}

// Cloud provider from which to migrate data. Use 'telnyx' if you want to migrate
// data from one Telnyx bucket to another.
type MigrationSourceParamsProvider string

const (
	MigrationSourceParamsProviderAws    MigrationSourceParamsProvider = "aws"
	MigrationSourceParamsProviderTelnyx MigrationSourceParamsProvider = "telnyx"
)

type MigrationSourceParamsProviderAuthResp struct {
	// AWS Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key here.
	AccessKey string `json:"access_key"`
	// AWS Secret Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key
	// here as well.
	SecretAccessKey string `json:"secret_access_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey       respjson.Field
		SecretAccessKey respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MigrationSourceParamsProviderAuthResp) RawJSON() string { return r.JSON.raw }
func (r *MigrationSourceParamsProviderAuthResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BucketName, Provider, ProviderAuth are required.
type MigrationSourceParams struct {
	// Bucket name to migrate the data from.
	BucketName string `json:"bucket_name" api:"required"`
	// Cloud provider from which to migrate data. Use 'telnyx' if you want to migrate
	// data from one Telnyx bucket to another.
	//
	// Any of "aws", "telnyx".
	Provider     MigrationSourceParamsProvider     `json:"provider,omitzero" api:"required"`
	ProviderAuth MigrationSourceParamsProviderAuth `json:"provider_auth,omitzero" api:"required"`
	// For intra-Telnyx buckets migration, specify the source bucket region in this
	// field.
	SourceRegion param.Opt[string] `json:"source_region,omitzero"`
	paramObj
}

func (r MigrationSourceParams) MarshalJSON() (data []byte, err error) {
	type shadow MigrationSourceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MigrationSourceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MigrationSourceParamsProviderAuth struct {
	// AWS Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key here.
	AccessKey param.Opt[string] `json:"access_key,omitzero"`
	// AWS Secret Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key
	// here as well.
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	paramObj
}

func (r MigrationSourceParamsProviderAuth) MarshalJSON() (data []byte, err error) {
	type shadow MigrationSourceParamsProviderAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MigrationSourceParamsProviderAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationSourceNewResponse struct {
	Data MigrationSourceParamsResp `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationSourceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationSourceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationSourceGetResponse struct {
	Data MigrationSourceParamsResp `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationSourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationSourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationSourceListResponse struct {
	Data []MigrationSourceParamsResp `json:"data"`
	Meta PaginationMetaSimple        `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationSourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationSourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationSourceDeleteResponse struct {
	Data MigrationSourceParamsResp `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationSourceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationSourceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageMigrationSourceNewParams struct {
	MigrationSourceParams MigrationSourceParams
	paramObj
}

func (r StorageMigrationSourceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MigrationSourceParams)
}
func (r *StorageMigrationSourceNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MigrationSourceParams)
}
