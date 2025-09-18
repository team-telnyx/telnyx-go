// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// StorageService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageService] method instead.
type StorageService struct {
	Options          []option.RequestOption
	Buckets          StorageBucketService
	MigrationSources StorageMigrationSourceService
	Migrations       StorageMigrationService
}

// NewStorageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStorageService(opts ...option.RequestOption) (r StorageService) {
	r = StorageService{}
	r.Options = opts
	r.Buckets = NewStorageBucketService(opts...)
	r.MigrationSources = NewStorageMigrationSourceService(opts...)
	r.Migrations = NewStorageMigrationService(opts...)
	return
}

// List Migration Source coverage
func (r *StorageService) ListMigrationSourceCoverage(ctx context.Context, opts ...option.RequestOption) (res *StorageListMigrationSourceCoverageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "storage/migration_source_coverage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StorageListMigrationSourceCoverageResponse struct {
	Data []StorageListMigrationSourceCoverageResponseData `json:"data"`
	Meta PaginationMetaSimple                             `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageListMigrationSourceCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageListMigrationSourceCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageListMigrationSourceCoverageResponseData struct {
	// Cloud provider from which to migrate data.
	//
	// Any of "aws".
	Provider string `json:"provider"`
	// Provider region from which to migrate data.
	SourceRegion string `json:"source_region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Provider     respjson.Field
		SourceRegion respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageListMigrationSourceCoverageResponseData) RawJSON() string { return r.JSON.raw }
func (r *StorageListMigrationSourceCoverageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
