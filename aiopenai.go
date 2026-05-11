// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
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

// Chat with a language model. This endpoint is consistent with the
// [OpenAI Responses API](https://platform.openai.com/docs/api-reference/responses)
// and may be used with the OpenAI JS or Python SDK. Response id parameter is not
// supported at the moment. Use 'conversation' parameter to leverage persistent
// conversations feature.
func (r *AIOpenAIService) NewResponse(ctx context.Context, body AIOpenAINewResponseParams, opts ...option.RequestOption) (res *AIOpenAINewResponseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists every model currently available to your account on Telnyx Inference,
// including SOTA open-source LLMs hosted on Telnyx GPUs (for example
// `moonshotai/Kimi-K2.6`, `zai-org/GLM-5.1-FP8`, and `MiniMaxAI/MiniMax-M2.7`),
// embedding models, and any fine-tuned models you have created.
//
// Each entry is a `ModelMetadata` object describing the model id, owner, task,
// context length, supported languages, billing tier, pricing per 1M tokens,
// deployment regions, and whether the model supports vision or fine-tuning. Use
// this endpoint to discover model ids you can pass to
// `POST /v2/ai/openai/chat/completions`.
//
// Model ids follow the `{organization}/{model_name}` convention from Hugging Face
// (for example `moonshotai/Kimi-K2.6`). This endpoint is OpenAI-compatible:
// clients pointed at `https://api.telnyx.com/v2/ai/openai` can call
// `client.models.list()` to retrieve the same payload.
func (r *AIOpenAIService) ListModels(ctx context.Context, opts ...option.RequestOption) (res *AIOpenAIListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AIOpenAINewResponseResponse map[string]any

type AIOpenAIListModelsResponse struct {
	Data   []ModelMetadata `json:"data" api:"required"`
	Object string          `json:"object"`
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

type AIOpenAINewResponseParams struct {
	Body map[string]any
	paramObj
}

func (r AIOpenAINewResponseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AIOpenAINewResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
