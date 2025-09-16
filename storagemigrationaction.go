// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// StorageMigrationActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageMigrationActionService] method instead.
type StorageMigrationActionService struct {
	Options []option.RequestOption
}

// NewStorageMigrationActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageMigrationActionService(opts ...option.RequestOption) (r StorageMigrationActionService) {
	r = StorageMigrationActionService{}
	r.Options = opts
	return
}

// Stop a Migration
func (r *StorageMigrationActionService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *StorageMigrationActionStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/migrations/%s/actions/stop", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type StorageMigrationActionStopResponse struct {
	Data MigrationParamsResp `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageMigrationActionStopResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageMigrationActionStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
