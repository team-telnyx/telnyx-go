// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// AIEmbeddingBucketService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIEmbeddingBucketService] method instead.
type AIEmbeddingBucketService struct {
	Options []option.RequestOption
}

// NewAIEmbeddingBucketService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIEmbeddingBucketService(opts ...option.RequestOption) (r AIEmbeddingBucketService) {
	r = AIEmbeddingBucketService{}
	r.Options = opts
	return
}

// Get all embedded files for a given user bucket, including their processing
// status.
func (r *AIEmbeddingBucketService) Get(ctx context.Context, bucketName string, opts ...option.RequestOption) (res *AIEmbeddingBucketGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("ai/embeddings/buckets/%s", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all embedding buckets for a user.
func (r *AIEmbeddingBucketService) List(ctx context.Context, opts ...option.RequestOption) (res *AIEmbeddingBucketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/embeddings/buckets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an entire bucket's embeddings and disables the bucket for AI-use,
// returning it to normal storage pricing.
func (r *AIEmbeddingBucketService) Delete(ctx context.Context, bucketName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("ai/embeddings/buckets/%s", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AIEmbeddingBucketGetResponse struct {
	Data []AIEmbeddingBucketGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingBucketGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingBucketGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingBucketGetResponseData struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	Filename       string    `json:"filename,required"`
	Status         string    `json:"status,required"`
	ErrorReason    string    `json:"error_reason"`
	LastEmbeddedAt time.Time `json:"last_embedded_at" format:"date-time"`
	UpdatedAt      time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Filename       respjson.Field
		Status         respjson.Field
		ErrorReason    respjson.Field
		LastEmbeddedAt respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingBucketGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingBucketGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingBucketListResponse struct {
	Data AIEmbeddingBucketListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingBucketListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingBucketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingBucketListResponseData struct {
	Buckets []string `json:"buckets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingBucketListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingBucketListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
