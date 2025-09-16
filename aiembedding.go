// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// AIEmbeddingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIEmbeddingService] method instead.
type AIEmbeddingService struct {
	Options []option.RequestOption
	Buckets AIEmbeddingBucketService
}

// NewAIEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIEmbeddingService(opts ...option.RequestOption) (r AIEmbeddingService) {
	r = AIEmbeddingService{}
	r.Options = opts
	r.Buckets = NewAIEmbeddingBucketService(opts...)
	return
}

// Perform embedding on a Telnyx Storage Bucket using the a embedding model. The
// current supported file types are:
//
//   - PDF
//   - HTML
//   - txt/unstructured text files
//   - json
//   - csv
//   - audio / video (mp3, mp4, mpeg, mpga, m4a, wav, or webm ) - Max of 100mb file
//     size.
//
// Any files not matching the above types will be attempted to be embedded as
// unstructured text.
//
// This process can be slow, so it runs in the background and the user can check
// the status of the task using the endpoint `/ai/embeddings/{task_id}`.
//
// **Important Note**: When you update documents in a Telnyx Storage bucket, their
// associated embeddings are automatically kept up to date. If you add or update a
// file, it is automatically embedded. If you delete a file, the embeddings are
// deleted for that particular file.
//
// You can also specify a custom `loader` param. Currently the only supported
// loader value is `intercom` which loads Intercom article jsons as specified by
// [the Intercom article API](https://developers.intercom.com/docs/references/rest-api/api.intercom.io/Articles/article/)
// This loader will split each article into paragraphs and save additional
// parameters relevant to Intercom docs, such as `article_url` and `heading`. These
// values will be returned by the `/v2/ai/embeddings/similarity-search` endpoint in
// the `loader_metadata` field.
func (r *AIEmbeddingService) New(ctx context.Context, body AIEmbeddingNewParams, opts ...option.RequestOption) (res *EmbeddingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check the status of a current embedding task. Will be one of the following:
//
//   - `queued` - Task is waiting to be picked up by a worker
//   - `processing` - The embedding task is running
//   - `success` - Task completed successfully and the bucket is embedded
//   - `failure` - Task failed and no files were embedded successfully
//   - `partial_success` - Some files were embedded successfully, but at least one
//     failed
func (r *AIEmbeddingService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *AIEmbeddingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("ai/embeddings/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve tasks for the user that are either `queued`, `processing`, `failed`,
// `success` or `partial_success` based on the query string. Defaults to `queued`
// and `processing`.
func (r *AIEmbeddingService) List(ctx context.Context, query AIEmbeddingListParams, opts ...option.RequestOption) (res *AIEmbeddingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Perform a similarity search on a Telnyx Storage Bucket, returning the most
// similar `num_docs` document chunks to the query.
//
// Currently the only available distance metric is cosine similarity which will
// return a `distance` between 0 and 1. The lower the distance, the more similar
// the returned document chunks are to the query. A `certainty` will also be
// returned, which is a value between 0 and 1 where the higher the certainty, the
// more similar the document. You can read more about Weaviate distance metrics
// here:
// [Weaviate Docs](https://weaviate.io/developers/weaviate/config-refs/distances)
//
// If a bucket was embedded using a custom loader, such as `intercom`, the
// additional metadata will be returned in the `loader_metadata` field.
func (r *AIEmbeddingService) SimilaritySearch(ctx context.Context, body AIEmbeddingSimilaritySearchParams, opts ...option.RequestOption) (res *AIEmbeddingSimilaritySearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/embeddings/similarity-search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Embed website content from a specified URL, including child pages up to 5 levels
// deep within the same domain. The process crawls and loads content from the main
// URL and its linked pages into a Telnyx Cloud Storage bucket. As soon as each
// webpage is added to the bucket, its content is immediately processed for
// embeddings, that can be used for
// [similarity search](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding-similarity-search)
// and [clustering](https://developers.telnyx.com/docs/inference/clusters).
func (r *AIEmbeddingService) URL(ctx context.Context, body AIEmbeddingURLParams, opts ...option.RequestOption) (res *EmbeddingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/embeddings/url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Status of an embeddings task.
type BackgroundTaskStatus string

const (
	BackgroundTaskStatusQueued         BackgroundTaskStatus = "queued"
	BackgroundTaskStatusProcessing     BackgroundTaskStatus = "processing"
	BackgroundTaskStatusSuccess        BackgroundTaskStatus = "success"
	BackgroundTaskStatusFailure        BackgroundTaskStatus = "failure"
	BackgroundTaskStatusPartialSuccess BackgroundTaskStatus = "partial_success"
)

type EmbeddingResponse struct {
	Data EmbeddingResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddingResponse) RawJSON() string { return r.JSON.raw }
func (r *EmbeddingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingResponseData struct {
	CreatedAt  string `json:"created_at"`
	FinishedAt string `json:"finished_at,nullable"`
	Status     string `json:"status"`
	TaskID     string `json:"task_id" format:"uuid"`
	TaskName   string `json:"task_name"`
	UserID     string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		FinishedAt  respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		TaskName    respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddingResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmbeddingResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingGetResponse struct {
	Data AIEmbeddingGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingGetResponseData struct {
	CreatedAt  string `json:"created_at"`
	FinishedAt string `json:"finished_at"`
	// Status of an embeddings task.
	//
	// Any of "queued", "processing", "success", "failure", "partial_success".
	Status   BackgroundTaskStatus `json:"status"`
	TaskID   string               `json:"task_id" format:"uuid"`
	TaskName string               `json:"task_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		FinishedAt  respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		TaskName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingListResponse struct {
	Data []AIEmbeddingListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingListResponseData struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Status of an embeddings task.
	//
	// Any of "queued", "processing", "success", "failure", "partial_success".
	Status     BackgroundTaskStatus `json:"status,required"`
	TaskID     string               `json:"task_id,required"`
	TaskName   string               `json:"task_name,required"`
	UserID     string               `json:"user_id,required"`
	Bucket     string               `json:"bucket"`
	FinishedAt time.Time            `json:"finished_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		TaskName    respjson.Field
		UserID      respjson.Field
		Bucket      respjson.Field
		FinishedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingSimilaritySearchResponse struct {
	Data []AIEmbeddingSimilaritySearchResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingSimilaritySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingSimilaritySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Example document response from embedding service { "document_chunk": "your
// status? This is Vanessa Bloome...", "distance": 0.18607724, "metadata": {
// "source": "https://us-central-1.telnyxstorage.com/scripts/bee_movie_script.txt",
// "checksum": "343054dd19bab39bbf6761a3d20f1daa", "embedding":
// "openai/text-embedding-ada-002", "filename": "bee_movie_script.txt",
// "certainty": 0.9069613814353943, "loader_metadata": {} } }
type AIEmbeddingSimilaritySearchResponseData struct {
	Distance      float64                                         `json:"distance,required"`
	DocumentChunk string                                          `json:"document_chunk,required"`
	Metadata      AIEmbeddingSimilaritySearchResponseDataMetadata `json:"metadata,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		DocumentChunk respjson.Field
		Metadata      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingSimilaritySearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingSimilaritySearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingSimilaritySearchResponseDataMetadata struct {
	Checksum       string         `json:"checksum,required"`
	Embedding      string         `json:"embedding,required"`
	Filename       string         `json:"filename,required"`
	Source         string         `json:"source,required"`
	Certainty      float64        `json:"certainty"`
	LoaderMetadata map[string]any `json:"loader_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checksum       respjson.Field
		Embedding      respjson.Field
		Filename       respjson.Field
		Source         respjson.Field
		Certainty      respjson.Field
		LoaderMetadata respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIEmbeddingSimilaritySearchResponseDataMetadata) RawJSON() string { return r.JSON.raw }
func (r *AIEmbeddingSimilaritySearchResponseDataMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingNewParams struct {
	BucketName               string           `json:"bucket_name,required"`
	DocumentChunkOverlapSize param.Opt[int64] `json:"document_chunk_overlap_size,omitzero"`
	DocumentChunkSize        param.Opt[int64] `json:"document_chunk_size,omitzero"`
	// Supported models to vectorize and embed documents.
	//
	// Any of "thenlper/gte-large", "intfloat/multilingual-e5-large".
	EmbeddingModel AIEmbeddingNewParamsEmbeddingModel `json:"embedding_model,omitzero"`
	// Supported types of custom document loaders for embeddings.
	//
	// Any of "default", "intercom".
	Loader AIEmbeddingNewParamsLoader `json:"loader,omitzero"`
	paramObj
}

func (r AIEmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIEmbeddingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIEmbeddingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Supported models to vectorize and embed documents.
type AIEmbeddingNewParamsEmbeddingModel string

const (
	AIEmbeddingNewParamsEmbeddingModelThenlperGteLarge            AIEmbeddingNewParamsEmbeddingModel = "thenlper/gte-large"
	AIEmbeddingNewParamsEmbeddingModelIntfloatMultilingualE5Large AIEmbeddingNewParamsEmbeddingModel = "intfloat/multilingual-e5-large"
)

// Supported types of custom document loaders for embeddings.
type AIEmbeddingNewParamsLoader string

const (
	AIEmbeddingNewParamsLoaderDefault  AIEmbeddingNewParamsLoader = "default"
	AIEmbeddingNewParamsLoaderIntercom AIEmbeddingNewParamsLoader = "intercom"
)

type AIEmbeddingListParams struct {
	// List of task statuses i.e. `status=queued&status=processing`
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIEmbeddingListParams]'s query parameters as `url.Values`.
func (r AIEmbeddingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIEmbeddingSimilaritySearchParams struct {
	BucketName string           `json:"bucket_name,required"`
	Query      string           `json:"query,required"`
	NumOfDocs  param.Opt[int64] `json:"num_of_docs,omitzero"`
	paramObj
}

func (r AIEmbeddingSimilaritySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow AIEmbeddingSimilaritySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIEmbeddingSimilaritySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIEmbeddingURLParams struct {
	// Name of the bucket to store the embeddings. This bucket must already exist.
	BucketName string `json:"bucket_name,required"`
	// The URL of the webpage to embed
	URL string `json:"url,required"`
	paramObj
}

func (r AIEmbeddingURLParams) MarshalJSON() (data []byte, err error) {
	type shadow AIEmbeddingURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIEmbeddingURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
