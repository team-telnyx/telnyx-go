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
func (r *AIOpenAIEmbeddingService) New(ctx context.Context, body AIOpenAIEmbeddingNewParams, opts ...option.RequestOption) (res *AIOpenAIEmbeddingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of available embedding models. This endpoint is compatible with
// the OpenAI Models API format.
func (r *AIOpenAIEmbeddingService) ListModels(ctx context.Context, opts ...option.RequestOption) (res *AIOpenAIEmbeddingListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/embeddings/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIOpenAIEmbeddingNewResponse struct {
	// List of embedding objects
	Data []AIOpenAIEmbeddingNewResponseData `json:"data,required"`
	// The model used for embedding
	Model string `json:"model,required"`
	// The object type, always 'list'
	Object string                            `json:"object,required"`
	Usage  AIOpenAIEmbeddingNewResponseUsage `json:"usage,required"`
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
func (r AIOpenAIEmbeddingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingNewResponseData struct {
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
func (r AIOpenAIEmbeddingNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingNewResponseUsage struct {
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
func (r AIOpenAIEmbeddingNewResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingNewResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingListModelsResponse struct {
	// List of available embedding models
	Data []AIOpenAIEmbeddingListModelsResponseData `json:"data,required"`
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
func (r AIOpenAIEmbeddingListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingListModelsResponseData struct {
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
func (r AIOpenAIEmbeddingListModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIEmbeddingListModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIEmbeddingNewParams struct {
	// Input text to embed. Can be a string or array of strings.
	Input AIOpenAIEmbeddingNewParamsInputUnion `json:"input,omitzero,required"`
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
	EncodingFormat AIOpenAIEmbeddingNewParamsEncodingFormat `json:"encoding_format,omitzero"`
	paramObj
}

func (r AIOpenAIEmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIEmbeddingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIEmbeddingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIOpenAIEmbeddingNewParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AIOpenAIEmbeddingNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AIOpenAIEmbeddingNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIOpenAIEmbeddingNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The format to return the embeddings in.
type AIOpenAIEmbeddingNewParamsEncodingFormat string

const (
	AIOpenAIEmbeddingNewParamsEncodingFormatFloat  AIOpenAIEmbeddingNewParamsEncodingFormat = "float"
	AIOpenAIEmbeddingNewParamsEncodingFormatBase64 AIOpenAIEmbeddingNewParamsEncodingFormat = "base64"
)
