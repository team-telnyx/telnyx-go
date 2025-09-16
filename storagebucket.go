// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// StorageBucketService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageBucketService] method instead.
type StorageBucketService struct {
	Options        []option.RequestOption
	SslCertificate StorageBucketSslCertificateService
	Usage          StorageBucketUsageService
}

// NewStorageBucketService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageBucketService(opts ...option.RequestOption) (r StorageBucketService) {
	r = StorageBucketService{}
	r.Options = opts
	r.SslCertificate = NewStorageBucketSslCertificateService(opts...)
	r.Usage = NewStorageBucketUsageService(opts...)
	return
}

// Returns a timed and authenticated URL to get an object. This is the equivalent
// to AWS S3’s “presigned” URL. Please note that Telnyx performs authentication
// differently from AWS S3 and you MUST NOT use the presign method of AWS s3api CLI
// or sdk to generate the presigned URL.
//
// Refer to: https://developers.telnyx.com/docs/cloud-storage/presigned-urls
func (r *StorageBucketService) NewPresignedURL(ctx context.Context, objectName string, params StorageBucketNewPresignedURLParams, opts ...option.RequestOption) (res *StorageBucketNewPresignedURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.BucketName == "" {
		err = errors.New("missing required bucketName parameter")
		return
	}
	if objectName == "" {
		err = errors.New("missing required objectName parameter")
		return
	}
	path := fmt.Sprintf("storage/buckets/%s/%s/presigned_url", params.BucketName, objectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type StorageBucketNewPresignedURLResponse struct {
	Content StorageBucketNewPresignedURLResponseContent `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketNewPresignedURLResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketNewPresignedURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketNewPresignedURLResponseContent struct {
	// The token for the object
	Token string `json:"token"`
	// The expiration time of the token
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The presigned URL for the object
	PresignedURL string `json:"presigned_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token        respjson.Field
		ExpiresAt    respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketNewPresignedURLResponseContent) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketNewPresignedURLResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketNewPresignedURLParams struct {
	BucketName string `path:"bucketName,required" json:"-"`
	// The time to live of the token in seconds
	Ttl param.Opt[int64] `json:"ttl,omitzero"`
	paramObj
}

func (r StorageBucketNewPresignedURLParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageBucketNewPresignedURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageBucketNewPresignedURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
