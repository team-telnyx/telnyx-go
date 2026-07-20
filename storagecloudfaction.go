// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
)

// Manage CloudFS filesystems — JuiceFS-compatible filesystems backed by Telnyx
// Cloud Storage
//
// StorageCloudfActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageCloudfActionService] method instead.
type StorageCloudfActionService struct {
	Options []option.RequestOption
}

// NewStorageCloudfActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageCloudfActionService(opts ...option.RequestOption) (r StorageCloudfActionService) {
	r = StorageCloudfActionService{}
	r.Options = opts
	return
}

// Issues a new metadata access token for the filesystem and returns the full
// filesystem, including the new `meta_token` and credential-bearing `meta_url`.
// The previous token stops authenticating immediately; the metadata database and
// S3 bucket are unchanged. The request takes no body. Allowed while the filesystem
// is `ready` or `needs_format`; otherwise returns a `409`. Retrying with the same
// `Idempotency-Key` within 24 hours replays the original response — including the
// same token — instead of rotating again.
func (r *StorageCloudfActionService) RotateMetaToken(ctx context.Context, id string, body StorageCloudfActionRotateMetaTokenParams, opts ...option.RequestOption) (res *CloudfsFilesystemResponseWrapper, err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/cloudfs/%s/actions/rotate-meta-token", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type StorageCloudfActionRotateMetaTokenParams struct {
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}
