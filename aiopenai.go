// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIOpenAIService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOpenAIService] method instead.
type AIOpenAIService struct {
	Options []option.RequestOption
	// OpenAI-compatible embeddings endpoints for generating vector representations of
	// text
	Embeddings AIOpenAIEmbeddingService
	Chat       AIOpenAIChatService
}

// NewAIOpenAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOpenAIService(opts ...option.RequestOption) (r AIOpenAIService) {
	r = AIOpenAIService{}
	r.Options = opts
	r.Embeddings = NewAIOpenAIEmbeddingService(opts...)
	r.Chat = NewAIOpenAIChatService(opts...)
	return
}

// This endpoint returns a list of Open Source and OpenAI models that are available
// for use. <br /><br /> **Note**: Model `id`'s will be in the form
// `{source}/{model_name}`. For example `openai/gpt-4` or
// `mistralai/Mistral-7B-Instruct-v0.1` consistent with HuggingFace naming
// conventions.
func (r *AIOpenAIService) ListModels(ctx context.Context, opts ...option.RequestOption) (res *AIOpenAIListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AIOpenAIListModelsResponse struct {
	Data   []AIOpenAIListModelsResponseData `json:"data" api:"required"`
	Object string                           `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIListModelsResponseData struct {
	ID      string `json:"id" api:"required"`
	Created int64  `json:"created" api:"required"`
	OwnedBy string `json:"owned_by" api:"required"`
	Object  string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		OwnedBy     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIListModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIListModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
