// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIOpenAIEmbeddingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOpenAIEmbeddingService] method instead.
type AIOpenAIEmbeddingService struct {
	Options []option.RequestOption
}

// NewAIOpenAIEmbeddingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOpenAIEmbeddingService(opts ...option.RequestOption) (r AIOpenAIEmbeddingService) {
	r = AIOpenAIEmbeddingService{}
	r.Options = opts
	return
}

// Creates an embedding vector representing the input text. This endpoint is
// compatible with the
// [OpenAI Embeddings API](https://platform.openai.com/docs/api-reference/embeddings)
// and may be used with the OpenAI JS or Python SDK by setting the base URL to
// `https://api.telnyx.com/v2/ai/openai`.
func (r *AIOpenAIEmbeddingService) NewEmbeddings(ctx context.Context, body AIOpenAIEmbeddingNewEmbeddingsParams, opts ...option.RequestOption) (res *AIOpenAIEmbeddingNewEmbeddingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of available embedding models. This endpoint is compatible with
// the OpenAI Models API format.
func (r *AIOpenAIEmbeddingService) ListEmbeddingModels(ctx context.Context, opts ...option.RequestOption) (res *AIOpenAIEmbeddingListEmbeddingModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/embeddings/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIOpenAIEmbeddingNewEmbeddingsResponse struct {
	// List of embedding objects
	Data []AIOpenAIEmbeddingNewEmbeddingsResponseData `json:"data,required"`
	// The model used for embedding
	Model string `json:"model,required"`
	// The object type, always 'list'
	Object string                                      `json:"object,required"`
	Usage  AIOpenAIEmbeddingNewEmbeddingsResponseUsage `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIEmbeddingNewEmbeddingsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingNewEmbeddingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingNewEmbeddingsResponseData struct {
	// The embedding vector
	Embedding []float64 `json:"embedding,required"`
	// The index of the embedding in the list of embeddings
	Index int64 `json:"index,required"`
	// The object type, always 'embedding'
	Object string `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding   respjson.Field
		Index       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIEmbeddingNewEmbeddingsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingNewEmbeddingsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingNewEmbeddingsResponseUsage struct {
	// Number of tokens in the input
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total number of tokens used
	TotalTokens int64 `json:"total_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptTokens respjson.Field
		TotalTokens  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIEmbeddingNewEmbeddingsResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingNewEmbeddingsResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingListEmbeddingModelsResponse struct {
	// List of available embedding models
	Data []AIOpenAIEmbeddingListEmbeddingModelsResponseData `json:"data,required"`
	// The object type, always 'list'
	Object string `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIEmbeddingListEmbeddingModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingListEmbeddingModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingListEmbeddingModelsResponseData struct {
	// The model identifier
	ID string `json:"id,required"`
	// Unix timestamp of when the model was created
	Created int64 `json:"created,required"`
	// The object type, always 'model'
	Object string `json:"object,required"`
	// The organization that owns the model
	OwnedBy string `json:"owned_by,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIEmbeddingListEmbeddingModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingListEmbeddingModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingNewEmbeddingsParams struct {
	// Input text to embed. Can be a string or array of strings.
	Input AIOpenAIEmbeddingNewEmbeddingsParamsInputUnion `json:"input,omitzero,required"`
	// ID of the model to use. Use the List embedding models endpoint to see available
	// models.
	Model string `json:"model,required"`
	// The number of dimensions the resulting output embeddings should have. Only
	// supported in some models.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// A unique identifier representing your end-user for monitoring and abuse
	// detection.
	User param.Opt[string] `json:"user,omitzero"`
	// The format to return the embeddings in.
	//
	// Any of "float", "base64".
	EncodingFormat AIOpenAIEmbeddingNewEmbeddingsParamsEncodingFormat `json:"encoding_format,omitzero"`
	paramObj
}

func (r AIOpenAIEmbeddingNewEmbeddingsParams) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIEmbeddingNewEmbeddingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIEmbeddingNewEmbeddingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIOpenAIEmbeddingNewEmbeddingsParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AIOpenAIEmbeddingNewEmbeddingsParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AIOpenAIEmbeddingNewEmbeddingsParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIOpenAIEmbeddingNewEmbeddingsParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The format to return the embeddings in.
type AIOpenAIEmbeddingNewEmbeddingsParamsEncodingFormat string

const (
	AIOpenAIEmbeddingNewEmbeddingsParamsEncodingFormatFloat  AIOpenAIEmbeddingNewEmbeddingsParamsEncodingFormat = "float"
	AIOpenAIEmbeddingNewEmbeddingsParamsEncodingFormatBase64 AIOpenAIEmbeddingNewEmbeddingsParamsEncodingFormat = "base64"
)
